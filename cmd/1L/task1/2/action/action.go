package action

type Human struct {
}

func (*Human) Age() int

// второй вариант
// если Human не инкапсулирован в Action
type Action struct {
	Human
}

func (*Action) Walk()
