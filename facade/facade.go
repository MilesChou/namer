package facade

import (
	"github.com/MilesChou/namer/provider"
)

var (
	GenerateFirstNameNum int
	GenerateGender string
	GenerateFirstNameOnly bool
)

type GenerateItemProcess func(item string, index int)

type QueryItemProcess func(item string)

func init() {
	reset()
}

func reset() {
	GenerateFirstNameNum = 2
	GenerateGender = ""
	GenerateFirstNameOnly = false
}

func Generate(path string, count int, process GenerateItemProcess) (err error) {
	res, _ := provider.ParseFile(path)

	generator := provider.Create()
	generator.Resource = res

	for i := 0; i < count; i++ {
		var name string

		if GenerateFirstNameOnly {
			name, err = generator.NameFirstNameOnly(GenerateGender, GenerateFirstNameNum)
		} else  {
			name, err = generator.Name(GenerateGender, GenerateFirstNameNum)
		}

		if err != nil {
			return err
		}

		process(name, i)
	}

	return nil
}

func Query(str string, process QueryItemProcess) error {
	dict, err := provider.Query(str)

	if err != nil {
		return err
	}

	for _, heteronym := range dict.Heteronyms {
		for _, definition := range heteronym.Definitions {
			process(definition)
		}
	}

	return nil
}