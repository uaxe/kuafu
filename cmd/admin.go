package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/uaxe/kuafu/provider"

	"github.com/uaxe/kuafu/provider/superadmin"
)

func admin(f *superadmin.AdminFlag) error {

	if len(f.Host) == 0 {
		return fmt.Errorf("please provide host using the -host flag")
	}

	if f.Port == 0 {
		return fmt.Errorf("please provide port using the -port flag")
	}

	provide := provider.New(context.Background(), f.Type)

	superAdminProvider, err := provide.SuperAdminProvider(f)
	if err != nil {
		return err
	}

	superAdmin, err := superAdminProvider.GetSuperAdmin()
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
