package pages

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/isacikgoz/tldr/pkg/config"
)

var (
	sep = string(os.PathSeparator)
	// a page should have md file extension
	ext = ".md"

	bold  = color.New(color.Bold)
	blue  = color.New(color.FgHiBlue)
	red   = color.New(color.FgRed)
	cyan  = color.New(color.FgCyan)
	white = color.New(color.FgWhite)
)

// Read finds and creates the Page, if it does not find, simply returns abstract
// contribution guide
func Read(seq []string) (p *Page, err error) {
	page := ""
	for i, l := range seq {
		if len(seq)-1 == i {
			page = page + l
			break
		} else {
			page = page + l + "-"
		}
	}
	// Common pages are more, so we have better luck there
	p, err = queryCommon(page)
	if err != nil {
		p, err = queryOS(page)
		if err != nil {
			return p, errors.New("This page (" + page + ") doesn't exist yet!\n" +
				"Submit new pages here: https://github.com/tldr-pages/tldr")
		}
	}
	return p, nil
}

// Queries from common folder
func queryCommon(page string) (p *Page, err error) {
	d := config.SourceDir + sep + "pages" + sep + "common" + sep
	b, err := ioutil.ReadFile(d + page + ".md")
	if err != nil {
		return p, err
	}
	p = ParsePage(string(b))
	return p, nil
}

// Queries from os specific folder
func queryOS(page string) (p *Page, err error) {
	d := config.SourceDir + sep + "pages" + sep + config.PageOSName() + sep
	b, err := ioutil.ReadFile(d + page + ".md")
	if err != nil {
		return p, err
	}
	p = ParsePage(string(b))
	return p, nil
}

func QueryRandom() (p *Page, err error) {
	d_cmn := config.SourceDir + sep + "pages" + sep + "common" + sep
	d_os := config.SourceDir + sep + "pages" + sep + config.PageOSName() + sep
	paths := []string{d_cmn, d_os}
	srcs := make([]string, 0)
	for _, p := range paths {
		files, err := ioutil.ReadDir(p)
		if err != nil {
			break
		}
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".md") {
				srcs = append(srcs, f.Name()[:len(f.Name())-3])
			}
		}
	}
	rand.Seed(time.Now().UTC().UnixNano())
	page := srcs[rand.Intn(len(srcs))]
	return Read([]string{page})
}

func ReadAll() (p *Page, err error) {
	d_cmn := config.SourceDir + sep + "pages" + sep + "common" + sep
	d_os := config.SourceDir + sep + "pages" + sep + config.PageOSName() + sep
	paths := []string{d_cmn, d_os}
	p = &Page{Name: "Search All"}
	p.Tips = make([]*Tip, 0)
	for _, pt := range paths {
		files, err := ioutil.ReadDir(pt)
		if err != nil {
			break
		}
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".md") {
				page, err := Read([]string{f.Name()[:len(f.Name())-3]})
				if err != nil {
					continue
				}
				p.Tips = append(p.Tips, page.Tips...)
			}
		}
	}
	return p, nil
}
