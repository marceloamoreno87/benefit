// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Marcelo Moreno",
            "email": "marceloamoreno87@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/benefit": {
            "get": {
                "description": "Return all benefits of the user by document (CPF) from http://extratoclube.com.br.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Benefits"
                ],
                "summary": "Get benefits by document (CPF)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Type the cpf of the user who wants to get the benefits",
                        "name": "doc",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Konsi API - Benefit",
	Description:      "Teste para vaga.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}