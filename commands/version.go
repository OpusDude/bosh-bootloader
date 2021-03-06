package commands

import (
	"fmt"
	"runtime"

	"github.com/cloudfoundry/bosh-bootloader/storage"
)

const (
	VersionCommand = "version"
	BBLDevVersion  = "dev"
)

type Version struct {
	logger  logger
	version string
}

func NewVersion(version string, logger logger) Version {
	if version == "" {
		version = BBLDevVersion
	}
	return Version{
		logger:  logger,
		version: fmt.Sprintf("%s (%s/%s)", version, runtime.GOOS, runtime.GOARCH),
	}
}

func (v Version) Execute(subcommandFlags []string, state storage.State) error {
	v.logger.Printf("bbl %s\n", v.version)
	return nil
}
