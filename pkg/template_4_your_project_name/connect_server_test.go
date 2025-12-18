package template_4_your_project_name

import (
	"context"
	"os"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/golog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	template_4_your_project_namev1 "github.com/your-github-account/template-4-your-project-name/gen/template_4_your_project_name/v1"
)

// =============================================================================
// Test Helpers
// =============================================================================

// Helper to create a test Connect server
func createTesttemplate4YourProjectNameConnectServer(mockStore *MockStorage, mockDB *MockDB) *template4YourProjectNameConnectServer {
	logger := golog.NewLogger("simple", os.Stdout, golog.InfoLevel, "test")
	businessService := NewBusinessService(mockStore, mockDB, logger, 50)
	return Newtemplate4YourProjectNameConnectServer(businessService, logger)
}

// Helper to create a test Typetemplate4YourProjectName Connect server
func createTestTypetemplate4YourProjectNameConnectServer(mockStore *MockStorage, mockDB *MockDB) *Typetemplate4YourProjectNameConnectServer {
	logger := golog.NewLogger("simple", os.Stdout, golog.InfoLevel, "test")
	businessService := NewBusinessService(mockStore, mockDB, logger, 50)
	return NewTypetemplate4YourProjectNameConnectServer(businessService, logger)
}

// Helper to create a context with user info (simulating what AuthInterceptor does)
func contextWithUser(userId int32, isAdmin bool) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDKey, userId)
	ctx = context.WithValue(ctx, isAdminKey, isAdmin)
	return ctx
}

// Helper to create a Connect request (no auth header needed since we inject via context)
func createConnectRequest[T any](msg *T) *connect.Request[T] {
	return connect.NewRequest(msg)
}

// =============================================================================
// template4YourProjectNameConnectServer Tests
// =============================================================================

func Testtemplate4YourProjectNameConnectServer_List(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTesttemplate4YourProjectNameConnectServer(mockStore, mockDB)

		// Setup mock storage
		now := time.Now()
		expectedList := []*template4YourProjectNameList{
			{Id: uuid.New(), Name: "template4YourProjectName 1", CreatedAt: &now},
			{Id: uuid.New(), Name: "template4YourProjectName 2", CreatedAt: &now},
		}
		mockStore.On("List", mock.Anytemplate_4_your_project_name, 0, 50, ListParams{}).Return(expectedList, nil)

		// Create request and context with user
		req := createConnectRequest(&template_4_your_project_namev1.ListRequest{Limit: 0, Offset: 0})
		ctx := contextWithUser(123, false)

		// Call handler
		resp, err := server.List(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Msg.template4YourProjectNames, 2)
		assert.Equal(t, "template4YourProjectName 1", resp.Msg.template4YourProjectNames[0].Name)
		assert.Equal(t, "template4YourProjectName 2", resp.Msg.template4YourProjectNames[1].Name)
		mockStore.AssertExpectations(t)
	})

	t.Run("list with pagination", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTesttemplate4YourProjectNameConnectServer(mockStore, mockDB)

		now := time.Now()
		expectedList := []*template4YourProjectNameList{
			{Id: uuid.New(), Name: "template4YourProjectName 3", CreatedAt: &now},
		}
		mockStore.On("List", mock.Anytemplate_4_your_project_name, 10, 5, ListParams{}).Return(expectedList, nil)

		req := createConnectRequest(&template_4_your_project_namev1.ListRequest{Limit: 5, Offset: 10})
		ctx := contextWithUser(123, false)

		resp, err := server.List(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Msg.template4YourProjectNames, 1)
		mockStore.AssertExpectations(t)
	})
}

