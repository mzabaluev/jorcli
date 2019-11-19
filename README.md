# jorcli

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/fb345d4e21584e38b13695438c733c79)](https://www.codacy.com/app/rinor/jorcli?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=rinor/jorcli&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/rinor/jorcli)](https://goreportcard.com/report/github.com/rinor/jorcli)
[![GoDoc](https://godoc.org/github.com/rinor/jorcli?status.svg)](https://godoc.org/github.com/rinor/jorcli)
[![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://github.com/emersion/stability-badges#experimental)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Frinor%2Fjorcli.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Frinor%2Fjorcli?ref=badge_shield)

[Jörmungandr](https://github.com/input-output-hk/jormungandr) tools in [Go](https://golang.org/) - (**experimental**)

Right now this is just a *Proof Of Concept* and may or may not become a *Prototype*.

The idea is to build:

- [x] a simple and small wrapper around *jcli* binary. (*beta*)
- [x] a simple and small wrapper around *Jörmungandr node* binary. (*beta*)
- [ ] *Jörmungandr rest* API. (*wip*)
- [ ] *Jörmungandr explorer node* graphql API. (*wip*)
- [ ] *Jörmungandr* grpc client. (*maybe*)

## Installation

Standard `go get`:

```shell
_$ go get github.com/rinor/jorcli
```

## Usage & Examples

For usage and examples see the [Godoc](https://godoc.org/github.com/rinor/jorcli).
Check also [jorcli_examples](https://github.com/rinor/jorcli_examples) repository.

## Status

**DONE** *jcli* :

- [x] **address** - Address tooling and helper
- [x] **certificate** - Certificate generation tool
- [x] **debug** - Debug tools for developers
- [x] **genesis** - Block tooling and helper
- [x] **key** - Key Generation
- [x] **transaction** - Build and view offline transaction
- [x] **utils** - Utilities that perform specialized tasks
- [x] **rest** - Send request to node REST API

**DONE** *Jörmungandr node*:

- [x] node controller
  - [x] Run node
  - [x] Stop node
  - [x] PID of running node
- [x] configs
  - [x] block0 (genesis) config
  - [x] secrets config
  - [x] node config
    - [ ] implement log output gelf (to be implemented)
    - [ ] implement log output file (to be implemented)

**TODO** *Jörmungandr* rest API:

- [ ] build from [openapi](https://github.com/input-output-hk/jormungandr/blob/master/doc/openapi.yaml)

**TODO** *Jörmungandr explorer node* graphql API:

- [ ] build from [explorer](https://github.com/input-output-hk/jormungandr/tree/master/jormungandr/src/rest/explorer)

**TODO** *Jörmungandr* grpc client:

- [ ] build from [node.proto](https://github.com/input-output-hk/chain-libs/blob/master/network-grpc/proto/node.proto)
