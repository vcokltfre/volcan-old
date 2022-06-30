package commands

import (
	"fmt"
	"strconv"
)

func ValidateInt(lower, upper int) Validator {
	return func(str string) error {
		i, err := strconv.Atoi(str)
		if err != nil {
			return err
		}

		if i >= lower && i <= upper {
			return nil
		}

		return fmt.Errorf("Integer must be within the bounds %d to %d", lower, upper)
	}
}
