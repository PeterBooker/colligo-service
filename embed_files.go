// +build ignore

package main

import (
	"log"

	"github.com/PeterBooker/colligo/internal/assets"
	"github.com/PeterBooker/colligo/internal/templates"
	"github.com/shurcooL/vfsgen"
)

func main() {
	var err error
	err = vfsgen.Generate(assets.Files, vfsgen.Options{
		Filename:     "internal/assets/embed.go",
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "Files",
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = vfsgen.Generate(templates.Files, vfsgen.Options{
		Filename:     "internal/templates/embed.go",
		PackageName:  "templates",
		BuildTags:    "!dev",
		VariableName: "Files",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
