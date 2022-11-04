package main

type Dictionary map[string]string

// func Search(dict map[string]string, value string) string {
// 	return dict[value]
// }

func (d Dictionary) Search(value string) string {
	return d[value]
}
