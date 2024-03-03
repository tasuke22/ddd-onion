package user

import (
	"github.com/tasuke/go-pkg/ulid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSkill(t *testing.T) {
	skillID := ulid.NewULID()
	tagID := ulid.NewULID()

	tests := []struct {
		name       string
		id         string
		tagID      string
		evaluation int32
		year       int32
		wantError  bool
	}{
		{
			name:       "有効な入力でSkillが正常に作成される",
			id:         skillID,
			tagID:      tagID,
			evaluation: 3,
			year:       2,
			wantError:  false,
		},
		{
			name:       "評価が範囲外の場合にエラーが発生する（下限）",
			id:         skillID,
			tagID:      tagID,
			evaluation: minEvaluation - 1,
			year:       2,
			wantError:  true,
		},
		{
			name:       "評価が範囲外の場合にエラーが発生する（上限）",
			id:         skillID,
			tagID:      tagID,
			evaluation: maxEvaluation + 1,
			year:       2,
			wantError:  true,
		},
		{
			name:       "年数が範囲外の場合にエラーが発生する（下限）",
			id:         skillID,
			tagID:      tagID,
			evaluation: 3,
			year:       minYear - 1,
			wantError:  true,
		},
		{
			name:       "年数が範囲外の場合にエラーが発生する（上限）",
			id:         skillID,
			tagID:      tagID,
			evaluation: 3,
			year:       maxYear + 1,
			wantError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewSkill(test.id, test.tagID, test.evaluation, test.year)
			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
