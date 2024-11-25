package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"delly.ioo/durian/cache/cache"
)

type Cache interface {
	Put(key string, values map[string]string) error
	Get(key string) (string, bool)
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
		attrs := map[string]string{}
		for i := 2; i < len(args); i += 2 {
			k, v := args[i], args[i+1]
			attrs[k] = v
		}
		if err := c.Put(key, attrs); err != nil {
			fmt.Println(err.Error())
		}
	case "delete":
		key := args[1]
		c.Delete(key)
	case "get":
		key := args[1]
		values, found := c.Get(key)
		if found {
			fmt.Println(values)
		} else {
			fmt.Printf("No entry found for %s\n", key)
		}
	case "search":
		attrKey, attrVal := args[1], args[2]
		fmt.Println(c.Search(attrKey, attrVal))
	case "keys":
		fmt.Println(c.Keys())
	}
}
