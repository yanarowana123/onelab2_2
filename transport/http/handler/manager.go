package handler

import (
	"github.com/yanarowana123/onelab2_1/internal/services"
)

type Manager struct {
	service services.Manager
}

func NewManager(service services.Manager) *Manager {
	return &Manager{service}
}
