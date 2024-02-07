package flags

import "github.com/uaxe/kuafu/internal/superadmin"

type Admin struct {
	Provider string `yaml:"provider" name:"provider" description:"device provider"`
	Addr     string `yaml:"addr" name:"addr" description:"route addr"`
}

func (a *Admin) Default() *Admin {
	return &Admin{
		Provider: string(superadmin.CMCCProvider),
		Addr:     "192.168.1.1",
	}
}
