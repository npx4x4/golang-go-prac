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

	//Solverを初期化
	s := ctx.NewSolver()
	defer s.Close()

	z3Members := make{map[string]}
	// 各メンバーを変数として登録
	for _, Member := range Members:
		Members[]
}