package raindrops

import "strconv"

func Convert(number int) string {
    r := map[int]string{3:"Pling", 5:"Plang", 7:"Plong"}
    raindrops := ""
	for _,value := range []int{3,5,7} {
        if number%value == 0 {
            raindrops += r[value]
        }
    }
	if raindrops == "" {
        return strconv.Itoa(number)
    }
	return raindrops
}
