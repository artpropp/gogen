package main

import (
	"fmt"
	"log"
	"text/template"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("gogen requires two arguments")
		fmt.Println("gogen [template] [typeName]")
		os.Exit(1)
	}

	templateFileName := os.Args[1]
	typeName := os.Args[2]

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Fatal(err)
	}

	outName := fmt.Sprintf("gogen_%s_gen.go", typeName)
	fd, err := os.OpenFile(outName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	data := struct {
		T string
	}{
		typeName,
	}
	t.Execute(fd, data)
}
