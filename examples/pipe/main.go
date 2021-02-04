package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"github.com/asmyasnikov/go-mavlink/mavlink"
	_ "github.com/asmyasnikov/go-mavlink/mavlink/dialects/ardupilotmega"
	"io"
	"log"
	"os"
	"regexp"
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
	dec := mavlink.NewDecoder(reader)
	log.Println("listening packets")
	for {
		if p, err := dec.Decode(); err != nil {
			log.Fatal("Error on decode:" + err.Error())
		} else {
			log.Println("<-", p.String())
		}
	}
}
