package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "http://localhost:8000/users/dsfsdfsdfsdfsdf4343123:2342"
	arr := strings.Split(str, "/")
	fmt.Println(arr[len(arr)-1])
	// 2024-01-24 18:02:07.000
}
