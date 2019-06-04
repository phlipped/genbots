package brain

type Env interface {
	Read(addr uint32) int32
}
