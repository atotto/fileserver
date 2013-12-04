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

	err := setServerRoot("/", *dir)
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("%s:%d", *addr, *port)
	log.Printf("Starting server on: %s", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("Server failed: ", err.Error())
	}
}

func setServerRoot(mount, dir string) error {
	if !IsDirExist(dir) {
		return errors.New(dir + ": No such directory")
	}

	path, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	log.Printf("Server root: %s at %s", filepath.ToSlash(path), mount)

	http.Handle(mount, http.FileServer(http.Dir(dir)))

	return nil
}

func IsDirExist(dir string) bool {
	f, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}

	if !f.IsDir() {
		return false
	}

	return true
}
