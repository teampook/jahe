package service

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/teampook/jahe/tmpl"
)

type Skeleton struct{}

func (s *Skeleton) CreateModels() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	modelPath := fmt.Sprint(path, "/models")
	err = os.MkdirAll(modelPath, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	StructModel(modelPath)
}

func StructModel(modelPath string) {

	// flag.StringVar(&d.Type, "type", "", "The subtype used for the queue being generated")
	// flag.StringVar(&d.Name, "name", "", "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
	// flag.Parse()
	x := tmpl.StructTmpl{
		StructName: "Test",
		FieldName: []tmpl.StructField{
			tmpl.StructField{
				Name:     "Mantap",
				DataType: "string",
			},
			tmpl.StructField{
				Name:     "Mantap1",
				DataType: "string",
			},
		},
	}

	mainFile, err := os.Create(fmt.Sprintf("%s/%s.go", modelPath, x.StructName))
	if err != nil {
	}
	defer mainFile.Close()

	t := template.Must(template.New("queue").Parse(string(tmpl.StructTemplate())))
	err = t.Execute(mainFile, x)
	if err != nil {
	}

	// _ = template.Must(template.New("queue").Parse(queueTemplate))
	// t.Execute(wc, d)
	// wc.Close()
	// io.Copy(os.Stdout, rc)
}
