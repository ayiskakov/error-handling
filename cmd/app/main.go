package main

import (
	"fmt"

	"github.com/khanfromasia/error-handling/internal/app"
	"github.com/khanfromasia/error-handling/internal/catch"
	"github.com/khanfromasia/error-handling/internal/cli"
)

func main() {

	err := fmt.Errorf("path already exists: %s", "path")
	catch.Error(cli.ErrServiceAlreadyExists.Wrap(err))

	a := app.NewApp()

	if err := a.Exec(); err != nil {
		catch.Error(err)
	}

}
