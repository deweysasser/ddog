# ddog

A DataDog CLI, focused on enabling IaC

## Overview

This program is a CLI for DataDog, focused on enabling IaC.

## Status:  Proof of Concept Only

This is current a proof of concept/initial implementation with effectively no features.

## Features

Initial features will include:
* save and restore monitors to local JSON files
* emit terraform compliant with the [DataDog provider](https://registry.terraform.io/providers/DataDog/datadog/latest/docs) to maintain monitors
* emit customizable terraform (almost certainly by golang template) to abstract monitors into
  terraform modules
* emit terraform import statements so that the above generated terraform can "take over" management
  of existing monitors

## Usage

```text
Usage: ddog <command>

Flags:
  -h, --help                      Show context-sensitive help.
      --version                   Show program version
      --datadog-api-key=STRING    Datadog API key ($DD_API_KEY)
      --datadog-app-key=STRING    Datadog APP key ($DD_APP_KEY)

Info
  --debug                   Show debugging information
  --output-format="auto"    How to show program output (auto|terminal|jsonl)
  --quiet                   Be less verbose than usual

Commands:
  monitor
    Manipulate monitors

Run "ddog <command> --help" for more information on a command.
```

## Examples
This will write monitors to the `/tmp/out` directory, with each monitor in JSON format in a separate
file.

```commandline
$ export DD_API_KEY=<generated datadog API key>
$ export DD_APP_KEY=<generated datadog APP key>
$ ddog monitor --save-to /tmp/out
```


## Installation

Binaries are not currently provided.  In the future, releases will be made with binaries for Linux,
MacOS, and Windows, and a homebrew tap will be provided.
