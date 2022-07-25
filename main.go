package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type patrimonio struct {
	Id         string `json:"id"`
	Tipo       string `json:"tipo"`
	Modelo     string `json:"modelo"`
	Observacao string `json:"observacao"`
}

type message struct {
	Success bool
	Message string
}

type resposta struct {
	Data    interface{}
	Message interface{}
}

func main() {

	//app := fiber.New()

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Use(cors.New())

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Rotas registradas",
			"Dados": []map[string]string{
				map[string]string{"route": "/api/patrimonios", "method": "GET", "description": "listar todos os patrimonios registrados na base de dados"},
				map[string]string{"route": "/api/patrimonio/add", "method": "POST", "description": "adicionar um novo patrimonio na base de dados"},
				map[string]string{"route": "/api/patrimonio/{id}/show", "method": "GET", "description": "exibir um patrimonio"},
				map[string]string{"route": "/api/patrimonio/update", "method": "PUT", "description": "atuaizar um patrimonio"},
				map[string]string{"route": "/api/patrimonio/{id}/delete", "method": "DELETE", "description": "excluir um patrimonio"},
			},
		})
	})

	app.Get("/api/patrimonios", func(c *fiber.Ctx) error {

		patrs := []patrimonio{}
		pat := patrimonio{}

		pat.Id = "010020010"
		pat.Tipo = "Computador"
		pat.Modelo = "DELL Inspire 755"
		pat.Observacao = "em perfeito estado"

		patrs = append(patrs, pat)

		pat.Id = "010020020"
		pat.Tipo = "Computador"
		pat.Modelo = "DELL Inspire 755"
		pat.Observacao = "em perfeito estado"

		patrs = append(patrs, pat)

		pat.Id = "01002030"
		pat.Tipo = "Monitor"
		pat.Modelo = "AOC 21\" "
		pat.Observacao = ""

		patrs = append(patrs, pat)

		pat.Id = "010020060"
		pat.Tipo = "Mouse"
		pat.Modelo = "Gistus"
		pat.Observacao = "Roda central"

		patrs = append(patrs, pat)

		return c.Status(c.Response().StatusCode()).JSON(patrs)
	})

	app.Post("/api/patrimonio/add", func(c *fiber.Ctx) error {

		pat := &patrimonio{}

		err := c.BodyParser(pat)
		if err != nil {
			dados := &message{
				Success: false,
				Message: "Erro ao tentar salvar registro.",
			}

			return c.Status(fiber.StatusOK).JSON(dados)
		}

		dados := &message{
			Success: true,
			Message: "registro salvo com sucesso.",
		}

		resp := &resposta{}
		resp.Data = pat
		resp.Message = dados

		//return c.JSON(resp)
		return c.Status(fiber.StatusOK).JSON(resp)

	})

	app.Put("/api/patrimonio/update", func(c *fiber.Ctx) error {

		pat := &patrimonio{}
		err := c.BodyParser(pat)

		if err != nil {

			msg := &message{
				Success: false,
				Message: "Erro ao tentar atualizar registro",
			}

			c.Status(fiber.StatusOK).JSON(msg)
		}

		msg := &message{
			Success: false,
			Message: "Registro atualizado com sucesso.",
		}

		resp := &resposta{}

		resp.Data = pat
		resp.Message = msg

		return c.Status(fiber.StatusOK).JSON(resp)
	})

	app.Delete("/api/patrimonio/delete", func(c *fiber.Ctx) error {

		pat := &patrimonio{}
		err := c.BodyParser(pat)

		if err != nil {
			msg := &message{
				Success: false,
				Message: fmt.Sprintf("Erro ao excluir registro %s.\n", err.Error()),
			}
			return c.Status(fiber.AcquireResponse().StatusCode()).JSON(msg)
		}

		msg := &message{
			Success: true,
			Message: "Registro excluido com sucesso",
		}

		resp := &resposta{}

		resp.Data = pat
		resp.Message = msg

		return c.Status(fiber.StatusOK).JSON(resp)
	})

	app.Get("/:param", func(c *fiber.Ctx) error {
		return c.SendString("Parametro: " + c.Params("param"))
	})

	app.Listen(":3000")

}