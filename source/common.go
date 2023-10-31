package downloader

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var config map[string]interface{}

func init() {
	config = make(map[string]interface{})
	// read config file
	projectDir := os.Getenv("MY_IEEE_ROOT")
	if projectDir == "" {
		projectDir = "../"
	}

	configFile, err := os.Open(path.Join(projectDir, "/config/downloader.yaml"))
	if err != nil {
		panic(err)
	}

	// parse config file
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)

}
