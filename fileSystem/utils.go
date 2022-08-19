package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getAllFilesWithEnding(startDir, fileEnding string) ([]string, error) {
	r := []string{}
	err := filepath.Walk(startDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(path, fileEnding) {
				r = append(r, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s", r)

}
