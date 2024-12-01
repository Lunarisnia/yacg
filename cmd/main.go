package main

import (
	"log"

	"github.com/lunarisnia/yacg"
	"github.com/lunarisnia/yacg/color"
)

func main() {
	yacg.PathTrace(10, 200, func(x int, y int, c *color.RGB) {
		log.Println(c, "a-sdasd0iad", x, y)
	})
}
