package template_4_your_project_name

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/golog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStorage is a mock implementation of the Storage interface for testing
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GeoJson(ctx context.Context, offset, limit int, params GeoJsonParams) (string, error) {
	args := m.Called(ctx, offset, limit, params)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) List(ctx context.Context, offset, limit int, params ListParams) ([]*template4YourProjectNameList, error) {
	args := m.Called(ctx, offset, limit, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*template4YourProjectNameList), args.Error(1)
}

func (m *MockStorage) ListByExternalId(ctx context.Context, offset, limit int, externalId int) ([]*template4YourProjectNameList, error) {
	args := m.Called(ctx, offset, limit, externalId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*template4YourProjectNameList), args.Error(1)
}

func (m *MockStorage) Search(ctx context.Context, offset, limit int, params SearchParams) ([]*template4YourProjectNameList, error) {
	args := m.Called(ctx, offset, limit, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*template4YourProjectNameList), args.Error(1)
}

func (m *MockStorage) Get(ctx context.Context, id uuid.UUID) (*template4YourProjectName, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*template4YourProjectName), args.Error(1)
}

func (m *MockStorage) Exist(ctx context.Context, id uuid.UUID) bool {
	args := m.Called(ctx, id)
	return args.Bool(0)
}

func (m *MockStorage) Count(ctx context.Context, params CountParams) (int32, error) {
	args := m.Called(ctx, params)
	return int32(args.Int(0)), args.Error(1)
}

func (m *MockStorage) Create(ctx context.Context, template_4_your_project_name template4YourProjectName) (*template4YourProjectName, error) {
	args := m.Called(ctx, template_4_your_project_name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*template4YourProjectName), args.Error(1)
}

func (m *MockStorage) Update(ctx context.Context, id uuid.UUID, template_4_your_project_name template4YourProjectName) (*template4YourProjectName, error) {
	args := m.Called(ctx, id, template_4_your_project_name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*template4YourProjectName), args.Error(1)
}

func (m *MockStorage) Delete(ctx context.Context, id uuid.UUID, userId int32) error {
	args := m.Called(ctx, id, userId)
	return args.Error(0)
}

func (m *MockStorage) Istemplate4YourProjectNameActive(ctx context.Context, id uuid.UUID) bool {
	args := m.Called(ctx, id)
	return args.Bool(0)
}

func (m *MockStorage) IsUserOwner(ctx context.Context, id uuid.UUID, userId int32) bool {
	args := m.Called(ctx, id, userId)
	return args.Bool(0)
}

func (m *MockStorage) CreateTypetemplate4YourProjectName(ctx context.Context, typetemplate4YourProjectName Typetemplate4YourProjectName) (*Typetemplate4YourProjectName, error) {
	args := m.Called(ctx, typetemplate4YourProjectName)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Typetemplate4YourProjectName), args.Error(1)
}

func (m *MockStorage) UpdateTypetemplate4YourProjectName(ctx context.Context, id int32, typetemplate4YourProjectName Typetemplate4YourProjectName) (*Typetemplate4YourProjectName, error) {
	args := m.Called(ctx, id, typetemplate4YourProjectName)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Typetemplate4YourProjectName), args.Error(1)
}

func (m *MockStorage) DeleteTypetemplate4YourProjectName(ctx context.Context, id int32, userId int32) error {
	args := m.Called(ctx, id, userId)
	return args.Error(0)
}

func (m *MockStorage) ListTypetemplate4YourProjectName(ctx context.Context, offset, limit int, params Typetemplate4YourProjectNameListParams) ([]*Typetemplate4YourProjectNameList, error) {
	args := m.Called(ctx, offset, limit, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Typetemplate4YourProjectNameList), args.Error(1)
}

func (m *MockStorage) GetTypetemplate4YourProjectName(ctx context.Context, id int32) (*Typetemplate4YourProjectName, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Typetemplate4YourProjectName), args.Error(1)
}

func (m *MockStorage) CountTypetemplate4YourProjectName(ctx context.Context, params Typetemplate4YourProjectNameCountParams) (int32, error) {
	args := m.Called(ctx, params)
	return int32(args.Int(0)), args.Error(1)
}

