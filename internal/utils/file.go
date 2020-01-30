package utils

import "io/ioutil"

func ReadFileToString(path string) (string, error) {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}
