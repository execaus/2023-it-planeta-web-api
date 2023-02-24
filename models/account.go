package models

type RegistrationAccountInput struct {
	FirstName string `json:"firstName" binding:"required,excludesall=' ',printascii"`
	LastName  string `json:"lastName" binding:"required,excludesall=' ',printascii"`
	Email     string `json:"email" binding:"required,email,excludesall=' ',printascii"`
	Password  string `json:"password" binding:"required,excludesall=' ',printascii"`
}

type RegistrationAccountOutput struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type GetAccountOutput struct {
	Id        int64  `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

type GetAccountsInput struct {
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
	Email     string `form:"email"`
	From      *int   `form:"from"`
	Size      *int   `form:"size"`
}

type GetAccountsOutput struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UpdateAccountInput struct {
	FirstName string `json:"firstName" binding:"required,excludesall=' ',printascii"`
	LastName  string `json:"lastName" binding:"required,excludesall=' ',printascii"`
	Email     string `json:"email" binding:"required,email,excludesall=' ',printascii"`
	Password  string `json:"password" binding:"required,excludesall=' ',printascii"`
}

type UpdateAccountOutput struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
