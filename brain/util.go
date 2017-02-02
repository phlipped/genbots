package brain

func NewRandomExpression() Expression {
	return &Constant{}
}

func NewRandomExpression2() Expression {

	MulSeven := &FuncCaller{
		f: Mul,
		args: []Expression{
			&Constant{7},
			GetArg{0},
		},
	}

	TwentyOne := &FuncCaller{
		f: MulSeven,
		args: []Expression{
			&Constant{3},
		},
	}

	return TwentyOne
}

func Fibs() Expression {

	Sub1 := &FuncCaller{
		f: Sub,
		args: []Expression {
			GetArg{0},
			&Constant{1},
		},
	}

	Fib := &FuncCaller{
		f: If,
	}

	Fib.args = []Expression{
		GetArg{0},
		&FuncCaller{
			f: Mul,
			args: []Expression{
				GetArg{0},
				&FuncCaller{
					f: Fib, 
						args: []Expression{
							&FuncCaller{
								f: Sub1,
								args: []Expression{
									GetArg{0},
								},
							},
						},
					},
				},
			},
		&Constant{1},
	}

	return &FuncCaller{
		f: Fib,
		args: []Expression{
			&Constant{3},
		},
	}
}
