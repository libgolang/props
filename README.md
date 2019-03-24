# Props
[![GoDoc](https://godoc.org/github.com/libgolang/props?status.svg)](https://godoc.org/github.com/libgolang/props)
[![Go Report Card](https://goreportcard.com/badge/github.com/libgolang/props)](https://goreportcard.com/report/github.com/libgolang/props)
[![Build Status](https://travis-ci.org/libgolang/props.svg?branch=master)](https://travis-ci.org/libgolang/props)

Command line flags made simple


## Installation
```bash
go get -u github.com/libgolang/props
```

## Usage
```go
import (
  "fmt"
  "github.com/libgolang/props"
)

func main() {
  if props.IsSet("name") {
    fmt.Printf("My Name is %s\n", props.GetProp("name"))
  } else {
    fmt.Printf("No name given\n")
  }
}
```

## Why
- Does not require pre-defining properties.  It dynamically figures out properties based on `os.Args`
- Does no require 3rd party libraries
- Just three functions `GetProp`, `IsSet` and `GetArgs`


## Features
- Arguments with two dashes and an equal sign are interpreted as properties with a value.
  e.g.:  `--prop-name=value`
- Arguments with two dashes and NO equal sign will interpret the next argument as a value.
  If no argument is followed by the property, then it will have an empty value.
  e.g.: `--prop-name value` then `GetProp("prop-name"): "value"` and `--prop-name --some-other-property` then `GetProp("prop-name"): ""`
- Arguments with one dash are interperted as Flags.  If a value is passed to the flag it is igned and used as an argument in `GetArgs()`.
  e.g.: `-p` then `IsSet("p") : true`
- Any argument not part of a flag or a property is added to the argument list and they can be retrived by calling `GetArgs()`.


