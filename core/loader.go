package core

import (
	
	"os"
	"gopkg.in/yaml.v3"
)

func LoadTools(path string)([]Tool,error){
	data , err := os.ReadFile(path)
	if err != nil {
		return nil,err 
	}

	var tools []Tool
	err = yaml.Unmarshal(data,&tools)
	if err != nil {
		return nil,err
	}

	return tools,nil





}
