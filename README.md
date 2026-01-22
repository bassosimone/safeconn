# Safe Conn

[![GoDoc](https://pkg.go.dev/badge/github.com/bassosimone/safeconn)](https://pkg.go.dev/github.com/bassosimone/safeconn) [![Build Status](https://github.com/bassosimone/safeconn/actions/workflows/go.yml/badge.svg)](https://github.com/bassosimone/safeconn/actions) [![codecov](https://codecov.io/gh/bassosimone/safeconn/branch/main/graph/badge.svg)](https://codecov.io/gh/bassosimone/safeconn)

The `safeconn` Go package provides nil-safe accessors for `net.Conn` properties.

These functions exist because gvisor-based network stacks can cause the
connection's addresses to disappear very quickly after a connection is closed,
leading to segfaults when accessing LocalAddr or RemoteAddr. By centralizing
nil checks here, we defend against this platform-specific behavior.

## Usage

```Go
import "github.com/bassosimone/safeconn"

// Get local address safely (returns "" if conn or addr is nil)
localAddr := safeconn.LocalAddr(conn)

// Get remote address safely
remoteAddr := safeconn.RemoteAddr(conn)

// Get network type safely
network := safeconn.Network(conn)
```

## Installation

```sh
go get github.com/bassosimone/safeconn
```

## Development

```sh
go test -v .
go test -v -cover .
```

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```

## History

Adapted from:
- [rbmk-project/rbmk](https://github.com/rbmk-project/rbmk/blob/v0.17.0/pkg/x/netcore/conn.go)
- [ooni/probe-cli](https://github.com/ooni/probe-cli/blob/v3.20.1/internal/measurexlite/conn.go)
