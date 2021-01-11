// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Article struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Brief      string    `json:"brief"`
	Content    string    `json:"content"`
	CategoryID string    `json:"categoryID"`
	Category   *Category `json:"category"`
	Tags       []*Tag    `json:"tags"`
	AuthorID   string    `json:"authorID"`
	Author     *User     `json:"author"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type ArticleInput struct {
	Name       string      `json:"name"`
	CategoryID string      `json:"categoryID"`
	Brief      string      `json:"brief"`
	Content    string      `json:"content"`
	Tags       []*TagInput `json:"tags"`
	AuthorID   string      `json:"authorID"`
}

type ArticleResponse struct {
	Message  string     `json:"message"`
	Status   int        `json:"status"`
	Data     *Article   `json:"data"`
	DataList []*Article `json:"dataList"`
}

type Category struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Articles []*Article `json:"articles"`
}

type CategoryInput struct {
	Name string `json:"name"`
}

type CategoryResponse struct {
	Message  string      `json:"message"`
	Status   int         `json:"status"`
	Data     *Category   `json:"data"`
	DataList []*Category `json:"dataList"`
}

type LoginResponse struct {
	Message string  `json:"message"`
	Status  int     `json:"status"`
	Token   *string `json:"token"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TagInput struct {
	Name string `json:"name"`
}

type User struct {
	ID           string     `json:"id"`
	Username     string     `json:"username"`
	EmailAddress string     `json:"emailAddress"`
	Password     string     `json:"password"`
	Role         Role       `json:"role"`
	Articles     []*Article `json:"articles"`
}

type UserInput struct {
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	Role         Role   `json:"role"`
}

type UserResponse struct {
	Message  string  `json:"message"`
	Status   int     `json:"status"`
	Data     *User   `json:"data"`
	DataList []*User `json:"dataList"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
