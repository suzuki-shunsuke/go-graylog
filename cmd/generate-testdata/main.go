package main

// Generate test data
// 1. call Graylog API by API browser
// 2. copy and paste the response to testdata/***.json
// 3. add type and update data
// 4. run this command and output test data
// 5. copy and paste the test data to testdata/***.go

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
		"users":      Users{},
		"user":       User{},
		"roles":      Roles{},
		"role":       Role{},
		"index_sets": IndexSets{},
		"index_set":  IndexSet{},
		"inputs":     Inputs{},
		"input":      Input{},
	}
)

type (
	Users struct {
		users graylog.UsersBody
	}

	User struct {
		user graylog.User
	}

	Roles struct {
		roles graylog.RolesBody
	}

	Role struct {
		role graylog.Role
	}

	IndexSets struct {
		data graylog.IndexSetsBody
	}

	IndexSet struct {
		data graylog.IndexSet
	}

	Inputs struct {
		data graylog.InputsBody
	}

	Input struct {
		data graylog.Input
	}
)

func (users Users) dump(input string) error {
	return dump(input, &users.users)
}

func (user User) dump(input string) error {
	return dump(input, &user.user)
}

func (roles Roles) dump(input string) error {
	return dump(input, &roles.roles)
}

func (role Role) dump(input string) error {
	return dump(input, &role.role)
}

func (is IndexSets) dump(input string) error {
	return dump(input, &is.data)
}

func (is IndexSet) dump(input string) error {
	return dump(input, &is.data)
}

func (ip Inputs) dump(input string) error {
	return dump(input, &ip.data)
}

func (ip Input) dump(input string) error {
	return dump(input, &ip.data)
}
