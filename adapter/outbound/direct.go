package outbound

type Direct struct{}

func (x *Direct) Name() string { return "direct" }
