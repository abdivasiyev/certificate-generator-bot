package viper

import "time"

func (c *Viper) Get(key string) interface{} {
	return c.Viper.Get(key)
}

func (c *Viper) GetBool(key string) bool {
	return c.Viper.GetBool(key)
}

func (c *Viper) GetFloat64(key string) float64 {
	return c.Viper.GetFloat64(key)
}

func (c *Viper) GetInt(key string) int {
	return c.Viper.GetInt(key)
}

func (c *Viper) GetUInt32(key string) uint32 {
	return c.Viper.GetUint32(key)
}

func (c *Viper) GetUInt8(key string) uint8 {
	return uint8(c.Viper.GetUint(key))
}

func (c *Viper) GetIntSlice(key string) []int {
	return c.Viper.GetIntSlice(key)
}

func (c *Viper) GetString(key string) string {
	return c.Viper.GetString(key)
}

func (c *Viper) GetStringSlice(key string) []string {
	return c.Viper.GetStringSlice(key)
}

func (c *Viper) GetStringMap(key string) map[string]interface{} {
	return c.Viper.GetStringMap(key)
}
func (c *Viper) GetStringMapString(key string) map[string]string {
	return c.Viper.GetStringMapString(key)
}

func (c *Viper) UnmarshalKey(key string, val interface{}) error {
	return c.Viper.UnmarshalKey(key, &val)
}

func (c *Viper) GetDuration(key string) time.Duration {
	return c.Viper.GetDuration(key)
}
