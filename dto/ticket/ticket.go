package ticketdto

import "landtick/models"

type CreateTicketRequest struct {
	Name_Train             string `json:"name_train" form:"name_train" validate:"required"`
	TicketID               int    `json:"ticket_id" form:"ticket_id"`
	Type_Train             string `json:"type_train" form:"type_train" validate:"required"`
	Start_Date             string `json:"start_date" form:"start_date" validate:"required"`
	Start_Station          string `json:"start_station" form:"start_station" validate:"required"`
	Start_Time             string `json:"start_time" form:"start_time" validate:"required"`
	Destination_Station_Id string `json:"destination_station_id" form:"destination_station_id" validate:"required"`
	Arrival_Time           string `json:"arrival_time" form:"arrival_time" validate:"required"`
	Price                  string `json:"price" form:"price" validate:"required"`
	Qty                    string `json:"qty" form:"qty" validate:"required"`
}

type UpdateTicketRequest struct {
	Name_Train             string `json:"name_train" form:"name_train"`
	TicketID               int    `json:"ticket_id" form:"ticket_id"`
	Type_Train             string `json:"type_train" form:"type_train"`
	Start_Date             string `json:"start_date" form:"start_date"`
	Start_Station          string `json:"start_station" form:"start_station"`
	Start_Time             string `json:"start_time" form:"start_time"`
	Destination_Station_Id string `json:"destination_station_id" form:"destination_station_id"`
	Arrival_Time           string `json:"arrival_time" form:"arrival_time"`
	Price                  string `json:"price" form:"price"`
	Qty                    string `json:"qty" form:"qty"`
}

type TicketResponse struct {
	ID                     int                 `json:"id"`
	TicketID               int                 `json:"ticket_id"`
	Name_Train             string              `json:"name_train"`
	Type_Train             string              `json:"type_train"`
	Start_Date             string              `json:"start_date"`
	Start_Station          string              `json:"start_station"`
	Start_Time             string              `json:"start_time" `
	Destination_Station_Id string              `json:"destination_station_id" `
	Arrival_Time           string              `json:"arrival_time"`
	Price                  string              `json:"price"`
	Qty                    string              `json:"qty"`
	User                   models.UserResponse `json:"user"`
	Transaction            models.Transaction  `json:"transaction"`
}
