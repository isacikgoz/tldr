package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	cp "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/isacikgoz/tldr/pages"
)

var (
	pathTo = "path/to/"
	ext    = ".*"
)

// SuggestCommand lets user to fill args with suggestions
func SuggestCommand(item interface{}) string {
	answers := make([]string, 0)
	t, ok := item.(*pages.Tip)
	if !ok {
		return ""
	}
	fmt.Println()
	for _, arg := range t.Cmd.Args {
		// cs, _ := suggestCompleterFunc(arg)
		ext = getFileExtension(arg)
		answers = append(answers, cp.Input(
			"$"+" "+arg+" -> ",
			fileExtCompleterFunc,
			cp.OptionPreviewSuggestionTextColor(cp.Cyan),
			cp.OptionInputTextColor(cp.Cyan),
			cp.OptionAddKeyBind(cp.KeyBind{
				Key: cp.ControlC,
				Fn: func(buf *cp.Buffer) {
					os.Exit(0)
					return
				}}),
			cp.OptionAddKeyBind(cp.KeyBind{
				Key: cp.Escape,
				Fn: func(buf *cp.Buffer) {

					return
				}}),
			cp.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
		))
	}

	fs := t.Cmd.String()
	for i, arg := range t.Cmd.Args {
		// since we have double curlybraces on args
		fs = strings.Replace(fs, "{{"+arg+"}}", answers[i], 1)
	}
	return fs
}

// if the arg extension is matched, suggested values moves top of the slice
// implementation could be beter
func fileExtCompleterFunc(t cp.Document) []cp.Suggest {
	s := make([]cp.Suggest, 0)
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
func filePathCompleterFunc(d cp.Document) []cp.Suggest {
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
	}
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
	}
	return completer.FilePathCompleter{}, errors.New("Can't suggest file")
}

// returns the file extension of the argument
func getFileExtension(arg string) string {
	// probably not a file. hence, wont have an extension
	if strings.Contains(arg, "..") || len(arg) < 2 {
		return ""
	}
	// since the args is surrounded with "}}"
	r := []rune(arg)
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
func removeDuplicates(elements []cp.Suggest) []cp.Suggest {
	// Use map to record duplicates as we find them.
	encountered := map[cp.Suggest]bool{}
	result := []cp.Suggest{}

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
