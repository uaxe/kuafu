package superadmin

import (
	"context"
	"github.com/uaxe/infra/hook"
)

type Provider interface {
	Type() string

	New(context.Context, *AdminFlag) error

	GetSuperAdmin() (*SuperAdmin, error)
}

var (
	defaultHook = &hook.IHook[Provider]{}
	Register    = defaultHook.Register
	Get         = defaultHook.Get
)
