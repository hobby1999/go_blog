package main

import (
	"go_blog/plugin"
	"go_blog/routers"
)

func main() {
	plugin.InitAdmin()
	routers.Start()

}
