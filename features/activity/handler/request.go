package handler

type ActivityRequest struct {
	Judul     string `json:"judul" form:"judul"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
}