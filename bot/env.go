package bot

type Env struct {
	bot *Bot
}

func (e *Env) Read(addr uint32) int32 {
	return 0
}
