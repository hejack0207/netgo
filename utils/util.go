package utils

import (
	"github.com/c-bata/go-prompt"
	"io"
	"log"
)

func Transform(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func TransformPrompt(dst io.Writer, src io.Reader) {
	input := prompt.Input(nil, nil, nil)
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
