package length

import (
	"fmt"
	"strconv"

	"github.com/amasses/govalidate/helper"
	"github.com/amasses/govalidate/rules"
)

func init() {
	rules.Add("Length", Length)
}

// Validates that a string is N characters long
func Length(data rules.ValidationData) error {
	v, err := helper.ToString(data.Value)
	if err != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        "is not a string",
			Message:        data.Message,
		}
	}

	// We should always be provided with a length to validate against
	if len(data.Args) == 0 {
		return fmt.Errorf("No argument found in the validation struct (eg 'Length:5')")
	}

	// Typecast our argument and test
	var length int
	if length, err = strconv.Atoi(data.Args[0]); err != nil {
		return err
	}

	if len(v) != length {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        fmt.Sprintf("must be %d characters long", length),
			Message:        data.Message,
		}

	}

	return nil
}
