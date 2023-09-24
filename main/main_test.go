package main

import (
	"log"
	"testing"

	"github.com/google/uuid"
)

func Test_UUID(t *testing.T) {
	uuid := uuid.New()
	log.Println(uuid)
}
