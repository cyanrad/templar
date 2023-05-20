package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/template"
)

type Template struct {
	Name   string
	Out    string
	Values []map[string]interface{}
}

func main() {
	// getting input data anc converting it into a map
	inputJson, err := os.Open("./operations.json")
	handleError(err)
	inputJsonData, err := io.ReadAll(inputJson)
	handleError(err)
	operations := make([]Template, 0)
	json.Unmarshal(inputJsonData, &operations)

	// executing templates
	for i, input := range operations {
		// reading template data
		file, err := os.Open("./templates/" + input.Name)
		handleError(err)
		templateData, err := io.ReadAll(file)
		handleError(err)

		// parsing template
		templateObj, err := template.New(input.Name + "_" + fmt.Sprint(i)).Parse(string(templateData))
		handleError(err)

		// getting file name
		nameComponents := splitFileName(input.Out)

		// creating output folder
		os.Mkdir("./output/"+input.Name, 0777)

		for i, value := range input.Values {
			filePath := fmt.Sprintf("output/%s/%s_%d%s", nameComponents[0], nameComponents[0], i, nameComponents[1])
			file, err := os.Create(filePath)
			handleError(err)
			defer file.Close()

			templateObj.Execute(file, value)
		}
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 0: name, 1: ext
func splitFileName(s string) (ret [2]string) {
	ok := false
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == '.' {
			ret[0] = s[:i]
			ret[1] = s[i:]
			ok = true
			break
		}
	}

	if ok {
		return
	} else {
		ret[0] = s
		ret[1] = ""
		return
	}
}
