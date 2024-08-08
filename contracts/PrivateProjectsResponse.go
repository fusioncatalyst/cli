package contracts

type PrivateProjectsResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	IsPrivate     bool   `json:"is_private"`
	Description   string `json:"description"`
	CreatedByType string `json:"created_by_type"`
	CreatedByID   string `json:"created_by_id"`
	CreatedByName string `json:"created_by_name"`
	Schemas       int64  `json:"schemas"`
}
