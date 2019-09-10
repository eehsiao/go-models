// Author :		Eric<eehsiao@gmail.com>

package sqlbuilder

type UpdateBuilder struct {
	Sets   map[string]interface{}
	Froms  string
	Wheres string
}
