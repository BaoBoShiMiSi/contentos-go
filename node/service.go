package node

type ServiceContext struct {
	config   *Config
	services map[string]Service
}

func (ctx *ServiceContext) ResolvePath(path string) string {
	return ctx.config.ResolvePath(path)
}

func (ctx *ServiceContext) Service(name string) (interface{}, error) {
	if running, ok := ctx.services[name]; ok {
		return running, nil
	}
	return nil, ErrServiceUnknown
}

func (ctx *ServiceContext) Config() Config {
	return *ctx.config
}

type ServiceConstructor func(ctx *ServiceContext) (Service, error)

type Service interface {
	Start(node *Node) error

	// stop all goroutines belonging to the service,
	// blocking until all of them are terminated.
	Stop() error
}

func (ctx *ServiceContext) ResetConfig(cfg *Config) {
	ctx.config = cfg
}

func (ctx *ServiceContext) ResetServices(s map[string]Service) {
	ctx.services = s
}