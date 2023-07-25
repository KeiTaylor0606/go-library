package toml

import "github.com/BurntSushi/toml"

type Config struct{}

func ReadToml(path string) *Config {
	var config *Config
	if _, err := toml.DecodeFile(path, config); err != nil {
		panic(err)
	}
	return config
}
