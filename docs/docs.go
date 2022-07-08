// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Alby",
            "url": "https://getalby.com",
            "email": "hello@getalby.com"
        },
        "license": {
            "name": "GNU GPL",
            "url": "https://www.gnu.org/licenses/gpl-3.0.en.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "post": {
                "description": "Exchanges a login + password for a token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Authenticate",
                "parameters": [
                    {
                        "description": "Login and password",
                        "name": "AuthRequestBody",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/balance": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Current user's balance in satoshi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Retrieve balance",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.BalanceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/invoices": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Returns a new bolt11 invoice",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "summary": "Generate a new invoice",
                "parameters": [
                    {
                        "description": "Add Invoice",
                        "name": "invoice",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v2controllers.AddInvoiceRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.AddInvoiceResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/invoices/incoming": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Returns a list of incoming invoices for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "summary": "Retrieve incoming invoices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/v2controllers.Invoice"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/invoices/outgoing": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Returns a list of outgoing payments for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "summary": "Retrieve outgoing payments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/v2controllers.Invoice"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/invoices/{payment_hash}": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Retrieve information about a specific invoice by payment hash",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "summary": "Get a specific invoice",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Payment hash",
                        "name": "payment_hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.Invoice"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/payments/bolt11": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Pay a bolt11 invoice",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment"
                ],
                "summary": "Pay an invoice",
                "parameters": [
                    {
                        "description": "Invoice to pay",
                        "name": "PayInvoiceRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v2controllers.PayInvoiceRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.PayInvoiceResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/payments/keysend": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Pay a node without an invoice using it's public key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment"
                ],
                "summary": "Make a keysend payment",
                "parameters": [
                    {
                        "description": "Invoice to pay",
                        "name": "KeySendRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v2controllers.KeySendRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.KeySendResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/users": {
            "post": {
                "description": "Create a new account with a login and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Create an account",
                "parameters": [
                    {
                        "description": "Create User",
                        "name": "account",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.CreateUserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.CreateUserResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.AuthRequestBody": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "controllers.AuthResponseBody": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "responses.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v2controllers.AddInvoiceRequestBody": {
            "type": "object",
            "required": [
                "amount"
            ],
            "properties": {
                "amount": {
                    "type": "integer",
                    "minimum": 0
                },
                "description": {
                    "type": "string"
                },
                "description_hash": {
                    "type": "string"
                }
            }
        },
        "v2controllers.AddInvoiceResponseBody": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "payment_hash": {
                    "type": "string"
                },
                "payment_request": {
                    "type": "string"
                }
            }
        },
        "v2controllers.BalanceResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "currency": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                }
            }
        },
        "v2controllers.CreateUserRequestBody": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "v2controllers.CreateUserResponseBody": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "v2controllers.Invoice": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "custom_records": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                },
                "description": {
                    "type": "string"
                },
                "description_hash": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "error_message": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "fee": {
                    "type": "integer"
                },
                "is_paid": {
                    "type": "boolean"
                },
                "keysend": {
                    "type": "boolean"
                },
                "payment_hash": {
                    "type": "string"
                },
                "payment_preimage": {
                    "type": "string"
                },
                "payment_request": {
                    "type": "string"
                },
                "settled_at": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "v2controllers.KeySendRequestBody": {
            "type": "object",
            "required": [
                "amount",
                "destination"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "customRecords": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "destination": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                }
            }
        },
        "v2controllers.KeySendResponseBody": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "description_hash": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "fee": {
                    "type": "integer"
                },
                "payment_hash": {
                    "type": "string"
                },
                "payment_preimage": {
                    "type": "string"
                }
            }
        },
        "v2controllers.PayInvoiceRequestBody": {
            "type": "object",
            "required": [
                "invoice"
            ],
            "properties": {
                "amount": {
                    "type": "integer",
                    "minimum": 0
                },
                "invoice": {
                    "type": "string"
                }
            }
        },
        "v2controllers.PayInvoiceResponseBody": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "description_hash": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "fee": {
                    "type": "integer"
                },
                "payment_hash": {
                    "type": "string"
                },
                "payment_preimage": {
                    "type": "string"
                },
                "payment_request": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "/auth"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.6.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "LNDhub.go",
	Description:      "Accounting wrapper for the Lightning Network providing separate accounts for end-users.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
