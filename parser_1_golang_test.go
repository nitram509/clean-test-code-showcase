package parser

import (
	"testing"
)

// See https://golang.org/pkg/testing/
// The quasi standard ;)

func Test_betaflight_message_can_be_parsed__with_GOLANG(t *testing.T) {
	info := ParseFirmwareInformation("# Betaflight / SPRACINGF3EVO (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39")

	if info.FirmwareName != "Betaflight" {
		t.Errorf("FirmwareName = %s; want Betaflight", info.FirmwareName)
	}
	if info.TargetName != "SPRACINGF3EVO" {
		t.Errorf("TargetName = %s; want SPRACINGF3EVO", info.TargetName)
	}
	if info.TargetDetail != "SPEV" {
		t.Errorf("TargetDetail = %s; want SPEV", info.TargetDetail)
	}
	if !(info.Version >= int64(3)) {
		t.Errorf("Version = %d; want less than 3", info.Version)
	}
	if info.ReleaseDateStr != "Apr 17 2018" {
		t.Errorf("ReleaseDateStr = %s; want Apr 17 2018", info.ReleaseDateStr)
	}
	if info.ReleaseTime != "14:00:13" {
		t.Errorf("ReleaseTime = %s; want 14:00:13", info.ReleaseTime)
	}
	if info.GitHash != "b2c247d34" {
		t.Errorf("GitHash = %s; want b2c247d34", info.GitHash)
	}
}

func Test_xparsing_multiple_lines_returns_first_hit__with_GOLANG(t *testing.T) {
	info := ParseFirmwareInformation(
		"# bla blubber\n" +
			"# Betaflight1 / SPRACINGF3EVO1 (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39\n" +
			"# Betaflight2 / SPRACINGF3EVO2 (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39",
	)

	if info.FirmwareName != "Betaflight1" {
		t.Errorf("FirmwareName = %s; want Betaflight1", info.FirmwareName)
	}
	if info.TargetName != "SPRACINGF3EVO1" {
		t.Errorf("TargetName = %s; want SPRACINGF3EVO1", info.TargetName)
	}
}
