package cli

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"
)

// ListDags executes the "airflow list_dags" command and returns a []string containing the id of all dags 
func ListDags() ([]string, error) {
	output, err := exec.Command("airflow", "list_dags").Output()

	if err != nil {
		return nil, err
	}

	reg, err := regexp.Compile(`(-*\nDAGS\n-*\n)`)

	if err != nil {
		return nil, err
	}

	indices := reg.FindIndex(output)

	if indices == nil {
		return nil, errors.New("Incompatible command output")
	}

	dagsString := strings.Trim(string(output[indices[1]:]), " \n")
	dags := strings.Split(dagsString, "\n")

	return dags, nil
}
