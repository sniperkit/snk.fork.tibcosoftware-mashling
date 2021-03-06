---
title: Mashling Gateway Model
weight: 4110
pre: "<i class=\"fa fa-asterisk\" aria-hidden=\"true\"></i> "
---

A gateway configuration file is what contains all details related to the runtime behavior of a mashling-gateway instance. The file can be named anything and pointed to via the -c or --config flag.

A gateway configuration specifies the appropriate schema version to load and validate against via the mashling_schema key. This is located at the top level of the configuration JSON schema. All other components specifying runtime behavior are contained within a gateway key and will be explained in detail below.

Example configuration files for the 1.0 schema version can be found in the example recipes folder. The corresponding schema can be found here.

At the root level, the Mashling gateway model contains the following objects:

1. Triggers
1. Dispatches
1. Routes
1. Steps
1. Services
1. Responses
1. Policies

#### Triggers
Triggers in Mashling are, currently, just Flogo triggers. Any Flogo trigger should work with the 1.0 schema specification. For the purposes of most of our examples, Mashling implemented triggers that conform to Flogo's specification are used.

An example trigger that listens for and dispatches HTTP requests looks like:

```
{
  "name": "MyProxy",
  "description": "Animals rest trigger - PUT animal details",
  "type": "github.com/TIBCOSoftware/mashling/ext/flogo/trigger/gorillamuxtrigger",
  "settings": {
    "port": "9096"
  },
  "handlers": [
    {
      "dispatch": "Pets",
      "settings": {
        "autoIdReply": "false",
        "method": "GET",
        "path": "/pets/{petId}",
        "useReplyHandler": "false"
      }
    }
  ]
}
```

You can map the execution of a trigger to a specific dispatch via the handlers array in the trigger configuration. This allows you to send the execution context to a different process flow based off of some specific settings. The dispatch value must map to a name in the dispatches array.

#### Dispatches
Dispatches are used to map trigger invocation with a set of possible execution routes. A dispatch has a name and receives the execution context from a trigger when that name is mapped via the trigger's handler. A dispatch is simple a name and an array of routes. A simple dispatch looks like:

```
{
  "name": "Pets",
  "routes": ["..."]
}
```

#### Routes
Routes define the actual exection logic of a dispatch. Each route in a dispatch comes with a condition value in the if key. The mashling engine will evaluate this condition within the trigger context. The first route with a condition to evaulate to true will then be the route executed. Only one route is executed per triggered flow. Once a route is selected by the mashling engine the steps defined therein will be evaluated and executed in the order in which they are defined. If a route is marked as "async": true then the execution will be asynchronous and the trigger will immediately be returned a response.

A simple route looks like:

```
{
  "if": "payload.pathParams.petId >= 8 && payload.pathParams.petId <= 15",
  "async": false,
  "steps": ["..."]
}
```

#### Steps
Each route is composed of a number of steps. Each step is evaluated in the order in which it is defined via an optional if condition. If the condition is true, that step is executed. If that condition is false the execution context moves onto the next step in the process and evaluates that one. A blank or omitted if condition always evaluates to true.

A simple step looks like:

```
{
  "if": "payload.pathParams.petId == 9",
  "service": "PetStorePets",
  "input": {
    "method": "GET",
    "pathParams.id": "${payload.pathParams.petId}"
  }
}
```

As you can see above, a step consists of a simple condition, a service reference, input parameters, and (not shown) output parameters. The service must map to a service defined in the services array that is defined outside of a dispatch. Input key and value pairs are translated and handed off to the service execution. Output key value pairs are translated and retained after the service has executed. Values wrapped with ${} are evaluated as variables within the context of the execution.

#### Services
A service defines a function or activity of some sort that will be utilized in a step within an execution flow. Services have names, types, and settings. Currently supported types are http, js, and flogoActivity. Services may call external endpoints like HTTP servers or may stay within the context of the mashling gateway, like the js service. Once a service is defined it can be used as many times as needed within your routes and steps.

A simple http service looks like:

```
{
  "name": "PetStorePets",
  "description": "Make calls to find pets",
  "type": "http",
  "settings": {
    "url": "http://petstore.swagger.io/v2/pet/:id"
  }
}
```

#### Responses
Each route has an optional set of responses that can be evaluated and returned to the invoking trigger. Much like routes, the first response with an if condition evaluating to true is the reponse that gets executed and returned. A response contains an if condition, an error boolean, a complex boolean, and an output object. The error boolean dictates whether or not an error should be returned to the engine. The complex boolean dictates whether to use the Reply or ReplyWithData function. A value of true causes the ReplyWithData function to be used when sending the response back to the trigger. The output is evaluated within the context of the execution and then sent back to the trigger as well.

