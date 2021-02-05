/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package paparazzi

import (
	"encoding/binary"
	"fmt"
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"math"
	"strings"
)

// ScriptItem struct (generated typeinfo)
// Message encoding a mission script item. This message is emitted upon a request for the next script item.
type ScriptItem struct {
	Seq             uint16 // Sequence
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
	Name            string `len:"50" ` // The name of the mission script, NULL terminated.
}

// MsgID (generated function)
func (m *ScriptItem) MsgID() message.MessageID {
	return MSG_ID_SCRIPT_ITEM
}

// String (generated function)
func (m *ScriptItem) String() string {
	return fmt.Sprintf(
		"&paparazzi.ScriptItem{ Seq: %+v, TargetSystem: %+v, TargetComponent: %+v, Name: \"%s\" }",
		m.Seq,
		m.TargetSystem,
		m.TargetComponent,
		m.Name,
	)
}

// Marshal (generated function)
func (m *ScriptItem) Marshal() ([]byte, error) {
	payload := make([]byte, 54)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Seq))
	payload[2] = byte(m.TargetSystem)
	payload[3] = byte(m.TargetComponent)
	copy(payload[4:], m.Name)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScriptItem) Unmarshal(data []byte) error {
	payload := make([]byte, 54)
	copy(payload[0:], data)
	m.Seq = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetSystem = uint8(payload[2])
	m.TargetComponent = uint8(payload[3])
	m.Name = strings.TrimRight(string(payload[4:54]), string(byte(0)))
	return nil
}

// ScriptRequest struct (generated typeinfo)
// Request script item with the sequence number seq. The response of the system to this message should be a SCRIPT_ITEM message.
type ScriptRequest struct {
	Seq             uint16 // Sequence
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
}

// MsgID (generated function)
func (m *ScriptRequest) MsgID() message.MessageID {
	return MSG_ID_SCRIPT_REQUEST
}

// String (generated function)
func (m *ScriptRequest) String() string {
	return fmt.Sprintf(
		"&paparazzi.ScriptRequest{ Seq: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.Seq,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *ScriptRequest) Marshal() ([]byte, error) {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Seq))
	payload[2] = byte(m.TargetSystem)
	payload[3] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScriptRequest) Unmarshal(data []byte) error {
	payload := make([]byte, 4)
	copy(payload[0:], data)
	m.Seq = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetSystem = uint8(payload[2])
	m.TargetComponent = uint8(payload[3])
	return nil
}

// ScriptRequestList struct (generated typeinfo)
// Request the overall list of mission items from the system/component.
type ScriptRequestList struct {
	TargetSystem    uint8 // System ID
	TargetComponent uint8 // Component ID
}

// MsgID (generated function)
func (m *ScriptRequestList) MsgID() message.MessageID {
	return MSG_ID_SCRIPT_REQUEST_LIST
}

// String (generated function)
func (m *ScriptRequestList) String() string {
	return fmt.Sprintf(
		"&paparazzi.ScriptRequestList{ TargetSystem: %+v, TargetComponent: %+v }",
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *ScriptRequestList) Marshal() ([]byte, error) {
	payload := make([]byte, 2)
	payload[0] = byte(m.TargetSystem)
	payload[1] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScriptRequestList) Unmarshal(data []byte) error {
	payload := make([]byte, 2)
	copy(payload[0:], data)
	m.TargetSystem = uint8(payload[0])
	m.TargetComponent = uint8(payload[1])
	return nil
}

// ScriptCount struct (generated typeinfo)
// This message is emitted as response to SCRIPT_REQUEST_LIST by the MAV to get the number of mission scripts.
type ScriptCount struct {
	Count           uint16 // Number of script items in the sequence
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
}

// MsgID (generated function)
func (m *ScriptCount) MsgID() message.MessageID {
	return MSG_ID_SCRIPT_COUNT
}

// String (generated function)
func (m *ScriptCount) String() string {
	return fmt.Sprintf(
		"&paparazzi.ScriptCount{ Count: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.Count,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *ScriptCount) Marshal() ([]byte, error) {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Count))
	payload[2] = byte(m.TargetSystem)
	payload[3] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScriptCount) Unmarshal(data []byte) error {
	payload := make([]byte, 4)
	copy(payload[0:], data)
	m.Count = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetSystem = uint8(payload[2])
	m.TargetComponent = uint8(payload[3])
	return nil
}

// ScriptCurrent struct (generated typeinfo)
// This message informs about the currently active SCRIPT.
type ScriptCurrent struct {
	Seq uint16 // Active Sequence
}

// MsgID (generated function)
func (m *ScriptCurrent) MsgID() message.MessageID {
	return MSG_ID_SCRIPT_CURRENT
}

// String (generated function)
func (m *ScriptCurrent) String() string {
	return fmt.Sprintf(
		"&paparazzi.ScriptCurrent{ Seq: %+v }",
		m.Seq,
	)
}

// Marshal (generated function)
func (m *ScriptCurrent) Marshal() ([]byte, error) {
	payload := make([]byte, 2)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Seq))
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScriptCurrent) Unmarshal(data []byte) error {
	payload := make([]byte, 2)
	copy(payload[0:], data)
	m.Seq = uint16(binary.LittleEndian.Uint16(payload[0:]))
	return nil
}

// SysStatus struct (generated typeinfo)
// The general system state. If the system is following the MAVLink standard, the system state is mainly defined by three orthogonal states/modes: The system mode, which is either LOCKED (motors shut down and locked), MANUAL (system under RC control), GUIDED (system with autonomous position control, position setpoint controlled manually) or AUTO (system guided by path/waypoint planner). The NAV_MODE defined the current flight state: LIFTOFF (often an open-loop maneuver), LANDING, WAYPOINTS or VECTOR. This represents the internal navigation state machine. The system status shows whether the system is currently active or not and if an emergency occurred. During the CRITICAL and EMERGENCY states the MAV is still considered to be active, but should start emergency procedures autonomously. After a failure occurred it should first move from active to critical to allow manual intervention and then move to emergency after a certain timeout.
type SysStatus struct {
	OnboardControlSensorsPresent MAV_SYS_STATUS_SENSOR // Bitmap showing which onboard controllers and sensors are present. Value of 0: not present. Value of 1: present.
	OnboardControlSensorsEnabled MAV_SYS_STATUS_SENSOR // Bitmap showing which onboard controllers and sensors are enabled:  Value of 0: not enabled. Value of 1: enabled.
	OnboardControlSensorsHealth  MAV_SYS_STATUS_SENSOR // Bitmap showing which onboard controllers and sensors have an error (or are operational). Value of 0: error. Value of 1: healthy.
	Load                         uint16                // Maximum usage in percent of the mainloop time. Values: [0-1000] - should always be below 1000
	VoltageBattery               uint16                // Battery voltage, UINT16_MAX: Voltage not sent by autopilot
	CurrentBattery               int16                 // Battery current, -1: Current not sent by autopilot
	DropRateComm                 uint16                // Communication drop rate, (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV)
	ErrorsComm                   uint16                // Communication errors (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV)
	ErrorsCount1                 uint16                // Autopilot-specific errors
	ErrorsCount2                 uint16                // Autopilot-specific errors
	ErrorsCount3                 uint16                // Autopilot-specific errors
	ErrorsCount4                 uint16                // Autopilot-specific errors
	BatteryRemaining             int8                  // Battery energy remaining, -1: Battery remaining energy not sent by autopilot
}

// MsgID (generated function)
func (m *SysStatus) MsgID() message.MessageID {
	return MSG_ID_SYS_STATUS
}

// String (generated function)
func (m *SysStatus) String() string {
	return fmt.Sprintf(
		"&common.SysStatus{ OnboardControlSensorsPresent: %+v (%032b), OnboardControlSensorsEnabled: %+v (%032b), OnboardControlSensorsHealth: %+v (%032b), Load: %+v, VoltageBattery: %+v, CurrentBattery: %+v, DropRateComm: %+v, ErrorsComm: %+v, ErrorsCount1: %+v, ErrorsCount2: %+v, ErrorsCount3: %+v, ErrorsCount4: %+v, BatteryRemaining: %+v }",
		m.OnboardControlSensorsPresent.Bitmask(), uint64(m.OnboardControlSensorsPresent),
		m.OnboardControlSensorsEnabled.Bitmask(), uint64(m.OnboardControlSensorsEnabled),
		m.OnboardControlSensorsHealth.Bitmask(), uint64(m.OnboardControlSensorsHealth),
		m.Load,
		m.VoltageBattery,
		m.CurrentBattery,
		m.DropRateComm,
		m.ErrorsComm,
		m.ErrorsCount1,
		m.ErrorsCount2,
		m.ErrorsCount3,
		m.ErrorsCount4,
		m.BatteryRemaining,
	)
}

// Marshal (generated function)
func (m *SysStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 31)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.OnboardControlSensorsPresent))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.OnboardControlSensorsEnabled))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.OnboardControlSensorsHealth))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Load))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.VoltageBattery))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.CurrentBattery))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.DropRateComm))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.ErrorsComm))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.ErrorsCount1))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.ErrorsCount2))
	binary.LittleEndian.PutUint16(payload[26:], uint16(m.ErrorsCount3))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.ErrorsCount4))
	payload[30] = byte(m.BatteryRemaining)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SysStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 31)
	copy(payload[0:], data)
	m.OnboardControlSensorsPresent = MAV_SYS_STATUS_SENSOR(binary.LittleEndian.Uint32(payload[0:]))
	m.OnboardControlSensorsEnabled = MAV_SYS_STATUS_SENSOR(binary.LittleEndian.Uint32(payload[4:]))
	m.OnboardControlSensorsHealth = MAV_SYS_STATUS_SENSOR(binary.LittleEndian.Uint32(payload[8:]))
	m.Load = uint16(binary.LittleEndian.Uint16(payload[12:]))
	m.VoltageBattery = uint16(binary.LittleEndian.Uint16(payload[14:]))
	m.CurrentBattery = int16(binary.LittleEndian.Uint16(payload[16:]))
	m.DropRateComm = uint16(binary.LittleEndian.Uint16(payload[18:]))
	m.ErrorsComm = uint16(binary.LittleEndian.Uint16(payload[20:]))
	m.ErrorsCount1 = uint16(binary.LittleEndian.Uint16(payload[22:]))
	m.ErrorsCount2 = uint16(binary.LittleEndian.Uint16(payload[24:]))
	m.ErrorsCount3 = uint16(binary.LittleEndian.Uint16(payload[26:]))
	m.ErrorsCount4 = uint16(binary.LittleEndian.Uint16(payload[28:]))
	m.BatteryRemaining = int8(payload[30])
	return nil
}

// SystemTime struct (generated typeinfo)
// The system time is the time of the master clock, typically the computer clock of the main onboard computer.
type SystemTime struct {
	TimeUnixUsec uint64 // Timestamp (UNIX epoch time).
	TimeBootMs   uint32 // Timestamp (time since system boot).
}

// MsgID (generated function)
func (m *SystemTime) MsgID() message.MessageID {
	return MSG_ID_SYSTEM_TIME
}

// String (generated function)
func (m *SystemTime) String() string {
	return fmt.Sprintf(
		"&common.SystemTime{ TimeUnixUsec: %+v, TimeBootMs: %+v }",
		m.TimeUnixUsec,
		m.TimeBootMs,
	)
}

// Marshal (generated function)
func (m *SystemTime) Marshal() ([]byte, error) {
	payload := make([]byte, 12)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUnixUsec))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.TimeBootMs))
	return payload, nil
}

// Unmarshal (generated function)
func (m *SystemTime) Unmarshal(data []byte) error {
	payload := make([]byte, 12)
	copy(payload[0:], data)
	m.TimeUnixUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[8:]))
	return nil
}

// Ping struct (generated typeinfo)
// A ping message either requesting or responding to a ping. This allows to measure the system latencies, including serial port, radio modem and UDP connections. The ping microservice is documented at https://mavlink.io/en/services/ping.html
type Ping struct {
	TimeUsec        uint64 // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Seq             uint32 // PING sequence
	TargetSystem    uint8  // 0: request ping from all receiving systems. If greater than 0: message is a ping response and number is the system id of the requesting system
	TargetComponent uint8  // 0: request ping from all receiving components. If greater than 0: message is a ping response and number is the component id of the requesting component.
}

// MsgID (generated function)
func (m *Ping) MsgID() message.MessageID {
	return MSG_ID_PING
}

// String (generated function)
func (m *Ping) String() string {
	return fmt.Sprintf(
		"&common.Ping{ TimeUsec: %+v, Seq: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.TimeUsec,
		m.Seq,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *Ping) Marshal() ([]byte, error) {
	payload := make([]byte, 14)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Seq))
	payload[12] = byte(m.TargetSystem)
	payload[13] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *Ping) Unmarshal(data []byte) error {
	payload := make([]byte, 14)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Seq = uint32(binary.LittleEndian.Uint32(payload[8:]))
	m.TargetSystem = uint8(payload[12])
	m.TargetComponent = uint8(payload[13])
	return nil
}

// ChangeOperatorControl struct (generated typeinfo)
// Request to control this MAV
type ChangeOperatorControl struct {
	TargetSystem   uint8  // System the GCS requests control for
	ControlRequest uint8  // 0: request control of this MAV, 1: Release control of this MAV
	Version        uint8  // 0: key as plaintext, 1-255: future, different hashing/encryption variants. The GCS should in general use the safest mode possible initially and then gradually move down the encryption level if it gets a NACK message indicating an encryption mismatch.
	Passkey        string `len:"25" ` // Password / Key, depending on version plaintext or encrypted. 25 or less characters, NULL terminated. The characters may involve A-Z, a-z, 0-9, and "!?,.-"
}

// MsgID (generated function)
func (m *ChangeOperatorControl) MsgID() message.MessageID {
	return MSG_ID_CHANGE_OPERATOR_CONTROL
}

// String (generated function)
func (m *ChangeOperatorControl) String() string {
	return fmt.Sprintf(
		"&common.ChangeOperatorControl{ TargetSystem: %+v, ControlRequest: %+v, Version: %+v, Passkey: \"%s\" }",
		m.TargetSystem,
		m.ControlRequest,
		m.Version,
		m.Passkey,
	)
}

// Marshal (generated function)
func (m *ChangeOperatorControl) Marshal() ([]byte, error) {
	payload := make([]byte, 28)
	payload[0] = byte(m.TargetSystem)
	payload[1] = byte(m.ControlRequest)
	payload[2] = byte(m.Version)
	copy(payload[3:], m.Passkey)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ChangeOperatorControl) Unmarshal(data []byte) error {
	payload := make([]byte, 28)
	copy(payload[0:], data)
	m.TargetSystem = uint8(payload[0])
	m.ControlRequest = uint8(payload[1])
	m.Version = uint8(payload[2])
	m.Passkey = strings.TrimRight(string(payload[3:28]), string(byte(0)))
	return nil
}

// ChangeOperatorControlAck struct (generated typeinfo)
// Accept / deny control of this MAV
type ChangeOperatorControlAck struct {
	GcsSystemID    uint8 // ID of the GCS this message
	ControlRequest uint8 // 0: request control of this MAV, 1: Release control of this MAV
	Ack            uint8 // 0: ACK, 1: NACK: Wrong passkey, 2: NACK: Unsupported passkey encryption method, 3: NACK: Already under control
}

// MsgID (generated function)
func (m *ChangeOperatorControlAck) MsgID() message.MessageID {
	return MSG_ID_CHANGE_OPERATOR_CONTROL_ACK
}

// String (generated function)
func (m *ChangeOperatorControlAck) String() string {
	return fmt.Sprintf(
		"&common.ChangeOperatorControlAck{ GcsSystemID: %+v, ControlRequest: %+v, Ack: %+v }",
		m.GcsSystemID,
		m.ControlRequest,
		m.Ack,
	)
}

// Marshal (generated function)
func (m *ChangeOperatorControlAck) Marshal() ([]byte, error) {
	payload := make([]byte, 3)
	payload[0] = byte(m.GcsSystemID)
	payload[1] = byte(m.ControlRequest)
	payload[2] = byte(m.Ack)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ChangeOperatorControlAck) Unmarshal(data []byte) error {
	payload := make([]byte, 3)
	copy(payload[0:], data)
	m.GcsSystemID = uint8(payload[0])
	m.ControlRequest = uint8(payload[1])
	m.Ack = uint8(payload[2])
	return nil
}

// AuthKey struct (generated typeinfo)
// Emit an encrypted signature / key identifying this system. PLEASE NOTE: This protocol has been kept simple, so transmitting the key requires an encrypted channel for true safety.
type AuthKey struct {
	Key string `len:"32" ` // key
}

// MsgID (generated function)
func (m *AuthKey) MsgID() message.MessageID {
	return MSG_ID_AUTH_KEY
}

// String (generated function)
func (m *AuthKey) String() string {
	return fmt.Sprintf(
		"&common.AuthKey{ Key: \"%s\" }",
		m.Key,
	)
}

// Marshal (generated function)
func (m *AuthKey) Marshal() ([]byte, error) {
	payload := make([]byte, 32)
	copy(payload[0:], m.Key)
	return payload, nil
}

// Unmarshal (generated function)
func (m *AuthKey) Unmarshal(data []byte) error {
	payload := make([]byte, 32)
	copy(payload[0:], data)
	m.Key = strings.TrimRight(string(payload[0:32]), string(byte(0)))
	return nil
}

// LinkNodeStatus struct (generated typeinfo)
// Status generated in each node in the communication chain and injected into MAVLink stream.
type LinkNodeStatus struct {
	Timestamp        uint64 // Timestamp (time since system boot).
	TxRate           uint32 // Transmit rate
	RxRate           uint32 // Receive rate
	MessagesSent     uint32 // Messages sent
	MessagesReceived uint32 // Messages received (estimated from counting seq)
	MessagesLost     uint32 // Messages lost (estimated from counting seq)
	RxParseErr       uint16 // Number of bytes that could not be parsed correctly.
	TxOverflows      uint16 // Transmit buffer overflows. This number wraps around as it reaches UINT16_MAX
	RxOverflows      uint16 // Receive buffer overflows. This number wraps around as it reaches UINT16_MAX
	TxBuf            uint8  // Remaining free transmit buffer space
	RxBuf            uint8  // Remaining free receive buffer space
}

// MsgID (generated function)
func (m *LinkNodeStatus) MsgID() message.MessageID {
	return MSG_ID_LINK_NODE_STATUS
}

// String (generated function)
func (m *LinkNodeStatus) String() string {
	return fmt.Sprintf(
		"&common.LinkNodeStatus{ Timestamp: %+v, TxRate: %+v, RxRate: %+v, MessagesSent: %+v, MessagesReceived: %+v, MessagesLost: %+v, RxParseErr: %+v, TxOverflows: %+v, RxOverflows: %+v, TxBuf: %+v, RxBuf: %+v }",
		m.Timestamp,
		m.TxRate,
		m.RxRate,
		m.MessagesSent,
		m.MessagesReceived,
		m.MessagesLost,
		m.RxParseErr,
		m.TxOverflows,
		m.RxOverflows,
		m.TxBuf,
		m.RxBuf,
	)
}

// Marshal (generated function)
func (m *LinkNodeStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 36)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.TxRate))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.RxRate))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.MessagesSent))
	binary.LittleEndian.PutUint32(payload[20:], uint32(m.MessagesReceived))
	binary.LittleEndian.PutUint32(payload[24:], uint32(m.MessagesLost))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.RxParseErr))
	binary.LittleEndian.PutUint16(payload[30:], uint16(m.TxOverflows))
	binary.LittleEndian.PutUint16(payload[32:], uint16(m.RxOverflows))
	payload[34] = byte(m.TxBuf)
	payload[35] = byte(m.RxBuf)
	return payload, nil
}

// Unmarshal (generated function)
func (m *LinkNodeStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 36)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.TxRate = uint32(binary.LittleEndian.Uint32(payload[8:]))
	m.RxRate = uint32(binary.LittleEndian.Uint32(payload[12:]))
	m.MessagesSent = uint32(binary.LittleEndian.Uint32(payload[16:]))
	m.MessagesReceived = uint32(binary.LittleEndian.Uint32(payload[20:]))
	m.MessagesLost = uint32(binary.LittleEndian.Uint32(payload[24:]))
	m.RxParseErr = uint16(binary.LittleEndian.Uint16(payload[28:]))
	m.TxOverflows = uint16(binary.LittleEndian.Uint16(payload[30:]))
	m.RxOverflows = uint16(binary.LittleEndian.Uint16(payload[32:]))
	m.TxBuf = uint8(payload[34])
	m.RxBuf = uint8(payload[35])
	return nil
}

// SetMode struct (generated typeinfo)
// Set the system mode, as defined by enum MAV_MODE. There is no target component id as the mode is by definition for the overall aircraft, not only for one component.
type SetMode struct {
	CustomMode   uint32   // The new autopilot-specific mode. This field can be ignored by an autopilot.
	TargetSystem uint8    // The system setting the mode
	BaseMode     MAV_MODE // The new base mode.
}

// MsgID (generated function)
func (m *SetMode) MsgID() message.MessageID {
	return MSG_ID_SET_MODE
}

// String (generated function)
func (m *SetMode) String() string {
	return fmt.Sprintf(
		"&common.SetMode{ CustomMode: %+v, TargetSystem: %+v, BaseMode: %+v }",
		m.CustomMode,
		m.TargetSystem,
		m.BaseMode,
	)
}

// Marshal (generated function)
func (m *SetMode) Marshal() ([]byte, error) {
	payload := make([]byte, 6)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.CustomMode))
	payload[4] = byte(m.TargetSystem)
	payload[5] = byte(m.BaseMode)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SetMode) Unmarshal(data []byte) error {
	payload := make([]byte, 6)
	copy(payload[0:], data)
	m.CustomMode = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.TargetSystem = uint8(payload[4])
	m.BaseMode = MAV_MODE(payload[5])
	return nil
}

// ParamAckTransaction struct (generated typeinfo)
// Response from a PARAM_SET message when it is used in a transaction.
type ParamAckTransaction struct {
	ParamValue      float32        // Parameter value (new value if PARAM_ACCEPTED, current value otherwise)
	TargetSystem    uint8          // Id of system that sent PARAM_SET message.
	TargetComponent uint8          // Id of system that sent PARAM_SET message.
	ParamID         string         `len:"16" ` // Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamType       MAV_PARAM_TYPE // Parameter type.
	ParamResult     PARAM_ACK      // Result code.
}

// MsgID (generated function)
func (m *ParamAckTransaction) MsgID() message.MessageID {
	return MSG_ID_PARAM_ACK_TRANSACTION
}

// String (generated function)
func (m *ParamAckTransaction) String() string {
	return fmt.Sprintf(
		"&common.ParamAckTransaction{ ParamValue: %+v, TargetSystem: %+v, TargetComponent: %+v, ParamID: \"%s\", ParamType: %+v, ParamResult: %+v }",
		m.ParamValue,
		m.TargetSystem,
		m.TargetComponent,
		m.ParamID,
		m.ParamType,
		m.ParamResult,
	)
}

// Marshal (generated function)
func (m *ParamAckTransaction) Marshal() ([]byte, error) {
	payload := make([]byte, 24)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.ParamValue))
	payload[4] = byte(m.TargetSystem)
	payload[5] = byte(m.TargetComponent)
	copy(payload[6:], m.ParamID)
	payload[22] = byte(m.ParamType)
	payload[23] = byte(m.ParamResult)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ParamAckTransaction) Unmarshal(data []byte) error {
	payload := make([]byte, 24)
	copy(payload[0:], data)
	m.ParamValue = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.TargetSystem = uint8(payload[4])
	m.TargetComponent = uint8(payload[5])
	m.ParamID = strings.TrimRight(string(payload[6:22]), string(byte(0)))
	m.ParamType = MAV_PARAM_TYPE(payload[22])
	m.ParamResult = PARAM_ACK(payload[23])
	return nil
}

// ParamRequestRead struct (generated typeinfo)
// Request to read the onboard parameter with the param_id string id. Onboard parameters are stored as key[const char*] -> value[float]. This allows to send a parameter to any other component (such as the GCS) without the need of previous knowledge of possible parameter names. Thus the same GCS can store different parameters for different autopilots. See also https://mavlink.io/en/services/parameter.html for a full documentation of QGroundControl and IMU code.
type ParamRequestRead struct {
	ParamIndex      int16  // Parameter index. Send -1 to use the param ID field as identifier (else the param id will be ignored)
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
	ParamID         string `len:"16" ` // Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
}

// MsgID (generated function)
func (m *ParamRequestRead) MsgID() message.MessageID {
	return MSG_ID_PARAM_REQUEST_READ
}

// String (generated function)
func (m *ParamRequestRead) String() string {
	return fmt.Sprintf(
		"&common.ParamRequestRead{ ParamIndex: %+v, TargetSystem: %+v, TargetComponent: %+v, ParamID: \"%s\" }",
		m.ParamIndex,
		m.TargetSystem,
		m.TargetComponent,
		m.ParamID,
	)
}

// Marshal (generated function)
func (m *ParamRequestRead) Marshal() ([]byte, error) {
	payload := make([]byte, 20)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.ParamIndex))
	payload[2] = byte(m.TargetSystem)
	payload[3] = byte(m.TargetComponent)
	copy(payload[4:], m.ParamID)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ParamRequestRead) Unmarshal(data []byte) error {
	payload := make([]byte, 20)
	copy(payload[0:], data)
	m.ParamIndex = int16(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetSystem = uint8(payload[2])
	m.TargetComponent = uint8(payload[3])
	m.ParamID = strings.TrimRight(string(payload[4:20]), string(byte(0)))
	return nil
}

// ParamRequestList struct (generated typeinfo)
// Request all parameters of this component. After this request, all parameters are emitted. The parameter microservice is documented at https://mavlink.io/en/services/parameter.html
type ParamRequestList struct {
	TargetSystem    uint8 // System ID
	TargetComponent uint8 // Component ID
}

// MsgID (generated function)
func (m *ParamRequestList) MsgID() message.MessageID {
	return MSG_ID_PARAM_REQUEST_LIST
}

// String (generated function)
func (m *ParamRequestList) String() string {
	return fmt.Sprintf(
		"&common.ParamRequestList{ TargetSystem: %+v, TargetComponent: %+v }",
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *ParamRequestList) Marshal() ([]byte, error) {
	payload := make([]byte, 2)
	payload[0] = byte(m.TargetSystem)
	payload[1] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ParamRequestList) Unmarshal(data []byte) error {
	payload := make([]byte, 2)
	copy(payload[0:], data)
	m.TargetSystem = uint8(payload[0])
	m.TargetComponent = uint8(payload[1])
	return nil
}

// ParamValue struct (generated typeinfo)
// Emit the value of a onboard parameter. The inclusion of param_count and param_index in the message allows the recipient to keep track of received parameters and allows him to re-request missing parameters after a loss or timeout. The parameter microservice is documented at https://mavlink.io/en/services/parameter.html
type ParamValue struct {
	ParamValue float32        // Onboard parameter value
	ParamCount uint16         // Total number of onboard parameters
	ParamIndex uint16         // Index of this onboard parameter
	ParamID    string         `len:"16" ` // Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamType  MAV_PARAM_TYPE // Onboard parameter type.
}

// MsgID (generated function)
func (m *ParamValue) MsgID() message.MessageID {
	return MSG_ID_PARAM_VALUE
}

// String (generated function)
func (m *ParamValue) String() string {
	return fmt.Sprintf(
		"&common.ParamValue{ ParamValue: %+v, ParamCount: %+v, ParamIndex: %+v, ParamID: \"%s\", ParamType: %+v }",
		m.ParamValue,
		m.ParamCount,
		m.ParamIndex,
		m.ParamID,
		m.ParamType,
	)
}

// Marshal (generated function)
func (m *ParamValue) Marshal() ([]byte, error) {
	payload := make([]byte, 25)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.ParamValue))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.ParamCount))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.ParamIndex))
	copy(payload[8:], m.ParamID)
	payload[24] = byte(m.ParamType)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ParamValue) Unmarshal(data []byte) error {
	payload := make([]byte, 25)
	copy(payload[0:], data)
	m.ParamValue = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.ParamCount = uint16(binary.LittleEndian.Uint16(payload[4:]))
	m.ParamIndex = uint16(binary.LittleEndian.Uint16(payload[6:]))
	m.ParamID = strings.TrimRight(string(payload[8:24]), string(byte(0)))
	m.ParamType = MAV_PARAM_TYPE(payload[24])
	return nil
}

// ParamSet struct (generated typeinfo)
// Set a parameter value (write new value to permanent storage).
//         The receiving component should acknowledge the new parameter value by broadcasting a PARAM_VALUE message (broadcasting ensures that multiple GCS all have an up-to-date list of all parameters). If the sending GCS did not receive a PARAM_VALUE within its timeout time, it should re-send the PARAM_SET message. The parameter microservice is documented at https://mavlink.io/en/services/parameter.html.
//         PARAM_SET may also be called within the context of a transaction (started with MAV_CMD_PARAM_TRANSACTION). Within a transaction the receiving component should respond with PARAM_ACK_TRANSACTION to the setter component (instead of broadcasting PARAM_VALUE), and PARAM_SET should be re-sent if this is ACK not received.
type ParamSet struct {
	ParamValue      float32        // Onboard parameter value
	TargetSystem    uint8          // System ID
	TargetComponent uint8          // Component ID
	ParamID         string         `len:"16" ` // Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamType       MAV_PARAM_TYPE // Onboard parameter type.
}

// MsgID (generated function)
func (m *ParamSet) MsgID() message.MessageID {
	return MSG_ID_PARAM_SET
}

// String (generated function)
func (m *ParamSet) String() string {
	return fmt.Sprintf(
		"&common.ParamSet{ ParamValue: %+v, TargetSystem: %+v, TargetComponent: %+v, ParamID: \"%s\", ParamType: %+v }",
		m.ParamValue,
		m.TargetSystem,
		m.TargetComponent,
		m.ParamID,
		m.ParamType,
	)
}

// Marshal (generated function)
func (m *ParamSet) Marshal() ([]byte, error) {
	payload := make([]byte, 23)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.ParamValue))
	payload[4] = byte(m.TargetSystem)
	payload[5] = byte(m.TargetComponent)
	copy(payload[6:], m.ParamID)
	payload[22] = byte(m.ParamType)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ParamSet) Unmarshal(data []byte) error {
	payload := make([]byte, 23)
	copy(payload[0:], data)
	m.ParamValue = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.TargetSystem = uint8(payload[4])
	m.TargetComponent = uint8(payload[5])
	m.ParamID = strings.TrimRight(string(payload[6:22]), string(byte(0)))
	m.ParamType = MAV_PARAM_TYPE(payload[22])
	return nil
}

// GpsRawInt struct (generated typeinfo)
// The global position, as returned by the Global Positioning System (GPS). This is
//                 NOT the global position estimate of the system, but rather a RAW sensor value. See message GLOBAL_POSITION for the global position estimate.
type GpsRawInt struct {
	TimeUsec          uint64       // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Lat               int32        // Latitude (WGS84, EGM96 ellipsoid)
	Lon               int32        // Longitude (WGS84, EGM96 ellipsoid)
	Alt               int32        // Altitude (MSL). Positive for up. Note that virtually all GPS modules provide the MSL altitude in addition to the WGS84 altitude.
	Eph               uint16       // GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX
	Epv               uint16       // GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX
	Vel               uint16       // GPS ground speed. If unknown, set to: UINT16_MAX
	Cog               uint16       // Course over ground (NOT heading, but direction of movement) in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	FixType           GPS_FIX_TYPE // GPS fix type.
	SatellitesVisible uint8        // Number of satellites visible. If unknown, set to 255
}

// MsgID (generated function)
func (m *GpsRawInt) MsgID() message.MessageID {
	return MSG_ID_GPS_RAW_INT
}

// String (generated function)
func (m *GpsRawInt) String() string {
	return fmt.Sprintf(
		"&common.GpsRawInt{ TimeUsec: %+v, Lat: %+v, Lon: %+v, Alt: %+v, Eph: %+v, Epv: %+v, Vel: %+v, Cog: %+v, FixType: %+v, SatellitesVisible: %+v }",
		m.TimeUsec,
		m.Lat,
		m.Lon,
		m.Alt,
		m.Eph,
		m.Epv,
		m.Vel,
		m.Cog,
		m.FixType,
		m.SatellitesVisible,
	)
}

// Marshal (generated function)
func (m *GpsRawInt) Marshal() ([]byte, error) {
	payload := make([]byte, 30)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.Alt))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Eph))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.Epv))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.Vel))
	binary.LittleEndian.PutUint16(payload[26:], uint16(m.Cog))
	payload[28] = byte(m.FixType)
	payload[29] = byte(m.SatellitesVisible)
	return payload, nil
}

// Unmarshal (generated function)
func (m *GpsRawInt) Unmarshal(data []byte) error {
	payload := make([]byte, 30)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[12:]))
	m.Alt = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.Eph = uint16(binary.LittleEndian.Uint16(payload[20:]))
	m.Epv = uint16(binary.LittleEndian.Uint16(payload[22:]))
	m.Vel = uint16(binary.LittleEndian.Uint16(payload[24:]))
	m.Cog = uint16(binary.LittleEndian.Uint16(payload[26:]))
	m.FixType = GPS_FIX_TYPE(payload[28])
	m.SatellitesVisible = uint8(payload[29])
	return nil
}

// GpsStatus struct (generated typeinfo)
// The positioning status, as reported by GPS. This message is intended to display status information about each satellite visible to the receiver. See message GLOBAL_POSITION for the global position estimate. This message can contain information for up to 20 satellites.
type GpsStatus struct {
	SatellitesVisible  uint8   // Number of satellites visible
	SatellitePrn       []uint8 `len:"20" ` // Global satellite ID
	SatelliteUsed      []uint8 `len:"20" ` // 0: Satellite not used, 1: used for localization
	SatelliteElevation []uint8 `len:"20" ` // Elevation (0: right on top of receiver, 90: on the horizon) of satellite
	SatelliteAzimuth   []uint8 `len:"20" ` // Direction of satellite, 0: 0 deg, 255: 360 deg.
	SatelliteSnr       []uint8 `len:"20" ` // Signal to noise ratio of satellite
}

// MsgID (generated function)
func (m *GpsStatus) MsgID() message.MessageID {
	return MSG_ID_GPS_STATUS
}

// String (generated function)
func (m *GpsStatus) String() string {
	return fmt.Sprintf(
		"&common.GpsStatus{ SatellitesVisible: %+v, SatellitePrn: %0X (\"%s\"), SatelliteUsed: %0X (\"%s\"), SatelliteElevation: %0X (\"%s\"), SatelliteAzimuth: %0X (\"%s\"), SatelliteSnr: %0X (\"%s\") }",
		m.SatellitesVisible,
		m.SatellitePrn, string(m.SatellitePrn[:]),
		m.SatelliteUsed, string(m.SatelliteUsed[:]),
		m.SatelliteElevation, string(m.SatelliteElevation[:]),
		m.SatelliteAzimuth, string(m.SatelliteAzimuth[:]),
		m.SatelliteSnr, string(m.SatelliteSnr[:]),
	)
}

// Marshal (generated function)
func (m *GpsStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 101)
	payload[0] = byte(m.SatellitesVisible)
	copy(payload[1:], m.SatellitePrn)
	copy(payload[21:], m.SatelliteUsed)
	copy(payload[41:], m.SatelliteElevation)
	copy(payload[61:], m.SatelliteAzimuth)
	copy(payload[81:], m.SatelliteSnr)
	return payload, nil
}

// Unmarshal (generated function)
func (m *GpsStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 101)
	copy(payload[0:], data)
	m.SatellitesVisible = uint8(payload[0])
	copy(m.SatellitePrn[:], payload[1:21])
	copy(m.SatelliteUsed[:], payload[21:41])
	copy(m.SatelliteElevation[:], payload[41:61])
	copy(m.SatelliteAzimuth[:], payload[61:81])
	copy(m.SatelliteSnr[:], payload[81:101])
	return nil
}

// ScaledImu struct (generated typeinfo)
// The RAW IMU readings for the usual 9DOF sensor setup. This message should contain the scaled values to the described units
type ScaledImu struct {
	TimeBootMs uint32 // Timestamp (time since system boot).
	Xacc       int16  // X acceleration
	Yacc       int16  // Y acceleration
	Zacc       int16  // Z acceleration
	Xgyro      int16  // Angular speed around X axis
	Ygyro      int16  // Angular speed around Y axis
	Zgyro      int16  // Angular speed around Z axis
	Xmag       int16  // X Magnetic field
	Ymag       int16  // Y Magnetic field
	Zmag       int16  // Z Magnetic field
}

// MsgID (generated function)
func (m *ScaledImu) MsgID() message.MessageID {
	return MSG_ID_SCALED_IMU
}

// String (generated function)
func (m *ScaledImu) String() string {
	return fmt.Sprintf(
		"&common.ScaledImu{ TimeBootMs: %+v, Xacc: %+v, Yacc: %+v, Zacc: %+v, Xgyro: %+v, Ygyro: %+v, Zgyro: %+v, Xmag: %+v, Ymag: %+v, Zmag: %+v }",
		m.TimeBootMs,
		m.Xacc,
		m.Yacc,
		m.Zacc,
		m.Xgyro,
		m.Ygyro,
		m.Zgyro,
		m.Xmag,
		m.Ymag,
		m.Zmag,
	)
}

// Marshal (generated function)
func (m *ScaledImu) Marshal() ([]byte, error) {
	payload := make([]byte, 22)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Xacc))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.Yacc))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Zacc))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.Xgyro))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Ygyro))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Zgyro))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Xmag))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Ymag))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Zmag))
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScaledImu) Unmarshal(data []byte) error {
	payload := make([]byte, 22)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Xacc = int16(binary.LittleEndian.Uint16(payload[4:]))
	m.Yacc = int16(binary.LittleEndian.Uint16(payload[6:]))
	m.Zacc = int16(binary.LittleEndian.Uint16(payload[8:]))
	m.Xgyro = int16(binary.LittleEndian.Uint16(payload[10:]))
	m.Ygyro = int16(binary.LittleEndian.Uint16(payload[12:]))
	m.Zgyro = int16(binary.LittleEndian.Uint16(payload[14:]))
	m.Xmag = int16(binary.LittleEndian.Uint16(payload[16:]))
	m.Ymag = int16(binary.LittleEndian.Uint16(payload[18:]))
	m.Zmag = int16(binary.LittleEndian.Uint16(payload[20:]))
	return nil
}

// RawImu struct (generated typeinfo)
// The RAW IMU readings for a 9DOF sensor, which is identified by the id (default IMU1). This message should always contain the true raw values without any scaling to allow data capture and system debugging.
type RawImu struct {
	TimeUsec uint64 // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Xacc     int16  // X acceleration (raw)
	Yacc     int16  // Y acceleration (raw)
	Zacc     int16  // Z acceleration (raw)
	Xgyro    int16  // Angular speed around X axis (raw)
	Ygyro    int16  // Angular speed around Y axis (raw)
	Zgyro    int16  // Angular speed around Z axis (raw)
	Xmag     int16  // X Magnetic field (raw)
	Ymag     int16  // Y Magnetic field (raw)
	Zmag     int16  // Z Magnetic field (raw)
}

// MsgID (generated function)
func (m *RawImu) MsgID() message.MessageID {
	return MSG_ID_RAW_IMU
}

// String (generated function)
func (m *RawImu) String() string {
	return fmt.Sprintf(
		"&common.RawImu{ TimeUsec: %+v, Xacc: %+v, Yacc: %+v, Zacc: %+v, Xgyro: %+v, Ygyro: %+v, Zgyro: %+v, Xmag: %+v, Ymag: %+v, Zmag: %+v }",
		m.TimeUsec,
		m.Xacc,
		m.Yacc,
		m.Zacc,
		m.Xgyro,
		m.Ygyro,
		m.Zgyro,
		m.Xmag,
		m.Ymag,
		m.Zmag,
	)
}

// Marshal (generated function)
func (m *RawImu) Marshal() ([]byte, error) {
	payload := make([]byte, 26)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Xacc))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.Yacc))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Zacc))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Xgyro))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Ygyro))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Zgyro))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Xmag))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.Ymag))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.Zmag))
	return payload, nil
}

// Unmarshal (generated function)
func (m *RawImu) Unmarshal(data []byte) error {
	payload := make([]byte, 26)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Xacc = int16(binary.LittleEndian.Uint16(payload[8:]))
	m.Yacc = int16(binary.LittleEndian.Uint16(payload[10:]))
	m.Zacc = int16(binary.LittleEndian.Uint16(payload[12:]))
	m.Xgyro = int16(binary.LittleEndian.Uint16(payload[14:]))
	m.Ygyro = int16(binary.LittleEndian.Uint16(payload[16:]))
	m.Zgyro = int16(binary.LittleEndian.Uint16(payload[18:]))
	m.Xmag = int16(binary.LittleEndian.Uint16(payload[20:]))
	m.Ymag = int16(binary.LittleEndian.Uint16(payload[22:]))
	m.Zmag = int16(binary.LittleEndian.Uint16(payload[24:]))
	return nil
}

// RawPressure struct (generated typeinfo)
// The RAW pressure readings for the typical setup of one absolute pressure and one differential pressure sensor. The sensor values should be the raw, UNSCALED ADC values.
type RawPressure struct {
	TimeUsec    uint64 // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	PressAbs    int16  // Absolute pressure (raw)
	PressDiff1  int16  // Differential pressure 1 (raw, 0 if nonexistent)
	PressDiff2  int16  // Differential pressure 2 (raw, 0 if nonexistent)
	Temperature int16  // Raw Temperature measurement (raw)
}

// MsgID (generated function)
func (m *RawPressure) MsgID() message.MessageID {
	return MSG_ID_RAW_PRESSURE
}

// String (generated function)
func (m *RawPressure) String() string {
	return fmt.Sprintf(
		"&common.RawPressure{ TimeUsec: %+v, PressAbs: %+v, PressDiff1: %+v, PressDiff2: %+v, Temperature: %+v }",
		m.TimeUsec,
		m.PressAbs,
		m.PressDiff1,
		m.PressDiff2,
		m.Temperature,
	)
}

// Marshal (generated function)
func (m *RawPressure) Marshal() ([]byte, error) {
	payload := make([]byte, 16)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.PressAbs))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.PressDiff1))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.PressDiff2))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Temperature))
	return payload, nil
}

// Unmarshal (generated function)
func (m *RawPressure) Unmarshal(data []byte) error {
	payload := make([]byte, 16)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.PressAbs = int16(binary.LittleEndian.Uint16(payload[8:]))
	m.PressDiff1 = int16(binary.LittleEndian.Uint16(payload[10:]))
	m.PressDiff2 = int16(binary.LittleEndian.Uint16(payload[12:]))
	m.Temperature = int16(binary.LittleEndian.Uint16(payload[14:]))
	return nil
}

// ScaledPressure struct (generated typeinfo)
// The pressure readings for the typical setup of one absolute and differential pressure sensor. The units are as specified in each field.
type ScaledPressure struct {
	TimeBootMs  uint32  // Timestamp (time since system boot).
	PressAbs    float32 // Absolute pressure
	PressDiff   float32 // Differential pressure 1
	Temperature int16   // Absolute pressure temperature
}

// MsgID (generated function)
func (m *ScaledPressure) MsgID() message.MessageID {
	return MSG_ID_SCALED_PRESSURE
}

// String (generated function)
func (m *ScaledPressure) String() string {
	return fmt.Sprintf(
		"&common.ScaledPressure{ TimeBootMs: %+v, PressAbs: %+v, PressDiff: %+v, Temperature: %+v }",
		m.TimeBootMs,
		m.PressAbs,
		m.PressDiff,
		m.Temperature,
	)
}

// Marshal (generated function)
func (m *ScaledPressure) Marshal() ([]byte, error) {
	payload := make([]byte, 14)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.PressAbs))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.PressDiff))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Temperature))
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScaledPressure) Unmarshal(data []byte) error {
	payload := make([]byte, 14)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.PressAbs = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.PressDiff = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Temperature = int16(binary.LittleEndian.Uint16(payload[12:]))
	return nil
}

// Attitude struct (generated typeinfo)
// The attitude in the aeronautical frame (right-handed, Z-down, X-front, Y-right).
type Attitude struct {
	TimeBootMs uint32  // Timestamp (time since system boot).
	Roll       float32 // Roll angle (-pi..+pi)
	Pitch      float32 // Pitch angle (-pi..+pi)
	Yaw        float32 // Yaw angle (-pi..+pi)
	Rollspeed  float32 // Roll angular speed
	Pitchspeed float32 // Pitch angular speed
	Yawspeed   float32 // Yaw angular speed
}

// MsgID (generated function)
func (m *Attitude) MsgID() message.MessageID {
	return MSG_ID_ATTITUDE
}

// String (generated function)
func (m *Attitude) String() string {
	return fmt.Sprintf(
		"&common.Attitude{ TimeBootMs: %+v, Roll: %+v, Pitch: %+v, Yaw: %+v, Rollspeed: %+v, Pitchspeed: %+v, Yawspeed: %+v }",
		m.TimeBootMs,
		m.Roll,
		m.Pitch,
		m.Yaw,
		m.Rollspeed,
		m.Pitchspeed,
		m.Yawspeed,
	)
}

// Marshal (generated function)
func (m *Attitude) Marshal() ([]byte, error) {
	payload := make([]byte, 28)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Roll))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Pitch))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Yaw))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Rollspeed))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Pitchspeed))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Yawspeed))
	return payload, nil
}

// Unmarshal (generated function)
func (m *Attitude) Unmarshal(data []byte) error {
	payload := make([]byte, 28)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Roll = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Pitch = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Rollspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Pitchspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Yawspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	return nil
}

// AttitudeQuaternion struct (generated typeinfo)
// The attitude in the aeronautical frame (right-handed, Z-down, X-front, Y-right), expressed as quaternion. Quaternion order is w, x, y, z and a zero rotation would be expressed as (1 0 0 0).
type AttitudeQuaternion struct {
	TimeBootMs uint32  // Timestamp (time since system boot).
	Q1         float32 // Quaternion component 1, w (1 in null-rotation)
	Q2         float32 // Quaternion component 2, x (0 in null-rotation)
	Q3         float32 // Quaternion component 3, y (0 in null-rotation)
	Q4         float32 // Quaternion component 4, z (0 in null-rotation)
	Rollspeed  float32 // Roll angular speed
	Pitchspeed float32 // Pitch angular speed
	Yawspeed   float32 // Yaw angular speed
}

// MsgID (generated function)
func (m *AttitudeQuaternion) MsgID() message.MessageID {
	return MSG_ID_ATTITUDE_QUATERNION
}

// String (generated function)
func (m *AttitudeQuaternion) String() string {
	return fmt.Sprintf(
		"&common.AttitudeQuaternion{ TimeBootMs: %+v, Q1: %+v, Q2: %+v, Q3: %+v, Q4: %+v, Rollspeed: %+v, Pitchspeed: %+v, Yawspeed: %+v }",
		m.TimeBootMs,
		m.Q1,
		m.Q2,
		m.Q3,
		m.Q4,
		m.Rollspeed,
		m.Pitchspeed,
		m.Yawspeed,
	)
}

// Marshal (generated function)
func (m *AttitudeQuaternion) Marshal() ([]byte, error) {
	payload := make([]byte, 32)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Q1))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Q2))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Q3))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Q4))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Rollspeed))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Pitchspeed))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Yawspeed))
	return payload, nil
}

// Unmarshal (generated function)
func (m *AttitudeQuaternion) Unmarshal(data []byte) error {
	payload := make([]byte, 32)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Q1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Q2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Q3 = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Q4 = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Rollspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Pitchspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Yawspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	return nil
}

// LocalPositionNed struct (generated typeinfo)
// The filtered local position (e.g. fused computer vision and accelerometers). Coordinate frame is right-handed, Z-axis down (aeronautical frame, NED / north-east-down convention)
type LocalPositionNed struct {
	TimeBootMs uint32  // Timestamp (time since system boot).
	X          float32 // X Position
	Y          float32 // Y Position
	Z          float32 // Z Position
	Vx         float32 // X Speed
	Vy         float32 // Y Speed
	Vz         float32 // Z Speed
}

// MsgID (generated function)
func (m *LocalPositionNed) MsgID() message.MessageID {
	return MSG_ID_LOCAL_POSITION_NED
}

// String (generated function)
func (m *LocalPositionNed) String() string {
	return fmt.Sprintf(
		"&common.LocalPositionNed{ TimeBootMs: %+v, X: %+v, Y: %+v, Z: %+v, Vx: %+v, Vy: %+v, Vz: %+v }",
		m.TimeBootMs,
		m.X,
		m.Y,
		m.Z,
		m.Vx,
		m.Vy,
		m.Vz,
	)
}

// Marshal (generated function)
func (m *LocalPositionNed) Marshal() ([]byte, error) {
	payload := make([]byte, 28)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Vx))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Vy))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Vz))
	return payload, nil
}

// Unmarshal (generated function)
func (m *LocalPositionNed) Unmarshal(data []byte) error {
	payload := make([]byte, 28)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Vx = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Vy = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Vz = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	return nil
}

// GlobalPositionInt struct (generated typeinfo)
// The filtered global position (e.g. fused GPS and accelerometers). The position is in GPS-frame (right-handed, Z-up). It
//                is designed as scaled integer message since the resolution of float is not sufficient.
type GlobalPositionInt struct {
	TimeBootMs  uint32 // Timestamp (time since system boot).
	Lat         int32  // Latitude, expressed
	Lon         int32  // Longitude, expressed
	Alt         int32  // Altitude (MSL). Note that virtually all GPS modules provide both WGS84 and MSL.
	RelativeAlt int32  // Altitude above ground
	Vx          int16  // Ground X Speed (Latitude, positive north)
	Vy          int16  // Ground Y Speed (Longitude, positive east)
	Vz          int16  // Ground Z Speed (Altitude, positive down)
	Hdg         uint16 // Vehicle heading (yaw angle), 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
}

// MsgID (generated function)
func (m *GlobalPositionInt) MsgID() message.MessageID {
	return MSG_ID_GLOBAL_POSITION_INT
}

// String (generated function)
func (m *GlobalPositionInt) String() string {
	return fmt.Sprintf(
		"&common.GlobalPositionInt{ TimeBootMs: %+v, Lat: %+v, Lon: %+v, Alt: %+v, RelativeAlt: %+v, Vx: %+v, Vy: %+v, Vz: %+v, Hdg: %+v }",
		m.TimeBootMs,
		m.Lat,
		m.Lon,
		m.Alt,
		m.RelativeAlt,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Hdg,
	)
}

// Marshal (generated function)
func (m *GlobalPositionInt) Marshal() ([]byte, error) {
	payload := make([]byte, 28)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.Alt))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.RelativeAlt))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Vx))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.Vy))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.Vz))
	binary.LittleEndian.PutUint16(payload[26:], uint16(m.Hdg))
	return payload, nil
}

// Unmarshal (generated function)
func (m *GlobalPositionInt) Unmarshal(data []byte) error {
	payload := make([]byte, 28)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.Alt = int32(binary.LittleEndian.Uint32(payload[12:]))
	m.RelativeAlt = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.Vx = int16(binary.LittleEndian.Uint16(payload[20:]))
	m.Vy = int16(binary.LittleEndian.Uint16(payload[22:]))
	m.Vz = int16(binary.LittleEndian.Uint16(payload[24:]))
	m.Hdg = uint16(binary.LittleEndian.Uint16(payload[26:]))
	return nil
}

// RcChannelsScaled struct (generated typeinfo)
// The scaled values of the RC channels received: (-100%) -10000, (0%) 0, (100%) 10000. Channels that are inactive should be set to UINT16_MAX.
type RcChannelsScaled struct {
	TimeBootMs  uint32 // Timestamp (time since system boot).
	Chan1Scaled int16  // RC channel 1 value scaled.
	Chan2Scaled int16  // RC channel 2 value scaled.
	Chan3Scaled int16  // RC channel 3 value scaled.
	Chan4Scaled int16  // RC channel 4 value scaled.
	Chan5Scaled int16  // RC channel 5 value scaled.
	Chan6Scaled int16  // RC channel 6 value scaled.
	Chan7Scaled int16  // RC channel 7 value scaled.
	Chan8Scaled int16  // RC channel 8 value scaled.
	Port        uint8  // Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
	Rssi        uint8  // Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
}

// MsgID (generated function)
func (m *RcChannelsScaled) MsgID() message.MessageID {
	return MSG_ID_RC_CHANNELS_SCALED
}

// String (generated function)
func (m *RcChannelsScaled) String() string {
	return fmt.Sprintf(
		"&common.RcChannelsScaled{ TimeBootMs: %+v, Chan1Scaled: %+v, Chan2Scaled: %+v, Chan3Scaled: %+v, Chan4Scaled: %+v, Chan5Scaled: %+v, Chan6Scaled: %+v, Chan7Scaled: %+v, Chan8Scaled: %+v, Port: %+v, Rssi: %+v }",
		m.TimeBootMs,
		m.Chan1Scaled,
		m.Chan2Scaled,
		m.Chan3Scaled,
		m.Chan4Scaled,
		m.Chan5Scaled,
		m.Chan6Scaled,
		m.Chan7Scaled,
		m.Chan8Scaled,
		m.Port,
		m.Rssi,
	)
}

// Marshal (generated function)
func (m *RcChannelsScaled) Marshal() ([]byte, error) {
	payload := make([]byte, 22)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Chan1Scaled))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.Chan2Scaled))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Chan3Scaled))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.Chan4Scaled))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Chan5Scaled))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Chan6Scaled))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Chan7Scaled))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Chan8Scaled))
	payload[20] = byte(m.Port)
	payload[21] = byte(m.Rssi)
	return payload, nil
}

// Unmarshal (generated function)
func (m *RcChannelsScaled) Unmarshal(data []byte) error {
	payload := make([]byte, 22)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Chan1Scaled = int16(binary.LittleEndian.Uint16(payload[4:]))
	m.Chan2Scaled = int16(binary.LittleEndian.Uint16(payload[6:]))
	m.Chan3Scaled = int16(binary.LittleEndian.Uint16(payload[8:]))
	m.Chan4Scaled = int16(binary.LittleEndian.Uint16(payload[10:]))
	m.Chan5Scaled = int16(binary.LittleEndian.Uint16(payload[12:]))
	m.Chan6Scaled = int16(binary.LittleEndian.Uint16(payload[14:]))
	m.Chan7Scaled = int16(binary.LittleEndian.Uint16(payload[16:]))
	m.Chan8Scaled = int16(binary.LittleEndian.Uint16(payload[18:]))
	m.Port = uint8(payload[20])
	m.Rssi = uint8(payload[21])
	return nil
}

// RcChannelsRaw struct (generated typeinfo)
// The RAW values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. A value of UINT16_MAX implies the channel is unused. Individual receivers/transmitters might violate this specification.
type RcChannelsRaw struct {
	TimeBootMs uint32 // Timestamp (time since system boot).
	Chan1Raw   uint16 // RC channel 1 value.
	Chan2Raw   uint16 // RC channel 2 value.
	Chan3Raw   uint16 // RC channel 3 value.
	Chan4Raw   uint16 // RC channel 4 value.
	Chan5Raw   uint16 // RC channel 5 value.
	Chan6Raw   uint16 // RC channel 6 value.
	Chan7Raw   uint16 // RC channel 7 value.
	Chan8Raw   uint16 // RC channel 8 value.
	Port       uint8  // Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
	Rssi       uint8  // Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
}

// MsgID (generated function)
func (m *RcChannelsRaw) MsgID() message.MessageID {
	return MSG_ID_RC_CHANNELS_RAW
}

// String (generated function)
func (m *RcChannelsRaw) String() string {
	return fmt.Sprintf(
		"&common.RcChannelsRaw{ TimeBootMs: %+v, Chan1Raw: %+v, Chan2Raw: %+v, Chan3Raw: %+v, Chan4Raw: %+v, Chan5Raw: %+v, Chan6Raw: %+v, Chan7Raw: %+v, Chan8Raw: %+v, Port: %+v, Rssi: %+v }",
		m.TimeBootMs,
		m.Chan1Raw,
		m.Chan2Raw,
		m.Chan3Raw,
		m.Chan4Raw,
		m.Chan5Raw,
		m.Chan6Raw,
		m.Chan7Raw,
		m.Chan8Raw,
		m.Port,
		m.Rssi,
	)
}

// Marshal (generated function)
func (m *RcChannelsRaw) Marshal() ([]byte, error) {
	payload := make([]byte, 22)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Chan1Raw))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.Chan2Raw))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Chan3Raw))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.Chan4Raw))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Chan5Raw))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Chan6Raw))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Chan7Raw))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Chan8Raw))
	payload[20] = byte(m.Port)
	payload[21] = byte(m.Rssi)
	return payload, nil
}

// Unmarshal (generated function)
func (m *RcChannelsRaw) Unmarshal(data []byte) error {
	payload := make([]byte, 22)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Chan1Raw = uint16(binary.LittleEndian.Uint16(payload[4:]))
	m.Chan2Raw = uint16(binary.LittleEndian.Uint16(payload[6:]))
	m.Chan3Raw = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.Chan4Raw = uint16(binary.LittleEndian.Uint16(payload[10:]))
	m.Chan5Raw = uint16(binary.LittleEndian.Uint16(payload[12:]))
	m.Chan6Raw = uint16(binary.LittleEndian.Uint16(payload[14:]))
	m.Chan7Raw = uint16(binary.LittleEndian.Uint16(payload[16:]))
	m.Chan8Raw = uint16(binary.LittleEndian.Uint16(payload[18:]))
	m.Port = uint8(payload[20])
	m.Rssi = uint8(payload[21])
	return nil
}

// ServoOutputRaw struct (generated typeinfo)
// Superseded by ACTUATOR_OUTPUT_STATUS. The RAW values of the servo outputs (for RC input from the remote, use the RC_CHANNELS messages). The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%.
type ServoOutputRaw struct {
	TimeUsec  uint32 // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Servo1Raw uint16 // Servo output 1 value
	Servo2Raw uint16 // Servo output 2 value
	Servo3Raw uint16 // Servo output 3 value
	Servo4Raw uint16 // Servo output 4 value
	Servo5Raw uint16 // Servo output 5 value
	Servo6Raw uint16 // Servo output 6 value
	Servo7Raw uint16 // Servo output 7 value
	Servo8Raw uint16 // Servo output 8 value
	Port      uint8  // Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
}

// MsgID (generated function)
func (m *ServoOutputRaw) MsgID() message.MessageID {
	return MSG_ID_SERVO_OUTPUT_RAW
}

// String (generated function)
func (m *ServoOutputRaw) String() string {
	return fmt.Sprintf(
		"&common.ServoOutputRaw{ TimeUsec: %+v, Servo1Raw: %+v, Servo2Raw: %+v, Servo3Raw: %+v, Servo4Raw: %+v, Servo5Raw: %+v, Servo6Raw: %+v, Servo7Raw: %+v, Servo8Raw: %+v, Port: %+v }",
		m.TimeUsec,
		m.Servo1Raw,
		m.Servo2Raw,
		m.Servo3Raw,
		m.Servo4Raw,
		m.Servo5Raw,
		m.Servo6Raw,
		m.Servo7Raw,
		m.Servo8Raw,
		m.Port,
	)
}

// Marshal (generated function)
func (m *ServoOutputRaw) Marshal() ([]byte, error) {
	payload := make([]byte, 21)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeUsec))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Servo1Raw))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.Servo2Raw))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Servo3Raw))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.Servo4Raw))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Servo5Raw))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Servo6Raw))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Servo7Raw))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Servo8Raw))
	payload[20] = byte(m.Port)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ServoOutputRaw) Unmarshal(data []byte) error {
	payload := make([]byte, 21)
	copy(payload[0:], data)
	m.TimeUsec = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Servo1Raw = uint16(binary.LittleEndian.Uint16(payload[4:]))
	m.Servo2Raw = uint16(binary.LittleEndian.Uint16(payload[6:]))
	m.Servo3Raw = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.Servo4Raw = uint16(binary.LittleEndian.Uint16(payload[10:]))
	m.Servo5Raw = uint16(binary.LittleEndian.Uint16(payload[12:]))
	m.Servo6Raw = uint16(binary.LittleEndian.Uint16(payload[14:]))
	m.Servo7Raw = uint16(binary.LittleEndian.Uint16(payload[16:]))
	m.Servo8Raw = uint16(binary.LittleEndian.Uint16(payload[18:]))
	m.Port = uint8(payload[20])
	return nil
}

// MissionRequestPartialList struct (generated typeinfo)
// Request a partial list of mission items from the system/component. https://mavlink.io/en/services/mission.html. If start and end index are the same, just send one waypoint.
type MissionRequestPartialList struct {
	StartIndex      int16 // Start index
	EndIndex        int16 // End index, -1 by default (-1: send list to end). Else a valid index of the list
	TargetSystem    uint8 // System ID
	TargetComponent uint8 // Component ID
}

// MsgID (generated function)
func (m *MissionRequestPartialList) MsgID() message.MessageID {
	return MSG_ID_MISSION_REQUEST_PARTIAL_LIST
}

// String (generated function)
func (m *MissionRequestPartialList) String() string {
	return fmt.Sprintf(
		"&common.MissionRequestPartialList{ StartIndex: %+v, EndIndex: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.StartIndex,
		m.EndIndex,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *MissionRequestPartialList) Marshal() ([]byte, error) {
	payload := make([]byte, 6)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.StartIndex))
	binary.LittleEndian.PutUint16(payload[2:], uint16(m.EndIndex))
	payload[4] = byte(m.TargetSystem)
	payload[5] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionRequestPartialList) Unmarshal(data []byte) error {
	payload := make([]byte, 6)
	copy(payload[0:], data)
	m.StartIndex = int16(binary.LittleEndian.Uint16(payload[0:]))
	m.EndIndex = int16(binary.LittleEndian.Uint16(payload[2:]))
	m.TargetSystem = uint8(payload[4])
	m.TargetComponent = uint8(payload[5])
	return nil
}

// MissionWritePartialList struct (generated typeinfo)
// This message is sent to the MAV to write a partial list. If start index == end index, only one item will be transmitted / updated. If the start index is NOT 0 and above the current list size, this request should be REJECTED!
type MissionWritePartialList struct {
	StartIndex      int16 // Start index. Must be smaller / equal to the largest index of the current onboard list.
	EndIndex        int16 // End index, equal or greater than start index.
	TargetSystem    uint8 // System ID
	TargetComponent uint8 // Component ID
}

// MsgID (generated function)
func (m *MissionWritePartialList) MsgID() message.MessageID {
	return MSG_ID_MISSION_WRITE_PARTIAL_LIST
}

// String (generated function)
func (m *MissionWritePartialList) String() string {
	return fmt.Sprintf(
		"&common.MissionWritePartialList{ StartIndex: %+v, EndIndex: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.StartIndex,
		m.EndIndex,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *MissionWritePartialList) Marshal() ([]byte, error) {
	payload := make([]byte, 6)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.StartIndex))
	binary.LittleEndian.PutUint16(payload[2:], uint16(m.EndIndex))
	payload[4] = byte(m.TargetSystem)
	payload[5] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionWritePartialList) Unmarshal(data []byte) error {
	payload := make([]byte, 6)
	copy(payload[0:], data)
	m.StartIndex = int16(binary.LittleEndian.Uint16(payload[0:]))
	m.EndIndex = int16(binary.LittleEndian.Uint16(payload[2:]))
	m.TargetSystem = uint8(payload[4])
	m.TargetComponent = uint8(payload[5])
	return nil
}

// MissionItem struct (generated typeinfo)
// Message encoding a mission item. This message is emitted to announce
//                 the presence of a mission item and to set a mission item on the system. The mission item can be either in x, y, z meters (type: LOCAL) or x:lat, y:lon, z:altitude. Local frame is Z-down, right handed (NED), global frame is Z-up, right handed (ENU). NaN may be used to indicate an optional/default value (e.g. to use the system's current latitude or yaw rather than a specific value). See also https://mavlink.io/en/services/mission.html.
type MissionItem struct {
	Param1          float32   // PARAM1, see MAV_CMD enum
	Param2          float32   // PARAM2, see MAV_CMD enum
	Param3          float32   // PARAM3, see MAV_CMD enum
	Param4          float32   // PARAM4, see MAV_CMD enum
	X               float32   // PARAM5 / local: X coordinate, global: latitude
	Y               float32   // PARAM6 / local: Y coordinate, global: longitude
	Z               float32   // PARAM7 / local: Z coordinate, global: altitude (relative or absolute, depending on frame).
	Seq             uint16    // Sequence
	Command         MAV_CMD   // The scheduled action for the waypoint.
	TargetSystem    uint8     // System ID
	TargetComponent uint8     // Component ID
	Frame           MAV_FRAME // The coordinate system of the waypoint.
	Current         uint8     // false:0, true:1
	Autocontinue    uint8     // Autocontinue to next waypoint
}

// MsgID (generated function)
func (m *MissionItem) MsgID() message.MessageID {
	return MSG_ID_MISSION_ITEM
}

// String (generated function)
func (m *MissionItem) String() string {
	return fmt.Sprintf(
		"&common.MissionItem{ Param1: %+v, Param2: %+v, Param3: %+v, Param4: %+v, X: %+v, Y: %+v, Z: %+v, Seq: %+v, Command: %+v, TargetSystem: %+v, TargetComponent: %+v, Frame: %+v, Current: %+v, Autocontinue: %+v }",
		m.Param1,
		m.Param2,
		m.Param3,
		m.Param4,
		m.X,
		m.Y,
		m.Z,
		m.Seq,
		m.Command,
		m.TargetSystem,
		m.TargetComponent,
		m.Frame,
		m.Current,
		m.Autocontinue,
	)
}

// Marshal (generated function)
func (m *MissionItem) Marshal() ([]byte, error) {
	payload := make([]byte, 37)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.Param1))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Param2))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Param3))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Param4))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Seq))
	binary.LittleEndian.PutUint16(payload[30:], uint16(m.Command))
	payload[32] = byte(m.TargetSystem)
	payload[33] = byte(m.TargetComponent)
	payload[34] = byte(m.Frame)
	payload[35] = byte(m.Current)
	payload[36] = byte(m.Autocontinue)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionItem) Unmarshal(data []byte) error {
	payload := make([]byte, 37)
	copy(payload[0:], data)
	m.Param1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.Param2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Param3 = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Param4 = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Seq = uint16(binary.LittleEndian.Uint16(payload[28:]))
	m.Command = MAV_CMD(binary.LittleEndian.Uint16(payload[30:]))
	m.TargetSystem = uint8(payload[32])
	m.TargetComponent = uint8(payload[33])
	m.Frame = MAV_FRAME(payload[34])
	m.Current = uint8(payload[35])
	m.Autocontinue = uint8(payload[36])
	return nil
}

// MissionRequest struct (generated typeinfo)
// Request the information of the mission item with the sequence number seq. The response of the system to this message should be a MISSION_ITEM message. https://mavlink.io/en/services/mission.html
type MissionRequest struct {
	Seq             uint16 // Sequence
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
}

// MsgID (generated function)
func (m *MissionRequest) MsgID() message.MessageID {
	return MSG_ID_MISSION_REQUEST
}

// String (generated function)
func (m *MissionRequest) String() string {
	return fmt.Sprintf(
		"&common.MissionRequest{ Seq: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.Seq,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *MissionRequest) Marshal() ([]byte, error) {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Seq))
	payload[2] = byte(m.TargetSystem)
	payload[3] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionRequest) Unmarshal(data []byte) error {
	payload := make([]byte, 4)
	copy(payload[0:], data)
	m.Seq = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetSystem = uint8(payload[2])
	m.TargetComponent = uint8(payload[3])
	return nil
}

// MissionSetCurrent struct (generated typeinfo)
// Set the mission item with sequence number seq as current item. This means that the MAV will continue to this mission item on the shortest path (not following the mission items in-between).
type MissionSetCurrent struct {
	Seq             uint16 // Sequence
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
}

// MsgID (generated function)
func (m *MissionSetCurrent) MsgID() message.MessageID {
	return MSG_ID_MISSION_SET_CURRENT
}

// String (generated function)
func (m *MissionSetCurrent) String() string {
	return fmt.Sprintf(
		"&common.MissionSetCurrent{ Seq: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.Seq,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *MissionSetCurrent) Marshal() ([]byte, error) {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Seq))
	payload[2] = byte(m.TargetSystem)
	payload[3] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionSetCurrent) Unmarshal(data []byte) error {
	payload := make([]byte, 4)
	copy(payload[0:], data)
	m.Seq = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetSystem = uint8(payload[2])
	m.TargetComponent = uint8(payload[3])
	return nil
}

// MissionCurrent struct (generated typeinfo)
// Message that announces the sequence number of the current active mission item. The MAV will fly towards this mission item.
type MissionCurrent struct {
	Seq uint16 // Sequence
}

// MsgID (generated function)
func (m *MissionCurrent) MsgID() message.MessageID {
	return MSG_ID_MISSION_CURRENT
}

// String (generated function)
func (m *MissionCurrent) String() string {
	return fmt.Sprintf(
		"&common.MissionCurrent{ Seq: %+v }",
		m.Seq,
	)
}

// Marshal (generated function)
func (m *MissionCurrent) Marshal() ([]byte, error) {
	payload := make([]byte, 2)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Seq))
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionCurrent) Unmarshal(data []byte) error {
	payload := make([]byte, 2)
	copy(payload[0:], data)
	m.Seq = uint16(binary.LittleEndian.Uint16(payload[0:]))
	return nil
}

// MissionRequestList struct (generated typeinfo)
// Request the overall list of mission items from the system/component.
type MissionRequestList struct {
	TargetSystem    uint8 // System ID
	TargetComponent uint8 // Component ID
}

// MsgID (generated function)
func (m *MissionRequestList) MsgID() message.MessageID {
	return MSG_ID_MISSION_REQUEST_LIST
}

// String (generated function)
func (m *MissionRequestList) String() string {
	return fmt.Sprintf(
		"&common.MissionRequestList{ TargetSystem: %+v, TargetComponent: %+v }",
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *MissionRequestList) Marshal() ([]byte, error) {
	payload := make([]byte, 2)
	payload[0] = byte(m.TargetSystem)
	payload[1] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionRequestList) Unmarshal(data []byte) error {
	payload := make([]byte, 2)
	copy(payload[0:], data)
	m.TargetSystem = uint8(payload[0])
	m.TargetComponent = uint8(payload[1])
	return nil
}

// MissionCount struct (generated typeinfo)
// This message is emitted as response to MISSION_REQUEST_LIST by the MAV and to initiate a write transaction. The GCS can then request the individual mission item based on the knowledge of the total number of waypoints.
type MissionCount struct {
	Count           uint16 // Number of mission items in the sequence
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
}

// MsgID (generated function)
func (m *MissionCount) MsgID() message.MessageID {
	return MSG_ID_MISSION_COUNT
}

// String (generated function)
func (m *MissionCount) String() string {
	return fmt.Sprintf(
		"&common.MissionCount{ Count: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.Count,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *MissionCount) Marshal() ([]byte, error) {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Count))
	payload[2] = byte(m.TargetSystem)
	payload[3] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionCount) Unmarshal(data []byte) error {
	payload := make([]byte, 4)
	copy(payload[0:], data)
	m.Count = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetSystem = uint8(payload[2])
	m.TargetComponent = uint8(payload[3])
	return nil
}

// MissionClearAll struct (generated typeinfo)
// Delete all mission items at once.
type MissionClearAll struct {
	TargetSystem    uint8 // System ID
	TargetComponent uint8 // Component ID
}

// MsgID (generated function)
func (m *MissionClearAll) MsgID() message.MessageID {
	return MSG_ID_MISSION_CLEAR_ALL
}

// String (generated function)
func (m *MissionClearAll) String() string {
	return fmt.Sprintf(
		"&common.MissionClearAll{ TargetSystem: %+v, TargetComponent: %+v }",
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *MissionClearAll) Marshal() ([]byte, error) {
	payload := make([]byte, 2)
	payload[0] = byte(m.TargetSystem)
	payload[1] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionClearAll) Unmarshal(data []byte) error {
	payload := make([]byte, 2)
	copy(payload[0:], data)
	m.TargetSystem = uint8(payload[0])
	m.TargetComponent = uint8(payload[1])
	return nil
}

// MissionItemReached struct (generated typeinfo)
// A certain mission item has been reached. The system will either hold this position (or circle on the orbit) or (if the autocontinue on the WP was set) continue to the next waypoint.
type MissionItemReached struct {
	Seq uint16 // Sequence
}

// MsgID (generated function)
func (m *MissionItemReached) MsgID() message.MessageID {
	return MSG_ID_MISSION_ITEM_REACHED
}

// String (generated function)
func (m *MissionItemReached) String() string {
	return fmt.Sprintf(
		"&common.MissionItemReached{ Seq: %+v }",
		m.Seq,
	)
}

// Marshal (generated function)
func (m *MissionItemReached) Marshal() ([]byte, error) {
	payload := make([]byte, 2)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Seq))
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionItemReached) Unmarshal(data []byte) error {
	payload := make([]byte, 2)
	copy(payload[0:], data)
	m.Seq = uint16(binary.LittleEndian.Uint16(payload[0:]))
	return nil
}

// MissionAck struct (generated typeinfo)
// Acknowledgment message during waypoint handling. The type field states if this message is a positive ack (type=0) or if an error happened (type=non-zero).
type MissionAck struct {
	TargetSystem    uint8              // System ID
	TargetComponent uint8              // Component ID
	Type            MAV_MISSION_RESULT // Mission result.
}

// MsgID (generated function)
func (m *MissionAck) MsgID() message.MessageID {
	return MSG_ID_MISSION_ACK
}

// String (generated function)
func (m *MissionAck) String() string {
	return fmt.Sprintf(
		"&common.MissionAck{ TargetSystem: %+v, TargetComponent: %+v, Type: %+v }",
		m.TargetSystem,
		m.TargetComponent,
		m.Type,
	)
}

// Marshal (generated function)
func (m *MissionAck) Marshal() ([]byte, error) {
	payload := make([]byte, 3)
	payload[0] = byte(m.TargetSystem)
	payload[1] = byte(m.TargetComponent)
	payload[2] = byte(m.Type)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionAck) Unmarshal(data []byte) error {
	payload := make([]byte, 3)
	copy(payload[0:], data)
	m.TargetSystem = uint8(payload[0])
	m.TargetComponent = uint8(payload[1])
	m.Type = MAV_MISSION_RESULT(payload[2])
	return nil
}

// SetGpsGlobalOrigin struct (generated typeinfo)
// Sets the GPS co-ordinates of the vehicle local origin (0,0,0) position. Vehicle should emit GPS_GLOBAL_ORIGIN irrespective of whether the origin is changed. This enables transform between the local coordinate frame and the global (GPS) coordinate frame, which may be necessary when (for example) indoor and outdoor settings are connected and the MAV should move from in- to outdoor.
type SetGpsGlobalOrigin struct {
	Latitude     int32 // Latitude (WGS84)
	Longitude    int32 // Longitude (WGS84)
	Altitude     int32 // Altitude (MSL). Positive for up.
	TargetSystem uint8 // System ID
}

// MsgID (generated function)
func (m *SetGpsGlobalOrigin) MsgID() message.MessageID {
	return MSG_ID_SET_GPS_GLOBAL_ORIGIN
}

// String (generated function)
func (m *SetGpsGlobalOrigin) String() string {
	return fmt.Sprintf(
		"&common.SetGpsGlobalOrigin{ Latitude: %+v, Longitude: %+v, Altitude: %+v, TargetSystem: %+v }",
		m.Latitude,
		m.Longitude,
		m.Altitude,
		m.TargetSystem,
	)
}

// Marshal (generated function)
func (m *SetGpsGlobalOrigin) Marshal() ([]byte, error) {
	payload := make([]byte, 13)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Latitude))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Longitude))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Altitude))
	payload[12] = byte(m.TargetSystem)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SetGpsGlobalOrigin) Unmarshal(data []byte) error {
	payload := make([]byte, 13)
	copy(payload[0:], data)
	m.Latitude = int32(binary.LittleEndian.Uint32(payload[0:]))
	m.Longitude = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.Altitude = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.TargetSystem = uint8(payload[12])
	return nil
}

// GpsGlobalOrigin struct (generated typeinfo)
// Publishes the GPS co-ordinates of the vehicle local origin (0,0,0) position. Emitted whenever a new GPS-Local position mapping is requested or set - e.g. following SET_GPS_GLOBAL_ORIGIN message.
type GpsGlobalOrigin struct {
	Latitude  int32 // Latitude (WGS84)
	Longitude int32 // Longitude (WGS84)
	Altitude  int32 // Altitude (MSL). Positive for up.
}

// MsgID (generated function)
func (m *GpsGlobalOrigin) MsgID() message.MessageID {
	return MSG_ID_GPS_GLOBAL_ORIGIN
}

// String (generated function)
func (m *GpsGlobalOrigin) String() string {
	return fmt.Sprintf(
		"&common.GpsGlobalOrigin{ Latitude: %+v, Longitude: %+v, Altitude: %+v }",
		m.Latitude,
		m.Longitude,
		m.Altitude,
	)
}

// Marshal (generated function)
func (m *GpsGlobalOrigin) Marshal() ([]byte, error) {
	payload := make([]byte, 12)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Latitude))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Longitude))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Altitude))
	return payload, nil
}

// Unmarshal (generated function)
func (m *GpsGlobalOrigin) Unmarshal(data []byte) error {
	payload := make([]byte, 12)
	copy(payload[0:], data)
	m.Latitude = int32(binary.LittleEndian.Uint32(payload[0:]))
	m.Longitude = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.Altitude = int32(binary.LittleEndian.Uint32(payload[8:]))
	return nil
}

// ParamMapRc struct (generated typeinfo)
// Bind a RC channel to a parameter. The parameter should change according to the RC channel value.
type ParamMapRc struct {
	ParamValue0             float32 // Initial parameter value
	Scale                   float32 // Scale, maps the RC range [-1, 1] to a parameter value
	ParamValueMin           float32 // Minimum param value. The protocol does not define if this overwrites an onboard minimum value. (Depends on implementation)
	ParamValueMax           float32 // Maximum param value. The protocol does not define if this overwrites an onboard maximum value. (Depends on implementation)
	ParamIndex              int16   // Parameter index. Send -1 to use the param ID field as identifier (else the param id will be ignored), send -2 to disable any existing map for this rc_channel_index.
	TargetSystem            uint8   // System ID
	TargetComponent         uint8   // Component ID
	ParamID                 string  `len:"16" ` // Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParameterRcChannelIndex uint8   // Index of parameter RC channel. Not equal to the RC channel id. Typically corresponds to a potentiometer-knob on the RC.
}

// MsgID (generated function)
func (m *ParamMapRc) MsgID() message.MessageID {
	return MSG_ID_PARAM_MAP_RC
}

// String (generated function)
func (m *ParamMapRc) String() string {
	return fmt.Sprintf(
		"&common.ParamMapRc{ ParamValue0: %+v, Scale: %+v, ParamValueMin: %+v, ParamValueMax: %+v, ParamIndex: %+v, TargetSystem: %+v, TargetComponent: %+v, ParamID: \"%s\", ParameterRcChannelIndex: %+v }",
		m.ParamValue0,
		m.Scale,
		m.ParamValueMin,
		m.ParamValueMax,
		m.ParamIndex,
		m.TargetSystem,
		m.TargetComponent,
		m.ParamID,
		m.ParameterRcChannelIndex,
	)
}

// Marshal (generated function)
func (m *ParamMapRc) Marshal() ([]byte, error) {
	payload := make([]byte, 37)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.ParamValue0))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Scale))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.ParamValueMin))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.ParamValueMax))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.ParamIndex))
	payload[18] = byte(m.TargetSystem)
	payload[19] = byte(m.TargetComponent)
	copy(payload[20:], m.ParamID)
	payload[36] = byte(m.ParameterRcChannelIndex)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ParamMapRc) Unmarshal(data []byte) error {
	payload := make([]byte, 37)
	copy(payload[0:], data)
	m.ParamValue0 = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.Scale = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.ParamValueMin = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.ParamValueMax = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.ParamIndex = int16(binary.LittleEndian.Uint16(payload[16:]))
	m.TargetSystem = uint8(payload[18])
	m.TargetComponent = uint8(payload[19])
	m.ParamID = strings.TrimRight(string(payload[20:36]), string(byte(0)))
	m.ParameterRcChannelIndex = uint8(payload[36])
	return nil
}

// MissionRequestInt struct (generated typeinfo)
// Request the information of the mission item with the sequence number seq. The response of the system to this message should be a MISSION_ITEM_INT message. https://mavlink.io/en/services/mission.html
type MissionRequestInt struct {
	Seq             uint16 // Sequence
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
}

// MsgID (generated function)
func (m *MissionRequestInt) MsgID() message.MessageID {
	return MSG_ID_MISSION_REQUEST_INT
}

// String (generated function)
func (m *MissionRequestInt) String() string {
	return fmt.Sprintf(
		"&common.MissionRequestInt{ Seq: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.Seq,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *MissionRequestInt) Marshal() ([]byte, error) {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Seq))
	payload[2] = byte(m.TargetSystem)
	payload[3] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionRequestInt) Unmarshal(data []byte) error {
	payload := make([]byte, 4)
	copy(payload[0:], data)
	m.Seq = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetSystem = uint8(payload[2])
	m.TargetComponent = uint8(payload[3])
	return nil
}

// MissionChanged struct (generated typeinfo)
// A broadcast message to notify any ground station or SDK if a mission, geofence or safe points have changed on the vehicle.
type MissionChanged struct {
	StartIndex   int16            // Start index for partial mission change (-1 for all items).
	EndIndex     int16            // End index of a partial mission change. -1 is a synonym for the last mission item (i.e. selects all items from start_index). Ignore field if start_index=-1.
	OriginSysid  uint8            // System ID of the author of the new mission.
	OriginCompid MAV_COMPONENT    // Compnent ID of the author of the new mission.
	MissionType  MAV_MISSION_TYPE // Mission type.
}

// MsgID (generated function)
func (m *MissionChanged) MsgID() message.MessageID {
	return MSG_ID_MISSION_CHANGED
}

// String (generated function)
func (m *MissionChanged) String() string {
	return fmt.Sprintf(
		"&common.MissionChanged{ StartIndex: %+v, EndIndex: %+v, OriginSysid: %+v, OriginCompid: %+v, MissionType: %+v }",
		m.StartIndex,
		m.EndIndex,
		m.OriginSysid,
		m.OriginCompid,
		m.MissionType,
	)
}

// Marshal (generated function)
func (m *MissionChanged) Marshal() ([]byte, error) {
	payload := make([]byte, 7)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.StartIndex))
	binary.LittleEndian.PutUint16(payload[2:], uint16(m.EndIndex))
	payload[4] = byte(m.OriginSysid)
	payload[5] = byte(m.OriginCompid)
	payload[6] = byte(m.MissionType)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionChanged) Unmarshal(data []byte) error {
	payload := make([]byte, 7)
	copy(payload[0:], data)
	m.StartIndex = int16(binary.LittleEndian.Uint16(payload[0:]))
	m.EndIndex = int16(binary.LittleEndian.Uint16(payload[2:]))
	m.OriginSysid = uint8(payload[4])
	m.OriginCompid = MAV_COMPONENT(payload[5])
	m.MissionType = MAV_MISSION_TYPE(payload[6])
	return nil
}

// SafetySetAllowedArea struct (generated typeinfo)
// Set a safety zone (volume), which is defined by two corners of a cube. This message can be used to tell the MAV which setpoints/waypoints to accept and which to reject. Safety areas are often enforced by national or competition regulations.
type SafetySetAllowedArea struct {
	P1x             float32   // x position 1 / Latitude 1
	P1y             float32   // y position 1 / Longitude 1
	P1z             float32   // z position 1 / Altitude 1
	P2x             float32   // x position 2 / Latitude 2
	P2y             float32   // y position 2 / Longitude 2
	P2z             float32   // z position 2 / Altitude 2
	TargetSystem    uint8     // System ID
	TargetComponent uint8     // Component ID
	Frame           MAV_FRAME // Coordinate frame. Can be either global, GPS, right-handed with Z axis up or local, right handed, Z axis down.
}

// MsgID (generated function)
func (m *SafetySetAllowedArea) MsgID() message.MessageID {
	return MSG_ID_SAFETY_SET_ALLOWED_AREA
}

// String (generated function)
func (m *SafetySetAllowedArea) String() string {
	return fmt.Sprintf(
		"&common.SafetySetAllowedArea{ P1x: %+v, P1y: %+v, P1z: %+v, P2x: %+v, P2y: %+v, P2z: %+v, TargetSystem: %+v, TargetComponent: %+v, Frame: %+v }",
		m.P1x,
		m.P1y,
		m.P1z,
		m.P2x,
		m.P2y,
		m.P2z,
		m.TargetSystem,
		m.TargetComponent,
		m.Frame,
	)
}

// Marshal (generated function)
func (m *SafetySetAllowedArea) Marshal() ([]byte, error) {
	payload := make([]byte, 27)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.P1x))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.P1y))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.P1z))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.P2x))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.P2y))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.P2z))
	payload[24] = byte(m.TargetSystem)
	payload[25] = byte(m.TargetComponent)
	payload[26] = byte(m.Frame)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SafetySetAllowedArea) Unmarshal(data []byte) error {
	payload := make([]byte, 27)
	copy(payload[0:], data)
	m.P1x = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.P1y = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.P1z = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.P2x = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.P2y = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.P2z = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.TargetSystem = uint8(payload[24])
	m.TargetComponent = uint8(payload[25])
	m.Frame = MAV_FRAME(payload[26])
	return nil
}

// SafetyAllowedArea struct (generated typeinfo)
// Read out the safety zone the MAV currently assumes.
type SafetyAllowedArea struct {
	P1x   float32   // x position 1 / Latitude 1
	P1y   float32   // y position 1 / Longitude 1
	P1z   float32   // z position 1 / Altitude 1
	P2x   float32   // x position 2 / Latitude 2
	P2y   float32   // y position 2 / Longitude 2
	P2z   float32   // z position 2 / Altitude 2
	Frame MAV_FRAME // Coordinate frame. Can be either global, GPS, right-handed with Z axis up or local, right handed, Z axis down.
}

// MsgID (generated function)
func (m *SafetyAllowedArea) MsgID() message.MessageID {
	return MSG_ID_SAFETY_ALLOWED_AREA
}

// String (generated function)
func (m *SafetyAllowedArea) String() string {
	return fmt.Sprintf(
		"&common.SafetyAllowedArea{ P1x: %+v, P1y: %+v, P1z: %+v, P2x: %+v, P2y: %+v, P2z: %+v, Frame: %+v }",
		m.P1x,
		m.P1y,
		m.P1z,
		m.P2x,
		m.P2y,
		m.P2z,
		m.Frame,
	)
}

// Marshal (generated function)
func (m *SafetyAllowedArea) Marshal() ([]byte, error) {
	payload := make([]byte, 25)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.P1x))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.P1y))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.P1z))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.P2x))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.P2y))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.P2z))
	payload[24] = byte(m.Frame)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SafetyAllowedArea) Unmarshal(data []byte) error {
	payload := make([]byte, 25)
	copy(payload[0:], data)
	m.P1x = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.P1y = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.P1z = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.P2x = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.P2y = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.P2z = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Frame = MAV_FRAME(payload[24])
	return nil
}

// AttitudeQuaternionCov struct (generated typeinfo)
// The attitude in the aeronautical frame (right-handed, Z-down, X-front, Y-right), expressed as quaternion. Quaternion order is w, x, y, z and a zero rotation would be expressed as (1 0 0 0).
type AttitudeQuaternionCov struct {
	TimeUsec   uint64    // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Q          []float32 `len:"4" ` // Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation)
	Rollspeed  float32   // Roll angular speed
	Pitchspeed float32   // Pitch angular speed
	Yawspeed   float32   // Yaw angular speed
	Covariance []float32 `len:"9" ` // Row-major representation of a 3x3 attitude covariance matrix (states: roll, pitch, yaw; first three entries are the first ROW, next three entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
}

// MsgID (generated function)
func (m *AttitudeQuaternionCov) MsgID() message.MessageID {
	return MSG_ID_ATTITUDE_QUATERNION_COV
}

// String (generated function)
func (m *AttitudeQuaternionCov) String() string {
	return fmt.Sprintf(
		"&common.AttitudeQuaternionCov{ TimeUsec: %+v, Q: %+v, Rollspeed: %+v, Pitchspeed: %+v, Yawspeed: %+v, Covariance: %+v }",
		m.TimeUsec,
		m.Q,
		m.Rollspeed,
		m.Pitchspeed,
		m.Yawspeed,
		m.Covariance,
	)
}

// Marshal (generated function)
func (m *AttitudeQuaternionCov) Marshal() ([]byte, error) {
	payload := make([]byte, 72)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	for i, v := range m.Q {
		binary.LittleEndian.PutUint32(payload[8+i*4:], math.Float32bits(v))
	}
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Rollspeed))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Pitchspeed))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Yawspeed))
	for i, v := range m.Covariance {
		binary.LittleEndian.PutUint32(payload[36+i*4:], math.Float32bits(v))
	}
	return payload, nil
}

// Unmarshal (generated function)
func (m *AttitudeQuaternionCov) Unmarshal(data []byte) error {
	payload := make([]byte, 72)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	for i := 0; i < len(m.Q); i++ {
		m.Q[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[8+i*4:]))
	}
	m.Rollspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Pitchspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Yawspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	for i := 0; i < len(m.Covariance); i++ {
		m.Covariance[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[36+i*4:]))
	}
	return nil
}

// NavControllerOutput struct (generated typeinfo)
// The state of the fixed wing navigation and position controller.
type NavControllerOutput struct {
	NavRoll       float32 // Current desired roll
	NavPitch      float32 // Current desired pitch
	AltError      float32 // Current altitude error
	AspdError     float32 // Current airspeed error
	XtrackError   float32 // Current crosstrack error on x-y plane
	NavBearing    int16   // Current desired heading
	TargetBearing int16   // Bearing to current waypoint/target
	WpDist        uint16  // Distance to active waypoint
}

// MsgID (generated function)
func (m *NavControllerOutput) MsgID() message.MessageID {
	return MSG_ID_NAV_CONTROLLER_OUTPUT
}

// String (generated function)
func (m *NavControllerOutput) String() string {
	return fmt.Sprintf(
		"&common.NavControllerOutput{ NavRoll: %+v, NavPitch: %+v, AltError: %+v, AspdError: %+v, XtrackError: %+v, NavBearing: %+v, TargetBearing: %+v, WpDist: %+v }",
		m.NavRoll,
		m.NavPitch,
		m.AltError,
		m.AspdError,
		m.XtrackError,
		m.NavBearing,
		m.TargetBearing,
		m.WpDist,
	)
}

// Marshal (generated function)
func (m *NavControllerOutput) Marshal() ([]byte, error) {
	payload := make([]byte, 26)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.NavRoll))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.NavPitch))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.AltError))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.AspdError))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.XtrackError))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.NavBearing))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.TargetBearing))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.WpDist))
	return payload, nil
}

// Unmarshal (generated function)
func (m *NavControllerOutput) Unmarshal(data []byte) error {
	payload := make([]byte, 26)
	copy(payload[0:], data)
	m.NavRoll = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.NavPitch = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.AltError = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.AspdError = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.XtrackError = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.NavBearing = int16(binary.LittleEndian.Uint16(payload[20:]))
	m.TargetBearing = int16(binary.LittleEndian.Uint16(payload[22:]))
	m.WpDist = uint16(binary.LittleEndian.Uint16(payload[24:]))
	return nil
}

// GlobalPositionIntCov struct (generated typeinfo)
// The filtered global position (e.g. fused GPS and accelerometers). The position is in GPS-frame (right-handed, Z-up). It  is designed as scaled integer message since the resolution of float is not sufficient. NOTE: This message is intended for onboard networks / companion computers and higher-bandwidth links and optimized for accuracy and completeness. Please use the GLOBAL_POSITION_INT message for a minimal subset.
type GlobalPositionIntCov struct {
	TimeUsec      uint64             // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Lat           int32              // Latitude
	Lon           int32              // Longitude
	Alt           int32              // Altitude in meters above MSL
	RelativeAlt   int32              // Altitude above ground
	Vx            float32            // Ground X Speed (Latitude)
	Vy            float32            // Ground Y Speed (Longitude)
	Vz            float32            // Ground Z Speed (Altitude)
	Covariance    []float32          `len:"36" ` // Row-major representation of a 6x6 position and velocity 6x6 cross-covariance matrix (states: lat, lon, alt, vx, vy, vz; first six entries are the first ROW, next six entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
	EstimatorType MAV_ESTIMATOR_TYPE // Class id of the estimator this estimate originated from.
}

// MsgID (generated function)
func (m *GlobalPositionIntCov) MsgID() message.MessageID {
	return MSG_ID_GLOBAL_POSITION_INT_COV
}

// String (generated function)
func (m *GlobalPositionIntCov) String() string {
	return fmt.Sprintf(
		"&common.GlobalPositionIntCov{ TimeUsec: %+v, Lat: %+v, Lon: %+v, Alt: %+v, RelativeAlt: %+v, Vx: %+v, Vy: %+v, Vz: %+v, Covariance: %+v, EstimatorType: %+v }",
		m.TimeUsec,
		m.Lat,
		m.Lon,
		m.Alt,
		m.RelativeAlt,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Covariance,
		m.EstimatorType,
	)
}

// Marshal (generated function)
func (m *GlobalPositionIntCov) Marshal() ([]byte, error) {
	payload := make([]byte, 181)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.Alt))
	binary.LittleEndian.PutUint32(payload[20:], uint32(m.RelativeAlt))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Vx))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Vy))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Vz))
	for i, v := range m.Covariance {
		binary.LittleEndian.PutUint32(payload[36+i*4:], math.Float32bits(v))
	}
	payload[180] = byte(m.EstimatorType)
	return payload, nil
}

// Unmarshal (generated function)
func (m *GlobalPositionIntCov) Unmarshal(data []byte) error {
	payload := make([]byte, 181)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[12:]))
	m.Alt = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.RelativeAlt = int32(binary.LittleEndian.Uint32(payload[20:]))
	m.Vx = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Vy = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Vz = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	for i := 0; i < len(m.Covariance); i++ {
		m.Covariance[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[36+i*4:]))
	}
	m.EstimatorType = MAV_ESTIMATOR_TYPE(payload[180])
	return nil
}

// LocalPositionNedCov struct (generated typeinfo)
// The filtered local position (e.g. fused computer vision and accelerometers). Coordinate frame is right-handed, Z-axis down (aeronautical frame, NED / north-east-down convention)
type LocalPositionNedCov struct {
	TimeUsec      uint64             // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	X             float32            // X Position
	Y             float32            // Y Position
	Z             float32            // Z Position
	Vx            float32            // X Speed
	Vy            float32            // Y Speed
	Vz            float32            // Z Speed
	Ax            float32            // X Acceleration
	Ay            float32            // Y Acceleration
	Az            float32            // Z Acceleration
	Covariance    []float32          `len:"45" ` // Row-major representation of position, velocity and acceleration 9x9 cross-covariance matrix upper right triangle (states: x, y, z, vx, vy, vz, ax, ay, az; first nine entries are the first ROW, next eight entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
	EstimatorType MAV_ESTIMATOR_TYPE // Class id of the estimator this estimate originated from.
}

// MsgID (generated function)
func (m *LocalPositionNedCov) MsgID() message.MessageID {
	return MSG_ID_LOCAL_POSITION_NED_COV
}

// String (generated function)
func (m *LocalPositionNedCov) String() string {
	return fmt.Sprintf(
		"&common.LocalPositionNedCov{ TimeUsec: %+v, X: %+v, Y: %+v, Z: %+v, Vx: %+v, Vy: %+v, Vz: %+v, Ax: %+v, Ay: %+v, Az: %+v, Covariance: %+v, EstimatorType: %+v }",
		m.TimeUsec,
		m.X,
		m.Y,
		m.Z,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Ax,
		m.Ay,
		m.Az,
		m.Covariance,
		m.EstimatorType,
	)
}

// Marshal (generated function)
func (m *LocalPositionNedCov) Marshal() ([]byte, error) {
	payload := make([]byte, 225)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Vx))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Vy))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Vz))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Ax))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Ay))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Az))
	for i, v := range m.Covariance {
		binary.LittleEndian.PutUint32(payload[44+i*4:], math.Float32bits(v))
	}
	payload[224] = byte(m.EstimatorType)
	return payload, nil
}

// Unmarshal (generated function)
func (m *LocalPositionNedCov) Unmarshal(data []byte) error {
	payload := make([]byte, 225)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Vx = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Vy = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Vz = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Ax = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Ay = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Az = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	for i := 0; i < len(m.Covariance); i++ {
		m.Covariance[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[44+i*4:]))
	}
	m.EstimatorType = MAV_ESTIMATOR_TYPE(payload[224])
	return nil
}

// RcChannels struct (generated typeinfo)
// The PPM values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%.  A value of UINT16_MAX implies the channel is unused. Individual receivers/transmitters might violate this specification.
type RcChannels struct {
	TimeBootMs uint32 // Timestamp (time since system boot).
	Chan1Raw   uint16 // RC channel 1 value.
	Chan2Raw   uint16 // RC channel 2 value.
	Chan3Raw   uint16 // RC channel 3 value.
	Chan4Raw   uint16 // RC channel 4 value.
	Chan5Raw   uint16 // RC channel 5 value.
	Chan6Raw   uint16 // RC channel 6 value.
	Chan7Raw   uint16 // RC channel 7 value.
	Chan8Raw   uint16 // RC channel 8 value.
	Chan9Raw   uint16 // RC channel 9 value.
	Chan10Raw  uint16 // RC channel 10 value.
	Chan11Raw  uint16 // RC channel 11 value.
	Chan12Raw  uint16 // RC channel 12 value.
	Chan13Raw  uint16 // RC channel 13 value.
	Chan14Raw  uint16 // RC channel 14 value.
	Chan15Raw  uint16 // RC channel 15 value.
	Chan16Raw  uint16 // RC channel 16 value.
	Chan17Raw  uint16 // RC channel 17 value.
	Chan18Raw  uint16 // RC channel 18 value.
	Chancount  uint8  // Total number of RC channels being received. This can be larger than 18, indicating that more channels are available but not given in this message. This value should be 0 when no RC channels are available.
	Rssi       uint8  // Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
}

// MsgID (generated function)
func (m *RcChannels) MsgID() message.MessageID {
	return MSG_ID_RC_CHANNELS
}

// String (generated function)
func (m *RcChannels) String() string {
	return fmt.Sprintf(
		"&common.RcChannels{ TimeBootMs: %+v, Chan1Raw: %+v, Chan2Raw: %+v, Chan3Raw: %+v, Chan4Raw: %+v, Chan5Raw: %+v, Chan6Raw: %+v, Chan7Raw: %+v, Chan8Raw: %+v, Chan9Raw: %+v, Chan10Raw: %+v, Chan11Raw: %+v, Chan12Raw: %+v, Chan13Raw: %+v, Chan14Raw: %+v, Chan15Raw: %+v, Chan16Raw: %+v, Chan17Raw: %+v, Chan18Raw: %+v, Chancount: %+v, Rssi: %+v }",
		m.TimeBootMs,
		m.Chan1Raw,
		m.Chan2Raw,
		m.Chan3Raw,
		m.Chan4Raw,
		m.Chan5Raw,
		m.Chan6Raw,
		m.Chan7Raw,
		m.Chan8Raw,
		m.Chan9Raw,
		m.Chan10Raw,
		m.Chan11Raw,
		m.Chan12Raw,
		m.Chan13Raw,
		m.Chan14Raw,
		m.Chan15Raw,
		m.Chan16Raw,
		m.Chan17Raw,
		m.Chan18Raw,
		m.Chancount,
		m.Rssi,
	)
}

// Marshal (generated function)
func (m *RcChannels) Marshal() ([]byte, error) {
	payload := make([]byte, 42)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Chan1Raw))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.Chan2Raw))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Chan3Raw))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.Chan4Raw))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Chan5Raw))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Chan6Raw))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Chan7Raw))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Chan8Raw))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Chan9Raw))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.Chan10Raw))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.Chan11Raw))
	binary.LittleEndian.PutUint16(payload[26:], uint16(m.Chan12Raw))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Chan13Raw))
	binary.LittleEndian.PutUint16(payload[30:], uint16(m.Chan14Raw))
	binary.LittleEndian.PutUint16(payload[32:], uint16(m.Chan15Raw))
	binary.LittleEndian.PutUint16(payload[34:], uint16(m.Chan16Raw))
	binary.LittleEndian.PutUint16(payload[36:], uint16(m.Chan17Raw))
	binary.LittleEndian.PutUint16(payload[38:], uint16(m.Chan18Raw))
	payload[40] = byte(m.Chancount)
	payload[41] = byte(m.Rssi)
	return payload, nil
}

// Unmarshal (generated function)
func (m *RcChannels) Unmarshal(data []byte) error {
	payload := make([]byte, 42)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Chan1Raw = uint16(binary.LittleEndian.Uint16(payload[4:]))
	m.Chan2Raw = uint16(binary.LittleEndian.Uint16(payload[6:]))
	m.Chan3Raw = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.Chan4Raw = uint16(binary.LittleEndian.Uint16(payload[10:]))
	m.Chan5Raw = uint16(binary.LittleEndian.Uint16(payload[12:]))
	m.Chan6Raw = uint16(binary.LittleEndian.Uint16(payload[14:]))
	m.Chan7Raw = uint16(binary.LittleEndian.Uint16(payload[16:]))
	m.Chan8Raw = uint16(binary.LittleEndian.Uint16(payload[18:]))
	m.Chan9Raw = uint16(binary.LittleEndian.Uint16(payload[20:]))
	m.Chan10Raw = uint16(binary.LittleEndian.Uint16(payload[22:]))
	m.Chan11Raw = uint16(binary.LittleEndian.Uint16(payload[24:]))
	m.Chan12Raw = uint16(binary.LittleEndian.Uint16(payload[26:]))
	m.Chan13Raw = uint16(binary.LittleEndian.Uint16(payload[28:]))
	m.Chan14Raw = uint16(binary.LittleEndian.Uint16(payload[30:]))
	m.Chan15Raw = uint16(binary.LittleEndian.Uint16(payload[32:]))
	m.Chan16Raw = uint16(binary.LittleEndian.Uint16(payload[34:]))
	m.Chan17Raw = uint16(binary.LittleEndian.Uint16(payload[36:]))
	m.Chan18Raw = uint16(binary.LittleEndian.Uint16(payload[38:]))
	m.Chancount = uint8(payload[40])
	m.Rssi = uint8(payload[41])
	return nil
}

// RequestDataStream struct (generated typeinfo)
// Request a data stream.
type RequestDataStream struct {
	ReqMessageRate  uint16 // The requested message rate
	TargetSystem    uint8  // The target requested to send the message stream.
	TargetComponent uint8  // The target requested to send the message stream.
	ReqStreamID     uint8  // The ID of the requested data stream
	StartStop       uint8  // 1 to start sending, 0 to stop sending.
}

// MsgID (generated function)
func (m *RequestDataStream) MsgID() message.MessageID {
	return MSG_ID_REQUEST_DATA_STREAM
}

// String (generated function)
func (m *RequestDataStream) String() string {
	return fmt.Sprintf(
		"&common.RequestDataStream{ ReqMessageRate: %+v, TargetSystem: %+v, TargetComponent: %+v, ReqStreamID: %+v, StartStop: %+v }",
		m.ReqMessageRate,
		m.TargetSystem,
		m.TargetComponent,
		m.ReqStreamID,
		m.StartStop,
	)
}

// Marshal (generated function)
func (m *RequestDataStream) Marshal() ([]byte, error) {
	payload := make([]byte, 6)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.ReqMessageRate))
	payload[2] = byte(m.TargetSystem)
	payload[3] = byte(m.TargetComponent)
	payload[4] = byte(m.ReqStreamID)
	payload[5] = byte(m.StartStop)
	return payload, nil
}

// Unmarshal (generated function)
func (m *RequestDataStream) Unmarshal(data []byte) error {
	payload := make([]byte, 6)
	copy(payload[0:], data)
	m.ReqMessageRate = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetSystem = uint8(payload[2])
	m.TargetComponent = uint8(payload[3])
	m.ReqStreamID = uint8(payload[4])
	m.StartStop = uint8(payload[5])
	return nil
}

// DataStream struct (generated typeinfo)
// Data stream status information.
type DataStream struct {
	MessageRate uint16 // The message rate
	StreamID    uint8  // The ID of the requested data stream
	OnOff       uint8  // 1 stream is enabled, 0 stream is stopped.
}

// MsgID (generated function)
func (m *DataStream) MsgID() message.MessageID {
	return MSG_ID_DATA_STREAM
}

// String (generated function)
func (m *DataStream) String() string {
	return fmt.Sprintf(
		"&common.DataStream{ MessageRate: %+v, StreamID: %+v, OnOff: %+v }",
		m.MessageRate,
		m.StreamID,
		m.OnOff,
	)
}

// Marshal (generated function)
func (m *DataStream) Marshal() ([]byte, error) {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.MessageRate))
	payload[2] = byte(m.StreamID)
	payload[3] = byte(m.OnOff)
	return payload, nil
}

// Unmarshal (generated function)
func (m *DataStream) Unmarshal(data []byte) error {
	payload := make([]byte, 4)
	copy(payload[0:], data)
	m.MessageRate = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.StreamID = uint8(payload[2])
	m.OnOff = uint8(payload[3])
	return nil
}

// ManualControl struct (generated typeinfo)
// This message provides an API for manually controlling the vehicle using standard joystick axes nomenclature, along with a joystick-like input device. Unused axes can be disabled an buttons are also transmit as boolean values of their
type ManualControl struct {
	X       int16  // X-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to forward(1000)-backward(-1000) movement on a joystick and the pitch of a vehicle.
	Y       int16  // Y-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to left(-1000)-right(1000) movement on a joystick and the roll of a vehicle.
	Z       int16  // Z-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to a separate slider movement with maximum being 1000 and minimum being -1000 on a joystick and the thrust of a vehicle. Positive values are positive thrust, negative values are negative thrust.
	R       int16  // R-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to a twisting of the joystick, with counter-clockwise being 1000 and clockwise being -1000, and the yaw of a vehicle.
	Buttons uint16 // A bitfield corresponding to the joystick buttons' current state, 1 for pressed, 0 for released. The lowest bit corresponds to Button 1.
	Target  uint8  // The system to be controlled.
}

// MsgID (generated function)
func (m *ManualControl) MsgID() message.MessageID {
	return MSG_ID_MANUAL_CONTROL
}

// String (generated function)
func (m *ManualControl) String() string {
	return fmt.Sprintf(
		"&common.ManualControl{ X: %+v, Y: %+v, Z: %+v, R: %+v, Buttons: %+v, Target: %+v }",
		m.X,
		m.Y,
		m.Z,
		m.R,
		m.Buttons,
		m.Target,
	)
}

// Marshal (generated function)
func (m *ManualControl) Marshal() ([]byte, error) {
	payload := make([]byte, 11)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.X))
	binary.LittleEndian.PutUint16(payload[2:], uint16(m.Y))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Z))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.R))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Buttons))
	payload[10] = byte(m.Target)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ManualControl) Unmarshal(data []byte) error {
	payload := make([]byte, 11)
	copy(payload[0:], data)
	m.X = int16(binary.LittleEndian.Uint16(payload[0:]))
	m.Y = int16(binary.LittleEndian.Uint16(payload[2:]))
	m.Z = int16(binary.LittleEndian.Uint16(payload[4:]))
	m.R = int16(binary.LittleEndian.Uint16(payload[6:]))
	m.Buttons = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.Target = uint8(payload[10])
	return nil
}

// RcChannelsOverride struct (generated typeinfo)
// The RAW values of the RC channels sent to the MAV to override info received from the RC radio. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. Individual receivers/transmitters might violate this specification.  Note carefully the semantic differences between the first 8 channels and the subsequent channels
type RcChannelsOverride struct {
	Chan1Raw        uint16 // RC channel 1 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan2Raw        uint16 // RC channel 2 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan3Raw        uint16 // RC channel 3 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan4Raw        uint16 // RC channel 4 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan5Raw        uint16 // RC channel 5 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan6Raw        uint16 // RC channel 6 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan7Raw        uint16 // RC channel 7 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan8Raw        uint16 // RC channel 8 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
}

// MsgID (generated function)
func (m *RcChannelsOverride) MsgID() message.MessageID {
	return MSG_ID_RC_CHANNELS_OVERRIDE
}

// String (generated function)
func (m *RcChannelsOverride) String() string {
	return fmt.Sprintf(
		"&common.RcChannelsOverride{ Chan1Raw: %+v, Chan2Raw: %+v, Chan3Raw: %+v, Chan4Raw: %+v, Chan5Raw: %+v, Chan6Raw: %+v, Chan7Raw: %+v, Chan8Raw: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.Chan1Raw,
		m.Chan2Raw,
		m.Chan3Raw,
		m.Chan4Raw,
		m.Chan5Raw,
		m.Chan6Raw,
		m.Chan7Raw,
		m.Chan8Raw,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *RcChannelsOverride) Marshal() ([]byte, error) {
	payload := make([]byte, 18)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Chan1Raw))
	binary.LittleEndian.PutUint16(payload[2:], uint16(m.Chan2Raw))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Chan3Raw))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.Chan4Raw))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Chan5Raw))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.Chan6Raw))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Chan7Raw))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Chan8Raw))
	payload[16] = byte(m.TargetSystem)
	payload[17] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *RcChannelsOverride) Unmarshal(data []byte) error {
	payload := make([]byte, 18)
	copy(payload[0:], data)
	m.Chan1Raw = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.Chan2Raw = uint16(binary.LittleEndian.Uint16(payload[2:]))
	m.Chan3Raw = uint16(binary.LittleEndian.Uint16(payload[4:]))
	m.Chan4Raw = uint16(binary.LittleEndian.Uint16(payload[6:]))
	m.Chan5Raw = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.Chan6Raw = uint16(binary.LittleEndian.Uint16(payload[10:]))
	m.Chan7Raw = uint16(binary.LittleEndian.Uint16(payload[12:]))
	m.Chan8Raw = uint16(binary.LittleEndian.Uint16(payload[14:]))
	m.TargetSystem = uint8(payload[16])
	m.TargetComponent = uint8(payload[17])
	return nil
}

// MissionItemInt struct (generated typeinfo)
// Message encoding a mission item. This message is emitted to announce
//                 the presence of a mission item and to set a mission item on the system. The mission item can be either in x, y, z meters (type: LOCAL) or x:lat, y:lon, z:altitude. Local frame is Z-down, right handed (NED), global frame is Z-up, right handed (ENU). NaN or INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current latitude, yaw rather than a specific value). See also https://mavlink.io/en/services/mission.html.
type MissionItemInt struct {
	Param1          float32   // PARAM1, see MAV_CMD enum
	Param2          float32   // PARAM2, see MAV_CMD enum
	Param3          float32   // PARAM3, see MAV_CMD enum
	Param4          float32   // PARAM4, see MAV_CMD enum
	X               int32     // PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
	Y               int32     // PARAM6 / y position: local: x position in meters * 1e4, global: longitude in degrees *10^7
	Z               float32   // PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame.
	Seq             uint16    // Waypoint ID (sequence number). Starts at zero. Increases monotonically for each waypoint, no gaps in the sequence (0,1,2,3,4).
	Command         MAV_CMD   // The scheduled action for the waypoint.
	TargetSystem    uint8     // System ID
	TargetComponent uint8     // Component ID
	Frame           MAV_FRAME // The coordinate system of the waypoint.
	Current         uint8     // false:0, true:1
	Autocontinue    uint8     // Autocontinue to next waypoint
}

// MsgID (generated function)
func (m *MissionItemInt) MsgID() message.MessageID {
	return MSG_ID_MISSION_ITEM_INT
}

// String (generated function)
func (m *MissionItemInt) String() string {
	return fmt.Sprintf(
		"&common.MissionItemInt{ Param1: %+v, Param2: %+v, Param3: %+v, Param4: %+v, X: %+v, Y: %+v, Z: %+v, Seq: %+v, Command: %+v, TargetSystem: %+v, TargetComponent: %+v, Frame: %+v, Current: %+v, Autocontinue: %+v }",
		m.Param1,
		m.Param2,
		m.Param3,
		m.Param4,
		m.X,
		m.Y,
		m.Z,
		m.Seq,
		m.Command,
		m.TargetSystem,
		m.TargetComponent,
		m.Frame,
		m.Current,
		m.Autocontinue,
	)
}

// Marshal (generated function)
func (m *MissionItemInt) Marshal() ([]byte, error) {
	payload := make([]byte, 37)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.Param1))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Param2))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Param3))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Param4))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.X))
	binary.LittleEndian.PutUint32(payload[20:], uint32(m.Y))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Seq))
	binary.LittleEndian.PutUint16(payload[30:], uint16(m.Command))
	payload[32] = byte(m.TargetSystem)
	payload[33] = byte(m.TargetComponent)
	payload[34] = byte(m.Frame)
	payload[35] = byte(m.Current)
	payload[36] = byte(m.Autocontinue)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MissionItemInt) Unmarshal(data []byte) error {
	payload := make([]byte, 37)
	copy(payload[0:], data)
	m.Param1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.Param2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Param3 = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Param4 = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.X = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.Y = int32(binary.LittleEndian.Uint32(payload[20:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Seq = uint16(binary.LittleEndian.Uint16(payload[28:]))
	m.Command = MAV_CMD(binary.LittleEndian.Uint16(payload[30:]))
	m.TargetSystem = uint8(payload[32])
	m.TargetComponent = uint8(payload[33])
	m.Frame = MAV_FRAME(payload[34])
	m.Current = uint8(payload[35])
	m.Autocontinue = uint8(payload[36])
	return nil
}

// VfrHud struct (generated typeinfo)
// Metrics typically displayed on a HUD for fixed wing aircraft.
type VfrHud struct {
	Airspeed    float32 // Vehicle speed in form appropriate for vehicle type. For standard aircraft this is typically calibrated airspeed (CAS) or indicated airspeed (IAS) - either of which can be used by a pilot to estimate stall speed.
	Groundspeed float32 // Current ground speed.
	Alt         float32 // Current altitude (MSL).
	Climb       float32 // Current climb rate.
	Heading     int16   // Current heading in compass units (0-360, 0=north).
	Throttle    uint16  // Current throttle setting (0 to 100).
}

// MsgID (generated function)
func (m *VfrHud) MsgID() message.MessageID {
	return MSG_ID_VFR_HUD
}

// String (generated function)
func (m *VfrHud) String() string {
	return fmt.Sprintf(
		"&common.VfrHud{ Airspeed: %+v, Groundspeed: %+v, Alt: %+v, Climb: %+v, Heading: %+v, Throttle: %+v }",
		m.Airspeed,
		m.Groundspeed,
		m.Alt,
		m.Climb,
		m.Heading,
		m.Throttle,
	)
}

// Marshal (generated function)
func (m *VfrHud) Marshal() ([]byte, error) {
	payload := make([]byte, 20)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.Airspeed))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Groundspeed))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Alt))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Climb))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Heading))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Throttle))
	return payload, nil
}

// Unmarshal (generated function)
func (m *VfrHud) Unmarshal(data []byte) error {
	payload := make([]byte, 20)
	copy(payload[0:], data)
	m.Airspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.Groundspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Alt = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Climb = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Heading = int16(binary.LittleEndian.Uint16(payload[16:]))
	m.Throttle = uint16(binary.LittleEndian.Uint16(payload[18:]))
	return nil
}

// CommandInt struct (generated typeinfo)
// Message encoding a command with parameters as scaled integers. Scaling depends on the actual command value. The command microservice is documented at https://mavlink.io/en/services/command.html
type CommandInt struct {
	Param1          float32   // PARAM1, see MAV_CMD enum
	Param2          float32   // PARAM2, see MAV_CMD enum
	Param3          float32   // PARAM3, see MAV_CMD enum
	Param4          float32   // PARAM4, see MAV_CMD enum
	X               int32     // PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
	Y               int32     // PARAM6 / local: y position in meters * 1e4, global: longitude in degrees * 10^7
	Z               float32   // PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame).
	Command         MAV_CMD   // The scheduled action for the mission item.
	TargetSystem    uint8     // System ID
	TargetComponent uint8     // Component ID
	Frame           MAV_FRAME // The coordinate system of the COMMAND.
	Current         uint8     // Not used.
	Autocontinue    uint8     // Not used (set 0).
}

// MsgID (generated function)
func (m *CommandInt) MsgID() message.MessageID {
	return MSG_ID_COMMAND_INT
}

// String (generated function)
func (m *CommandInt) String() string {
	return fmt.Sprintf(
		"&common.CommandInt{ Param1: %+v, Param2: %+v, Param3: %+v, Param4: %+v, X: %+v, Y: %+v, Z: %+v, Command: %+v, TargetSystem: %+v, TargetComponent: %+v, Frame: %+v, Current: %+v, Autocontinue: %+v }",
		m.Param1,
		m.Param2,
		m.Param3,
		m.Param4,
		m.X,
		m.Y,
		m.Z,
		m.Command,
		m.TargetSystem,
		m.TargetComponent,
		m.Frame,
		m.Current,
		m.Autocontinue,
	)
}

// Marshal (generated function)
func (m *CommandInt) Marshal() ([]byte, error) {
	payload := make([]byte, 35)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.Param1))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Param2))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Param3))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Param4))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.X))
	binary.LittleEndian.PutUint32(payload[20:], uint32(m.Y))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Command))
	payload[30] = byte(m.TargetSystem)
	payload[31] = byte(m.TargetComponent)
	payload[32] = byte(m.Frame)
	payload[33] = byte(m.Current)
	payload[34] = byte(m.Autocontinue)
	return payload, nil
}

// Unmarshal (generated function)
func (m *CommandInt) Unmarshal(data []byte) error {
	payload := make([]byte, 35)
	copy(payload[0:], data)
	m.Param1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.Param2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Param3 = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Param4 = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.X = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.Y = int32(binary.LittleEndian.Uint32(payload[20:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Command = MAV_CMD(binary.LittleEndian.Uint16(payload[28:]))
	m.TargetSystem = uint8(payload[30])
	m.TargetComponent = uint8(payload[31])
	m.Frame = MAV_FRAME(payload[32])
	m.Current = uint8(payload[33])
	m.Autocontinue = uint8(payload[34])
	return nil
}

// CommandLong struct (generated typeinfo)
// Send a command with up to seven parameters to the MAV. The command microservice is documented at https://mavlink.io/en/services/command.html
type CommandLong struct {
	Param1          float32 // Parameter 1 (for the specific command).
	Param2          float32 // Parameter 2 (for the specific command).
	Param3          float32 // Parameter 3 (for the specific command).
	Param4          float32 // Parameter 4 (for the specific command).
	Param5          float32 // Parameter 5 (for the specific command).
	Param6          float32 // Parameter 6 (for the specific command).
	Param7          float32 // Parameter 7 (for the specific command).
	Command         MAV_CMD // Command ID (of command to send).
	TargetSystem    uint8   // System which should execute the command
	TargetComponent uint8   // Component which should execute the command, 0 for all components
	Confirmation    uint8   // 0: First transmission of this command. 1-255: Confirmation transmissions (e.g. for kill command)
}

// MsgID (generated function)
func (m *CommandLong) MsgID() message.MessageID {
	return MSG_ID_COMMAND_LONG
}

// String (generated function)
func (m *CommandLong) String() string {
	return fmt.Sprintf(
		"&common.CommandLong{ Param1: %+v, Param2: %+v, Param3: %+v, Param4: %+v, Param5: %+v, Param6: %+v, Param7: %+v, Command: %+v, TargetSystem: %+v, TargetComponent: %+v, Confirmation: %+v }",
		m.Param1,
		m.Param2,
		m.Param3,
		m.Param4,
		m.Param5,
		m.Param6,
		m.Param7,
		m.Command,
		m.TargetSystem,
		m.TargetComponent,
		m.Confirmation,
	)
}

// Marshal (generated function)
func (m *CommandLong) Marshal() ([]byte, error) {
	payload := make([]byte, 33)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.Param1))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Param2))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Param3))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Param4))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Param5))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Param6))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Param7))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Command))
	payload[30] = byte(m.TargetSystem)
	payload[31] = byte(m.TargetComponent)
	payload[32] = byte(m.Confirmation)
	return payload, nil
}

// Unmarshal (generated function)
func (m *CommandLong) Unmarshal(data []byte) error {
	payload := make([]byte, 33)
	copy(payload[0:], data)
	m.Param1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.Param2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Param3 = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Param4 = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Param5 = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Param6 = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Param7 = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Command = MAV_CMD(binary.LittleEndian.Uint16(payload[28:]))
	m.TargetSystem = uint8(payload[30])
	m.TargetComponent = uint8(payload[31])
	m.Confirmation = uint8(payload[32])
	return nil
}

// CommandAck struct (generated typeinfo)
// Report status of a command. Includes feedback whether the command was executed. The command microservice is documented at https://mavlink.io/en/services/command.html
type CommandAck struct {
	Command MAV_CMD    // Command ID (of acknowledged command).
	Result  MAV_RESULT // Result of command.
}

// MsgID (generated function)
func (m *CommandAck) MsgID() message.MessageID {
	return MSG_ID_COMMAND_ACK
}

// String (generated function)
func (m *CommandAck) String() string {
	return fmt.Sprintf(
		"&common.CommandAck{ Command: %+v, Result: %+v }",
		m.Command,
		m.Result,
	)
}

// Marshal (generated function)
func (m *CommandAck) Marshal() ([]byte, error) {
	payload := make([]byte, 3)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Command))
	payload[2] = byte(m.Result)
	return payload, nil
}

// Unmarshal (generated function)
func (m *CommandAck) Unmarshal(data []byte) error {
	payload := make([]byte, 3)
	copy(payload[0:], data)
	m.Command = MAV_CMD(binary.LittleEndian.Uint16(payload[0:]))
	m.Result = MAV_RESULT(payload[2])
	return nil
}

// CommandCancel struct (generated typeinfo)
// Cancel a long running command. The target system should respond with a COMMAND_ACK to the original command with result=MAV_RESULT_CANCELLED if the long running process was cancelled. If it has already completed, the cancel action can be ignored. The cancel action can be retried until some sort of acknowledgement to the original command has been received. The command microservice is documented at https://mavlink.io/en/services/command.html
type CommandCancel struct {
	Command         MAV_CMD // Command ID (of command to cancel).
	TargetSystem    uint8   // System executing long running command. Should not be broadcast (0).
	TargetComponent uint8   // Component executing long running command.
}

// MsgID (generated function)
func (m *CommandCancel) MsgID() message.MessageID {
	return MSG_ID_COMMAND_CANCEL
}

// String (generated function)
func (m *CommandCancel) String() string {
	return fmt.Sprintf(
		"&common.CommandCancel{ Command: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.Command,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *CommandCancel) Marshal() ([]byte, error) {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Command))
	payload[2] = byte(m.TargetSystem)
	payload[3] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *CommandCancel) Unmarshal(data []byte) error {
	payload := make([]byte, 4)
	copy(payload[0:], data)
	m.Command = MAV_CMD(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetSystem = uint8(payload[2])
	m.TargetComponent = uint8(payload[3])
	return nil
}

// ManualSetpoint struct (generated typeinfo)
// Setpoint in roll, pitch, yaw and thrust from the operator
type ManualSetpoint struct {
	TimeBootMs           uint32  // Timestamp (time since system boot).
	Roll                 float32 // Desired roll rate
	Pitch                float32 // Desired pitch rate
	Yaw                  float32 // Desired yaw rate
	Thrust               float32 // Collective thrust, normalized to 0 .. 1
	ModeSwitch           uint8   // Flight mode switch position, 0.. 255
	ManualOverrideSwitch uint8   // Override mode switch position, 0.. 255
}

// MsgID (generated function)
func (m *ManualSetpoint) MsgID() message.MessageID {
	return MSG_ID_MANUAL_SETPOINT
}

// String (generated function)
func (m *ManualSetpoint) String() string {
	return fmt.Sprintf(
		"&common.ManualSetpoint{ TimeBootMs: %+v, Roll: %+v, Pitch: %+v, Yaw: %+v, Thrust: %+v, ModeSwitch: %+v, ManualOverrideSwitch: %+v }",
		m.TimeBootMs,
		m.Roll,
		m.Pitch,
		m.Yaw,
		m.Thrust,
		m.ModeSwitch,
		m.ManualOverrideSwitch,
	)
}

// Marshal (generated function)
func (m *ManualSetpoint) Marshal() ([]byte, error) {
	payload := make([]byte, 22)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Roll))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Pitch))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Yaw))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Thrust))
	payload[20] = byte(m.ModeSwitch)
	payload[21] = byte(m.ManualOverrideSwitch)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ManualSetpoint) Unmarshal(data []byte) error {
	payload := make([]byte, 22)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Roll = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Pitch = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Thrust = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.ModeSwitch = uint8(payload[20])
	m.ManualOverrideSwitch = uint8(payload[21])
	return nil
}

// SetAttitudeTarget struct (generated typeinfo)
// Sets a desired vehicle attitude. Used by an external controller to command the vehicle (manual controller or other system).
type SetAttitudeTarget struct {
	TimeBootMs      uint32                   // Timestamp (time since system boot).
	Q               []float32                `len:"4" ` // Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	BodyRollRate    float32                  // Body roll rate
	BodyPitchRate   float32                  // Body pitch rate
	BodyYawRate     float32                  // Body yaw rate
	Thrust          float32                  // Collective thrust, normalized to 0 .. 1 (-1 .. 1 for vehicles capable of reverse trust)
	TargetSystem    uint8                    // System ID
	TargetComponent uint8                    // Component ID
	TypeMask        ATTITUDE_TARGET_TYPEMASK // Bitmap to indicate which dimensions should be ignored by the vehicle.
}

// MsgID (generated function)
func (m *SetAttitudeTarget) MsgID() message.MessageID {
	return MSG_ID_SET_ATTITUDE_TARGET
}

// String (generated function)
func (m *SetAttitudeTarget) String() string {
	return fmt.Sprintf(
		"&common.SetAttitudeTarget{ TimeBootMs: %+v, Q: %+v, BodyRollRate: %+v, BodyPitchRate: %+v, BodyYawRate: %+v, Thrust: %+v, TargetSystem: %+v, TargetComponent: %+v, TypeMask: %+v (%08b) }",
		m.TimeBootMs,
		m.Q,
		m.BodyRollRate,
		m.BodyPitchRate,
		m.BodyYawRate,
		m.Thrust,
		m.TargetSystem,
		m.TargetComponent,
		m.TypeMask.Bitmask(), uint64(m.TypeMask),
	)
}

// Marshal (generated function)
func (m *SetAttitudeTarget) Marshal() ([]byte, error) {
	payload := make([]byte, 39)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	for i, v := range m.Q {
		binary.LittleEndian.PutUint32(payload[4+i*4:], math.Float32bits(v))
	}
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.BodyRollRate))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.BodyPitchRate))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.BodyYawRate))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Thrust))
	payload[36] = byte(m.TargetSystem)
	payload[37] = byte(m.TargetComponent)
	payload[38] = byte(m.TypeMask)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SetAttitudeTarget) Unmarshal(data []byte) error {
	payload := make([]byte, 39)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	for i := 0; i < len(m.Q); i++ {
		m.Q[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[4+i*4:]))
	}
	m.BodyRollRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.BodyPitchRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.BodyYawRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Thrust = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.TargetSystem = uint8(payload[36])
	m.TargetComponent = uint8(payload[37])
	m.TypeMask = ATTITUDE_TARGET_TYPEMASK(payload[38])
	return nil
}

// AttitudeTarget struct (generated typeinfo)
// Reports the current commanded attitude of the vehicle as specified by the autopilot. This should match the commands sent in a SET_ATTITUDE_TARGET message if the vehicle is being controlled this way.
type AttitudeTarget struct {
	TimeBootMs    uint32                   // Timestamp (time since system boot).
	Q             []float32                `len:"4" ` // Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	BodyRollRate  float32                  // Body roll rate
	BodyPitchRate float32                  // Body pitch rate
	BodyYawRate   float32                  // Body yaw rate
	Thrust        float32                  // Collective thrust, normalized to 0 .. 1 (-1 .. 1 for vehicles capable of reverse trust)
	TypeMask      ATTITUDE_TARGET_TYPEMASK // Bitmap to indicate which dimensions should be ignored by the vehicle.
}

// MsgID (generated function)
func (m *AttitudeTarget) MsgID() message.MessageID {
	return MSG_ID_ATTITUDE_TARGET
}

// String (generated function)
func (m *AttitudeTarget) String() string {
	return fmt.Sprintf(
		"&common.AttitudeTarget{ TimeBootMs: %+v, Q: %+v, BodyRollRate: %+v, BodyPitchRate: %+v, BodyYawRate: %+v, Thrust: %+v, TypeMask: %+v (%08b) }",
		m.TimeBootMs,
		m.Q,
		m.BodyRollRate,
		m.BodyPitchRate,
		m.BodyYawRate,
		m.Thrust,
		m.TypeMask.Bitmask(), uint64(m.TypeMask),
	)
}

// Marshal (generated function)
func (m *AttitudeTarget) Marshal() ([]byte, error) {
	payload := make([]byte, 37)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	for i, v := range m.Q {
		binary.LittleEndian.PutUint32(payload[4+i*4:], math.Float32bits(v))
	}
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.BodyRollRate))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.BodyPitchRate))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.BodyYawRate))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Thrust))
	payload[36] = byte(m.TypeMask)
	return payload, nil
}

// Unmarshal (generated function)
func (m *AttitudeTarget) Unmarshal(data []byte) error {
	payload := make([]byte, 37)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	for i := 0; i < len(m.Q); i++ {
		m.Q[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[4+i*4:]))
	}
	m.BodyRollRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.BodyPitchRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.BodyYawRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Thrust = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.TypeMask = ATTITUDE_TARGET_TYPEMASK(payload[36])
	return nil
}

// SetPositionTargetLocalNed struct (generated typeinfo)
// Sets a desired vehicle position in a local north-east-down coordinate frame. Used by an external controller to command the vehicle (manual controller or other system).
type SetPositionTargetLocalNed struct {
	TimeBootMs      uint32                   // Timestamp (time since system boot).
	X               float32                  // X Position in NED frame
	Y               float32                  // Y Position in NED frame
	Z               float32                  // Z Position in NED frame (note, altitude is negative in NED)
	Vx              float32                  // X velocity in NED frame
	Vy              float32                  // Y velocity in NED frame
	Vz              float32                  // Z velocity in NED frame
	Afx             float32                  // X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy             float32                  // Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz             float32                  // Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Yaw             float32                  // yaw setpoint
	YawRate         float32                  // yaw rate setpoint
	TypeMask        POSITION_TARGET_TYPEMASK // Bitmap to indicate which dimensions should be ignored by the vehicle.
	TargetSystem    uint8                    // System ID
	TargetComponent uint8                    // Component ID
	CoordinateFrame MAV_FRAME                // Valid options are: MAV_FRAME_LOCAL_NED = 1, MAV_FRAME_LOCAL_OFFSET_NED = 7, MAV_FRAME_BODY_NED = 8, MAV_FRAME_BODY_OFFSET_NED = 9
}

// MsgID (generated function)
func (m *SetPositionTargetLocalNed) MsgID() message.MessageID {
	return MSG_ID_SET_POSITION_TARGET_LOCAL_NED
}

// String (generated function)
func (m *SetPositionTargetLocalNed) String() string {
	return fmt.Sprintf(
		"&common.SetPositionTargetLocalNed{ TimeBootMs: %+v, X: %+v, Y: %+v, Z: %+v, Vx: %+v, Vy: %+v, Vz: %+v, Afx: %+v, Afy: %+v, Afz: %+v, Yaw: %+v, YawRate: %+v, TypeMask: %+v (%016b), TargetSystem: %+v, TargetComponent: %+v, CoordinateFrame: %+v }",
		m.TimeBootMs,
		m.X,
		m.Y,
		m.Z,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Afx,
		m.Afy,
		m.Afz,
		m.Yaw,
		m.YawRate,
		m.TypeMask.Bitmask(), uint64(m.TypeMask),
		m.TargetSystem,
		m.TargetComponent,
		m.CoordinateFrame,
	)
}

// Marshal (generated function)
func (m *SetPositionTargetLocalNed) Marshal() ([]byte, error) {
	payload := make([]byte, 53)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Vx))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Vy))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Vz))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Afx))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Afy))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Afz))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Yaw))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.YawRate))
	binary.LittleEndian.PutUint16(payload[48:], uint16(m.TypeMask))
	payload[50] = byte(m.TargetSystem)
	payload[51] = byte(m.TargetComponent)
	payload[52] = byte(m.CoordinateFrame)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SetPositionTargetLocalNed) Unmarshal(data []byte) error {
	payload := make([]byte, 53)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Vx = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Vy = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Vz = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Afx = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Afy = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Afz = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.YawRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.TypeMask = POSITION_TARGET_TYPEMASK(binary.LittleEndian.Uint16(payload[48:]))
	m.TargetSystem = uint8(payload[50])
	m.TargetComponent = uint8(payload[51])
	m.CoordinateFrame = MAV_FRAME(payload[52])
	return nil
}

// PositionTargetLocalNed struct (generated typeinfo)
// Reports the current commanded vehicle position, velocity, and acceleration as specified by the autopilot. This should match the commands sent in SET_POSITION_TARGET_LOCAL_NED if the vehicle is being controlled this way.
type PositionTargetLocalNed struct {
	TimeBootMs      uint32                   // Timestamp (time since system boot).
	X               float32                  // X Position in NED frame
	Y               float32                  // Y Position in NED frame
	Z               float32                  // Z Position in NED frame (note, altitude is negative in NED)
	Vx              float32                  // X velocity in NED frame
	Vy              float32                  // Y velocity in NED frame
	Vz              float32                  // Z velocity in NED frame
	Afx             float32                  // X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy             float32                  // Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz             float32                  // Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Yaw             float32                  // yaw setpoint
	YawRate         float32                  // yaw rate setpoint
	TypeMask        POSITION_TARGET_TYPEMASK // Bitmap to indicate which dimensions should be ignored by the vehicle.
	CoordinateFrame MAV_FRAME                // Valid options are: MAV_FRAME_LOCAL_NED = 1, MAV_FRAME_LOCAL_OFFSET_NED = 7, MAV_FRAME_BODY_NED = 8, MAV_FRAME_BODY_OFFSET_NED = 9
}

// MsgID (generated function)
func (m *PositionTargetLocalNed) MsgID() message.MessageID {
	return MSG_ID_POSITION_TARGET_LOCAL_NED
}

// String (generated function)
func (m *PositionTargetLocalNed) String() string {
	return fmt.Sprintf(
		"&common.PositionTargetLocalNed{ TimeBootMs: %+v, X: %+v, Y: %+v, Z: %+v, Vx: %+v, Vy: %+v, Vz: %+v, Afx: %+v, Afy: %+v, Afz: %+v, Yaw: %+v, YawRate: %+v, TypeMask: %+v (%016b), CoordinateFrame: %+v }",
		m.TimeBootMs,
		m.X,
		m.Y,
		m.Z,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Afx,
		m.Afy,
		m.Afz,
		m.Yaw,
		m.YawRate,
		m.TypeMask.Bitmask(), uint64(m.TypeMask),
		m.CoordinateFrame,
	)
}

// Marshal (generated function)
func (m *PositionTargetLocalNed) Marshal() ([]byte, error) {
	payload := make([]byte, 51)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Vx))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Vy))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Vz))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Afx))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Afy))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Afz))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Yaw))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.YawRate))
	binary.LittleEndian.PutUint16(payload[48:], uint16(m.TypeMask))
	payload[50] = byte(m.CoordinateFrame)
	return payload, nil
}

// Unmarshal (generated function)
func (m *PositionTargetLocalNed) Unmarshal(data []byte) error {
	payload := make([]byte, 51)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Vx = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Vy = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Vz = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Afx = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Afy = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Afz = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.YawRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.TypeMask = POSITION_TARGET_TYPEMASK(binary.LittleEndian.Uint16(payload[48:]))
	m.CoordinateFrame = MAV_FRAME(payload[50])
	return nil
}

// SetPositionTargetGlobalInt struct (generated typeinfo)
// Sets a desired vehicle position, velocity, and/or acceleration in a global coordinate system (WGS84). Used by an external controller to command the vehicle (manual controller or other system).
type SetPositionTargetGlobalInt struct {
	TimeBootMs      uint32                   // Timestamp (time since system boot). The rationale for the timestamp in the setpoint is to allow the system to compensate for the transport delay of the setpoint. This allows the system to compensate processing latency.
	LatInt          int32                    // X Position in WGS84 frame
	LonInt          int32                    // Y Position in WGS84 frame
	Alt             float32                  // Altitude (MSL, Relative to home, or AGL - depending on frame)
	Vx              float32                  // X velocity in NED frame
	Vy              float32                  // Y velocity in NED frame
	Vz              float32                  // Z velocity in NED frame
	Afx             float32                  // X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy             float32                  // Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz             float32                  // Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Yaw             float32                  // yaw setpoint
	YawRate         float32                  // yaw rate setpoint
	TypeMask        POSITION_TARGET_TYPEMASK // Bitmap to indicate which dimensions should be ignored by the vehicle.
	TargetSystem    uint8                    // System ID
	TargetComponent uint8                    // Component ID
	CoordinateFrame MAV_FRAME                // Valid options are: MAV_FRAME_GLOBAL_INT = 5, MAV_FRAME_GLOBAL_RELATIVE_ALT_INT = 6, MAV_FRAME_GLOBAL_TERRAIN_ALT_INT = 11
}

// MsgID (generated function)
func (m *SetPositionTargetGlobalInt) MsgID() message.MessageID {
	return MSG_ID_SET_POSITION_TARGET_GLOBAL_INT
}

// String (generated function)
func (m *SetPositionTargetGlobalInt) String() string {
	return fmt.Sprintf(
		"&common.SetPositionTargetGlobalInt{ TimeBootMs: %+v, LatInt: %+v, LonInt: %+v, Alt: %+v, Vx: %+v, Vy: %+v, Vz: %+v, Afx: %+v, Afy: %+v, Afz: %+v, Yaw: %+v, YawRate: %+v, TypeMask: %+v (%016b), TargetSystem: %+v, TargetComponent: %+v, CoordinateFrame: %+v }",
		m.TimeBootMs,
		m.LatInt,
		m.LonInt,
		m.Alt,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Afx,
		m.Afy,
		m.Afz,
		m.Yaw,
		m.YawRate,
		m.TypeMask.Bitmask(), uint64(m.TypeMask),
		m.TargetSystem,
		m.TargetComponent,
		m.CoordinateFrame,
	)
}

// Marshal (generated function)
func (m *SetPositionTargetGlobalInt) Marshal() ([]byte, error) {
	payload := make([]byte, 53)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.LatInt))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.LonInt))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Alt))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Vx))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Vy))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Vz))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Afx))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Afy))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Afz))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Yaw))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.YawRate))
	binary.LittleEndian.PutUint16(payload[48:], uint16(m.TypeMask))
	payload[50] = byte(m.TargetSystem)
	payload[51] = byte(m.TargetComponent)
	payload[52] = byte(m.CoordinateFrame)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SetPositionTargetGlobalInt) Unmarshal(data []byte) error {
	payload := make([]byte, 53)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.LatInt = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.LonInt = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.Alt = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Vx = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Vy = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Vz = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Afx = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Afy = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Afz = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.YawRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.TypeMask = POSITION_TARGET_TYPEMASK(binary.LittleEndian.Uint16(payload[48:]))
	m.TargetSystem = uint8(payload[50])
	m.TargetComponent = uint8(payload[51])
	m.CoordinateFrame = MAV_FRAME(payload[52])
	return nil
}

// PositionTargetGlobalInt struct (generated typeinfo)
// Reports the current commanded vehicle position, velocity, and acceleration as specified by the autopilot. This should match the commands sent in SET_POSITION_TARGET_GLOBAL_INT if the vehicle is being controlled this way.
type PositionTargetGlobalInt struct {
	TimeBootMs      uint32                   // Timestamp (time since system boot). The rationale for the timestamp in the setpoint is to allow the system to compensate for the transport delay of the setpoint. This allows the system to compensate processing latency.
	LatInt          int32                    // X Position in WGS84 frame
	LonInt          int32                    // Y Position in WGS84 frame
	Alt             float32                  // Altitude (MSL, AGL or relative to home altitude, depending on frame)
	Vx              float32                  // X velocity in NED frame
	Vy              float32                  // Y velocity in NED frame
	Vz              float32                  // Z velocity in NED frame
	Afx             float32                  // X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy             float32                  // Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz             float32                  // Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Yaw             float32                  // yaw setpoint
	YawRate         float32                  // yaw rate setpoint
	TypeMask        POSITION_TARGET_TYPEMASK // Bitmap to indicate which dimensions should be ignored by the vehicle.
	CoordinateFrame MAV_FRAME                // Valid options are: MAV_FRAME_GLOBAL_INT = 5, MAV_FRAME_GLOBAL_RELATIVE_ALT_INT = 6, MAV_FRAME_GLOBAL_TERRAIN_ALT_INT = 11
}

// MsgID (generated function)
func (m *PositionTargetGlobalInt) MsgID() message.MessageID {
	return MSG_ID_POSITION_TARGET_GLOBAL_INT
}

// String (generated function)
func (m *PositionTargetGlobalInt) String() string {
	return fmt.Sprintf(
		"&common.PositionTargetGlobalInt{ TimeBootMs: %+v, LatInt: %+v, LonInt: %+v, Alt: %+v, Vx: %+v, Vy: %+v, Vz: %+v, Afx: %+v, Afy: %+v, Afz: %+v, Yaw: %+v, YawRate: %+v, TypeMask: %+v (%016b), CoordinateFrame: %+v }",
		m.TimeBootMs,
		m.LatInt,
		m.LonInt,
		m.Alt,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Afx,
		m.Afy,
		m.Afz,
		m.Yaw,
		m.YawRate,
		m.TypeMask.Bitmask(), uint64(m.TypeMask),
		m.CoordinateFrame,
	)
}

// Marshal (generated function)
func (m *PositionTargetGlobalInt) Marshal() ([]byte, error) {
	payload := make([]byte, 51)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.LatInt))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.LonInt))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Alt))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Vx))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Vy))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Vz))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Afx))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Afy))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Afz))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Yaw))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.YawRate))
	binary.LittleEndian.PutUint16(payload[48:], uint16(m.TypeMask))
	payload[50] = byte(m.CoordinateFrame)
	return payload, nil
}

// Unmarshal (generated function)
func (m *PositionTargetGlobalInt) Unmarshal(data []byte) error {
	payload := make([]byte, 51)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.LatInt = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.LonInt = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.Alt = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Vx = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Vy = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Vz = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Afx = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Afy = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Afz = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.YawRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.TypeMask = POSITION_TARGET_TYPEMASK(binary.LittleEndian.Uint16(payload[48:]))
	m.CoordinateFrame = MAV_FRAME(payload[50])
	return nil
}

// LocalPositionNedSystemGlobalOffset struct (generated typeinfo)
// The offset in X, Y, Z and yaw between the LOCAL_POSITION_NED messages of MAV X and the global coordinate frame in NED coordinates. Coordinate frame is right-handed, Z-axis down (aeronautical frame, NED / north-east-down convention)
type LocalPositionNedSystemGlobalOffset struct {
	TimeBootMs uint32  // Timestamp (time since system boot).
	X          float32 // X Position
	Y          float32 // Y Position
	Z          float32 // Z Position
	Roll       float32 // Roll
	Pitch      float32 // Pitch
	Yaw        float32 // Yaw
}

// MsgID (generated function)
func (m *LocalPositionNedSystemGlobalOffset) MsgID() message.MessageID {
	return MSG_ID_LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET
}

// String (generated function)
func (m *LocalPositionNedSystemGlobalOffset) String() string {
	return fmt.Sprintf(
		"&common.LocalPositionNedSystemGlobalOffset{ TimeBootMs: %+v, X: %+v, Y: %+v, Z: %+v, Roll: %+v, Pitch: %+v, Yaw: %+v }",
		m.TimeBootMs,
		m.X,
		m.Y,
		m.Z,
		m.Roll,
		m.Pitch,
		m.Yaw,
	)
}

// Marshal (generated function)
func (m *LocalPositionNedSystemGlobalOffset) Marshal() ([]byte, error) {
	payload := make([]byte, 28)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Roll))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Pitch))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Yaw))
	return payload, nil
}

// Unmarshal (generated function)
func (m *LocalPositionNedSystemGlobalOffset) Unmarshal(data []byte) error {
	payload := make([]byte, 28)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Roll = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Pitch = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	return nil
}

// HilState struct (generated typeinfo)
// Sent from simulation to autopilot. This packet is useful for high throughput applications such as hardware in the loop simulations.
type HilState struct {
	TimeUsec   uint64  // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Roll       float32 // Roll angle
	Pitch      float32 // Pitch angle
	Yaw        float32 // Yaw angle
	Rollspeed  float32 // Body frame roll / phi angular speed
	Pitchspeed float32 // Body frame pitch / theta angular speed
	Yawspeed   float32 // Body frame yaw / psi angular speed
	Lat        int32   // Latitude
	Lon        int32   // Longitude
	Alt        int32   // Altitude
	Vx         int16   // Ground X Speed (Latitude)
	Vy         int16   // Ground Y Speed (Longitude)
	Vz         int16   // Ground Z Speed (Altitude)
	Xacc       int16   // X acceleration
	Yacc       int16   // Y acceleration
	Zacc       int16   // Z acceleration
}

// MsgID (generated function)
func (m *HilState) MsgID() message.MessageID {
	return MSG_ID_HIL_STATE
}

// String (generated function)
func (m *HilState) String() string {
	return fmt.Sprintf(
		"&common.HilState{ TimeUsec: %+v, Roll: %+v, Pitch: %+v, Yaw: %+v, Rollspeed: %+v, Pitchspeed: %+v, Yawspeed: %+v, Lat: %+v, Lon: %+v, Alt: %+v, Vx: %+v, Vy: %+v, Vz: %+v, Xacc: %+v, Yacc: %+v, Zacc: %+v }",
		m.TimeUsec,
		m.Roll,
		m.Pitch,
		m.Yaw,
		m.Rollspeed,
		m.Pitchspeed,
		m.Yawspeed,
		m.Lat,
		m.Lon,
		m.Alt,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Xacc,
		m.Yacc,
		m.Zacc,
	)
}

// Marshal (generated function)
func (m *HilState) Marshal() ([]byte, error) {
	payload := make([]byte, 56)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Roll))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Pitch))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Yaw))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Rollspeed))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Pitchspeed))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Yawspeed))
	binary.LittleEndian.PutUint32(payload[32:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[36:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[40:], uint32(m.Alt))
	binary.LittleEndian.PutUint16(payload[44:], uint16(m.Vx))
	binary.LittleEndian.PutUint16(payload[46:], uint16(m.Vy))
	binary.LittleEndian.PutUint16(payload[48:], uint16(m.Vz))
	binary.LittleEndian.PutUint16(payload[50:], uint16(m.Xacc))
	binary.LittleEndian.PutUint16(payload[52:], uint16(m.Yacc))
	binary.LittleEndian.PutUint16(payload[54:], uint16(m.Zacc))
	return payload, nil
}

// Unmarshal (generated function)
func (m *HilState) Unmarshal(data []byte) error {
	payload := make([]byte, 56)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Roll = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Pitch = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Rollspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Pitchspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Yawspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[32:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[36:]))
	m.Alt = int32(binary.LittleEndian.Uint32(payload[40:]))
	m.Vx = int16(binary.LittleEndian.Uint16(payload[44:]))
	m.Vy = int16(binary.LittleEndian.Uint16(payload[46:]))
	m.Vz = int16(binary.LittleEndian.Uint16(payload[48:]))
	m.Xacc = int16(binary.LittleEndian.Uint16(payload[50:]))
	m.Yacc = int16(binary.LittleEndian.Uint16(payload[52:]))
	m.Zacc = int16(binary.LittleEndian.Uint16(payload[54:]))
	return nil
}

// HilControls struct (generated typeinfo)
// Sent from autopilot to simulation. Hardware in the loop control outputs
type HilControls struct {
	TimeUsec      uint64   // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	RollAilerons  float32  // Control output -1 .. 1
	PitchElevator float32  // Control output -1 .. 1
	YawRudder     float32  // Control output -1 .. 1
	Throttle      float32  // Throttle 0 .. 1
	Aux1          float32  // Aux 1, -1 .. 1
	Aux2          float32  // Aux 2, -1 .. 1
	Aux3          float32  // Aux 3, -1 .. 1
	Aux4          float32  // Aux 4, -1 .. 1
	Mode          MAV_MODE // System mode.
	NavMode       uint8    // Navigation mode (MAV_NAV_MODE)
}

// MsgID (generated function)
func (m *HilControls) MsgID() message.MessageID {
	return MSG_ID_HIL_CONTROLS
}

// String (generated function)
func (m *HilControls) String() string {
	return fmt.Sprintf(
		"&common.HilControls{ TimeUsec: %+v, RollAilerons: %+v, PitchElevator: %+v, YawRudder: %+v, Throttle: %+v, Aux1: %+v, Aux2: %+v, Aux3: %+v, Aux4: %+v, Mode: %+v, NavMode: %+v }",
		m.TimeUsec,
		m.RollAilerons,
		m.PitchElevator,
		m.YawRudder,
		m.Throttle,
		m.Aux1,
		m.Aux2,
		m.Aux3,
		m.Aux4,
		m.Mode,
		m.NavMode,
	)
}

// Marshal (generated function)
func (m *HilControls) Marshal() ([]byte, error) {
	payload := make([]byte, 42)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.RollAilerons))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.PitchElevator))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.YawRudder))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Throttle))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Aux1))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Aux2))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Aux3))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Aux4))
	payload[40] = byte(m.Mode)
	payload[41] = byte(m.NavMode)
	return payload, nil
}

// Unmarshal (generated function)
func (m *HilControls) Unmarshal(data []byte) error {
	payload := make([]byte, 42)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.RollAilerons = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.PitchElevator = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.YawRudder = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Throttle = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Aux1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Aux2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Aux3 = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Aux4 = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Mode = MAV_MODE(payload[40])
	m.NavMode = uint8(payload[41])
	return nil
}

// HilRcInputsRaw struct (generated typeinfo)
// Sent from simulation to autopilot. The RAW values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. Individual receivers/transmitters might violate this specification.
type HilRcInputsRaw struct {
	TimeUsec  uint64 // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Chan1Raw  uint16 // RC channel 1 value
	Chan2Raw  uint16 // RC channel 2 value
	Chan3Raw  uint16 // RC channel 3 value
	Chan4Raw  uint16 // RC channel 4 value
	Chan5Raw  uint16 // RC channel 5 value
	Chan6Raw  uint16 // RC channel 6 value
	Chan7Raw  uint16 // RC channel 7 value
	Chan8Raw  uint16 // RC channel 8 value
	Chan9Raw  uint16 // RC channel 9 value
	Chan10Raw uint16 // RC channel 10 value
	Chan11Raw uint16 // RC channel 11 value
	Chan12Raw uint16 // RC channel 12 value
	Rssi      uint8  // Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
}

// MsgID (generated function)
func (m *HilRcInputsRaw) MsgID() message.MessageID {
	return MSG_ID_HIL_RC_INPUTS_RAW
}

// String (generated function)
func (m *HilRcInputsRaw) String() string {
	return fmt.Sprintf(
		"&common.HilRcInputsRaw{ TimeUsec: %+v, Chan1Raw: %+v, Chan2Raw: %+v, Chan3Raw: %+v, Chan4Raw: %+v, Chan5Raw: %+v, Chan6Raw: %+v, Chan7Raw: %+v, Chan8Raw: %+v, Chan9Raw: %+v, Chan10Raw: %+v, Chan11Raw: %+v, Chan12Raw: %+v, Rssi: %+v }",
		m.TimeUsec,
		m.Chan1Raw,
		m.Chan2Raw,
		m.Chan3Raw,
		m.Chan4Raw,
		m.Chan5Raw,
		m.Chan6Raw,
		m.Chan7Raw,
		m.Chan8Raw,
		m.Chan9Raw,
		m.Chan10Raw,
		m.Chan11Raw,
		m.Chan12Raw,
		m.Rssi,
	)
}

// Marshal (generated function)
func (m *HilRcInputsRaw) Marshal() ([]byte, error) {
	payload := make([]byte, 33)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Chan1Raw))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.Chan2Raw))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Chan3Raw))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Chan4Raw))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Chan5Raw))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Chan6Raw))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Chan7Raw))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.Chan8Raw))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.Chan9Raw))
	binary.LittleEndian.PutUint16(payload[26:], uint16(m.Chan10Raw))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Chan11Raw))
	binary.LittleEndian.PutUint16(payload[30:], uint16(m.Chan12Raw))
	payload[32] = byte(m.Rssi)
	return payload, nil
}

// Unmarshal (generated function)
func (m *HilRcInputsRaw) Unmarshal(data []byte) error {
	payload := make([]byte, 33)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Chan1Raw = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.Chan2Raw = uint16(binary.LittleEndian.Uint16(payload[10:]))
	m.Chan3Raw = uint16(binary.LittleEndian.Uint16(payload[12:]))
	m.Chan4Raw = uint16(binary.LittleEndian.Uint16(payload[14:]))
	m.Chan5Raw = uint16(binary.LittleEndian.Uint16(payload[16:]))
	m.Chan6Raw = uint16(binary.LittleEndian.Uint16(payload[18:]))
	m.Chan7Raw = uint16(binary.LittleEndian.Uint16(payload[20:]))
	m.Chan8Raw = uint16(binary.LittleEndian.Uint16(payload[22:]))
	m.Chan9Raw = uint16(binary.LittleEndian.Uint16(payload[24:]))
	m.Chan10Raw = uint16(binary.LittleEndian.Uint16(payload[26:]))
	m.Chan11Raw = uint16(binary.LittleEndian.Uint16(payload[28:]))
	m.Chan12Raw = uint16(binary.LittleEndian.Uint16(payload[30:]))
	m.Rssi = uint8(payload[32])
	return nil
}

// HilActuatorControls struct (generated typeinfo)
// Sent from autopilot to simulation. Hardware in the loop control outputs (replacement for HIL_CONTROLS)
type HilActuatorControls struct {
	TimeUsec uint64        // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Flags    uint64        // Flags as bitfield, 1: indicate simulation using lockstep.
	Controls []float32     `len:"16" ` // Control outputs -1 .. 1. Channel assignment depends on the simulated hardware.
	Mode     MAV_MODE_FLAG // System mode. Includes arming state.
}

// MsgID (generated function)
func (m *HilActuatorControls) MsgID() message.MessageID {
	return MSG_ID_HIL_ACTUATOR_CONTROLS
}

// String (generated function)
func (m *HilActuatorControls) String() string {
	return fmt.Sprintf(
		"&common.HilActuatorControls{ TimeUsec: %+v, Flags: %+v, Controls: %+v, Mode: %+v (%08b) }",
		m.TimeUsec,
		m.Flags,
		m.Controls,
		m.Mode.Bitmask(), uint64(m.Mode),
	)
}

// Marshal (generated function)
func (m *HilActuatorControls) Marshal() ([]byte, error) {
	payload := make([]byte, 81)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint64(payload[8:], uint64(m.Flags))
	for i, v := range m.Controls {
		binary.LittleEndian.PutUint32(payload[16+i*4:], math.Float32bits(v))
	}
	payload[80] = byte(m.Mode)
	return payload, nil
}

// Unmarshal (generated function)
func (m *HilActuatorControls) Unmarshal(data []byte) error {
	payload := make([]byte, 81)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Flags = uint64(binary.LittleEndian.Uint64(payload[8:]))
	for i := 0; i < len(m.Controls); i++ {
		m.Controls[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[16+i*4:]))
	}
	m.Mode = MAV_MODE_FLAG(payload[80])
	return nil
}

// OpticalFlow struct (generated typeinfo)
// Optical flow from a flow sensor (e.g. optical mouse sensor)
type OpticalFlow struct {
	TimeUsec       uint64  // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	FlowCompMX     float32 // Flow in x-sensor direction, angular-speed compensated
	FlowCompMY     float32 // Flow in y-sensor direction, angular-speed compensated
	GroundDistance float32 // Ground distance. Positive value: distance known. Negative value: Unknown distance
	FlowX          int16   // Flow in x-sensor direction
	FlowY          int16   // Flow in y-sensor direction
	SensorID       uint8   // Sensor ID
	Quality        uint8   // Optical flow quality / confidence. 0: bad, 255: maximum quality
}

// MsgID (generated function)
func (m *OpticalFlow) MsgID() message.MessageID {
	return MSG_ID_OPTICAL_FLOW
}

// String (generated function)
func (m *OpticalFlow) String() string {
	return fmt.Sprintf(
		"&common.OpticalFlow{ TimeUsec: %+v, FlowCompMX: %+v, FlowCompMY: %+v, GroundDistance: %+v, FlowX: %+v, FlowY: %+v, SensorID: %+v, Quality: %+v }",
		m.TimeUsec,
		m.FlowCompMX,
		m.FlowCompMY,
		m.GroundDistance,
		m.FlowX,
		m.FlowY,
		m.SensorID,
		m.Quality,
	)
}

// Marshal (generated function)
func (m *OpticalFlow) Marshal() ([]byte, error) {
	payload := make([]byte, 26)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.FlowCompMX))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.FlowCompMY))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.GroundDistance))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.FlowX))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.FlowY))
	payload[24] = byte(m.SensorID)
	payload[25] = byte(m.Quality)
	return payload, nil
}

// Unmarshal (generated function)
func (m *OpticalFlow) Unmarshal(data []byte) error {
	payload := make([]byte, 26)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.FlowCompMX = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.FlowCompMY = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.GroundDistance = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.FlowX = int16(binary.LittleEndian.Uint16(payload[20:]))
	m.FlowY = int16(binary.LittleEndian.Uint16(payload[22:]))
	m.SensorID = uint8(payload[24])
	m.Quality = uint8(payload[25])
	return nil
}

// GlobalVisionPositionEstimate struct (generated typeinfo)
// Global position/attitude estimate from a vision source.
type GlobalVisionPositionEstimate struct {
	Usec  uint64  // Timestamp (UNIX time or since system boot)
	X     float32 // Global X position
	Y     float32 // Global Y position
	Z     float32 // Global Z position
	Roll  float32 // Roll angle
	Pitch float32 // Pitch angle
	Yaw   float32 // Yaw angle
}

// MsgID (generated function)
func (m *GlobalVisionPositionEstimate) MsgID() message.MessageID {
	return MSG_ID_GLOBAL_VISION_POSITION_ESTIMATE
}

// String (generated function)
func (m *GlobalVisionPositionEstimate) String() string {
	return fmt.Sprintf(
		"&common.GlobalVisionPositionEstimate{ Usec: %+v, X: %+v, Y: %+v, Z: %+v, Roll: %+v, Pitch: %+v, Yaw: %+v }",
		m.Usec,
		m.X,
		m.Y,
		m.Z,
		m.Roll,
		m.Pitch,
		m.Yaw,
	)
}

// Marshal (generated function)
func (m *GlobalVisionPositionEstimate) Marshal() ([]byte, error) {
	payload := make([]byte, 32)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Usec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Roll))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Pitch))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Yaw))
	return payload, nil
}

// Unmarshal (generated function)
func (m *GlobalVisionPositionEstimate) Unmarshal(data []byte) error {
	payload := make([]byte, 32)
	copy(payload[0:], data)
	m.Usec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Roll = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Pitch = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	return nil
}

// VisionPositionEstimate struct (generated typeinfo)
// Local position/attitude estimate from a vision source.
type VisionPositionEstimate struct {
	Usec  uint64  // Timestamp (UNIX time or time since system boot)
	X     float32 // Local X position
	Y     float32 // Local Y position
	Z     float32 // Local Z position
	Roll  float32 // Roll angle
	Pitch float32 // Pitch angle
	Yaw   float32 // Yaw angle
}

// MsgID (generated function)
func (m *VisionPositionEstimate) MsgID() message.MessageID {
	return MSG_ID_VISION_POSITION_ESTIMATE
}

// String (generated function)
func (m *VisionPositionEstimate) String() string {
	return fmt.Sprintf(
		"&common.VisionPositionEstimate{ Usec: %+v, X: %+v, Y: %+v, Z: %+v, Roll: %+v, Pitch: %+v, Yaw: %+v }",
		m.Usec,
		m.X,
		m.Y,
		m.Z,
		m.Roll,
		m.Pitch,
		m.Yaw,
	)
}

// Marshal (generated function)
func (m *VisionPositionEstimate) Marshal() ([]byte, error) {
	payload := make([]byte, 32)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Usec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Roll))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Pitch))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Yaw))
	return payload, nil
}

// Unmarshal (generated function)
func (m *VisionPositionEstimate) Unmarshal(data []byte) error {
	payload := make([]byte, 32)
	copy(payload[0:], data)
	m.Usec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Roll = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Pitch = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	return nil
}

// VisionSpeedEstimate struct (generated typeinfo)
// Speed estimate from a vision source.
type VisionSpeedEstimate struct {
	Usec uint64  // Timestamp (UNIX time or time since system boot)
	X    float32 // Global X speed
	Y    float32 // Global Y speed
	Z    float32 // Global Z speed
}

// MsgID (generated function)
func (m *VisionSpeedEstimate) MsgID() message.MessageID {
	return MSG_ID_VISION_SPEED_ESTIMATE
}

// String (generated function)
func (m *VisionSpeedEstimate) String() string {
	return fmt.Sprintf(
		"&common.VisionSpeedEstimate{ Usec: %+v, X: %+v, Y: %+v, Z: %+v }",
		m.Usec,
		m.X,
		m.Y,
		m.Z,
	)
}

// Marshal (generated function)
func (m *VisionSpeedEstimate) Marshal() ([]byte, error) {
	payload := make([]byte, 20)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Usec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Z))
	return payload, nil
}

// Unmarshal (generated function)
func (m *VisionSpeedEstimate) Unmarshal(data []byte) error {
	payload := make([]byte, 20)
	copy(payload[0:], data)
	m.Usec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	return nil
}

// ViconPositionEstimate struct (generated typeinfo)
// Global position estimate from a Vicon motion system source.
type ViconPositionEstimate struct {
	Usec  uint64  // Timestamp (UNIX time or time since system boot)
	X     float32 // Global X position
	Y     float32 // Global Y position
	Z     float32 // Global Z position
	Roll  float32 // Roll angle
	Pitch float32 // Pitch angle
	Yaw   float32 // Yaw angle
}

// MsgID (generated function)
func (m *ViconPositionEstimate) MsgID() message.MessageID {
	return MSG_ID_VICON_POSITION_ESTIMATE
}

// String (generated function)
func (m *ViconPositionEstimate) String() string {
	return fmt.Sprintf(
		"&common.ViconPositionEstimate{ Usec: %+v, X: %+v, Y: %+v, Z: %+v, Roll: %+v, Pitch: %+v, Yaw: %+v }",
		m.Usec,
		m.X,
		m.Y,
		m.Z,
		m.Roll,
		m.Pitch,
		m.Yaw,
	)
}

// Marshal (generated function)
func (m *ViconPositionEstimate) Marshal() ([]byte, error) {
	payload := make([]byte, 32)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Usec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Roll))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Pitch))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Yaw))
	return payload, nil
}

// Unmarshal (generated function)
func (m *ViconPositionEstimate) Unmarshal(data []byte) error {
	payload := make([]byte, 32)
	copy(payload[0:], data)
	m.Usec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Roll = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Pitch = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	return nil
}

// HighresImu struct (generated typeinfo)
// The IMU readings in SI units in NED body frame
type HighresImu struct {
	TimeUsec      uint64  // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Xacc          float32 // X acceleration
	Yacc          float32 // Y acceleration
	Zacc          float32 // Z acceleration
	Xgyro         float32 // Angular speed around X axis
	Ygyro         float32 // Angular speed around Y axis
	Zgyro         float32 // Angular speed around Z axis
	Xmag          float32 // X Magnetic field
	Ymag          float32 // Y Magnetic field
	Zmag          float32 // Z Magnetic field
	AbsPressure   float32 // Absolute pressure
	DiffPressure  float32 // Differential pressure
	PressureAlt   float32 // Altitude calculated from pressure
	Temperature   float32 // Temperature
	FieldsUpdated uint16  // Bitmap for fields that have updated since last message, bit 0 = xacc, bit 12: temperature
}

// MsgID (generated function)
func (m *HighresImu) MsgID() message.MessageID {
	return MSG_ID_HIGHRES_IMU
}

// String (generated function)
func (m *HighresImu) String() string {
	return fmt.Sprintf(
		"&common.HighresImu{ TimeUsec: %+v, Xacc: %+v, Yacc: %+v, Zacc: %+v, Xgyro: %+v, Ygyro: %+v, Zgyro: %+v, Xmag: %+v, Ymag: %+v, Zmag: %+v, AbsPressure: %+v, DiffPressure: %+v, PressureAlt: %+v, Temperature: %+v, FieldsUpdated: %+v }",
		m.TimeUsec,
		m.Xacc,
		m.Yacc,
		m.Zacc,
		m.Xgyro,
		m.Ygyro,
		m.Zgyro,
		m.Xmag,
		m.Ymag,
		m.Zmag,
		m.AbsPressure,
		m.DiffPressure,
		m.PressureAlt,
		m.Temperature,
		m.FieldsUpdated,
	)
}

// Marshal (generated function)
func (m *HighresImu) Marshal() ([]byte, error) {
	payload := make([]byte, 62)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Xacc))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Yacc))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Zacc))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Xgyro))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Ygyro))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Zgyro))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Xmag))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Ymag))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Zmag))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.AbsPressure))
	binary.LittleEndian.PutUint32(payload[48:], math.Float32bits(m.DiffPressure))
	binary.LittleEndian.PutUint32(payload[52:], math.Float32bits(m.PressureAlt))
	binary.LittleEndian.PutUint32(payload[56:], math.Float32bits(m.Temperature))
	binary.LittleEndian.PutUint16(payload[60:], uint16(m.FieldsUpdated))
	return payload, nil
}

// Unmarshal (generated function)
func (m *HighresImu) Unmarshal(data []byte) error {
	payload := make([]byte, 62)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Xacc = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Yacc = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Zacc = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Xgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Ygyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Zgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Xmag = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Ymag = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Zmag = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.AbsPressure = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.DiffPressure = math.Float32frombits(binary.LittleEndian.Uint32(payload[48:]))
	m.PressureAlt = math.Float32frombits(binary.LittleEndian.Uint32(payload[52:]))
	m.Temperature = math.Float32frombits(binary.LittleEndian.Uint32(payload[56:]))
	m.FieldsUpdated = uint16(binary.LittleEndian.Uint16(payload[60:]))
	return nil
}

// OpticalFlowRad struct (generated typeinfo)
// Optical flow from an angular rate flow sensor (e.g. PX4FLOW or mouse sensor)
type OpticalFlowRad struct {
	TimeUsec            uint64  // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	IntegrationTimeUs   uint32  // Integration time. Divide integrated_x and integrated_y by the integration time to obtain average flow. The integration time also indicates the.
	IntegratedX         float32 // Flow around X axis (Sensor RH rotation about the X axis induces a positive flow. Sensor linear motion along the positive Y axis induces a negative flow.)
	IntegratedY         float32 // Flow around Y axis (Sensor RH rotation about the Y axis induces a positive flow. Sensor linear motion along the positive X axis induces a positive flow.)
	IntegratedXgyro     float32 // RH rotation around X axis
	IntegratedYgyro     float32 // RH rotation around Y axis
	IntegratedZgyro     float32 // RH rotation around Z axis
	TimeDeltaDistanceUs uint32  // Time since the distance was sampled.
	Distance            float32 // Distance to the center of the flow field. Positive value (including zero): distance known. Negative value: Unknown distance.
	Temperature         int16   // Temperature
	SensorID            uint8   // Sensor ID
	Quality             uint8   // Optical flow quality / confidence. 0: no valid flow, 255: maximum quality
}

// MsgID (generated function)
func (m *OpticalFlowRad) MsgID() message.MessageID {
	return MSG_ID_OPTICAL_FLOW_RAD
}

// String (generated function)
func (m *OpticalFlowRad) String() string {
	return fmt.Sprintf(
		"&common.OpticalFlowRad{ TimeUsec: %+v, IntegrationTimeUs: %+v, IntegratedX: %+v, IntegratedY: %+v, IntegratedXgyro: %+v, IntegratedYgyro: %+v, IntegratedZgyro: %+v, TimeDeltaDistanceUs: %+v, Distance: %+v, Temperature: %+v, SensorID: %+v, Quality: %+v }",
		m.TimeUsec,
		m.IntegrationTimeUs,
		m.IntegratedX,
		m.IntegratedY,
		m.IntegratedXgyro,
		m.IntegratedYgyro,
		m.IntegratedZgyro,
		m.TimeDeltaDistanceUs,
		m.Distance,
		m.Temperature,
		m.SensorID,
		m.Quality,
	)
}

// Marshal (generated function)
func (m *OpticalFlowRad) Marshal() ([]byte, error) {
	payload := make([]byte, 44)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.IntegrationTimeUs))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.IntegratedX))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.IntegratedY))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.IntegratedXgyro))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.IntegratedYgyro))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.IntegratedZgyro))
	binary.LittleEndian.PutUint32(payload[32:], uint32(m.TimeDeltaDistanceUs))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Distance))
	binary.LittleEndian.PutUint16(payload[40:], uint16(m.Temperature))
	payload[42] = byte(m.SensorID)
	payload[43] = byte(m.Quality)
	return payload, nil
}

// Unmarshal (generated function)
func (m *OpticalFlowRad) Unmarshal(data []byte) error {
	payload := make([]byte, 44)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.IntegrationTimeUs = uint32(binary.LittleEndian.Uint32(payload[8:]))
	m.IntegratedX = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.IntegratedY = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.IntegratedXgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.IntegratedYgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.IntegratedZgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.TimeDeltaDistanceUs = uint32(binary.LittleEndian.Uint32(payload[32:]))
	m.Distance = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Temperature = int16(binary.LittleEndian.Uint16(payload[40:]))
	m.SensorID = uint8(payload[42])
	m.Quality = uint8(payload[43])
	return nil
}

// HilSensor struct (generated typeinfo)
// The IMU readings in SI units in NED body frame
type HilSensor struct {
	TimeUsec      uint64  // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Xacc          float32 // X acceleration
	Yacc          float32 // Y acceleration
	Zacc          float32 // Z acceleration
	Xgyro         float32 // Angular speed around X axis in body frame
	Ygyro         float32 // Angular speed around Y axis in body frame
	Zgyro         float32 // Angular speed around Z axis in body frame
	Xmag          float32 // X Magnetic field
	Ymag          float32 // Y Magnetic field
	Zmag          float32 // Z Magnetic field
	AbsPressure   float32 // Absolute pressure
	DiffPressure  float32 // Differential pressure (airspeed)
	PressureAlt   float32 // Altitude calculated from pressure
	Temperature   float32 // Temperature
	FieldsUpdated uint32  // Bitmap for fields that have updated since last message, bit 0 = xacc, bit 12: temperature, bit 31: full reset of attitude/position/velocities/etc was performed in sim.
}

// MsgID (generated function)
func (m *HilSensor) MsgID() message.MessageID {
	return MSG_ID_HIL_SENSOR
}

// String (generated function)
func (m *HilSensor) String() string {
	return fmt.Sprintf(
		"&common.HilSensor{ TimeUsec: %+v, Xacc: %+v, Yacc: %+v, Zacc: %+v, Xgyro: %+v, Ygyro: %+v, Zgyro: %+v, Xmag: %+v, Ymag: %+v, Zmag: %+v, AbsPressure: %+v, DiffPressure: %+v, PressureAlt: %+v, Temperature: %+v, FieldsUpdated: %+v }",
		m.TimeUsec,
		m.Xacc,
		m.Yacc,
		m.Zacc,
		m.Xgyro,
		m.Ygyro,
		m.Zgyro,
		m.Xmag,
		m.Ymag,
		m.Zmag,
		m.AbsPressure,
		m.DiffPressure,
		m.PressureAlt,
		m.Temperature,
		m.FieldsUpdated,
	)
}

// Marshal (generated function)
func (m *HilSensor) Marshal() ([]byte, error) {
	payload := make([]byte, 64)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Xacc))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Yacc))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Zacc))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Xgyro))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Ygyro))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Zgyro))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Xmag))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Ymag))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Zmag))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.AbsPressure))
	binary.LittleEndian.PutUint32(payload[48:], math.Float32bits(m.DiffPressure))
	binary.LittleEndian.PutUint32(payload[52:], math.Float32bits(m.PressureAlt))
	binary.LittleEndian.PutUint32(payload[56:], math.Float32bits(m.Temperature))
	binary.LittleEndian.PutUint32(payload[60:], uint32(m.FieldsUpdated))
	return payload, nil
}

// Unmarshal (generated function)
func (m *HilSensor) Unmarshal(data []byte) error {
	payload := make([]byte, 64)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Xacc = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Yacc = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Zacc = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Xgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Ygyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Zgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Xmag = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Ymag = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Zmag = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.AbsPressure = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.DiffPressure = math.Float32frombits(binary.LittleEndian.Uint32(payload[48:]))
	m.PressureAlt = math.Float32frombits(binary.LittleEndian.Uint32(payload[52:]))
	m.Temperature = math.Float32frombits(binary.LittleEndian.Uint32(payload[56:]))
	m.FieldsUpdated = uint32(binary.LittleEndian.Uint32(payload[60:]))
	return nil
}

// SimState struct (generated typeinfo)
// Status of simulation environment, if used
type SimState struct {
	Q1         float32 // True attitude quaternion component 1, w (1 in null-rotation)
	Q2         float32 // True attitude quaternion component 2, x (0 in null-rotation)
	Q3         float32 // True attitude quaternion component 3, y (0 in null-rotation)
	Q4         float32 // True attitude quaternion component 4, z (0 in null-rotation)
	Roll       float32 // Attitude roll expressed as Euler angles, not recommended except for human-readable outputs
	Pitch      float32 // Attitude pitch expressed as Euler angles, not recommended except for human-readable outputs
	Yaw        float32 // Attitude yaw expressed as Euler angles, not recommended except for human-readable outputs
	Xacc       float32 // X acceleration
	Yacc       float32 // Y acceleration
	Zacc       float32 // Z acceleration
	Xgyro      float32 // Angular speed around X axis
	Ygyro      float32 // Angular speed around Y axis
	Zgyro      float32 // Angular speed around Z axis
	Lat        float32 // Latitude
	Lon        float32 // Longitude
	Alt        float32 // Altitude
	StdDevHorz float32 // Horizontal position standard deviation
	StdDevVert float32 // Vertical position standard deviation
	Vn         float32 // True velocity in north direction in earth-fixed NED frame
	Ve         float32 // True velocity in east direction in earth-fixed NED frame
	Vd         float32 // True velocity in down direction in earth-fixed NED frame
}

// MsgID (generated function)
func (m *SimState) MsgID() message.MessageID {
	return MSG_ID_SIM_STATE
}

// String (generated function)
func (m *SimState) String() string {
	return fmt.Sprintf(
		"&common.SimState{ Q1: %+v, Q2: %+v, Q3: %+v, Q4: %+v, Roll: %+v, Pitch: %+v, Yaw: %+v, Xacc: %+v, Yacc: %+v, Zacc: %+v, Xgyro: %+v, Ygyro: %+v, Zgyro: %+v, Lat: %+v, Lon: %+v, Alt: %+v, StdDevHorz: %+v, StdDevVert: %+v, Vn: %+v, Ve: %+v, Vd: %+v }",
		m.Q1,
		m.Q2,
		m.Q3,
		m.Q4,
		m.Roll,
		m.Pitch,
		m.Yaw,
		m.Xacc,
		m.Yacc,
		m.Zacc,
		m.Xgyro,
		m.Ygyro,
		m.Zgyro,
		m.Lat,
		m.Lon,
		m.Alt,
		m.StdDevHorz,
		m.StdDevVert,
		m.Vn,
		m.Ve,
		m.Vd,
	)
}

// Marshal (generated function)
func (m *SimState) Marshal() ([]byte, error) {
	payload := make([]byte, 84)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.Q1))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Q2))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Q3))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Q4))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Roll))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Pitch))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Yaw))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Xacc))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Yacc))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Zacc))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Xgyro))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.Ygyro))
	binary.LittleEndian.PutUint32(payload[48:], math.Float32bits(m.Zgyro))
	binary.LittleEndian.PutUint32(payload[52:], math.Float32bits(m.Lat))
	binary.LittleEndian.PutUint32(payload[56:], math.Float32bits(m.Lon))
	binary.LittleEndian.PutUint32(payload[60:], math.Float32bits(m.Alt))
	binary.LittleEndian.PutUint32(payload[64:], math.Float32bits(m.StdDevHorz))
	binary.LittleEndian.PutUint32(payload[68:], math.Float32bits(m.StdDevVert))
	binary.LittleEndian.PutUint32(payload[72:], math.Float32bits(m.Vn))
	binary.LittleEndian.PutUint32(payload[76:], math.Float32bits(m.Ve))
	binary.LittleEndian.PutUint32(payload[80:], math.Float32bits(m.Vd))
	return payload, nil
}

// Unmarshal (generated function)
func (m *SimState) Unmarshal(data []byte) error {
	payload := make([]byte, 84)
	copy(payload[0:], data)
	m.Q1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.Q2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Q3 = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Q4 = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Roll = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Pitch = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Xacc = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Yacc = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Zacc = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Xgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.Ygyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.Zgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[48:]))
	m.Lat = math.Float32frombits(binary.LittleEndian.Uint32(payload[52:]))
	m.Lon = math.Float32frombits(binary.LittleEndian.Uint32(payload[56:]))
	m.Alt = math.Float32frombits(binary.LittleEndian.Uint32(payload[60:]))
	m.StdDevHorz = math.Float32frombits(binary.LittleEndian.Uint32(payload[64:]))
	m.StdDevVert = math.Float32frombits(binary.LittleEndian.Uint32(payload[68:]))
	m.Vn = math.Float32frombits(binary.LittleEndian.Uint32(payload[72:]))
	m.Ve = math.Float32frombits(binary.LittleEndian.Uint32(payload[76:]))
	m.Vd = math.Float32frombits(binary.LittleEndian.Uint32(payload[80:]))
	return nil
}

// RadioStatus struct (generated typeinfo)
// Status generated by radio and injected into MAVLink stream.
type RadioStatus struct {
	Rxerrors uint16 // Count of radio packet receive errors (since boot).
	Fixed    uint16 // Count of error corrected radio packets (since boot).
	Rssi     uint8  // Local (message sender) recieved signal strength indication in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
	Remrssi  uint8  // Remote (message receiver) signal strength indication in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
	Txbuf    uint8  // Remaining free transmitter buffer space.
	Noise    uint8  // Local background noise level. These are device dependent RSSI values (scale as approx 2x dB on SiK radios). Values: [0-254], 255: invalid/unknown.
	Remnoise uint8  // Remote background noise level. These are device dependent RSSI values (scale as approx 2x dB on SiK radios). Values: [0-254], 255: invalid/unknown.
}

// MsgID (generated function)
func (m *RadioStatus) MsgID() message.MessageID {
	return MSG_ID_RADIO_STATUS
}

// String (generated function)
func (m *RadioStatus) String() string {
	return fmt.Sprintf(
		"&common.RadioStatus{ Rxerrors: %+v, Fixed: %+v, Rssi: %+v, Remrssi: %+v, Txbuf: %+v, Noise: %+v, Remnoise: %+v }",
		m.Rxerrors,
		m.Fixed,
		m.Rssi,
		m.Remrssi,
		m.Txbuf,
		m.Noise,
		m.Remnoise,
	)
}

// Marshal (generated function)
func (m *RadioStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 9)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Rxerrors))
	binary.LittleEndian.PutUint16(payload[2:], uint16(m.Fixed))
	payload[4] = byte(m.Rssi)
	payload[5] = byte(m.Remrssi)
	payload[6] = byte(m.Txbuf)
	payload[7] = byte(m.Noise)
	payload[8] = byte(m.Remnoise)
	return payload, nil
}

// Unmarshal (generated function)
func (m *RadioStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 9)
	copy(payload[0:], data)
	m.Rxerrors = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.Fixed = uint16(binary.LittleEndian.Uint16(payload[2:]))
	m.Rssi = uint8(payload[4])
	m.Remrssi = uint8(payload[5])
	m.Txbuf = uint8(payload[6])
	m.Noise = uint8(payload[7])
	m.Remnoise = uint8(payload[8])
	return nil
}

// FileTransferProtocol struct (generated typeinfo)
// File transfer message
type FileTransferProtocol struct {
	TargetNetwork   uint8   // Network ID (0 for broadcast)
	TargetSystem    uint8   // System ID (0 for broadcast)
	TargetComponent uint8   // Component ID (0 for broadcast)
	Payload         []uint8 `len:"251" ` // Variable length payload. The length is defined by the remaining message length when subtracting the header and other fields.  The entire content of this block is opaque unless you understand any the encoding message_type.  The particular encoding used can be extension specific and might not always be documented as part of the mavlink specification.
}

// MsgID (generated function)
func (m *FileTransferProtocol) MsgID() message.MessageID {
	return MSG_ID_FILE_TRANSFER_PROTOCOL
}

// String (generated function)
func (m *FileTransferProtocol) String() string {
	return fmt.Sprintf(
		"&common.FileTransferProtocol{ TargetNetwork: %+v, TargetSystem: %+v, TargetComponent: %+v, Payload: %0X (\"%s\") }",
		m.TargetNetwork,
		m.TargetSystem,
		m.TargetComponent,
		m.Payload, string(m.Payload[:]),
	)
}

// Marshal (generated function)
func (m *FileTransferProtocol) Marshal() ([]byte, error) {
	payload := make([]byte, 254)
	payload[0] = byte(m.TargetNetwork)
	payload[1] = byte(m.TargetSystem)
	payload[2] = byte(m.TargetComponent)
	copy(payload[3:], m.Payload)
	return payload, nil
}

// Unmarshal (generated function)
func (m *FileTransferProtocol) Unmarshal(data []byte) error {
	payload := make([]byte, 254)
	copy(payload[0:], data)
	m.TargetNetwork = uint8(payload[0])
	m.TargetSystem = uint8(payload[1])
	m.TargetComponent = uint8(payload[2])
	copy(m.Payload[:], payload[3:254])
	return nil
}

// Timesync struct (generated typeinfo)
// Time synchronization message.
type Timesync struct {
	Tc1 int64 // Time sync timestamp 1
	Ts1 int64 // Time sync timestamp 2
}

// MsgID (generated function)
func (m *Timesync) MsgID() message.MessageID {
	return MSG_ID_TIMESYNC
}

// String (generated function)
func (m *Timesync) String() string {
	return fmt.Sprintf(
		"&common.Timesync{ Tc1: %+v, Ts1: %+v }",
		m.Tc1,
		m.Ts1,
	)
}

// Marshal (generated function)
func (m *Timesync) Marshal() ([]byte, error) {
	payload := make([]byte, 16)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Tc1))
	binary.LittleEndian.PutUint64(payload[8:], uint64(m.Ts1))
	return payload, nil
}

// Unmarshal (generated function)
func (m *Timesync) Unmarshal(data []byte) error {
	payload := make([]byte, 16)
	copy(payload[0:], data)
	m.Tc1 = int64(binary.LittleEndian.Uint64(payload[0:]))
	m.Ts1 = int64(binary.LittleEndian.Uint64(payload[8:]))
	return nil
}

// CameraTrigger struct (generated typeinfo)
// Camera-IMU triggering and synchronisation message.
type CameraTrigger struct {
	TimeUsec uint64 // Timestamp for image frame (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Seq      uint32 // Image frame sequence
}

// MsgID (generated function)
func (m *CameraTrigger) MsgID() message.MessageID {
	return MSG_ID_CAMERA_TRIGGER
}

// String (generated function)
func (m *CameraTrigger) String() string {
	return fmt.Sprintf(
		"&common.CameraTrigger{ TimeUsec: %+v, Seq: %+v }",
		m.TimeUsec,
		m.Seq,
	)
}

// Marshal (generated function)
func (m *CameraTrigger) Marshal() ([]byte, error) {
	payload := make([]byte, 12)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Seq))
	return payload, nil
}

// Unmarshal (generated function)
func (m *CameraTrigger) Unmarshal(data []byte) error {
	payload := make([]byte, 12)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Seq = uint32(binary.LittleEndian.Uint32(payload[8:]))
	return nil
}

// HilGps struct (generated typeinfo)
// The global position, as returned by the Global Positioning System (GPS). This is
//                  NOT the global position estimate of the sytem, but rather a RAW sensor value. See message GLOBAL_POSITION for the global position estimate.
type HilGps struct {
	TimeUsec          uint64 // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Lat               int32  // Latitude (WGS84)
	Lon               int32  // Longitude (WGS84)
	Alt               int32  // Altitude (MSL). Positive for up.
	Eph               uint16 // GPS HDOP horizontal dilution of position. If unknown, set to: 65535
	Epv               uint16 // GPS VDOP vertical dilution of position. If unknown, set to: 65535
	Vel               uint16 // GPS ground speed. If unknown, set to: 65535
	Vn                int16  // GPS velocity in north direction in earth-fixed NED frame
	Ve                int16  // GPS velocity in east direction in earth-fixed NED frame
	Vd                int16  // GPS velocity in down direction in earth-fixed NED frame
	Cog               uint16 // Course over ground (NOT heading, but direction of movement), 0.0..359.99 degrees. If unknown, set to: 65535
	FixType           uint8  // 0-1: no fix, 2: 2D fix, 3: 3D fix. Some applications will not use the value of this field unless it is at least two, so always correctly fill in the fix.
	SatellitesVisible uint8  // Number of satellites visible. If unknown, set to 255
}

// MsgID (generated function)
func (m *HilGps) MsgID() message.MessageID {
	return MSG_ID_HIL_GPS
}

// String (generated function)
func (m *HilGps) String() string {
	return fmt.Sprintf(
		"&common.HilGps{ TimeUsec: %+v, Lat: %+v, Lon: %+v, Alt: %+v, Eph: %+v, Epv: %+v, Vel: %+v, Vn: %+v, Ve: %+v, Vd: %+v, Cog: %+v, FixType: %+v, SatellitesVisible: %+v }",
		m.TimeUsec,
		m.Lat,
		m.Lon,
		m.Alt,
		m.Eph,
		m.Epv,
		m.Vel,
		m.Vn,
		m.Ve,
		m.Vd,
		m.Cog,
		m.FixType,
		m.SatellitesVisible,
	)
}

// Marshal (generated function)
func (m *HilGps) Marshal() ([]byte, error) {
	payload := make([]byte, 36)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.Alt))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Eph))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.Epv))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.Vel))
	binary.LittleEndian.PutUint16(payload[26:], uint16(m.Vn))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Ve))
	binary.LittleEndian.PutUint16(payload[30:], uint16(m.Vd))
	binary.LittleEndian.PutUint16(payload[32:], uint16(m.Cog))
	payload[34] = byte(m.FixType)
	payload[35] = byte(m.SatellitesVisible)
	return payload, nil
}

// Unmarshal (generated function)
func (m *HilGps) Unmarshal(data []byte) error {
	payload := make([]byte, 36)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[12:]))
	m.Alt = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.Eph = uint16(binary.LittleEndian.Uint16(payload[20:]))
	m.Epv = uint16(binary.LittleEndian.Uint16(payload[22:]))
	m.Vel = uint16(binary.LittleEndian.Uint16(payload[24:]))
	m.Vn = int16(binary.LittleEndian.Uint16(payload[26:]))
	m.Ve = int16(binary.LittleEndian.Uint16(payload[28:]))
	m.Vd = int16(binary.LittleEndian.Uint16(payload[30:]))
	m.Cog = uint16(binary.LittleEndian.Uint16(payload[32:]))
	m.FixType = uint8(payload[34])
	m.SatellitesVisible = uint8(payload[35])
	return nil
}

// HilOpticalFlow struct (generated typeinfo)
// Simulated optical flow from a flow sensor (e.g. PX4FLOW or optical mouse sensor)
type HilOpticalFlow struct {
	TimeUsec            uint64  // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	IntegrationTimeUs   uint32  // Integration time. Divide integrated_x and integrated_y by the integration time to obtain average flow. The integration time also indicates the.
	IntegratedX         float32 // Flow in radians around X axis (Sensor RH rotation about the X axis induces a positive flow. Sensor linear motion along the positive Y axis induces a negative flow.)
	IntegratedY         float32 // Flow in radians around Y axis (Sensor RH rotation about the Y axis induces a positive flow. Sensor linear motion along the positive X axis induces a positive flow.)
	IntegratedXgyro     float32 // RH rotation around X axis
	IntegratedYgyro     float32 // RH rotation around Y axis
	IntegratedZgyro     float32 // RH rotation around Z axis
	TimeDeltaDistanceUs uint32  // Time since the distance was sampled.
	Distance            float32 // Distance to the center of the flow field. Positive value (including zero): distance known. Negative value: Unknown distance.
	Temperature         int16   // Temperature
	SensorID            uint8   // Sensor ID
	Quality             uint8   // Optical flow quality / confidence. 0: no valid flow, 255: maximum quality
}

// MsgID (generated function)
func (m *HilOpticalFlow) MsgID() message.MessageID {
	return MSG_ID_HIL_OPTICAL_FLOW
}

// String (generated function)
func (m *HilOpticalFlow) String() string {
	return fmt.Sprintf(
		"&common.HilOpticalFlow{ TimeUsec: %+v, IntegrationTimeUs: %+v, IntegratedX: %+v, IntegratedY: %+v, IntegratedXgyro: %+v, IntegratedYgyro: %+v, IntegratedZgyro: %+v, TimeDeltaDistanceUs: %+v, Distance: %+v, Temperature: %+v, SensorID: %+v, Quality: %+v }",
		m.TimeUsec,
		m.IntegrationTimeUs,
		m.IntegratedX,
		m.IntegratedY,
		m.IntegratedXgyro,
		m.IntegratedYgyro,
		m.IntegratedZgyro,
		m.TimeDeltaDistanceUs,
		m.Distance,
		m.Temperature,
		m.SensorID,
		m.Quality,
	)
}

// Marshal (generated function)
func (m *HilOpticalFlow) Marshal() ([]byte, error) {
	payload := make([]byte, 44)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.IntegrationTimeUs))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.IntegratedX))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.IntegratedY))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.IntegratedXgyro))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.IntegratedYgyro))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.IntegratedZgyro))
	binary.LittleEndian.PutUint32(payload[32:], uint32(m.TimeDeltaDistanceUs))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Distance))
	binary.LittleEndian.PutUint16(payload[40:], uint16(m.Temperature))
	payload[42] = byte(m.SensorID)
	payload[43] = byte(m.Quality)
	return payload, nil
}

// Unmarshal (generated function)
func (m *HilOpticalFlow) Unmarshal(data []byte) error {
	payload := make([]byte, 44)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.IntegrationTimeUs = uint32(binary.LittleEndian.Uint32(payload[8:]))
	m.IntegratedX = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.IntegratedY = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.IntegratedXgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.IntegratedYgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.IntegratedZgyro = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.TimeDeltaDistanceUs = uint32(binary.LittleEndian.Uint32(payload[32:]))
	m.Distance = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Temperature = int16(binary.LittleEndian.Uint16(payload[40:]))
	m.SensorID = uint8(payload[42])
	m.Quality = uint8(payload[43])
	return nil
}

// HilStateQuaternion struct (generated typeinfo)
// Sent from simulation to autopilot, avoids in contrast to HIL_STATE singularities. This packet is useful for high throughput applications such as hardware in the loop simulations.
type HilStateQuaternion struct {
	TimeUsec           uint64    // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	AttitudeQuaternion []float32 `len:"4" ` // Vehicle attitude expressed as normalized quaternion in w, x, y, z order (with 1 0 0 0 being the null-rotation)
	Rollspeed          float32   // Body frame roll / phi angular speed
	Pitchspeed         float32   // Body frame pitch / theta angular speed
	Yawspeed           float32   // Body frame yaw / psi angular speed
	Lat                int32     // Latitude
	Lon                int32     // Longitude
	Alt                int32     // Altitude
	Vx                 int16     // Ground X Speed (Latitude)
	Vy                 int16     // Ground Y Speed (Longitude)
	Vz                 int16     // Ground Z Speed (Altitude)
	IndAirspeed        uint16    // Indicated airspeed
	TrueAirspeed       uint16    // True airspeed
	Xacc               int16     // X acceleration
	Yacc               int16     // Y acceleration
	Zacc               int16     // Z acceleration
}

// MsgID (generated function)
func (m *HilStateQuaternion) MsgID() message.MessageID {
	return MSG_ID_HIL_STATE_QUATERNION
}

// String (generated function)
func (m *HilStateQuaternion) String() string {
	return fmt.Sprintf(
		"&common.HilStateQuaternion{ TimeUsec: %+v, AttitudeQuaternion: %+v, Rollspeed: %+v, Pitchspeed: %+v, Yawspeed: %+v, Lat: %+v, Lon: %+v, Alt: %+v, Vx: %+v, Vy: %+v, Vz: %+v, IndAirspeed: %+v, TrueAirspeed: %+v, Xacc: %+v, Yacc: %+v, Zacc: %+v }",
		m.TimeUsec,
		m.AttitudeQuaternion,
		m.Rollspeed,
		m.Pitchspeed,
		m.Yawspeed,
		m.Lat,
		m.Lon,
		m.Alt,
		m.Vx,
		m.Vy,
		m.Vz,
		m.IndAirspeed,
		m.TrueAirspeed,
		m.Xacc,
		m.Yacc,
		m.Zacc,
	)
}

// Marshal (generated function)
func (m *HilStateQuaternion) Marshal() ([]byte, error) {
	payload := make([]byte, 64)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	for i, v := range m.AttitudeQuaternion {
		binary.LittleEndian.PutUint32(payload[8+i*4:], math.Float32bits(v))
	}
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Rollspeed))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Pitchspeed))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Yawspeed))
	binary.LittleEndian.PutUint32(payload[36:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[40:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[44:], uint32(m.Alt))
	binary.LittleEndian.PutUint16(payload[48:], uint16(m.Vx))
	binary.LittleEndian.PutUint16(payload[50:], uint16(m.Vy))
	binary.LittleEndian.PutUint16(payload[52:], uint16(m.Vz))
	binary.LittleEndian.PutUint16(payload[54:], uint16(m.IndAirspeed))
	binary.LittleEndian.PutUint16(payload[56:], uint16(m.TrueAirspeed))
	binary.LittleEndian.PutUint16(payload[58:], uint16(m.Xacc))
	binary.LittleEndian.PutUint16(payload[60:], uint16(m.Yacc))
	binary.LittleEndian.PutUint16(payload[62:], uint16(m.Zacc))
	return payload, nil
}

// Unmarshal (generated function)
func (m *HilStateQuaternion) Unmarshal(data []byte) error {
	payload := make([]byte, 64)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	for i := 0; i < len(m.AttitudeQuaternion); i++ {
		m.AttitudeQuaternion[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[8+i*4:]))
	}
	m.Rollspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Pitchspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Yawspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[36:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[40:]))
	m.Alt = int32(binary.LittleEndian.Uint32(payload[44:]))
	m.Vx = int16(binary.LittleEndian.Uint16(payload[48:]))
	m.Vy = int16(binary.LittleEndian.Uint16(payload[50:]))
	m.Vz = int16(binary.LittleEndian.Uint16(payload[52:]))
	m.IndAirspeed = uint16(binary.LittleEndian.Uint16(payload[54:]))
	m.TrueAirspeed = uint16(binary.LittleEndian.Uint16(payload[56:]))
	m.Xacc = int16(binary.LittleEndian.Uint16(payload[58:]))
	m.Yacc = int16(binary.LittleEndian.Uint16(payload[60:]))
	m.Zacc = int16(binary.LittleEndian.Uint16(payload[62:]))
	return nil
}

// ScaledImu2 struct (generated typeinfo)
// The RAW IMU readings for secondary 9DOF sensor setup. This message should contain the scaled values to the described units
type ScaledImu2 struct {
	TimeBootMs uint32 // Timestamp (time since system boot).
	Xacc       int16  // X acceleration
	Yacc       int16  // Y acceleration
	Zacc       int16  // Z acceleration
	Xgyro      int16  // Angular speed around X axis
	Ygyro      int16  // Angular speed around Y axis
	Zgyro      int16  // Angular speed around Z axis
	Xmag       int16  // X Magnetic field
	Ymag       int16  // Y Magnetic field
	Zmag       int16  // Z Magnetic field
}

// MsgID (generated function)
func (m *ScaledImu2) MsgID() message.MessageID {
	return MSG_ID_SCALED_IMU2
}

// String (generated function)
func (m *ScaledImu2) String() string {
	return fmt.Sprintf(
		"&common.ScaledImu2{ TimeBootMs: %+v, Xacc: %+v, Yacc: %+v, Zacc: %+v, Xgyro: %+v, Ygyro: %+v, Zgyro: %+v, Xmag: %+v, Ymag: %+v, Zmag: %+v }",
		m.TimeBootMs,
		m.Xacc,
		m.Yacc,
		m.Zacc,
		m.Xgyro,
		m.Ygyro,
		m.Zgyro,
		m.Xmag,
		m.Ymag,
		m.Zmag,
	)
}

// Marshal (generated function)
func (m *ScaledImu2) Marshal() ([]byte, error) {
	payload := make([]byte, 22)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Xacc))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.Yacc))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Zacc))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.Xgyro))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Ygyro))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Zgyro))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Xmag))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Ymag))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Zmag))
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScaledImu2) Unmarshal(data []byte) error {
	payload := make([]byte, 22)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Xacc = int16(binary.LittleEndian.Uint16(payload[4:]))
	m.Yacc = int16(binary.LittleEndian.Uint16(payload[6:]))
	m.Zacc = int16(binary.LittleEndian.Uint16(payload[8:]))
	m.Xgyro = int16(binary.LittleEndian.Uint16(payload[10:]))
	m.Ygyro = int16(binary.LittleEndian.Uint16(payload[12:]))
	m.Zgyro = int16(binary.LittleEndian.Uint16(payload[14:]))
	m.Xmag = int16(binary.LittleEndian.Uint16(payload[16:]))
	m.Ymag = int16(binary.LittleEndian.Uint16(payload[18:]))
	m.Zmag = int16(binary.LittleEndian.Uint16(payload[20:]))
	return nil
}

// LogRequestList struct (generated typeinfo)
// Request a list of available logs. On some systems calling this may stop on-board logging until LOG_REQUEST_END is called. If there are no log files available this request shall be answered with one LOG_ENTRY message with id = 0 and num_logs = 0.
type LogRequestList struct {
	Start           uint16 // First log id (0 for first available)
	End             uint16 // Last log id (0xffff for last available)
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
}

// MsgID (generated function)
func (m *LogRequestList) MsgID() message.MessageID {
	return MSG_ID_LOG_REQUEST_LIST
}

// String (generated function)
func (m *LogRequestList) String() string {
	return fmt.Sprintf(
		"&common.LogRequestList{ Start: %+v, End: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.Start,
		m.End,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *LogRequestList) Marshal() ([]byte, error) {
	payload := make([]byte, 6)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Start))
	binary.LittleEndian.PutUint16(payload[2:], uint16(m.End))
	payload[4] = byte(m.TargetSystem)
	payload[5] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *LogRequestList) Unmarshal(data []byte) error {
	payload := make([]byte, 6)
	copy(payload[0:], data)
	m.Start = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.End = uint16(binary.LittleEndian.Uint16(payload[2:]))
	m.TargetSystem = uint8(payload[4])
	m.TargetComponent = uint8(payload[5])
	return nil
}

// LogEntry struct (generated typeinfo)
// Reply to LOG_REQUEST_LIST
type LogEntry struct {
	TimeUtc    uint32 // UTC timestamp of log since 1970, or 0 if not available
	Size       uint32 // Size of the log (may be approximate)
	ID         uint16 // Log id
	NumLogs    uint16 // Total number of logs
	LastLogNum uint16 // High log number
}

// MsgID (generated function)
func (m *LogEntry) MsgID() message.MessageID {
	return MSG_ID_LOG_ENTRY
}

// String (generated function)
func (m *LogEntry) String() string {
	return fmt.Sprintf(
		"&common.LogEntry{ TimeUtc: %+v, Size: %+v, ID: %+v, NumLogs: %+v, LastLogNum: %+v }",
		m.TimeUtc,
		m.Size,
		m.ID,
		m.NumLogs,
		m.LastLogNum,
	)
}

// Marshal (generated function)
func (m *LogEntry) Marshal() ([]byte, error) {
	payload := make([]byte, 14)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeUtc))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Size))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.ID))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.NumLogs))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.LastLogNum))
	return payload, nil
}

// Unmarshal (generated function)
func (m *LogEntry) Unmarshal(data []byte) error {
	payload := make([]byte, 14)
	copy(payload[0:], data)
	m.TimeUtc = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Size = uint32(binary.LittleEndian.Uint32(payload[4:]))
	m.ID = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.NumLogs = uint16(binary.LittleEndian.Uint16(payload[10:]))
	m.LastLogNum = uint16(binary.LittleEndian.Uint16(payload[12:]))
	return nil
}

// LogRequestData struct (generated typeinfo)
// Request a chunk of a log
type LogRequestData struct {
	Ofs             uint32 // Offset into the log
	Count           uint32 // Number of bytes
	ID              uint16 // Log id (from LOG_ENTRY reply)
	TargetSystem    uint8  // System ID
	TargetComponent uint8  // Component ID
}

// MsgID (generated function)
func (m *LogRequestData) MsgID() message.MessageID {
	return MSG_ID_LOG_REQUEST_DATA
}

// String (generated function)
func (m *LogRequestData) String() string {
	return fmt.Sprintf(
		"&common.LogRequestData{ Ofs: %+v, Count: %+v, ID: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.Ofs,
		m.Count,
		m.ID,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *LogRequestData) Marshal() ([]byte, error) {
	payload := make([]byte, 12)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Ofs))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Count))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.ID))
	payload[10] = byte(m.TargetSystem)
	payload[11] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *LogRequestData) Unmarshal(data []byte) error {
	payload := make([]byte, 12)
	copy(payload[0:], data)
	m.Ofs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Count = uint32(binary.LittleEndian.Uint32(payload[4:]))
	m.ID = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.TargetSystem = uint8(payload[10])
	m.TargetComponent = uint8(payload[11])
	return nil
}

// LogData struct (generated typeinfo)
// Reply to LOG_REQUEST_DATA
type LogData struct {
	Ofs   uint32  // Offset into the log
	ID    uint16  // Log id (from LOG_ENTRY reply)
	Count uint8   // Number of bytes (zero for end of log)
	Data  []uint8 `len:"90" ` // log data
}

// MsgID (generated function)
func (m *LogData) MsgID() message.MessageID {
	return MSG_ID_LOG_DATA
}

// String (generated function)
func (m *LogData) String() string {
	return fmt.Sprintf(
		"&common.LogData{ Ofs: %+v, ID: %+v, Count: %+v, Data: %0X (\"%s\") }",
		m.Ofs,
		m.ID,
		m.Count,
		m.Data, string(m.Data[:]),
	)
}

// Marshal (generated function)
func (m *LogData) Marshal() ([]byte, error) {
	payload := make([]byte, 97)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Ofs))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.ID))
	payload[6] = byte(m.Count)
	copy(payload[7:], m.Data)
	return payload, nil
}

// Unmarshal (generated function)
func (m *LogData) Unmarshal(data []byte) error {
	payload := make([]byte, 97)
	copy(payload[0:], data)
	m.Ofs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.ID = uint16(binary.LittleEndian.Uint16(payload[4:]))
	m.Count = uint8(payload[6])
	copy(m.Data[:], payload[7:97])
	return nil
}

// LogErase struct (generated typeinfo)
// Erase all logs
type LogErase struct {
	TargetSystem    uint8 // System ID
	TargetComponent uint8 // Component ID
}

// MsgID (generated function)
func (m *LogErase) MsgID() message.MessageID {
	return MSG_ID_LOG_ERASE
}

// String (generated function)
func (m *LogErase) String() string {
	return fmt.Sprintf(
		"&common.LogErase{ TargetSystem: %+v, TargetComponent: %+v }",
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *LogErase) Marshal() ([]byte, error) {
	payload := make([]byte, 2)
	payload[0] = byte(m.TargetSystem)
	payload[1] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *LogErase) Unmarshal(data []byte) error {
	payload := make([]byte, 2)
	copy(payload[0:], data)
	m.TargetSystem = uint8(payload[0])
	m.TargetComponent = uint8(payload[1])
	return nil
}

// LogRequestEnd struct (generated typeinfo)
// Stop log transfer and resume normal logging
type LogRequestEnd struct {
	TargetSystem    uint8 // System ID
	TargetComponent uint8 // Component ID
}

// MsgID (generated function)
func (m *LogRequestEnd) MsgID() message.MessageID {
	return MSG_ID_LOG_REQUEST_END
}

// String (generated function)
func (m *LogRequestEnd) String() string {
	return fmt.Sprintf(
		"&common.LogRequestEnd{ TargetSystem: %+v, TargetComponent: %+v }",
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *LogRequestEnd) Marshal() ([]byte, error) {
	payload := make([]byte, 2)
	payload[0] = byte(m.TargetSystem)
	payload[1] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *LogRequestEnd) Unmarshal(data []byte) error {
	payload := make([]byte, 2)
	copy(payload[0:], data)
	m.TargetSystem = uint8(payload[0])
	m.TargetComponent = uint8(payload[1])
	return nil
}

// GpsInjectData struct (generated typeinfo)
// Data for injecting into the onboard GPS (used for DGPS)
type GpsInjectData struct {
	TargetSystem    uint8   // System ID
	TargetComponent uint8   // Component ID
	Len             uint8   // Data length
	Data            []uint8 `len:"110" ` // Raw data (110 is enough for 12 satellites of RTCMv2)
}

// MsgID (generated function)
func (m *GpsInjectData) MsgID() message.MessageID {
	return MSG_ID_GPS_INJECT_DATA
}

// String (generated function)
func (m *GpsInjectData) String() string {
	return fmt.Sprintf(
		"&common.GpsInjectData{ TargetSystem: %+v, TargetComponent: %+v, Len: %+v, Data: %0X (\"%s\") }",
		m.TargetSystem,
		m.TargetComponent,
		m.Len,
		m.Data, string(m.Data[:]),
	)
}

// Marshal (generated function)
func (m *GpsInjectData) Marshal() ([]byte, error) {
	payload := make([]byte, 113)
	payload[0] = byte(m.TargetSystem)
	payload[1] = byte(m.TargetComponent)
	payload[2] = byte(m.Len)
	copy(payload[3:], m.Data)
	return payload, nil
}

// Unmarshal (generated function)
func (m *GpsInjectData) Unmarshal(data []byte) error {
	payload := make([]byte, 113)
	copy(payload[0:], data)
	m.TargetSystem = uint8(payload[0])
	m.TargetComponent = uint8(payload[1])
	m.Len = uint8(payload[2])
	copy(m.Data[:], payload[3:113])
	return nil
}

// Gps2Raw struct (generated typeinfo)
// Second GPS data.
type Gps2Raw struct {
	TimeUsec          uint64       // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Lat               int32        // Latitude (WGS84)
	Lon               int32        // Longitude (WGS84)
	Alt               int32        // Altitude (MSL). Positive for up.
	DgpsAge           uint32       // Age of DGPS info
	Eph               uint16       // GPS HDOP horizontal dilution of position. If unknown, set to: UINT16_MAX
	Epv               uint16       // GPS VDOP vertical dilution of position. If unknown, set to: UINT16_MAX
	Vel               uint16       // GPS ground speed. If unknown, set to: UINT16_MAX
	Cog               uint16       // Course over ground (NOT heading, but direction of movement): 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	FixType           GPS_FIX_TYPE // GPS fix type.
	SatellitesVisible uint8        // Number of satellites visible. If unknown, set to 255
	DgpsNumch         uint8        // Number of DGPS satellites
}

// MsgID (generated function)
func (m *Gps2Raw) MsgID() message.MessageID {
	return MSG_ID_GPS2_RAW
}

// String (generated function)
func (m *Gps2Raw) String() string {
	return fmt.Sprintf(
		"&common.Gps2Raw{ TimeUsec: %+v, Lat: %+v, Lon: %+v, Alt: %+v, DgpsAge: %+v, Eph: %+v, Epv: %+v, Vel: %+v, Cog: %+v, FixType: %+v, SatellitesVisible: %+v, DgpsNumch: %+v }",
		m.TimeUsec,
		m.Lat,
		m.Lon,
		m.Alt,
		m.DgpsAge,
		m.Eph,
		m.Epv,
		m.Vel,
		m.Cog,
		m.FixType,
		m.SatellitesVisible,
		m.DgpsNumch,
	)
}

// Marshal (generated function)
func (m *Gps2Raw) Marshal() ([]byte, error) {
	payload := make([]byte, 35)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.Alt))
	binary.LittleEndian.PutUint32(payload[20:], uint32(m.DgpsAge))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.Eph))
	binary.LittleEndian.PutUint16(payload[26:], uint16(m.Epv))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Vel))
	binary.LittleEndian.PutUint16(payload[30:], uint16(m.Cog))
	payload[32] = byte(m.FixType)
	payload[33] = byte(m.SatellitesVisible)
	payload[34] = byte(m.DgpsNumch)
	return payload, nil
}

// Unmarshal (generated function)
func (m *Gps2Raw) Unmarshal(data []byte) error {
	payload := make([]byte, 35)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[12:]))
	m.Alt = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.DgpsAge = uint32(binary.LittleEndian.Uint32(payload[20:]))
	m.Eph = uint16(binary.LittleEndian.Uint16(payload[24:]))
	m.Epv = uint16(binary.LittleEndian.Uint16(payload[26:]))
	m.Vel = uint16(binary.LittleEndian.Uint16(payload[28:]))
	m.Cog = uint16(binary.LittleEndian.Uint16(payload[30:]))
	m.FixType = GPS_FIX_TYPE(payload[32])
	m.SatellitesVisible = uint8(payload[33])
	m.DgpsNumch = uint8(payload[34])
	return nil
}

// PowerStatus struct (generated typeinfo)
// Power supply status
type PowerStatus struct {
	Vcc    uint16           // 5V rail voltage.
	Vservo uint16           // Servo rail voltage.
	Flags  MAV_POWER_STATUS // Bitmap of power supply status flags.
}

// MsgID (generated function)
func (m *PowerStatus) MsgID() message.MessageID {
	return MSG_ID_POWER_STATUS
}

// String (generated function)
func (m *PowerStatus) String() string {
	return fmt.Sprintf(
		"&common.PowerStatus{ Vcc: %+v, Vservo: %+v, Flags: %+v (%016b) }",
		m.Vcc,
		m.Vservo,
		m.Flags.Bitmask(), uint64(m.Flags),
	)
}

// Marshal (generated function)
func (m *PowerStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 6)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Vcc))
	binary.LittleEndian.PutUint16(payload[2:], uint16(m.Vservo))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Flags))
	return payload, nil
}

// Unmarshal (generated function)
func (m *PowerStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 6)
	copy(payload[0:], data)
	m.Vcc = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.Vservo = uint16(binary.LittleEndian.Uint16(payload[2:]))
	m.Flags = MAV_POWER_STATUS(binary.LittleEndian.Uint16(payload[4:]))
	return nil
}

// SerialControl struct (generated typeinfo)
// Control a serial port. This can be used for raw access to an onboard serial peripheral such as a GPS or telemetry radio. It is designed to make it possible to update the devices firmware via MAVLink messages or change the devices settings. A message with zero bytes can be used to change just the baudrate.
type SerialControl struct {
	Baudrate uint32              // Baudrate of transfer. Zero means no change.
	Timeout  uint16              // Timeout for reply data
	Device   SERIAL_CONTROL_DEV  // Serial control device type.
	Flags    SERIAL_CONTROL_FLAG // Bitmap of serial control flags.
	Count    uint8               // how many bytes in this transfer
	Data     []uint8             `len:"70" ` // serial data
}

// MsgID (generated function)
func (m *SerialControl) MsgID() message.MessageID {
	return MSG_ID_SERIAL_CONTROL
}

// String (generated function)
func (m *SerialControl) String() string {
	return fmt.Sprintf(
		"&common.SerialControl{ Baudrate: %+v, Timeout: %+v, Device: %+v, Flags: %+v (%08b), Count: %+v, Data: %0X (\"%s\") }",
		m.Baudrate,
		m.Timeout,
		m.Device,
		m.Flags.Bitmask(), uint64(m.Flags),
		m.Count,
		m.Data, string(m.Data[:]),
	)
}

// Marshal (generated function)
func (m *SerialControl) Marshal() ([]byte, error) {
	payload := make([]byte, 79)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Baudrate))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Timeout))
	payload[6] = byte(m.Device)
	payload[7] = byte(m.Flags)
	payload[8] = byte(m.Count)
	copy(payload[9:], m.Data)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SerialControl) Unmarshal(data []byte) error {
	payload := make([]byte, 79)
	copy(payload[0:], data)
	m.Baudrate = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Timeout = uint16(binary.LittleEndian.Uint16(payload[4:]))
	m.Device = SERIAL_CONTROL_DEV(payload[6])
	m.Flags = SERIAL_CONTROL_FLAG(payload[7])
	m.Count = uint8(payload[8])
	copy(m.Data[:], payload[9:79])
	return nil
}

// GpsRtk struct (generated typeinfo)
// RTK GPS data. Gives information on the relative baseline calculation the GPS is reporting
type GpsRtk struct {
	TimeLastBaselineMs uint32                         // Time since boot of last baseline message received.
	Tow                uint32                         // GPS Time of Week of last baseline
	BaselineAMm        int32                          // Current baseline in ECEF x or NED north component.
	BaselineBMm        int32                          // Current baseline in ECEF y or NED east component.
	BaselineCMm        int32                          // Current baseline in ECEF z or NED down component.
	Accuracy           uint32                         // Current estimate of baseline accuracy.
	IarNumHypotheses   int32                          // Current number of integer ambiguity hypotheses.
	Wn                 uint16                         // GPS Week Number of last baseline
	RtkReceiverID      uint8                          // Identification of connected RTK receiver.
	RtkHealth          uint8                          // GPS-specific health report for RTK data.
	RtkRate            uint8                          // Rate of baseline messages being received by GPS
	Nsats              uint8                          // Current number of sats used for RTK calculation.
	BaselineCoordsType RTK_BASELINE_COORDINATE_SYSTEM // Coordinate system of baseline
}

// MsgID (generated function)
func (m *GpsRtk) MsgID() message.MessageID {
	return MSG_ID_GPS_RTK
}

// String (generated function)
func (m *GpsRtk) String() string {
	return fmt.Sprintf(
		"&common.GpsRtk{ TimeLastBaselineMs: %+v, Tow: %+v, BaselineAMm: %+v, BaselineBMm: %+v, BaselineCMm: %+v, Accuracy: %+v, IarNumHypotheses: %+v, Wn: %+v, RtkReceiverID: %+v, RtkHealth: %+v, RtkRate: %+v, Nsats: %+v, BaselineCoordsType: %+v }",
		m.TimeLastBaselineMs,
		m.Tow,
		m.BaselineAMm,
		m.BaselineBMm,
		m.BaselineCMm,
		m.Accuracy,
		m.IarNumHypotheses,
		m.Wn,
		m.RtkReceiverID,
		m.RtkHealth,
		m.RtkRate,
		m.Nsats,
		m.BaselineCoordsType,
	)
}

// Marshal (generated function)
func (m *GpsRtk) Marshal() ([]byte, error) {
	payload := make([]byte, 35)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeLastBaselineMs))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Tow))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.BaselineAMm))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.BaselineBMm))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.BaselineCMm))
	binary.LittleEndian.PutUint32(payload[20:], uint32(m.Accuracy))
	binary.LittleEndian.PutUint32(payload[24:], uint32(m.IarNumHypotheses))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Wn))
	payload[30] = byte(m.RtkReceiverID)
	payload[31] = byte(m.RtkHealth)
	payload[32] = byte(m.RtkRate)
	payload[33] = byte(m.Nsats)
	payload[34] = byte(m.BaselineCoordsType)
	return payload, nil
}

// Unmarshal (generated function)
func (m *GpsRtk) Unmarshal(data []byte) error {
	payload := make([]byte, 35)
	copy(payload[0:], data)
	m.TimeLastBaselineMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Tow = uint32(binary.LittleEndian.Uint32(payload[4:]))
	m.BaselineAMm = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.BaselineBMm = int32(binary.LittleEndian.Uint32(payload[12:]))
	m.BaselineCMm = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.Accuracy = uint32(binary.LittleEndian.Uint32(payload[20:]))
	m.IarNumHypotheses = int32(binary.LittleEndian.Uint32(payload[24:]))
	m.Wn = uint16(binary.LittleEndian.Uint16(payload[28:]))
	m.RtkReceiverID = uint8(payload[30])
	m.RtkHealth = uint8(payload[31])
	m.RtkRate = uint8(payload[32])
	m.Nsats = uint8(payload[33])
	m.BaselineCoordsType = RTK_BASELINE_COORDINATE_SYSTEM(payload[34])
	return nil
}

// Gps2Rtk struct (generated typeinfo)
// RTK GPS data. Gives information on the relative baseline calculation the GPS is reporting
type Gps2Rtk struct {
	TimeLastBaselineMs uint32                         // Time since boot of last baseline message received.
	Tow                uint32                         // GPS Time of Week of last baseline
	BaselineAMm        int32                          // Current baseline in ECEF x or NED north component.
	BaselineBMm        int32                          // Current baseline in ECEF y or NED east component.
	BaselineCMm        int32                          // Current baseline in ECEF z or NED down component.
	Accuracy           uint32                         // Current estimate of baseline accuracy.
	IarNumHypotheses   int32                          // Current number of integer ambiguity hypotheses.
	Wn                 uint16                         // GPS Week Number of last baseline
	RtkReceiverID      uint8                          // Identification of connected RTK receiver.
	RtkHealth          uint8                          // GPS-specific health report for RTK data.
	RtkRate            uint8                          // Rate of baseline messages being received by GPS
	Nsats              uint8                          // Current number of sats used for RTK calculation.
	BaselineCoordsType RTK_BASELINE_COORDINATE_SYSTEM // Coordinate system of baseline
}

// MsgID (generated function)
func (m *Gps2Rtk) MsgID() message.MessageID {
	return MSG_ID_GPS2_RTK
}

// String (generated function)
func (m *Gps2Rtk) String() string {
	return fmt.Sprintf(
		"&common.Gps2Rtk{ TimeLastBaselineMs: %+v, Tow: %+v, BaselineAMm: %+v, BaselineBMm: %+v, BaselineCMm: %+v, Accuracy: %+v, IarNumHypotheses: %+v, Wn: %+v, RtkReceiverID: %+v, RtkHealth: %+v, RtkRate: %+v, Nsats: %+v, BaselineCoordsType: %+v }",
		m.TimeLastBaselineMs,
		m.Tow,
		m.BaselineAMm,
		m.BaselineBMm,
		m.BaselineCMm,
		m.Accuracy,
		m.IarNumHypotheses,
		m.Wn,
		m.RtkReceiverID,
		m.RtkHealth,
		m.RtkRate,
		m.Nsats,
		m.BaselineCoordsType,
	)
}

// Marshal (generated function)
func (m *Gps2Rtk) Marshal() ([]byte, error) {
	payload := make([]byte, 35)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeLastBaselineMs))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Tow))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.BaselineAMm))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.BaselineBMm))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.BaselineCMm))
	binary.LittleEndian.PutUint32(payload[20:], uint32(m.Accuracy))
	binary.LittleEndian.PutUint32(payload[24:], uint32(m.IarNumHypotheses))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Wn))
	payload[30] = byte(m.RtkReceiverID)
	payload[31] = byte(m.RtkHealth)
	payload[32] = byte(m.RtkRate)
	payload[33] = byte(m.Nsats)
	payload[34] = byte(m.BaselineCoordsType)
	return payload, nil
}

// Unmarshal (generated function)
func (m *Gps2Rtk) Unmarshal(data []byte) error {
	payload := make([]byte, 35)
	copy(payload[0:], data)
	m.TimeLastBaselineMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Tow = uint32(binary.LittleEndian.Uint32(payload[4:]))
	m.BaselineAMm = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.BaselineBMm = int32(binary.LittleEndian.Uint32(payload[12:]))
	m.BaselineCMm = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.Accuracy = uint32(binary.LittleEndian.Uint32(payload[20:]))
	m.IarNumHypotheses = int32(binary.LittleEndian.Uint32(payload[24:]))
	m.Wn = uint16(binary.LittleEndian.Uint16(payload[28:]))
	m.RtkReceiverID = uint8(payload[30])
	m.RtkHealth = uint8(payload[31])
	m.RtkRate = uint8(payload[32])
	m.Nsats = uint8(payload[33])
	m.BaselineCoordsType = RTK_BASELINE_COORDINATE_SYSTEM(payload[34])
	return nil
}

// ScaledImu3 struct (generated typeinfo)
// The RAW IMU readings for 3rd 9DOF sensor setup. This message should contain the scaled values to the described units
type ScaledImu3 struct {
	TimeBootMs uint32 // Timestamp (time since system boot).
	Xacc       int16  // X acceleration
	Yacc       int16  // Y acceleration
	Zacc       int16  // Z acceleration
	Xgyro      int16  // Angular speed around X axis
	Ygyro      int16  // Angular speed around Y axis
	Zgyro      int16  // Angular speed around Z axis
	Xmag       int16  // X Magnetic field
	Ymag       int16  // Y Magnetic field
	Zmag       int16  // Z Magnetic field
}

// MsgID (generated function)
func (m *ScaledImu3) MsgID() message.MessageID {
	return MSG_ID_SCALED_IMU3
}

// String (generated function)
func (m *ScaledImu3) String() string {
	return fmt.Sprintf(
		"&common.ScaledImu3{ TimeBootMs: %+v, Xacc: %+v, Yacc: %+v, Zacc: %+v, Xgyro: %+v, Ygyro: %+v, Zgyro: %+v, Xmag: %+v, Ymag: %+v, Zmag: %+v }",
		m.TimeBootMs,
		m.Xacc,
		m.Yacc,
		m.Zacc,
		m.Xgyro,
		m.Ygyro,
		m.Zgyro,
		m.Xmag,
		m.Ymag,
		m.Zmag,
	)
}

// Marshal (generated function)
func (m *ScaledImu3) Marshal() ([]byte, error) {
	payload := make([]byte, 22)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Xacc))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.Yacc))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Zacc))
	binary.LittleEndian.PutUint16(payload[10:], uint16(m.Xgyro))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Ygyro))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Zgyro))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Xmag))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Ymag))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Zmag))
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScaledImu3) Unmarshal(data []byte) error {
	payload := make([]byte, 22)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Xacc = int16(binary.LittleEndian.Uint16(payload[4:]))
	m.Yacc = int16(binary.LittleEndian.Uint16(payload[6:]))
	m.Zacc = int16(binary.LittleEndian.Uint16(payload[8:]))
	m.Xgyro = int16(binary.LittleEndian.Uint16(payload[10:]))
	m.Ygyro = int16(binary.LittleEndian.Uint16(payload[12:]))
	m.Zgyro = int16(binary.LittleEndian.Uint16(payload[14:]))
	m.Xmag = int16(binary.LittleEndian.Uint16(payload[16:]))
	m.Ymag = int16(binary.LittleEndian.Uint16(payload[18:]))
	m.Zmag = int16(binary.LittleEndian.Uint16(payload[20:]))
	return nil
}

// DataTransmissionHandshake struct (generated typeinfo)
// Handshake message to initiate, control and stop image streaming when using the Image Transmission Protocol: https://mavlink.io/en/services/image_transmission.html.
type DataTransmissionHandshake struct {
	Size       uint32                   // total data size (set on ACK only).
	Width      uint16                   // Width of a matrix or image.
	Height     uint16                   // Height of a matrix or image.
	Packets    uint16                   // Number of packets being sent (set on ACK only).
	Type       MAVLINK_DATA_STREAM_TYPE // Type of requested/acknowledged data.
	Payload    uint8                    // Payload size per packet (normally 253 byte, see DATA field size in message ENCAPSULATED_DATA) (set on ACK only).
	JpgQuality uint8                    // JPEG quality. Values: [1-100].
}

// MsgID (generated function)
func (m *DataTransmissionHandshake) MsgID() message.MessageID {
	return MSG_ID_DATA_TRANSMISSION_HANDSHAKE
}

// String (generated function)
func (m *DataTransmissionHandshake) String() string {
	return fmt.Sprintf(
		"&common.DataTransmissionHandshake{ Size: %+v, Width: %+v, Height: %+v, Packets: %+v, Type: %+v, Payload: %+v, JpgQuality: %+v }",
		m.Size,
		m.Width,
		m.Height,
		m.Packets,
		m.Type,
		m.Payload,
		m.JpgQuality,
	)
}

// Marshal (generated function)
func (m *DataTransmissionHandshake) Marshal() ([]byte, error) {
	payload := make([]byte, 13)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Size))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.Width))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.Height))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Packets))
	payload[10] = byte(m.Type)
	payload[11] = byte(m.Payload)
	payload[12] = byte(m.JpgQuality)
	return payload, nil
}

// Unmarshal (generated function)
func (m *DataTransmissionHandshake) Unmarshal(data []byte) error {
	payload := make([]byte, 13)
	copy(payload[0:], data)
	m.Size = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Width = uint16(binary.LittleEndian.Uint16(payload[4:]))
	m.Height = uint16(binary.LittleEndian.Uint16(payload[6:]))
	m.Packets = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.Type = MAVLINK_DATA_STREAM_TYPE(payload[10])
	m.Payload = uint8(payload[11])
	m.JpgQuality = uint8(payload[12])
	return nil
}

// EncapsulatedData struct (generated typeinfo)
// Data packet for images sent using the Image Transmission Protocol: https://mavlink.io/en/services/image_transmission.html.
type EncapsulatedData struct {
	Seqnr uint16  // sequence number (starting with 0 on every transmission)
	Data  []uint8 `len:"253" ` // image data bytes
}

// MsgID (generated function)
func (m *EncapsulatedData) MsgID() message.MessageID {
	return MSG_ID_ENCAPSULATED_DATA
}

// String (generated function)
func (m *EncapsulatedData) String() string {
	return fmt.Sprintf(
		"&common.EncapsulatedData{ Seqnr: %+v, Data: %0X (\"%s\") }",
		m.Seqnr,
		m.Data, string(m.Data[:]),
	)
}

// Marshal (generated function)
func (m *EncapsulatedData) Marshal() ([]byte, error) {
	payload := make([]byte, 255)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Seqnr))
	copy(payload[2:], m.Data)
	return payload, nil
}

// Unmarshal (generated function)
func (m *EncapsulatedData) Unmarshal(data []byte) error {
	payload := make([]byte, 255)
	copy(payload[0:], data)
	m.Seqnr = uint16(binary.LittleEndian.Uint16(payload[0:]))
	copy(m.Data[:], payload[2:255])
	return nil
}

// DistanceSensor struct (generated typeinfo)
// Distance sensor information for an onboard rangefinder.
type DistanceSensor struct {
	TimeBootMs      uint32                 // Timestamp (time since system boot).
	MinDistance     uint16                 // Minimum distance the sensor can measure
	MaxDistance     uint16                 // Maximum distance the sensor can measure
	CurrentDistance uint16                 // Current distance reading
	Type            MAV_DISTANCE_SENSOR    // Type of distance sensor.
	ID              uint8                  // Onboard ID of the sensor
	Orientation     MAV_SENSOR_ORIENTATION // Direction the sensor faces. downward-facing: ROTATION_PITCH_270, upward-facing: ROTATION_PITCH_90, backward-facing: ROTATION_PITCH_180, forward-facing: ROTATION_NONE, left-facing: ROTATION_YAW_90, right-facing: ROTATION_YAW_270
	Covariance      uint8                  // Measurement variance. Max standard deviation is 6cm. 255 if unknown.
}

// MsgID (generated function)
func (m *DistanceSensor) MsgID() message.MessageID {
	return MSG_ID_DISTANCE_SENSOR
}

// String (generated function)
func (m *DistanceSensor) String() string {
	return fmt.Sprintf(
		"&common.DistanceSensor{ TimeBootMs: %+v, MinDistance: %+v, MaxDistance: %+v, CurrentDistance: %+v, Type: %+v, ID: %+v, Orientation: %+v, Covariance: %+v }",
		m.TimeBootMs,
		m.MinDistance,
		m.MaxDistance,
		m.CurrentDistance,
		m.Type,
		m.ID,
		m.Orientation,
		m.Covariance,
	)
}

// Marshal (generated function)
func (m *DistanceSensor) Marshal() ([]byte, error) {
	payload := make([]byte, 14)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.MinDistance))
	binary.LittleEndian.PutUint16(payload[6:], uint16(m.MaxDistance))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.CurrentDistance))
	payload[10] = byte(m.Type)
	payload[11] = byte(m.ID)
	payload[12] = byte(m.Orientation)
	payload[13] = byte(m.Covariance)
	return payload, nil
}

// Unmarshal (generated function)
func (m *DistanceSensor) Unmarshal(data []byte) error {
	payload := make([]byte, 14)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.MinDistance = uint16(binary.LittleEndian.Uint16(payload[4:]))
	m.MaxDistance = uint16(binary.LittleEndian.Uint16(payload[6:]))
	m.CurrentDistance = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.Type = MAV_DISTANCE_SENSOR(payload[10])
	m.ID = uint8(payload[11])
	m.Orientation = MAV_SENSOR_ORIENTATION(payload[12])
	m.Covariance = uint8(payload[13])
	return nil
}

// TerrainRequest struct (generated typeinfo)
// Request for terrain data and terrain status. See terrain protocol docs: https://mavlink.io/en/services/terrain.html
type TerrainRequest struct {
	Mask        uint64 // Bitmask of requested 4x4 grids (row major 8x7 array of grids, 56 bits)
	Lat         int32  // Latitude of SW corner of first grid
	Lon         int32  // Longitude of SW corner of first grid
	GridSpacing uint16 // Grid spacing
}

// MsgID (generated function)
func (m *TerrainRequest) MsgID() message.MessageID {
	return MSG_ID_TERRAIN_REQUEST
}

// String (generated function)
func (m *TerrainRequest) String() string {
	return fmt.Sprintf(
		"&common.TerrainRequest{ Mask: %+v, Lat: %+v, Lon: %+v, GridSpacing: %+v }",
		m.Mask,
		m.Lat,
		m.Lon,
		m.GridSpacing,
	)
}

// Marshal (generated function)
func (m *TerrainRequest) Marshal() ([]byte, error) {
	payload := make([]byte, 18)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Mask))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.Lon))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.GridSpacing))
	return payload, nil
}

// Unmarshal (generated function)
func (m *TerrainRequest) Unmarshal(data []byte) error {
	payload := make([]byte, 18)
	copy(payload[0:], data)
	m.Mask = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[12:]))
	m.GridSpacing = uint16(binary.LittleEndian.Uint16(payload[16:]))
	return nil
}

// TerrainData struct (generated typeinfo)
// Terrain data sent from GCS. The lat/lon and grid_spacing must be the same as a lat/lon from a TERRAIN_REQUEST. See terrain protocol docs: https://mavlink.io/en/services/terrain.html
type TerrainData struct {
	Lat         int32   // Latitude of SW corner of first grid
	Lon         int32   // Longitude of SW corner of first grid
	GridSpacing uint16  // Grid spacing
	Data        []int16 `len:"16" ` // Terrain data MSL
	Gridbit     uint8   // bit within the terrain request mask
}

// MsgID (generated function)
func (m *TerrainData) MsgID() message.MessageID {
	return MSG_ID_TERRAIN_DATA
}

// String (generated function)
func (m *TerrainData) String() string {
	return fmt.Sprintf(
		"&common.TerrainData{ Lat: %+v, Lon: %+v, GridSpacing: %+v, Data: %+v, Gridbit: %+v }",
		m.Lat,
		m.Lon,
		m.GridSpacing,
		m.Data,
		m.Gridbit,
	)
}

// Marshal (generated function)
func (m *TerrainData) Marshal() ([]byte, error) {
	payload := make([]byte, 43)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Lon))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.GridSpacing))
	for i, v := range m.Data {
		binary.LittleEndian.PutUint16(payload[10+i*2:], uint16(v))
	}
	payload[42] = byte(m.Gridbit)
	return payload, nil
}

// Unmarshal (generated function)
func (m *TerrainData) Unmarshal(data []byte) error {
	payload := make([]byte, 43)
	copy(payload[0:], data)
	m.Lat = int32(binary.LittleEndian.Uint32(payload[0:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.GridSpacing = uint16(binary.LittleEndian.Uint16(payload[8:]))
	for i := 0; i < len(m.Data); i++ {
		m.Data[i] = int16(binary.LittleEndian.Uint16(payload[10+i*2:]))
	}
	m.Gridbit = uint8(payload[42])
	return nil
}

// TerrainCheck struct (generated typeinfo)
// Request that the vehicle report terrain height at the given location (expected response is a TERRAIN_REPORT). Used by GCS to check if vehicle has all terrain data needed for a mission.
type TerrainCheck struct {
	Lat int32 // Latitude
	Lon int32 // Longitude
}

// MsgID (generated function)
func (m *TerrainCheck) MsgID() message.MessageID {
	return MSG_ID_TERRAIN_CHECK
}

// String (generated function)
func (m *TerrainCheck) String() string {
	return fmt.Sprintf(
		"&common.TerrainCheck{ Lat: %+v, Lon: %+v }",
		m.Lat,
		m.Lon,
	)
}

// Marshal (generated function)
func (m *TerrainCheck) Marshal() ([]byte, error) {
	payload := make([]byte, 8)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Lon))
	return payload, nil
}

// Unmarshal (generated function)
func (m *TerrainCheck) Unmarshal(data []byte) error {
	payload := make([]byte, 8)
	copy(payload[0:], data)
	m.Lat = int32(binary.LittleEndian.Uint32(payload[0:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[4:]))
	return nil
}

// TerrainReport struct (generated typeinfo)
// Streamed from drone to report progress of terrain map download (initiated by TERRAIN_REQUEST), or sent as a response to a TERRAIN_CHECK request. See terrain protocol docs: https://mavlink.io/en/services/terrain.html
type TerrainReport struct {
	Lat           int32   // Latitude
	Lon           int32   // Longitude
	TerrainHeight float32 // Terrain height MSL
	CurrentHeight float32 // Current vehicle height above lat/lon terrain height
	Spacing       uint16  // grid spacing (zero if terrain at this location unavailable)
	Pending       uint16  // Number of 4x4 terrain blocks waiting to be received or read from disk
	Loaded        uint16  // Number of 4x4 terrain blocks in memory
}

// MsgID (generated function)
func (m *TerrainReport) MsgID() message.MessageID {
	return MSG_ID_TERRAIN_REPORT
}

// String (generated function)
func (m *TerrainReport) String() string {
	return fmt.Sprintf(
		"&common.TerrainReport{ Lat: %+v, Lon: %+v, TerrainHeight: %+v, CurrentHeight: %+v, Spacing: %+v, Pending: %+v, Loaded: %+v }",
		m.Lat,
		m.Lon,
		m.TerrainHeight,
		m.CurrentHeight,
		m.Spacing,
		m.Pending,
		m.Loaded,
	)
}

// Marshal (generated function)
func (m *TerrainReport) Marshal() ([]byte, error) {
	payload := make([]byte, 22)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.TerrainHeight))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.CurrentHeight))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Spacing))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.Pending))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Loaded))
	return payload, nil
}

// Unmarshal (generated function)
func (m *TerrainReport) Unmarshal(data []byte) error {
	payload := make([]byte, 22)
	copy(payload[0:], data)
	m.Lat = int32(binary.LittleEndian.Uint32(payload[0:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.TerrainHeight = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.CurrentHeight = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Spacing = uint16(binary.LittleEndian.Uint16(payload[16:]))
	m.Pending = uint16(binary.LittleEndian.Uint16(payload[18:]))
	m.Loaded = uint16(binary.LittleEndian.Uint16(payload[20:]))
	return nil
}

// ScaledPressure2 struct (generated typeinfo)
// Barometer readings for 2nd barometer
type ScaledPressure2 struct {
	TimeBootMs  uint32  // Timestamp (time since system boot).
	PressAbs    float32 // Absolute pressure
	PressDiff   float32 // Differential pressure
	Temperature int16   // Absolute pressure temperature
}

// MsgID (generated function)
func (m *ScaledPressure2) MsgID() message.MessageID {
	return MSG_ID_SCALED_PRESSURE2
}

// String (generated function)
func (m *ScaledPressure2) String() string {
	return fmt.Sprintf(
		"&common.ScaledPressure2{ TimeBootMs: %+v, PressAbs: %+v, PressDiff: %+v, Temperature: %+v }",
		m.TimeBootMs,
		m.PressAbs,
		m.PressDiff,
		m.Temperature,
	)
}

// Marshal (generated function)
func (m *ScaledPressure2) Marshal() ([]byte, error) {
	payload := make([]byte, 14)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.PressAbs))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.PressDiff))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Temperature))
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScaledPressure2) Unmarshal(data []byte) error {
	payload := make([]byte, 14)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.PressAbs = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.PressDiff = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Temperature = int16(binary.LittleEndian.Uint16(payload[12:]))
	return nil
}

// AttPosMocap struct (generated typeinfo)
// Motion capture attitude and position
type AttPosMocap struct {
	TimeUsec uint64    // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Q        []float32 `len:"4" ` // Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	X        float32   // X position (NED)
	Y        float32   // Y position (NED)
	Z        float32   // Z position (NED)
}

// MsgID (generated function)
func (m *AttPosMocap) MsgID() message.MessageID {
	return MSG_ID_ATT_POS_MOCAP
}

// String (generated function)
func (m *AttPosMocap) String() string {
	return fmt.Sprintf(
		"&common.AttPosMocap{ TimeUsec: %+v, Q: %+v, X: %+v, Y: %+v, Z: %+v }",
		m.TimeUsec,
		m.Q,
		m.X,
		m.Y,
		m.Z,
	)
}

// Marshal (generated function)
func (m *AttPosMocap) Marshal() ([]byte, error) {
	payload := make([]byte, 36)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	for i, v := range m.Q {
		binary.LittleEndian.PutUint32(payload[8+i*4:], math.Float32bits(v))
	}
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Z))
	return payload, nil
}

// Unmarshal (generated function)
func (m *AttPosMocap) Unmarshal(data []byte) error {
	payload := make([]byte, 36)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	for i := 0; i < len(m.Q); i++ {
		m.Q[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[8+i*4:]))
	}
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	return nil
}

// SetActuatorControlTarget struct (generated typeinfo)
// Set the vehicle attitude and body angular rates.
type SetActuatorControlTarget struct {
	TimeUsec        uint64    // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Controls        []float32 `len:"8" ` // Actuator controls. Normed to -1..+1 where 0 is neutral position. Throttle for single rotation direction motors is 0..1, negative range for reverse direction. Standard mapping for attitude controls (group 0): (index 0-7): roll, pitch, yaw, throttle, flaps, spoilers, airbrakes, landing gear. Load a pass-through mixer to repurpose them as generic outputs.
	GroupMlx        uint8     // Actuator group. The "_mlx" indicates this is a multi-instance message and a MAVLink parser should use this field to difference between instances.
	TargetSystem    uint8     // System ID
	TargetComponent uint8     // Component ID
}

// MsgID (generated function)
func (m *SetActuatorControlTarget) MsgID() message.MessageID {
	return MSG_ID_SET_ACTUATOR_CONTROL_TARGET
}

// String (generated function)
func (m *SetActuatorControlTarget) String() string {
	return fmt.Sprintf(
		"&common.SetActuatorControlTarget{ TimeUsec: %+v, Controls: %+v, GroupMlx: %+v, TargetSystem: %+v, TargetComponent: %+v }",
		m.TimeUsec,
		m.Controls,
		m.GroupMlx,
		m.TargetSystem,
		m.TargetComponent,
	)
}

// Marshal (generated function)
func (m *SetActuatorControlTarget) Marshal() ([]byte, error) {
	payload := make([]byte, 43)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	for i, v := range m.Controls {
		binary.LittleEndian.PutUint32(payload[8+i*4:], math.Float32bits(v))
	}
	payload[40] = byte(m.GroupMlx)
	payload[41] = byte(m.TargetSystem)
	payload[42] = byte(m.TargetComponent)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SetActuatorControlTarget) Unmarshal(data []byte) error {
	payload := make([]byte, 43)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	for i := 0; i < len(m.Controls); i++ {
		m.Controls[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[8+i*4:]))
	}
	m.GroupMlx = uint8(payload[40])
	m.TargetSystem = uint8(payload[41])
	m.TargetComponent = uint8(payload[42])
	return nil
}

// ActuatorControlTarget struct (generated typeinfo)
// Set the vehicle attitude and body angular rates.
type ActuatorControlTarget struct {
	TimeUsec uint64    // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Controls []float32 `len:"8" ` // Actuator controls. Normed to -1..+1 where 0 is neutral position. Throttle for single rotation direction motors is 0..1, negative range for reverse direction. Standard mapping for attitude controls (group 0): (index 0-7): roll, pitch, yaw, throttle, flaps, spoilers, airbrakes, landing gear. Load a pass-through mixer to repurpose them as generic outputs.
	GroupMlx uint8     // Actuator group. The "_mlx" indicates this is a multi-instance message and a MAVLink parser should use this field to difference between instances.
}

// MsgID (generated function)
func (m *ActuatorControlTarget) MsgID() message.MessageID {
	return MSG_ID_ACTUATOR_CONTROL_TARGET
}

// String (generated function)
func (m *ActuatorControlTarget) String() string {
	return fmt.Sprintf(
		"&common.ActuatorControlTarget{ TimeUsec: %+v, Controls: %+v, GroupMlx: %+v }",
		m.TimeUsec,
		m.Controls,
		m.GroupMlx,
	)
}

// Marshal (generated function)
func (m *ActuatorControlTarget) Marshal() ([]byte, error) {
	payload := make([]byte, 41)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	for i, v := range m.Controls {
		binary.LittleEndian.PutUint32(payload[8+i*4:], math.Float32bits(v))
	}
	payload[40] = byte(m.GroupMlx)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ActuatorControlTarget) Unmarshal(data []byte) error {
	payload := make([]byte, 41)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	for i := 0; i < len(m.Controls); i++ {
		m.Controls[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[8+i*4:]))
	}
	m.GroupMlx = uint8(payload[40])
	return nil
}

// Altitude struct (generated typeinfo)
// The current system altitude.
type Altitude struct {
	TimeUsec          uint64  // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	AltitudeMonotonic float32 // This altitude measure is initialized on system boot and monotonic (it is never reset, but represents the local altitude change). The only guarantee on this field is that it will never be reset and is consistent within a flight. The recommended value for this field is the uncorrected barometric altitude at boot time. This altitude will also drift and vary between flights.
	AltitudeAmsl      float32 // This altitude measure is strictly above mean sea level and might be non-monotonic (it might reset on events like GPS lock or when a new QNH value is set). It should be the altitude to which global altitude waypoints are compared to. Note that it is *not* the GPS altitude, however, most GPS modules already output MSL by default and not the WGS84 altitude.
	AltitudeLocal     float32 // This is the local altitude in the local coordinate frame. It is not the altitude above home, but in reference to the coordinate origin (0, 0, 0). It is up-positive.
	AltitudeRelative  float32 // This is the altitude above the home position. It resets on each change of the current home position.
	AltitudeTerrain   float32 // This is the altitude above terrain. It might be fed by a terrain database or an altimeter. Values smaller than -1000 should be interpreted as unknown.
	BottomClearance   float32 // This is not the altitude, but the clear space below the system according to the fused clearance estimate. It generally should max out at the maximum range of e.g. the laser altimeter. It is generally a moving target. A negative value indicates no measurement available.
}

// MsgID (generated function)
func (m *Altitude) MsgID() message.MessageID {
	return MSG_ID_ALTITUDE
}

// String (generated function)
func (m *Altitude) String() string {
	return fmt.Sprintf(
		"&common.Altitude{ TimeUsec: %+v, AltitudeMonotonic: %+v, AltitudeAmsl: %+v, AltitudeLocal: %+v, AltitudeRelative: %+v, AltitudeTerrain: %+v, BottomClearance: %+v }",
		m.TimeUsec,
		m.AltitudeMonotonic,
		m.AltitudeAmsl,
		m.AltitudeLocal,
		m.AltitudeRelative,
		m.AltitudeTerrain,
		m.BottomClearance,
	)
}

// Marshal (generated function)
func (m *Altitude) Marshal() ([]byte, error) {
	payload := make([]byte, 32)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.AltitudeMonotonic))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.AltitudeAmsl))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.AltitudeLocal))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.AltitudeRelative))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.AltitudeTerrain))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.BottomClearance))
	return payload, nil
}

// Unmarshal (generated function)
func (m *Altitude) Unmarshal(data []byte) error {
	payload := make([]byte, 32)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.AltitudeMonotonic = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.AltitudeAmsl = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.AltitudeLocal = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.AltitudeRelative = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.AltitudeTerrain = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.BottomClearance = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	return nil
}

// ResourceRequest struct (generated typeinfo)
// The autopilot is requesting a resource (file, binary, other type of data)
type ResourceRequest struct {
	RequestID    uint8   // Request ID. This ID should be re-used when sending back URI contents
	URIType      uint8   // The type of requested URI. 0 = a file via URL. 1 = a UAVCAN binary
	URI          []uint8 `len:"120" ` // The requested unique resource identifier (URI). It is not necessarily a straight domain name (depends on the URI type enum)
	TransferType uint8   // The way the autopilot wants to receive the URI. 0 = MAVLink FTP. 1 = binary stream.
	Storage      []uint8 `len:"120" ` // The storage path the autopilot wants the URI to be stored in. Will only be valid if the transfer_type has a storage associated (e.g. MAVLink FTP).
}

// MsgID (generated function)
func (m *ResourceRequest) MsgID() message.MessageID {
	return MSG_ID_RESOURCE_REQUEST
}

// String (generated function)
func (m *ResourceRequest) String() string {
	return fmt.Sprintf(
		"&common.ResourceRequest{ RequestID: %+v, URIType: %+v, URI: %0X (\"%s\"), TransferType: %+v, Storage: %0X (\"%s\") }",
		m.RequestID,
		m.URIType,
		m.URI, string(m.URI[:]),
		m.TransferType,
		m.Storage, string(m.Storage[:]),
	)
}

// Marshal (generated function)
func (m *ResourceRequest) Marshal() ([]byte, error) {
	payload := make([]byte, 243)
	payload[0] = byte(m.RequestID)
	payload[1] = byte(m.URIType)
	copy(payload[2:], m.URI)
	payload[122] = byte(m.TransferType)
	copy(payload[123:], m.Storage)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ResourceRequest) Unmarshal(data []byte) error {
	payload := make([]byte, 243)
	copy(payload[0:], data)
	m.RequestID = uint8(payload[0])
	m.URIType = uint8(payload[1])
	copy(m.URI[:], payload[2:122])
	m.TransferType = uint8(payload[122])
	copy(m.Storage[:], payload[123:243])
	return nil
}

// ScaledPressure3 struct (generated typeinfo)
// Barometer readings for 3rd barometer
type ScaledPressure3 struct {
	TimeBootMs  uint32  // Timestamp (time since system boot).
	PressAbs    float32 // Absolute pressure
	PressDiff   float32 // Differential pressure
	Temperature int16   // Absolute pressure temperature
}

// MsgID (generated function)
func (m *ScaledPressure3) MsgID() message.MessageID {
	return MSG_ID_SCALED_PRESSURE3
}

// String (generated function)
func (m *ScaledPressure3) String() string {
	return fmt.Sprintf(
		"&common.ScaledPressure3{ TimeBootMs: %+v, PressAbs: %+v, PressDiff: %+v, Temperature: %+v }",
		m.TimeBootMs,
		m.PressAbs,
		m.PressDiff,
		m.Temperature,
	)
}

// Marshal (generated function)
func (m *ScaledPressure3) Marshal() ([]byte, error) {
	payload := make([]byte, 14)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.PressAbs))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.PressDiff))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Temperature))
	return payload, nil
}

// Unmarshal (generated function)
func (m *ScaledPressure3) Unmarshal(data []byte) error {
	payload := make([]byte, 14)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.PressAbs = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.PressDiff = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Temperature = int16(binary.LittleEndian.Uint16(payload[12:]))
	return nil
}

// FollowTarget struct (generated typeinfo)
// Current motion information from a designated system
type FollowTarget struct {
	Timestamp       uint64    // Timestamp (time since system boot).
	CustomState     uint64    // button states or switches of a tracker device
	Lat             int32     // Latitude (WGS84)
	Lon             int32     // Longitude (WGS84)
	Alt             float32   // Altitude (MSL)
	Vel             []float32 `len:"3" ` // target velocity (0,0,0) for unknown
	Acc             []float32 `len:"3" ` // linear target acceleration (0,0,0) for unknown
	AttitudeQ       []float32 `len:"4" ` // (1 0 0 0 for unknown)
	Rates           []float32 `len:"3" ` // (0 0 0 for unknown)
	PositionCov     []float32 `len:"3" ` // eph epv
	EstCapabilities uint8     // bit positions for tracker reporting capabilities (POS = 0, VEL = 1, ACCEL = 2, ATT + RATES = 3)
}

// MsgID (generated function)
func (m *FollowTarget) MsgID() message.MessageID {
	return MSG_ID_FOLLOW_TARGET
}

// String (generated function)
func (m *FollowTarget) String() string {
	return fmt.Sprintf(
		"&common.FollowTarget{ Timestamp: %+v, CustomState: %+v, Lat: %+v, Lon: %+v, Alt: %+v, Vel: %+v, Acc: %+v, AttitudeQ: %+v, Rates: %+v, PositionCov: %+v, EstCapabilities: %+v }",
		m.Timestamp,
		m.CustomState,
		m.Lat,
		m.Lon,
		m.Alt,
		m.Vel,
		m.Acc,
		m.AttitudeQ,
		m.Rates,
		m.PositionCov,
		m.EstCapabilities,
	)
}

// Marshal (generated function)
func (m *FollowTarget) Marshal() ([]byte, error) {
	payload := make([]byte, 93)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint64(payload[8:], uint64(m.CustomState))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[20:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Alt))
	for i, v := range m.Vel {
		binary.LittleEndian.PutUint32(payload[28+i*4:], math.Float32bits(v))
	}
	for i, v := range m.Acc {
		binary.LittleEndian.PutUint32(payload[40+i*4:], math.Float32bits(v))
	}
	for i, v := range m.AttitudeQ {
		binary.LittleEndian.PutUint32(payload[52+i*4:], math.Float32bits(v))
	}
	for i, v := range m.Rates {
		binary.LittleEndian.PutUint32(payload[68+i*4:], math.Float32bits(v))
	}
	for i, v := range m.PositionCov {
		binary.LittleEndian.PutUint32(payload[80+i*4:], math.Float32bits(v))
	}
	payload[92] = byte(m.EstCapabilities)
	return payload, nil
}

// Unmarshal (generated function)
func (m *FollowTarget) Unmarshal(data []byte) error {
	payload := make([]byte, 93)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.CustomState = uint64(binary.LittleEndian.Uint64(payload[8:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[20:]))
	m.Alt = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	for i := 0; i < len(m.Vel); i++ {
		m.Vel[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[28+i*4:]))
	}
	for i := 0; i < len(m.Acc); i++ {
		m.Acc[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[40+i*4:]))
	}
	for i := 0; i < len(m.AttitudeQ); i++ {
		m.AttitudeQ[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[52+i*4:]))
	}
	for i := 0; i < len(m.Rates); i++ {
		m.Rates[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[68+i*4:]))
	}
	for i := 0; i < len(m.PositionCov); i++ {
		m.PositionCov[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[80+i*4:]))
	}
	m.EstCapabilities = uint8(payload[92])
	return nil
}

// ControlSystemState struct (generated typeinfo)
// The smoothed, monotonic system state used to feed the control loops of the system.
type ControlSystemState struct {
	TimeUsec    uint64    // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	XAcc        float32   // X acceleration in body frame
	YAcc        float32   // Y acceleration in body frame
	ZAcc        float32   // Z acceleration in body frame
	XVel        float32   // X velocity in body frame
	YVel        float32   // Y velocity in body frame
	ZVel        float32   // Z velocity in body frame
	XPos        float32   // X position in local frame
	YPos        float32   // Y position in local frame
	ZPos        float32   // Z position in local frame
	Airspeed    float32   // Airspeed, set to -1 if unknown
	VelVariance []float32 `len:"3" ` // Variance of body velocity estimate
	PosVariance []float32 `len:"3" ` // Variance in local position
	Q           []float32 `len:"4" ` // The attitude, represented as Quaternion
	RollRate    float32   // Angular rate in roll axis
	PitchRate   float32   // Angular rate in pitch axis
	YawRate     float32   // Angular rate in yaw axis
}

// MsgID (generated function)
func (m *ControlSystemState) MsgID() message.MessageID {
	return MSG_ID_CONTROL_SYSTEM_STATE
}

// String (generated function)
func (m *ControlSystemState) String() string {
	return fmt.Sprintf(
		"&common.ControlSystemState{ TimeUsec: %+v, XAcc: %+v, YAcc: %+v, ZAcc: %+v, XVel: %+v, YVel: %+v, ZVel: %+v, XPos: %+v, YPos: %+v, ZPos: %+v, Airspeed: %+v, VelVariance: %+v, PosVariance: %+v, Q: %+v, RollRate: %+v, PitchRate: %+v, YawRate: %+v }",
		m.TimeUsec,
		m.XAcc,
		m.YAcc,
		m.ZAcc,
		m.XVel,
		m.YVel,
		m.ZVel,
		m.XPos,
		m.YPos,
		m.ZPos,
		m.Airspeed,
		m.VelVariance,
		m.PosVariance,
		m.Q,
		m.RollRate,
		m.PitchRate,
		m.YawRate,
	)
}

// Marshal (generated function)
func (m *ControlSystemState) Marshal() ([]byte, error) {
	payload := make([]byte, 100)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.XAcc))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.YAcc))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.ZAcc))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.XVel))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.YVel))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.ZVel))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.XPos))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.YPos))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.ZPos))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.Airspeed))
	for i, v := range m.VelVariance {
		binary.LittleEndian.PutUint32(payload[48+i*4:], math.Float32bits(v))
	}
	for i, v := range m.PosVariance {
		binary.LittleEndian.PutUint32(payload[60+i*4:], math.Float32bits(v))
	}
	for i, v := range m.Q {
		binary.LittleEndian.PutUint32(payload[72+i*4:], math.Float32bits(v))
	}
	binary.LittleEndian.PutUint32(payload[88:], math.Float32bits(m.RollRate))
	binary.LittleEndian.PutUint32(payload[92:], math.Float32bits(m.PitchRate))
	binary.LittleEndian.PutUint32(payload[96:], math.Float32bits(m.YawRate))
	return payload, nil
}

// Unmarshal (generated function)
func (m *ControlSystemState) Unmarshal(data []byte) error {
	payload := make([]byte, 100)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.XAcc = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.YAcc = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.ZAcc = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.XVel = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.YVel = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.ZVel = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.XPos = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.YPos = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.ZPos = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.Airspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	for i := 0; i < len(m.VelVariance); i++ {
		m.VelVariance[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[48+i*4:]))
	}
	for i := 0; i < len(m.PosVariance); i++ {
		m.PosVariance[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[60+i*4:]))
	}
	for i := 0; i < len(m.Q); i++ {
		m.Q[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[72+i*4:]))
	}
	m.RollRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[88:]))
	m.PitchRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[92:]))
	m.YawRate = math.Float32frombits(binary.LittleEndian.Uint32(payload[96:]))
	return nil
}

// BatteryStatus struct (generated typeinfo)
// Battery information. Updates GCS with flight controller battery status. Smart batteries also use this message, but may additionally send SMART_BATTERY_INFO.
type BatteryStatus struct {
	CurrentConsumed  int32                // Consumed charge, -1: autopilot does not provide consumption estimate
	EnergyConsumed   int32                // Consumed energy, -1: autopilot does not provide energy consumption estimate
	Temperature      int16                // Temperature of the battery. INT16_MAX for unknown temperature.
	Voltages         []uint16             `len:"10" ` // Battery voltage of cells 1 to 10 (see voltages_ext for cells 11-14). Cells in this field above the valid cell count for this battery should have the UINT16_MAX value. If individual cell voltages are unknown or not measured for this battery, then the overall battery voltage should be filled in cell 0, with all others set to UINT16_MAX. If the voltage of the battery is greater than (UINT16_MAX - 1), then cell 0 should be set to (UINT16_MAX - 1), and cell 1 to the remaining voltage. This can be extended to multiple cells if the total voltage is greater than 2 * (UINT16_MAX - 1).
	CurrentBattery   int16                // Battery current, -1: autopilot does not measure the current
	ID               uint8                // Battery ID
	BatteryFunction  MAV_BATTERY_FUNCTION // Function of the battery
	Type             MAV_BATTERY_TYPE     // Type (chemistry) of the battery
	BatteryRemaining int8                 // Remaining battery energy. Values: [0-100], -1: autopilot does not estimate the remaining battery.
}

// MsgID (generated function)
func (m *BatteryStatus) MsgID() message.MessageID {
	return MSG_ID_BATTERY_STATUS
}

// String (generated function)
func (m *BatteryStatus) String() string {
	return fmt.Sprintf(
		"&common.BatteryStatus{ CurrentConsumed: %+v, EnergyConsumed: %+v, Temperature: %+v, Voltages: %+v, CurrentBattery: %+v, ID: %+v, BatteryFunction: %+v, Type: %+v, BatteryRemaining: %+v }",
		m.CurrentConsumed,
		m.EnergyConsumed,
		m.Temperature,
		m.Voltages,
		m.CurrentBattery,
		m.ID,
		m.BatteryFunction,
		m.Type,
		m.BatteryRemaining,
	)
}

// Marshal (generated function)
func (m *BatteryStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 36)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.CurrentConsumed))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.EnergyConsumed))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.Temperature))
	for i, v := range m.Voltages {
		binary.LittleEndian.PutUint16(payload[10+i*2:], uint16(v))
	}
	binary.LittleEndian.PutUint16(payload[30:], uint16(m.CurrentBattery))
	payload[32] = byte(m.ID)
	payload[33] = byte(m.BatteryFunction)
	payload[34] = byte(m.Type)
	payload[35] = byte(m.BatteryRemaining)
	return payload, nil
}

// Unmarshal (generated function)
func (m *BatteryStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 36)
	copy(payload[0:], data)
	m.CurrentConsumed = int32(binary.LittleEndian.Uint32(payload[0:]))
	m.EnergyConsumed = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.Temperature = int16(binary.LittleEndian.Uint16(payload[8:]))
	for i := 0; i < len(m.Voltages); i++ {
		m.Voltages[i] = uint16(binary.LittleEndian.Uint16(payload[10+i*2:]))
	}
	m.CurrentBattery = int16(binary.LittleEndian.Uint16(payload[30:]))
	m.ID = uint8(payload[32])
	m.BatteryFunction = MAV_BATTERY_FUNCTION(payload[33])
	m.Type = MAV_BATTERY_TYPE(payload[34])
	m.BatteryRemaining = int8(payload[35])
	return nil
}

// AutopilotVersion struct (generated typeinfo)
// Version and capability of autopilot software. This should be emitted in response to a request with MAV_CMD_REQUEST_MESSAGE.
type AutopilotVersion struct {
	Capabilities            MAV_PROTOCOL_CAPABILITY // Bitmap of capabilities
	UID                     uint64                  // UID if provided by hardware (see uid2)
	FlightSwVersion         uint32                  // Firmware version number
	MiddlewareSwVersion     uint32                  // Middleware version number
	OsSwVersion             uint32                  // Operating system version number
	BoardVersion            uint32                  // HW / board version (last 8 bytes should be silicon ID, if any)
	VendorID                uint16                  // ID of the board vendor
	ProductID               uint16                  // ID of the product
	FlightCustomVersion     []uint8                 `len:"8" ` // Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
	MiddlewareCustomVersion []uint8                 `len:"8" ` // Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
	OsCustomVersion         []uint8                 `len:"8" ` // Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
}

// MsgID (generated function)
func (m *AutopilotVersion) MsgID() message.MessageID {
	return MSG_ID_AUTOPILOT_VERSION
}

// String (generated function)
func (m *AutopilotVersion) String() string {
	return fmt.Sprintf(
		"&common.AutopilotVersion{ Capabilities: %+v (%064b), UID: %+v, FlightSwVersion: %+v, MiddlewareSwVersion: %+v, OsSwVersion: %+v, BoardVersion: %+v, VendorID: %+v, ProductID: %+v, FlightCustomVersion: %0X (\"%s\"), MiddlewareCustomVersion: %0X (\"%s\"), OsCustomVersion: %0X (\"%s\") }",
		m.Capabilities.Bitmask(), uint64(m.Capabilities),
		m.UID,
		m.FlightSwVersion,
		m.MiddlewareSwVersion,
		m.OsSwVersion,
		m.BoardVersion,
		m.VendorID,
		m.ProductID,
		m.FlightCustomVersion, string(m.FlightCustomVersion[:]),
		m.MiddlewareCustomVersion, string(m.MiddlewareCustomVersion[:]),
		m.OsCustomVersion, string(m.OsCustomVersion[:]),
	)
}

// Marshal (generated function)
func (m *AutopilotVersion) Marshal() ([]byte, error) {
	payload := make([]byte, 60)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Capabilities))
	binary.LittleEndian.PutUint64(payload[8:], uint64(m.UID))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.FlightSwVersion))
	binary.LittleEndian.PutUint32(payload[20:], uint32(m.MiddlewareSwVersion))
	binary.LittleEndian.PutUint32(payload[24:], uint32(m.OsSwVersion))
	binary.LittleEndian.PutUint32(payload[28:], uint32(m.BoardVersion))
	binary.LittleEndian.PutUint16(payload[32:], uint16(m.VendorID))
	binary.LittleEndian.PutUint16(payload[34:], uint16(m.ProductID))
	copy(payload[36:], m.FlightCustomVersion)
	copy(payload[44:], m.MiddlewareCustomVersion)
	copy(payload[52:], m.OsCustomVersion)
	return payload, nil
}

// Unmarshal (generated function)
func (m *AutopilotVersion) Unmarshal(data []byte) error {
	payload := make([]byte, 60)
	copy(payload[0:], data)
	m.Capabilities = MAV_PROTOCOL_CAPABILITY(binary.LittleEndian.Uint64(payload[0:]))
	m.UID = uint64(binary.LittleEndian.Uint64(payload[8:]))
	m.FlightSwVersion = uint32(binary.LittleEndian.Uint32(payload[16:]))
	m.MiddlewareSwVersion = uint32(binary.LittleEndian.Uint32(payload[20:]))
	m.OsSwVersion = uint32(binary.LittleEndian.Uint32(payload[24:]))
	m.BoardVersion = uint32(binary.LittleEndian.Uint32(payload[28:]))
	m.VendorID = uint16(binary.LittleEndian.Uint16(payload[32:]))
	m.ProductID = uint16(binary.LittleEndian.Uint16(payload[34:]))
	copy(m.FlightCustomVersion[:], payload[36:44])
	copy(m.MiddlewareCustomVersion[:], payload[44:52])
	copy(m.OsCustomVersion[:], payload[52:60])
	return nil
}

// LandingTarget struct (generated typeinfo)
// The location of a landing target. See: https://mavlink.io/en/services/landing_target.html
type LandingTarget struct {
	TimeUsec  uint64    // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	AngleX    float32   // X-axis angular offset of the target from the center of the image
	AngleY    float32   // Y-axis angular offset of the target from the center of the image
	Distance  float32   // Distance to the target from the vehicle
	SizeX     float32   // Size of target along x-axis
	SizeY     float32   // Size of target along y-axis
	TargetNum uint8     // The ID of the target if multiple targets are present
	Frame     MAV_FRAME // Coordinate frame used for following fields.
}

// MsgID (generated function)
func (m *LandingTarget) MsgID() message.MessageID {
	return MSG_ID_LANDING_TARGET
}

// String (generated function)
func (m *LandingTarget) String() string {
	return fmt.Sprintf(
		"&common.LandingTarget{ TimeUsec: %+v, AngleX: %+v, AngleY: %+v, Distance: %+v, SizeX: %+v, SizeY: %+v, TargetNum: %+v, Frame: %+v }",
		m.TimeUsec,
		m.AngleX,
		m.AngleY,
		m.Distance,
		m.SizeX,
		m.SizeY,
		m.TargetNum,
		m.Frame,
	)
}

// Marshal (generated function)
func (m *LandingTarget) Marshal() ([]byte, error) {
	payload := make([]byte, 30)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.AngleX))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.AngleY))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Distance))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.SizeX))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.SizeY))
	payload[28] = byte(m.TargetNum)
	payload[29] = byte(m.Frame)
	return payload, nil
}

// Unmarshal (generated function)
func (m *LandingTarget) Unmarshal(data []byte) error {
	payload := make([]byte, 30)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.AngleX = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.AngleY = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Distance = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.SizeX = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.SizeY = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.TargetNum = uint8(payload[28])
	m.Frame = MAV_FRAME(payload[29])
	return nil
}

// FenceStatus struct (generated typeinfo)
// Status of geo-fencing. Sent in extended status stream when fencing enabled.
type FenceStatus struct {
	BreachTime   uint32       // Time (since boot) of last breach.
	BreachCount  uint16       // Number of fence breaches.
	BreachStatus uint8        // Breach status (0 if currently inside fence, 1 if outside).
	BreachType   FENCE_BREACH // Last breach type.
}

// MsgID (generated function)
func (m *FenceStatus) MsgID() message.MessageID {
	return MSG_ID_FENCE_STATUS
}

// String (generated function)
func (m *FenceStatus) String() string {
	return fmt.Sprintf(
		"&common.FenceStatus{ BreachTime: %+v, BreachCount: %+v, BreachStatus: %+v, BreachType: %+v }",
		m.BreachTime,
		m.BreachCount,
		m.BreachStatus,
		m.BreachType,
	)
}

// Marshal (generated function)
func (m *FenceStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 8)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.BreachTime))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.BreachCount))
	payload[6] = byte(m.BreachStatus)
	payload[7] = byte(m.BreachType)
	return payload, nil
}

// Unmarshal (generated function)
func (m *FenceStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 8)
	copy(payload[0:], data)
	m.BreachTime = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.BreachCount = uint16(binary.LittleEndian.Uint16(payload[4:]))
	m.BreachStatus = uint8(payload[6])
	m.BreachType = FENCE_BREACH(payload[7])
	return nil
}

// MagCalReport struct (generated typeinfo)
// Reports results of completed compass calibration. Sent until MAG_CAL_ACK received.
type MagCalReport struct {
	Fitness   float32        // RMS milligauss residuals.
	OfsX      float32        // X offset.
	OfsY      float32        // Y offset.
	OfsZ      float32        // Z offset.
	DiagX     float32        // X diagonal (matrix 11).
	DiagY     float32        // Y diagonal (matrix 22).
	DiagZ     float32        // Z diagonal (matrix 33).
	OffdiagX  float32        // X off-diagonal (matrix 12 and 21).
	OffdiagY  float32        // Y off-diagonal (matrix 13 and 31).
	OffdiagZ  float32        // Z off-diagonal (matrix 32 and 23).
	CompassID uint8          // Compass being calibrated.
	CalMask   uint8          // Bitmask of compasses being calibrated.
	CalStatus MAG_CAL_STATUS // Calibration Status.
	Autosaved uint8          // 0=requires a MAV_CMD_DO_ACCEPT_MAG_CAL, 1=saved to parameters.
}

// MsgID (generated function)
func (m *MagCalReport) MsgID() message.MessageID {
	return MSG_ID_MAG_CAL_REPORT
}

// String (generated function)
func (m *MagCalReport) String() string {
	return fmt.Sprintf(
		"&common.MagCalReport{ Fitness: %+v, OfsX: %+v, OfsY: %+v, OfsZ: %+v, DiagX: %+v, DiagY: %+v, DiagZ: %+v, OffdiagX: %+v, OffdiagY: %+v, OffdiagZ: %+v, CompassID: %+v, CalMask: %+v, CalStatus: %+v, Autosaved: %+v }",
		m.Fitness,
		m.OfsX,
		m.OfsY,
		m.OfsZ,
		m.DiagX,
		m.DiagY,
		m.DiagZ,
		m.OffdiagX,
		m.OffdiagY,
		m.OffdiagZ,
		m.CompassID,
		m.CalMask,
		m.CalStatus,
		m.Autosaved,
	)
}

// Marshal (generated function)
func (m *MagCalReport) Marshal() ([]byte, error) {
	payload := make([]byte, 44)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.Fitness))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.OfsX))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.OfsY))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.OfsZ))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.DiagX))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.DiagY))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.DiagZ))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.OffdiagX))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.OffdiagY))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.OffdiagZ))
	payload[40] = byte(m.CompassID)
	payload[41] = byte(m.CalMask)
	payload[42] = byte(m.CalStatus)
	payload[43] = byte(m.Autosaved)
	return payload, nil
}

// Unmarshal (generated function)
func (m *MagCalReport) Unmarshal(data []byte) error {
	payload := make([]byte, 44)
	copy(payload[0:], data)
	m.Fitness = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.OfsX = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.OfsY = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.OfsZ = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.DiagX = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.DiagY = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.DiagZ = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.OffdiagX = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.OffdiagY = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.OffdiagZ = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.CompassID = uint8(payload[40])
	m.CalMask = uint8(payload[41])
	m.CalStatus = MAG_CAL_STATUS(payload[42])
	m.Autosaved = uint8(payload[43])
	return nil
}

// EfiStatus struct (generated typeinfo)
// EFI status output
type EfiStatus struct {
	EcuIndex                  float32 // ECU index
	Rpm                       float32 // RPM
	FuelConsumed              float32 // Fuel consumed
	FuelFlow                  float32 // Fuel flow rate
	EngineLoad                float32 // Engine load
	ThrottlePosition          float32 // Throttle position
	SparkDwellTime            float32 // Spark dwell time
	BarometricPressure        float32 // Barometric pressure
	IntakeManifoldPressure    float32 // Intake manifold pressure(
	IntakeManifoldTemperature float32 // Intake manifold temperature
	CylinderHeadTemperature   float32 // Cylinder head temperature
	IgnitionTiming            float32 // Ignition timing (Crank angle degrees)
	InjectionTime             float32 // Injection time
	ExhaustGasTemperature     float32 // Exhaust gas temperature
	ThrottleOut               float32 // Output throttle
	PtCompensation            float32 // Pressure/temperature compensation
	Health                    uint8   // EFI health status
}

// MsgID (generated function)
func (m *EfiStatus) MsgID() message.MessageID {
	return MSG_ID_EFI_STATUS
}

// String (generated function)
func (m *EfiStatus) String() string {
	return fmt.Sprintf(
		"&common.EfiStatus{ EcuIndex: %+v, Rpm: %+v, FuelConsumed: %+v, FuelFlow: %+v, EngineLoad: %+v, ThrottlePosition: %+v, SparkDwellTime: %+v, BarometricPressure: %+v, IntakeManifoldPressure: %+v, IntakeManifoldTemperature: %+v, CylinderHeadTemperature: %+v, IgnitionTiming: %+v, InjectionTime: %+v, ExhaustGasTemperature: %+v, ThrottleOut: %+v, PtCompensation: %+v, Health: %+v }",
		m.EcuIndex,
		m.Rpm,
		m.FuelConsumed,
		m.FuelFlow,
		m.EngineLoad,
		m.ThrottlePosition,
		m.SparkDwellTime,
		m.BarometricPressure,
		m.IntakeManifoldPressure,
		m.IntakeManifoldTemperature,
		m.CylinderHeadTemperature,
		m.IgnitionTiming,
		m.InjectionTime,
		m.ExhaustGasTemperature,
		m.ThrottleOut,
		m.PtCompensation,
		m.Health,
	)
}

// Marshal (generated function)
func (m *EfiStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 65)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.EcuIndex))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Rpm))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.FuelConsumed))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.FuelFlow))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.EngineLoad))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.ThrottlePosition))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.SparkDwellTime))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.BarometricPressure))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.IntakeManifoldPressure))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.IntakeManifoldTemperature))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.CylinderHeadTemperature))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.IgnitionTiming))
	binary.LittleEndian.PutUint32(payload[48:], math.Float32bits(m.InjectionTime))
	binary.LittleEndian.PutUint32(payload[52:], math.Float32bits(m.ExhaustGasTemperature))
	binary.LittleEndian.PutUint32(payload[56:], math.Float32bits(m.ThrottleOut))
	binary.LittleEndian.PutUint32(payload[60:], math.Float32bits(m.PtCompensation))
	payload[64] = byte(m.Health)
	return payload, nil
}

// Unmarshal (generated function)
func (m *EfiStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 65)
	copy(payload[0:], data)
	m.EcuIndex = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.Rpm = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.FuelConsumed = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.FuelFlow = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.EngineLoad = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.ThrottlePosition = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.SparkDwellTime = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.BarometricPressure = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.IntakeManifoldPressure = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.IntakeManifoldTemperature = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.CylinderHeadTemperature = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.IgnitionTiming = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.InjectionTime = math.Float32frombits(binary.LittleEndian.Uint32(payload[48:]))
	m.ExhaustGasTemperature = math.Float32frombits(binary.LittleEndian.Uint32(payload[52:]))
	m.ThrottleOut = math.Float32frombits(binary.LittleEndian.Uint32(payload[56:]))
	m.PtCompensation = math.Float32frombits(binary.LittleEndian.Uint32(payload[60:]))
	m.Health = uint8(payload[64])
	return nil
}

// EstimatorStatus struct (generated typeinfo)
// Estimator status message including flags, innovation test ratios and estimated accuracies. The flags message is an integer bitmask containing information on which EKF outputs are valid. See the ESTIMATOR_STATUS_FLAGS enum definition for further information. The innovation test ratios show the magnitude of the sensor innovation divided by the innovation check threshold. Under normal operation the innovation test ratios should be below 0.5 with occasional values up to 1.0. Values greater than 1.0 should be rare under normal operation and indicate that a measurement has been rejected by the filter. The user should be notified if an innovation test ratio greater than 1.0 is recorded. Notifications for values in the range between 0.5 and 1.0 should be optional and controllable by the user.
type EstimatorStatus struct {
	TimeUsec         uint64                 // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	VelRatio         float32                // Velocity innovation test ratio
	PosHorizRatio    float32                // Horizontal position innovation test ratio
	PosVertRatio     float32                // Vertical position innovation test ratio
	MagRatio         float32                // Magnetometer innovation test ratio
	HaglRatio        float32                // Height above terrain innovation test ratio
	TasRatio         float32                // True airspeed innovation test ratio
	PosHorizAccuracy float32                // Horizontal position 1-STD accuracy relative to the EKF local origin
	PosVertAccuracy  float32                // Vertical position 1-STD accuracy relative to the EKF local origin
	Flags            ESTIMATOR_STATUS_FLAGS // Bitmap indicating which EKF outputs are valid.
}

// MsgID (generated function)
func (m *EstimatorStatus) MsgID() message.MessageID {
	return MSG_ID_ESTIMATOR_STATUS
}

// String (generated function)
func (m *EstimatorStatus) String() string {
	return fmt.Sprintf(
		"&common.EstimatorStatus{ TimeUsec: %+v, VelRatio: %+v, PosHorizRatio: %+v, PosVertRatio: %+v, MagRatio: %+v, HaglRatio: %+v, TasRatio: %+v, PosHorizAccuracy: %+v, PosVertAccuracy: %+v, Flags: %+v (%016b) }",
		m.TimeUsec,
		m.VelRatio,
		m.PosHorizRatio,
		m.PosVertRatio,
		m.MagRatio,
		m.HaglRatio,
		m.TasRatio,
		m.PosHorizAccuracy,
		m.PosVertAccuracy,
		m.Flags.Bitmask(), uint64(m.Flags),
	)
}

// Marshal (generated function)
func (m *EstimatorStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 42)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.VelRatio))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.PosHorizRatio))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.PosVertRatio))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.MagRatio))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.HaglRatio))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.TasRatio))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.PosHorizAccuracy))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.PosVertAccuracy))
	binary.LittleEndian.PutUint16(payload[40:], uint16(m.Flags))
	return payload, nil
}

// Unmarshal (generated function)
func (m *EstimatorStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 42)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.VelRatio = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.PosHorizRatio = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.PosVertRatio = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.MagRatio = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.HaglRatio = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.TasRatio = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.PosHorizAccuracy = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.PosVertAccuracy = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Flags = ESTIMATOR_STATUS_FLAGS(binary.LittleEndian.Uint16(payload[40:]))
	return nil
}

// WindCov struct (generated typeinfo)
// Wind covariance estimate from vehicle.
type WindCov struct {
	TimeUsec      uint64  // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	WindX         float32 // Wind in X (NED) direction
	WindY         float32 // Wind in Y (NED) direction
	WindZ         float32 // Wind in Z (NED) direction
	VarHoriz      float32 // Variability of the wind in XY. RMS of a 1 Hz lowpassed wind estimate.
	VarVert       float32 // Variability of the wind in Z. RMS of a 1 Hz lowpassed wind estimate.
	WindAlt       float32 // Altitude (MSL) that this measurement was taken at
	HorizAccuracy float32 // Horizontal speed 1-STD accuracy
	VertAccuracy  float32 // Vertical speed 1-STD accuracy
}

// MsgID (generated function)
func (m *WindCov) MsgID() message.MessageID {
	return MSG_ID_WIND_COV
}

// String (generated function)
func (m *WindCov) String() string {
	return fmt.Sprintf(
		"&common.WindCov{ TimeUsec: %+v, WindX: %+v, WindY: %+v, WindZ: %+v, VarHoriz: %+v, VarVert: %+v, WindAlt: %+v, HorizAccuracy: %+v, VertAccuracy: %+v }",
		m.TimeUsec,
		m.WindX,
		m.WindY,
		m.WindZ,
		m.VarHoriz,
		m.VarVert,
		m.WindAlt,
		m.HorizAccuracy,
		m.VertAccuracy,
	)
}

// Marshal (generated function)
func (m *WindCov) Marshal() ([]byte, error) {
	payload := make([]byte, 40)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.WindX))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.WindY))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.WindZ))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.VarHoriz))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.VarVert))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.WindAlt))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.HorizAccuracy))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.VertAccuracy))
	return payload, nil
}

// Unmarshal (generated function)
func (m *WindCov) Unmarshal(data []byte) error {
	payload := make([]byte, 40)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.WindX = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.WindY = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.WindZ = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.VarHoriz = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.VarVert = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.WindAlt = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.HorizAccuracy = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.VertAccuracy = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	return nil
}

// GpsInput struct (generated typeinfo)
// GPS sensor input message.  This is a raw sensor value sent by the GPS. This is NOT the global position estimate of the system.
type GpsInput struct {
	TimeUsec          uint64                 // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeWeekMs        uint32                 // GPS time (from start of GPS week)
	Lat               int32                  // Latitude (WGS84)
	Lon               int32                  // Longitude (WGS84)
	Alt               float32                // Altitude (MSL). Positive for up.
	Hdop              float32                // GPS HDOP horizontal dilution of position
	Vdop              float32                // GPS VDOP vertical dilution of position
	Vn                float32                // GPS velocity in north direction in earth-fixed NED frame
	Ve                float32                // GPS velocity in east direction in earth-fixed NED frame
	Vd                float32                // GPS velocity in down direction in earth-fixed NED frame
	SpeedAccuracy     float32                // GPS speed accuracy
	HorizAccuracy     float32                // GPS horizontal accuracy
	VertAccuracy      float32                // GPS vertical accuracy
	IgnoreFlags       GPS_INPUT_IGNORE_FLAGS // Bitmap indicating which GPS input flags fields to ignore.  All other fields must be provided.
	TimeWeek          uint16                 // GPS week number
	GpsID             uint8                  // ID of the GPS for multiple GPS inputs
	FixType           uint8                  // 0-1: no fix, 2: 2D fix, 3: 3D fix. 4: 3D with DGPS. 5: 3D with RTK
	SatellitesVisible uint8                  // Number of satellites visible.
}

// MsgID (generated function)
func (m *GpsInput) MsgID() message.MessageID {
	return MSG_ID_GPS_INPUT
}

// String (generated function)
func (m *GpsInput) String() string {
	return fmt.Sprintf(
		"&common.GpsInput{ TimeUsec: %+v, TimeWeekMs: %+v, Lat: %+v, Lon: %+v, Alt: %+v, Hdop: %+v, Vdop: %+v, Vn: %+v, Ve: %+v, Vd: %+v, SpeedAccuracy: %+v, HorizAccuracy: %+v, VertAccuracy: %+v, IgnoreFlags: %+v (%016b), TimeWeek: %+v, GpsID: %+v, FixType: %+v, SatellitesVisible: %+v }",
		m.TimeUsec,
		m.TimeWeekMs,
		m.Lat,
		m.Lon,
		m.Alt,
		m.Hdop,
		m.Vdop,
		m.Vn,
		m.Ve,
		m.Vd,
		m.SpeedAccuracy,
		m.HorizAccuracy,
		m.VertAccuracy,
		m.IgnoreFlags.Bitmask(), uint64(m.IgnoreFlags),
		m.TimeWeek,
		m.GpsID,
		m.FixType,
		m.SatellitesVisible,
	)
}

// Marshal (generated function)
func (m *GpsInput) Marshal() ([]byte, error) {
	payload := make([]byte, 63)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.TimeWeekMs))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Alt))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Hdop))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Vdop))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Vn))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Ve))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Vd))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.SpeedAccuracy))
	binary.LittleEndian.PutUint32(payload[48:], math.Float32bits(m.HorizAccuracy))
	binary.LittleEndian.PutUint32(payload[52:], math.Float32bits(m.VertAccuracy))
	binary.LittleEndian.PutUint16(payload[56:], uint16(m.IgnoreFlags))
	binary.LittleEndian.PutUint16(payload[58:], uint16(m.TimeWeek))
	payload[60] = byte(m.GpsID)
	payload[61] = byte(m.FixType)
	payload[62] = byte(m.SatellitesVisible)
	return payload, nil
}

// Unmarshal (generated function)
func (m *GpsInput) Unmarshal(data []byte) error {
	payload := make([]byte, 63)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.TimeWeekMs = uint32(binary.LittleEndian.Uint32(payload[8:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[12:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[16:]))
	m.Alt = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Hdop = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Vdop = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Vn = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Ve = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Vd = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.SpeedAccuracy = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.HorizAccuracy = math.Float32frombits(binary.LittleEndian.Uint32(payload[48:]))
	m.VertAccuracy = math.Float32frombits(binary.LittleEndian.Uint32(payload[52:]))
	m.IgnoreFlags = GPS_INPUT_IGNORE_FLAGS(binary.LittleEndian.Uint16(payload[56:]))
	m.TimeWeek = uint16(binary.LittleEndian.Uint16(payload[58:]))
	m.GpsID = uint8(payload[60])
	m.FixType = uint8(payload[61])
	m.SatellitesVisible = uint8(payload[62])
	return nil
}

// GpsRtcmData struct (generated typeinfo)
// RTCM message for injecting into the onboard GPS (used for DGPS)
type GpsRtcmData struct {
	Flags uint8   // LSB: 1 means message is fragmented, next 2 bits are the fragment ID, the remaining 5 bits are used for the sequence ID. Messages are only to be flushed to the GPS when the entire message has been reconstructed on the autopilot. The fragment ID specifies which order the fragments should be assembled into a buffer, while the sequence ID is used to detect a mismatch between different buffers. The buffer is considered fully reconstructed when either all 4 fragments are present, or all the fragments before the first fragment with a non full payload is received. This management is used to ensure that normal GPS operation doesn't corrupt RTCM data, and to recover from a unreliable transport delivery order.
	Len   uint8   // data length
	Data  []uint8 `len:"180" ` // RTCM message (may be fragmented)
}

// MsgID (generated function)
func (m *GpsRtcmData) MsgID() message.MessageID {
	return MSG_ID_GPS_RTCM_DATA
}

// String (generated function)
func (m *GpsRtcmData) String() string {
	return fmt.Sprintf(
		"&common.GpsRtcmData{ Flags: %+v, Len: %+v, Data: %0X (\"%s\") }",
		m.Flags,
		m.Len,
		m.Data, string(m.Data[:]),
	)
}

// Marshal (generated function)
func (m *GpsRtcmData) Marshal() ([]byte, error) {
	payload := make([]byte, 182)
	payload[0] = byte(m.Flags)
	payload[1] = byte(m.Len)
	copy(payload[2:], m.Data)
	return payload, nil
}

// Unmarshal (generated function)
func (m *GpsRtcmData) Unmarshal(data []byte) error {
	payload := make([]byte, 182)
	copy(payload[0:], data)
	m.Flags = uint8(payload[0])
	m.Len = uint8(payload[1])
	copy(m.Data[:], payload[2:182])
	return nil
}

// HighLatency struct (generated typeinfo)
// Message appropriate for high latency connections like Iridium
type HighLatency struct {
	CustomMode       uint32           // A bitfield for use for autopilot-specific flags.
	Latitude         int32            // Latitude
	Longitude        int32            // Longitude
	Roll             int16            // roll
	Pitch            int16            // pitch
	Heading          uint16           // heading
	HeadingSp        int16            // heading setpoint
	AltitudeAmsl     int16            // Altitude above mean sea level
	AltitudeSp       int16            // Altitude setpoint relative to the home position
	WpDistance       uint16           // distance to target
	BaseMode         MAV_MODE_FLAG    // Bitmap of enabled system modes.
	LandedState      MAV_LANDED_STATE // The landed state. Is set to MAV_LANDED_STATE_UNDEFINED if landed state is unknown.
	Throttle         int8             // throttle (percentage)
	Airspeed         uint8            // airspeed
	AirspeedSp       uint8            // airspeed setpoint
	Groundspeed      uint8            // groundspeed
	ClimbRate        int8             // climb rate
	GpsNsat          uint8            // Number of satellites visible. If unknown, set to 255
	GpsFixType       GPS_FIX_TYPE     // GPS Fix type.
	BatteryRemaining uint8            // Remaining battery (percentage)
	Temperature      int8             // Autopilot temperature (degrees C)
	TemperatureAir   int8             // Air temperature (degrees C) from airspeed sensor
	Failsafe         uint8            // failsafe (each bit represents a failsafe where 0=ok, 1=failsafe active (bit0:RC, bit1:batt, bit2:GPS, bit3:GCS, bit4:fence)
	WpNum            uint8            // current waypoint number
}

// MsgID (generated function)
func (m *HighLatency) MsgID() message.MessageID {
	return MSG_ID_HIGH_LATENCY
}

// String (generated function)
func (m *HighLatency) String() string {
	return fmt.Sprintf(
		"&common.HighLatency{ CustomMode: %+v, Latitude: %+v, Longitude: %+v, Roll: %+v, Pitch: %+v, Heading: %+v, HeadingSp: %+v, AltitudeAmsl: %+v, AltitudeSp: %+v, WpDistance: %+v, BaseMode: %+v (%08b), LandedState: %+v, Throttle: %+v, Airspeed: %+v, AirspeedSp: %+v, Groundspeed: %+v, ClimbRate: %+v, GpsNsat: %+v, GpsFixType: %+v, BatteryRemaining: %+v, Temperature: %+v, TemperatureAir: %+v, Failsafe: %+v, WpNum: %+v }",
		m.CustomMode,
		m.Latitude,
		m.Longitude,
		m.Roll,
		m.Pitch,
		m.Heading,
		m.HeadingSp,
		m.AltitudeAmsl,
		m.AltitudeSp,
		m.WpDistance,
		m.BaseMode.Bitmask(), uint64(m.BaseMode),
		m.LandedState,
		m.Throttle,
		m.Airspeed,
		m.AirspeedSp,
		m.Groundspeed,
		m.ClimbRate,
		m.GpsNsat,
		m.GpsFixType,
		m.BatteryRemaining,
		m.Temperature,
		m.TemperatureAir,
		m.Failsafe,
		m.WpNum,
	)
}

// Marshal (generated function)
func (m *HighLatency) Marshal() ([]byte, error) {
	payload := make([]byte, 40)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.CustomMode))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Latitude))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Longitude))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.Roll))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Pitch))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Heading))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.HeadingSp))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.AltitudeAmsl))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.AltitudeSp))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.WpDistance))
	payload[26] = byte(m.BaseMode)
	payload[27] = byte(m.LandedState)
	payload[28] = byte(m.Throttle)
	payload[29] = byte(m.Airspeed)
	payload[30] = byte(m.AirspeedSp)
	payload[31] = byte(m.Groundspeed)
	payload[32] = byte(m.ClimbRate)
	payload[33] = byte(m.GpsNsat)
	payload[34] = byte(m.GpsFixType)
	payload[35] = byte(m.BatteryRemaining)
	payload[36] = byte(m.Temperature)
	payload[37] = byte(m.TemperatureAir)
	payload[38] = byte(m.Failsafe)
	payload[39] = byte(m.WpNum)
	return payload, nil
}

// Unmarshal (generated function)
func (m *HighLatency) Unmarshal(data []byte) error {
	payload := make([]byte, 40)
	copy(payload[0:], data)
	m.CustomMode = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Latitude = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.Longitude = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.Roll = int16(binary.LittleEndian.Uint16(payload[12:]))
	m.Pitch = int16(binary.LittleEndian.Uint16(payload[14:]))
	m.Heading = uint16(binary.LittleEndian.Uint16(payload[16:]))
	m.HeadingSp = int16(binary.LittleEndian.Uint16(payload[18:]))
	m.AltitudeAmsl = int16(binary.LittleEndian.Uint16(payload[20:]))
	m.AltitudeSp = int16(binary.LittleEndian.Uint16(payload[22:]))
	m.WpDistance = uint16(binary.LittleEndian.Uint16(payload[24:]))
	m.BaseMode = MAV_MODE_FLAG(payload[26])
	m.LandedState = MAV_LANDED_STATE(payload[27])
	m.Throttle = int8(payload[28])
	m.Airspeed = uint8(payload[29])
	m.AirspeedSp = uint8(payload[30])
	m.Groundspeed = uint8(payload[31])
	m.ClimbRate = int8(payload[32])
	m.GpsNsat = uint8(payload[33])
	m.GpsFixType = GPS_FIX_TYPE(payload[34])
	m.BatteryRemaining = uint8(payload[35])
	m.Temperature = int8(payload[36])
	m.TemperatureAir = int8(payload[37])
	m.Failsafe = uint8(payload[38])
	m.WpNum = uint8(payload[39])
	return nil
}

// HighLatency2 struct (generated typeinfo)
// Message appropriate for high latency connections like Iridium (version 2)
type HighLatency2 struct {
	Timestamp      uint32          // Timestamp (milliseconds since boot or Unix epoch)
	Latitude       int32           // Latitude
	Longitude      int32           // Longitude
	CustomMode     uint16          // A bitfield for use for autopilot-specific flags (2 byte version).
	Altitude       int16           // Altitude above mean sea level
	TargetAltitude int16           // Altitude setpoint
	TargetDistance uint16          // Distance to target waypoint or position
	WpNum          uint16          // Current waypoint number
	FailureFlags   HL_FAILURE_FLAG // Bitmap of failure flags.
	Type           MAV_TYPE        // Type of the MAV (quadrotor, helicopter, etc.)
	Autopilot      MAV_AUTOPILOT   // Autopilot type / class. Use MAV_AUTOPILOT_INVALID for components that are not flight controllers.
	Heading        uint8           // Heading
	TargetHeading  uint8           // Heading setpoint
	Throttle       uint8           // Throttle
	Airspeed       uint8           // Airspeed
	AirspeedSp     uint8           // Airspeed setpoint
	Groundspeed    uint8           // Groundspeed
	Windspeed      uint8           // Windspeed
	WindHeading    uint8           // Wind heading
	Eph            uint8           // Maximum error horizontal position since last message
	Epv            uint8           // Maximum error vertical position since last message
	TemperatureAir int8            // Air temperature from airspeed sensor
	ClimbRate      int8            // Maximum climb rate magnitude since last message
	Battery        int8            // Battery level (-1 if field not provided).
	Custom0        int8            // Field for custom payload.
	Custom1        int8            // Field for custom payload.
	Custom2        int8            // Field for custom payload.
}

// MsgID (generated function)
func (m *HighLatency2) MsgID() message.MessageID {
	return MSG_ID_HIGH_LATENCY2
}

// String (generated function)
func (m *HighLatency2) String() string {
	return fmt.Sprintf(
		"&common.HighLatency2{ Timestamp: %+v, Latitude: %+v, Longitude: %+v, CustomMode: %+v, Altitude: %+v, TargetAltitude: %+v, TargetDistance: %+v, WpNum: %+v, FailureFlags: %+v (%016b), Type: %+v, Autopilot: %+v, Heading: %+v, TargetHeading: %+v, Throttle: %+v, Airspeed: %+v, AirspeedSp: %+v, Groundspeed: %+v, Windspeed: %+v, WindHeading: %+v, Eph: %+v, Epv: %+v, TemperatureAir: %+v, ClimbRate: %+v, Battery: %+v, Custom0: %+v, Custom1: %+v, Custom2: %+v }",
		m.Timestamp,
		m.Latitude,
		m.Longitude,
		m.CustomMode,
		m.Altitude,
		m.TargetAltitude,
		m.TargetDistance,
		m.WpNum,
		m.FailureFlags.Bitmask(), uint64(m.FailureFlags),
		m.Type,
		m.Autopilot,
		m.Heading,
		m.TargetHeading,
		m.Throttle,
		m.Airspeed,
		m.AirspeedSp,
		m.Groundspeed,
		m.Windspeed,
		m.WindHeading,
		m.Eph,
		m.Epv,
		m.TemperatureAir,
		m.ClimbRate,
		m.Battery,
		m.Custom0,
		m.Custom1,
		m.Custom2,
	)
}

// Marshal (generated function)
func (m *HighLatency2) Marshal() ([]byte, error) {
	payload := make([]byte, 42)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Timestamp))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Latitude))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Longitude))
	binary.LittleEndian.PutUint16(payload[12:], uint16(m.CustomMode))
	binary.LittleEndian.PutUint16(payload[14:], uint16(m.Altitude))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.TargetAltitude))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.TargetDistance))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.WpNum))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.FailureFlags))
	payload[24] = byte(m.Type)
	payload[25] = byte(m.Autopilot)
	payload[26] = byte(m.Heading)
	payload[27] = byte(m.TargetHeading)
	payload[28] = byte(m.Throttle)
	payload[29] = byte(m.Airspeed)
	payload[30] = byte(m.AirspeedSp)
	payload[31] = byte(m.Groundspeed)
	payload[32] = byte(m.Windspeed)
	payload[33] = byte(m.WindHeading)
	payload[34] = byte(m.Eph)
	payload[35] = byte(m.Epv)
	payload[36] = byte(m.TemperatureAir)
	payload[37] = byte(m.ClimbRate)
	payload[38] = byte(m.Battery)
	payload[39] = byte(m.Custom0)
	payload[40] = byte(m.Custom1)
	payload[41] = byte(m.Custom2)
	return payload, nil
}

// Unmarshal (generated function)
func (m *HighLatency2) Unmarshal(data []byte) error {
	payload := make([]byte, 42)
	copy(payload[0:], data)
	m.Timestamp = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Latitude = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.Longitude = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.CustomMode = uint16(binary.LittleEndian.Uint16(payload[12:]))
	m.Altitude = int16(binary.LittleEndian.Uint16(payload[14:]))
	m.TargetAltitude = int16(binary.LittleEndian.Uint16(payload[16:]))
	m.TargetDistance = uint16(binary.LittleEndian.Uint16(payload[18:]))
	m.WpNum = uint16(binary.LittleEndian.Uint16(payload[20:]))
	m.FailureFlags = HL_FAILURE_FLAG(binary.LittleEndian.Uint16(payload[22:]))
	m.Type = MAV_TYPE(payload[24])
	m.Autopilot = MAV_AUTOPILOT(payload[25])
	m.Heading = uint8(payload[26])
	m.TargetHeading = uint8(payload[27])
	m.Throttle = uint8(payload[28])
	m.Airspeed = uint8(payload[29])
	m.AirspeedSp = uint8(payload[30])
	m.Groundspeed = uint8(payload[31])
	m.Windspeed = uint8(payload[32])
	m.WindHeading = uint8(payload[33])
	m.Eph = uint8(payload[34])
	m.Epv = uint8(payload[35])
	m.TemperatureAir = int8(payload[36])
	m.ClimbRate = int8(payload[37])
	m.Battery = int8(payload[38])
	m.Custom0 = int8(payload[39])
	m.Custom1 = int8(payload[40])
	m.Custom2 = int8(payload[41])
	return nil
}

// Vibration struct (generated typeinfo)
// Vibration levels and accelerometer clipping
type Vibration struct {
	TimeUsec   uint64  // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	VibrationX float32 // Vibration levels on X-axis
	VibrationY float32 // Vibration levels on Y-axis
	VibrationZ float32 // Vibration levels on Z-axis
	Clipping0  uint32  // first accelerometer clipping count
	Clipping1  uint32  // second accelerometer clipping count
	Clipping2  uint32  // third accelerometer clipping count
}

// MsgID (generated function)
func (m *Vibration) MsgID() message.MessageID {
	return MSG_ID_VIBRATION
}

// String (generated function)
func (m *Vibration) String() string {
	return fmt.Sprintf(
		"&common.Vibration{ TimeUsec: %+v, VibrationX: %+v, VibrationY: %+v, VibrationZ: %+v, Clipping0: %+v, Clipping1: %+v, Clipping2: %+v }",
		m.TimeUsec,
		m.VibrationX,
		m.VibrationY,
		m.VibrationZ,
		m.Clipping0,
		m.Clipping1,
		m.Clipping2,
	)
}

// Marshal (generated function)
func (m *Vibration) Marshal() ([]byte, error) {
	payload := make([]byte, 32)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.VibrationX))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.VibrationY))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.VibrationZ))
	binary.LittleEndian.PutUint32(payload[20:], uint32(m.Clipping0))
	binary.LittleEndian.PutUint32(payload[24:], uint32(m.Clipping1))
	binary.LittleEndian.PutUint32(payload[28:], uint32(m.Clipping2))
	return payload, nil
}

// Unmarshal (generated function)
func (m *Vibration) Unmarshal(data []byte) error {
	payload := make([]byte, 32)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.VibrationX = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.VibrationY = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.VibrationZ = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Clipping0 = uint32(binary.LittleEndian.Uint32(payload[20:]))
	m.Clipping1 = uint32(binary.LittleEndian.Uint32(payload[24:]))
	m.Clipping2 = uint32(binary.LittleEndian.Uint32(payload[28:]))
	return nil
}

// HomePosition struct (generated typeinfo)
// This message can be requested by sending the MAV_CMD_GET_HOME_POSITION command. The position the system will return to and land on. The position is set automatically by the system during the takeoff in case it was not explicitly set by the operator before or after. The global and local positions encode the position in the respective coordinate frames, while the q parameter encodes the orientation of the surface. Under normal conditions it describes the heading and terrain slope, which can be used by the aircraft to adjust the approach. The approach 3D vector describes the point to which the system should fly in normal flight mode and then perform a landing sequence along the vector.
type HomePosition struct {
	Latitude  int32     // Latitude (WGS84)
	Longitude int32     // Longitude (WGS84)
	Altitude  int32     // Altitude (MSL). Positive for up.
	X         float32   // Local X position of this position in the local coordinate frame
	Y         float32   // Local Y position of this position in the local coordinate frame
	Z         float32   // Local Z position of this position in the local coordinate frame
	Q         []float32 `len:"4" ` // World to surface normal and heading transformation of the takeoff position. Used to indicate the heading and slope of the ground
	ApproachX float32   // Local X position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachY float32   // Local Y position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachZ float32   // Local Z position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
}

// MsgID (generated function)
func (m *HomePosition) MsgID() message.MessageID {
	return MSG_ID_HOME_POSITION
}

// String (generated function)
func (m *HomePosition) String() string {
	return fmt.Sprintf(
		"&common.HomePosition{ Latitude: %+v, Longitude: %+v, Altitude: %+v, X: %+v, Y: %+v, Z: %+v, Q: %+v, ApproachX: %+v, ApproachY: %+v, ApproachZ: %+v }",
		m.Latitude,
		m.Longitude,
		m.Altitude,
		m.X,
		m.Y,
		m.Z,
		m.Q,
		m.ApproachX,
		m.ApproachY,
		m.ApproachZ,
	)
}

// Marshal (generated function)
func (m *HomePosition) Marshal() ([]byte, error) {
	payload := make([]byte, 52)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Latitude))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Longitude))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Altitude))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Z))
	for i, v := range m.Q {
		binary.LittleEndian.PutUint32(payload[24+i*4:], math.Float32bits(v))
	}
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.ApproachX))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.ApproachY))
	binary.LittleEndian.PutUint32(payload[48:], math.Float32bits(m.ApproachZ))
	return payload, nil
}

// Unmarshal (generated function)
func (m *HomePosition) Unmarshal(data []byte) error {
	payload := make([]byte, 52)
	copy(payload[0:], data)
	m.Latitude = int32(binary.LittleEndian.Uint32(payload[0:]))
	m.Longitude = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.Altitude = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	for i := 0; i < len(m.Q); i++ {
		m.Q[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[24+i*4:]))
	}
	m.ApproachX = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.ApproachY = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.ApproachZ = math.Float32frombits(binary.LittleEndian.Uint32(payload[48:]))
	return nil
}

// SetHomePosition struct (generated typeinfo)
// The position the system will return to and land on. The position is set automatically by the system during the takeoff in case it was not explicitly set by the operator before or after. The global and local positions encode the position in the respective coordinate frames, while the q parameter encodes the orientation of the surface. Under normal conditions it describes the heading and terrain slope, which can be used by the aircraft to adjust the approach. The approach 3D vector describes the point to which the system should fly in normal flight mode and then perform a landing sequence along the vector.
type SetHomePosition struct {
	Latitude     int32     // Latitude (WGS84)
	Longitude    int32     // Longitude (WGS84)
	Altitude     int32     // Altitude (MSL). Positive for up.
	X            float32   // Local X position of this position in the local coordinate frame
	Y            float32   // Local Y position of this position in the local coordinate frame
	Z            float32   // Local Z position of this position in the local coordinate frame
	Q            []float32 `len:"4" ` // World to surface normal and heading transformation of the takeoff position. Used to indicate the heading and slope of the ground
	ApproachX    float32   // Local X position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachY    float32   // Local Y position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachZ    float32   // Local Z position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	TargetSystem uint8     // System ID.
}

// MsgID (generated function)
func (m *SetHomePosition) MsgID() message.MessageID {
	return MSG_ID_SET_HOME_POSITION
}

// String (generated function)
func (m *SetHomePosition) String() string {
	return fmt.Sprintf(
		"&common.SetHomePosition{ Latitude: %+v, Longitude: %+v, Altitude: %+v, X: %+v, Y: %+v, Z: %+v, Q: %+v, ApproachX: %+v, ApproachY: %+v, ApproachZ: %+v, TargetSystem: %+v }",
		m.Latitude,
		m.Longitude,
		m.Altitude,
		m.X,
		m.Y,
		m.Z,
		m.Q,
		m.ApproachX,
		m.ApproachY,
		m.ApproachZ,
		m.TargetSystem,
	)
}

// Marshal (generated function)
func (m *SetHomePosition) Marshal() ([]byte, error) {
	payload := make([]byte, 53)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Latitude))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Longitude))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Altitude))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Z))
	for i, v := range m.Q {
		binary.LittleEndian.PutUint32(payload[24+i*4:], math.Float32bits(v))
	}
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.ApproachX))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.ApproachY))
	binary.LittleEndian.PutUint32(payload[48:], math.Float32bits(m.ApproachZ))
	payload[52] = byte(m.TargetSystem)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SetHomePosition) Unmarshal(data []byte) error {
	payload := make([]byte, 53)
	copy(payload[0:], data)
	m.Latitude = int32(binary.LittleEndian.Uint32(payload[0:]))
	m.Longitude = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.Altitude = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	for i := 0; i < len(m.Q); i++ {
		m.Q[i] = math.Float32frombits(binary.LittleEndian.Uint32(payload[24+i*4:]))
	}
	m.ApproachX = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.ApproachY = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.ApproachZ = math.Float32frombits(binary.LittleEndian.Uint32(payload[48:]))
	m.TargetSystem = uint8(payload[52])
	return nil
}

// MessageInterval struct (generated typeinfo)
// The interval between messages for a particular MAVLink message ID. This message is the response to the MAV_CMD_GET_MESSAGE_INTERVAL command. This interface replaces DATA_STREAM.
type MessageInterval struct {
	IntervalUs int32  // The interval between two messages. A value of -1 indicates this stream is disabled, 0 indicates it is not available, &gt; 0 indicates the interval at which it is sent.
	MessageID  uint16 // The ID of the requested MAVLink message. v1.0 is limited to 254 messages.
}

// MsgID (generated function)
func (m *MessageInterval) MsgID() message.MessageID {
	return MSG_ID_MESSAGE_INTERVAL
}

// String (generated function)
func (m *MessageInterval) String() string {
	return fmt.Sprintf(
		"&common.MessageInterval{ IntervalUs: %+v, MessageID: %+v }",
		m.IntervalUs,
		m.MessageID,
	)
}

// Marshal (generated function)
func (m *MessageInterval) Marshal() ([]byte, error) {
	payload := make([]byte, 6)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.IntervalUs))
	binary.LittleEndian.PutUint16(payload[4:], uint16(m.MessageID))
	return payload, nil
}

// Unmarshal (generated function)
func (m *MessageInterval) Unmarshal(data []byte) error {
	payload := make([]byte, 6)
	copy(payload[0:], data)
	m.IntervalUs = int32(binary.LittleEndian.Uint32(payload[0:]))
	m.MessageID = uint16(binary.LittleEndian.Uint16(payload[4:]))
	return nil
}

// ExtendedSysState struct (generated typeinfo)
// Provides state for additional features
type ExtendedSysState struct {
	VtolState   MAV_VTOL_STATE   // The VTOL state if applicable. Is set to MAV_VTOL_STATE_UNDEFINED if UAV is not in VTOL configuration.
	LandedState MAV_LANDED_STATE // The landed state. Is set to MAV_LANDED_STATE_UNDEFINED if landed state is unknown.
}

// MsgID (generated function)
func (m *ExtendedSysState) MsgID() message.MessageID {
	return MSG_ID_EXTENDED_SYS_STATE
}

// String (generated function)
func (m *ExtendedSysState) String() string {
	return fmt.Sprintf(
		"&common.ExtendedSysState{ VtolState: %+v, LandedState: %+v }",
		m.VtolState,
		m.LandedState,
	)
}

// Marshal (generated function)
func (m *ExtendedSysState) Marshal() ([]byte, error) {
	payload := make([]byte, 2)
	payload[0] = byte(m.VtolState)
	payload[1] = byte(m.LandedState)
	return payload, nil
}

// Unmarshal (generated function)
func (m *ExtendedSysState) Unmarshal(data []byte) error {
	payload := make([]byte, 2)
	copy(payload[0:], data)
	m.VtolState = MAV_VTOL_STATE(payload[0])
	m.LandedState = MAV_LANDED_STATE(payload[1])
	return nil
}

// AdsbVehicle struct (generated typeinfo)
// The location and information of an ADSB vehicle
type AdsbVehicle struct {
	IcaoAddress  uint32             // ICAO address
	Lat          int32              // Latitude
	Lon          int32              // Longitude
	Altitude     int32              // Altitude(ASL)
	Heading      uint16             // Course over ground
	HorVelocity  uint16             // The horizontal velocity
	VerVelocity  int16              // The vertical velocity. Positive is up
	Flags        ADSB_FLAGS         // Bitmap to indicate various statuses including valid data fields
	Squawk       uint16             // Squawk code
	AltitudeType ADSB_ALTITUDE_TYPE // ADSB altitude type.
	Callsign     string             `len:"9" ` // The callsign, 8+null
	EmitterType  ADSB_EMITTER_TYPE  // ADSB emitter type.
	Tslc         uint8              // Time since last communication in seconds
}

// MsgID (generated function)
func (m *AdsbVehicle) MsgID() message.MessageID {
	return MSG_ID_ADSB_VEHICLE
}

// String (generated function)
func (m *AdsbVehicle) String() string {
	return fmt.Sprintf(
		"&common.AdsbVehicle{ IcaoAddress: %+v, Lat: %+v, Lon: %+v, Altitude: %+v, Heading: %+v, HorVelocity: %+v, VerVelocity: %+v, Flags: %+v (%016b), Squawk: %+v, AltitudeType: %+v, Callsign: \"%s\", EmitterType: %+v, Tslc: %+v }",
		m.IcaoAddress,
		m.Lat,
		m.Lon,
		m.Altitude,
		m.Heading,
		m.HorVelocity,
		m.VerVelocity,
		m.Flags.Bitmask(), uint64(m.Flags),
		m.Squawk,
		m.AltitudeType,
		m.Callsign,
		m.EmitterType,
		m.Tslc,
	)
}

// Marshal (generated function)
func (m *AdsbVehicle) Marshal() ([]byte, error) {
	payload := make([]byte, 38)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.IcaoAddress))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Lat))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.Lon))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.Altitude))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.Heading))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.HorVelocity))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.VerVelocity))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.Flags))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.Squawk))
	payload[26] = byte(m.AltitudeType)
	copy(payload[27:], m.Callsign)
	payload[36] = byte(m.EmitterType)
	payload[37] = byte(m.Tslc)
	return payload, nil
}

// Unmarshal (generated function)
func (m *AdsbVehicle) Unmarshal(data []byte) error {
	payload := make([]byte, 38)
	copy(payload[0:], data)
	m.IcaoAddress = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Lat = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.Lon = int32(binary.LittleEndian.Uint32(payload[8:]))
	m.Altitude = int32(binary.LittleEndian.Uint32(payload[12:]))
	m.Heading = uint16(binary.LittleEndian.Uint16(payload[16:]))
	m.HorVelocity = uint16(binary.LittleEndian.Uint16(payload[18:]))
	m.VerVelocity = int16(binary.LittleEndian.Uint16(payload[20:]))
	m.Flags = ADSB_FLAGS(binary.LittleEndian.Uint16(payload[22:]))
	m.Squawk = uint16(binary.LittleEndian.Uint16(payload[24:]))
	m.AltitudeType = ADSB_ALTITUDE_TYPE(payload[26])
	m.Callsign = strings.TrimRight(string(payload[27:36]), string(byte(0)))
	m.EmitterType = ADSB_EMITTER_TYPE(payload[36])
	m.Tslc = uint8(payload[37])
	return nil
}

// Collision struct (generated typeinfo)
// Information about a potential collision
type Collision struct {
	ID                     uint32                     // Unique identifier, domain based on src field
	TimeToMinimumDelta     float32                    // Estimated time until collision occurs
	AltitudeMinimumDelta   float32                    // Closest vertical distance between vehicle and object
	HorizontalMinimumDelta float32                    // Closest horizontal distance between vehicle and object
	Src                    MAV_COLLISION_SRC          // Collision data source
	Action                 MAV_COLLISION_ACTION       // Action that is being taken to avoid this collision
	ThreatLevel            MAV_COLLISION_THREAT_LEVEL // How concerned the aircraft is about this collision
}

// MsgID (generated function)
func (m *Collision) MsgID() message.MessageID {
	return MSG_ID_COLLISION
}

// String (generated function)
func (m *Collision) String() string {
	return fmt.Sprintf(
		"&common.Collision{ ID: %+v, TimeToMinimumDelta: %+v, AltitudeMinimumDelta: %+v, HorizontalMinimumDelta: %+v, Src: %+v, Action: %+v, ThreatLevel: %+v }",
		m.ID,
		m.TimeToMinimumDelta,
		m.AltitudeMinimumDelta,
		m.HorizontalMinimumDelta,
		m.Src,
		m.Action,
		m.ThreatLevel,
	)
}

// Marshal (generated function)
func (m *Collision) Marshal() ([]byte, error) {
	payload := make([]byte, 19)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.ID))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.TimeToMinimumDelta))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.AltitudeMinimumDelta))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.HorizontalMinimumDelta))
	payload[16] = byte(m.Src)
	payload[17] = byte(m.Action)
	payload[18] = byte(m.ThreatLevel)
	return payload, nil
}

// Unmarshal (generated function)
func (m *Collision) Unmarshal(data []byte) error {
	payload := make([]byte, 19)
	copy(payload[0:], data)
	m.ID = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.TimeToMinimumDelta = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.AltitudeMinimumDelta = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.HorizontalMinimumDelta = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Src = MAV_COLLISION_SRC(payload[16])
	m.Action = MAV_COLLISION_ACTION(payload[17])
	m.ThreatLevel = MAV_COLLISION_THREAT_LEVEL(payload[18])
	return nil
}

// V2Extension struct (generated typeinfo)
// Message implementing parts of the V2 payload specs in V1 frames for transitional support.
type V2Extension struct {
	MessageType     uint16  // A code that identifies the software component that understands this message (analogous to USB device classes or mime type strings). If this code is less than 32768, it is considered a 'registered' protocol extension and the corresponding entry should be added to https://github.com/mavlink/mavlink/definition_files/extension_message_ids.xml. Software creators can register blocks of message IDs as needed (useful for GCS specific metadata, etc...). Message_types greater than 32767 are considered local experiments and should not be checked in to any widely distributed codebase.
	TargetNetwork   uint8   // Network ID (0 for broadcast)
	TargetSystem    uint8   // System ID (0 for broadcast)
	TargetComponent uint8   // Component ID (0 for broadcast)
	Payload         []uint8 `len:"249" ` // Variable length payload. The length must be encoded in the payload as part of the message_type protocol, e.g. by including the length as payload data, or by terminating the payload data with a non-zero marker. This is required in order to reconstruct zero-terminated payloads that are (or otherwise would be) trimmed by MAVLink 2 empty-byte truncation. The entire content of the payload block is opaque unless you understand the encoding message_type. The particular encoding used can be extension specific and might not always be documented as part of the MAVLink specification.
}

// MsgID (generated function)
func (m *V2Extension) MsgID() message.MessageID {
	return MSG_ID_V2_EXTENSION
}

// String (generated function)
func (m *V2Extension) String() string {
	return fmt.Sprintf(
		"&common.V2Extension{ MessageType: %+v, TargetNetwork: %+v, TargetSystem: %+v, TargetComponent: %+v, Payload: %0X (\"%s\") }",
		m.MessageType,
		m.TargetNetwork,
		m.TargetSystem,
		m.TargetComponent,
		m.Payload, string(m.Payload[:]),
	)
}

// Marshal (generated function)
func (m *V2Extension) Marshal() ([]byte, error) {
	payload := make([]byte, 254)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.MessageType))
	payload[2] = byte(m.TargetNetwork)
	payload[3] = byte(m.TargetSystem)
	payload[4] = byte(m.TargetComponent)
	copy(payload[5:], m.Payload)
	return payload, nil
}

// Unmarshal (generated function)
func (m *V2Extension) Unmarshal(data []byte) error {
	payload := make([]byte, 254)
	copy(payload[0:], data)
	m.MessageType = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.TargetNetwork = uint8(payload[2])
	m.TargetSystem = uint8(payload[3])
	m.TargetComponent = uint8(payload[4])
	copy(m.Payload[:], payload[5:254])
	return nil
}

// MemoryVect struct (generated typeinfo)
// Send raw controller memory. The use of this message is discouraged for normal packets, but a quite efficient way for testing new messages and getting experimental debug output.
type MemoryVect struct {
	Address uint16 // Starting address of the debug variables
	Ver     uint8  // Version code of the type variable. 0=unknown, type ignored and assumed int16_t. 1=as below
	Type    uint8  // Type code of the memory variables. for ver = 1: 0=16 x int16_t, 1=16 x uint16_t, 2=16 x Q15, 3=16 x 1Q14
	Value   []int8 `len:"32" ` // Memory contents at specified address
}

// MsgID (generated function)
func (m *MemoryVect) MsgID() message.MessageID {
	return MSG_ID_MEMORY_VECT
}

// String (generated function)
func (m *MemoryVect) String() string {
	return fmt.Sprintf(
		"&common.MemoryVect{ Address: %+v, Ver: %+v, Type: %+v, Value: %+v }",
		m.Address,
		m.Ver,
		m.Type,
		m.Value,
	)
}

// Marshal (generated function)
func (m *MemoryVect) Marshal() ([]byte, error) {
	payload := make([]byte, 36)
	binary.LittleEndian.PutUint16(payload[0:], uint16(m.Address))
	payload[2] = byte(m.Ver)
	payload[3] = byte(m.Type)
	for i, v := range m.Value {
		payload[4+i*1] = byte(v)
	}
	return payload, nil
}

// Unmarshal (generated function)
func (m *MemoryVect) Unmarshal(data []byte) error {
	payload := make([]byte, 36)
	copy(payload[0:], data)
	m.Address = uint16(binary.LittleEndian.Uint16(payload[0:]))
	m.Ver = uint8(payload[2])
	m.Type = uint8(payload[3])
	for i := 0; i < len(m.Value); i++ {
		m.Value[i] = int8(payload[4+i*1])
	}
	return nil
}

// DebugVect struct (generated typeinfo)
// To debug something using a named 3D vector.
type DebugVect struct {
	TimeUsec uint64  // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	X        float32 // x
	Y        float32 // y
	Z        float32 // z
	Name     string  `len:"10" ` // Name
}

// MsgID (generated function)
func (m *DebugVect) MsgID() message.MessageID {
	return MSG_ID_DEBUG_VECT
}

// String (generated function)
func (m *DebugVect) String() string {
	return fmt.Sprintf(
		"&common.DebugVect{ TimeUsec: %+v, X: %+v, Y: %+v, Z: %+v, Name: \"%s\" }",
		m.TimeUsec,
		m.X,
		m.Y,
		m.Z,
		m.Name,
	)
}

// Marshal (generated function)
func (m *DebugVect) Marshal() ([]byte, error) {
	payload := make([]byte, 30)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.TimeUsec))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.X))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Y))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Z))
	copy(payload[20:], m.Name)
	return payload, nil
}

// Unmarshal (generated function)
func (m *DebugVect) Unmarshal(data []byte) error {
	payload := make([]byte, 30)
	copy(payload[0:], data)
	m.TimeUsec = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Name = strings.TrimRight(string(payload[20:30]), string(byte(0)))
	return nil
}

// NamedValueFloat struct (generated typeinfo)
// Send a key-value pair as float. The use of this message is discouraged for normal packets, but a quite efficient way for testing new messages and getting experimental debug output.
type NamedValueFloat struct {
	TimeBootMs uint32  // Timestamp (time since system boot).
	Value      float32 // Floating point value
	Name       string  `len:"10" ` // Name of the debug variable
}

// MsgID (generated function)
func (m *NamedValueFloat) MsgID() message.MessageID {
	return MSG_ID_NAMED_VALUE_FLOAT
}

// String (generated function)
func (m *NamedValueFloat) String() string {
	return fmt.Sprintf(
		"&common.NamedValueFloat{ TimeBootMs: %+v, Value: %+v, Name: \"%s\" }",
		m.TimeBootMs,
		m.Value,
		m.Name,
	)
}

// Marshal (generated function)
func (m *NamedValueFloat) Marshal() ([]byte, error) {
	payload := make([]byte, 18)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Value))
	copy(payload[8:], m.Name)
	return payload, nil
}

// Unmarshal (generated function)
func (m *NamedValueFloat) Unmarshal(data []byte) error {
	payload := make([]byte, 18)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Value = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Name = strings.TrimRight(string(payload[8:18]), string(byte(0)))
	return nil
}

// NamedValueInt struct (generated typeinfo)
// Send a key-value pair as integer. The use of this message is discouraged for normal packets, but a quite efficient way for testing new messages and getting experimental debug output.
type NamedValueInt struct {
	TimeBootMs uint32 // Timestamp (time since system boot).
	Value      int32  // Signed integer value
	Name       string `len:"10" ` // Name of the debug variable
}

// MsgID (generated function)
func (m *NamedValueInt) MsgID() message.MessageID {
	return MSG_ID_NAMED_VALUE_INT
}

// String (generated function)
func (m *NamedValueInt) String() string {
	return fmt.Sprintf(
		"&common.NamedValueInt{ TimeBootMs: %+v, Value: %+v, Name: \"%s\" }",
		m.TimeBootMs,
		m.Value,
		m.Name,
	)
}

// Marshal (generated function)
func (m *NamedValueInt) Marshal() ([]byte, error) {
	payload := make([]byte, 18)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], uint32(m.Value))
	copy(payload[8:], m.Name)
	return payload, nil
}

// Unmarshal (generated function)
func (m *NamedValueInt) Unmarshal(data []byte) error {
	payload := make([]byte, 18)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Value = int32(binary.LittleEndian.Uint32(payload[4:]))
	m.Name = strings.TrimRight(string(payload[8:18]), string(byte(0)))
	return nil
}

// Statustext struct (generated typeinfo)
// Status text message. These messages are printed in yellow in the COMM console of QGroundControl. WARNING: They consume quite some bandwidth, so use only for important status and error messages. If implemented wisely, these messages are buffered on the MCU and sent only at a limited rate (e.g. 10 Hz).
type Statustext struct {
	Severity MAV_SEVERITY // Severity of status. Relies on the definitions within RFC-5424.
	Text     string       `len:"50" ` // Status text message, without null termination character
}

// MsgID (generated function)
func (m *Statustext) MsgID() message.MessageID {
	return MSG_ID_STATUSTEXT
}

// String (generated function)
func (m *Statustext) String() string {
	return fmt.Sprintf(
		"&common.Statustext{ Severity: %+v, Text: \"%s\" }",
		m.Severity,
		m.Text,
	)
}

// Marshal (generated function)
func (m *Statustext) Marshal() ([]byte, error) {
	payload := make([]byte, 51)
	payload[0] = byte(m.Severity)
	copy(payload[1:], m.Text)
	return payload, nil
}

// Unmarshal (generated function)
func (m *Statustext) Unmarshal(data []byte) error {
	payload := make([]byte, 51)
	copy(payload[0:], data)
	m.Severity = MAV_SEVERITY(payload[0])
	m.Text = strings.TrimRight(string(payload[1:51]), string(byte(0)))
	return nil
}

// Debug struct (generated typeinfo)
// Send a debug value. The index is used to discriminate between values. These values show up in the plot of QGroundControl as DEBUG N.
type Debug struct {
	TimeBootMs uint32  // Timestamp (time since system boot).
	Value      float32 // DEBUG value
	Ind        uint8   // index of debug variable
}

// MsgID (generated function)
func (m *Debug) MsgID() message.MessageID {
	return MSG_ID_DEBUG
}

// String (generated function)
func (m *Debug) String() string {
	return fmt.Sprintf(
		"&common.Debug{ TimeBootMs: %+v, Value: %+v, Ind: %+v }",
		m.TimeBootMs,
		m.Value,
		m.Ind,
	)
}

// Marshal (generated function)
func (m *Debug) Marshal() ([]byte, error) {
	payload := make([]byte, 9)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.TimeBootMs))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Value))
	payload[8] = byte(m.Ind)
	return payload, nil
}

// Unmarshal (generated function)
func (m *Debug) Unmarshal(data []byte) error {
	payload := make([]byte, 9)
	copy(payload[0:], data)
	m.TimeBootMs = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Value = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Ind = uint8(payload[8])
	return nil
}

// Heartbeat struct (generated typeinfo)
// The heartbeat message shows that a system or component is present and responding. The type and autopilot fields (along with the message component id), allow the receiving system to treat further messages from this system appropriately (e.g. by laying out the user interface based on the autopilot). This microservice is documented at https://mavlink.io/en/services/heartbeat.html
type Heartbeat struct {
	CustomMode     uint32        // A bitfield for use for autopilot-specific flags
	Type           MAV_TYPE      // Vehicle or component type. For a flight controller component the vehicle type (quadrotor, helicopter, etc.). For other components the component type (e.g. camera, gimbal, etc.). This should be used in preference to component id for identifying the component type.
	Autopilot      MAV_AUTOPILOT // Autopilot type / class. Use MAV_AUTOPILOT_INVALID for components that are not flight controllers.
	BaseMode       MAV_MODE_FLAG // System mode bitmap.
	SystemStatus   MAV_STATE     // System status flag.
	MavlinkVersion uint8         // MAVLink version, not writable by user, gets added by protocol because of magic data type: uint8_t_mavlink_version
}

// MsgID (generated function)
func (m *Heartbeat) MsgID() message.MessageID {
	return MSG_ID_HEARTBEAT
}

// String (generated function)
func (m *Heartbeat) String() string {
	return fmt.Sprintf(
		"&minimal.Heartbeat{ CustomMode: %+v, Type: %+v, Autopilot: %+v, BaseMode: %+v (%08b), SystemStatus: %+v, MavlinkVersion: %+v }",
		m.CustomMode,
		m.Type,
		m.Autopilot,
		m.BaseMode.Bitmask(), uint64(m.BaseMode),
		m.SystemStatus,
		m.MavlinkVersion,
	)
}

// Marshal (generated function)
func (m *Heartbeat) Marshal() ([]byte, error) {
	payload := make([]byte, 9)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.CustomMode))
	payload[4] = byte(m.Type)
	payload[5] = byte(m.Autopilot)
	payload[6] = byte(m.BaseMode)
	payload[7] = byte(m.SystemStatus)
	payload[8] = byte(m.MavlinkVersion)
	return payload, nil
}

// Unmarshal (generated function)
func (m *Heartbeat) Unmarshal(data []byte) error {
	payload := make([]byte, 9)
	copy(payload[0:], data)
	m.CustomMode = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.Type = MAV_TYPE(payload[4])
	m.Autopilot = MAV_AUTOPILOT(payload[5])
	m.BaseMode = MAV_MODE_FLAG(payload[6])
	m.SystemStatus = MAV_STATE(payload[7])
	m.MavlinkVersion = uint8(payload[8])
	return nil
}
