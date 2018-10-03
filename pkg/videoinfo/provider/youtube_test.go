package provider

import (
	"fmt"
	"testing"
)

var testCases = []string{
	"https://www.youtube.com/watch?v=wwDa2_pUxyU",
	"https://www.youtube.com/embed/YEkeXsmgtz0",
	"https://youtu.be/QFOivmZrnxM",
	"https://youtu.be/yb76aHz9YPs?t=6m43s",
	"nDiFzmGFjS8",
}

var youtube = NewProvider()

func TestParseVideoID(t *testing.T) {
	for _, tc := range testCases {
		youtube.ParseVideoID(tc)
	}
}

func TestGetInfoURL(t *testing.T) {
	for _, tc := range testCases {
		videoID, _ := youtube.ParseVideoID(tc)
		fmt.Println(youtube.GetInfoURL(videoID))
	}
}
