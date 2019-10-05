package main

import (
	"net/http"

	"github.com/goproxy/goproxy"
	"github.com/goproxy/goproxy/cacher"
)

func main() {
	p := goproxy.New()
	p.Cacher = &cacher.Disk{Root: "/var/db/goproxy"}
	http.ListenAndServe(":http", p)
}
