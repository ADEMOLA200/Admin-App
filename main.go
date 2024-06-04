package main

import (

	"github.com/ADEMOLA200/Admin-App.git/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(app.Module).Run()
}
 