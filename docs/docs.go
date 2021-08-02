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
        "/allochosts": {
            "post": {
                "description": "按指定的配置分配主机资源",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "分配主机接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "主机分配请求",
                        "name": "Alloc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hostapi.AllocHostsReq"
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
                                            "$ref": "#/definitions/hostapi.AllocHostsRsp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/download_template/": {
            "get": {
                "description": "将主机信息文件导出到本地",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "导出主机信息模板文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/host": {
            "post": {
                "description": "将给定的主机信息导入系统",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "导入主机接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "待导入的主机信息",
                        "name": "host",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hostapi.HostInfo"
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
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
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
                                                "$ref": "#/definitions/hostapi.DemoHostInfo"
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
        "/host/{hostId}": {
            "get": {
                "description": "展示指定的主机的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "查询主机详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "主机ID",
                        "name": "hostId",
                        "in": "path",
                        "required": true
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
                                            "$ref": "#/definitions/hostapi.HostInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "删除指定的主机",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "删除指定的主机",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "待删除的主机ID",
                        "name": "hostId",
                        "in": "path",
                        "required": true
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
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/hosts": {
            "get": {
                "description": "展示目前所有主机",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "查询主机列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "查询特定用途的主机列表",
                        "name": "purpose",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "查询特定状态的主机列表",
                        "name": "status",
                        "in": "query"
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
            },
            "post": {
                "description": "通过文件批量导入主机",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "通过文件批量导入主机",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "包含待导入主机信息的文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
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
                                            "type": "array",
                                            "items": {
                                                "type": "string"
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
        "/hosts/": {
            "delete": {
                "description": "批量删除指定的主机",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource"
                ],
                "summary": "批量删除指定的主机",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "待删除的主机ID数组",
                        "name": "hostIds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
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
                                            "type": "string"
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
        "hostapi.AllocHostsReq": {
            "type": "object",
            "properties": {
                "pdReq": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hostapi.Allocation"
                    }
                },
                "tidbReq": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hostapi.Allocation"
                    }
                },
                "tikvReq": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hostapi.Allocation"
                    }
                }
            }
        },
        "hostapi.AllocHostsRsp": {
            "type": "object",
            "properties": {
                "pdHosts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hostapi.AllocateRsp"
                    }
                },
                "tidbHosts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hostapi.AllocateRsp"
                    }
                },
                "tikvHosts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hostapi.AllocateRsp"
                    }
                }
            }
        },
        "hostapi.AllocateRsp": {
            "type": "object",
            "properties": {
                "cpuCore": {
                    "type": "integer"
                },
                "disk": {
                    "$ref": "#/definitions/hostapi.Disk"
                },
                "hostName": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "memory": {
                    "type": "integer"
                }
            }
        },
        "hostapi.Allocation": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "cpuCores": {
                    "type": "integer"
                },
                "failureDomain": {
                    "type": "string"
                },
                "memory": {
                    "type": "integer"
                }
            }
        },
        "hostapi.DemoHostInfo": {
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
        "hostapi.Disk": {
            "type": "object",
            "properties": {
                "capacity": {
                    "description": "Disk size, Unit: GB",
                    "type": "integer"
                },
                "diskId": {
                    "type": "string"
                },
                "name": {
                    "description": "[sda/sdb/nvmep0...]",
                    "type": "string"
                },
                "path": {
                    "description": "Disk mount path: [/data1]",
                    "type": "string"
                },
                "status": {
                    "description": "Disk Status, 0 for available, 1 for inused",
                    "type": "integer"
                }
            }
        },
        "hostapi.HostInfo": {
            "type": "object",
            "properties": {
                "az": {
                    "type": "string"
                },
                "cpuCores": {
                    "type": "integer"
                },
                "dc": {
                    "type": "string"
                },
                "disks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hostapi.Disk"
                    }
                },
                "hostId": {
                    "type": "string"
                },
                "hostName": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "kernel": {
                    "type": "string"
                },
                "memory": {
                    "description": "Host memory size, Unit:GB",
                    "type": "integer"
                },
                "nic": {
                    "description": "Host network type: 1GE or 10GE",
                    "type": "string"
                },
                "os": {
                    "type": "string"
                },
                "purpose": {
                    "description": "What Purpose is the host used for? [compute/storage or both]",
                    "type": "string"
                },
                "rack": {
                    "type": "string"
                },
                "status": {
                    "description": "Host Status, 0 for Online, 1 for offline",
                    "type": "integer"
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
                    "type": "string"
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
