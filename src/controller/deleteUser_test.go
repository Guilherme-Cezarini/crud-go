package controller

import (
	//"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"

)

func (uc *userControllerInterface) TestDelete(t *testing.T) {
	r := SetUpRouter()
	r.POST("/delete/:userId", uc.Insert)


	req, _ := http.NewRequest("DELETE", "/delete", nil)

	w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}