func Testtemplate4YourProjectNameConnectServer_Get(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTesttemplate4YourProjectNameConnectServer(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		expectedtemplate4YourProjectName := &template4YourProjectName{
			Id:   template_4_your_project_nameID,
			Name: "Test template4YourProjectName",
		}

		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(true)
		mockStore.On("Get", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(expectedtemplate4YourProjectName, nil)

		req := createConnectRequest(&template_4_your_project_namev1.GetRequest{Id: template_4_your_project_nameID.String()})
		ctx := contextWithUser(123, false)

		resp, err := server.Get(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, template_4_your_project_nameID.String(), resp.Msg.template4YourProjectName.Id)
		assert.Equal(t, "Test template4YourProjectName", resp.Msg.template4YourProjectName.Name)
		mockStore.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTesttemplate4YourProjectNameConnectServer(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()

		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(false)

		req := createConnectRequest(&template_4_your_project_namev1.GetRequest{Id: template_4_your_project_nameID.String()})
		ctx := contextWithUser(123, false)

		resp, err := server.Get(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		connectErr, ok := err.(*connect.Error)
		assert.True(t, ok)
		assert.Equal(t, connect.CodeNotFound, connectErr.Code())
		mockStore.AssertExpectations(t)
	})

	t.Run("invalid UUID format", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTesttemplate4YourProjectNameConnectServer(mockStore, mockDB)

		req := createConnectRequest(&template_4_your_project_namev1.GetRequest{Id: "not-a-uuid"})
		ctx := contextWithUser(123, false)

		resp, err := server.Get(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		connectErr, ok := err.(*connect.Error)
		assert.True(t, ok)
		assert.Equal(t, connect.CodeInvalidArgument, connectErr.Code())
	})
}

func Testtemplate4YourProjectNameConnectServer_Create(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTesttemplate4YourProjectNameConnectServer(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		expectedtemplate4YourProjectName := &template4YourProjectName{
			Id:        template_4_your_project_nameID,
			Name:      "New template4YourProjectName",
			CreatedBy: 123,
		}

		mockDB.On("GetQueryInt", mock.Anytemplate_4_your_project_name, existTypetemplate4YourProjectName, mock.Anytemplate_4_your_project_name).Return(1, nil)
		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, mock.Anytemplate_4_your_project_nameOfType("uuid.UUID")).Return(false)
		mockStore.On("Create", mock.Anytemplate_4_your_project_name, mock.Anytemplate_4_your_project_nameOfType("template4YourProjectName")).Return(expectedtemplate4YourProjectName, nil)

		req := createConnectRequest(&template_4_your_project_namev1.CreateRequest{
			template4YourProjectName: &template_4_your_project_namev1.template4YourProjectName{
				Name: "New template4YourProjectName",
			},
		})
		ctx := contextWithUser(123, false)

		resp, err := server.Create(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "New template4YourProjectName", resp.Msg.template4YourProjectName.Name)
		mockStore.AssertExpectations(t)
	})

	t.Run("validation error - missing template_4_your_project_name", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTesttemplate4YourProjectNameConnectServer(mockStore, mockDB)

		req := createConnectRequest(&template_4_your_project_namev1.CreateRequest{template4YourProjectName: nil})
		ctx := contextWithUser(123, false)

		resp, err := server.Create(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		connectErr, ok := err.(*connect.Error)
		assert.True(t, ok)
		assert.Equal(t, connect.CodeInvalidArgument, connectErr.Code())
	})
}

func Testtemplate4YourProjectNameConnectServer_Delete(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTesttemplate4YourProjectNameConnectServer(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		userID := int32(123)

		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(true)
		mockStore.On("IsUserOwner", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID, userID).Return(true)
		mockStore.On("Delete", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID, userID).Return(nil)

		req := createConnectRequest(&template_4_your_project_namev1.DeleteRequest{Id: template_4_your_project_nameID.String()})
		ctx := contextWithUser(userID, false)

		resp, err := server.Delete(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		mockStore.AssertExpectations(t)
	})

	t.Run("permission denied - not owner", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTesttemplate4YourProjectNameConnectServer(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		userID := int32(123)

		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(true)
		mockStore.On("IsUserOwner", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID, userID).Return(false)

		req := createConnectRequest(&template_4_your_project_namev1.DeleteRequest{Id: template_4_your_project_nameID.String()})
		ctx := contextWithUser(userID, false)

		resp, err := server.Delete(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		connectErr, ok := err.(*connect.Error)
		assert.True(t, ok)
		assert.Equal(t, connect.CodePermissionDenied, connectErr.Code())
		mockStore.AssertExpectations(t)
	})
}

func Testtemplate4YourProjectNameConnectServer_Count(t *testing.T) {
	t.Run("successful count", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTesttemplate4YourProjectNameConnectServer(mockStore, mockDB)

		mockStore.On("Count", mock.Anytemplate_4_your_project_name, CountParams{}).Return(42, nil)

		req := createConnectRequest(&template_4_your_project_namev1.CountRequest{})
		ctx := contextWithUser(123, false)

		resp, err := server.Count(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, int32(42), resp.Msg.Count)
		mockStore.AssertExpectations(t)
	})
}

