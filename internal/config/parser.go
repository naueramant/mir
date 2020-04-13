package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

const (
	Filename = "screen.yaml"
)

func Load() (Configuration, error) {
	t := Configuration{}

	data, err := ioutil.ReadFile(Filename)
	if err != nil {
		return t, errors.Wrap(err, "No configuration file found")
	}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return t, errors.Wrap(err, "Failed to unmarshal configuration file")
	}

	err = Validate(t)

	return t, errors.Wrap(err, "Configuration file invalid")
}
