package node

import (
	"errors"
	"github.com/coschain/contentos-go/p2p"
	"github.com/coschain/contentos-go/util/flock"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

// Node is a container and manager of services
type Node struct {
	config *Config

	instanceDirLock flock.Releaser // prevent concurrency
	serverConfig    p2p.Config
	server          *p2p.Server // running p2p network

	services     map[reflect.Type]Service // Currently running nodes
	serviceFuncs []ServiceConstructor     // node constructors

	httpEndpoint string // HTTP endpoint(host + port) to listen at

	stop chan struct{}
	lock sync.RWMutex

	log log.Logger
}

func New(conf *Config) (*Node, error) {
	// Copy config
	confCopy := *conf
	conf = &confCopy
	if conf.DataDir != "" {
		dir, err := filepath.Abs(conf.DataDir)
		if err != nil {
			return nil, err
		}
		conf.DataDir = dir
	}
	// Ensure that the instance name doesn't cause weird conflicts with
	// other files in the data directory.
	if strings.ContainsAny(conf.Name, `/\`) {
		return nil, errors.New(`Config.Name must not contain '/' or '\'`)
	}
	return &Node{
		config:       conf,
		httpEndpoint: conf.HTTPEndpoint(),
		log:          conf.Logger,
	}, nil
}

func (n *Node) Register(constructor ServiceConstructor) error {
	n.lock.Lock()
	defer n.lock.Unlock()

	if n.server != nil {
		return ErrNodeRunning
	}
	n.serviceFuncs = append(n.serviceFuncs, constructor)
	return nil
}

func (n *Node) Start() error {
	n.lock.Lock()
	defer n.lock.Unlock()

	if n.server != nil {
		return ErrNodeRunning
	}

	if err := n.openDataDir(); err != nil {
		return err
	}

	running := &p2p.Server{Config: n.serverConfig}
	services := make(map[reflect.Type]Service)
	for _, constructor := range n.serviceFuncs {
		ctx := &ServiceContext{
			config:   n.config,
			services: make(map[reflect.Type]Service),
		}

		// copy
		for kind, n := range services {
			ctx.services[kind] = n
		}

		service, err := constructor(ctx)
		if err != nil {
			return err
		}
		kind := reflect.TypeOf(service)
		if _, exists := services[kind]; exists {
			return &DuplicateServiceError{Kind: kind}
		}
		services[kind] = service

	}

	if err := running.Start(); err != nil {
		return ErrNodeRunning
	}

	// Start each of the services
	started := []reflect.Type{}
	for kind, service := range services {
		if err := service.Start(running); err != nil {
			for _, kind := range started {
				services[kind].Stop()
			}
			running.Stop()

			return err
		}
		started = append(started, kind)
	}

	n.services = services
	n.server = running
	n.stop = make(chan struct{})

	return nil

}

func (n *Node) openDataDir() error {
	if n.config.DataDir == "" {
		return nil
	}

	confdir := filepath.Join(n.config.DataDir, n.config.name())
	if err := os.MkdirAll(confdir, 0700); err != nil {
		return err
	}

	release, _, err := flock.New(filepath.Join(confdir, "LOCK"))
	if err != nil {
		return convertFileLockError(err)
	}
	n.instanceDirLock = release
	return nil
}

func (n *Node) Stop() error {
	n.lock.Lock()
	defer n.lock.Unlock()

	if n.server == nil {
		return ErrNodeStopped
	}

	failure := &StopError{
		Services: make(map[reflect.Type]error),
	}

	for kind, service := range n.services {
		if err := service.Stop(); err != nil {
			failure.Services[kind] = err
		}
	}
	n.server.Stop()
	n.services = nil
	n.server = nil

	if n.instanceDirLock != nil {
		if err := n.instanceDirLock.Release(); err != nil {
			n.log.Fatal("Can't release datadir lock", "err", err)
		}
		n.instanceDirLock = nil
	}

	close(n.stop)

	if len(failure.Services) > 0 {
		return failure
	}

	return nil
}

func (n *Node) Wait() {
	n.lock.RLock()
	if n.server == nil {
		n.lock.RUnlock()
		return
	}

	stop := n.stop
	n.lock.RUnlock()
	<-stop
}

func (n *Node) Restart() error {
	if err := n.Stop(); err != nil {
		return err
	}

	if err := n.Start(); err != nil {
		return err
	}

	return nil
}

func (n *Node) Service(service interface{}) error {
	n.lock.RLock()
	defer n.lock.RUnlock()

	if n.server == nil {
		return ErrNodeStopped
	}

	// Otherwise try to find the service to return
	element := reflect.ValueOf(service).Elem()
	if running, ok := n.services[element.Type()]; ok {
		element.Set(reflect.ValueOf(running))
		return nil
	}
	return ErrServiceUnknown
}