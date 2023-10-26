package repository

import (
	"time"

	"gorm.io/gorm"
)

type MachineType struct {
	ID   uint   `gorm:"primaryKey;comment:PK"`
	Name string `gorm:"unique;not null;comment:ประเภทการซัก "`
	gorm.Model
}

type WashingMachine struct {
	ID            uint        `gorm:"primaryKey;comment:PK"`
	Name          string      `gorm:"size:255;not null;comment:ชื่อเครื่อง"`
	MachineType   MachineType `gorm:"foreignKey:MachineTypeID"`
	MachineTypeID uint        `gorm:"not null;comment:ประเภทการซัก "`
	gorm.Model
}

type WorkingMachine struct {
	ID               uint      `gorm:"primaryKey;comment:PK"`
	WashingMachineID uint      `gorm:"not null;comment:FK table washing_machine.id "`
	NameMachine      string    `gorm:"size:255;not null;comment:ชื่อเครื่อง"`
	Strat            time.Time `gorm:"not null;comment:เวลาเริ่มต้น "`
	End              time.Time `gorm:"not null;comment:เวลาสิ้นสุด "`
	gorm.Model
}

type FilterGetAll struct {
	Q     string
	Limit int
	Page  int
}

type Pagination struct {
	Page      int
	TotalRow  int64
	TotalPage int
}

type LaundromatRepository interface {
	GetAllMachineTypeRepo() ([]MachineType, error)

	CreateWashingMachineRepo(WashingMachine) (string, error)
	GetAlleWashingMachineRepo(FilterGetAll) ([]WashingMachine, *Pagination, error)

	CreateWorkingMachineRepo(WorkingMachine) (string, error)
	GetAllCheckWorkingMachineRepo(FilterGetAll) ([]WorkingMachine, *Pagination, error)
	CheckWorkingMachineRepo(int) (*WorkingMachine, error)
}
