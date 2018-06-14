package utils

import (
	"github.com/c-bata/go-prompt"
	"github.com/grt1st/netgo/logging"
	"io"
	"log"
	"os"
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
	executor := func(input string) {
		os.Stdout.WriteString("got input:" + input)
		logging.Debug("got input:" + input)
		dst.Write([]byte(input + "\n"))
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("netgo-prompt"),
	)
	p.Run()
}
