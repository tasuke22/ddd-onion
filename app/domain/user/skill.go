package user

import (
	errDomain "github.com/tasuke/go-onion/domain/error"
	"github.com/tasuke/go-pkg/ulid"
	"strconv"
)

type Skill struct {
	id         string
	tagID      string
	evaluation int32
	years      int32
}

func NewSkill(tagID string, evaluation, year int32) (*Skill, error) {
	return newSkill(
		ulid.NewULID(),
		tagID,
		evaluation,
		year,
	)
}

func ReconstructSkill(id string, tagID string, evaluation, year int32) (*Skill, error) {
	return newSkill(
		id,
		tagID,
		evaluation,
		year,
	)
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
		return nil, errDomain.NewError("評価は" + strconv.Itoa(minEvaluation) + "から" + strconv.Itoa(maxEvaluation) + "の間でなければなりません")
	}

	// 年数（years）のバリデーション
	if year < minYear || year > maxYear {
		return nil, errDomain.NewError("年数は" + strconv.Itoa(minYear) + "から" + strconv.Itoa(maxYear) + "の間でなければなりません")
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
