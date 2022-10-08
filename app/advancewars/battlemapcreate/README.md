# Start an advance wars battle

By making a POST request to this endpoint the client can start a new battle between the authenticated user and an AI.

The API responds with the `map_id` which the API client can use to initialize new battles on this map

Example
```
POST /aw/map
Authenticated-User XXX-XXXXX-XXXXX-XX

{
    
}
```