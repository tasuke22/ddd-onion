package user

import (
	"context"
	"errors"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/stretchr/testify/assert"
	mockTag "github.com/tasuke/go-onion/domain/tag"
	mockUser "github.com/tasuke/go-onion/domain/user"

	userDomain "github.com/tasuke/go-onion/domain/user"
)

func TestUpdateUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockUser.NewMockUserRepository(ctrl)
	mockTagRepo := mockTag.NewMockTagRepository(ctrl)

	useCase := NewUpdateUserUseCase(mockUserRepo, mockTagRepo)

	tests := []struct {
		name      string
		input     *UpdateUserUseCaseInputDto
		setupMock func()
		wantErr   bool
	}{
		{
			name: "ユーザー更新成功",
			input: &UpdateUserUseCaseInputDto{
				UserID:   "test_user_id",
				Name:     "Updated User",
				Email:    "updated@example.com",
				Password: "newpassword",
				Profile:  "updated profile",
				Skills:   []UpdateSkillDto{},
				Careers:  []UpdateCareerDto{},
			},
			setupMock: func() {
				existingUser := &userDomain.User{}
				mockUserRepo.EXPECT().FindByUserID(context.Background(), "test_user_id").Return(existingUser, nil)
				mockUserRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ユーザーが見つからない",
			input: &UpdateUserUseCaseInputDto{
				UserID: "nonexistent_user_id",
			},
			setupMock: func() {
				mockUserRepo.EXPECT().FindByUserID(context.Background(), "nonexistent_user_id").Return(nil, errors.New("user not found"))
			},
			wantErr: true,
		},
		{
			name: "更新時にリポジトリエラー",
			input: &UpdateUserUseCaseInputDto{
				UserID:   "test_user_id",
				Name:     "Another User",
				Email:    "another@example.com",
				Password: "anotherpassword",
				Profile:  "another profile",
			},
			setupMock: func() {
				existingUser := &userDomain.User{}
				mockUserRepo.EXPECT().FindByUserID(context.Background(), "test_user_id").Return(existingUser, nil)
				mockUserRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(errors.New("repository error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			err := useCase.Run(context.Background(), tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
