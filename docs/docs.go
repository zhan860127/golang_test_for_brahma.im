// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "zhan860127.gmail.com"
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
        "/api/v1/del/{City}/{month}/{tempture}": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Weather"
                ],
                "summary": "delect the tempture in city on month",
                "parameters": [
                    {
                        "type": "string",
                        "description": "City",
                        "name": "City",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "month",
                        "name": "month",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功刪除",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "刪除失敗",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/modify/{City}/{month}/{tempture}": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Weather"
                ],
                "summary": "Modify the tempture in city on month",
                "parameters": [
                    {
                        "type": "string",
                        "description": "City",
                        "name": "City",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "month",
                        "name": "month",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "tempture",
                        "name": "tempture",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功修改",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "修改失敗",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/new/{City}/{month}/{tempture}": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Weather"
                ],
                "summary": "create the tempture in city on month",
                "parameters": [
                    {
                        "type": "string",
                        "description": "City",
                        "name": "City",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "month",
                        "name": "month",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "tempture",
                        "name": "tempture",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功新增",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "查詢失敗",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/weather/{City}/{month}": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Weather"
                ],
                "summary": "Get City_month Tempture",
                "parameters": [
                    {
                        "type": "string",
                        "description": "City",
                        "name": "City",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "month",
                        "name": "month",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查詢月份之城市溫度",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "404": {
                        "description": "查詢失敗",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7800",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Weather CRUD API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
