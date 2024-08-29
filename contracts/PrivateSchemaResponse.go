package contracts

// Object
type PrivateSchemaResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	CreatedByType string `json:"created_by_type"`
	CreatedByID   string `json:"created_by_id"`
	ProjectID     string `json:"project_id"`
	CreatedByName string `json:"created_by_name"`
	Schema        string `json:"schema"`
	Version       int    `json:"version"`
}
