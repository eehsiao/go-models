// Author :		Eric<eehsiao@gmail.com>

package sqlbuilder

import (
	"strconv"
	"strings"

	"github.com/eehsiao/go-models/lib"
)

type SQLBuilder struct {
	driver_type string

	// for select , delete
	distinct bool
	selects  []string
	froms    []string
	joins    []string
	wheres   []string
	orders   []string
	groups   []string
	havings  string
	limit    string
	top      string

	// TODO : for insert
	into   string
	fields []string
	values []string

	// TODO : for update
	sets map[string]interface{}
}

func NewSQLBuilder() (b *SQLBuilder) {
	b = &SQLBuilder{
		driver_type: "mysql",
	}
	b.ClearBuilder()

	return
}

func (sb *SQLBuilder) SetDriverType(t string) {
	switch t {
	case "mysql":
		fallthrough
	case "mssql":
		fallthrough
	case "oracle":
		sb.driver_type = t
	}
}

func (sb *SQLBuilder) ClearBuilder() {
	sb.distinct = false
	sb.selects = make([]string, 0)
	sb.froms = make([]string, 0)
	sb.joins = make([]string, 0)
	sb.wheres = make([]string, 0)
	sb.orders = make([]string, 0)
	sb.groups = make([]string, 0)
	sb.havings = ""
	sb.limit = ""
	sb.top = ""
	sb.into = ""
	sb.fields = make([]string, 0)
	sb.values = make([]string, 0)
	sb.sets = make(map[string]interface{}, 0)
}

func (sb *SQLBuilder) BuildDeleteSQL() (sql string) {
	if len(sb.froms) != 1 {
		panic("must be have only one from table")
	}

	sql = "DELETE FROM " + sb.froms[0]
	if len(sb.wheres) > 0 {
		sql += " WHERE " + strings.Join(sb.wheres, " ")
	}
	return
}

func (sb *SQLBuilder) BuildSelectSQL() (sql string) {
	if len(sb.selects) == 0 || len(sb.froms) == 0 {
		panic("Without selects or from table is not set")
	}

	sql = "SELECT " + lib.Iif(sb.distinct, "DISTINCT ", "").(string)

	if sb.driver_type == "mssql" && sb.top != "" {
		sql += "TOP " + sb.top + " "
	}

	sql += strings.Join(sb.selects, ",")
	sql += " FROM " + strings.Join(sb.froms, ",")

	if len(sb.joins) > 0 {
		sql += " " + strings.Join(sb.joins, " ")
	}

	if len(sb.wheres) > 0 {
		sql += " WHERE " + strings.Join(sb.wheres, " ")
	}

	if len(sb.orders) > 0 {
		sql += " ORDER BY " + strings.Join(sb.orders, ",")
	}

	if len(sb.groups) > 0 {
		sql += " GROUP BY " + strings.Join(sb.groups, ",")
	}

	if sb.havings != "" {
		sql += " HAVING BY " + sb.havings
	}

	if sb.driver_type == "mysql" && sb.limit != "" {
		sql += " LIMIT " + sb.limit
	}

	return
}

func (sb *SQLBuilder) CanBuildSelect() bool {
	return len(sb.selects) > 0 && len(sb.froms) > 0
}

func (sb *SQLBuilder) CanBuildDelete() bool {
	return len(sb.froms) > 0
}

func (sb *SQLBuilder) CanBuildUpdate() bool {
	return len(sb.sets) > 0
}

func (sb *SQLBuilder) CanBuildInsert() bool {
	return len(sb.fields) > 0 && len(sb.values) > 0 && sb.into != ""
}

func (sb *SQLBuilder) Distinct(b bool) *SQLBuilder {
	sb.distinct = b

	return sb
}

func (sb *SQLBuilder) Limit(i ...int) *SQLBuilder {
	if sb.driver_type != "mysql" {
		panic("limit only support mysql")
	}
	if len(i) == 0 {
		panic("must have value for limit")
	}

	if len(i) == 1 {
		sb.limit = strconv.Itoa(i[0])
	} else if len(i) > 1 {
		sb.limit = strconv.Itoa(i[0]) + "," + strconv.Itoa(i[1])
	}

	return sb
}

