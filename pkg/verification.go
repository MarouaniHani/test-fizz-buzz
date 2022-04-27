package pkg

import (
	"github.com/pkg/errors"
	"test-fizz-buzz/dto"
)

func ValidateRequest(request dto.Request) error {
	if len(request.FirstString) == 0 {
		return errors.New("you should enter the first string")
	}
	if len(request.FirstString) == 0 {
		return errors.New("you should enter the second string")
	}
	if request.FirstNumber == request.SecondNumber {
		return errors.New("InvalidParameters: The 2 integer should be different")
	}
	if request.Limit < 1 {
		return errors.New("InvalidParameters: Limit should be greater than 1")
	}
	if request.FirstString == request.SecondString {
		return errors.New("InvalidParameters: The 2 strings should be different")
	}
	return nil
}
