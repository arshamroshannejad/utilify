package utilify

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestStruct struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func TestReadJson(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		contentType    string
		expectedError  error
		expectedResult TestStruct
	}{
		{
			name:          "Valid JSON",
			body:          `{"name":"John","age":30,"email":"john@example.com"}`,
			contentType:   contentTypeJson,
			expectedError: nil,
			expectedResult: TestStruct{
				Name:  "John",
				Age:   30,
				Email: "john@example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tt.body))
			req.Header.Set(contentType, tt.contentType)
			var result TestStruct
			err := ReadJson(req, &result)
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err == nil && tt.expectedError == nil {
				if result != tt.expectedResult {
					t.Errorf("expected result %v, got %v", tt.expectedResult, result)
				}
			}
		})
	}
}

func TestWriteJson(t *testing.T) {
	recorder := httptest.NewRecorder()
	response := map[string]string{"message": "success"}
	if err := WriteJson(recorder, http.StatusOK, response); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}
	contentType := recorder.Header().Get(contentType)
	if contentType != contentTypeJson {
		t.Errorf("expected Content-Type %s, got %s", contentTypeJson, contentType)
	}
	expectedResponse := `{"message":"success"}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("expected body %s, got %s", expectedResponse, recorder.Body.String())
	}
}
