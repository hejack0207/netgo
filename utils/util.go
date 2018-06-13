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

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func TransformWithPrompt(dst io.Writer, src io.Reader) {
	for {
		input := prompt.Input("", completer)
		log.Println("got input:" + input)
		dst.Write([]byte(input))
	}
}
