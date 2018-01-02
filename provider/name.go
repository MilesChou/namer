package provider

import (
	"math/rand"
	"time"
	"errors"
	"fmt"
)

type Generator struct {
	rand     *rand.Rand
	Resource NamesResource
}

func (generator *Generator) Name(gender string, firstNameNum int) (name string, err error) {
	switch gender {
	case "":
		name = generator.LastName() + generator.firstName(firstNameNum)
	case "male":
		name = generator.LastName() + generator.firstNameMale(firstNameNum)
	case "female":
		name = generator.LastName() + generator.firstNameFemale(firstNameNum)
	default:
		err = errors.New(fmt.Sprintf("gender '%s' is invalid", gender))
	}

	return
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
	return generator.firstName(0)
}

func (generator *Generator) FirstNameMale() string {
	return generator.firstNameMale(0)
}

func (generator *Generator) FirstNameFemale() string {
	return generator.firstNameFemale(0)
}

func (generator *Generator) FirstNameSingle() string {
	return generator.firstName(1)
}

func (generator *Generator) FirstNameMaleSingle() string {
	return generator.firstNameMale(1)
}

func (generator *Generator) FirstNameFemaleSingle() string {
	return generator.firstNameFemale(1)
}

func (generator *Generator) FirstNameDouble() string {
	return generator.firstName(2)
}

func (generator *Generator) FirstNameMaleDouble() string {
	return generator.firstNameMale(2)
}

func (generator *Generator) FirstNameFemaleDouble() string {
	return generator.firstNameFemale(2)
}

func (generator *Generator) firstName(count int) string {
	if count == 0 {
		count = generator.rand.Intn(2) + 1
	}

	merge := append(generator.Resource.CharacterMale, generator.Resource.CharacterFemale...)

	return generator.pickCharacter(merge, count)
}

func (generator *Generator) firstNameMale(count int) string {
	if count == 0 {
		count = generator.rand.Intn(2) + 1
	}

	return generator.pickCharacter(generator.Resource.CharacterMale, count)
}

func (generator *Generator) firstNameFemale(count int) string {
	if count == 0 {
		count = generator.rand.Intn(2) + 1
	}

	return generator.pickCharacter(generator.Resource.CharacterFemale, count)
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
