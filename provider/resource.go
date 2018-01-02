package provider

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)

type NamesResource struct {
	LastNames       []string `yaml:"lastNames"`
	CharacterMale   []string `yaml:"characterMale"`
	CharacterFemale []string `yaml:"characterFemale"`
}

const DefaultTemplate = `
lastNames:
- 李
- 王
- 張
- 劉
- 陳
- 楊
- 趙
- 黃

characterMale:
- 家
- 豪
- 志
- 明

characterFemale:
- 雅
- 婷
- 春
- 嬌
`

func ParseFile(file string) (res NamesResource, err error) {
	_, err = os.Stat(file)
	if err != nil {
		return res, err
	}

	r, err := ioutil.ReadFile(file)
	if err != nil {
		return res, err
	}

	if err := yaml.Unmarshal(r, &res); err != nil {
		return res, err
	}

	return res, nil
}

func InitFileDefault(filename string) error {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return ioutil.WriteFile(filename, []byte(DefaultTemplate), 0644)
	}

	return nil
}
