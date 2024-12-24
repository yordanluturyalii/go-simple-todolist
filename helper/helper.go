package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfNil(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfNil(errorCommit)
	}
}
