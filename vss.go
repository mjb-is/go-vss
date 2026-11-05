//go:build !windows
// +build !windows

package vss

import (
	"errors"
)

type Snapshotter struct{}

func (v *Snapshotter) CreateSnapshot(drive string, timeout int, opts ...SnapshotterOption) (*Snapshot, error) {
	return nil, errors.ErrUnsupported
}

func (v *Snapshotter) Release() error {
	return errors.ErrUnsupported
}
