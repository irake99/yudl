package config

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/irake99/yudl/util/log"
)

// Config is the Configuration Data for whole program
type Config struct {
	urls   []string
	outdir string
}

var cfg Config

func init() {
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	flag.StringVar(&cfg.outdir, "outdir", "",
		"Directory where the video files save to. (Default: "+currentDir+")")
	flag.Parse()

	if flag.NArg() <= 0 {
		log.Error.Println("no video ID or URL(s) given")
	}

	cfg.urls = flag.Args()
}

func (cfg *Config) String() string {
	return ""
}

// ParseFlags parses flags and arguments and save to config
func ParseFlags() *Config {
	return &cfg
}

// GetURLs returns video URLs
func (cfg *Config) GetURLs() []string {
	return cfg.urls
}

// GetOutDir returns the directory to save video files
func (cfg *Config) GetOutDir() string {
	return cfg.outdir
}
