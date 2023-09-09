## Arbitrary Code Execution Demo

Have you ever wondered what happens behind the scenes when you hit "Run" on a code snippet in online development environments like [Go Playground](https://go.dev/play/) or [Repl.it](https://replit.com/)?

This demonstration provides a basic implementation using approximately 100 lines of code.

The main considerations for building a service such as this are:

1.  **Security**: Since we allow users to execute arbitrary code, we need a way to sandbox the execution to prevent malicious users from gaining unauthorized access to our servers.

2.  **Limits**: We want to limit the resources allocated to the execution so that it does not tax our finite server resources.

3.  **Scalability**: We need to be able to linearly scale this service as the number of users grows.

These considerations are addressed in this demonstration by utilizing [Tork](https://github.com/runabol/tork), a distributed workflow engine, to handle all the heavy lifting

## Running the demo

Make sure you have [Go](https://golang.org/) version 1.19 or better installed.

Start the server:

```bash
go run main.go run standalone
```

Execute a code snippet. Example

```bash
curl \
  -s \
  -X POST \
  -H "content-type:application/json" \
  -d '{"language":"python3","code":"print(\"hello world\")"}' \
  http://localhost:8000/exec
```

Should output:

```bash
hello world
```

You can try changing the `language` to `golang` or `bash`.
