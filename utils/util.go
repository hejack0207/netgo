package utils

import (
	"bufio"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/grt1st/netgo/logging"
	"io"
	"strings"
	//"io/ioutil"
	"log"
	"os"
)

var (
	historyFilepath = os.Getenv("HOME") + "/.config/netgo/history"
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
	histFile, err := os.OpenFile(historyFilepath, os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("error open file %s", historyFilepath)
		return
	}
	defer histFile.Close()

	executor := func(input string) {
		logging.Debug("got input:" + input)
		histFile.WriteString(input + "\n")
		dst.Write([]byte(input + "\n"))
	}

	histReader := bufio.NewReader(histFile)
	histories := []string{}

	for {
		inputStr, err := histReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		histories = append(histories, strings.TrimSuffix(inputStr, "\n"))
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("ga?>"),
		prompt.OptionHistory(histories),
		prompt.OptionTitle("netgo-prompt"),
	)
	p.Run()
}
