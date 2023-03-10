package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/syahmiabbas/OneCV_Internship_test/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MockJsonGet(c *gin.Context, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")

	// set path params
	// c.Params = params

	// set query params
	c.Request.URL.RawQuery = u.Encode()
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestRegisterStudentsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	body := struct {
		Teacher  string
		Students []string
	}{
		Teacher:  "teacherken@gmail.com",
		Students: []string{"studenthon@gmail.com", "studentjon@gmail.com"},
	}

	MockJsonPost(ctx, body)

	TeachersCreate(ctx)

	assert.EqualValues(t, 204, w.Code)
}

func TestCommonStudentsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	u := url.Values{}
	u.Add("teacher", "teacherken@gmail.com")

	MockJsonGet(ctx, u)

	CommonStudents(ctx)

	assert.EqualValues(t, 200, w.Code)
}

func TestRetrieveFornotificationHandler(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	type result struct {
		Recipients []string
	}

	body := struct {
		Teacher      string
		Notification string
	}{
		Teacher:      "teacherken@gmail.com",
		Notification: "Hello students! @studenthon@gmail.com @studentjon@gmail.com",
	}

	MockJsonPost(ctx, body)

	RetrieveFornotification(ctx)

	assert.EqualValues(t, 200, w.Code)

}

func TestSuspendStudentHandler(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	body := struct {
		Student string
	}{
		Student: "studenthon@gmail.com",
	}

	MockJsonPost(ctx, body)

	SuspendStudent(ctx)

	assert.EqualValues(t, 204, w.Code)
}
