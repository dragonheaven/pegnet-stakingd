package pegnet

const createTableTransactionBatchHolding = `CREATE TABLE IF NOT EXISTS "pn_transaction_batch_holding" (
        "id"            	INTEGER PRIMARY KEY,
        "entry_hash"    	BLOB NOT NULL UNIQUE,
        "entry_data"    	BLOB NOT NULL,
        "height"        	INTEGER NOT NULL,
        "eblock_keymr"  	BLOB NOT NULL,
        "unix_timestamp"	INTEGER NOT NULL
);
`
