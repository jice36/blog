package server

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigBlog struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}
type ConfigDB struct {
	Database struct {
		Dbhost     string `yaml:"dbhost"`
		Dbport     string `yaml:"dbport"`
		Dbname     string `yaml:"dbname"`
		Dbuser     string `yaml:"dbuser"`
		Dbpassword string `yaml:"dbpassword"`
	} `yaml:"database"`
}

func NewConfigBlog(configPath string) (*ConfigBlog, error) {
	config := &ConfigBlog{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func NewConfigDB(configPath string) (*ConfigDB, error) {
	config := &ConfigDB{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func ParseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}
	return configPath, nil
}
