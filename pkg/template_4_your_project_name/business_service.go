package template_4_your_project_name

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/database"
)

// BusinessService Business Service contains the transport-agnostic business logic for template4YourProjectName operations
type BusinessService struct {
	Log              *slog.Logger
	DbConn           database.DB
	Store            Storage
	ListDefaultLimit int
}

// NewBusinessService creates a new instance of BusinessService
func NewBusinessService(store Storage, dbConn database.DB, log *slog.Logger, listDefaultLimit int) *BusinessService {
	return &BusinessService{
		Log:              log,
		DbConn:           dbConn,
		Store:            store,
		ListDefaultLimit: listDefaultLimit,
	}
}

// validateName validates the name field according to business rules
func validateName(name string) error {
	if len(strings.Trim(name, " ")) < 1 {
		return fmt.Errorf(FieldCannotBeEmpty, "name")
	}
	if len(name) < MinNameLength {
		return fmt.Errorf(FieldMinLengthIsN, "name", MinNameLength)
	}
	return nil
}

// GeoJson returns a geoJson representation of template_4_your_project_names based on the given parameters
func (s *BusinessService) GeoJson(ctx context.Context, offset, limit int, params GeoJsonParams) (string, error) {
	jsonResult, err := s.Store.GeoJson(ctx, offset, limit, params)
	if err != nil {
		return "", fmt.Errorf("error retrieving geoJson: %w", err)
	}
	if jsonResult == "" {
		return "empty", nil
	}
	return jsonResult, nil
}

// List returns the list of template_4_your_project_names based on the given parameters
func (s *BusinessService) List(ctx context.Context, offset, limit int, params ListParams) ([]*template4YourProjectNameList, error) {
	list, err := s.Store.List(ctx, offset, limit, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// No rows is not an error, return empty slice
			return make([]*template4YourProjectNameList, 0), nil
		}
		return nil, fmt.Errorf("error listing template_4_your_project_names: %w", err)
	}
	if list == nil {
		return make([]*template4YourProjectNameList, 0), nil
	}
	return list, nil
}

// Create creates a new template_4_your_project_name with the given data
func (s *BusinessService) Create(ctx context.Context, currentUserId int32, newtemplate4YourProjectName template4YourProjectName) (*template4YourProjectName, error) {
	// Validate name
	if err := validateName(newtemplate4YourProjectName.Name); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	// Validate TypeId
	typetemplate4YourProjectNameCount, err := s.DbConn.GetQueryInt(ctx, existTypetemplate4YourProjectName, newtemplate4YourProjectName.TypeId)
	if err != nil || typetemplate4YourProjectNameCount < 1 {
		return nil, fmt.Errorf("%w: typeId %v", ErrTypetemplate4YourProjectNameNotFound, newtemplate4YourProjectName.TypeId)
	}

	// Check if template_4_your_project_name already exists
	if s.Store.Exist(ctx, newtemplate4YourProjectName.Id) {
		return nil, fmt.Errorf("%w: id %v", ErrAlreadyExists, newtemplate4YourProjectName.Id)
	}

	// Set creator
	newtemplate4YourProjectName.CreatedBy = currentUserId

	// Create in storage
	template_4_your_project_nameCreated, err := s.Store.Create(ctx, newtemplate4YourProjectName)
	if err != nil {
		return nil, fmt.Errorf("error creating template_4_your_project_name: %w", err)
	}

	s.Log.Info("Created template_4_your_project_name", "id", template_4_your_project_nameCreated.Id, "userId", currentUserId)
	return template_4_your_project_nameCreated, nil
}

// Count returns the number of template_4_your_project_names based on the given parameters
func (s *BusinessService) Count(ctx context.Context, params CountParams) (int32, error) {
	numtemplate4YourProjectNames, err := s.Store.Count(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("error counting template_4_your_project_names: %w", err)
	}
	return numtemplate4YourProjectNames, nil
}

// Delete removes a template_4_your_project_name with the given ID
func (s *BusinessService) Delete(ctx context.Context, currentUserId int32, template_4_your_project_nameId uuid.UUID) error {
	// Check if template_4_your_project_name exists
	if !s.Store.Exist(ctx, template_4_your_project_nameId) {
		return fmt.Errorf("%w: id %v", ErrNotFound, template_4_your_project_nameId)
	}

	// Check if user is owner
	if !s.Store.IsUserOwner(ctx, template_4_your_project_nameId, currentUserId) {
		return fmt.Errorf("%w: user %d is not owner of template_4_your_project_name %v", ErrUnauthorized, currentUserId, template_4_your_project_nameId)
	}

	// Delete from storage
	err := s.Store.Delete(ctx, template_4_your_project_nameId, currentUserId)
	if err != nil {
		return fmt.Errorf("error deleting template_4_your_project_name: %w", err)
	}

	s.Log.Info("Deleted template_4_your_project_name", "id", template_4_your_project_nameId, "userId", currentUserId)
	return nil
}

