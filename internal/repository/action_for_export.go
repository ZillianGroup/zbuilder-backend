package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/illacloud/builder-backend/internal/idconvertor"
)

type ActionForExport struct {
	ID          string                 `json:"actionId"`
	UID         uuid.UUID              `json:"uid"`
	TeamID      string                 `json:"teamID"`
	App         string                 `json:"-"`
	Version     int                    `json:"-"`
	Resource    string                 `json:"resourceId,omitempty"`
	DisplayName string                 `json:"displayName"`
	Type        string                 `json:"actionType"`
	Template    map[string]interface{} `json:"content"`
	Transformer map[string]interface{} `json:"transformer"`
	TriggerMode string                 `json:"triggerMode"`
	Config      *ActionConfig          `json:"config"`
	CreatedAt   time.Time              `json:"createdAt,omitempty"`
	CreatedBy   string                 `json:"createdBy,omitempty"`
	UpdatedAt   time.Time              `json:"updatedAt,omitempty"`
	UpdatedBy   string                 `json:"updatedBy,omitempty"`
}

func NewActionForExport(action *Action) *ActionForExport {
	return &ActionForExport{
		ID:          idconvertor.ConvertIntToString(action.ID),
		UID:         action.UID,
		TeamID:      idconvertor.ConvertIntToString(action.TeamID),
		App:         idconvertor.ConvertIntToString(action.App),
		Version:     action.Version,
		Resource:    idconvertor.ConvertIntToString(action.Resource),
		DisplayName: action.ExportDisplayName(),
		Type:        action.ExportTypeInString(),
		Template:    action.Template,
		Transformer: action.Transformer,
		TriggerMode: action.TriggerMode,
		Config:      action.ExportConfig(),
		CreatedAt:   action.CreatedAt,
		CreatedBy:   idconvertor.ConvertIntToString(action.CreatedBy),
		UpdatedAt:   action.UpdatedAt,
		UpdatedBy:   idconvertor.ConvertIntToString(action.UpdatedBy),
	}
}
