package user

import "github.com/tasuke/go-pkg/ulid"

type CareerDto struct {
	Detail    string
	StartYear int
	EndYear   int
}
type SkillDto struct {
	Evaluation int
	Year       int
	TagID      string
}

func Create(
	name string,
	email string,
	reqPassword string,
	profile string,
	careersDto []CareerDto,
	skillsDto []SkillDto,
) (*User, error) {

	// キャリアのインスタンスを作成
	careers := make([]Career, len(careersDto))
	for i, rc := range careersDto {
		c, err := newCareer(ulid.NewULID(), rc.Detail, rc.StartYear, rc.EndYear)
		if err != nil {
			return nil, err
		}
		careers[i] = *c
	}

	// スキルのインスタンスを作成
	// タグIDのインスタンスを作成
	skills := make([]Skill, len(skillsDto))
	for i, sd := range skillsDto {
		s, err := newSkill(ulid.NewULID(), sd.TagID, sd.Evaluation, sd.Year)
		if err != nil {
			return nil, err
		}
		skills[i] = *s
	}

	// パスワードのインスタンスを作成
	password, err := newUserPassword(reqPassword)
	if err != nil {
		return nil, err
	}

	return newUser(
		ulid.NewULID(),
		name,
		email,
		password,
		profile,
		skills,
		careers,
	)
}
