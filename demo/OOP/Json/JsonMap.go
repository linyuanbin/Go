package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	jsonArr:="{\"data\": [{\"id\": \"27\",\"name\": \"an\"}]}"
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonArr), &m)
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		d:=m["data"]
		if v, ok := d.([]interface{})[0].(map[string]interface{}); ok {
			fmt.Println(ok, v["id"], v["name"])
		}
	}
}
