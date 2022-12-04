package route

import (
	"github.com/gofiber/fiber/v2"
	itementrypoint "so-cheap/internal/item/entrypoint"
	userentrypoint "so-cheap/internal/user/entrypoint"
)

func Routes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("users", userentrypoint.CreateUser)
	route.Put("users/:id", userentrypoint.UpdateUser)
	route.Get("users/:id", userentrypoint.GetUser)
	route.Get("users", userentrypoint.GetUsers)
	route.Delete("users/:id", userentrypoint.DeleteUser)

	route.Post("itens", itementrypoint.CreateItem)
	route.Put("itens/:id", itementrypoint.UpdateItem)
	route.Get("itens/:id", itementrypoint.GetItem)
	route.Get("itens", itementrypoint.GetItens)
	route.Delete("itens/:id", itementrypoint.DeleteItem)
}
