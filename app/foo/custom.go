package foo

type Custom struct {
	Key string `json:"key"`
}

func (c *Custom) Clone() *Custom {
	return &Custom{Key: c.Key}
}
