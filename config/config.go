package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-yaml/yaml"
	"github.com/mtricolici/file-hub-go/helpers"
)

const (
	CONFIG_PATH_ENV_VAR = "CONFIG_PATH"
	DEFAULT_CONFIG_PATH = "config.yml"
	DEFAULT_INTERFACE   = "127.0.0.1"
	DEFAULT_PORT        = 8080
)

type YamlConfigFile struct {
	Listen struct {
		Interface string `yaml:"interface"`
		Port      int    `yaml:"port"`
	} `yaml:"listen"`
}

type Config struct {
	yaml *YamlConfigFile
}

var (
	cfg_instance *Config
	cfg_once     sync.Once
)

func Get() *Config {
	cfg_once.Do(func() {
		cfg_instance = &Config{
			yaml: &YamlConfigFile{},
		}
		cfg_instance.loadConfig()
	})
	return cfg_instance
}

func (c *Config) loadConfig() {
	var filePath string
	if path, ok := os.LookupEnv(CONFIG_PATH_ENV_VAR); ok {
		filePath = helpers.ExpandHomePath(path)
	} else {
		filePath = DEFAULT_CONFIG_PATH
	}

	fmt.Printf("Loading config '%s'...\n", filePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("error reading config file: %s", err))
	}

	if err := yaml.Unmarshal(data, c.yaml); err != nil {
		panic(fmt.Errorf("error unmarshalling config file: %s", err))
	}
}

func (c *Config) ListenInterface() string {
	v := c.yaml.Listen.Interface
	if len(v) == 0 {
		return DEFAULT_INTERFACE
	}

	return v
}

func (c *Config) ListenPort() int {
	if c.yaml.Listen.Port == 0 {
		return DEFAULT_PORT
	}
	return c.yaml.Listen.Port
}
