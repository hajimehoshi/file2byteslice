# file2byteslice

A dead simple tool to embed a file to Go.

```
Usage of file2byteslice:
  -buildtags string
        build tags
  -compress
        use gzip compression
  -input string
        input filename
  -output string
        output filename
  -package string
        package name (default "main")
  -var string
        variable name (default "_")
```

## How to use

```sh
file2byteslice -input INPUT_FILE -output OUTPUT_FILE -package PACKAGE_NAME -var VARIABLE_NAME
```
