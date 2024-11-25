package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"delly.ioo/durian/cache/cache"
)

type Cache interface {
	Put(key string, attributeKeys, attributeValues []string) error
	Get(key string) (string, error)
	Search(attributeKey, attributeValue string) string
	Delete(key string)
	Keys() string
}

func main() {
	var c Cache = cache.NewCache()
	fmt.Println("insert command:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")
		if len(args) > 0 {
			if args[0] == "exit" {
				break
			}
			processArgs(args, c)
		}
	}
}

func processArgs(args []string, c Cache) {
	switch args[0] {
	case "put":
		key := args[1]
		attributeKeys, attributeValues := []string{}, []string{}
		for i := 2; i < len(args); i += 2 {
			k, v := args[i], args[i+1]
			attributeKeys = append(attributeKeys, k)
			attributeValues = append(attributeValues, v)
		}
		if err := c.Put(key, attributeKeys, attributeValues); err != nil {
			fmt.Println(err.Error())
		}
	case "delete":
		key := args[1]
		c.Delete(key)
	case "get":
		key := args[1]
		values, err := c.Get(key)
		if err == nil {
			fmt.Println(values)
		} else {
			fmt.Println(err)
		}
	case "search":
		attrKey, attrVal := args[1], args[2]
		fmt.Println(c.Search(attrKey, attrVal))
	case "keys":
		fmt.Println(c.Keys())
	}
}
