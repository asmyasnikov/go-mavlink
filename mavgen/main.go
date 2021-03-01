package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	mavgenVersion   = "devel"
	schemeFile      = flag.String("f", "", "mavlink xml-definition file input")
	commonPackage   = flag.String("c", "", "common mavlink package path used for import from go code. For example \"github.com/asmyasnikov/go-mavlink/generated/mavlink1\"")
	generateCommons = flag.Bool("p", false, "generate common mavlink package")
	json            = flag.Bool("json", false, "generate json marshalers/unmarshalers")
)

type templateData struct {
	MavlinkVersion   int
	CommonPackageURL string
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("mavgen: ")
	flag.Parse()

	if len(*schemeFile) == 0 {
		fmt.Println("Usage: " + os.Args[0] + " [-c \"github.com/my/project/mavlink/path\"] [-p] [-f ./ardupilotmega.xml] [-v \"ardupilotmega-v1.2.3\"]")
		flag.PrintDefaults()
		return
	}

	if len(*commonPackage) == 0 {
		*commonPackage = "github.com/asmyasnikov/go-mavlink/mavlink"
	}

	if err := generateDialect(nil, *commonPackage, *schemeFile, *json); err != nil {
		log.Fatal(err)
	}

	if *generateCommons {
		if err := generateCommonPackage(*commonPackage); err != nil {
			log.Fatal(err)
		}
	}
}
