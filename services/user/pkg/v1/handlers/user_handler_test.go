package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	mock "gitlab.com/scalent/goxpert/services/user/mocks"
	handler "gitlab.com/scalent/goxpert/services/user/pkg/v1/handlers"
	//handler "gitlab.com/scalent/goxpert/user/pkg/v1/handlers"
)

func TestUserHandlersImpl_CreateUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mock.NewMockUsersService(ctrl)
	userHttpHandlers := handler.NewUserHandlerImpl(mockUsersService)

	type args struct {
		Method string
		URL    string
		Body   []byte
	}
	tests := []struct {
		name             string
		userHttpHandlers *handler.UserHandlersImpl
		args             args
		fu               func()
		wantStatus       int
	}{
		{
			"Test 1 : Status Sucess",
			userHttpHandlers,
			args{
				"GET",
				"users/1",
				[]byte(`{{
					"email": "testUser1@test.com",
					"name": "Test User 1000",
					"password": "pass",
					"designation": "Manager",
					"empID": "EMP1000",
					"userType": "USER",
				}}`),
			},
			func() {
				mockUsersService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return("1", nil).Times(1)
			},
			400,
		},
	}
	for _, tt := range tests {

		req, err := http.NewRequest(tt.args.Method, tt.args.URL, bytes.NewBuffer(tt.args.Body))
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tt.userHttpHandlers.CreateUser)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != tt.wantStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// // Check the response body is what we expect.
		// expected := `1`
		// if rr.Body.String() == expected {
		// 	t.Errorf("handler returned unexpected body: got %v want %v",
		// 		rr.Body.String(), expected)
		// }
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
}
