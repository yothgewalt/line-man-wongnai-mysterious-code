# LINE MAN Wongnai Mysterious Code

# ![image](https://user-images.githubusercontent.com/108649272/178962304-75130546-d912-41fe-95e2-44a1623f7c83.png)

## Quiz Code (Mysterious Code)
```go
var whatIsIt string
secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
sd, _ := base64.StdEncoding.DecodeString(secret)
```

## My Answer Code (Mysterious Code)
```go
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
```

# <b>Final Answer: "Join:us:at:LINE:MAN:Wongnai"</b>