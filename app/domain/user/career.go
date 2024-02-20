package user

import (
	"golang.org/x/xerrors"
	"unicode/utf8"
)

type Career struct {
	id        string
	detail    string
	startYear int
	endYear   int
}

func (c *Career) ID() string {
	return c.id
}

func (c *Career) Detail() string {
	return c.detail
}

func (c *Career) StartYear() int {
	return c.startYear
}

func (c *Career) EndYear() int {
	return c.endYear
}

func newCareer(id string, detail string, startYear, endYear int) (*Career, error) {
	// 詳細の長さのバリデーション
	if utf8.RuneCountInString(detail) > maxDetailLength {
		return nil, xerrors.Errorf("詳細は%d文字以下でなければなりません", maxDetailLength)
	}

	// 西暦のバリデーション
	if startYear < minValidYear {
		return nil, xerrors.Errorf("開始年は%d年以上でなければなりません", minValidYear)
	}
	if endYear < minValidYear {
		return nil, xerrors.Errorf("終了年は%d年以上でなければなりません", minValidYear)
	}

	// 西暦の範囲のバリデーション
	if endYear < startYear {
		return nil, xerrors.New("終了年は開始年以上でなければなりません")
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
