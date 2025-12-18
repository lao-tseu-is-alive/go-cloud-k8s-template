// Package template_4_your_project_name provides Connect RPC handlers for the template4YourProjectNameService.
package template_4_your_project_name

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	template_4_your_project_namev1 "github.com/your-github-account/template-4-your-project-name/gen/template_4_your_project_name/v1"
	"github.com/your-github-account/template-4-your-project-name/gen/template_4_your_project_name/v1/template_4_your_project_namev1connect"
)

// template4YourProjectNameConnectServer implements the template4YourProjectNameServiceHandler interface.
// Authentication is handled by the AuthInterceptor, which injects user info into context.
type template4YourProjectNameConnectServer struct {
	BusinessService *BusinessService
	Log             *slog.Logger

	// Embed the unimplemented handler for forward compatibility
	template_4_your_project_namev1connect.Unimplementedtemplate4YourProjectNameServiceHandler
}

// Newtemplate4YourProjectNameConnectServer creates a new template4YourProjectNameConnectServer.
// Note: Authentication is handled by the AuthInterceptor, not by this server.
func Newtemplate4YourProjectNameConnectServer(business *BusinessService, log *slog.Logger) *template4YourProjectNameConnectServer {
	return &template4YourProjectNameConnectServer{
		BusinessService: business,
		Log:             log,
	}
}

// =============================================================================
// Helper Methods
// =============================================================================

// mapErrorToConnect converts business errors to Connect errors
func (s *template4YourProjectNameConnectServer) mapErrorToConnect(err error) *connect.Error {
	switch {
	case errors.Is(err, ErrNotFound):
		return connect.NewError(connect.CodeNotFound, err)
	case errors.Is(err, ErrTypetemplate4YourProjectNameNotFound):
		return connect.NewError(connect.CodeNotFound, err)
	case errors.Is(err, ErrAlreadyExists):
		return connect.NewError(connect.CodeAlreadyExists, err)
	case errors.Is(err, ErrUnauthorized):
		return connect.NewError(connect.CodePermissionDenied, err)
	case errors.Is(err, ErrNotOwner):
		return connect.NewError(connect.CodePermissionDenied, err)
	case errors.Is(err, ErrAdminRequired):
		return connect.NewError(connect.CodePermissionDenied, errors.New(OnlyAdminCanManageTypetemplate4YourProjectNames))
	case errors.Is(err, ErrInvalidInput):
		return connect.NewError(connect.CodeInvalidArgument, err)
	case errors.Is(err, pgx.ErrNoRows):
		return connect.NewError(connect.CodeNotFound, errors.New("not found"))
	default:
		s.Log.Error("internal error", "error", err)
		return connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}
}

// =============================================================================
// template4YourProjectNameService RPC Methods
// =============================================================================

// List returns a list of template_4_your_project_names
func (s *template4YourProjectNameConnectServer) List(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.ListRequest],
) (*connect.Response[template_4_your_project_namev1.ListResponse], error) {
	s.Log.Info("Connect: List called")

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("List", "userId", userId)

	// Build domain params from proto request
	msg := req.Msg
	params := ListParams{}
	if msg.Type != 0 {
		params.Type = &msg.Type
	}
	if msg.CreatedBy != 0 {
		params.CreatedBy = &msg.CreatedBy
	}
	if msg.Inactivated {
		params.Inactivated = &msg.Inactivated
	}
	if msg.Validated {
		params.Validated = &msg.Validated
	}

	// Handle pagination with defaults
	limit := s.BusinessService.ListDefaultLimit
	if msg.Limit > 0 {
		limit = int(msg.Limit)
	}
	offset := 0
	if msg.Offset > 0 {
		offset = int(msg.Offset)
	}

	// Call business logic
	list, err := s.BusinessService.List(ctx, offset, limit, params)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	// Convert to proto and return
	response := &template_4_your_project_namev1.ListResponse{
		template4YourProjectNames: Domaintemplate4YourProjectNameListSliceToProto(list),
	}
	return connect.NewResponse(response), nil
}

