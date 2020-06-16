package pegnet

import "time"

// This file can be used for node administration related functions

var (
	// PegnetdSyncVersion is an indicator of the version of pegnetd
	// at each height synced. This version number can differ from the tagged
	// version, and is likely only to be updated at hard forks. It is used to
	// detect if a pegnetd was updated late, and therefore has an invalid state.
	//
	// Each fork should increment this number by at least 1
	PegnetdSyncVersion = 1
)

// createTableSyncVersion is a SQL string that creates the
// "pn_sync_version" table. This table tracks which heights are synced
// with what version of pegnetd. This will allow pegnetd to detect if
// it was updated before or after a hardfork and respond appropriately.
const createTableSyncVersion = `CREATE TABLE IF NOT EXISTS "pn_sync_version" (
        "height"    		INTEGER NOT NULL,
        "version"       	INTEGER NOT NULL,
        "unix_timestamp"	INTEGER NOT NULL,

        PRIMARY KEY("height")
);
`

func (p Pegnet) MarkHeightSynced(tx QueryAble, height uint32) error {
	return p.markHeightSyncedVersion(tx, height, PegnetdSyncVersion)
}

func (Pegnet) markHeightSyncedVersion(tx QueryAble, height uint32, version int) error {
	stmtStringFmt := `INSERT INTO "pn_sync_version" 
			("height", "version", "unix_timestamp")
			VALUES (?, ?, ?);`

	stmt, err := tx.Prepare(stmtStringFmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(height, version, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}
