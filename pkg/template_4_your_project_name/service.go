package template_4_your_project_name

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/database"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/goHttpEcho"
)

type Permission int8 // enum
const (
	R Permission = iota // Read implies List (SELECT in DB, or GET in API)
	W                   // implies INSERT,UPDATE, DELETE
	M                   // Update or Put only
	D                   // Delete only
	C                   // Create only (Insert, Post)
	P                   // change Permissions of one template_4_your_project_name
	O                   // change Owner of one template4YourProjectName
	A                   // Audit log of changes of one template_4_your_project_name and read only special _fields like _created_by
)

func (s Permission) String() string {
	switch s {
	case R:
		return "R"
	case W:
		return "W"
	case M:
		return "M"
	case D:
		return "D"
	case C:
		return "C"
	case P:
		return "P"
	case O:
		return "O"
	case A:
		return "A"
	}
	return "ErrorPermissionUnknown"
}

type Service struct {
	Log              *slog.Logger
	DbConn           database.DB
	Store            Storage
	Server           *goHttpEcho.Server
	ListDefaultLimit int
}

func (s Service) GeoJson(ctx echo.Context, params GeoJsonParams) error {
	handlerName := "GeoJson"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	limit := s.ListDefaultLimit
	if params.Limit != nil {
		limit = int(*params.Limit)
	}
	offset := 0
	if params.Offset != nil {
		offset = int(*params.Offset)
	}
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	jsonResult, err := s.Store.GeoJson(reqCtx, offset, limit, params)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("there was a problem when calling store.List :%v", err))
		} else {
			jsonResult = "empty"
			return ctx.JSONBlob(http.StatusOK, []byte(jsonResult))
		}
	}
	return ctx.JSONBlob(http.StatusOK, []byte(jsonResult))
}

// List sends a list of template_4_your_project_names in the store based on the given parameters filters
// curl -s -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" 'http://localhost:9090/goapi/v1/template_4_your_project_name?limit=3&ofset=0' |jq
// curl -s -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" 'http://localhost:9090/goapi/v1/template_4_your_project_name?limit=3&type=112' |jq
func (s Service) List(ctx echo.Context, params ListParams) error {
	handlerName := "List"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	limit := s.ListDefaultLimit
	if params.Limit != nil {
		limit = int(*params.Limit)
	}
	offset := 0
	if params.Offset != nil {
		offset = int(*params.Offset)
	}
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	list, err := s.Store.List(reqCtx, offset, limit, params)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("there was a problem when calling store.List :%v", err))
		} else {
			list = make([]*template4YourProjectNameList, 0)
			return ctx.JSON(http.StatusOK, list)
		}
	}
	return ctx.JSON(http.StatusOK, list)
}

