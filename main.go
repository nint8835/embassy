package main

import (
	"fmt"
	"os"

	"github.com/nint8835/embassy/app"
)

func main() {
	err := app.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error starting embassy: %s", err)
	}
}
