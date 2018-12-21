package pages

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/isacikgoz/tldr/pkg/config"
)

var (
	sep = string(os.PathSeparator)
	ext = ".md"

	bold  = color.New(color.Bold)
	blue  = color.New(color.FgBlue)
	red   = color.New(color.FgRed)
	cyan  = color.New(color.FgCyan)
	white = color.New(color.FgWhite)
)

func Read(page string) (p *Page, err error) {
	p, err = queryCommon(page)
	if err != nil {
		p, err = queryOS(page)
		if err != nil {
			return p, errors.New("This page doesn't exist yet!\n" +
				"Submit new pages here: https://github.com/tldr-pages/tldr")
		}
	}
	return p, nil
}

func queryCommon(page string) (p *Page, err error) {
	d := config.SourceDir + sep + "pages" + sep + "common" + sep
	b, err := ioutil.ReadFile(d + page + ".md")
	if err != nil {
		return p, err
	}
	p = ParsePage(string(b))
	return p, nil
}

func queryOS(page string) (p *Page, err error) {
	d := config.SourceDir + sep + "pages" + sep + config.OSName() + sep
	b, err := ioutil.ReadFile(d + page + ".md")
	if err != nil {
		return p, err
	}
	p = ParsePage(string(b))
	return p, nil
}
