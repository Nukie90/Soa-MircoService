// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "nukie.nxk@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/": {
            "get": {
                "description": "Get all account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Get all account",
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Create account",
                "parameters": [
                    {
                        "description": "Account information",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateAccount"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/account/delete": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete account by verifying ID and PIN",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Delete account",
                "parameters": [
                    {
                        "description": "Delete account information",
                        "name": "deleteAccount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DeleteAccount"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/account/topup": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Top up account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Top up account",
                "parameters": [
                    {
                        "description": "Top up information",
                        "name": "topUp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TopUp"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/account/{id}": {
            "get": {
                "description": "Get account by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Get account by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.CreateAccount": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "pin": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.DeleteAccount": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "pin": {
                    "type": "string"
                }
            }
        },
        "model.TopUp": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3200",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Account Service",
	Description:      "This is the Account service API documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
