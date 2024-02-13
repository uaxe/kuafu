package provider

import (
	"github.com/uaxe/kuafu/provider/modem"

	// cmcc driver
	_ "github.com/uaxe/kuafu/provider/modem/cmcc"
	// cucc driver
	_ "github.com/uaxe/kuafu/provider/modem/cucc"
)

type Provider interface {
	Type() string

	SuperAdminProvider(f *modem.AdminFlag) (modem.Provider, error)
}
