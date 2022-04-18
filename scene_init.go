package rayframe

type InitScene interface {
	Init(*RayFrame) interface{}
}
