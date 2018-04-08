# Go ABN

[![Build Status](https://travis-ci.org/sjauld/go-abn.svg?branch=master)](https://travis-ci.org/sjauld/go-abn)

A Go wrapper for the [Australian Business Register](https://abr.business.gov.au/abrxmlsearch/abrxmlsearch.asmx)

## Usage

1. [Register](https://www.abr.business.gov.au/RegisterAgreement.aspx) for a GUID
2. Set the ABR_GUID environment variable to your GUID

## Testing

1. [Register](https://www.abr.business.gov.au/RegisterAgreement.aspx) for a GUID
2. Run `TEST_ABR_GUID=<your guid> go test ./abr`
