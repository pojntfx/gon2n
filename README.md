# gon2n

Go bindings, management daemons and CLIs for n2n edges and supernodes.

[![pipeline status](https://gitlab.com/pojntfx/gon2n/badges/master/pipeline.svg)](https://gitlab.com/pojntfx/gon2n/commits/master)

## Overview

`gon2n` is a collection of Go bindings, management daemons and CLIs for the n2n peer-to-peer VPN. n2n is built of two main components:

- `edge`s, which are the "VPN clients" that manage the TUN/TAP interfaces on every device that is part of a community (a overlay network)
- `supernode`s, which are responsible for both keeping track of the `edge`s of a community as well routing traffic to `edge`s which can't communicate to each other with a peer-to-peer connection

In a similar way, `gon2n` is built of multiple components. The components are:

- `edged`, a n2n edge management daemon with a gRPC interface
- `supernoded`, a n2n supernode management daemon with a gRPC interface
- `edgectl`, a CLI for `edged`
- `supernodectl`, a CLI for `supernoded`

## Installation

### Prebuilt Binaries

Prebuilt binaries are available on the [releases page](https://github.com/pojntfx/gon2n/releases/latest).

### Go Package

A Go package [is available](https://pkg.go.dev/github.com/pojntfx/gon2n).

### Docker Image

#### `supernoded`

A Docker image is available on [Docker Hub](https://hub.docker.com/r/pojntfx/supernoded).

#### `edged`

A Docker image is available on [Docker Hub](https://hub.docker.com/r/pojntfx/edged).

### Helm Chart

Helm charts for `supernoded` and `edged` are available in [@pojntfx's Helm chart repository](https://pojntfx.github.io/charts/).

## Usage

### Daemons

There are two daemons, `supernoded` and `edged`; the latter requires `CAP_NET_ADMIN` capabilities to manage the TUN/TAP interfaces.

#### `supernoded`

You may also set the flags by setting env variables in the format `SUPERNODED_[FLAG]` (i.e. `SUPERNODED_SUPERNODED_CONFIGFILE=examples/supernoded.yaml`) or by using a [configuration file](examples/supernoded.yaml).

```bash
% supernoded --help
supernoded is the n2n supernode management daemon.

Find more information at:
https://pojntfx.github.io/gon2n/

Usage:
  supernoded [flags]

Flags:
  -h, --help                               help for supernoded
  -f, --supernoded.configFile string       Configuration file to use.
  -l, --supernoded.listenHostPort string   TCP listen host:port. (default ":1050")
```

#### `edged`

You may also set the flags by setting env variables in the format `EDGED_[FLAG]` (i.e. `EDGED_EDGED_CONFIGFILE=examples/edged.yaml`) or by using a [configuration file](examples/edged.yaml).

```bash
% edged --help
edged is the n2n edge management daemon.

Find more information at:
https://pojntfx.github.io/gon2n/

Usage:
  edged [flags]

Flags:
  -f, --edged.configFile string       Configuration file to use.
  -l, --edged.listenHostPort string   TCP listen host:port. (default ":1060")
  -h, --help                          help for edged
```

### Client CLIs

There are two client CLIs, `supernodectl` and `edgectl`.

#### `supernodectl`

You may also set the flags by setting env variables in the format `SUPERNODE_[FLAG]` (i.e. `SUPERNODE_SUPERNODE_CONFIGFILE=examples/supernode.yaml`) or by using a [configuration file](examples/supernode.yaml).

```bash
% supernodectl
supernodectl manages supernoded, the n2n supernode management daemon.

Find more information at:
https://pojntfx.github.io/gon2n/

Usage:
  supernodectl [command]

Available Commands:
  apply       Apply a supernode
  delete      Delete one or more supernode(s)
  get         Get one or all supernode(s)
  help        Help about any command

Flags:
  -h, --help   help for supernodectl

Use "supernodectl [command] --help" for more information about a command.
```

#### `edgectl`

You may also set the flags by setting env variables in the format `EDGE_[FLAG]` (i.e. `EDGE_EDGE_CONFIGFILE=examples/edge.yaml`) or by using a [configuration file](examples/edge.yaml) ([alternative with DHCP instead of static IPs](examples/edge-dhcp.yaml)).

```bash
% edgectl
edgectl manages edged, the n2n edge management daemon.

Find more information at:
https://pojntfx.github.io/gon2n/

Usage:
  edgectl [command]

Available Commands:
  apply       Apply an edge
  delete      Delete one or more edge(s)
  get         Get one or all edge(s)
  help        Help about any command

Flags:
  -h, --help   help for edgectl

Use "edgectl [command] --help" for more information about a command.
```

## License

gon2n (c) 2020 Felix Pojtinger

SPDX-License-Identifier: AGPL-3.0
