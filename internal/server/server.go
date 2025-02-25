package server

import (
	"context"
	"SCA/internal/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	addr string
	ctx  context.Context
	server *fiber.App
	storage interfaces.Storage
}

func New(addr string, ctx context.Context, storage interfaces.Storage) Server {
	return Server{
		addr: addr,
		ctx: ctx,
		server: fiber.New(),
		storage: storage,
	}
}

func (s *Server) setRouter() {
	s.server.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	s.server.Post("/newCat/:name/:salary/:experience/:bread", s.addCat)
	s.server.Get("/cats", s.listCat)
	s.server.Delete("/deleteCat/:catID", s.deleteCat)
	s.server.Post("/updateSalary/:catID/:salary", s.updateSalary)
	s.server.Get("/cat/:id", s.getCat)

	s.server.Post("/newMission", s.addMision)
	s.server.Get("/missions", s.listMission)
	s.server.Delete("/deleteMission/:id", s.deleteMission)
	s.server.Post("/assign/:catID/:missionID", s.assignMission)
	s.server.Post("/updateMission/:id", s.updateMission)

	s.server.Post("/newTarget/:idMission/:name/:country", s.addTarget)
	s.server.Delete("/deleteTarget/:nameTarget", s.deleteTarget)
	s.server.Post("/updateTarget/:nameTarget/", s.updateTarget)
	s.server.Post("/updateNote/:nameTarget/:note", s.updateTargetNote)
}

func (s *Server) Conn() error {
	s.setRouter()
	return s.server.Listen(s.addr)
}

func (s *Server) Close() error {
	return s.server.Shutdown()
}