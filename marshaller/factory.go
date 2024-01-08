package marshaller

type Factory struct {
	Encoder    func(query any, cur, field *Field) error
	SqlFactory func(field *Field) any
}

var factories = map[string]*Factory{}

func reg(est string, encoder func(query any, cur, field *Field) error, f func(field *Field) any) {
	factories[est] = &Factory{
		Encoder:    encoder,
		SqlFactory: f,
	}
}

func getFactory(est string) (*Factory, bool) {
	f, exist := factories[est]
	return f, exist
}

func (ins *Factory) Create(cur *Field) (func(field *Field) error, any) {
	sql := ins.SqlFactory(cur)
	return func(field *Field) error {
		return ins.Encoder(sql, cur, field)
	}, sql
}
