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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "get": {
                "description": "check user auth",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Auth",
                "responses": {
                    "200": {
                        "description": "success auth",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pkg.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/models.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "401": {
                        "description": "no cookie",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/feed": {
            "get": {
                "description": "Feed",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Feed",
                "responses": {
                    "200": {
                        "description": "success get feed",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pkg.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Post"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "405": {
                        "description": "invalid http method",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/friends/add": {
            "post": {
                "description": "add friend",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "AddFriend",
                "parameters": [
                    {
                        "description": "friends info",
                        "name": "friends",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Friends"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "friend added"
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "friend or user doesn't exist",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "409": {
                        "description": "friend already exists",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/friends/delete/{id_user}/{id_friend}": {
            "delete": {
                "description": "delete friend",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "DeleteFriend",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id_user",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Friend ID",
                        "name": "id_friend",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "friend deleted, body is empty"
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "friend/user/friendship doesn't exist",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/image/upload": {
            "post": {
                "description": "Upload image",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "image"
                ],
                "summary": "Upload image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "image file",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success upload image"
                    },
                    "405": {
                        "description": "invalid http method",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/image/{id}": {
            "get": {
                "description": "Get image by id",
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "image"
                ],
                "summary": "Get image by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "image ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success get image"
                    },
                    "405": {
                        "description": "invalid http method",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/logout": {
            "delete": {
                "description": "user logout",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout",
                "responses": {
                    "204": {
                        "description": "success logout, body is empty"
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "401": {
                        "description": "no cookie",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/post/create": {
            "post": {
                "description": "Create a post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Create a post",
                "parameters": [
                    {
                        "description": "post info",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success get post",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pkg.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/models.Post"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "405": {
                        "description": "invalid http method",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/post/update": {
            "post": {
                "description": "Update a post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Update a post",
                "parameters": [
                    {
                        "description": "post info",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success get post",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pkg.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/models.Post"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "405": {
                        "description": "invalid http method",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/post/{id}": {
            "get": {
                "description": "Get post by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Show a post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success get post",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pkg.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/models.Post"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "405": {
                        "description": "invalid http method",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a post",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Delete a post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "405": {
                        "description": "invalid http method",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "description": "user sign in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignIn",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserSignIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success sign in",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pkg.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/models.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "401": {
                        "description": "invalid password",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "user doesn't exist",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "user sign up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignUp",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "user created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pkg.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/models.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "409": {
                        "description": "user with this email already exists",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "get user's profile",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "GetProfile",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pkg.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/models.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "can't find user with such id",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/{id}/posts": {
            "get": {
                "description": "Get all users posts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get users posts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success get feed",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pkg.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Post"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Post not found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "405": {
                        "description": "invalid http method",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "models.Friends": {
            "type": "object",
            "required": [
                "id_friend",
                "id_user"
            ],
            "properties": {
                "id_friend": {
                    "type": "integer"
                },
                "id_user": {
                    "type": "integer"
                }
            }
        },
        "models.Image": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.Post": {
            "type": "object",
            "required": [
                "message",
                "user_id"
            ],
            "properties": {
                "avatar_id": {
                    "type": "integer"
                },
                "create_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Image"
                    }
                },
                "message": {
                    "type": "string"
                },
                "user_first_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_last_name": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "nick_name",
                "password"
            ],
            "properties": {
                "avatar": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "readOnly": true
                },
                "last_name": {
                    "type": "string"
                },
                "nick_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.UserSignIn": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "pkg.Response": {
            "type": "object",
            "properties": {
                "body": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "89.208.197.127:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "WS Swagger API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
