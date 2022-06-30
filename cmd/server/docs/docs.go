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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/countries": {
            "get": {
                "description": "get all details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all countries details",
                "operationId": "get-countries-details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Country"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/countries/{id}": {
            "get": {
                "description": "get all details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get country details",
                "operationId": "get-country-details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Country"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update country",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update country based on parameter",
                "operationId": "update-country-details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Country"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Country": {
            "type": "object",
            "properties": {
                "code2": {
                    "description": "Code2 holds the value of the \"code2\" field.",
                    "type": "string"
                },
                "gnp": {
                    "description": "Gnp holds the value of the \"gnp\" field.",
                    "type": "number"
                },
                "gnp_old": {
                    "description": "GnpOld holds the value of the \"gnp_old\" field.",
                    "type": "number"
                },
                "government_form": {
                    "description": "GovernmentForm holds the value of the \"government_form\" field.",
                    "type": "string"
                },
                "head_of_state": {
                    "description": "HeadOfState holds the value of the \"head_of_state\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "indep_year": {
                    "description": "IndepYear holds the value of the \"indep_year\" field.",
                    "type": "integer"
                },
                "life_expectancy": {
                    "description": "LifeExpectancy holds the value of the \"life_expectancy\" field.",
                    "type": "number"
                },
                "local_name": {
                    "description": "LocalName holds the value of the \"local_name\" field.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "population": {
                    "description": "Population holds the value of the \"population\" field.",
                    "type": "integer"
                },
                "region": {
                    "description": "Region holds the value of the \"region\" field.",
                    "type": "string"
                },
                "surface_area": {
                    "description": "SurfaceArea holds the value of the \"surface_area\" field.",
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9090",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a server that gives information about countries",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}