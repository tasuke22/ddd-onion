package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuke/go-onion/presentation/settings"
	userUseCase "github.com/tasuke/go-onion/usecase/user"
)

type handler struct {
	saveUserUseCase *userUseCase.SaveUserUseCase
}

func NewHandler(
	saveUserUseCase *userUseCase.SaveUserUseCase,
) handler {
	return handler{
		saveUserUseCase: saveUserUseCase,
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
