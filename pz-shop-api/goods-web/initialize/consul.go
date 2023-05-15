package initialize

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {
		Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	
}

