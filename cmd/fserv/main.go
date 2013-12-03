package main

import (
	"net/http"

	"errors"
	"flag"
	"os"

	"fmt"
	"log"
	"path/filepath"
)

var addr = flag.String("addr", "127.0.0.1", "TCP network address")
var port = flag.Int("port", 8000, "port number")
var dir = flag.String("root", "./", "server root dir")

func main() {
	flag.Parse()

	setServerRoot("/", *dir)

	addr := fmt.Sprintf("%s:%d", *addr, *port)
	log.Printf("Starting server on: %s", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("Server failed: ", err.Error())
	}
}

func setServerRoot(mount, dir string) {
	err := mustDirExist(dir)
	if err != nil {
		log.Fatal(err)
	}

	path, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server root: %s", filepath.ToSlash(path))

	http.Handle(mount, http.FileServer(http.Dir(dir)))
}

func mustDirExist(dir string) error {
	f, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return errors.New(dir + ": No such directory")
	}

	if !f.IsDir() {
		return errors.New(dir + ": Should be directory")
	}

	return nil
}
