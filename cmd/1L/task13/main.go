// Реализовать пересечение двух неупорядоченных множеств.

package main

type space struct{}

func main() {

}

func set(list []string) map[string]space {
	m := make(map[string]space)
	for _, n := range list {
		m[n] = space{}
	}
	return m
}
