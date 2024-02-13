package modem

import (
	"context"

	"github.com/uaxe/infra/hook"
	"golang.org/x/net/html"
)

var (
	defaultHook = &hook.IHook[Provider]{}
	Register    = defaultHook.Register
	Get         = defaultHook.Get
)

type Provider interface {
	Type() string

	IsMe(*html.Node) bool

	New(context.Context, *AdminFlag) (Provider, error)

	GetSuperAdmin() (*SuperAdmin, error)
}
