package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	length := len(sd)
	runes := make([]rune, length)
	for _, rune := range string(sd) {
		length--
		runes[length] = rune
	}

	whatIsIt = string(runes[length:])

	fmt.Println(whatIsIt)
}
