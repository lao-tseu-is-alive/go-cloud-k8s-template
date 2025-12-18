// Package template_4_your_project_name provides Connect RPC handlers for the Typetemplate4YourProjectNameService.
package template_4_your_project_name

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5"
	template_4_your_project_namev1 "github.com/your-github-account/template-4-your-project-name/gen/template_4_your_project_name/v1"
	"github.com/your-github-account/template-4-your-project-name/gen/template_4_your_project_name/v1/template_4_your_project_namev1connect"
)

// Typetemplate4YourProjectNameConnectServer implements the Typetemplate4YourProjectNameServiceHandler interface.
// Authentication is handled by the AuthInterceptor, which injects user info into context.
type Typetemplate4YourProjectNameConnectServer struct {
	BusinessService *BusinessService
	Log             *slog.Logger

	// Embed the unimplemented handler for forward compatibility
	template_4_your_project_namev1connect.UnimplementedTypetemplate4YourProjectNameServiceHandler
}

// NewTypetemplate4YourProjectNameConnectServer creates a new Typetemplate4YourProjectNameConnectServer.
// Note: Authentication is handled by the AuthInterceptor, not by this server.
func NewTypetemplate4YourProjectNameConnectServer(business *BusinessService, log *slog.Logger) *Typetemplate4YourProjectNameConnectServer {
	return &Typetemplate4YourProjectNameConnectServer{
		BusinessService: business,
		Log:             log,
	}
}

// =============================================================================
// Helper Methods
// =============================================================================

// mapErrorToConnect converts business errors to Connect errors
func (s *Typetemplate4YourProjectNameConnectServer) mapErrorToConnect(err error) *connect.Error {
	switch {
	case errors.Is(err, ErrTypetemplate4YourProjectNameNotFound):
		return connect.NewError(connect.CodeNotFound, err)
	case errors.Is(err, ErrAlreadyExists):
		return connect.NewError(connect.CodeAlreadyExists, err)
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
// Typetemplate4YourProjectNameService RPC Methods
// =============================================================================

// List returns a list of type template_4_your_project_names
func (s *Typetemplate4YourProjectNameConnectServer) List(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.Typetemplate4YourProjectNameListRequest],
) (*connect.Response[template_4_your_project_namev1.Typetemplate4YourProjectNameListResponse], error) {
	s.Log.Info("Connect: Typetemplate4YourProjectName.List called")

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("Typetemplate4YourProjectName.List", "userId", userId)

	msg := req.Msg
	params := Typetemplate4YourProjectNameListParams{}
	if msg.Keywords != "" {
		params.Keywords = &msg.Keywords
	}
	if msg.CreatedBy != 0 {
		params.CreatedBy = &msg.CreatedBy
	}
	if msg.ExternalId != 0 {
		params.ExternalId = &msg.ExternalId
	}
	if msg.Inactivated {
		params.Inactivated = &msg.Inactivated
	}

	// Handle pagination
	limit := 250 // Default for Typetemplate4YourProjectName as in HTTP handler
	if msg.Limit > 0 {
		limit = int(msg.Limit)
	}
	offset := 0
	if msg.Offset > 0 {
		offset = int(msg.Offset)
	}

	list, err := s.BusinessService.ListTypetemplate4YourProjectNames(ctx, offset, limit, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Return empty list instead of error
			return connect.NewResponse(&template_4_your_project_namev1.Typetemplate4YourProjectNameListResponse{
				Typetemplate4YourProjectNames: []*template_4_your_project_namev1.Typetemplate4YourProjectNameList{},
			}), nil
		}
		return nil, s.mapErrorToConnect(err)
	}

	response := &template_4_your_project_namev1.Typetemplate4YourProjectNameListResponse{
		Typetemplate4YourProjectNames: DomainTypetemplate4YourProjectNameListSliceToProto(list),
	}
	return connect.NewResponse(response), nil
}

// Create creates a new type template_4_your_project_name
func (s *Typetemplate4YourProjectNameConnectServer) Create(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.Typetemplate4YourProjectNameCreateRequest],
) (*connect.Response[template_4_your_project_namev1.Typetemplate4YourProjectNameCreateResponse], error) {
	s.Log.Info("Connect: Typetemplate4YourProjectName.Create called")

	// User info injected by AuthInterceptor
	userId, isAdmin := GetUserFromContext(ctx)
	s.Log.Info("Typetemplate4YourProjectName.Create", "userId", userId, "isAdmin", isAdmin)

	protoTypetemplate4YourProjectName := req.Msg.Typetemplate4YourProjectName
	if protoTypetemplate4YourProjectName == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("type_template_4_your_project_name is required"))
	}

	domainTypetemplate4YourProjectName := ProtoTypetemplate4YourProjectNameToDomain(protoTypetemplate4YourProjectName)

	createdTypetemplate4YourProjectName, err := s.BusinessService.CreateTypetemplate4YourProjectName(ctx, userId, isAdmin, *domainTypetemplate4YourProjectName)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	response := &template_4_your_project_namev1.Typetemplate4YourProjectNameCreateResponse{
		Typetemplate4YourProjectName: DomainTypetemplate4YourProjectNameToProto(createdTypetemplate4YourProjectName),
	}
	return connect.NewResponse(response), nil
}

