package character

import "tde/internal/utilities"

type Character struct{}

func (c *Character) Develop(code string) string {
	index := c.PickCutPoint(code)
	fragment := string(ProduceRandomFragment())
	seperator := " "
	return code[:index] + seperator + fragment + seperator + code[index+1:]
}

func (c *Character) PickCutPoint(code string) int {
	// 1 for excluding last item, 1 is for staying in boundary
	return utilities.URandIntN(len(code) - 2)
}