A simple response looks like:

```
{
  "if": "PetStorePets.response.body.status == 'available'",
  "error": false,
  "complex": false,
  "output": {
    "code": 200,
    "format": "json",
    "body.pet": "${PetStorePets.response.body}",
    "body.inventory": "${PetStoreInventory.response.body}"
  }
}
```

#### Policies
Policies are called out in the JSON Schema and the types for the V2 package, however, they are not yet implemented. This section of the document outlines the third iteration of a proposed policy design. This has been reworked following feedback from two previous sessions with the team.

The new proposed implementation is to treat policies as distinct entities from services and to make each policy invocation atomic. The notion of hooks for policies are also introducted in this design. As with most entities in the model, a conditional expression is optional and is mostly useful for after policy hooks and for feedback into a policy that is invoked in the corresponding before hook. Lifecycle hook specification is optional. If it is omitted the behavior for all policies specified under that policies key is the same as if the before hook was used.

This iteration of the policy design adds a policy block to dispatches and also expands the schema definition of the policy invocation blocks to introduce the notion of hooks. These hooks look like beforeRoute, afterRoute, beforeStep, afterStep, etc... and dictate the invocation order for the included policies. The ability to add a one off lower level invocation can be achieived by adding the policy to the policies key in that lower level entity.

Unlike the previous proposals, an interrupt is not required to achieve any of the example policies outlined below. An interrupt is left in the example below simply because it is a useful flow construct, but it is not required for policies to function.

Providing these entry points to polices allows one to support something simple like a rate limiter that returns a simple yes or no before executing the steps. It also provides the ability to wrap a call in a circuit breaker via the after[Dispatch|Route|Step] policy hook.

##### Simple Policy Example
A simple HTTP proxy example with two policies (rate limiter and API key validation) added before the HTTP backend invocation happens is below. This example also demonstrates a simplified way of declaring a policy invocation: the before and after lifecycle hooks are excluded resulting in a default of before invocation behavior.

```
{
  "mashling_schema": "1.0",
  "gateway": {
    "name": "MyProxy",
    "version": "1.0.0",
    "description": "This is a simple proxy.",
    "triggers": [
      {
        "name": "MyProxy",
        "description": "Animals rest trigger - PUT animal details",
        "type": "github.com/TIBCOSoftware/mashling/ext/flogo/trigger/gorillamuxtrigger",
        "settings": {
          "port": "9096"
        },
        "handlers": [
          {
            "dispatch": "Pets"
          }
        ]
      }
    ],
    "dispatches": [
      {
        "name": "Pets",
        "routes": [
          {
            "policies": [
              {
                "policy": "GlobalRateLimiter",
                "input": {
                  "key": "${payload.ipAddress}"
                }
              },
              {
                "policy": "APIKeyAuth",
                "input": {
                  "key": "${payload.headers.APIKey}"
                }
              }
            ],
            "steps": [
              {
                "service": "MySpecialBackend",
                "input": {
                  "pathParams.id": "${payload.pathParams.petId}"
                }
              }
            ],
            "responses": [
              {
                "output": {
                  "code": "${MySpecialBackend.response.code}",
                  "format": "json",
                  "body.pet": "${MySpecialBackend.response.body}",
                  "body.inventory": "${MySpecialBackend.response.body}"
                }
              }
            ]
          }
        ]
      }
    ],
    "services": [
      {
        "name": "MySpecialBackend",
        "description": "Make calls to do stuff",
        "type": "http",
        "settings": {
          "url": "http://petstore.swagger.io/v2/pet/:id"
        }
      }
    ],
    "policies": [
      {
        "name": "GlobalRateLimiter",
        "description": "Rate limit all requests",
        "type": "rateLimiter",
        "settings": {
          "perSecond": 100
        }
      },
      {
        "name": "APIKeyAuth",
        "description": "Test API key.",
        "type": "apiKeyAuth",
        "settings": {
          "url": "https://www.somewherespecial.com"
        }
      }
    ]
  }
}
```

##### Complex Policy Example
A complex configuration file that has a contrived example using all of the hooks is as follows:

