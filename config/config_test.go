package config

import "testing"

func TestSnapDir(t *testing.T) {
	tests := map[string]string{
		"/":            "/snap",
		"/var/lib/etc": "/var/lib/etc/snap",
	}
	for dd, w := range tests {
		cfg := ServerConfig{
			DataDir: dd,
		}
		if g := cfg.SnapDir(); g != w {
			t.Errorf("DataDir=%q: SnapDir()=%q, want=%q", dd, g, w)
		}
	}
}

func TestWALDir(t *testing.T) {
	tests := map[string]string{
		"/":            "/wal",
		"/var/lib/etc": "/var/lib/etc/wal",
	}
	for dd, w := range tests {
		cfg := ServerConfig{
			DataDir: dd,
		}
		if g := cfg.WALDir(); g != w {
			t.Errorf("DataDir=%q: WALDir()=%q, want=%q", dd, g, w)
		}
	}
}

func TestShouldDiscover(t *testing.T) {
	tests := map[string]bool{
		"":    false,
		"foo": true,
		"http://discovery.etcd.io/asdf": true,
	}
	for durl, w := range tests {
		cfg := ServerConfig{
			DiscoveryURL: durl,
		}
		if g := cfg.ShouldDiscover(); g != w {
			t.Errorf("durl=%q: ShouldDiscover()=%t, want=%t", durl, g, w)
		}
	}
}
