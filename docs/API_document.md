## Gardens

`GET /v1/gardens/{id}` return id's garden data as JSON respons

example respons JSON data :

```
{
    "_id" : "61262d8686f74f225c1dd0ae"
    "name" : "Fat Cat's garden"
    "description" : "Garden for my Fat Cat"
    "fireBaseUserId" : "jEazVdPDhqec0tnEOG7vM5wbDyU2"
}
```

`GET /v1/gardens/user/{userid}` return search userid's garden id, and return those ids as JSON data

example respons JSON data :

```
{
    "ids" : [61262d8686f74f225c1dd0ae, 52455shv4jji4fdfqd5opj5q]
}
```

`POST /v1/gardens/{id}` create gardens data using request JSON data, and If API success this, return objest id as JSON.

example request JSON data :

```
{
    "name" : "Fat Cat's garden"
    "description" : "Garden for my Fat Cat"
    "fireBaseUserId" : "jEazVdPDhqec0tnEOG7vM5wbDyU2"
}
```

example respons JSON data :

```
{
    "_id" : "61262d8686f74f225c1dd0a"
```
