package action

import "github.com/wb/cmd/1L/task1/1/human"

// если Human  инкапсулирован в Action
type Action struct {
	h human.Human
}

func (a *Action) Walk()

// обарачиваем методы human.Human в методы Action
func (a *Action) Age() {
	a.h.Age()
}
