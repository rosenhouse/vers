// Package vers models semantic versioning for a runtime / plugin system.
//
// The system should be Forwards Compatible within Minor and Patch versions.
//
// A runtime calls a plugin passing in Config and expecting to receive Results.
// A single Spec defines the data types and behavior for Config and Results,
// but the Spec may evolve over time and be versioned.
//
// Breaking changes to the Spec, such as removal of a field from either the
// Config or Results, should cause a Major version bump.
// Adding a field to either Config or Results should cause a Minor version bump.
//
// Any given plugin will require Config at a particular Semantic Version.
// This means that if the runtime provides Config with greater Minor
// (or Patch) versions (but with equal Major versions) then the plugin should work.
//
// Similarly, if the plugin provides a Result with a Minor (or Patch) version
// higher than that required by the runtime, then the runtime should be able to
// accept that data (again given equality of Major version numbers)
//
// Plugins are assumed to be lightweight but a runtime may support many Modes,
// each of which could drive plugins at a single Major version.
package vers

import (
	"errors"

	"github.com/blang/semver"
)

// Spec represents the semantic version of the plugin specification
type Spec semver.Version

var ErrorIncompatible = errors.New("incompatible versions")

// RuntimeMode represents one (of possibly several) modes of operation
// of the runtime.  If one mode fails to check against a plugin, then the
// runtime may opt to fall back to an older mode which does check ok.
type RuntimeMode struct {
	ProvidesConfig  semver.Version
	RequiresResults semver.Version
}

// Plugin represents the required version of input Config and provided version
// of output Results for a single plugin.
type Plugin struct {
	RequiresConfig  semver.Version
	ProvidesResults semver.Version
}

func meetsReqs(provides semver.Version, requirement semver.Version) bool {
	if provides.Major != requirement.Major {
		return false
	}
	return provides.GTE(requirement)
}

// Check returns nil if the given RuntimeMode is compatible with the given
// Plugin.  Otherwise it returns an error.
//
// Semantic versioning applies, so Major version differences always cause an
// error.  Minor and patch version differences are allowed if and only if
// the provider version meets or exceeds the required version.
func Check(runtime RuntimeMode, plugin Plugin) error {
	configOK := meetsReqs(runtime.ProvidesConfig, plugin.RequiresConfig)
	resultOK := meetsReqs(plugin.ProvidesResults, runtime.RequiresResults)
	ok := configOK && resultOK

	if !ok {
		return ErrorIncompatible
	}

	return nil
}
