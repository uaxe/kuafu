package cmd

import (
	"context"
	"fmt"
	"github.com/uaxe/kuafu/provider"
	"github.com/uaxe/kuafu/provider/modem"
)

func _admin(f *modem.AdminFlag) error {

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

	fmt.Println(fmt.Sprintf("mac_addr: %s", superAdmin.MacAddr))

	fmt.Println(fmt.Sprintf("admin_name: %s", superAdmin.Name))

	fmt.Println(fmt.Sprintf("admin_pwd: %s", superAdmin.Pwd))

	return nil
}
