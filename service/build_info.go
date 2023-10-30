package service

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/leliuga/validation"
)

var (
	CommitRegex       = regexp.MustCompile(`^[a-f0-9]{7,40}$`)
	PlatformRegex     = regexp.MustCompile(`^(linux\/(amd64|arm64))$`)
	ArchitectureRegex = regexp.MustCompile(`^(amd64|arm64)$`)
	OSRegex           = regexp.MustCompile(`^linux$`)
)

const (
	InvalidCommit       = "A commit must be a valid git commit hash."
	InvalidPlatform     = "A platform must be a valid, supported platforms are: linux/amd64, linux/arm64."
	InvalidOS           = "An operating system must be a valid, supported operating systems are: linux."
	InvalidArchitecture = "An architecture must be a valid, supported architectures are: amd64, arm64."
)

// NewBuildInfo creates a new BuildInfo.
func NewBuildInfo(repository, commit, when string) *BuildInfo {
	return &BuildInfo{
		Repository:   repository,
		Commit:       commit,
		When:         when,
		GoVersion:    strings.Trim(runtime.Version(), "go"),
		Platform:     runtime.GOOS + "/" + runtime.GOARCH,
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
	}
}

// Validate returns true if the BuildInfo is valid.
func (b *BuildInfo) Validate() error {
	return validation.ValidateStruct(b,
		validation.Field(&b.Repository, validation.Required),
		validation.Field(&b.Commit, validation.Required, validation.Length(7, 40), validation.Match(CommitRegex).Error(InvalidCommit)),
		validation.Field(&b.When, validation.Required),
		validation.Field(&b.GoVersion, validation.Required),
		validation.Field(&b.Platform, validation.Required, validation.Match(PlatformRegex).Error(InvalidPlatform)),
		validation.Field(&b.Architecture, validation.Required, validation.Match(ArchitectureRegex).Error(InvalidArchitecture)),
		validation.Field(&b.OS, validation.Required, validation.Match(OSRegex).Error(InvalidOS)),
	)
}
