package config

import (
	"github.com/stellar/go/network"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const anchorFile = "anchor.json"

var _ = Describe("config.go test", func() {

	Describe("#GetConfig", func() {

		BeforeEach(func() {
			os.Unsetenv("MP_ENV")
		})

		It("makes test config", func() {
			const testURL = "http://test"

			os.Setenv("MP_ENV", "local")

			c := &BaseConfig{}
			c.Parse()
			Expect(c.IsProduction()).To(BeFalse())
			Expect(c.IsQA()).To(BeFalse())
			Expect(c.IsDev()).To(BeFalse())
			Expect(c.IsLocal()).To(BeTrue())
			Expect(c.NetworkPassphrase).To(Equal(network.TestNetworkPassphrase))
		})

		It("makes prod config", func() {
			const prodURL = "http://prod"

			os.Setenv("MP_ENV", "prod")

			c := &BaseConfig{}
			c.Parse()
			Expect(c.IsProduction()).To(BeTrue())
			Expect(c.IsQA()).To(BeFalse())
			Expect(c.IsDev()).To(BeFalse())
			Expect(c.IsLocal()).To(BeFalse())
			Expect(c.NetworkPassphrase).To(Equal(network.PublicNetworkPassphrase))
		})
	})

})