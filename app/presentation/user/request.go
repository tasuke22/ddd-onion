package user

type SaveUserRequest struct {
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Profile  string          `json:"profile"`
	Skills   []SkillRequest  `json:"skills"`
	Careers  []CareerRequest `json:"careers"`
}

type CareerRequest struct {
	Detail    string `json:"detail"`
	StartYear int32  `json:"start_year"`
	EndYear   int32  `json:"end_year"`
}

type SkillRequest struct {
	Evaluation int32  `json:"evaluation"`
	Years      int32  `json:"years"`
	TagName    string `json:"tag_name"`
}
