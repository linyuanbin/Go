package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"strings"
)

type model struct {
	Name   string
	Type   string
	DBName string
	DBType string
}

type data struct {
	Name   string  `json:"name"`
	Models []model `json:"models"`
}

func outFile(tmplFilePath string, writeFilePath string, data interface{}) error {
	t := template.Must(template.New(Base(tmplFilePath)).ParseFiles(tmplFilePath))
	buff := bytes.NewBufferString("")
	err := t.Execute(buff, data)
	if err != nil {
		return err
	}
	src, er := format.Source(buff.Bytes())
	if er != nil {
		return er
	}
	e := ioutil.WriteFile(writeFilePath, src, 0766)
	if e != nil {
		return e
	}
	return nil
}

func Base(path string) string {
	if path == "" {
		return "."
	}
	for len(path) > 0 && path[len(path)-1] == '/' {
		path = path[0 : len(path)-1]
	}
	if i := strings.LastIndex(path, "/"); i >= 0 {
		path = path[i+1:]
	}
	if path == "" {
		return "/"
	}
	return path
}

func ParseJsonFile(filePath string, v interface{}) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	if err = json.Unmarshal(data, v); err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
}

func main()  {
	da := data{}
	ParseJsonFile("../../../config/model.json", &da)
	d := map[string]interface{}{
		"name":   da.Name,
		"models": da.Models,
	}
	er := outFile("auto/model.go.tmpl", "out/"+da.Name+".go", d)
	fmt.Println(er)
}