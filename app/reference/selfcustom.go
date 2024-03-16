package reference

type SelfCustom struct {
	Key string `json:"key"`
}

func (c *SelfCustom) Clone() *SelfCustom {
	return &SelfCustom{Key: c.Key}
}

func (c *SelfCustom) String() string {
	return c.Key
}
