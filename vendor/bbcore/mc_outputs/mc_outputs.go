package mc_outputs

import(
 .	"bbcore/types"

	"bbcore/db"
)

type(
	DBConnT		= db.DBConnT
	refDBConnT	= *DBConnT

	CalcEarningsCbT struct{
		IfError		func (err string)
	}
)

func CalcEarnings(conn refDBConnT, type_ string, from_main_chain_index MCIndexT, to_main_chain_index MCIndexT,
			address AddressT, callbacks CalcEarningsCbT) {
	panic("[tbd] CalcEarnings")
}
