# Base64 encoder/decoder

This is a simple CLI command that encodes and decodes files using base64.

## Get and compile code

Before you start, you need to have Golang installed. For more help see [here][1]

```bash
go get github.com/c0nstantx/base64encdec
go build github.com/c0nstantx/base64encdec
```

## Usage

The command encodes and decodes files and print the result to standard output

Basic usage:
```bash
base64encdec [encode|decode] /path/to/file
```

## Examples

```bash
base64encdec encode /path/to/file								Print to stdout	(encode)
```
```bash
base64encdec encode /path/to/file > /path/to/encoded/file		Save to file	(encode)
```
```bash
base64encdec decode /path/to/file								Print to stdout	(decode)
```
```bash
base64encdec decode /path/to/file > /path/to/decoded/file		Save to file	(decode)
```


[1]:https://golang.org/doc/install