package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FilePath string
}

func NewStorage[T any](filePath string) *Storage[T] {
	return &Storage[T]{FilePath: filePath}
}

func (s *Storage[T]) Load(data *T) error {
	file, err := os.Open(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No file exists, nothing to load
		}
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(data)
}

func (s *Storage[T]) Save(data T) error {
	file, err := os.Create(s.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(data)
}
