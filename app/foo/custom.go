package foo

type Custom struct {
	Key string `json:"key"`
}

func (c *Custom) Clone() *Custom {
	return &Custom{Key: c.Key}
}

func (c *Custom) String() string {
	return c.Key
}
