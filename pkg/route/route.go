package route

import (
	"github.com/gofiber/fiber/v2"
	itementrypoint "so-cheap/internal/item/entrypoint"
	orderentrypoint "so-cheap/internal/order/entrypoint"
	userentrypoint "so-cheap/internal/user/entrypoint"
	"so-cheap/pkg/mid"
)

func Routes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("users", userentrypoint.CreateUser)
	route.Post("users/authenticate", userentrypoint.AuthenticateUser)
	route.Put("users/:id", mid.JWTProtected(), userentrypoint.UpdateUser)
	route.Get("users/:id", mid.JWTProtected(), userentrypoint.GetUser)
	route.Get("users", mid.JWTProtected(), userentrypoint.GetUsers)
	route.Delete("users/:id", mid.JWTProtected(), userentrypoint.DeleteUser)

	route.Post("itens", mid.JWTProtected(), itementrypoint.CreateItem)
	route.Put("itens/:id", mid.JWTProtected(), itementrypoint.UpdateItem)
	route.Get("itens/:id", mid.JWTProtected(), itementrypoint.GetItem)
	route.Get("itens", mid.JWTProtected(), itementrypoint.GetItens)
	route.Delete("itens/:id", mid.JWTProtected(), itementrypoint.DeleteItem)

	route.Post("orders", mid.JWTProtected(), orderentrypoint.CreateOrder)
	route.Put("orders/:id", mid.JWTProtected(), orderentrypoint.UpdateOrder)
	route.Put("orders/:id/pay", mid.JWTProtected(), orderentrypoint.PayOrder)
	route.Get("orders/:id", mid.JWTProtected(), orderentrypoint.GetOrder)
	route.Get("orders", mid.JWTProtected(), orderentrypoint.GetOrders)
	route.Delete("orders/:id", mid.JWTProtected(), orderentrypoint.DeleteOrder)
}
