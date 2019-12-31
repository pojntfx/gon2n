# gon2n

Go bindings and CLI for n2n.

## Installation

A Go package [is available](https://godoc.org/github.com/pojntfx/gon2n). In order to use it, you have to `go generate` it first.

## Usage

You may also set the flags by setting env variables in the format `GON2N_[FLAG]` (i.e. `GON2N_EDGE_DEVICEIP=10.0.0.2`) or by using a [command-specific configuration file](examples/edge.yaml).

```bash
% gon2n --help
Go bindings and CLI for n2n.

Find more information at:
https://pojntfx.github.io/gon2n/

Usage:
  gon2n [command]

Available Commands:
  create      Create gon2n resources
  help        Help about any command

Flags:
  -h, --help   help for gon2n

Use "gon2n [command] --help" for more information about a command.
```

```bash
% gon2n create supernode --help
Create a supernode

Usage:
  gon2n create supernode [flags]

Flags:
  -h, --help                           help for supernode
  -f, --supernode.configFile string    Configuration file to use.
  -l, --supernode.listenPort int       UDP listen port. (default 1234)
  -m, --supernode.managementPort int   UDP management port. (default 5645)
```

```bash
% gon2n create edge --help
Create an edge

Usage:
  gon2n create edge [flags]

Flags:
  -u, --edge.MTU int                    The MTU to use. (default 1290)
  -z, --edge.addressMode string         Mode of IP address assignment. "static" is a static assignment, "dhcp" uses the DHCP server at --device-ip (see --dynamic-ip-node). If the edge is running the network's DHCP server, this must be "static". (default "static")
  -p, --edge.allowP2P                   Whether to allow peer-to-peer connections. If false, all traffic will be routed through the supernode. (default true)
  -r, --edge.allowRouting               Whether to allow the node to route traffic to other nodes. (default true)
  -c, --edge.communityName string       The name of the n2n community to join. (default "mynetwork")
  -f, --edge.configFile string          Configuration file to use.
  -i, --edge.deviceIP string            IP address to set. Set to "0.0.0.0" if you are using "dhcp" as --address-mode. If the edge is running the network's DHCP server this must be set explicitly; i.e. to "192.168.1.0" if the DHCP server should give out addresses from "192.168.1.10" to "192.168.1.100". (default "10.0.0.1")
  -x, --edge.deviceMACAddress string    The MAC address to use. Must be unique per edge. (default "DE:AD:BE:EF:01:10")
  -v, --edge.deviceName string          Name of the TUN/TAP device to create. (default "edge0")
  -q, --edge.deviceNetmask string       The netmask to use. (default "255.255.255.0")
  -m, --edge.disableMulticast           Whether to disable multicast.
  -d, --edge.disablePMTUDiscovery       Whether to disable path MTU discovery.
  -y, --edge.dynamicIPMode              Whether the IP address is set dynamically (see --address-mode). If the edge is running the network's DHCP server, this must be false.
  -k, --edge.encryptionKey string       The key to use for encryption. (default "mysecretkey")
  -e, --edge.encryptionMethod int       Method of encryption to use. 1 is no encryption, 2 is Twofish encryption, 3 is AES-CBC encryption. Twofish encryption is the n2n default, but only due to legacy compatibility reasons; AES-CBC is up to ten times faster. (default 2)
  -l, --edge.localPort int              The local port to use. 0 uses any available port.
  -a, --edge.managementPort int         UDP management port. 5644 is the n2n default. (default 5644)
  -n, --edge.registerInterval int       Interval in seconds for both UDP NAT hole punching and registration of the edge on the supernode. (default 1)
  -t, --edge.registerTTL int            Interval in seconds for UDP NAT hole punching through the supernode. (default 1)
  -s, --edge.supernodeHostPort string   Host:port of the supernode to connect to. (default "localhost:1234")
  -o, --edge.typeOfService int          Type of service to use. (default 16)
  -h, --help                            help for edge
```

## License

gon2n (c) 2019 Felix Pojtinger

SPDX-License-Identifier: AGPL-3.0
