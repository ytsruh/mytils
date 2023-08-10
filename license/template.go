package license

import (
	"fmt"
	"os"
	"text/template"
)

func GenTemplate(input Result) error {
	tmplFile := fmt.Sprintf("license/templates/%s.txt", input.License)

	// Open a file for writing
	file, err := os.Create("LICENSE.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		return err
	}
	err = tmpl.Execute(file, input)
	if err != nil {
		return err
	}

	return nil
}
