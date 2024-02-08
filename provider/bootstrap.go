package provider

import (
	"context"

	"github.com/uaxe/kuafu/provider/superadmin"
)

func New(ctx context.Context, providerType string) Provider {
	return &_provider{
		ctx:          ctx,
		providerType: providerType,
	}
}

type _provider struct {
	ctx          context.Context
	providerType string
}

func (p *_provider) Type() string {
	return p.providerType
}

func (p *_provider) SuperAdminProvider(f *superadmin.AdminFlag) (superadmin.Provider, error) {
	pp, err := superadmin.Get(func(s superadmin.Provider) bool { return p.Type() == s.Type() })
	if err != nil {
		return nil, err
	}
	return pp.New(p.ctx, f)
}
