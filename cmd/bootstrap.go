package cmd

import (
	"fmt"
	"os"

	"github.com/uaxe/infra/color"
	"github.com/uaxe/infra/zcli"
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
		verText := fmt.Sprintf("%s %s", color.Green("Kuafu CLI"), color.Cyan(version))
		fmt.Println(verText)
		return nil
	})

	modem := app.NewSubCommand("modem", "The Kuafu CLI modem")
	modem.NewSubCommandFunction("admin", "The Kuafu modem admin", _admin)

	if err := app.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
