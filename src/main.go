package main

import (
	"context"
	"fmt"
	"main/src/common/cpuTest"
	"main/src/common/db"
	"net/http"

	"net/http/pprof"
	_ "net/http/pprof"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func main() {
	//서버 초기 셋팅
	if err := db.InitMongoDB(); err != nil{
		fmt.Println(err.Error())
		return
	}

	// Echo 인스턴스 생성
	e := echo.New()

	// 미들웨어 등록 - 로깅 및 기본 리커버리
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// 미들웨어 등록 - 프로파일링
	e.GET("/debug/pprof/*", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	e.GET("/debug/pprof/cmdline", echo.WrapHandler(http.HandlerFunc(pprof.Cmdline)))
	e.GET("/debug/pprof/profile", echo.WrapHandler(http.HandlerFunc(pprof.Profile)))
	e.GET("/debug/pprof/symbol", echo.WrapHandler(http.HandlerFunc(pprof.Symbol)))
	e.GET("/debug/pprof/trace", echo.WrapHandler(http.HandlerFunc(pprof.Trace)))
	// 라우터 등록
	e.GET("/test/cpu/2", func(c echo.Context) error {
		cpuTest.FibRecursive(50)
		return c.String(http.StatusOK, "Hello, Echo!")
	})
	e.POST("/user",func(c echo.Context) error{
		userDTO := db.UserDTO{Name: "ryan",Age: 30}
		result , err := db.UserCollection.InsertOne(context.Background(),userDTO)
		if err != nil{
			return err
		}
		id := result.InsertedID.(primitive.ObjectID)
		fmt.Println(id.Hex())
		return c.JSON(http.StatusOK, "ok")
	})
	e.GET("/test/cpu/1",func(c echo.Context) error{
		cpuTest.IncreaseInt()
		return c.String(http.StatusOK,"cpu stress start")
	})
	// 서버 시작
	e.Start(":8080")

}
