package common

import (
	"flag"
	"github.com/asmyasnikov/go-mavlink/mavlink"
)

var (
	// Mavlink1 common flag
	Mavlink1 = flag.Bool("1", false, "mavlink v1")
	// Mavlink2 common flag
	Mavlink2 = flag.Bool("2", false, "mavlink v2")
)

// MavlinkVersion process user flags and return mavlink version
func MavlinkVersion() mavlink.MAVLINK_VERSION {
	return mavlinkVersion(*Mavlink1, *Mavlink2)
}

func mavlinkVersion(mavlink1 bool, mavlink2 bool) mavlink.MAVLINK_VERSION {
	version := mavlink.MAVLINK_V1 | mavlink.MAVLINK_V2
	if !mavlink1 && mavlink2 {
		version &= ^mavlink.MAVLINK_V1
	}
	if !mavlink2 && mavlink1 {
		version &= ^mavlink.MAVLINK_V2
	}
	return version
}