// Create allows to insert a new template_4_your_project_name
// curl -s -XPOST -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" -d '{"id": "3999971f-53d7-4eb6-8898-97f257ea5f27","type_id": 3,"name": "Gil-Parcelle","description": "just a nice parcelle test","external_id": 345678912,"inactivated": false,"managed_by": 999, "more_data": NULL,"pos_x":2537603.0 ,"pos_y":1152613.0   }' 'http://localhost:9090/goapi/v1/template_4_your_project_name'
func (s Service) Create(ctx echo.Context) error {
	handlerName := "Create"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	/* TODO implement ACL & RBAC handling
	if !s.Store.IsUserAllowedToCreate(currentUserId, typetemplate4YourProjectName) {
		return echo.NewHTTPError(http.StatusUnauthorized, "current user has no create role privilege")
	}
	*/
	newtemplate4YourProjectName := &template4YourProjectName{
		CreatedBy: int32(currentUserId),
	}
	if err := ctx.Bind(newtemplate4YourProjectName); err != nil {
		msg := fmt.Sprintf("Create has invalid format [%v]", err)
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	s.Log.Info("Create template4YourProjectName Bind ok", "template_4_your_project_name", newtemplate4YourProjectName.Name)
	if len(strings.Trim(newtemplate4YourProjectName.Name, " ")) < 1 {

		msg := fmt.Sprintf(FieldCannotBeEmpty, "name")
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	if len(newtemplate4YourProjectName.Name) < MinNameLength {
		msg := fmt.Sprintf(FieldMinLengthIsN, "name", MinNameLength)
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	if s.Store.Exist(reqCtx, newtemplate4YourProjectName.Id) {
		msg := fmt.Sprintf("This id (%v) already exist !", newtemplate4YourProjectName.Id)
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	template_4_your_project_nameCreated, err := s.Store.Create(reqCtx, *newtemplate4YourProjectName)
	if err != nil {
		msg := fmt.Sprintf("Create had an error saving template_4_your_project_name:%#v, err:%#v", *newtemplate4YourProjectName, err)
		s.Log.Info(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	s.Log.Info("Create success", "template_4_your_project_nameId", template_4_your_project_nameCreated.Id)
	return ctx.JSON(http.StatusCreated, template_4_your_project_nameCreated)
}

// Count returns the number of template_4_your_project_names found after filtering data with any given CountParams
// curl -s -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" 'http://localhost:9090/goapi/v1/template_4_your_project_name/count' |jq
func (s Service) Count(ctx echo.Context, params CountParams) error {
	handlerName := "Count"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	numtemplate4YourProjectNames, err := s.Store.Count(reqCtx, params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("problem counting template_4_your_project_names :%v", err))
	}
	return ctx.JSON(http.StatusOK, numtemplate4YourProjectNames)
}

// Delete will remove the given template_4_your_project_nameId entry from the store, and if not present will return 400 Bad Request
// curl -v -XDELETE -H "Content-Type: application/json" -H "Authorization: Bearer $token" 'http://localhost:8888/api/users/3' ->  204 No Content if present and delete it
// curl -v -XDELETE -H "Content-Type: application/json"  -H "Authorization: Bearer $token" 'http://localhost:8888/users/93333' -> 400 Bad Request
func (s Service) Delete(ctx echo.Context, template_4_your_project_nameId uuid.UUID) error {
	handlerName := "GeoJson"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := int32(claims.User.UserId)
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	if s.Store.Exist(reqCtx, template_4_your_project_nameId) == false {
		msg := fmt.Sprintf("Delete(%v) cannot delete this id, it does not exist !", template_4_your_project_nameId)
		s.Log.Warn(msg)
		return ctx.JSON(http.StatusNotFound, msg)
	}
	// IF USER IS NOT OWNER OF RECORD RETURN 401 Unauthorized
	if !s.Store.IsUserOwner(reqCtx, template_4_your_project_nameId, currentUserId) {
		return echo.NewHTTPError(http.StatusUnauthorized, "current user is not owner of this template_4_your_project_name")
	}
	/* TODO implement ACL & RBAC handling
	if !s.Store.IsUserAllowedToDelete(currentUserId, typetemplate4YourProjectName) {
		return echo.NewHTTPError(http.StatusUnauthorized, "current user has no create role privilege")
	}
	*/
	err := s.Store.Delete(reqCtx, template_4_your_project_nameId, currentUserId)
	if err != nil {
		msg := fmt.Sprintf("Delete(%v) got an error: %#v ", template_4_your_project_nameId, err)
		s.Log.Error(msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}
	return ctx.NoContent(http.StatusNoContent)

}

// Get will retrieve the template4YourProjectName with the given id in the store and return it
// curl -s -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" 'http://localhost:9090/goapi/v1/template_4_your_project_name/9999971f-53d7-4eb6-8898-97f257ea5f27' |jq
func (s Service) Get(ctx echo.Context, template_4_your_project_nameId uuid.UUID) error {
	handlerName := "Get"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	if s.Store.Exist(reqCtx, template_4_your_project_nameId) == false {
		msg := fmt.Sprintf("Get(%v) cannot get this id, it does not exist !", template_4_your_project_nameId)
		s.Log.Info(msg)
		return ctx.JSON(http.StatusNotFound, msg)
	}
	/* TODO implement ACL & RBAC handling
	if !s.Store.IsUserAllowedToGet(currentUserId, typetemplate4YourProjectName) {
		return echo.NewHTTPError(http.StatusUnauthorized, "current user has no create role privilege")
	}
	*/
	template_4_your_project_name, err := s.Store.Get(reqCtx, template_4_your_project_nameId)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("problem retrieving template_4_your_project_name :%v", err))
		} else {
			msg := fmt.Sprintf("Get(%v) no rows found in db", template_4_your_project_nameId)
			s.Log.Info(msg)
			return ctx.JSON(http.StatusNotFound, msg)
		}
	}
	return ctx.JSON(http.StatusOK, template_4_your_project_name)
}

// Update will change the attributes values for the template_4_your_project_name identified by the given template_4_your_project_nameId
// curl -s -XPUT -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" -d '{"id": "3999971f-53d7-4eb6-8898-97f257ea5f27","type_id": 3,"name": "Gil-Parcelle","description": "just a nice parcelle test by GIL","external_id": 345678912,"inactivated": false,"managed_by": 999, "more_data": {"info_value": 3230 },"pos_x":2537603.0 ,"pos_y":1152613.0   }' 'http://localhost:9090/goapi/v1/template_4_your_project_name/3999971f-53d7-4eb6-8898-97f257ea5f27' |jq
func (s Service) Update(ctx echo.Context, template_4_your_project_nameId uuid.UUID) error {
	handlerName := "GeoJson"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := int32(claims.User.UserId)
	s.Log.Info("handler called", "handler", handlerName, "template_4_your_project_nameId", template_4_your_project_nameId, "userId", currentUserId)
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	if s.Store.Exist(reqCtx, template_4_your_project_nameId) == false {
		msg := fmt.Sprintf("Update(%v) cannot update this id, it does not exist !", template_4_your_project_nameId)
		s.Log.Warn(msg)
		return ctx.JSON(http.StatusNotFound, msg)
	}
	if !s.Store.IsUserOwner(reqCtx, template_4_your_project_nameId, currentUserId) {
		return echo.NewHTTPError(http.StatusUnauthorized, "current user is not owner of this template_4_your_project_name")
	}
	/* TODO implement ACL & RBAC handling
	if !s.Store.IsUserAllowedToUpdate(currentUserId, typetemplate4YourProjectName) {
		return echo.NewHTTPError(http.StatusUnauthorized, "current user has no create role privilege")
	}
	*/

	updatetemplate4YourProjectName := new(template4YourProjectName)
	if err := ctx.Bind(updatetemplate4YourProjectName); err != nil {
		msg := fmt.Sprintf("Update has invalid format error:[%v]", err)
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	if len(strings.Trim(updatetemplate4YourProjectName.Name, " ")) < 1 {
		msg := fmt.Sprintf(FieldCannotBeEmpty, "name")
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	if len(updatetemplate4YourProjectName.Name) < MinNameLength {

		msg := fmt.Sprintf(FieldMinLengthIsN+FoundNum, "name", MinNameLength, len(updatetemplate4YourProjectName.Name))
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	updatetemplate4YourProjectName.LastModifiedBy = &currentUserId
	//TODO handle update of validated field correctly by adding validated time & user
	// handle update of managed_by field correctly by checking if user is a valid active one
	template_4_your_project_nameUpdated, err := s.Store.Update(reqCtx, template_4_your_project_nameId, *updatetemplate4YourProjectName)
	if err != nil {
		msg := fmt.Sprintf("Update had an error saving template_4_your_project_name:%#v, err:%#v", *updatetemplate4YourProjectName, err)
		s.Log.Info(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	s.Log.Info("Update success", "template_4_your_project_nameId", template_4_your_project_nameUpdated.Id)
	return ctx.JSON(http.StatusOK, template_4_your_project_nameUpdated)
}

// ListByExternalId sends a list of template_4_your_project_names in the store as json based of the given filters
// curl -s -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" 'http://localhost:9090/goapi/v1/template_4_your_project_name/by-external-id/345678912?limit=3&ofset=0' |jq
func (s Service) ListByExternalId(ctx echo.Context, externalId int32, params ListByExternalIdParams) error {
	handlerName := "ListByExternalId"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	limit := s.ListDefaultLimit
	if params.Limit != nil {
		limit = int(*params.Limit)
	}
	offset := 0
	if params.Offset != nil {
		offset = int(*params.Offset)
	}
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	list, err := s.Store.ListByExternalId(reqCtx, offset, limit, int(externalId))
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("there was a problem when calling store.ListByExternalId :%v", err))
		} else {
			list = make([]*template4YourProjectNameList, 0)
			return ctx.JSON(http.StatusNotFound, list)
		}
	}
	return ctx.JSON(http.StatusOK, list)
}

// Search returns a list of template_4_your_project_names in the store as json based of the given search criteria
// curl -s -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" 'http://localhost:9090/goapi/v1/template_4_your_project_name/search?limit=3&ofset=0' |jq
// curl -s -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" 'http://localhost:9090/goapi/v1/template_4_your_project_name/search?limit=3&type=112' |jq
func (s Service) Search(ctx echo.Context, params SearchParams) error {
	handlerName := "Search"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	limit := s.ListDefaultLimit
	if params.Limit != nil {
		limit = int(*params.Limit)
	}
	offset := 0
	if params.Offset != nil {
		offset = int(*params.Offset)
	}
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	list, err := s.Store.Search(reqCtx, offset, limit, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			list = make([]*template4YourProjectNameList, 0)
			return ctx.JSON(http.StatusOK, list)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("there was a problem when calling store.Search :%v", err))
		}
	}
	return ctx.JSON(http.StatusOK, list)
}

// Typetemplate4YourProjectNameList sends a list of Typetemplate4YourProjectName based on the given Typetemplate4YourProjectNameListParams parameters filters
func (s Service) Typetemplate4YourProjectNameList(ctx echo.Context, params Typetemplate4YourProjectNameListParams) error {
	handlerName := "Typetemplate4YourProjectNameList"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	limit := 250
	if params.Limit != nil {
		limit = int(*params.Limit)
	}
	offset := 0
	if params.Offset != nil {
		offset = int(*params.Offset)
	}
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	list, err := s.Store.ListTypetemplate4YourProjectName(reqCtx, offset, limit, params)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("there was a problem when calling store.ListTypetemplate4YourProjectName :%v", err))
		} else {
			list = make([]*Typetemplate4YourProjectNameList, 0)
			return ctx.JSON(http.StatusNotFound, list)
		}
	}
	return ctx.JSON(http.StatusOK, list)
}

