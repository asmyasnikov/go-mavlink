/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package asluav

import (
	"encoding/binary"
	"fmt"
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"math"
	"strings"
)

// CommandIntStamped struct (generated typeinfo)
// Message encoding a command with parameters as scaled integers and additional metadata. Scaling depends on the actual command value.
type CommandIntStamped struct {
	VehicleTimestamp uint64    `json:"VehicleTimestamp" ` // Microseconds elapsed since vehicle boot
	UtcTime          uint32    `json:"UtcTime" `          // UTC time, seconds elapsed since 01.01.1970
	Param1           float32   `json:"Param1" `           // PARAM1, see MAV_CMD enum
	Param2           float32   `json:"Param2" `           // PARAM2, see MAV_CMD enum
	Param3           float32   `json:"Param3" `           // PARAM3, see MAV_CMD enum
	Param4           float32   `json:"Param4" `           // PARAM4, see MAV_CMD enum
	X                int32     `json:"X" `                // PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
	Y                int32     `json:"Y" `                // PARAM6 / local: y position in meters * 1e4, global: longitude in degrees * 10^7
	Z                float32   `json:"Z" `                // PARAM7 / z position: global: altitude in meters (MSL, WGS84, AGL or relative to home - depending on frame).
	Command          MAV_CMD   `json:"Command" `          // The scheduled action for the mission item, as defined by MAV_CMD enum
	TargetSystem     uint8     `json:"TargetSystem" `     // System ID
	TargetComponent  uint8     `json:"TargetComponent" `  // Component ID
	Frame            MAV_FRAME `json:"Frame" `            // The coordinate system of the COMMAND, as defined by MAV_FRAME enum
	Current          uint8     `json:"Current" `          // false:0, true:1
	Autocontinue     uint8     `json:"Autocontinue" `     // autocontinue to next wp
}

// MsgID (generated function)
func (m *CommandIntStamped) MsgID() message.MessageID {
	return MSG_ID_COMMAND_INT_STAMPED
}

// String (generated function)
func (m *CommandIntStamped) String() string {
	return fmt.Sprintf(
		"&ASLUAV.CommandIntStamped{ VehicleTimestamp: %+v, UtcTime: %+v, Param1: %+v, Param2: %+v, Param3: %+v, Param4: %+v, X: %+v, Y: %+v, Z: %+v, Command: %+v, TargetSystem: %+v, TargetComponent: %+v, Frame: %+v, Current: %+v, Autocontinue: %+v }",
		m.VehicleTimestamp,
		m.UtcTime,
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

const (
	COMMAND_INT_STAMPED_FIELD_VEHICLE_TIMESTAMP = "CommandIntStamped.VehicleTimestamp"
	COMMAND_INT_STAMPED_FIELD_UTC_TIME          = "CommandIntStamped.UtcTime"
	COMMAND_INT_STAMPED_FIELD_PARAM1            = "CommandIntStamped.Param1"
	COMMAND_INT_STAMPED_FIELD_PARAM2            = "CommandIntStamped.Param2"
	COMMAND_INT_STAMPED_FIELD_PARAM3            = "CommandIntStamped.Param3"
	COMMAND_INT_STAMPED_FIELD_PARAM4            = "CommandIntStamped.Param4"
	COMMAND_INT_STAMPED_FIELD_X                 = "CommandIntStamped.X"
	COMMAND_INT_STAMPED_FIELD_Y                 = "CommandIntStamped.Y"
	COMMAND_INT_STAMPED_FIELD_Z                 = "CommandIntStamped.Z"
	COMMAND_INT_STAMPED_FIELD_COMMAND           = "CommandIntStamped.Command"
	COMMAND_INT_STAMPED_FIELD_TARGET_SYSTEM     = "CommandIntStamped.TargetSystem"
	COMMAND_INT_STAMPED_FIELD_TARGET_COMPONENT  = "CommandIntStamped.TargetComponent"
	COMMAND_INT_STAMPED_FIELD_FRAME             = "CommandIntStamped.Frame"
	COMMAND_INT_STAMPED_FIELD_CURRENT           = "CommandIntStamped.Current"
	COMMAND_INT_STAMPED_FIELD_AUTOCONTINUE      = "CommandIntStamped.Autocontinue"
)

// ToMap (generated function)
func (m *CommandIntStamped) Dict() map[string]interface{} {
	return map[string]interface{}{
		COMMAND_INT_STAMPED_FIELD_VEHICLE_TIMESTAMP: m.VehicleTimestamp,
		COMMAND_INT_STAMPED_FIELD_UTC_TIME:          m.UtcTime,
		COMMAND_INT_STAMPED_FIELD_PARAM1:            m.Param1,
		COMMAND_INT_STAMPED_FIELD_PARAM2:            m.Param2,
		COMMAND_INT_STAMPED_FIELD_PARAM3:            m.Param3,
		COMMAND_INT_STAMPED_FIELD_PARAM4:            m.Param4,
		COMMAND_INT_STAMPED_FIELD_X:                 m.X,
		COMMAND_INT_STAMPED_FIELD_Y:                 m.Y,
		COMMAND_INT_STAMPED_FIELD_Z:                 m.Z,
		COMMAND_INT_STAMPED_FIELD_COMMAND:           m.Command,
		COMMAND_INT_STAMPED_FIELD_TARGET_SYSTEM:     m.TargetSystem,
		COMMAND_INT_STAMPED_FIELD_TARGET_COMPONENT:  m.TargetComponent,
		COMMAND_INT_STAMPED_FIELD_FRAME:             m.Frame,
		COMMAND_INT_STAMPED_FIELD_CURRENT:           m.Current,
		COMMAND_INT_STAMPED_FIELD_AUTOCONTINUE:      m.Autocontinue,
	}
}

// Marshal (generated function)
func (m *CommandIntStamped) Marshal() ([]byte, error) {
	payload := make([]byte, 47)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.VehicleTimestamp))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.UtcTime))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Param1))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Param2))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Param3))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Param4))
	binary.LittleEndian.PutUint32(payload[28:], uint32(m.X))
	binary.LittleEndian.PutUint32(payload[32:], uint32(m.Y))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Z))
	binary.LittleEndian.PutUint16(payload[40:], uint16(m.Command))
	payload[42] = byte(m.TargetSystem)
	payload[43] = byte(m.TargetComponent)
	payload[44] = byte(m.Frame)
	payload[45] = byte(m.Current)
	payload[46] = byte(m.Autocontinue)
	return payload, nil
}

// Unmarshal (generated function)
func (m *CommandIntStamped) Unmarshal(data []byte) error {
	payload := make([]byte, 47)
	copy(payload[0:], data)
	m.VehicleTimestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.UtcTime = uint32(binary.LittleEndian.Uint32(payload[8:]))
	m.Param1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Param2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Param3 = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Param4 = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.X = int32(binary.LittleEndian.Uint32(payload[28:]))
	m.Y = int32(binary.LittleEndian.Uint32(payload[32:]))
	m.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Command = MAV_CMD(binary.LittleEndian.Uint16(payload[40:]))
	m.TargetSystem = uint8(payload[42])
	m.TargetComponent = uint8(payload[43])
	m.Frame = MAV_FRAME(payload[44])
	m.Current = uint8(payload[45])
	m.Autocontinue = uint8(payload[46])
	return nil
}

// CommandLongStamped struct (generated typeinfo)
// Send a command with up to seven parameters to the MAV and additional metadata
type CommandLongStamped struct {
	VehicleTimestamp uint64  `json:"VehicleTimestamp" ` // Microseconds elapsed since vehicle boot
	UtcTime          uint32  `json:"UtcTime" `          // UTC time, seconds elapsed since 01.01.1970
	Param1           float32 `json:"Param1" `           // Parameter 1, as defined by MAV_CMD enum.
	Param2           float32 `json:"Param2" `           // Parameter 2, as defined by MAV_CMD enum.
	Param3           float32 `json:"Param3" `           // Parameter 3, as defined by MAV_CMD enum.
	Param4           float32 `json:"Param4" `           // Parameter 4, as defined by MAV_CMD enum.
	Param5           float32 `json:"Param5" `           // Parameter 5, as defined by MAV_CMD enum.
	Param6           float32 `json:"Param6" `           // Parameter 6, as defined by MAV_CMD enum.
	Param7           float32 `json:"Param7" `           // Parameter 7, as defined by MAV_CMD enum.
	Command          MAV_CMD `json:"Command" `          // Command ID, as defined by MAV_CMD enum.
	TargetSystem     uint8   `json:"TargetSystem" `     // System which should execute the command
	TargetComponent  uint8   `json:"TargetComponent" `  // Component which should execute the command, 0 for all components
	Confirmation     uint8   `json:"Confirmation" `     // 0: First transmission of this command. 1-255: Confirmation transmissions (e.g. for kill command)
}

// MsgID (generated function)
func (m *CommandLongStamped) MsgID() message.MessageID {
	return MSG_ID_COMMAND_LONG_STAMPED
}

// String (generated function)
func (m *CommandLongStamped) String() string {
	return fmt.Sprintf(
		"&ASLUAV.CommandLongStamped{ VehicleTimestamp: %+v, UtcTime: %+v, Param1: %+v, Param2: %+v, Param3: %+v, Param4: %+v, Param5: %+v, Param6: %+v, Param7: %+v, Command: %+v, TargetSystem: %+v, TargetComponent: %+v, Confirmation: %+v }",
		m.VehicleTimestamp,
		m.UtcTime,
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

const (
	COMMAND_LONG_STAMPED_FIELD_VEHICLE_TIMESTAMP = "CommandLongStamped.VehicleTimestamp"
	COMMAND_LONG_STAMPED_FIELD_UTC_TIME          = "CommandLongStamped.UtcTime"
	COMMAND_LONG_STAMPED_FIELD_PARAM1            = "CommandLongStamped.Param1"
	COMMAND_LONG_STAMPED_FIELD_PARAM2            = "CommandLongStamped.Param2"
	COMMAND_LONG_STAMPED_FIELD_PARAM3            = "CommandLongStamped.Param3"
	COMMAND_LONG_STAMPED_FIELD_PARAM4            = "CommandLongStamped.Param4"
	COMMAND_LONG_STAMPED_FIELD_PARAM5            = "CommandLongStamped.Param5"
	COMMAND_LONG_STAMPED_FIELD_PARAM6            = "CommandLongStamped.Param6"
	COMMAND_LONG_STAMPED_FIELD_PARAM7            = "CommandLongStamped.Param7"
	COMMAND_LONG_STAMPED_FIELD_COMMAND           = "CommandLongStamped.Command"
	COMMAND_LONG_STAMPED_FIELD_TARGET_SYSTEM     = "CommandLongStamped.TargetSystem"
	COMMAND_LONG_STAMPED_FIELD_TARGET_COMPONENT  = "CommandLongStamped.TargetComponent"
	COMMAND_LONG_STAMPED_FIELD_CONFIRMATION      = "CommandLongStamped.Confirmation"
)

// ToMap (generated function)
func (m *CommandLongStamped) Dict() map[string]interface{} {
	return map[string]interface{}{
		COMMAND_LONG_STAMPED_FIELD_VEHICLE_TIMESTAMP: m.VehicleTimestamp,
		COMMAND_LONG_STAMPED_FIELD_UTC_TIME:          m.UtcTime,
		COMMAND_LONG_STAMPED_FIELD_PARAM1:            m.Param1,
		COMMAND_LONG_STAMPED_FIELD_PARAM2:            m.Param2,
		COMMAND_LONG_STAMPED_FIELD_PARAM3:            m.Param3,
		COMMAND_LONG_STAMPED_FIELD_PARAM4:            m.Param4,
		COMMAND_LONG_STAMPED_FIELD_PARAM5:            m.Param5,
		COMMAND_LONG_STAMPED_FIELD_PARAM6:            m.Param6,
		COMMAND_LONG_STAMPED_FIELD_PARAM7:            m.Param7,
		COMMAND_LONG_STAMPED_FIELD_COMMAND:           m.Command,
		COMMAND_LONG_STAMPED_FIELD_TARGET_SYSTEM:     m.TargetSystem,
		COMMAND_LONG_STAMPED_FIELD_TARGET_COMPONENT:  m.TargetComponent,
		COMMAND_LONG_STAMPED_FIELD_CONFIRMATION:      m.Confirmation,
	}
}

// Marshal (generated function)
func (m *CommandLongStamped) Marshal() ([]byte, error) {
	payload := make([]byte, 45)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.VehicleTimestamp))
	binary.LittleEndian.PutUint32(payload[8:], uint32(m.UtcTime))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Param1))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Param2))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Param3))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Param4))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Param5))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Param6))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Param7))
	binary.LittleEndian.PutUint16(payload[40:], uint16(m.Command))
	payload[42] = byte(m.TargetSystem)
	payload[43] = byte(m.TargetComponent)
	payload[44] = byte(m.Confirmation)
	return payload, nil
}

// Unmarshal (generated function)
func (m *CommandLongStamped) Unmarshal(data []byte) error {
	payload := make([]byte, 45)
	copy(payload[0:], data)
	m.VehicleTimestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.UtcTime = uint32(binary.LittleEndian.Uint32(payload[8:]))
	m.Param1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Param2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Param3 = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Param4 = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Param5 = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Param6 = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Param7 = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Command = MAV_CMD(binary.LittleEndian.Uint16(payload[40:]))
	m.TargetSystem = uint8(payload[42])
	m.TargetComponent = uint8(payload[43])
	m.Confirmation = uint8(payload[44])
	return nil
}

// SensPower struct (generated typeinfo)
// Voltage and current sensor data
type SensPower struct {
	Adc121VspbVolt float32 `json:"Adc121VspbVolt" ` // [ V ]  Power board voltage sensor reading
	Adc121CspbAmp  float32 `json:"Adc121CspbAmp" `  // [ A ]  Power board current sensor reading
	Adc121Cs1Amp   float32 `json:"Adc121Cs1Amp" `   // [ A ]  Board current sensor 1 reading
	Adc121Cs2Amp   float32 `json:"Adc121Cs2Amp" `   // [ A ]  Board current sensor 2 reading
}

// MsgID (generated function)
func (m *SensPower) MsgID() message.MessageID {
	return MSG_ID_SENS_POWER
}

// String (generated function)
func (m *SensPower) String() string {
	return fmt.Sprintf(
		"&ASLUAV.SensPower{ Adc121VspbVolt: %+v, Adc121CspbAmp: %+v, Adc121Cs1Amp: %+v, Adc121Cs2Amp: %+v }",
		m.Adc121VspbVolt,
		m.Adc121CspbAmp,
		m.Adc121Cs1Amp,
		m.Adc121Cs2Amp,
	)
}

const (
	SENS_POWER_FIELD_ADC121_VSPB_VOLT = "SensPower.Adc121VspbVolt"
	SENS_POWER_FIELD_ADC121_CSPB_AMP  = "SensPower.Adc121CspbAmp"
	SENS_POWER_FIELD_ADC121_CS1_AMP   = "SensPower.Adc121Cs1Amp"
	SENS_POWER_FIELD_ADC121_CS2_AMP   = "SensPower.Adc121Cs2Amp"
)

// ToMap (generated function)
func (m *SensPower) Dict() map[string]interface{} {
	return map[string]interface{}{
		SENS_POWER_FIELD_ADC121_VSPB_VOLT: m.Adc121VspbVolt,
		SENS_POWER_FIELD_ADC121_CSPB_AMP:  m.Adc121CspbAmp,
		SENS_POWER_FIELD_ADC121_CS1_AMP:   m.Adc121Cs1Amp,
		SENS_POWER_FIELD_ADC121_CS2_AMP:   m.Adc121Cs2Amp,
	}
}

// Marshal (generated function)
func (m *SensPower) Marshal() ([]byte, error) {
	payload := make([]byte, 16)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.Adc121VspbVolt))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.Adc121CspbAmp))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Adc121Cs1Amp))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Adc121Cs2Amp))
	return payload, nil
}

// Unmarshal (generated function)
func (m *SensPower) Unmarshal(data []byte) error {
	payload := make([]byte, 16)
	copy(payload[0:], data)
	m.Adc121VspbVolt = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.Adc121CspbAmp = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.Adc121Cs1Amp = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Adc121Cs2Amp = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	return nil
}

// SensMppt struct (generated typeinfo)
// Maximum Power Point Tracker (MPPT) sensor data for solar module power performance tracking
type SensMppt struct {
	MpptTimestamp uint64  `json:"MpptTimestamp" ` // [ us ]  MPPT last timestamp
	Mppt1Volt     float32 `json:"Mppt1Volt" `     // [ V ]  MPPT1 voltage
	Mppt1Amp      float32 `json:"Mppt1Amp" `      // [ A ]  MPPT1 current
	Mppt2Volt     float32 `json:"Mppt2Volt" `     // [ V ]  MPPT2 voltage
	Mppt2Amp      float32 `json:"Mppt2Amp" `      // [ A ]  MPPT2 current
	Mppt3Volt     float32 `json:"Mppt3Volt" `     // [ V ] MPPT3 voltage
	Mppt3Amp      float32 `json:"Mppt3Amp" `      // [ A ]  MPPT3 current
	Mppt1Pwm      uint16  `json:"Mppt1Pwm" `      // [ us ]  MPPT1 pwm
	Mppt2Pwm      uint16  `json:"Mppt2Pwm" `      // [ us ]  MPPT2 pwm
	Mppt3Pwm      uint16  `json:"Mppt3Pwm" `      // [ us ]  MPPT3 pwm
	Mppt1Status   uint8   `json:"Mppt1Status" `   //  MPPT1 status
	Mppt2Status   uint8   `json:"Mppt2Status" `   //  MPPT2 status
	Mppt3Status   uint8   `json:"Mppt3Status" `   //  MPPT3 status
}

// MsgID (generated function)
func (m *SensMppt) MsgID() message.MessageID {
	return MSG_ID_SENS_MPPT
}

// String (generated function)
func (m *SensMppt) String() string {
	return fmt.Sprintf(
		"&ASLUAV.SensMppt{ MpptTimestamp: %+v, Mppt1Volt: %+v, Mppt1Amp: %+v, Mppt2Volt: %+v, Mppt2Amp: %+v, Mppt3Volt: %+v, Mppt3Amp: %+v, Mppt1Pwm: %+v, Mppt2Pwm: %+v, Mppt3Pwm: %+v, Mppt1Status: %+v, Mppt2Status: %+v, Mppt3Status: %+v }",
		m.MpptTimestamp,
		m.Mppt1Volt,
		m.Mppt1Amp,
		m.Mppt2Volt,
		m.Mppt2Amp,
		m.Mppt3Volt,
		m.Mppt3Amp,
		m.Mppt1Pwm,
		m.Mppt2Pwm,
		m.Mppt3Pwm,
		m.Mppt1Status,
		m.Mppt2Status,
		m.Mppt3Status,
	)
}

const (
	SENS_MPPT_FIELD_MPPT_TIMESTAMP = "SensMppt.MpptTimestamp"
	SENS_MPPT_FIELD_MPPT1_VOLT     = "SensMppt.Mppt1Volt"
	SENS_MPPT_FIELD_MPPT1_AMP      = "SensMppt.Mppt1Amp"
	SENS_MPPT_FIELD_MPPT2_VOLT     = "SensMppt.Mppt2Volt"
	SENS_MPPT_FIELD_MPPT2_AMP      = "SensMppt.Mppt2Amp"
	SENS_MPPT_FIELD_MPPT3_VOLT     = "SensMppt.Mppt3Volt"
	SENS_MPPT_FIELD_MPPT3_AMP      = "SensMppt.Mppt3Amp"
	SENS_MPPT_FIELD_MPPT1_PWM      = "SensMppt.Mppt1Pwm"
	SENS_MPPT_FIELD_MPPT2_PWM      = "SensMppt.Mppt2Pwm"
	SENS_MPPT_FIELD_MPPT3_PWM      = "SensMppt.Mppt3Pwm"
	SENS_MPPT_FIELD_MPPT1_STATUS   = "SensMppt.Mppt1Status"
	SENS_MPPT_FIELD_MPPT2_STATUS   = "SensMppt.Mppt2Status"
	SENS_MPPT_FIELD_MPPT3_STATUS   = "SensMppt.Mppt3Status"
)

// ToMap (generated function)
func (m *SensMppt) Dict() map[string]interface{} {
	return map[string]interface{}{
		SENS_MPPT_FIELD_MPPT_TIMESTAMP: m.MpptTimestamp,
		SENS_MPPT_FIELD_MPPT1_VOLT:     m.Mppt1Volt,
		SENS_MPPT_FIELD_MPPT1_AMP:      m.Mppt1Amp,
		SENS_MPPT_FIELD_MPPT2_VOLT:     m.Mppt2Volt,
		SENS_MPPT_FIELD_MPPT2_AMP:      m.Mppt2Amp,
		SENS_MPPT_FIELD_MPPT3_VOLT:     m.Mppt3Volt,
		SENS_MPPT_FIELD_MPPT3_AMP:      m.Mppt3Amp,
		SENS_MPPT_FIELD_MPPT1_PWM:      m.Mppt1Pwm,
		SENS_MPPT_FIELD_MPPT2_PWM:      m.Mppt2Pwm,
		SENS_MPPT_FIELD_MPPT3_PWM:      m.Mppt3Pwm,
		SENS_MPPT_FIELD_MPPT1_STATUS:   m.Mppt1Status,
		SENS_MPPT_FIELD_MPPT2_STATUS:   m.Mppt2Status,
		SENS_MPPT_FIELD_MPPT3_STATUS:   m.Mppt3Status,
	}
}

// Marshal (generated function)
func (m *SensMppt) Marshal() ([]byte, error) {
	payload := make([]byte, 41)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.MpptTimestamp))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Mppt1Volt))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Mppt1Amp))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Mppt2Volt))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Mppt2Amp))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Mppt3Volt))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Mppt3Amp))
	binary.LittleEndian.PutUint16(payload[32:], uint16(m.Mppt1Pwm))
	binary.LittleEndian.PutUint16(payload[34:], uint16(m.Mppt2Pwm))
	binary.LittleEndian.PutUint16(payload[36:], uint16(m.Mppt3Pwm))
	payload[38] = byte(m.Mppt1Status)
	payload[39] = byte(m.Mppt2Status)
	payload[40] = byte(m.Mppt3Status)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SensMppt) Unmarshal(data []byte) error {
	payload := make([]byte, 41)
	copy(payload[0:], data)
	m.MpptTimestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Mppt1Volt = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Mppt1Amp = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Mppt2Volt = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Mppt2Amp = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Mppt3Volt = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Mppt3Amp = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Mppt1Pwm = uint16(binary.LittleEndian.Uint16(payload[32:]))
	m.Mppt2Pwm = uint16(binary.LittleEndian.Uint16(payload[34:]))
	m.Mppt3Pwm = uint16(binary.LittleEndian.Uint16(payload[36:]))
	m.Mppt1Status = uint8(payload[38])
	m.Mppt2Status = uint8(payload[39])
	m.Mppt3Status = uint8(payload[40])
	return nil
}

// AslctrlData struct (generated typeinfo)
// ASL-fixed-wing controller data
type AslctrlData struct {
	Timestamp       uint64  `json:"Timestamp" `       // [ us ]  Timestamp
	H               float32 `json:"H" `               //  See sourcecode for a description of these values...
	Href            float32 `json:"Href" `            //
	HrefT           float32 `json:"HrefT" `           //
	Pitchangle      float32 `json:"Pitchangle" `      // [ deg ] Pitch angle
	Pitchangleref   float32 `json:"Pitchangleref" `   // [ deg ] Pitch angle reference
	Q               float32 `json:"Q" `               //
	Qref            float32 `json:"Qref" `            //
	Uelev           float32 `json:"Uelev" `           //
	Uthrot          float32 `json:"Uthrot" `          //
	Uthrot2         float32 `json:"Uthrot2" `         //
	Nz              float32 `json:"Nz" `              //
	Airspeedref     float32 `json:"Airspeedref" `     // [ m/s ] Airspeed reference
	Yawangle        float32 `json:"Yawangle" `        // [ deg ] Yaw angle
	Yawangleref     float32 `json:"Yawangleref" `     // [ deg ] Yaw angle reference
	Rollangle       float32 `json:"Rollangle" `       // [ deg ] Roll angle
	Rollangleref    float32 `json:"Rollangleref" `    // [ deg ] Roll angle reference
	P               float32 `json:"P" `               //
	Pref            float32 `json:"Pref" `            //
	R               float32 `json:"R" `               //
	Rref            float32 `json:"Rref" `            //
	Uail            float32 `json:"Uail" `            //
	Urud            float32 `json:"Urud" `            //
	AslctrlMode     uint8   `json:"AslctrlMode" `     //  ASLCTRL control-mode (manual, stabilized, auto, etc...)
	Spoilersengaged uint8   `json:"Spoilersengaged" ` //
}

// MsgID (generated function)
func (m *AslctrlData) MsgID() message.MessageID {
	return MSG_ID_ASLCTRL_DATA
}

// String (generated function)
func (m *AslctrlData) String() string {
	return fmt.Sprintf(
		"&ASLUAV.AslctrlData{ Timestamp: %+v, H: %+v, Href: %+v, HrefT: %+v, Pitchangle: %+v, Pitchangleref: %+v, Q: %+v, Qref: %+v, Uelev: %+v, Uthrot: %+v, Uthrot2: %+v, Nz: %+v, Airspeedref: %+v, Yawangle: %+v, Yawangleref: %+v, Rollangle: %+v, Rollangleref: %+v, P: %+v, Pref: %+v, R: %+v, Rref: %+v, Uail: %+v, Urud: %+v, AslctrlMode: %+v, Spoilersengaged: %+v }",
		m.Timestamp,
		m.H,
		m.Href,
		m.HrefT,
		m.Pitchangle,
		m.Pitchangleref,
		m.Q,
		m.Qref,
		m.Uelev,
		m.Uthrot,
		m.Uthrot2,
		m.Nz,
		m.Airspeedref,
		m.Yawangle,
		m.Yawangleref,
		m.Rollangle,
		m.Rollangleref,
		m.P,
		m.Pref,
		m.R,
		m.Rref,
		m.Uail,
		m.Urud,
		m.AslctrlMode,
		m.Spoilersengaged,
	)
}

const (
	ASLCTRL_DATA_FIELD_TIMESTAMP       = "AslctrlData.Timestamp"
	ASLCTRL_DATA_FIELD_H               = "AslctrlData.H"
	ASLCTRL_DATA_FIELD_HREF            = "AslctrlData.Href"
	ASLCTRL_DATA_FIELD_HREF_T          = "AslctrlData.HrefT"
	ASLCTRL_DATA_FIELD_PITCHANGLE      = "AslctrlData.Pitchangle"
	ASLCTRL_DATA_FIELD_PITCHANGLEREF   = "AslctrlData.Pitchangleref"
	ASLCTRL_DATA_FIELD_Q               = "AslctrlData.Q"
	ASLCTRL_DATA_FIELD_QREF            = "AslctrlData.Qref"
	ASLCTRL_DATA_FIELD_UELEV           = "AslctrlData.Uelev"
	ASLCTRL_DATA_FIELD_UTHROT          = "AslctrlData.Uthrot"
	ASLCTRL_DATA_FIELD_UTHROT2         = "AslctrlData.Uthrot2"
	ASLCTRL_DATA_FIELD_NZ              = "AslctrlData.Nz"
	ASLCTRL_DATA_FIELD_AIRSPEEDREF     = "AslctrlData.Airspeedref"
	ASLCTRL_DATA_FIELD_YAWANGLE        = "AslctrlData.Yawangle"
	ASLCTRL_DATA_FIELD_YAWANGLEREF     = "AslctrlData.Yawangleref"
	ASLCTRL_DATA_FIELD_ROLLANGLE       = "AslctrlData.Rollangle"
	ASLCTRL_DATA_FIELD_ROLLANGLEREF    = "AslctrlData.Rollangleref"
	ASLCTRL_DATA_FIELD_P               = "AslctrlData.P"
	ASLCTRL_DATA_FIELD_PREF            = "AslctrlData.Pref"
	ASLCTRL_DATA_FIELD_R               = "AslctrlData.R"
	ASLCTRL_DATA_FIELD_RREF            = "AslctrlData.Rref"
	ASLCTRL_DATA_FIELD_UAIL            = "AslctrlData.Uail"
	ASLCTRL_DATA_FIELD_URUD            = "AslctrlData.Urud"
	ASLCTRL_DATA_FIELD_ASLCTRL_MODE    = "AslctrlData.AslctrlMode"
	ASLCTRL_DATA_FIELD_SPOILERSENGAGED = "AslctrlData.Spoilersengaged"
)

// ToMap (generated function)
func (m *AslctrlData) Dict() map[string]interface{} {
	return map[string]interface{}{
		ASLCTRL_DATA_FIELD_TIMESTAMP:       m.Timestamp,
		ASLCTRL_DATA_FIELD_H:               m.H,
		ASLCTRL_DATA_FIELD_HREF:            m.Href,
		ASLCTRL_DATA_FIELD_HREF_T:          m.HrefT,
		ASLCTRL_DATA_FIELD_PITCHANGLE:      m.Pitchangle,
		ASLCTRL_DATA_FIELD_PITCHANGLEREF:   m.Pitchangleref,
		ASLCTRL_DATA_FIELD_Q:               m.Q,
		ASLCTRL_DATA_FIELD_QREF:            m.Qref,
		ASLCTRL_DATA_FIELD_UELEV:           m.Uelev,
		ASLCTRL_DATA_FIELD_UTHROT:          m.Uthrot,
		ASLCTRL_DATA_FIELD_UTHROT2:         m.Uthrot2,
		ASLCTRL_DATA_FIELD_NZ:              m.Nz,
		ASLCTRL_DATA_FIELD_AIRSPEEDREF:     m.Airspeedref,
		ASLCTRL_DATA_FIELD_YAWANGLE:        m.Yawangle,
		ASLCTRL_DATA_FIELD_YAWANGLEREF:     m.Yawangleref,
		ASLCTRL_DATA_FIELD_ROLLANGLE:       m.Rollangle,
		ASLCTRL_DATA_FIELD_ROLLANGLEREF:    m.Rollangleref,
		ASLCTRL_DATA_FIELD_P:               m.P,
		ASLCTRL_DATA_FIELD_PREF:            m.Pref,
		ASLCTRL_DATA_FIELD_R:               m.R,
		ASLCTRL_DATA_FIELD_RREF:            m.Rref,
		ASLCTRL_DATA_FIELD_UAIL:            m.Uail,
		ASLCTRL_DATA_FIELD_URUD:            m.Urud,
		ASLCTRL_DATA_FIELD_ASLCTRL_MODE:    m.AslctrlMode,
		ASLCTRL_DATA_FIELD_SPOILERSENGAGED: m.Spoilersengaged,
	}
}

// Marshal (generated function)
func (m *AslctrlData) Marshal() ([]byte, error) {
	payload := make([]byte, 98)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.H))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Href))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.HrefT))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Pitchangle))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Pitchangleref))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Q))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Qref))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Uelev))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Uthrot))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.Uthrot2))
	binary.LittleEndian.PutUint32(payload[48:], math.Float32bits(m.Nz))
	binary.LittleEndian.PutUint32(payload[52:], math.Float32bits(m.Airspeedref))
	binary.LittleEndian.PutUint32(payload[56:], math.Float32bits(m.Yawangle))
	binary.LittleEndian.PutUint32(payload[60:], math.Float32bits(m.Yawangleref))
	binary.LittleEndian.PutUint32(payload[64:], math.Float32bits(m.Rollangle))
	binary.LittleEndian.PutUint32(payload[68:], math.Float32bits(m.Rollangleref))
	binary.LittleEndian.PutUint32(payload[72:], math.Float32bits(m.P))
	binary.LittleEndian.PutUint32(payload[76:], math.Float32bits(m.Pref))
	binary.LittleEndian.PutUint32(payload[80:], math.Float32bits(m.R))
	binary.LittleEndian.PutUint32(payload[84:], math.Float32bits(m.Rref))
	binary.LittleEndian.PutUint32(payload[88:], math.Float32bits(m.Uail))
	binary.LittleEndian.PutUint32(payload[92:], math.Float32bits(m.Urud))
	payload[96] = byte(m.AslctrlMode)
	payload[97] = byte(m.Spoilersengaged)
	return payload, nil
}

// Unmarshal (generated function)
func (m *AslctrlData) Unmarshal(data []byte) error {
	payload := make([]byte, 98)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.H = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Href = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.HrefT = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Pitchangle = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Pitchangleref = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Q = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Qref = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Uelev = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Uthrot = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.Uthrot2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.Nz = math.Float32frombits(binary.LittleEndian.Uint32(payload[48:]))
	m.Airspeedref = math.Float32frombits(binary.LittleEndian.Uint32(payload[52:]))
	m.Yawangle = math.Float32frombits(binary.LittleEndian.Uint32(payload[56:]))
	m.Yawangleref = math.Float32frombits(binary.LittleEndian.Uint32(payload[60:]))
	m.Rollangle = math.Float32frombits(binary.LittleEndian.Uint32(payload[64:]))
	m.Rollangleref = math.Float32frombits(binary.LittleEndian.Uint32(payload[68:]))
	m.P = math.Float32frombits(binary.LittleEndian.Uint32(payload[72:]))
	m.Pref = math.Float32frombits(binary.LittleEndian.Uint32(payload[76:]))
	m.R = math.Float32frombits(binary.LittleEndian.Uint32(payload[80:]))
	m.Rref = math.Float32frombits(binary.LittleEndian.Uint32(payload[84:]))
	m.Uail = math.Float32frombits(binary.LittleEndian.Uint32(payload[88:]))
	m.Urud = math.Float32frombits(binary.LittleEndian.Uint32(payload[92:]))
	m.AslctrlMode = uint8(payload[96])
	m.Spoilersengaged = uint8(payload[97])
	return nil
}

// AslctrlDebug struct (generated typeinfo)
// ASL-fixed-wing controller debug data
type AslctrlDebug struct {
	I321 uint32  `json:"I321" ` //  Debug data
	F1   float32 `json:"F1" `   //  Debug data
	F2   float32 `json:"F2" `   //  Debug data
	F3   float32 `json:"F3" `   //  Debug data
	F4   float32 `json:"F4" `   //  Debug data
	F5   float32 `json:"F5" `   //  Debug data
	F6   float32 `json:"F6" `   //  Debug data
	F7   float32 `json:"F7" `   //  Debug data
	F8   float32 `json:"F8" `   //  Debug data
	I81  uint8   `json:"I81" `  //  Debug data
	I82  uint8   `json:"I82" `  //  Debug data
}

// MsgID (generated function)
func (m *AslctrlDebug) MsgID() message.MessageID {
	return MSG_ID_ASLCTRL_DEBUG
}

// String (generated function)
func (m *AslctrlDebug) String() string {
	return fmt.Sprintf(
		"&ASLUAV.AslctrlDebug{ I321: %+v, F1: %+v, F2: %+v, F3: %+v, F4: %+v, F5: %+v, F6: %+v, F7: %+v, F8: %+v, I81: %+v, I82: %+v }",
		m.I321,
		m.F1,
		m.F2,
		m.F3,
		m.F4,
		m.F5,
		m.F6,
		m.F7,
		m.F8,
		m.I81,
		m.I82,
	)
}

const (
	ASLCTRL_DEBUG_FIELD_I32_1 = "AslctrlDebug.I321"
	ASLCTRL_DEBUG_FIELD_F_1   = "AslctrlDebug.F1"
	ASLCTRL_DEBUG_FIELD_F_2   = "AslctrlDebug.F2"
	ASLCTRL_DEBUG_FIELD_F_3   = "AslctrlDebug.F3"
	ASLCTRL_DEBUG_FIELD_F_4   = "AslctrlDebug.F4"
	ASLCTRL_DEBUG_FIELD_F_5   = "AslctrlDebug.F5"
	ASLCTRL_DEBUG_FIELD_F_6   = "AslctrlDebug.F6"
	ASLCTRL_DEBUG_FIELD_F_7   = "AslctrlDebug.F7"
	ASLCTRL_DEBUG_FIELD_F_8   = "AslctrlDebug.F8"
	ASLCTRL_DEBUG_FIELD_I8_1  = "AslctrlDebug.I81"
	ASLCTRL_DEBUG_FIELD_I8_2  = "AslctrlDebug.I82"
)

// ToMap (generated function)
func (m *AslctrlDebug) Dict() map[string]interface{} {
	return map[string]interface{}{
		ASLCTRL_DEBUG_FIELD_I32_1: m.I321,
		ASLCTRL_DEBUG_FIELD_F_1:   m.F1,
		ASLCTRL_DEBUG_FIELD_F_2:   m.F2,
		ASLCTRL_DEBUG_FIELD_F_3:   m.F3,
		ASLCTRL_DEBUG_FIELD_F_4:   m.F4,
		ASLCTRL_DEBUG_FIELD_F_5:   m.F5,
		ASLCTRL_DEBUG_FIELD_F_6:   m.F6,
		ASLCTRL_DEBUG_FIELD_F_7:   m.F7,
		ASLCTRL_DEBUG_FIELD_F_8:   m.F8,
		ASLCTRL_DEBUG_FIELD_I8_1:  m.I81,
		ASLCTRL_DEBUG_FIELD_I8_2:  m.I82,
	}
}

// Marshal (generated function)
func (m *AslctrlDebug) Marshal() ([]byte, error) {
	payload := make([]byte, 38)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.I321))
	binary.LittleEndian.PutUint32(payload[4:], math.Float32bits(m.F1))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.F2))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.F3))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.F4))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.F5))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.F6))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.F7))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.F8))
	payload[36] = byte(m.I81)
	payload[37] = byte(m.I82)
	return payload, nil
}

// Unmarshal (generated function)
func (m *AslctrlDebug) Unmarshal(data []byte) error {
	payload := make([]byte, 38)
	copy(payload[0:], data)
	m.I321 = uint32(binary.LittleEndian.Uint32(payload[0:]))
	m.F1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[4:]))
	m.F2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.F3 = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.F4 = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.F5 = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.F6 = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.F7 = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.F8 = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.I81 = uint8(payload[36])
	m.I82 = uint8(payload[37])
	return nil
}

// AsluavStatus struct (generated typeinfo)
// Extended state information for ASLUAVs
type AsluavStatus struct {
	MotorRpm     float32 `json:"MotorRpm" `            //  Motor RPM
	LedStatus    uint8   `json:"LedStatus" `           //  Status of the position-indicator LEDs
	SatcomStatus uint8   `json:"SatcomStatus" `        //  Status of the IRIDIUM satellite communication system
	ServoStatus  []uint8 `json:"ServoStatus" len:"8" ` //  Status vector for up to 8 servos
}

// MsgID (generated function)
func (m *AsluavStatus) MsgID() message.MessageID {
	return MSG_ID_ASLUAV_STATUS
}

// String (generated function)
func (m *AsluavStatus) String() string {
	return fmt.Sprintf(
		"&ASLUAV.AsluavStatus{ MotorRpm: %+v, LedStatus: %+v, SatcomStatus: %+v, ServoStatus: %0X (\"%s\") }",
		m.MotorRpm,
		m.LedStatus,
		m.SatcomStatus,
		m.ServoStatus, string(m.ServoStatus[:]),
	)
}

const (
	ASLUAV_STATUS_FIELD_MOTOR_RPM     = "AsluavStatus.MotorRpm"
	ASLUAV_STATUS_FIELD_LED_STATUS    = "AsluavStatus.LedStatus"
	ASLUAV_STATUS_FIELD_SATCOM_STATUS = "AsluavStatus.SatcomStatus"
	ASLUAV_STATUS_FIELD_SERVO_STATUS  = "AsluavStatus.ServoStatus"
)

// ToMap (generated function)
func (m *AsluavStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		ASLUAV_STATUS_FIELD_MOTOR_RPM:     m.MotorRpm,
		ASLUAV_STATUS_FIELD_LED_STATUS:    m.LedStatus,
		ASLUAV_STATUS_FIELD_SATCOM_STATUS: m.SatcomStatus,
		ASLUAV_STATUS_FIELD_SERVO_STATUS:  m.ServoStatus,
	}
}

// Marshal (generated function)
func (m *AsluavStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 14)
	binary.LittleEndian.PutUint32(payload[0:], math.Float32bits(m.MotorRpm))
	payload[4] = byte(m.LedStatus)
	payload[5] = byte(m.SatcomStatus)
	copy(payload[6:], m.ServoStatus)
	return payload, nil
}

// Unmarshal (generated function)
func (m *AsluavStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 14)
	copy(payload[0:], data)
	m.MotorRpm = math.Float32frombits(binary.LittleEndian.Uint32(payload[0:]))
	m.LedStatus = uint8(payload[4])
	m.SatcomStatus = uint8(payload[5])
	copy(m.ServoStatus[:], payload[6:14])
	return nil
}

// EkfExt struct (generated typeinfo)
// Extended EKF state estimates for ASLUAVs
type EkfExt struct {
	Timestamp uint64  `json:"Timestamp" ` // [ us ]  Time since system start
	Windspeed float32 `json:"Windspeed" ` // [ m/s ]  Magnitude of wind velocity (in lateral inertial plane)
	Winddir   float32 `json:"Winddir" `   // [ rad ]  Wind heading angle from North
	Windz     float32 `json:"Windz" `     // [ m/s ]  Z (Down) component of inertial wind velocity
	Airspeed  float32 `json:"Airspeed" `  // [ m/s ]  Magnitude of air velocity
	Beta      float32 `json:"Beta" `      // [ rad ]  Sideslip angle
	Alpha     float32 `json:"Alpha" `     // [ rad ]  Angle of attack
}

// MsgID (generated function)
func (m *EkfExt) MsgID() message.MessageID {
	return MSG_ID_EKF_EXT
}

// String (generated function)
func (m *EkfExt) String() string {
	return fmt.Sprintf(
		"&ASLUAV.EkfExt{ Timestamp: %+v, Windspeed: %+v, Winddir: %+v, Windz: %+v, Airspeed: %+v, Beta: %+v, Alpha: %+v }",
		m.Timestamp,
		m.Windspeed,
		m.Winddir,
		m.Windz,
		m.Airspeed,
		m.Beta,
		m.Alpha,
	)
}

const (
	EKF_EXT_FIELD_TIMESTAMP = "EkfExt.Timestamp"
	EKF_EXT_FIELD_WINDSPEED = "EkfExt.Windspeed"
	EKF_EXT_FIELD_WINDDIR   = "EkfExt.Winddir"
	EKF_EXT_FIELD_WINDZ     = "EkfExt.Windz"
	EKF_EXT_FIELD_AIRSPEED  = "EkfExt.Airspeed"
	EKF_EXT_FIELD_BETA      = "EkfExt.Beta"
	EKF_EXT_FIELD_ALPHA     = "EkfExt.Alpha"
)

// ToMap (generated function)
func (m *EkfExt) Dict() map[string]interface{} {
	return map[string]interface{}{
		EKF_EXT_FIELD_TIMESTAMP: m.Timestamp,
		EKF_EXT_FIELD_WINDSPEED: m.Windspeed,
		EKF_EXT_FIELD_WINDDIR:   m.Winddir,
		EKF_EXT_FIELD_WINDZ:     m.Windz,
		EKF_EXT_FIELD_AIRSPEED:  m.Airspeed,
		EKF_EXT_FIELD_BETA:      m.Beta,
		EKF_EXT_FIELD_ALPHA:     m.Alpha,
	}
}

// Marshal (generated function)
func (m *EkfExt) Marshal() ([]byte, error) {
	payload := make([]byte, 32)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Windspeed))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Winddir))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Windz))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Airspeed))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Beta))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Alpha))
	return payload, nil
}

// Unmarshal (generated function)
func (m *EkfExt) Unmarshal(data []byte) error {
	payload := make([]byte, 32)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Windspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Winddir = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Windz = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Airspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Beta = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Alpha = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	return nil
}

// AslObctrl struct (generated typeinfo)
// Off-board controls/commands for ASLUAVs
type AslObctrl struct {
	Timestamp    uint64  `json:"Timestamp" `    // [ us ]  Time since system start
	Uelev        float32 `json:"Uelev" `        //  Elevator command [~]
	Uthrot       float32 `json:"Uthrot" `       //  Throttle command [~]
	Uthrot2      float32 `json:"Uthrot2" `      //  Throttle 2 command [~]
	Uaill        float32 `json:"Uaill" `        //  Left aileron command [~]
	Uailr        float32 `json:"Uailr" `        //  Right aileron command [~]
	Urud         float32 `json:"Urud" `         //  Rudder command [~]
	ObctrlStatus uint8   `json:"ObctrlStatus" ` //  Off-board computer status
}

// MsgID (generated function)
func (m *AslObctrl) MsgID() message.MessageID {
	return MSG_ID_ASL_OBCTRL
}

// String (generated function)
func (m *AslObctrl) String() string {
	return fmt.Sprintf(
		"&ASLUAV.AslObctrl{ Timestamp: %+v, Uelev: %+v, Uthrot: %+v, Uthrot2: %+v, Uaill: %+v, Uailr: %+v, Urud: %+v, ObctrlStatus: %+v }",
		m.Timestamp,
		m.Uelev,
		m.Uthrot,
		m.Uthrot2,
		m.Uaill,
		m.Uailr,
		m.Urud,
		m.ObctrlStatus,
	)
}

const (
	ASL_OBCTRL_FIELD_TIMESTAMP     = "AslObctrl.Timestamp"
	ASL_OBCTRL_FIELD_UELEV         = "AslObctrl.Uelev"
	ASL_OBCTRL_FIELD_UTHROT        = "AslObctrl.Uthrot"
	ASL_OBCTRL_FIELD_UTHROT2       = "AslObctrl.Uthrot2"
	ASL_OBCTRL_FIELD_UAILL         = "AslObctrl.Uaill"
	ASL_OBCTRL_FIELD_UAILR         = "AslObctrl.Uailr"
	ASL_OBCTRL_FIELD_URUD          = "AslObctrl.Urud"
	ASL_OBCTRL_FIELD_OBCTRL_STATUS = "AslObctrl.ObctrlStatus"
)

// ToMap (generated function)
func (m *AslObctrl) Dict() map[string]interface{} {
	return map[string]interface{}{
		ASL_OBCTRL_FIELD_TIMESTAMP:     m.Timestamp,
		ASL_OBCTRL_FIELD_UELEV:         m.Uelev,
		ASL_OBCTRL_FIELD_UTHROT:        m.Uthrot,
		ASL_OBCTRL_FIELD_UTHROT2:       m.Uthrot2,
		ASL_OBCTRL_FIELD_UAILL:         m.Uaill,
		ASL_OBCTRL_FIELD_UAILR:         m.Uailr,
		ASL_OBCTRL_FIELD_URUD:          m.Urud,
		ASL_OBCTRL_FIELD_OBCTRL_STATUS: m.ObctrlStatus,
	}
}

// Marshal (generated function)
func (m *AslObctrl) Marshal() ([]byte, error) {
	payload := make([]byte, 33)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Uelev))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Uthrot))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Uthrot2))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Uaill))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Uailr))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Urud))
	payload[32] = byte(m.ObctrlStatus)
	return payload, nil
}

// Unmarshal (generated function)
func (m *AslObctrl) Unmarshal(data []byte) error {
	payload := make([]byte, 33)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Uelev = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Uthrot = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.Uthrot2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Uaill = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Uailr = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Urud = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.ObctrlStatus = uint8(payload[32])
	return nil
}

// SensAtmos struct (generated typeinfo)
// Atmospheric sensors (temperature, humidity, ...)
type SensAtmos struct {
	Timestamp   uint64  `json:"Timestamp" `   // [ us ] Time since system boot
	Tempambient float32 `json:"Tempambient" ` // [ degC ]  Ambient temperature
	Humidity    float32 `json:"Humidity" `    // [ % ]  Relative humidity
}

// MsgID (generated function)
func (m *SensAtmos) MsgID() message.MessageID {
	return MSG_ID_SENS_ATMOS
}

// String (generated function)
func (m *SensAtmos) String() string {
	return fmt.Sprintf(
		"&ASLUAV.SensAtmos{ Timestamp: %+v, Tempambient: %+v, Humidity: %+v }",
		m.Timestamp,
		m.Tempambient,
		m.Humidity,
	)
}

const (
	SENS_ATMOS_FIELD_TIMESTAMP   = "SensAtmos.Timestamp"
	SENS_ATMOS_FIELD_TEMPAMBIENT = "SensAtmos.Tempambient"
	SENS_ATMOS_FIELD_HUMIDITY    = "SensAtmos.Humidity"
)

// ToMap (generated function)
func (m *SensAtmos) Dict() map[string]interface{} {
	return map[string]interface{}{
		SENS_ATMOS_FIELD_TIMESTAMP:   m.Timestamp,
		SENS_ATMOS_FIELD_TEMPAMBIENT: m.Tempambient,
		SENS_ATMOS_FIELD_HUMIDITY:    m.Humidity,
	}
}

// Marshal (generated function)
func (m *SensAtmos) Marshal() ([]byte, error) {
	payload := make([]byte, 16)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Tempambient))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Humidity))
	return payload, nil
}

// Unmarshal (generated function)
func (m *SensAtmos) Unmarshal(data []byte) error {
	payload := make([]byte, 16)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Tempambient = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Humidity = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	return nil
}

// SensBatmon struct (generated typeinfo)
// Battery pack monitoring data for Li-Ion batteries
type SensBatmon struct {
	BatmonTimestamp uint64  `json:"BatmonTimestamp" ` // [ us ] Time since system start
	Temperature     float32 `json:"Temperature" `     // [ degC ] Battery pack temperature
	Safetystatus    uint32  `json:"Safetystatus" `    // Battery monitor safetystatus report bits in Hex
	Operationstatus uint32  `json:"Operationstatus" ` // Battery monitor operation status report bits in Hex
	Voltage         uint16  `json:"Voltage" `         // [ mV ] Battery pack voltage
	Current         int16   `json:"Current" `         // [ mA ] Battery pack current
	Batterystatus   uint16  `json:"Batterystatus" `   // Battery monitor status report bits in Hex
	Serialnumber    uint16  `json:"Serialnumber" `    // Battery monitor serial number in Hex
	Cellvoltage1    uint16  `json:"Cellvoltage1" `    // [ mV ] Battery pack cell 1 voltage
	Cellvoltage2    uint16  `json:"Cellvoltage2" `    // [ mV ] Battery pack cell 2 voltage
	Cellvoltage3    uint16  `json:"Cellvoltage3" `    // [ mV ] Battery pack cell 3 voltage
	Cellvoltage4    uint16  `json:"Cellvoltage4" `    // [ mV ] Battery pack cell 4 voltage
	Cellvoltage5    uint16  `json:"Cellvoltage5" `    // [ mV ] Battery pack cell 5 voltage
	Cellvoltage6    uint16  `json:"Cellvoltage6" `    // [ mV ] Battery pack cell 6 voltage
	Soc             uint8   `json:"Soc" `             // Battery pack state-of-charge
}

// MsgID (generated function)
func (m *SensBatmon) MsgID() message.MessageID {
	return MSG_ID_SENS_BATMON
}

// String (generated function)
func (m *SensBatmon) String() string {
	return fmt.Sprintf(
		"&ASLUAV.SensBatmon{ BatmonTimestamp: %+v, Temperature: %+v, Safetystatus: %+v, Operationstatus: %+v, Voltage: %+v, Current: %+v, Batterystatus: %+v, Serialnumber: %+v, Cellvoltage1: %+v, Cellvoltage2: %+v, Cellvoltage3: %+v, Cellvoltage4: %+v, Cellvoltage5: %+v, Cellvoltage6: %+v, Soc: %+v }",
		m.BatmonTimestamp,
		m.Temperature,
		m.Safetystatus,
		m.Operationstatus,
		m.Voltage,
		m.Current,
		m.Batterystatus,
		m.Serialnumber,
		m.Cellvoltage1,
		m.Cellvoltage2,
		m.Cellvoltage3,
		m.Cellvoltage4,
		m.Cellvoltage5,
		m.Cellvoltage6,
		m.Soc,
	)
}

const (
	SENS_BATMON_FIELD_BATMON_TIMESTAMP = "SensBatmon.BatmonTimestamp"
	SENS_BATMON_FIELD_TEMPERATURE      = "SensBatmon.Temperature"
	SENS_BATMON_FIELD_SAFETYSTATUS     = "SensBatmon.Safetystatus"
	SENS_BATMON_FIELD_OPERATIONSTATUS  = "SensBatmon.Operationstatus"
	SENS_BATMON_FIELD_VOLTAGE          = "SensBatmon.Voltage"
	SENS_BATMON_FIELD_CURRENT          = "SensBatmon.Current"
	SENS_BATMON_FIELD_BATTERYSTATUS    = "SensBatmon.Batterystatus"
	SENS_BATMON_FIELD_SERIALNUMBER     = "SensBatmon.Serialnumber"
	SENS_BATMON_FIELD_CELLVOLTAGE1     = "SensBatmon.Cellvoltage1"
	SENS_BATMON_FIELD_CELLVOLTAGE2     = "SensBatmon.Cellvoltage2"
	SENS_BATMON_FIELD_CELLVOLTAGE3     = "SensBatmon.Cellvoltage3"
	SENS_BATMON_FIELD_CELLVOLTAGE4     = "SensBatmon.Cellvoltage4"
	SENS_BATMON_FIELD_CELLVOLTAGE5     = "SensBatmon.Cellvoltage5"
	SENS_BATMON_FIELD_CELLVOLTAGE6     = "SensBatmon.Cellvoltage6"
	SENS_BATMON_FIELD_SOC              = "SensBatmon.Soc"
)

// ToMap (generated function)
func (m *SensBatmon) Dict() map[string]interface{} {
	return map[string]interface{}{
		SENS_BATMON_FIELD_BATMON_TIMESTAMP: m.BatmonTimestamp,
		SENS_BATMON_FIELD_TEMPERATURE:      m.Temperature,
		SENS_BATMON_FIELD_SAFETYSTATUS:     m.Safetystatus,
		SENS_BATMON_FIELD_OPERATIONSTATUS:  m.Operationstatus,
		SENS_BATMON_FIELD_VOLTAGE:          m.Voltage,
		SENS_BATMON_FIELD_CURRENT:          m.Current,
		SENS_BATMON_FIELD_BATTERYSTATUS:    m.Batterystatus,
		SENS_BATMON_FIELD_SERIALNUMBER:     m.Serialnumber,
		SENS_BATMON_FIELD_CELLVOLTAGE1:     m.Cellvoltage1,
		SENS_BATMON_FIELD_CELLVOLTAGE2:     m.Cellvoltage2,
		SENS_BATMON_FIELD_CELLVOLTAGE3:     m.Cellvoltage3,
		SENS_BATMON_FIELD_CELLVOLTAGE4:     m.Cellvoltage4,
		SENS_BATMON_FIELD_CELLVOLTAGE5:     m.Cellvoltage5,
		SENS_BATMON_FIELD_CELLVOLTAGE6:     m.Cellvoltage6,
		SENS_BATMON_FIELD_SOC:              m.Soc,
	}
}

// Marshal (generated function)
func (m *SensBatmon) Marshal() ([]byte, error) {
	payload := make([]byte, 41)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.BatmonTimestamp))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Temperature))
	binary.LittleEndian.PutUint32(payload[12:], uint32(m.Safetystatus))
	binary.LittleEndian.PutUint32(payload[16:], uint32(m.Operationstatus))
	binary.LittleEndian.PutUint16(payload[20:], uint16(m.Voltage))
	binary.LittleEndian.PutUint16(payload[22:], uint16(m.Current))
	binary.LittleEndian.PutUint16(payload[24:], uint16(m.Batterystatus))
	binary.LittleEndian.PutUint16(payload[26:], uint16(m.Serialnumber))
	binary.LittleEndian.PutUint16(payload[28:], uint16(m.Cellvoltage1))
	binary.LittleEndian.PutUint16(payload[30:], uint16(m.Cellvoltage2))
	binary.LittleEndian.PutUint16(payload[32:], uint16(m.Cellvoltage3))
	binary.LittleEndian.PutUint16(payload[34:], uint16(m.Cellvoltage4))
	binary.LittleEndian.PutUint16(payload[36:], uint16(m.Cellvoltage5))
	binary.LittleEndian.PutUint16(payload[38:], uint16(m.Cellvoltage6))
	payload[40] = byte(m.Soc)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SensBatmon) Unmarshal(data []byte) error {
	payload := make([]byte, 41)
	copy(payload[0:], data)
	m.BatmonTimestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Temperature = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Safetystatus = uint32(binary.LittleEndian.Uint32(payload[12:]))
	m.Operationstatus = uint32(binary.LittleEndian.Uint32(payload[16:]))
	m.Voltage = uint16(binary.LittleEndian.Uint16(payload[20:]))
	m.Current = int16(binary.LittleEndian.Uint16(payload[22:]))
	m.Batterystatus = uint16(binary.LittleEndian.Uint16(payload[24:]))
	m.Serialnumber = uint16(binary.LittleEndian.Uint16(payload[26:]))
	m.Cellvoltage1 = uint16(binary.LittleEndian.Uint16(payload[28:]))
	m.Cellvoltage2 = uint16(binary.LittleEndian.Uint16(payload[30:]))
	m.Cellvoltage3 = uint16(binary.LittleEndian.Uint16(payload[32:]))
	m.Cellvoltage4 = uint16(binary.LittleEndian.Uint16(payload[34:]))
	m.Cellvoltage5 = uint16(binary.LittleEndian.Uint16(payload[36:]))
	m.Cellvoltage6 = uint16(binary.LittleEndian.Uint16(payload[38:]))
	m.Soc = uint8(payload[40])
	return nil
}

// FwSoaringData struct (generated typeinfo)
// Fixed-wing soaring (i.e. thermal seeking) data
type FwSoaringData struct {
	Timestamp            uint64  `json:"Timestamp" `            // [ ms ] Timestamp
	Timestampmodechanged uint64  `json:"Timestampmodechanged" ` // [ ms ] Timestamp since last mode change
	Xw                   float32 `json:"Xw" `                   // [ m/s ] Thermal core updraft strength
	Xr                   float32 `json:"Xr" `                   // [ m ] Thermal radius
	Xlat                 float32 `json:"Xlat" `                 // [ deg ] Thermal center latitude
	Xlon                 float32 `json:"Xlon" `                 // [ deg ] Thermal center longitude
	Varw                 float32 `json:"Varw" `                 // Variance W
	Varr                 float32 `json:"Varr" `                 // Variance R
	Varlat               float32 `json:"Varlat" `               // Variance Lat
	Varlon               float32 `json:"Varlon" `               // Variance Lon
	Loiterradius         float32 `json:"Loiterradius" `         // [ m ] Suggested loiter radius
	Loiterdirection      float32 `json:"Loiterdirection" `      // Suggested loiter direction
	Disttosoarpoint      float32 `json:"Disttosoarpoint" `      // [ m ] Distance to soar point
	Vsinkexp             float32 `json:"Vsinkexp" `             // [ m/s ] Expected sink rate at current airspeed, roll and throttle
	Z1Localupdraftspeed  float32 `json:"Z1Localupdraftspeed" `  // [ m/s ] Measurement / updraft speed at current/local airplane position
	Z2Deltaroll          float32 `json:"Z2Deltaroll" `          // [ deg ] Measurement / roll angle tracking error
	Z1Exp                float32 `json:"Z1Exp" `                // Expected measurement 1
	Z2Exp                float32 `json:"Z2Exp" `                // Expected measurement 2
	Thermalgsnorth       float32 `json:"Thermalgsnorth" `       // [ m/s ] Thermal drift (from estimator prediction step only)
	Thermalgseast        float32 `json:"Thermalgseast" `        // [ m/s ] Thermal drift (from estimator prediction step only)
	TseDot               float32 `json:"TseDot" `               // [ m/s ]  Total specific energy change (filtered)
	Debugvar1            float32 `json:"Debugvar1" `            //  Debug variable 1
	Debugvar2            float32 `json:"Debugvar2" `            //  Debug variable 2
	Controlmode          uint8   `json:"Controlmode" `          // Control Mode [-]
	Valid                uint8   `json:"Valid" `                // Data valid [-]
}

// MsgID (generated function)
func (m *FwSoaringData) MsgID() message.MessageID {
	return MSG_ID_FW_SOARING_DATA
}

// String (generated function)
func (m *FwSoaringData) String() string {
	return fmt.Sprintf(
		"&ASLUAV.FwSoaringData{ Timestamp: %+v, Timestampmodechanged: %+v, Xw: %+v, Xr: %+v, Xlat: %+v, Xlon: %+v, Varw: %+v, Varr: %+v, Varlat: %+v, Varlon: %+v, Loiterradius: %+v, Loiterdirection: %+v, Disttosoarpoint: %+v, Vsinkexp: %+v, Z1Localupdraftspeed: %+v, Z2Deltaroll: %+v, Z1Exp: %+v, Z2Exp: %+v, Thermalgsnorth: %+v, Thermalgseast: %+v, TseDot: %+v, Debugvar1: %+v, Debugvar2: %+v, Controlmode: %+v, Valid: %+v }",
		m.Timestamp,
		m.Timestampmodechanged,
		m.Xw,
		m.Xr,
		m.Xlat,
		m.Xlon,
		m.Varw,
		m.Varr,
		m.Varlat,
		m.Varlon,
		m.Loiterradius,
		m.Loiterdirection,
		m.Disttosoarpoint,
		m.Vsinkexp,
		m.Z1Localupdraftspeed,
		m.Z2Deltaroll,
		m.Z1Exp,
		m.Z2Exp,
		m.Thermalgsnorth,
		m.Thermalgseast,
		m.TseDot,
		m.Debugvar1,
		m.Debugvar2,
		m.Controlmode,
		m.Valid,
	)
}

const (
	FW_SOARING_DATA_FIELD_TIMESTAMP            = "FwSoaringData.Timestamp"
	FW_SOARING_DATA_FIELD_TIMESTAMPMODECHANGED = "FwSoaringData.Timestampmodechanged"
	FW_SOARING_DATA_FIELD_XW                   = "FwSoaringData.Xw"
	FW_SOARING_DATA_FIELD_XR                   = "FwSoaringData.Xr"
	FW_SOARING_DATA_FIELD_XLAT                 = "FwSoaringData.Xlat"
	FW_SOARING_DATA_FIELD_XLON                 = "FwSoaringData.Xlon"
	FW_SOARING_DATA_FIELD_VARW                 = "FwSoaringData.Varw"
	FW_SOARING_DATA_FIELD_VARR                 = "FwSoaringData.Varr"
	FW_SOARING_DATA_FIELD_VARLAT               = "FwSoaringData.Varlat"
	FW_SOARING_DATA_FIELD_VARLON               = "FwSoaringData.Varlon"
	FW_SOARING_DATA_FIELD_LOITERRADIUS         = "FwSoaringData.Loiterradius"
	FW_SOARING_DATA_FIELD_LOITERDIRECTION      = "FwSoaringData.Loiterdirection"
	FW_SOARING_DATA_FIELD_DISTTOSOARPOINT      = "FwSoaringData.Disttosoarpoint"
	FW_SOARING_DATA_FIELD_VSINKEXP             = "FwSoaringData.Vsinkexp"
	FW_SOARING_DATA_FIELD_Z1_LOCALUPDRAFTSPEED = "FwSoaringData.Z1Localupdraftspeed"
	FW_SOARING_DATA_FIELD_Z2_DELTAROLL         = "FwSoaringData.Z2Deltaroll"
	FW_SOARING_DATA_FIELD_Z1_EXP               = "FwSoaringData.Z1Exp"
	FW_SOARING_DATA_FIELD_Z2_EXP               = "FwSoaringData.Z2Exp"
	FW_SOARING_DATA_FIELD_THERMALGSNORTH       = "FwSoaringData.Thermalgsnorth"
	FW_SOARING_DATA_FIELD_THERMALGSEAST        = "FwSoaringData.Thermalgseast"
	FW_SOARING_DATA_FIELD_TSE_DOT              = "FwSoaringData.TseDot"
	FW_SOARING_DATA_FIELD_DEBUGVAR1            = "FwSoaringData.Debugvar1"
	FW_SOARING_DATA_FIELD_DEBUGVAR2            = "FwSoaringData.Debugvar2"
	FW_SOARING_DATA_FIELD_CONTROLMODE          = "FwSoaringData.Controlmode"
	FW_SOARING_DATA_FIELD_VALID                = "FwSoaringData.Valid"
)

// ToMap (generated function)
func (m *FwSoaringData) Dict() map[string]interface{} {
	return map[string]interface{}{
		FW_SOARING_DATA_FIELD_TIMESTAMP:            m.Timestamp,
		FW_SOARING_DATA_FIELD_TIMESTAMPMODECHANGED: m.Timestampmodechanged,
		FW_SOARING_DATA_FIELD_XW:                   m.Xw,
		FW_SOARING_DATA_FIELD_XR:                   m.Xr,
		FW_SOARING_DATA_FIELD_XLAT:                 m.Xlat,
		FW_SOARING_DATA_FIELD_XLON:                 m.Xlon,
		FW_SOARING_DATA_FIELD_VARW:                 m.Varw,
		FW_SOARING_DATA_FIELD_VARR:                 m.Varr,
		FW_SOARING_DATA_FIELD_VARLAT:               m.Varlat,
		FW_SOARING_DATA_FIELD_VARLON:               m.Varlon,
		FW_SOARING_DATA_FIELD_LOITERRADIUS:         m.Loiterradius,
		FW_SOARING_DATA_FIELD_LOITERDIRECTION:      m.Loiterdirection,
		FW_SOARING_DATA_FIELD_DISTTOSOARPOINT:      m.Disttosoarpoint,
		FW_SOARING_DATA_FIELD_VSINKEXP:             m.Vsinkexp,
		FW_SOARING_DATA_FIELD_Z1_LOCALUPDRAFTSPEED: m.Z1Localupdraftspeed,
		FW_SOARING_DATA_FIELD_Z2_DELTAROLL:         m.Z2Deltaroll,
		FW_SOARING_DATA_FIELD_Z1_EXP:               m.Z1Exp,
		FW_SOARING_DATA_FIELD_Z2_EXP:               m.Z2Exp,
		FW_SOARING_DATA_FIELD_THERMALGSNORTH:       m.Thermalgsnorth,
		FW_SOARING_DATA_FIELD_THERMALGSEAST:        m.Thermalgseast,
		FW_SOARING_DATA_FIELD_TSE_DOT:              m.TseDot,
		FW_SOARING_DATA_FIELD_DEBUGVAR1:            m.Debugvar1,
		FW_SOARING_DATA_FIELD_DEBUGVAR2:            m.Debugvar2,
		FW_SOARING_DATA_FIELD_CONTROLMODE:          m.Controlmode,
		FW_SOARING_DATA_FIELD_VALID:                m.Valid,
	}
}

// Marshal (generated function)
func (m *FwSoaringData) Marshal() ([]byte, error) {
	payload := make([]byte, 102)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint64(payload[8:], uint64(m.Timestampmodechanged))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.Xw))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.Xr))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.Xlat))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.Xlon))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.Varw))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.Varr))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.Varlat))
	binary.LittleEndian.PutUint32(payload[44:], math.Float32bits(m.Varlon))
	binary.LittleEndian.PutUint32(payload[48:], math.Float32bits(m.Loiterradius))
	binary.LittleEndian.PutUint32(payload[52:], math.Float32bits(m.Loiterdirection))
	binary.LittleEndian.PutUint32(payload[56:], math.Float32bits(m.Disttosoarpoint))
	binary.LittleEndian.PutUint32(payload[60:], math.Float32bits(m.Vsinkexp))
	binary.LittleEndian.PutUint32(payload[64:], math.Float32bits(m.Z1Localupdraftspeed))
	binary.LittleEndian.PutUint32(payload[68:], math.Float32bits(m.Z2Deltaroll))
	binary.LittleEndian.PutUint32(payload[72:], math.Float32bits(m.Z1Exp))
	binary.LittleEndian.PutUint32(payload[76:], math.Float32bits(m.Z2Exp))
	binary.LittleEndian.PutUint32(payload[80:], math.Float32bits(m.Thermalgsnorth))
	binary.LittleEndian.PutUint32(payload[84:], math.Float32bits(m.Thermalgseast))
	binary.LittleEndian.PutUint32(payload[88:], math.Float32bits(m.TseDot))
	binary.LittleEndian.PutUint32(payload[92:], math.Float32bits(m.Debugvar1))
	binary.LittleEndian.PutUint32(payload[96:], math.Float32bits(m.Debugvar2))
	payload[100] = byte(m.Controlmode)
	payload[101] = byte(m.Valid)
	return payload, nil
}

// Unmarshal (generated function)
func (m *FwSoaringData) Unmarshal(data []byte) error {
	payload := make([]byte, 102)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Timestampmodechanged = uint64(binary.LittleEndian.Uint64(payload[8:]))
	m.Xw = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.Xr = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.Xlat = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.Xlon = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.Varw = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.Varr = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.Varlat = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.Varlon = math.Float32frombits(binary.LittleEndian.Uint32(payload[44:]))
	m.Loiterradius = math.Float32frombits(binary.LittleEndian.Uint32(payload[48:]))
	m.Loiterdirection = math.Float32frombits(binary.LittleEndian.Uint32(payload[52:]))
	m.Disttosoarpoint = math.Float32frombits(binary.LittleEndian.Uint32(payload[56:]))
	m.Vsinkexp = math.Float32frombits(binary.LittleEndian.Uint32(payload[60:]))
	m.Z1Localupdraftspeed = math.Float32frombits(binary.LittleEndian.Uint32(payload[64:]))
	m.Z2Deltaroll = math.Float32frombits(binary.LittleEndian.Uint32(payload[68:]))
	m.Z1Exp = math.Float32frombits(binary.LittleEndian.Uint32(payload[72:]))
	m.Z2Exp = math.Float32frombits(binary.LittleEndian.Uint32(payload[76:]))
	m.Thermalgsnorth = math.Float32frombits(binary.LittleEndian.Uint32(payload[80:]))
	m.Thermalgseast = math.Float32frombits(binary.LittleEndian.Uint32(payload[84:]))
	m.TseDot = math.Float32frombits(binary.LittleEndian.Uint32(payload[88:]))
	m.Debugvar1 = math.Float32frombits(binary.LittleEndian.Uint32(payload[92:]))
	m.Debugvar2 = math.Float32frombits(binary.LittleEndian.Uint32(payload[96:]))
	m.Controlmode = uint8(payload[100])
	m.Valid = uint8(payload[101])
	return nil
}

// SensorpodStatus struct (generated typeinfo)
// Monitoring of sensorpod status
type SensorpodStatus struct {
	Timestamp           uint64 `json:"Timestamp" `           // [ ms ] Timestamp in linuxtime (since 1.1.1970)
	FreeSpace           uint16 `json:"FreeSpace" `           // Free space available in recordings directory in [Gb] * 1e2
	VisensorRate1       uint8  `json:"VisensorRate1" `       // Rate of ROS topic 1
	VisensorRate2       uint8  `json:"VisensorRate2" `       // Rate of ROS topic 2
	VisensorRate3       uint8  `json:"VisensorRate3" `       // Rate of ROS topic 3
	VisensorRate4       uint8  `json:"VisensorRate4" `       // Rate of ROS topic 4
	RecordingNodesCount uint8  `json:"RecordingNodesCount" ` // Number of recording nodes
	CPUTemp             uint8  `json:"CPUTemp" `             // [ degC ] Temperature of sensorpod CPU in
}

// MsgID (generated function)
func (m *SensorpodStatus) MsgID() message.MessageID {
	return MSG_ID_SENSORPOD_STATUS
}

// String (generated function)
func (m *SensorpodStatus) String() string {
	return fmt.Sprintf(
		"&ASLUAV.SensorpodStatus{ Timestamp: %+v, FreeSpace: %+v, VisensorRate1: %+v, VisensorRate2: %+v, VisensorRate3: %+v, VisensorRate4: %+v, RecordingNodesCount: %+v, CPUTemp: %+v }",
		m.Timestamp,
		m.FreeSpace,
		m.VisensorRate1,
		m.VisensorRate2,
		m.VisensorRate3,
		m.VisensorRate4,
		m.RecordingNodesCount,
		m.CPUTemp,
	)
}

const (
	SENSORPOD_STATUS_FIELD_TIMESTAMP             = "SensorpodStatus.Timestamp"
	SENSORPOD_STATUS_FIELD_FREE_SPACE            = "SensorpodStatus.FreeSpace"
	SENSORPOD_STATUS_FIELD_VISENSOR_RATE_1       = "SensorpodStatus.VisensorRate1"
	SENSORPOD_STATUS_FIELD_VISENSOR_RATE_2       = "SensorpodStatus.VisensorRate2"
	SENSORPOD_STATUS_FIELD_VISENSOR_RATE_3       = "SensorpodStatus.VisensorRate3"
	SENSORPOD_STATUS_FIELD_VISENSOR_RATE_4       = "SensorpodStatus.VisensorRate4"
	SENSORPOD_STATUS_FIELD_RECORDING_NODES_COUNT = "SensorpodStatus.RecordingNodesCount"
	SENSORPOD_STATUS_FIELD_CPU_TEMP              = "SensorpodStatus.CPUTemp"
)

// ToMap (generated function)
func (m *SensorpodStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		SENSORPOD_STATUS_FIELD_TIMESTAMP:             m.Timestamp,
		SENSORPOD_STATUS_FIELD_FREE_SPACE:            m.FreeSpace,
		SENSORPOD_STATUS_FIELD_VISENSOR_RATE_1:       m.VisensorRate1,
		SENSORPOD_STATUS_FIELD_VISENSOR_RATE_2:       m.VisensorRate2,
		SENSORPOD_STATUS_FIELD_VISENSOR_RATE_3:       m.VisensorRate3,
		SENSORPOD_STATUS_FIELD_VISENSOR_RATE_4:       m.VisensorRate4,
		SENSORPOD_STATUS_FIELD_RECORDING_NODES_COUNT: m.RecordingNodesCount,
		SENSORPOD_STATUS_FIELD_CPU_TEMP:              m.CPUTemp,
	}
}

// Marshal (generated function)
func (m *SensorpodStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 16)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint16(payload[8:], uint16(m.FreeSpace))
	payload[10] = byte(m.VisensorRate1)
	payload[11] = byte(m.VisensorRate2)
	payload[12] = byte(m.VisensorRate3)
	payload[13] = byte(m.VisensorRate4)
	payload[14] = byte(m.RecordingNodesCount)
	payload[15] = byte(m.CPUTemp)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SensorpodStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 16)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.FreeSpace = uint16(binary.LittleEndian.Uint16(payload[8:]))
	m.VisensorRate1 = uint8(payload[10])
	m.VisensorRate2 = uint8(payload[11])
	m.VisensorRate3 = uint8(payload[12])
	m.VisensorRate4 = uint8(payload[13])
	m.RecordingNodesCount = uint8(payload[14])
	m.CPUTemp = uint8(payload[15])
	return nil
}

// SensPowerBoard struct (generated typeinfo)
// Monitoring of power board status
type SensPowerBoard struct {
	Timestamp         uint64  `json:"Timestamp" `         // [ us ] Timestamp
	PwrBrdSystemVolt  float32 `json:"PwrBrdSystemVolt" `  // [ V ] Power board system voltage
	PwrBrdServoVolt   float32 `json:"PwrBrdServoVolt" `   // [ V ] Power board servo voltage
	PwrBrdDigitalVolt float32 `json:"PwrBrdDigitalVolt" ` // [ V ] Power board digital voltage
	PwrBrdMotLAmp     float32 `json:"PwrBrdMotLAmp" `     // [ A ] Power board left motor current sensor
	PwrBrdMotRAmp     float32 `json:"PwrBrdMotRAmp" `     // [ A ] Power board right motor current sensor
	PwrBrdAnalogAmp   float32 `json:"PwrBrdAnalogAmp" `   // [ A ] Power board analog current sensor
	PwrBrdDigitalAmp  float32 `json:"PwrBrdDigitalAmp" `  // [ A ] Power board digital current sensor
	PwrBrdExtAmp      float32 `json:"PwrBrdExtAmp" `      // [ A ] Power board extension current sensor
	PwrBrdAuxAmp      float32 `json:"PwrBrdAuxAmp" `      // [ A ] Power board aux current sensor
	PwrBrdStatus      uint8   `json:"PwrBrdStatus" `      // Power board status register
	PwrBrdLedStatus   uint8   `json:"PwrBrdLedStatus" `   // Power board leds status
}

// MsgID (generated function)
func (m *SensPowerBoard) MsgID() message.MessageID {
	return MSG_ID_SENS_POWER_BOARD
}

// String (generated function)
func (m *SensPowerBoard) String() string {
	return fmt.Sprintf(
		"&ASLUAV.SensPowerBoard{ Timestamp: %+v, PwrBrdSystemVolt: %+v, PwrBrdServoVolt: %+v, PwrBrdDigitalVolt: %+v, PwrBrdMotLAmp: %+v, PwrBrdMotRAmp: %+v, PwrBrdAnalogAmp: %+v, PwrBrdDigitalAmp: %+v, PwrBrdExtAmp: %+v, PwrBrdAuxAmp: %+v, PwrBrdStatus: %+v, PwrBrdLedStatus: %+v }",
		m.Timestamp,
		m.PwrBrdSystemVolt,
		m.PwrBrdServoVolt,
		m.PwrBrdDigitalVolt,
		m.PwrBrdMotLAmp,
		m.PwrBrdMotRAmp,
		m.PwrBrdAnalogAmp,
		m.PwrBrdDigitalAmp,
		m.PwrBrdExtAmp,
		m.PwrBrdAuxAmp,
		m.PwrBrdStatus,
		m.PwrBrdLedStatus,
	)
}

const (
	SENS_POWER_BOARD_FIELD_TIMESTAMP            = "SensPowerBoard.Timestamp"
	SENS_POWER_BOARD_FIELD_PWR_BRD_SYSTEM_VOLT  = "SensPowerBoard.PwrBrdSystemVolt"
	SENS_POWER_BOARD_FIELD_PWR_BRD_SERVO_VOLT   = "SensPowerBoard.PwrBrdServoVolt"
	SENS_POWER_BOARD_FIELD_PWR_BRD_DIGITAL_VOLT = "SensPowerBoard.PwrBrdDigitalVolt"
	SENS_POWER_BOARD_FIELD_PWR_BRD_MOT_L_AMP    = "SensPowerBoard.PwrBrdMotLAmp"
	SENS_POWER_BOARD_FIELD_PWR_BRD_MOT_R_AMP    = "SensPowerBoard.PwrBrdMotRAmp"
	SENS_POWER_BOARD_FIELD_PWR_BRD_ANALOG_AMP   = "SensPowerBoard.PwrBrdAnalogAmp"
	SENS_POWER_BOARD_FIELD_PWR_BRD_DIGITAL_AMP  = "SensPowerBoard.PwrBrdDigitalAmp"
	SENS_POWER_BOARD_FIELD_PWR_BRD_EXT_AMP      = "SensPowerBoard.PwrBrdExtAmp"
	SENS_POWER_BOARD_FIELD_PWR_BRD_AUX_AMP      = "SensPowerBoard.PwrBrdAuxAmp"
	SENS_POWER_BOARD_FIELD_PWR_BRD_STATUS       = "SensPowerBoard.PwrBrdStatus"
	SENS_POWER_BOARD_FIELD_PWR_BRD_LED_STATUS   = "SensPowerBoard.PwrBrdLedStatus"
)

// ToMap (generated function)
func (m *SensPowerBoard) Dict() map[string]interface{} {
	return map[string]interface{}{
		SENS_POWER_BOARD_FIELD_TIMESTAMP:            m.Timestamp,
		SENS_POWER_BOARD_FIELD_PWR_BRD_SYSTEM_VOLT:  m.PwrBrdSystemVolt,
		SENS_POWER_BOARD_FIELD_PWR_BRD_SERVO_VOLT:   m.PwrBrdServoVolt,
		SENS_POWER_BOARD_FIELD_PWR_BRD_DIGITAL_VOLT: m.PwrBrdDigitalVolt,
		SENS_POWER_BOARD_FIELD_PWR_BRD_MOT_L_AMP:    m.PwrBrdMotLAmp,
		SENS_POWER_BOARD_FIELD_PWR_BRD_MOT_R_AMP:    m.PwrBrdMotRAmp,
		SENS_POWER_BOARD_FIELD_PWR_BRD_ANALOG_AMP:   m.PwrBrdAnalogAmp,
		SENS_POWER_BOARD_FIELD_PWR_BRD_DIGITAL_AMP:  m.PwrBrdDigitalAmp,
		SENS_POWER_BOARD_FIELD_PWR_BRD_EXT_AMP:      m.PwrBrdExtAmp,
		SENS_POWER_BOARD_FIELD_PWR_BRD_AUX_AMP:      m.PwrBrdAuxAmp,
		SENS_POWER_BOARD_FIELD_PWR_BRD_STATUS:       m.PwrBrdStatus,
		SENS_POWER_BOARD_FIELD_PWR_BRD_LED_STATUS:   m.PwrBrdLedStatus,
	}
}

// Marshal (generated function)
func (m *SensPowerBoard) Marshal() ([]byte, error) {
	payload := make([]byte, 46)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.PwrBrdSystemVolt))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.PwrBrdServoVolt))
	binary.LittleEndian.PutUint32(payload[16:], math.Float32bits(m.PwrBrdDigitalVolt))
	binary.LittleEndian.PutUint32(payload[20:], math.Float32bits(m.PwrBrdMotLAmp))
	binary.LittleEndian.PutUint32(payload[24:], math.Float32bits(m.PwrBrdMotRAmp))
	binary.LittleEndian.PutUint32(payload[28:], math.Float32bits(m.PwrBrdAnalogAmp))
	binary.LittleEndian.PutUint32(payload[32:], math.Float32bits(m.PwrBrdDigitalAmp))
	binary.LittleEndian.PutUint32(payload[36:], math.Float32bits(m.PwrBrdExtAmp))
	binary.LittleEndian.PutUint32(payload[40:], math.Float32bits(m.PwrBrdAuxAmp))
	payload[44] = byte(m.PwrBrdStatus)
	payload[45] = byte(m.PwrBrdLedStatus)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SensPowerBoard) Unmarshal(data []byte) error {
	payload := make([]byte, 46)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.PwrBrdSystemVolt = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.PwrBrdServoVolt = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.PwrBrdDigitalVolt = math.Float32frombits(binary.LittleEndian.Uint32(payload[16:]))
	m.PwrBrdMotLAmp = math.Float32frombits(binary.LittleEndian.Uint32(payload[20:]))
	m.PwrBrdMotRAmp = math.Float32frombits(binary.LittleEndian.Uint32(payload[24:]))
	m.PwrBrdAnalogAmp = math.Float32frombits(binary.LittleEndian.Uint32(payload[28:]))
	m.PwrBrdDigitalAmp = math.Float32frombits(binary.LittleEndian.Uint32(payload[32:]))
	m.PwrBrdExtAmp = math.Float32frombits(binary.LittleEndian.Uint32(payload[36:]))
	m.PwrBrdAuxAmp = math.Float32frombits(binary.LittleEndian.Uint32(payload[40:]))
	m.PwrBrdStatus = uint8(payload[44])
	m.PwrBrdLedStatus = uint8(payload[45])
	return nil
}

// GsmLinkStatus struct (generated typeinfo)
// Status of GSM modem (connected to onboard computer)
type GsmLinkStatus struct {
	Timestamp    uint64         `json:"Timestamp" `    // [ us ] Timestamp (of OBC)
	GsmModemType GSM_MODEM_TYPE `json:"GsmModemType" ` // GSM modem used
	GsmLinkType  GSM_LINK_TYPE  `json:"GsmLinkType" `  // GSM link type
	Rssi         uint8          `json:"Rssi" `         // RSSI as reported by modem (unconverted)
	RsrpRscp     uint8          `json:"RsrpRscp" `     // RSRP (LTE) or RSCP (WCDMA) as reported by modem (unconverted)
	SinrEcio     uint8          `json:"SinrEcio" `     // SINR (LTE) or ECIO (WCDMA) as reported by modem (unconverted)
	Rsrq         uint8          `json:"Rsrq" `         // RSRQ (LTE only) as reported by modem (unconverted)
}

// MsgID (generated function)
func (m *GsmLinkStatus) MsgID() message.MessageID {
	return MSG_ID_GSM_LINK_STATUS
}

// String (generated function)
func (m *GsmLinkStatus) String() string {
	return fmt.Sprintf(
		"&ASLUAV.GsmLinkStatus{ Timestamp: %+v, GsmModemType: %+v, GsmLinkType: %+v, Rssi: %+v, RsrpRscp: %+v, SinrEcio: %+v, Rsrq: %+v }",
		m.Timestamp,
		m.GsmModemType,
		m.GsmLinkType,
		m.Rssi,
		m.RsrpRscp,
		m.SinrEcio,
		m.Rsrq,
	)
}

const (
	GSM_LINK_STATUS_FIELD_TIMESTAMP      = "GsmLinkStatus.Timestamp"
	GSM_LINK_STATUS_FIELD_GSM_MODEM_TYPE = "GsmLinkStatus.GsmModemType"
	GSM_LINK_STATUS_FIELD_GSM_LINK_TYPE  = "GsmLinkStatus.GsmLinkType"
	GSM_LINK_STATUS_FIELD_RSSI           = "GsmLinkStatus.Rssi"
	GSM_LINK_STATUS_FIELD_RSRP_RSCP      = "GsmLinkStatus.RsrpRscp"
	GSM_LINK_STATUS_FIELD_SINR_ECIO      = "GsmLinkStatus.SinrEcio"
	GSM_LINK_STATUS_FIELD_RSRQ           = "GsmLinkStatus.Rsrq"
)

// ToMap (generated function)
func (m *GsmLinkStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		GSM_LINK_STATUS_FIELD_TIMESTAMP:      m.Timestamp,
		GSM_LINK_STATUS_FIELD_GSM_MODEM_TYPE: m.GsmModemType,
		GSM_LINK_STATUS_FIELD_GSM_LINK_TYPE:  m.GsmLinkType,
		GSM_LINK_STATUS_FIELD_RSSI:           m.Rssi,
		GSM_LINK_STATUS_FIELD_RSRP_RSCP:      m.RsrpRscp,
		GSM_LINK_STATUS_FIELD_SINR_ECIO:      m.SinrEcio,
		GSM_LINK_STATUS_FIELD_RSRQ:           m.Rsrq,
	}
}

// Marshal (generated function)
func (m *GsmLinkStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 14)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	payload[8] = byte(m.GsmModemType)
	payload[9] = byte(m.GsmLinkType)
	payload[10] = byte(m.Rssi)
	payload[11] = byte(m.RsrpRscp)
	payload[12] = byte(m.SinrEcio)
	payload[13] = byte(m.Rsrq)
	return payload, nil
}

// Unmarshal (generated function)
func (m *GsmLinkStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 14)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.GsmModemType = GSM_MODEM_TYPE(payload[8])
	m.GsmLinkType = GSM_LINK_TYPE(payload[9])
	m.Rssi = uint8(payload[10])
	m.RsrpRscp = uint8(payload[11])
	m.SinrEcio = uint8(payload[12])
	m.Rsrq = uint8(payload[13])
	return nil
}

// SatcomLinkStatus struct (generated typeinfo)
// Status of the SatCom link
type SatcomLinkStatus struct {
	Timestamp          uint64 `json:"Timestamp" `          // [ us ] Timestamp
	LastHeartbeat      uint64 `json:"LastHeartbeat" `      // [ us ] Timestamp of the last successful sbd session
	FailedSessions     uint16 `json:"FailedSessions" `     // Number of failed sessions
	SuccessfulSessions uint16 `json:"SuccessfulSessions" ` // Number of successful sessions
	SignalQuality      uint8  `json:"SignalQuality" `      // Signal quality
	RingPending        uint8  `json:"RingPending" `        // Ring call pending
	TxSessionPending   uint8  `json:"TxSessionPending" `   // Transmission session pending
	RxSessionPending   uint8  `json:"RxSessionPending" `   // Receiving session pending
}

// MsgID (generated function)
func (m *SatcomLinkStatus) MsgID() message.MessageID {
	return MSG_ID_SATCOM_LINK_STATUS
}

// String (generated function)
func (m *SatcomLinkStatus) String() string {
	return fmt.Sprintf(
		"&ASLUAV.SatcomLinkStatus{ Timestamp: %+v, LastHeartbeat: %+v, FailedSessions: %+v, SuccessfulSessions: %+v, SignalQuality: %+v, RingPending: %+v, TxSessionPending: %+v, RxSessionPending: %+v }",
		m.Timestamp,
		m.LastHeartbeat,
		m.FailedSessions,
		m.SuccessfulSessions,
		m.SignalQuality,
		m.RingPending,
		m.TxSessionPending,
		m.RxSessionPending,
	)
}

const (
	SATCOM_LINK_STATUS_FIELD_TIMESTAMP           = "SatcomLinkStatus.Timestamp"
	SATCOM_LINK_STATUS_FIELD_LAST_HEARTBEAT      = "SatcomLinkStatus.LastHeartbeat"
	SATCOM_LINK_STATUS_FIELD_FAILED_SESSIONS     = "SatcomLinkStatus.FailedSessions"
	SATCOM_LINK_STATUS_FIELD_SUCCESSFUL_SESSIONS = "SatcomLinkStatus.SuccessfulSessions"
	SATCOM_LINK_STATUS_FIELD_SIGNAL_QUALITY      = "SatcomLinkStatus.SignalQuality"
	SATCOM_LINK_STATUS_FIELD_RING_PENDING        = "SatcomLinkStatus.RingPending"
	SATCOM_LINK_STATUS_FIELD_TX_SESSION_PENDING  = "SatcomLinkStatus.TxSessionPending"
	SATCOM_LINK_STATUS_FIELD_RX_SESSION_PENDING  = "SatcomLinkStatus.RxSessionPending"
)

// ToMap (generated function)
func (m *SatcomLinkStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		SATCOM_LINK_STATUS_FIELD_TIMESTAMP:           m.Timestamp,
		SATCOM_LINK_STATUS_FIELD_LAST_HEARTBEAT:      m.LastHeartbeat,
		SATCOM_LINK_STATUS_FIELD_FAILED_SESSIONS:     m.FailedSessions,
		SATCOM_LINK_STATUS_FIELD_SUCCESSFUL_SESSIONS: m.SuccessfulSessions,
		SATCOM_LINK_STATUS_FIELD_SIGNAL_QUALITY:      m.SignalQuality,
		SATCOM_LINK_STATUS_FIELD_RING_PENDING:        m.RingPending,
		SATCOM_LINK_STATUS_FIELD_TX_SESSION_PENDING:  m.TxSessionPending,
		SATCOM_LINK_STATUS_FIELD_RX_SESSION_PENDING:  m.RxSessionPending,
	}
}

// Marshal (generated function)
func (m *SatcomLinkStatus) Marshal() ([]byte, error) {
	payload := make([]byte, 24)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint64(payload[8:], uint64(m.LastHeartbeat))
	binary.LittleEndian.PutUint16(payload[16:], uint16(m.FailedSessions))
	binary.LittleEndian.PutUint16(payload[18:], uint16(m.SuccessfulSessions))
	payload[20] = byte(m.SignalQuality)
	payload[21] = byte(m.RingPending)
	payload[22] = byte(m.TxSessionPending)
	payload[23] = byte(m.RxSessionPending)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SatcomLinkStatus) Unmarshal(data []byte) error {
	payload := make([]byte, 24)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.LastHeartbeat = uint64(binary.LittleEndian.Uint64(payload[8:]))
	m.FailedSessions = uint16(binary.LittleEndian.Uint16(payload[16:]))
	m.SuccessfulSessions = uint16(binary.LittleEndian.Uint16(payload[18:]))
	m.SignalQuality = uint8(payload[20])
	m.RingPending = uint8(payload[21])
	m.TxSessionPending = uint8(payload[22])
	m.RxSessionPending = uint8(payload[23])
	return nil
}

// SensorAirflowAngles struct (generated typeinfo)
// Calibrated airflow angle measurements
type SensorAirflowAngles struct {
	Timestamp          uint64  `json:"Timestamp" `          // [ us ] Timestamp
	Angleofattack      float32 `json:"Angleofattack" `      // [ deg ] Angle of attack
	Sideslip           float32 `json:"Sideslip" `           // [ deg ] Sideslip angle
	AngleofattackValid uint8   `json:"AngleofattackValid" ` // Angle of attack measurement valid
	SideslipValid      uint8   `json:"SideslipValid" `      // Sideslip angle measurement valid
}

// MsgID (generated function)
func (m *SensorAirflowAngles) MsgID() message.MessageID {
	return MSG_ID_SENSOR_AIRFLOW_ANGLES
}

// String (generated function)
func (m *SensorAirflowAngles) String() string {
	return fmt.Sprintf(
		"&ASLUAV.SensorAirflowAngles{ Timestamp: %+v, Angleofattack: %+v, Sideslip: %+v, AngleofattackValid: %+v, SideslipValid: %+v }",
		m.Timestamp,
		m.Angleofattack,
		m.Sideslip,
		m.AngleofattackValid,
		m.SideslipValid,
	)
}

const (
	SENSOR_AIRFLOW_ANGLES_FIELD_TIMESTAMP           = "SensorAirflowAngles.Timestamp"
	SENSOR_AIRFLOW_ANGLES_FIELD_ANGLEOFATTACK       = "SensorAirflowAngles.Angleofattack"
	SENSOR_AIRFLOW_ANGLES_FIELD_SIDESLIP            = "SensorAirflowAngles.Sideslip"
	SENSOR_AIRFLOW_ANGLES_FIELD_ANGLEOFATTACK_VALID = "SensorAirflowAngles.AngleofattackValid"
	SENSOR_AIRFLOW_ANGLES_FIELD_SIDESLIP_VALID      = "SensorAirflowAngles.SideslipValid"
)

// ToMap (generated function)
func (m *SensorAirflowAngles) Dict() map[string]interface{} {
	return map[string]interface{}{
		SENSOR_AIRFLOW_ANGLES_FIELD_TIMESTAMP:           m.Timestamp,
		SENSOR_AIRFLOW_ANGLES_FIELD_ANGLEOFATTACK:       m.Angleofattack,
		SENSOR_AIRFLOW_ANGLES_FIELD_SIDESLIP:            m.Sideslip,
		SENSOR_AIRFLOW_ANGLES_FIELD_ANGLEOFATTACK_VALID: m.AngleofattackValid,
		SENSOR_AIRFLOW_ANGLES_FIELD_SIDESLIP_VALID:      m.SideslipValid,
	}
}

// Marshal (generated function)
func (m *SensorAirflowAngles) Marshal() ([]byte, error) {
	payload := make([]byte, 18)
	binary.LittleEndian.PutUint64(payload[0:], uint64(m.Timestamp))
	binary.LittleEndian.PutUint32(payload[8:], math.Float32bits(m.Angleofattack))
	binary.LittleEndian.PutUint32(payload[12:], math.Float32bits(m.Sideslip))
	payload[16] = byte(m.AngleofattackValid)
	payload[17] = byte(m.SideslipValid)
	return payload, nil
}

// Unmarshal (generated function)
func (m *SensorAirflowAngles) Unmarshal(data []byte) error {
	payload := make([]byte, 18)
	copy(payload[0:], data)
	m.Timestamp = uint64(binary.LittleEndian.Uint64(payload[0:]))
	m.Angleofattack = math.Float32frombits(binary.LittleEndian.Uint32(payload[8:]))
	m.Sideslip = math.Float32frombits(binary.LittleEndian.Uint32(payload[12:]))
	m.AngleofattackValid = uint8(payload[16])
	m.SideslipValid = uint8(payload[17])
	return nil
}

// SysStatus struct (generated typeinfo)
// The general system state. If the system is following the MAVLink standard, the system state is mainly defined by three orthogonal states/modes: The system mode, which is either LOCKED (motors shut down and locked), MANUAL (system under RC control), GUIDED (system with autonomous position control, position setpoint controlled manually) or AUTO (system guided by path/waypoint planner). The NAV_MODE defined the current flight state: LIFTOFF (often an open-loop maneuver), LANDING, WAYPOINTS or VECTOR. This represents the internal navigation state machine. The system status shows whether the system is currently active or not and if an emergency occurred. During the CRITICAL and EMERGENCY states the MAV is still considered to be active, but should start emergency procedures autonomously. After a failure occurred it should first move from active to critical to allow manual intervention and then move to emergency after a certain timeout.
type SysStatus struct {
	OnboardControlSensorsPresent MAV_SYS_STATUS_SENSOR `json:"OnboardControlSensorsPresent" ` // Bitmap showing which onboard controllers and sensors are present. Value of 0: not present. Value of 1: present.
	OnboardControlSensorsEnabled MAV_SYS_STATUS_SENSOR `json:"OnboardControlSensorsEnabled" ` // Bitmap showing which onboard controllers and sensors are enabled:  Value of 0: not enabled. Value of 1: enabled.
	OnboardControlSensorsHealth  MAV_SYS_STATUS_SENSOR `json:"OnboardControlSensorsHealth" `  // Bitmap showing which onboard controllers and sensors have an error (or are operational). Value of 0: error. Value of 1: healthy.
	Load                         uint16                `json:"Load" `                         // [ d% ] Maximum usage in percent of the mainloop time. Values: [0-1000] - should always be below 1000
	VoltageBattery               uint16                `json:"VoltageBattery" `               // [ mV ] Battery voltage, UINT16_MAX: Voltage not sent by autopilot
	CurrentBattery               int16                 `json:"CurrentBattery" `               // [ cA ] Battery current, -1: Current not sent by autopilot
	DropRateComm                 uint16                `json:"DropRateComm" `                 // [ c% ] Communication drop rate, (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV)
	ErrorsComm                   uint16                `json:"ErrorsComm" `                   // Communication errors (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV)
	ErrorsCount1                 uint16                `json:"ErrorsCount1" `                 // Autopilot-specific errors
	ErrorsCount2                 uint16                `json:"ErrorsCount2" `                 // Autopilot-specific errors
	ErrorsCount3                 uint16                `json:"ErrorsCount3" `                 // Autopilot-specific errors
	ErrorsCount4                 uint16                `json:"ErrorsCount4" `                 // Autopilot-specific errors
	BatteryRemaining             int8                  `json:"BatteryRemaining" `             // [ % ] Battery energy remaining, -1: Battery remaining energy not sent by autopilot
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

const (
	SYS_STATUS_FIELD_ONBOARD_CONTROL_SENSORS_PRESENT = "SysStatus.OnboardControlSensorsPresent"
	SYS_STATUS_FIELD_ONBOARD_CONTROL_SENSORS_ENABLED = "SysStatus.OnboardControlSensorsEnabled"
	SYS_STATUS_FIELD_ONBOARD_CONTROL_SENSORS_HEALTH  = "SysStatus.OnboardControlSensorsHealth"
	SYS_STATUS_FIELD_LOAD                            = "SysStatus.Load"
	SYS_STATUS_FIELD_VOLTAGE_BATTERY                 = "SysStatus.VoltageBattery"
	SYS_STATUS_FIELD_CURRENT_BATTERY                 = "SysStatus.CurrentBattery"
	SYS_STATUS_FIELD_DROP_RATE_COMM                  = "SysStatus.DropRateComm"
	SYS_STATUS_FIELD_ERRORS_COMM                     = "SysStatus.ErrorsComm"
	SYS_STATUS_FIELD_ERRORS_COUNT1                   = "SysStatus.ErrorsCount1"
	SYS_STATUS_FIELD_ERRORS_COUNT2                   = "SysStatus.ErrorsCount2"
	SYS_STATUS_FIELD_ERRORS_COUNT3                   = "SysStatus.ErrorsCount3"
	SYS_STATUS_FIELD_ERRORS_COUNT4                   = "SysStatus.ErrorsCount4"
	SYS_STATUS_FIELD_BATTERY_REMAINING               = "SysStatus.BatteryRemaining"
)

// ToMap (generated function)
func (m *SysStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		SYS_STATUS_FIELD_ONBOARD_CONTROL_SENSORS_PRESENT: m.OnboardControlSensorsPresent,
		SYS_STATUS_FIELD_ONBOARD_CONTROL_SENSORS_ENABLED: m.OnboardControlSensorsEnabled,
		SYS_STATUS_FIELD_ONBOARD_CONTROL_SENSORS_HEALTH:  m.OnboardControlSensorsHealth,
		SYS_STATUS_FIELD_LOAD:                            m.Load,
		SYS_STATUS_FIELD_VOLTAGE_BATTERY:                 m.VoltageBattery,
		SYS_STATUS_FIELD_CURRENT_BATTERY:                 m.CurrentBattery,
		SYS_STATUS_FIELD_DROP_RATE_COMM:                  m.DropRateComm,
		SYS_STATUS_FIELD_ERRORS_COMM:                     m.ErrorsComm,
		SYS_STATUS_FIELD_ERRORS_COUNT1:                   m.ErrorsCount1,
		SYS_STATUS_FIELD_ERRORS_COUNT2:                   m.ErrorsCount2,
		SYS_STATUS_FIELD_ERRORS_COUNT3:                   m.ErrorsCount3,
		SYS_STATUS_FIELD_ERRORS_COUNT4:                   m.ErrorsCount4,
		SYS_STATUS_FIELD_BATTERY_REMAINING:               m.BatteryRemaining,
	}
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
	TimeUnixUsec uint64 `json:"TimeUnixUsec" ` // [ us ] Timestamp (UNIX epoch time).
	TimeBootMs   uint32 `json:"TimeBootMs" `   // [ ms ] Timestamp (time since system boot).
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

const (
	SYSTEM_TIME_FIELD_TIME_UNIX_USEC = "SystemTime.TimeUnixUsec"
	SYSTEM_TIME_FIELD_TIME_BOOT_MS   = "SystemTime.TimeBootMs"
)

// ToMap (generated function)
func (m *SystemTime) Dict() map[string]interface{} {
	return map[string]interface{}{
		SYSTEM_TIME_FIELD_TIME_UNIX_USEC: m.TimeUnixUsec,
		SYSTEM_TIME_FIELD_TIME_BOOT_MS:   m.TimeBootMs,
	}
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
	TimeUsec        uint64 `json:"TimeUsec" `        // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Seq             uint32 `json:"Seq" `             // PING sequence
	TargetSystem    uint8  `json:"TargetSystem" `    // 0: request ping from all receiving systems. If greater than 0: message is a ping response and number is the system id of the requesting system
	TargetComponent uint8  `json:"TargetComponent" ` // 0: request ping from all receiving components. If greater than 0: message is a ping response and number is the component id of the requesting component.
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

const (
	PING_FIELD_TIME_USEC        = "Ping.TimeUsec"
	PING_FIELD_SEQ              = "Ping.Seq"
	PING_FIELD_TARGET_SYSTEM    = "Ping.TargetSystem"
	PING_FIELD_TARGET_COMPONENT = "Ping.TargetComponent"
)

// ToMap (generated function)
func (m *Ping) Dict() map[string]interface{} {
	return map[string]interface{}{
		PING_FIELD_TIME_USEC:        m.TimeUsec,
		PING_FIELD_SEQ:              m.Seq,
		PING_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		PING_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	TargetSystem   uint8  `json:"TargetSystem" `     // System the GCS requests control for
	ControlRequest uint8  `json:"ControlRequest" `   // 0: request control of this MAV, 1: Release control of this MAV
	Version        uint8  `json:"Version" `          // [ rad ] 0: key as plaintext, 1-255: future, different hashing/encryption variants. The GCS should in general use the safest mode possible initially and then gradually move down the encryption level if it gets a NACK message indicating an encryption mismatch.
	Passkey        string `json:"Passkey" len:"25" ` // Password / Key, depending on version plaintext or encrypted. 25 or less characters, NULL terminated. The characters may involve A-Z, a-z, 0-9, and "!?,.-"
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
		strings.TrimRight(m.Passkey, string(byte(0))),
	)
}

const (
	CHANGE_OPERATOR_CONTROL_FIELD_TARGET_SYSTEM   = "ChangeOperatorControl.TargetSystem"
	CHANGE_OPERATOR_CONTROL_FIELD_CONTROL_REQUEST = "ChangeOperatorControl.ControlRequest"
	CHANGE_OPERATOR_CONTROL_FIELD_VERSION         = "ChangeOperatorControl.Version"
	CHANGE_OPERATOR_CONTROL_FIELD_PASSKEY         = "ChangeOperatorControl.Passkey"
)

// ToMap (generated function)
func (m *ChangeOperatorControl) Dict() map[string]interface{} {
	return map[string]interface{}{
		CHANGE_OPERATOR_CONTROL_FIELD_TARGET_SYSTEM:   m.TargetSystem,
		CHANGE_OPERATOR_CONTROL_FIELD_CONTROL_REQUEST: m.ControlRequest,
		CHANGE_OPERATOR_CONTROL_FIELD_VERSION:         m.Version,
		CHANGE_OPERATOR_CONTROL_FIELD_PASSKEY:         m.Passkey,
	}
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
	GcsSystemID    uint8 `json:"GcsSystemID" `    // ID of the GCS this message
	ControlRequest uint8 `json:"ControlRequest" ` // 0: request control of this MAV, 1: Release control of this MAV
	Ack            uint8 `json:"Ack" `            // 0: ACK, 1: NACK: Wrong passkey, 2: NACK: Unsupported passkey encryption method, 3: NACK: Already under control
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

const (
	CHANGE_OPERATOR_CONTROL_ACK_FIELD_GCS_SYSTEM_ID   = "ChangeOperatorControlAck.GcsSystemID"
	CHANGE_OPERATOR_CONTROL_ACK_FIELD_CONTROL_REQUEST = "ChangeOperatorControlAck.ControlRequest"
	CHANGE_OPERATOR_CONTROL_ACK_FIELD_ACK             = "ChangeOperatorControlAck.Ack"
)

// ToMap (generated function)
func (m *ChangeOperatorControlAck) Dict() map[string]interface{} {
	return map[string]interface{}{
		CHANGE_OPERATOR_CONTROL_ACK_FIELD_GCS_SYSTEM_ID:   m.GcsSystemID,
		CHANGE_OPERATOR_CONTROL_ACK_FIELD_CONTROL_REQUEST: m.ControlRequest,
		CHANGE_OPERATOR_CONTROL_ACK_FIELD_ACK:             m.Ack,
	}
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
	Key string `json:"Key" len:"32" ` // key
}

// MsgID (generated function)
func (m *AuthKey) MsgID() message.MessageID {
	return MSG_ID_AUTH_KEY
}

// String (generated function)
func (m *AuthKey) String() string {
	return fmt.Sprintf(
		"&common.AuthKey{ Key: \"%s\" }",
		strings.TrimRight(m.Key, string(byte(0))),
	)
}

const (
	AUTH_KEY_FIELD_KEY = "AuthKey.Key"
)

// ToMap (generated function)
func (m *AuthKey) Dict() map[string]interface{} {
	return map[string]interface{}{
		AUTH_KEY_FIELD_KEY: m.Key,
	}
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
	Timestamp        uint64 `json:"Timestamp" `        // [ ms ] Timestamp (time since system boot).
	TxRate           uint32 `json:"TxRate" `           // [ bytes/s ] Transmit rate
	RxRate           uint32 `json:"RxRate" `           // [ bytes/s ] Receive rate
	MessagesSent     uint32 `json:"MessagesSent" `     // Messages sent
	MessagesReceived uint32 `json:"MessagesReceived" ` // Messages received (estimated from counting seq)
	MessagesLost     uint32 `json:"MessagesLost" `     // Messages lost (estimated from counting seq)
	RxParseErr       uint16 `json:"RxParseErr" `       // [ bytes ] Number of bytes that could not be parsed correctly.
	TxOverflows      uint16 `json:"TxOverflows" `      // [ bytes ] Transmit buffer overflows. This number wraps around as it reaches UINT16_MAX
	RxOverflows      uint16 `json:"RxOverflows" `      // [ bytes ] Receive buffer overflows. This number wraps around as it reaches UINT16_MAX
	TxBuf            uint8  `json:"TxBuf" `            // [ % ] Remaining free transmit buffer space
	RxBuf            uint8  `json:"RxBuf" `            // [ % ] Remaining free receive buffer space
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

const (
	LINK_NODE_STATUS_FIELD_TIMESTAMP         = "LinkNodeStatus.Timestamp"
	LINK_NODE_STATUS_FIELD_TX_RATE           = "LinkNodeStatus.TxRate"
	LINK_NODE_STATUS_FIELD_RX_RATE           = "LinkNodeStatus.RxRate"
	LINK_NODE_STATUS_FIELD_MESSAGES_SENT     = "LinkNodeStatus.MessagesSent"
	LINK_NODE_STATUS_FIELD_MESSAGES_RECEIVED = "LinkNodeStatus.MessagesReceived"
	LINK_NODE_STATUS_FIELD_MESSAGES_LOST     = "LinkNodeStatus.MessagesLost"
	LINK_NODE_STATUS_FIELD_RX_PARSE_ERR      = "LinkNodeStatus.RxParseErr"
	LINK_NODE_STATUS_FIELD_TX_OVERFLOWS      = "LinkNodeStatus.TxOverflows"
	LINK_NODE_STATUS_FIELD_RX_OVERFLOWS      = "LinkNodeStatus.RxOverflows"
	LINK_NODE_STATUS_FIELD_TX_BUF            = "LinkNodeStatus.TxBuf"
	LINK_NODE_STATUS_FIELD_RX_BUF            = "LinkNodeStatus.RxBuf"
)

// ToMap (generated function)
func (m *LinkNodeStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		LINK_NODE_STATUS_FIELD_TIMESTAMP:         m.Timestamp,
		LINK_NODE_STATUS_FIELD_TX_RATE:           m.TxRate,
		LINK_NODE_STATUS_FIELD_RX_RATE:           m.RxRate,
		LINK_NODE_STATUS_FIELD_MESSAGES_SENT:     m.MessagesSent,
		LINK_NODE_STATUS_FIELD_MESSAGES_RECEIVED: m.MessagesReceived,
		LINK_NODE_STATUS_FIELD_MESSAGES_LOST:     m.MessagesLost,
		LINK_NODE_STATUS_FIELD_RX_PARSE_ERR:      m.RxParseErr,
		LINK_NODE_STATUS_FIELD_TX_OVERFLOWS:      m.TxOverflows,
		LINK_NODE_STATUS_FIELD_RX_OVERFLOWS:      m.RxOverflows,
		LINK_NODE_STATUS_FIELD_TX_BUF:            m.TxBuf,
		LINK_NODE_STATUS_FIELD_RX_BUF:            m.RxBuf,
	}
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
	CustomMode   uint32   `json:"CustomMode" `   // The new autopilot-specific mode. This field can be ignored by an autopilot.
	TargetSystem uint8    `json:"TargetSystem" ` // The system setting the mode
	BaseMode     MAV_MODE `json:"BaseMode" `     // The new base mode.
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

const (
	SET_MODE_FIELD_CUSTOM_MODE   = "SetMode.CustomMode"
	SET_MODE_FIELD_TARGET_SYSTEM = "SetMode.TargetSystem"
	SET_MODE_FIELD_BASE_MODE     = "SetMode.BaseMode"
)

// ToMap (generated function)
func (m *SetMode) Dict() map[string]interface{} {
	return map[string]interface{}{
		SET_MODE_FIELD_CUSTOM_MODE:   m.CustomMode,
		SET_MODE_FIELD_TARGET_SYSTEM: m.TargetSystem,
		SET_MODE_FIELD_BASE_MODE:     m.BaseMode,
	}
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
	ParamValue      float32        `json:"ParamValue" `       // Parameter value (new value if PARAM_ACCEPTED, current value otherwise)
	TargetSystem    uint8          `json:"TargetSystem" `     // Id of system that sent PARAM_SET message.
	TargetComponent uint8          `json:"TargetComponent" `  // Id of system that sent PARAM_SET message.
	ParamID         string         `json:"ParamID" len:"16" ` // Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamType       MAV_PARAM_TYPE `json:"ParamType" `        // Parameter type.
	ParamResult     PARAM_ACK      `json:"ParamResult" `      // Result code.
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
		strings.TrimRight(m.ParamID, string(byte(0))),
		m.ParamType,
		m.ParamResult,
	)
}

const (
	PARAM_ACK_TRANSACTION_FIELD_PARAM_VALUE      = "ParamAckTransaction.ParamValue"
	PARAM_ACK_TRANSACTION_FIELD_TARGET_SYSTEM    = "ParamAckTransaction.TargetSystem"
	PARAM_ACK_TRANSACTION_FIELD_TARGET_COMPONENT = "ParamAckTransaction.TargetComponent"
	PARAM_ACK_TRANSACTION_FIELD_PARAM_ID         = "ParamAckTransaction.ParamID"
	PARAM_ACK_TRANSACTION_FIELD_PARAM_TYPE       = "ParamAckTransaction.ParamType"
	PARAM_ACK_TRANSACTION_FIELD_PARAM_RESULT     = "ParamAckTransaction.ParamResult"
)

// ToMap (generated function)
func (m *ParamAckTransaction) Dict() map[string]interface{} {
	return map[string]interface{}{
		PARAM_ACK_TRANSACTION_FIELD_PARAM_VALUE:      m.ParamValue,
		PARAM_ACK_TRANSACTION_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		PARAM_ACK_TRANSACTION_FIELD_TARGET_COMPONENT: m.TargetComponent,
		PARAM_ACK_TRANSACTION_FIELD_PARAM_ID:         m.ParamID,
		PARAM_ACK_TRANSACTION_FIELD_PARAM_TYPE:       m.ParamType,
		PARAM_ACK_TRANSACTION_FIELD_PARAM_RESULT:     m.ParamResult,
	}
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
	ParamIndex      int16  `json:"ParamIndex" `       // Parameter index. Send -1 to use the param ID field as identifier (else the param id will be ignored)
	TargetSystem    uint8  `json:"TargetSystem" `     // System ID
	TargetComponent uint8  `json:"TargetComponent" `  // Component ID
	ParamID         string `json:"ParamID" len:"16" ` // Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
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
		strings.TrimRight(m.ParamID, string(byte(0))),
	)
}

const (
	PARAM_REQUEST_READ_FIELD_PARAM_INDEX      = "ParamRequestRead.ParamIndex"
	PARAM_REQUEST_READ_FIELD_TARGET_SYSTEM    = "ParamRequestRead.TargetSystem"
	PARAM_REQUEST_READ_FIELD_TARGET_COMPONENT = "ParamRequestRead.TargetComponent"
	PARAM_REQUEST_READ_FIELD_PARAM_ID         = "ParamRequestRead.ParamID"
)

// ToMap (generated function)
func (m *ParamRequestRead) Dict() map[string]interface{} {
	return map[string]interface{}{
		PARAM_REQUEST_READ_FIELD_PARAM_INDEX:      m.ParamIndex,
		PARAM_REQUEST_READ_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		PARAM_REQUEST_READ_FIELD_TARGET_COMPONENT: m.TargetComponent,
		PARAM_REQUEST_READ_FIELD_PARAM_ID:         m.ParamID,
	}
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
	TargetSystem    uint8 `json:"TargetSystem" `    // System ID
	TargetComponent uint8 `json:"TargetComponent" ` // Component ID
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

const (
	PARAM_REQUEST_LIST_FIELD_TARGET_SYSTEM    = "ParamRequestList.TargetSystem"
	PARAM_REQUEST_LIST_FIELD_TARGET_COMPONENT = "ParamRequestList.TargetComponent"
)

// ToMap (generated function)
func (m *ParamRequestList) Dict() map[string]interface{} {
	return map[string]interface{}{
		PARAM_REQUEST_LIST_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		PARAM_REQUEST_LIST_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	ParamValue float32        `json:"ParamValue" `       // Onboard parameter value
	ParamCount uint16         `json:"ParamCount" `       // Total number of onboard parameters
	ParamIndex uint16         `json:"ParamIndex" `       // Index of this onboard parameter
	ParamID    string         `json:"ParamID" len:"16" ` // Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamType  MAV_PARAM_TYPE `json:"ParamType" `        // Onboard parameter type.
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
		strings.TrimRight(m.ParamID, string(byte(0))),
		m.ParamType,
	)
}

const (
	PARAM_VALUE_FIELD_PARAM_VALUE = "ParamValue.ParamValue"
	PARAM_VALUE_FIELD_PARAM_COUNT = "ParamValue.ParamCount"
	PARAM_VALUE_FIELD_PARAM_INDEX = "ParamValue.ParamIndex"
	PARAM_VALUE_FIELD_PARAM_ID    = "ParamValue.ParamID"
	PARAM_VALUE_FIELD_PARAM_TYPE  = "ParamValue.ParamType"
)

// ToMap (generated function)
func (m *ParamValue) Dict() map[string]interface{} {
	return map[string]interface{}{
		PARAM_VALUE_FIELD_PARAM_VALUE: m.ParamValue,
		PARAM_VALUE_FIELD_PARAM_COUNT: m.ParamCount,
		PARAM_VALUE_FIELD_PARAM_INDEX: m.ParamIndex,
		PARAM_VALUE_FIELD_PARAM_ID:    m.ParamID,
		PARAM_VALUE_FIELD_PARAM_TYPE:  m.ParamType,
	}
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
	ParamValue      float32        `json:"ParamValue" `       // Onboard parameter value
	TargetSystem    uint8          `json:"TargetSystem" `     // System ID
	TargetComponent uint8          `json:"TargetComponent" `  // Component ID
	ParamID         string         `json:"ParamID" len:"16" ` // Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamType       MAV_PARAM_TYPE `json:"ParamType" `        // Onboard parameter type.
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
		strings.TrimRight(m.ParamID, string(byte(0))),
		m.ParamType,
	)
}

const (
	PARAM_SET_FIELD_PARAM_VALUE      = "ParamSet.ParamValue"
	PARAM_SET_FIELD_TARGET_SYSTEM    = "ParamSet.TargetSystem"
	PARAM_SET_FIELD_TARGET_COMPONENT = "ParamSet.TargetComponent"
	PARAM_SET_FIELD_PARAM_ID         = "ParamSet.ParamID"
	PARAM_SET_FIELD_PARAM_TYPE       = "ParamSet.ParamType"
)

// ToMap (generated function)
func (m *ParamSet) Dict() map[string]interface{} {
	return map[string]interface{}{
		PARAM_SET_FIELD_PARAM_VALUE:      m.ParamValue,
		PARAM_SET_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		PARAM_SET_FIELD_TARGET_COMPONENT: m.TargetComponent,
		PARAM_SET_FIELD_PARAM_ID:         m.ParamID,
		PARAM_SET_FIELD_PARAM_TYPE:       m.ParamType,
	}
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
	TimeUsec          uint64       `json:"TimeUsec" `          // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Lat               int32        `json:"Lat" `               // [ degE7 ] Latitude (WGS84, EGM96 ellipsoid)
	Lon               int32        `json:"Lon" `               // [ degE7 ] Longitude (WGS84, EGM96 ellipsoid)
	Alt               int32        `json:"Alt" `               // [ mm ] Altitude (MSL). Positive for up. Note that virtually all GPS modules provide the MSL altitude in addition to the WGS84 altitude.
	Eph               uint16       `json:"Eph" `               // GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX
	Epv               uint16       `json:"Epv" `               // GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX
	Vel               uint16       `json:"Vel" `               // [ cm/s ] GPS ground speed. If unknown, set to: UINT16_MAX
	Cog               uint16       `json:"Cog" `               // [ cdeg ] Course over ground (NOT heading, but direction of movement) in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	FixType           GPS_FIX_TYPE `json:"FixType" `           // GPS fix type.
	SatellitesVisible uint8        `json:"SatellitesVisible" ` // Number of satellites visible. If unknown, set to 255
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

const (
	GPS_RAW_INT_FIELD_TIME_USEC          = "GpsRawInt.TimeUsec"
	GPS_RAW_INT_FIELD_LAT                = "GpsRawInt.Lat"
	GPS_RAW_INT_FIELD_LON                = "GpsRawInt.Lon"
	GPS_RAW_INT_FIELD_ALT                = "GpsRawInt.Alt"
	GPS_RAW_INT_FIELD_EPH                = "GpsRawInt.Eph"
	GPS_RAW_INT_FIELD_EPV                = "GpsRawInt.Epv"
	GPS_RAW_INT_FIELD_VEL                = "GpsRawInt.Vel"
	GPS_RAW_INT_FIELD_COG                = "GpsRawInt.Cog"
	GPS_RAW_INT_FIELD_FIX_TYPE           = "GpsRawInt.FixType"
	GPS_RAW_INT_FIELD_SATELLITES_VISIBLE = "GpsRawInt.SatellitesVisible"
)

// ToMap (generated function)
func (m *GpsRawInt) Dict() map[string]interface{} {
	return map[string]interface{}{
		GPS_RAW_INT_FIELD_TIME_USEC:          m.TimeUsec,
		GPS_RAW_INT_FIELD_LAT:                m.Lat,
		GPS_RAW_INT_FIELD_LON:                m.Lon,
		GPS_RAW_INT_FIELD_ALT:                m.Alt,
		GPS_RAW_INT_FIELD_EPH:                m.Eph,
		GPS_RAW_INT_FIELD_EPV:                m.Epv,
		GPS_RAW_INT_FIELD_VEL:                m.Vel,
		GPS_RAW_INT_FIELD_COG:                m.Cog,
		GPS_RAW_INT_FIELD_FIX_TYPE:           m.FixType,
		GPS_RAW_INT_FIELD_SATELLITES_VISIBLE: m.SatellitesVisible,
	}
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
	SatellitesVisible  uint8   `json:"SatellitesVisible" `           // Number of satellites visible
	SatellitePrn       []uint8 `json:"SatellitePrn" len:"20" `       // Global satellite ID
	SatelliteUsed      []uint8 `json:"SatelliteUsed" len:"20" `      // 0: Satellite not used, 1: used for localization
	SatelliteElevation []uint8 `json:"SatelliteElevation" len:"20" ` // [ deg ] Elevation (0: right on top of receiver, 90: on the horizon) of satellite
	SatelliteAzimuth   []uint8 `json:"SatelliteAzimuth" len:"20" `   // [ deg ] Direction of satellite, 0: 0 deg, 255: 360 deg.
	SatelliteSnr       []uint8 `json:"SatelliteSnr" len:"20" `       // [ dB ] Signal to noise ratio of satellite
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

const (
	GPS_STATUS_FIELD_SATELLITES_VISIBLE  = "GpsStatus.SatellitesVisible"
	GPS_STATUS_FIELD_SATELLITE_PRN       = "GpsStatus.SatellitePrn"
	GPS_STATUS_FIELD_SATELLITE_USED      = "GpsStatus.SatelliteUsed"
	GPS_STATUS_FIELD_SATELLITE_ELEVATION = "GpsStatus.SatelliteElevation"
	GPS_STATUS_FIELD_SATELLITE_AZIMUTH   = "GpsStatus.SatelliteAzimuth"
	GPS_STATUS_FIELD_SATELLITE_SNR       = "GpsStatus.SatelliteSnr"
)

// ToMap (generated function)
func (m *GpsStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		GPS_STATUS_FIELD_SATELLITES_VISIBLE:  m.SatellitesVisible,
		GPS_STATUS_FIELD_SATELLITE_PRN:       m.SatellitePrn,
		GPS_STATUS_FIELD_SATELLITE_USED:      m.SatelliteUsed,
		GPS_STATUS_FIELD_SATELLITE_ELEVATION: m.SatelliteElevation,
		GPS_STATUS_FIELD_SATELLITE_AZIMUTH:   m.SatelliteAzimuth,
		GPS_STATUS_FIELD_SATELLITE_SNR:       m.SatelliteSnr,
	}
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
	TimeBootMs uint32 `json:"TimeBootMs" ` // [ ms ] Timestamp (time since system boot).
	Xacc       int16  `json:"Xacc" `       // [ mG ] X acceleration
	Yacc       int16  `json:"Yacc" `       // [ mG ] Y acceleration
	Zacc       int16  `json:"Zacc" `       // [ mG ] Z acceleration
	Xgyro      int16  `json:"Xgyro" `      // [ mrad/s ] Angular speed around X axis
	Ygyro      int16  `json:"Ygyro" `      // [ mrad/s ] Angular speed around Y axis
	Zgyro      int16  `json:"Zgyro" `      // [ mrad/s ] Angular speed around Z axis
	Xmag       int16  `json:"Xmag" `       // [ mgauss ] X Magnetic field
	Ymag       int16  `json:"Ymag" `       // [ mgauss ] Y Magnetic field
	Zmag       int16  `json:"Zmag" `       // [ mgauss ] Z Magnetic field
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

const (
	SCALED_IMU_FIELD_TIME_BOOT_MS = "ScaledImu.TimeBootMs"
	SCALED_IMU_FIELD_XACC         = "ScaledImu.Xacc"
	SCALED_IMU_FIELD_YACC         = "ScaledImu.Yacc"
	SCALED_IMU_FIELD_ZACC         = "ScaledImu.Zacc"
	SCALED_IMU_FIELD_XGYRO        = "ScaledImu.Xgyro"
	SCALED_IMU_FIELD_YGYRO        = "ScaledImu.Ygyro"
	SCALED_IMU_FIELD_ZGYRO        = "ScaledImu.Zgyro"
	SCALED_IMU_FIELD_XMAG         = "ScaledImu.Xmag"
	SCALED_IMU_FIELD_YMAG         = "ScaledImu.Ymag"
	SCALED_IMU_FIELD_ZMAG         = "ScaledImu.Zmag"
)

// ToMap (generated function)
func (m *ScaledImu) Dict() map[string]interface{} {
	return map[string]interface{}{
		SCALED_IMU_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		SCALED_IMU_FIELD_XACC:         m.Xacc,
		SCALED_IMU_FIELD_YACC:         m.Yacc,
		SCALED_IMU_FIELD_ZACC:         m.Zacc,
		SCALED_IMU_FIELD_XGYRO:        m.Xgyro,
		SCALED_IMU_FIELD_YGYRO:        m.Ygyro,
		SCALED_IMU_FIELD_ZGYRO:        m.Zgyro,
		SCALED_IMU_FIELD_XMAG:         m.Xmag,
		SCALED_IMU_FIELD_YMAG:         m.Ymag,
		SCALED_IMU_FIELD_ZMAG:         m.Zmag,
	}
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
	TimeUsec uint64 `json:"TimeUsec" ` // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Xacc     int16  `json:"Xacc" `     // X acceleration (raw)
	Yacc     int16  `json:"Yacc" `     // Y acceleration (raw)
	Zacc     int16  `json:"Zacc" `     // Z acceleration (raw)
	Xgyro    int16  `json:"Xgyro" `    // Angular speed around X axis (raw)
	Ygyro    int16  `json:"Ygyro" `    // Angular speed around Y axis (raw)
	Zgyro    int16  `json:"Zgyro" `    // Angular speed around Z axis (raw)
	Xmag     int16  `json:"Xmag" `     // X Magnetic field (raw)
	Ymag     int16  `json:"Ymag" `     // Y Magnetic field (raw)
	Zmag     int16  `json:"Zmag" `     // Z Magnetic field (raw)
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

const (
	RAW_IMU_FIELD_TIME_USEC = "RawImu.TimeUsec"
	RAW_IMU_FIELD_XACC      = "RawImu.Xacc"
	RAW_IMU_FIELD_YACC      = "RawImu.Yacc"
	RAW_IMU_FIELD_ZACC      = "RawImu.Zacc"
	RAW_IMU_FIELD_XGYRO     = "RawImu.Xgyro"
	RAW_IMU_FIELD_YGYRO     = "RawImu.Ygyro"
	RAW_IMU_FIELD_ZGYRO     = "RawImu.Zgyro"
	RAW_IMU_FIELD_XMAG      = "RawImu.Xmag"
	RAW_IMU_FIELD_YMAG      = "RawImu.Ymag"
	RAW_IMU_FIELD_ZMAG      = "RawImu.Zmag"
)

// ToMap (generated function)
func (m *RawImu) Dict() map[string]interface{} {
	return map[string]interface{}{
		RAW_IMU_FIELD_TIME_USEC: m.TimeUsec,
		RAW_IMU_FIELD_XACC:      m.Xacc,
		RAW_IMU_FIELD_YACC:      m.Yacc,
		RAW_IMU_FIELD_ZACC:      m.Zacc,
		RAW_IMU_FIELD_XGYRO:     m.Xgyro,
		RAW_IMU_FIELD_YGYRO:     m.Ygyro,
		RAW_IMU_FIELD_ZGYRO:     m.Zgyro,
		RAW_IMU_FIELD_XMAG:      m.Xmag,
		RAW_IMU_FIELD_YMAG:      m.Ymag,
		RAW_IMU_FIELD_ZMAG:      m.Zmag,
	}
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
	TimeUsec    uint64 `json:"TimeUsec" `    // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	PressAbs    int16  `json:"PressAbs" `    // Absolute pressure (raw)
	PressDiff1  int16  `json:"PressDiff1" `  // Differential pressure 1 (raw, 0 if nonexistent)
	PressDiff2  int16  `json:"PressDiff2" `  // Differential pressure 2 (raw, 0 if nonexistent)
	Temperature int16  `json:"Temperature" ` // Raw Temperature measurement (raw)
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

const (
	RAW_PRESSURE_FIELD_TIME_USEC   = "RawPressure.TimeUsec"
	RAW_PRESSURE_FIELD_PRESS_ABS   = "RawPressure.PressAbs"
	RAW_PRESSURE_FIELD_PRESS_DIFF1 = "RawPressure.PressDiff1"
	RAW_PRESSURE_FIELD_PRESS_DIFF2 = "RawPressure.PressDiff2"
	RAW_PRESSURE_FIELD_TEMPERATURE = "RawPressure.Temperature"
)

// ToMap (generated function)
func (m *RawPressure) Dict() map[string]interface{} {
	return map[string]interface{}{
		RAW_PRESSURE_FIELD_TIME_USEC:   m.TimeUsec,
		RAW_PRESSURE_FIELD_PRESS_ABS:   m.PressAbs,
		RAW_PRESSURE_FIELD_PRESS_DIFF1: m.PressDiff1,
		RAW_PRESSURE_FIELD_PRESS_DIFF2: m.PressDiff2,
		RAW_PRESSURE_FIELD_TEMPERATURE: m.Temperature,
	}
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
	TimeBootMs  uint32  `json:"TimeBootMs" `  // [ ms ] Timestamp (time since system boot).
	PressAbs    float32 `json:"PressAbs" `    // [ hPa ] Absolute pressure
	PressDiff   float32 `json:"PressDiff" `   // [ hPa ] Differential pressure 1
	Temperature int16   `json:"Temperature" ` // [ cdegC ] Absolute pressure temperature
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

const (
	SCALED_PRESSURE_FIELD_TIME_BOOT_MS = "ScaledPressure.TimeBootMs"
	SCALED_PRESSURE_FIELD_PRESS_ABS    = "ScaledPressure.PressAbs"
	SCALED_PRESSURE_FIELD_PRESS_DIFF   = "ScaledPressure.PressDiff"
	SCALED_PRESSURE_FIELD_TEMPERATURE  = "ScaledPressure.Temperature"
)

// ToMap (generated function)
func (m *ScaledPressure) Dict() map[string]interface{} {
	return map[string]interface{}{
		SCALED_PRESSURE_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		SCALED_PRESSURE_FIELD_PRESS_ABS:    m.PressAbs,
		SCALED_PRESSURE_FIELD_PRESS_DIFF:   m.PressDiff,
		SCALED_PRESSURE_FIELD_TEMPERATURE:  m.Temperature,
	}
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
	TimeBootMs uint32  `json:"TimeBootMs" ` // [ ms ] Timestamp (time since system boot).
	Roll       float32 `json:"Roll" `       // [ rad ] Roll angle (-pi..+pi)
	Pitch      float32 `json:"Pitch" `      // [ rad ] Pitch angle (-pi..+pi)
	Yaw        float32 `json:"Yaw" `        // [ rad ] Yaw angle (-pi..+pi)
	Rollspeed  float32 `json:"Rollspeed" `  // [ rad/s ] Roll angular speed
	Pitchspeed float32 `json:"Pitchspeed" ` // [ rad/s ] Pitch angular speed
	Yawspeed   float32 `json:"Yawspeed" `   // [ rad/s ] Yaw angular speed
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

const (
	ATTITUDE_FIELD_TIME_BOOT_MS = "Attitude.TimeBootMs"
	ATTITUDE_FIELD_ROLL         = "Attitude.Roll"
	ATTITUDE_FIELD_PITCH        = "Attitude.Pitch"
	ATTITUDE_FIELD_YAW          = "Attitude.Yaw"
	ATTITUDE_FIELD_ROLLSPEED    = "Attitude.Rollspeed"
	ATTITUDE_FIELD_PITCHSPEED   = "Attitude.Pitchspeed"
	ATTITUDE_FIELD_YAWSPEED     = "Attitude.Yawspeed"
)

// ToMap (generated function)
func (m *Attitude) Dict() map[string]interface{} {
	return map[string]interface{}{
		ATTITUDE_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		ATTITUDE_FIELD_ROLL:         m.Roll,
		ATTITUDE_FIELD_PITCH:        m.Pitch,
		ATTITUDE_FIELD_YAW:          m.Yaw,
		ATTITUDE_FIELD_ROLLSPEED:    m.Rollspeed,
		ATTITUDE_FIELD_PITCHSPEED:   m.Pitchspeed,
		ATTITUDE_FIELD_YAWSPEED:     m.Yawspeed,
	}
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
	TimeBootMs uint32  `json:"TimeBootMs" ` // [ ms ] Timestamp (time since system boot).
	Q1         float32 `json:"Q1" `         // Quaternion component 1, w (1 in null-rotation)
	Q2         float32 `json:"Q2" `         // Quaternion component 2, x (0 in null-rotation)
	Q3         float32 `json:"Q3" `         // Quaternion component 3, y (0 in null-rotation)
	Q4         float32 `json:"Q4" `         // Quaternion component 4, z (0 in null-rotation)
	Rollspeed  float32 `json:"Rollspeed" `  // [ rad/s ] Roll angular speed
	Pitchspeed float32 `json:"Pitchspeed" ` // [ rad/s ] Pitch angular speed
	Yawspeed   float32 `json:"Yawspeed" `   // [ rad/s ] Yaw angular speed
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

const (
	ATTITUDE_QUATERNION_FIELD_TIME_BOOT_MS = "AttitudeQuaternion.TimeBootMs"
	ATTITUDE_QUATERNION_FIELD_Q1           = "AttitudeQuaternion.Q1"
	ATTITUDE_QUATERNION_FIELD_Q2           = "AttitudeQuaternion.Q2"
	ATTITUDE_QUATERNION_FIELD_Q3           = "AttitudeQuaternion.Q3"
	ATTITUDE_QUATERNION_FIELD_Q4           = "AttitudeQuaternion.Q4"
	ATTITUDE_QUATERNION_FIELD_ROLLSPEED    = "AttitudeQuaternion.Rollspeed"
	ATTITUDE_QUATERNION_FIELD_PITCHSPEED   = "AttitudeQuaternion.Pitchspeed"
	ATTITUDE_QUATERNION_FIELD_YAWSPEED     = "AttitudeQuaternion.Yawspeed"
)

// ToMap (generated function)
func (m *AttitudeQuaternion) Dict() map[string]interface{} {
	return map[string]interface{}{
		ATTITUDE_QUATERNION_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		ATTITUDE_QUATERNION_FIELD_Q1:           m.Q1,
		ATTITUDE_QUATERNION_FIELD_Q2:           m.Q2,
		ATTITUDE_QUATERNION_FIELD_Q3:           m.Q3,
		ATTITUDE_QUATERNION_FIELD_Q4:           m.Q4,
		ATTITUDE_QUATERNION_FIELD_ROLLSPEED:    m.Rollspeed,
		ATTITUDE_QUATERNION_FIELD_PITCHSPEED:   m.Pitchspeed,
		ATTITUDE_QUATERNION_FIELD_YAWSPEED:     m.Yawspeed,
	}
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
	TimeBootMs uint32  `json:"TimeBootMs" ` // [ ms ] Timestamp (time since system boot).
	X          float32 `json:"X" `          // [ m ] X Position
	Y          float32 `json:"Y" `          // [ m ] Y Position
	Z          float32 `json:"Z" `          // [ m ] Z Position
	Vx         float32 `json:"Vx" `         // [ m/s ] X Speed
	Vy         float32 `json:"Vy" `         // [ m/s ] Y Speed
	Vz         float32 `json:"Vz" `         // [ m/s ] Z Speed
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

const (
	LOCAL_POSITION_NED_FIELD_TIME_BOOT_MS = "LocalPositionNed.TimeBootMs"
	LOCAL_POSITION_NED_FIELD_X            = "LocalPositionNed.X"
	LOCAL_POSITION_NED_FIELD_Y            = "LocalPositionNed.Y"
	LOCAL_POSITION_NED_FIELD_Z            = "LocalPositionNed.Z"
	LOCAL_POSITION_NED_FIELD_VX           = "LocalPositionNed.Vx"
	LOCAL_POSITION_NED_FIELD_VY           = "LocalPositionNed.Vy"
	LOCAL_POSITION_NED_FIELD_VZ           = "LocalPositionNed.Vz"
)

// ToMap (generated function)
func (m *LocalPositionNed) Dict() map[string]interface{} {
	return map[string]interface{}{
		LOCAL_POSITION_NED_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		LOCAL_POSITION_NED_FIELD_X:            m.X,
		LOCAL_POSITION_NED_FIELD_Y:            m.Y,
		LOCAL_POSITION_NED_FIELD_Z:            m.Z,
		LOCAL_POSITION_NED_FIELD_VX:           m.Vx,
		LOCAL_POSITION_NED_FIELD_VY:           m.Vy,
		LOCAL_POSITION_NED_FIELD_VZ:           m.Vz,
	}
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
	TimeBootMs  uint32 `json:"TimeBootMs" `  // [ ms ] Timestamp (time since system boot).
	Lat         int32  `json:"Lat" `         // [ degE7 ] Latitude, expressed
	Lon         int32  `json:"Lon" `         // [ degE7 ] Longitude, expressed
	Alt         int32  `json:"Alt" `         // [ mm ] Altitude (MSL). Note that virtually all GPS modules provide both WGS84 and MSL.
	RelativeAlt int32  `json:"RelativeAlt" ` // [ mm ] Altitude above ground
	Vx          int16  `json:"Vx" `          // [ cm/s ] Ground X Speed (Latitude, positive north)
	Vy          int16  `json:"Vy" `          // [ cm/s ] Ground Y Speed (Longitude, positive east)
	Vz          int16  `json:"Vz" `          // [ cm/s ] Ground Z Speed (Altitude, positive down)
	Hdg         uint16 `json:"Hdg" `         // [ cdeg ] Vehicle heading (yaw angle), 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
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

const (
	GLOBAL_POSITION_INT_FIELD_TIME_BOOT_MS = "GlobalPositionInt.TimeBootMs"
	GLOBAL_POSITION_INT_FIELD_LAT          = "GlobalPositionInt.Lat"
	GLOBAL_POSITION_INT_FIELD_LON          = "GlobalPositionInt.Lon"
	GLOBAL_POSITION_INT_FIELD_ALT          = "GlobalPositionInt.Alt"
	GLOBAL_POSITION_INT_FIELD_RELATIVE_ALT = "GlobalPositionInt.RelativeAlt"
	GLOBAL_POSITION_INT_FIELD_VX           = "GlobalPositionInt.Vx"
	GLOBAL_POSITION_INT_FIELD_VY           = "GlobalPositionInt.Vy"
	GLOBAL_POSITION_INT_FIELD_VZ           = "GlobalPositionInt.Vz"
	GLOBAL_POSITION_INT_FIELD_HDG          = "GlobalPositionInt.Hdg"
)

// ToMap (generated function)
func (m *GlobalPositionInt) Dict() map[string]interface{} {
	return map[string]interface{}{
		GLOBAL_POSITION_INT_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		GLOBAL_POSITION_INT_FIELD_LAT:          m.Lat,
		GLOBAL_POSITION_INT_FIELD_LON:          m.Lon,
		GLOBAL_POSITION_INT_FIELD_ALT:          m.Alt,
		GLOBAL_POSITION_INT_FIELD_RELATIVE_ALT: m.RelativeAlt,
		GLOBAL_POSITION_INT_FIELD_VX:           m.Vx,
		GLOBAL_POSITION_INT_FIELD_VY:           m.Vy,
		GLOBAL_POSITION_INT_FIELD_VZ:           m.Vz,
		GLOBAL_POSITION_INT_FIELD_HDG:          m.Hdg,
	}
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
	TimeBootMs  uint32 `json:"TimeBootMs" `  // [ ms ] Timestamp (time since system boot).
	Chan1Scaled int16  `json:"Chan1Scaled" ` // RC channel 1 value scaled.
	Chan2Scaled int16  `json:"Chan2Scaled" ` // RC channel 2 value scaled.
	Chan3Scaled int16  `json:"Chan3Scaled" ` // RC channel 3 value scaled.
	Chan4Scaled int16  `json:"Chan4Scaled" ` // RC channel 4 value scaled.
	Chan5Scaled int16  `json:"Chan5Scaled" ` // RC channel 5 value scaled.
	Chan6Scaled int16  `json:"Chan6Scaled" ` // RC channel 6 value scaled.
	Chan7Scaled int16  `json:"Chan7Scaled" ` // RC channel 7 value scaled.
	Chan8Scaled int16  `json:"Chan8Scaled" ` // RC channel 8 value scaled.
	Port        uint8  `json:"Port" `        // Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
	Rssi        uint8  `json:"Rssi" `        // Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
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

const (
	RC_CHANNELS_SCALED_FIELD_TIME_BOOT_MS = "RcChannelsScaled.TimeBootMs"
	RC_CHANNELS_SCALED_FIELD_CHAN1_SCALED = "RcChannelsScaled.Chan1Scaled"
	RC_CHANNELS_SCALED_FIELD_CHAN2_SCALED = "RcChannelsScaled.Chan2Scaled"
	RC_CHANNELS_SCALED_FIELD_CHAN3_SCALED = "RcChannelsScaled.Chan3Scaled"
	RC_CHANNELS_SCALED_FIELD_CHAN4_SCALED = "RcChannelsScaled.Chan4Scaled"
	RC_CHANNELS_SCALED_FIELD_CHAN5_SCALED = "RcChannelsScaled.Chan5Scaled"
	RC_CHANNELS_SCALED_FIELD_CHAN6_SCALED = "RcChannelsScaled.Chan6Scaled"
	RC_CHANNELS_SCALED_FIELD_CHAN7_SCALED = "RcChannelsScaled.Chan7Scaled"
	RC_CHANNELS_SCALED_FIELD_CHAN8_SCALED = "RcChannelsScaled.Chan8Scaled"
	RC_CHANNELS_SCALED_FIELD_PORT         = "RcChannelsScaled.Port"
	RC_CHANNELS_SCALED_FIELD_RSSI         = "RcChannelsScaled.Rssi"
)

// ToMap (generated function)
func (m *RcChannelsScaled) Dict() map[string]interface{} {
	return map[string]interface{}{
		RC_CHANNELS_SCALED_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		RC_CHANNELS_SCALED_FIELD_CHAN1_SCALED: m.Chan1Scaled,
		RC_CHANNELS_SCALED_FIELD_CHAN2_SCALED: m.Chan2Scaled,
		RC_CHANNELS_SCALED_FIELD_CHAN3_SCALED: m.Chan3Scaled,
		RC_CHANNELS_SCALED_FIELD_CHAN4_SCALED: m.Chan4Scaled,
		RC_CHANNELS_SCALED_FIELD_CHAN5_SCALED: m.Chan5Scaled,
		RC_CHANNELS_SCALED_FIELD_CHAN6_SCALED: m.Chan6Scaled,
		RC_CHANNELS_SCALED_FIELD_CHAN7_SCALED: m.Chan7Scaled,
		RC_CHANNELS_SCALED_FIELD_CHAN8_SCALED: m.Chan8Scaled,
		RC_CHANNELS_SCALED_FIELD_PORT:         m.Port,
		RC_CHANNELS_SCALED_FIELD_RSSI:         m.Rssi,
	}
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
	TimeBootMs uint32 `json:"TimeBootMs" ` // [ ms ] Timestamp (time since system boot).
	Chan1Raw   uint16 `json:"Chan1Raw" `   // [ us ] RC channel 1 value.
	Chan2Raw   uint16 `json:"Chan2Raw" `   // [ us ] RC channel 2 value.
	Chan3Raw   uint16 `json:"Chan3Raw" `   // [ us ] RC channel 3 value.
	Chan4Raw   uint16 `json:"Chan4Raw" `   // [ us ] RC channel 4 value.
	Chan5Raw   uint16 `json:"Chan5Raw" `   // [ us ] RC channel 5 value.
	Chan6Raw   uint16 `json:"Chan6Raw" `   // [ us ] RC channel 6 value.
	Chan7Raw   uint16 `json:"Chan7Raw" `   // [ us ] RC channel 7 value.
	Chan8Raw   uint16 `json:"Chan8Raw" `   // [ us ] RC channel 8 value.
	Port       uint8  `json:"Port" `       // Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
	Rssi       uint8  `json:"Rssi" `       // Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
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

const (
	RC_CHANNELS_RAW_FIELD_TIME_BOOT_MS = "RcChannelsRaw.TimeBootMs"
	RC_CHANNELS_RAW_FIELD_CHAN1_RAW    = "RcChannelsRaw.Chan1Raw"
	RC_CHANNELS_RAW_FIELD_CHAN2_RAW    = "RcChannelsRaw.Chan2Raw"
	RC_CHANNELS_RAW_FIELD_CHAN3_RAW    = "RcChannelsRaw.Chan3Raw"
	RC_CHANNELS_RAW_FIELD_CHAN4_RAW    = "RcChannelsRaw.Chan4Raw"
	RC_CHANNELS_RAW_FIELD_CHAN5_RAW    = "RcChannelsRaw.Chan5Raw"
	RC_CHANNELS_RAW_FIELD_CHAN6_RAW    = "RcChannelsRaw.Chan6Raw"
	RC_CHANNELS_RAW_FIELD_CHAN7_RAW    = "RcChannelsRaw.Chan7Raw"
	RC_CHANNELS_RAW_FIELD_CHAN8_RAW    = "RcChannelsRaw.Chan8Raw"
	RC_CHANNELS_RAW_FIELD_PORT         = "RcChannelsRaw.Port"
	RC_CHANNELS_RAW_FIELD_RSSI         = "RcChannelsRaw.Rssi"
)

// ToMap (generated function)
func (m *RcChannelsRaw) Dict() map[string]interface{} {
	return map[string]interface{}{
		RC_CHANNELS_RAW_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		RC_CHANNELS_RAW_FIELD_CHAN1_RAW:    m.Chan1Raw,
		RC_CHANNELS_RAW_FIELD_CHAN2_RAW:    m.Chan2Raw,
		RC_CHANNELS_RAW_FIELD_CHAN3_RAW:    m.Chan3Raw,
		RC_CHANNELS_RAW_FIELD_CHAN4_RAW:    m.Chan4Raw,
		RC_CHANNELS_RAW_FIELD_CHAN5_RAW:    m.Chan5Raw,
		RC_CHANNELS_RAW_FIELD_CHAN6_RAW:    m.Chan6Raw,
		RC_CHANNELS_RAW_FIELD_CHAN7_RAW:    m.Chan7Raw,
		RC_CHANNELS_RAW_FIELD_CHAN8_RAW:    m.Chan8Raw,
		RC_CHANNELS_RAW_FIELD_PORT:         m.Port,
		RC_CHANNELS_RAW_FIELD_RSSI:         m.Rssi,
	}
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
	TimeUsec  uint32 `json:"TimeUsec" `  // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Servo1Raw uint16 `json:"Servo1Raw" ` // [ us ] Servo output 1 value
	Servo2Raw uint16 `json:"Servo2Raw" ` // [ us ] Servo output 2 value
	Servo3Raw uint16 `json:"Servo3Raw" ` // [ us ] Servo output 3 value
	Servo4Raw uint16 `json:"Servo4Raw" ` // [ us ] Servo output 4 value
	Servo5Raw uint16 `json:"Servo5Raw" ` // [ us ] Servo output 5 value
	Servo6Raw uint16 `json:"Servo6Raw" ` // [ us ] Servo output 6 value
	Servo7Raw uint16 `json:"Servo7Raw" ` // [ us ] Servo output 7 value
	Servo8Raw uint16 `json:"Servo8Raw" ` // [ us ] Servo output 8 value
	Port      uint8  `json:"Port" `      // Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
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

const (
	SERVO_OUTPUT_RAW_FIELD_TIME_USEC  = "ServoOutputRaw.TimeUsec"
	SERVO_OUTPUT_RAW_FIELD_SERVO1_RAW = "ServoOutputRaw.Servo1Raw"
	SERVO_OUTPUT_RAW_FIELD_SERVO2_RAW = "ServoOutputRaw.Servo2Raw"
	SERVO_OUTPUT_RAW_FIELD_SERVO3_RAW = "ServoOutputRaw.Servo3Raw"
	SERVO_OUTPUT_RAW_FIELD_SERVO4_RAW = "ServoOutputRaw.Servo4Raw"
	SERVO_OUTPUT_RAW_FIELD_SERVO5_RAW = "ServoOutputRaw.Servo5Raw"
	SERVO_OUTPUT_RAW_FIELD_SERVO6_RAW = "ServoOutputRaw.Servo6Raw"
	SERVO_OUTPUT_RAW_FIELD_SERVO7_RAW = "ServoOutputRaw.Servo7Raw"
	SERVO_OUTPUT_RAW_FIELD_SERVO8_RAW = "ServoOutputRaw.Servo8Raw"
	SERVO_OUTPUT_RAW_FIELD_PORT       = "ServoOutputRaw.Port"
)

// ToMap (generated function)
func (m *ServoOutputRaw) Dict() map[string]interface{} {
	return map[string]interface{}{
		SERVO_OUTPUT_RAW_FIELD_TIME_USEC:  m.TimeUsec,
		SERVO_OUTPUT_RAW_FIELD_SERVO1_RAW: m.Servo1Raw,
		SERVO_OUTPUT_RAW_FIELD_SERVO2_RAW: m.Servo2Raw,
		SERVO_OUTPUT_RAW_FIELD_SERVO3_RAW: m.Servo3Raw,
		SERVO_OUTPUT_RAW_FIELD_SERVO4_RAW: m.Servo4Raw,
		SERVO_OUTPUT_RAW_FIELD_SERVO5_RAW: m.Servo5Raw,
		SERVO_OUTPUT_RAW_FIELD_SERVO6_RAW: m.Servo6Raw,
		SERVO_OUTPUT_RAW_FIELD_SERVO7_RAW: m.Servo7Raw,
		SERVO_OUTPUT_RAW_FIELD_SERVO8_RAW: m.Servo8Raw,
		SERVO_OUTPUT_RAW_FIELD_PORT:       m.Port,
	}
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
	StartIndex      int16 `json:"StartIndex" `      // Start index
	EndIndex        int16 `json:"EndIndex" `        // End index, -1 by default (-1: send list to end). Else a valid index of the list
	TargetSystem    uint8 `json:"TargetSystem" `    // System ID
	TargetComponent uint8 `json:"TargetComponent" ` // Component ID
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

const (
	MISSION_REQUEST_PARTIAL_LIST_FIELD_START_INDEX      = "MissionRequestPartialList.StartIndex"
	MISSION_REQUEST_PARTIAL_LIST_FIELD_END_INDEX        = "MissionRequestPartialList.EndIndex"
	MISSION_REQUEST_PARTIAL_LIST_FIELD_TARGET_SYSTEM    = "MissionRequestPartialList.TargetSystem"
	MISSION_REQUEST_PARTIAL_LIST_FIELD_TARGET_COMPONENT = "MissionRequestPartialList.TargetComponent"
)

// ToMap (generated function)
func (m *MissionRequestPartialList) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_REQUEST_PARTIAL_LIST_FIELD_START_INDEX:      m.StartIndex,
		MISSION_REQUEST_PARTIAL_LIST_FIELD_END_INDEX:        m.EndIndex,
		MISSION_REQUEST_PARTIAL_LIST_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_REQUEST_PARTIAL_LIST_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	StartIndex      int16 `json:"StartIndex" `      // Start index. Must be smaller / equal to the largest index of the current onboard list.
	EndIndex        int16 `json:"EndIndex" `        // End index, equal or greater than start index.
	TargetSystem    uint8 `json:"TargetSystem" `    // System ID
	TargetComponent uint8 `json:"TargetComponent" ` // Component ID
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

const (
	MISSION_WRITE_PARTIAL_LIST_FIELD_START_INDEX      = "MissionWritePartialList.StartIndex"
	MISSION_WRITE_PARTIAL_LIST_FIELD_END_INDEX        = "MissionWritePartialList.EndIndex"
	MISSION_WRITE_PARTIAL_LIST_FIELD_TARGET_SYSTEM    = "MissionWritePartialList.TargetSystem"
	MISSION_WRITE_PARTIAL_LIST_FIELD_TARGET_COMPONENT = "MissionWritePartialList.TargetComponent"
)

// ToMap (generated function)
func (m *MissionWritePartialList) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_WRITE_PARTIAL_LIST_FIELD_START_INDEX:      m.StartIndex,
		MISSION_WRITE_PARTIAL_LIST_FIELD_END_INDEX:        m.EndIndex,
		MISSION_WRITE_PARTIAL_LIST_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_WRITE_PARTIAL_LIST_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	Param1          float32   `json:"Param1" `          // PARAM1, see MAV_CMD enum
	Param2          float32   `json:"Param2" `          // PARAM2, see MAV_CMD enum
	Param3          float32   `json:"Param3" `          // PARAM3, see MAV_CMD enum
	Param4          float32   `json:"Param4" `          // PARAM4, see MAV_CMD enum
	X               float32   `json:"X" `               // PARAM5 / local: X coordinate, global: latitude
	Y               float32   `json:"Y" `               // PARAM6 / local: Y coordinate, global: longitude
	Z               float32   `json:"Z" `               // PARAM7 / local: Z coordinate, global: altitude (relative or absolute, depending on frame).
	Seq             uint16    `json:"Seq" `             // Sequence
	Command         MAV_CMD   `json:"Command" `         // The scheduled action for the waypoint.
	TargetSystem    uint8     `json:"TargetSystem" `    // System ID
	TargetComponent uint8     `json:"TargetComponent" ` // Component ID
	Frame           MAV_FRAME `json:"Frame" `           // The coordinate system of the waypoint.
	Current         uint8     `json:"Current" `         // false:0, true:1
	Autocontinue    uint8     `json:"Autocontinue" `    // Autocontinue to next waypoint
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

const (
	MISSION_ITEM_FIELD_PARAM1           = "MissionItem.Param1"
	MISSION_ITEM_FIELD_PARAM2           = "MissionItem.Param2"
	MISSION_ITEM_FIELD_PARAM3           = "MissionItem.Param3"
	MISSION_ITEM_FIELD_PARAM4           = "MissionItem.Param4"
	MISSION_ITEM_FIELD_X                = "MissionItem.X"
	MISSION_ITEM_FIELD_Y                = "MissionItem.Y"
	MISSION_ITEM_FIELD_Z                = "MissionItem.Z"
	MISSION_ITEM_FIELD_SEQ              = "MissionItem.Seq"
	MISSION_ITEM_FIELD_COMMAND          = "MissionItem.Command"
	MISSION_ITEM_FIELD_TARGET_SYSTEM    = "MissionItem.TargetSystem"
	MISSION_ITEM_FIELD_TARGET_COMPONENT = "MissionItem.TargetComponent"
	MISSION_ITEM_FIELD_FRAME            = "MissionItem.Frame"
	MISSION_ITEM_FIELD_CURRENT          = "MissionItem.Current"
	MISSION_ITEM_FIELD_AUTOCONTINUE     = "MissionItem.Autocontinue"
)

// ToMap (generated function)
func (m *MissionItem) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_ITEM_FIELD_PARAM1:           m.Param1,
		MISSION_ITEM_FIELD_PARAM2:           m.Param2,
		MISSION_ITEM_FIELD_PARAM3:           m.Param3,
		MISSION_ITEM_FIELD_PARAM4:           m.Param4,
		MISSION_ITEM_FIELD_X:                m.X,
		MISSION_ITEM_FIELD_Y:                m.Y,
		MISSION_ITEM_FIELD_Z:                m.Z,
		MISSION_ITEM_FIELD_SEQ:              m.Seq,
		MISSION_ITEM_FIELD_COMMAND:          m.Command,
		MISSION_ITEM_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_ITEM_FIELD_TARGET_COMPONENT: m.TargetComponent,
		MISSION_ITEM_FIELD_FRAME:            m.Frame,
		MISSION_ITEM_FIELD_CURRENT:          m.Current,
		MISSION_ITEM_FIELD_AUTOCONTINUE:     m.Autocontinue,
	}
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
	Seq             uint16 `json:"Seq" `             // Sequence
	TargetSystem    uint8  `json:"TargetSystem" `    // System ID
	TargetComponent uint8  `json:"TargetComponent" ` // Component ID
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

const (
	MISSION_REQUEST_FIELD_SEQ              = "MissionRequest.Seq"
	MISSION_REQUEST_FIELD_TARGET_SYSTEM    = "MissionRequest.TargetSystem"
	MISSION_REQUEST_FIELD_TARGET_COMPONENT = "MissionRequest.TargetComponent"
)

// ToMap (generated function)
func (m *MissionRequest) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_REQUEST_FIELD_SEQ:              m.Seq,
		MISSION_REQUEST_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_REQUEST_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	Seq             uint16 `json:"Seq" `             // Sequence
	TargetSystem    uint8  `json:"TargetSystem" `    // System ID
	TargetComponent uint8  `json:"TargetComponent" ` // Component ID
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

const (
	MISSION_SET_CURRENT_FIELD_SEQ              = "MissionSetCurrent.Seq"
	MISSION_SET_CURRENT_FIELD_TARGET_SYSTEM    = "MissionSetCurrent.TargetSystem"
	MISSION_SET_CURRENT_FIELD_TARGET_COMPONENT = "MissionSetCurrent.TargetComponent"
)

// ToMap (generated function)
func (m *MissionSetCurrent) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_SET_CURRENT_FIELD_SEQ:              m.Seq,
		MISSION_SET_CURRENT_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_SET_CURRENT_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	Seq uint16 `json:"Seq" ` // Sequence
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

const (
	MISSION_CURRENT_FIELD_SEQ = "MissionCurrent.Seq"
)

// ToMap (generated function)
func (m *MissionCurrent) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_CURRENT_FIELD_SEQ: m.Seq,
	}
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
	TargetSystem    uint8 `json:"TargetSystem" `    // System ID
	TargetComponent uint8 `json:"TargetComponent" ` // Component ID
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

const (
	MISSION_REQUEST_LIST_FIELD_TARGET_SYSTEM    = "MissionRequestList.TargetSystem"
	MISSION_REQUEST_LIST_FIELD_TARGET_COMPONENT = "MissionRequestList.TargetComponent"
)

// ToMap (generated function)
func (m *MissionRequestList) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_REQUEST_LIST_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_REQUEST_LIST_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	Count           uint16 `json:"Count" `           // Number of mission items in the sequence
	TargetSystem    uint8  `json:"TargetSystem" `    // System ID
	TargetComponent uint8  `json:"TargetComponent" ` // Component ID
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

const (
	MISSION_COUNT_FIELD_COUNT            = "MissionCount.Count"
	MISSION_COUNT_FIELD_TARGET_SYSTEM    = "MissionCount.TargetSystem"
	MISSION_COUNT_FIELD_TARGET_COMPONENT = "MissionCount.TargetComponent"
)

// ToMap (generated function)
func (m *MissionCount) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_COUNT_FIELD_COUNT:            m.Count,
		MISSION_COUNT_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_COUNT_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	TargetSystem    uint8 `json:"TargetSystem" `    // System ID
	TargetComponent uint8 `json:"TargetComponent" ` // Component ID
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

const (
	MISSION_CLEAR_ALL_FIELD_TARGET_SYSTEM    = "MissionClearAll.TargetSystem"
	MISSION_CLEAR_ALL_FIELD_TARGET_COMPONENT = "MissionClearAll.TargetComponent"
)

// ToMap (generated function)
func (m *MissionClearAll) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_CLEAR_ALL_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_CLEAR_ALL_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	Seq uint16 `json:"Seq" ` // Sequence
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

const (
	MISSION_ITEM_REACHED_FIELD_SEQ = "MissionItemReached.Seq"
)

// ToMap (generated function)
func (m *MissionItemReached) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_ITEM_REACHED_FIELD_SEQ: m.Seq,
	}
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
	TargetSystem    uint8              `json:"TargetSystem" `    // System ID
	TargetComponent uint8              `json:"TargetComponent" ` // Component ID
	Type            MAV_MISSION_RESULT `json:"Type" `            // Mission result.
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

const (
	MISSION_ACK_FIELD_TARGET_SYSTEM    = "MissionAck.TargetSystem"
	MISSION_ACK_FIELD_TARGET_COMPONENT = "MissionAck.TargetComponent"
	MISSION_ACK_FIELD_TYPE             = "MissionAck.Type"
)

// ToMap (generated function)
func (m *MissionAck) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_ACK_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_ACK_FIELD_TARGET_COMPONENT: m.TargetComponent,
		MISSION_ACK_FIELD_TYPE:             m.Type,
	}
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
	Latitude     int32 `json:"Latitude" `     // [ degE7 ] Latitude (WGS84)
	Longitude    int32 `json:"Longitude" `    // [ degE7 ] Longitude (WGS84)
	Altitude     int32 `json:"Altitude" `     // [ mm ] Altitude (MSL). Positive for up.
	TargetSystem uint8 `json:"TargetSystem" ` // System ID
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

const (
	SET_GPS_GLOBAL_ORIGIN_FIELD_LATITUDE      = "SetGpsGlobalOrigin.Latitude"
	SET_GPS_GLOBAL_ORIGIN_FIELD_LONGITUDE     = "SetGpsGlobalOrigin.Longitude"
	SET_GPS_GLOBAL_ORIGIN_FIELD_ALTITUDE      = "SetGpsGlobalOrigin.Altitude"
	SET_GPS_GLOBAL_ORIGIN_FIELD_TARGET_SYSTEM = "SetGpsGlobalOrigin.TargetSystem"
)

// ToMap (generated function)
func (m *SetGpsGlobalOrigin) Dict() map[string]interface{} {
	return map[string]interface{}{
		SET_GPS_GLOBAL_ORIGIN_FIELD_LATITUDE:      m.Latitude,
		SET_GPS_GLOBAL_ORIGIN_FIELD_LONGITUDE:     m.Longitude,
		SET_GPS_GLOBAL_ORIGIN_FIELD_ALTITUDE:      m.Altitude,
		SET_GPS_GLOBAL_ORIGIN_FIELD_TARGET_SYSTEM: m.TargetSystem,
	}
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
	Latitude  int32 `json:"Latitude" `  // [ degE7 ] Latitude (WGS84)
	Longitude int32 `json:"Longitude" ` // [ degE7 ] Longitude (WGS84)
	Altitude  int32 `json:"Altitude" `  // [ mm ] Altitude (MSL). Positive for up.
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

const (
	GPS_GLOBAL_ORIGIN_FIELD_LATITUDE  = "GpsGlobalOrigin.Latitude"
	GPS_GLOBAL_ORIGIN_FIELD_LONGITUDE = "GpsGlobalOrigin.Longitude"
	GPS_GLOBAL_ORIGIN_FIELD_ALTITUDE  = "GpsGlobalOrigin.Altitude"
)

// ToMap (generated function)
func (m *GpsGlobalOrigin) Dict() map[string]interface{} {
	return map[string]interface{}{
		GPS_GLOBAL_ORIGIN_FIELD_LATITUDE:  m.Latitude,
		GPS_GLOBAL_ORIGIN_FIELD_LONGITUDE: m.Longitude,
		GPS_GLOBAL_ORIGIN_FIELD_ALTITUDE:  m.Altitude,
	}
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
	ParamValue0             float32 `json:"ParamValue0" `             // Initial parameter value
	Scale                   float32 `json:"Scale" `                   // Scale, maps the RC range [-1, 1] to a parameter value
	ParamValueMin           float32 `json:"ParamValueMin" `           // Minimum param value. The protocol does not define if this overwrites an onboard minimum value. (Depends on implementation)
	ParamValueMax           float32 `json:"ParamValueMax" `           // Maximum param value. The protocol does not define if this overwrites an onboard maximum value. (Depends on implementation)
	ParamIndex              int16   `json:"ParamIndex" `              // Parameter index. Send -1 to use the param ID field as identifier (else the param id will be ignored), send -2 to disable any existing map for this rc_channel_index.
	TargetSystem            uint8   `json:"TargetSystem" `            // System ID
	TargetComponent         uint8   `json:"TargetComponent" `         // Component ID
	ParamID                 string  `json:"ParamID" len:"16" `        // Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParameterRcChannelIndex uint8   `json:"ParameterRcChannelIndex" ` // Index of parameter RC channel. Not equal to the RC channel id. Typically corresponds to a potentiometer-knob on the RC.
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
		strings.TrimRight(m.ParamID, string(byte(0))),
		m.ParameterRcChannelIndex,
	)
}

const (
	PARAM_MAP_RC_FIELD_PARAM_VALUE0               = "ParamMapRc.ParamValue0"
	PARAM_MAP_RC_FIELD_SCALE                      = "ParamMapRc.Scale"
	PARAM_MAP_RC_FIELD_PARAM_VALUE_MIN            = "ParamMapRc.ParamValueMin"
	PARAM_MAP_RC_FIELD_PARAM_VALUE_MAX            = "ParamMapRc.ParamValueMax"
	PARAM_MAP_RC_FIELD_PARAM_INDEX                = "ParamMapRc.ParamIndex"
	PARAM_MAP_RC_FIELD_TARGET_SYSTEM              = "ParamMapRc.TargetSystem"
	PARAM_MAP_RC_FIELD_TARGET_COMPONENT           = "ParamMapRc.TargetComponent"
	PARAM_MAP_RC_FIELD_PARAM_ID                   = "ParamMapRc.ParamID"
	PARAM_MAP_RC_FIELD_PARAMETER_RC_CHANNEL_INDEX = "ParamMapRc.ParameterRcChannelIndex"
)

// ToMap (generated function)
func (m *ParamMapRc) Dict() map[string]interface{} {
	return map[string]interface{}{
		PARAM_MAP_RC_FIELD_PARAM_VALUE0:               m.ParamValue0,
		PARAM_MAP_RC_FIELD_SCALE:                      m.Scale,
		PARAM_MAP_RC_FIELD_PARAM_VALUE_MIN:            m.ParamValueMin,
		PARAM_MAP_RC_FIELD_PARAM_VALUE_MAX:            m.ParamValueMax,
		PARAM_MAP_RC_FIELD_PARAM_INDEX:                m.ParamIndex,
		PARAM_MAP_RC_FIELD_TARGET_SYSTEM:              m.TargetSystem,
		PARAM_MAP_RC_FIELD_TARGET_COMPONENT:           m.TargetComponent,
		PARAM_MAP_RC_FIELD_PARAM_ID:                   m.ParamID,
		PARAM_MAP_RC_FIELD_PARAMETER_RC_CHANNEL_INDEX: m.ParameterRcChannelIndex,
	}
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
	Seq             uint16 `json:"Seq" `             // Sequence
	TargetSystem    uint8  `json:"TargetSystem" `    // System ID
	TargetComponent uint8  `json:"TargetComponent" ` // Component ID
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

const (
	MISSION_REQUEST_INT_FIELD_SEQ              = "MissionRequestInt.Seq"
	MISSION_REQUEST_INT_FIELD_TARGET_SYSTEM    = "MissionRequestInt.TargetSystem"
	MISSION_REQUEST_INT_FIELD_TARGET_COMPONENT = "MissionRequestInt.TargetComponent"
)

// ToMap (generated function)
func (m *MissionRequestInt) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_REQUEST_INT_FIELD_SEQ:              m.Seq,
		MISSION_REQUEST_INT_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_REQUEST_INT_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	StartIndex   int16            `json:"StartIndex" `   // Start index for partial mission change (-1 for all items).
	EndIndex     int16            `json:"EndIndex" `     // End index of a partial mission change. -1 is a synonym for the last mission item (i.e. selects all items from start_index). Ignore field if start_index=-1.
	OriginSysid  uint8            `json:"OriginSysid" `  // System ID of the author of the new mission.
	OriginCompid MAV_COMPONENT    `json:"OriginCompid" ` // Compnent ID of the author of the new mission.
	MissionType  MAV_MISSION_TYPE `json:"MissionType" `  // Mission type.
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

const (
	MISSION_CHANGED_FIELD_START_INDEX   = "MissionChanged.StartIndex"
	MISSION_CHANGED_FIELD_END_INDEX     = "MissionChanged.EndIndex"
	MISSION_CHANGED_FIELD_ORIGIN_SYSID  = "MissionChanged.OriginSysid"
	MISSION_CHANGED_FIELD_ORIGIN_COMPID = "MissionChanged.OriginCompid"
	MISSION_CHANGED_FIELD_MISSION_TYPE  = "MissionChanged.MissionType"
)

// ToMap (generated function)
func (m *MissionChanged) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_CHANGED_FIELD_START_INDEX:   m.StartIndex,
		MISSION_CHANGED_FIELD_END_INDEX:     m.EndIndex,
		MISSION_CHANGED_FIELD_ORIGIN_SYSID:  m.OriginSysid,
		MISSION_CHANGED_FIELD_ORIGIN_COMPID: m.OriginCompid,
		MISSION_CHANGED_FIELD_MISSION_TYPE:  m.MissionType,
	}
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
	P1x             float32   `json:"P1x" `             // [ m ] x position 1 / Latitude 1
	P1y             float32   `json:"P1y" `             // [ m ] y position 1 / Longitude 1
	P1z             float32   `json:"P1z" `             // [ m ] z position 1 / Altitude 1
	P2x             float32   `json:"P2x" `             // [ m ] x position 2 / Latitude 2
	P2y             float32   `json:"P2y" `             // [ m ] y position 2 / Longitude 2
	P2z             float32   `json:"P2z" `             // [ m ] z position 2 / Altitude 2
	TargetSystem    uint8     `json:"TargetSystem" `    // System ID
	TargetComponent uint8     `json:"TargetComponent" ` // Component ID
	Frame           MAV_FRAME `json:"Frame" `           // Coordinate frame. Can be either global, GPS, right-handed with Z axis up or local, right handed, Z axis down.
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

const (
	SAFETY_SET_ALLOWED_AREA_FIELD_P1X              = "SafetySetAllowedArea.P1x"
	SAFETY_SET_ALLOWED_AREA_FIELD_P1Y              = "SafetySetAllowedArea.P1y"
	SAFETY_SET_ALLOWED_AREA_FIELD_P1Z              = "SafetySetAllowedArea.P1z"
	SAFETY_SET_ALLOWED_AREA_FIELD_P2X              = "SafetySetAllowedArea.P2x"
	SAFETY_SET_ALLOWED_AREA_FIELD_P2Y              = "SafetySetAllowedArea.P2y"
	SAFETY_SET_ALLOWED_AREA_FIELD_P2Z              = "SafetySetAllowedArea.P2z"
	SAFETY_SET_ALLOWED_AREA_FIELD_TARGET_SYSTEM    = "SafetySetAllowedArea.TargetSystem"
	SAFETY_SET_ALLOWED_AREA_FIELD_TARGET_COMPONENT = "SafetySetAllowedArea.TargetComponent"
	SAFETY_SET_ALLOWED_AREA_FIELD_FRAME            = "SafetySetAllowedArea.Frame"
)

// ToMap (generated function)
func (m *SafetySetAllowedArea) Dict() map[string]interface{} {
	return map[string]interface{}{
		SAFETY_SET_ALLOWED_AREA_FIELD_P1X:              m.P1x,
		SAFETY_SET_ALLOWED_AREA_FIELD_P1Y:              m.P1y,
		SAFETY_SET_ALLOWED_AREA_FIELD_P1Z:              m.P1z,
		SAFETY_SET_ALLOWED_AREA_FIELD_P2X:              m.P2x,
		SAFETY_SET_ALLOWED_AREA_FIELD_P2Y:              m.P2y,
		SAFETY_SET_ALLOWED_AREA_FIELD_P2Z:              m.P2z,
		SAFETY_SET_ALLOWED_AREA_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		SAFETY_SET_ALLOWED_AREA_FIELD_TARGET_COMPONENT: m.TargetComponent,
		SAFETY_SET_ALLOWED_AREA_FIELD_FRAME:            m.Frame,
	}
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
	P1x   float32   `json:"P1x" `   // [ m ] x position 1 / Latitude 1
	P1y   float32   `json:"P1y" `   // [ m ] y position 1 / Longitude 1
	P1z   float32   `json:"P1z" `   // [ m ] z position 1 / Altitude 1
	P2x   float32   `json:"P2x" `   // [ m ] x position 2 / Latitude 2
	P2y   float32   `json:"P2y" `   // [ m ] y position 2 / Longitude 2
	P2z   float32   `json:"P2z" `   // [ m ] z position 2 / Altitude 2
	Frame MAV_FRAME `json:"Frame" ` // Coordinate frame. Can be either global, GPS, right-handed with Z axis up or local, right handed, Z axis down.
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

const (
	SAFETY_ALLOWED_AREA_FIELD_P1X   = "SafetyAllowedArea.P1x"
	SAFETY_ALLOWED_AREA_FIELD_P1Y   = "SafetyAllowedArea.P1y"
	SAFETY_ALLOWED_AREA_FIELD_P1Z   = "SafetyAllowedArea.P1z"
	SAFETY_ALLOWED_AREA_FIELD_P2X   = "SafetyAllowedArea.P2x"
	SAFETY_ALLOWED_AREA_FIELD_P2Y   = "SafetyAllowedArea.P2y"
	SAFETY_ALLOWED_AREA_FIELD_P2Z   = "SafetyAllowedArea.P2z"
	SAFETY_ALLOWED_AREA_FIELD_FRAME = "SafetyAllowedArea.Frame"
)

// ToMap (generated function)
func (m *SafetyAllowedArea) Dict() map[string]interface{} {
	return map[string]interface{}{
		SAFETY_ALLOWED_AREA_FIELD_P1X:   m.P1x,
		SAFETY_ALLOWED_AREA_FIELD_P1Y:   m.P1y,
		SAFETY_ALLOWED_AREA_FIELD_P1Z:   m.P1z,
		SAFETY_ALLOWED_AREA_FIELD_P2X:   m.P2x,
		SAFETY_ALLOWED_AREA_FIELD_P2Y:   m.P2y,
		SAFETY_ALLOWED_AREA_FIELD_P2Z:   m.P2z,
		SAFETY_ALLOWED_AREA_FIELD_FRAME: m.Frame,
	}
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
	TimeUsec   uint64    `json:"TimeUsec" `           // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Q          []float32 `json:"Q" len:"4" `          // Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation)
	Rollspeed  float32   `json:"Rollspeed" `          // [ rad/s ] Roll angular speed
	Pitchspeed float32   `json:"Pitchspeed" `         // [ rad/s ] Pitch angular speed
	Yawspeed   float32   `json:"Yawspeed" `           // [ rad/s ] Yaw angular speed
	Covariance []float32 `json:"Covariance" len:"9" ` // Row-major representation of a 3x3 attitude covariance matrix (states: roll, pitch, yaw; first three entries are the first ROW, next three entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
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

const (
	ATTITUDE_QUATERNION_COV_FIELD_TIME_USEC  = "AttitudeQuaternionCov.TimeUsec"
	ATTITUDE_QUATERNION_COV_FIELD_Q          = "AttitudeQuaternionCov.Q"
	ATTITUDE_QUATERNION_COV_FIELD_ROLLSPEED  = "AttitudeQuaternionCov.Rollspeed"
	ATTITUDE_QUATERNION_COV_FIELD_PITCHSPEED = "AttitudeQuaternionCov.Pitchspeed"
	ATTITUDE_QUATERNION_COV_FIELD_YAWSPEED   = "AttitudeQuaternionCov.Yawspeed"
	ATTITUDE_QUATERNION_COV_FIELD_COVARIANCE = "AttitudeQuaternionCov.Covariance"
)

// ToMap (generated function)
func (m *AttitudeQuaternionCov) Dict() map[string]interface{} {
	return map[string]interface{}{
		ATTITUDE_QUATERNION_COV_FIELD_TIME_USEC:  m.TimeUsec,
		ATTITUDE_QUATERNION_COV_FIELD_Q:          m.Q,
		ATTITUDE_QUATERNION_COV_FIELD_ROLLSPEED:  m.Rollspeed,
		ATTITUDE_QUATERNION_COV_FIELD_PITCHSPEED: m.Pitchspeed,
		ATTITUDE_QUATERNION_COV_FIELD_YAWSPEED:   m.Yawspeed,
		ATTITUDE_QUATERNION_COV_FIELD_COVARIANCE: m.Covariance,
	}
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
	NavRoll       float32 `json:"NavRoll" `       // [ deg ] Current desired roll
	NavPitch      float32 `json:"NavPitch" `      // [ deg ] Current desired pitch
	AltError      float32 `json:"AltError" `      // [ m ] Current altitude error
	AspdError     float32 `json:"AspdError" `     // [ m/s ] Current airspeed error
	XtrackError   float32 `json:"XtrackError" `   // [ m ] Current crosstrack error on x-y plane
	NavBearing    int16   `json:"NavBearing" `    // [ deg ] Current desired heading
	TargetBearing int16   `json:"TargetBearing" ` // [ deg ] Bearing to current waypoint/target
	WpDist        uint16  `json:"WpDist" `        // [ m ] Distance to active waypoint
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

const (
	NAV_CONTROLLER_OUTPUT_FIELD_NAV_ROLL       = "NavControllerOutput.NavRoll"
	NAV_CONTROLLER_OUTPUT_FIELD_NAV_PITCH      = "NavControllerOutput.NavPitch"
	NAV_CONTROLLER_OUTPUT_FIELD_ALT_ERROR      = "NavControllerOutput.AltError"
	NAV_CONTROLLER_OUTPUT_FIELD_ASPD_ERROR     = "NavControllerOutput.AspdError"
	NAV_CONTROLLER_OUTPUT_FIELD_XTRACK_ERROR   = "NavControllerOutput.XtrackError"
	NAV_CONTROLLER_OUTPUT_FIELD_NAV_BEARING    = "NavControllerOutput.NavBearing"
	NAV_CONTROLLER_OUTPUT_FIELD_TARGET_BEARING = "NavControllerOutput.TargetBearing"
	NAV_CONTROLLER_OUTPUT_FIELD_WP_DIST        = "NavControllerOutput.WpDist"
)

// ToMap (generated function)
func (m *NavControllerOutput) Dict() map[string]interface{} {
	return map[string]interface{}{
		NAV_CONTROLLER_OUTPUT_FIELD_NAV_ROLL:       m.NavRoll,
		NAV_CONTROLLER_OUTPUT_FIELD_NAV_PITCH:      m.NavPitch,
		NAV_CONTROLLER_OUTPUT_FIELD_ALT_ERROR:      m.AltError,
		NAV_CONTROLLER_OUTPUT_FIELD_ASPD_ERROR:     m.AspdError,
		NAV_CONTROLLER_OUTPUT_FIELD_XTRACK_ERROR:   m.XtrackError,
		NAV_CONTROLLER_OUTPUT_FIELD_NAV_BEARING:    m.NavBearing,
		NAV_CONTROLLER_OUTPUT_FIELD_TARGET_BEARING: m.TargetBearing,
		NAV_CONTROLLER_OUTPUT_FIELD_WP_DIST:        m.WpDist,
	}
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
	TimeUsec      uint64             `json:"TimeUsec" `            // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Lat           int32              `json:"Lat" `                 // [ degE7 ] Latitude
	Lon           int32              `json:"Lon" `                 // [ degE7 ] Longitude
	Alt           int32              `json:"Alt" `                 // [ mm ] Altitude in meters above MSL
	RelativeAlt   int32              `json:"RelativeAlt" `         // [ mm ] Altitude above ground
	Vx            float32            `json:"Vx" `                  // [ m/s ] Ground X Speed (Latitude)
	Vy            float32            `json:"Vy" `                  // [ m/s ] Ground Y Speed (Longitude)
	Vz            float32            `json:"Vz" `                  // [ m/s ] Ground Z Speed (Altitude)
	Covariance    []float32          `json:"Covariance" len:"36" ` // Row-major representation of a 6x6 position and velocity 6x6 cross-covariance matrix (states: lat, lon, alt, vx, vy, vz; first six entries are the first ROW, next six entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
	EstimatorType MAV_ESTIMATOR_TYPE `json:"EstimatorType" `       // Class id of the estimator this estimate originated from.
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

const (
	GLOBAL_POSITION_INT_COV_FIELD_TIME_USEC      = "GlobalPositionIntCov.TimeUsec"
	GLOBAL_POSITION_INT_COV_FIELD_LAT            = "GlobalPositionIntCov.Lat"
	GLOBAL_POSITION_INT_COV_FIELD_LON            = "GlobalPositionIntCov.Lon"
	GLOBAL_POSITION_INT_COV_FIELD_ALT            = "GlobalPositionIntCov.Alt"
	GLOBAL_POSITION_INT_COV_FIELD_RELATIVE_ALT   = "GlobalPositionIntCov.RelativeAlt"
	GLOBAL_POSITION_INT_COV_FIELD_VX             = "GlobalPositionIntCov.Vx"
	GLOBAL_POSITION_INT_COV_FIELD_VY             = "GlobalPositionIntCov.Vy"
	GLOBAL_POSITION_INT_COV_FIELD_VZ             = "GlobalPositionIntCov.Vz"
	GLOBAL_POSITION_INT_COV_FIELD_COVARIANCE     = "GlobalPositionIntCov.Covariance"
	GLOBAL_POSITION_INT_COV_FIELD_ESTIMATOR_TYPE = "GlobalPositionIntCov.EstimatorType"
)

// ToMap (generated function)
func (m *GlobalPositionIntCov) Dict() map[string]interface{} {
	return map[string]interface{}{
		GLOBAL_POSITION_INT_COV_FIELD_TIME_USEC:      m.TimeUsec,
		GLOBAL_POSITION_INT_COV_FIELD_LAT:            m.Lat,
		GLOBAL_POSITION_INT_COV_FIELD_LON:            m.Lon,
		GLOBAL_POSITION_INT_COV_FIELD_ALT:            m.Alt,
		GLOBAL_POSITION_INT_COV_FIELD_RELATIVE_ALT:   m.RelativeAlt,
		GLOBAL_POSITION_INT_COV_FIELD_VX:             m.Vx,
		GLOBAL_POSITION_INT_COV_FIELD_VY:             m.Vy,
		GLOBAL_POSITION_INT_COV_FIELD_VZ:             m.Vz,
		GLOBAL_POSITION_INT_COV_FIELD_COVARIANCE:     m.Covariance,
		GLOBAL_POSITION_INT_COV_FIELD_ESTIMATOR_TYPE: m.EstimatorType,
	}
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
	TimeUsec      uint64             `json:"TimeUsec" `            // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	X             float32            `json:"X" `                   // [ m ] X Position
	Y             float32            `json:"Y" `                   // [ m ] Y Position
	Z             float32            `json:"Z" `                   // [ m ] Z Position
	Vx            float32            `json:"Vx" `                  // [ m/s ] X Speed
	Vy            float32            `json:"Vy" `                  // [ m/s ] Y Speed
	Vz            float32            `json:"Vz" `                  // [ m/s ] Z Speed
	Ax            float32            `json:"Ax" `                  // [ m/s/s ] X Acceleration
	Ay            float32            `json:"Ay" `                  // [ m/s/s ] Y Acceleration
	Az            float32            `json:"Az" `                  // [ m/s/s ] Z Acceleration
	Covariance    []float32          `json:"Covariance" len:"45" ` // Row-major representation of position, velocity and acceleration 9x9 cross-covariance matrix upper right triangle (states: x, y, z, vx, vy, vz, ax, ay, az; first nine entries are the first ROW, next eight entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
	EstimatorType MAV_ESTIMATOR_TYPE `json:"EstimatorType" `       // Class id of the estimator this estimate originated from.
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

const (
	LOCAL_POSITION_NED_COV_FIELD_TIME_USEC      = "LocalPositionNedCov.TimeUsec"
	LOCAL_POSITION_NED_COV_FIELD_X              = "LocalPositionNedCov.X"
	LOCAL_POSITION_NED_COV_FIELD_Y              = "LocalPositionNedCov.Y"
	LOCAL_POSITION_NED_COV_FIELD_Z              = "LocalPositionNedCov.Z"
	LOCAL_POSITION_NED_COV_FIELD_VX             = "LocalPositionNedCov.Vx"
	LOCAL_POSITION_NED_COV_FIELD_VY             = "LocalPositionNedCov.Vy"
	LOCAL_POSITION_NED_COV_FIELD_VZ             = "LocalPositionNedCov.Vz"
	LOCAL_POSITION_NED_COV_FIELD_AX             = "LocalPositionNedCov.Ax"
	LOCAL_POSITION_NED_COV_FIELD_AY             = "LocalPositionNedCov.Ay"
	LOCAL_POSITION_NED_COV_FIELD_AZ             = "LocalPositionNedCov.Az"
	LOCAL_POSITION_NED_COV_FIELD_COVARIANCE     = "LocalPositionNedCov.Covariance"
	LOCAL_POSITION_NED_COV_FIELD_ESTIMATOR_TYPE = "LocalPositionNedCov.EstimatorType"
)

// ToMap (generated function)
func (m *LocalPositionNedCov) Dict() map[string]interface{} {
	return map[string]interface{}{
		LOCAL_POSITION_NED_COV_FIELD_TIME_USEC:      m.TimeUsec,
		LOCAL_POSITION_NED_COV_FIELD_X:              m.X,
		LOCAL_POSITION_NED_COV_FIELD_Y:              m.Y,
		LOCAL_POSITION_NED_COV_FIELD_Z:              m.Z,
		LOCAL_POSITION_NED_COV_FIELD_VX:             m.Vx,
		LOCAL_POSITION_NED_COV_FIELD_VY:             m.Vy,
		LOCAL_POSITION_NED_COV_FIELD_VZ:             m.Vz,
		LOCAL_POSITION_NED_COV_FIELD_AX:             m.Ax,
		LOCAL_POSITION_NED_COV_FIELD_AY:             m.Ay,
		LOCAL_POSITION_NED_COV_FIELD_AZ:             m.Az,
		LOCAL_POSITION_NED_COV_FIELD_COVARIANCE:     m.Covariance,
		LOCAL_POSITION_NED_COV_FIELD_ESTIMATOR_TYPE: m.EstimatorType,
	}
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
	TimeBootMs uint32 `json:"TimeBootMs" ` // [ ms ] Timestamp (time since system boot).
	Chan1Raw   uint16 `json:"Chan1Raw" `   // [ us ] RC channel 1 value.
	Chan2Raw   uint16 `json:"Chan2Raw" `   // [ us ] RC channel 2 value.
	Chan3Raw   uint16 `json:"Chan3Raw" `   // [ us ] RC channel 3 value.
	Chan4Raw   uint16 `json:"Chan4Raw" `   // [ us ] RC channel 4 value.
	Chan5Raw   uint16 `json:"Chan5Raw" `   // [ us ] RC channel 5 value.
	Chan6Raw   uint16 `json:"Chan6Raw" `   // [ us ] RC channel 6 value.
	Chan7Raw   uint16 `json:"Chan7Raw" `   // [ us ] RC channel 7 value.
	Chan8Raw   uint16 `json:"Chan8Raw" `   // [ us ] RC channel 8 value.
	Chan9Raw   uint16 `json:"Chan9Raw" `   // [ us ] RC channel 9 value.
	Chan10Raw  uint16 `json:"Chan10Raw" `  // [ us ] RC channel 10 value.
	Chan11Raw  uint16 `json:"Chan11Raw" `  // [ us ] RC channel 11 value.
	Chan12Raw  uint16 `json:"Chan12Raw" `  // [ us ] RC channel 12 value.
	Chan13Raw  uint16 `json:"Chan13Raw" `  // [ us ] RC channel 13 value.
	Chan14Raw  uint16 `json:"Chan14Raw" `  // [ us ] RC channel 14 value.
	Chan15Raw  uint16 `json:"Chan15Raw" `  // [ us ] RC channel 15 value.
	Chan16Raw  uint16 `json:"Chan16Raw" `  // [ us ] RC channel 16 value.
	Chan17Raw  uint16 `json:"Chan17Raw" `  // [ us ] RC channel 17 value.
	Chan18Raw  uint16 `json:"Chan18Raw" `  // [ us ] RC channel 18 value.
	Chancount  uint8  `json:"Chancount" `  // Total number of RC channels being received. This can be larger than 18, indicating that more channels are available but not given in this message. This value should be 0 when no RC channels are available.
	Rssi       uint8  `json:"Rssi" `       // Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
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

const (
	RC_CHANNELS_FIELD_TIME_BOOT_MS = "RcChannels.TimeBootMs"
	RC_CHANNELS_FIELD_CHAN1_RAW    = "RcChannels.Chan1Raw"
	RC_CHANNELS_FIELD_CHAN2_RAW    = "RcChannels.Chan2Raw"
	RC_CHANNELS_FIELD_CHAN3_RAW    = "RcChannels.Chan3Raw"
	RC_CHANNELS_FIELD_CHAN4_RAW    = "RcChannels.Chan4Raw"
	RC_CHANNELS_FIELD_CHAN5_RAW    = "RcChannels.Chan5Raw"
	RC_CHANNELS_FIELD_CHAN6_RAW    = "RcChannels.Chan6Raw"
	RC_CHANNELS_FIELD_CHAN7_RAW    = "RcChannels.Chan7Raw"
	RC_CHANNELS_FIELD_CHAN8_RAW    = "RcChannels.Chan8Raw"
	RC_CHANNELS_FIELD_CHAN9_RAW    = "RcChannels.Chan9Raw"
	RC_CHANNELS_FIELD_CHAN10_RAW   = "RcChannels.Chan10Raw"
	RC_CHANNELS_FIELD_CHAN11_RAW   = "RcChannels.Chan11Raw"
	RC_CHANNELS_FIELD_CHAN12_RAW   = "RcChannels.Chan12Raw"
	RC_CHANNELS_FIELD_CHAN13_RAW   = "RcChannels.Chan13Raw"
	RC_CHANNELS_FIELD_CHAN14_RAW   = "RcChannels.Chan14Raw"
	RC_CHANNELS_FIELD_CHAN15_RAW   = "RcChannels.Chan15Raw"
	RC_CHANNELS_FIELD_CHAN16_RAW   = "RcChannels.Chan16Raw"
	RC_CHANNELS_FIELD_CHAN17_RAW   = "RcChannels.Chan17Raw"
	RC_CHANNELS_FIELD_CHAN18_RAW   = "RcChannels.Chan18Raw"
	RC_CHANNELS_FIELD_CHANCOUNT    = "RcChannels.Chancount"
	RC_CHANNELS_FIELD_RSSI         = "RcChannels.Rssi"
)

// ToMap (generated function)
func (m *RcChannels) Dict() map[string]interface{} {
	return map[string]interface{}{
		RC_CHANNELS_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		RC_CHANNELS_FIELD_CHAN1_RAW:    m.Chan1Raw,
		RC_CHANNELS_FIELD_CHAN2_RAW:    m.Chan2Raw,
		RC_CHANNELS_FIELD_CHAN3_RAW:    m.Chan3Raw,
		RC_CHANNELS_FIELD_CHAN4_RAW:    m.Chan4Raw,
		RC_CHANNELS_FIELD_CHAN5_RAW:    m.Chan5Raw,
		RC_CHANNELS_FIELD_CHAN6_RAW:    m.Chan6Raw,
		RC_CHANNELS_FIELD_CHAN7_RAW:    m.Chan7Raw,
		RC_CHANNELS_FIELD_CHAN8_RAW:    m.Chan8Raw,
		RC_CHANNELS_FIELD_CHAN9_RAW:    m.Chan9Raw,
		RC_CHANNELS_FIELD_CHAN10_RAW:   m.Chan10Raw,
		RC_CHANNELS_FIELD_CHAN11_RAW:   m.Chan11Raw,
		RC_CHANNELS_FIELD_CHAN12_RAW:   m.Chan12Raw,
		RC_CHANNELS_FIELD_CHAN13_RAW:   m.Chan13Raw,
		RC_CHANNELS_FIELD_CHAN14_RAW:   m.Chan14Raw,
		RC_CHANNELS_FIELD_CHAN15_RAW:   m.Chan15Raw,
		RC_CHANNELS_FIELD_CHAN16_RAW:   m.Chan16Raw,
		RC_CHANNELS_FIELD_CHAN17_RAW:   m.Chan17Raw,
		RC_CHANNELS_FIELD_CHAN18_RAW:   m.Chan18Raw,
		RC_CHANNELS_FIELD_CHANCOUNT:    m.Chancount,
		RC_CHANNELS_FIELD_RSSI:         m.Rssi,
	}
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
	ReqMessageRate  uint16 `json:"ReqMessageRate" `  // [ Hz ] The requested message rate
	TargetSystem    uint8  `json:"TargetSystem" `    // The target requested to send the message stream.
	TargetComponent uint8  `json:"TargetComponent" ` // The target requested to send the message stream.
	ReqStreamID     uint8  `json:"ReqStreamID" `     // The ID of the requested data stream
	StartStop       uint8  `json:"StartStop" `       // 1 to start sending, 0 to stop sending.
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

const (
	REQUEST_DATA_STREAM_FIELD_REQ_MESSAGE_RATE = "RequestDataStream.ReqMessageRate"
	REQUEST_DATA_STREAM_FIELD_TARGET_SYSTEM    = "RequestDataStream.TargetSystem"
	REQUEST_DATA_STREAM_FIELD_TARGET_COMPONENT = "RequestDataStream.TargetComponent"
	REQUEST_DATA_STREAM_FIELD_REQ_STREAM_ID    = "RequestDataStream.ReqStreamID"
	REQUEST_DATA_STREAM_FIELD_START_STOP       = "RequestDataStream.StartStop"
)

// ToMap (generated function)
func (m *RequestDataStream) Dict() map[string]interface{} {
	return map[string]interface{}{
		REQUEST_DATA_STREAM_FIELD_REQ_MESSAGE_RATE: m.ReqMessageRate,
		REQUEST_DATA_STREAM_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		REQUEST_DATA_STREAM_FIELD_TARGET_COMPONENT: m.TargetComponent,
		REQUEST_DATA_STREAM_FIELD_REQ_STREAM_ID:    m.ReqStreamID,
		REQUEST_DATA_STREAM_FIELD_START_STOP:       m.StartStop,
	}
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
	MessageRate uint16 `json:"MessageRate" ` // [ Hz ] The message rate
	StreamID    uint8  `json:"StreamID" `    // The ID of the requested data stream
	OnOff       uint8  `json:"OnOff" `       // 1 stream is enabled, 0 stream is stopped.
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

const (
	DATA_STREAM_FIELD_MESSAGE_RATE = "DataStream.MessageRate"
	DATA_STREAM_FIELD_STREAM_ID    = "DataStream.StreamID"
	DATA_STREAM_FIELD_ON_OFF       = "DataStream.OnOff"
)

// ToMap (generated function)
func (m *DataStream) Dict() map[string]interface{} {
	return map[string]interface{}{
		DATA_STREAM_FIELD_MESSAGE_RATE: m.MessageRate,
		DATA_STREAM_FIELD_STREAM_ID:    m.StreamID,
		DATA_STREAM_FIELD_ON_OFF:       m.OnOff,
	}
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
	X       int16  `json:"X" `       // X-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to forward(1000)-backward(-1000) movement on a joystick and the pitch of a vehicle.
	Y       int16  `json:"Y" `       // Y-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to left(-1000)-right(1000) movement on a joystick and the roll of a vehicle.
	Z       int16  `json:"Z" `       // Z-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to a separate slider movement with maximum being 1000 and minimum being -1000 on a joystick and the thrust of a vehicle. Positive values are positive thrust, negative values are negative thrust.
	R       int16  `json:"R" `       // R-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to a twisting of the joystick, with counter-clockwise being 1000 and clockwise being -1000, and the yaw of a vehicle.
	Buttons uint16 `json:"Buttons" ` // A bitfield corresponding to the joystick buttons' current state, 1 for pressed, 0 for released. The lowest bit corresponds to Button 1.
	Target  uint8  `json:"Target" `  // The system to be controlled.
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

const (
	MANUAL_CONTROL_FIELD_X       = "ManualControl.X"
	MANUAL_CONTROL_FIELD_Y       = "ManualControl.Y"
	MANUAL_CONTROL_FIELD_Z       = "ManualControl.Z"
	MANUAL_CONTROL_FIELD_R       = "ManualControl.R"
	MANUAL_CONTROL_FIELD_BUTTONS = "ManualControl.Buttons"
	MANUAL_CONTROL_FIELD_TARGET  = "ManualControl.Target"
)

// ToMap (generated function)
func (m *ManualControl) Dict() map[string]interface{} {
	return map[string]interface{}{
		MANUAL_CONTROL_FIELD_X:       m.X,
		MANUAL_CONTROL_FIELD_Y:       m.Y,
		MANUAL_CONTROL_FIELD_Z:       m.Z,
		MANUAL_CONTROL_FIELD_R:       m.R,
		MANUAL_CONTROL_FIELD_BUTTONS: m.Buttons,
		MANUAL_CONTROL_FIELD_TARGET:  m.Target,
	}
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
	Chan1Raw        uint16 `json:"Chan1Raw" `        // [ us ] RC channel 1 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan2Raw        uint16 `json:"Chan2Raw" `        // [ us ] RC channel 2 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan3Raw        uint16 `json:"Chan3Raw" `        // [ us ] RC channel 3 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan4Raw        uint16 `json:"Chan4Raw" `        // [ us ] RC channel 4 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan5Raw        uint16 `json:"Chan5Raw" `        // [ us ] RC channel 5 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan6Raw        uint16 `json:"Chan6Raw" `        // [ us ] RC channel 6 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan7Raw        uint16 `json:"Chan7Raw" `        // [ us ] RC channel 7 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	Chan8Raw        uint16 `json:"Chan8Raw" `        // [ us ] RC channel 8 value. A value of UINT16_MAX means to ignore this field. A value of 0 means to release this channel back to the RC radio.
	TargetSystem    uint8  `json:"TargetSystem" `    // System ID
	TargetComponent uint8  `json:"TargetComponent" ` // Component ID
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

const (
	RC_CHANNELS_OVERRIDE_FIELD_CHAN1_RAW        = "RcChannelsOverride.Chan1Raw"
	RC_CHANNELS_OVERRIDE_FIELD_CHAN2_RAW        = "RcChannelsOverride.Chan2Raw"
	RC_CHANNELS_OVERRIDE_FIELD_CHAN3_RAW        = "RcChannelsOverride.Chan3Raw"
	RC_CHANNELS_OVERRIDE_FIELD_CHAN4_RAW        = "RcChannelsOverride.Chan4Raw"
	RC_CHANNELS_OVERRIDE_FIELD_CHAN5_RAW        = "RcChannelsOverride.Chan5Raw"
	RC_CHANNELS_OVERRIDE_FIELD_CHAN6_RAW        = "RcChannelsOverride.Chan6Raw"
	RC_CHANNELS_OVERRIDE_FIELD_CHAN7_RAW        = "RcChannelsOverride.Chan7Raw"
	RC_CHANNELS_OVERRIDE_FIELD_CHAN8_RAW        = "RcChannelsOverride.Chan8Raw"
	RC_CHANNELS_OVERRIDE_FIELD_TARGET_SYSTEM    = "RcChannelsOverride.TargetSystem"
	RC_CHANNELS_OVERRIDE_FIELD_TARGET_COMPONENT = "RcChannelsOverride.TargetComponent"
)

// ToMap (generated function)
func (m *RcChannelsOverride) Dict() map[string]interface{} {
	return map[string]interface{}{
		RC_CHANNELS_OVERRIDE_FIELD_CHAN1_RAW:        m.Chan1Raw,
		RC_CHANNELS_OVERRIDE_FIELD_CHAN2_RAW:        m.Chan2Raw,
		RC_CHANNELS_OVERRIDE_FIELD_CHAN3_RAW:        m.Chan3Raw,
		RC_CHANNELS_OVERRIDE_FIELD_CHAN4_RAW:        m.Chan4Raw,
		RC_CHANNELS_OVERRIDE_FIELD_CHAN5_RAW:        m.Chan5Raw,
		RC_CHANNELS_OVERRIDE_FIELD_CHAN6_RAW:        m.Chan6Raw,
		RC_CHANNELS_OVERRIDE_FIELD_CHAN7_RAW:        m.Chan7Raw,
		RC_CHANNELS_OVERRIDE_FIELD_CHAN8_RAW:        m.Chan8Raw,
		RC_CHANNELS_OVERRIDE_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		RC_CHANNELS_OVERRIDE_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	Param1          float32   `json:"Param1" `          // PARAM1, see MAV_CMD enum
	Param2          float32   `json:"Param2" `          // PARAM2, see MAV_CMD enum
	Param3          float32   `json:"Param3" `          // PARAM3, see MAV_CMD enum
	Param4          float32   `json:"Param4" `          // PARAM4, see MAV_CMD enum
	X               int32     `json:"X" `               // PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
	Y               int32     `json:"Y" `               // PARAM6 / y position: local: x position in meters * 1e4, global: longitude in degrees *10^7
	Z               float32   `json:"Z" `               // PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame.
	Seq             uint16    `json:"Seq" `             // Waypoint ID (sequence number). Starts at zero. Increases monotonically for each waypoint, no gaps in the sequence (0,1,2,3,4).
	Command         MAV_CMD   `json:"Command" `         // The scheduled action for the waypoint.
	TargetSystem    uint8     `json:"TargetSystem" `    // System ID
	TargetComponent uint8     `json:"TargetComponent" ` // Component ID
	Frame           MAV_FRAME `json:"Frame" `           // The coordinate system of the waypoint.
	Current         uint8     `json:"Current" `         // false:0, true:1
	Autocontinue    uint8     `json:"Autocontinue" `    // Autocontinue to next waypoint
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

const (
	MISSION_ITEM_INT_FIELD_PARAM1           = "MissionItemInt.Param1"
	MISSION_ITEM_INT_FIELD_PARAM2           = "MissionItemInt.Param2"
	MISSION_ITEM_INT_FIELD_PARAM3           = "MissionItemInt.Param3"
	MISSION_ITEM_INT_FIELD_PARAM4           = "MissionItemInt.Param4"
	MISSION_ITEM_INT_FIELD_X                = "MissionItemInt.X"
	MISSION_ITEM_INT_FIELD_Y                = "MissionItemInt.Y"
	MISSION_ITEM_INT_FIELD_Z                = "MissionItemInt.Z"
	MISSION_ITEM_INT_FIELD_SEQ              = "MissionItemInt.Seq"
	MISSION_ITEM_INT_FIELD_COMMAND          = "MissionItemInt.Command"
	MISSION_ITEM_INT_FIELD_TARGET_SYSTEM    = "MissionItemInt.TargetSystem"
	MISSION_ITEM_INT_FIELD_TARGET_COMPONENT = "MissionItemInt.TargetComponent"
	MISSION_ITEM_INT_FIELD_FRAME            = "MissionItemInt.Frame"
	MISSION_ITEM_INT_FIELD_CURRENT          = "MissionItemInt.Current"
	MISSION_ITEM_INT_FIELD_AUTOCONTINUE     = "MissionItemInt.Autocontinue"
)

// ToMap (generated function)
func (m *MissionItemInt) Dict() map[string]interface{} {
	return map[string]interface{}{
		MISSION_ITEM_INT_FIELD_PARAM1:           m.Param1,
		MISSION_ITEM_INT_FIELD_PARAM2:           m.Param2,
		MISSION_ITEM_INT_FIELD_PARAM3:           m.Param3,
		MISSION_ITEM_INT_FIELD_PARAM4:           m.Param4,
		MISSION_ITEM_INT_FIELD_X:                m.X,
		MISSION_ITEM_INT_FIELD_Y:                m.Y,
		MISSION_ITEM_INT_FIELD_Z:                m.Z,
		MISSION_ITEM_INT_FIELD_SEQ:              m.Seq,
		MISSION_ITEM_INT_FIELD_COMMAND:          m.Command,
		MISSION_ITEM_INT_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		MISSION_ITEM_INT_FIELD_TARGET_COMPONENT: m.TargetComponent,
		MISSION_ITEM_INT_FIELD_FRAME:            m.Frame,
		MISSION_ITEM_INT_FIELD_CURRENT:          m.Current,
		MISSION_ITEM_INT_FIELD_AUTOCONTINUE:     m.Autocontinue,
	}
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
	Airspeed    float32 `json:"Airspeed" `    // [ m/s ] Vehicle speed in form appropriate for vehicle type. For standard aircraft this is typically calibrated airspeed (CAS) or indicated airspeed (IAS) - either of which can be used by a pilot to estimate stall speed.
	Groundspeed float32 `json:"Groundspeed" ` // [ m/s ] Current ground speed.
	Alt         float32 `json:"Alt" `         // [ m ] Current altitude (MSL).
	Climb       float32 `json:"Climb" `       // [ m/s ] Current climb rate.
	Heading     int16   `json:"Heading" `     // [ deg ] Current heading in compass units (0-360, 0=north).
	Throttle    uint16  `json:"Throttle" `    // [ % ] Current throttle setting (0 to 100).
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

const (
	VFR_HUD_FIELD_AIRSPEED    = "VfrHud.Airspeed"
	VFR_HUD_FIELD_GROUNDSPEED = "VfrHud.Groundspeed"
	VFR_HUD_FIELD_ALT         = "VfrHud.Alt"
	VFR_HUD_FIELD_CLIMB       = "VfrHud.Climb"
	VFR_HUD_FIELD_HEADING     = "VfrHud.Heading"
	VFR_HUD_FIELD_THROTTLE    = "VfrHud.Throttle"
)

// ToMap (generated function)
func (m *VfrHud) Dict() map[string]interface{} {
	return map[string]interface{}{
		VFR_HUD_FIELD_AIRSPEED:    m.Airspeed,
		VFR_HUD_FIELD_GROUNDSPEED: m.Groundspeed,
		VFR_HUD_FIELD_ALT:         m.Alt,
		VFR_HUD_FIELD_CLIMB:       m.Climb,
		VFR_HUD_FIELD_HEADING:     m.Heading,
		VFR_HUD_FIELD_THROTTLE:    m.Throttle,
	}
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
	Param1          float32   `json:"Param1" `          // PARAM1, see MAV_CMD enum
	Param2          float32   `json:"Param2" `          // PARAM2, see MAV_CMD enum
	Param3          float32   `json:"Param3" `          // PARAM3, see MAV_CMD enum
	Param4          float32   `json:"Param4" `          // PARAM4, see MAV_CMD enum
	X               int32     `json:"X" `               // PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
	Y               int32     `json:"Y" `               // PARAM6 / local: y position in meters * 1e4, global: longitude in degrees * 10^7
	Z               float32   `json:"Z" `               // PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame).
	Command         MAV_CMD   `json:"Command" `         // The scheduled action for the mission item.
	TargetSystem    uint8     `json:"TargetSystem" `    // System ID
	TargetComponent uint8     `json:"TargetComponent" ` // Component ID
	Frame           MAV_FRAME `json:"Frame" `           // The coordinate system of the COMMAND.
	Current         uint8     `json:"Current" `         // Not used.
	Autocontinue    uint8     `json:"Autocontinue" `    // Not used (set 0).
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

const (
	COMMAND_INT_FIELD_PARAM1           = "CommandInt.Param1"
	COMMAND_INT_FIELD_PARAM2           = "CommandInt.Param2"
	COMMAND_INT_FIELD_PARAM3           = "CommandInt.Param3"
	COMMAND_INT_FIELD_PARAM4           = "CommandInt.Param4"
	COMMAND_INT_FIELD_X                = "CommandInt.X"
	COMMAND_INT_FIELD_Y                = "CommandInt.Y"
	COMMAND_INT_FIELD_Z                = "CommandInt.Z"
	COMMAND_INT_FIELD_COMMAND          = "CommandInt.Command"
	COMMAND_INT_FIELD_TARGET_SYSTEM    = "CommandInt.TargetSystem"
	COMMAND_INT_FIELD_TARGET_COMPONENT = "CommandInt.TargetComponent"
	COMMAND_INT_FIELD_FRAME            = "CommandInt.Frame"
	COMMAND_INT_FIELD_CURRENT          = "CommandInt.Current"
	COMMAND_INT_FIELD_AUTOCONTINUE     = "CommandInt.Autocontinue"
)

// ToMap (generated function)
func (m *CommandInt) Dict() map[string]interface{} {
	return map[string]interface{}{
		COMMAND_INT_FIELD_PARAM1:           m.Param1,
		COMMAND_INT_FIELD_PARAM2:           m.Param2,
		COMMAND_INT_FIELD_PARAM3:           m.Param3,
		COMMAND_INT_FIELD_PARAM4:           m.Param4,
		COMMAND_INT_FIELD_X:                m.X,
		COMMAND_INT_FIELD_Y:                m.Y,
		COMMAND_INT_FIELD_Z:                m.Z,
		COMMAND_INT_FIELD_COMMAND:          m.Command,
		COMMAND_INT_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		COMMAND_INT_FIELD_TARGET_COMPONENT: m.TargetComponent,
		COMMAND_INT_FIELD_FRAME:            m.Frame,
		COMMAND_INT_FIELD_CURRENT:          m.Current,
		COMMAND_INT_FIELD_AUTOCONTINUE:     m.Autocontinue,
	}
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
	Param1          float32 `json:"Param1" `          // Parameter 1 (for the specific command).
	Param2          float32 `json:"Param2" `          // Parameter 2 (for the specific command).
	Param3          float32 `json:"Param3" `          // Parameter 3 (for the specific command).
	Param4          float32 `json:"Param4" `          // Parameter 4 (for the specific command).
	Param5          float32 `json:"Param5" `          // Parameter 5 (for the specific command).
	Param6          float32 `json:"Param6" `          // Parameter 6 (for the specific command).
	Param7          float32 `json:"Param7" `          // Parameter 7 (for the specific command).
	Command         MAV_CMD `json:"Command" `         // Command ID (of command to send).
	TargetSystem    uint8   `json:"TargetSystem" `    // System which should execute the command
	TargetComponent uint8   `json:"TargetComponent" ` // Component which should execute the command, 0 for all components
	Confirmation    uint8   `json:"Confirmation" `    // 0: First transmission of this command. 1-255: Confirmation transmissions (e.g. for kill command)
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

const (
	COMMAND_LONG_FIELD_PARAM1           = "CommandLong.Param1"
	COMMAND_LONG_FIELD_PARAM2           = "CommandLong.Param2"
	COMMAND_LONG_FIELD_PARAM3           = "CommandLong.Param3"
	COMMAND_LONG_FIELD_PARAM4           = "CommandLong.Param4"
	COMMAND_LONG_FIELD_PARAM5           = "CommandLong.Param5"
	COMMAND_LONG_FIELD_PARAM6           = "CommandLong.Param6"
	COMMAND_LONG_FIELD_PARAM7           = "CommandLong.Param7"
	COMMAND_LONG_FIELD_COMMAND          = "CommandLong.Command"
	COMMAND_LONG_FIELD_TARGET_SYSTEM    = "CommandLong.TargetSystem"
	COMMAND_LONG_FIELD_TARGET_COMPONENT = "CommandLong.TargetComponent"
	COMMAND_LONG_FIELD_CONFIRMATION     = "CommandLong.Confirmation"
)

// ToMap (generated function)
func (m *CommandLong) Dict() map[string]interface{} {
	return map[string]interface{}{
		COMMAND_LONG_FIELD_PARAM1:           m.Param1,
		COMMAND_LONG_FIELD_PARAM2:           m.Param2,
		COMMAND_LONG_FIELD_PARAM3:           m.Param3,
		COMMAND_LONG_FIELD_PARAM4:           m.Param4,
		COMMAND_LONG_FIELD_PARAM5:           m.Param5,
		COMMAND_LONG_FIELD_PARAM6:           m.Param6,
		COMMAND_LONG_FIELD_PARAM7:           m.Param7,
		COMMAND_LONG_FIELD_COMMAND:          m.Command,
		COMMAND_LONG_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		COMMAND_LONG_FIELD_TARGET_COMPONENT: m.TargetComponent,
		COMMAND_LONG_FIELD_CONFIRMATION:     m.Confirmation,
	}
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
	Command MAV_CMD    `json:"Command" ` // Command ID (of acknowledged command).
	Result  MAV_RESULT `json:"Result" `  // Result of command.
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

const (
	COMMAND_ACK_FIELD_COMMAND = "CommandAck.Command"
	COMMAND_ACK_FIELD_RESULT  = "CommandAck.Result"
)

// ToMap (generated function)
func (m *CommandAck) Dict() map[string]interface{} {
	return map[string]interface{}{
		COMMAND_ACK_FIELD_COMMAND: m.Command,
		COMMAND_ACK_FIELD_RESULT:  m.Result,
	}
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
	Command         MAV_CMD `json:"Command" `         // Command ID (of command to cancel).
	TargetSystem    uint8   `json:"TargetSystem" `    // System executing long running command. Should not be broadcast (0).
	TargetComponent uint8   `json:"TargetComponent" ` // Component executing long running command.
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

const (
	COMMAND_CANCEL_FIELD_COMMAND          = "CommandCancel.Command"
	COMMAND_CANCEL_FIELD_TARGET_SYSTEM    = "CommandCancel.TargetSystem"
	COMMAND_CANCEL_FIELD_TARGET_COMPONENT = "CommandCancel.TargetComponent"
)

// ToMap (generated function)
func (m *CommandCancel) Dict() map[string]interface{} {
	return map[string]interface{}{
		COMMAND_CANCEL_FIELD_COMMAND:          m.Command,
		COMMAND_CANCEL_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		COMMAND_CANCEL_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	TimeBootMs           uint32  `json:"TimeBootMs" `           // [ ms ] Timestamp (time since system boot).
	Roll                 float32 `json:"Roll" `                 // [ rad/s ] Desired roll rate
	Pitch                float32 `json:"Pitch" `                // [ rad/s ] Desired pitch rate
	Yaw                  float32 `json:"Yaw" `                  // [ rad/s ] Desired yaw rate
	Thrust               float32 `json:"Thrust" `               // Collective thrust, normalized to 0 .. 1
	ModeSwitch           uint8   `json:"ModeSwitch" `           // Flight mode switch position, 0.. 255
	ManualOverrideSwitch uint8   `json:"ManualOverrideSwitch" ` // Override mode switch position, 0.. 255
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

const (
	MANUAL_SETPOINT_FIELD_TIME_BOOT_MS           = "ManualSetpoint.TimeBootMs"
	MANUAL_SETPOINT_FIELD_ROLL                   = "ManualSetpoint.Roll"
	MANUAL_SETPOINT_FIELD_PITCH                  = "ManualSetpoint.Pitch"
	MANUAL_SETPOINT_FIELD_YAW                    = "ManualSetpoint.Yaw"
	MANUAL_SETPOINT_FIELD_THRUST                 = "ManualSetpoint.Thrust"
	MANUAL_SETPOINT_FIELD_MODE_SWITCH            = "ManualSetpoint.ModeSwitch"
	MANUAL_SETPOINT_FIELD_MANUAL_OVERRIDE_SWITCH = "ManualSetpoint.ManualOverrideSwitch"
)

// ToMap (generated function)
func (m *ManualSetpoint) Dict() map[string]interface{} {
	return map[string]interface{}{
		MANUAL_SETPOINT_FIELD_TIME_BOOT_MS:           m.TimeBootMs,
		MANUAL_SETPOINT_FIELD_ROLL:                   m.Roll,
		MANUAL_SETPOINT_FIELD_PITCH:                  m.Pitch,
		MANUAL_SETPOINT_FIELD_YAW:                    m.Yaw,
		MANUAL_SETPOINT_FIELD_THRUST:                 m.Thrust,
		MANUAL_SETPOINT_FIELD_MODE_SWITCH:            m.ModeSwitch,
		MANUAL_SETPOINT_FIELD_MANUAL_OVERRIDE_SWITCH: m.ManualOverrideSwitch,
	}
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
	TimeBootMs      uint32                   `json:"TimeBootMs" `      // [ ms ] Timestamp (time since system boot).
	Q               []float32                `json:"Q" len:"4" `       // Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	BodyRollRate    float32                  `json:"BodyRollRate" `    // [ rad/s ] Body roll rate
	BodyPitchRate   float32                  `json:"BodyPitchRate" `   // [ rad/s ] Body pitch rate
	BodyYawRate     float32                  `json:"BodyYawRate" `     // [ rad/s ] Body yaw rate
	Thrust          float32                  `json:"Thrust" `          // Collective thrust, normalized to 0 .. 1 (-1 .. 1 for vehicles capable of reverse trust)
	TargetSystem    uint8                    `json:"TargetSystem" `    // System ID
	TargetComponent uint8                    `json:"TargetComponent" ` // Component ID
	TypeMask        ATTITUDE_TARGET_TYPEMASK `json:"TypeMask" `        // Bitmap to indicate which dimensions should be ignored by the vehicle.
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

const (
	SET_ATTITUDE_TARGET_FIELD_TIME_BOOT_MS     = "SetAttitudeTarget.TimeBootMs"
	SET_ATTITUDE_TARGET_FIELD_Q                = "SetAttitudeTarget.Q"
	SET_ATTITUDE_TARGET_FIELD_BODY_ROLL_RATE   = "SetAttitudeTarget.BodyRollRate"
	SET_ATTITUDE_TARGET_FIELD_BODY_PITCH_RATE  = "SetAttitudeTarget.BodyPitchRate"
	SET_ATTITUDE_TARGET_FIELD_BODY_YAW_RATE    = "SetAttitudeTarget.BodyYawRate"
	SET_ATTITUDE_TARGET_FIELD_THRUST           = "SetAttitudeTarget.Thrust"
	SET_ATTITUDE_TARGET_FIELD_TARGET_SYSTEM    = "SetAttitudeTarget.TargetSystem"
	SET_ATTITUDE_TARGET_FIELD_TARGET_COMPONENT = "SetAttitudeTarget.TargetComponent"
	SET_ATTITUDE_TARGET_FIELD_TYPE_MASK        = "SetAttitudeTarget.TypeMask"
)

// ToMap (generated function)
func (m *SetAttitudeTarget) Dict() map[string]interface{} {
	return map[string]interface{}{
		SET_ATTITUDE_TARGET_FIELD_TIME_BOOT_MS:     m.TimeBootMs,
		SET_ATTITUDE_TARGET_FIELD_Q:                m.Q,
		SET_ATTITUDE_TARGET_FIELD_BODY_ROLL_RATE:   m.BodyRollRate,
		SET_ATTITUDE_TARGET_FIELD_BODY_PITCH_RATE:  m.BodyPitchRate,
		SET_ATTITUDE_TARGET_FIELD_BODY_YAW_RATE:    m.BodyYawRate,
		SET_ATTITUDE_TARGET_FIELD_THRUST:           m.Thrust,
		SET_ATTITUDE_TARGET_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		SET_ATTITUDE_TARGET_FIELD_TARGET_COMPONENT: m.TargetComponent,
		SET_ATTITUDE_TARGET_FIELD_TYPE_MASK:        m.TypeMask,
	}
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
	TimeBootMs    uint32                   `json:"TimeBootMs" `    // [ ms ] Timestamp (time since system boot).
	Q             []float32                `json:"Q" len:"4" `     // Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	BodyRollRate  float32                  `json:"BodyRollRate" `  // [ rad/s ] Body roll rate
	BodyPitchRate float32                  `json:"BodyPitchRate" ` // [ rad/s ] Body pitch rate
	BodyYawRate   float32                  `json:"BodyYawRate" `   // [ rad/s ] Body yaw rate
	Thrust        float32                  `json:"Thrust" `        // Collective thrust, normalized to 0 .. 1 (-1 .. 1 for vehicles capable of reverse trust)
	TypeMask      ATTITUDE_TARGET_TYPEMASK `json:"TypeMask" `      // Bitmap to indicate which dimensions should be ignored by the vehicle.
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

const (
	ATTITUDE_TARGET_FIELD_TIME_BOOT_MS    = "AttitudeTarget.TimeBootMs"
	ATTITUDE_TARGET_FIELD_Q               = "AttitudeTarget.Q"
	ATTITUDE_TARGET_FIELD_BODY_ROLL_RATE  = "AttitudeTarget.BodyRollRate"
	ATTITUDE_TARGET_FIELD_BODY_PITCH_RATE = "AttitudeTarget.BodyPitchRate"
	ATTITUDE_TARGET_FIELD_BODY_YAW_RATE   = "AttitudeTarget.BodyYawRate"
	ATTITUDE_TARGET_FIELD_THRUST          = "AttitudeTarget.Thrust"
	ATTITUDE_TARGET_FIELD_TYPE_MASK       = "AttitudeTarget.TypeMask"
)

// ToMap (generated function)
func (m *AttitudeTarget) Dict() map[string]interface{} {
	return map[string]interface{}{
		ATTITUDE_TARGET_FIELD_TIME_BOOT_MS:    m.TimeBootMs,
		ATTITUDE_TARGET_FIELD_Q:               m.Q,
		ATTITUDE_TARGET_FIELD_BODY_ROLL_RATE:  m.BodyRollRate,
		ATTITUDE_TARGET_FIELD_BODY_PITCH_RATE: m.BodyPitchRate,
		ATTITUDE_TARGET_FIELD_BODY_YAW_RATE:   m.BodyYawRate,
		ATTITUDE_TARGET_FIELD_THRUST:          m.Thrust,
		ATTITUDE_TARGET_FIELD_TYPE_MASK:       m.TypeMask,
	}
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
	TimeBootMs      uint32                   `json:"TimeBootMs" `      // [ ms ] Timestamp (time since system boot).
	X               float32                  `json:"X" `               // [ m ] X Position in NED frame
	Y               float32                  `json:"Y" `               // [ m ] Y Position in NED frame
	Z               float32                  `json:"Z" `               // [ m ] Z Position in NED frame (note, altitude is negative in NED)
	Vx              float32                  `json:"Vx" `              // [ m/s ] X velocity in NED frame
	Vy              float32                  `json:"Vy" `              // [ m/s ] Y velocity in NED frame
	Vz              float32                  `json:"Vz" `              // [ m/s ] Z velocity in NED frame
	Afx             float32                  `json:"Afx" `             // [ m/s/s ] X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy             float32                  `json:"Afy" `             // [ m/s/s ] Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz             float32                  `json:"Afz" `             // [ m/s/s ] Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Yaw             float32                  `json:"Yaw" `             // [ rad ] yaw setpoint
	YawRate         float32                  `json:"YawRate" `         // [ rad/s ] yaw rate setpoint
	TypeMask        POSITION_TARGET_TYPEMASK `json:"TypeMask" `        // Bitmap to indicate which dimensions should be ignored by the vehicle.
	TargetSystem    uint8                    `json:"TargetSystem" `    // System ID
	TargetComponent uint8                    `json:"TargetComponent" ` // Component ID
	CoordinateFrame MAV_FRAME                `json:"CoordinateFrame" ` // Valid options are: MAV_FRAME_LOCAL_NED = 1, MAV_FRAME_LOCAL_OFFSET_NED = 7, MAV_FRAME_BODY_NED = 8, MAV_FRAME_BODY_OFFSET_NED = 9
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

const (
	SET_POSITION_TARGET_LOCAL_NED_FIELD_TIME_BOOT_MS     = "SetPositionTargetLocalNed.TimeBootMs"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_X                = "SetPositionTargetLocalNed.X"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_Y                = "SetPositionTargetLocalNed.Y"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_Z                = "SetPositionTargetLocalNed.Z"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_VX               = "SetPositionTargetLocalNed.Vx"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_VY               = "SetPositionTargetLocalNed.Vy"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_VZ               = "SetPositionTargetLocalNed.Vz"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_AFX              = "SetPositionTargetLocalNed.Afx"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_AFY              = "SetPositionTargetLocalNed.Afy"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_AFZ              = "SetPositionTargetLocalNed.Afz"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_YAW              = "SetPositionTargetLocalNed.Yaw"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_YAW_RATE         = "SetPositionTargetLocalNed.YawRate"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_TYPE_MASK        = "SetPositionTargetLocalNed.TypeMask"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_TARGET_SYSTEM    = "SetPositionTargetLocalNed.TargetSystem"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_TARGET_COMPONENT = "SetPositionTargetLocalNed.TargetComponent"
	SET_POSITION_TARGET_LOCAL_NED_FIELD_COORDINATE_FRAME = "SetPositionTargetLocalNed.CoordinateFrame"
)

// ToMap (generated function)
func (m *SetPositionTargetLocalNed) Dict() map[string]interface{} {
	return map[string]interface{}{
		SET_POSITION_TARGET_LOCAL_NED_FIELD_TIME_BOOT_MS:     m.TimeBootMs,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_X:                m.X,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_Y:                m.Y,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_Z:                m.Z,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_VX:               m.Vx,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_VY:               m.Vy,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_VZ:               m.Vz,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_AFX:              m.Afx,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_AFY:              m.Afy,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_AFZ:              m.Afz,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_YAW:              m.Yaw,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_YAW_RATE:         m.YawRate,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_TYPE_MASK:        m.TypeMask,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_TARGET_COMPONENT: m.TargetComponent,
		SET_POSITION_TARGET_LOCAL_NED_FIELD_COORDINATE_FRAME: m.CoordinateFrame,
	}
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
	TimeBootMs      uint32                   `json:"TimeBootMs" `      // [ ms ] Timestamp (time since system boot).
	X               float32                  `json:"X" `               // [ m ] X Position in NED frame
	Y               float32                  `json:"Y" `               // [ m ] Y Position in NED frame
	Z               float32                  `json:"Z" `               // [ m ] Z Position in NED frame (note, altitude is negative in NED)
	Vx              float32                  `json:"Vx" `              // [ m/s ] X velocity in NED frame
	Vy              float32                  `json:"Vy" `              // [ m/s ] Y velocity in NED frame
	Vz              float32                  `json:"Vz" `              // [ m/s ] Z velocity in NED frame
	Afx             float32                  `json:"Afx" `             // [ m/s/s ] X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy             float32                  `json:"Afy" `             // [ m/s/s ] Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz             float32                  `json:"Afz" `             // [ m/s/s ] Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Yaw             float32                  `json:"Yaw" `             // [ rad ] yaw setpoint
	YawRate         float32                  `json:"YawRate" `         // [ rad/s ] yaw rate setpoint
	TypeMask        POSITION_TARGET_TYPEMASK `json:"TypeMask" `        // Bitmap to indicate which dimensions should be ignored by the vehicle.
	CoordinateFrame MAV_FRAME                `json:"CoordinateFrame" ` // Valid options are: MAV_FRAME_LOCAL_NED = 1, MAV_FRAME_LOCAL_OFFSET_NED = 7, MAV_FRAME_BODY_NED = 8, MAV_FRAME_BODY_OFFSET_NED = 9
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

const (
	POSITION_TARGET_LOCAL_NED_FIELD_TIME_BOOT_MS     = "PositionTargetLocalNed.TimeBootMs"
	POSITION_TARGET_LOCAL_NED_FIELD_X                = "PositionTargetLocalNed.X"
	POSITION_TARGET_LOCAL_NED_FIELD_Y                = "PositionTargetLocalNed.Y"
	POSITION_TARGET_LOCAL_NED_FIELD_Z                = "PositionTargetLocalNed.Z"
	POSITION_TARGET_LOCAL_NED_FIELD_VX               = "PositionTargetLocalNed.Vx"
	POSITION_TARGET_LOCAL_NED_FIELD_VY               = "PositionTargetLocalNed.Vy"
	POSITION_TARGET_LOCAL_NED_FIELD_VZ               = "PositionTargetLocalNed.Vz"
	POSITION_TARGET_LOCAL_NED_FIELD_AFX              = "PositionTargetLocalNed.Afx"
	POSITION_TARGET_LOCAL_NED_FIELD_AFY              = "PositionTargetLocalNed.Afy"
	POSITION_TARGET_LOCAL_NED_FIELD_AFZ              = "PositionTargetLocalNed.Afz"
	POSITION_TARGET_LOCAL_NED_FIELD_YAW              = "PositionTargetLocalNed.Yaw"
	POSITION_TARGET_LOCAL_NED_FIELD_YAW_RATE         = "PositionTargetLocalNed.YawRate"
	POSITION_TARGET_LOCAL_NED_FIELD_TYPE_MASK        = "PositionTargetLocalNed.TypeMask"
	POSITION_TARGET_LOCAL_NED_FIELD_COORDINATE_FRAME = "PositionTargetLocalNed.CoordinateFrame"
)

// ToMap (generated function)
func (m *PositionTargetLocalNed) Dict() map[string]interface{} {
	return map[string]interface{}{
		POSITION_TARGET_LOCAL_NED_FIELD_TIME_BOOT_MS:     m.TimeBootMs,
		POSITION_TARGET_LOCAL_NED_FIELD_X:                m.X,
		POSITION_TARGET_LOCAL_NED_FIELD_Y:                m.Y,
		POSITION_TARGET_LOCAL_NED_FIELD_Z:                m.Z,
		POSITION_TARGET_LOCAL_NED_FIELD_VX:               m.Vx,
		POSITION_TARGET_LOCAL_NED_FIELD_VY:               m.Vy,
		POSITION_TARGET_LOCAL_NED_FIELD_VZ:               m.Vz,
		POSITION_TARGET_LOCAL_NED_FIELD_AFX:              m.Afx,
		POSITION_TARGET_LOCAL_NED_FIELD_AFY:              m.Afy,
		POSITION_TARGET_LOCAL_NED_FIELD_AFZ:              m.Afz,
		POSITION_TARGET_LOCAL_NED_FIELD_YAW:              m.Yaw,
		POSITION_TARGET_LOCAL_NED_FIELD_YAW_RATE:         m.YawRate,
		POSITION_TARGET_LOCAL_NED_FIELD_TYPE_MASK:        m.TypeMask,
		POSITION_TARGET_LOCAL_NED_FIELD_COORDINATE_FRAME: m.CoordinateFrame,
	}
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
	TimeBootMs      uint32                   `json:"TimeBootMs" `      // [ ms ] Timestamp (time since system boot). The rationale for the timestamp in the setpoint is to allow the system to compensate for the transport delay of the setpoint. This allows the system to compensate processing latency.
	LatInt          int32                    `json:"LatInt" `          // [ degE7 ] X Position in WGS84 frame
	LonInt          int32                    `json:"LonInt" `          // [ degE7 ] Y Position in WGS84 frame
	Alt             float32                  `json:"Alt" `             // [ m ] Altitude (MSL, Relative to home, or AGL - depending on frame)
	Vx              float32                  `json:"Vx" `              // [ m/s ] X velocity in NED frame
	Vy              float32                  `json:"Vy" `              // [ m/s ] Y velocity in NED frame
	Vz              float32                  `json:"Vz" `              // [ m/s ] Z velocity in NED frame
	Afx             float32                  `json:"Afx" `             // [ m/s/s ] X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy             float32                  `json:"Afy" `             // [ m/s/s ] Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz             float32                  `json:"Afz" `             // [ m/s/s ] Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Yaw             float32                  `json:"Yaw" `             // [ rad ] yaw setpoint
	YawRate         float32                  `json:"YawRate" `         // [ rad/s ] yaw rate setpoint
	TypeMask        POSITION_TARGET_TYPEMASK `json:"TypeMask" `        // Bitmap to indicate which dimensions should be ignored by the vehicle.
	TargetSystem    uint8                    `json:"TargetSystem" `    // System ID
	TargetComponent uint8                    `json:"TargetComponent" ` // Component ID
	CoordinateFrame MAV_FRAME                `json:"CoordinateFrame" ` // Valid options are: MAV_FRAME_GLOBAL_INT = 5, MAV_FRAME_GLOBAL_RELATIVE_ALT_INT = 6, MAV_FRAME_GLOBAL_TERRAIN_ALT_INT = 11
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

const (
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_TIME_BOOT_MS     = "SetPositionTargetGlobalInt.TimeBootMs"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_LAT_INT          = "SetPositionTargetGlobalInt.LatInt"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_LON_INT          = "SetPositionTargetGlobalInt.LonInt"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_ALT              = "SetPositionTargetGlobalInt.Alt"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_VX               = "SetPositionTargetGlobalInt.Vx"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_VY               = "SetPositionTargetGlobalInt.Vy"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_VZ               = "SetPositionTargetGlobalInt.Vz"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_AFX              = "SetPositionTargetGlobalInt.Afx"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_AFY              = "SetPositionTargetGlobalInt.Afy"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_AFZ              = "SetPositionTargetGlobalInt.Afz"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_YAW              = "SetPositionTargetGlobalInt.Yaw"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_YAW_RATE         = "SetPositionTargetGlobalInt.YawRate"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_TYPE_MASK        = "SetPositionTargetGlobalInt.TypeMask"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_TARGET_SYSTEM    = "SetPositionTargetGlobalInt.TargetSystem"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_TARGET_COMPONENT = "SetPositionTargetGlobalInt.TargetComponent"
	SET_POSITION_TARGET_GLOBAL_INT_FIELD_COORDINATE_FRAME = "SetPositionTargetGlobalInt.CoordinateFrame"
)

// ToMap (generated function)
func (m *SetPositionTargetGlobalInt) Dict() map[string]interface{} {
	return map[string]interface{}{
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_TIME_BOOT_MS:     m.TimeBootMs,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_LAT_INT:          m.LatInt,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_LON_INT:          m.LonInt,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_ALT:              m.Alt,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_VX:               m.Vx,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_VY:               m.Vy,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_VZ:               m.Vz,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_AFX:              m.Afx,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_AFY:              m.Afy,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_AFZ:              m.Afz,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_YAW:              m.Yaw,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_YAW_RATE:         m.YawRate,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_TYPE_MASK:        m.TypeMask,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_TARGET_COMPONENT: m.TargetComponent,
		SET_POSITION_TARGET_GLOBAL_INT_FIELD_COORDINATE_FRAME: m.CoordinateFrame,
	}
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
	TimeBootMs      uint32                   `json:"TimeBootMs" `      // [ ms ] Timestamp (time since system boot). The rationale for the timestamp in the setpoint is to allow the system to compensate for the transport delay of the setpoint. This allows the system to compensate processing latency.
	LatInt          int32                    `json:"LatInt" `          // [ degE7 ] X Position in WGS84 frame
	LonInt          int32                    `json:"LonInt" `          // [ degE7 ] Y Position in WGS84 frame
	Alt             float32                  `json:"Alt" `             // [ m ] Altitude (MSL, AGL or relative to home altitude, depending on frame)
	Vx              float32                  `json:"Vx" `              // [ m/s ] X velocity in NED frame
	Vy              float32                  `json:"Vy" `              // [ m/s ] Y velocity in NED frame
	Vz              float32                  `json:"Vz" `              // [ m/s ] Z velocity in NED frame
	Afx             float32                  `json:"Afx" `             // [ m/s/s ] X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy             float32                  `json:"Afy" `             // [ m/s/s ] Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz             float32                  `json:"Afz" `             // [ m/s/s ] Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Yaw             float32                  `json:"Yaw" `             // [ rad ] yaw setpoint
	YawRate         float32                  `json:"YawRate" `         // [ rad/s ] yaw rate setpoint
	TypeMask        POSITION_TARGET_TYPEMASK `json:"TypeMask" `        // Bitmap to indicate which dimensions should be ignored by the vehicle.
	CoordinateFrame MAV_FRAME                `json:"CoordinateFrame" ` // Valid options are: MAV_FRAME_GLOBAL_INT = 5, MAV_FRAME_GLOBAL_RELATIVE_ALT_INT = 6, MAV_FRAME_GLOBAL_TERRAIN_ALT_INT = 11
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

const (
	POSITION_TARGET_GLOBAL_INT_FIELD_TIME_BOOT_MS     = "PositionTargetGlobalInt.TimeBootMs"
	POSITION_TARGET_GLOBAL_INT_FIELD_LAT_INT          = "PositionTargetGlobalInt.LatInt"
	POSITION_TARGET_GLOBAL_INT_FIELD_LON_INT          = "PositionTargetGlobalInt.LonInt"
	POSITION_TARGET_GLOBAL_INT_FIELD_ALT              = "PositionTargetGlobalInt.Alt"
	POSITION_TARGET_GLOBAL_INT_FIELD_VX               = "PositionTargetGlobalInt.Vx"
	POSITION_TARGET_GLOBAL_INT_FIELD_VY               = "PositionTargetGlobalInt.Vy"
	POSITION_TARGET_GLOBAL_INT_FIELD_VZ               = "PositionTargetGlobalInt.Vz"
	POSITION_TARGET_GLOBAL_INT_FIELD_AFX              = "PositionTargetGlobalInt.Afx"
	POSITION_TARGET_GLOBAL_INT_FIELD_AFY              = "PositionTargetGlobalInt.Afy"
	POSITION_TARGET_GLOBAL_INT_FIELD_AFZ              = "PositionTargetGlobalInt.Afz"
	POSITION_TARGET_GLOBAL_INT_FIELD_YAW              = "PositionTargetGlobalInt.Yaw"
	POSITION_TARGET_GLOBAL_INT_FIELD_YAW_RATE         = "PositionTargetGlobalInt.YawRate"
	POSITION_TARGET_GLOBAL_INT_FIELD_TYPE_MASK        = "PositionTargetGlobalInt.TypeMask"
	POSITION_TARGET_GLOBAL_INT_FIELD_COORDINATE_FRAME = "PositionTargetGlobalInt.CoordinateFrame"
)

// ToMap (generated function)
func (m *PositionTargetGlobalInt) Dict() map[string]interface{} {
	return map[string]interface{}{
		POSITION_TARGET_GLOBAL_INT_FIELD_TIME_BOOT_MS:     m.TimeBootMs,
		POSITION_TARGET_GLOBAL_INT_FIELD_LAT_INT:          m.LatInt,
		POSITION_TARGET_GLOBAL_INT_FIELD_LON_INT:          m.LonInt,
		POSITION_TARGET_GLOBAL_INT_FIELD_ALT:              m.Alt,
		POSITION_TARGET_GLOBAL_INT_FIELD_VX:               m.Vx,
		POSITION_TARGET_GLOBAL_INT_FIELD_VY:               m.Vy,
		POSITION_TARGET_GLOBAL_INT_FIELD_VZ:               m.Vz,
		POSITION_TARGET_GLOBAL_INT_FIELD_AFX:              m.Afx,
		POSITION_TARGET_GLOBAL_INT_FIELD_AFY:              m.Afy,
		POSITION_TARGET_GLOBAL_INT_FIELD_AFZ:              m.Afz,
		POSITION_TARGET_GLOBAL_INT_FIELD_YAW:              m.Yaw,
		POSITION_TARGET_GLOBAL_INT_FIELD_YAW_RATE:         m.YawRate,
		POSITION_TARGET_GLOBAL_INT_FIELD_TYPE_MASK:        m.TypeMask,
		POSITION_TARGET_GLOBAL_INT_FIELD_COORDINATE_FRAME: m.CoordinateFrame,
	}
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
	TimeBootMs uint32  `json:"TimeBootMs" ` // [ ms ] Timestamp (time since system boot).
	X          float32 `json:"X" `          // [ m ] X Position
	Y          float32 `json:"Y" `          // [ m ] Y Position
	Z          float32 `json:"Z" `          // [ m ] Z Position
	Roll       float32 `json:"Roll" `       // [ rad ] Roll
	Pitch      float32 `json:"Pitch" `      // [ rad ] Pitch
	Yaw        float32 `json:"Yaw" `        // [ rad ] Yaw
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

const (
	LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_TIME_BOOT_MS = "LocalPositionNedSystemGlobalOffset.TimeBootMs"
	LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_X            = "LocalPositionNedSystemGlobalOffset.X"
	LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_Y            = "LocalPositionNedSystemGlobalOffset.Y"
	LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_Z            = "LocalPositionNedSystemGlobalOffset.Z"
	LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_ROLL         = "LocalPositionNedSystemGlobalOffset.Roll"
	LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_PITCH        = "LocalPositionNedSystemGlobalOffset.Pitch"
	LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_YAW          = "LocalPositionNedSystemGlobalOffset.Yaw"
)

// ToMap (generated function)
func (m *LocalPositionNedSystemGlobalOffset) Dict() map[string]interface{} {
	return map[string]interface{}{
		LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_X:            m.X,
		LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_Y:            m.Y,
		LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_Z:            m.Z,
		LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_ROLL:         m.Roll,
		LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_PITCH:        m.Pitch,
		LOCAL_POSITION_NED_SYSTEM_GLOBAL_OFFSET_FIELD_YAW:          m.Yaw,
	}
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
	TimeUsec   uint64  `json:"TimeUsec" `   // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Roll       float32 `json:"Roll" `       // [ rad ] Roll angle
	Pitch      float32 `json:"Pitch" `      // [ rad ] Pitch angle
	Yaw        float32 `json:"Yaw" `        // [ rad ] Yaw angle
	Rollspeed  float32 `json:"Rollspeed" `  // [ rad/s ] Body frame roll / phi angular speed
	Pitchspeed float32 `json:"Pitchspeed" ` // [ rad/s ] Body frame pitch / theta angular speed
	Yawspeed   float32 `json:"Yawspeed" `   // [ rad/s ] Body frame yaw / psi angular speed
	Lat        int32   `json:"Lat" `        // [ degE7 ] Latitude
	Lon        int32   `json:"Lon" `        // [ degE7 ] Longitude
	Alt        int32   `json:"Alt" `        // [ mm ] Altitude
	Vx         int16   `json:"Vx" `         // [ cm/s ] Ground X Speed (Latitude)
	Vy         int16   `json:"Vy" `         // [ cm/s ] Ground Y Speed (Longitude)
	Vz         int16   `json:"Vz" `         // [ cm/s ] Ground Z Speed (Altitude)
	Xacc       int16   `json:"Xacc" `       // [ mG ] X acceleration
	Yacc       int16   `json:"Yacc" `       // [ mG ] Y acceleration
	Zacc       int16   `json:"Zacc" `       // [ mG ] Z acceleration
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

const (
	HIL_STATE_FIELD_TIME_USEC  = "HilState.TimeUsec"
	HIL_STATE_FIELD_ROLL       = "HilState.Roll"
	HIL_STATE_FIELD_PITCH      = "HilState.Pitch"
	HIL_STATE_FIELD_YAW        = "HilState.Yaw"
	HIL_STATE_FIELD_ROLLSPEED  = "HilState.Rollspeed"
	HIL_STATE_FIELD_PITCHSPEED = "HilState.Pitchspeed"
	HIL_STATE_FIELD_YAWSPEED   = "HilState.Yawspeed"
	HIL_STATE_FIELD_LAT        = "HilState.Lat"
	HIL_STATE_FIELD_LON        = "HilState.Lon"
	HIL_STATE_FIELD_ALT        = "HilState.Alt"
	HIL_STATE_FIELD_VX         = "HilState.Vx"
	HIL_STATE_FIELD_VY         = "HilState.Vy"
	HIL_STATE_FIELD_VZ         = "HilState.Vz"
	HIL_STATE_FIELD_XACC       = "HilState.Xacc"
	HIL_STATE_FIELD_YACC       = "HilState.Yacc"
	HIL_STATE_FIELD_ZACC       = "HilState.Zacc"
)

// ToMap (generated function)
func (m *HilState) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIL_STATE_FIELD_TIME_USEC:  m.TimeUsec,
		HIL_STATE_FIELD_ROLL:       m.Roll,
		HIL_STATE_FIELD_PITCH:      m.Pitch,
		HIL_STATE_FIELD_YAW:        m.Yaw,
		HIL_STATE_FIELD_ROLLSPEED:  m.Rollspeed,
		HIL_STATE_FIELD_PITCHSPEED: m.Pitchspeed,
		HIL_STATE_FIELD_YAWSPEED:   m.Yawspeed,
		HIL_STATE_FIELD_LAT:        m.Lat,
		HIL_STATE_FIELD_LON:        m.Lon,
		HIL_STATE_FIELD_ALT:        m.Alt,
		HIL_STATE_FIELD_VX:         m.Vx,
		HIL_STATE_FIELD_VY:         m.Vy,
		HIL_STATE_FIELD_VZ:         m.Vz,
		HIL_STATE_FIELD_XACC:       m.Xacc,
		HIL_STATE_FIELD_YACC:       m.Yacc,
		HIL_STATE_FIELD_ZACC:       m.Zacc,
	}
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
	TimeUsec      uint64   `json:"TimeUsec" `      // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	RollAilerons  float32  `json:"RollAilerons" `  // Control output -1 .. 1
	PitchElevator float32  `json:"PitchElevator" ` // Control output -1 .. 1
	YawRudder     float32  `json:"YawRudder" `     // Control output -1 .. 1
	Throttle      float32  `json:"Throttle" `      // Throttle 0 .. 1
	Aux1          float32  `json:"Aux1" `          // Aux 1, -1 .. 1
	Aux2          float32  `json:"Aux2" `          // Aux 2, -1 .. 1
	Aux3          float32  `json:"Aux3" `          // Aux 3, -1 .. 1
	Aux4          float32  `json:"Aux4" `          // Aux 4, -1 .. 1
	Mode          MAV_MODE `json:"Mode" `          // System mode.
	NavMode       uint8    `json:"NavMode" `       // Navigation mode (MAV_NAV_MODE)
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

const (
	HIL_CONTROLS_FIELD_TIME_USEC      = "HilControls.TimeUsec"
	HIL_CONTROLS_FIELD_ROLL_AILERONS  = "HilControls.RollAilerons"
	HIL_CONTROLS_FIELD_PITCH_ELEVATOR = "HilControls.PitchElevator"
	HIL_CONTROLS_FIELD_YAW_RUDDER     = "HilControls.YawRudder"
	HIL_CONTROLS_FIELD_THROTTLE       = "HilControls.Throttle"
	HIL_CONTROLS_FIELD_AUX1           = "HilControls.Aux1"
	HIL_CONTROLS_FIELD_AUX2           = "HilControls.Aux2"
	HIL_CONTROLS_FIELD_AUX3           = "HilControls.Aux3"
	HIL_CONTROLS_FIELD_AUX4           = "HilControls.Aux4"
	HIL_CONTROLS_FIELD_MODE           = "HilControls.Mode"
	HIL_CONTROLS_FIELD_NAV_MODE       = "HilControls.NavMode"
)

// ToMap (generated function)
func (m *HilControls) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIL_CONTROLS_FIELD_TIME_USEC:      m.TimeUsec,
		HIL_CONTROLS_FIELD_ROLL_AILERONS:  m.RollAilerons,
		HIL_CONTROLS_FIELD_PITCH_ELEVATOR: m.PitchElevator,
		HIL_CONTROLS_FIELD_YAW_RUDDER:     m.YawRudder,
		HIL_CONTROLS_FIELD_THROTTLE:       m.Throttle,
		HIL_CONTROLS_FIELD_AUX1:           m.Aux1,
		HIL_CONTROLS_FIELD_AUX2:           m.Aux2,
		HIL_CONTROLS_FIELD_AUX3:           m.Aux3,
		HIL_CONTROLS_FIELD_AUX4:           m.Aux4,
		HIL_CONTROLS_FIELD_MODE:           m.Mode,
		HIL_CONTROLS_FIELD_NAV_MODE:       m.NavMode,
	}
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
	TimeUsec  uint64 `json:"TimeUsec" `  // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Chan1Raw  uint16 `json:"Chan1Raw" `  // [ us ] RC channel 1 value
	Chan2Raw  uint16 `json:"Chan2Raw" `  // [ us ] RC channel 2 value
	Chan3Raw  uint16 `json:"Chan3Raw" `  // [ us ] RC channel 3 value
	Chan4Raw  uint16 `json:"Chan4Raw" `  // [ us ] RC channel 4 value
	Chan5Raw  uint16 `json:"Chan5Raw" `  // [ us ] RC channel 5 value
	Chan6Raw  uint16 `json:"Chan6Raw" `  // [ us ] RC channel 6 value
	Chan7Raw  uint16 `json:"Chan7Raw" `  // [ us ] RC channel 7 value
	Chan8Raw  uint16 `json:"Chan8Raw" `  // [ us ] RC channel 8 value
	Chan9Raw  uint16 `json:"Chan9Raw" `  // [ us ] RC channel 9 value
	Chan10Raw uint16 `json:"Chan10Raw" ` // [ us ] RC channel 10 value
	Chan11Raw uint16 `json:"Chan11Raw" ` // [ us ] RC channel 11 value
	Chan12Raw uint16 `json:"Chan12Raw" ` // [ us ] RC channel 12 value
	Rssi      uint8  `json:"Rssi" `      // Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
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

const (
	HIL_RC_INPUTS_RAW_FIELD_TIME_USEC  = "HilRcInputsRaw.TimeUsec"
	HIL_RC_INPUTS_RAW_FIELD_CHAN1_RAW  = "HilRcInputsRaw.Chan1Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN2_RAW  = "HilRcInputsRaw.Chan2Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN3_RAW  = "HilRcInputsRaw.Chan3Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN4_RAW  = "HilRcInputsRaw.Chan4Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN5_RAW  = "HilRcInputsRaw.Chan5Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN6_RAW  = "HilRcInputsRaw.Chan6Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN7_RAW  = "HilRcInputsRaw.Chan7Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN8_RAW  = "HilRcInputsRaw.Chan8Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN9_RAW  = "HilRcInputsRaw.Chan9Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN10_RAW = "HilRcInputsRaw.Chan10Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN11_RAW = "HilRcInputsRaw.Chan11Raw"
	HIL_RC_INPUTS_RAW_FIELD_CHAN12_RAW = "HilRcInputsRaw.Chan12Raw"
	HIL_RC_INPUTS_RAW_FIELD_RSSI       = "HilRcInputsRaw.Rssi"
)

// ToMap (generated function)
func (m *HilRcInputsRaw) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIL_RC_INPUTS_RAW_FIELD_TIME_USEC:  m.TimeUsec,
		HIL_RC_INPUTS_RAW_FIELD_CHAN1_RAW:  m.Chan1Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN2_RAW:  m.Chan2Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN3_RAW:  m.Chan3Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN4_RAW:  m.Chan4Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN5_RAW:  m.Chan5Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN6_RAW:  m.Chan6Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN7_RAW:  m.Chan7Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN8_RAW:  m.Chan8Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN9_RAW:  m.Chan9Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN10_RAW: m.Chan10Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN11_RAW: m.Chan11Raw,
		HIL_RC_INPUTS_RAW_FIELD_CHAN12_RAW: m.Chan12Raw,
		HIL_RC_INPUTS_RAW_FIELD_RSSI:       m.Rssi,
	}
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
	TimeUsec uint64        `json:"TimeUsec" `          // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Flags    uint64        `json:"Flags" `             // Flags as bitfield, 1: indicate simulation using lockstep.
	Controls []float32     `json:"Controls" len:"16" ` // Control outputs -1 .. 1. Channel assignment depends on the simulated hardware.
	Mode     MAV_MODE_FLAG `json:"Mode" `              // System mode. Includes arming state.
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

const (
	HIL_ACTUATOR_CONTROLS_FIELD_TIME_USEC = "HilActuatorControls.TimeUsec"
	HIL_ACTUATOR_CONTROLS_FIELD_FLAGS     = "HilActuatorControls.Flags"
	HIL_ACTUATOR_CONTROLS_FIELD_CONTROLS  = "HilActuatorControls.Controls"
	HIL_ACTUATOR_CONTROLS_FIELD_MODE      = "HilActuatorControls.Mode"
)

// ToMap (generated function)
func (m *HilActuatorControls) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIL_ACTUATOR_CONTROLS_FIELD_TIME_USEC: m.TimeUsec,
		HIL_ACTUATOR_CONTROLS_FIELD_FLAGS:     m.Flags,
		HIL_ACTUATOR_CONTROLS_FIELD_CONTROLS:  m.Controls,
		HIL_ACTUATOR_CONTROLS_FIELD_MODE:      m.Mode,
	}
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
	TimeUsec       uint64  `json:"TimeUsec" `       // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	FlowCompMX     float32 `json:"FlowCompMX" `     // [ m/s ] Flow in x-sensor direction, angular-speed compensated
	FlowCompMY     float32 `json:"FlowCompMY" `     // [ m/s ] Flow in y-sensor direction, angular-speed compensated
	GroundDistance float32 `json:"GroundDistance" ` // [ m ] Ground distance. Positive value: distance known. Negative value: Unknown distance
	FlowX          int16   `json:"FlowX" `          // [ dpix ] Flow in x-sensor direction
	FlowY          int16   `json:"FlowY" `          // [ dpix ] Flow in y-sensor direction
	SensorID       uint8   `json:"SensorID" `       // Sensor ID
	Quality        uint8   `json:"Quality" `        // Optical flow quality / confidence. 0: bad, 255: maximum quality
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

const (
	OPTICAL_FLOW_FIELD_TIME_USEC       = "OpticalFlow.TimeUsec"
	OPTICAL_FLOW_FIELD_FLOW_COMP_M_X   = "OpticalFlow.FlowCompMX"
	OPTICAL_FLOW_FIELD_FLOW_COMP_M_Y   = "OpticalFlow.FlowCompMY"
	OPTICAL_FLOW_FIELD_GROUND_DISTANCE = "OpticalFlow.GroundDistance"
	OPTICAL_FLOW_FIELD_FLOW_X          = "OpticalFlow.FlowX"
	OPTICAL_FLOW_FIELD_FLOW_Y          = "OpticalFlow.FlowY"
	OPTICAL_FLOW_FIELD_SENSOR_ID       = "OpticalFlow.SensorID"
	OPTICAL_FLOW_FIELD_QUALITY         = "OpticalFlow.Quality"
)

// ToMap (generated function)
func (m *OpticalFlow) Dict() map[string]interface{} {
	return map[string]interface{}{
		OPTICAL_FLOW_FIELD_TIME_USEC:       m.TimeUsec,
		OPTICAL_FLOW_FIELD_FLOW_COMP_M_X:   m.FlowCompMX,
		OPTICAL_FLOW_FIELD_FLOW_COMP_M_Y:   m.FlowCompMY,
		OPTICAL_FLOW_FIELD_GROUND_DISTANCE: m.GroundDistance,
		OPTICAL_FLOW_FIELD_FLOW_X:          m.FlowX,
		OPTICAL_FLOW_FIELD_FLOW_Y:          m.FlowY,
		OPTICAL_FLOW_FIELD_SENSOR_ID:       m.SensorID,
		OPTICAL_FLOW_FIELD_QUALITY:         m.Quality,
	}
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
	Usec  uint64  `json:"Usec" `  // [ us ] Timestamp (UNIX time or since system boot)
	X     float32 `json:"X" `     // [ m ] Global X position
	Y     float32 `json:"Y" `     // [ m ] Global Y position
	Z     float32 `json:"Z" `     // [ m ] Global Z position
	Roll  float32 `json:"Roll" `  // [ rad ] Roll angle
	Pitch float32 `json:"Pitch" ` // [ rad ] Pitch angle
	Yaw   float32 `json:"Yaw" `   // [ rad ] Yaw angle
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

const (
	GLOBAL_VISION_POSITION_ESTIMATE_FIELD_USEC  = "GlobalVisionPositionEstimate.Usec"
	GLOBAL_VISION_POSITION_ESTIMATE_FIELD_X     = "GlobalVisionPositionEstimate.X"
	GLOBAL_VISION_POSITION_ESTIMATE_FIELD_Y     = "GlobalVisionPositionEstimate.Y"
	GLOBAL_VISION_POSITION_ESTIMATE_FIELD_Z     = "GlobalVisionPositionEstimate.Z"
	GLOBAL_VISION_POSITION_ESTIMATE_FIELD_ROLL  = "GlobalVisionPositionEstimate.Roll"
	GLOBAL_VISION_POSITION_ESTIMATE_FIELD_PITCH = "GlobalVisionPositionEstimate.Pitch"
	GLOBAL_VISION_POSITION_ESTIMATE_FIELD_YAW   = "GlobalVisionPositionEstimate.Yaw"
)

// ToMap (generated function)
func (m *GlobalVisionPositionEstimate) Dict() map[string]interface{} {
	return map[string]interface{}{
		GLOBAL_VISION_POSITION_ESTIMATE_FIELD_USEC:  m.Usec,
		GLOBAL_VISION_POSITION_ESTIMATE_FIELD_X:     m.X,
		GLOBAL_VISION_POSITION_ESTIMATE_FIELD_Y:     m.Y,
		GLOBAL_VISION_POSITION_ESTIMATE_FIELD_Z:     m.Z,
		GLOBAL_VISION_POSITION_ESTIMATE_FIELD_ROLL:  m.Roll,
		GLOBAL_VISION_POSITION_ESTIMATE_FIELD_PITCH: m.Pitch,
		GLOBAL_VISION_POSITION_ESTIMATE_FIELD_YAW:   m.Yaw,
	}
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
	Usec  uint64  `json:"Usec" `  // [ us ] Timestamp (UNIX time or time since system boot)
	X     float32 `json:"X" `     // [ m ] Local X position
	Y     float32 `json:"Y" `     // [ m ] Local Y position
	Z     float32 `json:"Z" `     // [ m ] Local Z position
	Roll  float32 `json:"Roll" `  // [ rad ] Roll angle
	Pitch float32 `json:"Pitch" ` // [ rad ] Pitch angle
	Yaw   float32 `json:"Yaw" `   // [ rad ] Yaw angle
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

const (
	VISION_POSITION_ESTIMATE_FIELD_USEC  = "VisionPositionEstimate.Usec"
	VISION_POSITION_ESTIMATE_FIELD_X     = "VisionPositionEstimate.X"
	VISION_POSITION_ESTIMATE_FIELD_Y     = "VisionPositionEstimate.Y"
	VISION_POSITION_ESTIMATE_FIELD_Z     = "VisionPositionEstimate.Z"
	VISION_POSITION_ESTIMATE_FIELD_ROLL  = "VisionPositionEstimate.Roll"
	VISION_POSITION_ESTIMATE_FIELD_PITCH = "VisionPositionEstimate.Pitch"
	VISION_POSITION_ESTIMATE_FIELD_YAW   = "VisionPositionEstimate.Yaw"
)

// ToMap (generated function)
func (m *VisionPositionEstimate) Dict() map[string]interface{} {
	return map[string]interface{}{
		VISION_POSITION_ESTIMATE_FIELD_USEC:  m.Usec,
		VISION_POSITION_ESTIMATE_FIELD_X:     m.X,
		VISION_POSITION_ESTIMATE_FIELD_Y:     m.Y,
		VISION_POSITION_ESTIMATE_FIELD_Z:     m.Z,
		VISION_POSITION_ESTIMATE_FIELD_ROLL:  m.Roll,
		VISION_POSITION_ESTIMATE_FIELD_PITCH: m.Pitch,
		VISION_POSITION_ESTIMATE_FIELD_YAW:   m.Yaw,
	}
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
	Usec uint64  `json:"Usec" ` // [ us ] Timestamp (UNIX time or time since system boot)
	X    float32 `json:"X" `    // [ m/s ] Global X speed
	Y    float32 `json:"Y" `    // [ m/s ] Global Y speed
	Z    float32 `json:"Z" `    // [ m/s ] Global Z speed
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

const (
	VISION_SPEED_ESTIMATE_FIELD_USEC = "VisionSpeedEstimate.Usec"
	VISION_SPEED_ESTIMATE_FIELD_X    = "VisionSpeedEstimate.X"
	VISION_SPEED_ESTIMATE_FIELD_Y    = "VisionSpeedEstimate.Y"
	VISION_SPEED_ESTIMATE_FIELD_Z    = "VisionSpeedEstimate.Z"
)

// ToMap (generated function)
func (m *VisionSpeedEstimate) Dict() map[string]interface{} {
	return map[string]interface{}{
		VISION_SPEED_ESTIMATE_FIELD_USEC: m.Usec,
		VISION_SPEED_ESTIMATE_FIELD_X:    m.X,
		VISION_SPEED_ESTIMATE_FIELD_Y:    m.Y,
		VISION_SPEED_ESTIMATE_FIELD_Z:    m.Z,
	}
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
	Usec  uint64  `json:"Usec" `  // [ us ] Timestamp (UNIX time or time since system boot)
	X     float32 `json:"X" `     // [ m ] Global X position
	Y     float32 `json:"Y" `     // [ m ] Global Y position
	Z     float32 `json:"Z" `     // [ m ] Global Z position
	Roll  float32 `json:"Roll" `  // [ rad ] Roll angle
	Pitch float32 `json:"Pitch" ` // [ rad ] Pitch angle
	Yaw   float32 `json:"Yaw" `   // [ rad ] Yaw angle
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

const (
	VICON_POSITION_ESTIMATE_FIELD_USEC  = "ViconPositionEstimate.Usec"
	VICON_POSITION_ESTIMATE_FIELD_X     = "ViconPositionEstimate.X"
	VICON_POSITION_ESTIMATE_FIELD_Y     = "ViconPositionEstimate.Y"
	VICON_POSITION_ESTIMATE_FIELD_Z     = "ViconPositionEstimate.Z"
	VICON_POSITION_ESTIMATE_FIELD_ROLL  = "ViconPositionEstimate.Roll"
	VICON_POSITION_ESTIMATE_FIELD_PITCH = "ViconPositionEstimate.Pitch"
	VICON_POSITION_ESTIMATE_FIELD_YAW   = "ViconPositionEstimate.Yaw"
)

// ToMap (generated function)
func (m *ViconPositionEstimate) Dict() map[string]interface{} {
	return map[string]interface{}{
		VICON_POSITION_ESTIMATE_FIELD_USEC:  m.Usec,
		VICON_POSITION_ESTIMATE_FIELD_X:     m.X,
		VICON_POSITION_ESTIMATE_FIELD_Y:     m.Y,
		VICON_POSITION_ESTIMATE_FIELD_Z:     m.Z,
		VICON_POSITION_ESTIMATE_FIELD_ROLL:  m.Roll,
		VICON_POSITION_ESTIMATE_FIELD_PITCH: m.Pitch,
		VICON_POSITION_ESTIMATE_FIELD_YAW:   m.Yaw,
	}
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
	TimeUsec      uint64  `json:"TimeUsec" `      // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Xacc          float32 `json:"Xacc" `          // [ m/s/s ] X acceleration
	Yacc          float32 `json:"Yacc" `          // [ m/s/s ] Y acceleration
	Zacc          float32 `json:"Zacc" `          // [ m/s/s ] Z acceleration
	Xgyro         float32 `json:"Xgyro" `         // [ rad/s ] Angular speed around X axis
	Ygyro         float32 `json:"Ygyro" `         // [ rad/s ] Angular speed around Y axis
	Zgyro         float32 `json:"Zgyro" `         // [ rad/s ] Angular speed around Z axis
	Xmag          float32 `json:"Xmag" `          // [ gauss ] X Magnetic field
	Ymag          float32 `json:"Ymag" `          // [ gauss ] Y Magnetic field
	Zmag          float32 `json:"Zmag" `          // [ gauss ] Z Magnetic field
	AbsPressure   float32 `json:"AbsPressure" `   // [ hPa ] Absolute pressure
	DiffPressure  float32 `json:"DiffPressure" `  // [ hPa ] Differential pressure
	PressureAlt   float32 `json:"PressureAlt" `   // Altitude calculated from pressure
	Temperature   float32 `json:"Temperature" `   // [ degC ] Temperature
	FieldsUpdated uint16  `json:"FieldsUpdated" ` // Bitmap for fields that have updated since last message, bit 0 = xacc, bit 12: temperature
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

const (
	HIGHRES_IMU_FIELD_TIME_USEC      = "HighresImu.TimeUsec"
	HIGHRES_IMU_FIELD_XACC           = "HighresImu.Xacc"
	HIGHRES_IMU_FIELD_YACC           = "HighresImu.Yacc"
	HIGHRES_IMU_FIELD_ZACC           = "HighresImu.Zacc"
	HIGHRES_IMU_FIELD_XGYRO          = "HighresImu.Xgyro"
	HIGHRES_IMU_FIELD_YGYRO          = "HighresImu.Ygyro"
	HIGHRES_IMU_FIELD_ZGYRO          = "HighresImu.Zgyro"
	HIGHRES_IMU_FIELD_XMAG           = "HighresImu.Xmag"
	HIGHRES_IMU_FIELD_YMAG           = "HighresImu.Ymag"
	HIGHRES_IMU_FIELD_ZMAG           = "HighresImu.Zmag"
	HIGHRES_IMU_FIELD_ABS_PRESSURE   = "HighresImu.AbsPressure"
	HIGHRES_IMU_FIELD_DIFF_PRESSURE  = "HighresImu.DiffPressure"
	HIGHRES_IMU_FIELD_PRESSURE_ALT   = "HighresImu.PressureAlt"
	HIGHRES_IMU_FIELD_TEMPERATURE    = "HighresImu.Temperature"
	HIGHRES_IMU_FIELD_FIELDS_UPDATED = "HighresImu.FieldsUpdated"
)

// ToMap (generated function)
func (m *HighresImu) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIGHRES_IMU_FIELD_TIME_USEC:      m.TimeUsec,
		HIGHRES_IMU_FIELD_XACC:           m.Xacc,
		HIGHRES_IMU_FIELD_YACC:           m.Yacc,
		HIGHRES_IMU_FIELD_ZACC:           m.Zacc,
		HIGHRES_IMU_FIELD_XGYRO:          m.Xgyro,
		HIGHRES_IMU_FIELD_YGYRO:          m.Ygyro,
		HIGHRES_IMU_FIELD_ZGYRO:          m.Zgyro,
		HIGHRES_IMU_FIELD_XMAG:           m.Xmag,
		HIGHRES_IMU_FIELD_YMAG:           m.Ymag,
		HIGHRES_IMU_FIELD_ZMAG:           m.Zmag,
		HIGHRES_IMU_FIELD_ABS_PRESSURE:   m.AbsPressure,
		HIGHRES_IMU_FIELD_DIFF_PRESSURE:  m.DiffPressure,
		HIGHRES_IMU_FIELD_PRESSURE_ALT:   m.PressureAlt,
		HIGHRES_IMU_FIELD_TEMPERATURE:    m.Temperature,
		HIGHRES_IMU_FIELD_FIELDS_UPDATED: m.FieldsUpdated,
	}
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
	TimeUsec            uint64  `json:"TimeUsec" `            // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	IntegrationTimeUs   uint32  `json:"IntegrationTimeUs" `   // [ us ] Integration time. Divide integrated_x and integrated_y by the integration time to obtain average flow. The integration time also indicates the.
	IntegratedX         float32 `json:"IntegratedX" `         // [ rad ] Flow around X axis (Sensor RH rotation about the X axis induces a positive flow. Sensor linear motion along the positive Y axis induces a negative flow.)
	IntegratedY         float32 `json:"IntegratedY" `         // [ rad ] Flow around Y axis (Sensor RH rotation about the Y axis induces a positive flow. Sensor linear motion along the positive X axis induces a positive flow.)
	IntegratedXgyro     float32 `json:"IntegratedXgyro" `     // [ rad ] RH rotation around X axis
	IntegratedYgyro     float32 `json:"IntegratedYgyro" `     // [ rad ] RH rotation around Y axis
	IntegratedZgyro     float32 `json:"IntegratedZgyro" `     // [ rad ] RH rotation around Z axis
	TimeDeltaDistanceUs uint32  `json:"TimeDeltaDistanceUs" ` // [ us ] Time since the distance was sampled.
	Distance            float32 `json:"Distance" `            // [ m ] Distance to the center of the flow field. Positive value (including zero): distance known. Negative value: Unknown distance.
	Temperature         int16   `json:"Temperature" `         // [ cdegC ] Temperature
	SensorID            uint8   `json:"SensorID" `            // Sensor ID
	Quality             uint8   `json:"Quality" `             // Optical flow quality / confidence. 0: no valid flow, 255: maximum quality
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

const (
	OPTICAL_FLOW_RAD_FIELD_TIME_USEC              = "OpticalFlowRad.TimeUsec"
	OPTICAL_FLOW_RAD_FIELD_INTEGRATION_TIME_US    = "OpticalFlowRad.IntegrationTimeUs"
	OPTICAL_FLOW_RAD_FIELD_INTEGRATED_X           = "OpticalFlowRad.IntegratedX"
	OPTICAL_FLOW_RAD_FIELD_INTEGRATED_Y           = "OpticalFlowRad.IntegratedY"
	OPTICAL_FLOW_RAD_FIELD_INTEGRATED_XGYRO       = "OpticalFlowRad.IntegratedXgyro"
	OPTICAL_FLOW_RAD_FIELD_INTEGRATED_YGYRO       = "OpticalFlowRad.IntegratedYgyro"
	OPTICAL_FLOW_RAD_FIELD_INTEGRATED_ZGYRO       = "OpticalFlowRad.IntegratedZgyro"
	OPTICAL_FLOW_RAD_FIELD_TIME_DELTA_DISTANCE_US = "OpticalFlowRad.TimeDeltaDistanceUs"
	OPTICAL_FLOW_RAD_FIELD_DISTANCE               = "OpticalFlowRad.Distance"
	OPTICAL_FLOW_RAD_FIELD_TEMPERATURE            = "OpticalFlowRad.Temperature"
	OPTICAL_FLOW_RAD_FIELD_SENSOR_ID              = "OpticalFlowRad.SensorID"
	OPTICAL_FLOW_RAD_FIELD_QUALITY                = "OpticalFlowRad.Quality"
)

// ToMap (generated function)
func (m *OpticalFlowRad) Dict() map[string]interface{} {
	return map[string]interface{}{
		OPTICAL_FLOW_RAD_FIELD_TIME_USEC:              m.TimeUsec,
		OPTICAL_FLOW_RAD_FIELD_INTEGRATION_TIME_US:    m.IntegrationTimeUs,
		OPTICAL_FLOW_RAD_FIELD_INTEGRATED_X:           m.IntegratedX,
		OPTICAL_FLOW_RAD_FIELD_INTEGRATED_Y:           m.IntegratedY,
		OPTICAL_FLOW_RAD_FIELD_INTEGRATED_XGYRO:       m.IntegratedXgyro,
		OPTICAL_FLOW_RAD_FIELD_INTEGRATED_YGYRO:       m.IntegratedYgyro,
		OPTICAL_FLOW_RAD_FIELD_INTEGRATED_ZGYRO:       m.IntegratedZgyro,
		OPTICAL_FLOW_RAD_FIELD_TIME_DELTA_DISTANCE_US: m.TimeDeltaDistanceUs,
		OPTICAL_FLOW_RAD_FIELD_DISTANCE:               m.Distance,
		OPTICAL_FLOW_RAD_FIELD_TEMPERATURE:            m.Temperature,
		OPTICAL_FLOW_RAD_FIELD_SENSOR_ID:              m.SensorID,
		OPTICAL_FLOW_RAD_FIELD_QUALITY:                m.Quality,
	}
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
	TimeUsec      uint64  `json:"TimeUsec" `      // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Xacc          float32 `json:"Xacc" `          // [ m/s/s ] X acceleration
	Yacc          float32 `json:"Yacc" `          // [ m/s/s ] Y acceleration
	Zacc          float32 `json:"Zacc" `          // [ m/s/s ] Z acceleration
	Xgyro         float32 `json:"Xgyro" `         // [ rad/s ] Angular speed around X axis in body frame
	Ygyro         float32 `json:"Ygyro" `         // [ rad/s ] Angular speed around Y axis in body frame
	Zgyro         float32 `json:"Zgyro" `         // [ rad/s ] Angular speed around Z axis in body frame
	Xmag          float32 `json:"Xmag" `          // [ gauss ] X Magnetic field
	Ymag          float32 `json:"Ymag" `          // [ gauss ] Y Magnetic field
	Zmag          float32 `json:"Zmag" `          // [ gauss ] Z Magnetic field
	AbsPressure   float32 `json:"AbsPressure" `   // [ hPa ] Absolute pressure
	DiffPressure  float32 `json:"DiffPressure" `  // [ hPa ] Differential pressure (airspeed)
	PressureAlt   float32 `json:"PressureAlt" `   // Altitude calculated from pressure
	Temperature   float32 `json:"Temperature" `   // [ degC ] Temperature
	FieldsUpdated uint32  `json:"FieldsUpdated" ` // Bitmap for fields that have updated since last message, bit 0 = xacc, bit 12: temperature, bit 31: full reset of attitude/position/velocities/etc was performed in sim.
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

const (
	HIL_SENSOR_FIELD_TIME_USEC      = "HilSensor.TimeUsec"
	HIL_SENSOR_FIELD_XACC           = "HilSensor.Xacc"
	HIL_SENSOR_FIELD_YACC           = "HilSensor.Yacc"
	HIL_SENSOR_FIELD_ZACC           = "HilSensor.Zacc"
	HIL_SENSOR_FIELD_XGYRO          = "HilSensor.Xgyro"
	HIL_SENSOR_FIELD_YGYRO          = "HilSensor.Ygyro"
	HIL_SENSOR_FIELD_ZGYRO          = "HilSensor.Zgyro"
	HIL_SENSOR_FIELD_XMAG           = "HilSensor.Xmag"
	HIL_SENSOR_FIELD_YMAG           = "HilSensor.Ymag"
	HIL_SENSOR_FIELD_ZMAG           = "HilSensor.Zmag"
	HIL_SENSOR_FIELD_ABS_PRESSURE   = "HilSensor.AbsPressure"
	HIL_SENSOR_FIELD_DIFF_PRESSURE  = "HilSensor.DiffPressure"
	HIL_SENSOR_FIELD_PRESSURE_ALT   = "HilSensor.PressureAlt"
	HIL_SENSOR_FIELD_TEMPERATURE    = "HilSensor.Temperature"
	HIL_SENSOR_FIELD_FIELDS_UPDATED = "HilSensor.FieldsUpdated"
)

// ToMap (generated function)
func (m *HilSensor) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIL_SENSOR_FIELD_TIME_USEC:      m.TimeUsec,
		HIL_SENSOR_FIELD_XACC:           m.Xacc,
		HIL_SENSOR_FIELD_YACC:           m.Yacc,
		HIL_SENSOR_FIELD_ZACC:           m.Zacc,
		HIL_SENSOR_FIELD_XGYRO:          m.Xgyro,
		HIL_SENSOR_FIELD_YGYRO:          m.Ygyro,
		HIL_SENSOR_FIELD_ZGYRO:          m.Zgyro,
		HIL_SENSOR_FIELD_XMAG:           m.Xmag,
		HIL_SENSOR_FIELD_YMAG:           m.Ymag,
		HIL_SENSOR_FIELD_ZMAG:           m.Zmag,
		HIL_SENSOR_FIELD_ABS_PRESSURE:   m.AbsPressure,
		HIL_SENSOR_FIELD_DIFF_PRESSURE:  m.DiffPressure,
		HIL_SENSOR_FIELD_PRESSURE_ALT:   m.PressureAlt,
		HIL_SENSOR_FIELD_TEMPERATURE:    m.Temperature,
		HIL_SENSOR_FIELD_FIELDS_UPDATED: m.FieldsUpdated,
	}
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
	Q1         float32 `json:"Q1" `         // True attitude quaternion component 1, w (1 in null-rotation)
	Q2         float32 `json:"Q2" `         // True attitude quaternion component 2, x (0 in null-rotation)
	Q3         float32 `json:"Q3" `         // True attitude quaternion component 3, y (0 in null-rotation)
	Q4         float32 `json:"Q4" `         // True attitude quaternion component 4, z (0 in null-rotation)
	Roll       float32 `json:"Roll" `       // Attitude roll expressed as Euler angles, not recommended except for human-readable outputs
	Pitch      float32 `json:"Pitch" `      // Attitude pitch expressed as Euler angles, not recommended except for human-readable outputs
	Yaw        float32 `json:"Yaw" `        // Attitude yaw expressed as Euler angles, not recommended except for human-readable outputs
	Xacc       float32 `json:"Xacc" `       // [ m/s/s ] X acceleration
	Yacc       float32 `json:"Yacc" `       // [ m/s/s ] Y acceleration
	Zacc       float32 `json:"Zacc" `       // [ m/s/s ] Z acceleration
	Xgyro      float32 `json:"Xgyro" `      // [ rad/s ] Angular speed around X axis
	Ygyro      float32 `json:"Ygyro" `      // [ rad/s ] Angular speed around Y axis
	Zgyro      float32 `json:"Zgyro" `      // [ rad/s ] Angular speed around Z axis
	Lat        float32 `json:"Lat" `        // [ deg ] Latitude
	Lon        float32 `json:"Lon" `        // [ deg ] Longitude
	Alt        float32 `json:"Alt" `        // [ m ] Altitude
	StdDevHorz float32 `json:"StdDevHorz" ` // Horizontal position standard deviation
	StdDevVert float32 `json:"StdDevVert" ` // Vertical position standard deviation
	Vn         float32 `json:"Vn" `         // [ m/s ] True velocity in north direction in earth-fixed NED frame
	Ve         float32 `json:"Ve" `         // [ m/s ] True velocity in east direction in earth-fixed NED frame
	Vd         float32 `json:"Vd" `         // [ m/s ] True velocity in down direction in earth-fixed NED frame
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

const (
	SIM_STATE_FIELD_Q1           = "SimState.Q1"
	SIM_STATE_FIELD_Q2           = "SimState.Q2"
	SIM_STATE_FIELD_Q3           = "SimState.Q3"
	SIM_STATE_FIELD_Q4           = "SimState.Q4"
	SIM_STATE_FIELD_ROLL         = "SimState.Roll"
	SIM_STATE_FIELD_PITCH        = "SimState.Pitch"
	SIM_STATE_FIELD_YAW          = "SimState.Yaw"
	SIM_STATE_FIELD_XACC         = "SimState.Xacc"
	SIM_STATE_FIELD_YACC         = "SimState.Yacc"
	SIM_STATE_FIELD_ZACC         = "SimState.Zacc"
	SIM_STATE_FIELD_XGYRO        = "SimState.Xgyro"
	SIM_STATE_FIELD_YGYRO        = "SimState.Ygyro"
	SIM_STATE_FIELD_ZGYRO        = "SimState.Zgyro"
	SIM_STATE_FIELD_LAT          = "SimState.Lat"
	SIM_STATE_FIELD_LON          = "SimState.Lon"
	SIM_STATE_FIELD_ALT          = "SimState.Alt"
	SIM_STATE_FIELD_STD_DEV_HORZ = "SimState.StdDevHorz"
	SIM_STATE_FIELD_STD_DEV_VERT = "SimState.StdDevVert"
	SIM_STATE_FIELD_VN           = "SimState.Vn"
	SIM_STATE_FIELD_VE           = "SimState.Ve"
	SIM_STATE_FIELD_VD           = "SimState.Vd"
)

// ToMap (generated function)
func (m *SimState) Dict() map[string]interface{} {
	return map[string]interface{}{
		SIM_STATE_FIELD_Q1:           m.Q1,
		SIM_STATE_FIELD_Q2:           m.Q2,
		SIM_STATE_FIELD_Q3:           m.Q3,
		SIM_STATE_FIELD_Q4:           m.Q4,
		SIM_STATE_FIELD_ROLL:         m.Roll,
		SIM_STATE_FIELD_PITCH:        m.Pitch,
		SIM_STATE_FIELD_YAW:          m.Yaw,
		SIM_STATE_FIELD_XACC:         m.Xacc,
		SIM_STATE_FIELD_YACC:         m.Yacc,
		SIM_STATE_FIELD_ZACC:         m.Zacc,
		SIM_STATE_FIELD_XGYRO:        m.Xgyro,
		SIM_STATE_FIELD_YGYRO:        m.Ygyro,
		SIM_STATE_FIELD_ZGYRO:        m.Zgyro,
		SIM_STATE_FIELD_LAT:          m.Lat,
		SIM_STATE_FIELD_LON:          m.Lon,
		SIM_STATE_FIELD_ALT:          m.Alt,
		SIM_STATE_FIELD_STD_DEV_HORZ: m.StdDevHorz,
		SIM_STATE_FIELD_STD_DEV_VERT: m.StdDevVert,
		SIM_STATE_FIELD_VN:           m.Vn,
		SIM_STATE_FIELD_VE:           m.Ve,
		SIM_STATE_FIELD_VD:           m.Vd,
	}
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
	Rxerrors uint16 `json:"Rxerrors" ` // Count of radio packet receive errors (since boot).
	Fixed    uint16 `json:"Fixed" `    // Count of error corrected radio packets (since boot).
	Rssi     uint8  `json:"Rssi" `     // Local (message sender) recieved signal strength indication in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
	Remrssi  uint8  `json:"Remrssi" `  // Remote (message receiver) signal strength indication in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
	Txbuf    uint8  `json:"Txbuf" `    // [ % ] Remaining free transmitter buffer space.
	Noise    uint8  `json:"Noise" `    // Local background noise level. These are device dependent RSSI values (scale as approx 2x dB on SiK radios). Values: [0-254], 255: invalid/unknown.
	Remnoise uint8  `json:"Remnoise" ` // Remote background noise level. These are device dependent RSSI values (scale as approx 2x dB on SiK radios). Values: [0-254], 255: invalid/unknown.
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

const (
	RADIO_STATUS_FIELD_RXERRORS = "RadioStatus.Rxerrors"
	RADIO_STATUS_FIELD_FIXED    = "RadioStatus.Fixed"
	RADIO_STATUS_FIELD_RSSI     = "RadioStatus.Rssi"
	RADIO_STATUS_FIELD_REMRSSI  = "RadioStatus.Remrssi"
	RADIO_STATUS_FIELD_TXBUF    = "RadioStatus.Txbuf"
	RADIO_STATUS_FIELD_NOISE    = "RadioStatus.Noise"
	RADIO_STATUS_FIELD_REMNOISE = "RadioStatus.Remnoise"
)

// ToMap (generated function)
func (m *RadioStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		RADIO_STATUS_FIELD_RXERRORS: m.Rxerrors,
		RADIO_STATUS_FIELD_FIXED:    m.Fixed,
		RADIO_STATUS_FIELD_RSSI:     m.Rssi,
		RADIO_STATUS_FIELD_REMRSSI:  m.Remrssi,
		RADIO_STATUS_FIELD_TXBUF:    m.Txbuf,
		RADIO_STATUS_FIELD_NOISE:    m.Noise,
		RADIO_STATUS_FIELD_REMNOISE: m.Remnoise,
	}
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
	TargetNetwork   uint8   `json:"TargetNetwork" `     // Network ID (0 for broadcast)
	TargetSystem    uint8   `json:"TargetSystem" `      // System ID (0 for broadcast)
	TargetComponent uint8   `json:"TargetComponent" `   // Component ID (0 for broadcast)
	Payload         []uint8 `json:"Payload" len:"251" ` // Variable length payload. The length is defined by the remaining message length when subtracting the header and other fields.  The entire content of this block is opaque unless you understand any the encoding message_type.  The particular encoding used can be extension specific and might not always be documented as part of the mavlink specification.
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

const (
	FILE_TRANSFER_PROTOCOL_FIELD_TARGET_NETWORK   = "FileTransferProtocol.TargetNetwork"
	FILE_TRANSFER_PROTOCOL_FIELD_TARGET_SYSTEM    = "FileTransferProtocol.TargetSystem"
	FILE_TRANSFER_PROTOCOL_FIELD_TARGET_COMPONENT = "FileTransferProtocol.TargetComponent"
	FILE_TRANSFER_PROTOCOL_FIELD_PAYLOAD          = "FileTransferProtocol.Payload"
)

// ToMap (generated function)
func (m *FileTransferProtocol) Dict() map[string]interface{} {
	return map[string]interface{}{
		FILE_TRANSFER_PROTOCOL_FIELD_TARGET_NETWORK:   m.TargetNetwork,
		FILE_TRANSFER_PROTOCOL_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		FILE_TRANSFER_PROTOCOL_FIELD_TARGET_COMPONENT: m.TargetComponent,
		FILE_TRANSFER_PROTOCOL_FIELD_PAYLOAD:          m.Payload,
	}
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
	Tc1 int64 `json:"Tc1" ` // Time sync timestamp 1
	Ts1 int64 `json:"Ts1" ` // Time sync timestamp 2
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

const (
	TIMESYNC_FIELD_TC1 = "Timesync.Tc1"
	TIMESYNC_FIELD_TS1 = "Timesync.Ts1"
)

// ToMap (generated function)
func (m *Timesync) Dict() map[string]interface{} {
	return map[string]interface{}{
		TIMESYNC_FIELD_TC1: m.Tc1,
		TIMESYNC_FIELD_TS1: m.Ts1,
	}
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
	TimeUsec uint64 `json:"TimeUsec" ` // [ us ] Timestamp for image frame (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Seq      uint32 `json:"Seq" `      // Image frame sequence
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

const (
	CAMERA_TRIGGER_FIELD_TIME_USEC = "CameraTrigger.TimeUsec"
	CAMERA_TRIGGER_FIELD_SEQ       = "CameraTrigger.Seq"
)

// ToMap (generated function)
func (m *CameraTrigger) Dict() map[string]interface{} {
	return map[string]interface{}{
		CAMERA_TRIGGER_FIELD_TIME_USEC: m.TimeUsec,
		CAMERA_TRIGGER_FIELD_SEQ:       m.Seq,
	}
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
	TimeUsec          uint64 `json:"TimeUsec" `          // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Lat               int32  `json:"Lat" `               // [ degE7 ] Latitude (WGS84)
	Lon               int32  `json:"Lon" `               // [ degE7 ] Longitude (WGS84)
	Alt               int32  `json:"Alt" `               // [ mm ] Altitude (MSL). Positive for up.
	Eph               uint16 `json:"Eph" `               // GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX
	Epv               uint16 `json:"Epv" `               // GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX
	Vel               uint16 `json:"Vel" `               // [ cm/s ] GPS ground speed. If unknown, set to: 65535
	Vn                int16  `json:"Vn" `                // [ cm/s ] GPS velocity in north direction in earth-fixed NED frame
	Ve                int16  `json:"Ve" `                // [ cm/s ] GPS velocity in east direction in earth-fixed NED frame
	Vd                int16  `json:"Vd" `                // [ cm/s ] GPS velocity in down direction in earth-fixed NED frame
	Cog               uint16 `json:"Cog" `               // [ cdeg ] Course over ground (NOT heading, but direction of movement), 0.0..359.99 degrees. If unknown, set to: 65535
	FixType           uint8  `json:"FixType" `           // 0-1: no fix, 2: 2D fix, 3: 3D fix. Some applications will not use the value of this field unless it is at least two, so always correctly fill in the fix.
	SatellitesVisible uint8  `json:"SatellitesVisible" ` // Number of satellites visible. If unknown, set to 255
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

const (
	HIL_GPS_FIELD_TIME_USEC          = "HilGps.TimeUsec"
	HIL_GPS_FIELD_LAT                = "HilGps.Lat"
	HIL_GPS_FIELD_LON                = "HilGps.Lon"
	HIL_GPS_FIELD_ALT                = "HilGps.Alt"
	HIL_GPS_FIELD_EPH                = "HilGps.Eph"
	HIL_GPS_FIELD_EPV                = "HilGps.Epv"
	HIL_GPS_FIELD_VEL                = "HilGps.Vel"
	HIL_GPS_FIELD_VN                 = "HilGps.Vn"
	HIL_GPS_FIELD_VE                 = "HilGps.Ve"
	HIL_GPS_FIELD_VD                 = "HilGps.Vd"
	HIL_GPS_FIELD_COG                = "HilGps.Cog"
	HIL_GPS_FIELD_FIX_TYPE           = "HilGps.FixType"
	HIL_GPS_FIELD_SATELLITES_VISIBLE = "HilGps.SatellitesVisible"
)

// ToMap (generated function)
func (m *HilGps) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIL_GPS_FIELD_TIME_USEC:          m.TimeUsec,
		HIL_GPS_FIELD_LAT:                m.Lat,
		HIL_GPS_FIELD_LON:                m.Lon,
		HIL_GPS_FIELD_ALT:                m.Alt,
		HIL_GPS_FIELD_EPH:                m.Eph,
		HIL_GPS_FIELD_EPV:                m.Epv,
		HIL_GPS_FIELD_VEL:                m.Vel,
		HIL_GPS_FIELD_VN:                 m.Vn,
		HIL_GPS_FIELD_VE:                 m.Ve,
		HIL_GPS_FIELD_VD:                 m.Vd,
		HIL_GPS_FIELD_COG:                m.Cog,
		HIL_GPS_FIELD_FIX_TYPE:           m.FixType,
		HIL_GPS_FIELD_SATELLITES_VISIBLE: m.SatellitesVisible,
	}
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
	TimeUsec            uint64  `json:"TimeUsec" `            // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	IntegrationTimeUs   uint32  `json:"IntegrationTimeUs" `   // [ us ] Integration time. Divide integrated_x and integrated_y by the integration time to obtain average flow. The integration time also indicates the.
	IntegratedX         float32 `json:"IntegratedX" `         // [ rad ] Flow in radians around X axis (Sensor RH rotation about the X axis induces a positive flow. Sensor linear motion along the positive Y axis induces a negative flow.)
	IntegratedY         float32 `json:"IntegratedY" `         // [ rad ] Flow in radians around Y axis (Sensor RH rotation about the Y axis induces a positive flow. Sensor linear motion along the positive X axis induces a positive flow.)
	IntegratedXgyro     float32 `json:"IntegratedXgyro" `     // [ rad ] RH rotation around X axis
	IntegratedYgyro     float32 `json:"IntegratedYgyro" `     // [ rad ] RH rotation around Y axis
	IntegratedZgyro     float32 `json:"IntegratedZgyro" `     // [ rad ] RH rotation around Z axis
	TimeDeltaDistanceUs uint32  `json:"TimeDeltaDistanceUs" ` // [ us ] Time since the distance was sampled.
	Distance            float32 `json:"Distance" `            // [ m ] Distance to the center of the flow field. Positive value (including zero): distance known. Negative value: Unknown distance.
	Temperature         int16   `json:"Temperature" `         // [ cdegC ] Temperature
	SensorID            uint8   `json:"SensorID" `            // Sensor ID
	Quality             uint8   `json:"Quality" `             // Optical flow quality / confidence. 0: no valid flow, 255: maximum quality
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

const (
	HIL_OPTICAL_FLOW_FIELD_TIME_USEC              = "HilOpticalFlow.TimeUsec"
	HIL_OPTICAL_FLOW_FIELD_INTEGRATION_TIME_US    = "HilOpticalFlow.IntegrationTimeUs"
	HIL_OPTICAL_FLOW_FIELD_INTEGRATED_X           = "HilOpticalFlow.IntegratedX"
	HIL_OPTICAL_FLOW_FIELD_INTEGRATED_Y           = "HilOpticalFlow.IntegratedY"
	HIL_OPTICAL_FLOW_FIELD_INTEGRATED_XGYRO       = "HilOpticalFlow.IntegratedXgyro"
	HIL_OPTICAL_FLOW_FIELD_INTEGRATED_YGYRO       = "HilOpticalFlow.IntegratedYgyro"
	HIL_OPTICAL_FLOW_FIELD_INTEGRATED_ZGYRO       = "HilOpticalFlow.IntegratedZgyro"
	HIL_OPTICAL_FLOW_FIELD_TIME_DELTA_DISTANCE_US = "HilOpticalFlow.TimeDeltaDistanceUs"
	HIL_OPTICAL_FLOW_FIELD_DISTANCE               = "HilOpticalFlow.Distance"
	HIL_OPTICAL_FLOW_FIELD_TEMPERATURE            = "HilOpticalFlow.Temperature"
	HIL_OPTICAL_FLOW_FIELD_SENSOR_ID              = "HilOpticalFlow.SensorID"
	HIL_OPTICAL_FLOW_FIELD_QUALITY                = "HilOpticalFlow.Quality"
)

// ToMap (generated function)
func (m *HilOpticalFlow) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIL_OPTICAL_FLOW_FIELD_TIME_USEC:              m.TimeUsec,
		HIL_OPTICAL_FLOW_FIELD_INTEGRATION_TIME_US:    m.IntegrationTimeUs,
		HIL_OPTICAL_FLOW_FIELD_INTEGRATED_X:           m.IntegratedX,
		HIL_OPTICAL_FLOW_FIELD_INTEGRATED_Y:           m.IntegratedY,
		HIL_OPTICAL_FLOW_FIELD_INTEGRATED_XGYRO:       m.IntegratedXgyro,
		HIL_OPTICAL_FLOW_FIELD_INTEGRATED_YGYRO:       m.IntegratedYgyro,
		HIL_OPTICAL_FLOW_FIELD_INTEGRATED_ZGYRO:       m.IntegratedZgyro,
		HIL_OPTICAL_FLOW_FIELD_TIME_DELTA_DISTANCE_US: m.TimeDeltaDistanceUs,
		HIL_OPTICAL_FLOW_FIELD_DISTANCE:               m.Distance,
		HIL_OPTICAL_FLOW_FIELD_TEMPERATURE:            m.Temperature,
		HIL_OPTICAL_FLOW_FIELD_SENSOR_ID:              m.SensorID,
		HIL_OPTICAL_FLOW_FIELD_QUALITY:                m.Quality,
	}
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
	TimeUsec           uint64    `json:"TimeUsec" `                   // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	AttitudeQuaternion []float32 `json:"AttitudeQuaternion" len:"4" ` // Vehicle attitude expressed as normalized quaternion in w, x, y, z order (with 1 0 0 0 being the null-rotation)
	Rollspeed          float32   `json:"Rollspeed" `                  // [ rad/s ] Body frame roll / phi angular speed
	Pitchspeed         float32   `json:"Pitchspeed" `                 // [ rad/s ] Body frame pitch / theta angular speed
	Yawspeed           float32   `json:"Yawspeed" `                   // [ rad/s ] Body frame yaw / psi angular speed
	Lat                int32     `json:"Lat" `                        // [ degE7 ] Latitude
	Lon                int32     `json:"Lon" `                        // [ degE7 ] Longitude
	Alt                int32     `json:"Alt" `                        // [ mm ] Altitude
	Vx                 int16     `json:"Vx" `                         // [ cm/s ] Ground X Speed (Latitude)
	Vy                 int16     `json:"Vy" `                         // [ cm/s ] Ground Y Speed (Longitude)
	Vz                 int16     `json:"Vz" `                         // [ cm/s ] Ground Z Speed (Altitude)
	IndAirspeed        uint16    `json:"IndAirspeed" `                // [ cm/s ] Indicated airspeed
	TrueAirspeed       uint16    `json:"TrueAirspeed" `               // [ cm/s ] True airspeed
	Xacc               int16     `json:"Xacc" `                       // [ mG ] X acceleration
	Yacc               int16     `json:"Yacc" `                       // [ mG ] Y acceleration
	Zacc               int16     `json:"Zacc" `                       // [ mG ] Z acceleration
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

const (
	HIL_STATE_QUATERNION_FIELD_TIME_USEC           = "HilStateQuaternion.TimeUsec"
	HIL_STATE_QUATERNION_FIELD_ATTITUDE_QUATERNION = "HilStateQuaternion.AttitudeQuaternion"
	HIL_STATE_QUATERNION_FIELD_ROLLSPEED           = "HilStateQuaternion.Rollspeed"
	HIL_STATE_QUATERNION_FIELD_PITCHSPEED          = "HilStateQuaternion.Pitchspeed"
	HIL_STATE_QUATERNION_FIELD_YAWSPEED            = "HilStateQuaternion.Yawspeed"
	HIL_STATE_QUATERNION_FIELD_LAT                 = "HilStateQuaternion.Lat"
	HIL_STATE_QUATERNION_FIELD_LON                 = "HilStateQuaternion.Lon"
	HIL_STATE_QUATERNION_FIELD_ALT                 = "HilStateQuaternion.Alt"
	HIL_STATE_QUATERNION_FIELD_VX                  = "HilStateQuaternion.Vx"
	HIL_STATE_QUATERNION_FIELD_VY                  = "HilStateQuaternion.Vy"
	HIL_STATE_QUATERNION_FIELD_VZ                  = "HilStateQuaternion.Vz"
	HIL_STATE_QUATERNION_FIELD_IND_AIRSPEED        = "HilStateQuaternion.IndAirspeed"
	HIL_STATE_QUATERNION_FIELD_TRUE_AIRSPEED       = "HilStateQuaternion.TrueAirspeed"
	HIL_STATE_QUATERNION_FIELD_XACC                = "HilStateQuaternion.Xacc"
	HIL_STATE_QUATERNION_FIELD_YACC                = "HilStateQuaternion.Yacc"
	HIL_STATE_QUATERNION_FIELD_ZACC                = "HilStateQuaternion.Zacc"
)

// ToMap (generated function)
func (m *HilStateQuaternion) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIL_STATE_QUATERNION_FIELD_TIME_USEC:           m.TimeUsec,
		HIL_STATE_QUATERNION_FIELD_ATTITUDE_QUATERNION: m.AttitudeQuaternion,
		HIL_STATE_QUATERNION_FIELD_ROLLSPEED:           m.Rollspeed,
		HIL_STATE_QUATERNION_FIELD_PITCHSPEED:          m.Pitchspeed,
		HIL_STATE_QUATERNION_FIELD_YAWSPEED:            m.Yawspeed,
		HIL_STATE_QUATERNION_FIELD_LAT:                 m.Lat,
		HIL_STATE_QUATERNION_FIELD_LON:                 m.Lon,
		HIL_STATE_QUATERNION_FIELD_ALT:                 m.Alt,
		HIL_STATE_QUATERNION_FIELD_VX:                  m.Vx,
		HIL_STATE_QUATERNION_FIELD_VY:                  m.Vy,
		HIL_STATE_QUATERNION_FIELD_VZ:                  m.Vz,
		HIL_STATE_QUATERNION_FIELD_IND_AIRSPEED:        m.IndAirspeed,
		HIL_STATE_QUATERNION_FIELD_TRUE_AIRSPEED:       m.TrueAirspeed,
		HIL_STATE_QUATERNION_FIELD_XACC:                m.Xacc,
		HIL_STATE_QUATERNION_FIELD_YACC:                m.Yacc,
		HIL_STATE_QUATERNION_FIELD_ZACC:                m.Zacc,
	}
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
	TimeBootMs uint32 `json:"TimeBootMs" ` // [ ms ] Timestamp (time since system boot).
	Xacc       int16  `json:"Xacc" `       // [ mG ] X acceleration
	Yacc       int16  `json:"Yacc" `       // [ mG ] Y acceleration
	Zacc       int16  `json:"Zacc" `       // [ mG ] Z acceleration
	Xgyro      int16  `json:"Xgyro" `      // [ mrad/s ] Angular speed around X axis
	Ygyro      int16  `json:"Ygyro" `      // [ mrad/s ] Angular speed around Y axis
	Zgyro      int16  `json:"Zgyro" `      // [ mrad/s ] Angular speed around Z axis
	Xmag       int16  `json:"Xmag" `       // [ mgauss ] X Magnetic field
	Ymag       int16  `json:"Ymag" `       // [ mgauss ] Y Magnetic field
	Zmag       int16  `json:"Zmag" `       // [ mgauss ] Z Magnetic field
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

const (
	SCALED_IMU2_FIELD_TIME_BOOT_MS = "ScaledImu2.TimeBootMs"
	SCALED_IMU2_FIELD_XACC         = "ScaledImu2.Xacc"
	SCALED_IMU2_FIELD_YACC         = "ScaledImu2.Yacc"
	SCALED_IMU2_FIELD_ZACC         = "ScaledImu2.Zacc"
	SCALED_IMU2_FIELD_XGYRO        = "ScaledImu2.Xgyro"
	SCALED_IMU2_FIELD_YGYRO        = "ScaledImu2.Ygyro"
	SCALED_IMU2_FIELD_ZGYRO        = "ScaledImu2.Zgyro"
	SCALED_IMU2_FIELD_XMAG         = "ScaledImu2.Xmag"
	SCALED_IMU2_FIELD_YMAG         = "ScaledImu2.Ymag"
	SCALED_IMU2_FIELD_ZMAG         = "ScaledImu2.Zmag"
)

// ToMap (generated function)
func (m *ScaledImu2) Dict() map[string]interface{} {
	return map[string]interface{}{
		SCALED_IMU2_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		SCALED_IMU2_FIELD_XACC:         m.Xacc,
		SCALED_IMU2_FIELD_YACC:         m.Yacc,
		SCALED_IMU2_FIELD_ZACC:         m.Zacc,
		SCALED_IMU2_FIELD_XGYRO:        m.Xgyro,
		SCALED_IMU2_FIELD_YGYRO:        m.Ygyro,
		SCALED_IMU2_FIELD_ZGYRO:        m.Zgyro,
		SCALED_IMU2_FIELD_XMAG:         m.Xmag,
		SCALED_IMU2_FIELD_YMAG:         m.Ymag,
		SCALED_IMU2_FIELD_ZMAG:         m.Zmag,
	}
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
	Start           uint16 `json:"Start" `           // First log id (0 for first available)
	End             uint16 `json:"End" `             // Last log id (0xffff for last available)
	TargetSystem    uint8  `json:"TargetSystem" `    // System ID
	TargetComponent uint8  `json:"TargetComponent" ` // Component ID
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

const (
	LOG_REQUEST_LIST_FIELD_START            = "LogRequestList.Start"
	LOG_REQUEST_LIST_FIELD_END              = "LogRequestList.End"
	LOG_REQUEST_LIST_FIELD_TARGET_SYSTEM    = "LogRequestList.TargetSystem"
	LOG_REQUEST_LIST_FIELD_TARGET_COMPONENT = "LogRequestList.TargetComponent"
)

// ToMap (generated function)
func (m *LogRequestList) Dict() map[string]interface{} {
	return map[string]interface{}{
		LOG_REQUEST_LIST_FIELD_START:            m.Start,
		LOG_REQUEST_LIST_FIELD_END:              m.End,
		LOG_REQUEST_LIST_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		LOG_REQUEST_LIST_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	TimeUtc    uint32 `json:"TimeUtc" `    // [ s ] UTC timestamp of log since 1970, or 0 if not available
	Size       uint32 `json:"Size" `       // [ bytes ] Size of the log (may be approximate)
	ID         uint16 `json:"ID" `         // Log id
	NumLogs    uint16 `json:"NumLogs" `    // Total number of logs
	LastLogNum uint16 `json:"LastLogNum" ` // High log number
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

const (
	LOG_ENTRY_FIELD_TIME_UTC     = "LogEntry.TimeUtc"
	LOG_ENTRY_FIELD_SIZE         = "LogEntry.Size"
	LOG_ENTRY_FIELD_ID           = "LogEntry.ID"
	LOG_ENTRY_FIELD_NUM_LOGS     = "LogEntry.NumLogs"
	LOG_ENTRY_FIELD_LAST_LOG_NUM = "LogEntry.LastLogNum"
)

// ToMap (generated function)
func (m *LogEntry) Dict() map[string]interface{} {
	return map[string]interface{}{
		LOG_ENTRY_FIELD_TIME_UTC:     m.TimeUtc,
		LOG_ENTRY_FIELD_SIZE:         m.Size,
		LOG_ENTRY_FIELD_ID:           m.ID,
		LOG_ENTRY_FIELD_NUM_LOGS:     m.NumLogs,
		LOG_ENTRY_FIELD_LAST_LOG_NUM: m.LastLogNum,
	}
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
	Ofs             uint32 `json:"Ofs" `             // Offset into the log
	Count           uint32 `json:"Count" `           // [ bytes ] Number of bytes
	ID              uint16 `json:"ID" `              // Log id (from LOG_ENTRY reply)
	TargetSystem    uint8  `json:"TargetSystem" `    // System ID
	TargetComponent uint8  `json:"TargetComponent" ` // Component ID
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

const (
	LOG_REQUEST_DATA_FIELD_OFS              = "LogRequestData.Ofs"
	LOG_REQUEST_DATA_FIELD_COUNT            = "LogRequestData.Count"
	LOG_REQUEST_DATA_FIELD_ID               = "LogRequestData.ID"
	LOG_REQUEST_DATA_FIELD_TARGET_SYSTEM    = "LogRequestData.TargetSystem"
	LOG_REQUEST_DATA_FIELD_TARGET_COMPONENT = "LogRequestData.TargetComponent"
)

// ToMap (generated function)
func (m *LogRequestData) Dict() map[string]interface{} {
	return map[string]interface{}{
		LOG_REQUEST_DATA_FIELD_OFS:              m.Ofs,
		LOG_REQUEST_DATA_FIELD_COUNT:            m.Count,
		LOG_REQUEST_DATA_FIELD_ID:               m.ID,
		LOG_REQUEST_DATA_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		LOG_REQUEST_DATA_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	Ofs   uint32  `json:"Ofs" `           // Offset into the log
	ID    uint16  `json:"ID" `            // Log id (from LOG_ENTRY reply)
	Count uint8   `json:"Count" `         // [ bytes ] Number of bytes (zero for end of log)
	Data  []uint8 `json:"Data" len:"90" ` // log data
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

const (
	LOG_DATA_FIELD_OFS   = "LogData.Ofs"
	LOG_DATA_FIELD_ID    = "LogData.ID"
	LOG_DATA_FIELD_COUNT = "LogData.Count"
	LOG_DATA_FIELD_DATA  = "LogData.Data"
)

// ToMap (generated function)
func (m *LogData) Dict() map[string]interface{} {
	return map[string]interface{}{
		LOG_DATA_FIELD_OFS:   m.Ofs,
		LOG_DATA_FIELD_ID:    m.ID,
		LOG_DATA_FIELD_COUNT: m.Count,
		LOG_DATA_FIELD_DATA:  m.Data,
	}
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
	TargetSystem    uint8 `json:"TargetSystem" `    // System ID
	TargetComponent uint8 `json:"TargetComponent" ` // Component ID
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

const (
	LOG_ERASE_FIELD_TARGET_SYSTEM    = "LogErase.TargetSystem"
	LOG_ERASE_FIELD_TARGET_COMPONENT = "LogErase.TargetComponent"
)

// ToMap (generated function)
func (m *LogErase) Dict() map[string]interface{} {
	return map[string]interface{}{
		LOG_ERASE_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		LOG_ERASE_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	TargetSystem    uint8 `json:"TargetSystem" `    // System ID
	TargetComponent uint8 `json:"TargetComponent" ` // Component ID
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

const (
	LOG_REQUEST_END_FIELD_TARGET_SYSTEM    = "LogRequestEnd.TargetSystem"
	LOG_REQUEST_END_FIELD_TARGET_COMPONENT = "LogRequestEnd.TargetComponent"
)

// ToMap (generated function)
func (m *LogRequestEnd) Dict() map[string]interface{} {
	return map[string]interface{}{
		LOG_REQUEST_END_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		LOG_REQUEST_END_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	TargetSystem    uint8   `json:"TargetSystem" `    // System ID
	TargetComponent uint8   `json:"TargetComponent" ` // Component ID
	Len             uint8   `json:"Len" `             // [ bytes ] Data length
	Data            []uint8 `json:"Data" len:"110" `  // Raw data (110 is enough for 12 satellites of RTCMv2)
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

const (
	GPS_INJECT_DATA_FIELD_TARGET_SYSTEM    = "GpsInjectData.TargetSystem"
	GPS_INJECT_DATA_FIELD_TARGET_COMPONENT = "GpsInjectData.TargetComponent"
	GPS_INJECT_DATA_FIELD_LEN              = "GpsInjectData.Len"
	GPS_INJECT_DATA_FIELD_DATA             = "GpsInjectData.Data"
)

// ToMap (generated function)
func (m *GpsInjectData) Dict() map[string]interface{} {
	return map[string]interface{}{
		GPS_INJECT_DATA_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		GPS_INJECT_DATA_FIELD_TARGET_COMPONENT: m.TargetComponent,
		GPS_INJECT_DATA_FIELD_LEN:              m.Len,
		GPS_INJECT_DATA_FIELD_DATA:             m.Data,
	}
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
	TimeUsec          uint64       `json:"TimeUsec" `          // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Lat               int32        `json:"Lat" `               // [ degE7 ] Latitude (WGS84)
	Lon               int32        `json:"Lon" `               // [ degE7 ] Longitude (WGS84)
	Alt               int32        `json:"Alt" `               // [ mm ] Altitude (MSL). Positive for up.
	DgpsAge           uint32       `json:"DgpsAge" `           // [ ms ] Age of DGPS info
	Eph               uint16       `json:"Eph" `               // GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX
	Epv               uint16       `json:"Epv" `               // GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX
	Vel               uint16       `json:"Vel" `               // [ cm/s ] GPS ground speed. If unknown, set to: UINT16_MAX
	Cog               uint16       `json:"Cog" `               // [ cdeg ] Course over ground (NOT heading, but direction of movement): 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	FixType           GPS_FIX_TYPE `json:"FixType" `           // GPS fix type.
	SatellitesVisible uint8        `json:"SatellitesVisible" ` // Number of satellites visible. If unknown, set to 255
	DgpsNumch         uint8        `json:"DgpsNumch" `         // Number of DGPS satellites
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

const (
	GPS2_RAW_FIELD_TIME_USEC          = "Gps2Raw.TimeUsec"
	GPS2_RAW_FIELD_LAT                = "Gps2Raw.Lat"
	GPS2_RAW_FIELD_LON                = "Gps2Raw.Lon"
	GPS2_RAW_FIELD_ALT                = "Gps2Raw.Alt"
	GPS2_RAW_FIELD_DGPS_AGE           = "Gps2Raw.DgpsAge"
	GPS2_RAW_FIELD_EPH                = "Gps2Raw.Eph"
	GPS2_RAW_FIELD_EPV                = "Gps2Raw.Epv"
	GPS2_RAW_FIELD_VEL                = "Gps2Raw.Vel"
	GPS2_RAW_FIELD_COG                = "Gps2Raw.Cog"
	GPS2_RAW_FIELD_FIX_TYPE           = "Gps2Raw.FixType"
	GPS2_RAW_FIELD_SATELLITES_VISIBLE = "Gps2Raw.SatellitesVisible"
	GPS2_RAW_FIELD_DGPS_NUMCH         = "Gps2Raw.DgpsNumch"
)

// ToMap (generated function)
func (m *Gps2Raw) Dict() map[string]interface{} {
	return map[string]interface{}{
		GPS2_RAW_FIELD_TIME_USEC:          m.TimeUsec,
		GPS2_RAW_FIELD_LAT:                m.Lat,
		GPS2_RAW_FIELD_LON:                m.Lon,
		GPS2_RAW_FIELD_ALT:                m.Alt,
		GPS2_RAW_FIELD_DGPS_AGE:           m.DgpsAge,
		GPS2_RAW_FIELD_EPH:                m.Eph,
		GPS2_RAW_FIELD_EPV:                m.Epv,
		GPS2_RAW_FIELD_VEL:                m.Vel,
		GPS2_RAW_FIELD_COG:                m.Cog,
		GPS2_RAW_FIELD_FIX_TYPE:           m.FixType,
		GPS2_RAW_FIELD_SATELLITES_VISIBLE: m.SatellitesVisible,
		GPS2_RAW_FIELD_DGPS_NUMCH:         m.DgpsNumch,
	}
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
	Vcc    uint16           `json:"Vcc" `    // [ mV ] 5V rail voltage.
	Vservo uint16           `json:"Vservo" ` // [ mV ] Servo rail voltage.
	Flags  MAV_POWER_STATUS `json:"Flags" `  // Bitmap of power supply status flags.
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

const (
	POWER_STATUS_FIELD_VCC    = "PowerStatus.Vcc"
	POWER_STATUS_FIELD_VSERVO = "PowerStatus.Vservo"
	POWER_STATUS_FIELD_FLAGS  = "PowerStatus.Flags"
)

// ToMap (generated function)
func (m *PowerStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		POWER_STATUS_FIELD_VCC:    m.Vcc,
		POWER_STATUS_FIELD_VSERVO: m.Vservo,
		POWER_STATUS_FIELD_FLAGS:  m.Flags,
	}
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
	Baudrate uint32              `json:"Baudrate" `      // [ bits/s ] Baudrate of transfer. Zero means no change.
	Timeout  uint16              `json:"Timeout" `       // [ ms ] Timeout for reply data
	Device   SERIAL_CONTROL_DEV  `json:"Device" `        // Serial control device type.
	Flags    SERIAL_CONTROL_FLAG `json:"Flags" `         // Bitmap of serial control flags.
	Count    uint8               `json:"Count" `         // [ bytes ] how many bytes in this transfer
	Data     []uint8             `json:"Data" len:"70" ` // serial data
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

const (
	SERIAL_CONTROL_FIELD_BAUDRATE = "SerialControl.Baudrate"
	SERIAL_CONTROL_FIELD_TIMEOUT  = "SerialControl.Timeout"
	SERIAL_CONTROL_FIELD_DEVICE   = "SerialControl.Device"
	SERIAL_CONTROL_FIELD_FLAGS    = "SerialControl.Flags"
	SERIAL_CONTROL_FIELD_COUNT    = "SerialControl.Count"
	SERIAL_CONTROL_FIELD_DATA     = "SerialControl.Data"
)

// ToMap (generated function)
func (m *SerialControl) Dict() map[string]interface{} {
	return map[string]interface{}{
		SERIAL_CONTROL_FIELD_BAUDRATE: m.Baudrate,
		SERIAL_CONTROL_FIELD_TIMEOUT:  m.Timeout,
		SERIAL_CONTROL_FIELD_DEVICE:   m.Device,
		SERIAL_CONTROL_FIELD_FLAGS:    m.Flags,
		SERIAL_CONTROL_FIELD_COUNT:    m.Count,
		SERIAL_CONTROL_FIELD_DATA:     m.Data,
	}
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
	TimeLastBaselineMs uint32                         `json:"TimeLastBaselineMs" ` // [ ms ] Time since boot of last baseline message received.
	Tow                uint32                         `json:"Tow" `                // [ ms ] GPS Time of Week of last baseline
	BaselineAMm        int32                          `json:"BaselineAMm" `        // [ mm ] Current baseline in ECEF x or NED north component.
	BaselineBMm        int32                          `json:"BaselineBMm" `        // [ mm ] Current baseline in ECEF y or NED east component.
	BaselineCMm        int32                          `json:"BaselineCMm" `        // [ mm ] Current baseline in ECEF z or NED down component.
	Accuracy           uint32                         `json:"Accuracy" `           // Current estimate of baseline accuracy.
	IarNumHypotheses   int32                          `json:"IarNumHypotheses" `   // Current number of integer ambiguity hypotheses.
	Wn                 uint16                         `json:"Wn" `                 // GPS Week Number of last baseline
	RtkReceiverID      uint8                          `json:"RtkReceiverID" `      // Identification of connected RTK receiver.
	RtkHealth          uint8                          `json:"RtkHealth" `          // GPS-specific health report for RTK data.
	RtkRate            uint8                          `json:"RtkRate" `            // [ Hz ] Rate of baseline messages being received by GPS
	Nsats              uint8                          `json:"Nsats" `              // Current number of sats used for RTK calculation.
	BaselineCoordsType RTK_BASELINE_COORDINATE_SYSTEM `json:"BaselineCoordsType" ` // Coordinate system of baseline
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

const (
	GPS_RTK_FIELD_TIME_LAST_BASELINE_MS = "GpsRtk.TimeLastBaselineMs"
	GPS_RTK_FIELD_TOW                   = "GpsRtk.Tow"
	GPS_RTK_FIELD_BASELINE_A_MM         = "GpsRtk.BaselineAMm"
	GPS_RTK_FIELD_BASELINE_B_MM         = "GpsRtk.BaselineBMm"
	GPS_RTK_FIELD_BASELINE_C_MM         = "GpsRtk.BaselineCMm"
	GPS_RTK_FIELD_ACCURACY              = "GpsRtk.Accuracy"
	GPS_RTK_FIELD_IAR_NUM_HYPOTHESES    = "GpsRtk.IarNumHypotheses"
	GPS_RTK_FIELD_WN                    = "GpsRtk.Wn"
	GPS_RTK_FIELD_RTK_RECEIVER_ID       = "GpsRtk.RtkReceiverID"
	GPS_RTK_FIELD_RTK_HEALTH            = "GpsRtk.RtkHealth"
	GPS_RTK_FIELD_RTK_RATE              = "GpsRtk.RtkRate"
	GPS_RTK_FIELD_NSATS                 = "GpsRtk.Nsats"
	GPS_RTK_FIELD_BASELINE_COORDS_TYPE  = "GpsRtk.BaselineCoordsType"
)

// ToMap (generated function)
func (m *GpsRtk) Dict() map[string]interface{} {
	return map[string]interface{}{
		GPS_RTK_FIELD_TIME_LAST_BASELINE_MS: m.TimeLastBaselineMs,
		GPS_RTK_FIELD_TOW:                   m.Tow,
		GPS_RTK_FIELD_BASELINE_A_MM:         m.BaselineAMm,
		GPS_RTK_FIELD_BASELINE_B_MM:         m.BaselineBMm,
		GPS_RTK_FIELD_BASELINE_C_MM:         m.BaselineCMm,
		GPS_RTK_FIELD_ACCURACY:              m.Accuracy,
		GPS_RTK_FIELD_IAR_NUM_HYPOTHESES:    m.IarNumHypotheses,
		GPS_RTK_FIELD_WN:                    m.Wn,
		GPS_RTK_FIELD_RTK_RECEIVER_ID:       m.RtkReceiverID,
		GPS_RTK_FIELD_RTK_HEALTH:            m.RtkHealth,
		GPS_RTK_FIELD_RTK_RATE:              m.RtkRate,
		GPS_RTK_FIELD_NSATS:                 m.Nsats,
		GPS_RTK_FIELD_BASELINE_COORDS_TYPE:  m.BaselineCoordsType,
	}
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
	TimeLastBaselineMs uint32                         `json:"TimeLastBaselineMs" ` // [ ms ] Time since boot of last baseline message received.
	Tow                uint32                         `json:"Tow" `                // [ ms ] GPS Time of Week of last baseline
	BaselineAMm        int32                          `json:"BaselineAMm" `        // [ mm ] Current baseline in ECEF x or NED north component.
	BaselineBMm        int32                          `json:"BaselineBMm" `        // [ mm ] Current baseline in ECEF y or NED east component.
	BaselineCMm        int32                          `json:"BaselineCMm" `        // [ mm ] Current baseline in ECEF z or NED down component.
	Accuracy           uint32                         `json:"Accuracy" `           // Current estimate of baseline accuracy.
	IarNumHypotheses   int32                          `json:"IarNumHypotheses" `   // Current number of integer ambiguity hypotheses.
	Wn                 uint16                         `json:"Wn" `                 // GPS Week Number of last baseline
	RtkReceiverID      uint8                          `json:"RtkReceiverID" `      // Identification of connected RTK receiver.
	RtkHealth          uint8                          `json:"RtkHealth" `          // GPS-specific health report for RTK data.
	RtkRate            uint8                          `json:"RtkRate" `            // [ Hz ] Rate of baseline messages being received by GPS
	Nsats              uint8                          `json:"Nsats" `              // Current number of sats used for RTK calculation.
	BaselineCoordsType RTK_BASELINE_COORDINATE_SYSTEM `json:"BaselineCoordsType" ` // Coordinate system of baseline
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

const (
	GPS2_RTK_FIELD_TIME_LAST_BASELINE_MS = "Gps2Rtk.TimeLastBaselineMs"
	GPS2_RTK_FIELD_TOW                   = "Gps2Rtk.Tow"
	GPS2_RTK_FIELD_BASELINE_A_MM         = "Gps2Rtk.BaselineAMm"
	GPS2_RTK_FIELD_BASELINE_B_MM         = "Gps2Rtk.BaselineBMm"
	GPS2_RTK_FIELD_BASELINE_C_MM         = "Gps2Rtk.BaselineCMm"
	GPS2_RTK_FIELD_ACCURACY              = "Gps2Rtk.Accuracy"
	GPS2_RTK_FIELD_IAR_NUM_HYPOTHESES    = "Gps2Rtk.IarNumHypotheses"
	GPS2_RTK_FIELD_WN                    = "Gps2Rtk.Wn"
	GPS2_RTK_FIELD_RTK_RECEIVER_ID       = "Gps2Rtk.RtkReceiverID"
	GPS2_RTK_FIELD_RTK_HEALTH            = "Gps2Rtk.RtkHealth"
	GPS2_RTK_FIELD_RTK_RATE              = "Gps2Rtk.RtkRate"
	GPS2_RTK_FIELD_NSATS                 = "Gps2Rtk.Nsats"
	GPS2_RTK_FIELD_BASELINE_COORDS_TYPE  = "Gps2Rtk.BaselineCoordsType"
)

// ToMap (generated function)
func (m *Gps2Rtk) Dict() map[string]interface{} {
	return map[string]interface{}{
		GPS2_RTK_FIELD_TIME_LAST_BASELINE_MS: m.TimeLastBaselineMs,
		GPS2_RTK_FIELD_TOW:                   m.Tow,
		GPS2_RTK_FIELD_BASELINE_A_MM:         m.BaselineAMm,
		GPS2_RTK_FIELD_BASELINE_B_MM:         m.BaselineBMm,
		GPS2_RTK_FIELD_BASELINE_C_MM:         m.BaselineCMm,
		GPS2_RTK_FIELD_ACCURACY:              m.Accuracy,
		GPS2_RTK_FIELD_IAR_NUM_HYPOTHESES:    m.IarNumHypotheses,
		GPS2_RTK_FIELD_WN:                    m.Wn,
		GPS2_RTK_FIELD_RTK_RECEIVER_ID:       m.RtkReceiverID,
		GPS2_RTK_FIELD_RTK_HEALTH:            m.RtkHealth,
		GPS2_RTK_FIELD_RTK_RATE:              m.RtkRate,
		GPS2_RTK_FIELD_NSATS:                 m.Nsats,
		GPS2_RTK_FIELD_BASELINE_COORDS_TYPE:  m.BaselineCoordsType,
	}
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
	TimeBootMs uint32 `json:"TimeBootMs" ` // [ ms ] Timestamp (time since system boot).
	Xacc       int16  `json:"Xacc" `       // [ mG ] X acceleration
	Yacc       int16  `json:"Yacc" `       // [ mG ] Y acceleration
	Zacc       int16  `json:"Zacc" `       // [ mG ] Z acceleration
	Xgyro      int16  `json:"Xgyro" `      // [ mrad/s ] Angular speed around X axis
	Ygyro      int16  `json:"Ygyro" `      // [ mrad/s ] Angular speed around Y axis
	Zgyro      int16  `json:"Zgyro" `      // [ mrad/s ] Angular speed around Z axis
	Xmag       int16  `json:"Xmag" `       // [ mgauss ] X Magnetic field
	Ymag       int16  `json:"Ymag" `       // [ mgauss ] Y Magnetic field
	Zmag       int16  `json:"Zmag" `       // [ mgauss ] Z Magnetic field
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

const (
	SCALED_IMU3_FIELD_TIME_BOOT_MS = "ScaledImu3.TimeBootMs"
	SCALED_IMU3_FIELD_XACC         = "ScaledImu3.Xacc"
	SCALED_IMU3_FIELD_YACC         = "ScaledImu3.Yacc"
	SCALED_IMU3_FIELD_ZACC         = "ScaledImu3.Zacc"
	SCALED_IMU3_FIELD_XGYRO        = "ScaledImu3.Xgyro"
	SCALED_IMU3_FIELD_YGYRO        = "ScaledImu3.Ygyro"
	SCALED_IMU3_FIELD_ZGYRO        = "ScaledImu3.Zgyro"
	SCALED_IMU3_FIELD_XMAG         = "ScaledImu3.Xmag"
	SCALED_IMU3_FIELD_YMAG         = "ScaledImu3.Ymag"
	SCALED_IMU3_FIELD_ZMAG         = "ScaledImu3.Zmag"
)

// ToMap (generated function)
func (m *ScaledImu3) Dict() map[string]interface{} {
	return map[string]interface{}{
		SCALED_IMU3_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		SCALED_IMU3_FIELD_XACC:         m.Xacc,
		SCALED_IMU3_FIELD_YACC:         m.Yacc,
		SCALED_IMU3_FIELD_ZACC:         m.Zacc,
		SCALED_IMU3_FIELD_XGYRO:        m.Xgyro,
		SCALED_IMU3_FIELD_YGYRO:        m.Ygyro,
		SCALED_IMU3_FIELD_ZGYRO:        m.Zgyro,
		SCALED_IMU3_FIELD_XMAG:         m.Xmag,
		SCALED_IMU3_FIELD_YMAG:         m.Ymag,
		SCALED_IMU3_FIELD_ZMAG:         m.Zmag,
	}
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
	Size       uint32                   `json:"Size" `       // [ bytes ] total data size (set on ACK only).
	Width      uint16                   `json:"Width" `      // Width of a matrix or image.
	Height     uint16                   `json:"Height" `     // Height of a matrix or image.
	Packets    uint16                   `json:"Packets" `    // Number of packets being sent (set on ACK only).
	Type       MAVLINK_DATA_STREAM_TYPE `json:"Type" `       // Type of requested/acknowledged data.
	Payload    uint8                    `json:"Payload" `    // [ bytes ] Payload size per packet (normally 253 byte, see DATA field size in message ENCAPSULATED_DATA) (set on ACK only).
	JpgQuality uint8                    `json:"JpgQuality" ` // [ % ] JPEG quality. Values: [1-100].
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

const (
	DATA_TRANSMISSION_HANDSHAKE_FIELD_SIZE        = "DataTransmissionHandshake.Size"
	DATA_TRANSMISSION_HANDSHAKE_FIELD_WIDTH       = "DataTransmissionHandshake.Width"
	DATA_TRANSMISSION_HANDSHAKE_FIELD_HEIGHT      = "DataTransmissionHandshake.Height"
	DATA_TRANSMISSION_HANDSHAKE_FIELD_PACKETS     = "DataTransmissionHandshake.Packets"
	DATA_TRANSMISSION_HANDSHAKE_FIELD_TYPE        = "DataTransmissionHandshake.Type"
	DATA_TRANSMISSION_HANDSHAKE_FIELD_PAYLOAD     = "DataTransmissionHandshake.Payload"
	DATA_TRANSMISSION_HANDSHAKE_FIELD_JPG_QUALITY = "DataTransmissionHandshake.JpgQuality"
)

// ToMap (generated function)
func (m *DataTransmissionHandshake) Dict() map[string]interface{} {
	return map[string]interface{}{
		DATA_TRANSMISSION_HANDSHAKE_FIELD_SIZE:        m.Size,
		DATA_TRANSMISSION_HANDSHAKE_FIELD_WIDTH:       m.Width,
		DATA_TRANSMISSION_HANDSHAKE_FIELD_HEIGHT:      m.Height,
		DATA_TRANSMISSION_HANDSHAKE_FIELD_PACKETS:     m.Packets,
		DATA_TRANSMISSION_HANDSHAKE_FIELD_TYPE:        m.Type,
		DATA_TRANSMISSION_HANDSHAKE_FIELD_PAYLOAD:     m.Payload,
		DATA_TRANSMISSION_HANDSHAKE_FIELD_JPG_QUALITY: m.JpgQuality,
	}
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
	Seqnr uint16  `json:"Seqnr" `          // sequence number (starting with 0 on every transmission)
	Data  []uint8 `json:"Data" len:"253" ` // image data bytes
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

const (
	ENCAPSULATED_DATA_FIELD_SEQNR = "EncapsulatedData.Seqnr"
	ENCAPSULATED_DATA_FIELD_DATA  = "EncapsulatedData.Data"
)

// ToMap (generated function)
func (m *EncapsulatedData) Dict() map[string]interface{} {
	return map[string]interface{}{
		ENCAPSULATED_DATA_FIELD_SEQNR: m.Seqnr,
		ENCAPSULATED_DATA_FIELD_DATA:  m.Data,
	}
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
	TimeBootMs      uint32                 `json:"TimeBootMs" `      // [ ms ] Timestamp (time since system boot).
	MinDistance     uint16                 `json:"MinDistance" `     // [ cm ] Minimum distance the sensor can measure
	MaxDistance     uint16                 `json:"MaxDistance" `     // [ cm ] Maximum distance the sensor can measure
	CurrentDistance uint16                 `json:"CurrentDistance" ` // [ cm ] Current distance reading
	Type            MAV_DISTANCE_SENSOR    `json:"Type" `            // Type of distance sensor.
	ID              uint8                  `json:"ID" `              // Onboard ID of the sensor
	Orientation     MAV_SENSOR_ORIENTATION `json:"Orientation" `     // Direction the sensor faces. downward-facing: ROTATION_PITCH_270, upward-facing: ROTATION_PITCH_90, backward-facing: ROTATION_PITCH_180, forward-facing: ROTATION_NONE, left-facing: ROTATION_YAW_90, right-facing: ROTATION_YAW_270
	Covariance      uint8                  `json:"Covariance" `      // [ cm^2 ] Measurement variance. Max standard deviation is 6cm. 255 if unknown.
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

const (
	DISTANCE_SENSOR_FIELD_TIME_BOOT_MS     = "DistanceSensor.TimeBootMs"
	DISTANCE_SENSOR_FIELD_MIN_DISTANCE     = "DistanceSensor.MinDistance"
	DISTANCE_SENSOR_FIELD_MAX_DISTANCE     = "DistanceSensor.MaxDistance"
	DISTANCE_SENSOR_FIELD_CURRENT_DISTANCE = "DistanceSensor.CurrentDistance"
	DISTANCE_SENSOR_FIELD_TYPE             = "DistanceSensor.Type"
	DISTANCE_SENSOR_FIELD_ID               = "DistanceSensor.ID"
	DISTANCE_SENSOR_FIELD_ORIENTATION      = "DistanceSensor.Orientation"
	DISTANCE_SENSOR_FIELD_COVARIANCE       = "DistanceSensor.Covariance"
)

// ToMap (generated function)
func (m *DistanceSensor) Dict() map[string]interface{} {
	return map[string]interface{}{
		DISTANCE_SENSOR_FIELD_TIME_BOOT_MS:     m.TimeBootMs,
		DISTANCE_SENSOR_FIELD_MIN_DISTANCE:     m.MinDistance,
		DISTANCE_SENSOR_FIELD_MAX_DISTANCE:     m.MaxDistance,
		DISTANCE_SENSOR_FIELD_CURRENT_DISTANCE: m.CurrentDistance,
		DISTANCE_SENSOR_FIELD_TYPE:             m.Type,
		DISTANCE_SENSOR_FIELD_ID:               m.ID,
		DISTANCE_SENSOR_FIELD_ORIENTATION:      m.Orientation,
		DISTANCE_SENSOR_FIELD_COVARIANCE:       m.Covariance,
	}
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
	Mask        uint64 `json:"Mask" `        // Bitmask of requested 4x4 grids (row major 8x7 array of grids, 56 bits)
	Lat         int32  `json:"Lat" `         // [ degE7 ] Latitude of SW corner of first grid
	Lon         int32  `json:"Lon" `         // [ degE7 ] Longitude of SW corner of first grid
	GridSpacing uint16 `json:"GridSpacing" ` // [ m ] Grid spacing
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

const (
	TERRAIN_REQUEST_FIELD_MASK         = "TerrainRequest.Mask"
	TERRAIN_REQUEST_FIELD_LAT          = "TerrainRequest.Lat"
	TERRAIN_REQUEST_FIELD_LON          = "TerrainRequest.Lon"
	TERRAIN_REQUEST_FIELD_GRID_SPACING = "TerrainRequest.GridSpacing"
)

// ToMap (generated function)
func (m *TerrainRequest) Dict() map[string]interface{} {
	return map[string]interface{}{
		TERRAIN_REQUEST_FIELD_MASK:         m.Mask,
		TERRAIN_REQUEST_FIELD_LAT:          m.Lat,
		TERRAIN_REQUEST_FIELD_LON:          m.Lon,
		TERRAIN_REQUEST_FIELD_GRID_SPACING: m.GridSpacing,
	}
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
	Lat         int32   `json:"Lat" `           // [ degE7 ] Latitude of SW corner of first grid
	Lon         int32   `json:"Lon" `           // [ degE7 ] Longitude of SW corner of first grid
	GridSpacing uint16  `json:"GridSpacing" `   // [ m ] Grid spacing
	Data        []int16 `json:"Data" len:"16" ` // [ m ] Terrain data MSL
	Gridbit     uint8   `json:"Gridbit" `       // bit within the terrain request mask
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

const (
	TERRAIN_DATA_FIELD_LAT          = "TerrainData.Lat"
	TERRAIN_DATA_FIELD_LON          = "TerrainData.Lon"
	TERRAIN_DATA_FIELD_GRID_SPACING = "TerrainData.GridSpacing"
	TERRAIN_DATA_FIELD_DATA         = "TerrainData.Data"
	TERRAIN_DATA_FIELD_GRIDBIT      = "TerrainData.Gridbit"
)

// ToMap (generated function)
func (m *TerrainData) Dict() map[string]interface{} {
	return map[string]interface{}{
		TERRAIN_DATA_FIELD_LAT:          m.Lat,
		TERRAIN_DATA_FIELD_LON:          m.Lon,
		TERRAIN_DATA_FIELD_GRID_SPACING: m.GridSpacing,
		TERRAIN_DATA_FIELD_DATA:         m.Data,
		TERRAIN_DATA_FIELD_GRIDBIT:      m.Gridbit,
	}
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
	Lat int32 `json:"Lat" ` // [ degE7 ] Latitude
	Lon int32 `json:"Lon" ` // [ degE7 ] Longitude
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

const (
	TERRAIN_CHECK_FIELD_LAT = "TerrainCheck.Lat"
	TERRAIN_CHECK_FIELD_LON = "TerrainCheck.Lon"
)

// ToMap (generated function)
func (m *TerrainCheck) Dict() map[string]interface{} {
	return map[string]interface{}{
		TERRAIN_CHECK_FIELD_LAT: m.Lat,
		TERRAIN_CHECK_FIELD_LON: m.Lon,
	}
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
	Lat           int32   `json:"Lat" `           // [ degE7 ] Latitude
	Lon           int32   `json:"Lon" `           // [ degE7 ] Longitude
	TerrainHeight float32 `json:"TerrainHeight" ` // [ m ] Terrain height MSL
	CurrentHeight float32 `json:"CurrentHeight" ` // [ m ] Current vehicle height above lat/lon terrain height
	Spacing       uint16  `json:"Spacing" `       // grid spacing (zero if terrain at this location unavailable)
	Pending       uint16  `json:"Pending" `       // Number of 4x4 terrain blocks waiting to be received or read from disk
	Loaded        uint16  `json:"Loaded" `        // Number of 4x4 terrain blocks in memory
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

const (
	TERRAIN_REPORT_FIELD_LAT            = "TerrainReport.Lat"
	TERRAIN_REPORT_FIELD_LON            = "TerrainReport.Lon"
	TERRAIN_REPORT_FIELD_TERRAIN_HEIGHT = "TerrainReport.TerrainHeight"
	TERRAIN_REPORT_FIELD_CURRENT_HEIGHT = "TerrainReport.CurrentHeight"
	TERRAIN_REPORT_FIELD_SPACING        = "TerrainReport.Spacing"
	TERRAIN_REPORT_FIELD_PENDING        = "TerrainReport.Pending"
	TERRAIN_REPORT_FIELD_LOADED         = "TerrainReport.Loaded"
)

// ToMap (generated function)
func (m *TerrainReport) Dict() map[string]interface{} {
	return map[string]interface{}{
		TERRAIN_REPORT_FIELD_LAT:            m.Lat,
		TERRAIN_REPORT_FIELD_LON:            m.Lon,
		TERRAIN_REPORT_FIELD_TERRAIN_HEIGHT: m.TerrainHeight,
		TERRAIN_REPORT_FIELD_CURRENT_HEIGHT: m.CurrentHeight,
		TERRAIN_REPORT_FIELD_SPACING:        m.Spacing,
		TERRAIN_REPORT_FIELD_PENDING:        m.Pending,
		TERRAIN_REPORT_FIELD_LOADED:         m.Loaded,
	}
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
	TimeBootMs  uint32  `json:"TimeBootMs" `  // [ ms ] Timestamp (time since system boot).
	PressAbs    float32 `json:"PressAbs" `    // [ hPa ] Absolute pressure
	PressDiff   float32 `json:"PressDiff" `   // [ hPa ] Differential pressure
	Temperature int16   `json:"Temperature" ` // [ cdegC ] Absolute pressure temperature
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

const (
	SCALED_PRESSURE2_FIELD_TIME_BOOT_MS = "ScaledPressure2.TimeBootMs"
	SCALED_PRESSURE2_FIELD_PRESS_ABS    = "ScaledPressure2.PressAbs"
	SCALED_PRESSURE2_FIELD_PRESS_DIFF   = "ScaledPressure2.PressDiff"
	SCALED_PRESSURE2_FIELD_TEMPERATURE  = "ScaledPressure2.Temperature"
)

// ToMap (generated function)
func (m *ScaledPressure2) Dict() map[string]interface{} {
	return map[string]interface{}{
		SCALED_PRESSURE2_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		SCALED_PRESSURE2_FIELD_PRESS_ABS:    m.PressAbs,
		SCALED_PRESSURE2_FIELD_PRESS_DIFF:   m.PressDiff,
		SCALED_PRESSURE2_FIELD_TEMPERATURE:  m.Temperature,
	}
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
	TimeUsec uint64    `json:"TimeUsec" `  // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Q        []float32 `json:"Q" len:"4" ` // Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	X        float32   `json:"X" `         // [ m ] X position (NED)
	Y        float32   `json:"Y" `         // [ m ] Y position (NED)
	Z        float32   `json:"Z" `         // [ m ] Z position (NED)
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

const (
	ATT_POS_MOCAP_FIELD_TIME_USEC = "AttPosMocap.TimeUsec"
	ATT_POS_MOCAP_FIELD_Q         = "AttPosMocap.Q"
	ATT_POS_MOCAP_FIELD_X         = "AttPosMocap.X"
	ATT_POS_MOCAP_FIELD_Y         = "AttPosMocap.Y"
	ATT_POS_MOCAP_FIELD_Z         = "AttPosMocap.Z"
)

// ToMap (generated function)
func (m *AttPosMocap) Dict() map[string]interface{} {
	return map[string]interface{}{
		ATT_POS_MOCAP_FIELD_TIME_USEC: m.TimeUsec,
		ATT_POS_MOCAP_FIELD_Q:         m.Q,
		ATT_POS_MOCAP_FIELD_X:         m.X,
		ATT_POS_MOCAP_FIELD_Y:         m.Y,
		ATT_POS_MOCAP_FIELD_Z:         m.Z,
	}
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
	TimeUsec        uint64    `json:"TimeUsec" `         // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Controls        []float32 `json:"Controls" len:"8" ` // Actuator controls. Normed to -1..+1 where 0 is neutral position. Throttle for single rotation direction motors is 0..1, negative range for reverse direction. Standard mapping for attitude controls (group 0): (index 0-7): roll, pitch, yaw, throttle, flaps, spoilers, airbrakes, landing gear. Load a pass-through mixer to repurpose them as generic outputs.
	GroupMlx        uint8     `json:"GroupMlx" `         // Actuator group. The "_mlx" indicates this is a multi-instance message and a MAVLink parser should use this field to difference between instances.
	TargetSystem    uint8     `json:"TargetSystem" `     // System ID
	TargetComponent uint8     `json:"TargetComponent" `  // Component ID
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

const (
	SET_ACTUATOR_CONTROL_TARGET_FIELD_TIME_USEC        = "SetActuatorControlTarget.TimeUsec"
	SET_ACTUATOR_CONTROL_TARGET_FIELD_CONTROLS         = "SetActuatorControlTarget.Controls"
	SET_ACTUATOR_CONTROL_TARGET_FIELD_GROUP_MLX        = "SetActuatorControlTarget.GroupMlx"
	SET_ACTUATOR_CONTROL_TARGET_FIELD_TARGET_SYSTEM    = "SetActuatorControlTarget.TargetSystem"
	SET_ACTUATOR_CONTROL_TARGET_FIELD_TARGET_COMPONENT = "SetActuatorControlTarget.TargetComponent"
)

// ToMap (generated function)
func (m *SetActuatorControlTarget) Dict() map[string]interface{} {
	return map[string]interface{}{
		SET_ACTUATOR_CONTROL_TARGET_FIELD_TIME_USEC:        m.TimeUsec,
		SET_ACTUATOR_CONTROL_TARGET_FIELD_CONTROLS:         m.Controls,
		SET_ACTUATOR_CONTROL_TARGET_FIELD_GROUP_MLX:        m.GroupMlx,
		SET_ACTUATOR_CONTROL_TARGET_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		SET_ACTUATOR_CONTROL_TARGET_FIELD_TARGET_COMPONENT: m.TargetComponent,
	}
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
	TimeUsec uint64    `json:"TimeUsec" `         // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Controls []float32 `json:"Controls" len:"8" ` // Actuator controls. Normed to -1..+1 where 0 is neutral position. Throttle for single rotation direction motors is 0..1, negative range for reverse direction. Standard mapping for attitude controls (group 0): (index 0-7): roll, pitch, yaw, throttle, flaps, spoilers, airbrakes, landing gear. Load a pass-through mixer to repurpose them as generic outputs.
	GroupMlx uint8     `json:"GroupMlx" `         // Actuator group. The "_mlx" indicates this is a multi-instance message and a MAVLink parser should use this field to difference between instances.
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

const (
	ACTUATOR_CONTROL_TARGET_FIELD_TIME_USEC = "ActuatorControlTarget.TimeUsec"
	ACTUATOR_CONTROL_TARGET_FIELD_CONTROLS  = "ActuatorControlTarget.Controls"
	ACTUATOR_CONTROL_TARGET_FIELD_GROUP_MLX = "ActuatorControlTarget.GroupMlx"
)

// ToMap (generated function)
func (m *ActuatorControlTarget) Dict() map[string]interface{} {
	return map[string]interface{}{
		ACTUATOR_CONTROL_TARGET_FIELD_TIME_USEC: m.TimeUsec,
		ACTUATOR_CONTROL_TARGET_FIELD_CONTROLS:  m.Controls,
		ACTUATOR_CONTROL_TARGET_FIELD_GROUP_MLX: m.GroupMlx,
	}
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
	TimeUsec          uint64  `json:"TimeUsec" `          // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	AltitudeMonotonic float32 `json:"AltitudeMonotonic" ` // [ m ] This altitude measure is initialized on system boot and monotonic (it is never reset, but represents the local altitude change). The only guarantee on this field is that it will never be reset and is consistent within a flight. The recommended value for this field is the uncorrected barometric altitude at boot time. This altitude will also drift and vary between flights.
	AltitudeAmsl      float32 `json:"AltitudeAmsl" `      // [ m ] This altitude measure is strictly above mean sea level and might be non-monotonic (it might reset on events like GPS lock or when a new QNH value is set). It should be the altitude to which global altitude waypoints are compared to. Note that it is *not* the GPS altitude, however, most GPS modules already output MSL by default and not the WGS84 altitude.
	AltitudeLocal     float32 `json:"AltitudeLocal" `     // [ m ] This is the local altitude in the local coordinate frame. It is not the altitude above home, but in reference to the coordinate origin (0, 0, 0). It is up-positive.
	AltitudeRelative  float32 `json:"AltitudeRelative" `  // [ m ] This is the altitude above the home position. It resets on each change of the current home position.
	AltitudeTerrain   float32 `json:"AltitudeTerrain" `   // [ m ] This is the altitude above terrain. It might be fed by a terrain database or an altimeter. Values smaller than -1000 should be interpreted as unknown.
	BottomClearance   float32 `json:"BottomClearance" `   // [ m ] This is not the altitude, but the clear space below the system according to the fused clearance estimate. It generally should max out at the maximum range of e.g. the laser altimeter. It is generally a moving target. A negative value indicates no measurement available.
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

const (
	ALTITUDE_FIELD_TIME_USEC          = "Altitude.TimeUsec"
	ALTITUDE_FIELD_ALTITUDE_MONOTONIC = "Altitude.AltitudeMonotonic"
	ALTITUDE_FIELD_ALTITUDE_AMSL      = "Altitude.AltitudeAmsl"
	ALTITUDE_FIELD_ALTITUDE_LOCAL     = "Altitude.AltitudeLocal"
	ALTITUDE_FIELD_ALTITUDE_RELATIVE  = "Altitude.AltitudeRelative"
	ALTITUDE_FIELD_ALTITUDE_TERRAIN   = "Altitude.AltitudeTerrain"
	ALTITUDE_FIELD_BOTTOM_CLEARANCE   = "Altitude.BottomClearance"
)

// ToMap (generated function)
func (m *Altitude) Dict() map[string]interface{} {
	return map[string]interface{}{
		ALTITUDE_FIELD_TIME_USEC:          m.TimeUsec,
		ALTITUDE_FIELD_ALTITUDE_MONOTONIC: m.AltitudeMonotonic,
		ALTITUDE_FIELD_ALTITUDE_AMSL:      m.AltitudeAmsl,
		ALTITUDE_FIELD_ALTITUDE_LOCAL:     m.AltitudeLocal,
		ALTITUDE_FIELD_ALTITUDE_RELATIVE:  m.AltitudeRelative,
		ALTITUDE_FIELD_ALTITUDE_TERRAIN:   m.AltitudeTerrain,
		ALTITUDE_FIELD_BOTTOM_CLEARANCE:   m.BottomClearance,
	}
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
	RequestID    uint8   `json:"RequestID" `         // Request ID. This ID should be re-used when sending back URI contents
	URIType      uint8   `json:"URIType" `           // The type of requested URI. 0 = a file via URL. 1 = a UAVCAN binary
	URI          []uint8 `json:"URI" len:"120" `     // The requested unique resource identifier (URI). It is not necessarily a straight domain name (depends on the URI type enum)
	TransferType uint8   `json:"TransferType" `      // The way the autopilot wants to receive the URI. 0 = MAVLink FTP. 1 = binary stream.
	Storage      []uint8 `json:"Storage" len:"120" ` // The storage path the autopilot wants the URI to be stored in. Will only be valid if the transfer_type has a storage associated (e.g. MAVLink FTP).
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

const (
	RESOURCE_REQUEST_FIELD_REQUEST_ID    = "ResourceRequest.RequestID"
	RESOURCE_REQUEST_FIELD_URI_TYPE      = "ResourceRequest.URIType"
	RESOURCE_REQUEST_FIELD_URI           = "ResourceRequest.URI"
	RESOURCE_REQUEST_FIELD_TRANSFER_TYPE = "ResourceRequest.TransferType"
	RESOURCE_REQUEST_FIELD_STORAGE       = "ResourceRequest.Storage"
)

// ToMap (generated function)
func (m *ResourceRequest) Dict() map[string]interface{} {
	return map[string]interface{}{
		RESOURCE_REQUEST_FIELD_REQUEST_ID:    m.RequestID,
		RESOURCE_REQUEST_FIELD_URI_TYPE:      m.URIType,
		RESOURCE_REQUEST_FIELD_URI:           m.URI,
		RESOURCE_REQUEST_FIELD_TRANSFER_TYPE: m.TransferType,
		RESOURCE_REQUEST_FIELD_STORAGE:       m.Storage,
	}
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
	TimeBootMs  uint32  `json:"TimeBootMs" `  // [ ms ] Timestamp (time since system boot).
	PressAbs    float32 `json:"PressAbs" `    // [ hPa ] Absolute pressure
	PressDiff   float32 `json:"PressDiff" `   // [ hPa ] Differential pressure
	Temperature int16   `json:"Temperature" ` // [ cdegC ] Absolute pressure temperature
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

const (
	SCALED_PRESSURE3_FIELD_TIME_BOOT_MS = "ScaledPressure3.TimeBootMs"
	SCALED_PRESSURE3_FIELD_PRESS_ABS    = "ScaledPressure3.PressAbs"
	SCALED_PRESSURE3_FIELD_PRESS_DIFF   = "ScaledPressure3.PressDiff"
	SCALED_PRESSURE3_FIELD_TEMPERATURE  = "ScaledPressure3.Temperature"
)

// ToMap (generated function)
func (m *ScaledPressure3) Dict() map[string]interface{} {
	return map[string]interface{}{
		SCALED_PRESSURE3_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		SCALED_PRESSURE3_FIELD_PRESS_ABS:    m.PressAbs,
		SCALED_PRESSURE3_FIELD_PRESS_DIFF:   m.PressDiff,
		SCALED_PRESSURE3_FIELD_TEMPERATURE:  m.Temperature,
	}
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
	Timestamp       uint64    `json:"Timestamp" `           // [ ms ] Timestamp (time since system boot).
	CustomState     uint64    `json:"CustomState" `         // button states or switches of a tracker device
	Lat             int32     `json:"Lat" `                 // [ degE7 ] Latitude (WGS84)
	Lon             int32     `json:"Lon" `                 // [ degE7 ] Longitude (WGS84)
	Alt             float32   `json:"Alt" `                 // [ m ] Altitude (MSL)
	Vel             []float32 `json:"Vel" len:"3" `         // [ m/s ] target velocity (0,0,0) for unknown
	Acc             []float32 `json:"Acc" len:"3" `         // [ m/s/s ] linear target acceleration (0,0,0) for unknown
	AttitudeQ       []float32 `json:"AttitudeQ" len:"4" `   // (1 0 0 0 for unknown)
	Rates           []float32 `json:"Rates" len:"3" `       // (0 0 0 for unknown)
	PositionCov     []float32 `json:"PositionCov" len:"3" ` // eph epv
	EstCapabilities uint8     `json:"EstCapabilities" `     // bit positions for tracker reporting capabilities (POS = 0, VEL = 1, ACCEL = 2, ATT + RATES = 3)
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

const (
	FOLLOW_TARGET_FIELD_TIMESTAMP        = "FollowTarget.Timestamp"
	FOLLOW_TARGET_FIELD_CUSTOM_STATE     = "FollowTarget.CustomState"
	FOLLOW_TARGET_FIELD_LAT              = "FollowTarget.Lat"
	FOLLOW_TARGET_FIELD_LON              = "FollowTarget.Lon"
	FOLLOW_TARGET_FIELD_ALT              = "FollowTarget.Alt"
	FOLLOW_TARGET_FIELD_VEL              = "FollowTarget.Vel"
	FOLLOW_TARGET_FIELD_ACC              = "FollowTarget.Acc"
	FOLLOW_TARGET_FIELD_ATTITUDE_Q       = "FollowTarget.AttitudeQ"
	FOLLOW_TARGET_FIELD_RATES            = "FollowTarget.Rates"
	FOLLOW_TARGET_FIELD_POSITION_COV     = "FollowTarget.PositionCov"
	FOLLOW_TARGET_FIELD_EST_CAPABILITIES = "FollowTarget.EstCapabilities"
)

// ToMap (generated function)
func (m *FollowTarget) Dict() map[string]interface{} {
	return map[string]interface{}{
		FOLLOW_TARGET_FIELD_TIMESTAMP:        m.Timestamp,
		FOLLOW_TARGET_FIELD_CUSTOM_STATE:     m.CustomState,
		FOLLOW_TARGET_FIELD_LAT:              m.Lat,
		FOLLOW_TARGET_FIELD_LON:              m.Lon,
		FOLLOW_TARGET_FIELD_ALT:              m.Alt,
		FOLLOW_TARGET_FIELD_VEL:              m.Vel,
		FOLLOW_TARGET_FIELD_ACC:              m.Acc,
		FOLLOW_TARGET_FIELD_ATTITUDE_Q:       m.AttitudeQ,
		FOLLOW_TARGET_FIELD_RATES:            m.Rates,
		FOLLOW_TARGET_FIELD_POSITION_COV:     m.PositionCov,
		FOLLOW_TARGET_FIELD_EST_CAPABILITIES: m.EstCapabilities,
	}
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
	TimeUsec    uint64    `json:"TimeUsec" `            // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	XAcc        float32   `json:"XAcc" `                // [ m/s/s ] X acceleration in body frame
	YAcc        float32   `json:"YAcc" `                // [ m/s/s ] Y acceleration in body frame
	ZAcc        float32   `json:"ZAcc" `                // [ m/s/s ] Z acceleration in body frame
	XVel        float32   `json:"XVel" `                // [ m/s ] X velocity in body frame
	YVel        float32   `json:"YVel" `                // [ m/s ] Y velocity in body frame
	ZVel        float32   `json:"ZVel" `                // [ m/s ] Z velocity in body frame
	XPos        float32   `json:"XPos" `                // [ m ] X position in local frame
	YPos        float32   `json:"YPos" `                // [ m ] Y position in local frame
	ZPos        float32   `json:"ZPos" `                // [ m ] Z position in local frame
	Airspeed    float32   `json:"Airspeed" `            // [ m/s ] Airspeed, set to -1 if unknown
	VelVariance []float32 `json:"VelVariance" len:"3" ` // Variance of body velocity estimate
	PosVariance []float32 `json:"PosVariance" len:"3" ` // Variance in local position
	Q           []float32 `json:"Q" len:"4" `           // The attitude, represented as Quaternion
	RollRate    float32   `json:"RollRate" `            // [ rad/s ] Angular rate in roll axis
	PitchRate   float32   `json:"PitchRate" `           // [ rad/s ] Angular rate in pitch axis
	YawRate     float32   `json:"YawRate" `             // [ rad/s ] Angular rate in yaw axis
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

const (
	CONTROL_SYSTEM_STATE_FIELD_TIME_USEC    = "ControlSystemState.TimeUsec"
	CONTROL_SYSTEM_STATE_FIELD_X_ACC        = "ControlSystemState.XAcc"
	CONTROL_SYSTEM_STATE_FIELD_Y_ACC        = "ControlSystemState.YAcc"
	CONTROL_SYSTEM_STATE_FIELD_Z_ACC        = "ControlSystemState.ZAcc"
	CONTROL_SYSTEM_STATE_FIELD_X_VEL        = "ControlSystemState.XVel"
	CONTROL_SYSTEM_STATE_FIELD_Y_VEL        = "ControlSystemState.YVel"
	CONTROL_SYSTEM_STATE_FIELD_Z_VEL        = "ControlSystemState.ZVel"
	CONTROL_SYSTEM_STATE_FIELD_X_POS        = "ControlSystemState.XPos"
	CONTROL_SYSTEM_STATE_FIELD_Y_POS        = "ControlSystemState.YPos"
	CONTROL_SYSTEM_STATE_FIELD_Z_POS        = "ControlSystemState.ZPos"
	CONTROL_SYSTEM_STATE_FIELD_AIRSPEED     = "ControlSystemState.Airspeed"
	CONTROL_SYSTEM_STATE_FIELD_VEL_VARIANCE = "ControlSystemState.VelVariance"
	CONTROL_SYSTEM_STATE_FIELD_POS_VARIANCE = "ControlSystemState.PosVariance"
	CONTROL_SYSTEM_STATE_FIELD_Q            = "ControlSystemState.Q"
	CONTROL_SYSTEM_STATE_FIELD_ROLL_RATE    = "ControlSystemState.RollRate"
	CONTROL_SYSTEM_STATE_FIELD_PITCH_RATE   = "ControlSystemState.PitchRate"
	CONTROL_SYSTEM_STATE_FIELD_YAW_RATE     = "ControlSystemState.YawRate"
)

// ToMap (generated function)
func (m *ControlSystemState) Dict() map[string]interface{} {
	return map[string]interface{}{
		CONTROL_SYSTEM_STATE_FIELD_TIME_USEC:    m.TimeUsec,
		CONTROL_SYSTEM_STATE_FIELD_X_ACC:        m.XAcc,
		CONTROL_SYSTEM_STATE_FIELD_Y_ACC:        m.YAcc,
		CONTROL_SYSTEM_STATE_FIELD_Z_ACC:        m.ZAcc,
		CONTROL_SYSTEM_STATE_FIELD_X_VEL:        m.XVel,
		CONTROL_SYSTEM_STATE_FIELD_Y_VEL:        m.YVel,
		CONTROL_SYSTEM_STATE_FIELD_Z_VEL:        m.ZVel,
		CONTROL_SYSTEM_STATE_FIELD_X_POS:        m.XPos,
		CONTROL_SYSTEM_STATE_FIELD_Y_POS:        m.YPos,
		CONTROL_SYSTEM_STATE_FIELD_Z_POS:        m.ZPos,
		CONTROL_SYSTEM_STATE_FIELD_AIRSPEED:     m.Airspeed,
		CONTROL_SYSTEM_STATE_FIELD_VEL_VARIANCE: m.VelVariance,
		CONTROL_SYSTEM_STATE_FIELD_POS_VARIANCE: m.PosVariance,
		CONTROL_SYSTEM_STATE_FIELD_Q:            m.Q,
		CONTROL_SYSTEM_STATE_FIELD_ROLL_RATE:    m.RollRate,
		CONTROL_SYSTEM_STATE_FIELD_PITCH_RATE:   m.PitchRate,
		CONTROL_SYSTEM_STATE_FIELD_YAW_RATE:     m.YawRate,
	}
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
	CurrentConsumed  int32                `json:"CurrentConsumed" `   // [ mAh ] Consumed charge, -1: autopilot does not provide consumption estimate
	EnergyConsumed   int32                `json:"EnergyConsumed" `    // [ hJ ] Consumed energy, -1: autopilot does not provide energy consumption estimate
	Temperature      int16                `json:"Temperature" `       // [ cdegC ] Temperature of the battery. INT16_MAX for unknown temperature.
	Voltages         []uint16             `json:"Voltages" len:"10" ` // [ mV ] Battery voltage of cells 1 to 10 (see voltages_ext for cells 11-14). Cells in this field above the valid cell count for this battery should have the UINT16_MAX value. If individual cell voltages are unknown or not measured for this battery, then the overall battery voltage should be filled in cell 0, with all others set to UINT16_MAX. If the voltage of the battery is greater than (UINT16_MAX - 1), then cell 0 should be set to (UINT16_MAX - 1), and cell 1 to the remaining voltage. This can be extended to multiple cells if the total voltage is greater than 2 * (UINT16_MAX - 1).
	CurrentBattery   int16                `json:"CurrentBattery" `    // [ cA ] Battery current, -1: autopilot does not measure the current
	ID               uint8                `json:"ID" `                // Battery ID
	BatteryFunction  MAV_BATTERY_FUNCTION `json:"BatteryFunction" `   // Function of the battery
	Type             MAV_BATTERY_TYPE     `json:"Type" `              // Type (chemistry) of the battery
	BatteryRemaining int8                 `json:"BatteryRemaining" `  // [ % ] Remaining battery energy. Values: [0-100], -1: autopilot does not estimate the remaining battery.
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

const (
	BATTERY_STATUS_FIELD_CURRENT_CONSUMED  = "BatteryStatus.CurrentConsumed"
	BATTERY_STATUS_FIELD_ENERGY_CONSUMED   = "BatteryStatus.EnergyConsumed"
	BATTERY_STATUS_FIELD_TEMPERATURE       = "BatteryStatus.Temperature"
	BATTERY_STATUS_FIELD_VOLTAGES          = "BatteryStatus.Voltages"
	BATTERY_STATUS_FIELD_CURRENT_BATTERY   = "BatteryStatus.CurrentBattery"
	BATTERY_STATUS_FIELD_ID                = "BatteryStatus.ID"
	BATTERY_STATUS_FIELD_BATTERY_FUNCTION  = "BatteryStatus.BatteryFunction"
	BATTERY_STATUS_FIELD_TYPE              = "BatteryStatus.Type"
	BATTERY_STATUS_FIELD_BATTERY_REMAINING = "BatteryStatus.BatteryRemaining"
)

// ToMap (generated function)
func (m *BatteryStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		BATTERY_STATUS_FIELD_CURRENT_CONSUMED:  m.CurrentConsumed,
		BATTERY_STATUS_FIELD_ENERGY_CONSUMED:   m.EnergyConsumed,
		BATTERY_STATUS_FIELD_TEMPERATURE:       m.Temperature,
		BATTERY_STATUS_FIELD_VOLTAGES:          m.Voltages,
		BATTERY_STATUS_FIELD_CURRENT_BATTERY:   m.CurrentBattery,
		BATTERY_STATUS_FIELD_ID:                m.ID,
		BATTERY_STATUS_FIELD_BATTERY_FUNCTION:  m.BatteryFunction,
		BATTERY_STATUS_FIELD_TYPE:              m.Type,
		BATTERY_STATUS_FIELD_BATTERY_REMAINING: m.BatteryRemaining,
	}
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
	Capabilities            MAV_PROTOCOL_CAPABILITY `json:"Capabilities" `                    // Bitmap of capabilities
	UID                     uint64                  `json:"UID" `                             // UID if provided by hardware (see uid2)
	FlightSwVersion         uint32                  `json:"FlightSwVersion" `                 // Firmware version number
	MiddlewareSwVersion     uint32                  `json:"MiddlewareSwVersion" `             // Middleware version number
	OsSwVersion             uint32                  `json:"OsSwVersion" `                     // Operating system version number
	BoardVersion            uint32                  `json:"BoardVersion" `                    // HW / board version (last 8 bytes should be silicon ID, if any)
	VendorID                uint16                  `json:"VendorID" `                        // ID of the board vendor
	ProductID               uint16                  `json:"ProductID" `                       // ID of the product
	FlightCustomVersion     []uint8                 `json:"FlightCustomVersion" len:"8" `     // Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
	MiddlewareCustomVersion []uint8                 `json:"MiddlewareCustomVersion" len:"8" ` // Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
	OsCustomVersion         []uint8                 `json:"OsCustomVersion" len:"8" `         // Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
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

const (
	AUTOPILOT_VERSION_FIELD_CAPABILITIES              = "AutopilotVersion.Capabilities"
	AUTOPILOT_VERSION_FIELD_UID                       = "AutopilotVersion.UID"
	AUTOPILOT_VERSION_FIELD_FLIGHT_SW_VERSION         = "AutopilotVersion.FlightSwVersion"
	AUTOPILOT_VERSION_FIELD_MIDDLEWARE_SW_VERSION     = "AutopilotVersion.MiddlewareSwVersion"
	AUTOPILOT_VERSION_FIELD_OS_SW_VERSION             = "AutopilotVersion.OsSwVersion"
	AUTOPILOT_VERSION_FIELD_BOARD_VERSION             = "AutopilotVersion.BoardVersion"
	AUTOPILOT_VERSION_FIELD_VENDOR_ID                 = "AutopilotVersion.VendorID"
	AUTOPILOT_VERSION_FIELD_PRODUCT_ID                = "AutopilotVersion.ProductID"
	AUTOPILOT_VERSION_FIELD_FLIGHT_CUSTOM_VERSION     = "AutopilotVersion.FlightCustomVersion"
	AUTOPILOT_VERSION_FIELD_MIDDLEWARE_CUSTOM_VERSION = "AutopilotVersion.MiddlewareCustomVersion"
	AUTOPILOT_VERSION_FIELD_OS_CUSTOM_VERSION         = "AutopilotVersion.OsCustomVersion"
)

// ToMap (generated function)
func (m *AutopilotVersion) Dict() map[string]interface{} {
	return map[string]interface{}{
		AUTOPILOT_VERSION_FIELD_CAPABILITIES:              m.Capabilities,
		AUTOPILOT_VERSION_FIELD_UID:                       m.UID,
		AUTOPILOT_VERSION_FIELD_FLIGHT_SW_VERSION:         m.FlightSwVersion,
		AUTOPILOT_VERSION_FIELD_MIDDLEWARE_SW_VERSION:     m.MiddlewareSwVersion,
		AUTOPILOT_VERSION_FIELD_OS_SW_VERSION:             m.OsSwVersion,
		AUTOPILOT_VERSION_FIELD_BOARD_VERSION:             m.BoardVersion,
		AUTOPILOT_VERSION_FIELD_VENDOR_ID:                 m.VendorID,
		AUTOPILOT_VERSION_FIELD_PRODUCT_ID:                m.ProductID,
		AUTOPILOT_VERSION_FIELD_FLIGHT_CUSTOM_VERSION:     m.FlightCustomVersion,
		AUTOPILOT_VERSION_FIELD_MIDDLEWARE_CUSTOM_VERSION: m.MiddlewareCustomVersion,
		AUTOPILOT_VERSION_FIELD_OS_CUSTOM_VERSION:         m.OsCustomVersion,
	}
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
	TimeUsec  uint64    `json:"TimeUsec" `  // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	AngleX    float32   `json:"AngleX" `    // [ rad ] X-axis angular offset of the target from the center of the image
	AngleY    float32   `json:"AngleY" `    // [ rad ] Y-axis angular offset of the target from the center of the image
	Distance  float32   `json:"Distance" `  // [ m ] Distance to the target from the vehicle
	SizeX     float32   `json:"SizeX" `     // [ rad ] Size of target along x-axis
	SizeY     float32   `json:"SizeY" `     // [ rad ] Size of target along y-axis
	TargetNum uint8     `json:"TargetNum" ` // The ID of the target if multiple targets are present
	Frame     MAV_FRAME `json:"Frame" `     // Coordinate frame used for following fields.
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

const (
	LANDING_TARGET_FIELD_TIME_USEC  = "LandingTarget.TimeUsec"
	LANDING_TARGET_FIELD_ANGLE_X    = "LandingTarget.AngleX"
	LANDING_TARGET_FIELD_ANGLE_Y    = "LandingTarget.AngleY"
	LANDING_TARGET_FIELD_DISTANCE   = "LandingTarget.Distance"
	LANDING_TARGET_FIELD_SIZE_X     = "LandingTarget.SizeX"
	LANDING_TARGET_FIELD_SIZE_Y     = "LandingTarget.SizeY"
	LANDING_TARGET_FIELD_TARGET_NUM = "LandingTarget.TargetNum"
	LANDING_TARGET_FIELD_FRAME      = "LandingTarget.Frame"
)

// ToMap (generated function)
func (m *LandingTarget) Dict() map[string]interface{} {
	return map[string]interface{}{
		LANDING_TARGET_FIELD_TIME_USEC:  m.TimeUsec,
		LANDING_TARGET_FIELD_ANGLE_X:    m.AngleX,
		LANDING_TARGET_FIELD_ANGLE_Y:    m.AngleY,
		LANDING_TARGET_FIELD_DISTANCE:   m.Distance,
		LANDING_TARGET_FIELD_SIZE_X:     m.SizeX,
		LANDING_TARGET_FIELD_SIZE_Y:     m.SizeY,
		LANDING_TARGET_FIELD_TARGET_NUM: m.TargetNum,
		LANDING_TARGET_FIELD_FRAME:      m.Frame,
	}
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
	BreachTime   uint32       `json:"BreachTime" `   // [ ms ] Time (since boot) of last breach.
	BreachCount  uint16       `json:"BreachCount" `  // Number of fence breaches.
	BreachStatus uint8        `json:"BreachStatus" ` // Breach status (0 if currently inside fence, 1 if outside).
	BreachType   FENCE_BREACH `json:"BreachType" `   // Last breach type.
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

const (
	FENCE_STATUS_FIELD_BREACH_TIME   = "FenceStatus.BreachTime"
	FENCE_STATUS_FIELD_BREACH_COUNT  = "FenceStatus.BreachCount"
	FENCE_STATUS_FIELD_BREACH_STATUS = "FenceStatus.BreachStatus"
	FENCE_STATUS_FIELD_BREACH_TYPE   = "FenceStatus.BreachType"
)

// ToMap (generated function)
func (m *FenceStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		FENCE_STATUS_FIELD_BREACH_TIME:   m.BreachTime,
		FENCE_STATUS_FIELD_BREACH_COUNT:  m.BreachCount,
		FENCE_STATUS_FIELD_BREACH_STATUS: m.BreachStatus,
		FENCE_STATUS_FIELD_BREACH_TYPE:   m.BreachType,
	}
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
	Fitness   float32        `json:"Fitness" `   // [ mgauss ] RMS milligauss residuals.
	OfsX      float32        `json:"OfsX" `      // X offset.
	OfsY      float32        `json:"OfsY" `      // Y offset.
	OfsZ      float32        `json:"OfsZ" `      // Z offset.
	DiagX     float32        `json:"DiagX" `     // X diagonal (matrix 11).
	DiagY     float32        `json:"DiagY" `     // Y diagonal (matrix 22).
	DiagZ     float32        `json:"DiagZ" `     // Z diagonal (matrix 33).
	OffdiagX  float32        `json:"OffdiagX" `  // X off-diagonal (matrix 12 and 21).
	OffdiagY  float32        `json:"OffdiagY" `  // Y off-diagonal (matrix 13 and 31).
	OffdiagZ  float32        `json:"OffdiagZ" `  // Z off-diagonal (matrix 32 and 23).
	CompassID uint8          `json:"CompassID" ` // Compass being calibrated.
	CalMask   uint8          `json:"CalMask" `   // Bitmask of compasses being calibrated.
	CalStatus MAG_CAL_STATUS `json:"CalStatus" ` // Calibration Status.
	Autosaved uint8          `json:"Autosaved" ` // 0=requires a MAV_CMD_DO_ACCEPT_MAG_CAL, 1=saved to parameters.
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

const (
	MAG_CAL_REPORT_FIELD_FITNESS    = "MagCalReport.Fitness"
	MAG_CAL_REPORT_FIELD_OFS_X      = "MagCalReport.OfsX"
	MAG_CAL_REPORT_FIELD_OFS_Y      = "MagCalReport.OfsY"
	MAG_CAL_REPORT_FIELD_OFS_Z      = "MagCalReport.OfsZ"
	MAG_CAL_REPORT_FIELD_DIAG_X     = "MagCalReport.DiagX"
	MAG_CAL_REPORT_FIELD_DIAG_Y     = "MagCalReport.DiagY"
	MAG_CAL_REPORT_FIELD_DIAG_Z     = "MagCalReport.DiagZ"
	MAG_CAL_REPORT_FIELD_OFFDIAG_X  = "MagCalReport.OffdiagX"
	MAG_CAL_REPORT_FIELD_OFFDIAG_Y  = "MagCalReport.OffdiagY"
	MAG_CAL_REPORT_FIELD_OFFDIAG_Z  = "MagCalReport.OffdiagZ"
	MAG_CAL_REPORT_FIELD_COMPASS_ID = "MagCalReport.CompassID"
	MAG_CAL_REPORT_FIELD_CAL_MASK   = "MagCalReport.CalMask"
	MAG_CAL_REPORT_FIELD_CAL_STATUS = "MagCalReport.CalStatus"
	MAG_CAL_REPORT_FIELD_AUTOSAVED  = "MagCalReport.Autosaved"
)

// ToMap (generated function)
func (m *MagCalReport) Dict() map[string]interface{} {
	return map[string]interface{}{
		MAG_CAL_REPORT_FIELD_FITNESS:    m.Fitness,
		MAG_CAL_REPORT_FIELD_OFS_X:      m.OfsX,
		MAG_CAL_REPORT_FIELD_OFS_Y:      m.OfsY,
		MAG_CAL_REPORT_FIELD_OFS_Z:      m.OfsZ,
		MAG_CAL_REPORT_FIELD_DIAG_X:     m.DiagX,
		MAG_CAL_REPORT_FIELD_DIAG_Y:     m.DiagY,
		MAG_CAL_REPORT_FIELD_DIAG_Z:     m.DiagZ,
		MAG_CAL_REPORT_FIELD_OFFDIAG_X:  m.OffdiagX,
		MAG_CAL_REPORT_FIELD_OFFDIAG_Y:  m.OffdiagY,
		MAG_CAL_REPORT_FIELD_OFFDIAG_Z:  m.OffdiagZ,
		MAG_CAL_REPORT_FIELD_COMPASS_ID: m.CompassID,
		MAG_CAL_REPORT_FIELD_CAL_MASK:   m.CalMask,
		MAG_CAL_REPORT_FIELD_CAL_STATUS: m.CalStatus,
		MAG_CAL_REPORT_FIELD_AUTOSAVED:  m.Autosaved,
	}
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
	EcuIndex                  float32 `json:"EcuIndex" `                  // ECU index
	Rpm                       float32 `json:"Rpm" `                       // RPM
	FuelConsumed              float32 `json:"FuelConsumed" `              // [ cm^3 ] Fuel consumed
	FuelFlow                  float32 `json:"FuelFlow" `                  // [ cm^3/min ] Fuel flow rate
	EngineLoad                float32 `json:"EngineLoad" `                // [ % ] Engine load
	ThrottlePosition          float32 `json:"ThrottlePosition" `          // [ % ] Throttle position
	SparkDwellTime            float32 `json:"SparkDwellTime" `            // [ ms ] Spark dwell time
	BarometricPressure        float32 `json:"BarometricPressure" `        // [ kPa ] Barometric pressure
	IntakeManifoldPressure    float32 `json:"IntakeManifoldPressure" `    // [ kPa ] Intake manifold pressure(
	IntakeManifoldTemperature float32 `json:"IntakeManifoldTemperature" ` // [ degC ] Intake manifold temperature
	CylinderHeadTemperature   float32 `json:"CylinderHeadTemperature" `   // [ degC ] Cylinder head temperature
	IgnitionTiming            float32 `json:"IgnitionTiming" `            // [ deg ] Ignition timing (Crank angle degrees)
	InjectionTime             float32 `json:"InjectionTime" `             // [ ms ] Injection time
	ExhaustGasTemperature     float32 `json:"ExhaustGasTemperature" `     // [ degC ] Exhaust gas temperature
	ThrottleOut               float32 `json:"ThrottleOut" `               // [ % ] Output throttle
	PtCompensation            float32 `json:"PtCompensation" `            // Pressure/temperature compensation
	Health                    uint8   `json:"Health" `                    // EFI health status
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

const (
	EFI_STATUS_FIELD_ECU_INDEX                   = "EfiStatus.EcuIndex"
	EFI_STATUS_FIELD_RPM                         = "EfiStatus.Rpm"
	EFI_STATUS_FIELD_FUEL_CONSUMED               = "EfiStatus.FuelConsumed"
	EFI_STATUS_FIELD_FUEL_FLOW                   = "EfiStatus.FuelFlow"
	EFI_STATUS_FIELD_ENGINE_LOAD                 = "EfiStatus.EngineLoad"
	EFI_STATUS_FIELD_THROTTLE_POSITION           = "EfiStatus.ThrottlePosition"
	EFI_STATUS_FIELD_SPARK_DWELL_TIME            = "EfiStatus.SparkDwellTime"
	EFI_STATUS_FIELD_BAROMETRIC_PRESSURE         = "EfiStatus.BarometricPressure"
	EFI_STATUS_FIELD_INTAKE_MANIFOLD_PRESSURE    = "EfiStatus.IntakeManifoldPressure"
	EFI_STATUS_FIELD_INTAKE_MANIFOLD_TEMPERATURE = "EfiStatus.IntakeManifoldTemperature"
	EFI_STATUS_FIELD_CYLINDER_HEAD_TEMPERATURE   = "EfiStatus.CylinderHeadTemperature"
	EFI_STATUS_FIELD_IGNITION_TIMING             = "EfiStatus.IgnitionTiming"
	EFI_STATUS_FIELD_INJECTION_TIME              = "EfiStatus.InjectionTime"
	EFI_STATUS_FIELD_EXHAUST_GAS_TEMPERATURE     = "EfiStatus.ExhaustGasTemperature"
	EFI_STATUS_FIELD_THROTTLE_OUT                = "EfiStatus.ThrottleOut"
	EFI_STATUS_FIELD_PT_COMPENSATION             = "EfiStatus.PtCompensation"
	EFI_STATUS_FIELD_HEALTH                      = "EfiStatus.Health"
)

// ToMap (generated function)
func (m *EfiStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		EFI_STATUS_FIELD_ECU_INDEX:                   m.EcuIndex,
		EFI_STATUS_FIELD_RPM:                         m.Rpm,
		EFI_STATUS_FIELD_FUEL_CONSUMED:               m.FuelConsumed,
		EFI_STATUS_FIELD_FUEL_FLOW:                   m.FuelFlow,
		EFI_STATUS_FIELD_ENGINE_LOAD:                 m.EngineLoad,
		EFI_STATUS_FIELD_THROTTLE_POSITION:           m.ThrottlePosition,
		EFI_STATUS_FIELD_SPARK_DWELL_TIME:            m.SparkDwellTime,
		EFI_STATUS_FIELD_BAROMETRIC_PRESSURE:         m.BarometricPressure,
		EFI_STATUS_FIELD_INTAKE_MANIFOLD_PRESSURE:    m.IntakeManifoldPressure,
		EFI_STATUS_FIELD_INTAKE_MANIFOLD_TEMPERATURE: m.IntakeManifoldTemperature,
		EFI_STATUS_FIELD_CYLINDER_HEAD_TEMPERATURE:   m.CylinderHeadTemperature,
		EFI_STATUS_FIELD_IGNITION_TIMING:             m.IgnitionTiming,
		EFI_STATUS_FIELD_INJECTION_TIME:              m.InjectionTime,
		EFI_STATUS_FIELD_EXHAUST_GAS_TEMPERATURE:     m.ExhaustGasTemperature,
		EFI_STATUS_FIELD_THROTTLE_OUT:                m.ThrottleOut,
		EFI_STATUS_FIELD_PT_COMPENSATION:             m.PtCompensation,
		EFI_STATUS_FIELD_HEALTH:                      m.Health,
	}
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
	TimeUsec         uint64                 `json:"TimeUsec" `         // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	VelRatio         float32                `json:"VelRatio" `         // Velocity innovation test ratio
	PosHorizRatio    float32                `json:"PosHorizRatio" `    // Horizontal position innovation test ratio
	PosVertRatio     float32                `json:"PosVertRatio" `     // Vertical position innovation test ratio
	MagRatio         float32                `json:"MagRatio" `         // Magnetometer innovation test ratio
	HaglRatio        float32                `json:"HaglRatio" `        // Height above terrain innovation test ratio
	TasRatio         float32                `json:"TasRatio" `         // True airspeed innovation test ratio
	PosHorizAccuracy float32                `json:"PosHorizAccuracy" ` // [ m ] Horizontal position 1-STD accuracy relative to the EKF local origin
	PosVertAccuracy  float32                `json:"PosVertAccuracy" `  // [ m ] Vertical position 1-STD accuracy relative to the EKF local origin
	Flags            ESTIMATOR_STATUS_FLAGS `json:"Flags" `            // Bitmap indicating which EKF outputs are valid.
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

const (
	ESTIMATOR_STATUS_FIELD_TIME_USEC          = "EstimatorStatus.TimeUsec"
	ESTIMATOR_STATUS_FIELD_VEL_RATIO          = "EstimatorStatus.VelRatio"
	ESTIMATOR_STATUS_FIELD_POS_HORIZ_RATIO    = "EstimatorStatus.PosHorizRatio"
	ESTIMATOR_STATUS_FIELD_POS_VERT_RATIO     = "EstimatorStatus.PosVertRatio"
	ESTIMATOR_STATUS_FIELD_MAG_RATIO          = "EstimatorStatus.MagRatio"
	ESTIMATOR_STATUS_FIELD_HAGL_RATIO         = "EstimatorStatus.HaglRatio"
	ESTIMATOR_STATUS_FIELD_TAS_RATIO          = "EstimatorStatus.TasRatio"
	ESTIMATOR_STATUS_FIELD_POS_HORIZ_ACCURACY = "EstimatorStatus.PosHorizAccuracy"
	ESTIMATOR_STATUS_FIELD_POS_VERT_ACCURACY  = "EstimatorStatus.PosVertAccuracy"
	ESTIMATOR_STATUS_FIELD_FLAGS              = "EstimatorStatus.Flags"
)

// ToMap (generated function)
func (m *EstimatorStatus) Dict() map[string]interface{} {
	return map[string]interface{}{
		ESTIMATOR_STATUS_FIELD_TIME_USEC:          m.TimeUsec,
		ESTIMATOR_STATUS_FIELD_VEL_RATIO:          m.VelRatio,
		ESTIMATOR_STATUS_FIELD_POS_HORIZ_RATIO:    m.PosHorizRatio,
		ESTIMATOR_STATUS_FIELD_POS_VERT_RATIO:     m.PosVertRatio,
		ESTIMATOR_STATUS_FIELD_MAG_RATIO:          m.MagRatio,
		ESTIMATOR_STATUS_FIELD_HAGL_RATIO:         m.HaglRatio,
		ESTIMATOR_STATUS_FIELD_TAS_RATIO:          m.TasRatio,
		ESTIMATOR_STATUS_FIELD_POS_HORIZ_ACCURACY: m.PosHorizAccuracy,
		ESTIMATOR_STATUS_FIELD_POS_VERT_ACCURACY:  m.PosVertAccuracy,
		ESTIMATOR_STATUS_FIELD_FLAGS:              m.Flags,
	}
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
	TimeUsec      uint64  `json:"TimeUsec" `      // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	WindX         float32 `json:"WindX" `         // [ m/s ] Wind in X (NED) direction
	WindY         float32 `json:"WindY" `         // [ m/s ] Wind in Y (NED) direction
	WindZ         float32 `json:"WindZ" `         // [ m/s ] Wind in Z (NED) direction
	VarHoriz      float32 `json:"VarHoriz" `      // [ m/s ] Variability of the wind in XY. RMS of a 1 Hz lowpassed wind estimate.
	VarVert       float32 `json:"VarVert" `       // [ m/s ] Variability of the wind in Z. RMS of a 1 Hz lowpassed wind estimate.
	WindAlt       float32 `json:"WindAlt" `       // [ m ] Altitude (MSL) that this measurement was taken at
	HorizAccuracy float32 `json:"HorizAccuracy" ` // [ m ] Horizontal speed 1-STD accuracy
	VertAccuracy  float32 `json:"VertAccuracy" `  // [ m ] Vertical speed 1-STD accuracy
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

const (
	WIND_COV_FIELD_TIME_USEC      = "WindCov.TimeUsec"
	WIND_COV_FIELD_WIND_X         = "WindCov.WindX"
	WIND_COV_FIELD_WIND_Y         = "WindCov.WindY"
	WIND_COV_FIELD_WIND_Z         = "WindCov.WindZ"
	WIND_COV_FIELD_VAR_HORIZ      = "WindCov.VarHoriz"
	WIND_COV_FIELD_VAR_VERT       = "WindCov.VarVert"
	WIND_COV_FIELD_WIND_ALT       = "WindCov.WindAlt"
	WIND_COV_FIELD_HORIZ_ACCURACY = "WindCov.HorizAccuracy"
	WIND_COV_FIELD_VERT_ACCURACY  = "WindCov.VertAccuracy"
)

// ToMap (generated function)
func (m *WindCov) Dict() map[string]interface{} {
	return map[string]interface{}{
		WIND_COV_FIELD_TIME_USEC:      m.TimeUsec,
		WIND_COV_FIELD_WIND_X:         m.WindX,
		WIND_COV_FIELD_WIND_Y:         m.WindY,
		WIND_COV_FIELD_WIND_Z:         m.WindZ,
		WIND_COV_FIELD_VAR_HORIZ:      m.VarHoriz,
		WIND_COV_FIELD_VAR_VERT:       m.VarVert,
		WIND_COV_FIELD_WIND_ALT:       m.WindAlt,
		WIND_COV_FIELD_HORIZ_ACCURACY: m.HorizAccuracy,
		WIND_COV_FIELD_VERT_ACCURACY:  m.VertAccuracy,
	}
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
	TimeUsec          uint64                 `json:"TimeUsec" `          // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeWeekMs        uint32                 `json:"TimeWeekMs" `        // [ ms ] GPS time (from start of GPS week)
	Lat               int32                  `json:"Lat" `               // [ degE7 ] Latitude (WGS84)
	Lon               int32                  `json:"Lon" `               // [ degE7 ] Longitude (WGS84)
	Alt               float32                `json:"Alt" `               // [ m ] Altitude (MSL). Positive for up.
	Hdop              float32                `json:"Hdop" `              // GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX
	Vdop              float32                `json:"Vdop" `              // GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX
	Vn                float32                `json:"Vn" `                // [ m/s ] GPS velocity in north direction in earth-fixed NED frame
	Ve                float32                `json:"Ve" `                // [ m/s ] GPS velocity in east direction in earth-fixed NED frame
	Vd                float32                `json:"Vd" `                // [ m/s ] GPS velocity in down direction in earth-fixed NED frame
	SpeedAccuracy     float32                `json:"SpeedAccuracy" `     // [ m/s ] GPS speed accuracy
	HorizAccuracy     float32                `json:"HorizAccuracy" `     // [ m ] GPS horizontal accuracy
	VertAccuracy      float32                `json:"VertAccuracy" `      // [ m ] GPS vertical accuracy
	IgnoreFlags       GPS_INPUT_IGNORE_FLAGS `json:"IgnoreFlags" `       // Bitmap indicating which GPS input flags fields to ignore.  All other fields must be provided.
	TimeWeek          uint16                 `json:"TimeWeek" `          // GPS week number
	GpsID             uint8                  `json:"GpsID" `             // ID of the GPS for multiple GPS inputs
	FixType           uint8                  `json:"FixType" `           // 0-1: no fix, 2: 2D fix, 3: 3D fix. 4: 3D with DGPS. 5: 3D with RTK
	SatellitesVisible uint8                  `json:"SatellitesVisible" ` // Number of satellites visible.
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

const (
	GPS_INPUT_FIELD_TIME_USEC          = "GpsInput.TimeUsec"
	GPS_INPUT_FIELD_TIME_WEEK_MS       = "GpsInput.TimeWeekMs"
	GPS_INPUT_FIELD_LAT                = "GpsInput.Lat"
	GPS_INPUT_FIELD_LON                = "GpsInput.Lon"
	GPS_INPUT_FIELD_ALT                = "GpsInput.Alt"
	GPS_INPUT_FIELD_HDOP               = "GpsInput.Hdop"
	GPS_INPUT_FIELD_VDOP               = "GpsInput.Vdop"
	GPS_INPUT_FIELD_VN                 = "GpsInput.Vn"
	GPS_INPUT_FIELD_VE                 = "GpsInput.Ve"
	GPS_INPUT_FIELD_VD                 = "GpsInput.Vd"
	GPS_INPUT_FIELD_SPEED_ACCURACY     = "GpsInput.SpeedAccuracy"
	GPS_INPUT_FIELD_HORIZ_ACCURACY     = "GpsInput.HorizAccuracy"
	GPS_INPUT_FIELD_VERT_ACCURACY      = "GpsInput.VertAccuracy"
	GPS_INPUT_FIELD_IGNORE_FLAGS       = "GpsInput.IgnoreFlags"
	GPS_INPUT_FIELD_TIME_WEEK          = "GpsInput.TimeWeek"
	GPS_INPUT_FIELD_GPS_ID             = "GpsInput.GpsID"
	GPS_INPUT_FIELD_FIX_TYPE           = "GpsInput.FixType"
	GPS_INPUT_FIELD_SATELLITES_VISIBLE = "GpsInput.SatellitesVisible"
)

// ToMap (generated function)
func (m *GpsInput) Dict() map[string]interface{} {
	return map[string]interface{}{
		GPS_INPUT_FIELD_TIME_USEC:          m.TimeUsec,
		GPS_INPUT_FIELD_TIME_WEEK_MS:       m.TimeWeekMs,
		GPS_INPUT_FIELD_LAT:                m.Lat,
		GPS_INPUT_FIELD_LON:                m.Lon,
		GPS_INPUT_FIELD_ALT:                m.Alt,
		GPS_INPUT_FIELD_HDOP:               m.Hdop,
		GPS_INPUT_FIELD_VDOP:               m.Vdop,
		GPS_INPUT_FIELD_VN:                 m.Vn,
		GPS_INPUT_FIELD_VE:                 m.Ve,
		GPS_INPUT_FIELD_VD:                 m.Vd,
		GPS_INPUT_FIELD_SPEED_ACCURACY:     m.SpeedAccuracy,
		GPS_INPUT_FIELD_HORIZ_ACCURACY:     m.HorizAccuracy,
		GPS_INPUT_FIELD_VERT_ACCURACY:      m.VertAccuracy,
		GPS_INPUT_FIELD_IGNORE_FLAGS:       m.IgnoreFlags,
		GPS_INPUT_FIELD_TIME_WEEK:          m.TimeWeek,
		GPS_INPUT_FIELD_GPS_ID:             m.GpsID,
		GPS_INPUT_FIELD_FIX_TYPE:           m.FixType,
		GPS_INPUT_FIELD_SATELLITES_VISIBLE: m.SatellitesVisible,
	}
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
	Flags uint8   `json:"Flags" `          // LSB: 1 means message is fragmented, next 2 bits are the fragment ID, the remaining 5 bits are used for the sequence ID. Messages are only to be flushed to the GPS when the entire message has been reconstructed on the autopilot. The fragment ID specifies which order the fragments should be assembled into a buffer, while the sequence ID is used to detect a mismatch between different buffers. The buffer is considered fully reconstructed when either all 4 fragments are present, or all the fragments before the first fragment with a non full payload is received. This management is used to ensure that normal GPS operation doesn't corrupt RTCM data, and to recover from a unreliable transport delivery order.
	Len   uint8   `json:"Len" `            // [ bytes ] data length
	Data  []uint8 `json:"Data" len:"180" ` // RTCM message (may be fragmented)
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

const (
	GPS_RTCM_DATA_FIELD_FLAGS = "GpsRtcmData.Flags"
	GPS_RTCM_DATA_FIELD_LEN   = "GpsRtcmData.Len"
	GPS_RTCM_DATA_FIELD_DATA  = "GpsRtcmData.Data"
)

// ToMap (generated function)
func (m *GpsRtcmData) Dict() map[string]interface{} {
	return map[string]interface{}{
		GPS_RTCM_DATA_FIELD_FLAGS: m.Flags,
		GPS_RTCM_DATA_FIELD_LEN:   m.Len,
		GPS_RTCM_DATA_FIELD_DATA:  m.Data,
	}
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
	CustomMode       uint32           `json:"CustomMode" `       // A bitfield for use for autopilot-specific flags.
	Latitude         int32            `json:"Latitude" `         // [ degE7 ] Latitude
	Longitude        int32            `json:"Longitude" `        // [ degE7 ] Longitude
	Roll             int16            `json:"Roll" `             // [ cdeg ] roll
	Pitch            int16            `json:"Pitch" `            // [ cdeg ] pitch
	Heading          uint16           `json:"Heading" `          // [ cdeg ] heading
	HeadingSp        int16            `json:"HeadingSp" `        // [ cdeg ] heading setpoint
	AltitudeAmsl     int16            `json:"AltitudeAmsl" `     // [ m ] Altitude above mean sea level
	AltitudeSp       int16            `json:"AltitudeSp" `       // [ m ] Altitude setpoint relative to the home position
	WpDistance       uint16           `json:"WpDistance" `       // [ m ] distance to target
	BaseMode         MAV_MODE_FLAG    `json:"BaseMode" `         // Bitmap of enabled system modes.
	LandedState      MAV_LANDED_STATE `json:"LandedState" `      // The landed state. Is set to MAV_LANDED_STATE_UNDEFINED if landed state is unknown.
	Throttle         int8             `json:"Throttle" `         // [ % ] throttle (percentage)
	Airspeed         uint8            `json:"Airspeed" `         // [ m/s ] airspeed
	AirspeedSp       uint8            `json:"AirspeedSp" `       // [ m/s ] airspeed setpoint
	Groundspeed      uint8            `json:"Groundspeed" `      // [ m/s ] groundspeed
	ClimbRate        int8             `json:"ClimbRate" `        // [ m/s ] climb rate
	GpsNsat          uint8            `json:"GpsNsat" `          // Number of satellites visible. If unknown, set to 255
	GpsFixType       GPS_FIX_TYPE     `json:"GpsFixType" `       // GPS Fix type.
	BatteryRemaining uint8            `json:"BatteryRemaining" ` // [ % ] Remaining battery (percentage)
	Temperature      int8             `json:"Temperature" `      // [ degC ] Autopilot temperature (degrees C)
	TemperatureAir   int8             `json:"TemperatureAir" `   // [ degC ] Air temperature (degrees C) from airspeed sensor
	Failsafe         uint8            `json:"Failsafe" `         // failsafe (each bit represents a failsafe where 0=ok, 1=failsafe active (bit0:RC, bit1:batt, bit2:GPS, bit3:GCS, bit4:fence)
	WpNum            uint8            `json:"WpNum" `            // current waypoint number
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

const (
	HIGH_LATENCY_FIELD_CUSTOM_MODE       = "HighLatency.CustomMode"
	HIGH_LATENCY_FIELD_LATITUDE          = "HighLatency.Latitude"
	HIGH_LATENCY_FIELD_LONGITUDE         = "HighLatency.Longitude"
	HIGH_LATENCY_FIELD_ROLL              = "HighLatency.Roll"
	HIGH_LATENCY_FIELD_PITCH             = "HighLatency.Pitch"
	HIGH_LATENCY_FIELD_HEADING           = "HighLatency.Heading"
	HIGH_LATENCY_FIELD_HEADING_SP        = "HighLatency.HeadingSp"
	HIGH_LATENCY_FIELD_ALTITUDE_AMSL     = "HighLatency.AltitudeAmsl"
	HIGH_LATENCY_FIELD_ALTITUDE_SP       = "HighLatency.AltitudeSp"
	HIGH_LATENCY_FIELD_WP_DISTANCE       = "HighLatency.WpDistance"
	HIGH_LATENCY_FIELD_BASE_MODE         = "HighLatency.BaseMode"
	HIGH_LATENCY_FIELD_LANDED_STATE      = "HighLatency.LandedState"
	HIGH_LATENCY_FIELD_THROTTLE          = "HighLatency.Throttle"
	HIGH_LATENCY_FIELD_AIRSPEED          = "HighLatency.Airspeed"
	HIGH_LATENCY_FIELD_AIRSPEED_SP       = "HighLatency.AirspeedSp"
	HIGH_LATENCY_FIELD_GROUNDSPEED       = "HighLatency.Groundspeed"
	HIGH_LATENCY_FIELD_CLIMB_RATE        = "HighLatency.ClimbRate"
	HIGH_LATENCY_FIELD_GPS_NSAT          = "HighLatency.GpsNsat"
	HIGH_LATENCY_FIELD_GPS_FIX_TYPE      = "HighLatency.GpsFixType"
	HIGH_LATENCY_FIELD_BATTERY_REMAINING = "HighLatency.BatteryRemaining"
	HIGH_LATENCY_FIELD_TEMPERATURE       = "HighLatency.Temperature"
	HIGH_LATENCY_FIELD_TEMPERATURE_AIR   = "HighLatency.TemperatureAir"
	HIGH_LATENCY_FIELD_FAILSAFE          = "HighLatency.Failsafe"
	HIGH_LATENCY_FIELD_WP_NUM            = "HighLatency.WpNum"
)

// ToMap (generated function)
func (m *HighLatency) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIGH_LATENCY_FIELD_CUSTOM_MODE:       m.CustomMode,
		HIGH_LATENCY_FIELD_LATITUDE:          m.Latitude,
		HIGH_LATENCY_FIELD_LONGITUDE:         m.Longitude,
		HIGH_LATENCY_FIELD_ROLL:              m.Roll,
		HIGH_LATENCY_FIELD_PITCH:             m.Pitch,
		HIGH_LATENCY_FIELD_HEADING:           m.Heading,
		HIGH_LATENCY_FIELD_HEADING_SP:        m.HeadingSp,
		HIGH_LATENCY_FIELD_ALTITUDE_AMSL:     m.AltitudeAmsl,
		HIGH_LATENCY_FIELD_ALTITUDE_SP:       m.AltitudeSp,
		HIGH_LATENCY_FIELD_WP_DISTANCE:       m.WpDistance,
		HIGH_LATENCY_FIELD_BASE_MODE:         m.BaseMode,
		HIGH_LATENCY_FIELD_LANDED_STATE:      m.LandedState,
		HIGH_LATENCY_FIELD_THROTTLE:          m.Throttle,
		HIGH_LATENCY_FIELD_AIRSPEED:          m.Airspeed,
		HIGH_LATENCY_FIELD_AIRSPEED_SP:       m.AirspeedSp,
		HIGH_LATENCY_FIELD_GROUNDSPEED:       m.Groundspeed,
		HIGH_LATENCY_FIELD_CLIMB_RATE:        m.ClimbRate,
		HIGH_LATENCY_FIELD_GPS_NSAT:          m.GpsNsat,
		HIGH_LATENCY_FIELD_GPS_FIX_TYPE:      m.GpsFixType,
		HIGH_LATENCY_FIELD_BATTERY_REMAINING: m.BatteryRemaining,
		HIGH_LATENCY_FIELD_TEMPERATURE:       m.Temperature,
		HIGH_LATENCY_FIELD_TEMPERATURE_AIR:   m.TemperatureAir,
		HIGH_LATENCY_FIELD_FAILSAFE:          m.Failsafe,
		HIGH_LATENCY_FIELD_WP_NUM:            m.WpNum,
	}
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
	Timestamp      uint32          `json:"Timestamp" `      // [ ms ] Timestamp (milliseconds since boot or Unix epoch)
	Latitude       int32           `json:"Latitude" `       // [ degE7 ] Latitude
	Longitude      int32           `json:"Longitude" `      // [ degE7 ] Longitude
	CustomMode     uint16          `json:"CustomMode" `     // A bitfield for use for autopilot-specific flags (2 byte version).
	Altitude       int16           `json:"Altitude" `       // [ m ] Altitude above mean sea level
	TargetAltitude int16           `json:"TargetAltitude" ` // [ m ] Altitude setpoint
	TargetDistance uint16          `json:"TargetDistance" ` // [ dam ] Distance to target waypoint or position
	WpNum          uint16          `json:"WpNum" `          // Current waypoint number
	FailureFlags   HL_FAILURE_FLAG `json:"FailureFlags" `   // Bitmap of failure flags.
	Type           MAV_TYPE        `json:"Type" `           // Type of the MAV (quadrotor, helicopter, etc.)
	Autopilot      MAV_AUTOPILOT   `json:"Autopilot" `      // Autopilot type / class. Use MAV_AUTOPILOT_INVALID for components that are not flight controllers.
	Heading        uint8           `json:"Heading" `        // [ deg/2 ] Heading
	TargetHeading  uint8           `json:"TargetHeading" `  // [ deg/2 ] Heading setpoint
	Throttle       uint8           `json:"Throttle" `       // [ % ] Throttle
	Airspeed       uint8           `json:"Airspeed" `       // [ m/s*5 ] Airspeed
	AirspeedSp     uint8           `json:"AirspeedSp" `     // [ m/s*5 ] Airspeed setpoint
	Groundspeed    uint8           `json:"Groundspeed" `    // [ m/s*5 ] Groundspeed
	Windspeed      uint8           `json:"Windspeed" `      // [ m/s*5 ] Windspeed
	WindHeading    uint8           `json:"WindHeading" `    // [ deg/2 ] Wind heading
	Eph            uint8           `json:"Eph" `            // [ dm ] Maximum error horizontal position since last message
	Epv            uint8           `json:"Epv" `            // [ dm ] Maximum error vertical position since last message
	TemperatureAir int8            `json:"TemperatureAir" ` // [ degC ] Air temperature from airspeed sensor
	ClimbRate      int8            `json:"ClimbRate" `      // [ dm/s ] Maximum climb rate magnitude since last message
	Battery        int8            `json:"Battery" `        // [ % ] Battery level (-1 if field not provided).
	Custom0        int8            `json:"Custom0" `        // Field for custom payload.
	Custom1        int8            `json:"Custom1" `        // Field for custom payload.
	Custom2        int8            `json:"Custom2" `        // Field for custom payload.
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

const (
	HIGH_LATENCY2_FIELD_TIMESTAMP       = "HighLatency2.Timestamp"
	HIGH_LATENCY2_FIELD_LATITUDE        = "HighLatency2.Latitude"
	HIGH_LATENCY2_FIELD_LONGITUDE       = "HighLatency2.Longitude"
	HIGH_LATENCY2_FIELD_CUSTOM_MODE     = "HighLatency2.CustomMode"
	HIGH_LATENCY2_FIELD_ALTITUDE        = "HighLatency2.Altitude"
	HIGH_LATENCY2_FIELD_TARGET_ALTITUDE = "HighLatency2.TargetAltitude"
	HIGH_LATENCY2_FIELD_TARGET_DISTANCE = "HighLatency2.TargetDistance"
	HIGH_LATENCY2_FIELD_WP_NUM          = "HighLatency2.WpNum"
	HIGH_LATENCY2_FIELD_FAILURE_FLAGS   = "HighLatency2.FailureFlags"
	HIGH_LATENCY2_FIELD_TYPE            = "HighLatency2.Type"
	HIGH_LATENCY2_FIELD_AUTOPILOT       = "HighLatency2.Autopilot"
	HIGH_LATENCY2_FIELD_HEADING         = "HighLatency2.Heading"
	HIGH_LATENCY2_FIELD_TARGET_HEADING  = "HighLatency2.TargetHeading"
	HIGH_LATENCY2_FIELD_THROTTLE        = "HighLatency2.Throttle"
	HIGH_LATENCY2_FIELD_AIRSPEED        = "HighLatency2.Airspeed"
	HIGH_LATENCY2_FIELD_AIRSPEED_SP     = "HighLatency2.AirspeedSp"
	HIGH_LATENCY2_FIELD_GROUNDSPEED     = "HighLatency2.Groundspeed"
	HIGH_LATENCY2_FIELD_WINDSPEED       = "HighLatency2.Windspeed"
	HIGH_LATENCY2_FIELD_WIND_HEADING    = "HighLatency2.WindHeading"
	HIGH_LATENCY2_FIELD_EPH             = "HighLatency2.Eph"
	HIGH_LATENCY2_FIELD_EPV             = "HighLatency2.Epv"
	HIGH_LATENCY2_FIELD_TEMPERATURE_AIR = "HighLatency2.TemperatureAir"
	HIGH_LATENCY2_FIELD_CLIMB_RATE      = "HighLatency2.ClimbRate"
	HIGH_LATENCY2_FIELD_BATTERY         = "HighLatency2.Battery"
	HIGH_LATENCY2_FIELD_CUSTOM0         = "HighLatency2.Custom0"
	HIGH_LATENCY2_FIELD_CUSTOM1         = "HighLatency2.Custom1"
	HIGH_LATENCY2_FIELD_CUSTOM2         = "HighLatency2.Custom2"
)

// ToMap (generated function)
func (m *HighLatency2) Dict() map[string]interface{} {
	return map[string]interface{}{
		HIGH_LATENCY2_FIELD_TIMESTAMP:       m.Timestamp,
		HIGH_LATENCY2_FIELD_LATITUDE:        m.Latitude,
		HIGH_LATENCY2_FIELD_LONGITUDE:       m.Longitude,
		HIGH_LATENCY2_FIELD_CUSTOM_MODE:     m.CustomMode,
		HIGH_LATENCY2_FIELD_ALTITUDE:        m.Altitude,
		HIGH_LATENCY2_FIELD_TARGET_ALTITUDE: m.TargetAltitude,
		HIGH_LATENCY2_FIELD_TARGET_DISTANCE: m.TargetDistance,
		HIGH_LATENCY2_FIELD_WP_NUM:          m.WpNum,
		HIGH_LATENCY2_FIELD_FAILURE_FLAGS:   m.FailureFlags,
		HIGH_LATENCY2_FIELD_TYPE:            m.Type,
		HIGH_LATENCY2_FIELD_AUTOPILOT:       m.Autopilot,
		HIGH_LATENCY2_FIELD_HEADING:         m.Heading,
		HIGH_LATENCY2_FIELD_TARGET_HEADING:  m.TargetHeading,
		HIGH_LATENCY2_FIELD_THROTTLE:        m.Throttle,
		HIGH_LATENCY2_FIELD_AIRSPEED:        m.Airspeed,
		HIGH_LATENCY2_FIELD_AIRSPEED_SP:     m.AirspeedSp,
		HIGH_LATENCY2_FIELD_GROUNDSPEED:     m.Groundspeed,
		HIGH_LATENCY2_FIELD_WINDSPEED:       m.Windspeed,
		HIGH_LATENCY2_FIELD_WIND_HEADING:    m.WindHeading,
		HIGH_LATENCY2_FIELD_EPH:             m.Eph,
		HIGH_LATENCY2_FIELD_EPV:             m.Epv,
		HIGH_LATENCY2_FIELD_TEMPERATURE_AIR: m.TemperatureAir,
		HIGH_LATENCY2_FIELD_CLIMB_RATE:      m.ClimbRate,
		HIGH_LATENCY2_FIELD_BATTERY:         m.Battery,
		HIGH_LATENCY2_FIELD_CUSTOM0:         m.Custom0,
		HIGH_LATENCY2_FIELD_CUSTOM1:         m.Custom1,
		HIGH_LATENCY2_FIELD_CUSTOM2:         m.Custom2,
	}
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
	TimeUsec   uint64  `json:"TimeUsec" `   // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	VibrationX float32 `json:"VibrationX" ` // Vibration levels on X-axis
	VibrationY float32 `json:"VibrationY" ` // Vibration levels on Y-axis
	VibrationZ float32 `json:"VibrationZ" ` // Vibration levels on Z-axis
	Clipping0  uint32  `json:"Clipping0" `  // first accelerometer clipping count
	Clipping1  uint32  `json:"Clipping1" `  // second accelerometer clipping count
	Clipping2  uint32  `json:"Clipping2" `  // third accelerometer clipping count
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

const (
	VIBRATION_FIELD_TIME_USEC   = "Vibration.TimeUsec"
	VIBRATION_FIELD_VIBRATION_X = "Vibration.VibrationX"
	VIBRATION_FIELD_VIBRATION_Y = "Vibration.VibrationY"
	VIBRATION_FIELD_VIBRATION_Z = "Vibration.VibrationZ"
	VIBRATION_FIELD_CLIPPING_0  = "Vibration.Clipping0"
	VIBRATION_FIELD_CLIPPING_1  = "Vibration.Clipping1"
	VIBRATION_FIELD_CLIPPING_2  = "Vibration.Clipping2"
)

// ToMap (generated function)
func (m *Vibration) Dict() map[string]interface{} {
	return map[string]interface{}{
		VIBRATION_FIELD_TIME_USEC:   m.TimeUsec,
		VIBRATION_FIELD_VIBRATION_X: m.VibrationX,
		VIBRATION_FIELD_VIBRATION_Y: m.VibrationY,
		VIBRATION_FIELD_VIBRATION_Z: m.VibrationZ,
		VIBRATION_FIELD_CLIPPING_0:  m.Clipping0,
		VIBRATION_FIELD_CLIPPING_1:  m.Clipping1,
		VIBRATION_FIELD_CLIPPING_2:  m.Clipping2,
	}
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
	Latitude  int32     `json:"Latitude" `  // [ degE7 ] Latitude (WGS84)
	Longitude int32     `json:"Longitude" ` // [ degE7 ] Longitude (WGS84)
	Altitude  int32     `json:"Altitude" `  // [ mm ] Altitude (MSL). Positive for up.
	X         float32   `json:"X" `         // [ m ] Local X position of this position in the local coordinate frame
	Y         float32   `json:"Y" `         // [ m ] Local Y position of this position in the local coordinate frame
	Z         float32   `json:"Z" `         // [ m ] Local Z position of this position in the local coordinate frame
	Q         []float32 `json:"Q" len:"4" ` // World to surface normal and heading transformation of the takeoff position. Used to indicate the heading and slope of the ground
	ApproachX float32   `json:"ApproachX" ` // [ m ] Local X position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachY float32   `json:"ApproachY" ` // [ m ] Local Y position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachZ float32   `json:"ApproachZ" ` // [ m ] Local Z position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
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

const (
	HOME_POSITION_FIELD_LATITUDE   = "HomePosition.Latitude"
	HOME_POSITION_FIELD_LONGITUDE  = "HomePosition.Longitude"
	HOME_POSITION_FIELD_ALTITUDE   = "HomePosition.Altitude"
	HOME_POSITION_FIELD_X          = "HomePosition.X"
	HOME_POSITION_FIELD_Y          = "HomePosition.Y"
	HOME_POSITION_FIELD_Z          = "HomePosition.Z"
	HOME_POSITION_FIELD_Q          = "HomePosition.Q"
	HOME_POSITION_FIELD_APPROACH_X = "HomePosition.ApproachX"
	HOME_POSITION_FIELD_APPROACH_Y = "HomePosition.ApproachY"
	HOME_POSITION_FIELD_APPROACH_Z = "HomePosition.ApproachZ"
)

// ToMap (generated function)
func (m *HomePosition) Dict() map[string]interface{} {
	return map[string]interface{}{
		HOME_POSITION_FIELD_LATITUDE:   m.Latitude,
		HOME_POSITION_FIELD_LONGITUDE:  m.Longitude,
		HOME_POSITION_FIELD_ALTITUDE:   m.Altitude,
		HOME_POSITION_FIELD_X:          m.X,
		HOME_POSITION_FIELD_Y:          m.Y,
		HOME_POSITION_FIELD_Z:          m.Z,
		HOME_POSITION_FIELD_Q:          m.Q,
		HOME_POSITION_FIELD_APPROACH_X: m.ApproachX,
		HOME_POSITION_FIELD_APPROACH_Y: m.ApproachY,
		HOME_POSITION_FIELD_APPROACH_Z: m.ApproachZ,
	}
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
	Latitude     int32     `json:"Latitude" `     // [ degE7 ] Latitude (WGS84)
	Longitude    int32     `json:"Longitude" `    // [ degE7 ] Longitude (WGS84)
	Altitude     int32     `json:"Altitude" `     // [ mm ] Altitude (MSL). Positive for up.
	X            float32   `json:"X" `            // [ m ] Local X position of this position in the local coordinate frame
	Y            float32   `json:"Y" `            // [ m ] Local Y position of this position in the local coordinate frame
	Z            float32   `json:"Z" `            // [ m ] Local Z position of this position in the local coordinate frame
	Q            []float32 `json:"Q" len:"4" `    // World to surface normal and heading transformation of the takeoff position. Used to indicate the heading and slope of the ground
	ApproachX    float32   `json:"ApproachX" `    // [ m ] Local X position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachY    float32   `json:"ApproachY" `    // [ m ] Local Y position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachZ    float32   `json:"ApproachZ" `    // [ m ] Local Z position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	TargetSystem uint8     `json:"TargetSystem" ` // System ID.
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

const (
	SET_HOME_POSITION_FIELD_LATITUDE      = "SetHomePosition.Latitude"
	SET_HOME_POSITION_FIELD_LONGITUDE     = "SetHomePosition.Longitude"
	SET_HOME_POSITION_FIELD_ALTITUDE      = "SetHomePosition.Altitude"
	SET_HOME_POSITION_FIELD_X             = "SetHomePosition.X"
	SET_HOME_POSITION_FIELD_Y             = "SetHomePosition.Y"
	SET_HOME_POSITION_FIELD_Z             = "SetHomePosition.Z"
	SET_HOME_POSITION_FIELD_Q             = "SetHomePosition.Q"
	SET_HOME_POSITION_FIELD_APPROACH_X    = "SetHomePosition.ApproachX"
	SET_HOME_POSITION_FIELD_APPROACH_Y    = "SetHomePosition.ApproachY"
	SET_HOME_POSITION_FIELD_APPROACH_Z    = "SetHomePosition.ApproachZ"
	SET_HOME_POSITION_FIELD_TARGET_SYSTEM = "SetHomePosition.TargetSystem"
)

// ToMap (generated function)
func (m *SetHomePosition) Dict() map[string]interface{} {
	return map[string]interface{}{
		SET_HOME_POSITION_FIELD_LATITUDE:      m.Latitude,
		SET_HOME_POSITION_FIELD_LONGITUDE:     m.Longitude,
		SET_HOME_POSITION_FIELD_ALTITUDE:      m.Altitude,
		SET_HOME_POSITION_FIELD_X:             m.X,
		SET_HOME_POSITION_FIELD_Y:             m.Y,
		SET_HOME_POSITION_FIELD_Z:             m.Z,
		SET_HOME_POSITION_FIELD_Q:             m.Q,
		SET_HOME_POSITION_FIELD_APPROACH_X:    m.ApproachX,
		SET_HOME_POSITION_FIELD_APPROACH_Y:    m.ApproachY,
		SET_HOME_POSITION_FIELD_APPROACH_Z:    m.ApproachZ,
		SET_HOME_POSITION_FIELD_TARGET_SYSTEM: m.TargetSystem,
	}
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
	IntervalUs int32  `json:"IntervalUs" ` // [ us ] The interval between two messages. A value of -1 indicates this stream is disabled, 0 indicates it is not available, &gt; 0 indicates the interval at which it is sent.
	MessageID  uint16 `json:"MessageID" `  // The ID of the requested MAVLink message. v1.0 is limited to 254 messages.
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

const (
	MESSAGE_INTERVAL_FIELD_INTERVAL_US = "MessageInterval.IntervalUs"
	MESSAGE_INTERVAL_FIELD_MESSAGE_ID  = "MessageInterval.MessageID"
)

// ToMap (generated function)
func (m *MessageInterval) Dict() map[string]interface{} {
	return map[string]interface{}{
		MESSAGE_INTERVAL_FIELD_INTERVAL_US: m.IntervalUs,
		MESSAGE_INTERVAL_FIELD_MESSAGE_ID:  m.MessageID,
	}
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
	VtolState   MAV_VTOL_STATE   `json:"VtolState" `   // The VTOL state if applicable. Is set to MAV_VTOL_STATE_UNDEFINED if UAV is not in VTOL configuration.
	LandedState MAV_LANDED_STATE `json:"LandedState" ` // The landed state. Is set to MAV_LANDED_STATE_UNDEFINED if landed state is unknown.
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

const (
	EXTENDED_SYS_STATE_FIELD_VTOL_STATE   = "ExtendedSysState.VtolState"
	EXTENDED_SYS_STATE_FIELD_LANDED_STATE = "ExtendedSysState.LandedState"
)

// ToMap (generated function)
func (m *ExtendedSysState) Dict() map[string]interface{} {
	return map[string]interface{}{
		EXTENDED_SYS_STATE_FIELD_VTOL_STATE:   m.VtolState,
		EXTENDED_SYS_STATE_FIELD_LANDED_STATE: m.LandedState,
	}
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
	IcaoAddress  uint32             `json:"IcaoAddress" `      // ICAO address
	Lat          int32              `json:"Lat" `              // [ degE7 ] Latitude
	Lon          int32              `json:"Lon" `              // [ degE7 ] Longitude
	Altitude     int32              `json:"Altitude" `         // [ mm ] Altitude(ASL)
	Heading      uint16             `json:"Heading" `          // [ cdeg ] Course over ground
	HorVelocity  uint16             `json:"HorVelocity" `      // [ cm/s ] The horizontal velocity
	VerVelocity  int16              `json:"VerVelocity" `      // [ cm/s ] The vertical velocity. Positive is up
	Flags        ADSB_FLAGS         `json:"Flags" `            // Bitmap to indicate various statuses including valid data fields
	Squawk       uint16             `json:"Squawk" `           // Squawk code
	AltitudeType ADSB_ALTITUDE_TYPE `json:"AltitudeType" `     // ADSB altitude type.
	Callsign     string             `json:"Callsign" len:"9" ` // The callsign, 8+null
	EmitterType  ADSB_EMITTER_TYPE  `json:"EmitterType" `      // ADSB emitter type.
	Tslc         uint8              `json:"Tslc" `             // [ s ] Time since last communication in seconds
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
		strings.TrimRight(m.Callsign, string(byte(0))),
		m.EmitterType,
		m.Tslc,
	)
}

const (
	ADSB_VEHICLE_FIELD_ICAO_ADDRESS  = "AdsbVehicle.IcaoAddress"
	ADSB_VEHICLE_FIELD_LAT           = "AdsbVehicle.Lat"
	ADSB_VEHICLE_FIELD_LON           = "AdsbVehicle.Lon"
	ADSB_VEHICLE_FIELD_ALTITUDE      = "AdsbVehicle.Altitude"
	ADSB_VEHICLE_FIELD_HEADING       = "AdsbVehicle.Heading"
	ADSB_VEHICLE_FIELD_HOR_VELOCITY  = "AdsbVehicle.HorVelocity"
	ADSB_VEHICLE_FIELD_VER_VELOCITY  = "AdsbVehicle.VerVelocity"
	ADSB_VEHICLE_FIELD_FLAGS         = "AdsbVehicle.Flags"
	ADSB_VEHICLE_FIELD_SQUAWK        = "AdsbVehicle.Squawk"
	ADSB_VEHICLE_FIELD_ALTITUDE_TYPE = "AdsbVehicle.AltitudeType"
	ADSB_VEHICLE_FIELD_CALLSIGN      = "AdsbVehicle.Callsign"
	ADSB_VEHICLE_FIELD_EMITTER_TYPE  = "AdsbVehicle.EmitterType"
	ADSB_VEHICLE_FIELD_TSLC          = "AdsbVehicle.Tslc"
)

// ToMap (generated function)
func (m *AdsbVehicle) Dict() map[string]interface{} {
	return map[string]interface{}{
		ADSB_VEHICLE_FIELD_ICAO_ADDRESS:  m.IcaoAddress,
		ADSB_VEHICLE_FIELD_LAT:           m.Lat,
		ADSB_VEHICLE_FIELD_LON:           m.Lon,
		ADSB_VEHICLE_FIELD_ALTITUDE:      m.Altitude,
		ADSB_VEHICLE_FIELD_HEADING:       m.Heading,
		ADSB_VEHICLE_FIELD_HOR_VELOCITY:  m.HorVelocity,
		ADSB_VEHICLE_FIELD_VER_VELOCITY:  m.VerVelocity,
		ADSB_VEHICLE_FIELD_FLAGS:         m.Flags,
		ADSB_VEHICLE_FIELD_SQUAWK:        m.Squawk,
		ADSB_VEHICLE_FIELD_ALTITUDE_TYPE: m.AltitudeType,
		ADSB_VEHICLE_FIELD_CALLSIGN:      m.Callsign,
		ADSB_VEHICLE_FIELD_EMITTER_TYPE:  m.EmitterType,
		ADSB_VEHICLE_FIELD_TSLC:          m.Tslc,
	}
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
	ID                     uint32                     `json:"ID" `                     // Unique identifier, domain based on src field
	TimeToMinimumDelta     float32                    `json:"TimeToMinimumDelta" `     // [ s ] Estimated time until collision occurs
	AltitudeMinimumDelta   float32                    `json:"AltitudeMinimumDelta" `   // [ m ] Closest vertical distance between vehicle and object
	HorizontalMinimumDelta float32                    `json:"HorizontalMinimumDelta" ` // [ m ] Closest horizontal distance between vehicle and object
	Src                    MAV_COLLISION_SRC          `json:"Src" `                    // Collision data source
	Action                 MAV_COLLISION_ACTION       `json:"Action" `                 // Action that is being taken to avoid this collision
	ThreatLevel            MAV_COLLISION_THREAT_LEVEL `json:"ThreatLevel" `            // How concerned the aircraft is about this collision
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

const (
	COLLISION_FIELD_ID                       = "Collision.ID"
	COLLISION_FIELD_TIME_TO_MINIMUM_DELTA    = "Collision.TimeToMinimumDelta"
	COLLISION_FIELD_ALTITUDE_MINIMUM_DELTA   = "Collision.AltitudeMinimumDelta"
	COLLISION_FIELD_HORIZONTAL_MINIMUM_DELTA = "Collision.HorizontalMinimumDelta"
	COLLISION_FIELD_SRC                      = "Collision.Src"
	COLLISION_FIELD_ACTION                   = "Collision.Action"
	COLLISION_FIELD_THREAT_LEVEL             = "Collision.ThreatLevel"
)

// ToMap (generated function)
func (m *Collision) Dict() map[string]interface{} {
	return map[string]interface{}{
		COLLISION_FIELD_ID:                       m.ID,
		COLLISION_FIELD_TIME_TO_MINIMUM_DELTA:    m.TimeToMinimumDelta,
		COLLISION_FIELD_ALTITUDE_MINIMUM_DELTA:   m.AltitudeMinimumDelta,
		COLLISION_FIELD_HORIZONTAL_MINIMUM_DELTA: m.HorizontalMinimumDelta,
		COLLISION_FIELD_SRC:                      m.Src,
		COLLISION_FIELD_ACTION:                   m.Action,
		COLLISION_FIELD_THREAT_LEVEL:             m.ThreatLevel,
	}
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
	MessageType     uint16  `json:"MessageType" `       // A code that identifies the software component that understands this message (analogous to USB device classes or mime type strings). If this code is less than 32768, it is considered a 'registered' protocol extension and the corresponding entry should be added to https://github.com/mavlink/mavlink/definition_files/extension_message_ids.xml. Software creators can register blocks of message IDs as needed (useful for GCS specific metadata, etc...). Message_types greater than 32767 are considered local experiments and should not be checked in to any widely distributed codebase.
	TargetNetwork   uint8   `json:"TargetNetwork" `     // Network ID (0 for broadcast)
	TargetSystem    uint8   `json:"TargetSystem" `      // System ID (0 for broadcast)
	TargetComponent uint8   `json:"TargetComponent" `   // Component ID (0 for broadcast)
	Payload         []uint8 `json:"Payload" len:"249" ` // Variable length payload. The length must be encoded in the payload as part of the message_type protocol, e.g. by including the length as payload data, or by terminating the payload data with a non-zero marker. This is required in order to reconstruct zero-terminated payloads that are (or otherwise would be) trimmed by MAVLink 2 empty-byte truncation. The entire content of the payload block is opaque unless you understand the encoding message_type. The particular encoding used can be extension specific and might not always be documented as part of the MAVLink specification.
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

const (
	V2_EXTENSION_FIELD_MESSAGE_TYPE     = "V2Extension.MessageType"
	V2_EXTENSION_FIELD_TARGET_NETWORK   = "V2Extension.TargetNetwork"
	V2_EXTENSION_FIELD_TARGET_SYSTEM    = "V2Extension.TargetSystem"
	V2_EXTENSION_FIELD_TARGET_COMPONENT = "V2Extension.TargetComponent"
	V2_EXTENSION_FIELD_PAYLOAD          = "V2Extension.Payload"
)

// ToMap (generated function)
func (m *V2Extension) Dict() map[string]interface{} {
	return map[string]interface{}{
		V2_EXTENSION_FIELD_MESSAGE_TYPE:     m.MessageType,
		V2_EXTENSION_FIELD_TARGET_NETWORK:   m.TargetNetwork,
		V2_EXTENSION_FIELD_TARGET_SYSTEM:    m.TargetSystem,
		V2_EXTENSION_FIELD_TARGET_COMPONENT: m.TargetComponent,
		V2_EXTENSION_FIELD_PAYLOAD:          m.Payload,
	}
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
	Address uint16 `json:"Address" `        // Starting address of the debug variables
	Ver     uint8  `json:"Ver" `            // Version code of the type variable. 0=unknown, type ignored and assumed int16_t. 1=as below
	Type    uint8  `json:"Type" `           // Type code of the memory variables. for ver = 1: 0=16 x int16_t, 1=16 x uint16_t, 2=16 x Q15, 3=16 x 1Q14
	Value   []int8 `json:"Value" len:"32" ` // Memory contents at specified address
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

const (
	MEMORY_VECT_FIELD_ADDRESS = "MemoryVect.Address"
	MEMORY_VECT_FIELD_VER     = "MemoryVect.Ver"
	MEMORY_VECT_FIELD_TYPE    = "MemoryVect.Type"
	MEMORY_VECT_FIELD_VALUE   = "MemoryVect.Value"
)

// ToMap (generated function)
func (m *MemoryVect) Dict() map[string]interface{} {
	return map[string]interface{}{
		MEMORY_VECT_FIELD_ADDRESS: m.Address,
		MEMORY_VECT_FIELD_VER:     m.Ver,
		MEMORY_VECT_FIELD_TYPE:    m.Type,
		MEMORY_VECT_FIELD_VALUE:   m.Value,
	}
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
	TimeUsec uint64  `json:"TimeUsec" `      // [ us ] Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	X        float32 `json:"X" `             // x
	Y        float32 `json:"Y" `             // y
	Z        float32 `json:"Z" `             // z
	Name     string  `json:"Name" len:"10" ` // Name
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
		strings.TrimRight(m.Name, string(byte(0))),
	)
}

const (
	DEBUG_VECT_FIELD_TIME_USEC = "DebugVect.TimeUsec"
	DEBUG_VECT_FIELD_X         = "DebugVect.X"
	DEBUG_VECT_FIELD_Y         = "DebugVect.Y"
	DEBUG_VECT_FIELD_Z         = "DebugVect.Z"
	DEBUG_VECT_FIELD_NAME      = "DebugVect.Name"
)

// ToMap (generated function)
func (m *DebugVect) Dict() map[string]interface{} {
	return map[string]interface{}{
		DEBUG_VECT_FIELD_TIME_USEC: m.TimeUsec,
		DEBUG_VECT_FIELD_X:         m.X,
		DEBUG_VECT_FIELD_Y:         m.Y,
		DEBUG_VECT_FIELD_Z:         m.Z,
		DEBUG_VECT_FIELD_NAME:      m.Name,
	}
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
	TimeBootMs uint32  `json:"TimeBootMs" `    // [ ms ] Timestamp (time since system boot).
	Value      float32 `json:"Value" `         // Floating point value
	Name       string  `json:"Name" len:"10" ` // Name of the debug variable
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
		strings.TrimRight(m.Name, string(byte(0))),
	)
}

const (
	NAMED_VALUE_FLOAT_FIELD_TIME_BOOT_MS = "NamedValueFloat.TimeBootMs"
	NAMED_VALUE_FLOAT_FIELD_VALUE        = "NamedValueFloat.Value"
	NAMED_VALUE_FLOAT_FIELD_NAME         = "NamedValueFloat.Name"
)

// ToMap (generated function)
func (m *NamedValueFloat) Dict() map[string]interface{} {
	return map[string]interface{}{
		NAMED_VALUE_FLOAT_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		NAMED_VALUE_FLOAT_FIELD_VALUE:        m.Value,
		NAMED_VALUE_FLOAT_FIELD_NAME:         m.Name,
	}
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
	TimeBootMs uint32 `json:"TimeBootMs" `    // [ ms ] Timestamp (time since system boot).
	Value      int32  `json:"Value" `         // Signed integer value
	Name       string `json:"Name" len:"10" ` // Name of the debug variable
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
		strings.TrimRight(m.Name, string(byte(0))),
	)
}

const (
	NAMED_VALUE_INT_FIELD_TIME_BOOT_MS = "NamedValueInt.TimeBootMs"
	NAMED_VALUE_INT_FIELD_VALUE        = "NamedValueInt.Value"
	NAMED_VALUE_INT_FIELD_NAME         = "NamedValueInt.Name"
)

// ToMap (generated function)
func (m *NamedValueInt) Dict() map[string]interface{} {
	return map[string]interface{}{
		NAMED_VALUE_INT_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		NAMED_VALUE_INT_FIELD_VALUE:        m.Value,
		NAMED_VALUE_INT_FIELD_NAME:         m.Name,
	}
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
	Severity MAV_SEVERITY `json:"Severity" `      // Severity of status. Relies on the definitions within RFC-5424.
	Text     string       `json:"Text" len:"50" ` // Status text message, without null termination character
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
		strings.TrimRight(m.Text, string(byte(0))),
	)
}

const (
	STATUSTEXT_FIELD_SEVERITY = "Statustext.Severity"
	STATUSTEXT_FIELD_TEXT     = "Statustext.Text"
)

// ToMap (generated function)
func (m *Statustext) Dict() map[string]interface{} {
	return map[string]interface{}{
		STATUSTEXT_FIELD_SEVERITY: m.Severity,
		STATUSTEXT_FIELD_TEXT:     m.Text,
	}
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
	TimeBootMs uint32  `json:"TimeBootMs" ` // [ ms ] Timestamp (time since system boot).
	Value      float32 `json:"Value" `      // DEBUG value
	Ind        uint8   `json:"Ind" `        // index of debug variable
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

const (
	DEBUG_FIELD_TIME_BOOT_MS = "Debug.TimeBootMs"
	DEBUG_FIELD_VALUE        = "Debug.Value"
	DEBUG_FIELD_IND          = "Debug.Ind"
)

// ToMap (generated function)
func (m *Debug) Dict() map[string]interface{} {
	return map[string]interface{}{
		DEBUG_FIELD_TIME_BOOT_MS: m.TimeBootMs,
		DEBUG_FIELD_VALUE:        m.Value,
		DEBUG_FIELD_IND:          m.Ind,
	}
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
	CustomMode     uint32        `json:"CustomMode" `     // A bitfield for use for autopilot-specific flags
	Type           MAV_TYPE      `json:"Type" `           // Vehicle or component type. For a flight controller component the vehicle type (quadrotor, helicopter, etc.). For other components the component type (e.g. camera, gimbal, etc.). This should be used in preference to component id for identifying the component type.
	Autopilot      MAV_AUTOPILOT `json:"Autopilot" `      // Autopilot type / class. Use MAV_AUTOPILOT_INVALID for components that are not flight controllers.
	BaseMode       MAV_MODE_FLAG `json:"BaseMode" `       // System mode bitmap.
	SystemStatus   MAV_STATE     `json:"SystemStatus" `   // System status flag.
	MavlinkVersion uint8         `json:"MavlinkVersion" ` // MAVLink version, not writable by user, gets added by protocol because of magic data type: uint8_t_mavlink_version
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

const (
	HEARTBEAT_FIELD_CUSTOM_MODE     = "Heartbeat.CustomMode"
	HEARTBEAT_FIELD_TYPE            = "Heartbeat.Type"
	HEARTBEAT_FIELD_AUTOPILOT       = "Heartbeat.Autopilot"
	HEARTBEAT_FIELD_BASE_MODE       = "Heartbeat.BaseMode"
	HEARTBEAT_FIELD_SYSTEM_STATUS   = "Heartbeat.SystemStatus"
	HEARTBEAT_FIELD_MAVLINK_VERSION = "Heartbeat.MavlinkVersion"
)

// ToMap (generated function)
func (m *Heartbeat) Dict() map[string]interface{} {
	return map[string]interface{}{
		HEARTBEAT_FIELD_CUSTOM_MODE:     m.CustomMode,
		HEARTBEAT_FIELD_TYPE:            m.Type,
		HEARTBEAT_FIELD_AUTOPILOT:       m.Autopilot,
		HEARTBEAT_FIELD_BASE_MODE:       m.BaseMode,
		HEARTBEAT_FIELD_SYSTEM_STATUS:   m.SystemStatus,
		HEARTBEAT_FIELD_MAVLINK_VERSION: m.MavlinkVersion,
	}
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
