package server

import (
	"SCA/internal/storage"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Bread struct {
	Name string `json:name`
}

func (s *Server) addCat(c *fiber.Ctx) error {
	sal, err := strconv.Atoi(c.Params("salary"))
	if err != nil {
		c.SendStatus(422)
		return err
	}

	exp, err := strconv.Atoi(c.Params("experience"))
	if err != nil {
		c.SendStatus(422)
		return err
	}

	cat := storage.Cat{
		Name: c.Params("name"),
		Salary: sal,
		Experience: exp,
		Bread: c.Params("bread"),
	}

	catsBody := fiber.Get("https://api.thecatapi.com/v1/breeds")
	_, data, _ := catsBody.Bytes()
	var cats []Bread
	err = json.Unmarshal(data, &cats)
	fmt.Println(cats[0])
	if err != nil {
		c.SendStatus(500)
		return err
	}

	valid := false
	for _, bread := range cats {
		if strings.ToLower(bread.Name) == strings.ToLower(cat.Bread) {
			valid = true
			break
		}
	}
	if !valid {
		fmt.Println("EGOR")
		c.SendStatus(422)
		return nil
	}

	err = s.storage.AddCat(cat)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) deleteCat(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("catID"))
	if err != nil {
		c.SendStatus(422)
		return err
	}

	err = s.storage.DeleteCat(id)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) listCat(c *fiber.Ctx) error {
	cats, err := s.storage.ListCat()
	if err != nil {
		c.SendStatus(500)
		return err
	}

	err = c.JSON(cats)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) getCat(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.SendStatus(422)
		return err
	}

	cats, err := s.storage.GetCat(id)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	err = c.JSON(cats)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) updateSalary(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("catID"))
	if err != nil {
		c.SendStatus(422)
		return err
	}

	salary, err := strconv.Atoi(c.Params("salary"))
	if err != nil {
		c.SendStatus(422)
		return err
	}

	err = s.storage.UpdateSalary(id, salary)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) addMision(c *fiber.Ctx) error {
	names := strings.Split(c.Queries()["names"], ",")
	countrys := strings.Split(c.Queries()["countrys"], ",")
	if len(names) != len(countrys) || len(names) == 0 || len(names) > 3 {
		c.SendStatus(422)		
		return nil
	}

	var targets []storage.Target = make([]storage.Target, 0, 1)
	for i:=0; i<len(names); i++ {
		targets = append(targets, storage.Target{Name: names[i], Country: countrys[i], Complete: false})
	}
	for _, target := range targets {
		err := s.storage.AddTarget(target)
		if err != nil {
			c.SendStatus(500)
			return err
		}
	}
	
	err := s.storage.AddMission(names)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) listMission(c *fiber.Ctx) error {
	missions, err := s.storage.ListMission()
	if err != nil {
		c.SendStatus(500)
		return err
	}

	err = c.JSON(missions)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) deleteMission(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.SendStatus(422)
		return err
	}

	err = s.storage.DeleteMission(id)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) assignMission(c *fiber.Ctx) error {
	idCat, err := strconv.Atoi(c.Params("catID"))
	if err != nil {
		c.SendStatus(422)
		return err
	}

	idMission, err := strconv.Atoi(c.Params("missionID"))
	if err != nil {
		c.SendStatus(422)
		return err
	}

	err = s.storage.AssignMission(idCat, idMission)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) updateMission(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.SendStatus(422)
		return err
	}
	_ = id

	err = s.storage.UpdateMission(id)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) addTarget(c *fiber.Ctx) error {
	target := storage.Target{
		Name: c.Params("name"),
		Country: c.Params("country"),
	}
	
	id, err := strconv.Atoi(c.Params("idMission"))
	if err != nil {
		c.SendStatus(422)
		return err
	}

	mission, err := s.storage.GetMission(id)
	if err != nil {
		c.SendStatus(422)
		return err
	}

	if len(mission.Targets) == 3 {
		c.SendStatus(412)
		return nil
	}
	var task = []string{target.Name} 
	for _, val := range mission.Targets {
		task = append(task, val.Name)
	}
	err = s.storage.AddTarget(target)
	if err != nil {
		c.SendStatus(500)
		return err
	}
	
	err = s.storage.UpdateMissionTarget(task)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) deleteTarget(c *fiber.Ctx) error {
	err := s.storage.DeleteTarget(c.Params("nameTarget"))
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) updateTargetNote(c *fiber.Ctx) error {
	err := s.storage.UpdateTargetNote(c.Params("nameTarget"), c.Params("note"))
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}

func (s *Server) updateTarget(c *fiber.Ctx) error {
	err := s.storage.UpdateTarget(c.Params("nameTarget"))
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return nil
}