// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "stas@kuznec.team"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/stuff/all": {
            "get": {
                "description": "health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stuff Api"
                ],
                "summary": "health check",
                "responses": {
                    "200": {
                        "description": "OK response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/routergin.StuffAnswer"
                            }
                        }
                    },
                    "500": {
                        "description": "BAD response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/routergin.StuffAnswer"
                            }
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public API"
                ],
                "summary": "health check",
                "responses": {
                    "200": {
                        "description": "OK response",
                        "schema": {
                            "$ref": "#/definitions/routergin.OKAnswer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "routergin.OKAnswer": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "routergin.StuffAnswer": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "meta": {
                    "type": "object",
                    "additionalProperties": true
                },
                "name": {
                    "type": "string",
                    "example": "Сергей"
                },
                "photo_id": {
                    "type": "integer",
                    "example": 0
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger API",
	Description:      "Swagger API for Golang FaceID.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
