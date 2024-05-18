// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package tcp represents the imported interface "wasi:sockets/tcp@0.2.0".
package tcp

import (
	monotonicclock "github.com/rajatjindal/wasip2-http-incoming-handler/wasi/clocks/monotonic-clock"
	"github.com/rajatjindal/wasip2-http-incoming-handler/wasi/io/poll"
	"github.com/rajatjindal/wasip2-http-incoming-handler/wasi/io/streams"
	"github.com/rajatjindal/wasip2-http-incoming-handler/wasi/sockets/network"
	"github.com/ydnar/wasm-tools-go/cm"
)

// ShutdownType represents the imported enum "wasi:sockets/tcp@0.2.0#shutdown-type".
//
//	enum shutdown-type {
//		receive,
//		send,
//		both
//	}
type ShutdownType uint8

const (
	// Similar to `SHUT_RD` in POSIX.
	ShutdownTypeReceive ShutdownType = iota

	// Similar to `SHUT_WR` in POSIX.
	ShutdownTypeSend

	// Similar to `SHUT_RDWR` in POSIX.
	ShutdownTypeBoth
)

// TCPSocket represents the imported resource "wasi:sockets/tcp@0.2.0#tcp-socket".
//
// A TCP socket resource.
//
// The socket can be in one of the following states:
// - `unbound`
// - `bind-in-progress`
// - `bound` (See note below)
// - `listen-in-progress`
// - `listening`
// - `connect-in-progress`
// - `connected`
// - `closed`
// See <https://github.com/WebAssembly/wasi-sockets/TcpSocketOperationalSemantics.md>
// for a more information.
//
// Note: Except where explicitly mentioned, whenever this documentation uses
// the term "bound" without backticks it actually means: in the `bound` state *or
// higher*.
// (i.e. `bound`, `listen-in-progress`, `listening`, `connect-in-progress` or `connected`)
//
// In addition to the general error codes documented on the
// `network::error-code` type, TCP socket methods may always return
// `error(invalid-state)` when in the `closed` state.
//
//	resource tcp-socket
type TCPSocket cm.Resource

// ResourceDrop represents the imported resource-drop for resource "tcp-socket".
//
// Drops a resource handle.
//
//go:nosplit
func (self TCPSocket) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [resource-drop]tcp-socket
//go:noescape
func (self TCPSocket) wasmimport_ResourceDrop()

