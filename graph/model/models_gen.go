// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthAdminOps struct {
	Login    interface{} `json:"login"`
	Register interface{} `json:"register"`
}

type RegisterAdmin struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
