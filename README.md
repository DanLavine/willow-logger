# willow-logger
[godoc](https://pkg.go.dev/github.com/DanLavine/willow-logger)

Willow Logger is a shared logging setup package that is used for the Willow project.
It provides a number of easy to use conventions for setting up trace IDs for all API
requests in the log messages.

As it currently sits, this package is being used as part of a much larger refactor for Willow
to allow for horizontal scalability and might eventually be pulled back into a common shared package
that defines all of the shared code conventrions across will. For now at least, puting the code here
make the current work a bit easier to manage, but it might not be the final place where this package
ends up living