package main 

import (
	"fmt"	
	"os"
	"github.com/codedloki/securestack/core"
	
)


// type Tool struct {
// 	Name string `yaml:"name"`
// 	Category string `yaml:"category"`
// 	Auth string `yaml:"auth"`
// 	Deployment string `yaml:"deployment"`
// 	Maintained bool `yaml:"maintained"`
	
// }






func main(){
	if len(os.Args)<3{
		fmt.Fprintln(os.Stderr,"usage : securestack validate <config path>")
		os.Exit(1)
	}

	command := os.Args[1]
	configPath := os.Args[2]

	if command != "validate" {
		fmt.Fprintln(os.Stderr,"unknown command : " ,command)
		os.Exit(1)
	}

	tools, err :=  core.LoadTools(configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr,"\033[31mERROR:",err)
		os.Exit(1)
	}

	err = core.Validator(tools)
	if err != nil {
		fmt.Fprintln(os.Stderr,"\033[31mError:",err)
		os.Exit(1)
	}
	fmt.Println("\033[32mvalidation successful")
}
