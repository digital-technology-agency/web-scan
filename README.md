# web-scan

<p><img src="https://rawcdn.githack.com/devicons/devicon/9c6bfdb9783cdfe1018666ed76adcfd3eab6fad6/icons/go/go-original.svg" alt="go" width="100" height="100"/></p>

[![Donate](https://img.shields.io/badge/Donat-For%20project-yellow)](https://www.donationalerts.com/r/dta_agency)
![Platform](https://img.shields.io/badge/platform-win%20%7C%20mac%20%7C%20linux-blue)

# Intro

[![GoDoc](https://godoc.org/github.com/digital-technology-agency/web-scan?status.svg)](https://godoc.org/github.com/digital-technology-agency/web-scan)
[![Go](https://github.com/digital-technology-agency/web-scan/actions/workflows/go.yml/badge.svg)](https://github.com/digital-technology-agency/web-scan/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/digital-technology-agency/web-scan)](https://goreportcard.com/report/github.com/digital-technology-agency/web-scan)
[![License](http://img.shields.io/badge/Licence-MIT-brightgreen.svg)](LICENSE)
[![Website dta.agency](https://img.shields.io/website-up-down-green-red/http/shields.io.svg)](https://dta.agency)
[![GitHub release](https://img.shields.io/github/v/release/digital-technology-agency/web-scan)](https://github.com/digital-technology-agency/web-scan/releases/latest)
<p>&nbsp;<img align="center" src="https://github-readme-stats.vercel.app/api?username=digitaltechnologyagency&show_icons=true&count_private=true" alt="Агентство цифровых технологий" width="80%"/></p>

A program that generates a combination of letters from the alphabet and composes an address for accessing the site.After
that, it reads the data from the title and description of this site and saves it to the structure. At the end, the
program displays all the titles and descriptions of the available sites.

## Usage

### [golang/cmd/go](https://golang.org/cmd/go/)

```bash
go get github.com/digital-technology-agency/web-scan
```

## Data store

* sqlite - SQLite
* jsoneachrow - Json Each row

## Generator type

- simple - Simple generator type

## Init configuration file

```bash
$ ./wscan init

```

---

#### Configuration file `config.json`

```json
{
  "process_count": 1,
  "alphabet": "abcdefgefghijklmnop",
  "url_len": 5,
  "concurrency_count": 5,
  "data_store_type": "sqlite",
  "generator_type": "simple",
  "protocol_types": [
    "http",
    "https"
  ]
}
```

## Run

```bash
$ ./wscan -configuration_file config.json
```
