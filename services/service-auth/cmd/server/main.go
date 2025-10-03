// service-auth → đăng nhập, đăng ký, phân quyền user/admin
//server/main.go để gọi hàm kết nối với các sql và kafka từ libs/db rồi điền tham số để kết nối với các sql và kafka
//server/main.go tạo fiber lắng nghe cổng 3001 cho service-auth
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DVT1609/tea-ordering-platform/services/service-auth/internal/handler"
)

func main() {
	app := fiber.New()

	app.Post("/login", handler.LoginHandler)
	app.Post("/register", handler.RegisterHandler)

	app.Listen(":3001")
}