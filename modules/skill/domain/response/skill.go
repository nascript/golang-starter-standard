package response

type (
	GreenSkillResponse struct {
		ID          string                     `json:"id"`
		Name        string                     `json:"name"`
		Description string                     `json:"description"`
		Groups      *[]GreenSkillGroupResponse `json:"groups"`
	}

	GreenSkillGroupResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
