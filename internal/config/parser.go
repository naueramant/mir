package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

func Load(filename string) (*Configuration, error) {
	c := Configuration{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return &c, errors.Wrap(err, "Failed to load configuration file")
	}

	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		return &c, errors.Wrap(err, "Failed to unmarshal configuration file")
	}

	err = Validate(c)

	return &c, errors.Wrap(err, "Configuration file invalid")
}
