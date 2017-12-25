package provider

import (
	"math/rand"
	"time"
)

type Generator struct {
	rand *rand.Rand
}

func (generator *Generator) Name() string {
	length := len(names)

	return names[generator.rand.Intn(length)]
}

func Create() Generator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return Generator{
		rand: r,
	}
}
