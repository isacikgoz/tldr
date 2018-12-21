package prompt

import (
	"errors"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
)

var (
	pathTo = "path/to/"
	ext    = ".*"
)

// if the arg extension is matched, suggested values moves top of the slice
// implementation could be beter
func fileExtCompleterFunc(t prompt.Document) []prompt.Suggest {
	s := make([]prompt.Suggest, 0)
	if len(ext) > 0 {
		filePathExtCompleter := completer.FilePathCompleter{
			IgnoreCase: true,
			Filter: func(fi os.FileInfo) bool {
				promoted := strings.HasSuffix(fi.Name(), ext)
				return promoted
			},
		}
		s = filePathExtCompleter.Complete(t)
	}
	f := filePathCompleterFunc(t)
	s = append(s, f...)

	return removeDuplicates(s)
}

// default file path completer, return all files
func filePathCompleterFunc(d prompt.Document) []prompt.Suggest {
	filePathCompleter := completer.FilePathCompleter{
		IgnoreCase: true,
	}
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
		// the arg should be longer than regular extension length such as "a.z"
		if len(arg) > 3 && len(ext) > 0 {
			filePathCompleter := completer.FilePathCompleter{
				IgnoreCase: true,
				Filter: func(fi os.FileInfo) bool {
					promoted := strings.HasSuffix(fi.Name(), ext)
					return promoted
				},
			}
			return filePathCompleter, nil
		} else {

			return completer.FilePathCompleter{}, errors.New("Can't suggest file")
		}
	}
}

// returns the file extension of the argument
func getFileExtension(arg string) string {
	// probably not a file. hence, wont have an extension
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
	// we expect a dot to determine if it is an extension
	if strings.Contains(ext, ".") {
		return ext
	}
	return ""
}

// removes duplicate entries from prompt.Suggest slice
func removeDuplicates(elements []prompt.Suggest) []prompt.Suggest {
	// Use map to record duplicates as we find them.
	encountered := map[prompt.Suggest]bool{}
	result := []prompt.Suggest{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
