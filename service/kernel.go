package service

import (
	"context"

	"github.com/leliuga/cdk/types"
)

// NewKernel returns a new kernel.
func NewKernel() *Kernel {
	return &Kernel{
		Instances: types.Map[any]{},
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

// Register a new instance to the kernel.
func (k *Kernel) Register(key string, instance any) {
	k.Instances.Set(key, instance)
}

// Get an instance from the kernel.
func (k *Kernel) Get(key string) any {
	return k.Instances.Get(key)
}

// Has an instance from the kernel.
func (k *Kernel) Has(key string) bool {
	return k.Instances.Has(key)
}
