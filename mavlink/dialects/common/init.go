/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package common

import (
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/asmyasnikov/go-mavlink/mavlink/register"
)

func init() {
	for msgID, info := range map[message.MessageID]register.MessageInfo{
		MSG_ID_SYS_STATUS: {
			"MSG_ID_SYS_STATUS",
			31,
			124,
			func(bytes []byte) (message.Message, error) {
				msg := new(SysStatus)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SYSTEM_TIME: {
			"MSG_ID_SYSTEM_TIME",
			12,
			137,
			func(bytes []byte) (message.Message, error) {
				msg := new(SystemTime)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_PING: {
			"MSG_ID_PING",
			14,
			237,
			func(bytes []byte) (message.Message, error) {
				msg := new(Ping)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_CHANGE_OPERATOR_CONTROL: {
			"MSG_ID_CHANGE_OPERATOR_CONTROL",
			28,
			217,
			func(bytes []byte) (message.Message, error) {
				msg := new(ChangeOperatorControl)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_CHANGE_OPERATOR_CONTROL_ACK: {
			"MSG_ID_CHANGE_OPERATOR_CONTROL_ACK",
			3,
			104,
			func(bytes []byte) (message.Message, error) {
				msg := new(ChangeOperatorControlAck)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_AUTH_KEY: {
			"MSG_ID_AUTH_KEY",
			32,
			119,
			func(bytes []byte) (message.Message, error) {
				msg := new(AuthKey)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LINK_NODE_STATUS: {
			"MSG_ID_LINK_NODE_STATUS",
			36,
			117,
			func(bytes []byte) (message.Message, error) {
				msg := new(LinkNodeStatus)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SET_MODE: {
			"MSG_ID_SET_MODE",
			6,
			89,
			func(bytes []byte) (message.Message, error) {
				msg := new(SetMode)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_PARAM_ACK_TRANSACTION: {
			"MSG_ID_PARAM_ACK_TRANSACTION",
			24,
			137,
			func(bytes []byte) (message.Message, error) {
				msg := new(ParamAckTransaction)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_PARAM_REQUEST_READ: {
			"MSG_ID_PARAM_REQUEST_READ",
			20,
			214,
			func(bytes []byte) (message.Message, error) {
				msg := new(ParamRequestRead)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_PARAM_REQUEST_LIST: {
			"MSG_ID_PARAM_REQUEST_LIST",
			2,
			159,
			func(bytes []byte) (message.Message, error) {
				msg := new(ParamRequestList)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_PARAM_VALUE: {
			"MSG_ID_PARAM_VALUE",
			25,
			220,
			func(bytes []byte) (message.Message, error) {
				msg := new(ParamValue)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_PARAM_SET: {
			"MSG_ID_PARAM_SET",
			23,
			168,
			func(bytes []byte) (message.Message, error) {
				msg := new(ParamSet)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GPS_RAW_INT: {
			"MSG_ID_GPS_RAW_INT",
			30,
			24,
			func(bytes []byte) (message.Message, error) {
				msg := new(GpsRawInt)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GPS_STATUS: {
			"MSG_ID_GPS_STATUS",
			101,
			23,
			func(bytes []byte) (message.Message, error) {
				msg := new(GpsStatus)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SCALED_IMU: {
			"MSG_ID_SCALED_IMU",
			22,
			170,
			func(bytes []byte) (message.Message, error) {
				msg := new(ScaledImu)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_RAW_IMU: {
			"MSG_ID_RAW_IMU",
			26,
			144,
			func(bytes []byte) (message.Message, error) {
				msg := new(RawImu)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_RAW_PRESSURE: {
			"MSG_ID_RAW_PRESSURE",
			16,
			67,
			func(bytes []byte) (message.Message, error) {
				msg := new(RawPressure)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SCALED_PRESSURE: {
			"MSG_ID_SCALED_PRESSURE",
			14,
			115,
			func(bytes []byte) (message.Message, error) {
				msg := new(ScaledPressure)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_ATTITUDE: {
			"MSG_ID_ATTITUDE",
			28,
			39,
			func(bytes []byte) (message.Message, error) {
				msg := new(Attitude)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_ATTITUDE_QUATERNION: {
			"MSG_ID_ATTITUDE_QUATERNION",
			32,
			246,
			func(bytes []byte) (message.Message, error) {
				msg := new(AttitudeQuaternion)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LOCAL_POSITION_NED: {
			"MSG_ID_LOCAL_POSITION_NED",
			28,
			185,
			func(bytes []byte) (message.Message, error) {
				msg := new(LocalPositionNed)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GLOBAL_POSITION_INT: {
			"MSG_ID_GLOBAL_POSITION_INT",
			28,
			104,
			func(bytes []byte) (message.Message, error) {
				msg := new(GlobalPositionInt)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_RC_CHANNELS_SCALED: {
			"MSG_ID_RC_CHANNELS_SCALED",
			22,
			237,
			func(bytes []byte) (message.Message, error) {
				msg := new(RcChannelsScaled)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_RC_CHANNELS_RAW: {
			"MSG_ID_RC_CHANNELS_RAW",
			22,
			244,
			func(bytes []byte) (message.Message, error) {
				msg := new(RcChannelsRaw)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SERVO_OUTPUT_RAW: {
			"MSG_ID_SERVO_OUTPUT_RAW",
			21,
			222,
			func(bytes []byte) (message.Message, error) {
				msg := new(ServoOutputRaw)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_REQUEST_PARTIAL_LIST: {
			"MSG_ID_MISSION_REQUEST_PARTIAL_LIST",
			6,
			212,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionRequestPartialList)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_WRITE_PARTIAL_LIST: {
			"MSG_ID_MISSION_WRITE_PARTIAL_LIST",
			6,
			9,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionWritePartialList)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_ITEM: {
			"MSG_ID_MISSION_ITEM",
			37,
			254,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionItem)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_REQUEST: {
			"MSG_ID_MISSION_REQUEST",
			4,
			230,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionRequest)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_SET_CURRENT: {
			"MSG_ID_MISSION_SET_CURRENT",
			4,
			28,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionSetCurrent)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_CURRENT: {
			"MSG_ID_MISSION_CURRENT",
			2,
			28,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionCurrent)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_REQUEST_LIST: {
			"MSG_ID_MISSION_REQUEST_LIST",
			2,
			132,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionRequestList)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_COUNT: {
			"MSG_ID_MISSION_COUNT",
			4,
			221,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionCount)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_CLEAR_ALL: {
			"MSG_ID_MISSION_CLEAR_ALL",
			2,
			232,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionClearAll)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_ITEM_REACHED: {
			"MSG_ID_MISSION_ITEM_REACHED",
			2,
			11,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionItemReached)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_ACK: {
			"MSG_ID_MISSION_ACK",
			3,
			153,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionAck)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SET_GPS_GLOBAL_ORIGIN: {
			"MSG_ID_SET_GPS_GLOBAL_ORIGIN",
			13,
			41,
			func(bytes []byte) (message.Message, error) {
				msg := new(SetGpsGlobalOrigin)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GPS_GLOBAL_ORIGIN: {
			"MSG_ID_GPS_GLOBAL_ORIGIN",
			12,
			39,
			func(bytes []byte) (message.Message, error) {
				msg := new(GpsGlobalOrigin)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_PARAM_MAP_RC: {
			"MSG_ID_PARAM_MAP_RC",
			37,
			78,
			func(bytes []byte) (message.Message, error) {
				msg := new(ParamMapRc)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_REQUEST_INT: {
			"MSG_ID_MISSION_REQUEST_INT",
			4,
			196,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionRequestInt)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_CHANGED: {
			"MSG_ID_MISSION_CHANGED",
			7,
			132,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionChanged)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SAFETY_SET_ALLOWED_AREA: {
			"MSG_ID_SAFETY_SET_ALLOWED_AREA",
			27,
			15,
			func(bytes []byte) (message.Message, error) {
				msg := new(SafetySetAllowedArea)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SAFETY_ALLOWED_AREA: {
			"MSG_ID_SAFETY_ALLOWED_AREA",
			25,
			3,
			func(bytes []byte) (message.Message, error) {
				msg := new(SafetyAllowedArea)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_ATTITUDE_QUATERNION_COV: {
			"MSG_ID_ATTITUDE_QUATERNION_COV",
			72,
			167,
			func(bytes []byte) (message.Message, error) {
				msg := new(AttitudeQuaternionCov)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_NAV_CONTROLLER_OUTPUT: {
			"MSG_ID_NAV_CONTROLLER_OUTPUT",
			26,
			183,
			func(bytes []byte) (message.Message, error) {
				msg := new(NavControllerOutput)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GLOBAL_POSITION_INT_COV: {
			"MSG_ID_GLOBAL_POSITION_INT_COV",
			181,
			119,
			func(bytes []byte) (message.Message, error) {
				msg := new(GlobalPositionIntCov)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LOCAL_POSITION_NED_COV: {
			"MSG_ID_LOCAL_POSITION_NED_COV",
			225,
			191,
			func(bytes []byte) (message.Message, error) {
				msg := new(LocalPositionNedCov)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_RC_CHANNELS: {
			"MSG_ID_RC_CHANNELS",
			42,
			118,
			func(bytes []byte) (message.Message, error) {
				msg := new(RcChannels)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_REQUEST_DATA_STREAM: {
			"MSG_ID_REQUEST_DATA_STREAM",
			6,
			148,
			func(bytes []byte) (message.Message, error) {
				msg := new(RequestDataStream)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_DATA_STREAM: {
			"MSG_ID_DATA_STREAM",
			4,
			21,
			func(bytes []byte) (message.Message, error) {
				msg := new(DataStream)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MANUAL_CONTROL: {
			"MSG_ID_MANUAL_CONTROL",
			11,
			243,
			func(bytes []byte) (message.Message, error) {
				msg := new(ManualControl)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_RC_CHANNELS_OVERRIDE: {
			"MSG_ID_RC_CHANNELS_OVERRIDE",
			18,
			124,
			func(bytes []byte) (message.Message, error) {
				msg := new(RcChannelsOverride)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MISSION_ITEM_INT: {
			"MSG_ID_MISSION_ITEM_INT",
			37,
			38,
			func(bytes []byte) (message.Message, error) {
				msg := new(MissionItemInt)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_VFR_HUD: {
			"MSG_ID_VFR_HUD",
			20,
			20,
			func(bytes []byte) (message.Message, error) {
				msg := new(VfrHud)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_COMMAND_INT: {
			"MSG_ID_COMMAND_INT",
			35,
			158,
			func(bytes []byte) (message.Message, error) {
				msg := new(CommandInt)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_COMMAND_LONG: {
			"MSG_ID_COMMAND_LONG",
			33,
			152,
			func(bytes []byte) (message.Message, error) {
				msg := new(CommandLong)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_COMMAND_ACK: {
			"MSG_ID_COMMAND_ACK",
			3,
			143,
			func(bytes []byte) (message.Message, error) {
				msg := new(CommandAck)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_COMMAND_CANCEL: {
			"MSG_ID_COMMAND_CANCEL",
			4,
			14,
			func(bytes []byte) (message.Message, error) {
				msg := new(CommandCancel)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MANUAL_SETPOINT: {
			"MSG_ID_MANUAL_SETPOINT",
			22,
			106,
			func(bytes []byte) (message.Message, error) {
				msg := new(ManualSetpoint)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SET_ATTITUDE_TARGET: {
			"MSG_ID_SET_ATTITUDE_TARGET",
			39,
			49,
			func(bytes []byte) (message.Message, error) {
				msg := new(SetAttitudeTarget)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_ATTITUDE_TARGET: {
			"MSG_ID_ATTITUDE_TARGET",
			37,
			22,
			func(bytes []byte) (message.Message, error) {
				msg := new(AttitudeTarget)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SET_POSITION_TARGET_LOCAL_NED: {
			"MSG_ID_SET_POSITION_TARGET_LOCAL_NED",
			53,
			143,
			func(bytes []byte) (message.Message, error) {
				msg := new(SetPositionTargetLocalNed)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_POSITION_TARGET_LOCAL_NED: {
			"MSG_ID_POSITION_TARGET_LOCAL_NED",
			51,
			140,
			func(bytes []byte) (message.Message, error) {
				msg := new(PositionTargetLocalNed)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SET_POSITION_TARGET_GLOBAL_INT: {
			"MSG_ID_SET_POSITION_TARGET_GLOBAL_INT",
			53,
			5,
			func(bytes []byte) (message.Message, error) {
				msg := new(SetPositionTargetGlobalInt)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_POSITION_TARGET_GLOBAL_INT: {
			"MSG_ID_POSITION_TARGET_GLOBAL_INT",
			51,
			150,
			func(bytes []byte) (message.Message, error) {
				msg := new(PositionTargetGlobalInt)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET: {
			"MSG_ID_LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET",
			28,
			231,
			func(bytes []byte) (message.Message, error) {
				msg := new(LocalPositionNedSystemGlobalOffset)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIL_STATE: {
			"MSG_ID_HIL_STATE",
			56,
			183,
			func(bytes []byte) (message.Message, error) {
				msg := new(HilState)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIL_CONTROLS: {
			"MSG_ID_HIL_CONTROLS",
			42,
			63,
			func(bytes []byte) (message.Message, error) {
				msg := new(HilControls)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIL_RC_INPUTS_RAW: {
			"MSG_ID_HIL_RC_INPUTS_RAW",
			33,
			54,
			func(bytes []byte) (message.Message, error) {
				msg := new(HilRcInputsRaw)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIL_ACTUATOR_CONTROLS: {
			"MSG_ID_HIL_ACTUATOR_CONTROLS",
			81,
			47,
			func(bytes []byte) (message.Message, error) {
				msg := new(HilActuatorControls)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_OPTICAL_FLOW: {
			"MSG_ID_OPTICAL_FLOW",
			26,
			175,
			func(bytes []byte) (message.Message, error) {
				msg := new(OpticalFlow)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GLOBAL_VISION_POSITION_ESTIMATE: {
			"MSG_ID_GLOBAL_VISION_POSITION_ESTIMATE",
			32,
			102,
			func(bytes []byte) (message.Message, error) {
				msg := new(GlobalVisionPositionEstimate)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_VISION_POSITION_ESTIMATE: {
			"MSG_ID_VISION_POSITION_ESTIMATE",
			32,
			158,
			func(bytes []byte) (message.Message, error) {
				msg := new(VisionPositionEstimate)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_VISION_SPEED_ESTIMATE: {
			"MSG_ID_VISION_SPEED_ESTIMATE",
			20,
			208,
			func(bytes []byte) (message.Message, error) {
				msg := new(VisionSpeedEstimate)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_VICON_POSITION_ESTIMATE: {
			"MSG_ID_VICON_POSITION_ESTIMATE",
			32,
			56,
			func(bytes []byte) (message.Message, error) {
				msg := new(ViconPositionEstimate)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIGHRES_IMU: {
			"MSG_ID_HIGHRES_IMU",
			62,
			93,
			func(bytes []byte) (message.Message, error) {
				msg := new(HighresImu)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_OPTICAL_FLOW_RAD: {
			"MSG_ID_OPTICAL_FLOW_RAD",
			44,
			138,
			func(bytes []byte) (message.Message, error) {
				msg := new(OpticalFlowRad)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIL_SENSOR: {
			"MSG_ID_HIL_SENSOR",
			64,
			108,
			func(bytes []byte) (message.Message, error) {
				msg := new(HilSensor)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SIM_STATE: {
			"MSG_ID_SIM_STATE",
			84,
			32,
			func(bytes []byte) (message.Message, error) {
				msg := new(SimState)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_RADIO_STATUS: {
			"MSG_ID_RADIO_STATUS",
			9,
			185,
			func(bytes []byte) (message.Message, error) {
				msg := new(RadioStatus)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_FILE_TRANSFER_PROTOCOL: {
			"MSG_ID_FILE_TRANSFER_PROTOCOL",
			254,
			84,
			func(bytes []byte) (message.Message, error) {
				msg := new(FileTransferProtocol)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_TIMESYNC: {
			"MSG_ID_TIMESYNC",
			16,
			34,
			func(bytes []byte) (message.Message, error) {
				msg := new(Timesync)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_CAMERA_TRIGGER: {
			"MSG_ID_CAMERA_TRIGGER",
			12,
			174,
			func(bytes []byte) (message.Message, error) {
				msg := new(CameraTrigger)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIL_GPS: {
			"MSG_ID_HIL_GPS",
			36,
			124,
			func(bytes []byte) (message.Message, error) {
				msg := new(HilGps)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIL_OPTICAL_FLOW: {
			"MSG_ID_HIL_OPTICAL_FLOW",
			44,
			237,
			func(bytes []byte) (message.Message, error) {
				msg := new(HilOpticalFlow)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIL_STATE_QUATERNION: {
			"MSG_ID_HIL_STATE_QUATERNION",
			64,
			4,
			func(bytes []byte) (message.Message, error) {
				msg := new(HilStateQuaternion)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SCALED_IMU2: {
			"MSG_ID_SCALED_IMU2",
			22,
			76,
			func(bytes []byte) (message.Message, error) {
				msg := new(ScaledImu2)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LOG_REQUEST_LIST: {
			"MSG_ID_LOG_REQUEST_LIST",
			6,
			128,
			func(bytes []byte) (message.Message, error) {
				msg := new(LogRequestList)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LOG_ENTRY: {
			"MSG_ID_LOG_ENTRY",
			14,
			56,
			func(bytes []byte) (message.Message, error) {
				msg := new(LogEntry)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LOG_REQUEST_DATA: {
			"MSG_ID_LOG_REQUEST_DATA",
			12,
			116,
			func(bytes []byte) (message.Message, error) {
				msg := new(LogRequestData)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LOG_DATA: {
			"MSG_ID_LOG_DATA",
			97,
			134,
			func(bytes []byte) (message.Message, error) {
				msg := new(LogData)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LOG_ERASE: {
			"MSG_ID_LOG_ERASE",
			2,
			237,
			func(bytes []byte) (message.Message, error) {
				msg := new(LogErase)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LOG_REQUEST_END: {
			"MSG_ID_LOG_REQUEST_END",
			2,
			203,
			func(bytes []byte) (message.Message, error) {
				msg := new(LogRequestEnd)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GPS_INJECT_DATA: {
			"MSG_ID_GPS_INJECT_DATA",
			113,
			250,
			func(bytes []byte) (message.Message, error) {
				msg := new(GpsInjectData)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GPS2_RAW: {
			"MSG_ID_GPS2_RAW",
			35,
			87,
			func(bytes []byte) (message.Message, error) {
				msg := new(Gps2Raw)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_POWER_STATUS: {
			"MSG_ID_POWER_STATUS",
			6,
			203,
			func(bytes []byte) (message.Message, error) {
				msg := new(PowerStatus)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SERIAL_CONTROL: {
			"MSG_ID_SERIAL_CONTROL",
			79,
			220,
			func(bytes []byte) (message.Message, error) {
				msg := new(SerialControl)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GPS_RTK: {
			"MSG_ID_GPS_RTK",
			35,
			25,
			func(bytes []byte) (message.Message, error) {
				msg := new(GpsRtk)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GPS2_RTK: {
			"MSG_ID_GPS2_RTK",
			35,
			226,
			func(bytes []byte) (message.Message, error) {
				msg := new(Gps2Rtk)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SCALED_IMU3: {
			"MSG_ID_SCALED_IMU3",
			22,
			46,
			func(bytes []byte) (message.Message, error) {
				msg := new(ScaledImu3)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_DATA_TRANSMISSION_HANDSHAKE: {
			"MSG_ID_DATA_TRANSMISSION_HANDSHAKE",
			13,
			29,
			func(bytes []byte) (message.Message, error) {
				msg := new(DataTransmissionHandshake)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_ENCAPSULATED_DATA: {
			"MSG_ID_ENCAPSULATED_DATA",
			255,
			223,
			func(bytes []byte) (message.Message, error) {
				msg := new(EncapsulatedData)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_DISTANCE_SENSOR: {
			"MSG_ID_DISTANCE_SENSOR",
			14,
			85,
			func(bytes []byte) (message.Message, error) {
				msg := new(DistanceSensor)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_TERRAIN_REQUEST: {
			"MSG_ID_TERRAIN_REQUEST",
			18,
			6,
			func(bytes []byte) (message.Message, error) {
				msg := new(TerrainRequest)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_TERRAIN_DATA: {
			"MSG_ID_TERRAIN_DATA",
			43,
			229,
			func(bytes []byte) (message.Message, error) {
				msg := new(TerrainData)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_TERRAIN_CHECK: {
			"MSG_ID_TERRAIN_CHECK",
			8,
			203,
			func(bytes []byte) (message.Message, error) {
				msg := new(TerrainCheck)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_TERRAIN_REPORT: {
			"MSG_ID_TERRAIN_REPORT",
			22,
			1,
			func(bytes []byte) (message.Message, error) {
				msg := new(TerrainReport)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SCALED_PRESSURE2: {
			"MSG_ID_SCALED_PRESSURE2",
			14,
			195,
			func(bytes []byte) (message.Message, error) {
				msg := new(ScaledPressure2)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_ATT_POS_MOCAP: {
			"MSG_ID_ATT_POS_MOCAP",
			36,
			109,
			func(bytes []byte) (message.Message, error) {
				msg := new(AttPosMocap)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SET_ACTUATOR_CONTROL_TARGET: {
			"MSG_ID_SET_ACTUATOR_CONTROL_TARGET",
			43,
			168,
			func(bytes []byte) (message.Message, error) {
				msg := new(SetActuatorControlTarget)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_ACTUATOR_CONTROL_TARGET: {
			"MSG_ID_ACTUATOR_CONTROL_TARGET",
			41,
			181,
			func(bytes []byte) (message.Message, error) {
				msg := new(ActuatorControlTarget)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_ALTITUDE: {
			"MSG_ID_ALTITUDE",
			32,
			47,
			func(bytes []byte) (message.Message, error) {
				msg := new(Altitude)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_RESOURCE_REQUEST: {
			"MSG_ID_RESOURCE_REQUEST",
			243,
			72,
			func(bytes []byte) (message.Message, error) {
				msg := new(ResourceRequest)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SCALED_PRESSURE3: {
			"MSG_ID_SCALED_PRESSURE3",
			14,
			131,
			func(bytes []byte) (message.Message, error) {
				msg := new(ScaledPressure3)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_FOLLOW_TARGET: {
			"MSG_ID_FOLLOW_TARGET",
			93,
			127,
			func(bytes []byte) (message.Message, error) {
				msg := new(FollowTarget)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_CONTROL_SYSTEM_STATE: {
			"MSG_ID_CONTROL_SYSTEM_STATE",
			100,
			103,
			func(bytes []byte) (message.Message, error) {
				msg := new(ControlSystemState)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_BATTERY_STATUS: {
			"MSG_ID_BATTERY_STATUS",
			36,
			154,
			func(bytes []byte) (message.Message, error) {
				msg := new(BatteryStatus)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_AUTOPILOT_VERSION: {
			"MSG_ID_AUTOPILOT_VERSION",
			60,
			178,
			func(bytes []byte) (message.Message, error) {
				msg := new(AutopilotVersion)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_LANDING_TARGET: {
			"MSG_ID_LANDING_TARGET",
			30,
			200,
			func(bytes []byte) (message.Message, error) {
				msg := new(LandingTarget)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_FENCE_STATUS: {
			"MSG_ID_FENCE_STATUS",
			8,
			189,
			func(bytes []byte) (message.Message, error) {
				msg := new(FenceStatus)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MAG_CAL_REPORT: {
			"MSG_ID_MAG_CAL_REPORT",
			44,
			36,
			func(bytes []byte) (message.Message, error) {
				msg := new(MagCalReport)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_EFI_STATUS: {
			"MSG_ID_EFI_STATUS",
			65,
			208,
			func(bytes []byte) (message.Message, error) {
				msg := new(EfiStatus)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_ESTIMATOR_STATUS: {
			"MSG_ID_ESTIMATOR_STATUS",
			42,
			163,
			func(bytes []byte) (message.Message, error) {
				msg := new(EstimatorStatus)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_WIND_COV: {
			"MSG_ID_WIND_COV",
			40,
			105,
			func(bytes []byte) (message.Message, error) {
				msg := new(WindCov)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GPS_INPUT: {
			"MSG_ID_GPS_INPUT",
			63,
			151,
			func(bytes []byte) (message.Message, error) {
				msg := new(GpsInput)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_GPS_RTCM_DATA: {
			"MSG_ID_GPS_RTCM_DATA",
			182,
			35,
			func(bytes []byte) (message.Message, error) {
				msg := new(GpsRtcmData)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIGH_LATENCY: {
			"MSG_ID_HIGH_LATENCY",
			40,
			150,
			func(bytes []byte) (message.Message, error) {
				msg := new(HighLatency)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HIGH_LATENCY2: {
			"MSG_ID_HIGH_LATENCY2",
			42,
			179,
			func(bytes []byte) (message.Message, error) {
				msg := new(HighLatency2)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_VIBRATION: {
			"MSG_ID_VIBRATION",
			32,
			90,
			func(bytes []byte) (message.Message, error) {
				msg := new(Vibration)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HOME_POSITION: {
			"MSG_ID_HOME_POSITION",
			52,
			104,
			func(bytes []byte) (message.Message, error) {
				msg := new(HomePosition)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_SET_HOME_POSITION: {
			"MSG_ID_SET_HOME_POSITION",
			53,
			85,
			func(bytes []byte) (message.Message, error) {
				msg := new(SetHomePosition)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MESSAGE_INTERVAL: {
			"MSG_ID_MESSAGE_INTERVAL",
			6,
			95,
			func(bytes []byte) (message.Message, error) {
				msg := new(MessageInterval)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_EXTENDED_SYS_STATE: {
			"MSG_ID_EXTENDED_SYS_STATE",
			2,
			130,
			func(bytes []byte) (message.Message, error) {
				msg := new(ExtendedSysState)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_ADSB_VEHICLE: {
			"MSG_ID_ADSB_VEHICLE",
			38,
			184,
			func(bytes []byte) (message.Message, error) {
				msg := new(AdsbVehicle)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_COLLISION: {
			"MSG_ID_COLLISION",
			19,
			81,
			func(bytes []byte) (message.Message, error) {
				msg := new(Collision)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_V2_EXTENSION: {
			"MSG_ID_V2_EXTENSION",
			254,
			8,
			func(bytes []byte) (message.Message, error) {
				msg := new(V2Extension)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_MEMORY_VECT: {
			"MSG_ID_MEMORY_VECT",
			36,
			204,
			func(bytes []byte) (message.Message, error) {
				msg := new(MemoryVect)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_DEBUG_VECT: {
			"MSG_ID_DEBUG_VECT",
			30,
			49,
			func(bytes []byte) (message.Message, error) {
				msg := new(DebugVect)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_NAMED_VALUE_FLOAT: {
			"MSG_ID_NAMED_VALUE_FLOAT",
			18,
			170,
			func(bytes []byte) (message.Message, error) {
				msg := new(NamedValueFloat)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_NAMED_VALUE_INT: {
			"MSG_ID_NAMED_VALUE_INT",
			18,
			44,
			func(bytes []byte) (message.Message, error) {
				msg := new(NamedValueInt)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_STATUSTEXT: {
			"MSG_ID_STATUSTEXT",
			51,
			83,
			func(bytes []byte) (message.Message, error) {
				msg := new(Statustext)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_DEBUG: {
			"MSG_ID_DEBUG",
			9,
			46,
			func(bytes []byte) (message.Message, error) {
				msg := new(Debug)
				return msg, msg.Unmarshal(bytes)
			},
		},
		MSG_ID_HEARTBEAT: {
			"MSG_ID_HEARTBEAT",
			9,
			50,
			func(bytes []byte) (message.Message, error) {
				msg := new(Heartbeat)
				return msg, msg.Unmarshal(bytes)
			},
		},
	} {
		register.Register(msgID, info.Name, info.Size, info.Extra, info.Constructor)
	}
}
