package route

import (
	"github.com/gofiber/fiber/v2"
	"so-cheap/internal/item/entrypoint"
	userentrypoint "so-cheap/internal/user/entrypoint"
)

func Routes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("itens", entrypoint.CreateItem)
	route.Put("itens/:id", entrypoint.UpdateItem)
	route.Get("itens/:id", entrypoint.GetItem)
	route.Get("itens", entrypoint.GetItens)
	route.Delete("itens/:id", entrypoint.DeleteItem)

	route.Post("users", userentrypoint.CreateUser)
	route.Put("users/:id", userentrypoint.UpdateUser)
	route.Get("users/:id", userentrypoint.GetUser)
	route.Get("users", userentrypoint.GetUsers)
	route.Delete("users/:id", userentrypoint.DeleteUser)
}
