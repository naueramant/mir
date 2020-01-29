package config

import (
	"errors"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	Filename = "screens.yaml"
)

func Load() (*Configuration, error) {
	return readFromFile(Filename)
}

func readFromFile(path string) (*Configuration, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("no " + Filename + " file found")
	}

	t := Configuration{}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return nil, err
	}

	err = Validate(t)

	return &t, err
}

func Exists() bool {
	if _, err := os.Stat(Filename); os.IsNotExist(err) {
		return false
	}
	return true
}
