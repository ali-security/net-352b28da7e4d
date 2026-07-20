package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <module> <version> <dir> <out.zip>")
		os.Exit(1)
	}
	modPath, version, dir, out := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	f, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, dir); err != nil {
		panic(err)
	}
	fmt.Println("wrote", out)
}
