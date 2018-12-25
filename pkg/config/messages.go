package config

import (
	"github.com/fatih/color"
)

func colorLogo() string {

	cyan := color.New(color.FgCyan)
	blue := color.New(color.FgHiBlue)
	logo := cyan.Sprint(`
   __  __    __               
  / /_/ /___/ /____  __    __ 
 / __/ / __  / ___/_/ /___/ /_`) + blue.Sprint(`
/ /_/ / /_/ / /  /_  __/_  __/
\__/_/\__,_/_/    /_/   /_/   
                              
`)
	return logo
}

func logo() string {
	logo := `
   __  __    __               
  / /_/ /___/ /____  __    __ 
 / __/ / __  / ___/_/ /___/ /_
/ /_/ / /_/ / /  /_  __/_  __/
\__/_/\__,_/_/    /_/   /_/   
                              
`
	return logo
}
