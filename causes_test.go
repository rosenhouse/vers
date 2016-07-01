package vers_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("how the Spec should evolve", func() {
	Context("when a new field is added", func() {
		It("should force a minor (or major) semver bump in the spec", func() {})
	})

	Context("when a field is removed from the spec inputs", func() {
		It("should force a major semver bump in the spec", func() {})
	})
})

var _ = Describe("Requires and Provides versions", func() {

	Context("when a new field is added to Config", func() {
		Context("when a plugin requires the new Config field", func() {
			It("should increase its RequiresConfig version to match", func() {})
		})

		Context("when a runtime provides the new Config field", func() {
			It("should increase its ProvidesConfig version to match", func() {})
		})
	})

	Context("when a new field is added to Results", func() {
		Context("when a plugin provides the new Results field", func() {
			It("should increase its ProvidesResult version to match", func() {})
		})

		Context("when a runtime requires the new Results field", func() {
			It("should increase its RequiresResult version to match", func() {})
		})
	})
})
