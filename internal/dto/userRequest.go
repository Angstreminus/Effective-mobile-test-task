package dto

type UserRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type UserEditRequest struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
	IsDeleted   bool   `json:"is_deleted"`
}
