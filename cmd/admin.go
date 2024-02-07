package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/uaxe/kuafu/cmd/flags"
	"github.com/uaxe/kuafu/internal/superadmin"
)

func admin(f *flags.Admin) error {
	if len(f.Addr) == 0 {
		return fmt.Errorf("please provide a addr using the -addr flag")
	}
	if len(f.Provider) == 0 {
		return fmt.Errorf("please provide a provider using the -provider flag")
	}
	provider, err := superadmin.NewProvider(context.Background(), superadmin.ProviderType(f.Provider), f.Addr)
	if err != nil {
		return err
	}
	superAdmin, err := provider.GetSuperAdmin()
	if err != nil {
		return err
	}
	raw, err := json.MarshalIndent(superAdmin, "", "	")
	if err != nil {
		return err
	}

	fmt.Println(string(raw))

	return nil
}
