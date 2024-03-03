package user

import (
	"github.com/tasuke/go-pkg/ulid"
)

type CareerDto struct {
	Detail    string
	StartYear int32
	EndYear   int32
}
type SkillDto struct {
	Evaluation int32
	Year       int32
	TagID      string
}

func Create(
	name string,
	email string,
	password string,
	profile string,
	careersDto []CareerDto,
	skillsDto []SkillDto,
) (*User, error) {

	// キャリアのインスタンスを作成
	careers := make([]Career, len(careersDto))
	for i, rc := range careersDto {
		c, err := NewCareer(
			rc.Detail,
			rc.StartYear,
			rc.EndYear,
		)
		if err != nil {
			return nil, err
		}
		careers[i] = *c
	}

	// スキルのインスタンスを作成
	// タグIDのインスタンスを作成
	skills := make([]Skill, len(skillsDto))
	for i, sd := range skillsDto {
		s, err := NewSkill(
			sd.TagID,
			sd.Evaluation,
			sd.Year,
		)
		if err != nil {
			return nil, err
		}
		skills[i] = *s
	}

	// パスワードのインスタンスを作成
	newUserPassword, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}

	return newUser(
		ulid.NewULID(),
		name,
		email,
		newUserPassword,
		profile,
		skills,
		careers,
	)
}
