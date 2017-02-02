package brain

func NewRandomExpression() Expression {
	return &Constant{}
}

func NewRandomExpression2() Expression {

	i := If

	fc := FuncCaller{
		f: &i,
		args: []Expression{
			&Constant{1},
			&Constant{3},
			&Constant{10},
		},
	}
	return &fc
}
