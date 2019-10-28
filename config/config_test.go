package config

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const anchorFile = "anchor.json"

var _ = Describe("config.go test", func() {

	Describe("#GetConfig", func() {

		BeforeEach(func() {
			os.Unsetenv("MP_PRODUCTION")
			os.Unsetenv("MP_HORIZON_URL")
			os.Unsetenv("MP_ANCHOR_FILE")
		})

		It("makes test config", func() {
			const testURL = "http://test"

			os.Setenv("MP_PRODUCTION", "false")
			os.Setenv("MP_HORIZON_URL", testURL)
			os.Setenv("MP_ANCHOR_FILE", anchorFile)

			makeConfig()
			c := GetConfig()
			Expect(c.IsProduction).To(BeFalse())
			Expect(c.HorizonURL).To(Equal(testURL))
			Expect(c.AnchorFile).To(Equal(anchorFile))
		})

		It("makes prod config", func() {
			const prodURL = "http://prod"

			os.Setenv("MP_PRODUCTION", "true")
			os.Setenv("MP_HORIZON_URL", prodURL)
			os.Setenv("MP_ANCHOR_FILE", anchorFile)

			makeConfig()
			c := GetConfig()
			Expect(c.IsProduction).To(BeTrue())
			Expect(c.HorizonURL).To(Equal(prodURL))
		})
	})

})