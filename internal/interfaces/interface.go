package interfaces

import "SCA/internal/storage"

type CatStorage interface {
	AddCat(storage.Cat) error
	DeleteCat(int) error
	UpdateSalary(int, int) error
	ListCat() ([]storage.Cat, error)
	GetCat(int) (storage.Cat, error)
}

type MissionStorage interface {
	AddMission([]string) error
	DeleteMission(int) error
	UpdateMission(int) error
	AssignMission(int, int) error
	ListMission() ([]storage.Mision, error)
	GetMission(int) (storage.Mision, error)
	UpdateMissionTarget([]string) error
}

type TargetStorage interface {
	AddTarget(storage.Target) error
	UpdateTarget(string) error
	UpdateTargetNote(string, string) error
	DeleteTarget(string) error
}

type Storage interface {
	CatStorage
	MissionStorage
	TargetStorage
}
