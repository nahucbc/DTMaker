package args

import (
	"flag"
	"fmt"
	"os"

	"github.com/nahucbc/DTMaker/utils"
)

type Config struct {
	Build struct {
		Version     int    `toml:"version"`
		Description string `toml:"description"`
	} `toml:"build"`
}

func Parse_flag() {
	init := flag.Bool("init", false, "Initialize a DTMaker project")
	name := flag.String("name", "template", "name of the datapack, by default is template")
	version := flag.Int("version", 4, "format version, by default is 4 (1.13–1.14.4)")
	description := flag.String("description", "generated datapack by DTMaker", "description of the datapack")
	flag.Parse()

	if *init {
		initialize_project(name, version, description)
	}
}

func initialize_project(name *string, version *int, description *string) {
	os.Mkdir(*name, 0755)

	build_file := fmt.Sprintf("%s/build.toml", *name)
	file, err := os.Create(build_file)

	if err != nil {
		fmt.Println(err)
	}

	var conf Config
	conf.Build.Description = *description
	conf.Build.Version = *version

	encoder, err := utils.CustomEncoder(conf)

	if err != nil {
		fmt.Println(err)
	}

	file.Write(encoder)
	defer file.Close()
}
