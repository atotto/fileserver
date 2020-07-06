package main

import (
	"crypto/tls"
	"net/http"

	"errors"
	"flag"
	"os"

	"fmt"
	"log"
	"path/filepath"
)

var (
	addr = flag.String("addr", "127.0.0.1", "TCP network address")
	port = flag.Int("port", 8080, "port number")
	dir  = flag.String("root", "./", "server root dir")

	useTls = flag.Bool("tls", false, `use tls (https)
example generate private key and certificate:
  go get -u github.com/FiloSottile/mkcert
  mkcert -install
  mkdir -p ~/.config/fserv && cd ~/.config/fserv
  mkcert -ecdsa fserv.local localhost
`)

	keyFile  = flag.String("key", os.ExpandEnv("${HOME}/.config/fserv/fserv.local+1-key.pem"), "key path")
	certFile = flag.String("cert", os.ExpandEnv("${HOME}/.config/fserv/fserv.local+1.pem"), "cert path")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	if err := setServerRoot(mux, "/", *dir); err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("%s:%d", *addr, *port)
	log.Printf("Starting server on: %s", addr)

	config := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		TLSConfig:    config,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	var err error
	if *useTls {
		err = server.ListenAndServeTLS(*certFile, *keyFile)
	} else {
		err = server.ListenAndServe()
	}

	if err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}

func setServerRoot(mux *http.ServeMux, mount, dir string) error {
	if !IsDirExist(dir) {
		return errors.New(dir + ": No such directory")
	}

	path, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	log.Printf("Server root: %s at %s", filepath.ToSlash(path), mount)

	mux.Handle(mount, http.FileServer(http.Dir(dir)))

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
