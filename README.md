# http-smoke-tests

A generic re-usable test suite using [ginkgo](https://github.com/onsi/ginkgo) that can be used for testing the HTTP response codes and body of a URL.

### Configuration

These should all be provided as environment variables.

* `RESPONSE_CODE`: *Required.* The acceptable response code expected from the HTTP request.
* `RESPONSE_BODY_REGEX`: *Required.* A REGEX matcher for the acceptable response body.
* `URL`: *Required.* The URL to test
* `REQUEST_METHOD`: *Optional.* Defaults to `GET`
* `SKIP_SSL_VERIFICATION`: *Optional.* If set to `true` skips SSL verification
* `HEADERS`: *Optional.* Set any required HTTP headers on the HTTP request. For example `{"HOST": "google.com"}`

### Usage

Example:

With the `ginkgo` CLI

```
RESPONSE_CODE=200 RESPONSE_BODY_REGEX=".*" URL="https://github.com" ginkgo -r
```

With `go test`

```
RESPONSE_CODE=200 RESPONSE_BODY_REGEX=".*" URL="https://github.com" go test -v ./...
```
