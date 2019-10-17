package server

import "net/http"

func helpCheckOrigin(_ *http.Request) bool {
	return true
}

func helpCheckString(variable interface{}) bool {
	_, err := variable.(string)
	return err
}
