// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cliente": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cliente"
                ],
                "summary": "cadastra um novo cliente",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Cliente"
                        }
                    }
                }
            }
        },
        "/clientes/{cpf}": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cliente"
                ],
                "summary": "pega um cliente por cpf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cpf do cliente",
                        "name": "cpf",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Cliente"
                        }
                    }
                }
            }
        },
        "/liveness": {
            "get": {
                "description": "get the status of http.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Show the status of http.",
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pedido": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pedido"
                ],
                "summary": "cadastra um novo pedido",
                "responses": {}
            }
        },
        "/pedido/detail/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pedido"
                ],
                "summary": "lista detalhes do pedido",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id do pedido a ser lista",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Pedido"
                        }
                    }
                }
            }
        },
        "/pedido/{id}": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pedido"
                ],
                "summary": "atualiza o status do pedido",
                "responses": {}
            }
        },
        "/pedidos/{statuses}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pedido"
                ],
                "summary": "lista pedido por status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "status dos pedidos a ser pesquisado",
                        "name": "statuses",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Pedido"
                            }
                        }
                    }
                }
            }
        },
        "/produto": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produto"
                ],
                "summary": "atualiza um novo produto",
                "responses": {}
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produto"
                ],
                "summary": "cadastra um novo produto",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Produto"
                        }
                    }
                }
            }
        },
        "/produto/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produto"
                ],
                "summary": "apaga produto por id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id do produto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/produtos/{categoria}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produto"
                ],
                "summary": "pega produtos por categoria",
                "parameters": [
                    {
                        "type": "string",
                        "description": "categoria do produto",
                        "name": "categoria",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Produto"
                            }
                        }
                    }
                }
            }
        },
        "/readiness": {
            "get": {
                "description": "get the status of http.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Show the status of http.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Cliente": {
            "type": "object",
            "required": [
                "cpf",
                "nome"
            ],
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "data_aniversario": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "telefone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.Pedido": {
            "type": "object",
            "properties": {
                "cliente": {
                    "$ref": "#/definitions/domain.Cliente"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "observacao": {
                    "type": "string"
                },
                "produtos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Produto"
                    }
                },
                "status": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "domain.Produto": {
            "type": "object",
            "properties": {
                "categoria": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "descricao": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "3.0.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Tech Challenge API",
	Description:      "This is a documentation of all endpoints in the API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
