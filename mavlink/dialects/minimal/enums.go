/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package minimal

import (
	"fmt"
)

// MAV_AUTOPILOT type. Micro air vehicle / autopilot classes. This identifies the individual model.
type MAV_AUTOPILOT int

const (
	// MAV_AUTOPILOT_GENERIC enum. Generic autopilot, full support for everything
	MAV_AUTOPILOT_GENERIC MAV_AUTOPILOT = 0
	// MAV_AUTOPILOT_RESERVED enum. Reserved for future use
	MAV_AUTOPILOT_RESERVED MAV_AUTOPILOT = 1
	// MAV_AUTOPILOT_SLUGS enum. SLUGS autopilot, http://slugsuav.soe.ucsc.edu
	MAV_AUTOPILOT_SLUGS MAV_AUTOPILOT = 2
	// MAV_AUTOPILOT_ARDUPILOTMEGA enum. ArduPilot - Plane/Copter/Rover/Sub/Tracker, https://ardupilot.org
	MAV_AUTOPILOT_ARDUPILOTMEGA MAV_AUTOPILOT = 3
	// MAV_AUTOPILOT_OPENPILOT enum. OpenPilot, http://openpilot.org
	MAV_AUTOPILOT_OPENPILOT MAV_AUTOPILOT = 4
	// MAV_AUTOPILOT_GENERIC_WAYPOINTS_ONLY enum. Generic autopilot only supporting simple waypoints
	MAV_AUTOPILOT_GENERIC_WAYPOINTS_ONLY MAV_AUTOPILOT = 5
	// MAV_AUTOPILOT_GENERIC_WAYPOINTS_AND_SIMPLE_NAVIGATION_ONLY enum. Generic autopilot supporting waypoints and other simple navigation commands
	MAV_AUTOPILOT_GENERIC_WAYPOINTS_AND_SIMPLE_NAVIGATION_ONLY MAV_AUTOPILOT = 6
	// MAV_AUTOPILOT_GENERIC_MISSION_FULL enum. Generic autopilot supporting the full mission command set
	MAV_AUTOPILOT_GENERIC_MISSION_FULL MAV_AUTOPILOT = 7
	// MAV_AUTOPILOT_INVALID enum. No valid autopilot, e.g. a GCS or other MAVLink component
	MAV_AUTOPILOT_INVALID MAV_AUTOPILOT = 8
	// MAV_AUTOPILOT_PPZ enum. PPZ UAV - http://nongnu.org/paparazzi
	MAV_AUTOPILOT_PPZ MAV_AUTOPILOT = 9
	// MAV_AUTOPILOT_UDB enum. UAV Dev Board
	MAV_AUTOPILOT_UDB MAV_AUTOPILOT = 10
	// MAV_AUTOPILOT_FP enum. FlexiPilot
	MAV_AUTOPILOT_FP MAV_AUTOPILOT = 11
	// MAV_AUTOPILOT_PX4 enum. PX4 Autopilot - http://px4.io/
	MAV_AUTOPILOT_PX4 MAV_AUTOPILOT = 12
	// MAV_AUTOPILOT_SMACCMPILOT enum. SMACCMPilot - http://smaccmpilot.org
	MAV_AUTOPILOT_SMACCMPILOT MAV_AUTOPILOT = 13
	// MAV_AUTOPILOT_AUTOQUAD enum. AutoQuad -- http://autoquad.org
	MAV_AUTOPILOT_AUTOQUAD MAV_AUTOPILOT = 14
	// MAV_AUTOPILOT_ARMAZILA enum. Armazila -- http://armazila.com
	MAV_AUTOPILOT_ARMAZILA MAV_AUTOPILOT = 15
	// MAV_AUTOPILOT_AEROB enum. Aerob -- http://aerob.ru
	MAV_AUTOPILOT_AEROB MAV_AUTOPILOT = 16
	// MAV_AUTOPILOT_ASLUAV enum. ASLUAV autopilot -- http://www.asl.ethz.ch
	MAV_AUTOPILOT_ASLUAV MAV_AUTOPILOT = 17
	// MAV_AUTOPILOT_SMARTAP enum. SmartAP Autopilot - http://sky-drones.com
	MAV_AUTOPILOT_SMARTAP MAV_AUTOPILOT = 18
	// MAV_AUTOPILOT_AIRRAILS enum. AirRails - http://uaventure.com
	MAV_AUTOPILOT_AIRRAILS MAV_AUTOPILOT = 19
)

