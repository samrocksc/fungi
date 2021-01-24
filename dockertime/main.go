package main

import (
	"flag"

	"github.com/samrocksc/fungi/dockertime/container"
)

func main() {
	image_name := flag.String("image", "hello-world", "the image you would like to launch")
	flag.Parse()
	container.CreateNewContainer(*image_name)
}
