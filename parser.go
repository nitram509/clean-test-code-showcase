package parser

import (
	"regexp"
	"strconv"
	"strings"
)

type FirmwareInformation struct {
	FirmwareName   string
	TargetName     string
	TargetDetail   string
	Version        int64
	ReleaseDateStr string
	GitHash        string
	ReleaseTime    string
}

func ParseFirmwareInformation(text string) *FirmwareInformation {
	lines := strings.Split(text, "\n")
	regex, _ := createSearchExpression()
	for _, line := range lines {
		information := parseLine(regex, line)
		if information != nil {
			return information
		}
	}
	return nil
}

func createSearchExpression() (*regexp.Regexp, error) {
	return regexp.Compile("^#\\s" +
		"(\\w+)?\\s/\\s" +
		"(\\w+)?\\s" +
		"\\((\\w+)\\)\\s" +
		"(\\d+)\\.\\d+\\.\\d+\\s" +
		"(\\w+\\s\\d+\\s\\d{4})?\\s/\\s" +
		"(\\d+:\\d+:\\d+)?\\s" +
		"\\((\\w+)\\)?")
}

func parseLine(regex *regexp.Regexp, text string) *FirmwareInformation {
	submatch := regex.FindSubmatch([]byte(text))
	if len(submatch) >= 2 {
		info := FirmwareInformation{}
		info.FirmwareName = string(submatch[1])
		if len(submatch) > 2 {
			info.TargetName = string(submatch[2])
		}
		if len(submatch) > 3 {
			info.TargetDetail = string(submatch[3])
		}
		if len(submatch) > 4 {
			info.Version, _ = strconv.ParseInt(string(submatch[4]), 10, 64)
		}
		if len(submatch) > 5 {
			info.ReleaseDateStr = string(submatch[5])
		}
		if len(submatch) > 6 {
			info.ReleaseTime = string(submatch[6])
		}
		if len(submatch) > 7 {
			info.GitHash = string(submatch[7])
		}
		return &info
	}
	return nil
}
