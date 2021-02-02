package main

import (
	"flag"
	"fmt"
	"github.com/asmyasnikov/go-mavlink/common"
	"log"
	"os"
)

var (
	mavgenVersion   = "devel"
	schemeFile      = flag.String("f", "", "mavlink xml-definition file input")
	version         = flag.String("v", mavgenVersion, "custom version of mavlink dialect")
	commonPackage   = flag.String("c", "", "common mavlink package path used for import from go code. For example \"github.com/asmyasnikov/go-mavlink/generated/mavlink1\"")
	generateCommons = flag.Bool("p", false, "generate common mavlink package")
)

type templateData struct {
	Version        string
	MavlinkVersion int
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("mavgen: ")
	flag.Parse()

	if len(*schemeFile) == 0 {
		fmt.Println("Usage: " + os.Args[0] + " (-1|-2) [-c \"github.com/my/project/mavlink/path\"] [-p] [-f ./ardupilotmega.xml] [-v \"ardupilotmega-v1.2.3\"]")
		flag.PrintDefaults()
		return
	}

	if (*common.Mavlink1 && *common.Mavlink2) || (!*common.Mavlink1 && !*common.Mavlink2) {
		fmt.Println("Select only one version of mavlink: -1 flag or -2 flag")
		return
	}

	if err := generateDialect(nil, *commonPackage, *schemeFile, common.MavlinkVersion()); err != nil {
		log.Fatal(err)
	}

	if *generateCommons {
		data := templateData{
			Version:        *version,
			MavlinkVersion: common.MavlinkVersion(),
		}
		if err := generateCommonPackage(data); err != nil {
			log.Fatal(err)
		}
	}
}
