package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"github.com/asmyasnikov/go-mavlink/common/decoder"
	_ "github.com/asmyasnikov/go-mavlink/generated/mavlink1/ardupilotmega"
	_ "github.com/asmyasnikov/go-mavlink/generated/mavlink2/ardupilotmega"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
)

//////////////////////////////////////
//
// mavlink pipe parser example
//
// listen to input pipe and prints received msgs, more info:
//
// run via `cat /dev/ttyUSB0 | go run main.go -version 1`
//
//////////////////////////////////////

var (
	hexInput = flag.Bool("h", false, "process input stream as hex dump")
)

type hexByteReader struct {
	r *bufio.Reader
}

func (r *hexByteReader) Read(p []byte) (n int, err error) {
	b, err := r.r.ReadBytes('\n')
	if err != nil {
		log.Fatal("Error on reader.ReadBytes(): " + err.Error())
		return 0, err
	}
	reg, err := regexp.Compile("[^a-fA-F0-9]+")
	if err != nil {
		log.Fatal("Error on regexp.Compile(): " + err.Error())
		return 0, err
	}
	b, err = hex.DecodeString(reg.ReplaceAllString(string(b), ""))
	if err != nil {
		log.Printf("%0X", b)
		log.Fatal("Error on hex.DecodeString(): " + err.Error())
		return 0, err
	}
	copy(p, b)
	return len(p), err
}

func main() {
	flag.Parse()
	listenAndServe()
}

func listenAndServe() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeNamedPipe == 0 {
		log.Println("The command is intended to work with pipes.")
		return
	}
	var reader io.Reader
	if *hexInput {
		reader = &hexByteReader{
			r: bufio.NewReader(os.Stdin),
		}
	} else {
		reader = bufio.NewReader(os.Stdin)
	}
	wg := sync.WaitGroup{}
	decs := decoder.Decoders(reader)
	for i := range decs {
		wg.Add(1)
		dec := decs[i]
		go func() {
			defer wg.Done()
			log.Println("listening packets from decoder " + dec.Name())
			for {
				if p, err := dec.Decode(); err != nil {
					log.Fatal("Error on " + dec.Name() + " decode:" + err.Error())
				} else {
					log.Println("<-", p.String())
				}
			}
		}()
	}
	wg.Wait()
}
