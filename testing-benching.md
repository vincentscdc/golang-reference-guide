# Testing and benchmarking

## Tests

### Unit tests

If you can't test it, your design is wrong. You might be missing an interface or did not split your code enough.

You should use go unit tests, if possible don't use external libraries because they look more like the tests you're doing in ruby or javascript.

vscode provides you with a "right click" on your function name, that allow you to auto generate unit tests for the function.

Also, always test for [goroutine leaks](https://github.com/uber-go/goleak) and -race., that should be detailed in the example folder.

### Fuzzing

Since go1.18, fuzzing is included in go.
If you want to test a wide variety of inputs, you now can do it from within your test file

### E2E tests

Try as much as possible to test your whole API at every level, then the e2e test will have not much to do.

## Benchmarks

Go has built in benchmarking (TODO: link to the example API).
Try to bench when you release, and keep track of the benchmark, it can help debug perf issues later on.

## k6

You can use k6 to bench and even better, profile your API endpoints.
