/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package asluav

import (
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
)

// Message IDs
const (
	MSG_ID_COMMAND_INT_STAMPED                     message.MessageID = 78
	MSG_ID_COMMAND_LONG_STAMPED                    message.MessageID = 79
	MSG_ID_SENS_POWER                              message.MessageID = 201
	MSG_ID_SENS_MPPT                               message.MessageID = 202
	MSG_ID_ASLCTRL_DATA                            message.MessageID = 203
	MSG_ID_ASLCTRL_DEBUG                           message.MessageID = 204
	MSG_ID_ASLUAV_STATUS                           message.MessageID = 205
	MSG_ID_EKF_EXT                                 message.MessageID = 206
	MSG_ID_ASL_OBCTRL                              message.MessageID = 207
	MSG_ID_SENS_ATMOS                              message.MessageID = 208
	MSG_ID_SENS_BATMON                             message.MessageID = 209
	MSG_ID_FW_SOARING_DATA                         message.MessageID = 210
	MSG_ID_SENSORPOD_STATUS                        message.MessageID = 211
	MSG_ID_SENS_POWER_BOARD                        message.MessageID = 212
	MSG_ID_GSM_LINK_STATUS                         message.MessageID = 213
	MSG_ID_SATCOM_LINK_STATUS                      message.MessageID = 214
	MSG_ID_SENSOR_AIRFLOW_ANGLES                   message.MessageID = 215
	MSG_ID_SYS_STATUS                              message.MessageID = 1
	MSG_ID_SYSTEM_TIME                             message.MessageID = 2
	MSG_ID_PING                                    message.MessageID = 4
	MSG_ID_CHANGE_OPERATOR_CONTROL                 message.MessageID = 5
	MSG_ID_CHANGE_OPERATOR_CONTROL_ACK             message.MessageID = 6
	MSG_ID_AUTH_KEY                                message.MessageID = 7
	MSG_ID_LINK_NODE_STATUS                        message.MessageID = 8
	MSG_ID_SET_MODE                                message.MessageID = 11
	MSG_ID_PARAM_ACK_TRANSACTION                   message.MessageID = 19
	MSG_ID_PARAM_REQUEST_READ                      message.MessageID = 20
	MSG_ID_PARAM_REQUEST_LIST                      message.MessageID = 21
	MSG_ID_PARAM_VALUE                             message.MessageID = 22
	MSG_ID_PARAM_SET                               message.MessageID = 23
	MSG_ID_GPS_RAW_INT                             message.MessageID = 24
	MSG_ID_GPS_STATUS                              message.MessageID = 25
	MSG_ID_SCALED_IMU                              message.MessageID = 26
	MSG_ID_RAW_IMU                                 message.MessageID = 27
	MSG_ID_RAW_PRESSURE                            message.MessageID = 28
	MSG_ID_SCALED_PRESSURE                         message.MessageID = 29
	MSG_ID_ATTITUDE                                message.MessageID = 30
	MSG_ID_ATTITUDE_QUATERNION                     message.MessageID = 31
	MSG_ID_LOCAL_POSITION_NED                      message.MessageID = 32
	MSG_ID_GLOBAL_POSITION_INT                     message.MessageID = 33
	MSG_ID_RC_CHANNELS_SCALED                      message.MessageID = 34
	MSG_ID_RC_CHANNELS_RAW                         message.MessageID = 35
	MSG_ID_SERVO_OUTPUT_RAW                        message.MessageID = 36
	MSG_ID_MISSION_REQUEST_PARTIAL_LIST            message.MessageID = 37
	MSG_ID_MISSION_WRITE_PARTIAL_LIST              message.MessageID = 38
	MSG_ID_MISSION_ITEM                            message.MessageID = 39
	MSG_ID_MISSION_REQUEST                         message.MessageID = 40
	MSG_ID_MISSION_SET_CURRENT                     message.MessageID = 41
	MSG_ID_MISSION_CURRENT                         message.MessageID = 42
	MSG_ID_MISSION_REQUEST_LIST                    message.MessageID = 43
	MSG_ID_MISSION_COUNT                           message.MessageID = 44
	MSG_ID_MISSION_CLEAR_ALL                       message.MessageID = 45
	MSG_ID_MISSION_ITEM_REACHED                    message.MessageID = 46
	MSG_ID_MISSION_ACK                             message.MessageID = 47
	MSG_ID_SET_GPS_GLOBAL_ORIGIN                   message.MessageID = 48
	MSG_ID_GPS_GLOBAL_ORIGIN                       message.MessageID = 49
	MSG_ID_PARAM_MAP_RC                            message.MessageID = 50
	MSG_ID_MISSION_REQUEST_INT                     message.MessageID = 51
	MSG_ID_MISSION_CHANGED                         message.MessageID = 52
	MSG_ID_SAFETY_SET_ALLOWED_AREA                 message.MessageID = 54
	MSG_ID_SAFETY_ALLOWED_AREA                     message.MessageID = 55
	MSG_ID_ATTITUDE_QUATERNION_COV                 message.MessageID = 61
	MSG_ID_NAV_CONTROLLER_OUTPUT                   message.MessageID = 62
	MSG_ID_GLOBAL_POSITION_INT_COV                 message.MessageID = 63
	MSG_ID_LOCAL_POSITION_NED_COV                  message.MessageID = 64
	MSG_ID_RC_CHANNELS                             message.MessageID = 65
	MSG_ID_REQUEST_DATA_STREAM                     message.MessageID = 66
	MSG_ID_DATA_STREAM                             message.MessageID = 67
	MSG_ID_MANUAL_CONTROL                          message.MessageID = 69
	MSG_ID_RC_CHANNELS_OVERRIDE                    message.MessageID = 70
	MSG_ID_MISSION_ITEM_INT                        message.MessageID = 73
	MSG_ID_VFR_HUD                                 message.MessageID = 74
	MSG_ID_COMMAND_INT                             message.MessageID = 75
	MSG_ID_COMMAND_LONG                            message.MessageID = 76
	MSG_ID_COMMAND_ACK                             message.MessageID = 77
	MSG_ID_COMMAND_CANCEL                          message.MessageID = 80
	MSG_ID_MANUAL_SETPOINT                         message.MessageID = 81
	MSG_ID_SET_ATTITUDE_TARGET                     message.MessageID = 82
	MSG_ID_ATTITUDE_TARGET                         message.MessageID = 83
	MSG_ID_SET_POSITION_TARGET_LOCAL_NED           message.MessageID = 84
	MSG_ID_POSITION_TARGET_LOCAL_NED               message.MessageID = 85
	MSG_ID_SET_POSITION_TARGET_GLOBAL_INT          message.MessageID = 86
	MSG_ID_POSITION_TARGET_GLOBAL_INT              message.MessageID = 87
	MSG_ID_LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET message.MessageID = 89
	MSG_ID_HIL_STATE                               message.MessageID = 90
	MSG_ID_HIL_CONTROLS                            message.MessageID = 91
	MSG_ID_HIL_RC_INPUTS_RAW                       message.MessageID = 92
	MSG_ID_HIL_ACTUATOR_CONTROLS                   message.MessageID = 93
	MSG_ID_OPTICAL_FLOW                            message.MessageID = 100
	MSG_ID_GLOBAL_VISION_POSITION_ESTIMATE         message.MessageID = 101
	MSG_ID_VISION_POSITION_ESTIMATE                message.MessageID = 102
	MSG_ID_VISION_SPEED_ESTIMATE                   message.MessageID = 103
	MSG_ID_VICON_POSITION_ESTIMATE                 message.MessageID = 104
	MSG_ID_HIGHRES_IMU                             message.MessageID = 105
	MSG_ID_OPTICAL_FLOW_RAD                        message.MessageID = 106
	MSG_ID_HIL_SENSOR                              message.MessageID = 107
	MSG_ID_SIM_STATE                               message.MessageID = 108
	MSG_ID_RADIO_STATUS                            message.MessageID = 109
	MSG_ID_FILE_TRANSFER_PROTOCOL                  message.MessageID = 110
	MSG_ID_TIMESYNC                                message.MessageID = 111
	MSG_ID_CAMERA_TRIGGER                          message.MessageID = 112
	MSG_ID_HIL_GPS                                 message.MessageID = 113
	MSG_ID_HIL_OPTICAL_FLOW                        message.MessageID = 114
	MSG_ID_HIL_STATE_QUATERNION                    message.MessageID = 115
	MSG_ID_SCALED_IMU2                             message.MessageID = 116
	MSG_ID_LOG_REQUEST_LIST                        message.MessageID = 117
	MSG_ID_LOG_ENTRY                               message.MessageID = 118
	MSG_ID_LOG_REQUEST_DATA                        message.MessageID = 119
	MSG_ID_LOG_DATA                                message.MessageID = 120
	MSG_ID_LOG_ERASE                               message.MessageID = 121
	MSG_ID_LOG_REQUEST_END                         message.MessageID = 122
	MSG_ID_GPS_INJECT_DATA                         message.MessageID = 123
	MSG_ID_GPS2_RAW                                message.MessageID = 124
	MSG_ID_POWER_STATUS                            message.MessageID = 125
	MSG_ID_SERIAL_CONTROL                          message.MessageID = 126
	MSG_ID_GPS_RTK                                 message.MessageID = 127
	MSG_ID_GPS2_RTK                                message.MessageID = 128
	MSG_ID_SCALED_IMU3                             message.MessageID = 129
	MSG_ID_DATA_TRANSMISSION_HANDSHAKE             message.MessageID = 130
	MSG_ID_ENCAPSULATED_DATA                       message.MessageID = 131
	MSG_ID_DISTANCE_SENSOR                         message.MessageID = 132
	MSG_ID_TERRAIN_REQUEST                         message.MessageID = 133
	MSG_ID_TERRAIN_DATA                            message.MessageID = 134
	MSG_ID_TERRAIN_CHECK                           message.MessageID = 135
	MSG_ID_TERRAIN_REPORT                          message.MessageID = 136
	MSG_ID_SCALED_PRESSURE2                        message.MessageID = 137
	MSG_ID_ATT_POS_MOCAP                           message.MessageID = 138
	MSG_ID_SET_ACTUATOR_CONTROL_TARGET             message.MessageID = 139
	MSG_ID_ACTUATOR_CONTROL_TARGET                 message.MessageID = 140
	MSG_ID_ALTITUDE                                message.MessageID = 141
	MSG_ID_RESOURCE_REQUEST                        message.MessageID = 142
	MSG_ID_SCALED_PRESSURE3                        message.MessageID = 143
	MSG_ID_FOLLOW_TARGET                           message.MessageID = 144
	MSG_ID_CONTROL_SYSTEM_STATE                    message.MessageID = 146
	MSG_ID_BATTERY_STATUS                          message.MessageID = 147
	MSG_ID_AUTOPILOT_VERSION                       message.MessageID = 148
	MSG_ID_LANDING_TARGET                          message.MessageID = 149
	MSG_ID_FENCE_STATUS                            message.MessageID = 162
	MSG_ID_MAG_CAL_REPORT                          message.MessageID = 192
	MSG_ID_EFI_STATUS                              message.MessageID = 225
	MSG_ID_ESTIMATOR_STATUS                        message.MessageID = 230
	MSG_ID_WIND_COV                                message.MessageID = 231
	MSG_ID_GPS_INPUT                               message.MessageID = 232
	MSG_ID_GPS_RTCM_DATA                           message.MessageID = 233
	MSG_ID_HIGH_LATENCY                            message.MessageID = 234
	MSG_ID_HIGH_LATENCY2                           message.MessageID = 235
	MSG_ID_VIBRATION                               message.MessageID = 241
	MSG_ID_HOME_POSITION                           message.MessageID = 242
	MSG_ID_SET_HOME_POSITION                       message.MessageID = 243
	MSG_ID_MESSAGE_INTERVAL                        message.MessageID = 244
	MSG_ID_EXTENDED_SYS_STATE                      message.MessageID = 245
	MSG_ID_ADSB_VEHICLE                            message.MessageID = 246
	MSG_ID_COLLISION                               message.MessageID = 247
	MSG_ID_V2_EXTENSION                            message.MessageID = 248
	MSG_ID_MEMORY_VECT                             message.MessageID = 249
	MSG_ID_DEBUG_VECT                              message.MessageID = 250
	MSG_ID_NAMED_VALUE_FLOAT                       message.MessageID = 251
	MSG_ID_NAMED_VALUE_INT                         message.MessageID = 252
	MSG_ID_STATUSTEXT                              message.MessageID = 253
	MSG_ID_DEBUG                                   message.MessageID = 254
	MSG_ID_HEARTBEAT                               message.MessageID = 0
)
