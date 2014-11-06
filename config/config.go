package config

import (
	"path"
)

type ServerConfig struct {
	Name         string
	DiscoveryURL string
	DataDir      string
	SnapCount    uint64
}

func (c *ServerConfig) WALDir() string { return path.Join(c.DataDir, "wal") }

func (c *ServerConfig) SnapDir() string { return path.Join(c.DataDir, "snap") }

func (c *ServerConfig) ShouldDiscover() bool {
	return c.DiscoveryURL != ""
}
