package provider

import (
	"context"
	"github.com/uaxe/kuafu/provider/modem"
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

func (p *_provider) SuperAdminProvider(f *modem.AdminFlag) (modem.Provider, error) {
	pp, err := modem.Get(func(s modem.Provider) bool { return p.Type() == s.Type() })
	if err != nil {
		return nil, err
	}
	return pp.New(p.ctx, f)
}
