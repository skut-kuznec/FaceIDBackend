package handler

import (
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
	"github.com/smart48ru/FaceIDApp/internal/service/imageservice"
	"github.com/smart48ru/FaceIDApp/internal/service/staffservice"
	"github.com/smart48ru/FaceIDApp/internal/service/timerecordservice"
)

var _ openapi.ServerInterface = &Handlers{}

type Handlers struct {
	imageService      *imageservice.Service
	staffService      *staffservice.Service
	timeRecordService *timerecordservice.Service
}

func New(
	imageService *imageservice.Service,
	staffService *staffservice.Service,
	timeRecordService *timerecordservice.Service,
) *Handlers {
	return &Handlers{
		imageService:      imageService,
		staffService:      staffService,
		timeRecordService: timeRecordService,
	}
}
