package core 

import (
	"fmt"
	"errors"
)

func Validator(tools []Tool) error {
	var errs []error
	for i ,t := range tools {
		if t.Name == "" {
			errs = append(errs,fmt.Errorf("tool[%d].name empty",i))

		}

	}
	if len(errs) > 0 {
			return errors.Join(errs...)

	}

	return nil
}
