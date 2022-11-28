package usecase

import (
	"errors"
	"fmt"
	"os"

	"github.com/khanfromasia/error-handling/internal/cli"
)

func (u *UseCase) CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)

	if err != nil {
		if errors.Is(err, os.ErrExist) {
			err = fmt.Errorf("path already exists: %s", path)
			return nil, cli.ErrServiceAlreadyExists.Wrap(err)
		}

		return nil, cli.ErrServiceInternal.Wrap(err)
	}

	return file, nil
}
