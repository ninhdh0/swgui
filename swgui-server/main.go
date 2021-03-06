package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/lazada/swgui"
)

var (
	host  string
	port  uint
	title string
)

func init() {
	flag.StringVar(&host, "host", "", "Host")
	flag.UintVar(&port, "port", 8080, "Port")
	flag.StringVar(&title, "title", "API Document", "Page title")

	flag.Parse()
}

func main() {
	http.Handle("/", swgui.NewHandler(title, "/swagger.json", "/"))
	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}