// Get retrieves a type template_4_your_project_name by ID
func (s *Typetemplate4YourProjectNameConnectServer) Get(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.Typetemplate4YourProjectNameGetRequest],
) (*connect.Response[template_4_your_project_namev1.Typetemplate4YourProjectNameGetResponse], error) {
	s.Log.Info("Connect: Typetemplate4YourProjectName.Get called", "id", req.Msg.Id)

	// User info injected by AuthInterceptor
	_, isAdmin := GetUserFromContext(ctx)

	typetemplate4YourProjectName, err := s.BusinessService.GetTypetemplate4YourProjectName(ctx, isAdmin, req.Msg.Id)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	response := &template_4_your_project_namev1.Typetemplate4YourProjectNameGetResponse{
		Typetemplate4YourProjectName: DomainTypetemplate4YourProjectNameToProto(typetemplate4YourProjectName),
	}
	return connect.NewResponse(response), nil
}

// Update updates a type template_4_your_project_name
func (s *Typetemplate4YourProjectNameConnectServer) Update(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.Typetemplate4YourProjectNameUpdateRequest],
) (*connect.Response[template_4_your_project_namev1.Typetemplate4YourProjectNameUpdateResponse], error) {
	s.Log.Info("Connect: Typetemplate4YourProjectName.Update called", "id", req.Msg.Id)

	// User info injected by AuthInterceptor
	userId, isAdmin := GetUserFromContext(ctx)
	s.Log.Info("Typetemplate4YourProjectName.Update", "userId", userId, "isAdmin", isAdmin)

	protoTypetemplate4YourProjectName := req.Msg.Typetemplate4YourProjectName
	if protoTypetemplate4YourProjectName == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("type_template_4_your_project_name data is required"))
	}

	domainTypetemplate4YourProjectName := ProtoTypetemplate4YourProjectNameToDomain(protoTypetemplate4YourProjectName)

	updatedTypetemplate4YourProjectName, err := s.BusinessService.UpdateTypetemplate4YourProjectName(ctx, userId, isAdmin, req.Msg.Id, *domainTypetemplate4YourProjectName)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	response := &template_4_your_project_namev1.Typetemplate4YourProjectNameUpdateResponse{
		Typetemplate4YourProjectName: DomainTypetemplate4YourProjectNameToProto(updatedTypetemplate4YourProjectName),
	}
	return connect.NewResponse(response), nil
}

// Delete deletes a type template_4_your_project_name
func (s *Typetemplate4YourProjectNameConnectServer) Delete(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.Typetemplate4YourProjectNameDeleteRequest],
) (*connect.Response[template_4_your_project_namev1.Typetemplate4YourProjectNameDeleteResponse], error) {
	s.Log.Info("Connect: Typetemplate4YourProjectName.Delete called", "id", req.Msg.Id)

	// User info injected by AuthInterceptor
	userId, isAdmin := GetUserFromContext(ctx)
	s.Log.Info("Typetemplate4YourProjectName.Delete", "userId", userId, "isAdmin", isAdmin)

	err := s.BusinessService.DeleteTypetemplate4YourProjectName(ctx, userId, isAdmin, req.Msg.Id)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	return connect.NewResponse(&template_4_your_project_namev1.Typetemplate4YourProjectNameDeleteResponse{}), nil
}

// Count returns the number of type template_4_your_project_names
func (s *Typetemplate4YourProjectNameConnectServer) Count(
	ctx context.Context,
	req *connect.Request[template_4_your_project_namev1.Typetemplate4YourProjectNameCountRequest],
) (*connect.Response[template_4_your_project_namev1.Typetemplate4YourProjectNameCountResponse], error) {
	s.Log.Info("Connect: Typetemplate4YourProjectName.Count called")

	// User info injected by AuthInterceptor
	userId, _ := GetUserFromContext(ctx)
	s.Log.Info("Typetemplate4YourProjectName.Count", "userId", userId)

	msg := req.Msg
	params := Typetemplate4YourProjectNameCountParams{}
	if msg.Keywords != "" {
		params.Keywords = &msg.Keywords
	}
	if msg.CreatedBy != 0 {
		params.CreatedBy = &msg.CreatedBy
	}
	if msg.Inactivated {
		params.Inactivated = &msg.Inactivated
	}

	count, err := s.BusinessService.CountTypetemplate4YourProjectNames(ctx, params)
	if err != nil {
		return nil, s.mapErrorToConnect(err)
	}

	response := &template_4_your_project_namev1.Typetemplate4YourProjectNameCountResponse{
		Count: count,
	}
	return connect.NewResponse(response), nil
}
