# path

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/rickb777/path)
[![Go Report Card](https://goreportcard.com/badge/github.com/rickb777/path)](https://goreportcard.com/report/github.com/rickb777/path)
[![Build](https://github.com/rickb777/path/actions/workflows/go.yml/badge.svg)](https://github.com/rickb777/path/actions)
[![Coverage](https://coveralls.io/repos/github/rickb777/path/badge.svg?branch=main)](https://coveralls.io/github/rickb777/path?branch=main)
[![Issues](https://img.shields.io/github/issues/rickb777/path.svg)](https://github.com/rickb777/path/issues)

This enhances the standard path API with some extra functions. This API is intended to be a drop-in replacement;
it merely calls through to the standard API where there is duplication.

There is also a type Path, which is a kind of string. Path provides a similar set of methods to the helper functions.

Please see the [GoDoc](https://godoc.org/github.com/rickb777/path) for more.

## Installation

    go get -u github.com/rickb777/path

## Status

This library has been in reliable production use for some time. Versioning follows the well-known semantic version pattern.

## Licence

[MIT](LICENSE)
