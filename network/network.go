package network

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanche-network-runner/network/node"
)

var ErrUndefined = errors.New("undefined network")
var ErrStopped = errors.New("network stopped")

// Network is an abstraction of an Avalanche network
type Network interface {
	// Initializes and starts the network using the given config
	// To be executed after network creation. Enables the other calls.
	LoadConfig(context.Context, Config) error
	// Initializes and starts network using the given snapshot
	// To be executed after network creation. Enables the other calls.
	LoadSnapshot(context.Context, string) error
	// Returns nil if all the nodes in the network are healthy.
	// A stopped network is considered unhealthy.
	// Timeout is given by the context parameter.
	Healthy(context.Context) error
	// Stop all the nodes.
	// Returns ErrStopped if Stop() was previously called.
	Stop(context.Context) error
	// Start a new node with the given config.
	// Returns ErrStopped if Stop() was previously called.
	AddNode(node.Config) (node.Node, error)
	// Stop the node with this name.
	// Returns ErrStopped if Stop() was previously called.
	RemoveNode(name string) error
	// Return the node with this name.
	// Returns ErrStopped if Stop() was previously called.
	GetNode(name string) (node.Node, error)
	// Return all the nodes in this network.
	// Node name --> Node.
	// Returns ErrStopped if Stop() was previously called.
	GetAllNodes() (map[string]node.Node, error)
	// Returns the names of all nodes in this network.
	// Returns ErrStopped if Stop() was previously called.
	GetNodeNames() ([]string, error)
	// Save network snapshot
	// Network is stopped in order to do a safe preservation
	SaveSnapshot(context.Context, string) error
	// Remove network snapshot
	RemoveSnapshot(string) error
	// Get name of available snapshots
	GetSnapshotNames() ([]string, error)
}
