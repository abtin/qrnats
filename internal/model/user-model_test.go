package model

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestNewUserFromJson(t *testing.T) {
	jsonString := `{"name": "abtin", "email" : "abtin@example.com", "twitter":"@ThisIsAbtin"}`
	jsonFileName := "user.json"
	file, err := os.Open(jsonFileName)
	if err != nil {
		t.Errorf("NewUserFromJson() error = %v", err)
		return
	}
	type testCase struct {
		name  string
		input io.Reader
		want  User
	}
	want := User{Name: "abtin", Email: "abtin@example.com", Twitter: "@ThisIsAbtin"}
	var testCases = []testCase{
		{"json string", strings.NewReader(jsonString), want},
		{"json file", file, want},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserFromJson(tt.input)
			if err != nil {
				t.Errorf("NewUserFromJson() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserFromJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
