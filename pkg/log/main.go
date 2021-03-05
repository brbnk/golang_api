package log

import "errors"

func LogMethodError(n string, e error) error {
	return errors.New("ERROR: Method" + n + " >>> " + e.Error())
}
