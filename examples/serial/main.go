package main

import (
	"flag"
	"github.com/asmyasnikov/go-mavlink/mavlink"
	"github.com/asmyasnikov/go-mavlink/mavlink/dialects/ardupilotmega"
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/tarm/serial"
	"io"
	"log"
	"sync"
	"time"
)

//////////////////////////////////////
//
// mavlink serial port device reader and parser example
//
// listen serial port device and prints received msgs, more info:
//
// run via `go run main.go -d /dev/ttyACM0 -b 57600`
//
//////////////////////////////////////

var (
	baudrate   = flag.Int("b", 57600, "baudrate of serial port connection")
	device     = flag.String("d", "/dev/ttyUSB0", "path of serial port device")
	ro         = flag.Bool("ro", false, "read-only mode")
	retryCount = flag.Int("retry-count", 1, "retry count for sending packets")
	sysID      = flag.Int("sysID", 1, "System ID (copter)")
)

func main() {
	flag.Parse()
	port, err := serial.OpenPort(&serial.Config{
		Name:        *device,
		Baud:        *baudrate,
		ReadTimeout: time.Second,
		Size:        8,
		Parity:      serial.ParityNone,
		StopBits:    1,
	})
	if err != nil {
		log.Fatalf("Error on opening device %s: %s\n", *device, err.Error())
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go listenAndServe(&wg, port, *ro)
	if !*ro {
		wg.Add(3)
		go sendLoop(&wg, port)
		go heartbeatLoop(&wg)
		go handshake(&wg)
	}
	wg.Wait()
}

func listenAndServe(wg *sync.WaitGroup, device io.ReadWriteCloser, ro bool) {
	defer wg.Done()
	dec := mavlink.NewDecoder(device)
	log.Println("listening packets")
	for {
		p, err := dec.Decode()
		if err != nil {
			log.Fatal("decode error: ", err.Error())
		} else if p != nil {
			log.Println("<-", p.String())
		}
		if p.MsgID() == ardupilotmega.MSG_ID_TIMESYNC {
			ts := ardupilotmega.Timesync{}
			if err := ts.Unmarshal(p.Payload()); err != nil {
				log.Fatal(err)
			} else if !ro {
				sendChan <- makeTimeSync(ts.Ts1)
			}
		}
	}
}

func makeHeartbeat() message.Message {
	return &ardupilotmega.Heartbeat{
		CustomMode:     0,
		Type:           ardupilotmega.MAV_TYPE_GCS,
		Autopilot:      ardupilotmega.MAV_AUTOPILOT_INVALID,
		BaseMode:       0,
		SystemStatus:   ardupilotmega.MAV_STATE_UNINIT,
		MavlinkVersion: 3,
	}
}

func makeRequestDataStream(msgID ardupilotmega.MAV_DATA_STREAM, rate uint16) message.Message {
	return &ardupilotmega.RequestDataStream{
		ReqMessageRate:  rate,
		TargetSystem:    1,
		TargetComponent: 1,
		ReqStreamID:     uint8(msgID),
		StartStop:       1,
	}
}

func makeStatustext(text string) message.Message {
	return &ardupilotmega.Statustext{
		Severity: ardupilotmega.MAV_SEVERITY_INFO,
		Text:     text,
	}
}

func makeCommandLong(cmd ardupilotmega.MAV_CMD, param1 uint32) message.Message {
	return &ardupilotmega.CommandLong{
		Param1:          float32(param1),
		Param2:          0,
		Param3:          0,
		Param4:          0,
		Param5:          0,
		Param6:          0,
		Param7:          0,
		Command:         cmd,
		TargetSystem:    1,
		TargetComponent: 1,
		Confirmation:    0,
	}
}

func makeParamRequestList() message.Message {
	return &ardupilotmega.ParamRequestList{
		TargetSystem:    1,
		TargetComponent: 1,
	}
}

func makeTimeSync(ts int64) message.Message {
	return &ardupilotmega.Timesync{
		Tc1: time.Now().UnixNano(),
		Ts1: ts,
	}
}

func makeFileTransferProtocol(payload []byte) message.Message {
	return &ardupilotmega.FileTransferProtocol{
		TargetNetwork:   0,
		TargetSystem:    1,
		TargetComponent: 1,
		Payload:         payload,
	}
}

var sendChan = make(chan message.Message)

func heartbeatLoop(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(time.Second)
		sendChan <- makeHeartbeat()
	}
}

func sendLoop(wg *sync.WaitGroup, writer io.Writer) {
	defer wg.Done()
	enc := mavlink.NewEncoder(writer)
	for m := range sendChan {
		for i := 0; i < *retryCount; i++ {
			p, err := mavlink.NewPacket(2, uint8(*sysID), uint8(ardupilotmega.MAV_COMP_ID_MISSIONPLANNER), m)
			if err != nil {
				log.Fatalf("Error on create packet: %+v", err)
			}
			if err := enc.Encode(p); err != nil {
				log.Fatalf("Error on send loop: %+v", err)
			} else {
				log.Println("->", p)
			}
		}
	}
}

// https://mavlink.io/en/guide/mavlink_version.html#version_handshaking
func handshake(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	sendChan <- makeHeartbeat()
	sendChan <- makeCommandLong(ardupilotmega.MAV_CMD_REQUEST_MESSAGE, uint32(ardupilotmega.MSG_ID_AUTOPILOT_VERSION_REQUEST))
	sendChan <- makeCommandLong(ardupilotmega.MAV_CMD_REQUEST_PROTOCOL_VERSION, 1)
	sendChan <- makeCommandLong(ardupilotmega.MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES, 1)
	sendChan <- makeRequestDataStream(ardupilotmega.MAV_DATA_STREAM_EXTENDED_STATUS, 2)
	sendChan <- makeRequestDataStream(ardupilotmega.MAV_DATA_STREAM_POSITION, 2)
	sendChan <- makeRequestDataStream(ardupilotmega.MAV_DATA_STREAM_EXTRA1, 4)
	sendChan <- makeRequestDataStream(ardupilotmega.MAV_DATA_STREAM_EXTRA2, 4)
	sendChan <- makeRequestDataStream(ardupilotmega.MAV_DATA_STREAM_EXTRA3, 4)
	sendChan <- makeRequestDataStream(ardupilotmega.MAV_DATA_STREAM_RAW_SENSORS, 2)
	sendChan <- makeRequestDataStream(ardupilotmega.MAV_DATA_STREAM_RC_CHANNELS, 2)
	sendChan <- makeHeartbeat()
	sendChan <- makeStatustext("Mission Planner 1.3.74")
	sendChan <- makeCommandLong(ardupilotmega.MAV_CMD_DO_SEND_BANNER, 0)
	sendChan <- makeFileTransferProtocol([]byte{0, 0, 0, 4, 16, 0, 0, 0, 0, 0, 0, 0, 64, 80, 65, 82, 65, 77, 47, 112, 97, 114, 97, 109, 46, 112, 99, 107, 0})
}
