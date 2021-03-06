// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Simple To Do List API",
    "version": "0.1.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "find",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "limit",
            "in": "formData",
            "required": true,
            "allowEmptyValue": true
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "X-Rate-Limit",
            "in": "header",
            "required": true
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "name": "words",
            "in": "formData",
            "required": true,
            "allowEmptyValue": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "todos"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "put": {
        "tags": [
          "todos"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "todos"
        ],
        "operationId": "destroyOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "properties": {
        "attributes": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "x-nullable": false
        },
        "code": {
          "type": "string",
          "x-nullable": false
        },
        "message": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Simple To Do List API",
    "version": "0.1.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "find",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "limit",
            "in": "formData",
            "required": true,
            "allowEmptyValue": true
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "X-Rate-Limit",
            "in": "header",
            "required": true
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "name": "words",
            "in": "formData",
            "required": true,
            "allowEmptyValue": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "todos"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "put": {
        "tags": [
          "todos"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "todos"
        ],
        "operationId": "destroyOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "properties": {
        "attributes": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "x-nullable": false
        },
        "code": {
          "type": "string",
          "x-nullable": false
        },
        "message": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  }
}`))
}