```
{
  "mashling_schema": "1.0",
  "gateway": {
    "name": "PolicyExample",
    "version": "1.0.0",
    "description": "This is a simple proxy.",
    "triggers": [
      {
        "name": "MyProxy",
        "description": "Animals rest trigger - PUT animal details",
        "type": "github.com/TIBCOSoftware/mashling/ext/flogo/trigger/gorillamuxtrigger",
        "settings": {
          "port": "9096"
        },
        "handlers": [
          {
            "dispatch": "Pets",
            "settings": {
              "autoIdReply": "false",
              "method": "GET",
              "path": "/pets/{petId}",
              "useReplyHandler": "false"
            }
          }
        ]
      }
    ],
    "dispatches": [
      {
        "name": "Pets",
        "policies": {
          "beforeDispatch": [
            {
              "policy": "Splunk"
            }
          ],
          "afterDispatch": [
            {
              "policy": "Splunk"
            }
          ]
        },
        "routes": [
          {
            "if": "payload.pathParams.petId >= 8 && payload.pathParams.petId <= 15",
            "policies": {
              "beforeRoute": [
                {
                  "policy": "GlobalRateLimiter",
                  "input": {
                    "key": "${payload.ipAddress}"
                  }
                },
                {
                  "policy": "CircuitBreaker"
                }
              ],
              "beforeStep": [
                {
                  "policy": "Splunk"
                }
              ],
              "afterStep": [
                {
                  "policy": "Splunk"
                }
              ],
              "beforeResponse": [
                {
                  "policy": "Splunk"
                }
              ],
              "afterResponse": [
                {
                  "policy": "Splunk"
                }
              ],
              "beforeInterrupt": [
                {
                  "policy": "Splunk"
                }
              ],
              "afterInterrupt": [
                {
                  "policy": "Splunk"
                }
              ],
              "afterRoute": [
                {
                  "if": "PetStorePets.response.error == true",
                  "policy": "CircuitBreaker",
                  "input": {
                    "failed": true
                  }
                }
              ]
            },
            "steps": [
              {
                "service": "PetStorePets",
                "input": {
                  "method": "GET",
                  "pathParams.id": "${payload.pathParams.petId}"
                },
                "interrupt": "PetStorePets.response.error == true"
              },
              {
                "if": "PetStorePets.response.body.status == 'available'",
                "policies": {
                  "beforeStep": [
                    {
                      "policy": "OneOffPolicyInvocationForJustThisStep"
                    }
                  ]
                },
                "service": "PetStoreInventory",
                "input": {
                  "method": "GET"
                }
              }
            ],
            "interrupts": [
              {
                "if": "PetStorePets.response.error == true",
                "service": "RemoteErrorNotification",
                "input": {
                  "body.message": "${PetStorePets.response.errorMessage}"
                }
              }
            ],
            "responses": [
              {
                "if": "payload.pathParams.petId == 13",
                "error": true,
                "output": {
                  "code": 404,
                  "format": "json",
                  "body": "petId is invalid"
                }
              },
              {
                "if": "PetStorePets.response.body.status != 'available'",
                "error": true,
                "output": {
                  "code": 403,
                  "format": "json",
                  "body": "Pet is unavailable."
                }
              },
              {
                "if": "PetStorePets.response.body.status == 'available'",
                "error": false,
                "output": {
                  "code": 200,
                  "format": "json",
                  "body.pet": "${PetStorePets.response.body}",
                  "body.inventory": "${PetStoreInventory.response.body}"
                }
              }
            ]
          }
        ]
      }
    ],
    "services": [
      {
        "name": "PetStorePets",
        "description": "Make calls to find pets",
        "type": "http",
        "settings": {
          "url": "http://petstore.swagger.io/v2/pet/:id"
        }
      },
      {
        "name": "PetStoreInventory",
        "description": "Get pet store inventory.",
        "type": "http",
        "settings": {
          "url": "http://petstore.swagger.io/v2/store/inventory"
        }
      },
      {
        "name": "RemoteErrorNotification",
        "description": "Send error details somewhere custom.",
        "type": "http",
        "settings": {
          "method": "POST",
          "url": "http://www.errorsarebad.io/report_error"
        }
      }
    ],
    "policies": [
      {
        "name": "GlobalRateLimiter",
        "description": "Rate limit all requests",
        "type": "rateLimiter",
        "settings": {
          "perSecond": 100
        }
      },
      {
        "name": "CircuitBreaker",
        "description": "Stop hitting broken routes.",
        "type": "circuitBreaker",
        "settings": {
          "maxFails": 5
        }
      },
      {
        "name": "Splunk",
        "description": "Send my information to Splunk.",
        "type": "splunk",
        "settings": {
          "format": "${time} - ${error} - ${message}"
        }
      }
    ]
  }
}
```
