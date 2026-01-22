//
// SPDX-License-Identifier: GPL-3.0-or-later
//

package safeconn

import (
	"net"
	"testing"

	"github.com/bassosimone/netstub"
	"github.com/stretchr/testify/assert"
)

func TestLocalAddr(t *testing.T) {
	tests := []struct {
		// name describes the test case.
		name string

		// conn is the connection to test.
		conn net.Conn

		// want is the expected result.
		want string
	}{
		{
			name: "nil conn returns empty string",
			conn: nil,
			want: "",
		},

		{
			name: "nil local addr returns empty string",
			conn: &netstub.FuncConn{
				LocalAddrFunc: func() net.Addr { return nil },
			},
			want: "",
		},

		{
			name: "valid local addr returns address string",
			conn: &netstub.FuncConn{
				LocalAddrFunc: func() net.Addr {
					return &netstub.FuncAddr{
						StringFunc: func() string { return "127.0.0.1:8080" },
					}
				},
			},
			want: "127.0.0.1:8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LocalAddr(tt.conn)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRemoteAddr(t *testing.T) {
	tests := []struct {
		// name describes the test case.
		name string

		// conn is the connection to test.
		conn net.Conn

		// want is the expected result.
		want string
	}{
		{
			name: "nil conn returns empty string",
			conn: nil,
			want: "",
		},

		{
			name: "nil remote addr returns empty string",
			conn: &netstub.FuncConn{
				RemoteAddrFunc: func() net.Addr { return nil },
			},
			want: "",
		},

		{
			name: "valid remote addr returns address string",
			conn: &netstub.FuncConn{
				RemoteAddrFunc: func() net.Addr {
					return &netstub.FuncAddr{
						StringFunc: func() string { return "192.168.1.1:443" },
					}
				},
			},
			want: "192.168.1.1:443",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoteAddr(tt.conn)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNetwork(t *testing.T) {
	tests := []struct {
		// name describes the test case.
		name string

		// conn is the connection to test.
		conn net.Conn

		// want is the expected result.
		want string
	}{
		{
			name: "nil conn returns empty string",
			conn: nil,
			want: "",
		},

		{
			name: "nil local addr returns empty string",
			conn: &netstub.FuncConn{
				LocalAddrFunc: func() net.Addr { return nil },
			},
			want: "",
		},

		{
			name: "valid local addr returns network string",
			conn: &netstub.FuncConn{
				LocalAddrFunc: func() net.Addr {
					return &netstub.FuncAddr{
						NetworkFunc: func() string { return "tcp" },
					}
				},
			},
			want: "tcp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Network(tt.conn)
			assert.Equal(t, tt.want, got)
		})
	}
}
