package character

import "tde/internal/utilities"

type Character struct{}

func (c *Character) Develop(code string) string {
	index := utilities.URandIntN(len(code) - 2) // 1 for excluding last item, 1 is for staying in boundary
	fragment := string(ProduceRandomFragment())
	seperator := " "
	return code[:index] + seperator + fragment + seperator + code[index+1:]
}
