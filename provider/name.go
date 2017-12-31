package provider

import (
	"math/rand"
	"time"
)

type Generator struct {
	rand *rand.Rand
	Resource NamesResource
}

func (generator *Generator) Name() string {
	return generator.LastName() + generator.FirstName()
}

func (generator *Generator) LastName() string {
	length := len(generator.Resource.LastNames)
	randomIndex := generator.rand.Intn(length)

	return generator.Resource.LastNames[randomIndex]
}

func (generator *Generator) FirstName() string {
	merge := append(generator.Resource.CharacterMale, generator.Resource.CharacterFemale...)
	length := len(merge)
	randomIndex := generator.rand.Intn(length)

	return merge[randomIndex]
}

func Create() Generator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return Generator{
		rand: r,
	}
}
