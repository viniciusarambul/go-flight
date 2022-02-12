package plane

type CreatePlaneInput struct {
	Model       string `json:"model"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreatePlaneOutput struct {
	ID          string `json:"id"`
	Model       string `json:"model"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ShowPlaneOutputDto struct {
	ID          string `json:"id"`
	Model       string `json:"model"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
