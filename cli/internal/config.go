package internal

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/harryvince/s/cli/constants"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Prefix  string `yaml:"prefix"`
	Profile string `yaml:"profile"`
	Region  string `yaml:"region"`
}

func NewConfig(prefix string, profile string, region string, env string) (*Config, error) {
	// First check if file exists
	if _, err := os.Stat(constants.CONFIG_NAME); err == nil {
		fmt.Println("Config already exists, failure to create.")
		return nil, err
	}

	// Create config file
	file, err := os.Create(constants.CONFIG_NAME)
	if err != nil {
		fmt.Println("Error creating config:", err)
		return nil, err
	}

	config := Config{Prefix: prefix, Profile: profile, Region: region}

	data, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Println("Error preparing config:", err)
		return nil, err
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing config:", err)
		return nil, err
	}
	file.Close()

	ssm_client := NewSSMClient()
	params, err := ssm_client.ListSecretNames(env)

	if err != nil {
		fmt.Println("Failed to list secret names during init:", err)
		return nil, err
	}

	if len(params) > 0 {
		err := os.Remove(constants.CONFIG_NAME)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		return nil, errors.New("Prefix already in use")
	}

	fmt.Println("Config initialized.")
	return &config, nil
}

func GetConfig() *Config {
	// First check if file exists
	if _, err := os.Stat(constants.CONFIG_NAME); err != nil {
		fmt.Println("Configuration does not exist, ensure to run `s init <prefix>`.")
		os.Exit(1)
	}

	// Read config file
	file, err := os.Open(constants.CONFIG_NAME)
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	return &config
}
