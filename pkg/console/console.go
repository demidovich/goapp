package console

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func InputString(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}
	value, _ := prompt.Run()

	return strings.Trim(value, " ")
}

func InputStringValidated(label string, validator func(string) error) string {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validator,
	}
	value, _ := prompt.Run()

	return strings.Trim(value, " ")
}

func InputInt(label string) int {
	re := regexp.MustCompile(`^\s*-?[0-9]+s*$`)
	validator := func(value string) error {
		if value == "" || re.MatchString(value) {
			return nil
		} else {
			return errors.New("the value must satisfy a regular expression \"-?[0-9]+\".")
		}
	}

	valueStr := InputStringValidated(label, validator)
	value, _ := strconv.Atoi(valueStr)

	return value
}

func Select(label string, options []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	_, value, _ := prompt.Run()

	return value
}