// Create creates a new template_4_your_project_name
func (s *template4YourProjectNameConnectServer) Create(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.CreateRequest],
) (*connect.Response[template_4_your_project_namev1.CreateResponse], error) {
	s.Log.Info("Connect: Create called")

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("Create", "userId", userId)

	// Convert proto to domain
	prototemplate4YourProjectName := req.Msg.template4YourProjectName
	if prototemplate4YourProjectName == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("template_4_your_project_name is required"))
	}

	domaintemplate4YourProjectName, err := Prototemplate4YourProjectNameToDomain(prototemplate4YourProjectName)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	// Call business logic
	createdtemplate4YourProjectName, err := s.BusinessService.Create(ctx, userId, *domaintemplate4YourProjectName)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	// Convert back to proto
	response := &template_4_your_project_namev1.CreateResponse{
		template4YourProjectName: Domaintemplate4YourProjectNameToProto(createdtemplate4YourProjectName),
	}
	return connect.NewResponse(response), nil
}

// Get retrieves a template_4_your_project_name by ID
func (s *template4YourProjectNameConnectServer) Get(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.GetRequest],
) (*connect.Response[template_4_your_project_namev1.GetResponse], error) {
	s.Log.Info("Connect: Get called", "id", req.Msg.Id)

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("Get", "userId", userId)

	// Parse UUID
	template_4_your_project_nameId, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid template_4_your_project_name ID format"))
	}

	// Call business logic
	template_4_your_project_name, err := s.BusinessService.Get(ctx, template_4_your_project_nameId)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	response := &template_4_your_project_namev1.GetResponse{
		template4YourProjectName: Domaintemplate4YourProjectNameToProto(template_4_your_project_name),
	}
	return connect.NewResponse(response), nil
}

// Update updates a template_4_your_project_name
func (s *template4YourProjectNameConnectServer) Update(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.UpdateRequest],
) (*connect.Response[template_4_your_project_namev1.UpdateResponse], error) {
	s.Log.Info("Connect: Update called", "id", req.Msg.Id)

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("Update", "userId", userId)

	// Parse UUID
	template_4_your_project_nameId, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid template_4_your_project_name ID format"))
	}

	// Convert proto to domain
	prototemplate4YourProjectName := req.Msg.template4YourProjectName
	if prototemplate4YourProjectName == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("template_4_your_project_name data is required"))
	}

	domaintemplate4YourProjectName, err := Prototemplate4YourProjectNameToDomain(prototemplate4YourProjectName)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	// Call business logic
	updatedtemplate4YourProjectName, err := s.BusinessService.Update(ctx, userId, template_4_your_project_nameId, *domaintemplate4YourProjectName)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	response := &template_4_your_project_namev1.UpdateResponse{
		template4YourProjectName: Domaintemplate4YourProjectNameToProto(updatedtemplate4YourProjectName),
	}
	return connect.NewResponse(response), nil
}

// Delete deletes a template_4_your_project_name
func (s *template4YourProjectNameConnectServer) Delete(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.DeleteRequest],
) (*connect.Response[template_4_your_project_namev1.DeleteResponse], error) {
	s.Log.Info("Connect: Delete called", "id", req.Msg.Id)

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("Delete", "userId", userId)

	// Parse UUID
	template_4_your_project_nameId, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid template_4_your_project_name ID format"))
	}

	// Call business logic
	err = s.BusinessService.Delete(ctx, userId, template_4_your_project_nameId)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	return connect.NewResponse(&template_4_your_project_namev1.DeleteResponse{}), nil
}

