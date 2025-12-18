package template_4_your_project_name

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/database"
)

// Storage is an interface to different implementation of persistence for template4YourProjectNames/Typetemplate4YourProjectName
type Storage interface {
	// GeoJson returns a geoJson of existing template_4_your_project_names with the given offset and limit.
	GeoJson(ctx context.Context, offset, limit int, params GeoJsonParams) (string, error)
	// List returns the list of existing template_4_your_project_names with the given offset and limit.
	List(ctx context.Context, offset, limit int, params ListParams) ([]*template4YourProjectNameList, error)
	// ListByExternalId returns the list of existing template_4_your_project_names having the given externalId with the given offset and limit.
	ListByExternalId(ctx context.Context, offset, limit int, externalId int) ([]*template4YourProjectNameList, error)
	// Search returns the list of existing template_4_your_project_names filtered by search params with the given offset and limit.
	Search(ctx context.Context, offset, limit int, params SearchParams) ([]*template4YourProjectNameList, error)
	// Get returns the template_4_your_project_name with the specified template_4_your_project_names ID.
	Get(ctx context.Context, id uuid.UUID) (*template4YourProjectName, error)
	// Exist returns true only if a template_4_your_project_names with the specified id exists in store.
	Exist(ctx context.Context, id uuid.UUID) bool
	// Count returns the total number of template_4_your_project_names.
	Count(ctx context.Context, params CountParams) (int32, error)
	// Create saves a new template_4_your_project_names in the storage.
	Create(ctx context.Context, template_4_your_project_name template4YourProjectName) (*template4YourProjectName, error)
	// Update updates the template_4_your_project_names with given ID in the storage.
	Update(ctx context.Context, id uuid.UUID, template_4_your_project_name template4YourProjectName) (*template4YourProjectName, error)
	// Delete removes the template_4_your_project_names with given ID from the storage.
	Delete(ctx context.Context, id uuid.UUID, userId int32) error
	// Istemplate4YourProjectNameActive returns true if the template_4_your_project_name with the specified id has the inactivated attribute set to false
	Istemplate4YourProjectNameActive(ctx context.Context, id uuid.UUID) bool
	// IsUserOwner returns true only if userId is the creator of the record (owner) of this template_4_your_project_name in store.
	IsUserOwner(ctx context.Context, id uuid.UUID, userId int32) bool
	// CreateTypetemplate4YourProjectName saves a new typetemplate4YourProjectName in the storage.
	CreateTypetemplate4YourProjectName(ctx context.Context, typetemplate4YourProjectName Typetemplate4YourProjectName) (*Typetemplate4YourProjectName, error)
	// UpdateTypetemplate4YourProjectName updates the typetemplate4YourProjectName with given ID in the storage.
	UpdateTypetemplate4YourProjectName(ctx context.Context, id int32, typetemplate4YourProjectName Typetemplate4YourProjectName) (*Typetemplate4YourProjectName, error)
	// DeleteTypetemplate4YourProjectName removes the typetemplate4YourProjectName with given ID from the storage.
	DeleteTypetemplate4YourProjectName(ctx context.Context, id int32, userId int32) error
	// ListTypetemplate4YourProjectName returns the list of active typetemplate4YourProjectNames with the given offset and limit.
	ListTypetemplate4YourProjectName(ctx context.Context, offset, limit int, params Typetemplate4YourProjectNameListParams) ([]*Typetemplate4YourProjectNameList, error)
	// GetTypetemplate4YourProjectName returns the typetemplate4YourProjectName with the specified template_4_your_project_names ID.
	GetTypetemplate4YourProjectName(ctx context.Context, id int32) (*Typetemplate4YourProjectName, error)
	// CountTypetemplate4YourProjectName returns the number of Typetemplate4YourProjectName based on search criteria
	CountTypetemplate4YourProjectName(ctx context.Context, params Typetemplate4YourProjectNameCountParams) (int32, error)
}

func GetStorageInstanceOrPanic(ctx context.Context, dbDriver string, db database.DB, l *slog.Logger) Storage {
	var store Storage
	var err error
	switch dbDriver {
	case "pgx":
		store, err = NewPgxDB(ctx, db, l)
		if err != nil {
			l.Error("error doing NewPgxDB", "error", err)
			panic(err)
		}

	default:
		panic("unsupported DB driver type")
	}
	return store
}
