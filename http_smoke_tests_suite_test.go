package http_smoke_tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHttpSmokeTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HttpSmokeTests Suite")
}