// Search returns template_4_your_project_names based on search criteria
func (s *template4YourProjectNameConnectServer) Search(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.SearchRequest],
) (*connect.Response[template_4_your_project_namev1.SearchResponse], error) {
	s.Log.Info("Connect: Search called")

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("Search", "userId", userId)

	msg := req.Msg
	params := SearchParams{}
	if msg.Keywords != "" {
		params.Keywords = &msg.Keywords
	}
	if msg.Type != 0 {
		params.Type = &msg.Type
	}
	if msg.CreatedBy != 0 {
		params.CreatedBy = &msg.CreatedBy
	}
	if msg.Inactivated {
		params.Inactivated = &msg.Inactivated
	}
	if msg.Validated {
		params.Validated = &msg.Validated
	}

	limit := s.BusinessService.ListDefaultLimit
	if msg.Limit > 0 {
		limit = int(msg.Limit)
	}
	offset := 0
	if msg.Offset > 0 {
		offset = int(msg.Offset)
	}

	list, err := s.BusinessService.Search(ctx, offset, limit, params)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	response := &template_4_your_project_namev1.SearchResponse{
		template4YourProjectNames: Domaintemplate4YourProjectNameListSliceToProto(list),
	}
	return connect.NewResponse(response), nil
}

// Count returns the number of template_4_your_project_names
func (s *template4YourProjectNameConnectServer) Count(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.CountRequest],
) (*connect.Response[template_4_your_project_namev1.CountResponse], error) {
	s.Log.Info("Connect: Count called")

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("Count", "userId", userId)

	msg := req.Msg
	params := CountParams{}
	if msg.Keywords != "" {
		params.Keywords = &msg.Keywords
	}
	if msg.Type != 0 {
		params.Type = &msg.Type
	}
	if msg.CreatedBy != 0 {
		params.CreatedBy = &msg.CreatedBy
	}
	if msg.Inactivated {
		params.Inactivated = &msg.Inactivated
	}
	if msg.Validated {
		params.Validated = &msg.Validated
	}

	count, err := s.BusinessService.Count(ctx, params)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	response := &template_4_your_project_namev1.CountResponse{
		Count: count,
	}
	return connect.NewResponse(response), nil
}

// GeoJson returns a GeoJSON representation of template_4_your_project_names
func (s *template4YourProjectNameConnectServer) GeoJson(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.GeoJsonRequest],
) (*connect.Response[template_4_your_project_namev1.GeoJsonResponse], error) {
	s.Log.Info("Connect: GeoJson called")

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("GeoJson", "userId", userId)

	msg := req.Msg
	params := GeoJsonParams{}
	if msg.Type != 0 {
		params.Type = &msg.Type
	}
	if msg.CreatedBy != 0 {
		params.CreatedBy = &msg.CreatedBy
	}
	if msg.Inactivated {
		params.Inactivated = &msg.Inactivated
	}
	if msg.Validated {
		params.Validated = &msg.Validated
	}

	limit := s.BusinessService.ListDefaultLimit
	if msg.Limit > 0 {
		limit = int(msg.Limit)
	}
	offset := 0
	if msg.Offset > 0 {
		offset = int(msg.Offset)
	}

	result, err := s.BusinessService.GeoJson(ctx, offset, limit, params)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	response := &template_4_your_project_namev1.GeoJsonResponse{
		Result: result,
	}
	return connect.NewResponse(response), nil
}

// ListByExternalId returns template_4_your_project_names filtered by external ID
func (s *template4YourProjectNameConnectServer) ListByExternalId(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.ListByExternalIdRequest],
) (*connect.Response[template_4_your_project_namev1.ListByExternalIdResponse], error) {
	s.Log.Info("Connect: ListByExternalId called", "externalId", req.Msg.ExternalId)

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("ListByExternalId", "userId", userId)

	msg := req.Msg
	limit := s.BusinessService.ListDefaultLimit
	if msg.Limit > 0 {
		limit = int(msg.Limit)
	}
	offset := 0
	if msg.Offset > 0 {
		offset = int(msg.Offset)
	}

	list, err := s.BusinessService.ListByExternalId(ctx, offset, limit, int(msg.ExternalId))
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	// Return NotFound if no results (matching HTTP handler behavior)
	if len(list) == 0 {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no template_4_your_project_names found with this external ID"))
	}

	response := &template_4_your_project_namev1.ListByExternalIdResponse{
		template4YourProjectNames: Domaintemplate4YourProjectNameListSliceToProto(list),
	}
	return connect.NewResponse(response), nil
}
