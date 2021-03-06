package log

import "errors"

func LogMethodError(n string, e error) error {
	return errors.New("ERROR: " + n + " >>> " + e.Error())
}

func Msg(s string) error {
	return errors.New(s)
}
