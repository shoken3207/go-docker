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
        "/api/auth/emailVerification/{email}": {
            "get": {
                "description": "リクエストからメールアドレス取得後、tokenTypeに応じてチェックし、メールアドレス宛にtokenを含めた画面URLをメールで送信",
                "tags": [
                    "auth"
                ],
                "summary": "メールアドレスの本人確認",
                "parameters": [
                    {
                        "type": "string",
                        "description": "メールアドレス",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "トークンタイプ register or reset",
                        "name": "tokenType",
                        "in": "query",
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
                "description": "メールアドレスとパスワードが合致したら、jwtトークンをクライアントに返却",
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
                "description": "メール内リンクで本人確認後、トークンと新しいパスワードをリクエストで取得し、パスワードを更新する",
                "tags": [
                    "auth"
                ],
                "summary": "ログアウト状態からパスワードを変更",
                "parameters": [
                    {
                        "description": "tokenと新しいパスワード",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ResetPassRequest"
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
        "/api/auth/updatePass/{userId}": {
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
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "メールアドレス",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.UpdatePassRequestBody"
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
                    "401": {
                        "description": "認証エラー",
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
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "画像をアップロードし、URLを返します。\u003cbr\u003eプロフィール、スタジアム、遠征など、格納フォルダを指定してください。\u003cbr\u003e画像は1枚から10枚アップロードできるが、Swagger UIでは1つしか選択できません。\u003cbr\u003eファイルの拡張子は、[\".jpg\", \".jpeg\", \".png\"]だけを受け付けています。ファイルサイズは最大5MBを上限としています。",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "upload"
                ],
                "summary": "画像をクラウドストレージ(imagekit)にアップロード",
                "parameters": [
                    {
                        "type": "string",
                        "description": "格納フォルダ",
                        "name": "folder",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "画像ファイル",
                        "name": "images",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse-upload_UploadImagesResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "401": {
                        "description": "認証エラー",
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
        "/api/user/isUnique/{username}": {
            "get": {
                "description": "リクエストと同じuserNameが登録済みかチェックする",
                "tags": [
                    "user"
                ],
                "summary": "ユーザーネームの重複チェック",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "一意かのフラグ",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse-user_IsUniqueUsernameResponse"
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
        "/api/user/logined": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "ヘッダーのトークンからユーザーを取得する",
                "tags": [
                    "user"
                ],
                "summary": "ログイン済みの場合、ログインユーザーの情報を取得",
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse-user_UserResponse"
                        }
                    },
                    "401": {
                        "description": "認証エラー",
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
        "/api/user/update/{userId}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "userIdが同じユーザーの情報を変更する",
                "tags": [
                    "user"
                ],
                "summary": "ユーザー情報変更",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "userId",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdateUserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ユーザー情報",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse-user_UserResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "401": {
                        "description": "認証エラー",
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
        "/api/user/userId/{userId}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "userIdからユーザーを1人取得",
                "tags": [
                    "user"
                ],
                "summary": "userIdからユーザー情報取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ユーザー情報",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse-user_UserResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "401": {
                        "description": "認証エラー",
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
        "/api/user/username/{username}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "usernameからユーザーを1人取得",
                "tags": [
                    "user"
                ],
                "summary": "usernameからユーザー情報取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ユーザー情報",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse-user_UserResponse"
                        }
                    },
                    "400": {
                        "description": "リクエストエラー",
                        "schema": {
                            "$ref": "#/definitions/utils.BasicResponse"
                        }
                    },
                    "401": {
                        "description": "認証エラー",
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
        }
    },
    "definitions": {
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
                "name",
                "password",
                "token",
                "username"
            ],
            "properties": {
                "description": {
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
                },
                "token": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 5
                }
            }
        },
        "auth.ResetPassRequest": {
            "type": "object",
            "required": [
                "afterPassword",
                "token"
            ],
            "properties": {
                "afterPassword": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "auth.UpdatePassRequestBody": {
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
        "upload.UploadImagesResponse": {
            "type": "object",
            "properties": {
                "urls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "user.IsUniqueUsernameResponse": {
            "type": "object",
            "properties": {
                "isUnique": {
                    "type": "boolean"
                }
            }
        },
        "user.UpdateUserRequestBody": {
            "type": "object",
            "required": [
                "description",
                "name",
                "profileImage"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "profileImage": {
                    "type": "string"
                }
            }
        },
        "user.UserResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "profileImage": {
                    "type": "string"
                },
                "username": {
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
        "utils.ApiResponse-upload_UploadImagesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/upload.UploadImagesResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "utils.ApiResponse-user_IsUniqueUsernameResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/user.IsUniqueUsernameResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "utils.ApiResponse-user_UserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/user.UserResponse"
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
