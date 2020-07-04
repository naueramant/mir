package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

func Load(filename string) (Configuration, error) {
	t := Configuration{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return t, errors.Wrap(err, "Failed to load configuration file")
	}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return t, errors.Wrap(err, "Failed to unmarshal configuration file")
	}

	err = Validate(t)

	return t, errors.Wrap(err, "Configuration file invalid")
}
