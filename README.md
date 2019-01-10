# Go ABN

[![Build Status](https://travis-ci.org/ace-teknologi/go-abn.svg?branch=master)](https://travis-ci.org/ace-teknologi/go-abn)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fsjauld%2Fgo-abn.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fsjauld%2Fgo-abn?ref=badge_shield)
[![GoDoc Status](https://godoc.org/github.com/ace-teknologi/go-abn?status.svg)](http://godoc.org/github.com/ace-teknologi/go-abn)

A Go wrapper for the
[Australian Business Register](https://abr.business.gov.au/abrxmlsearch/abrxmlsearch.asmx)

## Usage

1. [Register for a GUID](https://www.abr.business.gov.au/RegisterAgreement.aspx)
2. Set the `ABR_GUID` environment variable to the GUID issued to you.

### Search

Search by the name fields of the ABN entries.

```bash
goabn search -s "Bob's Country Bunker" --GUID 123-456-789
```

### Find by ABN

When you have an ABN you can get further information. For example:

```bash
goabn find-abn -s 33102417032 --GUID 123-456-789
```

### Find by ACN

```bash
goabn find-acn -s 102417032 --GUID 123-456-789
```

### Options

#### Output Types

There are three output types available:

* text
* json
* xml

Set via the `-f` or `--output-format` flag. Example:

```bash
goabn search -s "Bob's Country Bunker" -f json --GUID 123-456-789
```

#### Custom Text Output Template

Use go's [text/template](https://golang.org/pkg/text/template/) formatted
template to customise the output as required.

Set via the `-t` or `--text-output-template` flag. Example:

```bash
goabn search -s "Bob's Country Bunker" -f json -t "./tmp/my-custom-template.gtpl" --GUID 123-456-789
```

## Testing

1.  Run:
    ```
    go test ./abr
    ```

## Documenation

* [abr library documentation](https://godoc.org/github.com/ace-teknologi/go-abn/abr)
* [command line interface documentation](https://godoc.org/github.com/ace-teknologi/go-abn/cmd)

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fsjauld%2Fgo-abn.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fsjauld%2Fgo-abn?ref=badge_large)
