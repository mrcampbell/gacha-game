package util

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateID() string {
	id, err := gonanoid.New(8)
	if err != nil {
		log.Fatalf("failed to generate id | %s", err)
	}
	return id
}
