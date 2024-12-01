package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

var (
	targetPath string
)

func main() {

	dirFS := os.DirFS("./")
	err := fs.WalkDir(dirFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		pathSlice := strings.Split(path, "/")
		pathLen := len(pathSlice)
		fmt.Printf("%s%s\n", strings.Repeat("  ", pathLen), pathSlice[pathLen-1])
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

}
