package pegnet

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
