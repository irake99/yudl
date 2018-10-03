package provider

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/irake99/yudl/util/log"
)

// Youtube defines a video provider
type Youtube struct {
	// Regexp list for parsing URLs
	urlReList []*regexp.Regexp
	// URLs for retrieving video infomation
	infoURL string
}

func init() {
	y := NewProvider()
	ProviderList = append(ProviderList, y)
}

// NewProvider returns a new provider
func NewProvider() *Youtube {
	return &Youtube{
		urlReList: []*regexp.Regexp{
			regexp.MustCompile(`(?:v|embed|watch\?v)(?:=|/)([^"&?/=%]{11})`),
			regexp.MustCompile(`(?:=|/)([^"&?/=%]{11})`),
			regexp.MustCompile(`([^"&?/=%]{11})`),
		},
		infoURL: "https://youtube.com/get_video_info?video_id=",
	}
}

// ParseVideoID parses URLs and try to get Video ID
func (y *Youtube) ParseVideoID(url string) (string, error) {
	videoID := ""
	isMatch := false
	for _, re := range y.urlReList {
		if isMatch = re.MatchString(url); isMatch {
			subs := re.FindStringSubmatch(url)
			videoID = subs[1]
			break
		}
	}
	if !isMatch {
		return "", errors.New("no such video ID in: " + url)
	}

	log.Info.Printf("Found video ID: '%s'", videoID)
	if strings.ContainsAny(videoID, "?&/<%=") {
		return videoID, errors.New("invalid characters in video ID")
	}
	if len(videoID) < 10 {
		return videoID, errors.New("the video ID must be at least 10 characters long")
	}
	return videoID, nil
}

// GetInfoURL returns video infoURLs of video provider
func (y *Youtube) GetInfoURL(videoID string) string {
	return fmt.Sprintln(y.infoURL + videoID)
}

// DecodeDownloadURL decodes the video download URL from the video URL
func (y *Youtube) DecodeDownloadURL(url string) string {
	return ""
}
