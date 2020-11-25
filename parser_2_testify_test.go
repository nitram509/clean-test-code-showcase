package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_betaflight_message_can_be_parsed__with_TESTIFY(t *testing.T) {
	info := ParseFirmwareInformation("# Betaflight / SPRACINGF3EVO (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39")

	assert.Equal(t, info.FirmwareName, "Betaflight")
	assert.Equal(t, info.TargetName, "SPRACINGF3EVO")
	assert.Contains(t, info.TargetName, "RACING")
	assert.GreaterOrEqualf(t, info.Version, int64(3), "error message %s")
	assert.Equal(t, info.ReleaseDateStr, "Apr 17 2018")
	assert.Equal(t, info.ReleaseTime, "14:00:13")
	assert.Equal(t, info.GitHash, "b2c247d34")
}

func Test_parsing_multiple_lines_returns_first_hit__with_TESTIFY(t *testing.T) {
	info := ParseFirmwareInformation(
		"# bla blubber\n" +
			"# Betaflight1 / SPRACINGF3EVO1 (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39\n" +
			"# Betaflight2 / SPRACINGF3EVO2 (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39",
	)

	assert.Equal(t, "Betaflight1", info.FirmwareName)
	assert.Equal(t, "SPRACINGF3EVO1", info.TargetName)
}