func (sb *SQLBuilder) Top(i int) *SQLBuilder {
	if sb.driver_type != "mysql" {
		panic("limit only support mssql")
	}
	if i <= 0 {
		panic("must have >=1 value for top")
	}

	sb.limit = strconv.Itoa(i)

	return sb
}

func (sb *SQLBuilder) Select(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support fileds")
	}

	for _, v := range s {
		sb.selects = append(sb.selects, v)
	}

	return sb
}

func (sb *SQLBuilder) From(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support tables")
	}

	for _, v := range s {
		sb.froms = append(sb.froms, v)
	}

	return sb
}

func (sb *SQLBuilder) Where(s string) *SQLBuilder {
	if s == "" {
		panic("must be support conditions")
	}
	if len(sb.wheres) == 0 {
		sb.wheres = append(sb.wheres, s)
	} else {
		s = "AND " + s
		sb.wheres = append(sb.wheres, s)
	}

	return sb
}

func (sb *SQLBuilder) WhereAnd(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support conditions")
	}

	for _, v := range s {
		if len(sb.wheres) == 0 {
			sb.wheres = append(sb.wheres, v)
		} else {
			v = "AND " + v
			sb.wheres = append(sb.wheres, v)
		}
	}

	return sb
}

func (sb *SQLBuilder) WhereOr(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support conditions")
	}

	for _, v := range s {
		if len(sb.wheres) == 0 {
			sb.wheres = append(sb.wheres, v)
		} else {
			v = "OR " + v
			sb.wheres = append(sb.wheres, v)
		}
	}

	return sb
}

func (sb *SQLBuilder) Join(s string, c string) *SQLBuilder {
	if s == "" {
		panic("must be support join table")
	}

	if c != "" {
		s += " ON " + c
	}
	sb.joins = append(sb.joins, "JOIN "+s)

	return sb
}

func (sb *SQLBuilder) InnerJoin(s string, c string) *SQLBuilder {
	if s == "" {
		panic("must be support join table")
	}

	if c != "" {
		s += " ON " + c
	}
	sb.joins = append(sb.joins, "INNER JOIN "+s)

	return sb
}

func (sb *SQLBuilder) LeftJoin(s string, c string) *SQLBuilder {
	if s == "" {
		panic("must be support join table")
	}

	if c != "" {
		s += " ON " + c
	}
	sb.joins = append(sb.joins, "LEFT JOIN "+s)

	return sb
}

func (sb *SQLBuilder) RightJoin(s string, c string) *SQLBuilder {
	if s == "" {
		panic("must be support join table")
	}

	if c != "" {
		s += " ON " + c
	}
	sb.joins = append(sb.joins, "RIGHT JOIN "+s)

	return sb
}

func (sb *SQLBuilder) FullJoin(s string, c string) *SQLBuilder {
	if s == "" {
		panic("must be support join table")
	}

	if c != "" {
		s += " ON " + c
	}
	sb.joins = append(sb.joins, "FULL OUTER JOIN "+s)

	return sb
}

func (sb *SQLBuilder) GroupBy(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support group fileds")
	}

	for _, v := range s {
		sb.groups = append(sb.groups, v)
	}

	return sb
}

func (sb *SQLBuilder) OrderBy(s ...string) *SQLBuilder {
	return sb.OrderByAsc(s...)
}

func (sb *SQLBuilder) OrderByAsc(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support order fileds")
	}

	sb.orders = append(sb.orders, strings.Join(s, ",")+" ASC")

	return sb
}

func (sb *SQLBuilder) OrderByDesc(s ...string) *SQLBuilder {
	if len(s) == 0 {
		panic("must be support order fileds")
	}

	sb.orders = append(sb.orders, strings.Join(s, ",")+" DESC")

	return sb
}

func (sb *SQLBuilder) Having(s string) *SQLBuilder {
	if s == "" || len(sb.groups) == 0 {
		panic("must be support having condition or set group by first")
	}

	sb.havings = s

	return sb
}
