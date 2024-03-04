package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	mockTag "github.com/tasuke/go-onion/domain/tag"
	tagDomain "github.com/tasuke/go-onion/domain/tag"
	mockUser "github.com/tasuke/go-onion/domain/user"
	"github.com/tasuke/go-pkg/ulid"
	"go.uber.org/mock/gomock"
)

func TestSaveUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockUser.NewMockUserRepository(ctrl)
	mockTagRepo := mockTag.NewMockTagRepository(ctrl)

	useCase := NewUserUseCase(mockUserRepo, mockTagRepo)

	tests := []struct {
		name      string
		input     *SaveUserUseCaseInputDto
		setupMock func()
		wantErr   bool
	}{
		{
			name: "正常にユーザーが保存される場合",
			input: &SaveUserUseCaseInputDto{
				Name:     "新規ユーザー",
				Email:    "newuser@example.com",
				Password: "password1234",
				Profile:  "新しいプロフィール",
				Skills: []SkillInputDto{
					{Evaluation: 5, Years: 3, TagName: "Go"},
				},
				Careers: []CareerInputDto{
					{Detail: "エンジニア", StartYear: 2018, EndYear: 2021},
				},
			},
			setupMock: func() {
				mockTagRepo.EXPECT().FindByNames(context.Background(), []string{"Go"}).DoAndReturn(
					func(ctx context.Context, names []string) ([]*tagDomain.Tag, error) {
						var tags []*tagDomain.Tag
						for _, name := range names {
							tag, _ := tagDomain.ReconstructTag(ulid.NewULID(), name)
							tags = append(tags, tag)
						}
						return tags, nil
					},
				)
				mockUserRepo.EXPECT().Save(context.Background(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "タグの処理中にエラーが発生する場合",
			input: &SaveUserUseCaseInputDto{
				Name:     "新規ユーザー",
				Email:    "newuser@example.com",
				Password: "password123",
				Profile:  "新しいプロフィール",
				Skills: []SkillInputDto{
					{Evaluation: 5, Years: 3, TagName: "Python"},
				},
			},
			setupMock: func() {
				mockTagRepo.EXPECT().FindByNames(context.Background(), []string{"Python"}).Return(nil, errors.New("DBエラー"))
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
