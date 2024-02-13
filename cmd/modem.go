package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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

	b, err := json.MarshalIndent(superAdmin, "", " ")
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(os.Stdout, string(b), "\n")
	if err != nil {
		return err
	}

	return nil
}
