package gd

import ErrorType "graphics.gd/variant/Error"

type Error = ErrorType.Code

func ToError(err Error) error {
	if err == Error(0) {
		return nil
	}
	return err
}
