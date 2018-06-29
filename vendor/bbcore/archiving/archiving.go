package archiving

import(
 .	"bbcore/types"

	"bbcore/db"
)

type(
	DBConnT		= db.DBConnT
	refDBConnT	= *DBConnT

	refJointT	= *JointT

	refAsyncFunctorsT = *AsyncFunctorsT
)

//  (*db.DBConnT, *types.JointT, string, *[]func() error)

func GenerateQueriesToArchiveJoint_sync(conn refDBConnT, objJoint refJointT, s string, arrQueries refAsyncFunctorsT) {
	panic("[tbd] archiving.GenerateQueriesToArchiveJoint")
}
