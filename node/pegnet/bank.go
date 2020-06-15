package pegnet

const createTableBank = `CREATE TABLE IF NOT EXISTS "pn_bank" (
	"height" INTEGER PRIMARY KEY,
	"bank_amount" INTEGER NOT NULL DEFAULT 0,
	"bank_used" INTEGER NOT NULL DEFAULT 0,
	"total_requested" INTEGER NOT NULL DEFAULT 0,
	
	UNIQUE("height")
);
`
