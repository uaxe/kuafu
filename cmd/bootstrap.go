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

	ver := app.NewSubCommand("version", "The Kuafu CLI version")
	ver.Action(func() error {
		fmt.Println(fmt.Sprintf("%s %s", color.Green("Kuafu CLI"), color.Cyan(version)))
		return nil
	})

	modem := app.NewSubCommand("modem", "The Kuafu CLI modem")
	modem.NewSubCommandFunction("admin", "The Kuafu modem admin", _admin)

	if err := app.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
