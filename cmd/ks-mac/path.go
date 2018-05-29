package main

import (
	"fmt"
	"net/http"
	"os"
)

func ListenPath() {

	mux := http.NewServeMux()
	mux.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Query().Get("key")
		fmt.Fprintf(w, "hello, %s!\n\n%s\n", path, AddDir(path))
	})
	http.ListenAndServe(":16180", mux)
}

func AddDir(path string) error {

	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	infohash, err := FileSeed(path, torrentClient)
	if err != nil {
		return err
	}

	fmt.Println("ks-mac", "path", path, "infohash", infohash)

	SendHash(infohash)

	return nil
}
