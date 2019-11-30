// +build spec

package disclosure

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMoocro(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Moocro Suite")
}
