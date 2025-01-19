package main

import (
	"errors"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func Add(a, b int) int {
	return a + b
}

func GenerateID(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("invalid length: must be greater than 0")
	}

	id, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", length)
	if err != nil {
		return "", fmt.Errorf("failed to generate ID: %w", err)
	}

	return id, nil
}