// MockDB is a minimal mock for database connection
type MockDB struct {
	mock.Mock
}

func (m *MockDB) GetQueryInt(ctx context.Context, query string, args ...interface{}) (int, error) {
	callArgs := m.Called(ctx, query, args)
	return callArgs.Int(0), callArgs.Error(1)
}

func (m *MockDB) GetVersion(ctx context.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockDB) Close() {
	m.Called()
}

func (m *MockDB) HealthCheck(ctx context.Context) (bool, error) {
	args := m.Called(ctx)
	return args.Bool(0), args.Error(1)
}

func (m *MockDB) GetQueryBool(ctx context.Context, query string, args ...interface{}) (bool, error) {
	callArgs := m.Called(ctx, query, args)
	return callArgs.Bool(0), callArgs.Error(1)
}

func (m *MockDB) ExecActionQuery(ctx context.Context, query string, args ...interface{}) (int, error) {
	callArgs := m.Called(ctx, query, args)
	return callArgs.Int(0), callArgs.Error(1)
}

func (m *MockDB) DoesTableExist(ctx context.Context, schema, table string) bool {
	args := m.Called(ctx, schema, table)
	return args.Bool(0)
}

func (m *MockDB) GetPGConn() (*pgxpool.Pool, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pgxpool.Pool), args.Error(1)
}

func (m *MockDB) GetQueryString(ctx context.Context, query string, args ...interface{}) (string, error) {
	callArgs := m.Called(ctx, query, args)
	return callArgs.String(0), callArgs.Error(1)
}

func (m *MockDB) Insert(ctx context.Context, query string, args ...interface{}) (int, error) {
	callArgs := m.Called(ctx, query, args)
	return callArgs.Int(0), callArgs.Error(1)
}

// Helper function to create a test business service
func createTestBusinessService(mockStore *MockStorage, mockDB *MockDB) *BusinessService {
	logger := golog.NewLogger("simple", os.Stdout, golog.InfoLevel, "test")
	return NewBusinessService(mockStore, mockDB, logger, 50)
}

// Test Create operation
func TestBusinessService_Create(t *testing.T) {
	ctx := context.Background()

	t.Run("successful creation", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		newtemplate4YourProjectName := template4YourProjectName{
			Id:   template_4_your_project_nameID,
			Name: "Test template4YourProjectName",
		}

		expectedtemplate4YourProjectName := newtemplate4YourProjectName
		expectedtemplate4YourProjectName.CreatedBy = 123

		// Mock Typetemplate4YourProjectName existence check
		mockDB.On("GetQueryInt", mock.Anytemplate_4_your_project_name, existTypetemplate4YourProjectName, []interface{}{newtemplate4YourProjectName.TypeId}).Return(1, nil)
		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(false)
		mockStore.On("Create", mock.Anytemplate_4_your_project_name, mock.Anytemplate_4_your_project_nameOfType("template4YourProjectName")).Return(&expectedtemplate4YourProjectName, nil)

		result, err := service.Create(ctx, 123, newtemplate4YourProjectName)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int32(123), result.CreatedBy)
		mockStore.AssertExpectations(t)
	})

	t.Run("validation error - empty name", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		newtemplate4YourProjectName := template4YourProjectName{
			Id:   uuid.New(),
			Name: "  ", // Empty/whitespace name
		}

		result, err := service.Create(ctx, 123, newtemplate4YourProjectName)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, ErrInvalidInput)
	})

	t.Run("validation error - name too short", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		newtemplate4YourProjectName := template4YourProjectName{
			Id:   uuid.New(),
			Name: "ab", // Less than MinNameLength (5)
		}

		result, err := service.Create(ctx, 123, newtemplate4YourProjectName)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, ErrInvalidInput)
	})

	t.Run("validation error - invalid type id", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		newtemplate4YourProjectName := template4YourProjectName{
			Id:     uuid.New(),
			Name:   "Test template4YourProjectName",
			TypeId: 999,
		}

		// Mock Typetemplate4YourProjectName existence check failure
		mockDB.On("GetQueryInt", mock.Anytemplate_4_your_project_name, existTypetemplate4YourProjectName, []interface{}{newtemplate4YourProjectName.TypeId}).Return(0, nil)

		result, err := service.Create(ctx, 123, newtemplate4YourProjectName)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, ErrTypetemplate4YourProjectNameNotFound)
	})

	t.Run("already exists error", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		newtemplate4YourProjectName := template4YourProjectName{
			Id:   template_4_your_project_nameID,
			Name: "Test template4YourProjectName",
		}

		// Mock Typetemplate4YourProjectName existence check
		mockDB.On("GetQueryInt", mock.Anytemplate_4_your_project_name, existTypetemplate4YourProjectName, []interface{}{newtemplate4YourProjectName.TypeId}).Return(1, nil)
		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(true)

		result, err := service.Create(ctx, 123, newtemplate4YourProjectName)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, ErrAlreadyExists)
		mockStore.AssertExpectations(t)
	})
}

