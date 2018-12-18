# CHANGELOG.md

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v0.2.0

### Added

* Command Line Interface support for search by name as `search`
* Supports the HTTP POST ABRSearchByNameAdvancedSimpleProtocol2017 method
  including test cases

### Changed

* Command Line Interface refactor to provide `find-abn` and `find-acn` services.

## v0.1.0

### Added

* Support for ACN validation
* Support for the HTTP POST SearchByASICv201408 method
* Test cases for [./abr/abr.go](./abr/abr.go)

### Changed

* Mocking of HTTP requests in tests

## v0.0.1

* Initial release

### Added

* Support the HTTP POST SearchByABNv201408 method
* Support for ABN validation
