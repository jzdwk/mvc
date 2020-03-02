/*
@Time : 10/14/19
@Author : jzd
@Project: sigmaop
*/
package regex

import (
	"regexp"
)

const (
	//user
	USER_NAME  = `^[a-zA-Z][a-zA-Z0-9_.]{4,20}$`
	USER_EMAIL = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	USER_PHONE = `^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$`
)

func RegexCheck(reg string, param string) bool {
	exp := regexp.MustCompile(reg)
	return exp.MatchString(param)
}
