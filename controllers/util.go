package controllers

import "errors"

var (
	ErrUser = errors.New("illegal user")
)

func CheckUser(token int) (isAdmin bool, err error) {
	if token == 111 {
		isAdmin = true
		err = nil
		return
	} else if token == 222 {
		isAdmin = false
		err = nil
		return
	} else {
		isAdmin = false
		err = ErrUser
		return
	}
}
