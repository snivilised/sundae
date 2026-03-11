package assist_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"   
)

func TestAssist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assist Suite")
}
