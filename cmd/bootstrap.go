package cmd

import (
	"fmt"
	"github.com/uaxe/infra/color"
	"github.com/uaxe/infra/zcli"
	"os"
)

var (
	version = "0.0.1"
	app     *zcli.Cli
)

func banner(_ *zcli.Cli) string {
	return fmt.Sprintf("%s %s",
		color.Green("Kuafu CLI"),
		color.Cyan(version))
}

func Bootstrap() {
	app = zcli.NewCli("KuaFu", "The Go Kuafu", version)
	app.SetBannerFunction(banner)

	cmd := app.NewSubCommand("version", "The Kuafu CLI version")
	cmd.Action(func() error {
		fmt.Println(version)
		return nil
	})

	app.NewSubCommandFunction("admin", "The Kuafu admin", admin)

	if err := app.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
