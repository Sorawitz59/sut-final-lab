package unit

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"product.com/entity"
)

func TestPrice(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`Price is required`, func(t *testing.T) {
		product := entity.Products{
			Name: "ABC",
			Price: 0,
			SKU: "ABC12345",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price is required"))
	})
}