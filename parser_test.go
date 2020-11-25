package parser

import (
	"reflect"
	"testing"
)

// AUTO GENERATED code by IntelliJ
// TODO implement ;)

func TestParseFirmwareInformation(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want *FirmwareInformation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFirmwareInformation(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFirmwareInformation() = %v, want %v", got, tt.want)
			}
		})
	}
}
