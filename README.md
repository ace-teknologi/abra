# Abra

[![Build Status](https://travis-ci.org/ace-teknologi/abra.svg?branch=master)](https://travis-ci.org/ace-teknologi/abra)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Face-teknologi%2Fabra.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Face-teknologi%2Fabra?ref=badge_shield)
[![GoDoc Status](https://godoc.org/github.com/ace-teknologi/abra?status.svg)](http://godoc.org/github.com/ace-teknologi/abra)

A Go wrapper for the
[Australian Business Register](https://abr.business.gov.au/abrxmlsearch/abrxmlsearch.asmx)

![Australian Business Register Applicance](./abra.png)

## Usage

1. [Register for a GUID](https://www.abr.business.gov.au/RegisterAgreement.aspx)
2. Set the `ABR_GUID` environment variable to the GUID issued to you.

### Search

Search by the name fields of the ABN entries.

```bash
abra search -s "Bob's Country Bunker" --GUID 123-456-789
```

### Find by ABN

When you have an ABN you can get further information. For example:

```bash
abra find-abn -s 33102417032 --GUID 123-456-789
```

### Find by ACN

```bash
abra find-acn -s 102417032 --GUID 123-456-789
```

### Options

#### Output Types

There are three output types available:

* text
* json
* xml

Set via the `-f` or `--output-format` flag. Example:

```bash
abra search -s "Bob's Country Bunker" -f json --GUID 123-456-789
```

#### Custom Text Output Template

Use go's [text/template](https://golang.org/pkg/text/template/) formatted
template to customise the output as required.

Set via the `-t` or `--text-output-template` flag. Example:

```bash
abra search -s "Bob's Country Bunker" -f "text" \
  -t "./tmp/my-custom-template.gtpl" --GUID 123-456-789
```

With sample template:
```go
Name: {{.Name}}
Link: https://abr.business.gov.au/ABN/View?abn={{.ABN.IdentifierValue}}
```

## Testing

1.  Run:
    ```
    go test ./abra
    ```

## Documenation

* [abr library documentation](https://godoc.org/github.com/ace-teknologi/abra/abra)
* [command line interface documentation](https://godoc.org/github.com/ace-teknologi/abra/cmd)

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Face-teknologi%2Fabra.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Face-teknologi%2Fabra?ref=badge_large)
