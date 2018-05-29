package main

import (
	"log"
	"os"
)

func AddFileSeed(path string) error {

	// check file
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	// file seed
	infohash, err := FileSeed(path)
	if err != nil {
		return err
	}

	// send hash
	AddHash(infohash)

	log.Println("ks-mac AddFileSeed", path, "infohash", infohash)
	return nil
}

func DelFileSeed(path string) error {

	// check file
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	// cancel file seed
	_, err = CancelFileSeed(path)
	if err != nil {
		return err
	}

	// del hash TODO:要不要删除, 因为别人有可能也在做种

	log.Println("ks-mac DelFileSeed", path)
	return nil
}
