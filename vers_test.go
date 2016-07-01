package vers_test

import (
	"github.com/blang/semver"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/rosenhouse/vers"
)

var _ = Describe("checking compatibility", func() {
	var (
		runtime vers.RuntimeMode
		plugin  vers.Plugin
	)

	BeforeEach(func() {
		runtime = vers.RuntimeMode{
			ProvidesConfig:  semver.MustParse("2.5.7"),
			RequiresResults: semver.MustParse("2.5.7"),
		}
		plugin = vers.Plugin{
			RequiresConfig:  semver.MustParse("2.5.7"),
			ProvidesResults: semver.MustParse("2.5.7"),
		}
	})

	Context("when the versions are all equal", func() {
		It("succeeds", func() {
			Expect(vers.Check(runtime, plugin)).To(Succeed())
		})
	})

	Context("when the plugin requires config that is too new for the runtime", func() {
		BeforeEach(func() {
			plugin.RequiresConfig = semver.MustParse("2.5.8")
		})
		It("fails", func() {
			Expect(vers.Check(runtime, plugin)).To(Equal(vers.ErrorIncompatible))
		})
	})

	Context("when the runtime provides config that is newer than the plugin expects", func() {
		Context("when that new config is still within the same major version", func() {
			BeforeEach(func() {
				runtime.ProvidesConfig = semver.MustParse("2.5.8")
			})
			It("suceeds", func() {
				Expect(vers.Check(runtime, plugin)).To(Succeed())
			})
		})

		Context("when that new config is a major version ahead", func() {
			BeforeEach(func() {
				runtime.ProvidesConfig = semver.MustParse("3.0.0")
			})
			It("fails", func() {
				Expect(vers.Check(runtime, plugin)).To(Equal(vers.ErrorIncompatible))
			})
		})
	})

	Context("when the runtime requires results that are too new for the plugin", func() {
		BeforeEach(func() {
			runtime.RequiresResults = semver.MustParse("2.5.8")
		})
		It("fails", func() {
			Expect(vers.Check(runtime, plugin)).To(Equal(vers.ErrorIncompatible))
		})
	})

	Context("when the runtime requires results from an older version than the plugin provides", func() {
		Context("when they are still within the same major version", func() {
			BeforeEach(func() {
				runtime.RequiresResults = semver.MustParse("2.5.6")
			})
			It("succeeds", func() {
				Expect(vers.Check(runtime, plugin)).To(Succeed())
			})
		})

		Context("when the runtime requires a major version behind the plugin", func() {
			BeforeEach(func() {
				runtime.RequiresResults = semver.MustParse("1.9.9")
			})
			It("fails", func() {
				Expect(vers.Check(runtime, plugin)).To(Equal(vers.ErrorIncompatible))
			})
		})
	})
})
