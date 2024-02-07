package superadmin

import (
	"context"
	"errors"
	"github.com/uaxe/kuafu/internal/superadmin/cmcc"
)

type Provider interface {
	Name() string
	GetSuperAdmin() (*SuperAdmin, error)
}

var (
	ErrNotSupportProvider = errors.New(`not support provider`)
)

func NewProvider(ctx context.Context, t ProviderType, addr string) (Provider, error) {
	switch t {
	case CMCCProvider:
		return cmcc.NewCMCCProvider(ctx, cmcc.WithAddr(addr))
	default:
	}
	return nil, ErrNotSupportProvider
}
