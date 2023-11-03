package service

import (
	"context"

	"github.com/leliuga/cdk/types"
)

// NewKernel returns a new kernel.
func NewKernel() *Kernel {
	return &Kernel{
		instances: types.Map[any]{},
	}
}

// Boot the kernel.
func (k *Kernel) Boot(*Service) error {
	return nil
}

// Shutdown the kernel.
func (k *Kernel) Shutdown(context.Context) error {
	return nil
}

// Set an instance to the kernel.
func (k *Kernel) Set(key string, instance any) {
	k.instances.Set(key, instance)
}

// Get an instance from the kernel.
func (k *Kernel) Get(key string) any {
	return k.instances.Get(key)
}

// Has an instance from the kernel.
func (k *Kernel) Has(key string) bool {
	return k.instances.Has(key)
}

// Instances returns all instances from the kernel.
func (k *Kernel) Instances() types.Map[any] {
	return k.instances
}
