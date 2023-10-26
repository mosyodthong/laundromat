package repository

import (
	"errors"
	"math"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type laundromatRepositoryPG struct {
	db *gorm.DB
}

func NewLaundromatRepositoryPG(db *gorm.DB) LaundromatRepository {

	db.AutoMigrate(MachineType{})
	db.AutoMigrate(WashingMachine{})
	db.AutoMigrate(WorkingMachine{})
	SeedMachineType(db)
	return laundromatRepositoryPG{db: db}
}

func SeedMachineType(db *gorm.DB) error {

	var machineType = []MachineType{{Name: "ซักผ้า"}, {Name: "ปั่นผ้า"}, {Name: "อบผ้า"}}

	tx := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&machineType)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r laundromatRepositoryPG) GetAllMachineTypeRepo() ([]MachineType, error) {

	machineType := []MachineType{}

	tx := r.db.Order("id asc").Find(&machineType)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return machineType, nil
}

func (r laundromatRepositoryPG) CreateWashingMachineRepo(washingMachine WashingMachine) (string, error) {

	tx := r.db.Create(&washingMachine)
	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected != 1 {
		return "", errors.New("data is not creating")
	}
	return "Created.", nil

}

func (r laundromatRepositoryPG) GetAlleWashingMachineRepo(filter FilterGetAll) ([]WashingMachine, *Pagination, error) {
	washingMachine := []WashingMachine{}

	if filter.Page == 0 {
		filter.Page = 1
	}

	//จำนวนข้อมูล
	var count int64

	//แสดงข้อมูลในหนึ่งหน้า
	if filter.Limit == 0 {
		r.db.Model(&washingMachine).Count(&count)
		filter.Limit = int(count)
	}

	if filter.Q != "" {
		tx := r.db.Preload(clause.Associations).
			Where(" LOWER(REPLACE(name, ' ', '')) LIKE ?   ", "%"+strings.ToLower(filter.Q)+"%").
			Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit).
			Order("id asc").Find(&washingMachine)

		if tx.Error != nil {
			return nil, nil, tx.Error
		}

		//หาจำนวนข้อมูล
		tc := r.db.Model(&WashingMachine{}).Preload(clause.Associations).Where(" LOWER(REPLACE(name, ' ', '')) LIKE ?   ", "%"+strings.ToLower(filter.Q)+"%").Order("id asc").Count(&count)
		if tc.Error != nil {
			return nil, nil, tc.Error
		}

	} else {

		tx := r.db.Preload(clause.Associations).
			Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit).
			Order("id asc").Find(&washingMachine)

		if tx.Error != nil {
			return nil, nil, tx.Error
		}

		//หาจำนวนข้อมูล
		tc := r.db.Model(&WashingMachine{}).Preload(clause.Associations).Count(&count)
		if tc.Error != nil {
			return nil, nil, tc.Error
		}
	}

	count_page := math.Ceil(float64(count) / float64(filter.Limit))
	if int(count_page) < 0 {
		count_page = 0
	}
	pagination := Pagination{
		Page:      filter.Page,
		TotalRow:  count,
		TotalPage: int(count_page),
	}

	return washingMachine, &pagination, nil
}

func (r laundromatRepositoryPG) CreateWorkingMachineRepo(workingMachine WorkingMachine) (string, error) {

	tx := r.db.Create(&workingMachine)
	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected != 1 {
		return "", errors.New("data is not creating")
	}
	return "Created.", nil

}

func (r laundromatRepositoryPG) GetAllCheckWorkingMachineRepo(filter FilterGetAll) ([]WorkingMachine, *Pagination, error) {
	workingMachine := []WorkingMachine{}

	if filter.Page == 0 {
		filter.Page = 1
	}

	//จำนวนข้อมูล
	var count int64

	//แสดงข้อมูลในหนึ่งหน้า
	if filter.Limit == 0 {
		r.db.Model(&workingMachine).Count(&count)
		filter.Limit = int(count)
	}

	if filter.Q != "" {
		tx := r.db.Preload(clause.Associations).
			Where(" LOWER(REPLACE(name_machine, ' ', '')) LIKE ?   ", "%"+strings.ToLower(filter.Q)+"%").
			Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit).
			Order("id asc").Find(&workingMachine)

		if tx.Error != nil {
			return nil, nil, tx.Error
		}

		//หาจำนวนข้อมูล
		tc := r.db.Model(&WorkingMachine{}).Where(" LOWER(REPLACE(name_machine, ' ', '')) LIKE ?   ", "%"+strings.ToLower(filter.Q)+"%").Order("id asc").Count(&count)
		if tc.Error != nil {
			return nil, nil, tc.Error
		}

	} else {

		tx := r.db.Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit).
			Order("id asc").Find(&workingMachine)

		if tx.Error != nil {
			return nil, nil, tx.Error
		}

		//หาจำนวนข้อมูล
		tc := r.db.Model(&WorkingMachine{}).Count(&count)
		if tc.Error != nil {
			return nil, nil, tc.Error
		}
	}

	count_page := math.Ceil(float64(count) / float64(filter.Limit))
	if int(count_page) < 0 {
		count_page = 0
	}
	pagination := Pagination{
		Page:      filter.Page,
		TotalRow:  count,
		TotalPage: int(count_page),
	}

	return workingMachine, &pagination, nil
}

func (r laundromatRepositoryPG) CheckWorkingMachineRepo(id int) (*WorkingMachine, error) {

	workingMachine := WorkingMachine{}

	currentDateTimeSrt := time.Now().Format("2006-01-02 15:04:05")

	tx := r.db.Where("washing_machine_id = ? AND strat < ? AND \"end\" > ?", id, currentDateTimeSrt, currentDateTimeSrt).Order("id desc").First(&workingMachine)

	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	} else if tx.Error != nil {
		return nil, tx.Error
	}

	return &workingMachine, nil
}
