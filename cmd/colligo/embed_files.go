// +build ignore

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PeterBooker/colligo/internal/data"
	"github.com/shurcooL/vfsgen"
)

func main() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	var err error
	err = vfsgen.Generate(data.Assets, vfsgen.Options{
		Filename:     "../../internal/data/assets.go",
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = vfsgen.Generate(data.Templates, vfsgen.Options{
		Filename:     "../../internal/data/templates.go",
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Templates",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
