package provider

import (
	"github.com/uaxe/kuafu/provider/superadmin"

	_ "github.com/uaxe/kuafu/provider/superadmin/cmcc"
)

type Provider interface {
	Type() string

	SuperAdminProvider(f *superadmin.AdminFlag) (superadmin.Provider, error)
}