// Test Get operation
func TestBusinessService_Get(t *testing.T) {
	ctx := context.Background()

	t.Run("successful get", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		expectedtemplate4YourProjectName := &template4YourProjectName{
			Id:   template_4_your_project_nameID,
			Name: "Test template4YourProjectName",
		}

		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(true)
		mockStore.On("Get", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(expectedtemplate4YourProjectName, nil)

		result, err := service.Get(ctx, template_4_your_project_nameID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Test template4YourProjectName", result.Name)
		mockStore.AssertExpectations(t)
	})

	t.Run("template_4_your_project_name not found", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(false)

		result, err := service.Get(ctx, template_4_your_project_nameID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, ErrNotFound)
		mockStore.AssertExpectations(t)
	})
}

// Test Update operation
func TestBusinessService_Update(t *testing.T) {
	ctx := context.Background()

	t.Run("successful update", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		userID := int32(123)
		updatetemplate4YourProjectName := template4YourProjectName{
			Id:   template_4_your_project_nameID,
			Name: "Updated template4YourProjectName",
		}

		expectedtemplate4YourProjectName := updatetemplate4YourProjectName
		expectedtemplate4YourProjectName.LastModifiedBy = &userID

		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(true)
		mockStore.On("IsUserOwner", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID, userID).Return(true)
		// Mock Typetemplate4YourProjectName existence check
		mockDB.On("GetQueryInt", mock.Anytemplate_4_your_project_name, existTypetemplate4YourProjectName, []interface{}{updatetemplate4YourProjectName.TypeId}).Return(1, nil)
		mockStore.On("Update", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID, mock.Anytemplate_4_your_project_nameOfType("template4YourProjectName")).Return(&expectedtemplate4YourProjectName, nil)

		result, err := service.Update(ctx, userID, template_4_your_project_nameID, updatetemplate4YourProjectName)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Updated template4YourProjectName", result.Name)
		mockStore.AssertExpectations(t)
	})

	t.Run("not owner error", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		userID := int32(123)
		updatetemplate4YourProjectName := template4YourProjectName{
			Id:   template_4_your_project_nameID,
			Name: "Updated template4YourProjectName",
		}

		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(true)
		mockStore.On("IsUserOwner", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID, userID).Return(false)

		result, err := service.Update(ctx, userID, template_4_your_project_nameID, updatetemplate4YourProjectName)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, ErrUnauthorized)
		mockStore.AssertExpectations(t)
	})
	t.Run("validation error - invalid type id", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		userID := int32(123)
		updatetemplate4YourProjectName := template4YourProjectName{
			Id:     template_4_your_project_nameID,
			Name:   "Updated template4YourProjectName",
			TypeId: 999,
		}

		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(true)
		mockStore.On("IsUserOwner", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID, userID).Return(true)
		// Mock Typetemplate4YourProjectName existence check failure
		mockDB.On("GetQueryInt", mock.Anytemplate_4_your_project_name, existTypetemplate4YourProjectName, []interface{}{updatetemplate4YourProjectName.TypeId}).Return(0, nil)

		result, err := service.Update(ctx, userID, template_4_your_project_nameID, updatetemplate4YourProjectName)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, ErrTypetemplate4YourProjectNameNotFound)
		mockStore.AssertExpectations(t)
	})
}

