// Реализовать паттерн «адаптер» на любом примере.

package main

// нужный нам итерфэйс
type i interface {
	First()
}

// структура которую нельзя изменять,
// она реализует нужный интерфейс,
// но логика ее методав нам не подходит
type firsType struct{}

func (t *firsType) First() {}

func (t *firsType) Second() {}

func (t *firsType) Three() {}

// структура не реализующая нужный интерфейс
// но с методами чья логика нам подходят
type secondType struct{}

func (t *secondType) Four() {}

// структура реализующая нужный интерфейс
// с использованием нужной нам структуры
type Adapter struct {
	S *secondType
}

// реализация интерфейса
func (t *Adapter) First() {
	t.S.Four()
}

func main() {

}
