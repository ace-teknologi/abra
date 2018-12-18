# Go ABN

[![Build Status](https://travis-ci.org/ace-teknologi/go-abn.svg?branch=master)](https://travis-ci.org/ace-teknologi/go-abn)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fsjauld%2Fgo-abn.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fsjauld%2Fgo-abn?ref=badge_shield)
[![GoDoc Status](https://godoc.org/github.com/ace-teknologi/go-abn?status.svg)](http://godoc.org/github.com/ace-teknologi/go-abn)

A Go wrapper for the
[Australian Business Register](https://abr.business.gov.au/abrxmlsearch/abrxmlsearch.asmx)

## Requirements

* Go 1.11.x

## Usage

1. [Register for a GUID](https://www.abr.business.gov.au/RegisterAgreement.aspx)
2. Set the `ABR_GUID` environment variable to the GUID issued to you.

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
