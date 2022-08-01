package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

var arquivoCSV string

func init() {

	arquivoCSV = "patrimonio.csv"

}

func main() {

	//app := fiber.New()

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Use(cors.New())

	app.Static("/api", "./public", fiber.Static{
		Compress: true,
	})

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

		patrs, err := ReadCSV(arquivoCSV)
		checkErr(err)

		return c.Status(c.Response().StatusCode()).Send(patrs)
	})

	app.Post("/api/patrimonio/add", func(c *fiber.Ctx) error {

		pat := patrimonio{}

		var dados = &message{}

		err := c.BodyParser(&pat)
		if err != nil {
			dados.Success = false
			dados.Message = "Erro ao tentar salvar registro."
		} else {

			dados.Success = true
			dados.Message = "registro salvo com sucesso."

		}

		txt_csv := pat.Id + ";" + pat.Tipo + ";" + pat.Modelo + ";" + pat.Observacao + "\n"

		fmt.Println(txt_csv)

		data, err := WriteCSV(arquivoCSV, txt_csv)
		checkErr(err)

		resp := &resposta{}
		resp.Data = data
		resp.Message = dados

		//return c.JSON(resp)
		return c.Status(fiber.StatusOK).JSON(resp)

	})

	app.Put("/api/patrimonio/update", func(c *fiber.Ctx) error {

		pat := &patrimonio{}
		err := c.BodyParser(&pat)
		checkErr(err)

		bytes, err := ReadCSV(arquivoCSV)
		checkErr(err)

		var jsonresult []patrimonio
		err = json.Unmarshal(bytes, &jsonresult)
		checkErr(err)

		//fmt.Println(jsonresult)
		var novoconteudo []patrimonio

		for index := range jsonresult {

			if jsonresult[index].Id != pat.Id {
				novoconteudo = append(novoconteudo, jsonresult[index])
			}
		}

		os.Remove(arquivoCSV)

		WriteCSV(arquivoCSV, pat.Id+";"+pat.Tipo+";"+pat.Modelo+";"+pat.Observacao+"\n")
		
		for i := range novoconteudo {
			WriteCSV(arquivoCSV, novoconteudo[i].Id+";"+novoconteudo[i].Tipo+";"+novoconteudo[i].Modelo+";"+novoconteudo[i].Observacao+"\n")
		}

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

	app.Delete("/api/patrimonio/delete", func(c *fiber.Ctx) error {

		pat := &patrimonio{}
		err := c.BodyParser(&pat)
		checkErr(err)

		bytes, err := ReadCSV(arquivoCSV)
		checkErr(err)

		var jsonresult []patrimonio
		err = json.Unmarshal(bytes, &jsonresult)
		checkErr(err)

		//fmt.Println(jsonresult)
		var novoconteudo []patrimonio

		for index := range jsonresult {
			if pat.Id != jsonresult[index].Id {
				novoconteudo = append(novoconteudo, jsonresult[index])
			}
		}

		os.Remove(arquivoCSV)

		for i := range novoconteudo {
			WriteCSV(arquivoCSV, novoconteudo[i].Id+";"+novoconteudo[i].Tipo+";"+novoconteudo[i].Modelo+";"+novoconteudo[i].Observacao+"\n")
		}

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

	app.Get("/api/patrimonio/:patrimonio/exists", func(c *fiber.Ctx) error {

		var existe bool = false

		pat := c.Params("patrimonio")

		bytes, err := ReadCSV(arquivoCSV)
		checkErr(err)

		var jsonresult []patrimonio

		err = json.Unmarshal(bytes, &jsonresult)
		checkErr(err)

		for index := range jsonresult {
			if pat == jsonresult[index].Id {
				existe = true
			}
		}

		var msg = message{}

		if !existe {
			msg.Success = false
			msg.Message = fmt.Sprintf("Registro %s não existe", pat)

		} else {

			msg.Success = true
			msg.Message = "Registro encontrada"

		}

		resp := &resposta{}

		resp.Data = existe
		resp.Message = msg

		return c.Status(fiber.StatusOK).JSON(resp)

	})

	// Página para testar aplicação em SPA (simple page aplication)
	app.Get("/api/app", func(c *fiber.Ctx) error {

		return c.Render("app", fiber.Map{

			"Title": "Página Teste",
			"Dados": nil,
		})

	})

	app.Listen(":3000")

}
