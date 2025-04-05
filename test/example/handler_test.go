package example_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"starter-go/api/rest/example"
	entity "starter-go/internal/domain/example"
	"starter-go/internal/pkg/driver/httpserver/middleware"
	pkgErrors "starter-go/internal/pkg/errors"
)

type mockExampleService struct {
	getExample    func(context.Context, int) (*entity.Example, error)
	createExample func(context.Context, string) (*entity.Example, error)
}

func (m *mockExampleService) GetExample(ctx context.Context, id int) (*entity.Example, error) {
	return m.getExample(ctx, id)
}

func (m *mockExampleService) CreateExample(ctx context.Context, desc string) (*entity.Example, error) {
	return m.createExample(ctx, desc)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.ErrorHandler())
	return r
}

func TestGetExample(t *testing.T) {
	tests := []struct {
		name           string
		exampleID      string
		mockResponse   *entity.Example
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:      "Success",
			exampleID: "123",
			mockResponse: &entity.Example{
				ID:          123,
				Description: "Test Example",
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"id":          float64(123), // JSON numbers are parsed as float64
				"description": "Test Example",
			},
		},
		{
			name:           "Invalid ID",
			exampleID:      "abc",
			mockResponse:   nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Not Found",
			exampleID:      "999",
			mockResponse:   nil,
			mockError:      pkgErrors.ErrNotFound("example", errors.New("not found")),
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setupRouter()
			mockSvc := &mockExampleService{
				getExample: func(ctx context.Context, id int) (*entity.Example, error) {
					return tt.mockResponse, tt.mockError
				},
			}

			handler := example.NewHandler(mockSvc)
			example.RegisterRoutes(r, handler)

			req, _ := http.NewRequest("GET", "/api/v1/example/"+tt.exampleID, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedBody != nil {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}

func TestCreateExampleEndpoint(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		mockResponse   *entity.Example
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Success",
			requestBody: map[string]interface{}{
				"description": "New Example",
			},
			mockResponse: &entity.Example{
				ID:          1,
				Description: "New Example",
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"id":          float64(1),
				"description": "New Example",
			},
		},
		{
			name: "Missing Description",
			requestBody: map[string]interface{}{
				"description": "",
			},
			mockResponse:   nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Service Error",
			requestBody: map[string]interface{}{
				"description": "Error Example",
			},
			mockResponse:   nil,
			mockError:      errors.New("service error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setupRouter()
			mockSvc := &mockExampleService{
				createExample: func(ctx context.Context, desc string) (*entity.Example, error) {
					return tt.mockResponse, tt.mockError
				},
			}

			handler := example.NewHandler(mockSvc)
			example.RegisterRoutes(r, handler)

			jsonBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/api/v1/example/", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedBody != nil {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
