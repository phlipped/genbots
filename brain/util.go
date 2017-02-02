package brain

func NewRandomExpression() Expression {
	return &Constant{}
}

func NewRandomExpression2() Expression {

	i := If{
		&GetArg{0},
		&GetArg{1},
		&GetArg{2},
	}

	fc := FuncCaller{
		f: &i,
		args: []Expression{
			&Constant{0},
			&Constant{3},
			&Constant{10},
		},
	}
	return &fc
}