// =============================================================================
// Typetemplate4YourProjectNameConnectServer Tests
// =============================================================================

func TestTypetemplate4YourProjectNameConnectServer_List(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTestTypetemplate4YourProjectNameConnectServer(mockStore, mockDB)

		now := time.Now()
		expectedList := []*Typetemplate4YourProjectNameList{
			{Id: 1, Name: "Type 1", CreatedAt: now},
			{Id: 2, Name: "Type 2", CreatedAt: now},
		}

		mockStore.On("ListTypetemplate4YourProjectName", mock.Anytemplate_4_your_project_name, 0, 250, Typetemplate4YourProjectNameListParams{}).Return(expectedList, nil)

		req := createConnectRequest(&template_4_your_project_namev1.Typetemplate4YourProjectNameListRequest{})
		ctx := contextWithUser(123, false)

		resp, err := server.List(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Msg.Typetemplate4YourProjectNames, 2)
		mockStore.AssertExpectations(t)
	})
}

func TestTypetemplate4YourProjectNameConnectServer_Create(t *testing.T) {
	t.Run("admin can create", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTestTypetemplate4YourProjectNameConnectServer(mockStore, mockDB)

		expectedTypetemplate4YourProjectName := &Typetemplate4YourProjectName{
			Id:        1,
			Name:      "New Type",
			CreatedBy: 123,
		}

		mockStore.On("CreateTypetemplate4YourProjectName", mock.Anytemplate_4_your_project_name, mock.Anytemplate_4_your_project_nameOfType("Typetemplate4YourProjectName")).Return(expectedTypetemplate4YourProjectName, nil)

		req := createConnectRequest(&template_4_your_project_namev1.Typetemplate4YourProjectNameCreateRequest{
			Typetemplate4YourProjectName: &template_4_your_project_namev1.Typetemplate4YourProjectName{
				Name: "New Type",
			},
		})
		ctx := contextWithUser(123, true) // isAdmin = true

		resp, err := server.Create(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "New Type", resp.Msg.Typetemplate4YourProjectName.Name)
		mockStore.AssertExpectations(t)
	})

	t.Run("non-admin rejected", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		server := createTestTypetemplate4YourProjectNameConnectServer(mockStore, mockDB)

		req := createConnectRequest(&template_4_your_project_namev1.Typetemplate4YourProjectNameCreateRequest{
			Typetemplate4YourProjectName: &template_4_your_project_namev1.Typetemplate4YourProjectName{
				Name: "New Type",
			},
		})
		ctx := contextWithUser(123, false) // isAdmin = false

		resp, err := server.Create(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		connectErr, ok := err.(*connect.Error)
		assert.True(t, ok)
		assert.Equal(t, connect.CodePermissionDenied, connectErr.Code())
	})
}

// =============================================================================
// AuthInterceptor Tests
// =============================================================================

func TestGetUserFromContext(t *testing.T) {
	t.Run("user present in context", func(t *testing.T) {
		ctx := contextWithUser(456, true)

		userId, isAdmin := GetUserFromContext(ctx)

		assert.Equal(t, int32(456), userId)
		assert.True(t, isAdmin)
	})

	t.Run("user not present in context", func(t *testing.T) {
		ctx := context.Background()

		userId, isAdmin := GetUserFromContext(ctx)

		assert.Equal(t, int32(0), userId)
		assert.False(t, isAdmin)
	})
}
