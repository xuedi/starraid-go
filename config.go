package main

import (
	"fmt"
)

type config struct {
	fullscreen bool
	width      int32
	height     int32
}

func createConfig(path string) config {
	var cfg config
	cfg.width = 800
	cfg.height = 600
	cfg.fullscreen = false

	return cfg
}

func (e config) show() {
	fmt.Println("Show:", e.width, "x", e.height)
}
