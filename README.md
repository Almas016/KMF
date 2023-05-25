Write HTTP server for proxying HTTP-requests to 3rd-party services.

The server is waiting HTTP-request from client (curl, for example). In request's body there

should be message in JSON format. For example:

```md
{
    "method": "GET",
    "url": "http://google.com",
    "headers": {
        "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",
        ....
    }
}
```
Server forms valid HTTP-request to 3rd-party service with data from client's message and responses to client with JSON object:

```md
{
    "id": <generated unique id>,
    "status": <HTTP status of 3rd-party service response>,
    "headers": {
        <headers array from 3rd-party service response>
    },
    "length": <content length of 3rd-party service response>
}
```
Server should have map to store requests from client and responses from 3rd-party service.

