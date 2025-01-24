//go:build exclude

package csrf_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCSRF(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CSRF Policy Suite")
}
