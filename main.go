package main

import (
	"fmt"
	"os"

	"ytsruh.com/license-gen/license"
)

func main() {
	result, err := license.RunPrompt()
	if err != nil {
		fmt.Printf("Generator failed %v\n", err)
		os.Exit(1)
	}

	err = license.GenTemplate(result)

	if err != nil {
		fmt.Printf("Template failed to generate %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Template has been generated & saved to your path. \n")
	os.Exit(0)
}
