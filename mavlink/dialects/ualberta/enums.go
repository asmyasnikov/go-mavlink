/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package ualberta

import (
	"fmt"
	"strconv"
)

// UALBERTA_AUTOPILOT_MODE type. Available autopilot modes for ualberta uav
type UALBERTA_AUTOPILOT_MODE int

func (e UALBERTA_AUTOPILOT_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MODE_MANUAL_DIRECT enum. Raw input pulse widts sent to output
	MODE_MANUAL_DIRECT UALBERTA_AUTOPILOT_MODE = 1
	// MODE_MANUAL_SCALED enum. Inputs are normalized using calibration, the converted back to raw pulse widths for output
	MODE_MANUAL_SCALED UALBERTA_AUTOPILOT_MODE = 2
	// MODE_AUTO_PID_ATT enum. dfsdfs
	MODE_AUTO_PID_ATT UALBERTA_AUTOPILOT_MODE = 3
	// MODE_AUTO_PID_VEL enum. dfsfds
	MODE_AUTO_PID_VEL UALBERTA_AUTOPILOT_MODE = 4
	// MODE_AUTO_PID_POS enum. dfsdfsdfs
	MODE_AUTO_PID_POS UALBERTA_AUTOPILOT_MODE = 5
)

func (e UALBERTA_AUTOPILOT_MODE) String() string {
	if name, ok := map[UALBERTA_AUTOPILOT_MODE]string{
		MODE_MANUAL_DIRECT: "MODE_MANUAL_DIRECT",
		MODE_MANUAL_SCALED: "MODE_MANUAL_SCALED",
		MODE_AUTO_PID_ATT:  "MODE_AUTO_PID_ATT",
		MODE_AUTO_PID_VEL:  "MODE_AUTO_PID_VEL",
		MODE_AUTO_PID_POS:  "MODE_AUTO_PID_POS",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("UALBERTA_AUTOPILOT_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects UALBERTA_AUTOPILOT_MODE enums
func (e UALBERTA_AUTOPILOT_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []UALBERTA_AUTOPILOT_MODE{
		MODE_MANUAL_DIRECT,
		MODE_MANUAL_SCALED,
		MODE_AUTO_PID_ATT,
		MODE_AUTO_PID_VEL,
		MODE_AUTO_PID_POS,
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

// UALBERTA_NAV_MODE type. Navigation filter mode
type UALBERTA_NAV_MODE int

func (e UALBERTA_NAV_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// NAV_AHRS_INIT enum
	NAV_AHRS_INIT UALBERTA_NAV_MODE = 1
	// NAV_AHRS enum. AHRS mode
	NAV_AHRS UALBERTA_NAV_MODE = 2
	// NAV_INS_GPS_INIT enum. INS/GPS initialization mode
	NAV_INS_GPS_INIT UALBERTA_NAV_MODE = 3
	// NAV_INS_GPS enum. INS/GPS mode
	NAV_INS_GPS UALBERTA_NAV_MODE = 4
)

func (e UALBERTA_NAV_MODE) String() string {
	if name, ok := map[UALBERTA_NAV_MODE]string{
		NAV_AHRS_INIT:    "NAV_AHRS_INIT",
		NAV_AHRS:         "NAV_AHRS",
		NAV_INS_GPS_INIT: "NAV_INS_GPS_INIT",
		NAV_INS_GPS:      "NAV_INS_GPS",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("UALBERTA_NAV_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects UALBERTA_NAV_MODE enums
func (e UALBERTA_NAV_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []UALBERTA_NAV_MODE{
		NAV_AHRS_INIT,
		NAV_AHRS,
		NAV_INS_GPS_INIT,
		NAV_INS_GPS,
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

// UALBERTA_PILOT_MODE type. Mode currently commanded by pilot
type UALBERTA_PILOT_MODE int

func (e UALBERTA_PILOT_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// PILOT_MANUAL enum. sdf
	PILOT_MANUAL UALBERTA_PILOT_MODE = 1
	// PILOT_AUTO enum. dfs
	PILOT_AUTO UALBERTA_PILOT_MODE = 2
	// PILOT_ROTO enum. Rotomotion mode
	PILOT_ROTO UALBERTA_PILOT_MODE = 3
)

func (e UALBERTA_PILOT_MODE) String() string {
	if name, ok := map[UALBERTA_PILOT_MODE]string{
		PILOT_MANUAL: "PILOT_MANUAL",
		PILOT_AUTO:   "PILOT_AUTO",
		PILOT_ROTO:   "PILOT_ROTO",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("UALBERTA_PILOT_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects UALBERTA_PILOT_MODE enums
func (e UALBERTA_PILOT_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []UALBERTA_PILOT_MODE{
		PILOT_MANUAL,
		PILOT_AUTO,
		PILOT_ROTO,
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

// FIRMWARE_VERSION_TYPE type. These values define the type of firmware release.  These values indicate the first version or release of this type.  For example the first alpha release would be 64, the second would be 65.
type FIRMWARE_VERSION_TYPE int

func (e FIRMWARE_VERSION_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// FIRMWARE_VERSION_TYPE_DEV enum. development release
	FIRMWARE_VERSION_TYPE_DEV FIRMWARE_VERSION_TYPE = 0
	// FIRMWARE_VERSION_TYPE_ALPHA enum. alpha release
	FIRMWARE_VERSION_TYPE_ALPHA FIRMWARE_VERSION_TYPE = 64
	// FIRMWARE_VERSION_TYPE_BETA enum. beta release
	FIRMWARE_VERSION_TYPE_BETA FIRMWARE_VERSION_TYPE = 128
	// FIRMWARE_VERSION_TYPE_RC enum. release candidate
	FIRMWARE_VERSION_TYPE_RC FIRMWARE_VERSION_TYPE = 192
	// FIRMWARE_VERSION_TYPE_OFFICIAL enum. official stable release
	FIRMWARE_VERSION_TYPE_OFFICIAL FIRMWARE_VERSION_TYPE = 255
)

func (e FIRMWARE_VERSION_TYPE) String() string {
	if name, ok := map[FIRMWARE_VERSION_TYPE]string{
		FIRMWARE_VERSION_TYPE_DEV:      "FIRMWARE_VERSION_TYPE_DEV",
		FIRMWARE_VERSION_TYPE_ALPHA:    "FIRMWARE_VERSION_TYPE_ALPHA",
		FIRMWARE_VERSION_TYPE_BETA:     "FIRMWARE_VERSION_TYPE_BETA",
		FIRMWARE_VERSION_TYPE_RC:       "FIRMWARE_VERSION_TYPE_RC",
		FIRMWARE_VERSION_TYPE_OFFICIAL: "FIRMWARE_VERSION_TYPE_OFFICIAL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("FIRMWARE_VERSION_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects FIRMWARE_VERSION_TYPE enums
func (e FIRMWARE_VERSION_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []FIRMWARE_VERSION_TYPE{
		FIRMWARE_VERSION_TYPE_DEV,
		FIRMWARE_VERSION_TYPE_ALPHA,
		FIRMWARE_VERSION_TYPE_BETA,
		FIRMWARE_VERSION_TYPE_RC,
		FIRMWARE_VERSION_TYPE_OFFICIAL,
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

// HL_FAILURE_FLAG type. Flags to report failure cases over the high latency telemtry.
type HL_FAILURE_FLAG int

func (e HL_FAILURE_FLAG) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// HL_FAILURE_FLAG_GPS enum. GPS failure
	HL_FAILURE_FLAG_GPS HL_FAILURE_FLAG = 1
	// HL_FAILURE_FLAG_DIFFERENTIAL_PRESSURE enum. Differential pressure sensor failure
	HL_FAILURE_FLAG_DIFFERENTIAL_PRESSURE HL_FAILURE_FLAG = 2
	// HL_FAILURE_FLAG_ABSOLUTE_PRESSURE enum. Absolute pressure sensor failure
	HL_FAILURE_FLAG_ABSOLUTE_PRESSURE HL_FAILURE_FLAG = 4
	// HL_FAILURE_FLAG_3D_ACCEL enum. Accelerometer sensor failure
	HL_FAILURE_FLAG_3D_ACCEL HL_FAILURE_FLAG = 8
	// HL_FAILURE_FLAG_3D_GYRO enum. Gyroscope sensor failure
	HL_FAILURE_FLAG_3D_GYRO HL_FAILURE_FLAG = 16
	// HL_FAILURE_FLAG_3D_MAG enum. Magnetometer sensor failure
	HL_FAILURE_FLAG_3D_MAG HL_FAILURE_FLAG = 32
	// HL_FAILURE_FLAG_TERRAIN enum. Terrain subsystem failure
	HL_FAILURE_FLAG_TERRAIN HL_FAILURE_FLAG = 64
	// HL_FAILURE_FLAG_BATTERY enum. Battery failure/critical low battery
	HL_FAILURE_FLAG_BATTERY HL_FAILURE_FLAG = 128
	// HL_FAILURE_FLAG_RC_RECEIVER enum. RC receiver failure/no rc connection
	HL_FAILURE_FLAG_RC_RECEIVER HL_FAILURE_FLAG = 256
	// HL_FAILURE_FLAG_OFFBOARD_LINK enum. Offboard link failure
	HL_FAILURE_FLAG_OFFBOARD_LINK HL_FAILURE_FLAG = 512
	// HL_FAILURE_FLAG_ENGINE enum. Engine failure
	HL_FAILURE_FLAG_ENGINE HL_FAILURE_FLAG = 1024
	// HL_FAILURE_FLAG_GEOFENCE enum. Geofence violation
	HL_FAILURE_FLAG_GEOFENCE HL_FAILURE_FLAG = 2048
	// HL_FAILURE_FLAG_ESTIMATOR enum. Estimator failure, for example measurement rejection or large variances
	HL_FAILURE_FLAG_ESTIMATOR HL_FAILURE_FLAG = 4096
	// HL_FAILURE_FLAG_MISSION enum. Mission failure
	HL_FAILURE_FLAG_MISSION HL_FAILURE_FLAG = 8192
)

func (e HL_FAILURE_FLAG) String() string {
	if name, ok := map[HL_FAILURE_FLAG]string{
		HL_FAILURE_FLAG_GPS:                   "HL_FAILURE_FLAG_GPS",
		HL_FAILURE_FLAG_DIFFERENTIAL_PRESSURE: "HL_FAILURE_FLAG_DIFFERENTIAL_PRESSURE",
		HL_FAILURE_FLAG_ABSOLUTE_PRESSURE:     "HL_FAILURE_FLAG_ABSOLUTE_PRESSURE",
		HL_FAILURE_FLAG_3D_ACCEL:              "HL_FAILURE_FLAG_3D_ACCEL",
		HL_FAILURE_FLAG_3D_GYRO:               "HL_FAILURE_FLAG_3D_GYRO",
		HL_FAILURE_FLAG_3D_MAG:                "HL_FAILURE_FLAG_3D_MAG",
		HL_FAILURE_FLAG_TERRAIN:               "HL_FAILURE_FLAG_TERRAIN",
		HL_FAILURE_FLAG_BATTERY:               "HL_FAILURE_FLAG_BATTERY",
		HL_FAILURE_FLAG_RC_RECEIVER:           "HL_FAILURE_FLAG_RC_RECEIVER",
		HL_FAILURE_FLAG_OFFBOARD_LINK:         "HL_FAILURE_FLAG_OFFBOARD_LINK",
		HL_FAILURE_FLAG_ENGINE:                "HL_FAILURE_FLAG_ENGINE",
		HL_FAILURE_FLAG_GEOFENCE:              "HL_FAILURE_FLAG_GEOFENCE",
		HL_FAILURE_FLAG_ESTIMATOR:             "HL_FAILURE_FLAG_ESTIMATOR",
		HL_FAILURE_FLAG_MISSION:               "HL_FAILURE_FLAG_MISSION",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("HL_FAILURE_FLAG_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects HL_FAILURE_FLAG enums
func (e HL_FAILURE_FLAG) Bitmask() string {
	bitmap := ""
	for _, entry := range []HL_FAILURE_FLAG{
		HL_FAILURE_FLAG_GPS,
		HL_FAILURE_FLAG_DIFFERENTIAL_PRESSURE,
		HL_FAILURE_FLAG_ABSOLUTE_PRESSURE,
		HL_FAILURE_FLAG_3D_ACCEL,
		HL_FAILURE_FLAG_3D_GYRO,
		HL_FAILURE_FLAG_3D_MAG,
		HL_FAILURE_FLAG_TERRAIN,
		HL_FAILURE_FLAG_BATTERY,
		HL_FAILURE_FLAG_RC_RECEIVER,
		HL_FAILURE_FLAG_OFFBOARD_LINK,
		HL_FAILURE_FLAG_ENGINE,
		HL_FAILURE_FLAG_GEOFENCE,
		HL_FAILURE_FLAG_ESTIMATOR,
		HL_FAILURE_FLAG_MISSION,
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

// MAV_GOTO type. Actions that may be specified in MAV_CMD_OVERRIDE_GOTO to override mission execution.
type MAV_GOTO int

func (e MAV_GOTO) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_GOTO_DO_HOLD enum. Hold at the current position
	MAV_GOTO_DO_HOLD MAV_GOTO = 0
	// MAV_GOTO_DO_CONTINUE enum. Continue with the next item in mission execution
	MAV_GOTO_DO_CONTINUE MAV_GOTO = 1
	// MAV_GOTO_HOLD_AT_CURRENT_POSITION enum. Hold at the current position of the system
	MAV_GOTO_HOLD_AT_CURRENT_POSITION MAV_GOTO = 2
	// MAV_GOTO_HOLD_AT_SPECIFIED_POSITION enum. Hold at the position specified in the parameters of the DO_HOLD action
	MAV_GOTO_HOLD_AT_SPECIFIED_POSITION MAV_GOTO = 3
)

func (e MAV_GOTO) String() string {
	if name, ok := map[MAV_GOTO]string{
		MAV_GOTO_DO_HOLD:                    "MAV_GOTO_DO_HOLD",
		MAV_GOTO_DO_CONTINUE:                "MAV_GOTO_DO_CONTINUE",
		MAV_GOTO_HOLD_AT_CURRENT_POSITION:   "MAV_GOTO_HOLD_AT_CURRENT_POSITION",
		MAV_GOTO_HOLD_AT_SPECIFIED_POSITION: "MAV_GOTO_HOLD_AT_SPECIFIED_POSITION",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_GOTO_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_GOTO enums
func (e MAV_GOTO) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_GOTO{
		MAV_GOTO_DO_HOLD,
		MAV_GOTO_DO_CONTINUE,
		MAV_GOTO_HOLD_AT_CURRENT_POSITION,
		MAV_GOTO_HOLD_AT_SPECIFIED_POSITION,
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

// MAV_MODE type. These defines are predefined OR-combined mode flags. There is no need to use values from this enum, but it                simplifies the use of the mode flags. Note that manual input is enabled in all modes as a safety override.
type MAV_MODE int

func (e MAV_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_MODE_PREFLIGHT enum. System is not ready to fly, booting, calibrating, etc. No flag is set
	MAV_MODE_PREFLIGHT MAV_MODE = 0
	// MAV_MODE_STABILIZE_DISARMED enum. System is allowed to be active, under assisted RC control
	MAV_MODE_STABILIZE_DISARMED MAV_MODE = 80
	// MAV_MODE_STABILIZE_ARMED enum. System is allowed to be active, under assisted RC control
	MAV_MODE_STABILIZE_ARMED MAV_MODE = 208
	// MAV_MODE_MANUAL_DISARMED enum. System is allowed to be active, under manual (RC) control, no stabilization
	MAV_MODE_MANUAL_DISARMED MAV_MODE = 64
	// MAV_MODE_MANUAL_ARMED enum. System is allowed to be active, under manual (RC) control, no stabilization
	MAV_MODE_MANUAL_ARMED MAV_MODE = 192
	// MAV_MODE_GUIDED_DISARMED enum. System is allowed to be active, under autonomous control, manual setpoint
	MAV_MODE_GUIDED_DISARMED MAV_MODE = 88
	// MAV_MODE_GUIDED_ARMED enum. System is allowed to be active, under autonomous control, manual setpoint
	MAV_MODE_GUIDED_ARMED MAV_MODE = 216
	// MAV_MODE_AUTO_DISARMED enum. System is allowed to be active, under autonomous control and navigation (the trajectory is decided onboard and not pre-programmed by waypoints)
	MAV_MODE_AUTO_DISARMED MAV_MODE = 92
	// MAV_MODE_AUTO_ARMED enum. System is allowed to be active, under autonomous control and navigation (the trajectory is decided onboard and not pre-programmed by waypoints)
	MAV_MODE_AUTO_ARMED MAV_MODE = 220
	// MAV_MODE_TEST_DISARMED enum. UNDEFINED mode. This solely depends on the autopilot - use with caution, intended for developers only
	MAV_MODE_TEST_DISARMED MAV_MODE = 66
	// MAV_MODE_TEST_ARMED enum. UNDEFINED mode. This solely depends on the autopilot - use with caution, intended for developers only
	MAV_MODE_TEST_ARMED MAV_MODE = 194
)

func (e MAV_MODE) String() string {
	if name, ok := map[MAV_MODE]string{
		MAV_MODE_PREFLIGHT:          "MAV_MODE_PREFLIGHT",
		MAV_MODE_STABILIZE_DISARMED: "MAV_MODE_STABILIZE_DISARMED",
		MAV_MODE_STABILIZE_ARMED:    "MAV_MODE_STABILIZE_ARMED",
		MAV_MODE_MANUAL_DISARMED:    "MAV_MODE_MANUAL_DISARMED",
		MAV_MODE_MANUAL_ARMED:       "MAV_MODE_MANUAL_ARMED",
		MAV_MODE_GUIDED_DISARMED:    "MAV_MODE_GUIDED_DISARMED",
		MAV_MODE_GUIDED_ARMED:       "MAV_MODE_GUIDED_ARMED",
		MAV_MODE_AUTO_DISARMED:      "MAV_MODE_AUTO_DISARMED",
		MAV_MODE_AUTO_ARMED:         "MAV_MODE_AUTO_ARMED",
		MAV_MODE_TEST_DISARMED:      "MAV_MODE_TEST_DISARMED",
		MAV_MODE_TEST_ARMED:         "MAV_MODE_TEST_ARMED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_MODE enums
func (e MAV_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_MODE{
		MAV_MODE_PREFLIGHT,
		MAV_MODE_STABILIZE_DISARMED,
		MAV_MODE_STABILIZE_ARMED,
		MAV_MODE_MANUAL_DISARMED,
		MAV_MODE_MANUAL_ARMED,
		MAV_MODE_GUIDED_DISARMED,
		MAV_MODE_GUIDED_ARMED,
		MAV_MODE_AUTO_DISARMED,
		MAV_MODE_AUTO_ARMED,
		MAV_MODE_TEST_DISARMED,
		MAV_MODE_TEST_ARMED,
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

// MAV_SYS_STATUS_SENSOR type. These encode the sensors whose status is sent as part of the SYS_STATUS message.
type MAV_SYS_STATUS_SENSOR int

func (e MAV_SYS_STATUS_SENSOR) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_SYS_STATUS_SENSOR_3D_GYRO enum. 0x01 3D gyro
	MAV_SYS_STATUS_SENSOR_3D_GYRO MAV_SYS_STATUS_SENSOR = 1
	// MAV_SYS_STATUS_SENSOR_3D_ACCEL enum. 0x02 3D accelerometer
	MAV_SYS_STATUS_SENSOR_3D_ACCEL MAV_SYS_STATUS_SENSOR = 2
	// MAV_SYS_STATUS_SENSOR_3D_MAG enum. 0x04 3D magnetometer
	MAV_SYS_STATUS_SENSOR_3D_MAG MAV_SYS_STATUS_SENSOR = 4
	// MAV_SYS_STATUS_SENSOR_ABSOLUTE_PRESSURE enum. 0x08 absolute pressure
	MAV_SYS_STATUS_SENSOR_ABSOLUTE_PRESSURE MAV_SYS_STATUS_SENSOR = 8
	// MAV_SYS_STATUS_SENSOR_DIFFERENTIAL_PRESSURE enum. 0x10 differential pressure
	MAV_SYS_STATUS_SENSOR_DIFFERENTIAL_PRESSURE MAV_SYS_STATUS_SENSOR = 16
	// MAV_SYS_STATUS_SENSOR_GPS enum. 0x20 GPS
	MAV_SYS_STATUS_SENSOR_GPS MAV_SYS_STATUS_SENSOR = 32
	// MAV_SYS_STATUS_SENSOR_OPTICAL_FLOW enum. 0x40 optical flow
	MAV_SYS_STATUS_SENSOR_OPTICAL_FLOW MAV_SYS_STATUS_SENSOR = 64
	// MAV_SYS_STATUS_SENSOR_VISION_POSITION enum. 0x80 computer vision position
	MAV_SYS_STATUS_SENSOR_VISION_POSITION MAV_SYS_STATUS_SENSOR = 128
	// MAV_SYS_STATUS_SENSOR_LASER_POSITION enum. 0x100 laser based position
	MAV_SYS_STATUS_SENSOR_LASER_POSITION MAV_SYS_STATUS_SENSOR = 256
	// MAV_SYS_STATUS_SENSOR_EXTERNAL_GROUND_TRUTH enum. 0x200 external ground truth (Vicon or Leica)
	MAV_SYS_STATUS_SENSOR_EXTERNAL_GROUND_TRUTH MAV_SYS_STATUS_SENSOR = 512
	// MAV_SYS_STATUS_SENSOR_ANGULAR_RATE_CONTROL enum. 0x400 3D angular rate control
	MAV_SYS_STATUS_SENSOR_ANGULAR_RATE_CONTROL MAV_SYS_STATUS_SENSOR = 1024
	// MAV_SYS_STATUS_SENSOR_ATTITUDE_STABILIZATION enum. 0x800 attitude stabilization
	MAV_SYS_STATUS_SENSOR_ATTITUDE_STABILIZATION MAV_SYS_STATUS_SENSOR = 2048
	// MAV_SYS_STATUS_SENSOR_YAW_POSITION enum. 0x1000 yaw position
	MAV_SYS_STATUS_SENSOR_YAW_POSITION MAV_SYS_STATUS_SENSOR = 4096
	// MAV_SYS_STATUS_SENSOR_Z_ALTITUDE_CONTROL enum. 0x2000 z/altitude control
	MAV_SYS_STATUS_SENSOR_Z_ALTITUDE_CONTROL MAV_SYS_STATUS_SENSOR = 8192
	// MAV_SYS_STATUS_SENSOR_XY_POSITION_CONTROL enum. 0x4000 x/y position control
	MAV_SYS_STATUS_SENSOR_XY_POSITION_CONTROL MAV_SYS_STATUS_SENSOR = 16384
	// MAV_SYS_STATUS_SENSOR_MOTOR_OUTPUTS enum. 0x8000 motor outputs / control
	MAV_SYS_STATUS_SENSOR_MOTOR_OUTPUTS MAV_SYS_STATUS_SENSOR = 32768
	// MAV_SYS_STATUS_SENSOR_RC_RECEIVER enum. 0x10000 rc receiver
	MAV_SYS_STATUS_SENSOR_RC_RECEIVER MAV_SYS_STATUS_SENSOR = 65536
	// MAV_SYS_STATUS_SENSOR_3D_GYRO2 enum. 0x20000 2nd 3D gyro
	MAV_SYS_STATUS_SENSOR_3D_GYRO2 MAV_SYS_STATUS_SENSOR = 131072
	// MAV_SYS_STATUS_SENSOR_3D_ACCEL2 enum. 0x40000 2nd 3D accelerometer
	MAV_SYS_STATUS_SENSOR_3D_ACCEL2 MAV_SYS_STATUS_SENSOR = 262144
	// MAV_SYS_STATUS_SENSOR_3D_MAG2 enum. 0x80000 2nd 3D magnetometer
	MAV_SYS_STATUS_SENSOR_3D_MAG2 MAV_SYS_STATUS_SENSOR = 524288
	// MAV_SYS_STATUS_GEOFENCE enum. 0x100000 geofence
	MAV_SYS_STATUS_GEOFENCE MAV_SYS_STATUS_SENSOR = 1048576
	// MAV_SYS_STATUS_AHRS enum. 0x200000 AHRS subsystem health
	MAV_SYS_STATUS_AHRS MAV_SYS_STATUS_SENSOR = 2097152
	// MAV_SYS_STATUS_TERRAIN enum. 0x400000 Terrain subsystem health
	MAV_SYS_STATUS_TERRAIN MAV_SYS_STATUS_SENSOR = 4194304
	// MAV_SYS_STATUS_REVERSE_MOTOR enum. 0x800000 Motors are reversed
	MAV_SYS_STATUS_REVERSE_MOTOR MAV_SYS_STATUS_SENSOR = 8388608
	// MAV_SYS_STATUS_LOGGING enum. 0x1000000 Logging
	MAV_SYS_STATUS_LOGGING MAV_SYS_STATUS_SENSOR = 16777216
	// MAV_SYS_STATUS_SENSOR_BATTERY enum. 0x2000000 Battery
	MAV_SYS_STATUS_SENSOR_BATTERY MAV_SYS_STATUS_SENSOR = 33554432
	// MAV_SYS_STATUS_SENSOR_PROXIMITY enum. 0x4000000 Proximity
	MAV_SYS_STATUS_SENSOR_PROXIMITY MAV_SYS_STATUS_SENSOR = 67108864
	// MAV_SYS_STATUS_SENSOR_SATCOM enum. 0x8000000 Satellite Communication
	MAV_SYS_STATUS_SENSOR_SATCOM MAV_SYS_STATUS_SENSOR = 134217728
	// MAV_SYS_STATUS_PREARM_CHECK enum. 0x10000000 pre-arm check status. Always healthy when armed
	MAV_SYS_STATUS_PREARM_CHECK MAV_SYS_STATUS_SENSOR = 268435456
	// MAV_SYS_STATUS_OBSTACLE_AVOIDANCE enum. 0x20000000 Avoidance/collision prevention
	MAV_SYS_STATUS_OBSTACLE_AVOIDANCE MAV_SYS_STATUS_SENSOR = 536870912
)

func (e MAV_SYS_STATUS_SENSOR) String() string {
	if name, ok := map[MAV_SYS_STATUS_SENSOR]string{
		MAV_SYS_STATUS_SENSOR_3D_GYRO:                "MAV_SYS_STATUS_SENSOR_3D_GYRO",
		MAV_SYS_STATUS_SENSOR_3D_ACCEL:               "MAV_SYS_STATUS_SENSOR_3D_ACCEL",
		MAV_SYS_STATUS_SENSOR_3D_MAG:                 "MAV_SYS_STATUS_SENSOR_3D_MAG",
		MAV_SYS_STATUS_SENSOR_ABSOLUTE_PRESSURE:      "MAV_SYS_STATUS_SENSOR_ABSOLUTE_PRESSURE",
		MAV_SYS_STATUS_SENSOR_DIFFERENTIAL_PRESSURE:  "MAV_SYS_STATUS_SENSOR_DIFFERENTIAL_PRESSURE",
		MAV_SYS_STATUS_SENSOR_GPS:                    "MAV_SYS_STATUS_SENSOR_GPS",
		MAV_SYS_STATUS_SENSOR_OPTICAL_FLOW:           "MAV_SYS_STATUS_SENSOR_OPTICAL_FLOW",
		MAV_SYS_STATUS_SENSOR_VISION_POSITION:        "MAV_SYS_STATUS_SENSOR_VISION_POSITION",
		MAV_SYS_STATUS_SENSOR_LASER_POSITION:         "MAV_SYS_STATUS_SENSOR_LASER_POSITION",
		MAV_SYS_STATUS_SENSOR_EXTERNAL_GROUND_TRUTH:  "MAV_SYS_STATUS_SENSOR_EXTERNAL_GROUND_TRUTH",
		MAV_SYS_STATUS_SENSOR_ANGULAR_RATE_CONTROL:   "MAV_SYS_STATUS_SENSOR_ANGULAR_RATE_CONTROL",
		MAV_SYS_STATUS_SENSOR_ATTITUDE_STABILIZATION: "MAV_SYS_STATUS_SENSOR_ATTITUDE_STABILIZATION",
		MAV_SYS_STATUS_SENSOR_YAW_POSITION:           "MAV_SYS_STATUS_SENSOR_YAW_POSITION",
		MAV_SYS_STATUS_SENSOR_Z_ALTITUDE_CONTROL:     "MAV_SYS_STATUS_SENSOR_Z_ALTITUDE_CONTROL",
		MAV_SYS_STATUS_SENSOR_XY_POSITION_CONTROL:    "MAV_SYS_STATUS_SENSOR_XY_POSITION_CONTROL",
		MAV_SYS_STATUS_SENSOR_MOTOR_OUTPUTS:          "MAV_SYS_STATUS_SENSOR_MOTOR_OUTPUTS",
		MAV_SYS_STATUS_SENSOR_RC_RECEIVER:            "MAV_SYS_STATUS_SENSOR_RC_RECEIVER",
		MAV_SYS_STATUS_SENSOR_3D_GYRO2:               "MAV_SYS_STATUS_SENSOR_3D_GYRO2",
		MAV_SYS_STATUS_SENSOR_3D_ACCEL2:              "MAV_SYS_STATUS_SENSOR_3D_ACCEL2",
		MAV_SYS_STATUS_SENSOR_3D_MAG2:                "MAV_SYS_STATUS_SENSOR_3D_MAG2",
		MAV_SYS_STATUS_GEOFENCE:                      "MAV_SYS_STATUS_GEOFENCE",
		MAV_SYS_STATUS_AHRS:                          "MAV_SYS_STATUS_AHRS",
		MAV_SYS_STATUS_TERRAIN:                       "MAV_SYS_STATUS_TERRAIN",
		MAV_SYS_STATUS_REVERSE_MOTOR:                 "MAV_SYS_STATUS_REVERSE_MOTOR",
		MAV_SYS_STATUS_LOGGING:                       "MAV_SYS_STATUS_LOGGING",
		MAV_SYS_STATUS_SENSOR_BATTERY:                "MAV_SYS_STATUS_SENSOR_BATTERY",
		MAV_SYS_STATUS_SENSOR_PROXIMITY:              "MAV_SYS_STATUS_SENSOR_PROXIMITY",
		MAV_SYS_STATUS_SENSOR_SATCOM:                 "MAV_SYS_STATUS_SENSOR_SATCOM",
		MAV_SYS_STATUS_PREARM_CHECK:                  "MAV_SYS_STATUS_PREARM_CHECK",
		MAV_SYS_STATUS_OBSTACLE_AVOIDANCE:            "MAV_SYS_STATUS_OBSTACLE_AVOIDANCE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_SYS_STATUS_SENSOR_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_SYS_STATUS_SENSOR enums
func (e MAV_SYS_STATUS_SENSOR) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_SYS_STATUS_SENSOR{
		MAV_SYS_STATUS_SENSOR_3D_GYRO,
		MAV_SYS_STATUS_SENSOR_3D_ACCEL,
		MAV_SYS_STATUS_SENSOR_3D_MAG,
		MAV_SYS_STATUS_SENSOR_ABSOLUTE_PRESSURE,
		MAV_SYS_STATUS_SENSOR_DIFFERENTIAL_PRESSURE,
		MAV_SYS_STATUS_SENSOR_GPS,
		MAV_SYS_STATUS_SENSOR_OPTICAL_FLOW,
		MAV_SYS_STATUS_SENSOR_VISION_POSITION,
		MAV_SYS_STATUS_SENSOR_LASER_POSITION,
		MAV_SYS_STATUS_SENSOR_EXTERNAL_GROUND_TRUTH,
		MAV_SYS_STATUS_SENSOR_ANGULAR_RATE_CONTROL,
		MAV_SYS_STATUS_SENSOR_ATTITUDE_STABILIZATION,
		MAV_SYS_STATUS_SENSOR_YAW_POSITION,
		MAV_SYS_STATUS_SENSOR_Z_ALTITUDE_CONTROL,
		MAV_SYS_STATUS_SENSOR_XY_POSITION_CONTROL,
		MAV_SYS_STATUS_SENSOR_MOTOR_OUTPUTS,
		MAV_SYS_STATUS_SENSOR_RC_RECEIVER,
		MAV_SYS_STATUS_SENSOR_3D_GYRO2,
		MAV_SYS_STATUS_SENSOR_3D_ACCEL2,
		MAV_SYS_STATUS_SENSOR_3D_MAG2,
		MAV_SYS_STATUS_GEOFENCE,
		MAV_SYS_STATUS_AHRS,
		MAV_SYS_STATUS_TERRAIN,
		MAV_SYS_STATUS_REVERSE_MOTOR,
		MAV_SYS_STATUS_LOGGING,
		MAV_SYS_STATUS_SENSOR_BATTERY,
		MAV_SYS_STATUS_SENSOR_PROXIMITY,
		MAV_SYS_STATUS_SENSOR_SATCOM,
		MAV_SYS_STATUS_PREARM_CHECK,
		MAV_SYS_STATUS_OBSTACLE_AVOIDANCE,
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

// MAV_FRAME type
type MAV_FRAME int

func (e MAV_FRAME) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_FRAME_GLOBAL enum. Global (WGS84) coordinate frame + MSL altitude. First value / x: latitude, second value / y: longitude, third value / z: positive altitude over mean sea level (MSL)
	MAV_FRAME_GLOBAL MAV_FRAME = 0
	// MAV_FRAME_LOCAL_NED enum. Local coordinate frame, Z-down (x: North, y: East, z: Down)
	MAV_FRAME_LOCAL_NED MAV_FRAME = 1
	// MAV_FRAME_MISSION enum. NOT a coordinate frame, indicates a mission command
	MAV_FRAME_MISSION MAV_FRAME = 2
	// MAV_FRAME_GLOBAL_RELATIVE_ALT enum. Global (WGS84) coordinate frame + altitude relative to the home position. First value / x: latitude, second value / y: longitude, third value / z: positive altitude with 0 being at the altitude of the home location
	MAV_FRAME_GLOBAL_RELATIVE_ALT MAV_FRAME = 3
	// MAV_FRAME_LOCAL_ENU enum. Local coordinate frame, Z-up (x: East, y: North, z: Up)
	MAV_FRAME_LOCAL_ENU MAV_FRAME = 4
	// MAV_FRAME_GLOBAL_INT enum. Global (WGS84) coordinate frame (scaled) + MSL altitude. First value / x: latitude in degrees*1.0e-7, second value / y: longitude in degrees*1.0e-7, third value / z: positive altitude over mean sea level (MSL)
	MAV_FRAME_GLOBAL_INT MAV_FRAME = 5
	// MAV_FRAME_GLOBAL_RELATIVE_ALT_INT enum. Global (WGS84) coordinate frame (scaled) + altitude relative to the home position. First value / x: latitude in degrees*10e-7, second value / y: longitude in degrees*10e-7, third value / z: positive altitude with 0 being at the altitude of the home location
	MAV_FRAME_GLOBAL_RELATIVE_ALT_INT MAV_FRAME = 6
	// MAV_FRAME_LOCAL_OFFSET_NED enum. Offset to the current local frame. Anything expressed in this frame should be added to the current local frame position
	MAV_FRAME_LOCAL_OFFSET_NED MAV_FRAME = 7
	// MAV_FRAME_BODY_NED enum. Setpoint in body NED frame. This makes sense if all position control is externalized - e.g. useful to command 2 m/s^2 acceleration to the right
	MAV_FRAME_BODY_NED MAV_FRAME = 8
	// MAV_FRAME_BODY_OFFSET_NED enum. Offset in body NED frame. This makes sense if adding setpoints to the current flight path, to avoid an obstacle - e.g. useful to command 2 m/s^2 acceleration to the east
	MAV_FRAME_BODY_OFFSET_NED MAV_FRAME = 9
	// MAV_FRAME_GLOBAL_TERRAIN_ALT enum. Global (WGS84) coordinate frame with AGL altitude (at the waypoint coordinate). First value / x: latitude in degrees, second value / y: longitude in degrees, third value / z: positive altitude in meters with 0 being at ground level in terrain model
	MAV_FRAME_GLOBAL_TERRAIN_ALT MAV_FRAME = 10
	// MAV_FRAME_GLOBAL_TERRAIN_ALT_INT enum. Global (WGS84) coordinate frame (scaled) with AGL altitude (at the waypoint coordinate). First value / x: latitude in degrees*10e-7, second value / y: longitude in degrees*10e-7, third value / z: positive altitude in meters with 0 being at ground level in terrain model
	MAV_FRAME_GLOBAL_TERRAIN_ALT_INT MAV_FRAME = 11
	// MAV_FRAME_BODY_FRD enum. Body fixed frame of reference, Z-down (x: Forward, y: Right, z: Down)
	MAV_FRAME_BODY_FRD MAV_FRAME = 12
	// MAV_FRAME_RESERVED_13 enum. MAV_FRAME_BODY_FLU - Body fixed frame of reference, Z-up (x: Forward, y: Left, z: Up)
	MAV_FRAME_RESERVED_13 MAV_FRAME = 13
	// MAV_FRAME_RESERVED_14 enum. MAV_FRAME_MOCAP_NED - Odometry local coordinate frame of data given by a motion capture system, Z-down (x: North, y: East, z: Down)
	MAV_FRAME_RESERVED_14 MAV_FRAME = 14
	// MAV_FRAME_RESERVED_15 enum. MAV_FRAME_MOCAP_ENU - Odometry local coordinate frame of data given by a motion capture system, Z-up (x: East, y: North, z: Up)
	MAV_FRAME_RESERVED_15 MAV_FRAME = 15
	// MAV_FRAME_RESERVED_16 enum. MAV_FRAME_VISION_NED - Odometry local coordinate frame of data given by a vision estimation system, Z-down (x: North, y: East, z: Down)
	MAV_FRAME_RESERVED_16 MAV_FRAME = 16
	// MAV_FRAME_RESERVED_17 enum. MAV_FRAME_VISION_ENU - Odometry local coordinate frame of data given by a vision estimation system, Z-up (x: East, y: North, z: Up)
	MAV_FRAME_RESERVED_17 MAV_FRAME = 17
	// MAV_FRAME_RESERVED_18 enum. MAV_FRAME_ESTIM_NED - Odometry local coordinate frame of data given by an estimator running onboard the vehicle, Z-down (x: North, y: East, z: Down)
	MAV_FRAME_RESERVED_18 MAV_FRAME = 18
	// MAV_FRAME_RESERVED_19 enum. MAV_FRAME_ESTIM_ENU - Odometry local coordinate frame of data given by an estimator running onboard the vehicle, Z-up (x: East, y: North, z: Up)
	MAV_FRAME_RESERVED_19 MAV_FRAME = 19
	// MAV_FRAME_LOCAL_FRD enum. Forward, Right, Down coordinate frame. This is a local frame with Z-down and arbitrary F/R alignment (i.e. not aligned with NED/earth frame)
	MAV_FRAME_LOCAL_FRD MAV_FRAME = 20
	// MAV_FRAME_LOCAL_FLU enum. Forward, Left, Up coordinate frame. This is a local frame with Z-up and arbitrary F/L alignment (i.e. not aligned with ENU/earth frame)
	MAV_FRAME_LOCAL_FLU MAV_FRAME = 21
)

func (e MAV_FRAME) String() string {
	if name, ok := map[MAV_FRAME]string{
		MAV_FRAME_GLOBAL:                  "MAV_FRAME_GLOBAL",
		MAV_FRAME_LOCAL_NED:               "MAV_FRAME_LOCAL_NED",
		MAV_FRAME_MISSION:                 "MAV_FRAME_MISSION",
		MAV_FRAME_GLOBAL_RELATIVE_ALT:     "MAV_FRAME_GLOBAL_RELATIVE_ALT",
		MAV_FRAME_LOCAL_ENU:               "MAV_FRAME_LOCAL_ENU",
		MAV_FRAME_GLOBAL_INT:              "MAV_FRAME_GLOBAL_INT",
		MAV_FRAME_GLOBAL_RELATIVE_ALT_INT: "MAV_FRAME_GLOBAL_RELATIVE_ALT_INT",
		MAV_FRAME_LOCAL_OFFSET_NED:        "MAV_FRAME_LOCAL_OFFSET_NED",
		MAV_FRAME_BODY_NED:                "MAV_FRAME_BODY_NED",
		MAV_FRAME_BODY_OFFSET_NED:         "MAV_FRAME_BODY_OFFSET_NED",
		MAV_FRAME_GLOBAL_TERRAIN_ALT:      "MAV_FRAME_GLOBAL_TERRAIN_ALT",
		MAV_FRAME_GLOBAL_TERRAIN_ALT_INT:  "MAV_FRAME_GLOBAL_TERRAIN_ALT_INT",
		MAV_FRAME_BODY_FRD:                "MAV_FRAME_BODY_FRD",
		MAV_FRAME_RESERVED_13:             "MAV_FRAME_RESERVED_13",
		MAV_FRAME_RESERVED_14:             "MAV_FRAME_RESERVED_14",
		MAV_FRAME_RESERVED_15:             "MAV_FRAME_RESERVED_15",
		MAV_FRAME_RESERVED_16:             "MAV_FRAME_RESERVED_16",
		MAV_FRAME_RESERVED_17:             "MAV_FRAME_RESERVED_17",
		MAV_FRAME_RESERVED_18:             "MAV_FRAME_RESERVED_18",
		MAV_FRAME_RESERVED_19:             "MAV_FRAME_RESERVED_19",
		MAV_FRAME_LOCAL_FRD:               "MAV_FRAME_LOCAL_FRD",
		MAV_FRAME_LOCAL_FLU:               "MAV_FRAME_LOCAL_FLU",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_FRAME_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_FRAME enums
func (e MAV_FRAME) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_FRAME{
		MAV_FRAME_GLOBAL,
		MAV_FRAME_LOCAL_NED,
		MAV_FRAME_MISSION,
		MAV_FRAME_GLOBAL_RELATIVE_ALT,
		MAV_FRAME_LOCAL_ENU,
		MAV_FRAME_GLOBAL_INT,
		MAV_FRAME_GLOBAL_RELATIVE_ALT_INT,
		MAV_FRAME_LOCAL_OFFSET_NED,
		MAV_FRAME_BODY_NED,
		MAV_FRAME_BODY_OFFSET_NED,
		MAV_FRAME_GLOBAL_TERRAIN_ALT,
		MAV_FRAME_GLOBAL_TERRAIN_ALT_INT,
		MAV_FRAME_BODY_FRD,
		MAV_FRAME_RESERVED_13,
		MAV_FRAME_RESERVED_14,
		MAV_FRAME_RESERVED_15,
		MAV_FRAME_RESERVED_16,
		MAV_FRAME_RESERVED_17,
		MAV_FRAME_RESERVED_18,
		MAV_FRAME_RESERVED_19,
		MAV_FRAME_LOCAL_FRD,
		MAV_FRAME_LOCAL_FLU,
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

// MAVLINK_DATA_STREAM_TYPE type
type MAVLINK_DATA_STREAM_TYPE int

func (e MAVLINK_DATA_STREAM_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAVLINK_DATA_STREAM_IMG_JPEG enum
	MAVLINK_DATA_STREAM_IMG_JPEG MAVLINK_DATA_STREAM_TYPE = 0
	// MAVLINK_DATA_STREAM_IMG_BMP enum
	MAVLINK_DATA_STREAM_IMG_BMP MAVLINK_DATA_STREAM_TYPE = 1
	// MAVLINK_DATA_STREAM_IMG_RAW8U enum
	MAVLINK_DATA_STREAM_IMG_RAW8U MAVLINK_DATA_STREAM_TYPE = 2
	// MAVLINK_DATA_STREAM_IMG_RAW32U enum
	MAVLINK_DATA_STREAM_IMG_RAW32U MAVLINK_DATA_STREAM_TYPE = 3
	// MAVLINK_DATA_STREAM_IMG_PGM enum
	MAVLINK_DATA_STREAM_IMG_PGM MAVLINK_DATA_STREAM_TYPE = 4
	// MAVLINK_DATA_STREAM_IMG_PNG enum
	MAVLINK_DATA_STREAM_IMG_PNG MAVLINK_DATA_STREAM_TYPE = 5
)

func (e MAVLINK_DATA_STREAM_TYPE) String() string {
	if name, ok := map[MAVLINK_DATA_STREAM_TYPE]string{
		MAVLINK_DATA_STREAM_IMG_JPEG:   "MAVLINK_DATA_STREAM_IMG_JPEG",
		MAVLINK_DATA_STREAM_IMG_BMP:    "MAVLINK_DATA_STREAM_IMG_BMP",
		MAVLINK_DATA_STREAM_IMG_RAW8U:  "MAVLINK_DATA_STREAM_IMG_RAW8U",
		MAVLINK_DATA_STREAM_IMG_RAW32U: "MAVLINK_DATA_STREAM_IMG_RAW32U",
		MAVLINK_DATA_STREAM_IMG_PGM:    "MAVLINK_DATA_STREAM_IMG_PGM",
		MAVLINK_DATA_STREAM_IMG_PNG:    "MAVLINK_DATA_STREAM_IMG_PNG",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAVLINK_DATA_STREAM_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAVLINK_DATA_STREAM_TYPE enums
func (e MAVLINK_DATA_STREAM_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAVLINK_DATA_STREAM_TYPE{
		MAVLINK_DATA_STREAM_IMG_JPEG,
		MAVLINK_DATA_STREAM_IMG_BMP,
		MAVLINK_DATA_STREAM_IMG_RAW8U,
		MAVLINK_DATA_STREAM_IMG_RAW32U,
		MAVLINK_DATA_STREAM_IMG_PGM,
		MAVLINK_DATA_STREAM_IMG_PNG,
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

// FENCE_ACTION type
type FENCE_ACTION int

func (e FENCE_ACTION) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// FENCE_ACTION_NONE enum. Disable fenced mode
	FENCE_ACTION_NONE FENCE_ACTION = 0
	// FENCE_ACTION_GUIDED enum. Switched to guided mode to return point (fence point 0)
	FENCE_ACTION_GUIDED FENCE_ACTION = 1
	// FENCE_ACTION_REPORT enum. Report fence breach, but don't take action
	FENCE_ACTION_REPORT FENCE_ACTION = 2
	// FENCE_ACTION_GUIDED_THR_PASS enum. Switched to guided mode to return point (fence point 0) with manual throttle control
	FENCE_ACTION_GUIDED_THR_PASS FENCE_ACTION = 3
	// FENCE_ACTION_RTL enum. Switch to RTL (return to launch) mode and head for the return point
	FENCE_ACTION_RTL FENCE_ACTION = 4
)

func (e FENCE_ACTION) String() string {
	if name, ok := map[FENCE_ACTION]string{
		FENCE_ACTION_NONE:            "FENCE_ACTION_NONE",
		FENCE_ACTION_GUIDED:          "FENCE_ACTION_GUIDED",
		FENCE_ACTION_REPORT:          "FENCE_ACTION_REPORT",
		FENCE_ACTION_GUIDED_THR_PASS: "FENCE_ACTION_GUIDED_THR_PASS",
		FENCE_ACTION_RTL:             "FENCE_ACTION_RTL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("FENCE_ACTION_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects FENCE_ACTION enums
func (e FENCE_ACTION) Bitmask() string {
	bitmap := ""
	for _, entry := range []FENCE_ACTION{
		FENCE_ACTION_NONE,
		FENCE_ACTION_GUIDED,
		FENCE_ACTION_REPORT,
		FENCE_ACTION_GUIDED_THR_PASS,
		FENCE_ACTION_RTL,
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

// FENCE_BREACH type
type FENCE_BREACH int

func (e FENCE_BREACH) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// FENCE_BREACH_NONE enum. No last fence breach
	FENCE_BREACH_NONE FENCE_BREACH = 0
	// FENCE_BREACH_MINALT enum. Breached minimum altitude
	FENCE_BREACH_MINALT FENCE_BREACH = 1
	// FENCE_BREACH_MAXALT enum. Breached maximum altitude
	FENCE_BREACH_MAXALT FENCE_BREACH = 2
	// FENCE_BREACH_BOUNDARY enum. Breached fence boundary
	FENCE_BREACH_BOUNDARY FENCE_BREACH = 3
)

func (e FENCE_BREACH) String() string {
	if name, ok := map[FENCE_BREACH]string{
		FENCE_BREACH_NONE:     "FENCE_BREACH_NONE",
		FENCE_BREACH_MINALT:   "FENCE_BREACH_MINALT",
		FENCE_BREACH_MAXALT:   "FENCE_BREACH_MAXALT",
		FENCE_BREACH_BOUNDARY: "FENCE_BREACH_BOUNDARY",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("FENCE_BREACH_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects FENCE_BREACH enums
func (e FENCE_BREACH) Bitmask() string {
	bitmap := ""
	for _, entry := range []FENCE_BREACH{
		FENCE_BREACH_NONE,
		FENCE_BREACH_MINALT,
		FENCE_BREACH_MAXALT,
		FENCE_BREACH_BOUNDARY,
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

// FENCE_MITIGATE type. Actions being taken to mitigate/prevent fence breach
type FENCE_MITIGATE int

func (e FENCE_MITIGATE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// FENCE_MITIGATE_UNKNOWN enum. Unknown
	FENCE_MITIGATE_UNKNOWN FENCE_MITIGATE = 0
	// FENCE_MITIGATE_NONE enum. No actions being taken
	FENCE_MITIGATE_NONE FENCE_MITIGATE = 1
	// FENCE_MITIGATE_VEL_LIMIT enum. Velocity limiting active to prevent breach
	FENCE_MITIGATE_VEL_LIMIT FENCE_MITIGATE = 2
)

func (e FENCE_MITIGATE) String() string {
	if name, ok := map[FENCE_MITIGATE]string{
		FENCE_MITIGATE_UNKNOWN:   "FENCE_MITIGATE_UNKNOWN",
		FENCE_MITIGATE_NONE:      "FENCE_MITIGATE_NONE",
		FENCE_MITIGATE_VEL_LIMIT: "FENCE_MITIGATE_VEL_LIMIT",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("FENCE_MITIGATE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects FENCE_MITIGATE enums
func (e FENCE_MITIGATE) Bitmask() string {
	bitmap := ""
	for _, entry := range []FENCE_MITIGATE{
		FENCE_MITIGATE_UNKNOWN,
		FENCE_MITIGATE_NONE,
		FENCE_MITIGATE_VEL_LIMIT,
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

// MAV_MOUNT_MODE type. Enumeration of possible mount operation modes. This message is used by obsolete/deprecated gimbal messages.
type MAV_MOUNT_MODE int

func (e MAV_MOUNT_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_MOUNT_MODE_RETRACT enum. Load and keep safe position (Roll,Pitch,Yaw) from permant memory and stop stabilization
	MAV_MOUNT_MODE_RETRACT MAV_MOUNT_MODE = 0
	// MAV_MOUNT_MODE_NEUTRAL enum. Load and keep neutral position (Roll,Pitch,Yaw) from permanent memory
	MAV_MOUNT_MODE_NEUTRAL MAV_MOUNT_MODE = 1
	// MAV_MOUNT_MODE_MAVLINK_TARGETING enum. Load neutral position and start MAVLink Roll,Pitch,Yaw control with stabilization
	MAV_MOUNT_MODE_MAVLINK_TARGETING MAV_MOUNT_MODE = 2
	// MAV_MOUNT_MODE_RC_TARGETING enum. Load neutral position and start RC Roll,Pitch,Yaw control with stabilization
	MAV_MOUNT_MODE_RC_TARGETING MAV_MOUNT_MODE = 3
	// MAV_MOUNT_MODE_GPS_POINT enum. Load neutral position and start to point to Lat,Lon,Alt
	MAV_MOUNT_MODE_GPS_POINT MAV_MOUNT_MODE = 4
	// MAV_MOUNT_MODE_SYSID_TARGET enum. Gimbal tracks system with specified system ID
	MAV_MOUNT_MODE_SYSID_TARGET MAV_MOUNT_MODE = 5
	// MAV_MOUNT_MODE_HOME_LOCATION enum. Gimbal tracks home location
	MAV_MOUNT_MODE_HOME_LOCATION MAV_MOUNT_MODE = 6
)

func (e MAV_MOUNT_MODE) String() string {
	if name, ok := map[MAV_MOUNT_MODE]string{
		MAV_MOUNT_MODE_RETRACT:           "MAV_MOUNT_MODE_RETRACT",
		MAV_MOUNT_MODE_NEUTRAL:           "MAV_MOUNT_MODE_NEUTRAL",
		MAV_MOUNT_MODE_MAVLINK_TARGETING: "MAV_MOUNT_MODE_MAVLINK_TARGETING",
		MAV_MOUNT_MODE_RC_TARGETING:      "MAV_MOUNT_MODE_RC_TARGETING",
		MAV_MOUNT_MODE_GPS_POINT:         "MAV_MOUNT_MODE_GPS_POINT",
		MAV_MOUNT_MODE_SYSID_TARGET:      "MAV_MOUNT_MODE_SYSID_TARGET",
		MAV_MOUNT_MODE_HOME_LOCATION:     "MAV_MOUNT_MODE_HOME_LOCATION",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_MOUNT_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_MOUNT_MODE enums
func (e MAV_MOUNT_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_MOUNT_MODE{
		MAV_MOUNT_MODE_RETRACT,
		MAV_MOUNT_MODE_NEUTRAL,
		MAV_MOUNT_MODE_MAVLINK_TARGETING,
		MAV_MOUNT_MODE_RC_TARGETING,
		MAV_MOUNT_MODE_GPS_POINT,
		MAV_MOUNT_MODE_SYSID_TARGET,
		MAV_MOUNT_MODE_HOME_LOCATION,
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

// GIMBAL_DEVICE_CAP_FLAGS type. Gimbal device (low level) capability flags (bitmap)
type GIMBAL_DEVICE_CAP_FLAGS int

func (e GIMBAL_DEVICE_CAP_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT enum. Gimbal device supports a retracted position
	GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT GIMBAL_DEVICE_CAP_FLAGS = 1
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL enum. Gimbal device supports a horizontal, forward looking position, stabilized
	GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL GIMBAL_DEVICE_CAP_FLAGS = 2
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS enum. Gimbal device supports rotating around roll axis
	GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS GIMBAL_DEVICE_CAP_FLAGS = 4
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW enum. Gimbal device supports to follow a roll angle relative to the vehicle
	GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW GIMBAL_DEVICE_CAP_FLAGS = 8
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK enum. Gimbal device supports locking to an roll angle (generally that's the default with roll stabilized)
	GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK GIMBAL_DEVICE_CAP_FLAGS = 16
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS enum. Gimbal device supports rotating around pitch axis
	GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS GIMBAL_DEVICE_CAP_FLAGS = 32
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW enum. Gimbal device supports to follow a pitch angle relative to the vehicle
	GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW GIMBAL_DEVICE_CAP_FLAGS = 64
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK enum. Gimbal device supports locking to an pitch angle (generally that's the default with pitch stabilized)
	GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK GIMBAL_DEVICE_CAP_FLAGS = 128
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS enum. Gimbal device supports rotating around yaw axis
	GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS GIMBAL_DEVICE_CAP_FLAGS = 256
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW enum. Gimbal device supports to follow a yaw angle relative to the vehicle (generally that's the default)
	GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW GIMBAL_DEVICE_CAP_FLAGS = 512
	// GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK enum. Gimbal device supports locking to an absolute heading (often this is an option available)
	GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK GIMBAL_DEVICE_CAP_FLAGS = 1024
	// GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW enum. Gimbal device supports yawing/panning infinetely (e.g. using slip disk)
	GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW GIMBAL_DEVICE_CAP_FLAGS = 2048
)

func (e GIMBAL_DEVICE_CAP_FLAGS) String() string {
	if name, ok := map[GIMBAL_DEVICE_CAP_FLAGS]string{
		GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT:           "GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT",
		GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL:           "GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL",
		GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS:         "GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS",
		GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW:       "GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW",
		GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK:         "GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK",
		GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS:        "GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS",
		GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW:      "GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW",
		GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK:        "GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK",
		GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS:          "GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS",
		GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW:        "GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW",
		GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK:          "GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK",
		GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW: "GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("GIMBAL_DEVICE_CAP_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects GIMBAL_DEVICE_CAP_FLAGS enums
func (e GIMBAL_DEVICE_CAP_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []GIMBAL_DEVICE_CAP_FLAGS{
		GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT,
		GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL,
		GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS,
		GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW,
		GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK,
		GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS,
		GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW,
		GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK,
		GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS,
		GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW,
		GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK,
		GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW,
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

// GIMBAL_MANAGER_CAP_FLAGS type. Gimbal manager high level capability flags (bitmap). The first 16 bits are identical to the GIMBAL_DEVICE_CAP_FLAGS which are identical with GIMBAL_DEVICE_FLAGS. However, the gimbal manager does not need to copy the flags from the gimbal but can also enhance the capabilities and thus add flags.
type GIMBAL_MANAGER_CAP_FLAGS int

func (e GIMBAL_MANAGER_CAP_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT
	GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT GIMBAL_MANAGER_CAP_FLAGS = 1
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL
	GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL GIMBAL_MANAGER_CAP_FLAGS = 2
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS GIMBAL_MANAGER_CAP_FLAGS = 4
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = 8
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK GIMBAL_MANAGER_CAP_FLAGS = 16
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS GIMBAL_MANAGER_CAP_FLAGS = 32
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = 64
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK GIMBAL_MANAGER_CAP_FLAGS = 128
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS GIMBAL_MANAGER_CAP_FLAGS = 256
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = 512
	// GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK enum. Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK GIMBAL_MANAGER_CAP_FLAGS = 1024
	// GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW enum. Based on GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW GIMBAL_MANAGER_CAP_FLAGS = 2048
	// GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL enum. Gimbal manager supports to point to a local position
	GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL GIMBAL_MANAGER_CAP_FLAGS = 65536
	// GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL enum. Gimbal manager supports to point to a global latitude, longitude, altitude position
	GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL GIMBAL_MANAGER_CAP_FLAGS = 131072
)

func (e GIMBAL_MANAGER_CAP_FLAGS) String() string {
	if name, ok := map[GIMBAL_MANAGER_CAP_FLAGS]string{
		GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT:               "GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT",
		GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL:               "GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL",
		GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS:             "GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS",
		GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW:           "GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW",
		GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK:             "GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK",
		GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS:            "GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS",
		GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW:          "GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW",
		GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK:            "GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK",
		GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS:              "GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS",
		GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW:            "GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW",
		GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK:              "GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK",
		GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW:     "GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW",
		GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL:  "GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL",
		GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL: "GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("GIMBAL_MANAGER_CAP_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects GIMBAL_MANAGER_CAP_FLAGS enums
func (e GIMBAL_MANAGER_CAP_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []GIMBAL_MANAGER_CAP_FLAGS{
		GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT,
		GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL,
		GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS,
		GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW,
		GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK,
		GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS,
		GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW,
		GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK,
		GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS,
		GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW,
		GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK,
		GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW,
		GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL,
		GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL,
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

// GIMBAL_DEVICE_FLAGS type. Flags for gimbal device (lower level) operation.
type GIMBAL_DEVICE_FLAGS int

func (e GIMBAL_DEVICE_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// GIMBAL_DEVICE_FLAGS_RETRACT enum. Set to retracted safe position (no stabilization), takes presedence over all other flags
	GIMBAL_DEVICE_FLAGS_RETRACT GIMBAL_DEVICE_FLAGS = 1
	// GIMBAL_DEVICE_FLAGS_NEUTRAL enum. Set to neutral position (horizontal, forward looking, with stabiliziation), takes presedence over all other flags except RETRACT
	GIMBAL_DEVICE_FLAGS_NEUTRAL GIMBAL_DEVICE_FLAGS = 2
	// GIMBAL_DEVICE_FLAGS_ROLL_LOCK enum. Lock roll angle to absolute angle relative to horizon (not relative to drone). This is generally the default with a stabilizing gimbal
	GIMBAL_DEVICE_FLAGS_ROLL_LOCK GIMBAL_DEVICE_FLAGS = 4
	// GIMBAL_DEVICE_FLAGS_PITCH_LOCK enum. Lock pitch angle to absolute angle relative to horizon (not relative to drone). This is generally the default
	GIMBAL_DEVICE_FLAGS_PITCH_LOCK GIMBAL_DEVICE_FLAGS = 8
	// GIMBAL_DEVICE_FLAGS_YAW_LOCK enum. Lock yaw angle to absolute angle relative to North (not relative to drone). If this flag is set, the quaternion is in the Earth frame with the x-axis pointing North (yaw absolute). If this flag is not set, the quaternion frame is in the Earth frame rotated so that the x-axis is pointing forward (yaw relative to vehicle)
	GIMBAL_DEVICE_FLAGS_YAW_LOCK GIMBAL_DEVICE_FLAGS = 16
)

func (e GIMBAL_DEVICE_FLAGS) String() string {
	if name, ok := map[GIMBAL_DEVICE_FLAGS]string{
		GIMBAL_DEVICE_FLAGS_RETRACT:    "GIMBAL_DEVICE_FLAGS_RETRACT",
		GIMBAL_DEVICE_FLAGS_NEUTRAL:    "GIMBAL_DEVICE_FLAGS_NEUTRAL",
		GIMBAL_DEVICE_FLAGS_ROLL_LOCK:  "GIMBAL_DEVICE_FLAGS_ROLL_LOCK",
		GIMBAL_DEVICE_FLAGS_PITCH_LOCK: "GIMBAL_DEVICE_FLAGS_PITCH_LOCK",
		GIMBAL_DEVICE_FLAGS_YAW_LOCK:   "GIMBAL_DEVICE_FLAGS_YAW_LOCK",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("GIMBAL_DEVICE_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects GIMBAL_DEVICE_FLAGS enums
func (e GIMBAL_DEVICE_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []GIMBAL_DEVICE_FLAGS{
		GIMBAL_DEVICE_FLAGS_RETRACT,
		GIMBAL_DEVICE_FLAGS_NEUTRAL,
		GIMBAL_DEVICE_FLAGS_ROLL_LOCK,
		GIMBAL_DEVICE_FLAGS_PITCH_LOCK,
		GIMBAL_DEVICE_FLAGS_YAW_LOCK,
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

// GIMBAL_MANAGER_FLAGS type. Flags for high level gimbal manager operation The first 16 bytes are identical to the GIMBAL_DEVICE_FLAGS.
type GIMBAL_MANAGER_FLAGS int

func (e GIMBAL_MANAGER_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// GIMBAL_MANAGER_FLAGS_RETRACT enum. Based on GIMBAL_DEVICE_FLAGS_RETRACT
	GIMBAL_MANAGER_FLAGS_RETRACT GIMBAL_MANAGER_FLAGS = 1
	// GIMBAL_MANAGER_FLAGS_NEUTRAL enum. Based on GIMBAL_DEVICE_FLAGS_NEUTRAL
	GIMBAL_MANAGER_FLAGS_NEUTRAL GIMBAL_MANAGER_FLAGS = 2
	// GIMBAL_MANAGER_FLAGS_ROLL_LOCK enum. Based on GIMBAL_DEVICE_FLAGS_ROLL_LOCK
	GIMBAL_MANAGER_FLAGS_ROLL_LOCK GIMBAL_MANAGER_FLAGS = 4
	// GIMBAL_MANAGER_FLAGS_PITCH_LOCK enum. Based on GIMBAL_DEVICE_FLAGS_PITCH_LOCK
	GIMBAL_MANAGER_FLAGS_PITCH_LOCK GIMBAL_MANAGER_FLAGS = 8
	// GIMBAL_MANAGER_FLAGS_YAW_LOCK enum. Based on GIMBAL_DEVICE_FLAGS_YAW_LOCK
	GIMBAL_MANAGER_FLAGS_YAW_LOCK GIMBAL_MANAGER_FLAGS = 16
)

func (e GIMBAL_MANAGER_FLAGS) String() string {
	if name, ok := map[GIMBAL_MANAGER_FLAGS]string{
		GIMBAL_MANAGER_FLAGS_RETRACT:    "GIMBAL_MANAGER_FLAGS_RETRACT",
		GIMBAL_MANAGER_FLAGS_NEUTRAL:    "GIMBAL_MANAGER_FLAGS_NEUTRAL",
		GIMBAL_MANAGER_FLAGS_ROLL_LOCK:  "GIMBAL_MANAGER_FLAGS_ROLL_LOCK",
		GIMBAL_MANAGER_FLAGS_PITCH_LOCK: "GIMBAL_MANAGER_FLAGS_PITCH_LOCK",
		GIMBAL_MANAGER_FLAGS_YAW_LOCK:   "GIMBAL_MANAGER_FLAGS_YAW_LOCK",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("GIMBAL_MANAGER_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects GIMBAL_MANAGER_FLAGS enums
func (e GIMBAL_MANAGER_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []GIMBAL_MANAGER_FLAGS{
		GIMBAL_MANAGER_FLAGS_RETRACT,
		GIMBAL_MANAGER_FLAGS_NEUTRAL,
		GIMBAL_MANAGER_FLAGS_ROLL_LOCK,
		GIMBAL_MANAGER_FLAGS_PITCH_LOCK,
		GIMBAL_MANAGER_FLAGS_YAW_LOCK,
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

// GIMBAL_DEVICE_ERROR_FLAGS type. Gimbal device (low level) error flags (bitmap, 0 means no error)
type GIMBAL_DEVICE_ERROR_FLAGS int

func (e GIMBAL_DEVICE_ERROR_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// GIMBAL_DEVICE_ERROR_FLAGS_AT_ROLL_LIMIT enum. Gimbal device is limited by hardware roll limit
	GIMBAL_DEVICE_ERROR_FLAGS_AT_ROLL_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 1
	// GIMBAL_DEVICE_ERROR_FLAGS_AT_PITCH_LIMIT enum. Gimbal device is limited by hardware pitch limit
	GIMBAL_DEVICE_ERROR_FLAGS_AT_PITCH_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 2
	// GIMBAL_DEVICE_ERROR_FLAGS_AT_YAW_LIMIT enum. Gimbal device is limited by hardware yaw limit
	GIMBAL_DEVICE_ERROR_FLAGS_AT_YAW_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 4
	// GIMBAL_DEVICE_ERROR_FLAGS_ENCODER_ERROR enum. There is an error with the gimbal encoders
	GIMBAL_DEVICE_ERROR_FLAGS_ENCODER_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 8
	// GIMBAL_DEVICE_ERROR_FLAGS_POWER_ERROR enum. There is an error with the gimbal power source
	GIMBAL_DEVICE_ERROR_FLAGS_POWER_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 16
	// GIMBAL_DEVICE_ERROR_FLAGS_MOTOR_ERROR enum. There is an error with the gimbal motor's
	GIMBAL_DEVICE_ERROR_FLAGS_MOTOR_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 32
	// GIMBAL_DEVICE_ERROR_FLAGS_SOFTWARE_ERROR enum. There is an error with the gimbal's software
	GIMBAL_DEVICE_ERROR_FLAGS_SOFTWARE_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 64
	// GIMBAL_DEVICE_ERROR_FLAGS_COMMS_ERROR enum. There is an error with the gimbal's communication
	GIMBAL_DEVICE_ERROR_FLAGS_COMMS_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 128
	// GIMBAL_DEVICE_ERROR_FLAGS_CALIBRATION_RUNNING enum. Gimbal is currently calibrating
	GIMBAL_DEVICE_ERROR_FLAGS_CALIBRATION_RUNNING GIMBAL_DEVICE_ERROR_FLAGS = 256
)

func (e GIMBAL_DEVICE_ERROR_FLAGS) String() string {
	if name, ok := map[GIMBAL_DEVICE_ERROR_FLAGS]string{
		GIMBAL_DEVICE_ERROR_FLAGS_AT_ROLL_LIMIT:       "GIMBAL_DEVICE_ERROR_FLAGS_AT_ROLL_LIMIT",
		GIMBAL_DEVICE_ERROR_FLAGS_AT_PITCH_LIMIT:      "GIMBAL_DEVICE_ERROR_FLAGS_AT_PITCH_LIMIT",
		GIMBAL_DEVICE_ERROR_FLAGS_AT_YAW_LIMIT:        "GIMBAL_DEVICE_ERROR_FLAGS_AT_YAW_LIMIT",
		GIMBAL_DEVICE_ERROR_FLAGS_ENCODER_ERROR:       "GIMBAL_DEVICE_ERROR_FLAGS_ENCODER_ERROR",
		GIMBAL_DEVICE_ERROR_FLAGS_POWER_ERROR:         "GIMBAL_DEVICE_ERROR_FLAGS_POWER_ERROR",
		GIMBAL_DEVICE_ERROR_FLAGS_MOTOR_ERROR:         "GIMBAL_DEVICE_ERROR_FLAGS_MOTOR_ERROR",
		GIMBAL_DEVICE_ERROR_FLAGS_SOFTWARE_ERROR:      "GIMBAL_DEVICE_ERROR_FLAGS_SOFTWARE_ERROR",
		GIMBAL_DEVICE_ERROR_FLAGS_COMMS_ERROR:         "GIMBAL_DEVICE_ERROR_FLAGS_COMMS_ERROR",
		GIMBAL_DEVICE_ERROR_FLAGS_CALIBRATION_RUNNING: "GIMBAL_DEVICE_ERROR_FLAGS_CALIBRATION_RUNNING",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("GIMBAL_DEVICE_ERROR_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects GIMBAL_DEVICE_ERROR_FLAGS enums
func (e GIMBAL_DEVICE_ERROR_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []GIMBAL_DEVICE_ERROR_FLAGS{
		GIMBAL_DEVICE_ERROR_FLAGS_AT_ROLL_LIMIT,
		GIMBAL_DEVICE_ERROR_FLAGS_AT_PITCH_LIMIT,
		GIMBAL_DEVICE_ERROR_FLAGS_AT_YAW_LIMIT,
		GIMBAL_DEVICE_ERROR_FLAGS_ENCODER_ERROR,
		GIMBAL_DEVICE_ERROR_FLAGS_POWER_ERROR,
		GIMBAL_DEVICE_ERROR_FLAGS_MOTOR_ERROR,
		GIMBAL_DEVICE_ERROR_FLAGS_SOFTWARE_ERROR,
		GIMBAL_DEVICE_ERROR_FLAGS_COMMS_ERROR,
		GIMBAL_DEVICE_ERROR_FLAGS_CALIBRATION_RUNNING,
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

// GRIPPER_ACTIONS type. Gripper actions.
type GRIPPER_ACTIONS int

func (e GRIPPER_ACTIONS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// GRIPPER_ACTION_RELEASE enum. Gripper release cargo
	GRIPPER_ACTION_RELEASE GRIPPER_ACTIONS = 0
	// GRIPPER_ACTION_GRAB enum. Gripper grab onto cargo
	GRIPPER_ACTION_GRAB GRIPPER_ACTIONS = 1
)

func (e GRIPPER_ACTIONS) String() string {
	if name, ok := map[GRIPPER_ACTIONS]string{
		GRIPPER_ACTION_RELEASE: "GRIPPER_ACTION_RELEASE",
		GRIPPER_ACTION_GRAB:    "GRIPPER_ACTION_GRAB",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("GRIPPER_ACTIONS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects GRIPPER_ACTIONS enums
func (e GRIPPER_ACTIONS) Bitmask() string {
	bitmap := ""
	for _, entry := range []GRIPPER_ACTIONS{
		GRIPPER_ACTION_RELEASE,
		GRIPPER_ACTION_GRAB,
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

// WINCH_ACTIONS type. Winch actions.
type WINCH_ACTIONS int

func (e WINCH_ACTIONS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// WINCH_RELAXED enum. Relax winch
	WINCH_RELAXED WINCH_ACTIONS = 0
	// WINCH_RELATIVE_LENGTH_CONTROL enum. Wind or unwind specified length of cable, optionally using specified rate
	WINCH_RELATIVE_LENGTH_CONTROL WINCH_ACTIONS = 1
	// WINCH_RATE_CONTROL enum. Wind or unwind cable at specified rate
	WINCH_RATE_CONTROL WINCH_ACTIONS = 2
)

func (e WINCH_ACTIONS) String() string {
	if name, ok := map[WINCH_ACTIONS]string{
		WINCH_RELAXED:                 "WINCH_RELAXED",
		WINCH_RELATIVE_LENGTH_CONTROL: "WINCH_RELATIVE_LENGTH_CONTROL",
		WINCH_RATE_CONTROL:            "WINCH_RATE_CONTROL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("WINCH_ACTIONS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects WINCH_ACTIONS enums
func (e WINCH_ACTIONS) Bitmask() string {
	bitmap := ""
	for _, entry := range []WINCH_ACTIONS{
		WINCH_RELAXED,
		WINCH_RELATIVE_LENGTH_CONTROL,
		WINCH_RATE_CONTROL,
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

// UAVCAN_NODE_HEALTH type. Generalized UAVCAN node health
type UAVCAN_NODE_HEALTH int

func (e UAVCAN_NODE_HEALTH) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// UAVCAN_NODE_HEALTH_OK enum. The node is functioning properly
	UAVCAN_NODE_HEALTH_OK UAVCAN_NODE_HEALTH = 0
	// UAVCAN_NODE_HEALTH_WARNING enum. A critical parameter went out of range or the node has encountered a minor failure
	UAVCAN_NODE_HEALTH_WARNING UAVCAN_NODE_HEALTH = 1
	// UAVCAN_NODE_HEALTH_ERROR enum. The node has encountered a major failure
	UAVCAN_NODE_HEALTH_ERROR UAVCAN_NODE_HEALTH = 2
	// UAVCAN_NODE_HEALTH_CRITICAL enum. The node has suffered a fatal malfunction
	UAVCAN_NODE_HEALTH_CRITICAL UAVCAN_NODE_HEALTH = 3
)

func (e UAVCAN_NODE_HEALTH) String() string {
	if name, ok := map[UAVCAN_NODE_HEALTH]string{
		UAVCAN_NODE_HEALTH_OK:       "UAVCAN_NODE_HEALTH_OK",
		UAVCAN_NODE_HEALTH_WARNING:  "UAVCAN_NODE_HEALTH_WARNING",
		UAVCAN_NODE_HEALTH_ERROR:    "UAVCAN_NODE_HEALTH_ERROR",
		UAVCAN_NODE_HEALTH_CRITICAL: "UAVCAN_NODE_HEALTH_CRITICAL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("UAVCAN_NODE_HEALTH_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects UAVCAN_NODE_HEALTH enums
func (e UAVCAN_NODE_HEALTH) Bitmask() string {
	bitmap := ""
	for _, entry := range []UAVCAN_NODE_HEALTH{
		UAVCAN_NODE_HEALTH_OK,
		UAVCAN_NODE_HEALTH_WARNING,
		UAVCAN_NODE_HEALTH_ERROR,
		UAVCAN_NODE_HEALTH_CRITICAL,
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

// UAVCAN_NODE_MODE type. Generalized UAVCAN node mode
type UAVCAN_NODE_MODE int

func (e UAVCAN_NODE_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// UAVCAN_NODE_MODE_OPERATIONAL enum. The node is performing its primary functions
	UAVCAN_NODE_MODE_OPERATIONAL UAVCAN_NODE_MODE = 0
	// UAVCAN_NODE_MODE_INITIALIZATION enum. The node is initializing; this mode is entered immediately after startup
	UAVCAN_NODE_MODE_INITIALIZATION UAVCAN_NODE_MODE = 1
	// UAVCAN_NODE_MODE_MAINTENANCE enum. The node is under maintenance
	UAVCAN_NODE_MODE_MAINTENANCE UAVCAN_NODE_MODE = 2
	// UAVCAN_NODE_MODE_SOFTWARE_UPDATE enum. The node is in the process of updating its software
	UAVCAN_NODE_MODE_SOFTWARE_UPDATE UAVCAN_NODE_MODE = 3
	// UAVCAN_NODE_MODE_OFFLINE enum. The node is no longer available online
	UAVCAN_NODE_MODE_OFFLINE UAVCAN_NODE_MODE = 7
)

func (e UAVCAN_NODE_MODE) String() string {
	if name, ok := map[UAVCAN_NODE_MODE]string{
		UAVCAN_NODE_MODE_OPERATIONAL:     "UAVCAN_NODE_MODE_OPERATIONAL",
		UAVCAN_NODE_MODE_INITIALIZATION:  "UAVCAN_NODE_MODE_INITIALIZATION",
		UAVCAN_NODE_MODE_MAINTENANCE:     "UAVCAN_NODE_MODE_MAINTENANCE",
		UAVCAN_NODE_MODE_SOFTWARE_UPDATE: "UAVCAN_NODE_MODE_SOFTWARE_UPDATE",
		UAVCAN_NODE_MODE_OFFLINE:         "UAVCAN_NODE_MODE_OFFLINE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("UAVCAN_NODE_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects UAVCAN_NODE_MODE enums
func (e UAVCAN_NODE_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []UAVCAN_NODE_MODE{
		UAVCAN_NODE_MODE_OPERATIONAL,
		UAVCAN_NODE_MODE_INITIALIZATION,
		UAVCAN_NODE_MODE_MAINTENANCE,
		UAVCAN_NODE_MODE_SOFTWARE_UPDATE,
		UAVCAN_NODE_MODE_OFFLINE,
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

// ESC_CONNECTION_TYPE type. Indicates the ESC connection type.
type ESC_CONNECTION_TYPE int

func (e ESC_CONNECTION_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ESC_CONNECTION_TYPE_PPM enum. Traditional PPM ESC
	ESC_CONNECTION_TYPE_PPM ESC_CONNECTION_TYPE = 0
	// ESC_CONNECTION_TYPE_SERIAL enum. Serial Bus connected ESC
	ESC_CONNECTION_TYPE_SERIAL ESC_CONNECTION_TYPE = 1
	// ESC_CONNECTION_TYPE_ONESHOT enum. One Shot PPM ESC
	ESC_CONNECTION_TYPE_ONESHOT ESC_CONNECTION_TYPE = 2
	// ESC_CONNECTION_TYPE_I2C enum. I2C ESC
	ESC_CONNECTION_TYPE_I2C ESC_CONNECTION_TYPE = 3
	// ESC_CONNECTION_TYPE_CAN enum. CAN-Bus ESC
	ESC_CONNECTION_TYPE_CAN ESC_CONNECTION_TYPE = 4
	// ESC_CONNECTION_TYPE_DSHOT enum. DShot ESC
	ESC_CONNECTION_TYPE_DSHOT ESC_CONNECTION_TYPE = 5
)

func (e ESC_CONNECTION_TYPE) String() string {
	if name, ok := map[ESC_CONNECTION_TYPE]string{
		ESC_CONNECTION_TYPE_PPM:     "ESC_CONNECTION_TYPE_PPM",
		ESC_CONNECTION_TYPE_SERIAL:  "ESC_CONNECTION_TYPE_SERIAL",
		ESC_CONNECTION_TYPE_ONESHOT: "ESC_CONNECTION_TYPE_ONESHOT",
		ESC_CONNECTION_TYPE_I2C:     "ESC_CONNECTION_TYPE_I2C",
		ESC_CONNECTION_TYPE_CAN:     "ESC_CONNECTION_TYPE_CAN",
		ESC_CONNECTION_TYPE_DSHOT:   "ESC_CONNECTION_TYPE_DSHOT",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("ESC_CONNECTION_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects ESC_CONNECTION_TYPE enums
func (e ESC_CONNECTION_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []ESC_CONNECTION_TYPE{
		ESC_CONNECTION_TYPE_PPM,
		ESC_CONNECTION_TYPE_SERIAL,
		ESC_CONNECTION_TYPE_ONESHOT,
		ESC_CONNECTION_TYPE_I2C,
		ESC_CONNECTION_TYPE_CAN,
		ESC_CONNECTION_TYPE_DSHOT,
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

// ESC_FAILURE_FLAGS type. Flags to report ESC failures.
type ESC_FAILURE_FLAGS int

func (e ESC_FAILURE_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ESC_FAILURE_NONE enum. No ESC failure
	ESC_FAILURE_NONE ESC_FAILURE_FLAGS = 0
	// ESC_FAILURE_OVER_CURRENT enum. Over current failure
	ESC_FAILURE_OVER_CURRENT ESC_FAILURE_FLAGS = 1
	// ESC_FAILURE_OVER_VOLTAGE enum. Over voltage failure
	ESC_FAILURE_OVER_VOLTAGE ESC_FAILURE_FLAGS = 2
	// ESC_FAILURE_OVER_TEMPERATURE enum. Over temperature failure
	ESC_FAILURE_OVER_TEMPERATURE ESC_FAILURE_FLAGS = 4
	// ESC_FAILURE_OVER_RPM enum. Over RPM failure
	ESC_FAILURE_OVER_RPM ESC_FAILURE_FLAGS = 8
	// ESC_FAILURE_INCONSISTENT_CMD enum. Inconsistent command failure i.e. out of bounds
	ESC_FAILURE_INCONSISTENT_CMD ESC_FAILURE_FLAGS = 16
	// ESC_FAILURE_MOTOR_STUCK enum. Motor stuck failure
	ESC_FAILURE_MOTOR_STUCK ESC_FAILURE_FLAGS = 32
	// ESC_FAILURE_GENERIC enum. Generic ESC failure
	ESC_FAILURE_GENERIC ESC_FAILURE_FLAGS = 64
)

func (e ESC_FAILURE_FLAGS) String() string {
	if name, ok := map[ESC_FAILURE_FLAGS]string{
		ESC_FAILURE_NONE:             "ESC_FAILURE_NONE",
		ESC_FAILURE_OVER_CURRENT:     "ESC_FAILURE_OVER_CURRENT",
		ESC_FAILURE_OVER_VOLTAGE:     "ESC_FAILURE_OVER_VOLTAGE",
		ESC_FAILURE_OVER_TEMPERATURE: "ESC_FAILURE_OVER_TEMPERATURE",
		ESC_FAILURE_OVER_RPM:         "ESC_FAILURE_OVER_RPM",
		ESC_FAILURE_INCONSISTENT_CMD: "ESC_FAILURE_INCONSISTENT_CMD",
		ESC_FAILURE_MOTOR_STUCK:      "ESC_FAILURE_MOTOR_STUCK",
		ESC_FAILURE_GENERIC:          "ESC_FAILURE_GENERIC",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("ESC_FAILURE_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects ESC_FAILURE_FLAGS enums
func (e ESC_FAILURE_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []ESC_FAILURE_FLAGS{
		ESC_FAILURE_NONE,
		ESC_FAILURE_OVER_CURRENT,
		ESC_FAILURE_OVER_VOLTAGE,
		ESC_FAILURE_OVER_TEMPERATURE,
		ESC_FAILURE_OVER_RPM,
		ESC_FAILURE_INCONSISTENT_CMD,
		ESC_FAILURE_MOTOR_STUCK,
		ESC_FAILURE_GENERIC,
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

// STORAGE_STATUS type. Flags to indicate the status of camera storage.
type STORAGE_STATUS int

func (e STORAGE_STATUS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// STORAGE_STATUS_EMPTY enum. Storage is missing (no microSD card loaded for example.)
	STORAGE_STATUS_EMPTY STORAGE_STATUS = 0
	// STORAGE_STATUS_UNFORMATTED enum. Storage present but unformatted
	STORAGE_STATUS_UNFORMATTED STORAGE_STATUS = 1
	// STORAGE_STATUS_READY enum. Storage present and ready
	STORAGE_STATUS_READY STORAGE_STATUS = 2
	// STORAGE_STATUS_NOT_SUPPORTED enum. Camera does not supply storage status information. Capacity information in STORAGE_INFORMATION fields will be ignored
	STORAGE_STATUS_NOT_SUPPORTED STORAGE_STATUS = 3
)

func (e STORAGE_STATUS) String() string {
	if name, ok := map[STORAGE_STATUS]string{
		STORAGE_STATUS_EMPTY:         "STORAGE_STATUS_EMPTY",
		STORAGE_STATUS_UNFORMATTED:   "STORAGE_STATUS_UNFORMATTED",
		STORAGE_STATUS_READY:         "STORAGE_STATUS_READY",
		STORAGE_STATUS_NOT_SUPPORTED: "STORAGE_STATUS_NOT_SUPPORTED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("STORAGE_STATUS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects STORAGE_STATUS enums
func (e STORAGE_STATUS) Bitmask() string {
	bitmap := ""
	for _, entry := range []STORAGE_STATUS{
		STORAGE_STATUS_EMPTY,
		STORAGE_STATUS_UNFORMATTED,
		STORAGE_STATUS_READY,
		STORAGE_STATUS_NOT_SUPPORTED,
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

// STORAGE_TYPE type. Flags to indicate the type of storage.
type STORAGE_TYPE int

func (e STORAGE_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// STORAGE_TYPE_UNKNOWN enum. Storage type is not known
	STORAGE_TYPE_UNKNOWN STORAGE_TYPE = 0
	// STORAGE_TYPE_USB_STICK enum. Storage type is USB device
	STORAGE_TYPE_USB_STICK STORAGE_TYPE = 1
	// STORAGE_TYPE_SD enum. Storage type is SD card
	STORAGE_TYPE_SD STORAGE_TYPE = 2
	// STORAGE_TYPE_MICROSD enum. Storage type is microSD card
	STORAGE_TYPE_MICROSD STORAGE_TYPE = 3
	// STORAGE_TYPE_CF enum. Storage type is CFast
	STORAGE_TYPE_CF STORAGE_TYPE = 4
	// STORAGE_TYPE_CFE enum. Storage type is CFexpress
	STORAGE_TYPE_CFE STORAGE_TYPE = 5
	// STORAGE_TYPE_XQD enum. Storage type is XQD
	STORAGE_TYPE_XQD STORAGE_TYPE = 6
	// STORAGE_TYPE_HD enum. Storage type is HD mass storage type
	STORAGE_TYPE_HD STORAGE_TYPE = 7
	// STORAGE_TYPE_OTHER enum. Storage type is other, not listed type
	STORAGE_TYPE_OTHER STORAGE_TYPE = 254
)

func (e STORAGE_TYPE) String() string {
	if name, ok := map[STORAGE_TYPE]string{
		STORAGE_TYPE_UNKNOWN:   "STORAGE_TYPE_UNKNOWN",
		STORAGE_TYPE_USB_STICK: "STORAGE_TYPE_USB_STICK",
		STORAGE_TYPE_SD:        "STORAGE_TYPE_SD",
		STORAGE_TYPE_MICROSD:   "STORAGE_TYPE_MICROSD",
		STORAGE_TYPE_CF:        "STORAGE_TYPE_CF",
		STORAGE_TYPE_CFE:       "STORAGE_TYPE_CFE",
		STORAGE_TYPE_XQD:       "STORAGE_TYPE_XQD",
		STORAGE_TYPE_HD:        "STORAGE_TYPE_HD",
		STORAGE_TYPE_OTHER:     "STORAGE_TYPE_OTHER",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("STORAGE_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects STORAGE_TYPE enums
func (e STORAGE_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []STORAGE_TYPE{
		STORAGE_TYPE_UNKNOWN,
		STORAGE_TYPE_USB_STICK,
		STORAGE_TYPE_SD,
		STORAGE_TYPE_MICROSD,
		STORAGE_TYPE_CF,
		STORAGE_TYPE_CFE,
		STORAGE_TYPE_XQD,
		STORAGE_TYPE_HD,
		STORAGE_TYPE_OTHER,
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

// ORBIT_YAW_BEHAVIOUR type. Yaw behaviour during orbit flight.
type ORBIT_YAW_BEHAVIOUR int

func (e ORBIT_YAW_BEHAVIOUR) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TO_CIRCLE_CENTER enum. Vehicle front points to the center (default)
	ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TO_CIRCLE_CENTER ORBIT_YAW_BEHAVIOUR = 0
	// ORBIT_YAW_BEHAVIOUR_HOLD_INITIAL_HEADING enum. Vehicle front holds heading when message received
	ORBIT_YAW_BEHAVIOUR_HOLD_INITIAL_HEADING ORBIT_YAW_BEHAVIOUR = 1
	// ORBIT_YAW_BEHAVIOUR_UNCONTROLLED enum. Yaw uncontrolled
	ORBIT_YAW_BEHAVIOUR_UNCONTROLLED ORBIT_YAW_BEHAVIOUR = 2
	// ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TANGENT_TO_CIRCLE enum. Vehicle front follows flight path (tangential to circle)
	ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TANGENT_TO_CIRCLE ORBIT_YAW_BEHAVIOUR = 3
	// ORBIT_YAW_BEHAVIOUR_RC_CONTROLLED enum. Yaw controlled by RC input
	ORBIT_YAW_BEHAVIOUR_RC_CONTROLLED ORBIT_YAW_BEHAVIOUR = 4
)

func (e ORBIT_YAW_BEHAVIOUR) String() string {
	if name, ok := map[ORBIT_YAW_BEHAVIOUR]string{
		ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TO_CIRCLE_CENTER:  "ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TO_CIRCLE_CENTER",
		ORBIT_YAW_BEHAVIOUR_HOLD_INITIAL_HEADING:         "ORBIT_YAW_BEHAVIOUR_HOLD_INITIAL_HEADING",
		ORBIT_YAW_BEHAVIOUR_UNCONTROLLED:                 "ORBIT_YAW_BEHAVIOUR_UNCONTROLLED",
		ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TANGENT_TO_CIRCLE: "ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TANGENT_TO_CIRCLE",
		ORBIT_YAW_BEHAVIOUR_RC_CONTROLLED:                "ORBIT_YAW_BEHAVIOUR_RC_CONTROLLED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("ORBIT_YAW_BEHAVIOUR_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects ORBIT_YAW_BEHAVIOUR enums
func (e ORBIT_YAW_BEHAVIOUR) Bitmask() string {
	bitmap := ""
	for _, entry := range []ORBIT_YAW_BEHAVIOUR{
		ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TO_CIRCLE_CENTER,
		ORBIT_YAW_BEHAVIOUR_HOLD_INITIAL_HEADING,
		ORBIT_YAW_BEHAVIOUR_UNCONTROLLED,
		ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TANGENT_TO_CIRCLE,
		ORBIT_YAW_BEHAVIOUR_RC_CONTROLLED,
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

// WIFI_CONFIG_AP_RESPONSE type. Possible responses from a WIFI_CONFIG_AP message.
type WIFI_CONFIG_AP_RESPONSE int

func (e WIFI_CONFIG_AP_RESPONSE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// WIFI_CONFIG_AP_RESPONSE_UNDEFINED enum. Undefined response. Likely an indicative of a system that doesn't support this request
	WIFI_CONFIG_AP_RESPONSE_UNDEFINED WIFI_CONFIG_AP_RESPONSE = 0
	// WIFI_CONFIG_AP_RESPONSE_ACCEPTED enum. Changes accepted
	WIFI_CONFIG_AP_RESPONSE_ACCEPTED WIFI_CONFIG_AP_RESPONSE = 1
	// WIFI_CONFIG_AP_RESPONSE_REJECTED enum. Changes rejected
	WIFI_CONFIG_AP_RESPONSE_REJECTED WIFI_CONFIG_AP_RESPONSE = 2
	// WIFI_CONFIG_AP_RESPONSE_MODE_ERROR enum. Invalid Mode
	WIFI_CONFIG_AP_RESPONSE_MODE_ERROR WIFI_CONFIG_AP_RESPONSE = 3
	// WIFI_CONFIG_AP_RESPONSE_SSID_ERROR enum. Invalid SSID
	WIFI_CONFIG_AP_RESPONSE_SSID_ERROR WIFI_CONFIG_AP_RESPONSE = 4
	// WIFI_CONFIG_AP_RESPONSE_PASSWORD_ERROR enum. Invalid Password
	WIFI_CONFIG_AP_RESPONSE_PASSWORD_ERROR WIFI_CONFIG_AP_RESPONSE = 5
)

func (e WIFI_CONFIG_AP_RESPONSE) String() string {
	if name, ok := map[WIFI_CONFIG_AP_RESPONSE]string{
		WIFI_CONFIG_AP_RESPONSE_UNDEFINED:      "WIFI_CONFIG_AP_RESPONSE_UNDEFINED",
		WIFI_CONFIG_AP_RESPONSE_ACCEPTED:       "WIFI_CONFIG_AP_RESPONSE_ACCEPTED",
		WIFI_CONFIG_AP_RESPONSE_REJECTED:       "WIFI_CONFIG_AP_RESPONSE_REJECTED",
		WIFI_CONFIG_AP_RESPONSE_MODE_ERROR:     "WIFI_CONFIG_AP_RESPONSE_MODE_ERROR",
		WIFI_CONFIG_AP_RESPONSE_SSID_ERROR:     "WIFI_CONFIG_AP_RESPONSE_SSID_ERROR",
		WIFI_CONFIG_AP_RESPONSE_PASSWORD_ERROR: "WIFI_CONFIG_AP_RESPONSE_PASSWORD_ERROR",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("WIFI_CONFIG_AP_RESPONSE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects WIFI_CONFIG_AP_RESPONSE enums
func (e WIFI_CONFIG_AP_RESPONSE) Bitmask() string {
	bitmap := ""
	for _, entry := range []WIFI_CONFIG_AP_RESPONSE{
		WIFI_CONFIG_AP_RESPONSE_UNDEFINED,
		WIFI_CONFIG_AP_RESPONSE_ACCEPTED,
		WIFI_CONFIG_AP_RESPONSE_REJECTED,
		WIFI_CONFIG_AP_RESPONSE_MODE_ERROR,
		WIFI_CONFIG_AP_RESPONSE_SSID_ERROR,
		WIFI_CONFIG_AP_RESPONSE_PASSWORD_ERROR,
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

// CELLULAR_CONFIG_RESPONSE type. Possible responses from a CELLULAR_CONFIG message.
type CELLULAR_CONFIG_RESPONSE int

func (e CELLULAR_CONFIG_RESPONSE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// CELLULAR_CONFIG_RESPONSE_ACCEPTED enum. Changes accepted
	CELLULAR_CONFIG_RESPONSE_ACCEPTED CELLULAR_CONFIG_RESPONSE = 0
	// CELLULAR_CONFIG_RESPONSE_APN_ERROR enum. Invalid APN
	CELLULAR_CONFIG_RESPONSE_APN_ERROR CELLULAR_CONFIG_RESPONSE = 1
	// CELLULAR_CONFIG_RESPONSE_PIN_ERROR enum. Invalid PIN
	CELLULAR_CONFIG_RESPONSE_PIN_ERROR CELLULAR_CONFIG_RESPONSE = 2
	// CELLULAR_CONFIG_RESPONSE_REJECTED enum. Changes rejected
	CELLULAR_CONFIG_RESPONSE_REJECTED CELLULAR_CONFIG_RESPONSE = 3
	// CELLULAR_CONFIG_BLOCKED_PUK_REQUIRED enum. PUK is required to unblock SIM card
	CELLULAR_CONFIG_BLOCKED_PUK_REQUIRED CELLULAR_CONFIG_RESPONSE = 4
)

func (e CELLULAR_CONFIG_RESPONSE) String() string {
	if name, ok := map[CELLULAR_CONFIG_RESPONSE]string{
		CELLULAR_CONFIG_RESPONSE_ACCEPTED:    "CELLULAR_CONFIG_RESPONSE_ACCEPTED",
		CELLULAR_CONFIG_RESPONSE_APN_ERROR:   "CELLULAR_CONFIG_RESPONSE_APN_ERROR",
		CELLULAR_CONFIG_RESPONSE_PIN_ERROR:   "CELLULAR_CONFIG_RESPONSE_PIN_ERROR",
		CELLULAR_CONFIG_RESPONSE_REJECTED:    "CELLULAR_CONFIG_RESPONSE_REJECTED",
		CELLULAR_CONFIG_BLOCKED_PUK_REQUIRED: "CELLULAR_CONFIG_BLOCKED_PUK_REQUIRED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("CELLULAR_CONFIG_RESPONSE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects CELLULAR_CONFIG_RESPONSE enums
func (e CELLULAR_CONFIG_RESPONSE) Bitmask() string {
	bitmap := ""
	for _, entry := range []CELLULAR_CONFIG_RESPONSE{
		CELLULAR_CONFIG_RESPONSE_ACCEPTED,
		CELLULAR_CONFIG_RESPONSE_APN_ERROR,
		CELLULAR_CONFIG_RESPONSE_PIN_ERROR,
		CELLULAR_CONFIG_RESPONSE_REJECTED,
		CELLULAR_CONFIG_BLOCKED_PUK_REQUIRED,
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

// WIFI_CONFIG_AP_MODE type. WiFi Mode.
type WIFI_CONFIG_AP_MODE int

func (e WIFI_CONFIG_AP_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// WIFI_CONFIG_AP_MODE_UNDEFINED enum. WiFi mode is undefined
	WIFI_CONFIG_AP_MODE_UNDEFINED WIFI_CONFIG_AP_MODE = 0
	// WIFI_CONFIG_AP_MODE_AP enum. WiFi configured as an access point
	WIFI_CONFIG_AP_MODE_AP WIFI_CONFIG_AP_MODE = 1
	// WIFI_CONFIG_AP_MODE_STATION enum. WiFi configured as a station connected to an existing local WiFi network
	WIFI_CONFIG_AP_MODE_STATION WIFI_CONFIG_AP_MODE = 2
	// WIFI_CONFIG_AP_MODE_DISABLED enum. WiFi disabled
	WIFI_CONFIG_AP_MODE_DISABLED WIFI_CONFIG_AP_MODE = 3
)

func (e WIFI_CONFIG_AP_MODE) String() string {
	if name, ok := map[WIFI_CONFIG_AP_MODE]string{
		WIFI_CONFIG_AP_MODE_UNDEFINED: "WIFI_CONFIG_AP_MODE_UNDEFINED",
		WIFI_CONFIG_AP_MODE_AP:        "WIFI_CONFIG_AP_MODE_AP",
		WIFI_CONFIG_AP_MODE_STATION:   "WIFI_CONFIG_AP_MODE_STATION",
		WIFI_CONFIG_AP_MODE_DISABLED:  "WIFI_CONFIG_AP_MODE_DISABLED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("WIFI_CONFIG_AP_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects WIFI_CONFIG_AP_MODE enums
func (e WIFI_CONFIG_AP_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []WIFI_CONFIG_AP_MODE{
		WIFI_CONFIG_AP_MODE_UNDEFINED,
		WIFI_CONFIG_AP_MODE_AP,
		WIFI_CONFIG_AP_MODE_STATION,
		WIFI_CONFIG_AP_MODE_DISABLED,
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

// COMP_METADATA_TYPE type. Possible values for COMPONENT_INFORMATION.comp_metadata_type.
type COMP_METADATA_TYPE int

func (e COMP_METADATA_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// COMP_METADATA_TYPE_VERSION enum. Version information which also includes information on other optional supported COMP_METADATA_TYPE's. Must be supported. Only downloadable from vehicle
	COMP_METADATA_TYPE_VERSION COMP_METADATA_TYPE = 0
	// COMP_METADATA_TYPE_PARAMETER enum. Parameter meta data
	COMP_METADATA_TYPE_PARAMETER COMP_METADATA_TYPE = 1
	// COMP_METADATA_TYPE_COMMANDS enum. Meta data which specifies the commands the vehicle supports. (WIP)
	COMP_METADATA_TYPE_COMMANDS COMP_METADATA_TYPE = 2
)

func (e COMP_METADATA_TYPE) String() string {
	if name, ok := map[COMP_METADATA_TYPE]string{
		COMP_METADATA_TYPE_VERSION:   "COMP_METADATA_TYPE_VERSION",
		COMP_METADATA_TYPE_PARAMETER: "COMP_METADATA_TYPE_PARAMETER",
		COMP_METADATA_TYPE_COMMANDS:  "COMP_METADATA_TYPE_COMMANDS",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("COMP_METADATA_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects COMP_METADATA_TYPE enums
func (e COMP_METADATA_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []COMP_METADATA_TYPE{
		COMP_METADATA_TYPE_VERSION,
		COMP_METADATA_TYPE_PARAMETER,
		COMP_METADATA_TYPE_COMMANDS,
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

// PARAM_TRANSACTION_TRANSPORT type. Possible transport layers to set and get parameters via mavlink during a parameter transaction.
type PARAM_TRANSACTION_TRANSPORT int

func (e PARAM_TRANSACTION_TRANSPORT) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// PARAM_TRANSACTION_TRANSPORT_PARAM enum. Transaction over param transport
	PARAM_TRANSACTION_TRANSPORT_PARAM PARAM_TRANSACTION_TRANSPORT = 0
	// PARAM_TRANSACTION_TRANSPORT_PARAM_EXT enum. Transaction over param_ext transport
	PARAM_TRANSACTION_TRANSPORT_PARAM_EXT PARAM_TRANSACTION_TRANSPORT = 1
)

func (e PARAM_TRANSACTION_TRANSPORT) String() string {
	if name, ok := map[PARAM_TRANSACTION_TRANSPORT]string{
		PARAM_TRANSACTION_TRANSPORT_PARAM:     "PARAM_TRANSACTION_TRANSPORT_PARAM",
		PARAM_TRANSACTION_TRANSPORT_PARAM_EXT: "PARAM_TRANSACTION_TRANSPORT_PARAM_EXT",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("PARAM_TRANSACTION_TRANSPORT_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects PARAM_TRANSACTION_TRANSPORT enums
func (e PARAM_TRANSACTION_TRANSPORT) Bitmask() string {
	bitmap := ""
	for _, entry := range []PARAM_TRANSACTION_TRANSPORT{
		PARAM_TRANSACTION_TRANSPORT_PARAM,
		PARAM_TRANSACTION_TRANSPORT_PARAM_EXT,
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

// PARAM_TRANSACTION_ACTION type. Possible parameter transaction actions.
type PARAM_TRANSACTION_ACTION int

func (e PARAM_TRANSACTION_ACTION) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// PARAM_TRANSACTION_ACTION_START enum. Commit the current parameter transaction
	PARAM_TRANSACTION_ACTION_START PARAM_TRANSACTION_ACTION = 0
	// PARAM_TRANSACTION_ACTION_COMMIT enum. Commit the current parameter transaction
	PARAM_TRANSACTION_ACTION_COMMIT PARAM_TRANSACTION_ACTION = 1
	// PARAM_TRANSACTION_ACTION_CANCEL enum. Cancel the current parameter transaction
	PARAM_TRANSACTION_ACTION_CANCEL PARAM_TRANSACTION_ACTION = 2
)

func (e PARAM_TRANSACTION_ACTION) String() string {
	if name, ok := map[PARAM_TRANSACTION_ACTION]string{
		PARAM_TRANSACTION_ACTION_START:  "PARAM_TRANSACTION_ACTION_START",
		PARAM_TRANSACTION_ACTION_COMMIT: "PARAM_TRANSACTION_ACTION_COMMIT",
		PARAM_TRANSACTION_ACTION_CANCEL: "PARAM_TRANSACTION_ACTION_CANCEL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("PARAM_TRANSACTION_ACTION_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects PARAM_TRANSACTION_ACTION enums
func (e PARAM_TRANSACTION_ACTION) Bitmask() string {
	bitmap := ""
	for _, entry := range []PARAM_TRANSACTION_ACTION{
		PARAM_TRANSACTION_ACTION_START,
		PARAM_TRANSACTION_ACTION_COMMIT,
		PARAM_TRANSACTION_ACTION_CANCEL,
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

// MAV_CMD type. Commands to be executed by the MAV. They can be executed on user request, or as part of a mission script. If the action is used in a mission, the parameter mapping to the waypoint/mission message is as follows: Param 1, Param 2, Param 3, Param 4, X: Param 5, Y:Param 6, Z:Param 7. This command list is similar what ARINC 424 is for commercial aircraft: A data format how to interpret waypoint/mission data. NaN and INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current yaw or latitude rather than a specific value). See https://mavlink.io/en/guide/xml_schema.html#MAV_CMD for information about the structure of the MAV_CMD entries
type MAV_CMD int

func (e MAV_CMD) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_CMD_NAV_WAYPOINT enum. Navigate to waypoint. Params: 1) Hold time. (ignored by fixed wing, time to stay at waypoint for rotary wing); 2) Acceptance radius (if the sphere with this radius is hit, the waypoint counts as reached); 3) 0 to pass through the WP, if &gt; 0 radius to pass by WP. Positive value for clockwise orbit, negative value for counter-clockwise orbit. Allows trajectory control.; 4) Desired yaw angle at waypoint (rotary wing). NaN to use the current system yaw heading mode (e.g. yaw towards next waypoint, yaw to home, etc.).; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_WAYPOINT MAV_CMD = 16
	// MAV_CMD_NAV_LOITER_UNLIM enum. Loiter around this waypoint an unlimited amount of time. Params: 1) Empty; 2) Empty; 3) Loiter radius around waypoint for forward-only moving vehicles (not multicopters). If positive loiter clockwise, else counter-clockwise; 4) Desired yaw angle. NaN to use the current system yaw heading mode (e.g. yaw towards next waypoint, yaw to home, etc.).; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_LOITER_UNLIM MAV_CMD = 17
	// MAV_CMD_NAV_LOITER_TURNS enum. Loiter around this waypoint for X turns. Params: 1) Number of turns.; 2) Leave loiter circle only once heading towards the next waypoint (0 = False); 3) Loiter radius around waypoint for forward-only moving vehicles (not multicopters). If positive loiter clockwise, else counter-clockwise; 4) Loiter circle exit location and/or path to next waypoint ("xtrack") for forward-only moving vehicles (not multicopters). 0 for the vehicle to converge towards the center xtrack when it leaves the loiter (the line between the centers of the current and next waypoint), 1 to converge to the direct line between the location that the vehicle exits the loiter radius and the next waypoint. Otherwise the angle (in degrees) between the tangent of the loiter circle and the center xtrack at which the vehicle must leave the loiter (and converge to the center xtrack). NaN to use the current system default xtrack behaviour.; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_LOITER_TURNS MAV_CMD = 18
	// MAV_CMD_NAV_LOITER_TIME enum. Loiter at the specified latitude, longitude and altitude for a certain amount of time. Multicopter vehicles stop at the point (within a vehicle-specific acceptance radius). Forward-only moving vehicles (e.g. fixed-wing) circle the point with the specified radius/direction. If the Heading Required parameter (2) is non-zero forward moving aircraft will only leave the loiter circle once heading towards the next waypoint. Params: 1) Loiter time (only starts once Lat, Lon and Alt is reached).; 2) Leave loiter circle only once heading towards the next waypoint (0 = False); 3) Loiter radius around waypoint for forward-only moving vehicles (not multicopters). If positive loiter clockwise, else counter-clockwise.; 4) Loiter circle exit location and/or path to next waypoint ("xtrack") for forward-only moving vehicles (not multicopters). 0 for the vehicle to converge towards the center xtrack when it leaves the loiter (the line between the centers of the current and next waypoint), 1 to converge to the direct line between the location that the vehicle exits the loiter radius and the next waypoint. Otherwise the angle (in degrees) between the tangent of the loiter circle and the center xtrack at which the vehicle must leave the loiter (and converge to the center xtrack). NaN to use the current system default xtrack behaviour.; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_LOITER_TIME MAV_CMD = 19
	// MAV_CMD_NAV_RETURN_TO_LAUNCH enum. Return to launch location. Params: 1) Empty; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_NAV_RETURN_TO_LAUNCH MAV_CMD = 20
	// MAV_CMD_NAV_LAND enum. Land at location. Params: 1) Minimum target altitude if landing is aborted (0 = undefined/use system default).; 2) Precision land mode.; 3) Empty; 4) Desired yaw angle. NaN to use the current system yaw heading mode (e.g. yaw towards next waypoint, yaw to home, etc.).; 5) Latitude.; 6) Longitude.; 7) Landing altitude (ground level in current frame).;
	MAV_CMD_NAV_LAND MAV_CMD = 21
	// MAV_CMD_NAV_TAKEOFF enum. Takeoff from ground / hand. Vehicles that support multiple takeoff modes (e.g. VTOL quadplane) should take off using the currently configured mode. Params: 1) Minimum pitch (if airspeed sensor present), desired pitch without sensor; 2) Empty; 3) Empty; 4) Yaw angle (if magnetometer present), ignored without magnetometer. NaN to use the current system yaw heading mode (e.g. yaw towards next waypoint, yaw to home, etc.).; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_TAKEOFF MAV_CMD = 22
	// MAV_CMD_NAV_LAND_LOCAL enum. Land at local position (local frame only). Params: 1) Landing target number (if available); 2) Maximum accepted offset from desired landing position - computed magnitude from spherical coordinates: d = sqrt(x^2 + y^2 + z^2), which gives the maximum accepted distance between the desired landing position and the position where the vehicle is about to land; 3) Landing descend rate; 4) Desired yaw angle; 5) Y-axis position; 6) X-axis position; 7) Z-axis / ground level position;
	MAV_CMD_NAV_LAND_LOCAL MAV_CMD = 23
	// MAV_CMD_NAV_TAKEOFF_LOCAL enum. Takeoff from local position (local frame only). Params: 1) Minimum pitch (if airspeed sensor present), desired pitch without sensor; 2) Empty; 3) Takeoff ascend rate; 4) Yaw angle (if magnetometer or another yaw estimation source present), ignored without one of these; 5) Y-axis position; 6) X-axis position; 7) Z-axis position;
	MAV_CMD_NAV_TAKEOFF_LOCAL MAV_CMD = 24
	// MAV_CMD_NAV_FOLLOW enum. Vehicle following, i.e. this waypoint represents the position of a moving vehicle. Params: 1) Following logic to use (e.g. loitering or sinusoidal following) - depends on specific autopilot implementation; 2) Ground speed of vehicle to be followed; 3) Radius around waypoint. If positive loiter clockwise, else counter-clockwise; 4) Desired yaw angle.; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_FOLLOW MAV_CMD = 25
	// MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT enum. Continue on the current course and climb/descend to specified altitude.  When the altitude is reached continue to the next command (i.e., don't proceed to the next command until the desired altitude is reached. Params: 1) Climb or Descend (0 = Neutral, command completes when within 5m of this command's altitude, 1 = Climbing, command completes when at or above this command's altitude, 2 = Descending, command completes when at or below this command's altitude.; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Desired altitude;
	MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT MAV_CMD = 30
	// MAV_CMD_NAV_LOITER_TO_ALT enum. Begin loiter at the specified Latitude and Longitude.  If Lat=Lon=0, then loiter at the current position.  Don't consider the navigation command complete (don't leave loiter) until the altitude has been reached. Additionally, if the Heading Required parameter is non-zero the aircraft will not leave the loiter until heading toward the next waypoint. Params: 1) Leave loiter circle only once heading towards the next waypoint (0 = False); 2) Loiter radius around waypoint for forward-only moving vehicles (not multicopters). If positive loiter clockwise, negative counter-clockwise, 0 means no change to standard loiter.; 3) Empty; 4) Loiter circle exit location and/or path to next waypoint ("xtrack") for forward-only moving vehicles (not multicopters). 0 for the vehicle to converge towards the center xtrack when it leaves the loiter (the line between the centers of the current and next waypoint), 1 to converge to the direct line between the location that the vehicle exits the loiter radius and the next waypoint. Otherwise the angle (in degrees) between the tangent of the loiter circle and the center xtrack at which the vehicle must leave the loiter (and converge to the center xtrack). NaN to use the current system default xtrack behaviour.; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_LOITER_TO_ALT MAV_CMD = 31
	// MAV_CMD_DO_FOLLOW enum. Begin following a target. Params: 1) System ID (of the FOLLOW_TARGET beacon). Send 0 to disable follow-me and return to the default position hold mode.; 2) Reserved; 3) Reserved; 4) Altitude mode: 0: Keep current altitude, 1: keep altitude difference to target, 2: go to a fixed altitude above home.; 5) Altitude above home. (used if mode=2); 6) Reserved; 7) Time to land in which the MAV should go to the default position hold mode after a message RX timeout.;
	MAV_CMD_DO_FOLLOW MAV_CMD = 32
	// MAV_CMD_DO_FOLLOW_REPOSITION enum. Reposition the MAV after a follow target command has been sent. Params: 1) Camera q1 (where 0 is on the ray from the camera to the tracking device); 2) Camera q2; 3) Camera q3; 4) Camera q4; 5) altitude offset from target; 6) X offset from target; 7) Y offset from target;
	MAV_CMD_DO_FOLLOW_REPOSITION MAV_CMD = 33
	// MAV_CMD_DO_ORBIT enum. Start orbiting on the circumference of a circle defined by the parameters. Setting any value NaN results in using defaults. Params: 1) Radius of the circle. positive: Orbit clockwise. negative: Orbit counter-clockwise.; 2) Tangential Velocity. NaN: Vehicle configuration default.; 3) Yaw behavior of the vehicle.; 4) Reserved (e.g. for dynamic center beacon options); 5) Center point latitude (if no MAV_FRAME specified) / X coordinate according to MAV_FRAME. NaN: Use current vehicle position or current center if already orbiting.; 6) Center point longitude (if no MAV_FRAME specified) / Y coordinate according to MAV_FRAME. NaN: Use current vehicle position or current center if already orbiting.; 7) Center point altitude (MSL) (if no MAV_FRAME specified) / Z coordinate according to MAV_FRAME. NaN: Use current vehicle position or current center if already orbiting.;
	MAV_CMD_DO_ORBIT MAV_CMD = 34
	// MAV_CMD_NAV_ROI enum. Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. Params: 1) Region of interest mode.; 2) Waypoint index/ target ID. (see MAV_ROI enum); 3) ROI index (allows a vehicle to manage multiple ROI's); 4) Empty; 5) x the location of the fixed ROI (see MAV_FRAME); 6) y; 7) z;
	MAV_CMD_NAV_ROI MAV_CMD = 80
	// MAV_CMD_NAV_PATHPLANNING enum. Control autonomous path planning on the MAV. Params: 1) 0: Disable local obstacle avoidance / local path planning (without resetting map), 1: Enable local path planning, 2: Enable and reset local path planning; 2) 0: Disable full path planning (without resetting map), 1: Enable, 2: Enable and reset map/occupancy grid, 3: Enable and reset planned route, but not occupancy grid; 3) Empty; 4) Yaw angle at goal; 5) Latitude/X of goal; 6) Longitude/Y of goal; 7) Altitude/Z of goal;
	MAV_CMD_NAV_PATHPLANNING MAV_CMD = 81
	// MAV_CMD_NAV_SPLINE_WAYPOINT enum. Navigate to waypoint using a spline path. Params: 1) Hold time. (ignored by fixed wing, time to stay at waypoint for rotary wing); 2) Empty; 3) Empty; 4) Empty; 5) Latitude/X of goal; 6) Longitude/Y of goal; 7) Altitude/Z of goal;
	MAV_CMD_NAV_SPLINE_WAYPOINT MAV_CMD = 82
	// MAV_CMD_NAV_VTOL_TAKEOFF enum. Takeoff from ground using VTOL mode, and transition to forward flight with specified heading. The command should be ignored by vehicles that dont support both VTOL and fixed-wing flight (multicopters, boats,etc.). Params: 1) Empty; 2) Front transition heading.; 3) Empty; 4) Yaw angle. NaN to use the current system yaw heading mode (e.g. yaw towards next waypoint, yaw to home, etc.).; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_VTOL_TAKEOFF MAV_CMD = 84
	// MAV_CMD_NAV_VTOL_LAND enum. Land using VTOL mode. Params: 1) Empty; 2) Empty; 3) Approach altitude (with the same reference as the Altitude field). NaN if unspecified.; 4) Yaw angle. NaN to use the current system yaw heading mode (e.g. yaw towards next waypoint, yaw to home, etc.).; 5) Latitude; 6) Longitude; 7) Altitude (ground level);
	MAV_CMD_NAV_VTOL_LAND MAV_CMD = 85
	// MAV_CMD_NAV_GUIDED_ENABLE enum. hand control over to an external controller. Params: 1) On / Off (&gt; 0.5f on); 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_NAV_GUIDED_ENABLE MAV_CMD = 92
	// MAV_CMD_NAV_DELAY enum. Delay the next navigation command a number of seconds or until a specified time. Params: 1) Delay (-1 to enable time-of-day fields); 2) hour (24h format, UTC, -1 to ignore); 3) minute (24h format, UTC, -1 to ignore); 4) second (24h format, UTC, -1 to ignore); 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_NAV_DELAY MAV_CMD = 93
	// MAV_CMD_NAV_PAYLOAD_PLACE enum. Descend and place payload. Vehicle moves to specified location, descends until it detects a hanging payload has reached the ground, and then releases the payload. If ground is not detected before the reaching the maximum descent value (param1), the command will complete without releasing the payload. Params: 1) Maximum distance to descend.; 2) Empty; 3) Empty; 4) Empty; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_PAYLOAD_PLACE MAV_CMD = 94
	// MAV_CMD_NAV_LAST enum. NOP - This command is only used to mark the upper limit of the NAV/ACTION commands in the enumeration. Params: 1) Empty; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_NAV_LAST MAV_CMD = 95
	// MAV_CMD_CONDITION_DELAY enum. Delay mission state machine. Params: 1) Delay; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_CONDITION_DELAY MAV_CMD = 112
	// MAV_CMD_CONDITION_CHANGE_ALT enum. Ascend/descend to target altitude at specified rate. Delay mission state machine until desired altitude reached. Params: 1) Descent / Ascend rate.; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Target Altitude;
	MAV_CMD_CONDITION_CHANGE_ALT MAV_CMD = 113
	// MAV_CMD_CONDITION_DISTANCE enum. Delay mission state machine until within desired distance of next NAV point. Params: 1) Distance.; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_CONDITION_DISTANCE MAV_CMD = 114
	// MAV_CMD_CONDITION_YAW enum. Reach a certain target angle. Params: 1) target angle, 0 is north; 2) angular speed; 3) direction: -1: counter clockwise, 1: clockwise; 4) 0: absolute angle, 1: relative offset; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_CONDITION_YAW MAV_CMD = 115
	// MAV_CMD_CONDITION_LAST enum. NOP - This command is only used to mark the upper limit of the CONDITION commands in the enumeration. Params: 1) Empty; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_CONDITION_LAST MAV_CMD = 159
	// MAV_CMD_DO_SET_MODE enum. Set system mode. Params: 1) Mode; 2) Custom mode - this is system specific, please refer to the individual autopilot specifications for details.; 3) Custom sub mode - this is system specific, please refer to the individual autopilot specifications for details.; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_SET_MODE MAV_CMD = 176
	// MAV_CMD_DO_JUMP enum. Jump to the desired command in the mission list.  Repeat this action only the specified number of times. Params: 1) Sequence number; 2) Repeat count; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_JUMP MAV_CMD = 177
	// MAV_CMD_DO_CHANGE_SPEED enum. Change speed and/or throttle set points. Params: 1) Speed type (0=Airspeed, 1=Ground Speed, 2=Climb Speed, 3=Descent Speed); 2) Speed (-1 indicates no change); 3) Throttle (-1 indicates no change); 4) 0: absolute, 1: relative; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_CHANGE_SPEED MAV_CMD = 178
	// MAV_CMD_DO_SET_HOME enum. Changes the home location either to the current location or a specified location. Params: 1) Use current (1=use current location, 0=use specified location); 2) Empty; 3) Empty; 4) Yaw angle. NaN to use default heading; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_DO_SET_HOME MAV_CMD = 179
	// MAV_CMD_DO_SET_PARAMETER enum. Set a system parameter.  Caution!  Use of this command requires knowledge of the numeric enumeration value of the parameter. Params: 1) Parameter number; 2) Parameter value; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_SET_PARAMETER MAV_CMD = 180
	// MAV_CMD_DO_SET_RELAY enum. Set a relay to a condition. Params: 1) Relay instance number.; 2) Setting. (1=on, 0=off, others possible depending on system hardware); 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_SET_RELAY MAV_CMD = 181
	// MAV_CMD_DO_REPEAT_RELAY enum. Cycle a relay on and off for a desired number of cycles with a desired period. Params: 1) Relay instance number.; 2) Cycle count.; 3) Cycle time.; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_REPEAT_RELAY MAV_CMD = 182
	// MAV_CMD_DO_SET_SERVO enum. Set a servo to a desired PWM value. Params: 1) Servo instance number.; 2) Pulse Width Modulation.; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_SET_SERVO MAV_CMD = 183
	// MAV_CMD_DO_REPEAT_SERVO enum. Cycle a between its nominal setting and a desired PWM for a desired number of cycles with a desired period. Params: 1) Servo instance number.; 2) Pulse Width Modulation.; 3) Cycle count.; 4) Cycle time.; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_REPEAT_SERVO MAV_CMD = 184
	// MAV_CMD_DO_FLIGHTTERMINATION enum. Terminate flight immediately. Params: 1) Flight termination activated if &gt; 0.5; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_FLIGHTTERMINATION MAV_CMD = 185
	// MAV_CMD_DO_CHANGE_ALTITUDE enum. Change altitude set point. Params: 1) Altitude; 2) Frame of new altitude.; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_CHANGE_ALTITUDE MAV_CMD = 186
	// MAV_CMD_DO_SET_ACTUATOR enum. Sets actuators (e.g. servos) to a desired value. The actuator numbers are mapped to specific outputs (e.g. on any MAIN or AUX PWM or UAVCAN) using a flight-stack specific mechanism (i.e. a parameter). Params: 1) Actuator 1 value, scaled from [-1 to 1]. NaN to ignore.; 2) Actuator 2 value, scaled from [-1 to 1]. NaN to ignore.; 3) Actuator 3 value, scaled from [-1 to 1]. NaN to ignore.; 4) Actuator 4 value, scaled from [-1 to 1]. NaN to ignore.; 5) Actuator 5 value, scaled from [-1 to 1]. NaN to ignore.; 6) Actuator 6 value, scaled from [-1 to 1]. NaN to ignore.; 7) Index of actuator set (i.e if set to 1, Actuator 1 becomes Actuator 7);
	MAV_CMD_DO_SET_ACTUATOR MAV_CMD = 187
	// MAV_CMD_DO_LAND_START enum. Mission command to perform a landing. This is used as a marker in a mission to tell the autopilot where a sequence of mission items that represents a landing starts. It may also be sent via a COMMAND_LONG to trigger a landing, in which case the nearest (geographically) landing sequence in the mission will be used. The Latitude/Longitude is optional, and may be set to 0 if not needed. If specified then it will be used to help find the closest landing sequence. Params: 1) Empty; 2) Empty; 3) Empty; 4) Empty; 5) Latitude; 6) Longitude; 7) Empty;
	MAV_CMD_DO_LAND_START MAV_CMD = 189
	// MAV_CMD_DO_RALLY_LAND enum. Mission command to perform a landing from a rally point. Params: 1) Break altitude; 2) Landing speed; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_RALLY_LAND MAV_CMD = 190
	// MAV_CMD_DO_GO_AROUND enum. Mission command to safely abort an autonomous landing. Params: 1) Altitude; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_GO_AROUND MAV_CMD = 191
	// MAV_CMD_DO_REPOSITION enum. Reposition the vehicle to a specific WGS84 global position. Params: 1) Ground speed, less than 0 (-1) for default; 2) Bitmask of option flags.; 3) Reserved; 4) Yaw heading. NaN to use the current system yaw heading mode (e.g. yaw towards next waypoint, yaw to home, etc.). For planes indicates loiter direction (0: clockwise, 1: counter clockwise); 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_DO_REPOSITION MAV_CMD = 192
	// MAV_CMD_DO_PAUSE_CONTINUE enum. If in a GPS controlled position mode, hold the current position or continue. Params: 1) 0: Pause current mission or reposition command, hold current position. 1: Continue mission. A VTOL capable vehicle should enter hover mode (multicopter and VTOL planes). A plane should loiter with the default loiter radius.; 2) Reserved; 3) Reserved; 4) Reserved; 5) Reserved; 6) Reserved; 7) Reserved;
	MAV_CMD_DO_PAUSE_CONTINUE MAV_CMD = 193
	// MAV_CMD_DO_SET_REVERSE enum. Set moving direction to forward or reverse. Params: 1) Direction (0=Forward, 1=Reverse); 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_SET_REVERSE MAV_CMD = 194
	// MAV_CMD_DO_SET_ROI_LOCATION enum. Sets the region of interest (ROI) to a location. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal is not to react to this message. Params: 1) Component ID of gimbal device to address (or 1-6 for non-MAVLink gimbal), 0 for all gimbal device components. Send command multiple times for more than one gimbal (but not all gimbals).; 2) Empty; 3) Empty; 4) Empty; 5) Latitude of ROI location; 6) Longitude of ROI location; 7) Altitude of ROI location;
	MAV_CMD_DO_SET_ROI_LOCATION MAV_CMD = 195
	// MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET enum. Sets the region of interest (ROI) to be toward next waypoint, with optional pitch/roll/yaw offset. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message. Params: 1) Component ID of gimbal device to address (or 1-6 for non-MAVLink gimbal), 0 for all gimbal device components. Send command multiple times for more than one gimbal (but not all gimbals).; 2) Empty; 3) Empty; 4) Empty; 5) Pitch offset from next waypoint, positive pitching up; 6) roll offset from next waypoint, positive rolling to the right; 7) yaw offset from next waypoint, positive yawing to the right;
	MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET MAV_CMD = 196
	// MAV_CMD_DO_SET_ROI_NONE enum. Cancels any previous ROI command returning the vehicle/sensors to default flight characteristics. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message. After this command the gimbal manager should go back to manual input if available, and otherwise assume a neutral position. Params: 1) Component ID of gimbal device to address (or 1-6 for non-MAVLink gimbal), 0 for all gimbal device components. Send command multiple times for more than one gimbal (but not all gimbals).; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_SET_ROI_NONE MAV_CMD = 197
	// MAV_CMD_DO_SET_ROI_SYSID enum. Mount tracks system with specified system ID. Determination of target vehicle position may be done with GLOBAL_POSITION_INT or any other means. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message. Params: 1) System ID; 2) Component ID of gimbal device to address (or 1-6 for non-MAVLink gimbal), 0 for all gimbal device components. Send command multiple times for more than one gimbal (but not all gimbals).;
	MAV_CMD_DO_SET_ROI_SYSID MAV_CMD = 198
	// MAV_CMD_DO_CONTROL_VIDEO enum. Control onboard camera system. Params: 1) Camera ID (-1 for all); 2) Transmission: 0: disabled, 1: enabled compressed, 2: enabled raw; 3) Transmission mode: 0: video stream, &gt;0: single images every n seconds; 4) Recording: 0: disabled, 1: enabled compressed, 2: enabled raw; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_CONTROL_VIDEO MAV_CMD = 200
	// MAV_CMD_DO_SET_ROI enum. Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. Params: 1) Region of interest mode.; 2) Waypoint index/ target ID (depends on param 1).; 3) Region of interest index. (allows a vehicle to manage multiple ROI's); 4) Empty; 5) MAV_ROI_WPNEXT: pitch offset from next waypoint, MAV_ROI_LOCATION: latitude; 6) MAV_ROI_WPNEXT: roll offset from next waypoint, MAV_ROI_LOCATION: longitude; 7) MAV_ROI_WPNEXT: yaw offset from next waypoint, MAV_ROI_LOCATION: altitude;
	MAV_CMD_DO_SET_ROI MAV_CMD = 201
	// MAV_CMD_DO_DIGICAM_CONFIGURE enum. Configure digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ). Params: 1) Modes: P, TV, AV, M, Etc.; 2) Shutter speed: Divisor number for one second.; 3) Aperture: F stop number.; 4) ISO number e.g. 80, 100, 200, Etc.; 5) Exposure type enumerator.; 6) Command Identity.; 7) Main engine cut-off time before camera trigger. (0 means no cut-off);
	MAV_CMD_DO_DIGICAM_CONFIGURE MAV_CMD = 202
	// MAV_CMD_DO_DIGICAM_CONTROL enum. Control digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ). Params: 1) Session control e.g. show/hide lens; 2) Zoom's absolute position; 3) Zooming step value to offset zoom from the current position; 4) Focus Locking, Unlocking or Re-locking; 5) Shooting Command; 6) Command Identity; 7) Test shot identifier. If set to 1, image will only be captured, but not counted towards internal frame count.;
	MAV_CMD_DO_DIGICAM_CONTROL MAV_CMD = 203
	// MAV_CMD_DO_MOUNT_CONFIGURE enum. Mission command to configure a camera or antenna mount. Params: 1) Mount operation mode; 2) stabilize roll? (1 = yes, 0 = no); 3) stabilize pitch? (1 = yes, 0 = no); 4) stabilize yaw? (1 = yes, 0 = no); 5) roll input (0 = angle body frame, 1 = angular rate, 2 = angle absolute frame); 6) pitch input (0 = angle body frame, 1 = angular rate, 2 = angle absolute frame); 7) yaw input (0 = angle body frame, 1 = angular rate, 2 = angle absolute frame);
	MAV_CMD_DO_MOUNT_CONFIGURE MAV_CMD = 204
	// MAV_CMD_DO_MOUNT_CONTROL enum. Mission command to control a camera or antenna mount. Params: 1) pitch depending on mount mode (degrees or degrees/second depending on pitch input).; 2) roll depending on mount mode (degrees or degrees/second depending on roll input).; 3) yaw depending on mount mode (degrees or degrees/second depending on yaw input).; 4) altitude depending on mount mode.; 5) latitude, set if appropriate mount mode.; 6) longitude, set if appropriate mount mode.; 7) Mount mode.;
	MAV_CMD_DO_MOUNT_CONTROL MAV_CMD = 205
	// MAV_CMD_DO_SET_CAM_TRIGG_DIST enum. Mission command to set camera trigger distance for this flight. The camera is triggered each time this distance is exceeded. This command can also be used to set the shutter integration time for the camera. Params: 1) Camera trigger distance. 0 to stop triggering.; 2) Camera shutter integration time. -1 or 0 to ignore; 3) Trigger camera once immediately. (0 = no trigger, 1 = trigger); 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_SET_CAM_TRIGG_DIST MAV_CMD = 206
	// MAV_CMD_DO_FENCE_ENABLE enum. Mission command to enable the geofence. Params: 1) enable? (0=disable, 1=enable, 2=disable_floor_only); 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_FENCE_ENABLE MAV_CMD = 207
	// MAV_CMD_DO_PARACHUTE enum. Mission item/command to release a parachute or enable/disable auto release. Params: 1) Action; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_PARACHUTE MAV_CMD = 208
	// MAV_CMD_DO_MOTOR_TEST enum. Mission command to perform motor test. Params: 1) Motor instance number. (from 1 to max number of motors on the vehicle); 2) Throttle type.; 3) Throttle.; 4) Timeout.; 5) Motor count. (number of motors to test to test in sequence, waiting for the timeout above between them; 0=1 motor, 1=1 motor, 2=2 motors...); 6) Motor test order.; 7) Empty;
	MAV_CMD_DO_MOTOR_TEST MAV_CMD = 209
	// MAV_CMD_DO_INVERTED_FLIGHT enum. Change to/from inverted flight. Params: 1) Inverted flight. (0=normal, 1=inverted); 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_INVERTED_FLIGHT MAV_CMD = 210
	// MAV_CMD_DO_GRIPPER enum. Mission command to operate a gripper. Params: 1) Gripper instance number.; 2) Gripper action to perform.; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_GRIPPER MAV_CMD = 211
	// MAV_CMD_DO_AUTOTUNE_ENABLE enum. Enable/disable autotune. Params: 1) Enable (1: enable, 0:disable).; 2) Empty.; 3) Empty.; 4) Empty.; 5) Empty.; 6) Empty.; 7) Empty.;
	MAV_CMD_DO_AUTOTUNE_ENABLE MAV_CMD = 212
	// MAV_CMD_NAV_SET_YAW_SPEED enum. Sets a desired vehicle turn angle and speed change. Params: 1) Yaw angle to adjust steering by.; 2) Speed.; 3) Final angle. (0=absolute, 1=relative); 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_NAV_SET_YAW_SPEED MAV_CMD = 213
	// MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL enum. Mission command to set camera trigger interval for this flight. If triggering is enabled, the camera is triggered each time this interval expires. This command can also be used to set the shutter integration time for the camera. Params: 1) Camera trigger cycle time. -1 or 0 to ignore.; 2) Camera shutter integration time. Should be less than trigger cycle time. -1 or 0 to ignore.; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL MAV_CMD = 214
	// MAV_CMD_DO_MOUNT_CONTROL_QUAT enum. Mission command to control a camera or antenna mount, using a quaternion as reference. Params: 1) quaternion param q1, w (1 in null-rotation); 2) quaternion param q2, x (0 in null-rotation); 3) quaternion param q3, y (0 in null-rotation); 4) quaternion param q4, z (0 in null-rotation); 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_MOUNT_CONTROL_QUAT MAV_CMD = 220
	// MAV_CMD_DO_GUIDED_MASTER enum. set id of master controller. Params: 1) System ID; 2) Component ID; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_GUIDED_MASTER MAV_CMD = 221
	// MAV_CMD_DO_GUIDED_LIMITS enum. Set limits for external control. Params: 1) Timeout - maximum time that external controller will be allowed to control vehicle. 0 means no timeout.; 2) Altitude (MSL) min - if vehicle moves below this alt, the command will be aborted and the mission will continue. 0 means no lower altitude limit.; 3) Altitude (MSL) max - if vehicle moves above this alt, the command will be aborted and the mission will continue. 0 means no upper altitude limit.; 4) Horizontal move limit - if vehicle moves more than this distance from its location at the moment the command was executed, the command will be aborted and the mission will continue. 0 means no horizontal move limit.; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_GUIDED_LIMITS MAV_CMD = 222
	// MAV_CMD_DO_ENGINE_CONTROL enum. Control vehicle engine. This is interpreted by the vehicles engine controller to change the target engine state. It is intended for vehicles with internal combustion engines. Params: 1) 0: Stop engine, 1:Start Engine; 2) 0: Warm start, 1:Cold start. Controls use of choke where applicable; 3) Height delay. This is for commanding engine start only after the vehicle has gained the specified height. Used in VTOL vehicles during takeoff to start engine after the aircraft is off the ground. Zero for no delay.; 4) Empty; 5) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_ENGINE_CONTROL MAV_CMD = 223
	// MAV_CMD_DO_SET_MISSION_CURRENT enum. Set the mission item with sequence number seq as current item. This means that the MAV will continue to this mission item on the shortest path (not following the mission items in-between). Params: 1) Mission sequence value to set; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_SET_MISSION_CURRENT MAV_CMD = 224
	// MAV_CMD_DO_LAST enum. NOP - This command is only used to mark the upper limit of the DO commands in the enumeration. Params: 1) Empty; 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_DO_LAST MAV_CMD = 240
	// MAV_CMD_PREFLIGHT_CALIBRATION enum. Trigger calibration. This command will be only accepted if in pre-flight mode. Except for Temperature Calibration, only one sensor should be set in a single message and all others should be zero. Params: 1) 1: gyro calibration, 3: gyro temperature calibration; 2) 1: magnetometer calibration; 3) 1: ground pressure calibration; 4) 1: radio RC calibration, 2: RC trim calibration; 5) 1: accelerometer calibration, 2: board level calibration, 3: accelerometer temperature calibration, 4: simple accelerometer calibration; 6) 1: APM: compass/motor interference calibration (PX4: airspeed calibration, deprecated), 2: airspeed calibration; 7) 1: ESC calibration, 3: barometer temperature calibration;
	MAV_CMD_PREFLIGHT_CALIBRATION MAV_CMD = 241
	// MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS enum. Set sensor offsets. This command will be only accepted if in pre-flight mode. Params: 1) Sensor to adjust the offsets for: 0: gyros, 1: accelerometer, 2: magnetometer, 3: barometer, 4: optical flow, 5: second magnetometer, 6: third magnetometer; 2) X axis offset (or generic dimension 1), in the sensor's raw units; 3) Y axis offset (or generic dimension 2), in the sensor's raw units; 4) Z axis offset (or generic dimension 3), in the sensor's raw units; 5) Generic dimension 4, in the sensor's raw units; 6) Generic dimension 5, in the sensor's raw units; 7) Generic dimension 6, in the sensor's raw units;
	MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS MAV_CMD = 242
	// MAV_CMD_PREFLIGHT_UAVCAN enum. Trigger UAVCAN configuration (actuator ID assignment and direction mapping). Note that this maps to the legacy UAVCAN v0 function UAVCAN_ENUMERATE, which is intended to be executed just once during initial vehicle configuration (it is not a normal pre-flight command and has been poorly named). Params: 1) 1: Trigger actuator ID assignment and direction mapping. 0: Cancel command.; 2) Reserved; 3) Reserved; 4) Reserved; 5) Reserved; 6) Reserved; 7) Reserved;
	MAV_CMD_PREFLIGHT_UAVCAN MAV_CMD = 243
	// MAV_CMD_PREFLIGHT_STORAGE enum. Request storage of different parameter values and logs. This command will be only accepted if in pre-flight mode. Params: 1) Parameter storage: 0: READ FROM FLASH/EEPROM, 1: WRITE CURRENT TO FLASH/EEPROM, 2: Reset to defaults; 2) Mission storage: 0: READ FROM FLASH/EEPROM, 1: WRITE CURRENT TO FLASH/EEPROM, 2: Reset to defaults; 3) Onboard logging: 0: Ignore, 1: Start default rate logging, -1: Stop logging, &gt; 1: logging rate (e.g. set to 1000 for 1000 Hz logging); 4) Reserved; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_PREFLIGHT_STORAGE MAV_CMD = 245
	// MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN enum. Request the reboot or shutdown of system components. Params: 1) 0: Do nothing for autopilot, 1: Reboot autopilot, 2: Shutdown autopilot, 3: Reboot autopilot and keep it in the bootloader until upgraded.; 2) 0: Do nothing for onboard computer, 1: Reboot onboard computer, 2: Shutdown onboard computer, 3: Reboot onboard computer and keep it in the bootloader until upgraded.; 3) WIP: 0: Do nothing for camera, 1: Reboot onboard camera, 2: Shutdown onboard camera, 3: Reboot onboard camera and keep it in the bootloader until upgraded; 4) WIP: 0: Do nothing for mount (e.g. gimbal), 1: Reboot mount, 2: Shutdown mount, 3: Reboot mount and keep it in the bootloader until upgraded; 5) Reserved (set to 0); 6) Reserved (set to 0); 7) WIP: ID (e.g. camera ID -1 for all IDs);
	MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN MAV_CMD = 246
	// MAV_CMD_DO_UPGRADE enum. Request a target system to start an upgrade of one (or all) of its components. For example, the command might be sent to a companion computer to cause it to upgrade a connected flight controller. The system doing the upgrade will report progress using the normal command protocol sequence for a long running operation. Command protocol information: https://mavlink.io/en/services/command.html. Params: 1) Component id of the component to be upgraded. If set to 0, all components should be upgraded.; 2) 0: Do not reboot component after the action is executed, 1: Reboot component after the action is executed.; 3) Reserved; 4) Reserved; 5) Reserved; 6) Reserved; 7) WIP: upgrade progress report rate (can be used for more granular control).;
	MAV_CMD_DO_UPGRADE MAV_CMD = 247
	// MAV_CMD_OVERRIDE_GOTO enum. Override current mission with command to pause mission, pause mission and move to position, continue/resume mission. When param 1 indicates that the mission is paused (MAV_GOTO_DO_HOLD), param 2 defines whether it holds in place or moves to another position. Params: 1) MAV_GOTO_DO_HOLD: pause mission and either hold or move to specified position (depending on param2), MAV_GOTO_DO_CONTINUE: resume mission.; 2) MAV_GOTO_HOLD_AT_CURRENT_POSITION: hold at current position, MAV_GOTO_HOLD_AT_SPECIFIED_POSITION: hold at specified position.; 3) Coordinate frame of hold point.; 4) Desired yaw angle.; 5) Latitude/X position.; 6) Longitude/Y position.; 7) Altitude/Z position.;
	MAV_CMD_OVERRIDE_GOTO MAV_CMD = 252
	// MAV_CMD_OBLIQUE_SURVEY enum. Mission command to set a Camera Auto Mount Pivoting Oblique Survey (Replaces CAM_TRIGG_DIST for this purpose). The camera is triggered each time this distance is exceeded, then the mount moves to the next position. Params 4~6 set-up the angle limits and number of positions for oblique survey, where mount-enabled vehicles automatically roll the camera between shots to emulate an oblique camera setup (providing an increased HFOV). This command can also be used to set the shutter integration time for the camera. Params: 1) Camera trigger distance. 0 to stop triggering.; 2) Camera shutter integration time. 0 to ignore; 3) The minimum interval in which the camera is capable of taking subsequent pictures repeatedly. 0 to ignore.; 4) Total number of roll positions at which the camera will capture photos (images captures spread evenly across the limits defined by param5).; 5) Angle limits that the camera can be rolled to left and right of center.; 6) Fixed pitch angle that the camera will hold in oblique mode if the mount is actuated in the pitch axis.; 7) Empty;
	MAV_CMD_OBLIQUE_SURVEY MAV_CMD = 260
	// MAV_CMD_MISSION_START enum. start running a mission. Params: 1) first_item: the first mission item to run; 2) last_item:  the last mission item to run (after this item is run, the mission ends);
	MAV_CMD_MISSION_START MAV_CMD = 300
	// MAV_CMD_COMPONENT_ARM_DISARM enum. Arms / Disarms a component. Params: 1) 0: disarm, 1: arm; 2) 0: arm-disarm unless prevented by safety checks (i.e. when landed), 21196: force arming/disarming (e.g. allow arming to override preflight checks and disarming in flight);
	MAV_CMD_COMPONENT_ARM_DISARM MAV_CMD = 400
	// MAV_CMD_ILLUMINATOR_ON_OFF enum. Turns illuminators ON/OFF. An illuminator is a light source that is used for lighting up dark areas external to the sytstem: e.g. a torch or searchlight (as opposed to a light source for illuminating the system itself, e.g. an indicator light). Params: 1) 0: Illuminators OFF, 1: Illuminators ON;
	MAV_CMD_ILLUMINATOR_ON_OFF MAV_CMD = 405
	// MAV_CMD_GET_HOME_POSITION enum. Request the home position from the vehicle. Params: 1) Reserved; 2) Reserved; 3) Reserved; 4) Reserved; 5) Reserved; 6) Reserved; 7) Reserved;
	MAV_CMD_GET_HOME_POSITION MAV_CMD = 410
	// MAV_CMD_INJECT_FAILURE enum. Inject artificial failure for testing purposes. Note that autopilots should implement an additional protection before accepting this command such as a specific param setting. Params: 1) The unit which is affected by the failure.; 2) The type how the failure manifests itself.; 3) Instance affected by failure (0 to signal all).;
	MAV_CMD_INJECT_FAILURE MAV_CMD = 420
	// MAV_CMD_START_RX_PAIR enum. Starts receiver pairing. Params: 1) 0:Spektrum.; 2) RC type.;
	MAV_CMD_START_RX_PAIR MAV_CMD = 500
	// MAV_CMD_GET_MESSAGE_INTERVAL enum. Request the interval between messages for a particular MAVLink message ID. The receiver should ACK the command and then emit its response in a MESSAGE_INTERVAL message. Params: 1) The MAVLink message ID;
	MAV_CMD_GET_MESSAGE_INTERVAL MAV_CMD = 510
	// MAV_CMD_SET_MESSAGE_INTERVAL enum. Set the interval between messages for a particular MAVLink message ID. This interface replaces REQUEST_DATA_STREAM. Params: 1) The MAVLink message ID; 2) The interval between two messages. Set to -1 to disable and 0 to request default rate.; 7) Target address of message stream (if message has target address fields). 0: Flight-stack default (recommended), 1: address of requestor, 2: broadcast.;
	MAV_CMD_SET_MESSAGE_INTERVAL MAV_CMD = 511
	// MAV_CMD_REQUEST_MESSAGE enum. Request the target system(s) emit a single instance of a specified message (i.e. a "one-shot" version of MAV_CMD_SET_MESSAGE_INTERVAL). Params: 1) The MAVLink message ID of the requested message.; 2) Use for index ID, if required. Otherwise, the use of this parameter (if any) must be defined in the requested message. By default assumed not used (0).; 3) The use of this parameter (if any), must be defined in the requested message. By default assumed not used (0).; 4) The use of this parameter (if any), must be defined in the requested message. By default assumed not used (0).; 5) The use of this parameter (if any), must be defined in the requested message. By default assumed not used (0).; 6) The use of this parameter (if any), must be defined in the requested message. By default assumed not used (0).; 7) Target address for requested message (if message has target address fields). 0: Flight-stack default, 1: address of requestor, 2: broadcast.;
	MAV_CMD_REQUEST_MESSAGE MAV_CMD = 512
	// MAV_CMD_REQUEST_PROTOCOL_VERSION enum. Request MAVLink protocol version compatibility. All receivers should ACK the command and then emit their capabilities in an PROTOCOL_VERSION message. Params: 1) 1: Request supported protocol versions by all nodes on the network; 2) Reserved (all remaining params);
	MAV_CMD_REQUEST_PROTOCOL_VERSION MAV_CMD = 519
	// MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES enum. Request autopilot capabilities. The receiver should ACK the command and then emit its capabilities in an AUTOPILOT_VERSION message. Params: 1) 1: Request autopilot version; 2) Reserved (all remaining params);
	MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES MAV_CMD = 520
	// MAV_CMD_REQUEST_CAMERA_INFORMATION enum. Request camera information (CAMERA_INFORMATION). Params: 1) 0: No action 1: Request camera capabilities; 2) Reserved (all remaining params);
	MAV_CMD_REQUEST_CAMERA_INFORMATION MAV_CMD = 521
	// MAV_CMD_REQUEST_CAMERA_SETTINGS enum. Request camera settings (CAMERA_SETTINGS). Params: 1) 0: No Action 1: Request camera settings; 2) Reserved (all remaining params);
	MAV_CMD_REQUEST_CAMERA_SETTINGS MAV_CMD = 522
	// MAV_CMD_REQUEST_STORAGE_INFORMATION enum. Request storage information (STORAGE_INFORMATION). Use the command's target_component to target a specific component's storage. Params: 1) Storage ID (0 for all, 1 for first, 2 for second, etc.); 2) 0: No Action 1: Request storage information; 3) Reserved (all remaining params);
	MAV_CMD_REQUEST_STORAGE_INFORMATION MAV_CMD = 525
	// MAV_CMD_STORAGE_FORMAT enum. Format a storage medium. Once format is complete, a STORAGE_INFORMATION message is sent. Use the command's target_component to target a specific component's storage. Params: 1) Storage ID (1 for first, 2 for second, etc.); 2) Format storage (and reset image log). 0: No action 1: Format storage; 3) Reset Image Log (without formatting storage medium). This will reset CAMERA_CAPTURE_STATUS.image_count and CAMERA_IMAGE_CAPTURED.image_index. 0: No action 1: Reset Image Log; 4) Reserved (all remaining params);
	MAV_CMD_STORAGE_FORMAT MAV_CMD = 526
	// MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS enum. Request camera capture status (CAMERA_CAPTURE_STATUS). Params: 1) 0: No Action 1: Request camera capture status; 2) Reserved (all remaining params);
	MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS MAV_CMD = 527
	// MAV_CMD_REQUEST_FLIGHT_INFORMATION enum. Request flight information (FLIGHT_INFORMATION). Params: 1) 1: Request flight information; 2) Reserved (all remaining params);
	MAV_CMD_REQUEST_FLIGHT_INFORMATION MAV_CMD = 528
	// MAV_CMD_RESET_CAMERA_SETTINGS enum. Reset all camera settings to Factory Default. Params: 1) 0: No Action 1: Reset all settings; 2) Reserved (all remaining params);
	MAV_CMD_RESET_CAMERA_SETTINGS MAV_CMD = 529
	// MAV_CMD_SET_CAMERA_MODE enum. Set camera running mode. Use NaN for reserved values. GCS will send a MAV_CMD_REQUEST_VIDEO_STREAM_STATUS command after a mode change if the camera supports video streaming. Params: 1) Reserved (Set to 0); 2) Camera mode; 3) ; 4) ; 7) ;
	MAV_CMD_SET_CAMERA_MODE MAV_CMD = 530
	// MAV_CMD_SET_CAMERA_ZOOM enum. Set camera zoom. Camera must respond with a CAMERA_SETTINGS message (on success). Params: 1) Zoom type; 2) Zoom value. The range of valid values depend on the zoom type.; 3) ; 4) ; 7) ;
	MAV_CMD_SET_CAMERA_ZOOM MAV_CMD = 531
	// MAV_CMD_SET_CAMERA_FOCUS enum. Set camera focus. Camera must respond with a CAMERA_SETTINGS message (on success). Params: 1) Focus type; 2) Focus value; 3) ; 4) ; 7) ;
	MAV_CMD_SET_CAMERA_FOCUS MAV_CMD = 532
	// MAV_CMD_JUMP_TAG enum. Tagged jump target. Can be jumped to with MAV_CMD_DO_JUMP_TAG. Params: 1) Tag.;
	MAV_CMD_JUMP_TAG MAV_CMD = 600
	// MAV_CMD_DO_JUMP_TAG enum. Jump to the matching tag in the mission list. Repeat this action for the specified number of times. A mission should contain a single matching tag for each jump. If this is not the case then a jump to a missing tag should complete the mission, and a jump where there are multiple matching tags should always select the one with the lowest mission sequence number. Params: 1) Target tag to jump to.; 2) Repeat count.;
	MAV_CMD_DO_JUMP_TAG MAV_CMD = 601
	// MAV_CMD_PARAM_TRANSACTION enum. Request to start or end a parameter transaction. Multiple kinds of transport layers can be used to exchange parameters in the transaction (param, param_ext and mavftp). The command response can either be a success/failure or an in progress in case the receiving side takes some time to apply the parameters. Params: 1) Action to be performed (start, commit, cancel, etc.); 2) Possible transport layers to set and get parameters via mavlink during a parameter transaction.; 3) Identifier for a specific transaction.;
	MAV_CMD_PARAM_TRANSACTION MAV_CMD = 900
	// MAV_CMD_DO_GIMBAL_MANAGER_PITCHYAW enum. High level setpoint to be sent to a gimbal manager to set a gimbal attitude. It is possible to set combinations of the values below. E.g. an angle as well as a desired angular rate can be used to get to this angle at a certain angular rate, or an angular rate only will result in continuous turning. NaN is to be used to signal unset. Note: a gimbal is never to react to this command but only the gimbal manager. Params: 1) Pitch angle (positive to pitch up, relative to vehicle for FOLLOW mode, relative to world horizon for LOCK mode).; 2) Yaw angle (positive to yaw to the right, relative to vehicle for FOLLOW mode, absolute to North for LOCK mode).; 3) Pitch rate (positive to pitch up).; 4) Yaw rate (positive to yaw to the right).; 5) Gimbal manager flags to use.; 7) Component ID of gimbal device to address (or 1-6 for non-MAVLink gimbal), 0 for all gimbal device components. Send command multiple times for more than one gimbal (but not all gimbals).;
	MAV_CMD_DO_GIMBAL_MANAGER_PITCHYAW MAV_CMD = 1000
	// MAV_CMD_DO_GIMBAL_MANAGER_CONFIGURE enum. Gimbal configuration to set which sysid/compid is in primary and secondary control. Params: 1) Sysid for primary control (0: no one in control, -1: leave unchanged, -2: set itself in control (for missions where the own sysid is still unknown), -3: remove control if currently in control).; 2) Compid for primary control (0: no one in control, -1: leave unchanged, -2: set itself in control (for missions where the own sysid is still unknown), -3: remove control if currently in control).; 3) Sysid for secondary control (0: no one in control, -1: leave unchanged, -2: set itself in control (for missions where the own sysid is still unknown), -3: remove control if currently in control).; 4) Compid for secondary control (0: no one in control, -1: leave unchanged, -2: set itself in control (for missions where the own sysid is still unknown), -3: remove control if currently in control).; 7) Component ID of gimbal device to address (or 1-6 for non-MAVLink gimbal), 0 for all gimbal device components. Send command multiple times for more than one gimbal (but not all gimbals).;
	MAV_CMD_DO_GIMBAL_MANAGER_CONFIGURE MAV_CMD = 1001
	// MAV_CMD_IMAGE_START_CAPTURE enum. Start image capture sequence. Sends CAMERA_IMAGE_CAPTURED after each capture. Use NaN for reserved values. Params: 1) Reserved (Set to 0); 2) Desired elapsed time between two consecutive pictures (in seconds). Minimum values depend on hardware (typically greater than 2 seconds).; 3) Total number of images to capture. 0 to capture forever/until MAV_CMD_IMAGE_STOP_CAPTURE.; 4) Capture sequence number starting from 1. This is only valid for single-capture (param3 == 1), otherwise set to 0. Increment the capture ID for each capture command to prevent double captures when a command is re-transmitted.; 5) ; 6) ; 7) ;
	MAV_CMD_IMAGE_START_CAPTURE MAV_CMD = 2000
	// MAV_CMD_IMAGE_STOP_CAPTURE enum. Stop image capture sequence Use NaN for reserved values. Params: 1) Reserved (Set to 0); 2) ; 3) ; 4) ; 7) ;
	MAV_CMD_IMAGE_STOP_CAPTURE MAV_CMD = 2001
	// MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE enum. Re-request a CAMERA_IMAGE_CAPTURED message. Params: 1) Sequence number for missing CAMERA_IMAGE_CAPTURED message; 2) ; 3) ; 4) ; 7) ;
	MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE MAV_CMD = 2002
	// MAV_CMD_DO_TRIGGER_CONTROL enum. Enable or disable on-board camera triggering system. Params: 1) Trigger enable/disable (0 for disable, 1 for start), -1 to ignore; 2) 1 to reset the trigger sequence, -1 or 0 to ignore; 3) 1 to pause triggering, but without switching the camera off or retracting it. -1 to ignore;
	MAV_CMD_DO_TRIGGER_CONTROL MAV_CMD = 2003
	// MAV_CMD_CAMERA_TRACK_POINT enum. If the camera supports point visual tracking (CAMERA_CAP_FLAGS_HAS_TRACKING_POINT is set), this command allows to initiate the tracking. Params: 1) Point to track x value (normalized 0..1, 0 is left, 1 is right).; 2) Point to track y value (normalized 0..1, 0 is top, 1 is bottom).; 3) Point radius (normalized 0..1, 0 is image left, 1 is image right).;
	MAV_CMD_CAMERA_TRACK_POINT MAV_CMD = 2004
	// MAV_CMD_CAMERA_TRACK_RECTANGLE enum. If the camera supports rectangle visual tracking (CAMERA_CAP_FLAGS_HAS_TRACKING_RECTANGLE is set), this command allows to initiate the tracking. Params: 1) Top left corner of rectangle x value (normalized 0..1, 0 is left, 1 is right).; 2) Top left corner of rectangle y value (normalized 0..1, 0 is top, 1 is bottom).; 3) Bottom right corner of rectangle x value (normalized 0..1, 0 is left, 1 is right).; 4) Bottom right corner of rectangle y value (normalized 0..1, 0 is top, 1 is bottom).;
	MAV_CMD_CAMERA_TRACK_RECTANGLE MAV_CMD = 2005
	// MAV_CMD_CAMERA_STOP_TRACKING enum. Stops ongoing tracking
	MAV_CMD_CAMERA_STOP_TRACKING MAV_CMD = 2010
	// MAV_CMD_VIDEO_START_CAPTURE enum. Starts video capture (recording). Params: 1) Video Stream ID (0 for all streams); 2) Frequency CAMERA_CAPTURE_STATUS messages should be sent while recording (0 for no messages, otherwise frequency); 3) ; 4) ; 5) ; 6) ; 7) ;
	MAV_CMD_VIDEO_START_CAPTURE MAV_CMD = 2500
	// MAV_CMD_VIDEO_STOP_CAPTURE enum. Stop the current video capture (recording). Params: 1) Video Stream ID (0 for all streams); 2) ; 3) ; 4) ; 5) ; 6) ; 7) ;
	MAV_CMD_VIDEO_STOP_CAPTURE MAV_CMD = 2501
	// MAV_CMD_VIDEO_START_STREAMING enum. Start video streaming. Params: 1) Video Stream ID (0 for all streams, 1 for first, 2 for second, etc.);
	MAV_CMD_VIDEO_START_STREAMING MAV_CMD = 2502
	// MAV_CMD_VIDEO_STOP_STREAMING enum. Stop the given video stream. Params: 1) Video Stream ID (0 for all streams, 1 for first, 2 for second, etc.);
	MAV_CMD_VIDEO_STOP_STREAMING MAV_CMD = 2503
	// MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION enum. Request video stream information (VIDEO_STREAM_INFORMATION). Params: 1) Video Stream ID (0 for all streams, 1 for first, 2 for second, etc.);
	MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION MAV_CMD = 2504
	// MAV_CMD_REQUEST_VIDEO_STREAM_STATUS enum. Request video stream status (VIDEO_STREAM_STATUS). Params: 1) Video Stream ID (0 for all streams, 1 for first, 2 for second, etc.);
	MAV_CMD_REQUEST_VIDEO_STREAM_STATUS MAV_CMD = 2505
	// MAV_CMD_LOGGING_START enum. Request to start streaming logging data over MAVLink (see also LOGGING_DATA message). Params: 1) Format: 0: ULog; 2) Reserved (set to 0); 3) Reserved (set to 0); 4) Reserved (set to 0); 5) Reserved (set to 0); 6) Reserved (set to 0); 7) Reserved (set to 0);
	MAV_CMD_LOGGING_START MAV_CMD = 2510
	// MAV_CMD_LOGGING_STOP enum. Request to stop streaming log data over MAVLink. Params: 1) Reserved (set to 0); 2) Reserved (set to 0); 3) Reserved (set to 0); 4) Reserved (set to 0); 5) Reserved (set to 0); 6) Reserved (set to 0); 7) Reserved (set to 0);
	MAV_CMD_LOGGING_STOP MAV_CMD = 2511
	// MAV_CMD_AIRFRAME_CONFIGURATION enum. Params: 1) Landing gear ID (default: 0, -1 for all); 2) Landing gear position (Down: 0, Up: 1, NaN for no change); 3) ; 4) ; 5) ; 6) ; 7) ;
	MAV_CMD_AIRFRAME_CONFIGURATION MAV_CMD = 2520
	// MAV_CMD_CONTROL_HIGH_LATENCY enum. Request to start/stop transmitting over the high latency telemetry. Params: 1) Control transmission over high latency telemetry (0: stop, 1: start); 2) Empty; 3) Empty; 4) Empty; 5) Empty; 6) Empty; 7) Empty;
	MAV_CMD_CONTROL_HIGH_LATENCY MAV_CMD = 2600
	// MAV_CMD_PANORAMA_CREATE enum. Create a panorama at the current position. Params: 1) Viewing angle horizontal of the panorama (+- 0.5 the total angle); 2) Viewing angle vertical of panorama.; 3) Speed of the horizontal rotation.; 4) Speed of the vertical rotation.;
	MAV_CMD_PANORAMA_CREATE MAV_CMD = 2800
	// MAV_CMD_DO_VTOL_TRANSITION enum. Request VTOL transition. Params: 1) The target VTOL state. Only MAV_VTOL_STATE_MC and MAV_VTOL_STATE_FW can be used.;
	MAV_CMD_DO_VTOL_TRANSITION MAV_CMD = 3000
	// MAV_CMD_ARM_AUTHORIZATION_REQUEST enum. Request authorization to arm the vehicle to a external entity, the arm authorizer is responsible to request all data that is needs from the vehicle before authorize or deny the request. If approved the progress of command_ack message should be set with period of time that this authorization is valid in seconds or in case it was denied it should be set with one of the reasons in ARM_AUTH_DENIED_REASON. Params: 1) Vehicle system id, this way ground station can request arm authorization on behalf of any vehicle;
	MAV_CMD_ARM_AUTHORIZATION_REQUEST MAV_CMD = 3001
	// MAV_CMD_SET_GUIDED_SUBMODE_STANDARD enum. This command sets the submode to standard guided when vehicle is in guided mode. The vehicle holds position and altitude and the user can input the desired velocities along all three axes
	MAV_CMD_SET_GUIDED_SUBMODE_STANDARD MAV_CMD = 4000
	// MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE enum. This command sets submode circle when vehicle is in guided mode. Vehicle flies along a circle facing the center of the circle. The user can input the velocity along the circle and change the radius. If no input is given the vehicle will hold position. Params: 1) Radius of desired circle in CIRCLE_MODE; 2) User defined; 3) User defined; 4) User defined; 5) Target latitude of center of circle in CIRCLE_MODE; 6) Target longitude of center of circle in CIRCLE_MODE;
	MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE MAV_CMD = 4001
	// MAV_CMD_CONDITION_GATE enum. Delay mission state machine until gate has been reached. Params: 1) Geometry: 0: orthogonal to path between previous and next waypoint.; 2) Altitude: 0: ignore altitude; 3) Empty; 4) Empty; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_CONDITION_GATE MAV_CMD = 4501
	// MAV_CMD_NAV_FENCE_RETURN_POINT enum. Fence return point. There can only be one fence return point. Params: 1) Reserved; 2) Reserved; 3) Reserved; 4) Reserved; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_FENCE_RETURN_POINT MAV_CMD = 5000
	// MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION enum. Fence vertex for an inclusion polygon (the polygon must not be self-intersecting). The vehicle must stay within this area. Minimum of 3 vertices required. Params: 1) Polygon vertex count; 2) Vehicle must be inside ALL inclusion zones in a single group, vehicle must be inside at least one group, must be the same for all points in each polygon; 3) Reserved; 4) Reserved; 5) Latitude; 6) Longitude; 7) Reserved;
	MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION MAV_CMD = 5001
	// MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION enum. Fence vertex for an exclusion polygon (the polygon must not be self-intersecting). The vehicle must stay outside this area. Minimum of 3 vertices required. Params: 1) Polygon vertex count; 2) Reserved; 3) Reserved; 4) Reserved; 5) Latitude; 6) Longitude; 7) Reserved;
	MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION MAV_CMD = 5002
	// MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION enum. Circular fence area. The vehicle must stay inside this area. Params: 1) Radius.; 2) Vehicle must be inside ALL inclusion zones in a single group, vehicle must be inside at least one group; 3) Reserved; 4) Reserved; 5) Latitude; 6) Longitude; 7) Reserved;
	MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION MAV_CMD = 5003
	// MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION enum. Circular fence area. The vehicle must stay outside this area. Params: 1) Radius.; 2) Reserved; 3) Reserved; 4) Reserved; 5) Latitude; 6) Longitude; 7) Reserved;
	MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION MAV_CMD = 5004
	// MAV_CMD_NAV_RALLY_POINT enum. Rally point. You can have multiple rally points defined. Params: 1) Reserved; 2) Reserved; 3) Reserved; 4) Reserved; 5) Latitude; 6) Longitude; 7) Altitude;
	MAV_CMD_NAV_RALLY_POINT MAV_CMD = 5100
	// MAV_CMD_UAVCAN_GET_NODE_INFO enum. Commands the vehicle to respond with a sequence of messages UAVCAN_NODE_INFO, one message per every UAVCAN node that is online. Note that some of the response messages can be lost, which the receiver can detect easily by checking whether every received UAVCAN_NODE_STATUS has a matching message UAVCAN_NODE_INFO received earlier; if not, this command should be sent again in order to request re-transmission of the node information messages. Params: 1) Reserved (set to 0); 2) Reserved (set to 0); 3) Reserved (set to 0); 4) Reserved (set to 0); 5) Reserved (set to 0); 6) Reserved (set to 0); 7) Reserved (set to 0);
	MAV_CMD_UAVCAN_GET_NODE_INFO MAV_CMD = 5200
	// MAV_CMD_PAYLOAD_PREPARE_DEPLOY enum. Deploy payload on a Lat / Lon / Alt position. This includes the navigation to reach the required release position and velocity. Params: 1) Operation mode. 0: prepare single payload deploy (overwriting previous requests), but do not execute it. 1: execute payload deploy immediately (rejecting further deploy commands during execution, but allowing abort). 2: add payload deploy to existing deployment list.; 2) Desired approach vector in compass heading. A negative value indicates the system can define the approach vector at will.; 3) Desired ground speed at release time. This can be overridden by the airframe in case it needs to meet minimum airspeed. A negative value indicates the system can define the ground speed at will.; 4) Minimum altitude clearance to the release position. A negative value indicates the system can define the clearance at will.; 5) Latitude. Note, if used in MISSION_ITEM (deprecated) the units are degrees (unscaled); 6) Longitude. Note, if used in MISSION_ITEM (deprecated) the units are degrees (unscaled); 7) Altitude (MSL);
	MAV_CMD_PAYLOAD_PREPARE_DEPLOY MAV_CMD = 30001
	// MAV_CMD_PAYLOAD_CONTROL_DEPLOY enum. Control the payload deployment. Params: 1) Operation mode. 0: Abort deployment, continue normal mission. 1: switch to payload deployment mode. 100: delete first payload deployment request. 101: delete all payload deployment requests.; 2) Reserved; 3) Reserved; 4) Reserved; 5) Reserved; 6) Reserved; 7) Reserved;
	MAV_CMD_PAYLOAD_CONTROL_DEPLOY MAV_CMD = 30002
	// MAV_CMD_FIXED_MAG_CAL_YAW enum. Magnetometer calibration based on provided known yaw. This allows for fast calibration using WMM field tables in the vehicle, given only the known yaw of the vehicle. If Latitude and longitude are both zero then use the current vehicle location. Params: 1) Yaw of vehicle in earth frame.; 2) CompassMask, 0 for all.; 3) Latitude.; 4) Longitude.; 5) Empty.; 6) Empty.; 7) Empty.;
	MAV_CMD_FIXED_MAG_CAL_YAW MAV_CMD = 42006
	// MAV_CMD_DO_WINCH enum. Command to operate winch. Params: 1) Winch instance number.; 2) Action to perform.; 3) Length of cable to release (negative to wind).; 4) Release rate (negative to wind).; 5) Empty.; 6) Empty.; 7) Empty.;
	MAV_CMD_DO_WINCH MAV_CMD = 42600
	// MAV_CMD_WAYPOINT_USER_1 enum. User defined waypoint item. Ground Station will show the Vehicle as flying through this item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) Latitude unscaled; 6) Longitude unscaled; 7) Altitude (MSL);
	MAV_CMD_WAYPOINT_USER_1 MAV_CMD = 31000
	// MAV_CMD_WAYPOINT_USER_2 enum. User defined waypoint item. Ground Station will show the Vehicle as flying through this item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) Latitude unscaled; 6) Longitude unscaled; 7) Altitude (MSL);
	MAV_CMD_WAYPOINT_USER_2 MAV_CMD = 31001
	// MAV_CMD_WAYPOINT_USER_3 enum. User defined waypoint item. Ground Station will show the Vehicle as flying through this item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) Latitude unscaled; 6) Longitude unscaled; 7) Altitude (MSL);
	MAV_CMD_WAYPOINT_USER_3 MAV_CMD = 31002
	// MAV_CMD_WAYPOINT_USER_4 enum. User defined waypoint item. Ground Station will show the Vehicle as flying through this item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) Latitude unscaled; 6) Longitude unscaled; 7) Altitude (MSL);
	MAV_CMD_WAYPOINT_USER_4 MAV_CMD = 31003
	// MAV_CMD_WAYPOINT_USER_5 enum. User defined waypoint item. Ground Station will show the Vehicle as flying through this item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) Latitude unscaled; 6) Longitude unscaled; 7) Altitude (MSL);
	MAV_CMD_WAYPOINT_USER_5 MAV_CMD = 31004
	// MAV_CMD_SPATIAL_USER_1 enum. User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) Latitude unscaled; 6) Longitude unscaled; 7) Altitude (MSL);
	MAV_CMD_SPATIAL_USER_1 MAV_CMD = 31005
	// MAV_CMD_SPATIAL_USER_2 enum. User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) Latitude unscaled; 6) Longitude unscaled; 7) Altitude (MSL);
	MAV_CMD_SPATIAL_USER_2 MAV_CMD = 31006
	// MAV_CMD_SPATIAL_USER_3 enum. User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) Latitude unscaled; 6) Longitude unscaled; 7) Altitude (MSL);
	MAV_CMD_SPATIAL_USER_3 MAV_CMD = 31007
	// MAV_CMD_SPATIAL_USER_4 enum. User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) Latitude unscaled; 6) Longitude unscaled; 7) Altitude (MSL);
	MAV_CMD_SPATIAL_USER_4 MAV_CMD = 31008
	// MAV_CMD_SPATIAL_USER_5 enum. User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) Latitude unscaled; 6) Longitude unscaled; 7) Altitude (MSL);
	MAV_CMD_SPATIAL_USER_5 MAV_CMD = 31009
	// MAV_CMD_USER_1 enum. User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) User defined; 6) User defined; 7) User defined;
	MAV_CMD_USER_1 MAV_CMD = 31010
	// MAV_CMD_USER_2 enum. User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) User defined; 6) User defined; 7) User defined;
	MAV_CMD_USER_2 MAV_CMD = 31011
	// MAV_CMD_USER_3 enum. User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) User defined; 6) User defined; 7) User defined;
	MAV_CMD_USER_3 MAV_CMD = 31012
	// MAV_CMD_USER_4 enum. User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) User defined; 6) User defined; 7) User defined;
	MAV_CMD_USER_4 MAV_CMD = 31013
	// MAV_CMD_USER_5 enum. User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item. Params: 1) User defined; 2) User defined; 3) User defined; 4) User defined; 5) User defined; 6) User defined; 7) User defined;
	MAV_CMD_USER_5 MAV_CMD = 31014
)

func (e MAV_CMD) String() string {
	if name, ok := map[MAV_CMD]string{
		MAV_CMD_NAV_WAYPOINT:                       "MAV_CMD_NAV_WAYPOINT",
		MAV_CMD_NAV_LOITER_UNLIM:                   "MAV_CMD_NAV_LOITER_UNLIM",
		MAV_CMD_NAV_LOITER_TURNS:                   "MAV_CMD_NAV_LOITER_TURNS",
		MAV_CMD_NAV_LOITER_TIME:                    "MAV_CMD_NAV_LOITER_TIME",
		MAV_CMD_NAV_RETURN_TO_LAUNCH:               "MAV_CMD_NAV_RETURN_TO_LAUNCH",
		MAV_CMD_NAV_LAND:                           "MAV_CMD_NAV_LAND",
		MAV_CMD_NAV_TAKEOFF:                        "MAV_CMD_NAV_TAKEOFF",
		MAV_CMD_NAV_LAND_LOCAL:                     "MAV_CMD_NAV_LAND_LOCAL",
		MAV_CMD_NAV_TAKEOFF_LOCAL:                  "MAV_CMD_NAV_TAKEOFF_LOCAL",
		MAV_CMD_NAV_FOLLOW:                         "MAV_CMD_NAV_FOLLOW",
		MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT:        "MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT",
		MAV_CMD_NAV_LOITER_TO_ALT:                  "MAV_CMD_NAV_LOITER_TO_ALT",
		MAV_CMD_DO_FOLLOW:                          "MAV_CMD_DO_FOLLOW",
		MAV_CMD_DO_FOLLOW_REPOSITION:               "MAV_CMD_DO_FOLLOW_REPOSITION",
		MAV_CMD_DO_ORBIT:                           "MAV_CMD_DO_ORBIT",
		MAV_CMD_NAV_ROI:                            "MAV_CMD_NAV_ROI",
		MAV_CMD_NAV_PATHPLANNING:                   "MAV_CMD_NAV_PATHPLANNING",
		MAV_CMD_NAV_SPLINE_WAYPOINT:                "MAV_CMD_NAV_SPLINE_WAYPOINT",
		MAV_CMD_NAV_VTOL_TAKEOFF:                   "MAV_CMD_NAV_VTOL_TAKEOFF",
		MAV_CMD_NAV_VTOL_LAND:                      "MAV_CMD_NAV_VTOL_LAND",
		MAV_CMD_NAV_GUIDED_ENABLE:                  "MAV_CMD_NAV_GUIDED_ENABLE",
		MAV_CMD_NAV_DELAY:                          "MAV_CMD_NAV_DELAY",
		MAV_CMD_NAV_PAYLOAD_PLACE:                  "MAV_CMD_NAV_PAYLOAD_PLACE",
		MAV_CMD_NAV_LAST:                           "MAV_CMD_NAV_LAST",
		MAV_CMD_CONDITION_DELAY:                    "MAV_CMD_CONDITION_DELAY",
		MAV_CMD_CONDITION_CHANGE_ALT:               "MAV_CMD_CONDITION_CHANGE_ALT",
		MAV_CMD_CONDITION_DISTANCE:                 "MAV_CMD_CONDITION_DISTANCE",
		MAV_CMD_CONDITION_YAW:                      "MAV_CMD_CONDITION_YAW",
		MAV_CMD_CONDITION_LAST:                     "MAV_CMD_CONDITION_LAST",
		MAV_CMD_DO_SET_MODE:                        "MAV_CMD_DO_SET_MODE",
		MAV_CMD_DO_JUMP:                            "MAV_CMD_DO_JUMP",
		MAV_CMD_DO_CHANGE_SPEED:                    "MAV_CMD_DO_CHANGE_SPEED",
		MAV_CMD_DO_SET_HOME:                        "MAV_CMD_DO_SET_HOME",
		MAV_CMD_DO_SET_PARAMETER:                   "MAV_CMD_DO_SET_PARAMETER",
		MAV_CMD_DO_SET_RELAY:                       "MAV_CMD_DO_SET_RELAY",
		MAV_CMD_DO_REPEAT_RELAY:                    "MAV_CMD_DO_REPEAT_RELAY",
		MAV_CMD_DO_SET_SERVO:                       "MAV_CMD_DO_SET_SERVO",
		MAV_CMD_DO_REPEAT_SERVO:                    "MAV_CMD_DO_REPEAT_SERVO",
		MAV_CMD_DO_FLIGHTTERMINATION:               "MAV_CMD_DO_FLIGHTTERMINATION",
		MAV_CMD_DO_CHANGE_ALTITUDE:                 "MAV_CMD_DO_CHANGE_ALTITUDE",
		MAV_CMD_DO_SET_ACTUATOR:                    "MAV_CMD_DO_SET_ACTUATOR",
		MAV_CMD_DO_LAND_START:                      "MAV_CMD_DO_LAND_START",
		MAV_CMD_DO_RALLY_LAND:                      "MAV_CMD_DO_RALLY_LAND",
		MAV_CMD_DO_GO_AROUND:                       "MAV_CMD_DO_GO_AROUND",
		MAV_CMD_DO_REPOSITION:                      "MAV_CMD_DO_REPOSITION",
		MAV_CMD_DO_PAUSE_CONTINUE:                  "MAV_CMD_DO_PAUSE_CONTINUE",
		MAV_CMD_DO_SET_REVERSE:                     "MAV_CMD_DO_SET_REVERSE",
		MAV_CMD_DO_SET_ROI_LOCATION:                "MAV_CMD_DO_SET_ROI_LOCATION",
		MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET:           "MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET",
		MAV_CMD_DO_SET_ROI_NONE:                    "MAV_CMD_DO_SET_ROI_NONE",
		MAV_CMD_DO_SET_ROI_SYSID:                   "MAV_CMD_DO_SET_ROI_SYSID",
		MAV_CMD_DO_CONTROL_VIDEO:                   "MAV_CMD_DO_CONTROL_VIDEO",
		MAV_CMD_DO_SET_ROI:                         "MAV_CMD_DO_SET_ROI",
		MAV_CMD_DO_DIGICAM_CONFIGURE:               "MAV_CMD_DO_DIGICAM_CONFIGURE",
		MAV_CMD_DO_DIGICAM_CONTROL:                 "MAV_CMD_DO_DIGICAM_CONTROL",
		MAV_CMD_DO_MOUNT_CONFIGURE:                 "MAV_CMD_DO_MOUNT_CONFIGURE",
		MAV_CMD_DO_MOUNT_CONTROL:                   "MAV_CMD_DO_MOUNT_CONTROL",
		MAV_CMD_DO_SET_CAM_TRIGG_DIST:              "MAV_CMD_DO_SET_CAM_TRIGG_DIST",
		MAV_CMD_DO_FENCE_ENABLE:                    "MAV_CMD_DO_FENCE_ENABLE",
		MAV_CMD_DO_PARACHUTE:                       "MAV_CMD_DO_PARACHUTE",
		MAV_CMD_DO_MOTOR_TEST:                      "MAV_CMD_DO_MOTOR_TEST",
		MAV_CMD_DO_INVERTED_FLIGHT:                 "MAV_CMD_DO_INVERTED_FLIGHT",
		MAV_CMD_DO_GRIPPER:                         "MAV_CMD_DO_GRIPPER",
		MAV_CMD_DO_AUTOTUNE_ENABLE:                 "MAV_CMD_DO_AUTOTUNE_ENABLE",
		MAV_CMD_NAV_SET_YAW_SPEED:                  "MAV_CMD_NAV_SET_YAW_SPEED",
		MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL:          "MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL",
		MAV_CMD_DO_MOUNT_CONTROL_QUAT:              "MAV_CMD_DO_MOUNT_CONTROL_QUAT",
		MAV_CMD_DO_GUIDED_MASTER:                   "MAV_CMD_DO_GUIDED_MASTER",
		MAV_CMD_DO_GUIDED_LIMITS:                   "MAV_CMD_DO_GUIDED_LIMITS",
		MAV_CMD_DO_ENGINE_CONTROL:                  "MAV_CMD_DO_ENGINE_CONTROL",
		MAV_CMD_DO_SET_MISSION_CURRENT:             "MAV_CMD_DO_SET_MISSION_CURRENT",
		MAV_CMD_DO_LAST:                            "MAV_CMD_DO_LAST",
		MAV_CMD_PREFLIGHT_CALIBRATION:              "MAV_CMD_PREFLIGHT_CALIBRATION",
		MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS:       "MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS",
		MAV_CMD_PREFLIGHT_UAVCAN:                   "MAV_CMD_PREFLIGHT_UAVCAN",
		MAV_CMD_PREFLIGHT_STORAGE:                  "MAV_CMD_PREFLIGHT_STORAGE",
		MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN:          "MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN",
		MAV_CMD_DO_UPGRADE:                         "MAV_CMD_DO_UPGRADE",
		MAV_CMD_OVERRIDE_GOTO:                      "MAV_CMD_OVERRIDE_GOTO",
		MAV_CMD_OBLIQUE_SURVEY:                     "MAV_CMD_OBLIQUE_SURVEY",
		MAV_CMD_MISSION_START:                      "MAV_CMD_MISSION_START",
		MAV_CMD_COMPONENT_ARM_DISARM:               "MAV_CMD_COMPONENT_ARM_DISARM",
		MAV_CMD_ILLUMINATOR_ON_OFF:                 "MAV_CMD_ILLUMINATOR_ON_OFF",
		MAV_CMD_GET_HOME_POSITION:                  "MAV_CMD_GET_HOME_POSITION",
		MAV_CMD_INJECT_FAILURE:                     "MAV_CMD_INJECT_FAILURE",
		MAV_CMD_START_RX_PAIR:                      "MAV_CMD_START_RX_PAIR",
		MAV_CMD_GET_MESSAGE_INTERVAL:               "MAV_CMD_GET_MESSAGE_INTERVAL",
		MAV_CMD_SET_MESSAGE_INTERVAL:               "MAV_CMD_SET_MESSAGE_INTERVAL",
		MAV_CMD_REQUEST_MESSAGE:                    "MAV_CMD_REQUEST_MESSAGE",
		MAV_CMD_REQUEST_PROTOCOL_VERSION:           "MAV_CMD_REQUEST_PROTOCOL_VERSION",
		MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES:     "MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES",
		MAV_CMD_REQUEST_CAMERA_INFORMATION:         "MAV_CMD_REQUEST_CAMERA_INFORMATION",
		MAV_CMD_REQUEST_CAMERA_SETTINGS:            "MAV_CMD_REQUEST_CAMERA_SETTINGS",
		MAV_CMD_REQUEST_STORAGE_INFORMATION:        "MAV_CMD_REQUEST_STORAGE_INFORMATION",
		MAV_CMD_STORAGE_FORMAT:                     "MAV_CMD_STORAGE_FORMAT",
		MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS:      "MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS",
		MAV_CMD_REQUEST_FLIGHT_INFORMATION:         "MAV_CMD_REQUEST_FLIGHT_INFORMATION",
		MAV_CMD_RESET_CAMERA_SETTINGS:              "MAV_CMD_RESET_CAMERA_SETTINGS",
		MAV_CMD_SET_CAMERA_MODE:                    "MAV_CMD_SET_CAMERA_MODE",
		MAV_CMD_SET_CAMERA_ZOOM:                    "MAV_CMD_SET_CAMERA_ZOOM",
		MAV_CMD_SET_CAMERA_FOCUS:                   "MAV_CMD_SET_CAMERA_FOCUS",
		MAV_CMD_JUMP_TAG:                           "MAV_CMD_JUMP_TAG",
		MAV_CMD_DO_JUMP_TAG:                        "MAV_CMD_DO_JUMP_TAG",
		MAV_CMD_PARAM_TRANSACTION:                  "MAV_CMD_PARAM_TRANSACTION",
		MAV_CMD_DO_GIMBAL_MANAGER_PITCHYAW:         "MAV_CMD_DO_GIMBAL_MANAGER_PITCHYAW",
		MAV_CMD_DO_GIMBAL_MANAGER_CONFIGURE:        "MAV_CMD_DO_GIMBAL_MANAGER_CONFIGURE",
		MAV_CMD_IMAGE_START_CAPTURE:                "MAV_CMD_IMAGE_START_CAPTURE",
		MAV_CMD_IMAGE_STOP_CAPTURE:                 "MAV_CMD_IMAGE_STOP_CAPTURE",
		MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE:       "MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE",
		MAV_CMD_DO_TRIGGER_CONTROL:                 "MAV_CMD_DO_TRIGGER_CONTROL",
		MAV_CMD_CAMERA_TRACK_POINT:                 "MAV_CMD_CAMERA_TRACK_POINT",
		MAV_CMD_CAMERA_TRACK_RECTANGLE:             "MAV_CMD_CAMERA_TRACK_RECTANGLE",
		MAV_CMD_CAMERA_STOP_TRACKING:               "MAV_CMD_CAMERA_STOP_TRACKING",
		MAV_CMD_VIDEO_START_CAPTURE:                "MAV_CMD_VIDEO_START_CAPTURE",
		MAV_CMD_VIDEO_STOP_CAPTURE:                 "MAV_CMD_VIDEO_STOP_CAPTURE",
		MAV_CMD_VIDEO_START_STREAMING:              "MAV_CMD_VIDEO_START_STREAMING",
		MAV_CMD_VIDEO_STOP_STREAMING:               "MAV_CMD_VIDEO_STOP_STREAMING",
		MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION:   "MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION",
		MAV_CMD_REQUEST_VIDEO_STREAM_STATUS:        "MAV_CMD_REQUEST_VIDEO_STREAM_STATUS",
		MAV_CMD_LOGGING_START:                      "MAV_CMD_LOGGING_START",
		MAV_CMD_LOGGING_STOP:                       "MAV_CMD_LOGGING_STOP",
		MAV_CMD_AIRFRAME_CONFIGURATION:             "MAV_CMD_AIRFRAME_CONFIGURATION",
		MAV_CMD_CONTROL_HIGH_LATENCY:               "MAV_CMD_CONTROL_HIGH_LATENCY",
		MAV_CMD_PANORAMA_CREATE:                    "MAV_CMD_PANORAMA_CREATE",
		MAV_CMD_DO_VTOL_TRANSITION:                 "MAV_CMD_DO_VTOL_TRANSITION",
		MAV_CMD_ARM_AUTHORIZATION_REQUEST:          "MAV_CMD_ARM_AUTHORIZATION_REQUEST",
		MAV_CMD_SET_GUIDED_SUBMODE_STANDARD:        "MAV_CMD_SET_GUIDED_SUBMODE_STANDARD",
		MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE:          "MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE",
		MAV_CMD_CONDITION_GATE:                     "MAV_CMD_CONDITION_GATE",
		MAV_CMD_NAV_FENCE_RETURN_POINT:             "MAV_CMD_NAV_FENCE_RETURN_POINT",
		MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION: "MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION",
		MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION: "MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION",
		MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION:         "MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION",
		MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION:         "MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION",
		MAV_CMD_NAV_RALLY_POINT:                    "MAV_CMD_NAV_RALLY_POINT",
		MAV_CMD_UAVCAN_GET_NODE_INFO:               "MAV_CMD_UAVCAN_GET_NODE_INFO",
		MAV_CMD_PAYLOAD_PREPARE_DEPLOY:             "MAV_CMD_PAYLOAD_PREPARE_DEPLOY",
		MAV_CMD_PAYLOAD_CONTROL_DEPLOY:             "MAV_CMD_PAYLOAD_CONTROL_DEPLOY",
		MAV_CMD_FIXED_MAG_CAL_YAW:                  "MAV_CMD_FIXED_MAG_CAL_YAW",
		MAV_CMD_DO_WINCH:                           "MAV_CMD_DO_WINCH",
		MAV_CMD_WAYPOINT_USER_1:                    "MAV_CMD_WAYPOINT_USER_1",
		MAV_CMD_WAYPOINT_USER_2:                    "MAV_CMD_WAYPOINT_USER_2",
		MAV_CMD_WAYPOINT_USER_3:                    "MAV_CMD_WAYPOINT_USER_3",
		MAV_CMD_WAYPOINT_USER_4:                    "MAV_CMD_WAYPOINT_USER_4",
		MAV_CMD_WAYPOINT_USER_5:                    "MAV_CMD_WAYPOINT_USER_5",
		MAV_CMD_SPATIAL_USER_1:                     "MAV_CMD_SPATIAL_USER_1",
		MAV_CMD_SPATIAL_USER_2:                     "MAV_CMD_SPATIAL_USER_2",
		MAV_CMD_SPATIAL_USER_3:                     "MAV_CMD_SPATIAL_USER_3",
		MAV_CMD_SPATIAL_USER_4:                     "MAV_CMD_SPATIAL_USER_4",
		MAV_CMD_SPATIAL_USER_5:                     "MAV_CMD_SPATIAL_USER_5",
		MAV_CMD_USER_1:                             "MAV_CMD_USER_1",
		MAV_CMD_USER_2:                             "MAV_CMD_USER_2",
		MAV_CMD_USER_3:                             "MAV_CMD_USER_3",
		MAV_CMD_USER_4:                             "MAV_CMD_USER_4",
		MAV_CMD_USER_5:                             "MAV_CMD_USER_5",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_CMD_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_CMD enums
func (e MAV_CMD) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_CMD{
		MAV_CMD_NAV_WAYPOINT,
		MAV_CMD_NAV_LOITER_UNLIM,
		MAV_CMD_NAV_LOITER_TURNS,
		MAV_CMD_NAV_LOITER_TIME,
		MAV_CMD_NAV_RETURN_TO_LAUNCH,
		MAV_CMD_NAV_LAND,
		MAV_CMD_NAV_TAKEOFF,
		MAV_CMD_NAV_LAND_LOCAL,
		MAV_CMD_NAV_TAKEOFF_LOCAL,
		MAV_CMD_NAV_FOLLOW,
		MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT,
		MAV_CMD_NAV_LOITER_TO_ALT,
		MAV_CMD_DO_FOLLOW,
		MAV_CMD_DO_FOLLOW_REPOSITION,
		MAV_CMD_DO_ORBIT,
		MAV_CMD_NAV_ROI,
		MAV_CMD_NAV_PATHPLANNING,
		MAV_CMD_NAV_SPLINE_WAYPOINT,
		MAV_CMD_NAV_VTOL_TAKEOFF,
		MAV_CMD_NAV_VTOL_LAND,
		MAV_CMD_NAV_GUIDED_ENABLE,
		MAV_CMD_NAV_DELAY,
		MAV_CMD_NAV_PAYLOAD_PLACE,
		MAV_CMD_NAV_LAST,
		MAV_CMD_CONDITION_DELAY,
		MAV_CMD_CONDITION_CHANGE_ALT,
		MAV_CMD_CONDITION_DISTANCE,
		MAV_CMD_CONDITION_YAW,
		MAV_CMD_CONDITION_LAST,
		MAV_CMD_DO_SET_MODE,
		MAV_CMD_DO_JUMP,
		MAV_CMD_DO_CHANGE_SPEED,
		MAV_CMD_DO_SET_HOME,
		MAV_CMD_DO_SET_PARAMETER,
		MAV_CMD_DO_SET_RELAY,
		MAV_CMD_DO_REPEAT_RELAY,
		MAV_CMD_DO_SET_SERVO,
		MAV_CMD_DO_REPEAT_SERVO,
		MAV_CMD_DO_FLIGHTTERMINATION,
		MAV_CMD_DO_CHANGE_ALTITUDE,
		MAV_CMD_DO_SET_ACTUATOR,
		MAV_CMD_DO_LAND_START,
		MAV_CMD_DO_RALLY_LAND,
		MAV_CMD_DO_GO_AROUND,
		MAV_CMD_DO_REPOSITION,
		MAV_CMD_DO_PAUSE_CONTINUE,
		MAV_CMD_DO_SET_REVERSE,
		MAV_CMD_DO_SET_ROI_LOCATION,
		MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET,
		MAV_CMD_DO_SET_ROI_NONE,
		MAV_CMD_DO_SET_ROI_SYSID,
		MAV_CMD_DO_CONTROL_VIDEO,
		MAV_CMD_DO_SET_ROI,
		MAV_CMD_DO_DIGICAM_CONFIGURE,
		MAV_CMD_DO_DIGICAM_CONTROL,
		MAV_CMD_DO_MOUNT_CONFIGURE,
		MAV_CMD_DO_MOUNT_CONTROL,
		MAV_CMD_DO_SET_CAM_TRIGG_DIST,
		MAV_CMD_DO_FENCE_ENABLE,
		MAV_CMD_DO_PARACHUTE,
		MAV_CMD_DO_MOTOR_TEST,
		MAV_CMD_DO_INVERTED_FLIGHT,
		MAV_CMD_DO_GRIPPER,
		MAV_CMD_DO_AUTOTUNE_ENABLE,
		MAV_CMD_NAV_SET_YAW_SPEED,
		MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL,
		MAV_CMD_DO_MOUNT_CONTROL_QUAT,
		MAV_CMD_DO_GUIDED_MASTER,
		MAV_CMD_DO_GUIDED_LIMITS,
		MAV_CMD_DO_ENGINE_CONTROL,
		MAV_CMD_DO_SET_MISSION_CURRENT,
		MAV_CMD_DO_LAST,
		MAV_CMD_PREFLIGHT_CALIBRATION,
		MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS,
		MAV_CMD_PREFLIGHT_UAVCAN,
		MAV_CMD_PREFLIGHT_STORAGE,
		MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN,
		MAV_CMD_DO_UPGRADE,
		MAV_CMD_OVERRIDE_GOTO,
		MAV_CMD_OBLIQUE_SURVEY,
		MAV_CMD_MISSION_START,
		MAV_CMD_COMPONENT_ARM_DISARM,
		MAV_CMD_ILLUMINATOR_ON_OFF,
		MAV_CMD_GET_HOME_POSITION,
		MAV_CMD_INJECT_FAILURE,
		MAV_CMD_START_RX_PAIR,
		MAV_CMD_GET_MESSAGE_INTERVAL,
		MAV_CMD_SET_MESSAGE_INTERVAL,
		MAV_CMD_REQUEST_MESSAGE,
		MAV_CMD_REQUEST_PROTOCOL_VERSION,
		MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES,
		MAV_CMD_REQUEST_CAMERA_INFORMATION,
		MAV_CMD_REQUEST_CAMERA_SETTINGS,
		MAV_CMD_REQUEST_STORAGE_INFORMATION,
		MAV_CMD_STORAGE_FORMAT,
		MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS,
		MAV_CMD_REQUEST_FLIGHT_INFORMATION,
		MAV_CMD_RESET_CAMERA_SETTINGS,
		MAV_CMD_SET_CAMERA_MODE,
		MAV_CMD_SET_CAMERA_ZOOM,
		MAV_CMD_SET_CAMERA_FOCUS,
		MAV_CMD_JUMP_TAG,
		MAV_CMD_DO_JUMP_TAG,
		MAV_CMD_PARAM_TRANSACTION,
		MAV_CMD_DO_GIMBAL_MANAGER_PITCHYAW,
		MAV_CMD_DO_GIMBAL_MANAGER_CONFIGURE,
		MAV_CMD_IMAGE_START_CAPTURE,
		MAV_CMD_IMAGE_STOP_CAPTURE,
		MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE,
		MAV_CMD_DO_TRIGGER_CONTROL,
		MAV_CMD_CAMERA_TRACK_POINT,
		MAV_CMD_CAMERA_TRACK_RECTANGLE,
		MAV_CMD_CAMERA_STOP_TRACKING,
		MAV_CMD_VIDEO_START_CAPTURE,
		MAV_CMD_VIDEO_STOP_CAPTURE,
		MAV_CMD_VIDEO_START_STREAMING,
		MAV_CMD_VIDEO_STOP_STREAMING,
		MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION,
		MAV_CMD_REQUEST_VIDEO_STREAM_STATUS,
		MAV_CMD_LOGGING_START,
		MAV_CMD_LOGGING_STOP,
		MAV_CMD_AIRFRAME_CONFIGURATION,
		MAV_CMD_CONTROL_HIGH_LATENCY,
		MAV_CMD_PANORAMA_CREATE,
		MAV_CMD_DO_VTOL_TRANSITION,
		MAV_CMD_ARM_AUTHORIZATION_REQUEST,
		MAV_CMD_SET_GUIDED_SUBMODE_STANDARD,
		MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE,
		MAV_CMD_CONDITION_GATE,
		MAV_CMD_NAV_FENCE_RETURN_POINT,
		MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION,
		MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION,
		MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION,
		MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION,
		MAV_CMD_NAV_RALLY_POINT,
		MAV_CMD_UAVCAN_GET_NODE_INFO,
		MAV_CMD_PAYLOAD_PREPARE_DEPLOY,
		MAV_CMD_PAYLOAD_CONTROL_DEPLOY,
		MAV_CMD_FIXED_MAG_CAL_YAW,
		MAV_CMD_DO_WINCH,
		MAV_CMD_WAYPOINT_USER_1,
		MAV_CMD_WAYPOINT_USER_2,
		MAV_CMD_WAYPOINT_USER_3,
		MAV_CMD_WAYPOINT_USER_4,
		MAV_CMD_WAYPOINT_USER_5,
		MAV_CMD_SPATIAL_USER_1,
		MAV_CMD_SPATIAL_USER_2,
		MAV_CMD_SPATIAL_USER_3,
		MAV_CMD_SPATIAL_USER_4,
		MAV_CMD_SPATIAL_USER_5,
		MAV_CMD_USER_1,
		MAV_CMD_USER_2,
		MAV_CMD_USER_3,
		MAV_CMD_USER_4,
		MAV_CMD_USER_5,
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

// MAV_DATA_STREAM type. A data stream is not a fixed set of messages, but rather a      recommendation to the autopilot software. Individual autopilots may or may not obey      the recommended messages.
type MAV_DATA_STREAM int

func (e MAV_DATA_STREAM) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_DATA_STREAM_ALL enum. Enable all data streams
	MAV_DATA_STREAM_ALL MAV_DATA_STREAM = 0
	// MAV_DATA_STREAM_RAW_SENSORS enum. Enable IMU_RAW, GPS_RAW, GPS_STATUS packets
	MAV_DATA_STREAM_RAW_SENSORS MAV_DATA_STREAM = 1
	// MAV_DATA_STREAM_EXTENDED_STATUS enum. Enable GPS_STATUS, CONTROL_STATUS, AUX_STATUS
	MAV_DATA_STREAM_EXTENDED_STATUS MAV_DATA_STREAM = 2
	// MAV_DATA_STREAM_RC_CHANNELS enum. Enable RC_CHANNELS_SCALED, RC_CHANNELS_RAW, SERVO_OUTPUT_RAW
	MAV_DATA_STREAM_RC_CHANNELS MAV_DATA_STREAM = 3
	// MAV_DATA_STREAM_RAW_CONTROLLER enum. Enable ATTITUDE_CONTROLLER_OUTPUT, POSITION_CONTROLLER_OUTPUT, NAV_CONTROLLER_OUTPUT
	MAV_DATA_STREAM_RAW_CONTROLLER MAV_DATA_STREAM = 4
	// MAV_DATA_STREAM_POSITION enum. Enable LOCAL_POSITION, GLOBAL_POSITION/GLOBAL_POSITION_INT messages
	MAV_DATA_STREAM_POSITION MAV_DATA_STREAM = 6
	// MAV_DATA_STREAM_EXTRA1 enum. Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA1 MAV_DATA_STREAM = 10
	// MAV_DATA_STREAM_EXTRA2 enum. Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA2 MAV_DATA_STREAM = 11
	// MAV_DATA_STREAM_EXTRA3 enum. Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA3 MAV_DATA_STREAM = 12
)

func (e MAV_DATA_STREAM) String() string {
	if name, ok := map[MAV_DATA_STREAM]string{
		MAV_DATA_STREAM_ALL:             "MAV_DATA_STREAM_ALL",
		MAV_DATA_STREAM_RAW_SENSORS:     "MAV_DATA_STREAM_RAW_SENSORS",
		MAV_DATA_STREAM_EXTENDED_STATUS: "MAV_DATA_STREAM_EXTENDED_STATUS",
		MAV_DATA_STREAM_RC_CHANNELS:     "MAV_DATA_STREAM_RC_CHANNELS",
		MAV_DATA_STREAM_RAW_CONTROLLER:  "MAV_DATA_STREAM_RAW_CONTROLLER",
		MAV_DATA_STREAM_POSITION:        "MAV_DATA_STREAM_POSITION",
		MAV_DATA_STREAM_EXTRA1:          "MAV_DATA_STREAM_EXTRA1",
		MAV_DATA_STREAM_EXTRA2:          "MAV_DATA_STREAM_EXTRA2",
		MAV_DATA_STREAM_EXTRA3:          "MAV_DATA_STREAM_EXTRA3",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_DATA_STREAM_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_DATA_STREAM enums
func (e MAV_DATA_STREAM) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_DATA_STREAM{
		MAV_DATA_STREAM_ALL,
		MAV_DATA_STREAM_RAW_SENSORS,
		MAV_DATA_STREAM_EXTENDED_STATUS,
		MAV_DATA_STREAM_RC_CHANNELS,
		MAV_DATA_STREAM_RAW_CONTROLLER,
		MAV_DATA_STREAM_POSITION,
		MAV_DATA_STREAM_EXTRA1,
		MAV_DATA_STREAM_EXTRA2,
		MAV_DATA_STREAM_EXTRA3,
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

// MAV_ROI type. The ROI (region of interest) for the vehicle. This can be                 be used by the vehicle for camera/vehicle attitude alignment (see                 MAV_CMD_NAV_ROI).
type MAV_ROI int

func (e MAV_ROI) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ROI_NONE enum. No region of interest
	MAV_ROI_NONE MAV_ROI = 0
	// MAV_ROI_WPNEXT enum. Point toward next waypoint, with optional pitch/roll/yaw offset
	MAV_ROI_WPNEXT MAV_ROI = 1
	// MAV_ROI_WPINDEX enum. Point toward given waypoint
	MAV_ROI_WPINDEX MAV_ROI = 2
	// MAV_ROI_LOCATION enum. Point toward fixed location
	MAV_ROI_LOCATION MAV_ROI = 3
	// MAV_ROI_TARGET enum. Point toward of given id
	MAV_ROI_TARGET MAV_ROI = 4
)

func (e MAV_ROI) String() string {
	if name, ok := map[MAV_ROI]string{
		MAV_ROI_NONE:     "MAV_ROI_NONE",
		MAV_ROI_WPNEXT:   "MAV_ROI_WPNEXT",
		MAV_ROI_WPINDEX:  "MAV_ROI_WPINDEX",
		MAV_ROI_LOCATION: "MAV_ROI_LOCATION",
		MAV_ROI_TARGET:   "MAV_ROI_TARGET",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ROI_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ROI enums
func (e MAV_ROI) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ROI{
		MAV_ROI_NONE,
		MAV_ROI_WPNEXT,
		MAV_ROI_WPINDEX,
		MAV_ROI_LOCATION,
		MAV_ROI_TARGET,
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

// MAV_CMD_ACK type. ACK / NACK / ERROR values as a result of MAV_CMDs and for mission item transmission.
type MAV_CMD_ACK int

func (e MAV_CMD_ACK) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_CMD_ACK_OK enum. Command / mission item is ok
	MAV_CMD_ACK_OK MAV_CMD_ACK = 0
	// MAV_CMD_ACK_ERR_FAIL enum. Generic error message if none of the other reasons fails or if no detailed error reporting is implemented
	MAV_CMD_ACK_ERR_FAIL MAV_CMD_ACK = 1
	// MAV_CMD_ACK_ERR_ACCESS_DENIED enum. The system is refusing to accept this command from this source / communication partner
	MAV_CMD_ACK_ERR_ACCESS_DENIED MAV_CMD_ACK = 2
	// MAV_CMD_ACK_ERR_NOT_SUPPORTED enum. Command or mission item is not supported, other commands would be accepted
	MAV_CMD_ACK_ERR_NOT_SUPPORTED MAV_CMD_ACK = 3
	// MAV_CMD_ACK_ERR_COORDINATE_FRAME_NOT_SUPPORTED enum. The coordinate frame of this command / mission item is not supported
	MAV_CMD_ACK_ERR_COORDINATE_FRAME_NOT_SUPPORTED MAV_CMD_ACK = 4
	// MAV_CMD_ACK_ERR_COORDINATES_OUT_OF_RANGE enum. The coordinate frame of this command is ok, but he coordinate values exceed the safety limits of this system. This is a generic error, please use the more specific error messages below if possible
	MAV_CMD_ACK_ERR_COORDINATES_OUT_OF_RANGE MAV_CMD_ACK = 5
	// MAV_CMD_ACK_ERR_X_LAT_OUT_OF_RANGE enum. The X or latitude value is out of range
	MAV_CMD_ACK_ERR_X_LAT_OUT_OF_RANGE MAV_CMD_ACK = 6
	// MAV_CMD_ACK_ERR_Y_LON_OUT_OF_RANGE enum. The Y or longitude value is out of range
	MAV_CMD_ACK_ERR_Y_LON_OUT_OF_RANGE MAV_CMD_ACK = 7
	// MAV_CMD_ACK_ERR_Z_ALT_OUT_OF_RANGE enum. The Z or altitude value is out of range
	MAV_CMD_ACK_ERR_Z_ALT_OUT_OF_RANGE MAV_CMD_ACK = 8
)

func (e MAV_CMD_ACK) String() string {
	if name, ok := map[MAV_CMD_ACK]string{
		MAV_CMD_ACK_OK:                                 "MAV_CMD_ACK_OK",
		MAV_CMD_ACK_ERR_FAIL:                           "MAV_CMD_ACK_ERR_FAIL",
		MAV_CMD_ACK_ERR_ACCESS_DENIED:                  "MAV_CMD_ACK_ERR_ACCESS_DENIED",
		MAV_CMD_ACK_ERR_NOT_SUPPORTED:                  "MAV_CMD_ACK_ERR_NOT_SUPPORTED",
		MAV_CMD_ACK_ERR_COORDINATE_FRAME_NOT_SUPPORTED: "MAV_CMD_ACK_ERR_COORDINATE_FRAME_NOT_SUPPORTED",
		MAV_CMD_ACK_ERR_COORDINATES_OUT_OF_RANGE:       "MAV_CMD_ACK_ERR_COORDINATES_OUT_OF_RANGE",
		MAV_CMD_ACK_ERR_X_LAT_OUT_OF_RANGE:             "MAV_CMD_ACK_ERR_X_LAT_OUT_OF_RANGE",
		MAV_CMD_ACK_ERR_Y_LON_OUT_OF_RANGE:             "MAV_CMD_ACK_ERR_Y_LON_OUT_OF_RANGE",
		MAV_CMD_ACK_ERR_Z_ALT_OUT_OF_RANGE:             "MAV_CMD_ACK_ERR_Z_ALT_OUT_OF_RANGE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_CMD_ACK_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_CMD_ACK enums
func (e MAV_CMD_ACK) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_CMD_ACK{
		MAV_CMD_ACK_OK,
		MAV_CMD_ACK_ERR_FAIL,
		MAV_CMD_ACK_ERR_ACCESS_DENIED,
		MAV_CMD_ACK_ERR_NOT_SUPPORTED,
		MAV_CMD_ACK_ERR_COORDINATE_FRAME_NOT_SUPPORTED,
		MAV_CMD_ACK_ERR_COORDINATES_OUT_OF_RANGE,
		MAV_CMD_ACK_ERR_X_LAT_OUT_OF_RANGE,
		MAV_CMD_ACK_ERR_Y_LON_OUT_OF_RANGE,
		MAV_CMD_ACK_ERR_Z_ALT_OUT_OF_RANGE,
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

// MAV_PARAM_TYPE type. Specifies the datatype of a MAVLink parameter.
type MAV_PARAM_TYPE int

func (e MAV_PARAM_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_PARAM_TYPE_UINT8 enum. 8-bit unsigned integer
	MAV_PARAM_TYPE_UINT8 MAV_PARAM_TYPE = 1
	// MAV_PARAM_TYPE_INT8 enum. 8-bit signed integer
	MAV_PARAM_TYPE_INT8 MAV_PARAM_TYPE = 2
	// MAV_PARAM_TYPE_UINT16 enum. 16-bit unsigned integer
	MAV_PARAM_TYPE_UINT16 MAV_PARAM_TYPE = 3
	// MAV_PARAM_TYPE_INT16 enum. 16-bit signed integer
	MAV_PARAM_TYPE_INT16 MAV_PARAM_TYPE = 4
	// MAV_PARAM_TYPE_UINT32 enum. 32-bit unsigned integer
	MAV_PARAM_TYPE_UINT32 MAV_PARAM_TYPE = 5
	// MAV_PARAM_TYPE_INT32 enum. 32-bit signed integer
	MAV_PARAM_TYPE_INT32 MAV_PARAM_TYPE = 6
	// MAV_PARAM_TYPE_UINT64 enum. 64-bit unsigned integer
	MAV_PARAM_TYPE_UINT64 MAV_PARAM_TYPE = 7
	// MAV_PARAM_TYPE_INT64 enum. 64-bit signed integer
	MAV_PARAM_TYPE_INT64 MAV_PARAM_TYPE = 8
	// MAV_PARAM_TYPE_REAL32 enum. 32-bit floating-point
	MAV_PARAM_TYPE_REAL32 MAV_PARAM_TYPE = 9
	// MAV_PARAM_TYPE_REAL64 enum. 64-bit floating-point
	MAV_PARAM_TYPE_REAL64 MAV_PARAM_TYPE = 10
)

func (e MAV_PARAM_TYPE) String() string {
	if name, ok := map[MAV_PARAM_TYPE]string{
		MAV_PARAM_TYPE_UINT8:  "MAV_PARAM_TYPE_UINT8",
		MAV_PARAM_TYPE_INT8:   "MAV_PARAM_TYPE_INT8",
		MAV_PARAM_TYPE_UINT16: "MAV_PARAM_TYPE_UINT16",
		MAV_PARAM_TYPE_INT16:  "MAV_PARAM_TYPE_INT16",
		MAV_PARAM_TYPE_UINT32: "MAV_PARAM_TYPE_UINT32",
		MAV_PARAM_TYPE_INT32:  "MAV_PARAM_TYPE_INT32",
		MAV_PARAM_TYPE_UINT64: "MAV_PARAM_TYPE_UINT64",
		MAV_PARAM_TYPE_INT64:  "MAV_PARAM_TYPE_INT64",
		MAV_PARAM_TYPE_REAL32: "MAV_PARAM_TYPE_REAL32",
		MAV_PARAM_TYPE_REAL64: "MAV_PARAM_TYPE_REAL64",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_PARAM_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_PARAM_TYPE enums
func (e MAV_PARAM_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_PARAM_TYPE{
		MAV_PARAM_TYPE_UINT8,
		MAV_PARAM_TYPE_INT8,
		MAV_PARAM_TYPE_UINT16,
		MAV_PARAM_TYPE_INT16,
		MAV_PARAM_TYPE_UINT32,
		MAV_PARAM_TYPE_INT32,
		MAV_PARAM_TYPE_UINT64,
		MAV_PARAM_TYPE_INT64,
		MAV_PARAM_TYPE_REAL32,
		MAV_PARAM_TYPE_REAL64,
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

// MAV_PARAM_EXT_TYPE type. Specifies the datatype of a MAVLink extended parameter.
type MAV_PARAM_EXT_TYPE int

func (e MAV_PARAM_EXT_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_PARAM_EXT_TYPE_UINT8 enum. 8-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT8 MAV_PARAM_EXT_TYPE = 1
	// MAV_PARAM_EXT_TYPE_INT8 enum. 8-bit signed integer
	MAV_PARAM_EXT_TYPE_INT8 MAV_PARAM_EXT_TYPE = 2
	// MAV_PARAM_EXT_TYPE_UINT16 enum. 16-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT16 MAV_PARAM_EXT_TYPE = 3
	// MAV_PARAM_EXT_TYPE_INT16 enum. 16-bit signed integer
	MAV_PARAM_EXT_TYPE_INT16 MAV_PARAM_EXT_TYPE = 4
	// MAV_PARAM_EXT_TYPE_UINT32 enum. 32-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT32 MAV_PARAM_EXT_TYPE = 5
	// MAV_PARAM_EXT_TYPE_INT32 enum. 32-bit signed integer
	MAV_PARAM_EXT_TYPE_INT32 MAV_PARAM_EXT_TYPE = 6
	// MAV_PARAM_EXT_TYPE_UINT64 enum. 64-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT64 MAV_PARAM_EXT_TYPE = 7
	// MAV_PARAM_EXT_TYPE_INT64 enum. 64-bit signed integer
	MAV_PARAM_EXT_TYPE_INT64 MAV_PARAM_EXT_TYPE = 8
	// MAV_PARAM_EXT_TYPE_REAL32 enum. 32-bit floating-point
	MAV_PARAM_EXT_TYPE_REAL32 MAV_PARAM_EXT_TYPE = 9
	// MAV_PARAM_EXT_TYPE_REAL64 enum. 64-bit floating-point
	MAV_PARAM_EXT_TYPE_REAL64 MAV_PARAM_EXT_TYPE = 10
	// MAV_PARAM_EXT_TYPE_CUSTOM enum. Custom Type
	MAV_PARAM_EXT_TYPE_CUSTOM MAV_PARAM_EXT_TYPE = 11
)

func (e MAV_PARAM_EXT_TYPE) String() string {
	if name, ok := map[MAV_PARAM_EXT_TYPE]string{
		MAV_PARAM_EXT_TYPE_UINT8:  "MAV_PARAM_EXT_TYPE_UINT8",
		MAV_PARAM_EXT_TYPE_INT8:   "MAV_PARAM_EXT_TYPE_INT8",
		MAV_PARAM_EXT_TYPE_UINT16: "MAV_PARAM_EXT_TYPE_UINT16",
		MAV_PARAM_EXT_TYPE_INT16:  "MAV_PARAM_EXT_TYPE_INT16",
		MAV_PARAM_EXT_TYPE_UINT32: "MAV_PARAM_EXT_TYPE_UINT32",
		MAV_PARAM_EXT_TYPE_INT32:  "MAV_PARAM_EXT_TYPE_INT32",
		MAV_PARAM_EXT_TYPE_UINT64: "MAV_PARAM_EXT_TYPE_UINT64",
		MAV_PARAM_EXT_TYPE_INT64:  "MAV_PARAM_EXT_TYPE_INT64",
		MAV_PARAM_EXT_TYPE_REAL32: "MAV_PARAM_EXT_TYPE_REAL32",
		MAV_PARAM_EXT_TYPE_REAL64: "MAV_PARAM_EXT_TYPE_REAL64",
		MAV_PARAM_EXT_TYPE_CUSTOM: "MAV_PARAM_EXT_TYPE_CUSTOM",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_PARAM_EXT_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_PARAM_EXT_TYPE enums
func (e MAV_PARAM_EXT_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_PARAM_EXT_TYPE{
		MAV_PARAM_EXT_TYPE_UINT8,
		MAV_PARAM_EXT_TYPE_INT8,
		MAV_PARAM_EXT_TYPE_UINT16,
		MAV_PARAM_EXT_TYPE_INT16,
		MAV_PARAM_EXT_TYPE_UINT32,
		MAV_PARAM_EXT_TYPE_INT32,
		MAV_PARAM_EXT_TYPE_UINT64,
		MAV_PARAM_EXT_TYPE_INT64,
		MAV_PARAM_EXT_TYPE_REAL32,
		MAV_PARAM_EXT_TYPE_REAL64,
		MAV_PARAM_EXT_TYPE_CUSTOM,
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

// MAV_RESULT type. Result from a MAVLink command (MAV_CMD)
type MAV_RESULT int

func (e MAV_RESULT) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_RESULT_ACCEPTED enum. Command is valid (is supported and has valid parameters), and was executed
	MAV_RESULT_ACCEPTED MAV_RESULT = 0
	// MAV_RESULT_TEMPORARILY_REJECTED enum. Command is valid, but cannot be executed at this time. This is used to indicate a problem that should be fixed just by waiting (e.g. a state machine is busy, can't arm because have not got GPS lock, etc.). Retrying later should work
	MAV_RESULT_TEMPORARILY_REJECTED MAV_RESULT = 1
	// MAV_RESULT_DENIED enum. Command is invalid (is supported but has invalid parameters). Retrying same command and parameters will not work
	MAV_RESULT_DENIED MAV_RESULT = 2
	// MAV_RESULT_UNSUPPORTED enum. Command is not supported (unknown)
	MAV_RESULT_UNSUPPORTED MAV_RESULT = 3
	// MAV_RESULT_FAILED enum. Command is valid, but execution has failed. This is used to indicate any non-temporary or unexpected problem, i.e. any problem that must be fixed before the command can succeed/be retried. For example, attempting to write a file when out of memory, attempting to arm when sensors are not calibrated, etc
	MAV_RESULT_FAILED MAV_RESULT = 4
	// MAV_RESULT_IN_PROGRESS enum. Command is valid and is being executed. This will be followed by further progress updates, i.e. the component may send further COMMAND_ACK messages with result MAV_RESULT_IN_PROGRESS (at a rate decided by the implementation), and must terminate by sending a COMMAND_ACK message with final result of the operation. The COMMAND_ACK.progress field can be used to indicate the progress of the operation
	MAV_RESULT_IN_PROGRESS MAV_RESULT = 5
	// MAV_RESULT_CANCELLED enum. Command has been cancelled (as a result of receiving a COMMAND_CANCEL message)
	MAV_RESULT_CANCELLED MAV_RESULT = 6
)

func (e MAV_RESULT) String() string {
	if name, ok := map[MAV_RESULT]string{
		MAV_RESULT_ACCEPTED:             "MAV_RESULT_ACCEPTED",
		MAV_RESULT_TEMPORARILY_REJECTED: "MAV_RESULT_TEMPORARILY_REJECTED",
		MAV_RESULT_DENIED:               "MAV_RESULT_DENIED",
		MAV_RESULT_UNSUPPORTED:          "MAV_RESULT_UNSUPPORTED",
		MAV_RESULT_FAILED:               "MAV_RESULT_FAILED",
		MAV_RESULT_IN_PROGRESS:          "MAV_RESULT_IN_PROGRESS",
		MAV_RESULT_CANCELLED:            "MAV_RESULT_CANCELLED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_RESULT_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_RESULT enums
func (e MAV_RESULT) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_RESULT{
		MAV_RESULT_ACCEPTED,
		MAV_RESULT_TEMPORARILY_REJECTED,
		MAV_RESULT_DENIED,
		MAV_RESULT_UNSUPPORTED,
		MAV_RESULT_FAILED,
		MAV_RESULT_IN_PROGRESS,
		MAV_RESULT_CANCELLED,
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

// MAV_MISSION_RESULT type. Result of mission operation (in a MISSION_ACK message).
type MAV_MISSION_RESULT int

func (e MAV_MISSION_RESULT) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_MISSION_ACCEPTED enum. mission accepted OK
	MAV_MISSION_ACCEPTED MAV_MISSION_RESULT = 0
	// MAV_MISSION_ERROR enum. Generic error / not accepting mission commands at all right now
	MAV_MISSION_ERROR MAV_MISSION_RESULT = 1
	// MAV_MISSION_UNSUPPORTED_FRAME enum. Coordinate frame is not supported
	MAV_MISSION_UNSUPPORTED_FRAME MAV_MISSION_RESULT = 2
	// MAV_MISSION_UNSUPPORTED enum. Command is not supported
	MAV_MISSION_UNSUPPORTED MAV_MISSION_RESULT = 3
	// MAV_MISSION_NO_SPACE enum. Mission items exceed storage space
	MAV_MISSION_NO_SPACE MAV_MISSION_RESULT = 4
	// MAV_MISSION_INVALID enum. One of the parameters has an invalid value
	MAV_MISSION_INVALID MAV_MISSION_RESULT = 5
	// MAV_MISSION_INVALID_PARAM1 enum. param1 has an invalid value
	MAV_MISSION_INVALID_PARAM1 MAV_MISSION_RESULT = 6
	// MAV_MISSION_INVALID_PARAM2 enum. param2 has an invalid value
	MAV_MISSION_INVALID_PARAM2 MAV_MISSION_RESULT = 7
	// MAV_MISSION_INVALID_PARAM3 enum. param3 has an invalid value
	MAV_MISSION_INVALID_PARAM3 MAV_MISSION_RESULT = 8
	// MAV_MISSION_INVALID_PARAM4 enum. param4 has an invalid value
	MAV_MISSION_INVALID_PARAM4 MAV_MISSION_RESULT = 9
	// MAV_MISSION_INVALID_PARAM5_X enum. x / param5 has an invalid value
	MAV_MISSION_INVALID_PARAM5_X MAV_MISSION_RESULT = 10
	// MAV_MISSION_INVALID_PARAM6_Y enum. y / param6 has an invalid value
	MAV_MISSION_INVALID_PARAM6_Y MAV_MISSION_RESULT = 11
	// MAV_MISSION_INVALID_PARAM7 enum. z / param7 has an invalid value
	MAV_MISSION_INVALID_PARAM7 MAV_MISSION_RESULT = 12
	// MAV_MISSION_INVALID_SEQUENCE enum. Mission item received out of sequence
	MAV_MISSION_INVALID_SEQUENCE MAV_MISSION_RESULT = 13
	// MAV_MISSION_DENIED enum. Not accepting any mission commands from this communication partner
	MAV_MISSION_DENIED MAV_MISSION_RESULT = 14
	// MAV_MISSION_OPERATION_CANCELLED enum. Current mission operation cancelled (e.g. mission upload, mission download)
	MAV_MISSION_OPERATION_CANCELLED MAV_MISSION_RESULT = 15
)

func (e MAV_MISSION_RESULT) String() string {
	if name, ok := map[MAV_MISSION_RESULT]string{
		MAV_MISSION_ACCEPTED:            "MAV_MISSION_ACCEPTED",
		MAV_MISSION_ERROR:               "MAV_MISSION_ERROR",
		MAV_MISSION_UNSUPPORTED_FRAME:   "MAV_MISSION_UNSUPPORTED_FRAME",
		MAV_MISSION_UNSUPPORTED:         "MAV_MISSION_UNSUPPORTED",
		MAV_MISSION_NO_SPACE:            "MAV_MISSION_NO_SPACE",
		MAV_MISSION_INVALID:             "MAV_MISSION_INVALID",
		MAV_MISSION_INVALID_PARAM1:      "MAV_MISSION_INVALID_PARAM1",
		MAV_MISSION_INVALID_PARAM2:      "MAV_MISSION_INVALID_PARAM2",
		MAV_MISSION_INVALID_PARAM3:      "MAV_MISSION_INVALID_PARAM3",
		MAV_MISSION_INVALID_PARAM4:      "MAV_MISSION_INVALID_PARAM4",
		MAV_MISSION_INVALID_PARAM5_X:    "MAV_MISSION_INVALID_PARAM5_X",
		MAV_MISSION_INVALID_PARAM6_Y:    "MAV_MISSION_INVALID_PARAM6_Y",
		MAV_MISSION_INVALID_PARAM7:      "MAV_MISSION_INVALID_PARAM7",
		MAV_MISSION_INVALID_SEQUENCE:    "MAV_MISSION_INVALID_SEQUENCE",
		MAV_MISSION_DENIED:              "MAV_MISSION_DENIED",
		MAV_MISSION_OPERATION_CANCELLED: "MAV_MISSION_OPERATION_CANCELLED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_MISSION_RESULT_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_MISSION_RESULT enums
func (e MAV_MISSION_RESULT) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_MISSION_RESULT{
		MAV_MISSION_ACCEPTED,
		MAV_MISSION_ERROR,
		MAV_MISSION_UNSUPPORTED_FRAME,
		MAV_MISSION_UNSUPPORTED,
		MAV_MISSION_NO_SPACE,
		MAV_MISSION_INVALID,
		MAV_MISSION_INVALID_PARAM1,
		MAV_MISSION_INVALID_PARAM2,
		MAV_MISSION_INVALID_PARAM3,
		MAV_MISSION_INVALID_PARAM4,
		MAV_MISSION_INVALID_PARAM5_X,
		MAV_MISSION_INVALID_PARAM6_Y,
		MAV_MISSION_INVALID_PARAM7,
		MAV_MISSION_INVALID_SEQUENCE,
		MAV_MISSION_DENIED,
		MAV_MISSION_OPERATION_CANCELLED,
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

// MAV_SEVERITY type. Indicates the severity level, generally used for status messages to indicate their relative urgency. Based on RFC-5424 using expanded definitions at: http://www.kiwisyslog.com/kb/info:-syslog-message-levels/.
type MAV_SEVERITY int

func (e MAV_SEVERITY) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_SEVERITY_EMERGENCY enum. System is unusable. This is a "panic" condition
	MAV_SEVERITY_EMERGENCY MAV_SEVERITY = 0
	// MAV_SEVERITY_ALERT enum. Action should be taken immediately. Indicates error in non-critical systems
	MAV_SEVERITY_ALERT MAV_SEVERITY = 1
	// MAV_SEVERITY_CRITICAL enum. Action must be taken immediately. Indicates failure in a primary system
	MAV_SEVERITY_CRITICAL MAV_SEVERITY = 2
	// MAV_SEVERITY_ERROR enum. Indicates an error in secondary/redundant systems
	MAV_SEVERITY_ERROR MAV_SEVERITY = 3
	// MAV_SEVERITY_WARNING enum. Indicates about a possible future error if this is not resolved within a given timeframe. Example would be a low battery warning
	MAV_SEVERITY_WARNING MAV_SEVERITY = 4
	// MAV_SEVERITY_NOTICE enum. An unusual event has occurred, though not an error condition. This should be investigated for the root cause
	MAV_SEVERITY_NOTICE MAV_SEVERITY = 5
	// MAV_SEVERITY_INFO enum. Normal operational messages. Useful for logging. No action is required for these messages
	MAV_SEVERITY_INFO MAV_SEVERITY = 6
	// MAV_SEVERITY_DEBUG enum. Useful non-operational messages that can assist in debugging. These should not occur during normal operation
	MAV_SEVERITY_DEBUG MAV_SEVERITY = 7
)

func (e MAV_SEVERITY) String() string {
	if name, ok := map[MAV_SEVERITY]string{
		MAV_SEVERITY_EMERGENCY: "MAV_SEVERITY_EMERGENCY",
		MAV_SEVERITY_ALERT:     "MAV_SEVERITY_ALERT",
		MAV_SEVERITY_CRITICAL:  "MAV_SEVERITY_CRITICAL",
		MAV_SEVERITY_ERROR:     "MAV_SEVERITY_ERROR",
		MAV_SEVERITY_WARNING:   "MAV_SEVERITY_WARNING",
		MAV_SEVERITY_NOTICE:    "MAV_SEVERITY_NOTICE",
		MAV_SEVERITY_INFO:      "MAV_SEVERITY_INFO",
		MAV_SEVERITY_DEBUG:     "MAV_SEVERITY_DEBUG",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_SEVERITY_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_SEVERITY enums
func (e MAV_SEVERITY) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_SEVERITY{
		MAV_SEVERITY_EMERGENCY,
		MAV_SEVERITY_ALERT,
		MAV_SEVERITY_CRITICAL,
		MAV_SEVERITY_ERROR,
		MAV_SEVERITY_WARNING,
		MAV_SEVERITY_NOTICE,
		MAV_SEVERITY_INFO,
		MAV_SEVERITY_DEBUG,
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

// MAV_POWER_STATUS type. Power supply status flags (bitmask)
type MAV_POWER_STATUS int

func (e MAV_POWER_STATUS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_POWER_STATUS_BRICK_VALID enum. main brick power supply valid
	MAV_POWER_STATUS_BRICK_VALID MAV_POWER_STATUS = 1
	// MAV_POWER_STATUS_SERVO_VALID enum. main servo power supply valid for FMU
	MAV_POWER_STATUS_SERVO_VALID MAV_POWER_STATUS = 2
	// MAV_POWER_STATUS_USB_CONNECTED enum. USB power is connected
	MAV_POWER_STATUS_USB_CONNECTED MAV_POWER_STATUS = 4
	// MAV_POWER_STATUS_PERIPH_OVERCURRENT enum. peripheral supply is in over-current state
	MAV_POWER_STATUS_PERIPH_OVERCURRENT MAV_POWER_STATUS = 8
	// MAV_POWER_STATUS_PERIPH_HIPOWER_OVERCURRENT enum. hi-power peripheral supply is in over-current state
	MAV_POWER_STATUS_PERIPH_HIPOWER_OVERCURRENT MAV_POWER_STATUS = 16
	// MAV_POWER_STATUS_CHANGED enum. Power status has changed since boot
	MAV_POWER_STATUS_CHANGED MAV_POWER_STATUS = 32
)

func (e MAV_POWER_STATUS) String() string {
	if name, ok := map[MAV_POWER_STATUS]string{
		MAV_POWER_STATUS_BRICK_VALID:                "MAV_POWER_STATUS_BRICK_VALID",
		MAV_POWER_STATUS_SERVO_VALID:                "MAV_POWER_STATUS_SERVO_VALID",
		MAV_POWER_STATUS_USB_CONNECTED:              "MAV_POWER_STATUS_USB_CONNECTED",
		MAV_POWER_STATUS_PERIPH_OVERCURRENT:         "MAV_POWER_STATUS_PERIPH_OVERCURRENT",
		MAV_POWER_STATUS_PERIPH_HIPOWER_OVERCURRENT: "MAV_POWER_STATUS_PERIPH_HIPOWER_OVERCURRENT",
		MAV_POWER_STATUS_CHANGED:                    "MAV_POWER_STATUS_CHANGED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_POWER_STATUS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_POWER_STATUS enums
func (e MAV_POWER_STATUS) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_POWER_STATUS{
		MAV_POWER_STATUS_BRICK_VALID,
		MAV_POWER_STATUS_SERVO_VALID,
		MAV_POWER_STATUS_USB_CONNECTED,
		MAV_POWER_STATUS_PERIPH_OVERCURRENT,
		MAV_POWER_STATUS_PERIPH_HIPOWER_OVERCURRENT,
		MAV_POWER_STATUS_CHANGED,
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

// SERIAL_CONTROL_DEV type. SERIAL_CONTROL device types
type SERIAL_CONTROL_DEV int

func (e SERIAL_CONTROL_DEV) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// SERIAL_CONTROL_DEV_TELEM1 enum. First telemetry port
	SERIAL_CONTROL_DEV_TELEM1 SERIAL_CONTROL_DEV = 0
	// SERIAL_CONTROL_DEV_TELEM2 enum. Second telemetry port
	SERIAL_CONTROL_DEV_TELEM2 SERIAL_CONTROL_DEV = 1
	// SERIAL_CONTROL_DEV_GPS1 enum. First GPS port
	SERIAL_CONTROL_DEV_GPS1 SERIAL_CONTROL_DEV = 2
	// SERIAL_CONTROL_DEV_GPS2 enum. Second GPS port
	SERIAL_CONTROL_DEV_GPS2 SERIAL_CONTROL_DEV = 3
	// SERIAL_CONTROL_DEV_SHELL enum. system shell
	SERIAL_CONTROL_DEV_SHELL SERIAL_CONTROL_DEV = 10
	// SERIAL_CONTROL_SERIAL0 enum. SERIAL0
	SERIAL_CONTROL_SERIAL0 SERIAL_CONTROL_DEV = 100
	// SERIAL_CONTROL_SERIAL1 enum. SERIAL1
	SERIAL_CONTROL_SERIAL1 SERIAL_CONTROL_DEV = 101
	// SERIAL_CONTROL_SERIAL2 enum. SERIAL2
	SERIAL_CONTROL_SERIAL2 SERIAL_CONTROL_DEV = 102
	// SERIAL_CONTROL_SERIAL3 enum. SERIAL3
	SERIAL_CONTROL_SERIAL3 SERIAL_CONTROL_DEV = 103
	// SERIAL_CONTROL_SERIAL4 enum. SERIAL4
	SERIAL_CONTROL_SERIAL4 SERIAL_CONTROL_DEV = 104
	// SERIAL_CONTROL_SERIAL5 enum. SERIAL5
	SERIAL_CONTROL_SERIAL5 SERIAL_CONTROL_DEV = 105
	// SERIAL_CONTROL_SERIAL6 enum. SERIAL6
	SERIAL_CONTROL_SERIAL6 SERIAL_CONTROL_DEV = 106
	// SERIAL_CONTROL_SERIAL7 enum. SERIAL7
	SERIAL_CONTROL_SERIAL7 SERIAL_CONTROL_DEV = 107
	// SERIAL_CONTROL_SERIAL8 enum. SERIAL8
	SERIAL_CONTROL_SERIAL8 SERIAL_CONTROL_DEV = 108
	// SERIAL_CONTROL_SERIAL9 enum. SERIAL9
	SERIAL_CONTROL_SERIAL9 SERIAL_CONTROL_DEV = 109
)

func (e SERIAL_CONTROL_DEV) String() string {
	if name, ok := map[SERIAL_CONTROL_DEV]string{
		SERIAL_CONTROL_DEV_TELEM1: "SERIAL_CONTROL_DEV_TELEM1",
		SERIAL_CONTROL_DEV_TELEM2: "SERIAL_CONTROL_DEV_TELEM2",
		SERIAL_CONTROL_DEV_GPS1:   "SERIAL_CONTROL_DEV_GPS1",
		SERIAL_CONTROL_DEV_GPS2:   "SERIAL_CONTROL_DEV_GPS2",
		SERIAL_CONTROL_DEV_SHELL:  "SERIAL_CONTROL_DEV_SHELL",
		SERIAL_CONTROL_SERIAL0:    "SERIAL_CONTROL_SERIAL0",
		SERIAL_CONTROL_SERIAL1:    "SERIAL_CONTROL_SERIAL1",
		SERIAL_CONTROL_SERIAL2:    "SERIAL_CONTROL_SERIAL2",
		SERIAL_CONTROL_SERIAL3:    "SERIAL_CONTROL_SERIAL3",
		SERIAL_CONTROL_SERIAL4:    "SERIAL_CONTROL_SERIAL4",
		SERIAL_CONTROL_SERIAL5:    "SERIAL_CONTROL_SERIAL5",
		SERIAL_CONTROL_SERIAL6:    "SERIAL_CONTROL_SERIAL6",
		SERIAL_CONTROL_SERIAL7:    "SERIAL_CONTROL_SERIAL7",
		SERIAL_CONTROL_SERIAL8:    "SERIAL_CONTROL_SERIAL8",
		SERIAL_CONTROL_SERIAL9:    "SERIAL_CONTROL_SERIAL9",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("SERIAL_CONTROL_DEV_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects SERIAL_CONTROL_DEV enums
func (e SERIAL_CONTROL_DEV) Bitmask() string {
	bitmap := ""
	for _, entry := range []SERIAL_CONTROL_DEV{
		SERIAL_CONTROL_DEV_TELEM1,
		SERIAL_CONTROL_DEV_TELEM2,
		SERIAL_CONTROL_DEV_GPS1,
		SERIAL_CONTROL_DEV_GPS2,
		SERIAL_CONTROL_DEV_SHELL,
		SERIAL_CONTROL_SERIAL0,
		SERIAL_CONTROL_SERIAL1,
		SERIAL_CONTROL_SERIAL2,
		SERIAL_CONTROL_SERIAL3,
		SERIAL_CONTROL_SERIAL4,
		SERIAL_CONTROL_SERIAL5,
		SERIAL_CONTROL_SERIAL6,
		SERIAL_CONTROL_SERIAL7,
		SERIAL_CONTROL_SERIAL8,
		SERIAL_CONTROL_SERIAL9,
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

// SERIAL_CONTROL_FLAG type. SERIAL_CONTROL flags (bitmask)
type SERIAL_CONTROL_FLAG int

func (e SERIAL_CONTROL_FLAG) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// SERIAL_CONTROL_FLAG_REPLY enum. Set if this is a reply
	SERIAL_CONTROL_FLAG_REPLY SERIAL_CONTROL_FLAG = 1
	// SERIAL_CONTROL_FLAG_RESPOND enum. Set if the sender wants the receiver to send a response as another SERIAL_CONTROL message
	SERIAL_CONTROL_FLAG_RESPOND SERIAL_CONTROL_FLAG = 2
	// SERIAL_CONTROL_FLAG_EXCLUSIVE enum. Set if access to the serial port should be removed from whatever driver is currently using it, giving exclusive access to the SERIAL_CONTROL protocol. The port can be handed back by sending a request without this flag set
	SERIAL_CONTROL_FLAG_EXCLUSIVE SERIAL_CONTROL_FLAG = 4
	// SERIAL_CONTROL_FLAG_BLOCKING enum. Block on writes to the serial port
	SERIAL_CONTROL_FLAG_BLOCKING SERIAL_CONTROL_FLAG = 8
	// SERIAL_CONTROL_FLAG_MULTI enum. Send multiple replies until port is drained
	SERIAL_CONTROL_FLAG_MULTI SERIAL_CONTROL_FLAG = 16
)

func (e SERIAL_CONTROL_FLAG) String() string {
	if name, ok := map[SERIAL_CONTROL_FLAG]string{
		SERIAL_CONTROL_FLAG_REPLY:     "SERIAL_CONTROL_FLAG_REPLY",
		SERIAL_CONTROL_FLAG_RESPOND:   "SERIAL_CONTROL_FLAG_RESPOND",
		SERIAL_CONTROL_FLAG_EXCLUSIVE: "SERIAL_CONTROL_FLAG_EXCLUSIVE",
		SERIAL_CONTROL_FLAG_BLOCKING:  "SERIAL_CONTROL_FLAG_BLOCKING",
		SERIAL_CONTROL_FLAG_MULTI:     "SERIAL_CONTROL_FLAG_MULTI",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("SERIAL_CONTROL_FLAG_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects SERIAL_CONTROL_FLAG enums
func (e SERIAL_CONTROL_FLAG) Bitmask() string {
	bitmap := ""
	for _, entry := range []SERIAL_CONTROL_FLAG{
		SERIAL_CONTROL_FLAG_REPLY,
		SERIAL_CONTROL_FLAG_RESPOND,
		SERIAL_CONTROL_FLAG_EXCLUSIVE,
		SERIAL_CONTROL_FLAG_BLOCKING,
		SERIAL_CONTROL_FLAG_MULTI,
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

// MAV_DISTANCE_SENSOR type. Enumeration of distance sensor types
type MAV_DISTANCE_SENSOR int

func (e MAV_DISTANCE_SENSOR) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_DISTANCE_SENSOR_LASER enum. Laser rangefinder, e.g. LightWare SF02/F or PulsedLight units
	MAV_DISTANCE_SENSOR_LASER MAV_DISTANCE_SENSOR = 0
	// MAV_DISTANCE_SENSOR_ULTRASOUND enum. Ultrasound rangefinder, e.g. MaxBotix units
	MAV_DISTANCE_SENSOR_ULTRASOUND MAV_DISTANCE_SENSOR = 1
	// MAV_DISTANCE_SENSOR_INFRARED enum. Infrared rangefinder, e.g. Sharp units
	MAV_DISTANCE_SENSOR_INFRARED MAV_DISTANCE_SENSOR = 2
	// MAV_DISTANCE_SENSOR_RADAR enum. Radar type, e.g. uLanding units
	MAV_DISTANCE_SENSOR_RADAR MAV_DISTANCE_SENSOR = 3
	// MAV_DISTANCE_SENSOR_UNKNOWN enum. Broken or unknown type, e.g. analog units
	MAV_DISTANCE_SENSOR_UNKNOWN MAV_DISTANCE_SENSOR = 4
)

func (e MAV_DISTANCE_SENSOR) String() string {
	if name, ok := map[MAV_DISTANCE_SENSOR]string{
		MAV_DISTANCE_SENSOR_LASER:      "MAV_DISTANCE_SENSOR_LASER",
		MAV_DISTANCE_SENSOR_ULTRASOUND: "MAV_DISTANCE_SENSOR_ULTRASOUND",
		MAV_DISTANCE_SENSOR_INFRARED:   "MAV_DISTANCE_SENSOR_INFRARED",
		MAV_DISTANCE_SENSOR_RADAR:      "MAV_DISTANCE_SENSOR_RADAR",
		MAV_DISTANCE_SENSOR_UNKNOWN:    "MAV_DISTANCE_SENSOR_UNKNOWN",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_DISTANCE_SENSOR_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_DISTANCE_SENSOR enums
func (e MAV_DISTANCE_SENSOR) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_DISTANCE_SENSOR{
		MAV_DISTANCE_SENSOR_LASER,
		MAV_DISTANCE_SENSOR_ULTRASOUND,
		MAV_DISTANCE_SENSOR_INFRARED,
		MAV_DISTANCE_SENSOR_RADAR,
		MAV_DISTANCE_SENSOR_UNKNOWN,
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

// MAV_SENSOR_ORIENTATION type. Enumeration of sensor orientation, according to its rotations
type MAV_SENSOR_ORIENTATION int

func (e MAV_SENSOR_ORIENTATION) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_SENSOR_ROTATION_NONE enum. Roll: 0, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_NONE MAV_SENSOR_ORIENTATION = 0
	// MAV_SENSOR_ROTATION_YAW_45 enum. Roll: 0, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_YAW_45 MAV_SENSOR_ORIENTATION = 1
	// MAV_SENSOR_ROTATION_YAW_90 enum. Roll: 0, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_YAW_90 MAV_SENSOR_ORIENTATION = 2
	// MAV_SENSOR_ROTATION_YAW_135 enum. Roll: 0, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_YAW_135 MAV_SENSOR_ORIENTATION = 3
	// MAV_SENSOR_ROTATION_YAW_180 enum. Roll: 0, Pitch: 0, Yaw: 180
	MAV_SENSOR_ROTATION_YAW_180 MAV_SENSOR_ORIENTATION = 4
	// MAV_SENSOR_ROTATION_YAW_225 enum. Roll: 0, Pitch: 0, Yaw: 225
	MAV_SENSOR_ROTATION_YAW_225 MAV_SENSOR_ORIENTATION = 5
	// MAV_SENSOR_ROTATION_YAW_270 enum. Roll: 0, Pitch: 0, Yaw: 270
	MAV_SENSOR_ROTATION_YAW_270 MAV_SENSOR_ORIENTATION = 6
	// MAV_SENSOR_ROTATION_YAW_315 enum. Roll: 0, Pitch: 0, Yaw: 315
	MAV_SENSOR_ROTATION_YAW_315 MAV_SENSOR_ORIENTATION = 7
	// MAV_SENSOR_ROTATION_ROLL_180 enum. Roll: 180, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_180 MAV_SENSOR_ORIENTATION = 8
	// MAV_SENSOR_ROTATION_ROLL_180_YAW_45 enum. Roll: 180, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_ROLL_180_YAW_45 MAV_SENSOR_ORIENTATION = 9
	// MAV_SENSOR_ROTATION_ROLL_180_YAW_90 enum. Roll: 180, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_180_YAW_90 MAV_SENSOR_ORIENTATION = 10
	// MAV_SENSOR_ROTATION_ROLL_180_YAW_135 enum. Roll: 180, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_ROLL_180_YAW_135 MAV_SENSOR_ORIENTATION = 11
	// MAV_SENSOR_ROTATION_PITCH_180 enum. Roll: 0, Pitch: 180, Yaw: 0
	MAV_SENSOR_ROTATION_PITCH_180 MAV_SENSOR_ORIENTATION = 12
	// MAV_SENSOR_ROTATION_ROLL_180_YAW_225 enum. Roll: 180, Pitch: 0, Yaw: 225
	MAV_SENSOR_ROTATION_ROLL_180_YAW_225 MAV_SENSOR_ORIENTATION = 13
	// MAV_SENSOR_ROTATION_ROLL_180_YAW_270 enum. Roll: 180, Pitch: 0, Yaw: 270
	MAV_SENSOR_ROTATION_ROLL_180_YAW_270 MAV_SENSOR_ORIENTATION = 14
	// MAV_SENSOR_ROTATION_ROLL_180_YAW_315 enum. Roll: 180, Pitch: 0, Yaw: 315
	MAV_SENSOR_ROTATION_ROLL_180_YAW_315 MAV_SENSOR_ORIENTATION = 15
	// MAV_SENSOR_ROTATION_ROLL_90 enum. Roll: 90, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90 MAV_SENSOR_ORIENTATION = 16
	// MAV_SENSOR_ROTATION_ROLL_90_YAW_45 enum. Roll: 90, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_ROLL_90_YAW_45 MAV_SENSOR_ORIENTATION = 17
	// MAV_SENSOR_ROTATION_ROLL_90_YAW_90 enum. Roll: 90, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_90_YAW_90 MAV_SENSOR_ORIENTATION = 18
	// MAV_SENSOR_ROTATION_ROLL_90_YAW_135 enum. Roll: 90, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_ROLL_90_YAW_135 MAV_SENSOR_ORIENTATION = 19
	// MAV_SENSOR_ROTATION_ROLL_270 enum. Roll: 270, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270 MAV_SENSOR_ORIENTATION = 20
	// MAV_SENSOR_ROTATION_ROLL_270_YAW_45 enum. Roll: 270, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_ROLL_270_YAW_45 MAV_SENSOR_ORIENTATION = 21
	// MAV_SENSOR_ROTATION_ROLL_270_YAW_90 enum. Roll: 270, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_270_YAW_90 MAV_SENSOR_ORIENTATION = 22
	// MAV_SENSOR_ROTATION_ROLL_270_YAW_135 enum. Roll: 270, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_ROLL_270_YAW_135 MAV_SENSOR_ORIENTATION = 23
	// MAV_SENSOR_ROTATION_PITCH_90 enum. Roll: 0, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_PITCH_90 MAV_SENSOR_ORIENTATION = 24
	// MAV_SENSOR_ROTATION_PITCH_270 enum. Roll: 0, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_PITCH_270 MAV_SENSOR_ORIENTATION = 25
	// MAV_SENSOR_ROTATION_PITCH_180_YAW_90 enum. Roll: 0, Pitch: 180, Yaw: 90
	MAV_SENSOR_ROTATION_PITCH_180_YAW_90 MAV_SENSOR_ORIENTATION = 26
	// MAV_SENSOR_ROTATION_PITCH_180_YAW_270 enum. Roll: 0, Pitch: 180, Yaw: 270
	MAV_SENSOR_ROTATION_PITCH_180_YAW_270 MAV_SENSOR_ORIENTATION = 27
	// MAV_SENSOR_ROTATION_ROLL_90_PITCH_90 enum. Roll: 90, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_90 MAV_SENSOR_ORIENTATION = 28
	// MAV_SENSOR_ROTATION_ROLL_180_PITCH_90 enum. Roll: 180, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_180_PITCH_90 MAV_SENSOR_ORIENTATION = 29
	// MAV_SENSOR_ROTATION_ROLL_270_PITCH_90 enum. Roll: 270, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270_PITCH_90 MAV_SENSOR_ORIENTATION = 30
	// MAV_SENSOR_ROTATION_ROLL_90_PITCH_180 enum. Roll: 90, Pitch: 180, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_180 MAV_SENSOR_ORIENTATION = 31
	// MAV_SENSOR_ROTATION_ROLL_270_PITCH_180 enum. Roll: 270, Pitch: 180, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270_PITCH_180 MAV_SENSOR_ORIENTATION = 32
	// MAV_SENSOR_ROTATION_ROLL_90_PITCH_270 enum. Roll: 90, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_270 MAV_SENSOR_ORIENTATION = 33
	// MAV_SENSOR_ROTATION_ROLL_180_PITCH_270 enum. Roll: 180, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_180_PITCH_270 MAV_SENSOR_ORIENTATION = 34
	// MAV_SENSOR_ROTATION_ROLL_270_PITCH_270 enum. Roll: 270, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270_PITCH_270 MAV_SENSOR_ORIENTATION = 35
	// MAV_SENSOR_ROTATION_ROLL_90_PITCH_180_YAW_90 enum. Roll: 90, Pitch: 180, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_180_YAW_90 MAV_SENSOR_ORIENTATION = 36
	// MAV_SENSOR_ROTATION_ROLL_90_YAW_270 enum. Roll: 90, Pitch: 0, Yaw: 270
	MAV_SENSOR_ROTATION_ROLL_90_YAW_270 MAV_SENSOR_ORIENTATION = 37
	// MAV_SENSOR_ROTATION_ROLL_90_PITCH_68_YAW_293 enum. Roll: 90, Pitch: 68, Yaw: 293
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_68_YAW_293 MAV_SENSOR_ORIENTATION = 38
	// MAV_SENSOR_ROTATION_PITCH_315 enum. Pitch: 315
	MAV_SENSOR_ROTATION_PITCH_315 MAV_SENSOR_ORIENTATION = 39
	// MAV_SENSOR_ROTATION_ROLL_90_PITCH_315 enum. Roll: 90, Pitch: 315
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_315 MAV_SENSOR_ORIENTATION = 40
	// MAV_SENSOR_ROTATION_CUSTOM enum. Custom orientation
	MAV_SENSOR_ROTATION_CUSTOM MAV_SENSOR_ORIENTATION = 100
)

func (e MAV_SENSOR_ORIENTATION) String() string {
	if name, ok := map[MAV_SENSOR_ORIENTATION]string{
		MAV_SENSOR_ROTATION_NONE:                     "MAV_SENSOR_ROTATION_NONE",
		MAV_SENSOR_ROTATION_YAW_45:                   "MAV_SENSOR_ROTATION_YAW_45",
		MAV_SENSOR_ROTATION_YAW_90:                   "MAV_SENSOR_ROTATION_YAW_90",
		MAV_SENSOR_ROTATION_YAW_135:                  "MAV_SENSOR_ROTATION_YAW_135",
		MAV_SENSOR_ROTATION_YAW_180:                  "MAV_SENSOR_ROTATION_YAW_180",
		MAV_SENSOR_ROTATION_YAW_225:                  "MAV_SENSOR_ROTATION_YAW_225",
		MAV_SENSOR_ROTATION_YAW_270:                  "MAV_SENSOR_ROTATION_YAW_270",
		MAV_SENSOR_ROTATION_YAW_315:                  "MAV_SENSOR_ROTATION_YAW_315",
		MAV_SENSOR_ROTATION_ROLL_180:                 "MAV_SENSOR_ROTATION_ROLL_180",
		MAV_SENSOR_ROTATION_ROLL_180_YAW_45:          "MAV_SENSOR_ROTATION_ROLL_180_YAW_45",
		MAV_SENSOR_ROTATION_ROLL_180_YAW_90:          "MAV_SENSOR_ROTATION_ROLL_180_YAW_90",
		MAV_SENSOR_ROTATION_ROLL_180_YAW_135:         "MAV_SENSOR_ROTATION_ROLL_180_YAW_135",
		MAV_SENSOR_ROTATION_PITCH_180:                "MAV_SENSOR_ROTATION_PITCH_180",
		MAV_SENSOR_ROTATION_ROLL_180_YAW_225:         "MAV_SENSOR_ROTATION_ROLL_180_YAW_225",
		MAV_SENSOR_ROTATION_ROLL_180_YAW_270:         "MAV_SENSOR_ROTATION_ROLL_180_YAW_270",
		MAV_SENSOR_ROTATION_ROLL_180_YAW_315:         "MAV_SENSOR_ROTATION_ROLL_180_YAW_315",
		MAV_SENSOR_ROTATION_ROLL_90:                  "MAV_SENSOR_ROTATION_ROLL_90",
		MAV_SENSOR_ROTATION_ROLL_90_YAW_45:           "MAV_SENSOR_ROTATION_ROLL_90_YAW_45",
		MAV_SENSOR_ROTATION_ROLL_90_YAW_90:           "MAV_SENSOR_ROTATION_ROLL_90_YAW_90",
		MAV_SENSOR_ROTATION_ROLL_90_YAW_135:          "MAV_SENSOR_ROTATION_ROLL_90_YAW_135",
		MAV_SENSOR_ROTATION_ROLL_270:                 "MAV_SENSOR_ROTATION_ROLL_270",
		MAV_SENSOR_ROTATION_ROLL_270_YAW_45:          "MAV_SENSOR_ROTATION_ROLL_270_YAW_45",
		MAV_SENSOR_ROTATION_ROLL_270_YAW_90:          "MAV_SENSOR_ROTATION_ROLL_270_YAW_90",
		MAV_SENSOR_ROTATION_ROLL_270_YAW_135:         "MAV_SENSOR_ROTATION_ROLL_270_YAW_135",
		MAV_SENSOR_ROTATION_PITCH_90:                 "MAV_SENSOR_ROTATION_PITCH_90",
		MAV_SENSOR_ROTATION_PITCH_270:                "MAV_SENSOR_ROTATION_PITCH_270",
		MAV_SENSOR_ROTATION_PITCH_180_YAW_90:         "MAV_SENSOR_ROTATION_PITCH_180_YAW_90",
		MAV_SENSOR_ROTATION_PITCH_180_YAW_270:        "MAV_SENSOR_ROTATION_PITCH_180_YAW_270",
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_90:         "MAV_SENSOR_ROTATION_ROLL_90_PITCH_90",
		MAV_SENSOR_ROTATION_ROLL_180_PITCH_90:        "MAV_SENSOR_ROTATION_ROLL_180_PITCH_90",
		MAV_SENSOR_ROTATION_ROLL_270_PITCH_90:        "MAV_SENSOR_ROTATION_ROLL_270_PITCH_90",
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_180:        "MAV_SENSOR_ROTATION_ROLL_90_PITCH_180",
		MAV_SENSOR_ROTATION_ROLL_270_PITCH_180:       "MAV_SENSOR_ROTATION_ROLL_270_PITCH_180",
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_270:        "MAV_SENSOR_ROTATION_ROLL_90_PITCH_270",
		MAV_SENSOR_ROTATION_ROLL_180_PITCH_270:       "MAV_SENSOR_ROTATION_ROLL_180_PITCH_270",
		MAV_SENSOR_ROTATION_ROLL_270_PITCH_270:       "MAV_SENSOR_ROTATION_ROLL_270_PITCH_270",
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_180_YAW_90: "MAV_SENSOR_ROTATION_ROLL_90_PITCH_180_YAW_90",
		MAV_SENSOR_ROTATION_ROLL_90_YAW_270:          "MAV_SENSOR_ROTATION_ROLL_90_YAW_270",
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_68_YAW_293: "MAV_SENSOR_ROTATION_ROLL_90_PITCH_68_YAW_293",
		MAV_SENSOR_ROTATION_PITCH_315:                "MAV_SENSOR_ROTATION_PITCH_315",
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_315:        "MAV_SENSOR_ROTATION_ROLL_90_PITCH_315",
		MAV_SENSOR_ROTATION_CUSTOM:                   "MAV_SENSOR_ROTATION_CUSTOM",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_SENSOR_ORIENTATION_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_SENSOR_ORIENTATION enums
func (e MAV_SENSOR_ORIENTATION) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_SENSOR_ORIENTATION{
		MAV_SENSOR_ROTATION_NONE,
		MAV_SENSOR_ROTATION_YAW_45,
		MAV_SENSOR_ROTATION_YAW_90,
		MAV_SENSOR_ROTATION_YAW_135,
		MAV_SENSOR_ROTATION_YAW_180,
		MAV_SENSOR_ROTATION_YAW_225,
		MAV_SENSOR_ROTATION_YAW_270,
		MAV_SENSOR_ROTATION_YAW_315,
		MAV_SENSOR_ROTATION_ROLL_180,
		MAV_SENSOR_ROTATION_ROLL_180_YAW_45,
		MAV_SENSOR_ROTATION_ROLL_180_YAW_90,
		MAV_SENSOR_ROTATION_ROLL_180_YAW_135,
		MAV_SENSOR_ROTATION_PITCH_180,
		MAV_SENSOR_ROTATION_ROLL_180_YAW_225,
		MAV_SENSOR_ROTATION_ROLL_180_YAW_270,
		MAV_SENSOR_ROTATION_ROLL_180_YAW_315,
		MAV_SENSOR_ROTATION_ROLL_90,
		MAV_SENSOR_ROTATION_ROLL_90_YAW_45,
		MAV_SENSOR_ROTATION_ROLL_90_YAW_90,
		MAV_SENSOR_ROTATION_ROLL_90_YAW_135,
		MAV_SENSOR_ROTATION_ROLL_270,
		MAV_SENSOR_ROTATION_ROLL_270_YAW_45,
		MAV_SENSOR_ROTATION_ROLL_270_YAW_90,
		MAV_SENSOR_ROTATION_ROLL_270_YAW_135,
		MAV_SENSOR_ROTATION_PITCH_90,
		MAV_SENSOR_ROTATION_PITCH_270,
		MAV_SENSOR_ROTATION_PITCH_180_YAW_90,
		MAV_SENSOR_ROTATION_PITCH_180_YAW_270,
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_90,
		MAV_SENSOR_ROTATION_ROLL_180_PITCH_90,
		MAV_SENSOR_ROTATION_ROLL_270_PITCH_90,
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_180,
		MAV_SENSOR_ROTATION_ROLL_270_PITCH_180,
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_270,
		MAV_SENSOR_ROTATION_ROLL_180_PITCH_270,
		MAV_SENSOR_ROTATION_ROLL_270_PITCH_270,
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_180_YAW_90,
		MAV_SENSOR_ROTATION_ROLL_90_YAW_270,
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_68_YAW_293,
		MAV_SENSOR_ROTATION_PITCH_315,
		MAV_SENSOR_ROTATION_ROLL_90_PITCH_315,
		MAV_SENSOR_ROTATION_CUSTOM,
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

// MAV_PROTOCOL_CAPABILITY type. Bitmask of (optional) autopilot capabilities (64 bit). If a bit is set, the autopilot supports this capability.
type MAV_PROTOCOL_CAPABILITY int

func (e MAV_PROTOCOL_CAPABILITY) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_PROTOCOL_CAPABILITY_MISSION_FLOAT enum. Autopilot supports MISSION float message type
	MAV_PROTOCOL_CAPABILITY_MISSION_FLOAT MAV_PROTOCOL_CAPABILITY = 1
	// MAV_PROTOCOL_CAPABILITY_PARAM_FLOAT enum. Autopilot supports the new param float message type
	MAV_PROTOCOL_CAPABILITY_PARAM_FLOAT MAV_PROTOCOL_CAPABILITY = 2
	// MAV_PROTOCOL_CAPABILITY_MISSION_INT enum. Autopilot supports MISSION_ITEM_INT scaled integer message type
	MAV_PROTOCOL_CAPABILITY_MISSION_INT MAV_PROTOCOL_CAPABILITY = 4
	// MAV_PROTOCOL_CAPABILITY_COMMAND_INT enum. Autopilot supports COMMAND_INT scaled integer message type
	MAV_PROTOCOL_CAPABILITY_COMMAND_INT MAV_PROTOCOL_CAPABILITY = 8
	// MAV_PROTOCOL_CAPABILITY_PARAM_UNION enum. Autopilot supports the new param union message type
	MAV_PROTOCOL_CAPABILITY_PARAM_UNION MAV_PROTOCOL_CAPABILITY = 16
	// MAV_PROTOCOL_CAPABILITY_FTP enum. Autopilot supports the new FILE_TRANSFER_PROTOCOL message type
	MAV_PROTOCOL_CAPABILITY_FTP MAV_PROTOCOL_CAPABILITY = 32
	// MAV_PROTOCOL_CAPABILITY_SET_ATTITUDE_TARGET enum. Autopilot supports commanding attitude offboard
	MAV_PROTOCOL_CAPABILITY_SET_ATTITUDE_TARGET MAV_PROTOCOL_CAPABILITY = 64
	// MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_LOCAL_NED enum. Autopilot supports commanding position and velocity targets in local NED frame
	MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_LOCAL_NED MAV_PROTOCOL_CAPABILITY = 128
	// MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_GLOBAL_INT enum. Autopilot supports commanding position and velocity targets in global scaled integers
	MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_GLOBAL_INT MAV_PROTOCOL_CAPABILITY = 256
	// MAV_PROTOCOL_CAPABILITY_TERRAIN enum. Autopilot supports terrain protocol / data handling
	MAV_PROTOCOL_CAPABILITY_TERRAIN MAV_PROTOCOL_CAPABILITY = 512
	// MAV_PROTOCOL_CAPABILITY_SET_ACTUATOR_TARGET enum. Autopilot supports direct actuator control
	MAV_PROTOCOL_CAPABILITY_SET_ACTUATOR_TARGET MAV_PROTOCOL_CAPABILITY = 1024
	// MAV_PROTOCOL_CAPABILITY_FLIGHT_TERMINATION enum. Autopilot supports the flight termination command
	MAV_PROTOCOL_CAPABILITY_FLIGHT_TERMINATION MAV_PROTOCOL_CAPABILITY = 2048
	// MAV_PROTOCOL_CAPABILITY_COMPASS_CALIBRATION enum. Autopilot supports onboard compass calibration
	MAV_PROTOCOL_CAPABILITY_COMPASS_CALIBRATION MAV_PROTOCOL_CAPABILITY = 4096
	// MAV_PROTOCOL_CAPABILITY_MAVLINK2 enum. Autopilot supports MAVLink version 2
	MAV_PROTOCOL_CAPABILITY_MAVLINK2 MAV_PROTOCOL_CAPABILITY = 8192
	// MAV_PROTOCOL_CAPABILITY_MISSION_FENCE enum. Autopilot supports mission fence protocol
	MAV_PROTOCOL_CAPABILITY_MISSION_FENCE MAV_PROTOCOL_CAPABILITY = 16384
	// MAV_PROTOCOL_CAPABILITY_MISSION_RALLY enum. Autopilot supports mission rally point protocol
	MAV_PROTOCOL_CAPABILITY_MISSION_RALLY MAV_PROTOCOL_CAPABILITY = 32768
	// MAV_PROTOCOL_CAPABILITY_FLIGHT_INFORMATION enum. Autopilot supports the flight information protocol
	MAV_PROTOCOL_CAPABILITY_FLIGHT_INFORMATION MAV_PROTOCOL_CAPABILITY = 65536
)

func (e MAV_PROTOCOL_CAPABILITY) String() string {
	if name, ok := map[MAV_PROTOCOL_CAPABILITY]string{
		MAV_PROTOCOL_CAPABILITY_MISSION_FLOAT:                  "MAV_PROTOCOL_CAPABILITY_MISSION_FLOAT",
		MAV_PROTOCOL_CAPABILITY_PARAM_FLOAT:                    "MAV_PROTOCOL_CAPABILITY_PARAM_FLOAT",
		MAV_PROTOCOL_CAPABILITY_MISSION_INT:                    "MAV_PROTOCOL_CAPABILITY_MISSION_INT",
		MAV_PROTOCOL_CAPABILITY_COMMAND_INT:                    "MAV_PROTOCOL_CAPABILITY_COMMAND_INT",
		MAV_PROTOCOL_CAPABILITY_PARAM_UNION:                    "MAV_PROTOCOL_CAPABILITY_PARAM_UNION",
		MAV_PROTOCOL_CAPABILITY_FTP:                            "MAV_PROTOCOL_CAPABILITY_FTP",
		MAV_PROTOCOL_CAPABILITY_SET_ATTITUDE_TARGET:            "MAV_PROTOCOL_CAPABILITY_SET_ATTITUDE_TARGET",
		MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_LOCAL_NED:  "MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_LOCAL_NED",
		MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_GLOBAL_INT: "MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_GLOBAL_INT",
		MAV_PROTOCOL_CAPABILITY_TERRAIN:                        "MAV_PROTOCOL_CAPABILITY_TERRAIN",
		MAV_PROTOCOL_CAPABILITY_SET_ACTUATOR_TARGET:            "MAV_PROTOCOL_CAPABILITY_SET_ACTUATOR_TARGET",
		MAV_PROTOCOL_CAPABILITY_FLIGHT_TERMINATION:             "MAV_PROTOCOL_CAPABILITY_FLIGHT_TERMINATION",
		MAV_PROTOCOL_CAPABILITY_COMPASS_CALIBRATION:            "MAV_PROTOCOL_CAPABILITY_COMPASS_CALIBRATION",
		MAV_PROTOCOL_CAPABILITY_MAVLINK2:                       "MAV_PROTOCOL_CAPABILITY_MAVLINK2",
		MAV_PROTOCOL_CAPABILITY_MISSION_FENCE:                  "MAV_PROTOCOL_CAPABILITY_MISSION_FENCE",
		MAV_PROTOCOL_CAPABILITY_MISSION_RALLY:                  "MAV_PROTOCOL_CAPABILITY_MISSION_RALLY",
		MAV_PROTOCOL_CAPABILITY_FLIGHT_INFORMATION:             "MAV_PROTOCOL_CAPABILITY_FLIGHT_INFORMATION",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_PROTOCOL_CAPABILITY_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_PROTOCOL_CAPABILITY enums
func (e MAV_PROTOCOL_CAPABILITY) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_PROTOCOL_CAPABILITY{
		MAV_PROTOCOL_CAPABILITY_MISSION_FLOAT,
		MAV_PROTOCOL_CAPABILITY_PARAM_FLOAT,
		MAV_PROTOCOL_CAPABILITY_MISSION_INT,
		MAV_PROTOCOL_CAPABILITY_COMMAND_INT,
		MAV_PROTOCOL_CAPABILITY_PARAM_UNION,
		MAV_PROTOCOL_CAPABILITY_FTP,
		MAV_PROTOCOL_CAPABILITY_SET_ATTITUDE_TARGET,
		MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_LOCAL_NED,
		MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_GLOBAL_INT,
		MAV_PROTOCOL_CAPABILITY_TERRAIN,
		MAV_PROTOCOL_CAPABILITY_SET_ACTUATOR_TARGET,
		MAV_PROTOCOL_CAPABILITY_FLIGHT_TERMINATION,
		MAV_PROTOCOL_CAPABILITY_COMPASS_CALIBRATION,
		MAV_PROTOCOL_CAPABILITY_MAVLINK2,
		MAV_PROTOCOL_CAPABILITY_MISSION_FENCE,
		MAV_PROTOCOL_CAPABILITY_MISSION_RALLY,
		MAV_PROTOCOL_CAPABILITY_FLIGHT_INFORMATION,
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

// MAV_MISSION_TYPE type. Type of mission items being requested/sent in mission protocol.
type MAV_MISSION_TYPE int

func (e MAV_MISSION_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_MISSION_TYPE_MISSION enum. Items are mission commands for main mission
	MAV_MISSION_TYPE_MISSION MAV_MISSION_TYPE = 0
	// MAV_MISSION_TYPE_FENCE enum. Specifies GeoFence area(s). Items are MAV_CMD_NAV_FENCE_ GeoFence items
	MAV_MISSION_TYPE_FENCE MAV_MISSION_TYPE = 1
	// MAV_MISSION_TYPE_RALLY enum. Specifies the rally points for the vehicle. Rally points are alternative RTL points. Items are MAV_CMD_NAV_RALLY_POINT rally point items
	MAV_MISSION_TYPE_RALLY MAV_MISSION_TYPE = 2
	// MAV_MISSION_TYPE_ALL enum. Only used in MISSION_CLEAR_ALL to clear all mission types
	MAV_MISSION_TYPE_ALL MAV_MISSION_TYPE = 255
)

func (e MAV_MISSION_TYPE) String() string {
	if name, ok := map[MAV_MISSION_TYPE]string{
		MAV_MISSION_TYPE_MISSION: "MAV_MISSION_TYPE_MISSION",
		MAV_MISSION_TYPE_FENCE:   "MAV_MISSION_TYPE_FENCE",
		MAV_MISSION_TYPE_RALLY:   "MAV_MISSION_TYPE_RALLY",
		MAV_MISSION_TYPE_ALL:     "MAV_MISSION_TYPE_ALL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_MISSION_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_MISSION_TYPE enums
func (e MAV_MISSION_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_MISSION_TYPE{
		MAV_MISSION_TYPE_MISSION,
		MAV_MISSION_TYPE_FENCE,
		MAV_MISSION_TYPE_RALLY,
		MAV_MISSION_TYPE_ALL,
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

// MAV_ESTIMATOR_TYPE type. Enumeration of estimator types
type MAV_ESTIMATOR_TYPE int

func (e MAV_ESTIMATOR_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ESTIMATOR_TYPE_UNKNOWN enum. Unknown type of the estimator
	MAV_ESTIMATOR_TYPE_UNKNOWN MAV_ESTIMATOR_TYPE = 0
	// MAV_ESTIMATOR_TYPE_NAIVE enum. This is a naive estimator without any real covariance feedback
	MAV_ESTIMATOR_TYPE_NAIVE MAV_ESTIMATOR_TYPE = 1
	// MAV_ESTIMATOR_TYPE_VISION enum. Computer vision based estimate. Might be up to scale
	MAV_ESTIMATOR_TYPE_VISION MAV_ESTIMATOR_TYPE = 2
	// MAV_ESTIMATOR_TYPE_VIO enum. Visual-inertial estimate
	MAV_ESTIMATOR_TYPE_VIO MAV_ESTIMATOR_TYPE = 3
	// MAV_ESTIMATOR_TYPE_GPS enum. Plain GPS estimate
	MAV_ESTIMATOR_TYPE_GPS MAV_ESTIMATOR_TYPE = 4
	// MAV_ESTIMATOR_TYPE_GPS_INS enum. Estimator integrating GPS and inertial sensing
	MAV_ESTIMATOR_TYPE_GPS_INS MAV_ESTIMATOR_TYPE = 5
	// MAV_ESTIMATOR_TYPE_MOCAP enum. Estimate from external motion capturing system
	MAV_ESTIMATOR_TYPE_MOCAP MAV_ESTIMATOR_TYPE = 6
	// MAV_ESTIMATOR_TYPE_LIDAR enum. Estimator based on lidar sensor input
	MAV_ESTIMATOR_TYPE_LIDAR MAV_ESTIMATOR_TYPE = 7
	// MAV_ESTIMATOR_TYPE_AUTOPILOT enum. Estimator on autopilot
	MAV_ESTIMATOR_TYPE_AUTOPILOT MAV_ESTIMATOR_TYPE = 8
)

func (e MAV_ESTIMATOR_TYPE) String() string {
	if name, ok := map[MAV_ESTIMATOR_TYPE]string{
		MAV_ESTIMATOR_TYPE_UNKNOWN:   "MAV_ESTIMATOR_TYPE_UNKNOWN",
		MAV_ESTIMATOR_TYPE_NAIVE:     "MAV_ESTIMATOR_TYPE_NAIVE",
		MAV_ESTIMATOR_TYPE_VISION:    "MAV_ESTIMATOR_TYPE_VISION",
		MAV_ESTIMATOR_TYPE_VIO:       "MAV_ESTIMATOR_TYPE_VIO",
		MAV_ESTIMATOR_TYPE_GPS:       "MAV_ESTIMATOR_TYPE_GPS",
		MAV_ESTIMATOR_TYPE_GPS_INS:   "MAV_ESTIMATOR_TYPE_GPS_INS",
		MAV_ESTIMATOR_TYPE_MOCAP:     "MAV_ESTIMATOR_TYPE_MOCAP",
		MAV_ESTIMATOR_TYPE_LIDAR:     "MAV_ESTIMATOR_TYPE_LIDAR",
		MAV_ESTIMATOR_TYPE_AUTOPILOT: "MAV_ESTIMATOR_TYPE_AUTOPILOT",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ESTIMATOR_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ESTIMATOR_TYPE enums
func (e MAV_ESTIMATOR_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ESTIMATOR_TYPE{
		MAV_ESTIMATOR_TYPE_UNKNOWN,
		MAV_ESTIMATOR_TYPE_NAIVE,
		MAV_ESTIMATOR_TYPE_VISION,
		MAV_ESTIMATOR_TYPE_VIO,
		MAV_ESTIMATOR_TYPE_GPS,
		MAV_ESTIMATOR_TYPE_GPS_INS,
		MAV_ESTIMATOR_TYPE_MOCAP,
		MAV_ESTIMATOR_TYPE_LIDAR,
		MAV_ESTIMATOR_TYPE_AUTOPILOT,
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

// MAV_BATTERY_TYPE type. Enumeration of battery types
type MAV_BATTERY_TYPE int

func (e MAV_BATTERY_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_BATTERY_TYPE_UNKNOWN enum. Not specified
	MAV_BATTERY_TYPE_UNKNOWN MAV_BATTERY_TYPE = 0
	// MAV_BATTERY_TYPE_LIPO enum. Lithium polymer battery
	MAV_BATTERY_TYPE_LIPO MAV_BATTERY_TYPE = 1
	// MAV_BATTERY_TYPE_LIFE enum. Lithium-iron-phosphate battery
	MAV_BATTERY_TYPE_LIFE MAV_BATTERY_TYPE = 2
	// MAV_BATTERY_TYPE_LION enum. Lithium-ION battery
	MAV_BATTERY_TYPE_LION MAV_BATTERY_TYPE = 3
	// MAV_BATTERY_TYPE_NIMH enum. Nickel metal hydride battery
	MAV_BATTERY_TYPE_NIMH MAV_BATTERY_TYPE = 4
)

func (e MAV_BATTERY_TYPE) String() string {
	if name, ok := map[MAV_BATTERY_TYPE]string{
		MAV_BATTERY_TYPE_UNKNOWN: "MAV_BATTERY_TYPE_UNKNOWN",
		MAV_BATTERY_TYPE_LIPO:    "MAV_BATTERY_TYPE_LIPO",
		MAV_BATTERY_TYPE_LIFE:    "MAV_BATTERY_TYPE_LIFE",
		MAV_BATTERY_TYPE_LION:    "MAV_BATTERY_TYPE_LION",
		MAV_BATTERY_TYPE_NIMH:    "MAV_BATTERY_TYPE_NIMH",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_BATTERY_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_BATTERY_TYPE enums
func (e MAV_BATTERY_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_BATTERY_TYPE{
		MAV_BATTERY_TYPE_UNKNOWN,
		MAV_BATTERY_TYPE_LIPO,
		MAV_BATTERY_TYPE_LIFE,
		MAV_BATTERY_TYPE_LION,
		MAV_BATTERY_TYPE_NIMH,
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

// MAV_BATTERY_FUNCTION type. Enumeration of battery functions
type MAV_BATTERY_FUNCTION int

func (e MAV_BATTERY_FUNCTION) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_BATTERY_FUNCTION_UNKNOWN enum. Battery function is unknown
	MAV_BATTERY_FUNCTION_UNKNOWN MAV_BATTERY_FUNCTION = 0
	// MAV_BATTERY_FUNCTION_ALL enum. Battery supports all flight systems
	MAV_BATTERY_FUNCTION_ALL MAV_BATTERY_FUNCTION = 1
	// MAV_BATTERY_FUNCTION_PROPULSION enum. Battery for the propulsion system
	MAV_BATTERY_FUNCTION_PROPULSION MAV_BATTERY_FUNCTION = 2
	// MAV_BATTERY_FUNCTION_AVIONICS enum. Avionics battery
	MAV_BATTERY_FUNCTION_AVIONICS MAV_BATTERY_FUNCTION = 3
	// MAV_BATTERY_TYPE_PAYLOAD enum. Payload battery
	MAV_BATTERY_TYPE_PAYLOAD MAV_BATTERY_FUNCTION = 4
)

func (e MAV_BATTERY_FUNCTION) String() string {
	if name, ok := map[MAV_BATTERY_FUNCTION]string{
		MAV_BATTERY_FUNCTION_UNKNOWN:    "MAV_BATTERY_FUNCTION_UNKNOWN",
		MAV_BATTERY_FUNCTION_ALL:        "MAV_BATTERY_FUNCTION_ALL",
		MAV_BATTERY_FUNCTION_PROPULSION: "MAV_BATTERY_FUNCTION_PROPULSION",
		MAV_BATTERY_FUNCTION_AVIONICS:   "MAV_BATTERY_FUNCTION_AVIONICS",
		MAV_BATTERY_TYPE_PAYLOAD:        "MAV_BATTERY_TYPE_PAYLOAD",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_BATTERY_FUNCTION_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_BATTERY_FUNCTION enums
func (e MAV_BATTERY_FUNCTION) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_BATTERY_FUNCTION{
		MAV_BATTERY_FUNCTION_UNKNOWN,
		MAV_BATTERY_FUNCTION_ALL,
		MAV_BATTERY_FUNCTION_PROPULSION,
		MAV_BATTERY_FUNCTION_AVIONICS,
		MAV_BATTERY_TYPE_PAYLOAD,
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

// MAV_BATTERY_CHARGE_STATE type. Enumeration for battery charge states.
type MAV_BATTERY_CHARGE_STATE int

func (e MAV_BATTERY_CHARGE_STATE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_BATTERY_CHARGE_STATE_UNDEFINED enum. Low battery state is not provided
	MAV_BATTERY_CHARGE_STATE_UNDEFINED MAV_BATTERY_CHARGE_STATE = 0
	// MAV_BATTERY_CHARGE_STATE_OK enum. Battery is not in low state. Normal operation
	MAV_BATTERY_CHARGE_STATE_OK MAV_BATTERY_CHARGE_STATE = 1
	// MAV_BATTERY_CHARGE_STATE_LOW enum. Battery state is low, warn and monitor close
	MAV_BATTERY_CHARGE_STATE_LOW MAV_BATTERY_CHARGE_STATE = 2
	// MAV_BATTERY_CHARGE_STATE_CRITICAL enum. Battery state is critical, return or abort immediately
	MAV_BATTERY_CHARGE_STATE_CRITICAL MAV_BATTERY_CHARGE_STATE = 3
	// MAV_BATTERY_CHARGE_STATE_EMERGENCY enum. Battery state is too low for ordinary abort sequence. Perform fastest possible emergency stop to prevent damage
	MAV_BATTERY_CHARGE_STATE_EMERGENCY MAV_BATTERY_CHARGE_STATE = 4
	// MAV_BATTERY_CHARGE_STATE_FAILED enum. Battery failed, damage unavoidable. Possible causes (faults) are listed in MAV_BATTERY_FAULT
	MAV_BATTERY_CHARGE_STATE_FAILED MAV_BATTERY_CHARGE_STATE = 5
	// MAV_BATTERY_CHARGE_STATE_UNHEALTHY enum. Battery is diagnosed to be defective or an error occurred, usage is discouraged / prohibited. Possible causes (faults) are listed in MAV_BATTERY_FAULT
	MAV_BATTERY_CHARGE_STATE_UNHEALTHY MAV_BATTERY_CHARGE_STATE = 6
	// MAV_BATTERY_CHARGE_STATE_CHARGING enum. Battery is charging
	MAV_BATTERY_CHARGE_STATE_CHARGING MAV_BATTERY_CHARGE_STATE = 7
)

func (e MAV_BATTERY_CHARGE_STATE) String() string {
	if name, ok := map[MAV_BATTERY_CHARGE_STATE]string{
		MAV_BATTERY_CHARGE_STATE_UNDEFINED: "MAV_BATTERY_CHARGE_STATE_UNDEFINED",
		MAV_BATTERY_CHARGE_STATE_OK:        "MAV_BATTERY_CHARGE_STATE_OK",
		MAV_BATTERY_CHARGE_STATE_LOW:       "MAV_BATTERY_CHARGE_STATE_LOW",
		MAV_BATTERY_CHARGE_STATE_CRITICAL:  "MAV_BATTERY_CHARGE_STATE_CRITICAL",
		MAV_BATTERY_CHARGE_STATE_EMERGENCY: "MAV_BATTERY_CHARGE_STATE_EMERGENCY",
		MAV_BATTERY_CHARGE_STATE_FAILED:    "MAV_BATTERY_CHARGE_STATE_FAILED",
		MAV_BATTERY_CHARGE_STATE_UNHEALTHY: "MAV_BATTERY_CHARGE_STATE_UNHEALTHY",
		MAV_BATTERY_CHARGE_STATE_CHARGING:  "MAV_BATTERY_CHARGE_STATE_CHARGING",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_BATTERY_CHARGE_STATE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_BATTERY_CHARGE_STATE enums
func (e MAV_BATTERY_CHARGE_STATE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_BATTERY_CHARGE_STATE{
		MAV_BATTERY_CHARGE_STATE_UNDEFINED,
		MAV_BATTERY_CHARGE_STATE_OK,
		MAV_BATTERY_CHARGE_STATE_LOW,
		MAV_BATTERY_CHARGE_STATE_CRITICAL,
		MAV_BATTERY_CHARGE_STATE_EMERGENCY,
		MAV_BATTERY_CHARGE_STATE_FAILED,
		MAV_BATTERY_CHARGE_STATE_UNHEALTHY,
		MAV_BATTERY_CHARGE_STATE_CHARGING,
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

// MAV_BATTERY_MODE type. Battery mode. Note, the normal operation mode (i.e. when flying) should be reported as MAV_BATTERY_MODE_UNKNOWN to allow message trimming in normal flight.
type MAV_BATTERY_MODE int

func (e MAV_BATTERY_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_BATTERY_MODE_UNKNOWN enum. Battery mode not supported/unknown battery mode/normal operation
	MAV_BATTERY_MODE_UNKNOWN MAV_BATTERY_MODE = 0
	// MAV_BATTERY_MODE_AUTO_DISCHARGING enum. Battery is auto discharging (towards storage level)
	MAV_BATTERY_MODE_AUTO_DISCHARGING MAV_BATTERY_MODE = 1
	// MAV_BATTERY_MODE_HOT_SWAP enum. Battery in hot-swap mode (current limited to prevent spikes that might damage sensitive electrical circuits)
	MAV_BATTERY_MODE_HOT_SWAP MAV_BATTERY_MODE = 2
)

func (e MAV_BATTERY_MODE) String() string {
	if name, ok := map[MAV_BATTERY_MODE]string{
		MAV_BATTERY_MODE_UNKNOWN:          "MAV_BATTERY_MODE_UNKNOWN",
		MAV_BATTERY_MODE_AUTO_DISCHARGING: "MAV_BATTERY_MODE_AUTO_DISCHARGING",
		MAV_BATTERY_MODE_HOT_SWAP:         "MAV_BATTERY_MODE_HOT_SWAP",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_BATTERY_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_BATTERY_MODE enums
func (e MAV_BATTERY_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_BATTERY_MODE{
		MAV_BATTERY_MODE_UNKNOWN,
		MAV_BATTERY_MODE_AUTO_DISCHARGING,
		MAV_BATTERY_MODE_HOT_SWAP,
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

// MAV_BATTERY_FAULT type. Smart battery supply status/fault flags (bitmask) for health indication. The battery must also report either MAV_BATTERY_CHARGE_STATE_FAILED or MAV_BATTERY_CHARGE_STATE_UNHEALTHY if any of these are set.
type MAV_BATTERY_FAULT int

func (e MAV_BATTERY_FAULT) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_BATTERY_FAULT_DEEP_DISCHARGE enum. Battery has deep discharged
	MAV_BATTERY_FAULT_DEEP_DISCHARGE MAV_BATTERY_FAULT = 1
	// MAV_BATTERY_FAULT_SPIKES enum. Voltage spikes
	MAV_BATTERY_FAULT_SPIKES MAV_BATTERY_FAULT = 2
	// MAV_BATTERY_FAULT_CELL_FAIL enum. One or more cells have failed. Battery should also report MAV_BATTERY_CHARGE_STATE_FAILE (and should not be used)
	MAV_BATTERY_FAULT_CELL_FAIL MAV_BATTERY_FAULT = 4
	// MAV_BATTERY_FAULT_OVER_CURRENT enum. Over-current fault
	MAV_BATTERY_FAULT_OVER_CURRENT MAV_BATTERY_FAULT = 8
	// MAV_BATTERY_FAULT_OVER_TEMPERATURE enum. Over-temperature fault
	MAV_BATTERY_FAULT_OVER_TEMPERATURE MAV_BATTERY_FAULT = 16
	// MAV_BATTERY_FAULT_UNDER_TEMPERATURE enum. Under-temperature fault
	MAV_BATTERY_FAULT_UNDER_TEMPERATURE MAV_BATTERY_FAULT = 32
	// MAV_BATTERY_FAULT_INCOMPATIBLE_VOLTAGE enum. Vehicle voltage is not compatible with this battery (batteries on same power rail should have similar voltage)
	MAV_BATTERY_FAULT_INCOMPATIBLE_VOLTAGE MAV_BATTERY_FAULT = 64
)

func (e MAV_BATTERY_FAULT) String() string {
	if name, ok := map[MAV_BATTERY_FAULT]string{
		MAV_BATTERY_FAULT_DEEP_DISCHARGE:       "MAV_BATTERY_FAULT_DEEP_DISCHARGE",
		MAV_BATTERY_FAULT_SPIKES:               "MAV_BATTERY_FAULT_SPIKES",
		MAV_BATTERY_FAULT_CELL_FAIL:            "MAV_BATTERY_FAULT_CELL_FAIL",
		MAV_BATTERY_FAULT_OVER_CURRENT:         "MAV_BATTERY_FAULT_OVER_CURRENT",
		MAV_BATTERY_FAULT_OVER_TEMPERATURE:     "MAV_BATTERY_FAULT_OVER_TEMPERATURE",
		MAV_BATTERY_FAULT_UNDER_TEMPERATURE:    "MAV_BATTERY_FAULT_UNDER_TEMPERATURE",
		MAV_BATTERY_FAULT_INCOMPATIBLE_VOLTAGE: "MAV_BATTERY_FAULT_INCOMPATIBLE_VOLTAGE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_BATTERY_FAULT_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_BATTERY_FAULT enums
func (e MAV_BATTERY_FAULT) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_BATTERY_FAULT{
		MAV_BATTERY_FAULT_DEEP_DISCHARGE,
		MAV_BATTERY_FAULT_SPIKES,
		MAV_BATTERY_FAULT_CELL_FAIL,
		MAV_BATTERY_FAULT_OVER_CURRENT,
		MAV_BATTERY_FAULT_OVER_TEMPERATURE,
		MAV_BATTERY_FAULT_UNDER_TEMPERATURE,
		MAV_BATTERY_FAULT_INCOMPATIBLE_VOLTAGE,
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

// MAV_GENERATOR_STATUS_FLAG type. Flags to report status/failure cases for a power generator (used in GENERATOR_STATUS). Note that FAULTS are conditions that cause the generator to fail. Warnings are conditions that require attention before the next use (they indicate the system is not operating properly).
type MAV_GENERATOR_STATUS_FLAG int

func (e MAV_GENERATOR_STATUS_FLAG) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_GENERATOR_STATUS_FLAG_OFF enum. Generator is off
	MAV_GENERATOR_STATUS_FLAG_OFF MAV_GENERATOR_STATUS_FLAG = 1
	// MAV_GENERATOR_STATUS_FLAG_READY enum. Generator is ready to start generating power
	MAV_GENERATOR_STATUS_FLAG_READY MAV_GENERATOR_STATUS_FLAG = 2
	// MAV_GENERATOR_STATUS_FLAG_GENERATING enum. Generator is generating power
	MAV_GENERATOR_STATUS_FLAG_GENERATING MAV_GENERATOR_STATUS_FLAG = 4
	// MAV_GENERATOR_STATUS_FLAG_CHARGING enum. Generator is charging the batteries (generating enough power to charge and provide the load)
	MAV_GENERATOR_STATUS_FLAG_CHARGING MAV_GENERATOR_STATUS_FLAG = 8
	// MAV_GENERATOR_STATUS_FLAG_REDUCED_POWER enum. Generator is operating at a reduced maximum power
	MAV_GENERATOR_STATUS_FLAG_REDUCED_POWER MAV_GENERATOR_STATUS_FLAG = 16
	// MAV_GENERATOR_STATUS_FLAG_MAXPOWER enum. Generator is providing the maximum output
	MAV_GENERATOR_STATUS_FLAG_MAXPOWER MAV_GENERATOR_STATUS_FLAG = 32
	// MAV_GENERATOR_STATUS_FLAG_OVERTEMP_WARNING enum. Generator is near the maximum operating temperature, cooling is insufficient
	MAV_GENERATOR_STATUS_FLAG_OVERTEMP_WARNING MAV_GENERATOR_STATUS_FLAG = 64
	// MAV_GENERATOR_STATUS_FLAG_OVERTEMP_FAULT enum. Generator hit the maximum operating temperature and shutdown
	MAV_GENERATOR_STATUS_FLAG_OVERTEMP_FAULT MAV_GENERATOR_STATUS_FLAG = 128
	// MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING enum. Power electronics are near the maximum operating temperature, cooling is insufficient
	MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING MAV_GENERATOR_STATUS_FLAG = 256
	// MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT enum. Power electronics hit the maximum operating temperature and shutdown
	MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT MAV_GENERATOR_STATUS_FLAG = 512
	// MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_FAULT enum. Power electronics experienced a fault and shutdown
	MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_FAULT MAV_GENERATOR_STATUS_FLAG = 1024
	// MAV_GENERATOR_STATUS_FLAG_POWERSOURCE_FAULT enum. The power source supplying the generator failed e.g. mechanical generator stopped, tether is no longer providing power, solar cell is in shade, hydrogen reaction no longer happening
	MAV_GENERATOR_STATUS_FLAG_POWERSOURCE_FAULT MAV_GENERATOR_STATUS_FLAG = 2048
	// MAV_GENERATOR_STATUS_FLAG_COMMUNICATION_WARNING enum. Generator controller having communication problems
	MAV_GENERATOR_STATUS_FLAG_COMMUNICATION_WARNING MAV_GENERATOR_STATUS_FLAG = 4096
	// MAV_GENERATOR_STATUS_FLAG_COOLING_WARNING enum. Power electronic or generator cooling system error
	MAV_GENERATOR_STATUS_FLAG_COOLING_WARNING MAV_GENERATOR_STATUS_FLAG = 8192
	// MAV_GENERATOR_STATUS_FLAG_POWER_RAIL_FAULT enum. Generator controller power rail experienced a fault
	MAV_GENERATOR_STATUS_FLAG_POWER_RAIL_FAULT MAV_GENERATOR_STATUS_FLAG = 16384
	// MAV_GENERATOR_STATUS_FLAG_OVERCURRENT_FAULT enum. Generator controller exceeded the overcurrent threshold and shutdown to prevent damage
	MAV_GENERATOR_STATUS_FLAG_OVERCURRENT_FAULT MAV_GENERATOR_STATUS_FLAG = 32768
	// MAV_GENERATOR_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT enum. Generator controller detected a high current going into the batteries and shutdown to prevent battery damage
	MAV_GENERATOR_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT MAV_GENERATOR_STATUS_FLAG = 65536
	// MAV_GENERATOR_STATUS_FLAG_OVERVOLTAGE_FAULT enum. Generator controller exceeded it's overvoltage threshold and shutdown to prevent it exceeding the voltage rating
	MAV_GENERATOR_STATUS_FLAG_OVERVOLTAGE_FAULT MAV_GENERATOR_STATUS_FLAG = 131072
	// MAV_GENERATOR_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT enum. Batteries are under voltage (generator will not start)
	MAV_GENERATOR_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT MAV_GENERATOR_STATUS_FLAG = 262144
	// MAV_GENERATOR_STATUS_FLAG_START_INHIBITED enum. Generator start is inhibited by e.g. a safety switch
	MAV_GENERATOR_STATUS_FLAG_START_INHIBITED MAV_GENERATOR_STATUS_FLAG = 524288
	// MAV_GENERATOR_STATUS_FLAG_MAINTENANCE_REQUIRED enum. Generator requires maintenance
	MAV_GENERATOR_STATUS_FLAG_MAINTENANCE_REQUIRED MAV_GENERATOR_STATUS_FLAG = 1048576
	// MAV_GENERATOR_STATUS_FLAG_WARMING_UP enum. Generator is not ready to generate yet
	MAV_GENERATOR_STATUS_FLAG_WARMING_UP MAV_GENERATOR_STATUS_FLAG = 2097152
	// MAV_GENERATOR_STATUS_FLAG_IDLE enum. Generator is idle
	MAV_GENERATOR_STATUS_FLAG_IDLE MAV_GENERATOR_STATUS_FLAG = 4194304
)

func (e MAV_GENERATOR_STATUS_FLAG) String() string {
	if name, ok := map[MAV_GENERATOR_STATUS_FLAG]string{
		MAV_GENERATOR_STATUS_FLAG_OFF:                              "MAV_GENERATOR_STATUS_FLAG_OFF",
		MAV_GENERATOR_STATUS_FLAG_READY:                            "MAV_GENERATOR_STATUS_FLAG_READY",
		MAV_GENERATOR_STATUS_FLAG_GENERATING:                       "MAV_GENERATOR_STATUS_FLAG_GENERATING",
		MAV_GENERATOR_STATUS_FLAG_CHARGING:                         "MAV_GENERATOR_STATUS_FLAG_CHARGING",
		MAV_GENERATOR_STATUS_FLAG_REDUCED_POWER:                    "MAV_GENERATOR_STATUS_FLAG_REDUCED_POWER",
		MAV_GENERATOR_STATUS_FLAG_MAXPOWER:                         "MAV_GENERATOR_STATUS_FLAG_MAXPOWER",
		MAV_GENERATOR_STATUS_FLAG_OVERTEMP_WARNING:                 "MAV_GENERATOR_STATUS_FLAG_OVERTEMP_WARNING",
		MAV_GENERATOR_STATUS_FLAG_OVERTEMP_FAULT:                   "MAV_GENERATOR_STATUS_FLAG_OVERTEMP_FAULT",
		MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING:     "MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING",
		MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT:       "MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT",
		MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_FAULT:                "MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_FAULT",
		MAV_GENERATOR_STATUS_FLAG_POWERSOURCE_FAULT:                "MAV_GENERATOR_STATUS_FLAG_POWERSOURCE_FAULT",
		MAV_GENERATOR_STATUS_FLAG_COMMUNICATION_WARNING:            "MAV_GENERATOR_STATUS_FLAG_COMMUNICATION_WARNING",
		MAV_GENERATOR_STATUS_FLAG_COOLING_WARNING:                  "MAV_GENERATOR_STATUS_FLAG_COOLING_WARNING",
		MAV_GENERATOR_STATUS_FLAG_POWER_RAIL_FAULT:                 "MAV_GENERATOR_STATUS_FLAG_POWER_RAIL_FAULT",
		MAV_GENERATOR_STATUS_FLAG_OVERCURRENT_FAULT:                "MAV_GENERATOR_STATUS_FLAG_OVERCURRENT_FAULT",
		MAV_GENERATOR_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT: "MAV_GENERATOR_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT",
		MAV_GENERATOR_STATUS_FLAG_OVERVOLTAGE_FAULT:                "MAV_GENERATOR_STATUS_FLAG_OVERVOLTAGE_FAULT",
		MAV_GENERATOR_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT:          "MAV_GENERATOR_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT",
		MAV_GENERATOR_STATUS_FLAG_START_INHIBITED:                  "MAV_GENERATOR_STATUS_FLAG_START_INHIBITED",
		MAV_GENERATOR_STATUS_FLAG_MAINTENANCE_REQUIRED:             "MAV_GENERATOR_STATUS_FLAG_MAINTENANCE_REQUIRED",
		MAV_GENERATOR_STATUS_FLAG_WARMING_UP:                       "MAV_GENERATOR_STATUS_FLAG_WARMING_UP",
		MAV_GENERATOR_STATUS_FLAG_IDLE:                             "MAV_GENERATOR_STATUS_FLAG_IDLE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_GENERATOR_STATUS_FLAG_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_GENERATOR_STATUS_FLAG enums
func (e MAV_GENERATOR_STATUS_FLAG) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_GENERATOR_STATUS_FLAG{
		MAV_GENERATOR_STATUS_FLAG_OFF,
		MAV_GENERATOR_STATUS_FLAG_READY,
		MAV_GENERATOR_STATUS_FLAG_GENERATING,
		MAV_GENERATOR_STATUS_FLAG_CHARGING,
		MAV_GENERATOR_STATUS_FLAG_REDUCED_POWER,
		MAV_GENERATOR_STATUS_FLAG_MAXPOWER,
		MAV_GENERATOR_STATUS_FLAG_OVERTEMP_WARNING,
		MAV_GENERATOR_STATUS_FLAG_OVERTEMP_FAULT,
		MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING,
		MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT,
		MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_FAULT,
		MAV_GENERATOR_STATUS_FLAG_POWERSOURCE_FAULT,
		MAV_GENERATOR_STATUS_FLAG_COMMUNICATION_WARNING,
		MAV_GENERATOR_STATUS_FLAG_COOLING_WARNING,
		MAV_GENERATOR_STATUS_FLAG_POWER_RAIL_FAULT,
		MAV_GENERATOR_STATUS_FLAG_OVERCURRENT_FAULT,
		MAV_GENERATOR_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT,
		MAV_GENERATOR_STATUS_FLAG_OVERVOLTAGE_FAULT,
		MAV_GENERATOR_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT,
		MAV_GENERATOR_STATUS_FLAG_START_INHIBITED,
		MAV_GENERATOR_STATUS_FLAG_MAINTENANCE_REQUIRED,
		MAV_GENERATOR_STATUS_FLAG_WARMING_UP,
		MAV_GENERATOR_STATUS_FLAG_IDLE,
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

// MAV_VTOL_STATE type. Enumeration of VTOL states
type MAV_VTOL_STATE int

func (e MAV_VTOL_STATE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_VTOL_STATE_UNDEFINED enum. MAV is not configured as VTOL
	MAV_VTOL_STATE_UNDEFINED MAV_VTOL_STATE = 0
	// MAV_VTOL_STATE_TRANSITION_TO_FW enum. VTOL is in transition from multicopter to fixed-wing
	MAV_VTOL_STATE_TRANSITION_TO_FW MAV_VTOL_STATE = 1
	// MAV_VTOL_STATE_TRANSITION_TO_MC enum. VTOL is in transition from fixed-wing to multicopter
	MAV_VTOL_STATE_TRANSITION_TO_MC MAV_VTOL_STATE = 2
	// MAV_VTOL_STATE_MC enum. VTOL is in multicopter state
	MAV_VTOL_STATE_MC MAV_VTOL_STATE = 3
	// MAV_VTOL_STATE_FW enum. VTOL is in fixed-wing state
	MAV_VTOL_STATE_FW MAV_VTOL_STATE = 4
)

func (e MAV_VTOL_STATE) String() string {
	if name, ok := map[MAV_VTOL_STATE]string{
		MAV_VTOL_STATE_UNDEFINED:        "MAV_VTOL_STATE_UNDEFINED",
		MAV_VTOL_STATE_TRANSITION_TO_FW: "MAV_VTOL_STATE_TRANSITION_TO_FW",
		MAV_VTOL_STATE_TRANSITION_TO_MC: "MAV_VTOL_STATE_TRANSITION_TO_MC",
		MAV_VTOL_STATE_MC:               "MAV_VTOL_STATE_MC",
		MAV_VTOL_STATE_FW:               "MAV_VTOL_STATE_FW",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_VTOL_STATE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_VTOL_STATE enums
func (e MAV_VTOL_STATE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_VTOL_STATE{
		MAV_VTOL_STATE_UNDEFINED,
		MAV_VTOL_STATE_TRANSITION_TO_FW,
		MAV_VTOL_STATE_TRANSITION_TO_MC,
		MAV_VTOL_STATE_MC,
		MAV_VTOL_STATE_FW,
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

// MAV_LANDED_STATE type. Enumeration of landed detector states
type MAV_LANDED_STATE int

func (e MAV_LANDED_STATE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_LANDED_STATE_UNDEFINED enum. MAV landed state is unknown
	MAV_LANDED_STATE_UNDEFINED MAV_LANDED_STATE = 0
	// MAV_LANDED_STATE_ON_GROUND enum. MAV is landed (on ground)
	MAV_LANDED_STATE_ON_GROUND MAV_LANDED_STATE = 1
	// MAV_LANDED_STATE_IN_AIR enum. MAV is in air
	MAV_LANDED_STATE_IN_AIR MAV_LANDED_STATE = 2
	// MAV_LANDED_STATE_TAKEOFF enum. MAV currently taking off
	MAV_LANDED_STATE_TAKEOFF MAV_LANDED_STATE = 3
	// MAV_LANDED_STATE_LANDING enum. MAV currently landing
	MAV_LANDED_STATE_LANDING MAV_LANDED_STATE = 4
)

func (e MAV_LANDED_STATE) String() string {
	if name, ok := map[MAV_LANDED_STATE]string{
		MAV_LANDED_STATE_UNDEFINED: "MAV_LANDED_STATE_UNDEFINED",
		MAV_LANDED_STATE_ON_GROUND: "MAV_LANDED_STATE_ON_GROUND",
		MAV_LANDED_STATE_IN_AIR:    "MAV_LANDED_STATE_IN_AIR",
		MAV_LANDED_STATE_TAKEOFF:   "MAV_LANDED_STATE_TAKEOFF",
		MAV_LANDED_STATE_LANDING:   "MAV_LANDED_STATE_LANDING",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_LANDED_STATE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_LANDED_STATE enums
func (e MAV_LANDED_STATE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_LANDED_STATE{
		MAV_LANDED_STATE_UNDEFINED,
		MAV_LANDED_STATE_ON_GROUND,
		MAV_LANDED_STATE_IN_AIR,
		MAV_LANDED_STATE_TAKEOFF,
		MAV_LANDED_STATE_LANDING,
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

// ADSB_ALTITUDE_TYPE type. Enumeration of the ADSB altimeter types
type ADSB_ALTITUDE_TYPE int

func (e ADSB_ALTITUDE_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ADSB_ALTITUDE_TYPE_PRESSURE_QNH enum. Altitude reported from a Baro source using QNH reference
	ADSB_ALTITUDE_TYPE_PRESSURE_QNH ADSB_ALTITUDE_TYPE = 0
	// ADSB_ALTITUDE_TYPE_GEOMETRIC enum. Altitude reported from a GNSS source
	ADSB_ALTITUDE_TYPE_GEOMETRIC ADSB_ALTITUDE_TYPE = 1
)

func (e ADSB_ALTITUDE_TYPE) String() string {
	if name, ok := map[ADSB_ALTITUDE_TYPE]string{
		ADSB_ALTITUDE_TYPE_PRESSURE_QNH: "ADSB_ALTITUDE_TYPE_PRESSURE_QNH",
		ADSB_ALTITUDE_TYPE_GEOMETRIC:    "ADSB_ALTITUDE_TYPE_GEOMETRIC",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("ADSB_ALTITUDE_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects ADSB_ALTITUDE_TYPE enums
func (e ADSB_ALTITUDE_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []ADSB_ALTITUDE_TYPE{
		ADSB_ALTITUDE_TYPE_PRESSURE_QNH,
		ADSB_ALTITUDE_TYPE_GEOMETRIC,
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

// ADSB_EMITTER_TYPE type. ADSB classification for the type of vehicle emitting the transponder signal
type ADSB_EMITTER_TYPE int

func (e ADSB_EMITTER_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ADSB_EMITTER_TYPE_NO_INFO enum
	ADSB_EMITTER_TYPE_NO_INFO ADSB_EMITTER_TYPE = 0
	// ADSB_EMITTER_TYPE_LIGHT enum
	ADSB_EMITTER_TYPE_LIGHT ADSB_EMITTER_TYPE = 1
	// ADSB_EMITTER_TYPE_SMALL enum
	ADSB_EMITTER_TYPE_SMALL ADSB_EMITTER_TYPE = 2
	// ADSB_EMITTER_TYPE_LARGE enum
	ADSB_EMITTER_TYPE_LARGE ADSB_EMITTER_TYPE = 3
	// ADSB_EMITTER_TYPE_HIGH_VORTEX_LARGE enum
	ADSB_EMITTER_TYPE_HIGH_VORTEX_LARGE ADSB_EMITTER_TYPE = 4
	// ADSB_EMITTER_TYPE_HEAVY enum
	ADSB_EMITTER_TYPE_HEAVY ADSB_EMITTER_TYPE = 5
	// ADSB_EMITTER_TYPE_HIGHLY_MANUV enum
	ADSB_EMITTER_TYPE_HIGHLY_MANUV ADSB_EMITTER_TYPE = 6
	// ADSB_EMITTER_TYPE_ROTOCRAFT enum
	ADSB_EMITTER_TYPE_ROTOCRAFT ADSB_EMITTER_TYPE = 7
	// ADSB_EMITTER_TYPE_UNASSIGNED enum
	ADSB_EMITTER_TYPE_UNASSIGNED ADSB_EMITTER_TYPE = 8
	// ADSB_EMITTER_TYPE_GLIDER enum
	ADSB_EMITTER_TYPE_GLIDER ADSB_EMITTER_TYPE = 9
	// ADSB_EMITTER_TYPE_LIGHTER_AIR enum
	ADSB_EMITTER_TYPE_LIGHTER_AIR ADSB_EMITTER_TYPE = 10
	// ADSB_EMITTER_TYPE_PARACHUTE enum
	ADSB_EMITTER_TYPE_PARACHUTE ADSB_EMITTER_TYPE = 11
	// ADSB_EMITTER_TYPE_ULTRA_LIGHT enum
	ADSB_EMITTER_TYPE_ULTRA_LIGHT ADSB_EMITTER_TYPE = 12
	// ADSB_EMITTER_TYPE_UNASSIGNED2 enum
	ADSB_EMITTER_TYPE_UNASSIGNED2 ADSB_EMITTER_TYPE = 13
	// ADSB_EMITTER_TYPE_UAV enum
	ADSB_EMITTER_TYPE_UAV ADSB_EMITTER_TYPE = 14
	// ADSB_EMITTER_TYPE_SPACE enum
	ADSB_EMITTER_TYPE_SPACE ADSB_EMITTER_TYPE = 15
	// ADSB_EMITTER_TYPE_UNASSGINED3 enum
	ADSB_EMITTER_TYPE_UNASSGINED3 ADSB_EMITTER_TYPE = 16
	// ADSB_EMITTER_TYPE_EMERGENCY_SURFACE enum
	ADSB_EMITTER_TYPE_EMERGENCY_SURFACE ADSB_EMITTER_TYPE = 17
	// ADSB_EMITTER_TYPE_SERVICE_SURFACE enum
	ADSB_EMITTER_TYPE_SERVICE_SURFACE ADSB_EMITTER_TYPE = 18
	// ADSB_EMITTER_TYPE_POINT_OBSTACLE enum
	ADSB_EMITTER_TYPE_POINT_OBSTACLE ADSB_EMITTER_TYPE = 19
)

func (e ADSB_EMITTER_TYPE) String() string {
	if name, ok := map[ADSB_EMITTER_TYPE]string{
		ADSB_EMITTER_TYPE_NO_INFO:           "ADSB_EMITTER_TYPE_NO_INFO",
		ADSB_EMITTER_TYPE_LIGHT:             "ADSB_EMITTER_TYPE_LIGHT",
		ADSB_EMITTER_TYPE_SMALL:             "ADSB_EMITTER_TYPE_SMALL",
		ADSB_EMITTER_TYPE_LARGE:             "ADSB_EMITTER_TYPE_LARGE",
		ADSB_EMITTER_TYPE_HIGH_VORTEX_LARGE: "ADSB_EMITTER_TYPE_HIGH_VORTEX_LARGE",
		ADSB_EMITTER_TYPE_HEAVY:             "ADSB_EMITTER_TYPE_HEAVY",
		ADSB_EMITTER_TYPE_HIGHLY_MANUV:      "ADSB_EMITTER_TYPE_HIGHLY_MANUV",
		ADSB_EMITTER_TYPE_ROTOCRAFT:         "ADSB_EMITTER_TYPE_ROTOCRAFT",
		ADSB_EMITTER_TYPE_UNASSIGNED:        "ADSB_EMITTER_TYPE_UNASSIGNED",
		ADSB_EMITTER_TYPE_GLIDER:            "ADSB_EMITTER_TYPE_GLIDER",
		ADSB_EMITTER_TYPE_LIGHTER_AIR:       "ADSB_EMITTER_TYPE_LIGHTER_AIR",
		ADSB_EMITTER_TYPE_PARACHUTE:         "ADSB_EMITTER_TYPE_PARACHUTE",
		ADSB_EMITTER_TYPE_ULTRA_LIGHT:       "ADSB_EMITTER_TYPE_ULTRA_LIGHT",
		ADSB_EMITTER_TYPE_UNASSIGNED2:       "ADSB_EMITTER_TYPE_UNASSIGNED2",
		ADSB_EMITTER_TYPE_UAV:               "ADSB_EMITTER_TYPE_UAV",
		ADSB_EMITTER_TYPE_SPACE:             "ADSB_EMITTER_TYPE_SPACE",
		ADSB_EMITTER_TYPE_UNASSGINED3:       "ADSB_EMITTER_TYPE_UNASSGINED3",
		ADSB_EMITTER_TYPE_EMERGENCY_SURFACE: "ADSB_EMITTER_TYPE_EMERGENCY_SURFACE",
		ADSB_EMITTER_TYPE_SERVICE_SURFACE:   "ADSB_EMITTER_TYPE_SERVICE_SURFACE",
		ADSB_EMITTER_TYPE_POINT_OBSTACLE:    "ADSB_EMITTER_TYPE_POINT_OBSTACLE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("ADSB_EMITTER_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects ADSB_EMITTER_TYPE enums
func (e ADSB_EMITTER_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []ADSB_EMITTER_TYPE{
		ADSB_EMITTER_TYPE_NO_INFO,
		ADSB_EMITTER_TYPE_LIGHT,
		ADSB_EMITTER_TYPE_SMALL,
		ADSB_EMITTER_TYPE_LARGE,
		ADSB_EMITTER_TYPE_HIGH_VORTEX_LARGE,
		ADSB_EMITTER_TYPE_HEAVY,
		ADSB_EMITTER_TYPE_HIGHLY_MANUV,
		ADSB_EMITTER_TYPE_ROTOCRAFT,
		ADSB_EMITTER_TYPE_UNASSIGNED,
		ADSB_EMITTER_TYPE_GLIDER,
		ADSB_EMITTER_TYPE_LIGHTER_AIR,
		ADSB_EMITTER_TYPE_PARACHUTE,
		ADSB_EMITTER_TYPE_ULTRA_LIGHT,
		ADSB_EMITTER_TYPE_UNASSIGNED2,
		ADSB_EMITTER_TYPE_UAV,
		ADSB_EMITTER_TYPE_SPACE,
		ADSB_EMITTER_TYPE_UNASSGINED3,
		ADSB_EMITTER_TYPE_EMERGENCY_SURFACE,
		ADSB_EMITTER_TYPE_SERVICE_SURFACE,
		ADSB_EMITTER_TYPE_POINT_OBSTACLE,
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

// ADSB_FLAGS type. These flags indicate status such as data validity of each data source. Set = data valid
type ADSB_FLAGS int

func (e ADSB_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ADSB_FLAGS_VALID_COORDS enum
	ADSB_FLAGS_VALID_COORDS ADSB_FLAGS = 1
	// ADSB_FLAGS_VALID_ALTITUDE enum
	ADSB_FLAGS_VALID_ALTITUDE ADSB_FLAGS = 2
	// ADSB_FLAGS_VALID_HEADING enum
	ADSB_FLAGS_VALID_HEADING ADSB_FLAGS = 4
	// ADSB_FLAGS_VALID_VELOCITY enum
	ADSB_FLAGS_VALID_VELOCITY ADSB_FLAGS = 8
	// ADSB_FLAGS_VALID_CALLSIGN enum
	ADSB_FLAGS_VALID_CALLSIGN ADSB_FLAGS = 16
	// ADSB_FLAGS_VALID_SQUAWK enum
	ADSB_FLAGS_VALID_SQUAWK ADSB_FLAGS = 32
	// ADSB_FLAGS_SIMULATED enum
	ADSB_FLAGS_SIMULATED ADSB_FLAGS = 64
	// ADSB_FLAGS_VERTICAL_VELOCITY_VALID enum
	ADSB_FLAGS_VERTICAL_VELOCITY_VALID ADSB_FLAGS = 128
	// ADSB_FLAGS_BARO_VALID enum
	ADSB_FLAGS_BARO_VALID ADSB_FLAGS = 256
	// ADSB_FLAGS_SOURCE_UAT enum
	ADSB_FLAGS_SOURCE_UAT ADSB_FLAGS = 32768
)

func (e ADSB_FLAGS) String() string {
	if name, ok := map[ADSB_FLAGS]string{
		ADSB_FLAGS_VALID_COORDS:            "ADSB_FLAGS_VALID_COORDS",
		ADSB_FLAGS_VALID_ALTITUDE:          "ADSB_FLAGS_VALID_ALTITUDE",
		ADSB_FLAGS_VALID_HEADING:           "ADSB_FLAGS_VALID_HEADING",
		ADSB_FLAGS_VALID_VELOCITY:          "ADSB_FLAGS_VALID_VELOCITY",
		ADSB_FLAGS_VALID_CALLSIGN:          "ADSB_FLAGS_VALID_CALLSIGN",
		ADSB_FLAGS_VALID_SQUAWK:            "ADSB_FLAGS_VALID_SQUAWK",
		ADSB_FLAGS_SIMULATED:               "ADSB_FLAGS_SIMULATED",
		ADSB_FLAGS_VERTICAL_VELOCITY_VALID: "ADSB_FLAGS_VERTICAL_VELOCITY_VALID",
		ADSB_FLAGS_BARO_VALID:              "ADSB_FLAGS_BARO_VALID",
		ADSB_FLAGS_SOURCE_UAT:              "ADSB_FLAGS_SOURCE_UAT",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("ADSB_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects ADSB_FLAGS enums
func (e ADSB_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []ADSB_FLAGS{
		ADSB_FLAGS_VALID_COORDS,
		ADSB_FLAGS_VALID_ALTITUDE,
		ADSB_FLAGS_VALID_HEADING,
		ADSB_FLAGS_VALID_VELOCITY,
		ADSB_FLAGS_VALID_CALLSIGN,
		ADSB_FLAGS_VALID_SQUAWK,
		ADSB_FLAGS_SIMULATED,
		ADSB_FLAGS_VERTICAL_VELOCITY_VALID,
		ADSB_FLAGS_BARO_VALID,
		ADSB_FLAGS_SOURCE_UAT,
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

// MAV_DO_REPOSITION_FLAGS type. Bitmap of options for the MAV_CMD_DO_REPOSITION
type MAV_DO_REPOSITION_FLAGS int

func (e MAV_DO_REPOSITION_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_DO_REPOSITION_FLAGS_CHANGE_MODE enum. The aircraft should immediately transition into guided. This should not be set for follow me applications
	MAV_DO_REPOSITION_FLAGS_CHANGE_MODE MAV_DO_REPOSITION_FLAGS = 1
)

func (e MAV_DO_REPOSITION_FLAGS) String() string {
	if name, ok := map[MAV_DO_REPOSITION_FLAGS]string{
		MAV_DO_REPOSITION_FLAGS_CHANGE_MODE: "MAV_DO_REPOSITION_FLAGS_CHANGE_MODE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_DO_REPOSITION_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_DO_REPOSITION_FLAGS enums
func (e MAV_DO_REPOSITION_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_DO_REPOSITION_FLAGS{
		MAV_DO_REPOSITION_FLAGS_CHANGE_MODE,
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

// ESTIMATOR_STATUS_FLAGS type. Flags in ESTIMATOR_STATUS message
type ESTIMATOR_STATUS_FLAGS int

func (e ESTIMATOR_STATUS_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ESTIMATOR_ATTITUDE enum. True if the attitude estimate is good
	ESTIMATOR_ATTITUDE ESTIMATOR_STATUS_FLAGS = 1
	// ESTIMATOR_VELOCITY_HORIZ enum. True if the horizontal velocity estimate is good
	ESTIMATOR_VELOCITY_HORIZ ESTIMATOR_STATUS_FLAGS = 2
	// ESTIMATOR_VELOCITY_VERT enum. True if the  vertical velocity estimate is good
	ESTIMATOR_VELOCITY_VERT ESTIMATOR_STATUS_FLAGS = 4
	// ESTIMATOR_POS_HORIZ_REL enum. True if the horizontal position (relative) estimate is good
	ESTIMATOR_POS_HORIZ_REL ESTIMATOR_STATUS_FLAGS = 8
	// ESTIMATOR_POS_HORIZ_ABS enum. True if the horizontal position (absolute) estimate is good
	ESTIMATOR_POS_HORIZ_ABS ESTIMATOR_STATUS_FLAGS = 16
	// ESTIMATOR_POS_VERT_ABS enum. True if the vertical position (absolute) estimate is good
	ESTIMATOR_POS_VERT_ABS ESTIMATOR_STATUS_FLAGS = 32
	// ESTIMATOR_POS_VERT_AGL enum. True if the vertical position (above ground) estimate is good
	ESTIMATOR_POS_VERT_AGL ESTIMATOR_STATUS_FLAGS = 64
	// ESTIMATOR_CONST_POS_MODE enum. True if the EKF is in a constant position mode and is not using external measurements (eg GPS or optical flow)
	ESTIMATOR_CONST_POS_MODE ESTIMATOR_STATUS_FLAGS = 128
	// ESTIMATOR_PRED_POS_HORIZ_REL enum. True if the EKF has sufficient data to enter a mode that will provide a (relative) position estimate
	ESTIMATOR_PRED_POS_HORIZ_REL ESTIMATOR_STATUS_FLAGS = 256
	// ESTIMATOR_PRED_POS_HORIZ_ABS enum. True if the EKF has sufficient data to enter a mode that will provide a (absolute) position estimate
	ESTIMATOR_PRED_POS_HORIZ_ABS ESTIMATOR_STATUS_FLAGS = 512
	// ESTIMATOR_GPS_GLITCH enum. True if the EKF has detected a GPS glitch
	ESTIMATOR_GPS_GLITCH ESTIMATOR_STATUS_FLAGS = 1024
	// ESTIMATOR_ACCEL_ERROR enum. True if the EKF has detected bad accelerometer data
	ESTIMATOR_ACCEL_ERROR ESTIMATOR_STATUS_FLAGS = 2048
)

func (e ESTIMATOR_STATUS_FLAGS) String() string {
	if name, ok := map[ESTIMATOR_STATUS_FLAGS]string{
		ESTIMATOR_ATTITUDE:           "ESTIMATOR_ATTITUDE",
		ESTIMATOR_VELOCITY_HORIZ:     "ESTIMATOR_VELOCITY_HORIZ",
		ESTIMATOR_VELOCITY_VERT:      "ESTIMATOR_VELOCITY_VERT",
		ESTIMATOR_POS_HORIZ_REL:      "ESTIMATOR_POS_HORIZ_REL",
		ESTIMATOR_POS_HORIZ_ABS:      "ESTIMATOR_POS_HORIZ_ABS",
		ESTIMATOR_POS_VERT_ABS:       "ESTIMATOR_POS_VERT_ABS",
		ESTIMATOR_POS_VERT_AGL:       "ESTIMATOR_POS_VERT_AGL",
		ESTIMATOR_CONST_POS_MODE:     "ESTIMATOR_CONST_POS_MODE",
		ESTIMATOR_PRED_POS_HORIZ_REL: "ESTIMATOR_PRED_POS_HORIZ_REL",
		ESTIMATOR_PRED_POS_HORIZ_ABS: "ESTIMATOR_PRED_POS_HORIZ_ABS",
		ESTIMATOR_GPS_GLITCH:         "ESTIMATOR_GPS_GLITCH",
		ESTIMATOR_ACCEL_ERROR:        "ESTIMATOR_ACCEL_ERROR",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("ESTIMATOR_STATUS_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects ESTIMATOR_STATUS_FLAGS enums
func (e ESTIMATOR_STATUS_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []ESTIMATOR_STATUS_FLAGS{
		ESTIMATOR_ATTITUDE,
		ESTIMATOR_VELOCITY_HORIZ,
		ESTIMATOR_VELOCITY_VERT,
		ESTIMATOR_POS_HORIZ_REL,
		ESTIMATOR_POS_HORIZ_ABS,
		ESTIMATOR_POS_VERT_ABS,
		ESTIMATOR_POS_VERT_AGL,
		ESTIMATOR_CONST_POS_MODE,
		ESTIMATOR_PRED_POS_HORIZ_REL,
		ESTIMATOR_PRED_POS_HORIZ_ABS,
		ESTIMATOR_GPS_GLITCH,
		ESTIMATOR_ACCEL_ERROR,
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

// MOTOR_TEST_ORDER type
type MOTOR_TEST_ORDER int

func (e MOTOR_TEST_ORDER) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MOTOR_TEST_ORDER_DEFAULT enum. default autopilot motor test method
	MOTOR_TEST_ORDER_DEFAULT MOTOR_TEST_ORDER = 0
	// MOTOR_TEST_ORDER_SEQUENCE enum. motor numbers are specified as their index in a predefined vehicle-specific sequence
	MOTOR_TEST_ORDER_SEQUENCE MOTOR_TEST_ORDER = 1
	// MOTOR_TEST_ORDER_BOARD enum. motor numbers are specified as the output as labeled on the board
	MOTOR_TEST_ORDER_BOARD MOTOR_TEST_ORDER = 2
)

func (e MOTOR_TEST_ORDER) String() string {
	if name, ok := map[MOTOR_TEST_ORDER]string{
		MOTOR_TEST_ORDER_DEFAULT:  "MOTOR_TEST_ORDER_DEFAULT",
		MOTOR_TEST_ORDER_SEQUENCE: "MOTOR_TEST_ORDER_SEQUENCE",
		MOTOR_TEST_ORDER_BOARD:    "MOTOR_TEST_ORDER_BOARD",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MOTOR_TEST_ORDER_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MOTOR_TEST_ORDER enums
func (e MOTOR_TEST_ORDER) Bitmask() string {
	bitmap := ""
	for _, entry := range []MOTOR_TEST_ORDER{
		MOTOR_TEST_ORDER_DEFAULT,
		MOTOR_TEST_ORDER_SEQUENCE,
		MOTOR_TEST_ORDER_BOARD,
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

// MOTOR_TEST_THROTTLE_TYPE type
type MOTOR_TEST_THROTTLE_TYPE int

func (e MOTOR_TEST_THROTTLE_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MOTOR_TEST_THROTTLE_PERCENT enum. throttle as a percentage from 0 ~ 100
	MOTOR_TEST_THROTTLE_PERCENT MOTOR_TEST_THROTTLE_TYPE = 0
	// MOTOR_TEST_THROTTLE_PWM enum. throttle as an absolute PWM value (normally in range of 1000~2000)
	MOTOR_TEST_THROTTLE_PWM MOTOR_TEST_THROTTLE_TYPE = 1
	// MOTOR_TEST_THROTTLE_PILOT enum. throttle pass-through from pilot's transmitter
	MOTOR_TEST_THROTTLE_PILOT MOTOR_TEST_THROTTLE_TYPE = 2
	// MOTOR_TEST_COMPASS_CAL enum. per-motor compass calibration test
	MOTOR_TEST_COMPASS_CAL MOTOR_TEST_THROTTLE_TYPE = 3
)

func (e MOTOR_TEST_THROTTLE_TYPE) String() string {
	if name, ok := map[MOTOR_TEST_THROTTLE_TYPE]string{
		MOTOR_TEST_THROTTLE_PERCENT: "MOTOR_TEST_THROTTLE_PERCENT",
		MOTOR_TEST_THROTTLE_PWM:     "MOTOR_TEST_THROTTLE_PWM",
		MOTOR_TEST_THROTTLE_PILOT:   "MOTOR_TEST_THROTTLE_PILOT",
		MOTOR_TEST_COMPASS_CAL:      "MOTOR_TEST_COMPASS_CAL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MOTOR_TEST_THROTTLE_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MOTOR_TEST_THROTTLE_TYPE enums
func (e MOTOR_TEST_THROTTLE_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MOTOR_TEST_THROTTLE_TYPE{
		MOTOR_TEST_THROTTLE_PERCENT,
		MOTOR_TEST_THROTTLE_PWM,
		MOTOR_TEST_THROTTLE_PILOT,
		MOTOR_TEST_COMPASS_CAL,
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

// GPS_INPUT_IGNORE_FLAGS type
type GPS_INPUT_IGNORE_FLAGS int

func (e GPS_INPUT_IGNORE_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// GPS_INPUT_IGNORE_FLAG_ALT enum. ignore altitude field
	GPS_INPUT_IGNORE_FLAG_ALT GPS_INPUT_IGNORE_FLAGS = 1
	// GPS_INPUT_IGNORE_FLAG_HDOP enum. ignore hdop field
	GPS_INPUT_IGNORE_FLAG_HDOP GPS_INPUT_IGNORE_FLAGS = 2
	// GPS_INPUT_IGNORE_FLAG_VDOP enum. ignore vdop field
	GPS_INPUT_IGNORE_FLAG_VDOP GPS_INPUT_IGNORE_FLAGS = 4
	// GPS_INPUT_IGNORE_FLAG_VEL_HORIZ enum. ignore horizontal velocity field (vn and ve)
	GPS_INPUT_IGNORE_FLAG_VEL_HORIZ GPS_INPUT_IGNORE_FLAGS = 8
	// GPS_INPUT_IGNORE_FLAG_VEL_VERT enum. ignore vertical velocity field (vd)
	GPS_INPUT_IGNORE_FLAG_VEL_VERT GPS_INPUT_IGNORE_FLAGS = 16
	// GPS_INPUT_IGNORE_FLAG_SPEED_ACCURACY enum. ignore speed accuracy field
	GPS_INPUT_IGNORE_FLAG_SPEED_ACCURACY GPS_INPUT_IGNORE_FLAGS = 32
	// GPS_INPUT_IGNORE_FLAG_HORIZONTAL_ACCURACY enum. ignore horizontal accuracy field
	GPS_INPUT_IGNORE_FLAG_HORIZONTAL_ACCURACY GPS_INPUT_IGNORE_FLAGS = 64
	// GPS_INPUT_IGNORE_FLAG_VERTICAL_ACCURACY enum. ignore vertical accuracy field
	GPS_INPUT_IGNORE_FLAG_VERTICAL_ACCURACY GPS_INPUT_IGNORE_FLAGS = 128
)

func (e GPS_INPUT_IGNORE_FLAGS) String() string {
	if name, ok := map[GPS_INPUT_IGNORE_FLAGS]string{
		GPS_INPUT_IGNORE_FLAG_ALT:                 "GPS_INPUT_IGNORE_FLAG_ALT",
		GPS_INPUT_IGNORE_FLAG_HDOP:                "GPS_INPUT_IGNORE_FLAG_HDOP",
		GPS_INPUT_IGNORE_FLAG_VDOP:                "GPS_INPUT_IGNORE_FLAG_VDOP",
		GPS_INPUT_IGNORE_FLAG_VEL_HORIZ:           "GPS_INPUT_IGNORE_FLAG_VEL_HORIZ",
		GPS_INPUT_IGNORE_FLAG_VEL_VERT:            "GPS_INPUT_IGNORE_FLAG_VEL_VERT",
		GPS_INPUT_IGNORE_FLAG_SPEED_ACCURACY:      "GPS_INPUT_IGNORE_FLAG_SPEED_ACCURACY",
		GPS_INPUT_IGNORE_FLAG_HORIZONTAL_ACCURACY: "GPS_INPUT_IGNORE_FLAG_HORIZONTAL_ACCURACY",
		GPS_INPUT_IGNORE_FLAG_VERTICAL_ACCURACY:   "GPS_INPUT_IGNORE_FLAG_VERTICAL_ACCURACY",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("GPS_INPUT_IGNORE_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects GPS_INPUT_IGNORE_FLAGS enums
func (e GPS_INPUT_IGNORE_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []GPS_INPUT_IGNORE_FLAGS{
		GPS_INPUT_IGNORE_FLAG_ALT,
		GPS_INPUT_IGNORE_FLAG_HDOP,
		GPS_INPUT_IGNORE_FLAG_VDOP,
		GPS_INPUT_IGNORE_FLAG_VEL_HORIZ,
		GPS_INPUT_IGNORE_FLAG_VEL_VERT,
		GPS_INPUT_IGNORE_FLAG_SPEED_ACCURACY,
		GPS_INPUT_IGNORE_FLAG_HORIZONTAL_ACCURACY,
		GPS_INPUT_IGNORE_FLAG_VERTICAL_ACCURACY,
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

// MAV_COLLISION_ACTION type. Possible actions an aircraft can take to avoid a collision.
type MAV_COLLISION_ACTION int

func (e MAV_COLLISION_ACTION) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_COLLISION_ACTION_NONE enum. Ignore any potential collisions
	MAV_COLLISION_ACTION_NONE MAV_COLLISION_ACTION = 0
	// MAV_COLLISION_ACTION_REPORT enum. Report potential collision
	MAV_COLLISION_ACTION_REPORT MAV_COLLISION_ACTION = 1
	// MAV_COLLISION_ACTION_ASCEND_OR_DESCEND enum. Ascend or Descend to avoid threat
	MAV_COLLISION_ACTION_ASCEND_OR_DESCEND MAV_COLLISION_ACTION = 2
	// MAV_COLLISION_ACTION_MOVE_HORIZONTALLY enum. Move horizontally to avoid threat
	MAV_COLLISION_ACTION_MOVE_HORIZONTALLY MAV_COLLISION_ACTION = 3
	// MAV_COLLISION_ACTION_MOVE_PERPENDICULAR enum. Aircraft to move perpendicular to the collision's velocity vector
	MAV_COLLISION_ACTION_MOVE_PERPENDICULAR MAV_COLLISION_ACTION = 4
	// MAV_COLLISION_ACTION_RTL enum. Aircraft to fly directly back to its launch point
	MAV_COLLISION_ACTION_RTL MAV_COLLISION_ACTION = 5
	// MAV_COLLISION_ACTION_HOVER enum. Aircraft to stop in place
	MAV_COLLISION_ACTION_HOVER MAV_COLLISION_ACTION = 6
)

func (e MAV_COLLISION_ACTION) String() string {
	if name, ok := map[MAV_COLLISION_ACTION]string{
		MAV_COLLISION_ACTION_NONE:               "MAV_COLLISION_ACTION_NONE",
		MAV_COLLISION_ACTION_REPORT:             "MAV_COLLISION_ACTION_REPORT",
		MAV_COLLISION_ACTION_ASCEND_OR_DESCEND:  "MAV_COLLISION_ACTION_ASCEND_OR_DESCEND",
		MAV_COLLISION_ACTION_MOVE_HORIZONTALLY:  "MAV_COLLISION_ACTION_MOVE_HORIZONTALLY",
		MAV_COLLISION_ACTION_MOVE_PERPENDICULAR: "MAV_COLLISION_ACTION_MOVE_PERPENDICULAR",
		MAV_COLLISION_ACTION_RTL:                "MAV_COLLISION_ACTION_RTL",
		MAV_COLLISION_ACTION_HOVER:              "MAV_COLLISION_ACTION_HOVER",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_COLLISION_ACTION_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_COLLISION_ACTION enums
func (e MAV_COLLISION_ACTION) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_COLLISION_ACTION{
		MAV_COLLISION_ACTION_NONE,
		MAV_COLLISION_ACTION_REPORT,
		MAV_COLLISION_ACTION_ASCEND_OR_DESCEND,
		MAV_COLLISION_ACTION_MOVE_HORIZONTALLY,
		MAV_COLLISION_ACTION_MOVE_PERPENDICULAR,
		MAV_COLLISION_ACTION_RTL,
		MAV_COLLISION_ACTION_HOVER,
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

// MAV_COLLISION_THREAT_LEVEL type. Aircraft-rated danger from this threat.
type MAV_COLLISION_THREAT_LEVEL int

func (e MAV_COLLISION_THREAT_LEVEL) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_COLLISION_THREAT_LEVEL_NONE enum. Not a threat
	MAV_COLLISION_THREAT_LEVEL_NONE MAV_COLLISION_THREAT_LEVEL = 0
	// MAV_COLLISION_THREAT_LEVEL_LOW enum. Craft is mildly concerned about this threat
	MAV_COLLISION_THREAT_LEVEL_LOW MAV_COLLISION_THREAT_LEVEL = 1
	// MAV_COLLISION_THREAT_LEVEL_HIGH enum. Craft is panicking, and may take actions to avoid threat
	MAV_COLLISION_THREAT_LEVEL_HIGH MAV_COLLISION_THREAT_LEVEL = 2
)

func (e MAV_COLLISION_THREAT_LEVEL) String() string {
	if name, ok := map[MAV_COLLISION_THREAT_LEVEL]string{
		MAV_COLLISION_THREAT_LEVEL_NONE: "MAV_COLLISION_THREAT_LEVEL_NONE",
		MAV_COLLISION_THREAT_LEVEL_LOW:  "MAV_COLLISION_THREAT_LEVEL_LOW",
		MAV_COLLISION_THREAT_LEVEL_HIGH: "MAV_COLLISION_THREAT_LEVEL_HIGH",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_COLLISION_THREAT_LEVEL_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_COLLISION_THREAT_LEVEL enums
func (e MAV_COLLISION_THREAT_LEVEL) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_COLLISION_THREAT_LEVEL{
		MAV_COLLISION_THREAT_LEVEL_NONE,
		MAV_COLLISION_THREAT_LEVEL_LOW,
		MAV_COLLISION_THREAT_LEVEL_HIGH,
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

// MAV_COLLISION_SRC type. Source of information about this collision.
type MAV_COLLISION_SRC int

func (e MAV_COLLISION_SRC) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_COLLISION_SRC_ADSB enum. ID field references ADSB_VEHICLE packets
	MAV_COLLISION_SRC_ADSB MAV_COLLISION_SRC = 0
	// MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT enum. ID field references MAVLink SRC ID
	MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT MAV_COLLISION_SRC = 1
)

func (e MAV_COLLISION_SRC) String() string {
	if name, ok := map[MAV_COLLISION_SRC]string{
		MAV_COLLISION_SRC_ADSB:                   "MAV_COLLISION_SRC_ADSB",
		MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT: "MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_COLLISION_SRC_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_COLLISION_SRC enums
func (e MAV_COLLISION_SRC) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_COLLISION_SRC{
		MAV_COLLISION_SRC_ADSB,
		MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT,
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

// GPS_FIX_TYPE type. Type of GPS fix
type GPS_FIX_TYPE int

func (e GPS_FIX_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// GPS_FIX_TYPE_NO_GPS enum. No GPS connected
	GPS_FIX_TYPE_NO_GPS GPS_FIX_TYPE = 0
	// GPS_FIX_TYPE_NO_FIX enum. No position information, GPS is connected
	GPS_FIX_TYPE_NO_FIX GPS_FIX_TYPE = 1
	// GPS_FIX_TYPE_2D_FIX enum. 2D position
	GPS_FIX_TYPE_2D_FIX GPS_FIX_TYPE = 2
	// GPS_FIX_TYPE_3D_FIX enum. 3D position
	GPS_FIX_TYPE_3D_FIX GPS_FIX_TYPE = 3
	// GPS_FIX_TYPE_DGPS enum. DGPS/SBAS aided 3D position
	GPS_FIX_TYPE_DGPS GPS_FIX_TYPE = 4
	// GPS_FIX_TYPE_RTK_FLOAT enum. RTK float, 3D position
	GPS_FIX_TYPE_RTK_FLOAT GPS_FIX_TYPE = 5
	// GPS_FIX_TYPE_RTK_FIXED enum. RTK Fixed, 3D position
	GPS_FIX_TYPE_RTK_FIXED GPS_FIX_TYPE = 6
	// GPS_FIX_TYPE_STATIC enum. Static fixed, typically used for base stations
	GPS_FIX_TYPE_STATIC GPS_FIX_TYPE = 7
	// GPS_FIX_TYPE_PPP enum. PPP, 3D position
	GPS_FIX_TYPE_PPP GPS_FIX_TYPE = 8
)

func (e GPS_FIX_TYPE) String() string {
	if name, ok := map[GPS_FIX_TYPE]string{
		GPS_FIX_TYPE_NO_GPS:    "GPS_FIX_TYPE_NO_GPS",
		GPS_FIX_TYPE_NO_FIX:    "GPS_FIX_TYPE_NO_FIX",
		GPS_FIX_TYPE_2D_FIX:    "GPS_FIX_TYPE_2D_FIX",
		GPS_FIX_TYPE_3D_FIX:    "GPS_FIX_TYPE_3D_FIX",
		GPS_FIX_TYPE_DGPS:      "GPS_FIX_TYPE_DGPS",
		GPS_FIX_TYPE_RTK_FLOAT: "GPS_FIX_TYPE_RTK_FLOAT",
		GPS_FIX_TYPE_RTK_FIXED: "GPS_FIX_TYPE_RTK_FIXED",
		GPS_FIX_TYPE_STATIC:    "GPS_FIX_TYPE_STATIC",
		GPS_FIX_TYPE_PPP:       "GPS_FIX_TYPE_PPP",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("GPS_FIX_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects GPS_FIX_TYPE enums
func (e GPS_FIX_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []GPS_FIX_TYPE{
		GPS_FIX_TYPE_NO_GPS,
		GPS_FIX_TYPE_NO_FIX,
		GPS_FIX_TYPE_2D_FIX,
		GPS_FIX_TYPE_3D_FIX,
		GPS_FIX_TYPE_DGPS,
		GPS_FIX_TYPE_RTK_FLOAT,
		GPS_FIX_TYPE_RTK_FIXED,
		GPS_FIX_TYPE_STATIC,
		GPS_FIX_TYPE_PPP,
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

// RTK_BASELINE_COORDINATE_SYSTEM type. RTK GPS baseline coordinate system, used for RTK corrections
type RTK_BASELINE_COORDINATE_SYSTEM int

func (e RTK_BASELINE_COORDINATE_SYSTEM) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// RTK_BASELINE_COORDINATE_SYSTEM_ECEF enum. Earth-centered, Earth-fixed
	RTK_BASELINE_COORDINATE_SYSTEM_ECEF RTK_BASELINE_COORDINATE_SYSTEM = 0
	// RTK_BASELINE_COORDINATE_SYSTEM_NED enum. RTK basestation centered, north, east, down
	RTK_BASELINE_COORDINATE_SYSTEM_NED RTK_BASELINE_COORDINATE_SYSTEM = 1
)

func (e RTK_BASELINE_COORDINATE_SYSTEM) String() string {
	if name, ok := map[RTK_BASELINE_COORDINATE_SYSTEM]string{
		RTK_BASELINE_COORDINATE_SYSTEM_ECEF: "RTK_BASELINE_COORDINATE_SYSTEM_ECEF",
		RTK_BASELINE_COORDINATE_SYSTEM_NED:  "RTK_BASELINE_COORDINATE_SYSTEM_NED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("RTK_BASELINE_COORDINATE_SYSTEM_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects RTK_BASELINE_COORDINATE_SYSTEM enums
func (e RTK_BASELINE_COORDINATE_SYSTEM) Bitmask() string {
	bitmap := ""
	for _, entry := range []RTK_BASELINE_COORDINATE_SYSTEM{
		RTK_BASELINE_COORDINATE_SYSTEM_ECEF,
		RTK_BASELINE_COORDINATE_SYSTEM_NED,
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

// LANDING_TARGET_TYPE type. Type of landing target
type LANDING_TARGET_TYPE int

func (e LANDING_TARGET_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// LANDING_TARGET_TYPE_LIGHT_BEACON enum. Landing target signaled by light beacon (ex: IR-LOCK)
	LANDING_TARGET_TYPE_LIGHT_BEACON LANDING_TARGET_TYPE = 0
	// LANDING_TARGET_TYPE_RADIO_BEACON enum. Landing target signaled by radio beacon (ex: ILS, NDB)
	LANDING_TARGET_TYPE_RADIO_BEACON LANDING_TARGET_TYPE = 1
	// LANDING_TARGET_TYPE_VISION_FIDUCIAL enum. Landing target represented by a fiducial marker (ex: ARTag)
	LANDING_TARGET_TYPE_VISION_FIDUCIAL LANDING_TARGET_TYPE = 2
	// LANDING_TARGET_TYPE_VISION_OTHER enum. Landing target represented by a pre-defined visual shape/feature (ex: X-marker, H-marker, square)
	LANDING_TARGET_TYPE_VISION_OTHER LANDING_TARGET_TYPE = 3
)

func (e LANDING_TARGET_TYPE) String() string {
	if name, ok := map[LANDING_TARGET_TYPE]string{
		LANDING_TARGET_TYPE_LIGHT_BEACON:    "LANDING_TARGET_TYPE_LIGHT_BEACON",
		LANDING_TARGET_TYPE_RADIO_BEACON:    "LANDING_TARGET_TYPE_RADIO_BEACON",
		LANDING_TARGET_TYPE_VISION_FIDUCIAL: "LANDING_TARGET_TYPE_VISION_FIDUCIAL",
		LANDING_TARGET_TYPE_VISION_OTHER:    "LANDING_TARGET_TYPE_VISION_OTHER",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("LANDING_TARGET_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects LANDING_TARGET_TYPE enums
func (e LANDING_TARGET_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []LANDING_TARGET_TYPE{
		LANDING_TARGET_TYPE_LIGHT_BEACON,
		LANDING_TARGET_TYPE_RADIO_BEACON,
		LANDING_TARGET_TYPE_VISION_FIDUCIAL,
		LANDING_TARGET_TYPE_VISION_OTHER,
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

// VTOL_TRANSITION_HEADING type. Direction of VTOL transition
type VTOL_TRANSITION_HEADING int

func (e VTOL_TRANSITION_HEADING) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// VTOL_TRANSITION_HEADING_VEHICLE_DEFAULT enum. Respect the heading configuration of the vehicle
	VTOL_TRANSITION_HEADING_VEHICLE_DEFAULT VTOL_TRANSITION_HEADING = 0
	// VTOL_TRANSITION_HEADING_NEXT_WAYPOINT enum. Use the heading pointing towards the next waypoint
	VTOL_TRANSITION_HEADING_NEXT_WAYPOINT VTOL_TRANSITION_HEADING = 1
	// VTOL_TRANSITION_HEADING_TAKEOFF enum. Use the heading on takeoff (while sitting on the ground)
	VTOL_TRANSITION_HEADING_TAKEOFF VTOL_TRANSITION_HEADING = 2
	// VTOL_TRANSITION_HEADING_SPECIFIED enum. Use the specified heading in parameter 4
	VTOL_TRANSITION_HEADING_SPECIFIED VTOL_TRANSITION_HEADING = 3
	// VTOL_TRANSITION_HEADING_ANY enum. Use the current heading when reaching takeoff altitude (potentially facing the wind when weather-vaning is active)
	VTOL_TRANSITION_HEADING_ANY VTOL_TRANSITION_HEADING = 4
)

func (e VTOL_TRANSITION_HEADING) String() string {
	if name, ok := map[VTOL_TRANSITION_HEADING]string{
		VTOL_TRANSITION_HEADING_VEHICLE_DEFAULT: "VTOL_TRANSITION_HEADING_VEHICLE_DEFAULT",
		VTOL_TRANSITION_HEADING_NEXT_WAYPOINT:   "VTOL_TRANSITION_HEADING_NEXT_WAYPOINT",
		VTOL_TRANSITION_HEADING_TAKEOFF:         "VTOL_TRANSITION_HEADING_TAKEOFF",
		VTOL_TRANSITION_HEADING_SPECIFIED:       "VTOL_TRANSITION_HEADING_SPECIFIED",
		VTOL_TRANSITION_HEADING_ANY:             "VTOL_TRANSITION_HEADING_ANY",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("VTOL_TRANSITION_HEADING_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects VTOL_TRANSITION_HEADING enums
func (e VTOL_TRANSITION_HEADING) Bitmask() string {
	bitmap := ""
	for _, entry := range []VTOL_TRANSITION_HEADING{
		VTOL_TRANSITION_HEADING_VEHICLE_DEFAULT,
		VTOL_TRANSITION_HEADING_NEXT_WAYPOINT,
		VTOL_TRANSITION_HEADING_TAKEOFF,
		VTOL_TRANSITION_HEADING_SPECIFIED,
		VTOL_TRANSITION_HEADING_ANY,
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

// CAMERA_CAP_FLAGS type. Camera capability flags (Bitmap)
type CAMERA_CAP_FLAGS int

func (e CAMERA_CAP_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// CAMERA_CAP_FLAGS_CAPTURE_VIDEO enum. Camera is able to record video
	CAMERA_CAP_FLAGS_CAPTURE_VIDEO CAMERA_CAP_FLAGS = 1
	// CAMERA_CAP_FLAGS_CAPTURE_IMAGE enum. Camera is able to capture images
	CAMERA_CAP_FLAGS_CAPTURE_IMAGE CAMERA_CAP_FLAGS = 2
	// CAMERA_CAP_FLAGS_HAS_MODES enum. Camera has separate Video and Image/Photo modes (MAV_CMD_SET_CAMERA_MODE)
	CAMERA_CAP_FLAGS_HAS_MODES CAMERA_CAP_FLAGS = 4
	// CAMERA_CAP_FLAGS_CAN_CAPTURE_IMAGE_IN_VIDEO_MODE enum. Camera can capture images while in video mode
	CAMERA_CAP_FLAGS_CAN_CAPTURE_IMAGE_IN_VIDEO_MODE CAMERA_CAP_FLAGS = 8
	// CAMERA_CAP_FLAGS_CAN_CAPTURE_VIDEO_IN_IMAGE_MODE enum. Camera can capture videos while in Photo/Image mode
	CAMERA_CAP_FLAGS_CAN_CAPTURE_VIDEO_IN_IMAGE_MODE CAMERA_CAP_FLAGS = 16
	// CAMERA_CAP_FLAGS_HAS_IMAGE_SURVEY_MODE enum. Camera has image survey mode (MAV_CMD_SET_CAMERA_MODE)
	CAMERA_CAP_FLAGS_HAS_IMAGE_SURVEY_MODE CAMERA_CAP_FLAGS = 32
	// CAMERA_CAP_FLAGS_HAS_BASIC_ZOOM enum. Camera has basic zoom control (MAV_CMD_SET_CAMERA_ZOOM)
	CAMERA_CAP_FLAGS_HAS_BASIC_ZOOM CAMERA_CAP_FLAGS = 64
	// CAMERA_CAP_FLAGS_HAS_BASIC_FOCUS enum. Camera has basic focus control (MAV_CMD_SET_CAMERA_FOCUS)
	CAMERA_CAP_FLAGS_HAS_BASIC_FOCUS CAMERA_CAP_FLAGS = 128
	// CAMERA_CAP_FLAGS_HAS_VIDEO_STREAM enum. Camera has video streaming capabilities (request VIDEO_STREAM_INFORMATION with MAV_CMD_REQUEST_MESSAGE for video streaming info)
	CAMERA_CAP_FLAGS_HAS_VIDEO_STREAM CAMERA_CAP_FLAGS = 256
	// CAMERA_CAP_FLAGS_HAS_TRACKING_POINT enum. Camera supports tracking of a point on the camera view
	CAMERA_CAP_FLAGS_HAS_TRACKING_POINT CAMERA_CAP_FLAGS = 512
	// CAMERA_CAP_FLAGS_HAS_TRACKING_RECTANGLE enum. Camera supports tracking of a selection rectangle on the camera view
	CAMERA_CAP_FLAGS_HAS_TRACKING_RECTANGLE CAMERA_CAP_FLAGS = 1024
	// CAMERA_CAP_FLAGS_HAS_TRACKING_GEO_STATUS enum. Camera supports tracking geo status (CAMERA_TRACKING_GEO_STATUS)
	CAMERA_CAP_FLAGS_HAS_TRACKING_GEO_STATUS CAMERA_CAP_FLAGS = 2048
)

func (e CAMERA_CAP_FLAGS) String() string {
	if name, ok := map[CAMERA_CAP_FLAGS]string{
		CAMERA_CAP_FLAGS_CAPTURE_VIDEO:                   "CAMERA_CAP_FLAGS_CAPTURE_VIDEO",
		CAMERA_CAP_FLAGS_CAPTURE_IMAGE:                   "CAMERA_CAP_FLAGS_CAPTURE_IMAGE",
		CAMERA_CAP_FLAGS_HAS_MODES:                       "CAMERA_CAP_FLAGS_HAS_MODES",
		CAMERA_CAP_FLAGS_CAN_CAPTURE_IMAGE_IN_VIDEO_MODE: "CAMERA_CAP_FLAGS_CAN_CAPTURE_IMAGE_IN_VIDEO_MODE",
		CAMERA_CAP_FLAGS_CAN_CAPTURE_VIDEO_IN_IMAGE_MODE: "CAMERA_CAP_FLAGS_CAN_CAPTURE_VIDEO_IN_IMAGE_MODE",
		CAMERA_CAP_FLAGS_HAS_IMAGE_SURVEY_MODE:           "CAMERA_CAP_FLAGS_HAS_IMAGE_SURVEY_MODE",
		CAMERA_CAP_FLAGS_HAS_BASIC_ZOOM:                  "CAMERA_CAP_FLAGS_HAS_BASIC_ZOOM",
		CAMERA_CAP_FLAGS_HAS_BASIC_FOCUS:                 "CAMERA_CAP_FLAGS_HAS_BASIC_FOCUS",
		CAMERA_CAP_FLAGS_HAS_VIDEO_STREAM:                "CAMERA_CAP_FLAGS_HAS_VIDEO_STREAM",
		CAMERA_CAP_FLAGS_HAS_TRACKING_POINT:              "CAMERA_CAP_FLAGS_HAS_TRACKING_POINT",
		CAMERA_CAP_FLAGS_HAS_TRACKING_RECTANGLE:          "CAMERA_CAP_FLAGS_HAS_TRACKING_RECTANGLE",
		CAMERA_CAP_FLAGS_HAS_TRACKING_GEO_STATUS:         "CAMERA_CAP_FLAGS_HAS_TRACKING_GEO_STATUS",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("CAMERA_CAP_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects CAMERA_CAP_FLAGS enums
func (e CAMERA_CAP_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []CAMERA_CAP_FLAGS{
		CAMERA_CAP_FLAGS_CAPTURE_VIDEO,
		CAMERA_CAP_FLAGS_CAPTURE_IMAGE,
		CAMERA_CAP_FLAGS_HAS_MODES,
		CAMERA_CAP_FLAGS_CAN_CAPTURE_IMAGE_IN_VIDEO_MODE,
		CAMERA_CAP_FLAGS_CAN_CAPTURE_VIDEO_IN_IMAGE_MODE,
		CAMERA_CAP_FLAGS_HAS_IMAGE_SURVEY_MODE,
		CAMERA_CAP_FLAGS_HAS_BASIC_ZOOM,
		CAMERA_CAP_FLAGS_HAS_BASIC_FOCUS,
		CAMERA_CAP_FLAGS_HAS_VIDEO_STREAM,
		CAMERA_CAP_FLAGS_HAS_TRACKING_POINT,
		CAMERA_CAP_FLAGS_HAS_TRACKING_RECTANGLE,
		CAMERA_CAP_FLAGS_HAS_TRACKING_GEO_STATUS,
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

// VIDEO_STREAM_STATUS_FLAGS type. Stream status flags (Bitmap)
type VIDEO_STREAM_STATUS_FLAGS int

func (e VIDEO_STREAM_STATUS_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// VIDEO_STREAM_STATUS_FLAGS_RUNNING enum. Stream is active (running)
	VIDEO_STREAM_STATUS_FLAGS_RUNNING VIDEO_STREAM_STATUS_FLAGS = 1
	// VIDEO_STREAM_STATUS_FLAGS_THERMAL enum. Stream is thermal imaging
	VIDEO_STREAM_STATUS_FLAGS_THERMAL VIDEO_STREAM_STATUS_FLAGS = 2
)

func (e VIDEO_STREAM_STATUS_FLAGS) String() string {
	if name, ok := map[VIDEO_STREAM_STATUS_FLAGS]string{
		VIDEO_STREAM_STATUS_FLAGS_RUNNING: "VIDEO_STREAM_STATUS_FLAGS_RUNNING",
		VIDEO_STREAM_STATUS_FLAGS_THERMAL: "VIDEO_STREAM_STATUS_FLAGS_THERMAL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("VIDEO_STREAM_STATUS_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects VIDEO_STREAM_STATUS_FLAGS enums
func (e VIDEO_STREAM_STATUS_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []VIDEO_STREAM_STATUS_FLAGS{
		VIDEO_STREAM_STATUS_FLAGS_RUNNING,
		VIDEO_STREAM_STATUS_FLAGS_THERMAL,
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

// VIDEO_STREAM_TYPE type. Video stream types
type VIDEO_STREAM_TYPE int

func (e VIDEO_STREAM_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// VIDEO_STREAM_TYPE_RTSP enum. Stream is RTSP
	VIDEO_STREAM_TYPE_RTSP VIDEO_STREAM_TYPE = 0
	// VIDEO_STREAM_TYPE_RTPUDP enum. Stream is RTP UDP (URI gives the port number)
	VIDEO_STREAM_TYPE_RTPUDP VIDEO_STREAM_TYPE = 1
	// VIDEO_STREAM_TYPE_TCP_MPEG enum. Stream is MPEG on TCP
	VIDEO_STREAM_TYPE_TCP_MPEG VIDEO_STREAM_TYPE = 2
	// VIDEO_STREAM_TYPE_MPEG_TS_H264 enum. Stream is h.264 on MPEG TS (URI gives the port number)
	VIDEO_STREAM_TYPE_MPEG_TS_H264 VIDEO_STREAM_TYPE = 3
)

func (e VIDEO_STREAM_TYPE) String() string {
	if name, ok := map[VIDEO_STREAM_TYPE]string{
		VIDEO_STREAM_TYPE_RTSP:         "VIDEO_STREAM_TYPE_RTSP",
		VIDEO_STREAM_TYPE_RTPUDP:       "VIDEO_STREAM_TYPE_RTPUDP",
		VIDEO_STREAM_TYPE_TCP_MPEG:     "VIDEO_STREAM_TYPE_TCP_MPEG",
		VIDEO_STREAM_TYPE_MPEG_TS_H264: "VIDEO_STREAM_TYPE_MPEG_TS_H264",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("VIDEO_STREAM_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects VIDEO_STREAM_TYPE enums
func (e VIDEO_STREAM_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []VIDEO_STREAM_TYPE{
		VIDEO_STREAM_TYPE_RTSP,
		VIDEO_STREAM_TYPE_RTPUDP,
		VIDEO_STREAM_TYPE_TCP_MPEG,
		VIDEO_STREAM_TYPE_MPEG_TS_H264,
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

// CAMERA_TRACKING_STATUS_FLAGS type. Camera tracking status flags
type CAMERA_TRACKING_STATUS_FLAGS int

func (e CAMERA_TRACKING_STATUS_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// CAMERA_TRACKING_STATUS_FLAGS_IDLE enum. Camera is not tracking
	CAMERA_TRACKING_STATUS_FLAGS_IDLE CAMERA_TRACKING_STATUS_FLAGS = 0
	// CAMERA_TRACKING_STATUS_FLAGS_ACTIVE enum. Camera is tracking
	CAMERA_TRACKING_STATUS_FLAGS_ACTIVE CAMERA_TRACKING_STATUS_FLAGS = 1
	// CAMERA_TRACKING_STATUS_FLAGS_ERROR enum. Camera tracking in error state
	CAMERA_TRACKING_STATUS_FLAGS_ERROR CAMERA_TRACKING_STATUS_FLAGS = 2
)

func (e CAMERA_TRACKING_STATUS_FLAGS) String() string {
	if name, ok := map[CAMERA_TRACKING_STATUS_FLAGS]string{
		CAMERA_TRACKING_STATUS_FLAGS_IDLE:   "CAMERA_TRACKING_STATUS_FLAGS_IDLE",
		CAMERA_TRACKING_STATUS_FLAGS_ACTIVE: "CAMERA_TRACKING_STATUS_FLAGS_ACTIVE",
		CAMERA_TRACKING_STATUS_FLAGS_ERROR:  "CAMERA_TRACKING_STATUS_FLAGS_ERROR",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("CAMERA_TRACKING_STATUS_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects CAMERA_TRACKING_STATUS_FLAGS enums
func (e CAMERA_TRACKING_STATUS_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []CAMERA_TRACKING_STATUS_FLAGS{
		CAMERA_TRACKING_STATUS_FLAGS_IDLE,
		CAMERA_TRACKING_STATUS_FLAGS_ACTIVE,
		CAMERA_TRACKING_STATUS_FLAGS_ERROR,
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

// CAMERA_TRACKING_MODE type. Camera tracking modes
type CAMERA_TRACKING_MODE int

func (e CAMERA_TRACKING_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// CAMERA_TRACKING_NONE enum. Not tracking
	CAMERA_TRACKING_NONE CAMERA_TRACKING_MODE = 0
	// CAMERA_TRACKING_POINT enum. Target is a point
	CAMERA_TRACKING_POINT CAMERA_TRACKING_MODE = 1
	// CAMERA_TRACKING_RECTANGLE enum. Target is a rectangle
	CAMERA_TRACKING_RECTANGLE CAMERA_TRACKING_MODE = 2
)

func (e CAMERA_TRACKING_MODE) String() string {
	if name, ok := map[CAMERA_TRACKING_MODE]string{
		CAMERA_TRACKING_NONE:      "CAMERA_TRACKING_NONE",
		CAMERA_TRACKING_POINT:     "CAMERA_TRACKING_POINT",
		CAMERA_TRACKING_RECTANGLE: "CAMERA_TRACKING_RECTANGLE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("CAMERA_TRACKING_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects CAMERA_TRACKING_MODE enums
func (e CAMERA_TRACKING_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []CAMERA_TRACKING_MODE{
		CAMERA_TRACKING_NONE,
		CAMERA_TRACKING_POINT,
		CAMERA_TRACKING_RECTANGLE,
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

// CAMERA_TRACKING_TARGET_DATA type. Camera tracking target data (shows where tracked target is within image)
type CAMERA_TRACKING_TARGET_DATA int

func (e CAMERA_TRACKING_TARGET_DATA) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// CAMERA_TRACKING_TARGET_NONE enum. No target data
	CAMERA_TRACKING_TARGET_NONE CAMERA_TRACKING_TARGET_DATA = 0
	// CAMERA_TRACKING_TARGET_EMBEDDED enum. Target data embedded in image data (proprietary)
	CAMERA_TRACKING_TARGET_EMBEDDED CAMERA_TRACKING_TARGET_DATA = 1
	// CAMERA_TRACKING_TARGET_RENDERED enum. Target data rendered in image
	CAMERA_TRACKING_TARGET_RENDERED CAMERA_TRACKING_TARGET_DATA = 2
	// CAMERA_TRACKING_TARGET_IN_STATUS enum. Target data within status message (Point or Rectangle)
	CAMERA_TRACKING_TARGET_IN_STATUS CAMERA_TRACKING_TARGET_DATA = 4
)

func (e CAMERA_TRACKING_TARGET_DATA) String() string {
	if name, ok := map[CAMERA_TRACKING_TARGET_DATA]string{
		CAMERA_TRACKING_TARGET_NONE:      "CAMERA_TRACKING_TARGET_NONE",
		CAMERA_TRACKING_TARGET_EMBEDDED:  "CAMERA_TRACKING_TARGET_EMBEDDED",
		CAMERA_TRACKING_TARGET_RENDERED:  "CAMERA_TRACKING_TARGET_RENDERED",
		CAMERA_TRACKING_TARGET_IN_STATUS: "CAMERA_TRACKING_TARGET_IN_STATUS",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("CAMERA_TRACKING_TARGET_DATA_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects CAMERA_TRACKING_TARGET_DATA enums
func (e CAMERA_TRACKING_TARGET_DATA) Bitmask() string {
	bitmap := ""
	for _, entry := range []CAMERA_TRACKING_TARGET_DATA{
		CAMERA_TRACKING_TARGET_NONE,
		CAMERA_TRACKING_TARGET_EMBEDDED,
		CAMERA_TRACKING_TARGET_RENDERED,
		CAMERA_TRACKING_TARGET_IN_STATUS,
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

// CAMERA_ZOOM_TYPE type. Zoom types for MAV_CMD_SET_CAMERA_ZOOM
type CAMERA_ZOOM_TYPE int

func (e CAMERA_ZOOM_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ZOOM_TYPE_STEP enum. Zoom one step increment (-1 for wide, 1 for tele)
	ZOOM_TYPE_STEP CAMERA_ZOOM_TYPE = 0
	// ZOOM_TYPE_CONTINUOUS enum. Continuous zoom up/down until stopped (-1 for wide, 1 for tele, 0 to stop zooming)
	ZOOM_TYPE_CONTINUOUS CAMERA_ZOOM_TYPE = 1
	// ZOOM_TYPE_RANGE enum. Zoom value as proportion of full camera range (a value between 0.0 and 100.0)
	ZOOM_TYPE_RANGE CAMERA_ZOOM_TYPE = 2
	// ZOOM_TYPE_FOCAL_LENGTH enum. Zoom value/variable focal length in milimetres. Note that there is no message to get the valid zoom range of the camera, so this can type can only be used for cameras where the zoom range is known (implying that this cannot reliably be used in a GCS for an arbitrary camera)
	ZOOM_TYPE_FOCAL_LENGTH CAMERA_ZOOM_TYPE = 3
)

func (e CAMERA_ZOOM_TYPE) String() string {
	if name, ok := map[CAMERA_ZOOM_TYPE]string{
		ZOOM_TYPE_STEP:         "ZOOM_TYPE_STEP",
		ZOOM_TYPE_CONTINUOUS:   "ZOOM_TYPE_CONTINUOUS",
		ZOOM_TYPE_RANGE:        "ZOOM_TYPE_RANGE",
		ZOOM_TYPE_FOCAL_LENGTH: "ZOOM_TYPE_FOCAL_LENGTH",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("CAMERA_ZOOM_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects CAMERA_ZOOM_TYPE enums
func (e CAMERA_ZOOM_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []CAMERA_ZOOM_TYPE{
		ZOOM_TYPE_STEP,
		ZOOM_TYPE_CONTINUOUS,
		ZOOM_TYPE_RANGE,
		ZOOM_TYPE_FOCAL_LENGTH,
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

// SET_FOCUS_TYPE type. Focus types for MAV_CMD_SET_CAMERA_FOCUS
type SET_FOCUS_TYPE int

func (e SET_FOCUS_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// FOCUS_TYPE_STEP enum. Focus one step increment (-1 for focusing in, 1 for focusing out towards infinity)
	FOCUS_TYPE_STEP SET_FOCUS_TYPE = 0
	// FOCUS_TYPE_CONTINUOUS enum. Continuous focus up/down until stopped (-1 for focusing in, 1 for focusing out towards infinity, 0 to stop focusing)
	FOCUS_TYPE_CONTINUOUS SET_FOCUS_TYPE = 1
	// FOCUS_TYPE_RANGE enum. Focus value as proportion of full camera focus range (a value between 0.0 and 100.0)
	FOCUS_TYPE_RANGE SET_FOCUS_TYPE = 2
	// FOCUS_TYPE_METERS enum. Focus value in metres. Note that there is no message to get the valid focus range of the camera, so this can type can only be used for cameras where the range is known (implying that this cannot reliably be used in a GCS for an arbitrary camera)
	FOCUS_TYPE_METERS SET_FOCUS_TYPE = 3
)

func (e SET_FOCUS_TYPE) String() string {
	if name, ok := map[SET_FOCUS_TYPE]string{
		FOCUS_TYPE_STEP:       "FOCUS_TYPE_STEP",
		FOCUS_TYPE_CONTINUOUS: "FOCUS_TYPE_CONTINUOUS",
		FOCUS_TYPE_RANGE:      "FOCUS_TYPE_RANGE",
		FOCUS_TYPE_METERS:     "FOCUS_TYPE_METERS",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("SET_FOCUS_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects SET_FOCUS_TYPE enums
func (e SET_FOCUS_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []SET_FOCUS_TYPE{
		FOCUS_TYPE_STEP,
		FOCUS_TYPE_CONTINUOUS,
		FOCUS_TYPE_RANGE,
		FOCUS_TYPE_METERS,
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

// PARAM_ACK type. Result from PARAM_EXT_SET message (or a PARAM_SET within a transaction).
type PARAM_ACK int

func (e PARAM_ACK) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// PARAM_ACK_ACCEPTED enum. Parameter value ACCEPTED and SET
	PARAM_ACK_ACCEPTED PARAM_ACK = 0
	// PARAM_ACK_VALUE_UNSUPPORTED enum. Parameter value UNKNOWN/UNSUPPORTED
	PARAM_ACK_VALUE_UNSUPPORTED PARAM_ACK = 1
	// PARAM_ACK_FAILED enum. Parameter failed to set
	PARAM_ACK_FAILED PARAM_ACK = 2
	// PARAM_ACK_IN_PROGRESS enum. Parameter value received but not yet set/accepted. A subsequent PARAM_ACK_TRANSACTION or PARAM_EXT_ACK with the final result will follow once operation is completed. This is returned immediately for parameters that take longer to set, indicating taht the the parameter was recieved and does not need to be resent
	PARAM_ACK_IN_PROGRESS PARAM_ACK = 3
)

func (e PARAM_ACK) String() string {
	if name, ok := map[PARAM_ACK]string{
		PARAM_ACK_ACCEPTED:          "PARAM_ACK_ACCEPTED",
		PARAM_ACK_VALUE_UNSUPPORTED: "PARAM_ACK_VALUE_UNSUPPORTED",
		PARAM_ACK_FAILED:            "PARAM_ACK_FAILED",
		PARAM_ACK_IN_PROGRESS:       "PARAM_ACK_IN_PROGRESS",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("PARAM_ACK_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects PARAM_ACK enums
func (e PARAM_ACK) Bitmask() string {
	bitmap := ""
	for _, entry := range []PARAM_ACK{
		PARAM_ACK_ACCEPTED,
		PARAM_ACK_VALUE_UNSUPPORTED,
		PARAM_ACK_FAILED,
		PARAM_ACK_IN_PROGRESS,
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

// CAMERA_MODE type. Camera Modes.
type CAMERA_MODE int

func (e CAMERA_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// CAMERA_MODE_IMAGE enum. Camera is in image/photo capture mode
	CAMERA_MODE_IMAGE CAMERA_MODE = 0
	// CAMERA_MODE_VIDEO enum. Camera is in video capture mode
	CAMERA_MODE_VIDEO CAMERA_MODE = 1
	// CAMERA_MODE_IMAGE_SURVEY enum. Camera is in image survey capture mode. It allows for camera controller to do specific settings for surveys
	CAMERA_MODE_IMAGE_SURVEY CAMERA_MODE = 2
)

func (e CAMERA_MODE) String() string {
	if name, ok := map[CAMERA_MODE]string{
		CAMERA_MODE_IMAGE:        "CAMERA_MODE_IMAGE",
		CAMERA_MODE_VIDEO:        "CAMERA_MODE_VIDEO",
		CAMERA_MODE_IMAGE_SURVEY: "CAMERA_MODE_IMAGE_SURVEY",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("CAMERA_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects CAMERA_MODE enums
func (e CAMERA_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []CAMERA_MODE{
		CAMERA_MODE_IMAGE,
		CAMERA_MODE_VIDEO,
		CAMERA_MODE_IMAGE_SURVEY,
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

// MAV_ARM_AUTH_DENIED_REASON type
type MAV_ARM_AUTH_DENIED_REASON int

func (e MAV_ARM_AUTH_DENIED_REASON) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ARM_AUTH_DENIED_REASON_GENERIC enum. Not a specific reason
	MAV_ARM_AUTH_DENIED_REASON_GENERIC MAV_ARM_AUTH_DENIED_REASON = 0
	// MAV_ARM_AUTH_DENIED_REASON_NONE enum. Authorizer will send the error as string to GCS
	MAV_ARM_AUTH_DENIED_REASON_NONE MAV_ARM_AUTH_DENIED_REASON = 1
	// MAV_ARM_AUTH_DENIED_REASON_INVALID_WAYPOINT enum. At least one waypoint have a invalid value
	MAV_ARM_AUTH_DENIED_REASON_INVALID_WAYPOINT MAV_ARM_AUTH_DENIED_REASON = 2
	// MAV_ARM_AUTH_DENIED_REASON_TIMEOUT enum. Timeout in the authorizer process(in case it depends on network)
	MAV_ARM_AUTH_DENIED_REASON_TIMEOUT MAV_ARM_AUTH_DENIED_REASON = 3
	// MAV_ARM_AUTH_DENIED_REASON_AIRSPACE_IN_USE enum. Airspace of the mission in use by another vehicle, second result parameter can have the waypoint id that caused it to be denied
	MAV_ARM_AUTH_DENIED_REASON_AIRSPACE_IN_USE MAV_ARM_AUTH_DENIED_REASON = 4
	// MAV_ARM_AUTH_DENIED_REASON_BAD_WEATHER enum. Weather is not good to fly
	MAV_ARM_AUTH_DENIED_REASON_BAD_WEATHER MAV_ARM_AUTH_DENIED_REASON = 5
)

func (e MAV_ARM_AUTH_DENIED_REASON) String() string {
	if name, ok := map[MAV_ARM_AUTH_DENIED_REASON]string{
		MAV_ARM_AUTH_DENIED_REASON_GENERIC:          "MAV_ARM_AUTH_DENIED_REASON_GENERIC",
		MAV_ARM_AUTH_DENIED_REASON_NONE:             "MAV_ARM_AUTH_DENIED_REASON_NONE",
		MAV_ARM_AUTH_DENIED_REASON_INVALID_WAYPOINT: "MAV_ARM_AUTH_DENIED_REASON_INVALID_WAYPOINT",
		MAV_ARM_AUTH_DENIED_REASON_TIMEOUT:          "MAV_ARM_AUTH_DENIED_REASON_TIMEOUT",
		MAV_ARM_AUTH_DENIED_REASON_AIRSPACE_IN_USE:  "MAV_ARM_AUTH_DENIED_REASON_AIRSPACE_IN_USE",
		MAV_ARM_AUTH_DENIED_REASON_BAD_WEATHER:      "MAV_ARM_AUTH_DENIED_REASON_BAD_WEATHER",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ARM_AUTH_DENIED_REASON_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ARM_AUTH_DENIED_REASON enums
func (e MAV_ARM_AUTH_DENIED_REASON) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ARM_AUTH_DENIED_REASON{
		MAV_ARM_AUTH_DENIED_REASON_GENERIC,
		MAV_ARM_AUTH_DENIED_REASON_NONE,
		MAV_ARM_AUTH_DENIED_REASON_INVALID_WAYPOINT,
		MAV_ARM_AUTH_DENIED_REASON_TIMEOUT,
		MAV_ARM_AUTH_DENIED_REASON_AIRSPACE_IN_USE,
		MAV_ARM_AUTH_DENIED_REASON_BAD_WEATHER,
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

// RC_TYPE type. RC type
type RC_TYPE int

func (e RC_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// RC_TYPE_SPEKTRUM_DSM2 enum. Spektrum DSM2
	RC_TYPE_SPEKTRUM_DSM2 RC_TYPE = 0
	// RC_TYPE_SPEKTRUM_DSMX enum. Spektrum DSMX
	RC_TYPE_SPEKTRUM_DSMX RC_TYPE = 1
)

func (e RC_TYPE) String() string {
	if name, ok := map[RC_TYPE]string{
		RC_TYPE_SPEKTRUM_DSM2: "RC_TYPE_SPEKTRUM_DSM2",
		RC_TYPE_SPEKTRUM_DSMX: "RC_TYPE_SPEKTRUM_DSMX",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("RC_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects RC_TYPE enums
func (e RC_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []RC_TYPE{
		RC_TYPE_SPEKTRUM_DSM2,
		RC_TYPE_SPEKTRUM_DSMX,
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

// POSITION_TARGET_TYPEMASK type. Bitmap to indicate which dimensions should be ignored by the vehicle: a value of 0b0000000000000000 or 0b0000001000000000 indicates that none of the setpoint dimensions should be ignored. If bit 9 is set the floats afx afy afz should be interpreted as force instead of acceleration.
type POSITION_TARGET_TYPEMASK int

func (e POSITION_TARGET_TYPEMASK) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// POSITION_TARGET_TYPEMASK_X_IGNORE enum. Ignore position x
	POSITION_TARGET_TYPEMASK_X_IGNORE POSITION_TARGET_TYPEMASK = 1
	// POSITION_TARGET_TYPEMASK_Y_IGNORE enum. Ignore position y
	POSITION_TARGET_TYPEMASK_Y_IGNORE POSITION_TARGET_TYPEMASK = 2
	// POSITION_TARGET_TYPEMASK_Z_IGNORE enum. Ignore position z
	POSITION_TARGET_TYPEMASK_Z_IGNORE POSITION_TARGET_TYPEMASK = 4
	// POSITION_TARGET_TYPEMASK_VX_IGNORE enum. Ignore velocity x
	POSITION_TARGET_TYPEMASK_VX_IGNORE POSITION_TARGET_TYPEMASK = 8
	// POSITION_TARGET_TYPEMASK_VY_IGNORE enum. Ignore velocity y
	POSITION_TARGET_TYPEMASK_VY_IGNORE POSITION_TARGET_TYPEMASK = 16
	// POSITION_TARGET_TYPEMASK_VZ_IGNORE enum. Ignore velocity z
	POSITION_TARGET_TYPEMASK_VZ_IGNORE POSITION_TARGET_TYPEMASK = 32
	// POSITION_TARGET_TYPEMASK_AX_IGNORE enum. Ignore acceleration x
	POSITION_TARGET_TYPEMASK_AX_IGNORE POSITION_TARGET_TYPEMASK = 64
	// POSITION_TARGET_TYPEMASK_AY_IGNORE enum. Ignore acceleration y
	POSITION_TARGET_TYPEMASK_AY_IGNORE POSITION_TARGET_TYPEMASK = 128
	// POSITION_TARGET_TYPEMASK_AZ_IGNORE enum. Ignore acceleration z
	POSITION_TARGET_TYPEMASK_AZ_IGNORE POSITION_TARGET_TYPEMASK = 256
	// POSITION_TARGET_TYPEMASK_FORCE_SET enum. Use force instead of acceleration
	POSITION_TARGET_TYPEMASK_FORCE_SET POSITION_TARGET_TYPEMASK = 512
	// POSITION_TARGET_TYPEMASK_YAW_IGNORE enum. Ignore yaw
	POSITION_TARGET_TYPEMASK_YAW_IGNORE POSITION_TARGET_TYPEMASK = 1024
	// POSITION_TARGET_TYPEMASK_YAW_RATE_IGNORE enum. Ignore yaw rate
	POSITION_TARGET_TYPEMASK_YAW_RATE_IGNORE POSITION_TARGET_TYPEMASK = 2048
)

func (e POSITION_TARGET_TYPEMASK) String() string {
	if name, ok := map[POSITION_TARGET_TYPEMASK]string{
		POSITION_TARGET_TYPEMASK_X_IGNORE:        "POSITION_TARGET_TYPEMASK_X_IGNORE",
		POSITION_TARGET_TYPEMASK_Y_IGNORE:        "POSITION_TARGET_TYPEMASK_Y_IGNORE",
		POSITION_TARGET_TYPEMASK_Z_IGNORE:        "POSITION_TARGET_TYPEMASK_Z_IGNORE",
		POSITION_TARGET_TYPEMASK_VX_IGNORE:       "POSITION_TARGET_TYPEMASK_VX_IGNORE",
		POSITION_TARGET_TYPEMASK_VY_IGNORE:       "POSITION_TARGET_TYPEMASK_VY_IGNORE",
		POSITION_TARGET_TYPEMASK_VZ_IGNORE:       "POSITION_TARGET_TYPEMASK_VZ_IGNORE",
		POSITION_TARGET_TYPEMASK_AX_IGNORE:       "POSITION_TARGET_TYPEMASK_AX_IGNORE",
		POSITION_TARGET_TYPEMASK_AY_IGNORE:       "POSITION_TARGET_TYPEMASK_AY_IGNORE",
		POSITION_TARGET_TYPEMASK_AZ_IGNORE:       "POSITION_TARGET_TYPEMASK_AZ_IGNORE",
		POSITION_TARGET_TYPEMASK_FORCE_SET:       "POSITION_TARGET_TYPEMASK_FORCE_SET",
		POSITION_TARGET_TYPEMASK_YAW_IGNORE:      "POSITION_TARGET_TYPEMASK_YAW_IGNORE",
		POSITION_TARGET_TYPEMASK_YAW_RATE_IGNORE: "POSITION_TARGET_TYPEMASK_YAW_RATE_IGNORE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("POSITION_TARGET_TYPEMASK_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects POSITION_TARGET_TYPEMASK enums
func (e POSITION_TARGET_TYPEMASK) Bitmask() string {
	bitmap := ""
	for _, entry := range []POSITION_TARGET_TYPEMASK{
		POSITION_TARGET_TYPEMASK_X_IGNORE,
		POSITION_TARGET_TYPEMASK_Y_IGNORE,
		POSITION_TARGET_TYPEMASK_Z_IGNORE,
		POSITION_TARGET_TYPEMASK_VX_IGNORE,
		POSITION_TARGET_TYPEMASK_VY_IGNORE,
		POSITION_TARGET_TYPEMASK_VZ_IGNORE,
		POSITION_TARGET_TYPEMASK_AX_IGNORE,
		POSITION_TARGET_TYPEMASK_AY_IGNORE,
		POSITION_TARGET_TYPEMASK_AZ_IGNORE,
		POSITION_TARGET_TYPEMASK_FORCE_SET,
		POSITION_TARGET_TYPEMASK_YAW_IGNORE,
		POSITION_TARGET_TYPEMASK_YAW_RATE_IGNORE,
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

// ATTITUDE_TARGET_TYPEMASK type. Bitmap to indicate which dimensions should be ignored by the vehicle: a value of 0b00000000 indicates that none of the setpoint dimensions should be ignored.
type ATTITUDE_TARGET_TYPEMASK int

func (e ATTITUDE_TARGET_TYPEMASK) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// ATTITUDE_TARGET_TYPEMASK_BODY_ROLL_RATE_IGNORE enum. Ignore body roll rate
	ATTITUDE_TARGET_TYPEMASK_BODY_ROLL_RATE_IGNORE ATTITUDE_TARGET_TYPEMASK = 1
	// ATTITUDE_TARGET_TYPEMASK_BODY_PITCH_RATE_IGNORE enum. Ignore body pitch rate
	ATTITUDE_TARGET_TYPEMASK_BODY_PITCH_RATE_IGNORE ATTITUDE_TARGET_TYPEMASK = 2
	// ATTITUDE_TARGET_TYPEMASK_BODY_YAW_RATE_IGNORE enum. Ignore body yaw rate
	ATTITUDE_TARGET_TYPEMASK_BODY_YAW_RATE_IGNORE ATTITUDE_TARGET_TYPEMASK = 4
	// ATTITUDE_TARGET_TYPEMASK_THROTTLE_IGNORE enum. Ignore throttle
	ATTITUDE_TARGET_TYPEMASK_THROTTLE_IGNORE ATTITUDE_TARGET_TYPEMASK = 64
	// ATTITUDE_TARGET_TYPEMASK_ATTITUDE_IGNORE enum. Ignore attitude
	ATTITUDE_TARGET_TYPEMASK_ATTITUDE_IGNORE ATTITUDE_TARGET_TYPEMASK = 128
)

func (e ATTITUDE_TARGET_TYPEMASK) String() string {
	if name, ok := map[ATTITUDE_TARGET_TYPEMASK]string{
		ATTITUDE_TARGET_TYPEMASK_BODY_ROLL_RATE_IGNORE:  "ATTITUDE_TARGET_TYPEMASK_BODY_ROLL_RATE_IGNORE",
		ATTITUDE_TARGET_TYPEMASK_BODY_PITCH_RATE_IGNORE: "ATTITUDE_TARGET_TYPEMASK_BODY_PITCH_RATE_IGNORE",
		ATTITUDE_TARGET_TYPEMASK_BODY_YAW_RATE_IGNORE:   "ATTITUDE_TARGET_TYPEMASK_BODY_YAW_RATE_IGNORE",
		ATTITUDE_TARGET_TYPEMASK_THROTTLE_IGNORE:        "ATTITUDE_TARGET_TYPEMASK_THROTTLE_IGNORE",
		ATTITUDE_TARGET_TYPEMASK_ATTITUDE_IGNORE:        "ATTITUDE_TARGET_TYPEMASK_ATTITUDE_IGNORE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("ATTITUDE_TARGET_TYPEMASK_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects ATTITUDE_TARGET_TYPEMASK enums
func (e ATTITUDE_TARGET_TYPEMASK) Bitmask() string {
	bitmap := ""
	for _, entry := range []ATTITUDE_TARGET_TYPEMASK{
		ATTITUDE_TARGET_TYPEMASK_BODY_ROLL_RATE_IGNORE,
		ATTITUDE_TARGET_TYPEMASK_BODY_PITCH_RATE_IGNORE,
		ATTITUDE_TARGET_TYPEMASK_BODY_YAW_RATE_IGNORE,
		ATTITUDE_TARGET_TYPEMASK_THROTTLE_IGNORE,
		ATTITUDE_TARGET_TYPEMASK_ATTITUDE_IGNORE,
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

// UTM_FLIGHT_STATE type. Airborne status of UAS.
type UTM_FLIGHT_STATE int

func (e UTM_FLIGHT_STATE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// UTM_FLIGHT_STATE_UNKNOWN enum. The flight state can't be determined
	UTM_FLIGHT_STATE_UNKNOWN UTM_FLIGHT_STATE = 1
	// UTM_FLIGHT_STATE_GROUND enum. UAS on ground
	UTM_FLIGHT_STATE_GROUND UTM_FLIGHT_STATE = 2
	// UTM_FLIGHT_STATE_AIRBORNE enum. UAS airborne
	UTM_FLIGHT_STATE_AIRBORNE UTM_FLIGHT_STATE = 3
	// UTM_FLIGHT_STATE_EMERGENCY enum. UAS is in an emergency flight state
	UTM_FLIGHT_STATE_EMERGENCY UTM_FLIGHT_STATE = 16
	// UTM_FLIGHT_STATE_NOCTRL enum. UAS has no active controls
	UTM_FLIGHT_STATE_NOCTRL UTM_FLIGHT_STATE = 32
)

func (e UTM_FLIGHT_STATE) String() string {
	if name, ok := map[UTM_FLIGHT_STATE]string{
		UTM_FLIGHT_STATE_UNKNOWN:   "UTM_FLIGHT_STATE_UNKNOWN",
		UTM_FLIGHT_STATE_GROUND:    "UTM_FLIGHT_STATE_GROUND",
		UTM_FLIGHT_STATE_AIRBORNE:  "UTM_FLIGHT_STATE_AIRBORNE",
		UTM_FLIGHT_STATE_EMERGENCY: "UTM_FLIGHT_STATE_EMERGENCY",
		UTM_FLIGHT_STATE_NOCTRL:    "UTM_FLIGHT_STATE_NOCTRL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("UTM_FLIGHT_STATE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects UTM_FLIGHT_STATE enums
func (e UTM_FLIGHT_STATE) Bitmask() string {
	bitmap := ""
	for _, entry := range []UTM_FLIGHT_STATE{
		UTM_FLIGHT_STATE_UNKNOWN,
		UTM_FLIGHT_STATE_GROUND,
		UTM_FLIGHT_STATE_AIRBORNE,
		UTM_FLIGHT_STATE_EMERGENCY,
		UTM_FLIGHT_STATE_NOCTRL,
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

// UTM_DATA_AVAIL_FLAGS type. Flags for the global position report.
type UTM_DATA_AVAIL_FLAGS int

func (e UTM_DATA_AVAIL_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// UTM_DATA_AVAIL_FLAGS_TIME_VALID enum. The field time contains valid data
	UTM_DATA_AVAIL_FLAGS_TIME_VALID UTM_DATA_AVAIL_FLAGS = 1
	// UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE enum. The field uas_id contains valid data
	UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE UTM_DATA_AVAIL_FLAGS = 2
	// UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE enum. The fields lat, lon and h_acc contain valid data
	UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE UTM_DATA_AVAIL_FLAGS = 4
	// UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE enum. The fields alt and v_acc contain valid data
	UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE UTM_DATA_AVAIL_FLAGS = 8
	// UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE enum. The field relative_alt contains valid data
	UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE UTM_DATA_AVAIL_FLAGS = 16
	// UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE enum. The fields vx and vy contain valid data
	UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE UTM_DATA_AVAIL_FLAGS = 32
	// UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE enum. The field vz contains valid data
	UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE UTM_DATA_AVAIL_FLAGS = 64
	// UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE enum. The fields next_lat, next_lon and next_alt contain valid data
	UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE UTM_DATA_AVAIL_FLAGS = 128
)

func (e UTM_DATA_AVAIL_FLAGS) String() string {
	if name, ok := map[UTM_DATA_AVAIL_FLAGS]string{
		UTM_DATA_AVAIL_FLAGS_TIME_VALID:                  "UTM_DATA_AVAIL_FLAGS_TIME_VALID",
		UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE:            "UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE",
		UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE:          "UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE",
		UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE:          "UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE",
		UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE: "UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE",
		UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE:   "UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE",
		UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE:     "UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE",
		UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE:     "UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("UTM_DATA_AVAIL_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects UTM_DATA_AVAIL_FLAGS enums
func (e UTM_DATA_AVAIL_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []UTM_DATA_AVAIL_FLAGS{
		UTM_DATA_AVAIL_FLAGS_TIME_VALID,
		UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE,
		UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE,
		UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE,
		UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE,
		UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE,
		UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE,
		UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE,
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

// CELLULAR_NETWORK_RADIO_TYPE type. Cellular network radio type
type CELLULAR_NETWORK_RADIO_TYPE int

func (e CELLULAR_NETWORK_RADIO_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// CELLULAR_NETWORK_RADIO_TYPE_NONE enum
	CELLULAR_NETWORK_RADIO_TYPE_NONE CELLULAR_NETWORK_RADIO_TYPE = 0
	// CELLULAR_NETWORK_RADIO_TYPE_GSM enum
	CELLULAR_NETWORK_RADIO_TYPE_GSM CELLULAR_NETWORK_RADIO_TYPE = 1
	// CELLULAR_NETWORK_RADIO_TYPE_CDMA enum
	CELLULAR_NETWORK_RADIO_TYPE_CDMA CELLULAR_NETWORK_RADIO_TYPE = 2
	// CELLULAR_NETWORK_RADIO_TYPE_WCDMA enum
	CELLULAR_NETWORK_RADIO_TYPE_WCDMA CELLULAR_NETWORK_RADIO_TYPE = 3
	// CELLULAR_NETWORK_RADIO_TYPE_LTE enum
	CELLULAR_NETWORK_RADIO_TYPE_LTE CELLULAR_NETWORK_RADIO_TYPE = 4
)

func (e CELLULAR_NETWORK_RADIO_TYPE) String() string {
	if name, ok := map[CELLULAR_NETWORK_RADIO_TYPE]string{
		CELLULAR_NETWORK_RADIO_TYPE_NONE:  "CELLULAR_NETWORK_RADIO_TYPE_NONE",
		CELLULAR_NETWORK_RADIO_TYPE_GSM:   "CELLULAR_NETWORK_RADIO_TYPE_GSM",
		CELLULAR_NETWORK_RADIO_TYPE_CDMA:  "CELLULAR_NETWORK_RADIO_TYPE_CDMA",
		CELLULAR_NETWORK_RADIO_TYPE_WCDMA: "CELLULAR_NETWORK_RADIO_TYPE_WCDMA",
		CELLULAR_NETWORK_RADIO_TYPE_LTE:   "CELLULAR_NETWORK_RADIO_TYPE_LTE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("CELLULAR_NETWORK_RADIO_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects CELLULAR_NETWORK_RADIO_TYPE enums
func (e CELLULAR_NETWORK_RADIO_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []CELLULAR_NETWORK_RADIO_TYPE{
		CELLULAR_NETWORK_RADIO_TYPE_NONE,
		CELLULAR_NETWORK_RADIO_TYPE_GSM,
		CELLULAR_NETWORK_RADIO_TYPE_CDMA,
		CELLULAR_NETWORK_RADIO_TYPE_WCDMA,
		CELLULAR_NETWORK_RADIO_TYPE_LTE,
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

// CELLULAR_STATUS_FLAG type. These flags encode the cellular network status
type CELLULAR_STATUS_FLAG int

func (e CELLULAR_STATUS_FLAG) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// CELLULAR_STATUS_FLAG_UNKNOWN enum. State unknown or not reportable
	CELLULAR_STATUS_FLAG_UNKNOWN CELLULAR_STATUS_FLAG = 0
	// CELLULAR_STATUS_FLAG_FAILED enum. Modem is unusable
	CELLULAR_STATUS_FLAG_FAILED CELLULAR_STATUS_FLAG = 1
	// CELLULAR_STATUS_FLAG_INITIALIZING enum. Modem is being initialized
	CELLULAR_STATUS_FLAG_INITIALIZING CELLULAR_STATUS_FLAG = 2
	// CELLULAR_STATUS_FLAG_LOCKED enum. Modem is locked
	CELLULAR_STATUS_FLAG_LOCKED CELLULAR_STATUS_FLAG = 3
	// CELLULAR_STATUS_FLAG_DISABLED enum. Modem is not enabled and is powered down
	CELLULAR_STATUS_FLAG_DISABLED CELLULAR_STATUS_FLAG = 4
	// CELLULAR_STATUS_FLAG_DISABLING enum. Modem is currently transitioning to the CELLULAR_STATUS_FLAG_DISABLED state
	CELLULAR_STATUS_FLAG_DISABLING CELLULAR_STATUS_FLAG = 5
	// CELLULAR_STATUS_FLAG_ENABLING enum. Modem is currently transitioning to the CELLULAR_STATUS_FLAG_ENABLED state
	CELLULAR_STATUS_FLAG_ENABLING CELLULAR_STATUS_FLAG = 6
	// CELLULAR_STATUS_FLAG_ENABLED enum. Modem is enabled and powered on but not registered with a network provider and not available for data connections
	CELLULAR_STATUS_FLAG_ENABLED CELLULAR_STATUS_FLAG = 7
	// CELLULAR_STATUS_FLAG_SEARCHING enum. Modem is searching for a network provider to register
	CELLULAR_STATUS_FLAG_SEARCHING CELLULAR_STATUS_FLAG = 8
	// CELLULAR_STATUS_FLAG_REGISTERED enum. Modem is registered with a network provider, and data connections and messaging may be available for use
	CELLULAR_STATUS_FLAG_REGISTERED CELLULAR_STATUS_FLAG = 9
	// CELLULAR_STATUS_FLAG_DISCONNECTING enum. Modem is disconnecting and deactivating the last active packet data bearer. This state will not be entered if more than one packet data bearer is active and one of the active bearers is deactivated
	CELLULAR_STATUS_FLAG_DISCONNECTING CELLULAR_STATUS_FLAG = 10
	// CELLULAR_STATUS_FLAG_CONNECTING enum. Modem is activating and connecting the first packet data bearer. Subsequent bearer activations when another bearer is already active do not cause this state to be entered
	CELLULAR_STATUS_FLAG_CONNECTING CELLULAR_STATUS_FLAG = 11
	// CELLULAR_STATUS_FLAG_CONNECTED enum. One or more packet data bearers is active and connected
	CELLULAR_STATUS_FLAG_CONNECTED CELLULAR_STATUS_FLAG = 12
)

func (e CELLULAR_STATUS_FLAG) String() string {
	if name, ok := map[CELLULAR_STATUS_FLAG]string{
		CELLULAR_STATUS_FLAG_UNKNOWN:       "CELLULAR_STATUS_FLAG_UNKNOWN",
		CELLULAR_STATUS_FLAG_FAILED:        "CELLULAR_STATUS_FLAG_FAILED",
		CELLULAR_STATUS_FLAG_INITIALIZING:  "CELLULAR_STATUS_FLAG_INITIALIZING",
		CELLULAR_STATUS_FLAG_LOCKED:        "CELLULAR_STATUS_FLAG_LOCKED",
		CELLULAR_STATUS_FLAG_DISABLED:      "CELLULAR_STATUS_FLAG_DISABLED",
		CELLULAR_STATUS_FLAG_DISABLING:     "CELLULAR_STATUS_FLAG_DISABLING",
		CELLULAR_STATUS_FLAG_ENABLING:      "CELLULAR_STATUS_FLAG_ENABLING",
		CELLULAR_STATUS_FLAG_ENABLED:       "CELLULAR_STATUS_FLAG_ENABLED",
		CELLULAR_STATUS_FLAG_SEARCHING:     "CELLULAR_STATUS_FLAG_SEARCHING",
		CELLULAR_STATUS_FLAG_REGISTERED:    "CELLULAR_STATUS_FLAG_REGISTERED",
		CELLULAR_STATUS_FLAG_DISCONNECTING: "CELLULAR_STATUS_FLAG_DISCONNECTING",
		CELLULAR_STATUS_FLAG_CONNECTING:    "CELLULAR_STATUS_FLAG_CONNECTING",
		CELLULAR_STATUS_FLAG_CONNECTED:     "CELLULAR_STATUS_FLAG_CONNECTED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("CELLULAR_STATUS_FLAG_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects CELLULAR_STATUS_FLAG enums
func (e CELLULAR_STATUS_FLAG) Bitmask() string {
	bitmap := ""
	for _, entry := range []CELLULAR_STATUS_FLAG{
		CELLULAR_STATUS_FLAG_UNKNOWN,
		CELLULAR_STATUS_FLAG_FAILED,
		CELLULAR_STATUS_FLAG_INITIALIZING,
		CELLULAR_STATUS_FLAG_LOCKED,
		CELLULAR_STATUS_FLAG_DISABLED,
		CELLULAR_STATUS_FLAG_DISABLING,
		CELLULAR_STATUS_FLAG_ENABLING,
		CELLULAR_STATUS_FLAG_ENABLED,
		CELLULAR_STATUS_FLAG_SEARCHING,
		CELLULAR_STATUS_FLAG_REGISTERED,
		CELLULAR_STATUS_FLAG_DISCONNECTING,
		CELLULAR_STATUS_FLAG_CONNECTING,
		CELLULAR_STATUS_FLAG_CONNECTED,
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

// CELLULAR_NETWORK_FAILED_REASON type. These flags are used to diagnose the failure state of CELLULAR_STATUS
type CELLULAR_NETWORK_FAILED_REASON int

func (e CELLULAR_NETWORK_FAILED_REASON) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// CELLULAR_NETWORK_FAILED_REASON_NONE enum. No error
	CELLULAR_NETWORK_FAILED_REASON_NONE CELLULAR_NETWORK_FAILED_REASON = 0
	// CELLULAR_NETWORK_FAILED_REASON_UNKNOWN enum. Error state is unknown
	CELLULAR_NETWORK_FAILED_REASON_UNKNOWN CELLULAR_NETWORK_FAILED_REASON = 1
	// CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING enum. SIM is required for the modem but missing
	CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING CELLULAR_NETWORK_FAILED_REASON = 2
	// CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR enum. SIM is available, but not usuable for connection
	CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR CELLULAR_NETWORK_FAILED_REASON = 3
)

func (e CELLULAR_NETWORK_FAILED_REASON) String() string {
	if name, ok := map[CELLULAR_NETWORK_FAILED_REASON]string{
		CELLULAR_NETWORK_FAILED_REASON_NONE:        "CELLULAR_NETWORK_FAILED_REASON_NONE",
		CELLULAR_NETWORK_FAILED_REASON_UNKNOWN:     "CELLULAR_NETWORK_FAILED_REASON_UNKNOWN",
		CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING: "CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING",
		CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR:   "CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("CELLULAR_NETWORK_FAILED_REASON_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects CELLULAR_NETWORK_FAILED_REASON enums
func (e CELLULAR_NETWORK_FAILED_REASON) Bitmask() string {
	bitmap := ""
	for _, entry := range []CELLULAR_NETWORK_FAILED_REASON{
		CELLULAR_NETWORK_FAILED_REASON_NONE,
		CELLULAR_NETWORK_FAILED_REASON_UNKNOWN,
		CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING,
		CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR,
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

// PRECISION_LAND_MODE type. Precision land modes (used in MAV_CMD_NAV_LAND).
type PRECISION_LAND_MODE int

func (e PRECISION_LAND_MODE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// PRECISION_LAND_MODE_DISABLED enum. Normal (non-precision) landing
	PRECISION_LAND_MODE_DISABLED PRECISION_LAND_MODE = 0
	// PRECISION_LAND_MODE_OPPORTUNISTIC enum. Use precision landing if beacon detected when land command accepted, otherwise land normally
	PRECISION_LAND_MODE_OPPORTUNISTIC PRECISION_LAND_MODE = 1
	// PRECISION_LAND_MODE_REQUIRED enum. Use precision landing, searching for beacon if not found when land command accepted (land normally if beacon cannot be found)
	PRECISION_LAND_MODE_REQUIRED PRECISION_LAND_MODE = 2
)

func (e PRECISION_LAND_MODE) String() string {
	if name, ok := map[PRECISION_LAND_MODE]string{
		PRECISION_LAND_MODE_DISABLED:      "PRECISION_LAND_MODE_DISABLED",
		PRECISION_LAND_MODE_OPPORTUNISTIC: "PRECISION_LAND_MODE_OPPORTUNISTIC",
		PRECISION_LAND_MODE_REQUIRED:      "PRECISION_LAND_MODE_REQUIRED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("PRECISION_LAND_MODE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects PRECISION_LAND_MODE enums
func (e PRECISION_LAND_MODE) Bitmask() string {
	bitmap := ""
	for _, entry := range []PRECISION_LAND_MODE{
		PRECISION_LAND_MODE_DISABLED,
		PRECISION_LAND_MODE_OPPORTUNISTIC,
		PRECISION_LAND_MODE_REQUIRED,
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

// PARACHUTE_ACTION type. Parachute actions. Trigger release and enable/disable auto-release.
type PARACHUTE_ACTION int

func (e PARACHUTE_ACTION) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// PARACHUTE_DISABLE enum. Disable auto-release of parachute (i.e. release triggered by crash detectors)
	PARACHUTE_DISABLE PARACHUTE_ACTION = 0
	// PARACHUTE_ENABLE enum. Enable auto-release of parachute
	PARACHUTE_ENABLE PARACHUTE_ACTION = 1
	// PARACHUTE_RELEASE enum. Release parachute and kill motors
	PARACHUTE_RELEASE PARACHUTE_ACTION = 2
)

func (e PARACHUTE_ACTION) String() string {
	if name, ok := map[PARACHUTE_ACTION]string{
		PARACHUTE_DISABLE: "PARACHUTE_DISABLE",
		PARACHUTE_ENABLE:  "PARACHUTE_ENABLE",
		PARACHUTE_RELEASE: "PARACHUTE_RELEASE",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("PARACHUTE_ACTION_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects PARACHUTE_ACTION enums
func (e PARACHUTE_ACTION) Bitmask() string {
	bitmap := ""
	for _, entry := range []PARACHUTE_ACTION{
		PARACHUTE_DISABLE,
		PARACHUTE_ENABLE,
		PARACHUTE_RELEASE,
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

// MAV_TUNNEL_PAYLOAD_TYPE type
type MAV_TUNNEL_PAYLOAD_TYPE int

func (e MAV_TUNNEL_PAYLOAD_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_TUNNEL_PAYLOAD_TYPE_UNKNOWN enum. Encoding of payload unknown
	MAV_TUNNEL_PAYLOAD_TYPE_UNKNOWN MAV_TUNNEL_PAYLOAD_TYPE = 0
	// MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED0 enum. Registered for STorM32 gimbal controller
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED0 MAV_TUNNEL_PAYLOAD_TYPE = 200
	// MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED1 enum. Registered for STorM32 gimbal controller
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED1 MAV_TUNNEL_PAYLOAD_TYPE = 201
	// MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED2 enum. Registered for STorM32 gimbal controller
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED2 MAV_TUNNEL_PAYLOAD_TYPE = 202
	// MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED3 enum. Registered for STorM32 gimbal controller
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED3 MAV_TUNNEL_PAYLOAD_TYPE = 203
	// MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED4 enum. Registered for STorM32 gimbal controller
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED4 MAV_TUNNEL_PAYLOAD_TYPE = 204
	// MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED5 enum. Registered for STorM32 gimbal controller
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED5 MAV_TUNNEL_PAYLOAD_TYPE = 205
	// MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED6 enum. Registered for STorM32 gimbal controller
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED6 MAV_TUNNEL_PAYLOAD_TYPE = 206
	// MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED7 enum. Registered for STorM32 gimbal controller
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED7 MAV_TUNNEL_PAYLOAD_TYPE = 207
	// MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED8 enum. Registered for STorM32 gimbal controller
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED8 MAV_TUNNEL_PAYLOAD_TYPE = 208
	// MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED9 enum. Registered for STorM32 gimbal controller
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED9 MAV_TUNNEL_PAYLOAD_TYPE = 209
)

func (e MAV_TUNNEL_PAYLOAD_TYPE) String() string {
	if name, ok := map[MAV_TUNNEL_PAYLOAD_TYPE]string{
		MAV_TUNNEL_PAYLOAD_TYPE_UNKNOWN:           "MAV_TUNNEL_PAYLOAD_TYPE_UNKNOWN",
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED0: "MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED0",
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED1: "MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED1",
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED2: "MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED2",
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED3: "MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED3",
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED4: "MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED4",
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED5: "MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED5",
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED6: "MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED6",
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED7: "MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED7",
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED8: "MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED8",
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED9: "MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED9",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_TUNNEL_PAYLOAD_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_TUNNEL_PAYLOAD_TYPE enums
func (e MAV_TUNNEL_PAYLOAD_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_TUNNEL_PAYLOAD_TYPE{
		MAV_TUNNEL_PAYLOAD_TYPE_UNKNOWN,
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED0,
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED1,
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED2,
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED3,
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED4,
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED5,
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED6,
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED7,
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED8,
		MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED9,
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

// MAV_ODID_ID_TYPE type
type MAV_ODID_ID_TYPE int

func (e MAV_ODID_ID_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_ID_TYPE_NONE enum. No type defined
	MAV_ODID_ID_TYPE_NONE MAV_ODID_ID_TYPE = 0
	// MAV_ODID_ID_TYPE_SERIAL_NUMBER enum. Manufacturer Serial Number (ANSI/CTA-2063 format)
	MAV_ODID_ID_TYPE_SERIAL_NUMBER MAV_ODID_ID_TYPE = 1
	// MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID enum. CAA (Civil Aviation Authority) registered ID. Format: [ICAO Country Code].[CAA Assigned ID]
	MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID MAV_ODID_ID_TYPE = 2
	// MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID enum. UTM (Unmanned Traffic Management) assigned UUID (RFC4122)
	MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID MAV_ODID_ID_TYPE = 3
)

func (e MAV_ODID_ID_TYPE) String() string {
	if name, ok := map[MAV_ODID_ID_TYPE]string{
		MAV_ODID_ID_TYPE_NONE:                "MAV_ODID_ID_TYPE_NONE",
		MAV_ODID_ID_TYPE_SERIAL_NUMBER:       "MAV_ODID_ID_TYPE_SERIAL_NUMBER",
		MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID: "MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID",
		MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID:   "MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_ID_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_ID_TYPE enums
func (e MAV_ODID_ID_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_ID_TYPE{
		MAV_ODID_ID_TYPE_NONE,
		MAV_ODID_ID_TYPE_SERIAL_NUMBER,
		MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID,
		MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID,
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

// MAV_ODID_UA_TYPE type
type MAV_ODID_UA_TYPE int

func (e MAV_ODID_UA_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_UA_TYPE_NONE enum. No UA (Unmanned Aircraft) type defined
	MAV_ODID_UA_TYPE_NONE MAV_ODID_UA_TYPE = 0
	// MAV_ODID_UA_TYPE_AEROPLANE enum. Aeroplane/Airplane. Fixed wing
	MAV_ODID_UA_TYPE_AEROPLANE MAV_ODID_UA_TYPE = 1
	// MAV_ODID_UA_TYPE_HELICOPTER_OR_MULTIROTOR enum. Helicopter or multirotor
	MAV_ODID_UA_TYPE_HELICOPTER_OR_MULTIROTOR MAV_ODID_UA_TYPE = 2
	// MAV_ODID_UA_TYPE_GYROPLANE enum. Gyroplane
	MAV_ODID_UA_TYPE_GYROPLANE MAV_ODID_UA_TYPE = 3
	// MAV_ODID_UA_TYPE_HYBRID_LIFT enum. VTOL (Vertical Take-Off and Landing). Fixed wing aircraft that can take off vertically
	MAV_ODID_UA_TYPE_HYBRID_LIFT MAV_ODID_UA_TYPE = 4
	// MAV_ODID_UA_TYPE_ORNITHOPTER enum. Ornithopter
	MAV_ODID_UA_TYPE_ORNITHOPTER MAV_ODID_UA_TYPE = 5
	// MAV_ODID_UA_TYPE_GLIDER enum. Glider
	MAV_ODID_UA_TYPE_GLIDER MAV_ODID_UA_TYPE = 6
	// MAV_ODID_UA_TYPE_KITE enum. Kite
	MAV_ODID_UA_TYPE_KITE MAV_ODID_UA_TYPE = 7
	// MAV_ODID_UA_TYPE_FREE_BALLOON enum. Free Balloon
	MAV_ODID_UA_TYPE_FREE_BALLOON MAV_ODID_UA_TYPE = 8
	// MAV_ODID_UA_TYPE_CAPTIVE_BALLOON enum. Captive Balloon
	MAV_ODID_UA_TYPE_CAPTIVE_BALLOON MAV_ODID_UA_TYPE = 9
	// MAV_ODID_UA_TYPE_AIRSHIP enum. Airship. E.g. a blimp
	MAV_ODID_UA_TYPE_AIRSHIP MAV_ODID_UA_TYPE = 10
	// MAV_ODID_UA_TYPE_FREE_FALL_PARACHUTE enum. Free Fall/Parachute (unpowered)
	MAV_ODID_UA_TYPE_FREE_FALL_PARACHUTE MAV_ODID_UA_TYPE = 11
	// MAV_ODID_UA_TYPE_ROCKET enum. Rocket
	MAV_ODID_UA_TYPE_ROCKET MAV_ODID_UA_TYPE = 12
	// MAV_ODID_UA_TYPE_TETHERED_POWERED_AIRCRAFT enum. Tethered powered aircraft
	MAV_ODID_UA_TYPE_TETHERED_POWERED_AIRCRAFT MAV_ODID_UA_TYPE = 13
	// MAV_ODID_UA_TYPE_GROUND_OBSTACLE enum. Ground Obstacle
	MAV_ODID_UA_TYPE_GROUND_OBSTACLE MAV_ODID_UA_TYPE = 14
	// MAV_ODID_UA_TYPE_OTHER enum. Other type of aircraft not listed earlier
	MAV_ODID_UA_TYPE_OTHER MAV_ODID_UA_TYPE = 15
)

func (e MAV_ODID_UA_TYPE) String() string {
	if name, ok := map[MAV_ODID_UA_TYPE]string{
		MAV_ODID_UA_TYPE_NONE:                      "MAV_ODID_UA_TYPE_NONE",
		MAV_ODID_UA_TYPE_AEROPLANE:                 "MAV_ODID_UA_TYPE_AEROPLANE",
		MAV_ODID_UA_TYPE_HELICOPTER_OR_MULTIROTOR:  "MAV_ODID_UA_TYPE_HELICOPTER_OR_MULTIROTOR",
		MAV_ODID_UA_TYPE_GYROPLANE:                 "MAV_ODID_UA_TYPE_GYROPLANE",
		MAV_ODID_UA_TYPE_HYBRID_LIFT:               "MAV_ODID_UA_TYPE_HYBRID_LIFT",
		MAV_ODID_UA_TYPE_ORNITHOPTER:               "MAV_ODID_UA_TYPE_ORNITHOPTER",
		MAV_ODID_UA_TYPE_GLIDER:                    "MAV_ODID_UA_TYPE_GLIDER",
		MAV_ODID_UA_TYPE_KITE:                      "MAV_ODID_UA_TYPE_KITE",
		MAV_ODID_UA_TYPE_FREE_BALLOON:              "MAV_ODID_UA_TYPE_FREE_BALLOON",
		MAV_ODID_UA_TYPE_CAPTIVE_BALLOON:           "MAV_ODID_UA_TYPE_CAPTIVE_BALLOON",
		MAV_ODID_UA_TYPE_AIRSHIP:                   "MAV_ODID_UA_TYPE_AIRSHIP",
		MAV_ODID_UA_TYPE_FREE_FALL_PARACHUTE:       "MAV_ODID_UA_TYPE_FREE_FALL_PARACHUTE",
		MAV_ODID_UA_TYPE_ROCKET:                    "MAV_ODID_UA_TYPE_ROCKET",
		MAV_ODID_UA_TYPE_TETHERED_POWERED_AIRCRAFT: "MAV_ODID_UA_TYPE_TETHERED_POWERED_AIRCRAFT",
		MAV_ODID_UA_TYPE_GROUND_OBSTACLE:           "MAV_ODID_UA_TYPE_GROUND_OBSTACLE",
		MAV_ODID_UA_TYPE_OTHER:                     "MAV_ODID_UA_TYPE_OTHER",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_UA_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_UA_TYPE enums
func (e MAV_ODID_UA_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_UA_TYPE{
		MAV_ODID_UA_TYPE_NONE,
		MAV_ODID_UA_TYPE_AEROPLANE,
		MAV_ODID_UA_TYPE_HELICOPTER_OR_MULTIROTOR,
		MAV_ODID_UA_TYPE_GYROPLANE,
		MAV_ODID_UA_TYPE_HYBRID_LIFT,
		MAV_ODID_UA_TYPE_ORNITHOPTER,
		MAV_ODID_UA_TYPE_GLIDER,
		MAV_ODID_UA_TYPE_KITE,
		MAV_ODID_UA_TYPE_FREE_BALLOON,
		MAV_ODID_UA_TYPE_CAPTIVE_BALLOON,
		MAV_ODID_UA_TYPE_AIRSHIP,
		MAV_ODID_UA_TYPE_FREE_FALL_PARACHUTE,
		MAV_ODID_UA_TYPE_ROCKET,
		MAV_ODID_UA_TYPE_TETHERED_POWERED_AIRCRAFT,
		MAV_ODID_UA_TYPE_GROUND_OBSTACLE,
		MAV_ODID_UA_TYPE_OTHER,
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

// MAV_ODID_STATUS type
type MAV_ODID_STATUS int

func (e MAV_ODID_STATUS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_STATUS_UNDECLARED enum. The status of the (UA) Unmanned Aircraft is undefined
	MAV_ODID_STATUS_UNDECLARED MAV_ODID_STATUS = 0
	// MAV_ODID_STATUS_GROUND enum. The UA is on the ground
	MAV_ODID_STATUS_GROUND MAV_ODID_STATUS = 1
	// MAV_ODID_STATUS_AIRBORNE enum. The UA is in the air
	MAV_ODID_STATUS_AIRBORNE MAV_ODID_STATUS = 2
	// MAV_ODID_STATUS_EMERGENCY enum. The UA is having an emergency
	MAV_ODID_STATUS_EMERGENCY MAV_ODID_STATUS = 3
)

func (e MAV_ODID_STATUS) String() string {
	if name, ok := map[MAV_ODID_STATUS]string{
		MAV_ODID_STATUS_UNDECLARED: "MAV_ODID_STATUS_UNDECLARED",
		MAV_ODID_STATUS_GROUND:     "MAV_ODID_STATUS_GROUND",
		MAV_ODID_STATUS_AIRBORNE:   "MAV_ODID_STATUS_AIRBORNE",
		MAV_ODID_STATUS_EMERGENCY:  "MAV_ODID_STATUS_EMERGENCY",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_STATUS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_STATUS enums
func (e MAV_ODID_STATUS) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_STATUS{
		MAV_ODID_STATUS_UNDECLARED,
		MAV_ODID_STATUS_GROUND,
		MAV_ODID_STATUS_AIRBORNE,
		MAV_ODID_STATUS_EMERGENCY,
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

// MAV_ODID_HEIGHT_REF type
type MAV_ODID_HEIGHT_REF int

func (e MAV_ODID_HEIGHT_REF) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_HEIGHT_REF_OVER_TAKEOFF enum. The height field is relative to the take-off location
	MAV_ODID_HEIGHT_REF_OVER_TAKEOFF MAV_ODID_HEIGHT_REF = 0
	// MAV_ODID_HEIGHT_REF_OVER_GROUND enum. The height field is relative to ground
	MAV_ODID_HEIGHT_REF_OVER_GROUND MAV_ODID_HEIGHT_REF = 1
)

func (e MAV_ODID_HEIGHT_REF) String() string {
	if name, ok := map[MAV_ODID_HEIGHT_REF]string{
		MAV_ODID_HEIGHT_REF_OVER_TAKEOFF: "MAV_ODID_HEIGHT_REF_OVER_TAKEOFF",
		MAV_ODID_HEIGHT_REF_OVER_GROUND:  "MAV_ODID_HEIGHT_REF_OVER_GROUND",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_HEIGHT_REF_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_HEIGHT_REF enums
func (e MAV_ODID_HEIGHT_REF) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_HEIGHT_REF{
		MAV_ODID_HEIGHT_REF_OVER_TAKEOFF,
		MAV_ODID_HEIGHT_REF_OVER_GROUND,
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

// MAV_ODID_HOR_ACC type
type MAV_ODID_HOR_ACC int

func (e MAV_ODID_HOR_ACC) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_HOR_ACC_UNKNOWN enum. The horizontal accuracy is unknown
	MAV_ODID_HOR_ACC_UNKNOWN MAV_ODID_HOR_ACC = 0
	// MAV_ODID_HOR_ACC_10NM enum. The horizontal accuracy is smaller than 10 Nautical Miles. 18.52 km
	MAV_ODID_HOR_ACC_10NM MAV_ODID_HOR_ACC = 1
	// MAV_ODID_HOR_ACC_4NM enum. The horizontal accuracy is smaller than 4 Nautical Miles. 7.408 km
	MAV_ODID_HOR_ACC_4NM MAV_ODID_HOR_ACC = 2
	// MAV_ODID_HOR_ACC_2NM enum. The horizontal accuracy is smaller than 2 Nautical Miles. 3.704 km
	MAV_ODID_HOR_ACC_2NM MAV_ODID_HOR_ACC = 3
	// MAV_ODID_HOR_ACC_1NM enum. The horizontal accuracy is smaller than 1 Nautical Miles. 1.852 km
	MAV_ODID_HOR_ACC_1NM MAV_ODID_HOR_ACC = 4
	// MAV_ODID_HOR_ACC_0_5NM enum. The horizontal accuracy is smaller than 0.5 Nautical Miles. 926 m
	MAV_ODID_HOR_ACC_0_5NM MAV_ODID_HOR_ACC = 5
	// MAV_ODID_HOR_ACC_0_3NM enum. The horizontal accuracy is smaller than 0.3 Nautical Miles. 555.6 m
	MAV_ODID_HOR_ACC_0_3NM MAV_ODID_HOR_ACC = 6
	// MAV_ODID_HOR_ACC_0_1NM enum. The horizontal accuracy is smaller than 0.1 Nautical Miles. 185.2 m
	MAV_ODID_HOR_ACC_0_1NM MAV_ODID_HOR_ACC = 7
	// MAV_ODID_HOR_ACC_0_05NM enum. The horizontal accuracy is smaller than 0.05 Nautical Miles. 92.6 m
	MAV_ODID_HOR_ACC_0_05NM MAV_ODID_HOR_ACC = 8
	// MAV_ODID_HOR_ACC_30_METER enum. The horizontal accuracy is smaller than 30 meter
	MAV_ODID_HOR_ACC_30_METER MAV_ODID_HOR_ACC = 9
	// MAV_ODID_HOR_ACC_10_METER enum. The horizontal accuracy is smaller than 10 meter
	MAV_ODID_HOR_ACC_10_METER MAV_ODID_HOR_ACC = 10
	// MAV_ODID_HOR_ACC_3_METER enum. The horizontal accuracy is smaller than 3 meter
	MAV_ODID_HOR_ACC_3_METER MAV_ODID_HOR_ACC = 11
	// MAV_ODID_HOR_ACC_1_METER enum. The horizontal accuracy is smaller than 1 meter
	MAV_ODID_HOR_ACC_1_METER MAV_ODID_HOR_ACC = 12
)

func (e MAV_ODID_HOR_ACC) String() string {
	if name, ok := map[MAV_ODID_HOR_ACC]string{
		MAV_ODID_HOR_ACC_UNKNOWN:  "MAV_ODID_HOR_ACC_UNKNOWN",
		MAV_ODID_HOR_ACC_10NM:     "MAV_ODID_HOR_ACC_10NM",
		MAV_ODID_HOR_ACC_4NM:      "MAV_ODID_HOR_ACC_4NM",
		MAV_ODID_HOR_ACC_2NM:      "MAV_ODID_HOR_ACC_2NM",
		MAV_ODID_HOR_ACC_1NM:      "MAV_ODID_HOR_ACC_1NM",
		MAV_ODID_HOR_ACC_0_5NM:    "MAV_ODID_HOR_ACC_0_5NM",
		MAV_ODID_HOR_ACC_0_3NM:    "MAV_ODID_HOR_ACC_0_3NM",
		MAV_ODID_HOR_ACC_0_1NM:    "MAV_ODID_HOR_ACC_0_1NM",
		MAV_ODID_HOR_ACC_0_05NM:   "MAV_ODID_HOR_ACC_0_05NM",
		MAV_ODID_HOR_ACC_30_METER: "MAV_ODID_HOR_ACC_30_METER",
		MAV_ODID_HOR_ACC_10_METER: "MAV_ODID_HOR_ACC_10_METER",
		MAV_ODID_HOR_ACC_3_METER:  "MAV_ODID_HOR_ACC_3_METER",
		MAV_ODID_HOR_ACC_1_METER:  "MAV_ODID_HOR_ACC_1_METER",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_HOR_ACC_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_HOR_ACC enums
func (e MAV_ODID_HOR_ACC) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_HOR_ACC{
		MAV_ODID_HOR_ACC_UNKNOWN,
		MAV_ODID_HOR_ACC_10NM,
		MAV_ODID_HOR_ACC_4NM,
		MAV_ODID_HOR_ACC_2NM,
		MAV_ODID_HOR_ACC_1NM,
		MAV_ODID_HOR_ACC_0_5NM,
		MAV_ODID_HOR_ACC_0_3NM,
		MAV_ODID_HOR_ACC_0_1NM,
		MAV_ODID_HOR_ACC_0_05NM,
		MAV_ODID_HOR_ACC_30_METER,
		MAV_ODID_HOR_ACC_10_METER,
		MAV_ODID_HOR_ACC_3_METER,
		MAV_ODID_HOR_ACC_1_METER,
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

// MAV_ODID_VER_ACC type
type MAV_ODID_VER_ACC int

func (e MAV_ODID_VER_ACC) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_VER_ACC_UNKNOWN enum. The vertical accuracy is unknown
	MAV_ODID_VER_ACC_UNKNOWN MAV_ODID_VER_ACC = 0
	// MAV_ODID_VER_ACC_150_METER enum. The vertical accuracy is smaller than 150 meter
	MAV_ODID_VER_ACC_150_METER MAV_ODID_VER_ACC = 1
	// MAV_ODID_VER_ACC_45_METER enum. The vertical accuracy is smaller than 45 meter
	MAV_ODID_VER_ACC_45_METER MAV_ODID_VER_ACC = 2
	// MAV_ODID_VER_ACC_25_METER enum. The vertical accuracy is smaller than 25 meter
	MAV_ODID_VER_ACC_25_METER MAV_ODID_VER_ACC = 3
	// MAV_ODID_VER_ACC_10_METER enum. The vertical accuracy is smaller than 10 meter
	MAV_ODID_VER_ACC_10_METER MAV_ODID_VER_ACC = 4
	// MAV_ODID_VER_ACC_3_METER enum. The vertical accuracy is smaller than 3 meter
	MAV_ODID_VER_ACC_3_METER MAV_ODID_VER_ACC = 5
	// MAV_ODID_VER_ACC_1_METER enum. The vertical accuracy is smaller than 1 meter
	MAV_ODID_VER_ACC_1_METER MAV_ODID_VER_ACC = 6
)

func (e MAV_ODID_VER_ACC) String() string {
	if name, ok := map[MAV_ODID_VER_ACC]string{
		MAV_ODID_VER_ACC_UNKNOWN:   "MAV_ODID_VER_ACC_UNKNOWN",
		MAV_ODID_VER_ACC_150_METER: "MAV_ODID_VER_ACC_150_METER",
		MAV_ODID_VER_ACC_45_METER:  "MAV_ODID_VER_ACC_45_METER",
		MAV_ODID_VER_ACC_25_METER:  "MAV_ODID_VER_ACC_25_METER",
		MAV_ODID_VER_ACC_10_METER:  "MAV_ODID_VER_ACC_10_METER",
		MAV_ODID_VER_ACC_3_METER:   "MAV_ODID_VER_ACC_3_METER",
		MAV_ODID_VER_ACC_1_METER:   "MAV_ODID_VER_ACC_1_METER",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_VER_ACC_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_VER_ACC enums
func (e MAV_ODID_VER_ACC) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_VER_ACC{
		MAV_ODID_VER_ACC_UNKNOWN,
		MAV_ODID_VER_ACC_150_METER,
		MAV_ODID_VER_ACC_45_METER,
		MAV_ODID_VER_ACC_25_METER,
		MAV_ODID_VER_ACC_10_METER,
		MAV_ODID_VER_ACC_3_METER,
		MAV_ODID_VER_ACC_1_METER,
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

// MAV_ODID_SPEED_ACC type
type MAV_ODID_SPEED_ACC int

func (e MAV_ODID_SPEED_ACC) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_SPEED_ACC_UNKNOWN enum. The speed accuracy is unknown
	MAV_ODID_SPEED_ACC_UNKNOWN MAV_ODID_SPEED_ACC = 0
	// MAV_ODID_SPEED_ACC_10_METERS_PER_SECOND enum. The speed accuracy is smaller than 10 meters per second
	MAV_ODID_SPEED_ACC_10_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 1
	// MAV_ODID_SPEED_ACC_3_METERS_PER_SECOND enum. The speed accuracy is smaller than 3 meters per second
	MAV_ODID_SPEED_ACC_3_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 2
	// MAV_ODID_SPEED_ACC_1_METERS_PER_SECOND enum. The speed accuracy is smaller than 1 meters per second
	MAV_ODID_SPEED_ACC_1_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 3
	// MAV_ODID_SPEED_ACC_0_3_METERS_PER_SECOND enum. The speed accuracy is smaller than 0.3 meters per second
	MAV_ODID_SPEED_ACC_0_3_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 4
)

func (e MAV_ODID_SPEED_ACC) String() string {
	if name, ok := map[MAV_ODID_SPEED_ACC]string{
		MAV_ODID_SPEED_ACC_UNKNOWN:               "MAV_ODID_SPEED_ACC_UNKNOWN",
		MAV_ODID_SPEED_ACC_10_METERS_PER_SECOND:  "MAV_ODID_SPEED_ACC_10_METERS_PER_SECOND",
		MAV_ODID_SPEED_ACC_3_METERS_PER_SECOND:   "MAV_ODID_SPEED_ACC_3_METERS_PER_SECOND",
		MAV_ODID_SPEED_ACC_1_METERS_PER_SECOND:   "MAV_ODID_SPEED_ACC_1_METERS_PER_SECOND",
		MAV_ODID_SPEED_ACC_0_3_METERS_PER_SECOND: "MAV_ODID_SPEED_ACC_0_3_METERS_PER_SECOND",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_SPEED_ACC_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_SPEED_ACC enums
func (e MAV_ODID_SPEED_ACC) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_SPEED_ACC{
		MAV_ODID_SPEED_ACC_UNKNOWN,
		MAV_ODID_SPEED_ACC_10_METERS_PER_SECOND,
		MAV_ODID_SPEED_ACC_3_METERS_PER_SECOND,
		MAV_ODID_SPEED_ACC_1_METERS_PER_SECOND,
		MAV_ODID_SPEED_ACC_0_3_METERS_PER_SECOND,
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

// MAV_ODID_TIME_ACC type
type MAV_ODID_TIME_ACC int

func (e MAV_ODID_TIME_ACC) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_TIME_ACC_UNKNOWN enum. The timestamp accuracy is unknown
	MAV_ODID_TIME_ACC_UNKNOWN MAV_ODID_TIME_ACC = 0
	// MAV_ODID_TIME_ACC_0_1_SECOND enum. The timestamp accuracy is smaller than or equal to 0.1 second
	MAV_ODID_TIME_ACC_0_1_SECOND MAV_ODID_TIME_ACC = 1
	// MAV_ODID_TIME_ACC_0_2_SECOND enum. The timestamp accuracy is smaller than or equal to 0.2 second
	MAV_ODID_TIME_ACC_0_2_SECOND MAV_ODID_TIME_ACC = 2
	// MAV_ODID_TIME_ACC_0_3_SECOND enum. The timestamp accuracy is smaller than or equal to 0.3 second
	MAV_ODID_TIME_ACC_0_3_SECOND MAV_ODID_TIME_ACC = 3
	// MAV_ODID_TIME_ACC_0_4_SECOND enum. The timestamp accuracy is smaller than or equal to 0.4 second
	MAV_ODID_TIME_ACC_0_4_SECOND MAV_ODID_TIME_ACC = 4
	// MAV_ODID_TIME_ACC_0_5_SECOND enum. The timestamp accuracy is smaller than or equal to 0.5 second
	MAV_ODID_TIME_ACC_0_5_SECOND MAV_ODID_TIME_ACC = 5
	// MAV_ODID_TIME_ACC_0_6_SECOND enum. The timestamp accuracy is smaller than or equal to 0.6 second
	MAV_ODID_TIME_ACC_0_6_SECOND MAV_ODID_TIME_ACC = 6
	// MAV_ODID_TIME_ACC_0_7_SECOND enum. The timestamp accuracy is smaller than or equal to 0.7 second
	MAV_ODID_TIME_ACC_0_7_SECOND MAV_ODID_TIME_ACC = 7
	// MAV_ODID_TIME_ACC_0_8_SECOND enum. The timestamp accuracy is smaller than or equal to 0.8 second
	MAV_ODID_TIME_ACC_0_8_SECOND MAV_ODID_TIME_ACC = 8
	// MAV_ODID_TIME_ACC_0_9_SECOND enum. The timestamp accuracy is smaller than or equal to 0.9 second
	MAV_ODID_TIME_ACC_0_9_SECOND MAV_ODID_TIME_ACC = 9
	// MAV_ODID_TIME_ACC_1_0_SECOND enum. The timestamp accuracy is smaller than or equal to 1.0 second
	MAV_ODID_TIME_ACC_1_0_SECOND MAV_ODID_TIME_ACC = 10
	// MAV_ODID_TIME_ACC_1_1_SECOND enum. The timestamp accuracy is smaller than or equal to 1.1 second
	MAV_ODID_TIME_ACC_1_1_SECOND MAV_ODID_TIME_ACC = 11
	// MAV_ODID_TIME_ACC_1_2_SECOND enum. The timestamp accuracy is smaller than or equal to 1.2 second
	MAV_ODID_TIME_ACC_1_2_SECOND MAV_ODID_TIME_ACC = 12
	// MAV_ODID_TIME_ACC_1_3_SECOND enum. The timestamp accuracy is smaller than or equal to 1.3 second
	MAV_ODID_TIME_ACC_1_3_SECOND MAV_ODID_TIME_ACC = 13
	// MAV_ODID_TIME_ACC_1_4_SECOND enum. The timestamp accuracy is smaller than or equal to 1.4 second
	MAV_ODID_TIME_ACC_1_4_SECOND MAV_ODID_TIME_ACC = 14
	// MAV_ODID_TIME_ACC_1_5_SECOND enum. The timestamp accuracy is smaller than or equal to 1.5 second
	MAV_ODID_TIME_ACC_1_5_SECOND MAV_ODID_TIME_ACC = 15
)

func (e MAV_ODID_TIME_ACC) String() string {
	if name, ok := map[MAV_ODID_TIME_ACC]string{
		MAV_ODID_TIME_ACC_UNKNOWN:    "MAV_ODID_TIME_ACC_UNKNOWN",
		MAV_ODID_TIME_ACC_0_1_SECOND: "MAV_ODID_TIME_ACC_0_1_SECOND",
		MAV_ODID_TIME_ACC_0_2_SECOND: "MAV_ODID_TIME_ACC_0_2_SECOND",
		MAV_ODID_TIME_ACC_0_3_SECOND: "MAV_ODID_TIME_ACC_0_3_SECOND",
		MAV_ODID_TIME_ACC_0_4_SECOND: "MAV_ODID_TIME_ACC_0_4_SECOND",
		MAV_ODID_TIME_ACC_0_5_SECOND: "MAV_ODID_TIME_ACC_0_5_SECOND",
		MAV_ODID_TIME_ACC_0_6_SECOND: "MAV_ODID_TIME_ACC_0_6_SECOND",
		MAV_ODID_TIME_ACC_0_7_SECOND: "MAV_ODID_TIME_ACC_0_7_SECOND",
		MAV_ODID_TIME_ACC_0_8_SECOND: "MAV_ODID_TIME_ACC_0_8_SECOND",
		MAV_ODID_TIME_ACC_0_9_SECOND: "MAV_ODID_TIME_ACC_0_9_SECOND",
		MAV_ODID_TIME_ACC_1_0_SECOND: "MAV_ODID_TIME_ACC_1_0_SECOND",
		MAV_ODID_TIME_ACC_1_1_SECOND: "MAV_ODID_TIME_ACC_1_1_SECOND",
		MAV_ODID_TIME_ACC_1_2_SECOND: "MAV_ODID_TIME_ACC_1_2_SECOND",
		MAV_ODID_TIME_ACC_1_3_SECOND: "MAV_ODID_TIME_ACC_1_3_SECOND",
		MAV_ODID_TIME_ACC_1_4_SECOND: "MAV_ODID_TIME_ACC_1_4_SECOND",
		MAV_ODID_TIME_ACC_1_5_SECOND: "MAV_ODID_TIME_ACC_1_5_SECOND",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_TIME_ACC_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_TIME_ACC enums
func (e MAV_ODID_TIME_ACC) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_TIME_ACC{
		MAV_ODID_TIME_ACC_UNKNOWN,
		MAV_ODID_TIME_ACC_0_1_SECOND,
		MAV_ODID_TIME_ACC_0_2_SECOND,
		MAV_ODID_TIME_ACC_0_3_SECOND,
		MAV_ODID_TIME_ACC_0_4_SECOND,
		MAV_ODID_TIME_ACC_0_5_SECOND,
		MAV_ODID_TIME_ACC_0_6_SECOND,
		MAV_ODID_TIME_ACC_0_7_SECOND,
		MAV_ODID_TIME_ACC_0_8_SECOND,
		MAV_ODID_TIME_ACC_0_9_SECOND,
		MAV_ODID_TIME_ACC_1_0_SECOND,
		MAV_ODID_TIME_ACC_1_1_SECOND,
		MAV_ODID_TIME_ACC_1_2_SECOND,
		MAV_ODID_TIME_ACC_1_3_SECOND,
		MAV_ODID_TIME_ACC_1_4_SECOND,
		MAV_ODID_TIME_ACC_1_5_SECOND,
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

// MAV_ODID_AUTH_TYPE type
type MAV_ODID_AUTH_TYPE int

func (e MAV_ODID_AUTH_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_AUTH_TYPE_NONE enum. No authentication type is specified
	MAV_ODID_AUTH_TYPE_NONE MAV_ODID_AUTH_TYPE = 0
	// MAV_ODID_AUTH_TYPE_UAS_ID_SIGNATURE enum. Signature for the UAS (Unmanned Aircraft System) ID
	MAV_ODID_AUTH_TYPE_UAS_ID_SIGNATURE MAV_ODID_AUTH_TYPE = 1
	// MAV_ODID_AUTH_TYPE_OPERATOR_ID_SIGNATURE enum. Signature for the Operator ID
	MAV_ODID_AUTH_TYPE_OPERATOR_ID_SIGNATURE MAV_ODID_AUTH_TYPE = 2
	// MAV_ODID_AUTH_TYPE_MESSAGE_SET_SIGNATURE enum. Signature for the entire message set
	MAV_ODID_AUTH_TYPE_MESSAGE_SET_SIGNATURE MAV_ODID_AUTH_TYPE = 3
	// MAV_ODID_AUTH_TYPE_NETWORK_REMOTE_ID enum. Authentication is provided by Network Remote ID
	MAV_ODID_AUTH_TYPE_NETWORK_REMOTE_ID MAV_ODID_AUTH_TYPE = 4
)

func (e MAV_ODID_AUTH_TYPE) String() string {
	if name, ok := map[MAV_ODID_AUTH_TYPE]string{
		MAV_ODID_AUTH_TYPE_NONE:                  "MAV_ODID_AUTH_TYPE_NONE",
		MAV_ODID_AUTH_TYPE_UAS_ID_SIGNATURE:      "MAV_ODID_AUTH_TYPE_UAS_ID_SIGNATURE",
		MAV_ODID_AUTH_TYPE_OPERATOR_ID_SIGNATURE: "MAV_ODID_AUTH_TYPE_OPERATOR_ID_SIGNATURE",
		MAV_ODID_AUTH_TYPE_MESSAGE_SET_SIGNATURE: "MAV_ODID_AUTH_TYPE_MESSAGE_SET_SIGNATURE",
		MAV_ODID_AUTH_TYPE_NETWORK_REMOTE_ID:     "MAV_ODID_AUTH_TYPE_NETWORK_REMOTE_ID",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_AUTH_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_AUTH_TYPE enums
func (e MAV_ODID_AUTH_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_AUTH_TYPE{
		MAV_ODID_AUTH_TYPE_NONE,
		MAV_ODID_AUTH_TYPE_UAS_ID_SIGNATURE,
		MAV_ODID_AUTH_TYPE_OPERATOR_ID_SIGNATURE,
		MAV_ODID_AUTH_TYPE_MESSAGE_SET_SIGNATURE,
		MAV_ODID_AUTH_TYPE_NETWORK_REMOTE_ID,
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

// MAV_ODID_DESC_TYPE type
type MAV_ODID_DESC_TYPE int

func (e MAV_ODID_DESC_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_DESC_TYPE_TEXT enum. Free-form text description of the purpose of the flight
	MAV_ODID_DESC_TYPE_TEXT MAV_ODID_DESC_TYPE = 0
)

func (e MAV_ODID_DESC_TYPE) String() string {
	if name, ok := map[MAV_ODID_DESC_TYPE]string{
		MAV_ODID_DESC_TYPE_TEXT: "MAV_ODID_DESC_TYPE_TEXT",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_DESC_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_DESC_TYPE enums
func (e MAV_ODID_DESC_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_DESC_TYPE{
		MAV_ODID_DESC_TYPE_TEXT,
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

// MAV_ODID_OPERATOR_LOCATION_TYPE type
type MAV_ODID_OPERATOR_LOCATION_TYPE int

func (e MAV_ODID_OPERATOR_LOCATION_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_OPERATOR_LOCATION_TYPE_TAKEOFF enum. The location of the operator is the same as the take-off location
	MAV_ODID_OPERATOR_LOCATION_TYPE_TAKEOFF MAV_ODID_OPERATOR_LOCATION_TYPE = 0
	// MAV_ODID_OPERATOR_LOCATION_TYPE_LIVE_GNSS enum. The location of the operator is based on live GNSS data
	MAV_ODID_OPERATOR_LOCATION_TYPE_LIVE_GNSS MAV_ODID_OPERATOR_LOCATION_TYPE = 1
	// MAV_ODID_OPERATOR_LOCATION_TYPE_FIXED enum. The location of the operator is a fixed location
	MAV_ODID_OPERATOR_LOCATION_TYPE_FIXED MAV_ODID_OPERATOR_LOCATION_TYPE = 2
)

func (e MAV_ODID_OPERATOR_LOCATION_TYPE) String() string {
	if name, ok := map[MAV_ODID_OPERATOR_LOCATION_TYPE]string{
		MAV_ODID_OPERATOR_LOCATION_TYPE_TAKEOFF:   "MAV_ODID_OPERATOR_LOCATION_TYPE_TAKEOFF",
		MAV_ODID_OPERATOR_LOCATION_TYPE_LIVE_GNSS: "MAV_ODID_OPERATOR_LOCATION_TYPE_LIVE_GNSS",
		MAV_ODID_OPERATOR_LOCATION_TYPE_FIXED:     "MAV_ODID_OPERATOR_LOCATION_TYPE_FIXED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_OPERATOR_LOCATION_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_OPERATOR_LOCATION_TYPE enums
func (e MAV_ODID_OPERATOR_LOCATION_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_OPERATOR_LOCATION_TYPE{
		MAV_ODID_OPERATOR_LOCATION_TYPE_TAKEOFF,
		MAV_ODID_OPERATOR_LOCATION_TYPE_LIVE_GNSS,
		MAV_ODID_OPERATOR_LOCATION_TYPE_FIXED,
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

// MAV_ODID_CLASSIFICATION_TYPE type
type MAV_ODID_CLASSIFICATION_TYPE int

func (e MAV_ODID_CLASSIFICATION_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_CLASSIFICATION_TYPE_UNDECLARED enum. The classification type for the UA is undeclared
	MAV_ODID_CLASSIFICATION_TYPE_UNDECLARED MAV_ODID_CLASSIFICATION_TYPE = 0
	// MAV_ODID_CLASSIFICATION_TYPE_EU enum. The classification type for the UA follows EU (European Union) specifications
	MAV_ODID_CLASSIFICATION_TYPE_EU MAV_ODID_CLASSIFICATION_TYPE = 1
)

func (e MAV_ODID_CLASSIFICATION_TYPE) String() string {
	if name, ok := map[MAV_ODID_CLASSIFICATION_TYPE]string{
		MAV_ODID_CLASSIFICATION_TYPE_UNDECLARED: "MAV_ODID_CLASSIFICATION_TYPE_UNDECLARED",
		MAV_ODID_CLASSIFICATION_TYPE_EU:         "MAV_ODID_CLASSIFICATION_TYPE_EU",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_CLASSIFICATION_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_CLASSIFICATION_TYPE enums
func (e MAV_ODID_CLASSIFICATION_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_CLASSIFICATION_TYPE{
		MAV_ODID_CLASSIFICATION_TYPE_UNDECLARED,
		MAV_ODID_CLASSIFICATION_TYPE_EU,
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

// MAV_ODID_CATEGORY_EU type
type MAV_ODID_CATEGORY_EU int

func (e MAV_ODID_CATEGORY_EU) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_CATEGORY_EU_UNDECLARED enum. The category for the UA, according to the EU specification, is undeclared
	MAV_ODID_CATEGORY_EU_UNDECLARED MAV_ODID_CATEGORY_EU = 0
	// MAV_ODID_CATEGORY_EU_OPEN enum. The category for the UA, according to the EU specification, is the Open category
	MAV_ODID_CATEGORY_EU_OPEN MAV_ODID_CATEGORY_EU = 1
	// MAV_ODID_CATEGORY_EU_SPECIFIC enum. The category for the UA, according to the EU specification, is the Specific category
	MAV_ODID_CATEGORY_EU_SPECIFIC MAV_ODID_CATEGORY_EU = 2
	// MAV_ODID_CATEGORY_EU_CERTIFIED enum. The category for the UA, according to the EU specification, is the Certified category
	MAV_ODID_CATEGORY_EU_CERTIFIED MAV_ODID_CATEGORY_EU = 3
)

func (e MAV_ODID_CATEGORY_EU) String() string {
	if name, ok := map[MAV_ODID_CATEGORY_EU]string{
		MAV_ODID_CATEGORY_EU_UNDECLARED: "MAV_ODID_CATEGORY_EU_UNDECLARED",
		MAV_ODID_CATEGORY_EU_OPEN:       "MAV_ODID_CATEGORY_EU_OPEN",
		MAV_ODID_CATEGORY_EU_SPECIFIC:   "MAV_ODID_CATEGORY_EU_SPECIFIC",
		MAV_ODID_CATEGORY_EU_CERTIFIED:  "MAV_ODID_CATEGORY_EU_CERTIFIED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_CATEGORY_EU_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_CATEGORY_EU enums
func (e MAV_ODID_CATEGORY_EU) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_CATEGORY_EU{
		MAV_ODID_CATEGORY_EU_UNDECLARED,
		MAV_ODID_CATEGORY_EU_OPEN,
		MAV_ODID_CATEGORY_EU_SPECIFIC,
		MAV_ODID_CATEGORY_EU_CERTIFIED,
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

// MAV_ODID_CLASS_EU type
type MAV_ODID_CLASS_EU int

func (e MAV_ODID_CLASS_EU) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_CLASS_EU_UNDECLARED enum. The class for the UA, according to the EU specification, is undeclared
	MAV_ODID_CLASS_EU_UNDECLARED MAV_ODID_CLASS_EU = 0
	// MAV_ODID_CLASS_EU_CLASS_0 enum. The class for the UA, according to the EU specification, is Class 0
	MAV_ODID_CLASS_EU_CLASS_0 MAV_ODID_CLASS_EU = 1
	// MAV_ODID_CLASS_EU_CLASS_1 enum. The class for the UA, according to the EU specification, is Class 1
	MAV_ODID_CLASS_EU_CLASS_1 MAV_ODID_CLASS_EU = 2
	// MAV_ODID_CLASS_EU_CLASS_2 enum. The class for the UA, according to the EU specification, is Class 2
	MAV_ODID_CLASS_EU_CLASS_2 MAV_ODID_CLASS_EU = 3
	// MAV_ODID_CLASS_EU_CLASS_3 enum. The class for the UA, according to the EU specification, is Class 3
	MAV_ODID_CLASS_EU_CLASS_3 MAV_ODID_CLASS_EU = 4
	// MAV_ODID_CLASS_EU_CLASS_4 enum. The class for the UA, according to the EU specification, is Class 4
	MAV_ODID_CLASS_EU_CLASS_4 MAV_ODID_CLASS_EU = 5
	// MAV_ODID_CLASS_EU_CLASS_5 enum. The class for the UA, according to the EU specification, is Class 5
	MAV_ODID_CLASS_EU_CLASS_5 MAV_ODID_CLASS_EU = 6
	// MAV_ODID_CLASS_EU_CLASS_6 enum. The class for the UA, according to the EU specification, is Class 6
	MAV_ODID_CLASS_EU_CLASS_6 MAV_ODID_CLASS_EU = 7
)

func (e MAV_ODID_CLASS_EU) String() string {
	if name, ok := map[MAV_ODID_CLASS_EU]string{
		MAV_ODID_CLASS_EU_UNDECLARED: "MAV_ODID_CLASS_EU_UNDECLARED",
		MAV_ODID_CLASS_EU_CLASS_0:    "MAV_ODID_CLASS_EU_CLASS_0",
		MAV_ODID_CLASS_EU_CLASS_1:    "MAV_ODID_CLASS_EU_CLASS_1",
		MAV_ODID_CLASS_EU_CLASS_2:    "MAV_ODID_CLASS_EU_CLASS_2",
		MAV_ODID_CLASS_EU_CLASS_3:    "MAV_ODID_CLASS_EU_CLASS_3",
		MAV_ODID_CLASS_EU_CLASS_4:    "MAV_ODID_CLASS_EU_CLASS_4",
		MAV_ODID_CLASS_EU_CLASS_5:    "MAV_ODID_CLASS_EU_CLASS_5",
		MAV_ODID_CLASS_EU_CLASS_6:    "MAV_ODID_CLASS_EU_CLASS_6",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_CLASS_EU_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_CLASS_EU enums
func (e MAV_ODID_CLASS_EU) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_CLASS_EU{
		MAV_ODID_CLASS_EU_UNDECLARED,
		MAV_ODID_CLASS_EU_CLASS_0,
		MAV_ODID_CLASS_EU_CLASS_1,
		MAV_ODID_CLASS_EU_CLASS_2,
		MAV_ODID_CLASS_EU_CLASS_3,
		MAV_ODID_CLASS_EU_CLASS_4,
		MAV_ODID_CLASS_EU_CLASS_5,
		MAV_ODID_CLASS_EU_CLASS_6,
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

// MAV_ODID_OPERATOR_ID_TYPE type
type MAV_ODID_OPERATOR_ID_TYPE int

func (e MAV_ODID_OPERATOR_ID_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_ODID_OPERATOR_ID_TYPE_CAA enum. CAA (Civil Aviation Authority) registered operator ID
	MAV_ODID_OPERATOR_ID_TYPE_CAA MAV_ODID_OPERATOR_ID_TYPE = 0
)

func (e MAV_ODID_OPERATOR_ID_TYPE) String() string {
	if name, ok := map[MAV_ODID_OPERATOR_ID_TYPE]string{
		MAV_ODID_OPERATOR_ID_TYPE_CAA: "MAV_ODID_OPERATOR_ID_TYPE_CAA",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_ODID_OPERATOR_ID_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_ODID_OPERATOR_ID_TYPE enums
func (e MAV_ODID_OPERATOR_ID_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_ODID_OPERATOR_ID_TYPE{
		MAV_ODID_OPERATOR_ID_TYPE_CAA,
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

// TUNE_FORMAT type. Tune formats (used for vehicle buzzer/tone generation).
type TUNE_FORMAT int

func (e TUNE_FORMAT) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// TUNE_FORMAT_QBASIC1_1 enum. Format is QBasic 1.1 Play: https://www.qbasic.net/en/reference/qb11/Statement/PLAY-006.htm
	TUNE_FORMAT_QBASIC1_1 TUNE_FORMAT = 1
	// TUNE_FORMAT_MML_MODERN enum. Format is Modern Music Markup Language (MML): https://en.wikipedia.org/wiki/Music_Macro_Language#Modern_MML
	TUNE_FORMAT_MML_MODERN TUNE_FORMAT = 2
)

func (e TUNE_FORMAT) String() string {
	if name, ok := map[TUNE_FORMAT]string{
		TUNE_FORMAT_QBASIC1_1:  "TUNE_FORMAT_QBASIC1_1",
		TUNE_FORMAT_MML_MODERN: "TUNE_FORMAT_MML_MODERN",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("TUNE_FORMAT_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects TUNE_FORMAT enums
func (e TUNE_FORMAT) Bitmask() string {
	bitmap := ""
	for _, entry := range []TUNE_FORMAT{
		TUNE_FORMAT_QBASIC1_1,
		TUNE_FORMAT_MML_MODERN,
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

// COMPONENT_CAP_FLAGS type. Component capability flags (Bitmap)
type COMPONENT_CAP_FLAGS int

func (e COMPONENT_CAP_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// COMPONENT_CAP_FLAGS_PARAM enum. Component has parameters, and supports the parameter protocol (PARAM messages)
	COMPONENT_CAP_FLAGS_PARAM COMPONENT_CAP_FLAGS = 1
	// COMPONENT_CAP_FLAGS_PARAM_EXT enum. Component has parameters, and supports the extended parameter protocol (PARAM_EXT messages)
	COMPONENT_CAP_FLAGS_PARAM_EXT COMPONENT_CAP_FLAGS = 2
)

func (e COMPONENT_CAP_FLAGS) String() string {
	if name, ok := map[COMPONENT_CAP_FLAGS]string{
		COMPONENT_CAP_FLAGS_PARAM:     "COMPONENT_CAP_FLAGS_PARAM",
		COMPONENT_CAP_FLAGS_PARAM_EXT: "COMPONENT_CAP_FLAGS_PARAM_EXT",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("COMPONENT_CAP_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects COMPONENT_CAP_FLAGS enums
func (e COMPONENT_CAP_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []COMPONENT_CAP_FLAGS{
		COMPONENT_CAP_FLAGS_PARAM,
		COMPONENT_CAP_FLAGS_PARAM_EXT,
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

// AIS_TYPE type. Type of AIS vessel, enum duplicated from AIS standard, https://gpsd.gitlab.io/gpsd/AIVDM.html
type AIS_TYPE int

func (e AIS_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// AIS_TYPE_UNKNOWN enum. Not available (default)
	AIS_TYPE_UNKNOWN AIS_TYPE = 0
	// AIS_TYPE_RESERVED_1 enum
	AIS_TYPE_RESERVED_1 AIS_TYPE = 1
	// AIS_TYPE_RESERVED_2 enum
	AIS_TYPE_RESERVED_2 AIS_TYPE = 2
	// AIS_TYPE_RESERVED_3 enum
	AIS_TYPE_RESERVED_3 AIS_TYPE = 3
	// AIS_TYPE_RESERVED_4 enum
	AIS_TYPE_RESERVED_4 AIS_TYPE = 4
	// AIS_TYPE_RESERVED_5 enum
	AIS_TYPE_RESERVED_5 AIS_TYPE = 5
	// AIS_TYPE_RESERVED_6 enum
	AIS_TYPE_RESERVED_6 AIS_TYPE = 6
	// AIS_TYPE_RESERVED_7 enum
	AIS_TYPE_RESERVED_7 AIS_TYPE = 7
	// AIS_TYPE_RESERVED_8 enum
	AIS_TYPE_RESERVED_8 AIS_TYPE = 8
	// AIS_TYPE_RESERVED_9 enum
	AIS_TYPE_RESERVED_9 AIS_TYPE = 9
	// AIS_TYPE_RESERVED_10 enum
	AIS_TYPE_RESERVED_10 AIS_TYPE = 10
	// AIS_TYPE_RESERVED_11 enum
	AIS_TYPE_RESERVED_11 AIS_TYPE = 11
	// AIS_TYPE_RESERVED_12 enum
	AIS_TYPE_RESERVED_12 AIS_TYPE = 12
	// AIS_TYPE_RESERVED_13 enum
	AIS_TYPE_RESERVED_13 AIS_TYPE = 13
	// AIS_TYPE_RESERVED_14 enum
	AIS_TYPE_RESERVED_14 AIS_TYPE = 14
	// AIS_TYPE_RESERVED_15 enum
	AIS_TYPE_RESERVED_15 AIS_TYPE = 15
	// AIS_TYPE_RESERVED_16 enum
	AIS_TYPE_RESERVED_16 AIS_TYPE = 16
	// AIS_TYPE_RESERVED_17 enum
	AIS_TYPE_RESERVED_17 AIS_TYPE = 17
	// AIS_TYPE_RESERVED_18 enum
	AIS_TYPE_RESERVED_18 AIS_TYPE = 18
	// AIS_TYPE_RESERVED_19 enum
	AIS_TYPE_RESERVED_19 AIS_TYPE = 19
	// AIS_TYPE_WIG enum. Wing In Ground effect
	AIS_TYPE_WIG AIS_TYPE = 20
	// AIS_TYPE_WIG_HAZARDOUS_A enum
	AIS_TYPE_WIG_HAZARDOUS_A AIS_TYPE = 21
	// AIS_TYPE_WIG_HAZARDOUS_B enum
	AIS_TYPE_WIG_HAZARDOUS_B AIS_TYPE = 22
	// AIS_TYPE_WIG_HAZARDOUS_C enum
	AIS_TYPE_WIG_HAZARDOUS_C AIS_TYPE = 23
	// AIS_TYPE_WIG_HAZARDOUS_D enum
	AIS_TYPE_WIG_HAZARDOUS_D AIS_TYPE = 24
	// AIS_TYPE_WIG_RESERVED_1 enum
	AIS_TYPE_WIG_RESERVED_1 AIS_TYPE = 25
	// AIS_TYPE_WIG_RESERVED_2 enum
	AIS_TYPE_WIG_RESERVED_2 AIS_TYPE = 26
	// AIS_TYPE_WIG_RESERVED_3 enum
	AIS_TYPE_WIG_RESERVED_3 AIS_TYPE = 27
	// AIS_TYPE_WIG_RESERVED_4 enum
	AIS_TYPE_WIG_RESERVED_4 AIS_TYPE = 28
	// AIS_TYPE_WIG_RESERVED_5 enum
	AIS_TYPE_WIG_RESERVED_5 AIS_TYPE = 29
	// AIS_TYPE_FISHING enum
	AIS_TYPE_FISHING AIS_TYPE = 30
	// AIS_TYPE_TOWING enum
	AIS_TYPE_TOWING AIS_TYPE = 31
	// AIS_TYPE_TOWING_LARGE enum. Towing: length exceeds 200m or breadth exceeds 25m
	AIS_TYPE_TOWING_LARGE AIS_TYPE = 32
	// AIS_TYPE_DREDGING enum. Dredging or other underwater ops
	AIS_TYPE_DREDGING AIS_TYPE = 33
	// AIS_TYPE_DIVING enum
	AIS_TYPE_DIVING AIS_TYPE = 34
	// AIS_TYPE_MILITARY enum
	AIS_TYPE_MILITARY AIS_TYPE = 35
	// AIS_TYPE_SAILING enum
	AIS_TYPE_SAILING AIS_TYPE = 36
	// AIS_TYPE_PLEASURE enum
	AIS_TYPE_PLEASURE AIS_TYPE = 37
	// AIS_TYPE_RESERVED_20 enum
	AIS_TYPE_RESERVED_20 AIS_TYPE = 38
	// AIS_TYPE_RESERVED_21 enum
	AIS_TYPE_RESERVED_21 AIS_TYPE = 39
	// AIS_TYPE_HSC enum. High Speed Craft
	AIS_TYPE_HSC AIS_TYPE = 40
	// AIS_TYPE_HSC_HAZARDOUS_A enum
	AIS_TYPE_HSC_HAZARDOUS_A AIS_TYPE = 41
	// AIS_TYPE_HSC_HAZARDOUS_B enum
	AIS_TYPE_HSC_HAZARDOUS_B AIS_TYPE = 42
	// AIS_TYPE_HSC_HAZARDOUS_C enum
	AIS_TYPE_HSC_HAZARDOUS_C AIS_TYPE = 43
	// AIS_TYPE_HSC_HAZARDOUS_D enum
	AIS_TYPE_HSC_HAZARDOUS_D AIS_TYPE = 44
	// AIS_TYPE_HSC_RESERVED_1 enum
	AIS_TYPE_HSC_RESERVED_1 AIS_TYPE = 45
	// AIS_TYPE_HSC_RESERVED_2 enum
	AIS_TYPE_HSC_RESERVED_2 AIS_TYPE = 46
	// AIS_TYPE_HSC_RESERVED_3 enum
	AIS_TYPE_HSC_RESERVED_3 AIS_TYPE = 47
	// AIS_TYPE_HSC_RESERVED_4 enum
	AIS_TYPE_HSC_RESERVED_4 AIS_TYPE = 48
	// AIS_TYPE_HSC_UNKNOWN enum
	AIS_TYPE_HSC_UNKNOWN AIS_TYPE = 49
	// AIS_TYPE_PILOT enum
	AIS_TYPE_PILOT AIS_TYPE = 50
	// AIS_TYPE_SAR enum. Search And Rescue vessel
	AIS_TYPE_SAR AIS_TYPE = 51
	// AIS_TYPE_TUG enum
	AIS_TYPE_TUG AIS_TYPE = 52
	// AIS_TYPE_PORT_TENDER enum
	AIS_TYPE_PORT_TENDER AIS_TYPE = 53
	// AIS_TYPE_ANTI_POLLUTION enum. Anti-pollution equipment
	AIS_TYPE_ANTI_POLLUTION AIS_TYPE = 54
	// AIS_TYPE_LAW_ENFORCEMENT enum
	AIS_TYPE_LAW_ENFORCEMENT AIS_TYPE = 55
	// AIS_TYPE_SPARE_LOCAL_1 enum
	AIS_TYPE_SPARE_LOCAL_1 AIS_TYPE = 56
	// AIS_TYPE_SPARE_LOCAL_2 enum
	AIS_TYPE_SPARE_LOCAL_2 AIS_TYPE = 57
	// AIS_TYPE_MEDICAL_TRANSPORT enum
	AIS_TYPE_MEDICAL_TRANSPORT AIS_TYPE = 58
	// AIS_TYPE_NONECOMBATANT enum. Noncombatant ship according to RR Resolution No. 18
	AIS_TYPE_NONECOMBATANT AIS_TYPE = 59
	// AIS_TYPE_PASSENGER enum
	AIS_TYPE_PASSENGER AIS_TYPE = 60
	// AIS_TYPE_PASSENGER_HAZARDOUS_A enum
	AIS_TYPE_PASSENGER_HAZARDOUS_A AIS_TYPE = 61
	// AIS_TYPE_PASSENGER_HAZARDOUS_B enum
	AIS_TYPE_PASSENGER_HAZARDOUS_B AIS_TYPE = 62
	// AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C enum
	AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C AIS_TYPE = 63
	// AIS_TYPE_PASSENGER_HAZARDOUS_D enum
	AIS_TYPE_PASSENGER_HAZARDOUS_D AIS_TYPE = 64
	// AIS_TYPE_PASSENGER_RESERVED_1 enum
	AIS_TYPE_PASSENGER_RESERVED_1 AIS_TYPE = 65
	// AIS_TYPE_PASSENGER_RESERVED_2 enum
	AIS_TYPE_PASSENGER_RESERVED_2 AIS_TYPE = 66
	// AIS_TYPE_PASSENGER_RESERVED_3 enum
	AIS_TYPE_PASSENGER_RESERVED_3 AIS_TYPE = 67
	// AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4 enum
	AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4 AIS_TYPE = 68
	// AIS_TYPE_PASSENGER_UNKNOWN enum
	AIS_TYPE_PASSENGER_UNKNOWN AIS_TYPE = 69
	// AIS_TYPE_CARGO enum
	AIS_TYPE_CARGO AIS_TYPE = 70
	// AIS_TYPE_CARGO_HAZARDOUS_A enum
	AIS_TYPE_CARGO_HAZARDOUS_A AIS_TYPE = 71
	// AIS_TYPE_CARGO_HAZARDOUS_B enum
	AIS_TYPE_CARGO_HAZARDOUS_B AIS_TYPE = 72
	// AIS_TYPE_CARGO_HAZARDOUS_C enum
	AIS_TYPE_CARGO_HAZARDOUS_C AIS_TYPE = 73
	// AIS_TYPE_CARGO_HAZARDOUS_D enum
	AIS_TYPE_CARGO_HAZARDOUS_D AIS_TYPE = 74
	// AIS_TYPE_CARGO_RESERVED_1 enum
	AIS_TYPE_CARGO_RESERVED_1 AIS_TYPE = 75
	// AIS_TYPE_CARGO_RESERVED_2 enum
	AIS_TYPE_CARGO_RESERVED_2 AIS_TYPE = 76
	// AIS_TYPE_CARGO_RESERVED_3 enum
	AIS_TYPE_CARGO_RESERVED_3 AIS_TYPE = 77
	// AIS_TYPE_CARGO_RESERVED_4 enum
	AIS_TYPE_CARGO_RESERVED_4 AIS_TYPE = 78
	// AIS_TYPE_CARGO_UNKNOWN enum
	AIS_TYPE_CARGO_UNKNOWN AIS_TYPE = 79
	// AIS_TYPE_TANKER enum
	AIS_TYPE_TANKER AIS_TYPE = 80
	// AIS_TYPE_TANKER_HAZARDOUS_A enum
	AIS_TYPE_TANKER_HAZARDOUS_A AIS_TYPE = 81
	// AIS_TYPE_TANKER_HAZARDOUS_B enum
	AIS_TYPE_TANKER_HAZARDOUS_B AIS_TYPE = 82
	// AIS_TYPE_TANKER_HAZARDOUS_C enum
	AIS_TYPE_TANKER_HAZARDOUS_C AIS_TYPE = 83
	// AIS_TYPE_TANKER_HAZARDOUS_D enum
	AIS_TYPE_TANKER_HAZARDOUS_D AIS_TYPE = 84
	// AIS_TYPE_TANKER_RESERVED_1 enum
	AIS_TYPE_TANKER_RESERVED_1 AIS_TYPE = 85
	// AIS_TYPE_TANKER_RESERVED_2 enum
	AIS_TYPE_TANKER_RESERVED_2 AIS_TYPE = 86
	// AIS_TYPE_TANKER_RESERVED_3 enum
	AIS_TYPE_TANKER_RESERVED_3 AIS_TYPE = 87
	// AIS_TYPE_TANKER_RESERVED_4 enum
	AIS_TYPE_TANKER_RESERVED_4 AIS_TYPE = 88
	// AIS_TYPE_TANKER_UNKNOWN enum
	AIS_TYPE_TANKER_UNKNOWN AIS_TYPE = 89
	// AIS_TYPE_OTHER enum
	AIS_TYPE_OTHER AIS_TYPE = 90
	// AIS_TYPE_OTHER_HAZARDOUS_A enum
	AIS_TYPE_OTHER_HAZARDOUS_A AIS_TYPE = 91
	// AIS_TYPE_OTHER_HAZARDOUS_B enum
	AIS_TYPE_OTHER_HAZARDOUS_B AIS_TYPE = 92
	// AIS_TYPE_OTHER_HAZARDOUS_C enum
	AIS_TYPE_OTHER_HAZARDOUS_C AIS_TYPE = 93
	// AIS_TYPE_OTHER_HAZARDOUS_D enum
	AIS_TYPE_OTHER_HAZARDOUS_D AIS_TYPE = 94
	// AIS_TYPE_OTHER_RESERVED_1 enum
	AIS_TYPE_OTHER_RESERVED_1 AIS_TYPE = 95
	// AIS_TYPE_OTHER_RESERVED_2 enum
	AIS_TYPE_OTHER_RESERVED_2 AIS_TYPE = 96
	// AIS_TYPE_OTHER_RESERVED_3 enum
	AIS_TYPE_OTHER_RESERVED_3 AIS_TYPE = 97
	// AIS_TYPE_OTHER_RESERVED_4 enum
	AIS_TYPE_OTHER_RESERVED_4 AIS_TYPE = 98
	// AIS_TYPE_OTHER_UNKNOWN enum
	AIS_TYPE_OTHER_UNKNOWN AIS_TYPE = 99
)

func (e AIS_TYPE) String() string {
	if name, ok := map[AIS_TYPE]string{
		AIS_TYPE_UNKNOWN:                        "AIS_TYPE_UNKNOWN",
		AIS_TYPE_RESERVED_1:                     "AIS_TYPE_RESERVED_1",
		AIS_TYPE_RESERVED_2:                     "AIS_TYPE_RESERVED_2",
		AIS_TYPE_RESERVED_3:                     "AIS_TYPE_RESERVED_3",
		AIS_TYPE_RESERVED_4:                     "AIS_TYPE_RESERVED_4",
		AIS_TYPE_RESERVED_5:                     "AIS_TYPE_RESERVED_5",
		AIS_TYPE_RESERVED_6:                     "AIS_TYPE_RESERVED_6",
		AIS_TYPE_RESERVED_7:                     "AIS_TYPE_RESERVED_7",
		AIS_TYPE_RESERVED_8:                     "AIS_TYPE_RESERVED_8",
		AIS_TYPE_RESERVED_9:                     "AIS_TYPE_RESERVED_9",
		AIS_TYPE_RESERVED_10:                    "AIS_TYPE_RESERVED_10",
		AIS_TYPE_RESERVED_11:                    "AIS_TYPE_RESERVED_11",
		AIS_TYPE_RESERVED_12:                    "AIS_TYPE_RESERVED_12",
		AIS_TYPE_RESERVED_13:                    "AIS_TYPE_RESERVED_13",
		AIS_TYPE_RESERVED_14:                    "AIS_TYPE_RESERVED_14",
		AIS_TYPE_RESERVED_15:                    "AIS_TYPE_RESERVED_15",
		AIS_TYPE_RESERVED_16:                    "AIS_TYPE_RESERVED_16",
		AIS_TYPE_RESERVED_17:                    "AIS_TYPE_RESERVED_17",
		AIS_TYPE_RESERVED_18:                    "AIS_TYPE_RESERVED_18",
		AIS_TYPE_RESERVED_19:                    "AIS_TYPE_RESERVED_19",
		AIS_TYPE_WIG:                            "AIS_TYPE_WIG",
		AIS_TYPE_WIG_HAZARDOUS_A:                "AIS_TYPE_WIG_HAZARDOUS_A",
		AIS_TYPE_WIG_HAZARDOUS_B:                "AIS_TYPE_WIG_HAZARDOUS_B",
		AIS_TYPE_WIG_HAZARDOUS_C:                "AIS_TYPE_WIG_HAZARDOUS_C",
		AIS_TYPE_WIG_HAZARDOUS_D:                "AIS_TYPE_WIG_HAZARDOUS_D",
		AIS_TYPE_WIG_RESERVED_1:                 "AIS_TYPE_WIG_RESERVED_1",
		AIS_TYPE_WIG_RESERVED_2:                 "AIS_TYPE_WIG_RESERVED_2",
		AIS_TYPE_WIG_RESERVED_3:                 "AIS_TYPE_WIG_RESERVED_3",
		AIS_TYPE_WIG_RESERVED_4:                 "AIS_TYPE_WIG_RESERVED_4",
		AIS_TYPE_WIG_RESERVED_5:                 "AIS_TYPE_WIG_RESERVED_5",
		AIS_TYPE_FISHING:                        "AIS_TYPE_FISHING",
		AIS_TYPE_TOWING:                         "AIS_TYPE_TOWING",
		AIS_TYPE_TOWING_LARGE:                   "AIS_TYPE_TOWING_LARGE",
		AIS_TYPE_DREDGING:                       "AIS_TYPE_DREDGING",
		AIS_TYPE_DIVING:                         "AIS_TYPE_DIVING",
		AIS_TYPE_MILITARY:                       "AIS_TYPE_MILITARY",
		AIS_TYPE_SAILING:                        "AIS_TYPE_SAILING",
		AIS_TYPE_PLEASURE:                       "AIS_TYPE_PLEASURE",
		AIS_TYPE_RESERVED_20:                    "AIS_TYPE_RESERVED_20",
		AIS_TYPE_RESERVED_21:                    "AIS_TYPE_RESERVED_21",
		AIS_TYPE_HSC:                            "AIS_TYPE_HSC",
		AIS_TYPE_HSC_HAZARDOUS_A:                "AIS_TYPE_HSC_HAZARDOUS_A",
		AIS_TYPE_HSC_HAZARDOUS_B:                "AIS_TYPE_HSC_HAZARDOUS_B",
		AIS_TYPE_HSC_HAZARDOUS_C:                "AIS_TYPE_HSC_HAZARDOUS_C",
		AIS_TYPE_HSC_HAZARDOUS_D:                "AIS_TYPE_HSC_HAZARDOUS_D",
		AIS_TYPE_HSC_RESERVED_1:                 "AIS_TYPE_HSC_RESERVED_1",
		AIS_TYPE_HSC_RESERVED_2:                 "AIS_TYPE_HSC_RESERVED_2",
		AIS_TYPE_HSC_RESERVED_3:                 "AIS_TYPE_HSC_RESERVED_3",
		AIS_TYPE_HSC_RESERVED_4:                 "AIS_TYPE_HSC_RESERVED_4",
		AIS_TYPE_HSC_UNKNOWN:                    "AIS_TYPE_HSC_UNKNOWN",
		AIS_TYPE_PILOT:                          "AIS_TYPE_PILOT",
		AIS_TYPE_SAR:                            "AIS_TYPE_SAR",
		AIS_TYPE_TUG:                            "AIS_TYPE_TUG",
		AIS_TYPE_PORT_TENDER:                    "AIS_TYPE_PORT_TENDER",
		AIS_TYPE_ANTI_POLLUTION:                 "AIS_TYPE_ANTI_POLLUTION",
		AIS_TYPE_LAW_ENFORCEMENT:                "AIS_TYPE_LAW_ENFORCEMENT",
		AIS_TYPE_SPARE_LOCAL_1:                  "AIS_TYPE_SPARE_LOCAL_1",
		AIS_TYPE_SPARE_LOCAL_2:                  "AIS_TYPE_SPARE_LOCAL_2",
		AIS_TYPE_MEDICAL_TRANSPORT:              "AIS_TYPE_MEDICAL_TRANSPORT",
		AIS_TYPE_NONECOMBATANT:                  "AIS_TYPE_NONECOMBATANT",
		AIS_TYPE_PASSENGER:                      "AIS_TYPE_PASSENGER",
		AIS_TYPE_PASSENGER_HAZARDOUS_A:          "AIS_TYPE_PASSENGER_HAZARDOUS_A",
		AIS_TYPE_PASSENGER_HAZARDOUS_B:          "AIS_TYPE_PASSENGER_HAZARDOUS_B",
		AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C: "AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C",
		AIS_TYPE_PASSENGER_HAZARDOUS_D:          "AIS_TYPE_PASSENGER_HAZARDOUS_D",
		AIS_TYPE_PASSENGER_RESERVED_1:           "AIS_TYPE_PASSENGER_RESERVED_1",
		AIS_TYPE_PASSENGER_RESERVED_2:           "AIS_TYPE_PASSENGER_RESERVED_2",
		AIS_TYPE_PASSENGER_RESERVED_3:           "AIS_TYPE_PASSENGER_RESERVED_3",
		AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4:  "AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4",
		AIS_TYPE_PASSENGER_UNKNOWN:              "AIS_TYPE_PASSENGER_UNKNOWN",
		AIS_TYPE_CARGO:                          "AIS_TYPE_CARGO",
		AIS_TYPE_CARGO_HAZARDOUS_A:              "AIS_TYPE_CARGO_HAZARDOUS_A",
		AIS_TYPE_CARGO_HAZARDOUS_B:              "AIS_TYPE_CARGO_HAZARDOUS_B",
		AIS_TYPE_CARGO_HAZARDOUS_C:              "AIS_TYPE_CARGO_HAZARDOUS_C",
		AIS_TYPE_CARGO_HAZARDOUS_D:              "AIS_TYPE_CARGO_HAZARDOUS_D",
		AIS_TYPE_CARGO_RESERVED_1:               "AIS_TYPE_CARGO_RESERVED_1",
		AIS_TYPE_CARGO_RESERVED_2:               "AIS_TYPE_CARGO_RESERVED_2",
		AIS_TYPE_CARGO_RESERVED_3:               "AIS_TYPE_CARGO_RESERVED_3",
		AIS_TYPE_CARGO_RESERVED_4:               "AIS_TYPE_CARGO_RESERVED_4",
		AIS_TYPE_CARGO_UNKNOWN:                  "AIS_TYPE_CARGO_UNKNOWN",
		AIS_TYPE_TANKER:                         "AIS_TYPE_TANKER",
		AIS_TYPE_TANKER_HAZARDOUS_A:             "AIS_TYPE_TANKER_HAZARDOUS_A",
		AIS_TYPE_TANKER_HAZARDOUS_B:             "AIS_TYPE_TANKER_HAZARDOUS_B",
		AIS_TYPE_TANKER_HAZARDOUS_C:             "AIS_TYPE_TANKER_HAZARDOUS_C",
		AIS_TYPE_TANKER_HAZARDOUS_D:             "AIS_TYPE_TANKER_HAZARDOUS_D",
		AIS_TYPE_TANKER_RESERVED_1:              "AIS_TYPE_TANKER_RESERVED_1",
		AIS_TYPE_TANKER_RESERVED_2:              "AIS_TYPE_TANKER_RESERVED_2",
		AIS_TYPE_TANKER_RESERVED_3:              "AIS_TYPE_TANKER_RESERVED_3",
		AIS_TYPE_TANKER_RESERVED_4:              "AIS_TYPE_TANKER_RESERVED_4",
		AIS_TYPE_TANKER_UNKNOWN:                 "AIS_TYPE_TANKER_UNKNOWN",
		AIS_TYPE_OTHER:                          "AIS_TYPE_OTHER",
		AIS_TYPE_OTHER_HAZARDOUS_A:              "AIS_TYPE_OTHER_HAZARDOUS_A",
		AIS_TYPE_OTHER_HAZARDOUS_B:              "AIS_TYPE_OTHER_HAZARDOUS_B",
		AIS_TYPE_OTHER_HAZARDOUS_C:              "AIS_TYPE_OTHER_HAZARDOUS_C",
		AIS_TYPE_OTHER_HAZARDOUS_D:              "AIS_TYPE_OTHER_HAZARDOUS_D",
		AIS_TYPE_OTHER_RESERVED_1:               "AIS_TYPE_OTHER_RESERVED_1",
		AIS_TYPE_OTHER_RESERVED_2:               "AIS_TYPE_OTHER_RESERVED_2",
		AIS_TYPE_OTHER_RESERVED_3:               "AIS_TYPE_OTHER_RESERVED_3",
		AIS_TYPE_OTHER_RESERVED_4:               "AIS_TYPE_OTHER_RESERVED_4",
		AIS_TYPE_OTHER_UNKNOWN:                  "AIS_TYPE_OTHER_UNKNOWN",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("AIS_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects AIS_TYPE enums
func (e AIS_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []AIS_TYPE{
		AIS_TYPE_UNKNOWN,
		AIS_TYPE_RESERVED_1,
		AIS_TYPE_RESERVED_2,
		AIS_TYPE_RESERVED_3,
		AIS_TYPE_RESERVED_4,
		AIS_TYPE_RESERVED_5,
		AIS_TYPE_RESERVED_6,
		AIS_TYPE_RESERVED_7,
		AIS_TYPE_RESERVED_8,
		AIS_TYPE_RESERVED_9,
		AIS_TYPE_RESERVED_10,
		AIS_TYPE_RESERVED_11,
		AIS_TYPE_RESERVED_12,
		AIS_TYPE_RESERVED_13,
		AIS_TYPE_RESERVED_14,
		AIS_TYPE_RESERVED_15,
		AIS_TYPE_RESERVED_16,
		AIS_TYPE_RESERVED_17,
		AIS_TYPE_RESERVED_18,
		AIS_TYPE_RESERVED_19,
		AIS_TYPE_WIG,
		AIS_TYPE_WIG_HAZARDOUS_A,
		AIS_TYPE_WIG_HAZARDOUS_B,
		AIS_TYPE_WIG_HAZARDOUS_C,
		AIS_TYPE_WIG_HAZARDOUS_D,
		AIS_TYPE_WIG_RESERVED_1,
		AIS_TYPE_WIG_RESERVED_2,
		AIS_TYPE_WIG_RESERVED_3,
		AIS_TYPE_WIG_RESERVED_4,
		AIS_TYPE_WIG_RESERVED_5,
		AIS_TYPE_FISHING,
		AIS_TYPE_TOWING,
		AIS_TYPE_TOWING_LARGE,
		AIS_TYPE_DREDGING,
		AIS_TYPE_DIVING,
		AIS_TYPE_MILITARY,
		AIS_TYPE_SAILING,
		AIS_TYPE_PLEASURE,
		AIS_TYPE_RESERVED_20,
		AIS_TYPE_RESERVED_21,
		AIS_TYPE_HSC,
		AIS_TYPE_HSC_HAZARDOUS_A,
		AIS_TYPE_HSC_HAZARDOUS_B,
		AIS_TYPE_HSC_HAZARDOUS_C,
		AIS_TYPE_HSC_HAZARDOUS_D,
		AIS_TYPE_HSC_RESERVED_1,
		AIS_TYPE_HSC_RESERVED_2,
		AIS_TYPE_HSC_RESERVED_3,
		AIS_TYPE_HSC_RESERVED_4,
		AIS_TYPE_HSC_UNKNOWN,
		AIS_TYPE_PILOT,
		AIS_TYPE_SAR,
		AIS_TYPE_TUG,
		AIS_TYPE_PORT_TENDER,
		AIS_TYPE_ANTI_POLLUTION,
		AIS_TYPE_LAW_ENFORCEMENT,
		AIS_TYPE_SPARE_LOCAL_1,
		AIS_TYPE_SPARE_LOCAL_2,
		AIS_TYPE_MEDICAL_TRANSPORT,
		AIS_TYPE_NONECOMBATANT,
		AIS_TYPE_PASSENGER,
		AIS_TYPE_PASSENGER_HAZARDOUS_A,
		AIS_TYPE_PASSENGER_HAZARDOUS_B,
		AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C,
		AIS_TYPE_PASSENGER_HAZARDOUS_D,
		AIS_TYPE_PASSENGER_RESERVED_1,
		AIS_TYPE_PASSENGER_RESERVED_2,
		AIS_TYPE_PASSENGER_RESERVED_3,
		AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4,
		AIS_TYPE_PASSENGER_UNKNOWN,
		AIS_TYPE_CARGO,
		AIS_TYPE_CARGO_HAZARDOUS_A,
		AIS_TYPE_CARGO_HAZARDOUS_B,
		AIS_TYPE_CARGO_HAZARDOUS_C,
		AIS_TYPE_CARGO_HAZARDOUS_D,
		AIS_TYPE_CARGO_RESERVED_1,
		AIS_TYPE_CARGO_RESERVED_2,
		AIS_TYPE_CARGO_RESERVED_3,
		AIS_TYPE_CARGO_RESERVED_4,
		AIS_TYPE_CARGO_UNKNOWN,
		AIS_TYPE_TANKER,
		AIS_TYPE_TANKER_HAZARDOUS_A,
		AIS_TYPE_TANKER_HAZARDOUS_B,
		AIS_TYPE_TANKER_HAZARDOUS_C,
		AIS_TYPE_TANKER_HAZARDOUS_D,
		AIS_TYPE_TANKER_RESERVED_1,
		AIS_TYPE_TANKER_RESERVED_2,
		AIS_TYPE_TANKER_RESERVED_3,
		AIS_TYPE_TANKER_RESERVED_4,
		AIS_TYPE_TANKER_UNKNOWN,
		AIS_TYPE_OTHER,
		AIS_TYPE_OTHER_HAZARDOUS_A,
		AIS_TYPE_OTHER_HAZARDOUS_B,
		AIS_TYPE_OTHER_HAZARDOUS_C,
		AIS_TYPE_OTHER_HAZARDOUS_D,
		AIS_TYPE_OTHER_RESERVED_1,
		AIS_TYPE_OTHER_RESERVED_2,
		AIS_TYPE_OTHER_RESERVED_3,
		AIS_TYPE_OTHER_RESERVED_4,
		AIS_TYPE_OTHER_UNKNOWN,
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

// AIS_NAV_STATUS type. Navigational status of AIS vessel, enum duplicated from AIS standard, https://gpsd.gitlab.io/gpsd/AIVDM.html
type AIS_NAV_STATUS int

func (e AIS_NAV_STATUS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// UNDER_WAY enum. Under way using engine
	UNDER_WAY AIS_NAV_STATUS = 0
	// AIS_NAV_ANCHORED enum
	AIS_NAV_ANCHORED AIS_NAV_STATUS = 1
	// AIS_NAV_UN_COMMANDED enum
	AIS_NAV_UN_COMMANDED AIS_NAV_STATUS = 2
	// AIS_NAV_RESTRICTED_MANOEUVERABILITY enum
	AIS_NAV_RESTRICTED_MANOEUVERABILITY AIS_NAV_STATUS = 3
	// AIS_NAV_DRAUGHT_CONSTRAINED enum
	AIS_NAV_DRAUGHT_CONSTRAINED AIS_NAV_STATUS = 4
	// AIS_NAV_MOORED enum
	AIS_NAV_MOORED AIS_NAV_STATUS = 5
	// AIS_NAV_AGROUND enum
	AIS_NAV_AGROUND AIS_NAV_STATUS = 6
	// AIS_NAV_FISHING enum
	AIS_NAV_FISHING AIS_NAV_STATUS = 7
	// AIS_NAV_SAILING enum
	AIS_NAV_SAILING AIS_NAV_STATUS = 8
	// AIS_NAV_RESERVED_HSC enum
	AIS_NAV_RESERVED_HSC AIS_NAV_STATUS = 9
	// AIS_NAV_RESERVED_WIG enum
	AIS_NAV_RESERVED_WIG AIS_NAV_STATUS = 10
	// AIS_NAV_RESERVED_1 enum
	AIS_NAV_RESERVED_1 AIS_NAV_STATUS = 11
	// AIS_NAV_RESERVED_2 enum
	AIS_NAV_RESERVED_2 AIS_NAV_STATUS = 12
	// AIS_NAV_RESERVED_3 enum
	AIS_NAV_RESERVED_3 AIS_NAV_STATUS = 13
	// AIS_NAV_AIS_SART enum. Search And Rescue Transponder
	AIS_NAV_AIS_SART AIS_NAV_STATUS = 14
	// AIS_NAV_UNKNOWN enum. Not available (default)
	AIS_NAV_UNKNOWN AIS_NAV_STATUS = 15
)

func (e AIS_NAV_STATUS) String() string {
	if name, ok := map[AIS_NAV_STATUS]string{
		UNDER_WAY:                           "UNDER_WAY",
		AIS_NAV_ANCHORED:                    "AIS_NAV_ANCHORED",
		AIS_NAV_UN_COMMANDED:                "AIS_NAV_UN_COMMANDED",
		AIS_NAV_RESTRICTED_MANOEUVERABILITY: "AIS_NAV_RESTRICTED_MANOEUVERABILITY",
		AIS_NAV_DRAUGHT_CONSTRAINED:         "AIS_NAV_DRAUGHT_CONSTRAINED",
		AIS_NAV_MOORED:                      "AIS_NAV_MOORED",
		AIS_NAV_AGROUND:                     "AIS_NAV_AGROUND",
		AIS_NAV_FISHING:                     "AIS_NAV_FISHING",
		AIS_NAV_SAILING:                     "AIS_NAV_SAILING",
		AIS_NAV_RESERVED_HSC:                "AIS_NAV_RESERVED_HSC",
		AIS_NAV_RESERVED_WIG:                "AIS_NAV_RESERVED_WIG",
		AIS_NAV_RESERVED_1:                  "AIS_NAV_RESERVED_1",
		AIS_NAV_RESERVED_2:                  "AIS_NAV_RESERVED_2",
		AIS_NAV_RESERVED_3:                  "AIS_NAV_RESERVED_3",
		AIS_NAV_AIS_SART:                    "AIS_NAV_AIS_SART",
		AIS_NAV_UNKNOWN:                     "AIS_NAV_UNKNOWN",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("AIS_NAV_STATUS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects AIS_NAV_STATUS enums
func (e AIS_NAV_STATUS) Bitmask() string {
	bitmap := ""
	for _, entry := range []AIS_NAV_STATUS{
		UNDER_WAY,
		AIS_NAV_ANCHORED,
		AIS_NAV_UN_COMMANDED,
		AIS_NAV_RESTRICTED_MANOEUVERABILITY,
		AIS_NAV_DRAUGHT_CONSTRAINED,
		AIS_NAV_MOORED,
		AIS_NAV_AGROUND,
		AIS_NAV_FISHING,
		AIS_NAV_SAILING,
		AIS_NAV_RESERVED_HSC,
		AIS_NAV_RESERVED_WIG,
		AIS_NAV_RESERVED_1,
		AIS_NAV_RESERVED_2,
		AIS_NAV_RESERVED_3,
		AIS_NAV_AIS_SART,
		AIS_NAV_UNKNOWN,
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

// AIS_FLAGS type. These flags are used in the AIS_VESSEL.fields bitmask to indicate validity of data in the other message fields. When set, the data is valid.
type AIS_FLAGS int

func (e AIS_FLAGS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// AIS_FLAGS_POSITION_ACCURACY enum. 1 = Position accuracy less than 10m, 0 = position accuracy greater than 10m
	AIS_FLAGS_POSITION_ACCURACY AIS_FLAGS = 1
	// AIS_FLAGS_VALID_COG enum
	AIS_FLAGS_VALID_COG AIS_FLAGS = 2
	// AIS_FLAGS_VALID_VELOCITY enum
	AIS_FLAGS_VALID_VELOCITY AIS_FLAGS = 4
	// AIS_FLAGS_HIGH_VELOCITY enum. 1 = Velocity over 52.5765m/s (102.2 knots)
	AIS_FLAGS_HIGH_VELOCITY AIS_FLAGS = 8
	// AIS_FLAGS_VALID_TURN_RATE enum
	AIS_FLAGS_VALID_TURN_RATE AIS_FLAGS = 16
	// AIS_FLAGS_TURN_RATE_SIGN_ONLY enum. Only the sign of the returned turn rate value is valid, either greater than 5deg/30s or less than -5deg/30s
	AIS_FLAGS_TURN_RATE_SIGN_ONLY AIS_FLAGS = 32
	// AIS_FLAGS_VALID_DIMENSIONS enum
	AIS_FLAGS_VALID_DIMENSIONS AIS_FLAGS = 64
	// AIS_FLAGS_LARGE_BOW_DIMENSION enum. Distance to bow is larger than 511m
	AIS_FLAGS_LARGE_BOW_DIMENSION AIS_FLAGS = 128
	// AIS_FLAGS_LARGE_STERN_DIMENSION enum. Distance to stern is larger than 511m
	AIS_FLAGS_LARGE_STERN_DIMENSION AIS_FLAGS = 256
	// AIS_FLAGS_LARGE_PORT_DIMENSION enum. Distance to port side is larger than 63m
	AIS_FLAGS_LARGE_PORT_DIMENSION AIS_FLAGS = 512
	// AIS_FLAGS_LARGE_STARBOARD_DIMENSION enum. Distance to starboard side is larger than 63m
	AIS_FLAGS_LARGE_STARBOARD_DIMENSION AIS_FLAGS = 1024
	// AIS_FLAGS_VALID_CALLSIGN enum
	AIS_FLAGS_VALID_CALLSIGN AIS_FLAGS = 2048
	// AIS_FLAGS_VALID_NAME enum
	AIS_FLAGS_VALID_NAME AIS_FLAGS = 4096
)

func (e AIS_FLAGS) String() string {
	if name, ok := map[AIS_FLAGS]string{
		AIS_FLAGS_POSITION_ACCURACY:         "AIS_FLAGS_POSITION_ACCURACY",
		AIS_FLAGS_VALID_COG:                 "AIS_FLAGS_VALID_COG",
		AIS_FLAGS_VALID_VELOCITY:            "AIS_FLAGS_VALID_VELOCITY",
		AIS_FLAGS_HIGH_VELOCITY:             "AIS_FLAGS_HIGH_VELOCITY",
		AIS_FLAGS_VALID_TURN_RATE:           "AIS_FLAGS_VALID_TURN_RATE",
		AIS_FLAGS_TURN_RATE_SIGN_ONLY:       "AIS_FLAGS_TURN_RATE_SIGN_ONLY",
		AIS_FLAGS_VALID_DIMENSIONS:          "AIS_FLAGS_VALID_DIMENSIONS",
		AIS_FLAGS_LARGE_BOW_DIMENSION:       "AIS_FLAGS_LARGE_BOW_DIMENSION",
		AIS_FLAGS_LARGE_STERN_DIMENSION:     "AIS_FLAGS_LARGE_STERN_DIMENSION",
		AIS_FLAGS_LARGE_PORT_DIMENSION:      "AIS_FLAGS_LARGE_PORT_DIMENSION",
		AIS_FLAGS_LARGE_STARBOARD_DIMENSION: "AIS_FLAGS_LARGE_STARBOARD_DIMENSION",
		AIS_FLAGS_VALID_CALLSIGN:            "AIS_FLAGS_VALID_CALLSIGN",
		AIS_FLAGS_VALID_NAME:                "AIS_FLAGS_VALID_NAME",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("AIS_FLAGS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects AIS_FLAGS enums
func (e AIS_FLAGS) Bitmask() string {
	bitmap := ""
	for _, entry := range []AIS_FLAGS{
		AIS_FLAGS_POSITION_ACCURACY,
		AIS_FLAGS_VALID_COG,
		AIS_FLAGS_VALID_VELOCITY,
		AIS_FLAGS_HIGH_VELOCITY,
		AIS_FLAGS_VALID_TURN_RATE,
		AIS_FLAGS_TURN_RATE_SIGN_ONLY,
		AIS_FLAGS_VALID_DIMENSIONS,
		AIS_FLAGS_LARGE_BOW_DIMENSION,
		AIS_FLAGS_LARGE_STERN_DIMENSION,
		AIS_FLAGS_LARGE_PORT_DIMENSION,
		AIS_FLAGS_LARGE_STARBOARD_DIMENSION,
		AIS_FLAGS_VALID_CALLSIGN,
		AIS_FLAGS_VALID_NAME,
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

// FAILURE_UNIT type. List of possible units where failures can be injected.
type FAILURE_UNIT int

func (e FAILURE_UNIT) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// FAILURE_UNIT_SENSOR_GYRO enum
	FAILURE_UNIT_SENSOR_GYRO FAILURE_UNIT = 0
	// FAILURE_UNIT_SENSOR_ACCEL enum
	FAILURE_UNIT_SENSOR_ACCEL FAILURE_UNIT = 1
	// FAILURE_UNIT_SENSOR_MAG enum
	FAILURE_UNIT_SENSOR_MAG FAILURE_UNIT = 2
	// FAILURE_UNIT_SENSOR_BARO enum
	FAILURE_UNIT_SENSOR_BARO FAILURE_UNIT = 3
	// FAILURE_UNIT_SENSOR_GPS enum
	FAILURE_UNIT_SENSOR_GPS FAILURE_UNIT = 4
	// FAILURE_UNIT_SENSOR_OPTICAL_FLOW enum
	FAILURE_UNIT_SENSOR_OPTICAL_FLOW FAILURE_UNIT = 5
	// FAILURE_UNIT_SENSOR_VIO enum
	FAILURE_UNIT_SENSOR_VIO FAILURE_UNIT = 6
	// FAILURE_UNIT_SENSOR_DISTANCE_SENSOR enum
	FAILURE_UNIT_SENSOR_DISTANCE_SENSOR FAILURE_UNIT = 7
	// FAILURE_UNIT_SENSOR_AIRSPEED enum
	FAILURE_UNIT_SENSOR_AIRSPEED FAILURE_UNIT = 8
	// FAILURE_UNIT_SYSTEM_BATTERY enum
	FAILURE_UNIT_SYSTEM_BATTERY FAILURE_UNIT = 100
	// FAILURE_UNIT_SYSTEM_MOTOR enum
	FAILURE_UNIT_SYSTEM_MOTOR FAILURE_UNIT = 101
	// FAILURE_UNIT_SYSTEM_SERVO enum
	FAILURE_UNIT_SYSTEM_SERVO FAILURE_UNIT = 102
	// FAILURE_UNIT_SYSTEM_AVOIDANCE enum
	FAILURE_UNIT_SYSTEM_AVOIDANCE FAILURE_UNIT = 103
	// FAILURE_UNIT_SYSTEM_RC_SIGNAL enum
	FAILURE_UNIT_SYSTEM_RC_SIGNAL FAILURE_UNIT = 104
	// FAILURE_UNIT_SYSTEM_MAVLINK_SIGNAL enum
	FAILURE_UNIT_SYSTEM_MAVLINK_SIGNAL FAILURE_UNIT = 105
)

func (e FAILURE_UNIT) String() string {
	if name, ok := map[FAILURE_UNIT]string{
		FAILURE_UNIT_SENSOR_GYRO:            "FAILURE_UNIT_SENSOR_GYRO",
		FAILURE_UNIT_SENSOR_ACCEL:           "FAILURE_UNIT_SENSOR_ACCEL",
		FAILURE_UNIT_SENSOR_MAG:             "FAILURE_UNIT_SENSOR_MAG",
		FAILURE_UNIT_SENSOR_BARO:            "FAILURE_UNIT_SENSOR_BARO",
		FAILURE_UNIT_SENSOR_GPS:             "FAILURE_UNIT_SENSOR_GPS",
		FAILURE_UNIT_SENSOR_OPTICAL_FLOW:    "FAILURE_UNIT_SENSOR_OPTICAL_FLOW",
		FAILURE_UNIT_SENSOR_VIO:             "FAILURE_UNIT_SENSOR_VIO",
		FAILURE_UNIT_SENSOR_DISTANCE_SENSOR: "FAILURE_UNIT_SENSOR_DISTANCE_SENSOR",
		FAILURE_UNIT_SENSOR_AIRSPEED:        "FAILURE_UNIT_SENSOR_AIRSPEED",
		FAILURE_UNIT_SYSTEM_BATTERY:         "FAILURE_UNIT_SYSTEM_BATTERY",
		FAILURE_UNIT_SYSTEM_MOTOR:           "FAILURE_UNIT_SYSTEM_MOTOR",
		FAILURE_UNIT_SYSTEM_SERVO:           "FAILURE_UNIT_SYSTEM_SERVO",
		FAILURE_UNIT_SYSTEM_AVOIDANCE:       "FAILURE_UNIT_SYSTEM_AVOIDANCE",
		FAILURE_UNIT_SYSTEM_RC_SIGNAL:       "FAILURE_UNIT_SYSTEM_RC_SIGNAL",
		FAILURE_UNIT_SYSTEM_MAVLINK_SIGNAL:  "FAILURE_UNIT_SYSTEM_MAVLINK_SIGNAL",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("FAILURE_UNIT_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects FAILURE_UNIT enums
func (e FAILURE_UNIT) Bitmask() string {
	bitmap := ""
	for _, entry := range []FAILURE_UNIT{
		FAILURE_UNIT_SENSOR_GYRO,
		FAILURE_UNIT_SENSOR_ACCEL,
		FAILURE_UNIT_SENSOR_MAG,
		FAILURE_UNIT_SENSOR_BARO,
		FAILURE_UNIT_SENSOR_GPS,
		FAILURE_UNIT_SENSOR_OPTICAL_FLOW,
		FAILURE_UNIT_SENSOR_VIO,
		FAILURE_UNIT_SENSOR_DISTANCE_SENSOR,
		FAILURE_UNIT_SENSOR_AIRSPEED,
		FAILURE_UNIT_SYSTEM_BATTERY,
		FAILURE_UNIT_SYSTEM_MOTOR,
		FAILURE_UNIT_SYSTEM_SERVO,
		FAILURE_UNIT_SYSTEM_AVOIDANCE,
		FAILURE_UNIT_SYSTEM_RC_SIGNAL,
		FAILURE_UNIT_SYSTEM_MAVLINK_SIGNAL,
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

// FAILURE_TYPE type. List of possible failure type to inject.
type FAILURE_TYPE int

func (e FAILURE_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// FAILURE_TYPE_OK enum. No failure injected, used to reset a previous failure
	FAILURE_TYPE_OK FAILURE_TYPE = 0
	// FAILURE_TYPE_OFF enum. Sets unit off, so completely non-responsive
	FAILURE_TYPE_OFF FAILURE_TYPE = 1
	// FAILURE_TYPE_STUCK enum. Unit is stuck e.g. keeps reporting the same value
	FAILURE_TYPE_STUCK FAILURE_TYPE = 2
	// FAILURE_TYPE_GARBAGE enum. Unit is reporting complete garbage
	FAILURE_TYPE_GARBAGE FAILURE_TYPE = 3
	// FAILURE_TYPE_WRONG enum. Unit is consistently wrong
	FAILURE_TYPE_WRONG FAILURE_TYPE = 4
	// FAILURE_TYPE_SLOW enum. Unit is slow, so e.g. reporting at slower than expected rate
	FAILURE_TYPE_SLOW FAILURE_TYPE = 5
	// FAILURE_TYPE_DELAYED enum. Data of unit is delayed in time
	FAILURE_TYPE_DELAYED FAILURE_TYPE = 6
	// FAILURE_TYPE_INTERMITTENT enum. Unit is sometimes working, sometimes not
	FAILURE_TYPE_INTERMITTENT FAILURE_TYPE = 7
)

func (e FAILURE_TYPE) String() string {
	if name, ok := map[FAILURE_TYPE]string{
		FAILURE_TYPE_OK:           "FAILURE_TYPE_OK",
		FAILURE_TYPE_OFF:          "FAILURE_TYPE_OFF",
		FAILURE_TYPE_STUCK:        "FAILURE_TYPE_STUCK",
		FAILURE_TYPE_GARBAGE:      "FAILURE_TYPE_GARBAGE",
		FAILURE_TYPE_WRONG:        "FAILURE_TYPE_WRONG",
		FAILURE_TYPE_SLOW:         "FAILURE_TYPE_SLOW",
		FAILURE_TYPE_DELAYED:      "FAILURE_TYPE_DELAYED",
		FAILURE_TYPE_INTERMITTENT: "FAILURE_TYPE_INTERMITTENT",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("FAILURE_TYPE_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects FAILURE_TYPE enums
func (e FAILURE_TYPE) Bitmask() string {
	bitmap := ""
	for _, entry := range []FAILURE_TYPE{
		FAILURE_TYPE_OK,
		FAILURE_TYPE_OFF,
		FAILURE_TYPE_STUCK,
		FAILURE_TYPE_GARBAGE,
		FAILURE_TYPE_WRONG,
		FAILURE_TYPE_SLOW,
		FAILURE_TYPE_DELAYED,
		FAILURE_TYPE_INTERMITTENT,
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

// MAV_WINCH_STATUS_FLAG type. Winch status flags used in WINCH_STATUS
type MAV_WINCH_STATUS_FLAG int

func (e MAV_WINCH_STATUS_FLAG) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAV_WINCH_STATUS_HEALTHY enum. Winch is healthy
	MAV_WINCH_STATUS_HEALTHY MAV_WINCH_STATUS_FLAG = 1
	// MAV_WINCH_STATUS_FULLY_RETRACTED enum. Winch thread is fully retracted
	MAV_WINCH_STATUS_FULLY_RETRACTED MAV_WINCH_STATUS_FLAG = 2
	// MAV_WINCH_STATUS_MOVING enum. Winch motor is moving
	MAV_WINCH_STATUS_MOVING MAV_WINCH_STATUS_FLAG = 4
	// MAV_WINCH_STATUS_CLUTCH_ENGAGED enum. Winch clutch is engaged allowing motor to move freely
	MAV_WINCH_STATUS_CLUTCH_ENGAGED MAV_WINCH_STATUS_FLAG = 8
)

func (e MAV_WINCH_STATUS_FLAG) String() string {
	if name, ok := map[MAV_WINCH_STATUS_FLAG]string{
		MAV_WINCH_STATUS_HEALTHY:         "MAV_WINCH_STATUS_HEALTHY",
		MAV_WINCH_STATUS_FULLY_RETRACTED: "MAV_WINCH_STATUS_FULLY_RETRACTED",
		MAV_WINCH_STATUS_MOVING:          "MAV_WINCH_STATUS_MOVING",
		MAV_WINCH_STATUS_CLUTCH_ENGAGED:  "MAV_WINCH_STATUS_CLUTCH_ENGAGED",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAV_WINCH_STATUS_FLAG_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAV_WINCH_STATUS_FLAG enums
func (e MAV_WINCH_STATUS_FLAG) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAV_WINCH_STATUS_FLAG{
		MAV_WINCH_STATUS_HEALTHY,
		MAV_WINCH_STATUS_FULLY_RETRACTED,
		MAV_WINCH_STATUS_MOVING,
		MAV_WINCH_STATUS_CLUTCH_ENGAGED,
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

// MAG_CAL_STATUS type
type MAG_CAL_STATUS int

func (e MAG_CAL_STATUS) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

const (
	// MAG_CAL_NOT_STARTED enum
	MAG_CAL_NOT_STARTED MAG_CAL_STATUS = 0
	// MAG_CAL_WAITING_TO_START enum
	MAG_CAL_WAITING_TO_START MAG_CAL_STATUS = 1
	// MAG_CAL_RUNNING_STEP_ONE enum
	MAG_CAL_RUNNING_STEP_ONE MAG_CAL_STATUS = 2
	// MAG_CAL_RUNNING_STEP_TWO enum
	MAG_CAL_RUNNING_STEP_TWO MAG_CAL_STATUS = 3
	// MAG_CAL_SUCCESS enum
	MAG_CAL_SUCCESS MAG_CAL_STATUS = 4
	// MAG_CAL_FAILED enum
	MAG_CAL_FAILED MAG_CAL_STATUS = 5
	// MAG_CAL_BAD_ORIENTATION enum
	MAG_CAL_BAD_ORIENTATION MAG_CAL_STATUS = 6
	// MAG_CAL_BAD_RADIUS enum
	MAG_CAL_BAD_RADIUS MAG_CAL_STATUS = 7
)

func (e MAG_CAL_STATUS) String() string {
	if name, ok := map[MAG_CAL_STATUS]string{
		MAG_CAL_NOT_STARTED:      "MAG_CAL_NOT_STARTED",
		MAG_CAL_WAITING_TO_START: "MAG_CAL_WAITING_TO_START",
		MAG_CAL_RUNNING_STEP_ONE: "MAG_CAL_RUNNING_STEP_ONE",
		MAG_CAL_RUNNING_STEP_TWO: "MAG_CAL_RUNNING_STEP_TWO",
		MAG_CAL_SUCCESS:          "MAG_CAL_SUCCESS",
		MAG_CAL_FAILED:           "MAG_CAL_FAILED",
		MAG_CAL_BAD_ORIENTATION:  "MAG_CAL_BAD_ORIENTATION",
		MAG_CAL_BAD_RADIUS:       "MAG_CAL_BAD_RADIUS",
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("MAG_CAL_STATUS_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects MAG_CAL_STATUS enums
func (e MAG_CAL_STATUS) Bitmask() string {
	bitmap := ""
	for _, entry := range []MAG_CAL_STATUS{
		MAG_CAL_NOT_STARTED,
		MAG_CAL_WAITING_TO_START,
		MAG_CAL_RUNNING_STEP_ONE,
		MAG_CAL_RUNNING_STEP_TWO,
		MAG_CAL_SUCCESS,
		MAG_CAL_FAILED,
		MAG_CAL_BAD_ORIENTATION,
		MAG_CAL_BAD_RADIUS,
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

// MAV_AUTOPILOT type. Micro air vehicle / autopilot classes. This identifies the individual model.
type MAV_AUTOPILOT int

func (e MAV_AUTOPILOT) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

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

func (e MAV_TYPE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

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

func (e MAV_MODE_FLAG) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

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

func (e MAV_MODE_FLAG_DECODE_POSITION) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

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

func (e MAV_STATE) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

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

func (e MAV_COMPONENT) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(int(e))), nil
}

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
