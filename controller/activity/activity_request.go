package activity

type AddActivityRequest struct {
	UserHp      string `json:"user_hp" form:"user_hp" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required,min=10,max=100"`
}

type UpdateActivityRequest struct {
	UserHp      string `json:"user_hp" form:"user_hp" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required,min=10,max=100"`
}
