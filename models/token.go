/*
@Time : 20-6-3
@Author : jzd
@Project: mvc
*/
package models

type Token struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token"` // the token returned by azure container registry is called "access_token"
	ExpiresIn   int    `json:"expires_in"`
	IssuedAt    string `json:"issued_at"`
}
