package user

import (
	errDomain "github.com/tasuke/go-onion/domain/error"
	"github.com/tasuke/go-pkg/ulid"
	"strconv"
	"unicode/utf8"
)

type Career struct {
	id        string
	detail    string
	startYear int32
	endYear   int32
}

func (c *Career) ID() string {
	return c.id
}

func (c *Career) Detail() string {
	return c.detail
}

func (c *Career) StartYear() int32 {
	return c.startYear
}

func (c *Career) EndYear() int32 {
	return c.endYear
}

func NewCareer(detail string, startYear, endYear int32) (*Career, error) {
	return newCareer(
		ulid.NewULID(),
		detail,
		startYear,
		endYear,
	)
}

func ReconstructCareer(id string, detail string, startYear, endYear int32) (*Career, error) {
	return newCareer(
		id,
		detail,
		startYear,
		endYear,
	)
}

func newCareer(id string, detail string, startYear, endYear int32) (*Career, error) {
	// 詳細の長さのバリデーション
	if utf8.RuneCountInString(detail) > maxDetailLength {
		return nil, errDomain.NewError("詳細は" + strconv.Itoa(maxDetailLength) + "文字以下でなければなりません")
	}

	// 西暦のバリデーション
	if startYear < minValidYear {
		return nil, errDomain.NewError("開始年は" + strconv.Itoa(minValidYear) + "年以上でなければなりません")
	}
	if endYear < minValidYear {
		return nil, errDomain.NewError("終了年は" + strconv.Itoa(minValidYear) + "年以上でなければなりません")
	}

	// 西暦の範囲のバリデーション
	if endYear < startYear {
		return nil, errDomain.NewError("終了年は開始年以上でなければなりません")
	}

	// Career オブジェクトの生成
	return &Career{
		id:        id,
		detail:    detail,
		startYear: startYear,
		endYear:   endYear,
	}, nil
}

const (
	minValidYear    = 1970
	maxDetailLength = 1000
)