// Typetemplate4YourProjectNameCreate will insert a new Typetemplate4YourProjectName in the store
func (s Service) Typetemplate4YourProjectNameCreate(ctx echo.Context) error {
	handlerName := "Typetemplate4YourProjectNameCreate"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	if !claims.User.IsAdmin {
		return echo.NewHTTPError(http.StatusUnauthorized, OnlyAdminCanManageTypetemplate4YourProjectNames)
	}
	newTypetemplate4YourProjectName := &Typetemplate4YourProjectName{
		Comment:           nil,
		CreatedAt:         nil,
		CreatedBy:         int32(currentUserId),
		Deleted:           false,
		DeletedAt:         nil,
		DeletedBy:         nil,
		Description:       nil,
		ExternalId:        nil,
		GeometryType:      nil,
		Id:                0,
		Inactivated:       false,
		InactivatedBy:     nil,
		InactivatedReason: nil,
		InactivatedTime:   nil,
		LastModifiedAt:    nil,
		LastModifiedBy:    nil,
		ManagedBy:         nil,
		IconPath:          "",
		MoreDataSchema:    nil,
		Name:              "",
		TableName:         nil,
	}
	if err := ctx.Bind(newTypetemplate4YourProjectName); err != nil {
		msg := fmt.Sprintf("Typetemplate4YourProjectNameCreate has invalid format [%v]", err)
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	if len(strings.Trim(newTypetemplate4YourProjectName.Name, " ")) < 1 {
		msg := fmt.Sprintf(FieldCannotBeEmpty, "name")
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	if len(newTypetemplate4YourProjectName.Name) < MinNameLength {
		msg := fmt.Sprintf(FieldMinLengthIsN+", found %d", "name", MinNameLength, len(newTypetemplate4YourProjectName.Name))
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	//s.Log.Info("# Create() before Store.Typetemplate4YourProjectNameCreate newtemplate4YourProjectName : %#v\n", newtemplate4YourProjectName)
	typetemplate4YourProjectNameCreated, err := s.Store.CreateTypetemplate4YourProjectName(reqCtx, *newTypetemplate4YourProjectName)
	if err != nil {
		msg := fmt.Sprintf("Typetemplate4YourProjectNameCreate had an error saving template_4_your_project_name:%#v, err:%#v", *newTypetemplate4YourProjectName, err)
		s.Log.Info(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	s.Log.Info("Typetemplate4YourProjectNameCreate success", "typetemplate4YourProjectNameId", typetemplate4YourProjectNameCreated.Id)
	return ctx.JSON(http.StatusCreated, typetemplate4YourProjectNameCreated)
}

func (s Service) Typetemplate4YourProjectNameCount(ctx echo.Context, params Typetemplate4YourProjectNameCountParams) error {
	handlerName := "Typetemplate4YourProjectNameCount"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	// Use request context for cancellation and tracing support
	reqCtx := ctx.Request().Context()
	numtemplate4YourProjectNames, err := s.Store.CountTypetemplate4YourProjectName(reqCtx, params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("problem counting template_4_your_project_names :%v", err))
	}
	return ctx.JSON(http.StatusOK, numtemplate4YourProjectNames)
}

// Typetemplate4YourProjectNameDelete will remove the given Typetemplate4YourProjectName entry from the store, and if not present will return 400 Bad Request
func (s Service) Typetemplate4YourProjectNameDelete(ctx echo.Context, typetemplate4YourProjectNameId int32) error {
	handlerName := "Typetemplate4YourProjectNameDelete"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := int32(claims.User.UserId)
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	// IF USER IS NOT ADMIN  RETURN 401 Unauthorized
	if !claims.User.IsAdmin {
		return echo.NewHTTPError(http.StatusUnauthorized, OnlyAdminCanManageTypetemplate4YourProjectNames)
	}
	reqCtx := ctx.Request().Context()
	typetemplate4YourProjectNameCount, err := s.DbConn.GetQueryInt(reqCtx, existTypetemplate4YourProjectName, typetemplate4YourProjectNameId)
	if err != nil || typetemplate4YourProjectNameCount < 1 {
		msg := fmt.Sprintf("Typetemplate4YourProjectNameDelete(%v) cannot delete this id, it does not exist !", typetemplate4YourProjectNameId)
		s.Log.Warn(msg)
		return ctx.JSON(http.StatusNotFound, msg)
	} else {
		err := s.Store.DeleteTypetemplate4YourProjectName(reqCtx, typetemplate4YourProjectNameId, currentUserId)
		if err != nil {
			msg := fmt.Sprintf("Typetemplate4YourProjectNameDelete(%v) got an error: %#v ", typetemplate4YourProjectNameId, err)
			s.Log.Error(msg)
			return echo.NewHTTPError(http.StatusInternalServerError, msg)
		}
		return ctx.NoContent(http.StatusNoContent)
	}
}

// Typetemplate4YourProjectNameGet will retrieve the template4YourProjectName with the given id in the store and return it
func (s Service) Typetemplate4YourProjectNameGet(ctx echo.Context, typetemplate4YourProjectNameId int32) error {
	handlerName := "Typetemplate4YourProjectNameGet"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := claims.User.UserId
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	if !claims.User.IsAdmin {
		return echo.NewHTTPError(http.StatusUnauthorized, OnlyAdminCanManageTypetemplate4YourProjectNames)
	}
	reqCtx := ctx.Request().Context()
	typetemplate4YourProjectNameCount, err := s.DbConn.GetQueryInt(reqCtx, existTypetemplate4YourProjectName, typetemplate4YourProjectNameId)
	if err != nil || typetemplate4YourProjectNameCount < 1 {
		msg := fmt.Sprintf("Typetemplate4YourProjectNameGet(%v) cannot retrieve this id, it does not exist !", typetemplate4YourProjectNameId)
		s.Log.Warn(msg)
		return ctx.JSON(http.StatusNotFound, msg)
	}
	typetemplate4YourProjectName, err := s.Store.GetTypetemplate4YourProjectName(reqCtx, typetemplate4YourProjectNameId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("problem retrieving Typetemplate4YourProjectName :%v", err))
	}
	return ctx.JSON(http.StatusOK, typetemplate4YourProjectName)
}

func (s Service) Typetemplate4YourProjectNameUpdate(ctx echo.Context, typetemplate4YourProjectNameId int32) error {
	handlerName := "Typetemplate4YourProjectNameUpdate"
	goHttpEcho.TraceHttpRequest(handlerName, ctx.Request(), s.Log)
	// get the current user from JWT TOKEN
	claims := s.Server.JwtCheck.GetJwtCustomClaimsFromContext(ctx)
	currentUserId := int32(claims.User.UserId)
	s.Log.Info("handler called", "handler", handlerName, "userId", currentUserId)
	// IF USER IS NOT ADMIN  RETURN 401 Unauthorized
	if !claims.User.IsAdmin {
		return echo.NewHTTPError(http.StatusUnauthorized, OnlyAdminCanManageTypetemplate4YourProjectNames)
	}
	reqCtx := ctx.Request().Context()
	typetemplate4YourProjectNameCount, err := s.DbConn.GetQueryInt(reqCtx, existTypetemplate4YourProjectName, typetemplate4YourProjectNameId)
	if err != nil || typetemplate4YourProjectNameCount < 1 {
		msg := fmt.Sprintf("Typetemplate4YourProjectNameUpdate(%v) cannot update this id, it does not exist !", typetemplate4YourProjectNameId)
		s.Log.Warn(msg)
		return ctx.JSON(http.StatusNotFound, msg)
	}
	uTypetemplate4YourProjectName := new(Typetemplate4YourProjectName)
	if err := ctx.Bind(uTypetemplate4YourProjectName); err != nil {
		msg := fmt.Sprintf("Typetemplate4YourProjectNameUpdate has invalid format error:[%v]", err)
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	if len(strings.Trim(uTypetemplate4YourProjectName.Name, " ")) < 1 {
		msg := fmt.Sprintf(FieldCannotBeEmpty, "name")
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	if len(uTypetemplate4YourProjectName.Name) < MinNameLength {
		msg := fmt.Sprintf(FieldMinLengthIsN+", found %d", "name", MinNameLength, len(uTypetemplate4YourProjectName.Name))
		s.Log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	uTypetemplate4YourProjectName.LastModifiedBy = &currentUserId
	template_4_your_project_nameUpdated, err := s.Store.UpdateTypetemplate4YourProjectName(reqCtx, typetemplate4YourProjectNameId, *uTypetemplate4YourProjectName)
	if err != nil {
		msg := fmt.Sprintf("Typetemplate4YourProjectNameUpdate had an error saving typetemplate4YourProjectName:%#v, err:%#v", *uTypetemplate4YourProjectName, err)
		s.Log.Info(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	s.Log.Info("Typetemplate4YourProjectNameUpdate success", "typetemplate4YourProjectNameId", template_4_your_project_nameUpdated.Id)
	return ctx.JSON(http.StatusOK, template_4_your_project_nameUpdated)
}
