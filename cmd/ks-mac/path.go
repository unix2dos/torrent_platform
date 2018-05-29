package main

import (
	"fmt"
	"os"
)

func AddPath(path string) error {

	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	infohash, err := FileSeed(path, torrentClient)
	if err != nil {
		return err
	}

	SendHash(infohash)

	fmt.Println("ks-mac", "path", path, "infohash", infohash)

	return nil
}
