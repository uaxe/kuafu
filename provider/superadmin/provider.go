package superadmin

import (
	"context"
	"github.com/uaxe/infra/hook"
)

type Provider interface {
	Type() string

	New(context.Context, *AdminFlag) (Provider, error)

	GetSuperAdmin() (*SuperAdmin, error)
}

var (
	defaultHook = &hook.IHook[Provider]{}
	Register    = defaultHook.Register
	Get         = defaultHook.Get
)
