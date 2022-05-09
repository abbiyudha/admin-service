package routes

import (
	"CodingTest/delivery/handler/admin"
	"CodingTest/delivery/middleware"
	"github.com/labstack/echo/v4"
)

func AdminPath(e *echo.Echo, ah *admin.AdminHandler) {
	e.POST("create/admin", ah.CreateAdmin())
	e.POST("/login", ah.LoginHandler())
	e.GET("/admin", ah.GetAdminById(), middleware.JWTMiddleware())
	e.PUT("/update", ah.UpdateAdmin(), middleware.JWTMiddleware())
	e.DELETE("/delete", ah.DeleteAdmin(), middleware.JWTMiddleware())
}
