package service

type MachineTypeResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type WashingMachineRequest struct {
	Name          string `json:"name"`
	MachineTypeID int    `json:"machine_type_id"`
}

type WashingMachineResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type WorkingMachineRequest struct {
	Coin             int    `json:"coin"`
	NameMachine      string `json:"name_machine"`
	WashingMachineID uint   `json:"washing_machine_id"`
	Start            string `json:"start"`
}

type WorkingMachineResponse struct {
	ID          int    `json:"id"`
	NameMachine string `json:"name_machine"`
	Status      string `json:"status"`
}

type DataWashingMachineResponse struct {
	Pagination PaginationResponse       `json:"pagination"`
	Data       []WashingMachineResponse `json:"washing_machine"`
}

type DataWorkingMachineResponse struct {
	Pagination PaginationResponse       `json:"pagination"`
	Data       []WorkingMachineResponse `json:"working_machine"`
}

type ParamsFilter struct {
	Q     string
	Limit int
	Page  int
}

type PaginationResponse struct {
	Page      int   `json:"page"`
	TotalRow  int64 `json:"total_row"`
	TotalPage int   `json:"total_page"`
}

type LaundromatService interface {
	GetAllMachineType() ([]MachineTypeResponse, error)
	CreateWashingMachine(WashingMachineRequest) (string, error)
	GetAlleWashingMachine(ParamsFilter) (*DataWashingMachineResponse, error)
	CreateWorkingMachine(WorkingMachineRequest) (string, error)

	GetAllCheckWorkingMachine(ParamsFilter) (*DataWorkingMachineResponse, error)
}