// Get retrieves a template_4_your_project_name by its ID
func (s *BusinessService) Get(ctx context.Context, template_4_your_project_nameId uuid.UUID) (*template4YourProjectName, error) {
	// Check if template_4_your_project_name exists
	if !s.Store.Exist(ctx, template_4_your_project_nameId) {
		return nil, fmt.Errorf("%w: id %v", ErrNotFound, template_4_your_project_nameId)
	}

	// Get from storage
	template_4_your_project_name, err := s.Store.Get(ctx, template_4_your_project_nameId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving template_4_your_project_name: %w", err)
	}

	return template_4_your_project_name, nil
}

// Update updates a template_4_your_project_name with the given ID
func (s *BusinessService) Update(ctx context.Context, currentUserId int32, template_4_your_project_nameId uuid.UUID, updatetemplate4YourProjectName template4YourProjectName) (*template4YourProjectName, error) {
	// Check if template_4_your_project_name exists
	if !s.Store.Exist(ctx, template_4_your_project_nameId) {
		return nil, fmt.Errorf("%w: id %v", ErrNotFound, template_4_your_project_nameId)
	}

	// Check if user is owner
	if !s.Store.IsUserOwner(ctx, template_4_your_project_nameId, currentUserId) {
		return nil, fmt.Errorf("%w: user %d is not owner of template_4_your_project_name %v", ErrUnauthorized, currentUserId, template_4_your_project_nameId)
	}

	// Validate name
	if err := validateName(updatetemplate4YourProjectName.Name); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	// Validate TypeId
	typetemplate4YourProjectNameCount, err := s.DbConn.GetQueryInt(ctx, existTypetemplate4YourProjectName, updatetemplate4YourProjectName.TypeId)
	if err != nil || typetemplate4YourProjectNameCount < 1 {
		return nil, fmt.Errorf("%w: typeId %v", ErrTypetemplate4YourProjectNameNotFound, updatetemplate4YourProjectName.TypeId)
	}

	// Set last modifier
	updatetemplate4YourProjectName.LastModifiedBy = &currentUserId

	// Update in storage
	template_4_your_project_nameUpdated, err := s.Store.Update(ctx, template_4_your_project_nameId, updatetemplate4YourProjectName)
	if err != nil {
		return nil, fmt.Errorf("error updating template_4_your_project_name: %w", err)
	}

	s.Log.Info("Updated template_4_your_project_name", "id", template_4_your_project_nameId, "userId", currentUserId)
	return template_4_your_project_nameUpdated, nil
}

// ListByExternalId returns template_4_your_project_names filtered by external ID
func (s *BusinessService) ListByExternalId(ctx context.Context, offset, limit, externalId int) ([]*template4YourProjectNameList, error) {
	list, err := s.Store.ListByExternalId(ctx, offset, limit, externalId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// No rows is not an error, return empty slice
			return make([]*template4YourProjectNameList, 0), nil
		}
		return nil, fmt.Errorf("error listing template_4_your_project_names by external id: %w", err)
	}
	if list == nil {
		return make([]*template4YourProjectNameList, 0), nil
	}
	return list, nil
}

// Search returns template_4_your_project_names based on search criteria
func (s *BusinessService) Search(ctx context.Context, offset, limit int, params SearchParams) ([]*template4YourProjectNameList, error) {
	list, err := s.Store.Search(ctx, offset, limit, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// No rows is not an error, return empty slice
			return make([]*template4YourProjectNameList, 0), nil
		}
		return nil, fmt.Errorf("error searching template_4_your_project_names: %w", err)
	}
	if list == nil {
		return make([]*template4YourProjectNameList, 0), nil
	}
	return list, nil
}

// ListTypetemplate4YourProjectNames returns a list of Typetemplate4YourProjectName based on parameters
func (s *BusinessService) ListTypetemplate4YourProjectNames(ctx context.Context, offset, limit int, params Typetemplate4YourProjectNameListParams) ([]*Typetemplate4YourProjectNameList, error) {
	list, err := s.Store.ListTypetemplate4YourProjectName(ctx, offset, limit, params)
	if err != nil {
		return nil, fmt.Errorf("error listing type template_4_your_project_names: %w", err)
	}
	if list == nil {
		return make([]*Typetemplate4YourProjectNameList, 0), nil
	}
	return list, nil
}

// CreateTypetemplate4YourProjectName creates a new Typetemplate4YourProjectName
func (s *BusinessService) CreateTypetemplate4YourProjectName(ctx context.Context, currentUserId int32, isAdmin bool, newTypetemplate4YourProjectName Typetemplate4YourProjectName) (*Typetemplate4YourProjectName, error) {
	// Check admin privileges
	if !isAdmin {
		return nil, ErrAdminRequired
	}

	// Validate name
	if err := validateName(newTypetemplate4YourProjectName.Name); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	// Set creator
	newTypetemplate4YourProjectName.CreatedBy = currentUserId

	// Create in storage
	typetemplate4YourProjectNameCreated, err := s.Store.CreateTypetemplate4YourProjectName(ctx, newTypetemplate4YourProjectName)
	if err != nil {
		return nil, fmt.Errorf("error creating type template_4_your_project_name: %w", err)
	}

	s.Log.Info("Created Typetemplate4YourProjectName", "id", typetemplate4YourProjectNameCreated.Id, "userId", currentUserId)
	return typetemplate4YourProjectNameCreated, nil
}

