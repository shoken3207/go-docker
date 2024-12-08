// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
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
        "/api/auth/emailVerified/{email}": {
            "get": {
                "description": "リクエストからメールアドレス取得後、ユーザー登録されていないか確認し、メールアドレス宛に本登録URLをメールで送信",
                "tags": [
                    "auth"
                ],
                "summary": "メールアドレスの本人確認",
                "parameters": [
                    {
                        "type": "string",
                        "description": "メールアドレス",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "内部エラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/login": {
            "post": {
                "description": "メールアドレスとパスワードが合致したら、jwtトークンをCookieに保存",
                "tags": [
                    "auth"
                ],
                "summary": "ログイン",
                "parameters": [
                    {
                        "description": "ログイン情報",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse-auth_LoginResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "404": {
                        "description": "not foundエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "内部エラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "description": "メールアドレス確認後にリクエスト内容をユーザーテーブルに保存",
                "tags": [
                    "auth"
                ],
                "summary": "ユーザー登録",
                "parameters": [
                    {
                        "description": "ユーザー情報",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "内部エラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/resetPass": {
            "put": {
                "description": "メール内リンクで本人確認後、トークンと新しいパスワードをリクエストで取得し、",
                "tags": [
                    "auth"
                ],
                "summary": "ログアウト状態からパスワードを変更",
                "parameters": [
                    {
                        "description": "メールアドレス",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.EmailVerificationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "内部エラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/updatePass": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "現在のパスワードと新しいパスワードをリクエストで取得し、現在のパスワードが合致したら、新しいパスワードに更新する",
                "tags": [
                    "auth"
                ],
                "summary": "ログイン状態からパスワードを変更",
                "parameters": [
                    {
                        "description": "メールアドレス",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.UpdatePassRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "404": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "内部エラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/expedition/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "遠征、出費、試合、訪れた施設の情報を保存する。",
                "tags": [
                    "expedition"
                ],
                "summary": "遠征記録を作成",
                "responses": {
                    "200": {
                        "description": "アップロードした画像のURL",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "403": {
                        "description": "ユーザーが見つかりません",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "404": {
                        "description": "ユーザーが見つかりません",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "ユーザーが見つかりません",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/expedition/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "pathのidをもとに遠征記録を削除する。",
                "tags": [
                    "expedition"
                ],
                "summary": "遠征記録を削除",
                "responses": {
                    "200": {
                        "description": "アップロードした画像のURL",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "403": {
                        "description": "ユーザーが見つかりません",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "404": {
                        "description": "ユーザーが見つかりません",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "ユーザーが見つかりません",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/expedition/update/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "遠征、出費、試合、訪れた施設の情報を更新する。",
                "tags": [
                    "expedition"
                ],
                "summary": "遠征記録を更新",
                "responses": {
                    "200": {
                        "description": "アップロードした画像のURL",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "403": {
                        "description": "ユーザーが見つかりません",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "404": {
                        "description": "ユーザーが見つかりません",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "ユーザーが見つかりません",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/sample/helloWorld": {
            "get": {
                "tags": [
                    "sample"
                ],
                "summary": "サンプルAPI",
                "responses": {}
            }
        },
        "/api/sample/protectedHelloWorld": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "sample"
                ],
                "summary": "サンプルAPI",
                "responses": {}
            }
        },
        "/api/upload/images": {
            "post": {
                "description": "画像をアップロードし、URLを返します。",
                "tags": [
                    "upload"
                ],
                "summary": "画像をクラウドストレージにアップロード",
                "parameters": [
                    {
                        "type": "file",
                        "description": "画像ファイル",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/user/update/:id": {
            "put": {
                "description": "userIdが同じユーザーの情報を変更する",
                "tags": [
                    "user"
                ],
                "summary": "ユーザー情報変更",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ユーザー情報",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        },
        "/api/user/{id}": {
            "get": {
                "description": "userIdからユーザーを1人取得",
                "tags": [
                    "user"
                ],
                "summary": "ユーザー情報取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ユーザー情報",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/user.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "ユーザーが見つかりません",
                        "schema": {
                            "$ref": "#/definitions/user.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.EmailVerificationRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "auth.LoginRequest": {
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
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6
                }
            }
        },
        "auth.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "auth.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6
                },
                "profileImage": {
                    "type": "string"
                }
            }
        },
        "auth.UpdatePassRequest": {
            "type": "object",
            "required": [
                "afterPassword",
                "beforePassword"
            ],
            "properties": {
                "afterPassword": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6
                },
                "beforePassword": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6
                }
            }
        },
        "user.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "utils.ApiResponse-auth_LoginResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/auth.LoginResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "utils.BasicResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "ビジターゴーAPI",
	Description:      "このapiは、ビジターゴーのAPIで、ユーザー、スタジアム、遠征記録、などについて扱います。",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
