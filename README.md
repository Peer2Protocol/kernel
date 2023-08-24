# peer2

[![Release](https://img.shields.io/github/release/peer2/peer2.svg)](https://github.com/peer2/peer2/releases)
[![License](https://img.shields.io/github/license/peer2/peer2)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/peer2/peer2)](https://goreportcard.com/report/peer2/peer2)
[![GoDoc](https://godoc.org/github.com/peer2/peer2?status.svg)](https://pkg.go.dev/github.com/peer2/peer2)
[![Discord](https://img.shields.io/discord/973677117722202152?color=%235865f2&label=discord)](https://discord.gg/peer2)

`peer2` is powered by the upstream project by Taubyte. When interconnected, these nodes form an overlay network offering features such as:
 - Serverless WebAssembly Functions
 - Serverless WebAssemvly Large Language Models 
 - Monolithic Application Hosting
 - S3 Storage
 - Distributed KV Database
 - Pub-Sub Messaging/Events


## Getting started

### Requirements

#### Hardware
The requirements for `peer2` vary depending on the shape, but at a minimum, assuming an empty shape, it requires:

- 512MB of RAM
- 2GB of storage space

#### Operating system
While `peer2` can be cross-compiled to various operating systems, for a quick and seamless deployment we recommend using a Linux distribution that employs SystemD.


#### Network requirements
Depending on the enabled protocols, `peer2` will require certain network ports to be available and open:

- Three configurable TCP ports for peer-to-peer (P2P) communication
- Ports 80 and 443 for HTTP and HTTPS, respectively
- Ports 53 and 953 for DNS, supporting both TCP and UDP


#### Freeing DNS Ports
DNS needs to be available for seer protocol to start properly. Follow these steps to adjust DNS settings:

1. Open `/etc/systemd/resolved.conf` with a text editor (nano, vim)
2. Set `DNS=1.1.1.1` and `DNSStubListener=no`
3. Run `systemctl restart systemd-resolved.service`
4. Run `ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf`


### Installation

#### From source
```bash
$ go install github.com/peer2/peer2
```

#### From binary
To download the latest release run:
```
$ curl https://get.peer2.link/peer2 | sh
```

To install
```
$ curl https://get.peer2.link/peer2 | sh -s -- -i
```

Other options:
 - `-O <path/to/download/folder>` if not specefied a temporrary folder is used
 - `-r <path/to/root>` by default `/tb`
 
### Filesystem Structure
We at peer2 prioritize convention over configuration. Hence, we've pre-defined the filesystem structure and its location as follows:
```
/tb
├── bin
│   └── peer2
├── cache
├── config
│   ├── <shape-1>.yaml
│   ├── <shape-2>.yaml
│   └── keys
│       ├── private.key
│       ├── public.pem
│       └── swarm.key
├── logs
├── plugins
└── storage
```

> Note: If you prefer a different location, use the --root option.

### Configuration
Configuration files for `peer2` are located at `/tb/config/shape-name.yaml`. Here's an example:

```yaml
privatekey: CAESQJxQzCe/N/C8A5TIgrL9F0p5iG...KzYW9pygBCTJSuezIc6w/TT/unZKJ5mo=
swarmkey: keys/test_swarm.key
protocols: [patrick,substrate,tns,monkey,seer,auth]
p2p-listen: [/ip4/0.0.0.0/tcp/8100]
p2p-announce: [/ip4/127.0.0.1/tcp/8100]
ports:
  main: 8100
  lite: 8102
  ipfs: 8104
location:
  lat: 120
  long: 21
network-url: example.com
domains:
  key:
    private: keys/test.key
  services: ^[^.]+\.peer2\.example\.com$
  generated: g\.example\.com$
```

### Running `peer2`
Execute a `peer2` node with:
```bash
peer2 start --shape shape-name
```
For an alternative root to `/tb`:
```bash
$ peer2 start --shape shape-name --root path-to-root
```

### Systemd Configuration
To ensure that `peer2` runs as a service and starts automatically upon system boot, you can set it up as a `systemd` service. 

1. Create a new service file:
```bash
$ sudo nano /etc/systemd/system/peer2.service
```

2. Add the following content to the file:
```plaintext
[Unit]
Description=peer2 Node Service
After=network.target

[Service]
ExecStart=/path/to/peer2/bin peer2 start --shape shape-name --root path-to-root
User=username
Restart=on-failure

[Install]
WantedBy=multi-user.target
```
Replace `/path/to/peer2/bin` with the actual path to your `peer2` binary and `username` with the name of the user running `peer2`.

3. Enable and start the service:
```bash
$ sudo systemctl enable peer2
$ sudo systemctl start peer2
```

To check the status:
```bash
$ sudo systemctl status peer2
```

This ensures `peer2` runs consistently, even after system reboots.

### Setting up DNS for your network
Next, set up DNS records for your seer, generated domain, and service URLs, all with a 1-minute TTL.

### Seer
For each host running seer, add an `A Record`.   
Host: seer -> Value: 127.0.0.1     

### Generated URL
Point your generated domain to your seers by adding a `NS Record`.   
If `g.example.com` is your generated domain URL, the record would be:    
Host: `g` -> Value: `seer.example.com`

### Service URL
Add a `NS Record` which by default should be your URL prefixed with peer2.   
This record will also point to your seers.   
For `example.com` the record will be: `peer2` -> `seer.example.com`

### Connecting to the network
It may take a few minutes for the DNS to recognize your changes. The best way to check is to perform a dig on your network's seer fqdn Like so:
```bash
$ dig a seer.peer2.example.com
```
> replace `example.com` with your respective domain.

Then, on a client machine:

#### peer2-cli
1. Get the `peer2` cli if you don't have it already. Check [github.com/peer2/peer2-cli](https://github.com/peer2/peer2-cli).
2. Run `peer2 login`
3. Then `peer2 select network`
4. Choose `Remote` then type your network's domain

#### Web console
1. Go to [peer2 Web Console](https://console.peer2.com) 
2. Fill in your email and select 'Custom' in the network selector
3. Enter the network's domain and hit the checkmark
4. Login