// CountTypetemplate4YourProjectNames returns the count of Typetemplate4YourProjectNames based on parameters
func (s *BusinessService) CountTypetemplate4YourProjectNames(ctx context.Context, params Typetemplate4YourProjectNameCountParams) (int32, error) {
	numtemplate4YourProjectNames, err := s.Store.CountTypetemplate4YourProjectName(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("error counting type template_4_your_project_names: %w", err)
	}
	return numtemplate4YourProjectNames, nil
}

// DeleteTypetemplate4YourProjectName deletes a Typetemplate4YourProjectName by ID
func (s *BusinessService) DeleteTypetemplate4YourProjectName(ctx context.Context, currentUserId int32, isAdmin bool, typetemplate4YourProjectNameId int32) error {
	// Check admin privileges
	if !isAdmin {
		return ErrAdminRequired
	}

	// Check if Typetemplate4YourProjectName exists
	typetemplate4YourProjectNameCount, err := s.DbConn.GetQueryInt(ctx, existTypetemplate4YourProjectName, typetemplate4YourProjectNameId)
	if err != nil || typetemplate4YourProjectNameCount < 1 {
		return fmt.Errorf("%w: id %d", ErrTypetemplate4YourProjectNameNotFound, typetemplate4YourProjectNameId)
	}

	// Delete from storage
	err = s.Store.DeleteTypetemplate4YourProjectName(ctx, typetemplate4YourProjectNameId, currentUserId)
	if err != nil {
		return fmt.Errorf("error deleting type template_4_your_project_name: %w", err)
	}

	s.Log.Info("Deleted Typetemplate4YourProjectName", "id", typetemplate4YourProjectNameId, "userId", currentUserId)
	return nil
}

// GetTypetemplate4YourProjectName retrieves a Typetemplate4YourProjectName by ID
func (s *BusinessService) GetTypetemplate4YourProjectName(ctx context.Context, isAdmin bool, typetemplate4YourProjectNameId int32) (*Typetemplate4YourProjectName, error) {
	// Check admin privileges
	if !isAdmin {
		return nil, ErrAdminRequired
	}

	// Check if Typetemplate4YourProjectName exists
	typetemplate4YourProjectNameCount, err := s.DbConn.GetQueryInt(ctx, existTypetemplate4YourProjectName, typetemplate4YourProjectNameId)
	if err != nil || typetemplate4YourProjectNameCount < 1 {
		return nil, fmt.Errorf("%w: id %d", ErrTypetemplate4YourProjectNameNotFound, typetemplate4YourProjectNameId)
	}

	// Get from storage
	typetemplate4YourProjectName, err := s.Store.GetTypetemplate4YourProjectName(ctx, typetemplate4YourProjectNameId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving type template_4_your_project_name: %w", err)
	}

	return typetemplate4YourProjectName, nil
}

// UpdateTypetemplate4YourProjectName updates a Typetemplate4YourProjectName
func (s *BusinessService) UpdateTypetemplate4YourProjectName(ctx context.Context, currentUserId int32, isAdmin bool, typetemplate4YourProjectNameId int32, updateTypetemplate4YourProjectName Typetemplate4YourProjectName) (*Typetemplate4YourProjectName, error) {
	// Check admin privileges
	if !isAdmin {
		return nil, ErrAdminRequired
	}

	// Check if Typetemplate4YourProjectName exists
	typetemplate4YourProjectNameCount, err := s.DbConn.GetQueryInt(ctx, existTypetemplate4YourProjectName, typetemplate4YourProjectNameId)
	if err != nil || typetemplate4YourProjectNameCount < 1 {
		return nil, fmt.Errorf("%w: id %d", ErrTypetemplate4YourProjectNameNotFound, typetemplate4YourProjectNameId)
	}

	// Validate name
	if err := validateName(updateTypetemplate4YourProjectName.Name); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	// Set last modifier
	updateTypetemplate4YourProjectName.LastModifiedBy = &currentUserId

	// Update in storage
	template_4_your_project_nameUpdated, err := s.Store.UpdateTypetemplate4YourProjectName(ctx, typetemplate4YourProjectNameId, updateTypetemplate4YourProjectName)
	if err != nil {
		return nil, fmt.Errorf("error updating type template_4_your_project_name: %w", err)
	}

	s.Log.Info("Updated Typetemplate4YourProjectName", "id", typetemplate4YourProjectNameId, "userId", currentUserId)
	return template_4_your_project_nameUpdated, nil
}
