package repository

import (
	"context"
	"fmt"
	"github.com/tasuke/go-onion/domain/user"
	"github.com/tasuke/go-onion/infrastructure/db"
	"github.com/tasuke/go-onion/infrastructure/db/dbgen"
)

type userRepository struct{}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Save(ctx context.Context, u *user.User) error {
	query := db.GetQuery(ctx)

	if err := query.UpsertUser(ctx, dbgen.UpsertUserParams{
		ID:       u.ID(),
		Name:     u.Name(),
		Email:    u.Email(),
		Password: string(u.Password()),
		Profile:  u.Profile(),
	}); err != nil {
		return err
	}

	for _, career := range u.Careers() {
		if err := query.UpsertCareer(ctx, dbgen.UpsertCareerParams{
			ID:        career.ID(),
			UserID:    u.ID(),
			Detail:    career.Detail(),
			StartYear: career.StartYear(),
			EndYear:   career.EndYear(),
		}); err != nil {
			return err
		}
	}

	for _, skill := range u.Skills() {
		if err := query.UpsertSkill(ctx, dbgen.UpsertSkillParams{
			ID:         skill.ID(),
			UserID:     u.ID(),
			TagID:      skill.TagID(),
			Evaluation: skill.Evaluation(),
			Years:      skill.Years(),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (ur *userRepository) FindByName(ctx context.Context, name string) (*user.User, error) {
	query := db.GetQuery(ctx)
	tempUser, err := query.FindByName(ctx, name)

	if err != nil {
		return nil, err
	}
	fmt.Println(tempUser)

	return nil, err
}

func (ur *userRepository) FindByUserID(ctx context.Context, id string) (*user.User, error) {
	query := db.GetQuery(ctx)
	ud, err := query.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	careers, err := query.FindCareersByUserID(ctx, id)
	var careersDomain []user.Career
	for _, career := range careers {
		cd, err := user.ReconstructCareer(
			career.ID,
			career.Detail,
			career.StartYear,
			career.EndYear,
		)
		if err != nil {
			return nil, err
		}
		careersDomain = append(careersDomain, *cd)
	}
	fmt.Println("careeeeeee")

	skills, err := query.FindSkillsByUserID(ctx, id)
	var skillsDomain []user.Skill
	if err != nil {
		return nil, err
	}
	for _, skill := range skills {
		sd, err := user.ReconstructSkill(
			skill.ID,
			skill.TagID,
			skill.Evaluation,
			skill.Years,
		)
		skillsDomain = append(skillsDomain, *sd)
		if err != nil {
			return nil, err
		}
	}

	userDomain, err := user.ReconstructUser(
		ud.ID,
		ud.Name,
		ud.Email,
		user.UserPassword(ud.Password),
		ud.Profile,
		skillsDomain,
		careersDomain,
	)
	fmt.Println("userDomain")

	return userDomain, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, u *user.User) error {
	query := db.GetQuery(ctx)

	for _, career := range u.Careers() {
		if err := query.UpsertCareer(ctx, dbgen.UpsertCareerParams{
			ID:        career.ID(),
			UserID:    u.ID(),
			Detail:    career.Detail(),
			StartYear: career.StartYear(),
			EndYear:   career.EndYear(),
		}); err != nil {
			return err
		}
	}

	for _, skill := range u.Skills() {
		if err := query.UpsertSkill(ctx, dbgen.UpsertSkillParams{
			ID:         skill.ID(),
			UserID:     u.ID(),
			TagID:      skill.TagID(),
			Evaluation: skill.Evaluation(),
			Years:      skill.Years(),
		}); err != nil {
			return err
		}
	}

	if err := query.UpsertUser(ctx, dbgen.UpsertUserParams{
		ID:       u.ID(),
		Name:     u.Name(),
		Email:    u.Email(),
		Password: string(u.Password()),
		Profile:  u.Profile(),
	}); err != nil {
		return err
	}

	return nil
}
