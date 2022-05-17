package main

import "fmt"

func main() {
	// 定义一个map
	m := map[string]string{
		"name":    "Asura",
		"course":  "Golang",
		"site":    "https://imooc.com",
		"quality": "not bad",
	}
	fmt.Println(m)

	// 定义一个空map
	m2 := make(map[string]string) // m2 == empty map
	fmt.Println(m2)

	// 使用var定义一个空map
	var m3 map[string]string // m3 == nil
	fmt.Println(m3)

	fmt.Println("------------- Traversing map -------------")
	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	fmt.Println("------------- Getting values -------------")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if coursName, ok := m["cours"]; ok {
		fmt.Println(coursName)
	} else {
		fmt.Println("key doesn't exist")
	}

	fmt.Println("------------- Deleting values -------------")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