// Test Delete operation
func TestBusinessService_Delete(t *testing.T) {
	ctx := context.Background()

	t.Run("successful delete", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		userID := int32(123)

		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(true)
		mockStore.On("IsUserOwner", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID, userID).Return(true)
		mockStore.On("Delete", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID, userID).Return(nil)

		err := service.Delete(ctx, userID, template_4_your_project_nameID)

		assert.NoError(t, err)
		mockStore.AssertExpectations(t)
	})

	t.Run("not owner error", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		template_4_your_project_nameID := uuid.New()
		userID := int32(123)

		mockStore.On("Exist", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID).Return(true)
		mockStore.On("IsUserOwner", mock.Anytemplate_4_your_project_name, template_4_your_project_nameID, userID).Return(false)

		err := service.Delete(ctx, userID, template_4_your_project_nameID)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnauthorized)
		mockStore.AssertExpectations(t)
	})
}

// Test List operation
func TestBusinessService_List(t *testing.T) {
	ctx := context.Background()

	t.Run("successful list", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		expectedList := []*template4YourProjectNameList{
			{Id: uuid.New(), Name: "template4YourProjectName 1"},
			{Id: uuid.New(), Name: "template4YourProjectName 2"},
		}
		params := ListParams{}

		mockStore.On("List", mock.Anytemplate_4_your_project_name, 0, 10, params).Return(expectedList, nil)

		result, err := service.List(ctx, 0, 10, params)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		mockStore.AssertExpectations(t)
	})

	t.Run("empty list with pgx.ErrNoRows", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		params := ListParams{}
		mockStore.On("List", mock.Anytemplate_4_your_project_name, 0, 10, params).Return(nil, pgx.ErrNoRows)

		result, err := service.List(ctx, 0, 10, params)

		assert.NoError(t, err)
		assert.Empty(t, result)
		mockStore.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		params := ListParams{}
		dbError := errors.New("database connection failed")
		mockStore.On("List", mock.Anytemplate_4_your_project_name, 0, 10, params).Return(nil, dbError)

		result, err := service.List(ctx, 0, 10, params)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockStore.AssertExpectations(t)
	})
}

// Test Count operation
func TestBusinessService_Count(t *testing.T) {
	ctx := context.Background()

	t.Run("successful count", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		params := CountParams{}
		mockStore.On("Count", mock.Anytemplate_4_your_project_name, params).Return(42, nil)

		result, err := service.Count(ctx, params)

		assert.NoError(t, err)
		assert.Equal(t, int32(42), result)
		mockStore.AssertExpectations(t)
	})
}

// Test CreateTypetemplate4YourProjectName operation
func TestBusinessService_CreateTypetemplate4YourProjectName(t *testing.T) {
	ctx := context.Background()

	t.Run("successful creation by admin", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		newTypetemplate4YourProjectName := Typetemplate4YourProjectName{
			Name: "Test Type",
		}

		expectedTypetemplate4YourProjectName := newTypetemplate4YourProjectName
		expectedTypetemplate4YourProjectName.Id = 1
		expectedTypetemplate4YourProjectName.CreatedBy = 123

		mockStore.On("CreateTypetemplate4YourProjectName", mock.Anytemplate_4_your_project_name, mock.Anytemplate_4_your_project_nameOfType("Typetemplate4YourProjectName")).Return(&expectedTypetemplate4YourProjectName, nil)

		result, err := service.CreateTypetemplate4YourProjectName(ctx, 123, true, newTypetemplate4YourProjectName)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int32(123), result.CreatedBy)
		mockStore.AssertExpectations(t)
	})

	t.Run("non-admin rejection", func(t *testing.T) {
		mockStore := new(MockStorage)
		mockDB := new(MockDB)
		service := createTestBusinessService(mockStore, mockDB)

		newTypetemplate4YourProjectName := Typetemplate4YourProjectName{
			Name: "Test Type",
		}

		result, err := service.CreateTypetemplate4YourProjectName(ctx, 123, false, newTypetemplate4YourProjectName)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, ErrAdminRequired)
	})
}

// Test validation function
func TestValidateName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{"valid name", "Valid Name", false},
		{"empty string", "", true},
		{"only spaces", "   ", true},
		{"too short", "ab", true},
		{"exactly min length", "12345", false},
		{"longer than min", "Long Enough Name", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateName(tt.input)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
