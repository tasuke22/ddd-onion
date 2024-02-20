package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	validEmail := "example@example.com"
	validPassword := "ValidPassword123"
	validProfile := "Valid profile"
	validDetail := "Valid career detail"
	validTagID := "valid-tag-id"

	tests := []struct {
		testName    string
		name        string
		email       string
		reqPassword string
		profile     string
		careersDto  []CareerDto
		skillsDto   []SkillDto
		wantError   bool
	}{
		{
			testName:    "有効な入力でユーザーが正常に作成される",
			name:        "Valid Name",
			email:       validEmail,
			reqPassword: validPassword,
			profile:     validProfile,
			careersDto: []CareerDto{
				{Detail: validDetail, StartYear: 2010, EndYear: 2020},
			},
			skillsDto: []SkillDto{
				{Evaluation: 3, Year: 2, TagID: validTagID},
			},
			wantError: false,
		},
		{
			testName:    "無効なパスワードでエラーが発生する",
			name:        "Valid Name",
			email:       validEmail,
			reqPassword: "short",
			profile:     validProfile,
			careersDto:  []CareerDto{},
			skillsDto:   []SkillDto{},
			wantError:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := Create(test.name, test.email, test.reqPassword, test.profile, test.careersDto, test.skillsDto)
			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
