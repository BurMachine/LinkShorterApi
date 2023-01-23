package flagsHandling

import "errors"

func CheckFlagString(str string) (string, error) {
	if str == "postgres" {
		return str, nil
	} else if str == "inmemory" {
		return str, nil
	} else {
		return "", errors.New("storage flag handling failed")
	}
}
