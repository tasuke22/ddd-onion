package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuke/go-onion/presentation/settings"
	userUseCase "github.com/tasuke/go-onion/usecase/user"
)

type handler struct {
	saveUserUseCase   *userUseCase.SaveUserUseCase
	updateUserUseCase *userUseCase.UpdateUserUseCase
}

func NewHandler(
	saveUserUseCase *userUseCase.SaveUserUseCase,
	updateUserUseCase *userUseCase.UpdateUserUseCase,
) handler {
	return handler{
		saveUserUseCase:   saveUserUseCase,
		updateUserUseCase: updateUserUseCase,
	}
}

func (h handler) SaveUser(ctx *gin.Context) {
	var request SaveUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		settings.ReturnError(ctx, err)
		return
	}

	skillsRequests := h.convertSkillRequestsToSkillInputDtos(request.Skills)
	careersRequests := h.convertCareerRequestsToCareerInputDtos(request.Careers)

	err := h.saveUserUseCase.Run(ctx, &userUseCase.SaveUserUseCaseInputDto{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Profile:  request.Profile,
		Skills:   skillsRequests,
		Careers:  careersRequests,
	})
	if err != nil {
		settings.ReturnError(ctx, err)
		return
	}
	settings.ReturnStatusNoContent(ctx)
}

func (h handler) convertSkillRequestsToSkillInputDtos(requests []SkillRequest) []userUseCase.SkillInputDto {
	skills := make([]userUseCase.SkillInputDto, len(requests))
	for i, req := range requests {
		skills[i] = userUseCase.SkillInputDto{
			Evaluation: req.Evaluation,
			Years:      req.Years,
			TagName:    req.TagName,
		}
	}
	return skills
}

func (h handler) convertCareerRequestsToCareerInputDtos(requests []CareerRequest) []userUseCase.CareerInputDto {
	careers := make([]userUseCase.CareerInputDto, len(requests))
	for i, req := range requests {
		careers[i] = userUseCase.CareerInputDto{
			Detail:    req.Detail,
			StartYear: req.StartYear,
			EndYear:   req.EndYear,
		}
	}
	return careers
}

func (h handler) UpdateUser(ctx *gin.Context) {
	var request UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		settings.ReturnError(ctx, err)
		return
	}

	skillsRequests := h.convertUpdateSkillRequestsToUpdateSkillInputDtos(request.Skills)
	careersRequests := h.convertUpdateCareerRequestsToUpdateCareerInputDtos(request.Careers)

	err := h.updateUserUseCase.Run(ctx, &userUseCase.UpdateUserUseCaseInputDto{
		UserID:   request.UserID,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Profile:  request.Profile,
		Skills:   skillsRequests,
		Careers:  careersRequests,
	})
	if err != nil {
		settings.ReturnError(ctx, err)
		return
	}

	settings.ReturnStatusNoContent(ctx)
}

// UpdateSkillRequest から UpdateSkillDto への変換
func (h handler) convertUpdateSkillRequestsToUpdateSkillInputDtos(skills []UpdateSkillRequest) []userUseCase.UpdateSkillDto {
	var skillsDtos []userUseCase.UpdateSkillDto
	for _, skill := range skills {
		skillDto := userUseCase.UpdateSkillDto{
			ID:         skill.ID,
			TagID:      skill.TagID,
			Evaluation: skill.Evaluation,
			Years:      skill.Years,
		}
		skillsDtos = append(skillsDtos, skillDto)
	}
	return skillsDtos
}

// UpdateCareerRequest から UpdateCareerDto への変換
func (h handler) convertUpdateCareerRequestsToUpdateCareerInputDtos(careers []UpdateCareerRequest) []userUseCase.UpdateCareerDto {
	var careersDtos []userUseCase.UpdateCareerDto
	for _, career := range careers {
		careerDto := userUseCase.UpdateCareerDto{
			ID:        career.ID,
			Detail:    career.Detail,
			StartYear: career.StartYear,
			EndYear:   career.EndYear,
		}
		careersDtos = append(careersDtos, careerDto)
	}
	return careersDtos
}
