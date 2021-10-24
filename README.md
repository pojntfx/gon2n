# gon2n

Go bindings, management daemons and CLIs for n2n edges and supernodes.

[![hydrun CI](https://github.com/pojntfx/gon2n/actions/workflows/hydrun.yaml/badge.svg)](https://github.com/pojntfx/gon2n/actions/workflows/hydrun.yaml)
[![Matrix](https://img.shields.io/matrix/gon2n:matrix.org)](https://matrix.to/#/#gon2n:matrix.org?via=matrix.org)
[![Binary Downloads](https://img.shields.io/github/downloads/pojntfx/gon2n/total?label=binary%20downloads)](https://github.com/pojntfx/gon2n/releases)

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

Static binaries are available on [GitHub releases](https://github.com/pojntfx/gon2n/releases).

On Linux, you can install them like so:

```shell
$ curl -L -o /tmp/supernoded "https://github.com/pojntfx/gon2n/releases/latest/download/supernoded.linux-$(uname -m)"
$ curl -L -o /tmp/edged "https://github.com/pojntfx/gon2n/releases/latest/download/edged.linux-$(uname -m)"
$ curl -L -o /tmp/supernodectl "https://github.com/pojntfx/gon2n/releases/latest/download/supernodectl.linux-$(uname -m)"
$ curl -L -o /tmp/edgectl "https://github.com/pojntfx/gon2n/releases/latest/download/edgectl.linux-$(uname -m)"
$ sudo install /tmp/{supernoded,edged,supernodectl,edgectl} /usr/local/bin
$ sudo setcap cap_net_admin+ep /usr/local/bin/edged # This allows rootless execution
```

On macOS, you can use the following (macOS does not support TAP devices, so only the client CLIs work):

```shell
$ curl -L -o /tmp/supernodectl "https://github.com/pojntfx/gon2n/releases/latest/download/supernodectl.linux-$(uname -m)"
$ curl -L -o /tmp/edgectl "https://github.com/pojntfx/gon2n/releases/latest/download/edgectl.linux-$(uname -m)"
$ sudo install /tmp/{supernodectl,edgectl} /usr/local/bin
```

On Windows, the following should work (using PowerShell as administrator; Windows does not support TAP devices, so only the client CLIs work):

```shell
PS> Invoke-WebRequest https://github.com/pojntfx/gon2n/releases/latest/download/supernodectl.windows-x86_64.exe -OutFile \Windows\System32\supernodectl.exe
PS> Invoke-WebRequest https://github.com/pojntfx/gon2n/releases/latest/download/edgectl.windows-x86_64.exe -OutFile \Windows\System32\edgectl.exe
```

You can find binaries for more operating systems and architectures on [GitHub releases](https://github.com/pojntfx/gon2n/releases).

## Usage

### 1. Setting up the Supernode

First, start the supernode management daemon:

```shell
$ supernoded -f examples/supernoded.yaml
{"level":"info","timestamp":"2021-10-24T20:32:21Z","message":"Starting server"}
```

Now, in a new terminal, create a supernode:

```shell
$ supernodectl apply -f examples/supernode.yaml
supernode "69b92323-9384-46fb-8814-663ea3ff98fe" created
```

You can retrieve the running supernodes with `supernodectl get`:

```shell
$ supernodectl get
ID                                      LISTEN PORT
69b92323-9384-46fb-8814-663ea3ff98fe    1234
```

### 2. Setting up the Edges

In this example, we'll be creating two edges, which will both run on the same host - in a real-world scenario, you probably want to run an edge management daemon per host. First, start the edge management daemon:

```shell
$ edged examples/edged.yaml
{"level":"info","timestamp":"2021-10-24T20:38:34Z","message":"Starting server"}
```

Now, in a new terminal, create the edges:

```shell
$ edgectl apply -f examples/edge-1.yaml
edge "74125834-6ad7-4c90-8c28-f22edda10dad" created
$ edgectl apply -f examples/edge-2.yaml
edge "7f1c60ed-e298-47a1-abe7-cefa61e4d886" created
```

You can retrieve the running supernodes with `edgectl get`:

```shell
$ edgectl get
ID                                      COMMUNITY NAME  LOCAL PORT      SUPERNODE HOST:PORT ENCRYPTION METHOD       DEVICE NAME
74125834-6ad7-4c90-8c28-f22edda10dad    mynetwork       0               localhost:1234      2                       edge0
7f1c60ed-e298-47a1-abe7-cefa61e4d886    mynetwork       0               localhost:1234      2                       edge1
```

The log of the edge management daemon started earlier also shows the changes:

```shell
24/Oct/2021 22:41:07 [edge_utils.c:2578] Adding supernode[0] = localhost:1234
{"level":"info","timestamp":"2021-10-24T20:41:07Z","message":"Starting edge"}
24/Oct/2021 22:41:07 [edge_utils.c:211] supernode 0 => localhost:1234
24/Oct/2021 22:41:07 [edge_utils.c:2104] TOS set to 0x10
24/Oct/2021 22:41:07 [edge_utils.c:727] Successfully joined multicast group 224.0.0.68:1968
24/Oct/2021 22:41:07 [edge_utils.c:1803] [OK] Edge Peer <<< ================ >>> Super Node
24/Oct/2021 22:41:12 [edge_utils.c:2578] Adding supernode[0] = localhost:1234
{"level":"info","timestamp":"2021-10-24T20:41:12Z","message":"Starting edge"}
24/Oct/2021 22:41:12 [edge_utils.c:211] supernode 0 => localhost:1234
24/Oct/2021 22:41:12 [edge_utils.c:2104] TOS set to 0x10
24/Oct/2021 22:41:12 [edge_utils.c:727] Successfully joined multicast group 224.0.0.68:1968
24/Oct/2021 22:41:12 [edge_utils.c:1756] [P2P] Rx REGISTER from 100.64.154.252:53854
24/Oct/2021 22:41:12 [edge_utils.c:1756] [P2P] Rx REGISTER from 100.64.154.252:34474
24/Oct/2021 22:41:12 [edge_utils.c:1756] [P2P] Rx REGISTER from 100.64.154.252:34474
# ...
```

If we check with `ip a`, we can also see the created `edge0` and `edge1` interfaces:

```shell
$ ip a
# ...
553: edge0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether de:ad:be:ef:01:10 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.1/24 brd 192.168.1.255 scope global edge0
       valid_lft forever preferred_lft forever
    inet6 fe80::dcad:beff:feef:110/64 scope link
       valid_lft forever preferred_lft forever
554: edge1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether de:a0:be:ef:01:10 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.2/24 brd 192.168.1.255 scope global edge1
       valid_lft forever preferred_lft forever
    inet6 fe80::dca0:beff:feef:110/64 scope link
       valid_lft forever preferred_lft forever
# ...
```

🚀 **That's it!** We've successfully created a layer 2 overlay network.

Be sure to check out the [reference](#reference) for more information.

## Reference

### Daemons

There are two daemons, `supernoded` and `edged`; the latter requires `CAP_NET_ADMIN` capabilities to manage the TUN/TAP interfaces.

#### `supernoded`

You may also set the flags by setting env variables in the format `SUPERNODED_[FLAG]` (i.e. `SUPERNODED_SUPERNODED_CONFIGFILE=examples/supernoded.yaml`) or by using a [configuration file](examples/supernoded.yaml).

```bash
$ supernoded --help
supernoded is the n2n supernode management daemon.

Find more information at:
https://github.com/pojntfx/gon2n

Usage:
  supernoded [flags]

Flags:
  -h, --help                               help for supernoded
  -f, --supernoded.configFile string       Configuration file to use.
  -l, --supernoded.listenHostPort string   TCP listen host:port. (default "localhost:1050")
```

#### `edged`

You may also set the flags by setting env variables in the format `EDGED_[FLAG]` (i.e. `EDGED_EDGED_CONFIGFILE=examples/edged.yaml`) or by using a [configuration file](examples/edged.yaml).

```bash
$ edged --help
edged is the n2n edge management daemon.

Find more information at:
https://github.com/pojntfx/gon2n

Usage:
  edged [flags]

Flags:
  -f, --edged.configFile string       Configuration file to use.
  -l, --edged.listenHostPort string   TCP listen host:port. (default "localhost:1060")
  -h, --help                          help for edged
```

### Client CLIs

There are two client CLIs, `supernodectl` and `edgectl`.

#### `supernodectl`

You may also set the flags by setting env variables in the format `SUPERNODE_[FLAG]` (i.e. `SUPERNODE_SUPERNODE_CONFIGFILE=examples/supernode.yaml`) or by using a [configuration file](examples/supernode.yaml).

```bash
$ supernodectl --help
supernodectl manages supernoded, the n2n supernode management daemon.

Find more information at:
https://github.com/pojntfx/gon2n

Usage:
  supernodectl [command]

Available Commands:
  apply       Apply a supernode
  completion  generate the autocompletion script for the specified shell
  delete      Delete one or more supernode(s)
  get         Get one or all supernode(s)
  help        Help about any command

Flags:
  -h, --help   help for supernodectl

Use "supernodectl [command] --help" for more information about a command.
```

#### `edgectl`

You may also set the flags by setting env variables in the format `EDGE_[FLAG]` (i.e. `EDGE_EDGE_CONFIGFILE=examples/edge-1.yaml`) or by using a [configuration file](examples/edge-1.yaml) ([alternative with DHCP instead of static IPs](examples/edge-dhcp.yaml)).

```bash
$ edgectl --help
edgectl manages edged, the n2n edge management daemon.

Find more information at:
https://github.com/pojntfx/gon2n

Usage:
  edgectl [command]

Available Commands:
  apply       Apply an edge
  completion  generate the autocompletion script for the specified shell
  delete      Delete one or more edge(s)
  get         Get one or all edge(s)
  help        Help about any command

Flags:
  -h, --help   help for edgectl

Use "edgectl [command] --help" for more information about a command.
```

## Acknowledgements

- This project would not have been possible were it not for [@ntop](https://github.com/ntop)'s [n2n](https://github.com/ntop/n2n) VPN; be sure to check it out too!
- All the rest of the authors who worked on the dependencies used! Thanks a lot!

## Contributing

To contribute, please use the [GitHub flow](https://guides.github.com/introduction/flow/) and follow our [Code of Conduct](./CODE_OF_CONDUCT.md).

To build gon2n locally, run:

```shell
$ git clone https://github.com/pojntfx/gon2n.git
$ cd gon2n
$ make depend
$ make
$ sudo make run/edged # Or run/supernoded etc.
```

Have any questions or need help? Chat with us [on Matrix](https://matrix.to/#/#gon2n:matrix.org?via=matrix.org)!

## License

gon2n (c) 2021 Felix Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
