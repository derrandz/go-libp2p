package network

import (
	ma "github.com/multiformats/go-multiaddr"
)

// Notifiee is an interface for an object wishing to receive
// notifications from a Network.
type Notifiee interface {
	Listen(Network, ma.Multiaddr)      // called when network starts listening on an addr
	ListenClose(Network, ma.Multiaddr) // called when network stops listening on an addr
	Connected(Network, Conn)           // called when a connection opened
	Disconnected(Network, Conn)        // called when a connection closed
}

// NotifyBundle implements Notifiee by calling any of the functions set on it,
// and nop'ing if they are unset. This is the easy way to register for
// notifications.
type NotifyBundle struct {
	ListenF      func(Network, ma.Multiaddr)
	ListenCloseF func(Network, ma.Multiaddr)

	ConnectedF    func(Network, Conn)
	DisconnectedF func(Network, Conn)
}

var _ Notifiee = (*NotifyBundle)(nil)

// Listen calls ListenF if it is not null.
func (nb *NotifyBundle) Listen(n Network, a ma.Multiaddr) {
	if nb.ListenF != nil {
		nb.ListenF(n, a)
	}
}

// ListenClose calls ListenCloseF if it is not null.
func (nb *NotifyBundle) ListenClose(n Network, a ma.Multiaddr) {
	if nb.ListenCloseF != nil {
		nb.ListenCloseF(n, a)
	}
}

// Connected calls ConnectedF if it is not null.
func (nb *NotifyBundle) Connected(n Network, c Conn) {
	if nb.ConnectedF != nil {
		nb.ConnectedF(n, c)
	}
}

// Disconnected calls DisconnectedF if it is not null.
func (nb *NotifyBundle) Disconnected(n Network, c Conn) {
	if nb.DisconnectedF != nil {
		nb.DisconnectedF(n, c)
	}
}

// Global noop notifiee. Do not change.
var GlobalNoopNotifiee = &NoopNotifiee{}

type NoopNotifiee struct{}

var _ Notifiee = (*NoopNotifiee)(nil)

func (nn *NoopNotifiee) Connected(_ Network, _ Conn)           {}
func (nn *NoopNotifiee) Disconnected(_ Network, _ Conn)        {}
func (nn *NoopNotifiee) Listen(_ Network, _ ma.Multiaddr)      {}
func (nn *NoopNotifiee) ListenClose(_ Network, _ ma.Multiaddr) {}
