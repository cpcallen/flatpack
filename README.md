# flatpack

A Go utilitity to to convert arbitrary Go data, possibly including
unexported fields and shared and/or cyclic substructure, into (and
back from) a purely tree-structured datastructure without unexported
fields that can be serialised using encoding.json or the like.

[![GoDoc](https://godoc.org/github.com/cpcallen/flatpack?status.svg)](https://godoc.org/github.com/cpcallen/flatpack)

See [full documentation on
godoc.org](https://godoc.org/github.com/cpcallen/flatpack)

This was forked (with history, but minus all unrelated parts of the
project) from the [CodeCity
project](https://github.com/neilFraser/CodeCity/) when that project
abandoned development in Go.
