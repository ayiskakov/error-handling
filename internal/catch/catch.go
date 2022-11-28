package catch

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/khanfromasia/error-handling/internal/cli"
)

type Details struct {
	Fail string   `yaml:"fail"`
	What []string `yaml:"what"`
}

type Errs struct {
	Errs map[string]Details `yaml:"errors"`
}

func Error(err error) {
	if err == nil {
		return
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	file, errF := os.Open("./errors.yaml")
	if errF != nil {
		fmt.Fprintln(out, "error: no errors.yaml")
		return
	}

	var errs Errs

	if errD := yaml.NewDecoder(file).Decode(&errs); errD != nil {
		fmt.Fprintln(out, "error: decoding errors.yaml")
		return
	}

	var cliErr *cli.ErrorID
	//log.Printf("#%#v", err)
	if errors.As(err, &cliErr) {
		cliErr = err.(*cli.ErrorID)

		fmt.Fprintln(out, errs.Errs[cliErr.ID].Fail)
		fmt.Fprintln(out, "What to do:")
		for _, w := range errs.Errs[cliErr.ID].What {
			fmt.Fprintln(out, fmt.Sprintf("\t- %s", w))
		}
		return
	} else {
		log.Println("cliErr.ID")
	}
}
