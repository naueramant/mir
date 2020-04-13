package config

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

const (
	Filename = "screen.yaml"
)

func Load() (Configuration, error) {
	t := Configuration{}

	data, err := ioutil.ReadFile(Filename)
	if err != nil {
		return t, errors.New("no " + Filename + " file found")
	}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return t, err
	}

	err = Validate(t)

	return t, err
}
