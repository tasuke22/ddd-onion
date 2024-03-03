package user

import (
	errDomain "github.com/tasuke/go-onion/domain/error"
	"strconv"
	"unicode"
)

type UserPassword string

func NewUserPassword(password string) (UserPassword, error) {
	return newUserPassword(password)
}

func ConvertToUserPassword(password string) UserPassword {
	return UserPassword(password)
}

func newUserPassword(password string) (UserPassword, error) {
	var hasLetter, hasDigit bool
	// パスワードの長さのバリデーション
	if len(password) < minPasswordLength {
		return "", errDomain.NewError("パスワードは少なくとも" + strconv.Itoa(minPasswordLength) + "文字である必要があります")
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
		return "", errDomain.NewError("パスワードには少なくとも1つの英字と1つの数字が含まれている必要があります")
	}

	// UserPassword オブジェクトの生成
	return ConvertToUserPassword(password), nil
}

const minPasswordLength = 12
