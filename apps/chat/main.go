package main

import (
	"log"

	"github.com/heyrovsky/yolkchat/pkg/ui"
)

func main() {
	if err := ui.PaintUi(); err != nil {
		log.Fatalln(err)
	}
}
