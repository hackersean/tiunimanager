// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "zhangpeijin",
            "email": "zhangpeijin@pingcap.com"
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
        "/host/query": {
            "post": {
                "description": "查询主机",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "查询主机接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "查询请求",
                        "name": "query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hostapi.HostQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.ResultWithPage"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/hostapi.HostInfo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/instance/create": {
            "post": {
                "description": "创建实例",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instance"
                ],
                "summary": "创建实例接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建参数",
                        "name": "createInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/instanceapi.InstanceCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.CommonResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/instanceapi.InstanceInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/instance/query": {
            "post": {
                "description": "查询实例",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instance"
                ],
                "summary": "查询实例接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "查询参数",
                        "name": "query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/instanceapi.InstanceQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.ResultWithPage"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/instanceapi.InstanceInfo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "description": "登录用户信息",
                        "name": "loginInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userapi.LoginInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.CommonResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/userapi.UserIdentity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "description": "退出登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform"
                ],
                "summary": "退出登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "退出登录信息",
                        "name": "logoutInfo",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/userapi.LogoutInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.CommonResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/userapi.UserIdentity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CommonResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "controller.Page": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "controller.ResultWithPage": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "page": {
                    "$ref": "#/definitions/controller.Page"
                }
            }
        },
        "hostapi.HostInfo": {
            "type": "object",
            "properties": {
                "hostId": {
                    "type": "string"
                },
                "hostIp": {
                    "type": "string"
                },
                "hostName": {
                    "type": "string"
                }
            }
        },
        "hostapi.HostQuery": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "instanceapi.InstanceCreate": {
            "type": "object",
            "properties": {
                "dbPassword": {
                    "type": "string"
                },
                "instanceName": {
                    "type": "string"
                },
                "instanceVersion": {
                    "type": "string"
                },
                "pdCount": {
                    "type": "integer"
                },
                "tiDBCount": {
                    "type": "integer"
                },
                "tiKVCount": {
                    "type": "integer"
                }
            }
        },
        "instanceapi.InstanceInfo": {
            "type": "object",
            "properties": {
                "instanceId": {
                    "type": "string"
                },
                "instanceName": {
                    "type": "string"
                },
                "instanceStatus": {
                    "type": "integer"
                },
                "instanceVersion": {
                    "type": "integer"
                }
            }
        },
        "instanceapi.InstanceQuery": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "userapi.LoginInfo": {
            "type": "object",
            "properties": {
                "userName": {
                    "type": "string"
                },
                "userPassword": {
                    "type": "string"
                }
            }
        },
        "userapi.LogoutInfo": {
            "type": "object",
            "properties": {
                "userName": {
                    "type": "string"
                }
            }
        },
        "userapi.UserIdentity": {
            "type": "object",
            "properties": {
                "userName": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1/",
	Schemes:     []string{},
	Title:       "TiCP UI API",
	Description: "TiCP UI API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
