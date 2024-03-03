package user

import (
	"context"
	errDomain "github.com/tasuke/go-onion/domain/error"
	tagDomain "github.com/tasuke/go-onion/domain/tag"
	userDomain "github.com/tasuke/go-onion/domain/user"
)

type SaveUserUseCase struct {
	userRepo userDomain.UserRepository
	tagRepo  tagDomain.TagRepository
}

func NewUserUseCase(
	userRepo userDomain.UserRepository,
	tagRepo tagDomain.TagRepository,
) *SaveUserUseCase {
	return &SaveUserUseCase{
		userRepo,
		tagRepo,
	}
}

type SaveUserUseCaseInputDto struct {
	Name     string           `json:"name"`
	Email    string           `json:"email"`
	Password string           `json:"password"`
	Profile  string           `json:"profile"`
	Skills   []SkillInputDto  `json:"skills"`
	Careers  []CareerInputDto `json:"careers"`
}

type CareerInputDto struct {
	Detail    string `json:"detail"`
	StartYear int32  `json:"start_year"`
	EndYear   int32  `json:"end_year"`
}

type SkillInputDto struct {
	Evaluation int32  `json:"evaluation"`
	Years      int32  `json:"years"`
	TagName    string `json:"tag_name"`
}

func (uc *SaveUserUseCase) Run(ctx context.Context, input *SaveUserUseCaseInputDto) error {
	// ユーザー名の重複チェック TODO
	//if err := uc.checkUserExists(ctx, input.Name); err != nil {
	//	return err
	//}

	// タグの処理 ここを修正したい
	tagsMap, err := uc.processTags(ctx, input.Skills)
	//{
	//	"既存のタグ名(PHP)": "dsfak2jk3jk3",
	//	"既存のタグ名(Go)": "fasdfafdsaffsd",
	//	"新規のタグ名(Ruby)": "gef223423423",
	//}

	if err != nil {
		return err
	}

	// ユーザー情報の作成
	user, err := uc.createUserDomain(input, tagsMap)
	if err != nil {
		return err
	}

	// ユーザー情報の保存
	return uc.userRepo.Save(ctx, user)
}

func (uc *SaveUserUseCase) createUserDomain(request *SaveUserUseCaseInputDto, tagsMap map[string]string) (*userDomain.User, error) {

	// 最終的にはTagIDを持つSkillDtoのスライスを作成する
	//type SkillDto struct {
	//	Evaluation int32
	//	Year       int32
	//	TagID      string
	//}

	skillsDto := make([]userDomain.SkillDto, len(request.Skills))
	for i, skill := range request.Skills {
		// tagsMapには、TagNameをキーとしてTagIDが格納されている
		// そのため、TagNameをキーとしてTagIDを取得することができる
		tagID := tagsMap[skill.TagName]
		skillsDto[i] = userDomain.SkillDto{
			TagID:      tagID,
			Evaluation: skill.Evaluation,
			Year:       skill.Years,
		}
	}

	careersDto := make([]userDomain.CareerDto, len(request.Careers))
	for i, career := range request.Careers {
		careersDto[i] = userDomain.CareerDto{
			Detail:    career.Detail,
			StartYear: career.StartYear,
			EndYear:   career.EndYear,
		}
	}

	return userDomain.Create(
		request.Name,
		request.Email,
		request.Password,
		request.Profile,
		careersDto,
		skillsDto,
	)
}

func (uc *SaveUserUseCase) checkUserExists(ctx context.Context, name string) error {
	_, err := uc.userRepo.FindByName(ctx, name)
	return err
}

func (uc *SaveUserUseCase) processTags(ctx context.Context, skills []SkillInputDto) (map[string]string, error) {
	// リクエストからユニークなタグ名のスライスを抽出
	uniqueTagNames := extractUniqueTagNames(skills)

	// リクエストのユニークなタグ名を、DBから取得してくる
	existingTags, err := uc.tagRepo.FindByNames(ctx, uniqueTagNames)
	if err != nil {
		return nil, errDomain.NotFoundErr
	}

	// 1. 既存のタグをマップに追加
	tagsMap := make(map[string]string)
	for _, tag := range existingTags {
		tagsMap[tag.Name()] = tag.ID()
	}

	// 2. 既存のタグのマップリクエストからのユニークなタグ名がない場合は、DBに新規登録。そして、マップにも追加
	for _, tagName := range uniqueTagNames {
		if _, exists := tagsMap[tagName]; !exists {
			// タグが存在しない場合は新規登録
			newTag, err := tagDomain.NewTag(tagName)
			if err != nil {
				return nil, errDomain.NewError("新しいタグの作成に失敗しました: " + err.Error())
			}
			// タグをDBに保存
			if err = uc.tagRepo.Store(ctx, newTag); err != nil {
				return nil, errDomain.NewError("タグの保存に失敗しました: " + err.Error())
			}
			// タグをマップに追加
			tagsMap[tagName] = newTag.ID()
		}
	}

	// 既存のDBに入っているタグ名と、今回リクエストで送られてきたタグ名のマップを返却
	return tagsMap, nil
}

// extractUniqueTagNames はスキルのスライスからユニークなタグ名のスライスを抽出して返します。
func extractUniqueTagNames(skills []SkillInputDto) []string {
	uniqueTags := make(map[string]struct{}) // ユニークなタグ名を保持するマップ

	// スキルのスライスをループして、タグ名をマップのキーとして追加
	// マップのキーは一意であるため、重複するタグ名は除外されます。
	for _, skill := range skills {
		uniqueTags[skill.TagName] = struct{}{}
	}

	// ユニークなタグ名のスライスを準備
	tagNames := make([]string, 0, len(uniqueTags))
	for tagName := range uniqueTags {
		tagNames = append(tagNames, tagName)
	}

	return tagNames
}
