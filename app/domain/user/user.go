package user

import (
	errDomain "github.com/tasuke/go-onion/domain/error"
	"strconv"
	"unicode/utf8"
)

type User struct {
	id       string
	name     string
	email    string
	password UserPassword
	profile  string
	skills   []Skill
	careers  []Career
}

func ReconstructUser(
	id string,
	name string,
	email string,
	password UserPassword,
	profile string,
	skills []Skill,
	careers []Career,
) (*User, error) {
	return newUser(
		id,
		name,
		email,
		password,
		profile,
		skills,
		careers,
	)
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() UserPassword {
	return u.password
}

func (u *User) Profile() string {
	return u.profile
}
func (u *User) Skills() []Skill {
	return u.skills
}

func (u *User) Careers() []Career {
	return u.careers
}

func newUser(
	id string,
	name string,
	email string,
	password UserPassword,
	profile string,
	skills []Skill,
	careers []Career,
) (*User, error) {

	// 名前の長さのバリデーション
	if utf8.RuneCountInString(name) > maxNameLength {
		return nil, errDomain.NewError("名前は" + strconv.Itoa(maxNameLength) + "文字以下でなければなりません")
	}

	// メールアドレスの長さのバリデーション
	if utf8.RuneCountInString(email) > maxEmailLength {
		return nil, errDomain.NewError("メールアドレスは" + strconv.Itoa(maxEmailLength) + "文字以下でなければなりません")
	}

	// 自己紹介の長さのバリデーション
	if utf8.RuneCountInString(profile) > maxProfileLength {
		return nil, errDomain.NewError("自己紹介は" + strconv.Itoa(maxProfileLength) + "文字以下でなければなりません")
	}

	// スキルのバリデーション
	if len(skills) < minSkillsLength {
		return nil, errDomain.NewError("スキルは" + strconv.Itoa(minSkillsLength) + "つ以上でなければなりません")
	}

	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
		profile:  profile,
		skills:   skills,
		careers:  careers,
	}, nil
}

const (
	maxNameLength    = 255
	maxEmailLength   = 255
	maxProfileLength = 2000
	minSkillsLength  = 1
)
