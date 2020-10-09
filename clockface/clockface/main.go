package main

import (
	"os"
	"time"

	"github.com/learn-go-luke/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
