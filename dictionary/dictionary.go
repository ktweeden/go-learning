package main

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definiton, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definiton, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
