package models

type Ticket struct {
	ID                     int    `json:"id" gorm:"primary_key:auto_increment"`
	TicketID               int    `json:"ticket_id" gorm:"type: varchar(255)"`
	Name_Train             string `json:"name_train" gorm:"type: varchar(255)"`
	Type_Train             string `json:"type_train" gorm:"type: varchar(255)"`
	Start_Date             string `json:"start_date" gorm:"type: varchar(255)"`
	Start_Station          string `json:"start_station" gorm:"type: int"`
	Start_Time             string `json:"start_time" gorm:"type:string"`
	Destination_Station_Id string `json:"destination_station_id" gorm:"type: int"`
	Arrival_Time           string `json:"arrival_time" gorm:"type:string"`
	Price                  string `json:"price" gorm:"type: int"`
	Qty                    string `json:"qty" gorm:"type:int"`
}

type TicketResponse struct {
	ID                     int    `json:"id"`
	TicketID               int    `json:"ticket_id"`
	Name_Train             string `json:"name_train"`
	Type_Train             string `json:"type_train"`
	Start_Date             string `json:"start_date"`
	Start_Station          string `json:"start_station"`
	Start_Time             string `json:"start_time"`
	Destination_Station_Id string `json:"destination_station_id"`
	Arrival_Time           string `json:"arrival_time"`
	Price                  string `json:"price"`
	Qty                    string `json:"qty"`
}

func (TicketResponse) TableName() string {
	return "Tickets"
}
