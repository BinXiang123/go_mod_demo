package main

import (
	"fmt"

	"github.com/kumamonMe/go_mod_demo/math"
	"github.com/kumamonMe/go_mod_demo/mydata"
	// _ "dubbo.apache.org/dubbo-go/v3/imports"
	// "github.com/labstack/echo"
)

func test_interface(u mydata.User) {
	u.SayHello()
}

func main() {
	nums := []float64{4.62, 90.31, 18.4, 70, 498}
	avg := math.Average(nums)
	fmt.Println(avg)

	//var s mydata.Student
	var s = mydata.Student{100, "Tom"}
	s.SetAge(10)
	// s.SetName("Tom")
	s.ShowAge()

	var t = mydata.Teacher{1, 2, "Lucy"}
	t.SayHello()
	t.ShowAge()

	test_interface(&s)
	test_interface(&t)

	var u mydata.User
	u = &t
	u.SayHello()
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}
