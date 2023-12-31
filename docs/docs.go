// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/apple/": {
            "get": {
                "description": "get apple jobs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jobs"
                ],
                "summary": "Get Apple Jobs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MicrosoftResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/microsoft/": {
            "get": {
                "description": "get microsoft jobs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jobs"
                ],
                "summary": "Get microsoft Jobs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MicrosoftResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.MicrosoftResponse": {
            "type": "object",
            "properties": {
                "jobId": {
                    "type": "string"
                },
                "postedUrl": {
                    "type": "string"
                },
                "postingDate": {
                    "type": "string"
                },
                "properties": {
                    "$ref": "#/definitions/model.Properties"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.Properties": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "educationLevel": {
                    "type": "string"
                },
                "employmentType": {
                    "type": "string"
                },
                "jobType": {
                    "type": "string"
                },
                "locations": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "primaryLocation": {
                    "type": "string"
                },
                "profession": {
                    "type": "string"
                },
                "roleType": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
