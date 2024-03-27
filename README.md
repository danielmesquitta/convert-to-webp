# webpfyer CLI tool

## Overview

The webpfyer cli tool designed to convert png, jpg and jpeg files into webp. This can be particularly useful for website optimization.

## Features

- Convert entire folders;
- Set desired image quality;
- Command-line interface for easy integration into automation scripts.

## Installation

To install webpfyer, ensure you have Go installed on your system. Then, run the following command:

```bash
go install github.com/danielmesquitta/webpfyer@latest
```

## Usage

To use the cli tool, run the compiled binary with the "webpfyer" command and the necessary flags.

Example:

```bash
webpfyer -p path/to/images
```

## Flags:

| Shorthand | Command     | Description                                                                                  |
| --------- | ----------- | -------------------------------------------------------------------------------------------- |
| -h        | --help      | help for webpfyer                                                                            |
| -p        | --path      | string Path to a file folder or a single file to convert to webp. (default "./")             |
| -s        | --separator | uint Quality specify the compression factor for RGB channels between 0 and 100. (default 80) |
