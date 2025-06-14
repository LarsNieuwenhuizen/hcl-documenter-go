# HCL Documenter

[![Tests](https://github.com/LarsNieuwenhuizen/hcl-documenter-go/actions/workflows/test.yml/badge.svg)](https://github.com/LarsNieuwenhuizen/hcl-documenter-go/actions/workflows/test.yml)

This application parses HCL files and generates documentation in Markdown format.

# Usage

The first feature here is to parse a variables file and generate a markdown table for you.

## Installation

Download the release for your platform either linux or mac.

Go to the [releases page](https://github.com/LarsNieuwenhuizen/hcl-documenter-go/releases/tag/v1.0.0)

Download the tar.gz file for your platform and extract it.

Move it to a directory in your `$PATH`, for example `/usr/local/bin`:

```shell
sudo mv hcldoc /usr/local/bin
```

```shell

Then you can run the binary and see something like this:

```shell
> hcldoc

Use this app to document your terraform files

Usage:
  hcldoc [command]

Available Commands:
  completion     Generate the autocompletion script for the specified shell
  help           Help about any command
  variables-file Convert variables file to markdown table

Flags:
  -h, --help   help for hcl-doc

Use "hcldoc [command] --help" for more information about a command.
```

### Create a markdown table from a variables file

You can use the `variables-file` command to convert a variables file to a markdown table.

Example:

```shell
> hcldoc vf -f /path/to/variables.tf

| Name | Type | Description         | Default |
|------|------|---------------------|---------|
| test | bool | This is just a test | false   |
```