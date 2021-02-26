/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package icarous

import (
	"fmt"
	"strconv"
)

// ICAROUS_TRACK_BAND_TYPES type
type ICAROUS_TRACK_BAND_TYPES int

func (e ICAROUS_TRACK_BAND_TYPES) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ICAROUS_TRACK_BAND_TYPE_NONE enum
	ICAROUS_TRACK_BAND_TYPE_NONE ICAROUS_TRACK_BAND_TYPES = 0
	// ICAROUS_TRACK_BAND_TYPE_NEAR enum
	ICAROUS_TRACK_BAND_TYPE_NEAR ICAROUS_TRACK_BAND_TYPES = 1
	// ICAROUS_TRACK_BAND_TYPE_RECOVERY enum
	ICAROUS_TRACK_BAND_TYPE_RECOVERY ICAROUS_TRACK_BAND_TYPES = 2
)

func (e ICAROUS_TRACK_BAND_TYPES) String() string {
	if name, ok := map[ICAROUS_TRACK_BAND_TYPES]string{
		ICAROUS_TRACK_BAND_TYPE_NONE:     "ICAROUS_TRACK_BAND_TYPE_NONE",
		ICAROUS_TRACK_BAND_TYPE_NEAR:     "ICAROUS_TRACK_BAND_TYPE_NEAR",
		ICAROUS_TRACK_BAND_TYPE_RECOVERY: "ICAROUS_TRACK_BAND_TYPE_RECOVERY",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("ICAROUS_TRACK_BAND_TYPES_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects ICAROUS_TRACK_BAND_TYPES enums
func (e ICAROUS_TRACK_BAND_TYPES) Bitmask() string {
	bitmap := ""
	for _, entry := range []ICAROUS_TRACK_BAND_TYPES{
		ICAROUS_TRACK_BAND_TYPE_NONE,
		ICAROUS_TRACK_BAND_TYPE_NEAR,
		ICAROUS_TRACK_BAND_TYPE_RECOVERY,
	} {
		if e&entry > 0 {
			if len(bitmap) > 0 {
				bitmap += " | "
			}
			bitmap += entry.String()
		}
	}
	return bitmap
}

// ICAROUS_FMS_STATE type
type ICAROUS_FMS_STATE int

func (e ICAROUS_FMS_STATE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ICAROUS_FMS_STATE_IDLE enum
	ICAROUS_FMS_STATE_IDLE ICAROUS_FMS_STATE = 0
	// ICAROUS_FMS_STATE_TAKEOFF enum
	ICAROUS_FMS_STATE_TAKEOFF ICAROUS_FMS_STATE = 1
	// ICAROUS_FMS_STATE_CLIMB enum
	ICAROUS_FMS_STATE_CLIMB ICAROUS_FMS_STATE = 2
	// ICAROUS_FMS_STATE_CRUISE enum
	ICAROUS_FMS_STATE_CRUISE ICAROUS_FMS_STATE = 3
	// ICAROUS_FMS_STATE_APPROACH enum
	ICAROUS_FMS_STATE_APPROACH ICAROUS_FMS_STATE = 4
	// ICAROUS_FMS_STATE_LAND enum
	ICAROUS_FMS_STATE_LAND ICAROUS_FMS_STATE = 5
)

func (e ICAROUS_FMS_STATE) String() string {
	if name, ok := map[ICAROUS_FMS_STATE]string{
		ICAROUS_FMS_STATE_IDLE:     "ICAROUS_FMS_STATE_IDLE",
		ICAROUS_FMS_STATE_TAKEOFF:  "ICAROUS_FMS_STATE_TAKEOFF",
		ICAROUS_FMS_STATE_CLIMB:    "ICAROUS_FMS_STATE_CLIMB",
		ICAROUS_FMS_STATE_CRUISE:   "ICAROUS_FMS_STATE_CRUISE",
		ICAROUS_FMS_STATE_APPROACH: "ICAROUS_FMS_STATE_APPROACH",
		ICAROUS_FMS_STATE_LAND:     "ICAROUS_FMS_STATE_LAND",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("ICAROUS_FMS_STATE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects ICAROUS_FMS_STATE enums
func (e ICAROUS_FMS_STATE) Bitmask() string {
	bitmap := ""
	for _, entry := range []ICAROUS_FMS_STATE{
		ICAROUS_FMS_STATE_IDLE,
		ICAROUS_FMS_STATE_TAKEOFF,
		ICAROUS_FMS_STATE_CLIMB,
		ICAROUS_FMS_STATE_CRUISE,
		ICAROUS_FMS_STATE_APPROACH,
		ICAROUS_FMS_STATE_LAND,
	} {
		if e&entry > 0 {
			if len(bitmap) > 0 {
				bitmap += " | "
			}
			bitmap += entry.String()
		}
	}
	return bitmap
}
