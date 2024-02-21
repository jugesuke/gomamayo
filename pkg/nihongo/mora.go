package nihongo

import "regexp"

type Mora []string

func DivideMora(s string) []string {
	return regexp.MustCompile("([ウクスツヌフムユルグズヅブプヴ][ァィェォ]|[イキシチニヒミリギジヂビピ][ャュェョ]|[テデ][ィュ]|[ァ-ヴー])").
		FindAllString(s, -1)
}
