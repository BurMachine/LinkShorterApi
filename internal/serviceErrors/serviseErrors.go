package serviceErrors

import "errors"

var ErrorLinkExist error = errors.New("short link already exists")
var ErrorLinkNotExist error = errors.New("short link does not exist exists")
