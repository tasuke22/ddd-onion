package user

import (
	"context"
	tagDomain "github.com/tasuke/go-onion/domain/tag"
	userDomain "github.com/tasuke/go-onion/domain/user"
	"github.com/tasuke/go-pkg/ulid"
)

type UpdateUserUseCase struct {
	userRepo userDomain.UserRepository
	tagRepo  tagDomain.TagRepository
}

func NewUpdateUserUseCase(
	userRepo userDomain.UserRepository,
	tagRepo tagDomain.TagRepository,
) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepo,
		tagRepo,
	}
}

type UpdateUserUseCaseInputDto struct {
	UserID   string            `json:"user_id"`
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Password string            `json:"password"`
	Profile  string            `json:"profile"`
	Skills   []UpdateSkillDto  `json:"skills"`
	Careers  []UpdateCareerDto `json:"careers"`
}

type UpdateCareerDto struct {
	ID        string `json:"id"`
	Detail    string `json:"detail"`
	StartYear int32  `json:"start_year"`
	EndYear   int32  `json:"end_year"`
}

type UpdateSkillDto struct {
	ID         string `json:"id"`
	TagID      string `json:"tag_id"`
	Evaluation int32  `json:"evaluation"`
	Years      int32  `json:"years"`
}

func (uc *UpdateUserUseCase) Run(ctx context.Context, input *UpdateUserUseCaseInputDto) error {
	// ユーザーの存在確認
	user, err := uc.userRepo.FindByUserID(ctx, input.UserID)
	if err != nil {
		return err
	}

	// スキルの更新または追加
	var skills []userDomain.Skill
	for _, skillReq := range input.Skills {
		var skill *userDomain.Skill
		found := false

		// 既存のスキルを確認
		for _, existingSkill := range user.Skills() {
			// 更新
			if existingSkill.ID() == skillReq.ID {
				found = true
				skill, err = userDomain.NewSkill(
					skillReq.ID,
					skillReq.TagID,
					skillReq.Evaluation,
					skillReq.Years,
				)
				break
			}
		}

		// 新規作成
		if !found {
			skill, err = userDomain.NewSkill(
				ulid.NewULID(),
				skillReq.TagID,
				skillReq.Evaluation,
				skillReq.Years,
			)
			if err != nil {
				return err
			}
		}

		if skill != nil {
			skills = append(skills, *skill)
		}
	}

	// キャリアの更新または追加
	var careers []userDomain.Career
	for _, careerReq := range input.Careers {
		var career *userDomain.Career
		found := false

		// 既存のキャリアを確認
		for _, existingCareer := range user.Careers() {
			// 更新
			if existingCareer.ID() == careerReq.ID {
				found = true
				career, err = userDomain.NewCareer(
					careerReq.ID,
					careerReq.Detail,
					careerReq.StartYear,
					careerReq.EndYear,
				)
				break
			}
		}

		// 新規作成
		if !found {
			career, err = userDomain.NewCareer(
				ulid.NewULID(),
				careerReq.Detail,
				careerReq.StartYear,
				careerReq.EndYear,
			)
			if err != nil {
				return err
			}
		}

		if career != nil {
			careers = append(careers, *career)
		}
	}

	// ユーザーオブジェクトのスキルとキャリアを更新
	if err := user.Update(
		input.Name,
		input.Email,
		userDomain.UserPassword(input.Password),
		input.Profile,
		skills,
		careers,
	); err != nil {
		return err
	}

	// 更新されたユーザー情報をデータベースに保存
	if err := uc.userRepo.UpdateUser(ctx, user); err != nil {
		return err
	}
	return nil
}
