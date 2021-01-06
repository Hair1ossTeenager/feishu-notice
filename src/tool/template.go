package tool

import (
	"bytes"
	"fmt"
	"text/template"
)

const TemplateDir = "res/templates/"

func Render(data interface{}, templateFile string) bytes.Buffer {
	templatePath := TemplateDir + templateFile
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	var buff bytes.Buffer
	template.Execute(&buff, data)
	fmt.Println(&buff)
	return buff
}
