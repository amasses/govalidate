package alpha

import (
	"regexp"

	"github.com/amasses/govalidate/helper"
	"github.com/amasses/govalidate/rules"
)

func init() {
	rules.Add("Alpha", Alpha)
}

// Validates that a string only contains alphabetic characters
func Alpha(data rules.ValidationData) (err error) {
	v, ok := helper.ToString(data.Value)
	if ok != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        "is not a string",
			Message:        data.Message,
		}
	}

	if regexp.MustCompile(`[^a-zA-Z]+`).MatchString(v) {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        "contains non-alphabetic characters",
			Message:        data.Message,
		}
	}

	return nil
}
