# web-scan

<p><img src="https://rawcdn.githack.com/devicons/devicon/9c6bfdb9783cdfe1018666ed76adcfd3eab6fad6/icons/go/go-original.svg" alt="go" width="100" height="100"/></p>

# Intro
[![GoDoc](https://godoc.org/github.com/digital-technology-agency/web-scan?status.svg)](https://godoc.org/github.com/digital-technology-agency/web-scan)
[![Go](https://github.com/digital-technology-agency/web-scan/actions/workflows/go.yml/badge.svg)](https://github.com/digital-technology-agency/web-scan/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/digital-technology-agency/web-scan)](https://goreportcard.com/report/github.com/digital-technology-agency/web-scan)
[![License](http://img.shields.io/badge/Licence-MIT-brightgreen.svg)](LICENSE)
[![Website dta.agency](https://img.shields.io/website-up-down-green-red/http/shields.io.svg)](https://dta.agency)
[![GitHub release](https://img.shields.io/github/v/release/digital-technology-agency/web-scan)](https://github.com/digital-technology-agency/web-scan/releases/latest)

<p>&nbsp;<img align="center" src="https://github-readme-stats.vercel.app/api?username=digitaltechnologyagency&show_icons=true&count_private=true" alt="Агентство цифровых технологий" width="80%"/></p>

A program that generates a combination of letters from the alphabet and composes an address for accessing the site.After that, it reads the data from the title and description of this site and saves it to the structure. At the end, the program displays all the titles and descriptions of the available sites.


## Usage
### [golang/cmd/go](https://golang.org/cmd/go/)

```bash
go get github.com/digital-technology-agency/web-scan/pkg/models
go get github.com/digital-technology-agency/web-scan/pkg/utils
```

## Run on Linux

```bash
$ ./wscan -alphabet "abcdef" -len 2
```
