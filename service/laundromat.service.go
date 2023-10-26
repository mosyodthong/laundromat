package service

import (
	"errors"
	"laundromat/repository"
	"strings"
	"time"
)

type laundromatService struct {
	laundromatRepo repository.LaundromatRepository
}

func NewOrderService(laundromatRepo repository.LaundromatRepository) laundromatService {
	return laundromatService{laundromatRepo: laundromatRepo}
}

func (s laundromatService) GetAllMachineType() ([]MachineTypeResponse, error) {

	machineTypes := []MachineTypeResponse{}

	dataMachineType, err := s.laundromatRepo.GetAllMachineTypeRepo()

	if err != nil {
		return nil, err
	}

	for _, data := range dataMachineType {

		machineType := MachineTypeResponse{
			ID:   int(data.ID),
			Name: data.Name,
		}

		machineTypes = append(machineTypes, machineType)

	}

	return machineTypes, nil

}

func (s laundromatService) CreateWashingMachine(request WashingMachineRequest) (string, error) {

	washing := repository.WashingMachine{
		Name:          request.Name,
		MachineTypeID: uint(request.MachineTypeID),
	}

	message, err := s.laundromatRepo.CreateWashingMachineRepo(washing)
	if err != nil {
		return "", err
	}

	return message, nil

}

func (s laundromatService) GetAlleWashingMachine(params ParamsFilter) (*DataWashingMachineResponse, error) {

	qName := strings.ReplaceAll(params.Q, " ", "") //ตัดช่องว่างออกทั้งหมด

	filter := repository.FilterGetAll{
		Q:     qName,
		Limit: params.Limit,
		Page:  params.Page,
	}

	data_all, pagination, err := s.laundromatRepo.GetAlleWashingMachineRepo(filter)
	if err != nil {
		return nil, err
	}

	washingMachineResponses := []WashingMachineResponse{}
	for _, machine := range data_all {

		washingMachineResponse := WashingMachineResponse{
			ID:   int(machine.ID),
			Name: machine.Name,
			Type: machine.MachineType.Name,
		}
		washingMachineResponses = append(washingMachineResponses, washingMachineResponse)

	}

	paginationResponse := PaginationResponse{
		Page:      pagination.Page,
		TotalRow:  pagination.TotalRow,
		TotalPage: pagination.TotalPage,
	}

	Response := DataWashingMachineResponse{
		Data:       washingMachineResponses,
		Pagination: paginationResponse,
	}

	return &Response, nil
}

func (s laundromatService) CreateWorkingMachine(request WorkingMachineRequest) (string, error) {

	checkWorking, err := s.laundromatRepo.CheckWorkingMachineRepo(int(request.WashingMachineID))
	if err != nil {
		return "", err
	}

	if checkWorking != nil {
		return "", errors.New("เครื่องไม่ว่าง")
	}

	if request.Coin != 10 {
		return "", errors.New("กรุณาหยอดเหรียญ 10 บาท")
	}

	oneHourLater := time.Now().Add(time.Hour)

	layoutDate := "2006-01-02 15:04" // รูปแบบวันที่ตาม ปี/เดือน/วัน (dd/mm/yyyy)

	//โหลดโซนเวลาของประเทศไทย
	timeZoneThai, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return "", err
	}

	startDate, err := time.Parse(layoutDate, time.Now().Format("2006-01-02 15:04"))
	if err != nil {
		return "", err
	}

	endDate, err := time.Parse(layoutDate, oneHourLater.Format("2006-01-02 15:04"))
	if err != nil {
		return "", err
	}

	startDateThai := time.Date(startDate.Year(),
		startDate.Month(),
		startDate.Day(),
		startDate.Hour(),
		startDate.Minute(),
		0,
		0,
		timeZoneThai)

	EndDateThai := time.Date(endDate.Year(),
		endDate.Month(),
		endDate.Day(),
		endDate.Hour(),
		endDate.Minute(),
		0,
		0,
		timeZoneThai)

	// fmt.Println("startDateThai", startDateThai.UTC())
	// fmt.Println("EndDateThai", EndDateThai.UTC())

	working := repository.WorkingMachine{
		NameMachine:      request.NameMachine,
		WashingMachineID: request.WashingMachineID,
		Strat:            startDateThai.UTC(),
		End:              EndDateThai.UTC(),
	}

	message, err := s.laundromatRepo.CreateWorkingMachineRepo(working)
	if err != nil {
		return "", err
	}

	return message, nil

}

func (s laundromatService) GetAllCheckWorkingMachine(params ParamsFilter) (*DataWorkingMachineResponse, error) {

	qName := strings.ReplaceAll(params.Q, " ", "") //ตัดช่องว่างออกทั้งหมด

	filter := repository.FilterGetAll{
		Q:     qName,
		Limit: params.Limit,
		Page:  params.Page,
	}

	data_all, pagination, err := s.laundromatRepo.GetAllCheckWorkingMachineRepo(filter)
	if err != nil {
		return nil, err
	}

	currentDateTimeSrt := time.Now().Format("2006-01-02 15:04:05")

	workMachineResponses := []WorkingMachineResponse{}
	for _, machine := range data_all {

		layout := "2006-01-02 15:04:05"

		endDateTimeStr := machine.End.Format("2006-01-02 15:04:05")

		currentDateTime, err := time.Parse(layout, currentDateTimeSrt)
		if err != nil {
			return nil, err
		}

		endDateTime, err := time.Parse(layout, endDateTimeStr)
		if err != nil {
			return nil, err
		}

		difference := endDateTime.Sub(currentDateTime)

		var status string
		// fmt.Println(" currentDateTimeSrt", currentDateTimeSrt)
		// fmt.Println("", machine.End.Format("2006-01-02 15:04:05"))
		if difference.Minutes() < 1 && machine.Strat.Format("2006-01-02 15:04:05") < currentDateTimeSrt && machine.End.Format("2006-01-02 15:04:05") > currentDateTimeSrt { // currentDateTime เป็นน้อยกว่า 1 นาทีของ endDateTime
			status = "ส่งสัญญาณเตือน"

		} else if machine.Strat.Format("2006-01-02 15:04:05") < currentDateTimeSrt && machine.End.Format("2006-01-02 15:04:05") > currentDateTimeSrt {

			status = "เครื่องทำงานอยู่"

		} else {

			status = "เครื่องว่าง"

		}

		WorkMachineResponse := WorkingMachineResponse{
			ID:          int(machine.ID),
			NameMachine: machine.NameMachine,
			Status:      status,
		}
		workMachineResponses = append(workMachineResponses, WorkMachineResponse)

	}

	paginationResponse := PaginationResponse{
		Page:      pagination.Page,
		TotalRow:  pagination.TotalRow,
		TotalPage: pagination.TotalPage,
	}

	Response := DataWorkingMachineResponse{
		Data:       workMachineResponses,
		Pagination: paginationResponse,
	}

	return &Response, nil
}
