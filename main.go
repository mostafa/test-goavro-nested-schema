package main

import (
	"encoding/json"
	"fmt"

	"github.com/linkedin/goavro/v2"
)

func main() {
	// Create a new codec for the Avro schema
	codec, err := goavro.NewCodec(`{
"type": "record",
  "name": "EmailMessageEventRestrictedInValue",
  "namespace": "internalservices.itoperation.emailmessage",
  "fields": [
    {
      "name": "payload",
      "type": {
        "type": "record",
        "name": "EmailMessage",
        "fields": [
          {
            "name": "id",
            "type": [
              "null",
              {
                "type": "string",
                "avro.java.string": "String"
              }
            ],
            "default": null
          },
          {
            "name": "smtpConfigName",
            "type": [
              "null",
              {
                "type": "string",
                "avro.java.string": "String"
              }
            ],
            "default": null
          },
          {
            "name": "emailAddressFrom",
            "type": {
              "type": "record",
              "name": "EmailAddressBasicInfo",
              "fields": [
                {
                  "name": "emailAddress",
                  "type": [
                    "null",
                    {
                      "type": "string",
                      "avro.java.string": "String"
                    }
                  ],
                  "default": null
                },
                {
                  "name": "alias",
                  "type": [
                    "null",
                    {
                      "type": "string",
                      "avro.java.string": "String"
                    }
                  ],
                  "default": null
                }
              ]
            }
          },
          {
            "name": "emailAddressTo",
            "type": {
              "type": "array",
              "items": "EmailAddressBasicInfo"
            }
          },
          {
            "name": "emailAddressCc",
            "type": [
              "null",
              {
                "type": "array",
                "items": "EmailAddressBasicInfo"
              }
            ],
            "default": null
          },
          {
            "name": "emailAddressBcc",
            "type": [
              "null",
              {
                "type": "array",
                "items": "EmailAddressBasicInfo"
              }
            ],
            "default": null
          },
          {
            "name": "subject",
            "type": [
              "null",
              {
                "type": "string",
                "avro.java.string": "String"
              }
            ],
            "default": null
          },
          {
            "name": "body",
            "type": [
              "null",
              {
                "type": "string",
                "avro.java.string": "String"
              }
            ],
            "default": null
          },
          {
            "name": "isHtmlBody",
            "type": [
              "null",
              "boolean"
            ],
            "default": null
          },
          {
            "name": "emailAddressReplyTo",
            "type": [
              "null",
              "EmailAddressBasicInfo"
            ],
            "default": null
          },
          {
            "name": "attachments",
            "type": [
              "null",
              {
                "type": "array",
                "items": {
                  "type": "record",
                  "name": "EmailMessageAttachment",
                  "fields": [
                    {
                      "name": "name",
                      "type": [
                        "null",
                        {
                          "type": "string",
                          "avro.java.string": "String"
                        }
                      ],
                      "default": null
                    },
                    {
                      "name": "attachment",
                      "type": [
                        "null",
                        {
                          "type": "string",
                          "avro.java.string": "String"
                        }
                      ],
                      "default": null
                    }
                  ]
                }
              }
            ],
            "default": null
          },
          {
            "name": "priority",
            "type": [
              "null",
              "int"
            ],
            "default": null
          }
        ]
      }
    }
  ]
}`)
	if err != nil {
		panic(err)
	}

	// Encode the data
	jsonPayload := []byte(`{
  "payload": {
    "id": {
      "string": "1234567890"
    },
    "smtpConfigName": {
      "string": "Config 1"
    },
    "emailAddressFrom": {
      "emailAddress": {
        "string": "nameapp@mail.test.com"
      },
      "alias": {
        "string": "nameapp"
      }
    },
    "emailAddressTo": [
      {
        "emailAddress": {
          "string": "dest1@test.com"
        },
        "alias": {
          "string": "dest 1"
        }
      }
    ],
    "emailAddressCc": {
      "array": [
        {
          "emailAddress": {
            "string": "dest2@test.com"
          },
          "alias": {
            "string": "dest 2"
          }
        }
      ]
    },
    "emailAddressBcc": {
      "array": [
        {
          "emailAddress": {
            "string": "dest3@test.com"
          },
          "alias": {
            "string": "dest 3"
          }
        }
      ]
    },
    "subject": {
      "string": "test info"
    },
    "body": {
      "string": "test info body"
    },
    "isHtmlBody": {
      "boolean": true
    },
    "emailAddressReplyTo": {
      "internalservices.itoperation.emailmessage.EmailAddressBasicInfo": {
        "emailAddress": {
          "string": "app-nameapp@test.com"
        },
        "alias": {
          "string": "App name App"
        }
      }
    },
    "attachments": {
      "array": [
        {
          "name": {
            "string": "attachmentfile.png"
          },
          "attachment": {
            "string": "http://www.test.com"
          }
        }
      ]
    },
    "priority": {
      "int": 0
    }
  }
}`)

	// Unmarshal the JSON payload
	var native map[string]interface{}
	err = json.Unmarshal(jsonPayload, &native)
	if err != nil {
		panic(err)
	}

	// Encode the data using the Avro codec and schema
	binary, err := codec.BinaryFromNative(nil, native)
	if err != nil {
		panic(err)
	}

	// Decode the data back to JSON (map[string]interfac{})
	decoded, _, err := codec.NativeFromBinary(binary)
	if err != nil {
		panic(err)
	}

	// Print out the original and decoded data for comparison
	fmt.Println("Original: ", native)
	fmt.Println("Decoded:  ", decoded)

	// Check if the original and decoded data are equal
	fmt.Println("Original and Decoded are equal: ", fmt.Sprint(native) == fmt.Sprint(decoded))
}
