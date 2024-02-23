// Package config_util loads the configuration for the application

package utils

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func readYaml(filename string, v interface{}) error {
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// read the file
	err = yaml.NewDecoder(file).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func displayConfig(config map[string]interface{}) {
	// iterate over the config
	for key, value := range config {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func GetDeployConfig(configFile string) map[string]interface{} {
	// read the yaml configuration file
	// check if the file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// file does not exist
		errorStr := fmt.Sprintf("config file does not exist: %s", configFile)
		panic(errorStr)
	}

	configFile = "config/deploy_config.yaml"

	// read the file
	config := make(map[string]interface{})
	// read the file
	err := readYaml(configFile, &config)
	if err != nil {
		errorStr := fmt.Sprintf("failed to read config file: %s", err)
		panic(errorStr)
	}

	// display the config
	log.Default().Println("Loaded config:")
	displayConfig(config)

	return config
}
