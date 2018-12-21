package prompt

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
)

var (
	pathTo = "path/to/"
	ext    = ".ext"

	filePathCompleter = completer.FilePathCompleter{
		IgnoreCase: true,
		Filter: func(fi os.FileInfo) bool {
			return fi.IsDir() || strings.HasSuffix(fi.Name(), ".go")
		},
	}
)

func completerLegacyFunc(t prompt.Document) []prompt.Suggest {
	files, _ := ioutil.ReadDir("./")
	s := make([]prompt.Suggest, 0)
	for _, f := range files {
		s = append(s, prompt.Suggest{
			Text: f.Name(),
		})
	}
	return s
}

func filePathCompleterFunc(d prompt.Document) []prompt.Suggest {
	return filePathCompleter.Complete(d)
}

func suggestCompleterFunc(arg string) (completer.FilePathCompleter, error) {
	if strings.Contains(arg, pathTo) {
		filePathCompleter := completer.FilePathCompleter{
			IgnoreCase: true,
		}
		return filePathCompleter, nil
	} else {
		ext := getFileExtension(arg)
		if len(arg) > 3 && len(ext) > 0 {
			filePathCompleter := completer.FilePathCompleter{
				IgnoreCase: true,
				Filter: func(fi os.FileInfo) bool {
					return strings.HasSuffix(fi.Name(), ext)
				},
			}
			return filePathCompleter, nil
		} else {

			return completer.FilePathCompleter{}, errors.New("Can't suggest file")
		}
	}
}

func getFileExtension(arg string) string {
	if strings.Contains(arg, "..") {
		return ""
	}
	r := []rune(arg[:len(arg)-2])
	var ext string
	for i := len(r) - 1; i >= 0; i-- {
		ext = string(r[i]) + ext
		if r[i] == '.' || i < len(r)-3 {
			break
		}
	}
	if strings.Contains(ext, ".") {
		return ext
	}
	return ""
}
