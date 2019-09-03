// Author :		Eric<eehsiao@gmail.com>

package lib

// Iif logical condition
func Iif(l bool, a interface{}, b interface{}) interface{} {
	if l {
		return a
	}
	return b
}
