/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package matrixpilot

import (
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/asmyasnikov/go-mavlink/mavlink/packet"
	"github.com/asmyasnikov/go-mavlink/mavlink/register"
)

func init() {
	for msgID, info := range map[message.MessageID]register.MessageInfo{
		MSG_ID_FLEXIFUNCTION_SET: {
			"MSG_ID_FLEXIFUNCTION_SET",
			2,
			181,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FlexifunctionSet)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_FLEXIFUNCTION_READ_REQ: {
			"MSG_ID_FLEXIFUNCTION_READ_REQ",
			6,
			26,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FlexifunctionReadReq)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_FLEXIFUNCTION_BUFFER_FUNCTION: {
			"MSG_ID_FLEXIFUNCTION_BUFFER_FUNCTION",
			58,
			101,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FlexifunctionBufferFunction)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_FLEXIFUNCTION_BUFFER_FUNCTION_ACK: {
			"MSG_ID_FLEXIFUNCTION_BUFFER_FUNCTION_ACK",
			6,
			109,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FlexifunctionBufferFunctionAck)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_FLEXIFUNCTION_DIRECTORY: {
			"MSG_ID_FLEXIFUNCTION_DIRECTORY",
			53,
			12,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FlexifunctionDirectory)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_FLEXIFUNCTION_DIRECTORY_ACK: {
			"MSG_ID_FLEXIFUNCTION_DIRECTORY_ACK",
			7,
			218,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FlexifunctionDirectoryAck)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_FLEXIFUNCTION_COMMAND: {
			"MSG_ID_FLEXIFUNCTION_COMMAND",
			3,
			133,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FlexifunctionCommand)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_FLEXIFUNCTION_COMMAND_ACK: {
			"MSG_ID_FLEXIFUNCTION_COMMAND_ACK",
			4,
			208,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FlexifunctionCommandAck)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F2_A: {
			"MSG_ID_SERIAL_UDB_EXTRA_F2_A",
			61,
			103,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF2A)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F2_B: {
			"MSG_ID_SERIAL_UDB_EXTRA_F2_B",
			108,
			245,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF2B)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F4: {
			"MSG_ID_SERIAL_UDB_EXTRA_F4",
			10,
			191,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF4)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F5: {
			"MSG_ID_SERIAL_UDB_EXTRA_F5",
			16,
			54,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF5)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F6: {
			"MSG_ID_SERIAL_UDB_EXTRA_F6",
			20,
			54,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF6)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F7: {
			"MSG_ID_SERIAL_UDB_EXTRA_F7",
			24,
			171,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF7)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F8: {
			"MSG_ID_SERIAL_UDB_EXTRA_F8",
			28,
			142,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF8)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F13: {
			"MSG_ID_SERIAL_UDB_EXTRA_F13",
			14,
			249,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF13)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F14: {
			"MSG_ID_SERIAL_UDB_EXTRA_F14",
			17,
			123,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF14)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F15: {
			"MSG_ID_SERIAL_UDB_EXTRA_F15",
			60,
			7,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF15)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F16: {
			"MSG_ID_SERIAL_UDB_EXTRA_F16",
			110,
			222,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF16)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ALTITUDES: {
			"MSG_ID_ALTITUDES",
			28,
			55,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Altitudes)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_AIRSPEEDS: {
			"MSG_ID_AIRSPEEDS",
			16,
			154,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Airspeeds)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F17: {
			"MSG_ID_SERIAL_UDB_EXTRA_F17",
			12,
			175,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF17)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F18: {
			"MSG_ID_SERIAL_UDB_EXTRA_F18",
			20,
			41,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF18)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F19: {
			"MSG_ID_SERIAL_UDB_EXTRA_F19",
			8,
			87,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF19)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F20: {
			"MSG_ID_SERIAL_UDB_EXTRA_F20",
			25,
			144,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF20)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F21: {
			"MSG_ID_SERIAL_UDB_EXTRA_F21",
			12,
			134,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF21)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_UDB_EXTRA_F22: {
			"MSG_ID_SERIAL_UDB_EXTRA_F22",
			12,
			91,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialUdbExtraF22)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SYS_STATUS: {
			"MSG_ID_SYS_STATUS",
			31,
			124,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SysStatus)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SYSTEM_TIME: {
			"MSG_ID_SYSTEM_TIME",
			12,
			137,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SystemTime)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_PING: {
			"MSG_ID_PING",
			14,
			237,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Ping)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_CHANGE_OPERATOR_CONTROL: {
			"MSG_ID_CHANGE_OPERATOR_CONTROL",
			28,
			217,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ChangeOperatorControl)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_CHANGE_OPERATOR_CONTROL_ACK: {
			"MSG_ID_CHANGE_OPERATOR_CONTROL_ACK",
			3,
			104,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ChangeOperatorControlAck)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_AUTH_KEY: {
			"MSG_ID_AUTH_KEY",
			32,
			119,
			func(p packet.Packet) (message.Message, error) {
				msg := new(AuthKey)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LINK_NODE_STATUS: {
			"MSG_ID_LINK_NODE_STATUS",
			36,
			117,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LinkNodeStatus)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SET_MODE: {
			"MSG_ID_SET_MODE",
			6,
			89,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SetMode)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_PARAM_ACK_TRANSACTION: {
			"MSG_ID_PARAM_ACK_TRANSACTION",
			24,
			137,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ParamAckTransaction)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_PARAM_REQUEST_READ: {
			"MSG_ID_PARAM_REQUEST_READ",
			20,
			214,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ParamRequestRead)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_PARAM_REQUEST_LIST: {
			"MSG_ID_PARAM_REQUEST_LIST",
			2,
			159,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ParamRequestList)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_PARAM_VALUE: {
			"MSG_ID_PARAM_VALUE",
			25,
			220,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ParamValue)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_PARAM_SET: {
			"MSG_ID_PARAM_SET",
			23,
			168,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ParamSet)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GPS_RAW_INT: {
			"MSG_ID_GPS_RAW_INT",
			30,
			24,
			func(p packet.Packet) (message.Message, error) {
				msg := new(GpsRawInt)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GPS_STATUS: {
			"MSG_ID_GPS_STATUS",
			101,
			23,
			func(p packet.Packet) (message.Message, error) {
				msg := new(GpsStatus)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SCALED_IMU: {
			"MSG_ID_SCALED_IMU",
			22,
			170,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ScaledImu)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_RAW_IMU: {
			"MSG_ID_RAW_IMU",
			26,
			144,
			func(p packet.Packet) (message.Message, error) {
				msg := new(RawImu)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_RAW_PRESSURE: {
			"MSG_ID_RAW_PRESSURE",
			16,
			67,
			func(p packet.Packet) (message.Message, error) {
				msg := new(RawPressure)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SCALED_PRESSURE: {
			"MSG_ID_SCALED_PRESSURE",
			14,
			115,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ScaledPressure)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ATTITUDE: {
			"MSG_ID_ATTITUDE",
			28,
			39,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Attitude)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ATTITUDE_QUATERNION: {
			"MSG_ID_ATTITUDE_QUATERNION",
			32,
			246,
			func(p packet.Packet) (message.Message, error) {
				msg := new(AttitudeQuaternion)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LOCAL_POSITION_NED: {
			"MSG_ID_LOCAL_POSITION_NED",
			28,
			185,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LocalPositionNed)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GLOBAL_POSITION_INT: {
			"MSG_ID_GLOBAL_POSITION_INT",
			28,
			104,
			func(p packet.Packet) (message.Message, error) {
				msg := new(GlobalPositionInt)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_RC_CHANNELS_SCALED: {
			"MSG_ID_RC_CHANNELS_SCALED",
			22,
			237,
			func(p packet.Packet) (message.Message, error) {
				msg := new(RcChannelsScaled)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_RC_CHANNELS_RAW: {
			"MSG_ID_RC_CHANNELS_RAW",
			22,
			244,
			func(p packet.Packet) (message.Message, error) {
				msg := new(RcChannelsRaw)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERVO_OUTPUT_RAW: {
			"MSG_ID_SERVO_OUTPUT_RAW",
			21,
			222,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ServoOutputRaw)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_REQUEST_PARTIAL_LIST: {
			"MSG_ID_MISSION_REQUEST_PARTIAL_LIST",
			6,
			212,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionRequestPartialList)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_WRITE_PARTIAL_LIST: {
			"MSG_ID_MISSION_WRITE_PARTIAL_LIST",
			6,
			9,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionWritePartialList)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_ITEM: {
			"MSG_ID_MISSION_ITEM",
			37,
			254,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionItem)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_REQUEST: {
			"MSG_ID_MISSION_REQUEST",
			4,
			230,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionRequest)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_SET_CURRENT: {
			"MSG_ID_MISSION_SET_CURRENT",
			4,
			28,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionSetCurrent)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_CURRENT: {
			"MSG_ID_MISSION_CURRENT",
			2,
			28,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionCurrent)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_REQUEST_LIST: {
			"MSG_ID_MISSION_REQUEST_LIST",
			2,
			132,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionRequestList)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_COUNT: {
			"MSG_ID_MISSION_COUNT",
			4,
			221,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionCount)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_CLEAR_ALL: {
			"MSG_ID_MISSION_CLEAR_ALL",
			2,
			232,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionClearAll)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_ITEM_REACHED: {
			"MSG_ID_MISSION_ITEM_REACHED",
			2,
			11,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionItemReached)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_ACK: {
			"MSG_ID_MISSION_ACK",
			3,
			153,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionAck)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SET_GPS_GLOBAL_ORIGIN: {
			"MSG_ID_SET_GPS_GLOBAL_ORIGIN",
			13,
			41,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SetGpsGlobalOrigin)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GPS_GLOBAL_ORIGIN: {
			"MSG_ID_GPS_GLOBAL_ORIGIN",
			12,
			39,
			func(p packet.Packet) (message.Message, error) {
				msg := new(GpsGlobalOrigin)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_PARAM_MAP_RC: {
			"MSG_ID_PARAM_MAP_RC",
			37,
			78,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ParamMapRc)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_REQUEST_INT: {
			"MSG_ID_MISSION_REQUEST_INT",
			4,
			196,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionRequestInt)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_CHANGED: {
			"MSG_ID_MISSION_CHANGED",
			7,
			132,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionChanged)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SAFETY_SET_ALLOWED_AREA: {
			"MSG_ID_SAFETY_SET_ALLOWED_AREA",
			27,
			15,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SafetySetAllowedArea)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SAFETY_ALLOWED_AREA: {
			"MSG_ID_SAFETY_ALLOWED_AREA",
			25,
			3,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SafetyAllowedArea)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ATTITUDE_QUATERNION_COV: {
			"MSG_ID_ATTITUDE_QUATERNION_COV",
			72,
			167,
			func(p packet.Packet) (message.Message, error) {
				msg := new(AttitudeQuaternionCov)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_NAV_CONTROLLER_OUTPUT: {
			"MSG_ID_NAV_CONTROLLER_OUTPUT",
			26,
			183,
			func(p packet.Packet) (message.Message, error) {
				msg := new(NavControllerOutput)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GLOBAL_POSITION_INT_COV: {
			"MSG_ID_GLOBAL_POSITION_INT_COV",
			181,
			119,
			func(p packet.Packet) (message.Message, error) {
				msg := new(GlobalPositionIntCov)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LOCAL_POSITION_NED_COV: {
			"MSG_ID_LOCAL_POSITION_NED_COV",
			225,
			191,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LocalPositionNedCov)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_RC_CHANNELS: {
			"MSG_ID_RC_CHANNELS",
			42,
			118,
			func(p packet.Packet) (message.Message, error) {
				msg := new(RcChannels)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_REQUEST_DATA_STREAM: {
			"MSG_ID_REQUEST_DATA_STREAM",
			6,
			148,
			func(p packet.Packet) (message.Message, error) {
				msg := new(RequestDataStream)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_DATA_STREAM: {
			"MSG_ID_DATA_STREAM",
			4,
			21,
			func(p packet.Packet) (message.Message, error) {
				msg := new(DataStream)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MANUAL_CONTROL: {
			"MSG_ID_MANUAL_CONTROL",
			11,
			243,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ManualControl)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_RC_CHANNELS_OVERRIDE: {
			"MSG_ID_RC_CHANNELS_OVERRIDE",
			18,
			124,
			func(p packet.Packet) (message.Message, error) {
				msg := new(RcChannelsOverride)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MISSION_ITEM_INT: {
			"MSG_ID_MISSION_ITEM_INT",
			37,
			38,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MissionItemInt)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_VFR_HUD: {
			"MSG_ID_VFR_HUD",
			20,
			20,
			func(p packet.Packet) (message.Message, error) {
				msg := new(VfrHud)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_COMMAND_INT: {
			"MSG_ID_COMMAND_INT",
			35,
			158,
			func(p packet.Packet) (message.Message, error) {
				msg := new(CommandInt)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_COMMAND_LONG: {
			"MSG_ID_COMMAND_LONG",
			33,
			152,
			func(p packet.Packet) (message.Message, error) {
				msg := new(CommandLong)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_COMMAND_ACK: {
			"MSG_ID_COMMAND_ACK",
			3,
			143,
			func(p packet.Packet) (message.Message, error) {
				msg := new(CommandAck)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_COMMAND_CANCEL: {
			"MSG_ID_COMMAND_CANCEL",
			4,
			14,
			func(p packet.Packet) (message.Message, error) {
				msg := new(CommandCancel)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MANUAL_SETPOINT: {
			"MSG_ID_MANUAL_SETPOINT",
			22,
			106,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ManualSetpoint)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SET_ATTITUDE_TARGET: {
			"MSG_ID_SET_ATTITUDE_TARGET",
			39,
			49,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SetAttitudeTarget)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ATTITUDE_TARGET: {
			"MSG_ID_ATTITUDE_TARGET",
			37,
			22,
			func(p packet.Packet) (message.Message, error) {
				msg := new(AttitudeTarget)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SET_POSITION_TARGET_LOCAL_NED: {
			"MSG_ID_SET_POSITION_TARGET_LOCAL_NED",
			53,
			143,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SetPositionTargetLocalNed)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_POSITION_TARGET_LOCAL_NED: {
			"MSG_ID_POSITION_TARGET_LOCAL_NED",
			51,
			140,
			func(p packet.Packet) (message.Message, error) {
				msg := new(PositionTargetLocalNed)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SET_POSITION_TARGET_GLOBAL_INT: {
			"MSG_ID_SET_POSITION_TARGET_GLOBAL_INT",
			53,
			5,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SetPositionTargetGlobalInt)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_POSITION_TARGET_GLOBAL_INT: {
			"MSG_ID_POSITION_TARGET_GLOBAL_INT",
			51,
			150,
			func(p packet.Packet) (message.Message, error) {
				msg := new(PositionTargetGlobalInt)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET: {
			"MSG_ID_LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET",
			28,
			231,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LocalPositionNedSystemGlobalOffset)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIL_STATE: {
			"MSG_ID_HIL_STATE",
			56,
			183,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HilState)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIL_CONTROLS: {
			"MSG_ID_HIL_CONTROLS",
			42,
			63,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HilControls)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIL_RC_INPUTS_RAW: {
			"MSG_ID_HIL_RC_INPUTS_RAW",
			33,
			54,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HilRcInputsRaw)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIL_ACTUATOR_CONTROLS: {
			"MSG_ID_HIL_ACTUATOR_CONTROLS",
			81,
			47,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HilActuatorControls)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_OPTICAL_FLOW: {
			"MSG_ID_OPTICAL_FLOW",
			26,
			175,
			func(p packet.Packet) (message.Message, error) {
				msg := new(OpticalFlow)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GLOBAL_VISION_POSITION_ESTIMATE: {
			"MSG_ID_GLOBAL_VISION_POSITION_ESTIMATE",
			32,
			102,
			func(p packet.Packet) (message.Message, error) {
				msg := new(GlobalVisionPositionEstimate)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_VISION_POSITION_ESTIMATE: {
			"MSG_ID_VISION_POSITION_ESTIMATE",
			32,
			158,
			func(p packet.Packet) (message.Message, error) {
				msg := new(VisionPositionEstimate)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_VISION_SPEED_ESTIMATE: {
			"MSG_ID_VISION_SPEED_ESTIMATE",
			20,
			208,
			func(p packet.Packet) (message.Message, error) {
				msg := new(VisionSpeedEstimate)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_VICON_POSITION_ESTIMATE: {
			"MSG_ID_VICON_POSITION_ESTIMATE",
			32,
			56,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ViconPositionEstimate)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIGHRES_IMU: {
			"MSG_ID_HIGHRES_IMU",
			62,
			93,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HighresImu)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_OPTICAL_FLOW_RAD: {
			"MSG_ID_OPTICAL_FLOW_RAD",
			44,
			138,
			func(p packet.Packet) (message.Message, error) {
				msg := new(OpticalFlowRad)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIL_SENSOR: {
			"MSG_ID_HIL_SENSOR",
			64,
			108,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HilSensor)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SIM_STATE: {
			"MSG_ID_SIM_STATE",
			84,
			32,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SimState)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_RADIO_STATUS: {
			"MSG_ID_RADIO_STATUS",
			9,
			185,
			func(p packet.Packet) (message.Message, error) {
				msg := new(RadioStatus)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_FILE_TRANSFER_PROTOCOL: {
			"MSG_ID_FILE_TRANSFER_PROTOCOL",
			254,
			84,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FileTransferProtocol)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_TIMESYNC: {
			"MSG_ID_TIMESYNC",
			16,
			34,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Timesync)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_CAMERA_TRIGGER: {
			"MSG_ID_CAMERA_TRIGGER",
			12,
			174,
			func(p packet.Packet) (message.Message, error) {
				msg := new(CameraTrigger)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIL_GPS: {
			"MSG_ID_HIL_GPS",
			36,
			124,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HilGps)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIL_OPTICAL_FLOW: {
			"MSG_ID_HIL_OPTICAL_FLOW",
			44,
			237,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HilOpticalFlow)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIL_STATE_QUATERNION: {
			"MSG_ID_HIL_STATE_QUATERNION",
			64,
			4,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HilStateQuaternion)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SCALED_IMU2: {
			"MSG_ID_SCALED_IMU2",
			22,
			76,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ScaledImu2)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LOG_REQUEST_LIST: {
			"MSG_ID_LOG_REQUEST_LIST",
			6,
			128,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LogRequestList)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LOG_ENTRY: {
			"MSG_ID_LOG_ENTRY",
			14,
			56,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LogEntry)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LOG_REQUEST_DATA: {
			"MSG_ID_LOG_REQUEST_DATA",
			12,
			116,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LogRequestData)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LOG_DATA: {
			"MSG_ID_LOG_DATA",
			97,
			134,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LogData)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LOG_ERASE: {
			"MSG_ID_LOG_ERASE",
			2,
			237,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LogErase)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LOG_REQUEST_END: {
			"MSG_ID_LOG_REQUEST_END",
			2,
			203,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LogRequestEnd)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GPS_INJECT_DATA: {
			"MSG_ID_GPS_INJECT_DATA",
			113,
			250,
			func(p packet.Packet) (message.Message, error) {
				msg := new(GpsInjectData)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GPS2_RAW: {
			"MSG_ID_GPS2_RAW",
			35,
			87,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Gps2Raw)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_POWER_STATUS: {
			"MSG_ID_POWER_STATUS",
			6,
			203,
			func(p packet.Packet) (message.Message, error) {
				msg := new(PowerStatus)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SERIAL_CONTROL: {
			"MSG_ID_SERIAL_CONTROL",
			79,
			220,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SerialControl)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GPS_RTK: {
			"MSG_ID_GPS_RTK",
			35,
			25,
			func(p packet.Packet) (message.Message, error) {
				msg := new(GpsRtk)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GPS2_RTK: {
			"MSG_ID_GPS2_RTK",
			35,
			226,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Gps2Rtk)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SCALED_IMU3: {
			"MSG_ID_SCALED_IMU3",
			22,
			46,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ScaledImu3)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_DATA_TRANSMISSION_HANDSHAKE: {
			"MSG_ID_DATA_TRANSMISSION_HANDSHAKE",
			13,
			29,
			func(p packet.Packet) (message.Message, error) {
				msg := new(DataTransmissionHandshake)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ENCAPSULATED_DATA: {
			"MSG_ID_ENCAPSULATED_DATA",
			255,
			223,
			func(p packet.Packet) (message.Message, error) {
				msg := new(EncapsulatedData)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_DISTANCE_SENSOR: {
			"MSG_ID_DISTANCE_SENSOR",
			14,
			85,
			func(p packet.Packet) (message.Message, error) {
				msg := new(DistanceSensor)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_TERRAIN_REQUEST: {
			"MSG_ID_TERRAIN_REQUEST",
			18,
			6,
			func(p packet.Packet) (message.Message, error) {
				msg := new(TerrainRequest)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_TERRAIN_DATA: {
			"MSG_ID_TERRAIN_DATA",
			43,
			229,
			func(p packet.Packet) (message.Message, error) {
				msg := new(TerrainData)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_TERRAIN_CHECK: {
			"MSG_ID_TERRAIN_CHECK",
			8,
			203,
			func(p packet.Packet) (message.Message, error) {
				msg := new(TerrainCheck)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_TERRAIN_REPORT: {
			"MSG_ID_TERRAIN_REPORT",
			22,
			1,
			func(p packet.Packet) (message.Message, error) {
				msg := new(TerrainReport)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SCALED_PRESSURE2: {
			"MSG_ID_SCALED_PRESSURE2",
			14,
			195,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ScaledPressure2)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ATT_POS_MOCAP: {
			"MSG_ID_ATT_POS_MOCAP",
			36,
			109,
			func(p packet.Packet) (message.Message, error) {
				msg := new(AttPosMocap)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SET_ACTUATOR_CONTROL_TARGET: {
			"MSG_ID_SET_ACTUATOR_CONTROL_TARGET",
			43,
			168,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SetActuatorControlTarget)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ACTUATOR_CONTROL_TARGET: {
			"MSG_ID_ACTUATOR_CONTROL_TARGET",
			41,
			181,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ActuatorControlTarget)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ALTITUDE: {
			"MSG_ID_ALTITUDE",
			32,
			47,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Altitude)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_RESOURCE_REQUEST: {
			"MSG_ID_RESOURCE_REQUEST",
			243,
			72,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ResourceRequest)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SCALED_PRESSURE3: {
			"MSG_ID_SCALED_PRESSURE3",
			14,
			131,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ScaledPressure3)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_FOLLOW_TARGET: {
			"MSG_ID_FOLLOW_TARGET",
			93,
			127,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FollowTarget)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_CONTROL_SYSTEM_STATE: {
			"MSG_ID_CONTROL_SYSTEM_STATE",
			100,
			103,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ControlSystemState)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_BATTERY_STATUS: {
			"MSG_ID_BATTERY_STATUS",
			36,
			154,
			func(p packet.Packet) (message.Message, error) {
				msg := new(BatteryStatus)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_AUTOPILOT_VERSION: {
			"MSG_ID_AUTOPILOT_VERSION",
			60,
			178,
			func(p packet.Packet) (message.Message, error) {
				msg := new(AutopilotVersion)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_LANDING_TARGET: {
			"MSG_ID_LANDING_TARGET",
			30,
			200,
			func(p packet.Packet) (message.Message, error) {
				msg := new(LandingTarget)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_FENCE_STATUS: {
			"MSG_ID_FENCE_STATUS",
			8,
			189,
			func(p packet.Packet) (message.Message, error) {
				msg := new(FenceStatus)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MAG_CAL_REPORT: {
			"MSG_ID_MAG_CAL_REPORT",
			44,
			36,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MagCalReport)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_EFI_STATUS: {
			"MSG_ID_EFI_STATUS",
			65,
			208,
			func(p packet.Packet) (message.Message, error) {
				msg := new(EfiStatus)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ESTIMATOR_STATUS: {
			"MSG_ID_ESTIMATOR_STATUS",
			42,
			163,
			func(p packet.Packet) (message.Message, error) {
				msg := new(EstimatorStatus)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_WIND_COV: {
			"MSG_ID_WIND_COV",
			40,
			105,
			func(p packet.Packet) (message.Message, error) {
				msg := new(WindCov)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GPS_INPUT: {
			"MSG_ID_GPS_INPUT",
			63,
			151,
			func(p packet.Packet) (message.Message, error) {
				msg := new(GpsInput)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_GPS_RTCM_DATA: {
			"MSG_ID_GPS_RTCM_DATA",
			182,
			35,
			func(p packet.Packet) (message.Message, error) {
				msg := new(GpsRtcmData)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIGH_LATENCY: {
			"MSG_ID_HIGH_LATENCY",
			40,
			150,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HighLatency)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HIGH_LATENCY2: {
			"MSG_ID_HIGH_LATENCY2",
			42,
			179,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HighLatency2)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_VIBRATION: {
			"MSG_ID_VIBRATION",
			32,
			90,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Vibration)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HOME_POSITION: {
			"MSG_ID_HOME_POSITION",
			52,
			104,
			func(p packet.Packet) (message.Message, error) {
				msg := new(HomePosition)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_SET_HOME_POSITION: {
			"MSG_ID_SET_HOME_POSITION",
			53,
			85,
			func(p packet.Packet) (message.Message, error) {
				msg := new(SetHomePosition)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MESSAGE_INTERVAL: {
			"MSG_ID_MESSAGE_INTERVAL",
			6,
			95,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MessageInterval)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_EXTENDED_SYS_STATE: {
			"MSG_ID_EXTENDED_SYS_STATE",
			2,
			130,
			func(p packet.Packet) (message.Message, error) {
				msg := new(ExtendedSysState)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_ADSB_VEHICLE: {
			"MSG_ID_ADSB_VEHICLE",
			38,
			184,
			func(p packet.Packet) (message.Message, error) {
				msg := new(AdsbVehicle)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_COLLISION: {
			"MSG_ID_COLLISION",
			19,
			81,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Collision)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_V2_EXTENSION: {
			"MSG_ID_V2_EXTENSION",
			254,
			8,
			func(p packet.Packet) (message.Message, error) {
				msg := new(V2Extension)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_MEMORY_VECT: {
			"MSG_ID_MEMORY_VECT",
			36,
			204,
			func(p packet.Packet) (message.Message, error) {
				msg := new(MemoryVect)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_DEBUG_VECT: {
			"MSG_ID_DEBUG_VECT",
			30,
			49,
			func(p packet.Packet) (message.Message, error) {
				msg := new(DebugVect)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_NAMED_VALUE_FLOAT: {
			"MSG_ID_NAMED_VALUE_FLOAT",
			18,
			170,
			func(p packet.Packet) (message.Message, error) {
				msg := new(NamedValueFloat)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_NAMED_VALUE_INT: {
			"MSG_ID_NAMED_VALUE_INT",
			18,
			44,
			func(p packet.Packet) (message.Message, error) {
				msg := new(NamedValueInt)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_STATUSTEXT: {
			"MSG_ID_STATUSTEXT",
			51,
			83,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Statustext)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_DEBUG: {
			"MSG_ID_DEBUG",
			9,
			46,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Debug)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
		MSG_ID_HEARTBEAT: {
			"MSG_ID_HEARTBEAT",
			9,
			50,
			func(p packet.Packet) (message.Message, error) {
				msg := new(Heartbeat)
				if err := msg.Unmarshal(p.Payload()); err != nil {
					return nil, err
				}
				return msg, nil
			},
		},
	} {
		register.Register(msgID, info.Name, info.Size, info.Extra, info.Constructor)
	}
}
