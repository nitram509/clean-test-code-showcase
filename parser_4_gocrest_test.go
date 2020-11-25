package parser

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

// See https://github.com/corbym/gocrest
// (exactly 11 Stars on Github)

func Test_betaflight_message_can_be_parsed__with_GOCREST(t *testing.T) {
	info := ParseFirmwareInformation("# Betaflight / SPRACINGF3EVO (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39")

	then.AssertThat(t, info.FirmwareName, is.EqualTo("Betaflight"))
	then.AssertThat(t, info.TargetName, is.EqualTo("SPRACINGF3EVO"))
	then.AssertThat(t, info.TargetDetail, is.EqualTo("SPEV"))
	then.AssertThat(t, info.Version, is.GreaterThanOrEqualTo(int64(3)))
	then.AssertThat(t, info.ReleaseDateStr, is.EqualTo("Apr 17 2018"))
	then.AssertThat(t, info.ReleaseTime, is.EqualTo("14:00:13"))
	then.AssertThat(t, info.GitHash, is.EqualTo("b2c247d34"))
}

func Test_parsing_multiple_lines_returns_first_hit__with_GOCREST(t *testing.T) {
	info := ParseFirmwareInformation(
		"# bla blubber\n" +
			"# Betaflight1 / SPRACINGF3EVO1 (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39\n" +
			"# Betaflight2 / SPRACINGF3EVO2 (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39",
	)

	then.AssertThat(t, info.FirmwareName, is.EqualTo("Betaflight1"))
	then.AssertThat(t, info.TargetName, is.EqualTo("SPRACINGF3EVO1"))
}
