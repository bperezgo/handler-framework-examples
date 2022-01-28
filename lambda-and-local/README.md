# Lambda and Local example

To test how to use a function in a local server, with a url, and to be interchangable with execution in lambda function. And it is possible to see it can create more integrations with other architectures

To start up the local server, use the next script

```bash
DEV_ENV=LOCAL go run .
```

And you can use some request with curl or postman etc.
```bash
curl -X POST --data '{"value": "message"}' http://localhost:3000/
```