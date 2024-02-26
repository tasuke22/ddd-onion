package user

import (
	"golang.org/x/xerrors"
)

type Skill struct {
	id         string
	tagID      string
	evaluation int32
	years      int32
}

func (s *Skill) ID() string {
	return s.id
}

func (s *Skill) TagID() string {
	return s.tagID
}

func (s *Skill) Evaluation() int32 {
	return s.evaluation
}

func (s *Skill) Years() int32 {
	return s.years
}

func newSkill(id string, tagID string, evaluation, year int32) (*Skill, error) {
	// 評価（evaluation）のバリデーション
	if evaluation < minEvaluation || evaluation > maxEvaluation {
		return nil, xerrors.Errorf("評価は%dから%dの間でなければなりません", minEvaluation, maxEvaluation)
	}

	// 年数（years）のバリデーション
	if year < minYear || year > maxYear {
		return nil, xerrors.Errorf("年数は%dから%dの間でなければなりません", minYear, maxYear)
	}

	// スキルオブジェクトの生成
	return &Skill{
		id:         id,
		tagID:      tagID,
		evaluation: evaluation,
		years:      year,
	}, nil
}

const (
	minEvaluation = 1
	maxEvaluation = 5
	minYear       = 0
	maxYear       = 5
)
