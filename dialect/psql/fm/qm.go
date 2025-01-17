package fm

import (
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/clause"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
)

func Distinct() bob.Mod[*dialect.Function] {
	return bob.ModFunc[*dialect.Function](func(f *dialect.Function) {
		f.Distinct = true
	})
}

func OrderBy(e any) dialect.OrderBy[*dialect.Function] {
	return dialect.OrderBy[*dialect.Function](func() clause.OrderDef {
		return clause.OrderDef{
			Expression: e,
		}
	})
}

func WithinGroup() bob.Mod[*dialect.Function] {
	return bob.ModFunc[*dialect.Function](func(f *dialect.Function) {
		f.WithinGroup = true
	})
}

func Filter(e ...any) bob.Mod[*dialect.Function] {
	return bob.ModFunc[*dialect.Function](func(f *dialect.Function) {
		f.Filter = append(f.Filter, e...)
	})
}

func Over() dialect.WindowMod[*dialect.Function] {
	m := dialect.WindowMod[*dialect.Function]{}
	m.WindowChain = &dialect.WindowChain[*dialect.WindowMod[*dialect.Function]]{
		Wrap: &m,
	}
	return m
}

func As(alias string) bob.Mod[*dialect.Function] {
	return bob.ModFunc[*dialect.Function](func(f *dialect.Function) {
		f.Alias = alias
	})
}

func Columns(name, datatype string) bob.Mod[*dialect.Function] {
	return bob.ModFunc[*dialect.Function](func(f *dialect.Function) {
		f.AppendColumn(name, datatype)
	})
}
