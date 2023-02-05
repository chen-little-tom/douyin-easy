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
        "/douyin/feed/": {
            "get": {
                "description": "获取视频流",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "基础模块"
                ],
                "summary": "获取视频流",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "最后一个视频时间戳",
                        "name": "lastTime",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.FeedResponse"
                        }
                    }
                }
            }
        },
        "/douyin/publish/action/": {
            "post": {
                "description": "上传视频",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "基础模块"
                ],
                "summary": "上传视频",
                "parameters": [
                    {
                        "description": "用户名",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "file",
                        "description": "视频文件",
                        "name": "data",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        },
        "/douyin/user/": {
            "get": {
                "description": "用户详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "基础模块"
                ],
                "summary": "用户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.UserDetailResponse"
                        }
                    }
                }
            }
        },
        "/douyin/user/login/": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "基础模块"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.UserLoginResponse"
                        }
                    }
                }
            }
        },
        "/douyin/user/register/": {
            "post": {
                "description": "注册用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "基础模块"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.UserLoginResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.FeedResponse": {
            "type": "object",
            "properties": {
                "next_time": {
                    "type": "integer"
                },
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "video_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.VideoVo"
                    }
                }
            }
        },
        "api.Response": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                }
            }
        },
        "api.UserDetailResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/vo.UserVo"
                }
            }
        },
        "api.UserLoginResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "vo.UserVo": {
            "type": "object",
            "properties": {
                "follow_count": {
                    "description": "关注总数",
                    "type": "integer"
                },
                "follower_count": {
                    "description": "粉丝总数",
                    "type": "integer"
                },
                "id": {
                    "description": "用户id",
                    "type": "integer"
                },
                "is_follow": {
                    "description": "是否关注",
                    "type": "boolean"
                },
                "name": {
                    "description": "用户名称",
                    "type": "string"
                }
            }
        },
        "vo.VideoVo": {
            "type": "object",
            "properties": {
                "author": {
                    "description": "作者",
                    "$ref": "#/definitions/vo.UserVo"
                },
                "commentCount": {
                    "description": "评论数",
                    "type": "integer"
                },
                "cover_url": {
                    "description": "视频封面地址",
                    "type": "string"
                },
                "favoriteCount": {
                    "description": "收到的喜欢数目",
                    "type": "integer"
                },
                "id": {
                    "description": "视频id",
                    "type": "string"
                },
                "is_favorite": {
                    "description": "是否喜欢",
                    "type": "boolean"
                },
                "play_url": {
                    "description": "播放地址",
                    "type": "string"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
