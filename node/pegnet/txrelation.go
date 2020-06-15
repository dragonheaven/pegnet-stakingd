package pegnet

// createTableTransactions is a SQL string that creates the
// "pn_address_transactions" table.
//
// The "pn_address_transactions" table has a foreign key reference to the
// "pn_addresses" table, which must exist first.
//
// If the transaction is a conversion, both "to" and "conversion" are set to true
const createTableTransactions = `CREATE TABLE IF NOT EXISTS "pn_address_transactions" (
        "entry_hash"    BLOB NOT NULL,
        "address"       BLOB NOT NULL,
        "tx_index"      INTEGER NOT NULL,
        "to"            BOOL NOT NULL,
        "conversion"    BOOL NOT NULL,

        PRIMARY KEY("entry_hash", "address"),

        FOREIGN KEY("address") REFERENCES "pn_addresses"
);
CREATE INDEX IF NOT EXISTS "idx_address_transactions_address" ON "pn_address_transactions"("address");
`
