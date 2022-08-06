package request

type CreateProjectsParams struct {
	Name    string `json:"name" form:"name" binding:"required"`
	About   string `json:"about" form:"about" binding:"required"`
	Website string `json:"website" form:"website"`
	Repo    string `json:"repo" form:"repo"`
	Preview string `json:"preview" form:"preview"`
	Readme  string `json:"readme" form:"readme"`
}
