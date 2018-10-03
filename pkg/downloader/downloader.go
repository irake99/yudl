package downloader

import "github.com/irake99/yudl/pkg/videoinfo"

// Downloader downloads a video
type Downloader struct {
	url    string
	outdir string
}

// List is a list of Downloader
type List []Downloader

// NewDownloaderList returns a list of video downloader
func NewDownloaderList(info videoinfo.InfoList) List {
	return List{}
}

// NewDownloader returns a video downloader
func NewDownloader(info *videoinfo.Info) Downloader {
	return buildDownloader(info)
}

func buildDownloader(info *videoinfo.Info) Downloader {
	dlr := Downloader{}

	return dlr
}

// Run the downloader to start download the video
func (dlr *Downloader) Run() {

}

// Run all downloader in the list to start download all video
func (dlrList *List) Run() {

}
