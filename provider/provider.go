package provider

import (
	"github.com/uaxe/kuafu/provider/modem"
	_ "github.com/uaxe/kuafu/provider/modem/cmcc"
)

type Provider interface {
	Type() string

	SuperAdminProvider(f *modem.AdminFlag) (modem.Provider, error)
}
