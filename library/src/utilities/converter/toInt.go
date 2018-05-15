package converter

import "strconv"

/*
	Tar inn string (består kun av positivt tall) og returnerer uint64
*/
func StringToInt_plus1(s string) string {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	s = strconv.FormatInt(i+1, 10)
	return s
}
