package redis

type RedisConfig struct {
	Db_zero        uint64
	Db_uuid        uint64
	Db_ap          uint64
	Db_start       uint64
	Key_db_next    string
	Key_db_mapping string
	Hostname_port  string
}

func (c *RedisConfig) db_zero() uint64 { return 0 }

func (c *RedisConfig) db_uuid() uint64 { return 10 }

func (c *RedisConfig) db_ap() uint64 { return 11 }

func (c *RedisConfig) db_start() uint64 { return 100 }

func (c *RedisConfig) key_db_next() string { return "nextdb" }

func (c *RedisConfig) key_db_mapping() string { return "hm:accountid:db" }

func (c *RedisConfig) hostname_port() string { return ":6379" }
