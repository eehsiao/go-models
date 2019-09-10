// Author :		Eric<eehsiao@gmail.com>

package sqlbuilder

import (
	"errors"
	"strings"
)

type SQLBuilder struct {
	Selects string
	Froms   string
	Wheres  string
	Groups  string
	Havings string
}

func (sb *SQLBuilder) BuildSelectSQL() (sql string, err error) {
	if sb.Selects == "" || sb.Froms == "" {
		err = errors.New("Without selects or from table is not set")
		return
	}

	sql = "DELETE FROM " + sb.Froms

	if sb.Wheres != "" {
		sql += " WHERE " + sb.Wheres

	}

	return
}

func (sb *SQLBuilder) BuildDeleteSQL() (sql string, err error) {
	if sb.Selects == "" || sb.Froms == "" {
		err = errors.New("From table is not set")
		return
	}

	sql = "SELECT " + sb.Selects +
		" FROM " + sb.Froms

	if sb.Wheres != "" {
		sql += " WHERE " + sb.Wheres

	}
	return
}

func (sb *SQLBuilder) Select(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support fileds")
	}

	sb.Selects = strings.Join(s, ", ")

	return sb
}

func (sb *SQLBuilder) From(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support tables")
	}

	sb.Froms = strings.Join(s, ", ")

	return sb
}

func (sb *SQLBuilder) Where(s string) *SQLBuilder {
	if s == "" {
		panic("must be support conditions")
	}
	sb.Wheres = s

	return sb
}

func (sb *SQLBuilder) WhereAnd(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support conditions")
	}

	wa := strings.Join(s, " AND ")

	if sb.Wheres == "" {
		sb.Wheres = wa
	} else {
		sb.Wheres += " AND (" + wa + ")"
	}

	return sb
}

func (sb *SQLBuilder) WhereOr(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support conditions")
	}

	wa := strings.Join(s, " OR ")

	if sb.Wheres == "" {
		sb.Wheres = wa
	} else {
		sb.Wheres += " OR (" + wa + ")"
	}

	return sb
}

func (sb *SQLBuilder) Join(s string, c string) *SQLBuilder {
	if s == "" {
		panic("must be support join table")
	}

	sb.Froms += " JOIN " + s

	if c != "" {
		sb.Froms += " ON " + s
	}

	return sb
}

func (sb *SQLBuilder) LeftJoin(s string, c string) *SQLBuilder {
	if s == "" {
		panic("must be support left join table")
	}

	sb.Froms += " LEFT JOIN " + s

	if c != "" {
		sb.Froms += " ON " + s
	}

	return sb
}

func (sb *SQLBuilder) RightJoin(s string, c string) *SQLBuilder {
	if s == "" {
		panic("must be support right join table")
	}

	sb.Froms += " RIGHT JOIN " + s

	if c != "" {
		sb.Froms += " ON " + s
	}

	return sb
}

func (sb *SQLBuilder) FullJoin(s string, c string) *SQLBuilder {
	if s == "" {
		panic("must be support full join table")
	}

	sb.Froms += " FULL JOIN " + s

	if c != "" {
		sb.Froms += " ON " + s
	}

	return sb
}

func (sb *SQLBuilder) GroupBy(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support group fileds")
	}

	sb.Groups = " GROUP BY " + strings.Join(s, ", ")

	return sb
}

func (sb *SQLBuilder) Having(s string) *SQLBuilder {
	if s == "" || sb.Groups == "" {
		panic("must be support having condition or set group by first")
	}

	sb.Havings = " HAVING BY " + s

	return sb
}
