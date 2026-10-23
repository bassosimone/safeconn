//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/rbmk-project/rbmk/blob/v0.17.0/pkg/x/netcore/conn.go
// Adapted from: https://github.com/ooni/probe-cli/blob/v3.20.1/internal/measurexlite/conn.go
//

// Package safeconn provides nil-safe accessors for [net.Conn] properties.
//
// These functions exist because gvisor-based network stacks can cause the
// connection's addresses to disappear very quickly after a connection is closed,
// leading to segfaults when accessing LocalAddr or RemoteAddr. By centralizing
// nil checks here, we defend against this platform-specific behavior.
package safeconn

import "net"

// LocalAddrer allows extracting the local address.
type LocalAddrer interface {
	LocalAddr() net.Addr
}

// LocalAddr returns the conn local address or the empty string.
func LocalAddr(conn LocalAddrer) (value string) {
	if conn != nil {
		if addr := conn.LocalAddr(); addr != nil {
			value = addr.String()
		}
	}
	return
}

// RemoteAddrer allows extracting the remote address.
type RemoteAddrer interface {
	RemoteAddr() net.Addr
}

// RemoteAddr returns the conn remote address or the empty string.
func RemoteAddr(conn RemoteAddrer) (value string) {
	if conn != nil {
		if addr := conn.RemoteAddr(); addr != nil {
			value = addr.String()
		}
	}
	return
}

// Network returns the conn network or the empty string.
func Network(conn LocalAddrer) (value string) {
	if conn != nil {
		if addr := conn.LocalAddr(); addr != nil {
			value = addr.Network()
		}
	}
	return
}
