[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/UlboraTemplateService)](https://goreportcard.com/report/github.com/Ulbora/UlboraTemplateService)

Ulbora Template Service
==============

A template micro service for CMS and Shopping Cart use


## Headers
- Content-Type: application/json (for POST and PUT)
- Authorization: Bearer aToken (POST, PUT, and DELETE. No token required for get services.)
- clientId: clientId (example 33477)



## Add Template

```
POST:
URL: http://localhost:3009/rs/template/add

Example Request
{
   "name":"new template2",
   "application": "cms"   
}
  
```

```
Example Response   

{
    "success": true,
    "id": 19
}

```


## Update Template Set Active

```
PUT:
URL: http://localhost:3009/rs/template/updateActive

Example Request
{
   "id": 88,
   "application": "cms"
   
}
  
```

```
Example Response   

{
    "success": true,
    "id": 11
}

```


## Get Active Template

```
GET:
URL: http://localhost:3009/rs/template/get/cms/403
  
```

```
Example Response   

{
    "id": 88,
    "name": "newtemplate2",
    "application": "cms",
    "active": true,
    "clientId": 403
}

```


## Get Templates for a Client

```
GET:
URL: http://localhost:3009/rs/template/list/cms/403
  
```

```
Example Response   

[
    {
        "id": 86,
        "name": "newtemplate2",
        "application": "cms",
        "active": false,
        "clientId": 403
    },
    {
        "id": 87,
        "name": "newtemplate2",
        "application": "cms",
        "active": false,
        "clientId": 403
    },
    {
        "id": 88,
        "name": "newtemplate2",
        "application": "cms",
        "active": true,
        "clientId": 403
    }
]

```


## Delete Templates for a Client

```
DELETE:
URL: http://localhost:3009/rs/template/delete/107/403
  
```

```
Example Response   

{
    "success": true,
    "id": 107
}

```

# Docker usage

```
docker run --network=ulbora_bridge --name templates --log-opt max-size=50m --env DATABASE_HOST=someHost /
 --env DATABASE_USER_NAME=someName --env DATABASE_USER_PASSWORD=somePw --env DATABASE_NAME=ulbora_template_service /
 --env DATABASE_POOL_SIZE=5 --env OAUTH2_VALIDATION_URI=http://oauth2:8080/rs/token/validate --env PORT=8080 /
 --restart=always -d ulboralabs/templates sh
```

