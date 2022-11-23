package main

import "errors"

type Dictionary map[string]string

// func Search(dict map[string]string, value string) string {
// 	return dict[value]
// }

// Empty map
// dictionary = map[string]string{}
// dictionary = make(map[string]string)

var WordNotFoundErr error = errors.New("The word non exist")

func (d Dictionary) Search(value string) (string, error) {
	word, exist := d[value]
	if !exist {
		return "", WordNotFoundErr
	}

	return word, nil
}

func (d Dictionary) Add(word, value string) {
	d[word] = value
}