func (e MAV_AUTOPILOT) String() string {
	if name, ok := map[MAV_AUTOPILOT]string{
		MAV_AUTOPILOT_GENERIC:                                      "MAV_AUTOPILOT_GENERIC",
		MAV_AUTOPILOT_RESERVED:                                     "MAV_AUTOPILOT_RESERVED",
		MAV_AUTOPILOT_SLUGS:                                        "MAV_AUTOPILOT_SLUGS",
		MAV_AUTOPILOT_ARDUPILOTMEGA:                                "MAV_AUTOPILOT_ARDUPILOTMEGA",
		MAV_AUTOPILOT_OPENPILOT:                                    "MAV_AUTOPILOT_OPENPILOT",
		MAV_AUTOPILOT_GENERIC_WAYPOINTS_ONLY:                       "MAV_AUTOPILOT_GENERIC_WAYPOINTS_ONLY",
		MAV_AUTOPILOT_GENERIC_WAYPOINTS_AND_SIMPLE_NAVIGATION_ONLY: "MAV_AUTOPILOT_GENERIC_WAYPOINTS_AND_SIMPLE_NAVIGATION_ONLY",
		MAV_AUTOPILOT_GENERIC_MISSION_FULL:                         "MAV_AUTOPILOT_GENERIC_MISSION_FULL",
		MAV_AUTOPILOT_INVALID:                                      "MAV_AUTOPILOT_INVALID",
		MAV_AUTOPILOT_PPZ:                                          "MAV_AUTOPILOT_PPZ",
		MAV_AUTOPILOT_UDB:                                          "MAV_AUTOPILOT_UDB",
		MAV_AUTOPILOT_FP:                                           "MAV_AUTOPILOT_FP",
		MAV_AUTOPILOT_PX4:                                          "MAV_AUTOPILOT_PX4",
		MAV_AUTOPILOT_SMACCMPILOT:                                  "MAV_AUTOPILOT_SMACCMPILOT",
		MAV_AUTOPILOT_AUTOQUAD:                                     "MAV_AUTOPILOT_AUTOQUAD",
		MAV_AUTOPILOT_ARMAZILA:                                     "MAV_AUTOPILOT_ARMAZILA",
		MAV_AUTOPILOT_AEROB:                                        "MAV_AUTOPILOT_AEROB",
		MAV_AUTOPILOT_ASLUAV:                                       "MAV_AUTOPILOT_ASLUAV",
		MAV_AUTOPILOT_SMARTAP:                                      "MAV_AUTOPILOT_SMARTAP",
		MAV_AUTOPILOT_AIRRAILS:                                     "MAV_AUTOPILOT_AIRRAILS",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_AUTOPILOT_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_AUTOPILOT enums
func (e MAV_AUTOPILOT) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_AUTOPILOT{
		MAV_AUTOPILOT_GENERIC,
		MAV_AUTOPILOT_RESERVED,
		MAV_AUTOPILOT_SLUGS,
		MAV_AUTOPILOT_ARDUPILOTMEGA,
		MAV_AUTOPILOT_OPENPILOT,
		MAV_AUTOPILOT_GENERIC_WAYPOINTS_ONLY,
		MAV_AUTOPILOT_GENERIC_WAYPOINTS_AND_SIMPLE_NAVIGATION_ONLY,
		MAV_AUTOPILOT_GENERIC_MISSION_FULL,
		MAV_AUTOPILOT_INVALID,
		MAV_AUTOPILOT_PPZ,
		MAV_AUTOPILOT_UDB,
		MAV_AUTOPILOT_FP,
		MAV_AUTOPILOT_PX4,
		MAV_AUTOPILOT_SMACCMPILOT,
		MAV_AUTOPILOT_AUTOQUAD,
		MAV_AUTOPILOT_ARMAZILA,
		MAV_AUTOPILOT_AEROB,
		MAV_AUTOPILOT_ASLUAV,
		MAV_AUTOPILOT_SMARTAP,
		MAV_AUTOPILOT_AIRRAILS,
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

// MAV_TYPE type. MAVLINK component type reported in HEARTBEAT message. Flight controllers must report the type of the vehicle on which they are mounted (e.g. MAV_TYPE_OCTOROTOR). All other components must report a value appropriate for their type (e.g. a camera must use MAV_TYPE_CAMERA).
type MAV_TYPE int

const (
	// MAV_TYPE_GENERIC enum. Generic micro air vehicle
	MAV_TYPE_GENERIC MAV_TYPE = 0
	// MAV_TYPE_FIXED_WING enum. Fixed wing aircraft
	MAV_TYPE_FIXED_WING MAV_TYPE = 1
	// MAV_TYPE_QUADROTOR enum. Quadrotor
	MAV_TYPE_QUADROTOR MAV_TYPE = 2
	// MAV_TYPE_COAXIAL enum. Coaxial helicopter
	MAV_TYPE_COAXIAL MAV_TYPE = 3
	// MAV_TYPE_HELICOPTER enum. Normal helicopter with tail rotor
	MAV_TYPE_HELICOPTER MAV_TYPE = 4
	// MAV_TYPE_ANTENNA_TRACKER enum. Ground installation
	MAV_TYPE_ANTENNA_TRACKER MAV_TYPE = 5
	// MAV_TYPE_GCS enum. Operator control unit / ground control station
	MAV_TYPE_GCS MAV_TYPE = 6
	// MAV_TYPE_AIRSHIP enum. Airship, controlled
	MAV_TYPE_AIRSHIP MAV_TYPE = 7
	// MAV_TYPE_FREE_BALLOON enum. Free balloon, uncontrolled
	MAV_TYPE_FREE_BALLOON MAV_TYPE = 8
	// MAV_TYPE_ROCKET enum. Rocket
	MAV_TYPE_ROCKET MAV_TYPE = 9
	// MAV_TYPE_GROUND_ROVER enum. Ground rover
	MAV_TYPE_GROUND_ROVER MAV_TYPE = 10
	// MAV_TYPE_SURFACE_BOAT enum. Surface vessel, boat, ship
	MAV_TYPE_SURFACE_BOAT MAV_TYPE = 11
	// MAV_TYPE_SUBMARINE enum. Submarine
	MAV_TYPE_SUBMARINE MAV_TYPE = 12
	// MAV_TYPE_HEXAROTOR enum. Hexarotor
	MAV_TYPE_HEXAROTOR MAV_TYPE = 13
	// MAV_TYPE_OCTOROTOR enum. Octorotor
	MAV_TYPE_OCTOROTOR MAV_TYPE = 14
	// MAV_TYPE_TRICOPTER enum. Tricopter
	MAV_TYPE_TRICOPTER MAV_TYPE = 15
	// MAV_TYPE_FLAPPING_WING enum. Flapping wing
	MAV_TYPE_FLAPPING_WING MAV_TYPE = 16
	// MAV_TYPE_KITE enum. Kite
	MAV_TYPE_KITE MAV_TYPE = 17
	// MAV_TYPE_ONBOARD_CONTROLLER enum. Onboard companion controller
	MAV_TYPE_ONBOARD_CONTROLLER MAV_TYPE = 18
	// MAV_TYPE_VTOL_DUOROTOR enum. Two-rotor VTOL using control surfaces in vertical operation in addition. Tailsitter
	MAV_TYPE_VTOL_DUOROTOR MAV_TYPE = 19
	// MAV_TYPE_VTOL_QUADROTOR enum. Quad-rotor VTOL using a V-shaped quad config in vertical operation. Tailsitter
	MAV_TYPE_VTOL_QUADROTOR MAV_TYPE = 20
	// MAV_TYPE_VTOL_TILTROTOR enum. Tiltrotor VTOL
	MAV_TYPE_VTOL_TILTROTOR MAV_TYPE = 21
	// MAV_TYPE_VTOL_RESERVED2 enum. VTOL reserved 2
	MAV_TYPE_VTOL_RESERVED2 MAV_TYPE = 22
	// MAV_TYPE_VTOL_RESERVED3 enum. VTOL reserved 3
	MAV_TYPE_VTOL_RESERVED3 MAV_TYPE = 23
	// MAV_TYPE_VTOL_RESERVED4 enum. VTOL reserved 4
	MAV_TYPE_VTOL_RESERVED4 MAV_TYPE = 24
	// MAV_TYPE_VTOL_RESERVED5 enum. VTOL reserved 5
	MAV_TYPE_VTOL_RESERVED5 MAV_TYPE = 25
	// MAV_TYPE_GIMBAL enum. Gimbal
	MAV_TYPE_GIMBAL MAV_TYPE = 26
	// MAV_TYPE_ADSB enum. ADSB system
	MAV_TYPE_ADSB MAV_TYPE = 27
	// MAV_TYPE_PARAFOIL enum. Steerable, nonrigid airfoil
	MAV_TYPE_PARAFOIL MAV_TYPE = 28
	// MAV_TYPE_DODECAROTOR enum. Dodecarotor
	MAV_TYPE_DODECAROTOR MAV_TYPE = 29
	// MAV_TYPE_CAMERA enum. Camera
	MAV_TYPE_CAMERA MAV_TYPE = 30
	// MAV_TYPE_CHARGING_STATION enum. Charging station
	MAV_TYPE_CHARGING_STATION MAV_TYPE = 31
	// MAV_TYPE_FLARM enum. FLARM collision avoidance system
	MAV_TYPE_FLARM MAV_TYPE = 32
	// MAV_TYPE_SERVO enum. Servo
	MAV_TYPE_SERVO MAV_TYPE = 33
	// MAV_TYPE_ODID enum. Open Drone ID. See https://mavlink.io/en/services/opendroneid.html
	MAV_TYPE_ODID MAV_TYPE = 34
	// MAV_TYPE_DECAROTOR enum. Decarotor
	MAV_TYPE_DECAROTOR MAV_TYPE = 35
)

func (e MAV_TYPE) String() string {
	if name, ok := map[MAV_TYPE]string{
		MAV_TYPE_GENERIC:            "MAV_TYPE_GENERIC",
		MAV_TYPE_FIXED_WING:         "MAV_TYPE_FIXED_WING",
		MAV_TYPE_QUADROTOR:          "MAV_TYPE_QUADROTOR",
		MAV_TYPE_COAXIAL:            "MAV_TYPE_COAXIAL",
		MAV_TYPE_HELICOPTER:         "MAV_TYPE_HELICOPTER",
		MAV_TYPE_ANTENNA_TRACKER:    "MAV_TYPE_ANTENNA_TRACKER",
		MAV_TYPE_GCS:                "MAV_TYPE_GCS",
		MAV_TYPE_AIRSHIP:            "MAV_TYPE_AIRSHIP",
		MAV_TYPE_FREE_BALLOON:       "MAV_TYPE_FREE_BALLOON",
		MAV_TYPE_ROCKET:             "MAV_TYPE_ROCKET",
		MAV_TYPE_GROUND_ROVER:       "MAV_TYPE_GROUND_ROVER",
		MAV_TYPE_SURFACE_BOAT:       "MAV_TYPE_SURFACE_BOAT",
		MAV_TYPE_SUBMARINE:          "MAV_TYPE_SUBMARINE",
		MAV_TYPE_HEXAROTOR:          "MAV_TYPE_HEXAROTOR",
		MAV_TYPE_OCTOROTOR:          "MAV_TYPE_OCTOROTOR",
		MAV_TYPE_TRICOPTER:          "MAV_TYPE_TRICOPTER",
		MAV_TYPE_FLAPPING_WING:      "MAV_TYPE_FLAPPING_WING",
		MAV_TYPE_KITE:               "MAV_TYPE_KITE",
		MAV_TYPE_ONBOARD_CONTROLLER: "MAV_TYPE_ONBOARD_CONTROLLER",
		MAV_TYPE_VTOL_DUOROTOR:      "MAV_TYPE_VTOL_DUOROTOR",
		MAV_TYPE_VTOL_QUADROTOR:     "MAV_TYPE_VTOL_QUADROTOR",
		MAV_TYPE_VTOL_TILTROTOR:     "MAV_TYPE_VTOL_TILTROTOR",
		MAV_TYPE_VTOL_RESERVED2:     "MAV_TYPE_VTOL_RESERVED2",
		MAV_TYPE_VTOL_RESERVED3:     "MAV_TYPE_VTOL_RESERVED3",
		MAV_TYPE_VTOL_RESERVED4:     "MAV_TYPE_VTOL_RESERVED4",
		MAV_TYPE_VTOL_RESERVED5:     "MAV_TYPE_VTOL_RESERVED5",
		MAV_TYPE_GIMBAL:             "MAV_TYPE_GIMBAL",
		MAV_TYPE_ADSB:               "MAV_TYPE_ADSB",
		MAV_TYPE_PARAFOIL:           "MAV_TYPE_PARAFOIL",
		MAV_TYPE_DODECAROTOR:        "MAV_TYPE_DODECAROTOR",
		MAV_TYPE_CAMERA:             "MAV_TYPE_CAMERA",
		MAV_TYPE_CHARGING_STATION:   "MAV_TYPE_CHARGING_STATION",
		MAV_TYPE_FLARM:              "MAV_TYPE_FLARM",
		MAV_TYPE_SERVO:              "MAV_TYPE_SERVO",
		MAV_TYPE_ODID:               "MAV_TYPE_ODID",
		MAV_TYPE_DECAROTOR:          "MAV_TYPE_DECAROTOR",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_TYPE enums
func (e MAV_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_TYPE{
		MAV_TYPE_GENERIC,
		MAV_TYPE_FIXED_WING,
		MAV_TYPE_QUADROTOR,
		MAV_TYPE_COAXIAL,
		MAV_TYPE_HELICOPTER,
		MAV_TYPE_ANTENNA_TRACKER,
		MAV_TYPE_GCS,
		MAV_TYPE_AIRSHIP,
		MAV_TYPE_FREE_BALLOON,
		MAV_TYPE_ROCKET,
		MAV_TYPE_GROUND_ROVER,
		MAV_TYPE_SURFACE_BOAT,
		MAV_TYPE_SUBMARINE,
		MAV_TYPE_HEXAROTOR,
		MAV_TYPE_OCTOROTOR,
		MAV_TYPE_TRICOPTER,
		MAV_TYPE_FLAPPING_WING,
		MAV_TYPE_KITE,
		MAV_TYPE_ONBOARD_CONTROLLER,
		MAV_TYPE_VTOL_DUOROTOR,
		MAV_TYPE_VTOL_QUADROTOR,
		MAV_TYPE_VTOL_TILTROTOR,
		MAV_TYPE_VTOL_RESERVED2,
		MAV_TYPE_VTOL_RESERVED3,
		MAV_TYPE_VTOL_RESERVED4,
		MAV_TYPE_VTOL_RESERVED5,
		MAV_TYPE_GIMBAL,
		MAV_TYPE_ADSB,
		MAV_TYPE_PARAFOIL,
		MAV_TYPE_DODECAROTOR,
		MAV_TYPE_CAMERA,
		MAV_TYPE_CHARGING_STATION,
		MAV_TYPE_FLARM,
		MAV_TYPE_SERVO,
		MAV_TYPE_ODID,
		MAV_TYPE_DECAROTOR,
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

// MAV_MODE_FLAG type. These flags encode the MAV mode.
type MAV_MODE_FLAG int

const (
	// MAV_MODE_FLAG_SAFETY_ARMED enum. 0b10000000 MAV safety set to armed. Motors are enabled / running / can start. Ready to fly. Additional note: this flag is to be ignore when sent in the command MAV_CMD_DO_SET_MODE and MAV_CMD_COMPONENT_ARM_DISARM shall be used instead. The flag can still be used to report the armed state
	MAV_MODE_FLAG_SAFETY_ARMED MAV_MODE_FLAG = 128
	// MAV_MODE_FLAG_MANUAL_INPUT_ENABLED enum. 0b01000000 remote control input is enabled
	MAV_MODE_FLAG_MANUAL_INPUT_ENABLED MAV_MODE_FLAG = 64
	// MAV_MODE_FLAG_HIL_ENABLED enum. 0b00100000 hardware in the loop simulation. All motors / actuators are blocked, but internal software is full operational
	MAV_MODE_FLAG_HIL_ENABLED MAV_MODE_FLAG = 32
	// MAV_MODE_FLAG_STABILIZE_ENABLED enum. 0b00010000 system stabilizes electronically its attitude (and optionally position). It needs however further control inputs to move around
	MAV_MODE_FLAG_STABILIZE_ENABLED MAV_MODE_FLAG = 16
	// MAV_MODE_FLAG_GUIDED_ENABLED enum. 0b00001000 guided mode enabled, system flies waypoints / mission items
	MAV_MODE_FLAG_GUIDED_ENABLED MAV_MODE_FLAG = 8
	// MAV_MODE_FLAG_AUTO_ENABLED enum. 0b00000100 autonomous mode enabled, system finds its own goal positions. Guided flag can be set or not, depends on the actual implementation
	MAV_MODE_FLAG_AUTO_ENABLED MAV_MODE_FLAG = 4
	// MAV_MODE_FLAG_TEST_ENABLED enum. 0b00000010 system has a test mode enabled. This flag is intended for temporary system tests and should not be used for stable implementations
	MAV_MODE_FLAG_TEST_ENABLED MAV_MODE_FLAG = 2
	// MAV_MODE_FLAG_CUSTOM_MODE_ENABLED enum. 0b00000001 Reserved for future use
	MAV_MODE_FLAG_CUSTOM_MODE_ENABLED MAV_MODE_FLAG = 1
)

func (e MAV_MODE_FLAG) String() string {
	if name, ok := map[MAV_MODE_FLAG]string{
		MAV_MODE_FLAG_SAFETY_ARMED:         "MAV_MODE_FLAG_SAFETY_ARMED",
		MAV_MODE_FLAG_MANUAL_INPUT_ENABLED: "MAV_MODE_FLAG_MANUAL_INPUT_ENABLED",
		MAV_MODE_FLAG_HIL_ENABLED:          "MAV_MODE_FLAG_HIL_ENABLED",
		MAV_MODE_FLAG_STABILIZE_ENABLED:    "MAV_MODE_FLAG_STABILIZE_ENABLED",
		MAV_MODE_FLAG_GUIDED_ENABLED:       "MAV_MODE_FLAG_GUIDED_ENABLED",
		MAV_MODE_FLAG_AUTO_ENABLED:         "MAV_MODE_FLAG_AUTO_ENABLED",
		MAV_MODE_FLAG_TEST_ENABLED:         "MAV_MODE_FLAG_TEST_ENABLED",
		MAV_MODE_FLAG_CUSTOM_MODE_ENABLED:  "MAV_MODE_FLAG_CUSTOM_MODE_ENABLED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_MODE_FLAG_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_MODE_FLAG enums
func (e MAV_MODE_FLAG) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_MODE_FLAG{
		MAV_MODE_FLAG_SAFETY_ARMED,
		MAV_MODE_FLAG_MANUAL_INPUT_ENABLED,
		MAV_MODE_FLAG_HIL_ENABLED,
		MAV_MODE_FLAG_STABILIZE_ENABLED,
		MAV_MODE_FLAG_GUIDED_ENABLED,
		MAV_MODE_FLAG_AUTO_ENABLED,
		MAV_MODE_FLAG_TEST_ENABLED,
		MAV_MODE_FLAG_CUSTOM_MODE_ENABLED,
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

// MAV_MODE_FLAG_DECODE_POSITION type. These values encode the bit positions of the decode position. These values can be used to read the value of a flag bit by combining the base_mode variable with AND with the flag position value. The result will be either 0 or 1, depending on if the flag is set or not.
type MAV_MODE_FLAG_DECODE_POSITION int

const (
	// MAV_MODE_FLAG_DECODE_POSITION_SAFETY enum. First bit:  10000000
	MAV_MODE_FLAG_DECODE_POSITION_SAFETY MAV_MODE_FLAG_DECODE_POSITION = 128
	// MAV_MODE_FLAG_DECODE_POSITION_MANUAL enum. Second bit: 01000000
	MAV_MODE_FLAG_DECODE_POSITION_MANUAL MAV_MODE_FLAG_DECODE_POSITION = 64
	// MAV_MODE_FLAG_DECODE_POSITION_HIL enum. Third bit:  00100000
	MAV_MODE_FLAG_DECODE_POSITION_HIL MAV_MODE_FLAG_DECODE_POSITION = 32
	// MAV_MODE_FLAG_DECODE_POSITION_STABILIZE enum. Fourth bit: 00010000
	MAV_MODE_FLAG_DECODE_POSITION_STABILIZE MAV_MODE_FLAG_DECODE_POSITION = 16
	// MAV_MODE_FLAG_DECODE_POSITION_GUIDED enum. Fifth bit:  00001000
	MAV_MODE_FLAG_DECODE_POSITION_GUIDED MAV_MODE_FLAG_DECODE_POSITION = 8
	// MAV_MODE_FLAG_DECODE_POSITION_AUTO enum. Sixth bit:   00000100
	MAV_MODE_FLAG_DECODE_POSITION_AUTO MAV_MODE_FLAG_DECODE_POSITION = 4
	// MAV_MODE_FLAG_DECODE_POSITION_TEST enum. Seventh bit: 00000010
	MAV_MODE_FLAG_DECODE_POSITION_TEST MAV_MODE_FLAG_DECODE_POSITION = 2
	// MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE enum. Eighth bit: 00000001
	MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE MAV_MODE_FLAG_DECODE_POSITION = 1
)

func (e MAV_MODE_FLAG_DECODE_POSITION) String() string {
	if name, ok := map[MAV_MODE_FLAG_DECODE_POSITION]string{
		MAV_MODE_FLAG_DECODE_POSITION_SAFETY:      "MAV_MODE_FLAG_DECODE_POSITION_SAFETY",
		MAV_MODE_FLAG_DECODE_POSITION_MANUAL:      "MAV_MODE_FLAG_DECODE_POSITION_MANUAL",
		MAV_MODE_FLAG_DECODE_POSITION_HIL:         "MAV_MODE_FLAG_DECODE_POSITION_HIL",
		MAV_MODE_FLAG_DECODE_POSITION_STABILIZE:   "MAV_MODE_FLAG_DECODE_POSITION_STABILIZE",
		MAV_MODE_FLAG_DECODE_POSITION_GUIDED:      "MAV_MODE_FLAG_DECODE_POSITION_GUIDED",
		MAV_MODE_FLAG_DECODE_POSITION_AUTO:        "MAV_MODE_FLAG_DECODE_POSITION_AUTO",
		MAV_MODE_FLAG_DECODE_POSITION_TEST:        "MAV_MODE_FLAG_DECODE_POSITION_TEST",
		MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE: "MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_MODE_FLAG_DECODE_POSITION_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_MODE_FLAG_DECODE_POSITION enums
func (e MAV_MODE_FLAG_DECODE_POSITION) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_MODE_FLAG_DECODE_POSITION{
		MAV_MODE_FLAG_DECODE_POSITION_SAFETY,
		MAV_MODE_FLAG_DECODE_POSITION_MANUAL,
		MAV_MODE_FLAG_DECODE_POSITION_HIL,
		MAV_MODE_FLAG_DECODE_POSITION_STABILIZE,
		MAV_MODE_FLAG_DECODE_POSITION_GUIDED,
		MAV_MODE_FLAG_DECODE_POSITION_AUTO,
		MAV_MODE_FLAG_DECODE_POSITION_TEST,
		MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE,
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

// MAV_STATE type
type MAV_STATE int

const (
	// MAV_STATE_UNINIT enum. Uninitialized system, state is unknown
	MAV_STATE_UNINIT MAV_STATE = 0
	// MAV_STATE_BOOT enum. System is booting up
	MAV_STATE_BOOT MAV_STATE = 1
	// MAV_STATE_CALIBRATING enum. System is calibrating and not flight-ready
	MAV_STATE_CALIBRATING MAV_STATE = 2
	// MAV_STATE_STANDBY enum. System is grounded and on standby. It can be launched any time
	MAV_STATE_STANDBY MAV_STATE = 3
	// MAV_STATE_ACTIVE enum. System is active and might be already airborne. Motors are engaged
	MAV_STATE_ACTIVE MAV_STATE = 4
	// MAV_STATE_CRITICAL enum. System is in a non-normal flight mode. It can however still navigate
	MAV_STATE_CRITICAL MAV_STATE = 5
	// MAV_STATE_EMERGENCY enum. System is in a non-normal flight mode. It lost control over parts or over the whole airframe. It is in mayday and going down
	MAV_STATE_EMERGENCY MAV_STATE = 6
	// MAV_STATE_POWEROFF enum. System just initialized its power-down sequence, will shut down now
	MAV_STATE_POWEROFF MAV_STATE = 7
	// MAV_STATE_FLIGHT_TERMINATION enum. System is terminating itself
	MAV_STATE_FLIGHT_TERMINATION MAV_STATE = 8
)

func (e MAV_STATE) String() string {
	if name, ok := map[MAV_STATE]string{
		MAV_STATE_UNINIT:             "MAV_STATE_UNINIT",
		MAV_STATE_BOOT:               "MAV_STATE_BOOT",
		MAV_STATE_CALIBRATING:        "MAV_STATE_CALIBRATING",
		MAV_STATE_STANDBY:            "MAV_STATE_STANDBY",
		MAV_STATE_ACTIVE:             "MAV_STATE_ACTIVE",
		MAV_STATE_CRITICAL:           "MAV_STATE_CRITICAL",
		MAV_STATE_EMERGENCY:          "MAV_STATE_EMERGENCY",
		MAV_STATE_POWEROFF:           "MAV_STATE_POWEROFF",
		MAV_STATE_FLIGHT_TERMINATION: "MAV_STATE_FLIGHT_TERMINATION",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_STATE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_STATE enums
func (e MAV_STATE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_STATE{
		MAV_STATE_UNINIT,
		MAV_STATE_BOOT,
		MAV_STATE_CALIBRATING,
		MAV_STATE_STANDBY,
		MAV_STATE_ACTIVE,
		MAV_STATE_CRITICAL,
		MAV_STATE_EMERGENCY,
		MAV_STATE_POWEROFF,
		MAV_STATE_FLIGHT_TERMINATION,
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

// MAV_COMPONENT type. Component ids (values) for the different types and instances of onboard hardware/software that might make up a MAVLink system (autopilot, cameras, servos, GPS systems, avoidance systems etc.).       Components must use the appropriate ID in their source address when sending messages. Components can also use IDs to determine if they are the intended recipient of an incoming message. The MAV_COMP_ID_ALL value is used to indicate messages that must be processed by all components.       When creating new entries, components that can have multiple instances (e.g. cameras, servos etc.) should be allocated sequential values. An appropriate number of values should be left free after these components to allow the number of instances to be expanded.
type MAV_COMPONENT int

const (
	// MAV_COMP_ID_ALL enum. Target id (target_component) used to broadcast messages to all components of the receiving system. Components should attempt to process messages with this component ID and forward to components on any other interfaces. Note: This is not a valid *source* component id for a message
	MAV_COMP_ID_ALL MAV_COMPONENT = 0
	// MAV_COMP_ID_AUTOPILOT1 enum. System flight controller component ("autopilot"). Only one autopilot is expected in a particular system
	MAV_COMP_ID_AUTOPILOT1 MAV_COMPONENT = 1
	// MAV_COMP_ID_USER1 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER1 MAV_COMPONENT = 25
	// MAV_COMP_ID_USER2 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER2 MAV_COMPONENT = 26
	// MAV_COMP_ID_USER3 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER3 MAV_COMPONENT = 27
	// MAV_COMP_ID_USER4 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER4 MAV_COMPONENT = 28
	// MAV_COMP_ID_USER5 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER5 MAV_COMPONENT = 29
	// MAV_COMP_ID_USER6 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER6 MAV_COMPONENT = 30
	// MAV_COMP_ID_USER7 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER7 MAV_COMPONENT = 31
	// MAV_COMP_ID_USER8 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER8 MAV_COMPONENT = 32
	// MAV_COMP_ID_USER9 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER9 MAV_COMPONENT = 33
	// MAV_COMP_ID_USER10 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER10 MAV_COMPONENT = 34
	// MAV_COMP_ID_USER11 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER11 MAV_COMPONENT = 35
	// MAV_COMP_ID_USER12 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER12 MAV_COMPONENT = 36
	// MAV_COMP_ID_USER13 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER13 MAV_COMPONENT = 37
	// MAV_COMP_ID_USER14 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER14 MAV_COMPONENT = 38
	// MAV_COMP_ID_USER15 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER15 MAV_COMPONENT = 39
	// MAV_COMP_ID_USER16 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER16 MAV_COMPONENT = 40
	// MAV_COMP_ID_USER17 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER17 MAV_COMPONENT = 41
	// MAV_COMP_ID_USER18 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER18 MAV_COMPONENT = 42
	// MAV_COMP_ID_USER19 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER19 MAV_COMPONENT = 43
	// MAV_COMP_ID_USER20 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER20 MAV_COMPONENT = 44
	// MAV_COMP_ID_USER21 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER21 MAV_COMPONENT = 45
	// MAV_COMP_ID_USER22 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER22 MAV_COMPONENT = 46
	// MAV_COMP_ID_USER23 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER23 MAV_COMPONENT = 47
	// MAV_COMP_ID_USER24 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER24 MAV_COMPONENT = 48
	// MAV_COMP_ID_USER25 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER25 MAV_COMPONENT = 49
	// MAV_COMP_ID_USER26 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER26 MAV_COMPONENT = 50
	// MAV_COMP_ID_USER27 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER27 MAV_COMPONENT = 51
	// MAV_COMP_ID_USER28 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER28 MAV_COMPONENT = 52
	// MAV_COMP_ID_USER29 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER29 MAV_COMPONENT = 53
	// MAV_COMP_ID_USER30 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER30 MAV_COMPONENT = 54
	// MAV_COMP_ID_USER31 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER31 MAV_COMPONENT = 55
	// MAV_COMP_ID_USER32 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER32 MAV_COMPONENT = 56
	// MAV_COMP_ID_USER33 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER33 MAV_COMPONENT = 57
	// MAV_COMP_ID_USER34 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER34 MAV_COMPONENT = 58
	// MAV_COMP_ID_USER35 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER35 MAV_COMPONENT = 59
	// MAV_COMP_ID_USER36 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER36 MAV_COMPONENT = 60
	// MAV_COMP_ID_USER37 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER37 MAV_COMPONENT = 61
	// MAV_COMP_ID_USER38 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER38 MAV_COMPONENT = 62
	// MAV_COMP_ID_USER39 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER39 MAV_COMPONENT = 63
	// MAV_COMP_ID_USER40 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER40 MAV_COMPONENT = 64
	// MAV_COMP_ID_USER41 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER41 MAV_COMPONENT = 65
	// MAV_COMP_ID_USER42 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER42 MAV_COMPONENT = 66
	// MAV_COMP_ID_USER43 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER43 MAV_COMPONENT = 67
	// MAV_COMP_ID_TELEMETRY_RADIO enum. Telemetry radio (e.g. SiK radio, or other component that emits RADIO_STATUS messages)
	MAV_COMP_ID_TELEMETRY_RADIO MAV_COMPONENT = 68
	// MAV_COMP_ID_USER45 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER45 MAV_COMPONENT = 69
	// MAV_COMP_ID_USER46 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER46 MAV_COMPONENT = 70
	// MAV_COMP_ID_USER47 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER47 MAV_COMPONENT = 71
	// MAV_COMP_ID_USER48 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER48 MAV_COMPONENT = 72
	// MAV_COMP_ID_USER49 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER49 MAV_COMPONENT = 73
	// MAV_COMP_ID_USER50 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER50 MAV_COMPONENT = 74
	// MAV_COMP_ID_USER51 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER51 MAV_COMPONENT = 75
	// MAV_COMP_ID_USER52 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER52 MAV_COMPONENT = 76
	// MAV_COMP_ID_USER53 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER53 MAV_COMPONENT = 77
	// MAV_COMP_ID_USER54 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER54 MAV_COMPONENT = 78
	// MAV_COMP_ID_USER55 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER55 MAV_COMPONENT = 79
	// MAV_COMP_ID_USER56 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER56 MAV_COMPONENT = 80
	// MAV_COMP_ID_USER57 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER57 MAV_COMPONENT = 81
	// MAV_COMP_ID_USER58 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER58 MAV_COMPONENT = 82
	// MAV_COMP_ID_USER59 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER59 MAV_COMPONENT = 83
	// MAV_COMP_ID_USER60 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER60 MAV_COMPONENT = 84
	// MAV_COMP_ID_USER61 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER61 MAV_COMPONENT = 85
	// MAV_COMP_ID_USER62 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER62 MAV_COMPONENT = 86
	// MAV_COMP_ID_USER63 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER63 MAV_COMPONENT = 87
	// MAV_COMP_ID_USER64 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER64 MAV_COMPONENT = 88
	// MAV_COMP_ID_USER65 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER65 MAV_COMPONENT = 89
	// MAV_COMP_ID_USER66 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER66 MAV_COMPONENT = 90
	// MAV_COMP_ID_USER67 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER67 MAV_COMPONENT = 91
	// MAV_COMP_ID_USER68 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER68 MAV_COMPONENT = 92
	// MAV_COMP_ID_USER69 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER69 MAV_COMPONENT = 93
	// MAV_COMP_ID_USER70 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER70 MAV_COMPONENT = 94
	// MAV_COMP_ID_USER71 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER71 MAV_COMPONENT = 95
	// MAV_COMP_ID_USER72 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER72 MAV_COMPONENT = 96
	// MAV_COMP_ID_USER73 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER73 MAV_COMPONENT = 97
	// MAV_COMP_ID_USER74 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER74 MAV_COMPONENT = 98
	// MAV_COMP_ID_USER75 enum. Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network
	MAV_COMP_ID_USER75 MAV_COMPONENT = 99
	// MAV_COMP_ID_CAMERA enum. Camera #1
	MAV_COMP_ID_CAMERA MAV_COMPONENT = 100
	// MAV_COMP_ID_CAMERA2 enum. Camera #2
	MAV_COMP_ID_CAMERA2 MAV_COMPONENT = 101
	// MAV_COMP_ID_CAMERA3 enum. Camera #3
	MAV_COMP_ID_CAMERA3 MAV_COMPONENT = 102
	// MAV_COMP_ID_CAMERA4 enum. Camera #4
	MAV_COMP_ID_CAMERA4 MAV_COMPONENT = 103
	// MAV_COMP_ID_CAMERA5 enum. Camera #5
	MAV_COMP_ID_CAMERA5 MAV_COMPONENT = 104
	// MAV_COMP_ID_CAMERA6 enum. Camera #6
	MAV_COMP_ID_CAMERA6 MAV_COMPONENT = 105
	// MAV_COMP_ID_SERVO1 enum. Servo #1
	MAV_COMP_ID_SERVO1 MAV_COMPONENT = 140
	// MAV_COMP_ID_SERVO2 enum. Servo #2
	MAV_COMP_ID_SERVO2 MAV_COMPONENT = 141
	// MAV_COMP_ID_SERVO3 enum. Servo #3
	MAV_COMP_ID_SERVO3 MAV_COMPONENT = 142
	// MAV_COMP_ID_SERVO4 enum. Servo #4
	MAV_COMP_ID_SERVO4 MAV_COMPONENT = 143
	// MAV_COMP_ID_SERVO5 enum. Servo #5
	MAV_COMP_ID_SERVO5 MAV_COMPONENT = 144
	// MAV_COMP_ID_SERVO6 enum. Servo #6
	MAV_COMP_ID_SERVO6 MAV_COMPONENT = 145
	// MAV_COMP_ID_SERVO7 enum. Servo #7
	MAV_COMP_ID_SERVO7 MAV_COMPONENT = 146
	// MAV_COMP_ID_SERVO8 enum. Servo #8
	MAV_COMP_ID_SERVO8 MAV_COMPONENT = 147
	// MAV_COMP_ID_SERVO9 enum. Servo #9
	MAV_COMP_ID_SERVO9 MAV_COMPONENT = 148
	// MAV_COMP_ID_SERVO10 enum. Servo #10
	MAV_COMP_ID_SERVO10 MAV_COMPONENT = 149
	// MAV_COMP_ID_SERVO11 enum. Servo #11
	MAV_COMP_ID_SERVO11 MAV_COMPONENT = 150
	// MAV_COMP_ID_SERVO12 enum. Servo #12
	MAV_COMP_ID_SERVO12 MAV_COMPONENT = 151
	// MAV_COMP_ID_SERVO13 enum. Servo #13
	MAV_COMP_ID_SERVO13 MAV_COMPONENT = 152
	// MAV_COMP_ID_SERVO14 enum. Servo #14
	MAV_COMP_ID_SERVO14 MAV_COMPONENT = 153
	// MAV_COMP_ID_GIMBAL enum. Gimbal #1
	MAV_COMP_ID_GIMBAL MAV_COMPONENT = 154
	// MAV_COMP_ID_LOG enum. Logging component
	MAV_COMP_ID_LOG MAV_COMPONENT = 155
	// MAV_COMP_ID_ADSB enum. Automatic Dependent Surveillance-Broadcast (ADS-B) component
	MAV_COMP_ID_ADSB MAV_COMPONENT = 156
	// MAV_COMP_ID_OSD enum. On Screen Display (OSD) devices for video links
	MAV_COMP_ID_OSD MAV_COMPONENT = 157
	// MAV_COMP_ID_PERIPHERAL enum. Generic autopilot peripheral component ID. Meant for devices that do not implement the parameter microservice
	MAV_COMP_ID_PERIPHERAL MAV_COMPONENT = 158
	// MAV_COMP_ID_QX1_GIMBAL enum. Gimbal ID for QX1
	MAV_COMP_ID_QX1_GIMBAL MAV_COMPONENT = 159
	// MAV_COMP_ID_FLARM enum. FLARM collision alert component
	MAV_COMP_ID_FLARM MAV_COMPONENT = 160
	// MAV_COMP_ID_GIMBAL2 enum. Gimbal #2
	MAV_COMP_ID_GIMBAL2 MAV_COMPONENT = 171
	// MAV_COMP_ID_GIMBAL3 enum. Gimbal #3
	MAV_COMP_ID_GIMBAL3 MAV_COMPONENT = 172
	// MAV_COMP_ID_GIMBAL4 enum. Gimbal #4
	MAV_COMP_ID_GIMBAL4 MAV_COMPONENT = 173
	// MAV_COMP_ID_GIMBAL5 enum. Gimbal #5
	MAV_COMP_ID_GIMBAL5 MAV_COMPONENT = 174
	// MAV_COMP_ID_GIMBAL6 enum. Gimbal #6
	MAV_COMP_ID_GIMBAL6 MAV_COMPONENT = 175
	// MAV_COMP_ID_MISSIONPLANNER enum. Component that can generate/supply a mission flight plan (e.g. GCS or developer API)
	MAV_COMP_ID_MISSIONPLANNER MAV_COMPONENT = 190
	// MAV_COMP_ID_ONBOARD_COMPUTER enum. Component that lives on the onboard computer (companion computer) and has some generic functionalities, such as settings system parameters and monitoring the status of some processes that don't directly speak mavlink and so on
	MAV_COMP_ID_ONBOARD_COMPUTER MAV_COMPONENT = 191
	// MAV_COMP_ID_PATHPLANNER enum. Component that finds an optimal path between points based on a certain constraint (e.g. minimum snap, shortest path, cost, etc.)
	MAV_COMP_ID_PATHPLANNER MAV_COMPONENT = 195
	// MAV_COMP_ID_OBSTACLE_AVOIDANCE enum. Component that plans a collision free path between two points
	MAV_COMP_ID_OBSTACLE_AVOIDANCE MAV_COMPONENT = 196
	// MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY enum. Component that provides position estimates using VIO techniques
	MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY MAV_COMPONENT = 197
	// MAV_COMP_ID_PAIRING_MANAGER enum. Component that manages pairing of vehicle and GCS
	MAV_COMP_ID_PAIRING_MANAGER MAV_COMPONENT = 198
	// MAV_COMP_ID_IMU enum. Inertial Measurement Unit (IMU) #1
	MAV_COMP_ID_IMU MAV_COMPONENT = 200
	// MAV_COMP_ID_IMU_2 enum. Inertial Measurement Unit (IMU) #2
	MAV_COMP_ID_IMU_2 MAV_COMPONENT = 201
	// MAV_COMP_ID_IMU_3 enum. Inertial Measurement Unit (IMU) #3
	MAV_COMP_ID_IMU_3 MAV_COMPONENT = 202
	// MAV_COMP_ID_GPS enum. GPS #1
	MAV_COMP_ID_GPS MAV_COMPONENT = 220
	// MAV_COMP_ID_GPS2 enum. GPS #2
	MAV_COMP_ID_GPS2 MAV_COMPONENT = 221
	// MAV_COMP_ID_ODID_TXRX_1 enum. Open Drone ID transmitter/receiver (Bluetooth/WiFi/Internet)
	MAV_COMP_ID_ODID_TXRX_1 MAV_COMPONENT = 236
	// MAV_COMP_ID_ODID_TXRX_2 enum. Open Drone ID transmitter/receiver (Bluetooth/WiFi/Internet)
	MAV_COMP_ID_ODID_TXRX_2 MAV_COMPONENT = 237
	// MAV_COMP_ID_ODID_TXRX_3 enum. Open Drone ID transmitter/receiver (Bluetooth/WiFi/Internet)
	MAV_COMP_ID_ODID_TXRX_3 MAV_COMPONENT = 238
	// MAV_COMP_ID_UDP_BRIDGE enum. Component to bridge MAVLink to UDP (i.e. from a UART)
	MAV_COMP_ID_UDP_BRIDGE MAV_COMPONENT = 240
	// MAV_COMP_ID_UART_BRIDGE enum. Component to bridge to UART (i.e. from UDP)
	MAV_COMP_ID_UART_BRIDGE MAV_COMPONENT = 241
	// MAV_COMP_ID_TUNNEL_NODE enum. Component handling TUNNEL messages (e.g. vendor specific GUI of a component)
	MAV_COMP_ID_TUNNEL_NODE MAV_COMPONENT = 242
	// MAV_COMP_ID_SYSTEM_CONTROL enum. Component for handling system messages (e.g. to ARM, takeoff, etc.)
	MAV_COMP_ID_SYSTEM_CONTROL MAV_COMPONENT = 250
)

func (e MAV_COMPONENT) String() string {
	if name, ok := map[MAV_COMPONENT]string{
		MAV_COMP_ID_ALL:                      "MAV_COMP_ID_ALL",
		MAV_COMP_ID_AUTOPILOT1:               "MAV_COMP_ID_AUTOPILOT1",
		MAV_COMP_ID_USER1:                    "MAV_COMP_ID_USER1",
		MAV_COMP_ID_USER2:                    "MAV_COMP_ID_USER2",
		MAV_COMP_ID_USER3:                    "MAV_COMP_ID_USER3",
		MAV_COMP_ID_USER4:                    "MAV_COMP_ID_USER4",
		MAV_COMP_ID_USER5:                    "MAV_COMP_ID_USER5",
		MAV_COMP_ID_USER6:                    "MAV_COMP_ID_USER6",
		MAV_COMP_ID_USER7:                    "MAV_COMP_ID_USER7",
		MAV_COMP_ID_USER8:                    "MAV_COMP_ID_USER8",
		MAV_COMP_ID_USER9:                    "MAV_COMP_ID_USER9",
		MAV_COMP_ID_USER10:                   "MAV_COMP_ID_USER10",
		MAV_COMP_ID_USER11:                   "MAV_COMP_ID_USER11",
		MAV_COMP_ID_USER12:                   "MAV_COMP_ID_USER12",
		MAV_COMP_ID_USER13:                   "MAV_COMP_ID_USER13",
		MAV_COMP_ID_USER14:                   "MAV_COMP_ID_USER14",
		MAV_COMP_ID_USER15:                   "MAV_COMP_ID_USER15",
		MAV_COMP_ID_USER16:                   "MAV_COMP_ID_USER16",
		MAV_COMP_ID_USER17:                   "MAV_COMP_ID_USER17",
		MAV_COMP_ID_USER18:                   "MAV_COMP_ID_USER18",
		MAV_COMP_ID_USER19:                   "MAV_COMP_ID_USER19",
		MAV_COMP_ID_USER20:                   "MAV_COMP_ID_USER20",
		MAV_COMP_ID_USER21:                   "MAV_COMP_ID_USER21",
		MAV_COMP_ID_USER22:                   "MAV_COMP_ID_USER22",
		MAV_COMP_ID_USER23:                   "MAV_COMP_ID_USER23",
		MAV_COMP_ID_USER24:                   "MAV_COMP_ID_USER24",
		MAV_COMP_ID_USER25:                   "MAV_COMP_ID_USER25",
		MAV_COMP_ID_USER26:                   "MAV_COMP_ID_USER26",
		MAV_COMP_ID_USER27:                   "MAV_COMP_ID_USER27",
		MAV_COMP_ID_USER28:                   "MAV_COMP_ID_USER28",
		MAV_COMP_ID_USER29:                   "MAV_COMP_ID_USER29",
		MAV_COMP_ID_USER30:                   "MAV_COMP_ID_USER30",
		MAV_COMP_ID_USER31:                   "MAV_COMP_ID_USER31",
		MAV_COMP_ID_USER32:                   "MAV_COMP_ID_USER32",
		MAV_COMP_ID_USER33:                   "MAV_COMP_ID_USER33",
		MAV_COMP_ID_USER34:                   "MAV_COMP_ID_USER34",
		MAV_COMP_ID_USER35:                   "MAV_COMP_ID_USER35",
		MAV_COMP_ID_USER36:                   "MAV_COMP_ID_USER36",
		MAV_COMP_ID_USER37:                   "MAV_COMP_ID_USER37",
		MAV_COMP_ID_USER38:                   "MAV_COMP_ID_USER38",
		MAV_COMP_ID_USER39:                   "MAV_COMP_ID_USER39",
		MAV_COMP_ID_USER40:                   "MAV_COMP_ID_USER40",
		MAV_COMP_ID_USER41:                   "MAV_COMP_ID_USER41",
		MAV_COMP_ID_USER42:                   "MAV_COMP_ID_USER42",
		MAV_COMP_ID_USER43:                   "MAV_COMP_ID_USER43",
		MAV_COMP_ID_TELEMETRY_RADIO:          "MAV_COMP_ID_TELEMETRY_RADIO",
		MAV_COMP_ID_USER45:                   "MAV_COMP_ID_USER45",
		MAV_COMP_ID_USER46:                   "MAV_COMP_ID_USER46",
		MAV_COMP_ID_USER47:                   "MAV_COMP_ID_USER47",
		MAV_COMP_ID_USER48:                   "MAV_COMP_ID_USER48",
		MAV_COMP_ID_USER49:                   "MAV_COMP_ID_USER49",
		MAV_COMP_ID_USER50:                   "MAV_COMP_ID_USER50",
		MAV_COMP_ID_USER51:                   "MAV_COMP_ID_USER51",
		MAV_COMP_ID_USER52:                   "MAV_COMP_ID_USER52",
		MAV_COMP_ID_USER53:                   "MAV_COMP_ID_USER53",
		MAV_COMP_ID_USER54:                   "MAV_COMP_ID_USER54",
		MAV_COMP_ID_USER55:                   "MAV_COMP_ID_USER55",
		MAV_COMP_ID_USER56:                   "MAV_COMP_ID_USER56",
		MAV_COMP_ID_USER57:                   "MAV_COMP_ID_USER57",
		MAV_COMP_ID_USER58:                   "MAV_COMP_ID_USER58",
		MAV_COMP_ID_USER59:                   "MAV_COMP_ID_USER59",
		MAV_COMP_ID_USER60:                   "MAV_COMP_ID_USER60",
		MAV_COMP_ID_USER61:                   "MAV_COMP_ID_USER61",
		MAV_COMP_ID_USER62:                   "MAV_COMP_ID_USER62",
		MAV_COMP_ID_USER63:                   "MAV_COMP_ID_USER63",
		MAV_COMP_ID_USER64:                   "MAV_COMP_ID_USER64",
		MAV_COMP_ID_USER65:                   "MAV_COMP_ID_USER65",
		MAV_COMP_ID_USER66:                   "MAV_COMP_ID_USER66",
		MAV_COMP_ID_USER67:                   "MAV_COMP_ID_USER67",
		MAV_COMP_ID_USER68:                   "MAV_COMP_ID_USER68",
		MAV_COMP_ID_USER69:                   "MAV_COMP_ID_USER69",
		MAV_COMP_ID_USER70:                   "MAV_COMP_ID_USER70",
		MAV_COMP_ID_USER71:                   "MAV_COMP_ID_USER71",
		MAV_COMP_ID_USER72:                   "MAV_COMP_ID_USER72",
		MAV_COMP_ID_USER73:                   "MAV_COMP_ID_USER73",
		MAV_COMP_ID_USER74:                   "MAV_COMP_ID_USER74",
		MAV_COMP_ID_USER75:                   "MAV_COMP_ID_USER75",
		MAV_COMP_ID_CAMERA:                   "MAV_COMP_ID_CAMERA",
		MAV_COMP_ID_CAMERA2:                  "MAV_COMP_ID_CAMERA2",
		MAV_COMP_ID_CAMERA3:                  "MAV_COMP_ID_CAMERA3",
		MAV_COMP_ID_CAMERA4:                  "MAV_COMP_ID_CAMERA4",
		MAV_COMP_ID_CAMERA5:                  "MAV_COMP_ID_CAMERA5",
		MAV_COMP_ID_CAMERA6:                  "MAV_COMP_ID_CAMERA6",
		MAV_COMP_ID_SERVO1:                   "MAV_COMP_ID_SERVO1",
		MAV_COMP_ID_SERVO2:                   "MAV_COMP_ID_SERVO2",
		MAV_COMP_ID_SERVO3:                   "MAV_COMP_ID_SERVO3",
		MAV_COMP_ID_SERVO4:                   "MAV_COMP_ID_SERVO4",
		MAV_COMP_ID_SERVO5:                   "MAV_COMP_ID_SERVO5",
		MAV_COMP_ID_SERVO6:                   "MAV_COMP_ID_SERVO6",
		MAV_COMP_ID_SERVO7:                   "MAV_COMP_ID_SERVO7",
		MAV_COMP_ID_SERVO8:                   "MAV_COMP_ID_SERVO8",
		MAV_COMP_ID_SERVO9:                   "MAV_COMP_ID_SERVO9",
		MAV_COMP_ID_SERVO10:                  "MAV_COMP_ID_SERVO10",
		MAV_COMP_ID_SERVO11:                  "MAV_COMP_ID_SERVO11",
		MAV_COMP_ID_SERVO12:                  "MAV_COMP_ID_SERVO12",
		MAV_COMP_ID_SERVO13:                  "MAV_COMP_ID_SERVO13",
		MAV_COMP_ID_SERVO14:                  "MAV_COMP_ID_SERVO14",
		MAV_COMP_ID_GIMBAL:                   "MAV_COMP_ID_GIMBAL",
		MAV_COMP_ID_LOG:                      "MAV_COMP_ID_LOG",
		MAV_COMP_ID_ADSB:                     "MAV_COMP_ID_ADSB",
		MAV_COMP_ID_OSD:                      "MAV_COMP_ID_OSD",
		MAV_COMP_ID_PERIPHERAL:               "MAV_COMP_ID_PERIPHERAL",
		MAV_COMP_ID_QX1_GIMBAL:               "MAV_COMP_ID_QX1_GIMBAL",
		MAV_COMP_ID_FLARM:                    "MAV_COMP_ID_FLARM",
		MAV_COMP_ID_GIMBAL2:                  "MAV_COMP_ID_GIMBAL2",
		MAV_COMP_ID_GIMBAL3:                  "MAV_COMP_ID_GIMBAL3",
		MAV_COMP_ID_GIMBAL4:                  "MAV_COMP_ID_GIMBAL4",
		MAV_COMP_ID_GIMBAL5:                  "MAV_COMP_ID_GIMBAL5",
		MAV_COMP_ID_GIMBAL6:                  "MAV_COMP_ID_GIMBAL6",
		MAV_COMP_ID_MISSIONPLANNER:           "MAV_COMP_ID_MISSIONPLANNER",
		MAV_COMP_ID_ONBOARD_COMPUTER:         "MAV_COMP_ID_ONBOARD_COMPUTER",
		MAV_COMP_ID_PATHPLANNER:              "MAV_COMP_ID_PATHPLANNER",
		MAV_COMP_ID_OBSTACLE_AVOIDANCE:       "MAV_COMP_ID_OBSTACLE_AVOIDANCE",
		MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY: "MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY",
		MAV_COMP_ID_PAIRING_MANAGER:          "MAV_COMP_ID_PAIRING_MANAGER",
		MAV_COMP_ID_IMU:                      "MAV_COMP_ID_IMU",
		MAV_COMP_ID_IMU_2:                    "MAV_COMP_ID_IMU_2",
		MAV_COMP_ID_IMU_3:                    "MAV_COMP_ID_IMU_3",
		MAV_COMP_ID_GPS:                      "MAV_COMP_ID_GPS",
		MAV_COMP_ID_GPS2:                     "MAV_COMP_ID_GPS2",
		MAV_COMP_ID_ODID_TXRX_1:              "MAV_COMP_ID_ODID_TXRX_1",
		MAV_COMP_ID_ODID_TXRX_2:              "MAV_COMP_ID_ODID_TXRX_2",
		MAV_COMP_ID_ODID_TXRX_3:              "MAV_COMP_ID_ODID_TXRX_3",
		MAV_COMP_ID_UDP_BRIDGE:               "MAV_COMP_ID_UDP_BRIDGE",
		MAV_COMP_ID_UART_BRIDGE:              "MAV_COMP_ID_UART_BRIDGE",
		MAV_COMP_ID_TUNNEL_NODE:              "MAV_COMP_ID_TUNNEL_NODE",
		MAV_COMP_ID_SYSTEM_CONTROL:           "MAV_COMP_ID_SYSTEM_CONTROL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_COMPONENT_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_COMPONENT enums
func (e MAV_COMPONENT) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_COMPONENT{
		MAV_COMP_ID_ALL,
		MAV_COMP_ID_AUTOPILOT1,
		MAV_COMP_ID_USER1,
		MAV_COMP_ID_USER2,
		MAV_COMP_ID_USER3,
		MAV_COMP_ID_USER4,
		MAV_COMP_ID_USER5,
		MAV_COMP_ID_USER6,
		MAV_COMP_ID_USER7,
		MAV_COMP_ID_USER8,
		MAV_COMP_ID_USER9,
		MAV_COMP_ID_USER10,
		MAV_COMP_ID_USER11,
		MAV_COMP_ID_USER12,
		MAV_COMP_ID_USER13,
		MAV_COMP_ID_USER14,
		MAV_COMP_ID_USER15,
		MAV_COMP_ID_USER16,
		MAV_COMP_ID_USER17,
		MAV_COMP_ID_USER18,
		MAV_COMP_ID_USER19,
		MAV_COMP_ID_USER20,
		MAV_COMP_ID_USER21,
		MAV_COMP_ID_USER22,
		MAV_COMP_ID_USER23,
		MAV_COMP_ID_USER24,
		MAV_COMP_ID_USER25,
		MAV_COMP_ID_USER26,
		MAV_COMP_ID_USER27,
		MAV_COMP_ID_USER28,
		MAV_COMP_ID_USER29,
		MAV_COMP_ID_USER30,
		MAV_COMP_ID_USER31,
		MAV_COMP_ID_USER32,
		MAV_COMP_ID_USER33,
		MAV_COMP_ID_USER34,
		MAV_COMP_ID_USER35,
		MAV_COMP_ID_USER36,
		MAV_COMP_ID_USER37,
		MAV_COMP_ID_USER38,
		MAV_COMP_ID_USER39,
		MAV_COMP_ID_USER40,
		MAV_COMP_ID_USER41,
		MAV_COMP_ID_USER42,
		MAV_COMP_ID_USER43,
		MAV_COMP_ID_TELEMETRY_RADIO,
		MAV_COMP_ID_USER45,
		MAV_COMP_ID_USER46,
		MAV_COMP_ID_USER47,
		MAV_COMP_ID_USER48,
		MAV_COMP_ID_USER49,
		MAV_COMP_ID_USER50,
		MAV_COMP_ID_USER51,
		MAV_COMP_ID_USER52,
		MAV_COMP_ID_USER53,
		MAV_COMP_ID_USER54,
		MAV_COMP_ID_USER55,
		MAV_COMP_ID_USER56,
		MAV_COMP_ID_USER57,
		MAV_COMP_ID_USER58,
		MAV_COMP_ID_USER59,
		MAV_COMP_ID_USER60,
		MAV_COMP_ID_USER61,
		MAV_COMP_ID_USER62,
		MAV_COMP_ID_USER63,
		MAV_COMP_ID_USER64,
		MAV_COMP_ID_USER65,
		MAV_COMP_ID_USER66,
		MAV_COMP_ID_USER67,
		MAV_COMP_ID_USER68,
		MAV_COMP_ID_USER69,
		MAV_COMP_ID_USER70,
		MAV_COMP_ID_USER71,
		MAV_COMP_ID_USER72,
		MAV_COMP_ID_USER73,
		MAV_COMP_ID_USER74,
		MAV_COMP_ID_USER75,
		MAV_COMP_ID_CAMERA,
		MAV_COMP_ID_CAMERA2,
		MAV_COMP_ID_CAMERA3,
		MAV_COMP_ID_CAMERA4,
		MAV_COMP_ID_CAMERA5,
		MAV_COMP_ID_CAMERA6,
		MAV_COMP_ID_SERVO1,
		MAV_COMP_ID_SERVO2,
		MAV_COMP_ID_SERVO3,
		MAV_COMP_ID_SERVO4,
		MAV_COMP_ID_SERVO5,
		MAV_COMP_ID_SERVO6,
		MAV_COMP_ID_SERVO7,
		MAV_COMP_ID_SERVO8,
		MAV_COMP_ID_SERVO9,
		MAV_COMP_ID_SERVO10,
		MAV_COMP_ID_SERVO11,
		MAV_COMP_ID_SERVO12,
		MAV_COMP_ID_SERVO13,
		MAV_COMP_ID_SERVO14,
		MAV_COMP_ID_GIMBAL,
		MAV_COMP_ID_LOG,
		MAV_COMP_ID_ADSB,
		MAV_COMP_ID_OSD,
		MAV_COMP_ID_PERIPHERAL,
		MAV_COMP_ID_QX1_GIMBAL,
		MAV_COMP_ID_FLARM,
		MAV_COMP_ID_GIMBAL2,
		MAV_COMP_ID_GIMBAL3,
		MAV_COMP_ID_GIMBAL4,
		MAV_COMP_ID_GIMBAL5,
		MAV_COMP_ID_GIMBAL6,
		MAV_COMP_ID_MISSIONPLANNER,
		MAV_COMP_ID_ONBOARD_COMPUTER,
		MAV_COMP_ID_PATHPLANNER,
		MAV_COMP_ID_OBSTACLE_AVOIDANCE,
		MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY,
		MAV_COMP_ID_PAIRING_MANAGER,
		MAV_COMP_ID_IMU,
		MAV_COMP_ID_IMU_2,
		MAV_COMP_ID_IMU_3,
		MAV_COMP_ID_GPS,
		MAV_COMP_ID_GPS2,
		MAV_COMP_ID_ODID_TXRX_1,
		MAV_COMP_ID_ODID_TXRX_2,
		MAV_COMP_ID_ODID_TXRX_3,
		MAV_COMP_ID_UDP_BRIDGE,
		MAV_COMP_ID_UART_BRIDGE,
		MAV_COMP_ID_TUNNEL_NODE,
		MAV_COMP_ID_SYSTEM_CONTROL,
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
