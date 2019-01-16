[![Build Status](https://travis-ci.com/isacikgoz/tldr.svg?branch=master)](https://travis-ci.com/isacikgoz/tldr) [![MIT License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](/LICENSE) [![Join the chat at https://gitter.im/tldrpp/community](https://badges.gitter.im/tldrpp/community.svg)](https://gitter.im/tldrpp/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) ![Release](https://img.shields.io/github/release/isacikgoz/tldr.svg)

# tldr++
Community driven man pages improved with smart user interaction. **tldr++** seperates itself from any other tldr client with convenient user guidance feature.

![screenplay](img/screenplay.gif)

## Features
- Fully Interactive (fill the command arguments easily)
- Search from commands to find your desired command (exact + fuzzy search)
- Smart file suggestions (further suggestions will be added)
- Simple implementation
- One of the fastest clients, even fastest (see [Benchmarks](https://github.com/isacikgoz/tldr/wiki/Benchmarks))
- Easy to install. Supports all mainstream OS and platforms (Linux, MacOS, Windows)(arm, x86)
- Pure-go (*even contains built-in git*)

## Installation

Refer to [Release Page](https://github.com/isacikgoz/tldr/releases) for binaries.

Or, you can build from source: (min. **go 1.10** compiler is recommended)

```bash
go get -u github.com/isacikgoz/tldr
```

## Credits
- [tldr-pages](https://github.com/tldr-pages/tldr)
- [survey](https://github.com/AlecAivazis/survey)
- [go-prompt](https://github.com/c-bata/go-prompt)
- [fuzzy](https://github.com/sahilm/fuzzy)
- [go-git](https://github.com/src-d/go-git)
- [kingpin](https://github.com/alecthomas/kingpin)
- [color](https://github.com/fatih/color)
