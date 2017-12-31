package provider

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type NamesResource struct {
	LastNames       []string `yaml:"lastNames"`
	CharacterMale   []string `yaml:"characterMale"`
	CharacterFemale []string `yaml:"characterFemale"`
}

var names = []string{
	"金太郎",
	"金城武",
	"金智賢",
}

func ParseFile(file string) (res NamesResource, err error) {
	r, err := ioutil.ReadFile(file)
	if err != nil {
		return res, err
	}

	if err := yaml.Unmarshal(r, &res); err != nil {
		return res, err
	}

	return res, nil
}
