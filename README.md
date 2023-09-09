## Demo

Ever wondered what happens behind the scenes when you hit "Run" on a code snippet in online development environments like [Go Playground](https://go.dev/play/) or [Repl.it](https://replit.com/)?

This demo provides a basic implementation using ~100 lines of code.

The main considerations for building a service such as this are:

1. Security - since we are allowing the user to execute arbitrary code we need a way to sandbox the execution so a malicious user does not gain unauthorized access to our servers. We also want to limit the resources allocated for the execution so it does not tax our limits server resources.

2. Scalability - we need to be able to linearly scale this service as the number of users grow.

Both this considerations are addressed in this demo by using [Tork](https://github.com/runabol/tork) which is a distributed workflow engine to do all the heavy lifting.
