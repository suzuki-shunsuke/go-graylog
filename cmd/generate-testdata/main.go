package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sanity-io/litter"

	"github.com/suzuki-shunsuke/go-graylog"
)

type (
	dumper interface {
		dump(string) error
	}
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	options := make([]string, len(data))
	i := 0
	for k := range data {
		options[i] = k
		i++
	}
	file := ""
	prompt := &survey.Select{
		Message: "file",
		Options: options,
	}
	if err := survey.AskOne(prompt, &file); err != nil {
		return err
	}
	d, ok := data[file]
	if !ok {
		return errors.New("file is not found")
	}
	if err := d.dump("testdata/" + file + ".json"); err != nil {
		return err
	}
	return nil
}

func dump(input string, dest interface{}) error {
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(dest); err != nil {
		return err
	}
	litter.Dump(dest)
	return nil
}

var (
	data = map[string]dumper{
		"users": Users{},
		"user":  User{},
	}
)

type (
	Users struct {
		users graylog.UsersBody
	}

	User struct {
		user graylog.User
	}
)

func (users Users) dump(input string) error {
	return dump(input, &users.users)
}

func (user User) dump(input string) error {
	return dump(input, &user.user)
}
