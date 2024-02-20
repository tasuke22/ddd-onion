package user

import (
	"github.com/tasuke/go-pkg/ulid"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	validID := ulid.NewULID()
	validName := "Valid Name"
	validEmail := "valid@example.com"
	validPassword, _ := newUserPassword("ValidPassword123")
	validProfile := "Valid Profile"
	validSkills := []Skill{{}}
	validCareers := []Career{{}}

	longString := strings.Repeat("a", maxNameLength+1)
	longProfile := strings.Repeat("a", maxProfileLength+1)

	tests := []struct {
		testName  string
		id        string
		name      string
		email     string
		password  UserPassword
		profile   string
		skills    []Skill
		careers   []Career
		wantError bool
	}{
		{
			testName:  "有効な入力でユーザーが正常に作成される",
			id:        validID,
			name:      validName,
			email:     validEmail,
			password:  validPassword,
			profile:   validProfile,
			skills:    validSkills,
			careers:   validCareers,
			wantError: false,
		},
		{
			testName:  "名前が長すぎる場合エラーが発生する",
			id:        validID,
			name:      longString,
			email:     validEmail,
			password:  validPassword,
			profile:   validProfile,
			skills:    validSkills,
			careers:   validCareers,
			wantError: true,
		},
		{
			testName:  "メールアドレスが長すぎる場合エラーが発生する",
			id:        validID,
			name:      validName,
			email:     longString,
			password:  validPassword,
			profile:   validProfile,
			skills:    validSkills,
			careers:   validCareers,
			wantError: true,
		},
		{
			testName:  "自己紹介が長すぎる場合エラーが発生する",
			id:        validID,
			name:      validName,
			email:     validEmail,
			password:  validPassword,
			profile:   longProfile,
			skills:    validSkills,
			careers:   validCareers,
			wantError: true,
		},
		{
			testName:  "スキルが不足している場合エラーが発生する",
			id:        validID,
			name:      validName,
			email:     validEmail,
			password:  validPassword,
			profile:   validProfile,
			skills:    []Skill{}, // スキル不足
			careers:   validCareers,
			wantError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := newUser(test.id, test.name, test.email, test.password, test.profile, test.skills, test.careers)
			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
