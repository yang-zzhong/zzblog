package zzblog

import (
	"github.com/go-yaml/yaml"
	"os"
)

var globalConfig Config

// Config ...
type Config struct {
	Port           string   `yaml:"port"`
	Domain         string   `yaml:"domain"`
	Root           string   `yaml:"root"`
	LogPath        string   `yaml:"log_path"`
	DocRoot        string   `yaml:"doc_root"`
	Bots           []string `yaml:"bots"`
	Renderer       string   `yaml:"renderer"`
	RenderCacheDir string   `yaml:"render_cache_dir"`
	Passport       string   `yaml:"passport"`
	AllowCors      bool     `yaml:"allow_cors"`
}

func InitConfig(pathfile string) {
	file, err := os.Open(pathfile)
	if err != nil {
		panic(err)
	}
	if err := yaml.NewDecoder(file).Decode(&globalConfig); err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return &globalConfig
}
