package cmcc

import (
	"context"
	"github.com/uaxe/kuafu/internal/superadmin"
)

var _ superadmin.Provider = (*CMCCProvider)(nil)

type CMCCProvider struct {
	*option
}

func NewCMCCProvider(ctx context.Context, opts ...OptionFunc) (*CMCCProvider, error) {
	c := &CMCCProvider{option: defaultOption(ctx)}
	for k := range opts {
		if err := opts[k](c.option); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (s *CMCCProvider) Name() string {
	return s.name
}

func (s *CMCCProvider) GetSuperAdmin() (ret *superadmin.SuperAdmin, e error) {

	return
}
