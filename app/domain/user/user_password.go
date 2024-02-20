package user

import (
	"golang.org/x/xerrors"
	"unicode"
)

type UserPassword string

func newUserPassword(password string) (UserPassword, error) {
	var hasLetter, hasDigit bool
	// パスワードの長さのバリデーション
	if len(password) < minPasswordLength {
		return "", xerrors.Errorf("パスワードは少なくとも%d文字である必要があります", minPasswordLength)
	}

	// パスワード内の文字の種類のバリデーション
	for _, c := range password {
		switch {
		case unicode.IsLetter(c):
			hasLetter = true
		case unicode.IsDigit(c):
			hasDigit = true
		}
	}

	// 英字と数字が最低1文字ずつ含まれているかのチェック
	if !hasLetter || !hasDigit {
		return "", xerrors.New("パスワードには少なくとも1つの英字と1つの数字が含まれている必要があります")
	}

	// UserPassword オブジェクトの生成
	return UserPassword(password), nil
}

const minPasswordLength = 12
