package user

import (
	"golang.org/x/xerrors"
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
		return nil, xerrors.Errorf("名前は%d文字以下でなければなりません", maxNameLength)
	}

	// メールアドレスの長さのバリデーション
	if utf8.RuneCountInString(email) > maxEmailLength {
		return nil, xerrors.Errorf("メールアドレスは%d文字以下でなければなりません", maxEmailLength)
	}

	// 自己紹介の長さのバリデーション
	if utf8.RuneCountInString(profile) > maxProfileLength {
		return nil, xerrors.Errorf("自己紹介は%d文字以下でなければなりません", maxProfileLength)
	}

	// スキルのバリデーション
	if len(skills) < minSkillsLength {
		return nil, xerrors.Errorf("スキルは%dつ以上でなければなりません", minSkillsLength)
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
