package main

import (
	"fmt"
	"github.com/alexedwards/argon2id"
	"log"
)

func main() {
	p := &argon2id.Params{
		Memory:      128 * 1024,
		Iterations:  4,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := argon2id.CreateHash("helloworld", p)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("hash", hash)

	match, err := argon2id.ComparePasswordAndHash("helloworldd", hash)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Match", match)
}
