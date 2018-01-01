package facade

import "github.com/MilesChou/namer/provider"

type GenerateItemProcess func(item string, index int)

func Generate(path string, count int, process GenerateItemProcess) error {
	res, _ := provider.ParseFile(path)

	generator := provider.Create()
	generator.Resource = res

	for i := 0; i < count; i++ {
		process(generator.Name(), i)
	}

	return nil
}
