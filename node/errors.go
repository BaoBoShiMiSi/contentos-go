package node

import (
	"errors"
	"fmt"
	//"reflect"
	"syscall"
)

var (
	ErrDatadirUsed    = errors.New("datadir already used by another process")
	ErrNodeStopped    = errors.New("node not started")
	ErrNodeRunning    = errors.New("node already running")
	ErrServiceUnknown = errors.New("unknown service")

	datadirInUseErrnos = map[uint]bool{11: true, 32: true, 35: true}
)

type DuplicateServiceError struct {
	//Kind reflect.Type
	Kind string
}

func (e *DuplicateServiceError) Error() string {
	return fmt.Sprintf("duplicate service: %v", e.Kind)
}

func convertFileLockError(err error) error {
	if errno, ok := err.(syscall.Errno); ok && datadirInUseErrnos[uint(errno)] {
		return ErrDatadirUsed
	}
	return err
}

// StopError is returned if a Node fails to stop either any of its registered
// services or itself.
//type StopError struct {
//	Server   error
//	Services map[reflect.Type]error
//}
type StopError struct {
	Server   error
	Services map[string]error
}

// Error generates a textual representation of the stop error.
func (e *StopError) Error() string {
	return fmt.Sprintf("server: %v, services: %v", e.Server, e.Services)
}
