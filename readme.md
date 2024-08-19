# Learn Go With Tests

## Writing Tests in Go - Rules
- It needs to be in a file with the `_test.go` suffix
- The test function must start with the word `Test*`
- The test function takes one argument only, `t *testing.T` 
- To use the `*testing.T` argument, you need to import the `testing` package. Same as 'fmt' package.

Think of 't' as your hook into the testing framework. You can use it to assert things, and report errors.

## Quality of Life - Go doc

Another quality-of-life feature of Go is the documentation. 
You can launch the docs locally by running:
```terminal
godoc -http=localhost:8000
``` 
If you go to localhost:8000/pkg, you will see all the packages installed on your system.

If you don't have godoc command, then maybe you are using the newer version of Go (1.14 or later) which no longer includes godoc. 
You can manually install it using:

```terminal 
go install golang.org/x/tools/cmd/godoc@latest
```

