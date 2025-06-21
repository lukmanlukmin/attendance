package attendance

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"attendance/bootstrap"
	"attendance/bootstrap/service"
	"attendance/constant"
	svc "attendance/service"
	mockAttendance "attendance/service/attendance/mocks"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSubmitAttendance(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	attendanceMock := mockAttendance.NewMockIAttendance(ctrl)

	tests := []struct {
		name         string
		mockSetup    func()
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success",
			mockSetup: func() {
				attendanceMock.EXPECT().
					SubmitAttendance(gomock.Any()).
					Return(nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: "",
		},
		{
			name: "Already submitted",
			mockSetup: func() {
				attendanceMock.EXPECT().
					SubmitAttendance(gomock.Any()).
					Return(constant.ErrAlreadySubmitAttendance)
			},
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: `{"error":"already submitted attendance"}`,
		},
		{
			name: "Internal server error",
			mockSetup: func() {
				attendanceMock.EXPECT().
					SubmitAttendance(gomock.Any()).
					Return(errors.New("internal error"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"internal error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			// Setup Fiber app
			app := fiber.New()
			h := NewHandler(&bootstrap.Bootstrap{
				Service: &service.Service{
					Service: svc.Service{
						Attendance: attendanceMock,
					},
				},
			})

			app.Post("/v1/attendances", h.Submit)

			// Send request
			req := httptest.NewRequest(http.MethodPost, "/v1/attendances", strings.NewReader(""))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			if tt.expectedBody != "" {
				buf := make([]byte, resp.ContentLength)
				resp.Body.Read(buf)
				assert.JSONEq(t, tt.expectedBody, string(buf))
			}
		})
	}
}
