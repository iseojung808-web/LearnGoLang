// Command finger-gun-game serves the webcam finger-gun game on localhost.
// Browsers only allow camera access on https or localhost, so opening
// index.html straight from disk may not work — run this instead:
//
//	go run .
package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed index.html
var site embed.FS

func main() {
	addr := flag.String("addr", "localhost:8080", "address to listen on")
	flag.Parse()

	root, err := fs.Sub(site, ".")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(root)))

	fmt.Printf("🔫 핑거건 좀비 사냥: http://%s 로 접속하세요 (Ctrl+C 로 종료)\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
