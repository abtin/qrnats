package model

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Twitter string `json:"twitter"`
}

func NewUserFromJson(jsonInput io.Reader) (User, error) {
	b, err := ioutil.ReadAll(jsonInput)
	if err != nil {
		return User{}, err
	}
	var u User
	err = json.Unmarshal(b, &u)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (u User) String() string {
	return fmt.Sprintf("User: Name=%s, Email=%s, Twitter=%s", u.Name, u.Email, u.Twitter)
}
