package user

import (
	"github.com/tasuke/go-pkg/ulid"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCareer(t *testing.T) {
	careerID := ulid.NewULID()
	longString := strings.Repeat("a", maxDetailLength+1)
	tests := []struct {
		name      string
		id        string
		detail    string
		startYear int32
		endYear   int32
		wantError bool
	}{
		{
			name:      "有効な入力でCareerが正常に作成される",
			id:        careerID,
			detail:    "有効なキャリア詳細",
			startYear: 2000,
			endYear:   2005,
			wantError: false,
		},
		{
			name:      "詳細が長すぎる場合にエラーが発生する",
			id:        careerID,
			detail:    longString,
			startYear: 2000,
			endYear:   2005,
			wantError: true,
		},
		{
			name:      "開始年が無効な場合にエラーが発生する",
			id:        careerID,
			detail:    "有効なキャリア詳細",
			startYear: minValidYear - 1,
			endYear:   2005,
			wantError: true,
		},
		{
			name:      "終了年が無効な場合にエラーが発生する",
			id:        careerID,
			detail:    "有効なキャリア詳細",
			startYear: 2000,
			endYear:   minValidYear - 1,
			wantError: true,
		},
		{
			name:      "終了年が開始年より前の場合にエラーが発生する",
			id:        careerID,
			detail:    "有効なキャリア詳細",
			startYear: 2005,
			endYear:   2000,
			wantError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := newCareer(test.id, test.detail, test.startYear, test.endYear)
			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
