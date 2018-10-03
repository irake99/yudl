package main

import (
	"github.com/irake99/yudl/pkg/config"
	"github.com/irake99/yudl/pkg/downloader"
	"github.com/irake99/yudl/pkg/videoinfo"
)

func main() {
	cfg := config.ParseFlags()
	viList, _ := videoinfo.NewInfoList(cfg)
	dlrList := downloader.NewDownloaderList(viList)
	dlrList.Run()
}
