package stationdto

type CreateStationRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type StationResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
