package provider

import (
	"math/rand"
	"time"
)

type Generator struct {
	rand     *rand.Rand
	Resource NamesResource
}

func (generator *Generator) Name() string {
	return generator.LastName() + generator.FirstName()
}

func (generator *Generator) NameSingle() string {
	return generator.LastName() + generator.FirstNameSingle()
}

func (generator *Generator) NameDouble() string {
	return generator.LastName() + generator.FirstNameDouble()
}

func (generator *Generator) LastName() string {
	return generator.pickCharacter(generator.Resource.LastNames, 1)
}

func (generator *Generator) FirstName() string {
	return generator.firstName(generator.rand.Intn(2) + 1)
}

func (generator *Generator) FirstNameSingle() string {
	return generator.firstName(1)
}

func (generator *Generator) FirstNameDouble() string {
	return generator.firstName(2)
}

func (generator *Generator) firstName(count int) string {
	merge := append(generator.Resource.CharacterMale, generator.Resource.CharacterFemale...)

	return generator.pickCharacter(merge, count)
}

func (generator *Generator) pickCharacter(collection []string, count int) (chars string) {
	length := len(collection)

	for i := 0; i < count; i++ {
		randomIndex := generator.rand.Intn(length)
		chars = chars + collection[randomIndex]
	}

	return chars
}

func Create() Generator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return Generator{
		rand: r,
	}
}
