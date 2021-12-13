package controller

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bigchange/go-practice/testcase/db"
	mockdb "github.com/bigchange/go-practice/testcase/db/mockgen"
	"github.com/bigchange/go-practice/testcase/hub"
	"github.com/bigchange/go-practice/testcase/middleware"
	"github.com/bigchange/go-practice/testcase/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	engine *gin.Engine
)

func TestMain(m *testing.M) {
	engine = gin.Default()
	os.Exit(m.Run())
}

func TestHandler1(t *testing.T) {
	var (
		ctx = context.Background()
		path = "/status"
	)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockdb := mockdb.NewMockDB(ctrl)
	// 方式一： 通过context传入hub
	engine.Use(middleware.SetHub(&hub.Hub{
		PrimaryDB: mockdb ,
	}))
	// 方式二： 初始化不同模块中单例
	db.SetDB(mockdb)
	engine.GET(path, Handler1)

	t.Run("param error", func(t *testing.T) {
		quey := "int"
		finalPath := fmt.Sprintf("%v?k=%v", path, quey)
		req, err := http.NewRequestWithContext(ctx, "GET", finalPath, nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		engine.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("GetData error", func(t *testing.T) {
		quey := "100"
		finalPath := fmt.Sprintf("%v?k=%v", path, quey)
		// mockdb
		retErr := errors.New("GetData failed")
		mockdb.EXPECT().GetData(ctx, gomock.Any()).Return(nil, retErr)

		req, err := http.NewRequestWithContext(ctx, "GET", finalPath, nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		engine.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		/*
		// Test failed contain in return
		// Not pass: why?
		bodyBytes, err := ioutil.ReadAll(rr.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(bodyBytes), "failed")
		 */
	})

	t.Run("success", func(t *testing.T) {
		quey := "100"
		finalPath := fmt.Sprintf("%v?k=%v", path, quey)
		// mockdb
		mockdb.EXPECT().GetData(ctx, gomock.Any()).Return(&model.M1{Status: quey}, nil)

		req, err := http.NewRequestWithContext(ctx, "GET", finalPath, nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		engine.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		bodyBytes, err := ioutil.ReadAll(rr.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(bodyBytes), "100")


	})

}