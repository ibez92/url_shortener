package shorturl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ibez92/url_shortener/internal/pkg/shorturl"
)

var _ = Describe("Bijection alg", Label("bijection"), func() {
	It("Generate id by string", func(ctx SpecContext) {
		result := shorturl.IdByShortURL("cb")
		Expect(result).To(Equal(uint64(125)))
	})

	It("Generate string by id", func(ctx SpecContext) {
		result := shorturl.ShortURLByID(125)
		Expect(result).To(Equal("cb"))
	})
})

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bijection alg")
}
