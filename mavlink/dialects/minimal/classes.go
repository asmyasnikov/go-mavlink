/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package minimal

import (
	"encoding/binary"
	"fmt"
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
)

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
