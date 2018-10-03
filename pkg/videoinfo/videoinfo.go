package videoinfo

import (
	"errors"

	"github.com/irake99/yudl/pkg/config"
	"github.com/irake99/yudl/pkg/videoinfo/provider"
)

// Info collects all information of a video
type Info struct {
	videoID  string
	provider provider.Provider
}

// InfoList is a list of VideoInfo
type InfoList []Info

// NewInfo creates a new video info by parsing an URL
func NewInfo(url string) (*Info, error) {
	info := new(Info)
	var videoID string
	var err error
	for _, p := range provider.ProviderList {
		videoID, err = p.ParseVideoID(url)
		if err == nil {
			info.videoID = videoID
			info.provider = p
			break
		}
	}
	if err != nil {
		return nil, err
	}

	info.DecodeDownloadURL()

	return info, nil
}

// NewInfoList creates a new list of video info
func NewInfoList(cfg *config.Config) (InfoList, error) {
	il, err := parseConfig(cfg)
	if err != nil {
		return nil, err
	}

	return il, nil
}

func parseConfig(cfg *config.Config) (InfoList, error) {
	il := InfoList{}
	urls := cfg.GetURLs()
	for _, url := range urls {
		info, err := NewInfo(url)
		if err != nil {
			il = append(il, *info)
		}
	}

	if len(il) == 0 {
		return nil, errors.New("no such any available URL")
	}

	return il, nil
}

// DecodeDownloadURL decodes the video download URL from video infomation
func (info *Info) DecodeDownloadURL() {

}

func (info *Info) String() string {
	return ""
}

func (il *InfoList) String() string {
	return ""
}
