package parser

import (
	"github.com/rdrdr/hamcrest/asserter"
	"github.com/rdrdr/hamcrest/core"
	"testing"
)

// See https://github.com/corbym/gocrest
// (exactly 27 Stars on Github)

// BROKEN ?!?! Last Release from 2011?

func Test_betaflight_message_can_be_parsed__with_HAMCREST(t *testing.T) {
	info := ParseFirmwareInformation("# Betaflight / SPRACINGF3EVO (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39")

	we := asserter.Using(t)
	we.AssertThat(info.FirmwareName, core.EqualTo("Betaflight").Comment("FirmwareName field"))
	we.AssertThat(info.TargetName, core.EqualTo("SPRACINGF3EVO"))
	we.AssertThat(info.TargetDetail, core.EqualTo("SPEV"))
	we.AssertThat(info.Version, core.GreaterThanOrEqualTo(int64(3)))
	we.AssertThat(info.ReleaseDateStr, core.EqualTo("Apr 17 2018"))
	we.AssertThat(info.ReleaseTime, core.EqualTo("14:00:13"))
	we.AssertThat(info.GitHash, core.EqualTo("b2c247d34"))
}

func Test_parsing_multiple_lines_returns_first_hit__with_HAMCREST(t *testing.T) {
	info := ParseFirmwareInformation(
		"# bla blubber\n" +
			"# Betaflight1 / SPRACINGF3EVO1 (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39\n" +
			"# Betaflight2 / SPRACINGF3EVO2 (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39",
	)

	we := asserter.Using(t)
	we.AssertThat(info.FirmwareName, core.EqualTo("Betaflight1").Comment("FirmwareName"))
	we.AssertThat(info.TargetName, core.EqualTo("SPRACINGF3EVO1"))
}
