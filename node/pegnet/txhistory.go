package pegnet

// in the context of tables, `history_txbatch` is the table that holds the unique reference hash
// and `transaction` is the table that holds the actions associated with that unique reference hash
// `lookup` is an outside reference that indexes the addresses involved in the actions
//
// associations are:
// 	* history_txbach : history_transaction is `1:n`
// 	* history_transaction : lookup is `1:n`
//	* lookup : (transaction.outputs + transaction.inputs) is `1:n` (unique addresses only)
const createTableTxHistoryBatch = `CREATE TABLE IF NOT EXISTS "pn_history_txbatch" (
	"history_id"	INTEGER PRIMARY KEY,
	"entry_hash"    BLOB NOT NULL,
	"height"        INTEGER NOT NULL, -- height the tx is in
	"blockorder"	INTEGER NOT NULL,
	"timestamp"		INTEGER NOT NULL,
	"executed"		INTEGER NOT NULL, -- -1 if failed, 0 if pending, height it was applied at otherwise

	UNIQUE("entry_hash", "height")
);
CREATE INDEX IF NOT EXISTS "idx_history_txbatch_entry_hash" ON "pn_history_txbatch"("entry_hash");
CREATE INDEX IF NOT EXISTS "idx_history_txbatch_timestamp" ON "pn_history_txbatch"("timestamp");
CREATE INDEX IF NOT EXISTS "idx_history_txbatch_height" ON "pn_history_txbatch"("height");
`
const createTableTxHistoryTx = `CREATE TABLE IF NOT EXISTS "pn_history_transaction" (
	"entry_hash"	BLOB NOT NULL,
	"tx_index"		INTEGER NOT NULL,	-- the batch index
	"action_type"	INTEGER NOT NULL,
	"from_address"  BLOB NOT NULL,
	"from_asset"	STRING NOT NULL,
	"from_amount"	INTEGER NOT NULL,
	"to_asset"		STRING NOT NULL,	-- used for NOT transfers
	"to_amount"		INTEGER NOT NULL,	-- used for NOT transfers
	"outputs"		BLOB NOT NULL,		-- used for transfers only

	PRIMARY KEY("entry_hash", "tx_index"),
	FOREIGN KEY("entry_hash") REFERENCES "pn_history_txbatch"
);
CREATE INDEX IF NOT EXISTS "idx_history_transaction_entry_hash" ON "pn_history_transaction"("entry_hash");
CREATE INDEX IF NOT EXISTS "idx_history_transaction_tx_index" ON "pn_history_transaction"("tx_index");
`

const createTableTxHistoryLookup = `CREATE TABLE IF NOT EXISTS "pn_history_lookup" (
	"entry_hash"	BLOB NOT NULL,
	"tx_index"		INTEGER NOT NULL,
	"address"		BLOB NOT NULL,

	PRIMARY KEY("entry_hash", "tx_index", "address"),
	FOREIGN KEY("entry_hash", "tx_index") REFERENCES "pn_history_transaction"
);
CREATE INDEX IF NOT EXISTS "idx_history_lookup_address" ON "pn_history_lookup"("address");
CREATE INDEX IF NOT EXISTS "idx_history_lookup_entry_index" ON "pn_history_lookup"("entry_hash", "tx_index");`
