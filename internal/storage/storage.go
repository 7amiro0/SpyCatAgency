package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	// Cat quary
	addCat       = "insert into cat (name, salary, experience, bread) values ($1, $2, $3, $4);"
	listCat      = "select * from cat;"
	getCat       = "select * from cat where id=$1;"
	deleteCat    = "delete from cat where id=$1;"
	salaryUpdate = "update cat set salary=$2 where cat.id=$1;"

	// Mission quary
	addMision     = "insert into mission (targets, complete, catID) values ($1, $2, $3);"
	listMission   = "select * from mission;"
	deleteM       = "delete from mission where (id=$1 and catID=0);"
	updateMission = "update mission set complete=$2, catID=0 where mission.id=$1;"
	assaignCat    = "update mission set catID=$2 where mission.id=$1;"
	getMission = "select * from mission where mission.id=$1"
	updateTargetList = "update mission set targets=$1;"

	// Target quary
	addTarget    = "insert into targets (name, country, notes, complete) values ($1, $2, $3, $4);"
	updateTarget = "update targets set complete=$2 where targets.name=$1;"
	updateTargetNote = "update targets set notes=$3 where (targets.name=$1 and targets.complete=$2);"
	listTargetByName  = "select * from targets where targets.name=$1;"
	deleteTarget =  "delete from targets where targets.name=$1;"
)

type Storage struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func New(ctx context.Context) *Storage {
	return &Storage{
		ctx: ctx,
	}
}

func (s *Storage) Conn() (err error) {
	s.db, err = pgxpool.New(s.ctx, initDB().getLink())
	return
}

func (s *Storage) Close() error {
	s.db.Close()
	return nil
}

func (s *Storage) AddCat(cat Cat) (err error) {
	_, err = s.db.Exec(s.ctx, addCat, cat.Name, cat.Salary, cat.Experience, cat.Bread)
	return err
}

func listening(rows pgx.Rows) (cats []Cat, err error) {
	var cat Cat

	for rows.Next() {
		err = rows.Scan(
			&cat.id,
			&cat.Name,
			&cat.Salary,
			&cat.Experience,
			&cat.Bread,
		)
		if err != nil {
			return nil, err
		}

		cats = append(cats, cat)
	}

	return cats, err
}

func (s *Storage) ListCat() ([]Cat, error) {
	rows, err := s.db.Query(s.ctx, listCat)
	if err != nil {
		return nil, err
	}

	return listening(rows)
}

func (s *Storage) UpdateSalary(id, salary int) error {
	_, err := s.db.Exec(s.ctx, salaryUpdate, id, salary)
	return err
}

func (s *Storage) GetCat(id int) (Cat, error) {
	rows, err := s.db.Query(s.ctx, getCat, id)
	if err != nil {
		return Cat{}, err
	}

	cats, err := listening(rows)
	if len(cats) == 0 || err != nil {
		return Cat{}, err
	}

	return cats[0], nil
}

func (s *Storage) DeleteCat(id int) error {
	_, err := s.db.Exec(s.ctx, deleteCat, id)
	return err
}

func (s *Storage) AddTarget(target Target) error {
	_, err := s.db.Exec(s.ctx, addTarget, target.Name, target.Country, target.Notes, target.Complete)
	return err
}

func (s *Storage) AddMission(targets []string) error {
	_, err := s.db.Exec(s.ctx, addMision, targets, false, 0)
	return err
}

func listingMission(rows pgx.Rows, ctx context.Context, db *pgxpool.Pool) ([]Mision, error) {
	var (
		mission  Mision
		missions []Mision
		target	Target
		targets []Target
		test    []string
		err 	error
	)

	for rows.Next() {
		err = rows.Scan(
			&mission.id,
			&mission.Cat,
			&mission.Complete,
			&test,
		)
		if err != nil {
			return nil, err
		}

		for _, x := range test {
			rowsTarget, err := db.Query(ctx, listTargetByName, x)
			if err != nil {
				return nil, err
			}

			for rowsTarget.Next() {
				err = rowsTarget.Scan(
					&target.Name,
					&target.Country,
					&target.Notes,
					&target.Complete,
				)
				if err != nil {
					return nil, err
				}
			
				targets = append(targets, target)
			}
		}

		mission.Targets = targets
		missions = append(missions, mission)
	}

	return missions, err
}

func (s *Storage) UpdateMissionTarget(targets []string) error {
	_, err := s.db.Exec(s.ctx, updateTargetList, targets)
	return err
}

func (s *Storage) GetMission(id int) (Mision, error) {
	rows, err := s.db.Query(s.ctx, getCat, id)
	if err != nil {
		return Mision{}, err
	}

	mission, err := listingMission(rows, s.ctx, s.db)
	if len(mission) == 0 || err != nil {
		return Mision{}, err
	}

	return mission[0], nil
}

func (s *Storage) ListMission() ([]Mision, error) {
	rows, err := s.db.Query(s.ctx, listMission)
	if err != nil {
		return nil, err
	}
	
	return listingMission(rows, s.ctx, s.db)
}

func (s *Storage) DeleteMission(id int) error {
	_, err := s.db.Exec(s.ctx, deleteM, id)
	return err
}

func (s *Storage) AssignMission(cat, mission int) error {
	_, err := s.db.Exec(s.ctx, assaignCat, cat, mission)
	return err
}

func (s *Storage) UpdateMission(id int) error {
	_, err := s.db.Exec(s.ctx, updateMission, id, true)
	return err
}

func (s *Storage) UpdateTargetNote(name string, text string) error {
	_, err := s.db.Exec(s.ctx, updateTargetNote, name, false, text)
	return err
}

func (s *Storage) UpdateTarget(name string) error {
	_, err := s.db.Exec(s.ctx, updateTarget, name, true)
	return err
}

func (s *Storage) DeleteTarget(name string) error {
	_, err := s.db.Exec(s.ctx, deleteTarget, name)
	return err
}