package user

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserPassword(t *testing.T) {
	validLongPassword := strings.Repeat("A1", minPasswordLength/2)

	tests := []struct {
		name      string
		password  string
		wantError bool
	}{
		{
			name:      "パスワードが短すぎる場合エラーが発生する",
			password:  "Short1",
			wantError: true,
		},
		{
			name:      "英字と数字が含まれている有効なパスワード",
			password:  validLongPassword,
			wantError: false,
		},
		{
			name:      "パスワードに英字が含まれていない場合エラーが発生する",
			password:  strings.Repeat("1", minPasswordLength), // 数字のみ
			wantError: true,
		},
		{
			name:      "パスワードに数字が含まれていない場合エラーが発生する",
			password:  strings.Repeat("A", minPasswordLength), // 英字のみ
			wantError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewUserPassword(test.password)
			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
