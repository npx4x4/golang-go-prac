package solver

import (
	"../model"
	"github.com/mitchellh/go-z3"
)
_ Uwasa = None

func SolveRanking(Members[],  ) {
	// contextを初期化
	config := z3.NewConfig()
	ctx := z3.NewContext(config)
	config.Close()
	defer ctx.Close()

	// Solverを初期化
	s := ctx.NewSolver()
	defer s.Close()

	z3Members := make{map[string]z3.Int}
	// 各メンバーを変数として登録
	for _, Member := range Members {
		z3Members[Member] = ctx.Const(ctx.Symbol(Member), ctx.IntSort())
		s.Assert(z3Members[Member].Gt(zero))
		s.Assert(z3Members(Member).Lt(RankingSize+1))
	}

	for _, Uwasa := range UwasaBox {

	}

	if v := s.Check(); v != True {
		fmt.println("いけてます")
		return
	}
}