// Accept represents the imported method "accept".
//
// Accept a new client socket.
//
// The returned socket is bound and in the `connected` state. The following properties
// are inherited from the listener socket:
// - `address-family`
// - `keep-alive-enabled`
// - `keep-alive-idle-time`
// - `keep-alive-interval`
// - `keep-alive-count`
// - `hop-limit`
// - `receive-buffer-size`
// - `send-buffer-size`
//
// On success, this function returns the newly accepted client socket along with
// a pair of streams that can be used to read & write to the connection.
//
// # Typical errors
// - `invalid-state`:      Socket is not in the `listening` state. (EINVAL)
// - `would-block`:        No pending connections at the moment. (EWOULDBLOCK, EAGAIN)
// - `connection-aborted`: An incoming connection was pending, but was terminated
// by the client before this listener could accept it. (ECONNABORTED)
// - `new-socket-limit`:   The new socket resource could not be created because of
// a system limit. (EMFILE, ENFILE)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/accept.html>
// - <https://man7.org/linux/man-pages/man2/accept.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-accept>
// - <https://man.freebsd.org/cgi/man.cgi?query=accept&sektion=2>
//
//	accept: func() -> result<tuple<tcp-socket, input-stream, output-stream>, error-code>
//
//go:nosplit
func (self TCPSocket) Accept() cm.OKResult[cm.Tuple3[TCPSocket, streams.InputStream, streams.OutputStream], network.ErrorCode] {
	var result cm.OKResult[cm.Tuple3[TCPSocket, streams.InputStream, streams.OutputStream], network.ErrorCode]
	self.wasmimport_Accept(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.accept
//go:noescape
func (self TCPSocket) wasmimport_Accept(result *cm.OKResult[cm.Tuple3[TCPSocket, streams.InputStream, streams.OutputStream], network.ErrorCode])

// AddressFamily represents the imported method "address-family".
//
// Whether this is a IPv4 or IPv6 socket.
//
// Equivalent to the SO_DOMAIN socket option.
//
//	address-family: func() -> ip-address-family
//
//go:nosplit
func (self TCPSocket) AddressFamily() network.IPAddressFamily {
	return self.wasmimport_AddressFamily()
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.address-family
//go:noescape
func (self TCPSocket) wasmimport_AddressFamily() network.IPAddressFamily

// FinishBind represents the imported method "finish-bind".
//
//	finish-bind: func() -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) FinishBind() cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_FinishBind(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.finish-bind
//go:noescape
func (self TCPSocket) wasmimport_FinishBind(result *cm.ErrResult[struct{}, network.ErrorCode])

// FinishConnect represents the imported method "finish-connect".
//
//	finish-connect: func() -> result<tuple<input-stream, output-stream>, error-code>
//
//go:nosplit
func (self TCPSocket) FinishConnect() cm.OKResult[cm.Tuple[streams.InputStream, streams.OutputStream], network.ErrorCode] {
	var result cm.OKResult[cm.Tuple[streams.InputStream, streams.OutputStream], network.ErrorCode]
	self.wasmimport_FinishConnect(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.finish-connect
//go:noescape
func (self TCPSocket) wasmimport_FinishConnect(result *cm.OKResult[cm.Tuple[streams.InputStream, streams.OutputStream], network.ErrorCode])

// FinishListen represents the imported method "finish-listen".
//
//	finish-listen: func() -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) FinishListen() cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_FinishListen(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.finish-listen
//go:noescape
func (self TCPSocket) wasmimport_FinishListen(result *cm.ErrResult[struct{}, network.ErrorCode])

// HopLimit represents the imported method "hop-limit".
//
// Equivalent to the IP_TTL & IPV6_UNICAST_HOPS socket options.
//
// If the provided value is 0, an `invalid-argument` error is returned.
//
// # Typical errors
// - `invalid-argument`:     (set) The TTL value must be 1 or higher.
//
//	hop-limit: func() -> result<u8, error-code>
//
//go:nosplit
func (self TCPSocket) HopLimit() cm.OKResult[uint8, network.ErrorCode] {
	var result cm.OKResult[uint8, network.ErrorCode]
	self.wasmimport_HopLimit(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.hop-limit
//go:noescape
func (self TCPSocket) wasmimport_HopLimit(result *cm.OKResult[uint8, network.ErrorCode])

// IsListening represents the imported method "is-listening".
//
// Whether the socket is in the `listening` state.
//
// Equivalent to the SO_ACCEPTCONN socket option.
//
//	is-listening: func() -> bool
//
//go:nosplit
func (self TCPSocket) IsListening() bool {
	return self.wasmimport_IsListening()
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.is-listening
//go:noescape
func (self TCPSocket) wasmimport_IsListening() bool

// KeepAliveCount represents the imported method "keep-alive-count".
//
// The maximum amount of keepalive packets TCP should send before aborting the connection.
//
// If the provided value is 0, an `invalid-argument` error is returned.
// Any other value will never cause an error, but it might be silently clamped and/or
// rounded.
// I.e. after setting a value, reading the same setting back may return a different
// value.
//
// Equivalent to the TCP_KEEPCNT socket option.
//
// # Typical errors
// - `invalid-argument`:     (set) The provided value was 0.
//
//	keep-alive-count: func() -> result<u32, error-code>
//
//go:nosplit
func (self TCPSocket) KeepAliveCount() cm.OKResult[uint32, network.ErrorCode] {
	var result cm.OKResult[uint32, network.ErrorCode]
	self.wasmimport_KeepAliveCount(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.keep-alive-count
//go:noescape
func (self TCPSocket) wasmimport_KeepAliveCount(result *cm.OKResult[uint32, network.ErrorCode])

// KeepAliveEnabled represents the imported method "keep-alive-enabled".
//
// Enables or disables keepalive.
//
// The keepalive behavior can be adjusted using:
// - `keep-alive-idle-time`
// - `keep-alive-interval`
// - `keep-alive-count`
// These properties can be configured while `keep-alive-enabled` is false, but only
// come into effect when `keep-alive-enabled` is true.
//
// Equivalent to the SO_KEEPALIVE socket option.
//
//	keep-alive-enabled: func() -> result<bool, error-code>
//
//go:nosplit
func (self TCPSocket) KeepAliveEnabled() cm.OKResult[bool, network.ErrorCode] {
	var result cm.OKResult[bool, network.ErrorCode]
	self.wasmimport_KeepAliveEnabled(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.keep-alive-enabled
//go:noescape
func (self TCPSocket) wasmimport_KeepAliveEnabled(result *cm.OKResult[bool, network.ErrorCode])

// KeepAliveIdleTime represents the imported method "keep-alive-idle-time".
//
// Amount of time the connection has to be idle before TCP starts sending keepalive
// packets.
//
// If the provided value is 0, an `invalid-argument` error is returned.
// Any other value will never cause an error, but it might be silently clamped and/or
// rounded.
// I.e. after setting a value, reading the same setting back may return a different
// value.
//
// Equivalent to the TCP_KEEPIDLE socket option. (TCP_KEEPALIVE on MacOS)
//
// # Typical errors
// - `invalid-argument`:     (set) The provided value was 0.
//
//	keep-alive-idle-time: func() -> result<duration, error-code>
//
//go:nosplit
func (self TCPSocket) KeepAliveIdleTime() cm.OKResult[monotonicclock.Duration, network.ErrorCode] {
	var result cm.OKResult[monotonicclock.Duration, network.ErrorCode]
	self.wasmimport_KeepAliveIdleTime(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.keep-alive-idle-time
//go:noescape
func (self TCPSocket) wasmimport_KeepAliveIdleTime(result *cm.OKResult[monotonicclock.Duration, network.ErrorCode])

// KeepAliveInterval represents the imported method "keep-alive-interval".
//
// The time between keepalive packets.
//
// If the provided value is 0, an `invalid-argument` error is returned.
// Any other value will never cause an error, but it might be silently clamped and/or
// rounded.
// I.e. after setting a value, reading the same setting back may return a different
// value.
//
// Equivalent to the TCP_KEEPINTVL socket option.
//
// # Typical errors
// - `invalid-argument`:     (set) The provided value was 0.
//
//	keep-alive-interval: func() -> result<duration, error-code>
//
//go:nosplit
func (self TCPSocket) KeepAliveInterval() cm.OKResult[monotonicclock.Duration, network.ErrorCode] {
	var result cm.OKResult[monotonicclock.Duration, network.ErrorCode]
	self.wasmimport_KeepAliveInterval(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.keep-alive-interval
//go:noescape
func (self TCPSocket) wasmimport_KeepAliveInterval(result *cm.OKResult[monotonicclock.Duration, network.ErrorCode])

// LocalAddress represents the imported method "local-address".
//
// Get the bound local address.
//
// POSIX mentions:
// > If the socket has not been bound to a local name, the value
// > stored in the object pointed to by `address` is unspecified.
//
// WASI is stricter and requires `local-address` to return `invalid-state` when the
// socket hasn't been bound yet.
//
// # Typical errors
// - `invalid-state`: The socket is not bound to any local address.
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/getsockname.html>
// - <https://man7.org/linux/man-pages/man2/getsockname.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-getsockname>
// - <https://man.freebsd.org/cgi/man.cgi?getsockname>
//
//	local-address: func() -> result<ip-socket-address, error-code>
//
//go:nosplit
func (self TCPSocket) LocalAddress() cm.OKResult[network.IPSocketAddress, network.ErrorCode] {
	var result cm.OKResult[network.IPSocketAddress, network.ErrorCode]
	self.wasmimport_LocalAddress(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.local-address
//go:noescape
func (self TCPSocket) wasmimport_LocalAddress(result *cm.OKResult[network.IPSocketAddress, network.ErrorCode])

// ReceiveBufferSize represents the imported method "receive-buffer-size".
//
// The kernel buffer space reserved for sends/receives on this socket.
//
// If the provided value is 0, an `invalid-argument` error is returned.
// Any other value will never cause an error, but it might be silently clamped and/or
// rounded.
// I.e. after setting a value, reading the same setting back may return a different
// value.
//
// Equivalent to the SO_RCVBUF and SO_SNDBUF socket options.
//
// # Typical errors
// - `invalid-argument`:     (set) The provided value was 0.
//
//	receive-buffer-size: func() -> result<u64, error-code>
//
//go:nosplit
func (self TCPSocket) ReceiveBufferSize() cm.OKResult[uint64, network.ErrorCode] {
	var result cm.OKResult[uint64, network.ErrorCode]
	self.wasmimport_ReceiveBufferSize(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.receive-buffer-size
//go:noescape
func (self TCPSocket) wasmimport_ReceiveBufferSize(result *cm.OKResult[uint64, network.ErrorCode])

// RemoteAddress represents the imported method "remote-address".
//
// Get the remote address.
//
// # Typical errors
// - `invalid-state`: The socket is not connected to a remote address. (ENOTCONN)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/getpeername.html>
// - <https://man7.org/linux/man-pages/man2/getpeername.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-getpeername>
// - <https://man.freebsd.org/cgi/man.cgi?query=getpeername&sektion=2&n=1>
//
//	remote-address: func() -> result<ip-socket-address, error-code>
//
//go:nosplit
func (self TCPSocket) RemoteAddress() cm.OKResult[network.IPSocketAddress, network.ErrorCode] {
	var result cm.OKResult[network.IPSocketAddress, network.ErrorCode]
	self.wasmimport_RemoteAddress(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.remote-address
//go:noescape
func (self TCPSocket) wasmimport_RemoteAddress(result *cm.OKResult[network.IPSocketAddress, network.ErrorCode])

// SendBufferSize represents the imported method "send-buffer-size".
//
//	send-buffer-size: func() -> result<u64, error-code>
//
//go:nosplit
func (self TCPSocket) SendBufferSize() cm.OKResult[uint64, network.ErrorCode] {
	var result cm.OKResult[uint64, network.ErrorCode]
	self.wasmimport_SendBufferSize(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.send-buffer-size
//go:noescape
func (self TCPSocket) wasmimport_SendBufferSize(result *cm.OKResult[uint64, network.ErrorCode])

// SetHopLimit represents the imported method "set-hop-limit".
//
//	set-hop-limit: func(value: u8) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetHopLimit(value uint8) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_SetHopLimit(value, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.set-hop-limit
//go:noescape
func (self TCPSocket) wasmimport_SetHopLimit(value uint8, result *cm.ErrResult[struct{}, network.ErrorCode])

// SetKeepAliveCount represents the imported method "set-keep-alive-count".
//
//	set-keep-alive-count: func(value: u32) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetKeepAliveCount(value uint32) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_SetKeepAliveCount(value, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.set-keep-alive-count
//go:noescape
func (self TCPSocket) wasmimport_SetKeepAliveCount(value uint32, result *cm.ErrResult[struct{}, network.ErrorCode])

// SetKeepAliveEnabled represents the imported method "set-keep-alive-enabled".
//
//	set-keep-alive-enabled: func(value: bool) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetKeepAliveEnabled(value bool) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_SetKeepAliveEnabled(value, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.set-keep-alive-enabled
//go:noescape
func (self TCPSocket) wasmimport_SetKeepAliveEnabled(value bool, result *cm.ErrResult[struct{}, network.ErrorCode])

// SetKeepAliveIdleTime represents the imported method "set-keep-alive-idle-time".
//
//	set-keep-alive-idle-time: func(value: duration) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetKeepAliveIdleTime(value monotonicclock.Duration) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_SetKeepAliveIdleTime(value, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.set-keep-alive-idle-time
//go:noescape
func (self TCPSocket) wasmimport_SetKeepAliveIdleTime(value monotonicclock.Duration, result *cm.ErrResult[struct{}, network.ErrorCode])

// SetKeepAliveInterval represents the imported method "set-keep-alive-interval".
//
//	set-keep-alive-interval: func(value: duration) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetKeepAliveInterval(value monotonicclock.Duration) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_SetKeepAliveInterval(value, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.set-keep-alive-interval
//go:noescape
func (self TCPSocket) wasmimport_SetKeepAliveInterval(value monotonicclock.Duration, result *cm.ErrResult[struct{}, network.ErrorCode])

// SetListenBacklogSize represents the imported method "set-listen-backlog-size".
//
// Hints the desired listen queue size. Implementations are free to ignore this.
//
// If the provided value is 0, an `invalid-argument` error is returned.
// Any other value will never cause an error, but it might be silently clamped and/or
// rounded.
//
// # Typical errors
// - `not-supported`:        (set) The platform does not support changing the backlog
// size after the initial listen.
// - `invalid-argument`:     (set) The provided value was 0.
// - `invalid-state`:        (set) The socket is in the `connect-in-progress` or `connected`
// state.
//
//	set-listen-backlog-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetListenBacklogSize(value uint64) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_SetListenBacklogSize(value, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.set-listen-backlog-size
//go:noescape
func (self TCPSocket) wasmimport_SetListenBacklogSize(value uint64, result *cm.ErrResult[struct{}, network.ErrorCode])

// SetReceiveBufferSize represents the imported method "set-receive-buffer-size".
//
//	set-receive-buffer-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetReceiveBufferSize(value uint64) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_SetReceiveBufferSize(value, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.set-receive-buffer-size
//go:noescape
func (self TCPSocket) wasmimport_SetReceiveBufferSize(value uint64, result *cm.ErrResult[struct{}, network.ErrorCode])

// SetSendBufferSize represents the imported method "set-send-buffer-size".
//
//	set-send-buffer-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetSendBufferSize(value uint64) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_SetSendBufferSize(value, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.set-send-buffer-size
//go:noescape
func (self TCPSocket) wasmimport_SetSendBufferSize(value uint64, result *cm.ErrResult[struct{}, network.ErrorCode])

// Shutdown represents the imported method "shutdown".
//
// Initiate a graceful shutdown.
//
// - `receive`: The socket is not expecting to receive any data from
// the peer. The `input-stream` associated with this socket will be
// closed. Any data still in the receive queue at time of calling
// this method will be discarded.
// - `send`: The socket has no more data to send to the peer. The `output-stream`
// associated with this socket will be closed and a FIN packet will be sent.
// - `both`: Same effect as `receive` & `send` combined.
//
// This function is idempotent. Shutting a down a direction more than once
// has no effect and returns `ok`.
//
// The shutdown function does not close (drop) the socket.
//
// # Typical errors
// - `invalid-state`: The socket is not in the `connected` state. (ENOTCONN)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/shutdown.html>
// - <https://man7.org/linux/man-pages/man2/shutdown.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-shutdown>
// - <https://man.freebsd.org/cgi/man.cgi?query=shutdown&sektion=2>
//
//	shutdown: func(shutdown-type: shutdown-type) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) Shutdown(shutdownType ShutdownType) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_Shutdown(shutdownType, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.shutdown
//go:noescape
func (self TCPSocket) wasmimport_Shutdown(shutdownType ShutdownType, result *cm.ErrResult[struct{}, network.ErrorCode])

// StartBind represents the imported method "start-bind".
//
// Bind the socket to a specific network on the provided IP address and port.
//
// If the IP address is zero (`0.0.0.0` in IPv4, `::` in IPv6), it is left to the
// implementation to decide which
// network interface(s) to bind to.
// If the TCP/UDP port is zero, the socket will be bound to a random free port.
//
// Bind can be attempted multiple times on the same socket, even with
// different arguments on each iteration. But never concurrently and
// only as long as the previous bind failed. Once a bind succeeds, the
// binding can't be changed anymore.
//
// # Typical errors
// - `invalid-argument`:          The `local-address` has the wrong address family.
// (EAFNOSUPPORT, EFAULT on Windows)
// - `invalid-argument`:          `local-address` is not a unicast address. (EINVAL)
// - `invalid-argument`:          `local-address` is an IPv4-mapped IPv6 address.
// (EINVAL)
// - `invalid-state`:             The socket is already bound. (EINVAL)
// - `address-in-use`:            No ephemeral ports available. (EADDRINUSE, ENOBUFS
// on Windows)
// - `address-in-use`:            Address is already in use. (EADDRINUSE)
// - `address-not-bindable`:      `local-address` is not an address that the `network`
// can bind to. (EADDRNOTAVAIL)
// - `not-in-progress`:           A `bind` operation is not in progress.
// - `would-block`:               Can't finish the operation, it is still in progress.
// (EWOULDBLOCK, EAGAIN)
//
// # Implementors note
// When binding to a non-zero port, this bind operation shouldn't be affected by the
// TIME_WAIT
// state of a recently closed socket on the same local address. In practice this means
// that the SO_REUSEADDR
// socket option should be set implicitly on all platforms, except on Windows where
// this is the default behavior
// and SO_REUSEADDR performs something different entirely.
//
// Unlike in POSIX, in WASI the bind operation is async. This enables
// interactive WASI hosts to inject permission prompts. Runtimes that
// don't want to make use of this ability can simply call the native
// `bind` as part of either `start-bind` or `finish-bind`.
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/bind.html>
// - <https://man7.org/linux/man-pages/man2/bind.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-bind>
// - <https://man.freebsd.org/cgi/man.cgi?query=bind&sektion=2&format=html>
//
//	start-bind: func(network: borrow<network>, local-address: ip-socket-address) ->
//	result<_, error-code>
//
//go:nosplit
func (self TCPSocket) StartBind(network_ network.Network, localAddress network.IPSocketAddress) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_StartBind(network_, localAddress, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.start-bind
//go:noescape
func (self TCPSocket) wasmimport_StartBind(network_ network.Network, localAddress network.IPSocketAddress, result *cm.ErrResult[struct{}, network.ErrorCode])

// StartConnect represents the imported method "start-connect".
//
// Connect to a remote endpoint.
//
// On success:
// - the socket is transitioned into the `connection` state.
// - a pair of streams is returned that can be used to read & write to the connection
//
// After a failed connection attempt, the socket will be in the `closed`
// state and the only valid action left is to `drop` the socket. A single
// socket can not be used to connect more than once.
//
// # Typical errors
// - `invalid-argument`:          The `remote-address` has the wrong address family.
// (EAFNOSUPPORT)
// - `invalid-argument`:          `remote-address` is not a unicast address. (EINVAL,
// ENETUNREACH on Linux, EAFNOSUPPORT on MacOS)
// - `invalid-argument`:          `remote-address` is an IPv4-mapped IPv6 address.
// (EINVAL, EADDRNOTAVAIL on Illumos)
// - `invalid-argument`:          The IP address in `remote-address` is set to INADDR_ANY
// (`0.0.0.0` / `::`). (EADDRNOTAVAIL on Windows)
// - `invalid-argument`:          The port in `remote-address` is set to 0. (EADDRNOTAVAIL
// on Windows)
// - `invalid-argument`:          The socket is already attached to a different network.
// The `network` passed to `connect` must be identical to the one passed to `bind`.
// - `invalid-state`:             The socket is already in the `connected` state.
// (EISCONN)
// - `invalid-state`:             The socket is already in the `listening` state.
// (EOPNOTSUPP, EINVAL on Windows)
// - `timeout`:                   Connection timed out. (ETIMEDOUT)
// - `connection-refused`:        The connection was forcefully rejected. (ECONNREFUSED)
// - `connection-reset`:          The connection was reset. (ECONNRESET)
// - `connection-aborted`:        The connection was aborted. (ECONNABORTED)
// - `remote-unreachable`:        The remote address is not reachable. (EHOSTUNREACH,
// EHOSTDOWN, ENETUNREACH, ENETDOWN, ENONET)
// - `address-in-use`:            Tried to perform an implicit bind, but there were
// no ephemeral ports available. (EADDRINUSE, EADDRNOTAVAIL on Linux, EAGAIN on BSD)
// - `not-in-progress`:           A connect operation is not in progress.
// - `would-block`:               Can't finish the operation, it is still in progress.
// (EWOULDBLOCK, EAGAIN)
//
// # Implementors note
// The POSIX equivalent of `start-connect` is the regular `connect` syscall.
// Because all WASI sockets are non-blocking this is expected to return
// EINPROGRESS, which should be translated to `ok()` in WASI.
//
// The POSIX equivalent of `finish-connect` is a `poll` for event `POLLOUT`
// with a timeout of 0 on the socket descriptor. Followed by a check for
// the `SO_ERROR` socket option, in case the poll signaled readiness.
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/connect.html>
// - <https://man7.org/linux/man-pages/man2/connect.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-connect>
// - <https://man.freebsd.org/cgi/man.cgi?connect>
//
//	start-connect: func(network: borrow<network>, remote-address: ip-socket-address)
//	-> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) StartConnect(network_ network.Network, remoteAddress network.IPSocketAddress) cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_StartConnect(network_, remoteAddress, &result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.start-connect
//go:noescape
func (self TCPSocket) wasmimport_StartConnect(network_ network.Network, remoteAddress network.IPSocketAddress, result *cm.ErrResult[struct{}, network.ErrorCode])

// StartListen represents the imported method "start-listen".
//
// Start listening for new connections.
//
// Transitions the socket into the `listening` state.
//
// Unlike POSIX, the socket must already be explicitly bound.
//
// # Typical errors
// - `invalid-state`:             The socket is not bound to any local address. (EDESTADDRREQ)
// - `invalid-state`:             The socket is already in the `connected` state.
// (EISCONN, EINVAL on BSD)
// - `invalid-state`:             The socket is already in the `listening` state.
// - `address-in-use`:            Tried to perform an implicit bind, but there were
// no ephemeral ports available. (EADDRINUSE)
// - `not-in-progress`:           A listen operation is not in progress.
// - `would-block`:               Can't finish the operation, it is still in progress.
// (EWOULDBLOCK, EAGAIN)
//
// # Implementors note
// Unlike in POSIX, in WASI the listen operation is async. This enables
// interactive WASI hosts to inject permission prompts. Runtimes that
// don't want to make use of this ability can simply call the native
// `listen` as part of either `start-listen` or `finish-listen`.
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/listen.html>
// - <https://man7.org/linux/man-pages/man2/listen.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-listen>
// - <https://man.freebsd.org/cgi/man.cgi?query=listen&sektion=2>
//
//	start-listen: func() -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) StartListen() cm.ErrResult[struct{}, network.ErrorCode] {
	var result cm.ErrResult[struct{}, network.ErrorCode]
	self.wasmimport_StartListen(&result)
	return result
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.start-listen
//go:noescape
func (self TCPSocket) wasmimport_StartListen(result *cm.ErrResult[struct{}, network.ErrorCode])

// Subscribe represents the imported method "subscribe".
//
// Create a `pollable` which can be used to poll for, or block on,
// completion of any of the asynchronous operations of this socket.
//
// When `finish-bind`, `finish-listen`, `finish-connect` or `accept`
// return `error(would-block)`, this pollable can be used to wait for
// their success or failure, after which the method can be retried.
//
// The pollable is not limited to the async operation that happens to be
// in progress at the time of calling `subscribe` (if any). Theoretically,
// `subscribe` only has to be called once per socket and can then be
// (re)used for the remainder of the socket's lifetime.
//
// See <https://github.com/WebAssembly/wasi-sockets/TcpSocketOperationalSemantics.md#Pollable-readiness>
// for a more information.
//
// Note: this function is here for WASI Preview2 only.
// It's planned to be removed when `future` is natively supported in Preview3.
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self TCPSocket) Subscribe() poll.Pollable {
	return self.wasmimport_Subscribe()
}

//go:wasmimport wasi:sockets/tcp@0.2.0 [method]tcp-socket.subscribe
//go:noescape
func (self TCPSocket) wasmimport_Subscribe() poll.Pollable
