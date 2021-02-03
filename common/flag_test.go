package common

import (
	"github.com/asmyasnikov/go-mavlink/mavlink"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMavlinkVersion(t *testing.T) {
	require.Equal(t, mavlink.MAVLINK_V1, mavlinkVersion(true, false))
	require.Equal(t, mavlink.MAVLINK_V2, mavlinkVersion(false, true))
	require.Equal(t, mavlink.MAVLINK_V1|mavlink.MAVLINK_V2, mavlinkVersion(true, true))
	require.Equal(t, mavlink.MAVLINK_V1|mavlink.MAVLINK_V2, mavlinkVersion(false, false))
}
