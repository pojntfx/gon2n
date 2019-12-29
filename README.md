# gon2n

Go bindings and CLI for n2n.

## Installation

A Go package [is available](https://godoc.org/github.com/pojntfx/gon2n). In order to use it, you have to `go generate` it first.

## Usage

You may also set the flags by setting env variables in the format `GON2N_[FLAG]` (i.e. `GON2N_EDGE_DEVICEIP=10.0.0.2`) or by using a [command-specific configuration file](examples/edge.yaml).

```bash
% gon2n
Go bindings and CLI for n2n.

Find more information at:
https://pojntfx.github.io/gon2n/

Usage:
  gon2n [command]

Available Commands:
  edge        Start an edge
  help        Help about any command
  supernode   Start a supernode

Flags:
  -h, --help   help for gon2n

Use "gon2n [command] --help" for more information about a command.
```

## License

gon2n (c) 2019 Felix Pojtinger

SPDX-License-Identifier: AGPL-3.0
