package pegnet

const createTableAddresses = `CREATE TABLE IF NOT EXISTS "pn_addresses" (
        "id"            INTEGER PRIMARY KEY,
        "address"       BLOB NOT NULL UNIQUE,
        "peg_balance"   INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("peg_balance" >= 0),
        "pusd_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pusd_balance" >= 0),
        "peur_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("peur_balance" >= 0),
        "pjpy_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pjpy_balance" >= 0),
        "pgbp_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pgbp_balance" >= 0),
        "pcad_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pcad_balance" >= 0),
        "pchf_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pchf_balance" >= 0),
        "pinr_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pinr_balance" >= 0),
        "psgd_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("psgd_balance" >= 0),
        "pcny_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pcny_balance" >= 0),
        "phkd_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("phkd_balance" >= 0),
        "pkrw_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pkrw_balance" >= 0),
        "pbrl_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pbrl_balance" >= 0),
        "pphp_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pphp_balance" >= 0),
        "pmxn_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pmxn_balance" >= 0),
        "pxau_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pxau_balance" >= 0),
        "pxag_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pxag_balance" >= 0),
        "pxbt_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pxbt_balance" >= 0),
        "peth_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("peth_balance" >= 0),
        "pltc_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pltc_balance" >= 0),
        "prvn_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("prvn_balance" >= 0),
        "pxbc_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pxbc_balance" >= 0),
        "pfct_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pfct_balance" >= 0),
        "pbnb_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pbnb_balance" >= 0),
        "pxlm_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pxlm_balance" >= 0),
        "pada_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pada_balance" >= 0),
        "pxmr_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pxmr_balance" >= 0),
        "pdash_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pdash_balance" >= 0),
        "pzec_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pzec_balance" >= 0),
        "pdcr_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pdcr_balance" >= 0),
        -- v4 additions
        "paud_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("paud_balance" >= 0),
        "pnzd_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pnzd_balance" >= 0),
        "psek_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("psek_balance" >= 0),
        "pnok_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pnok_balance" >= 0),
        "prub_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("prub_balance" >= 0),
        "pzar_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pzar_balance" >= 0),
        "ptry_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("ptry_balance" >= 0),
        "peos_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("peos_balance" >= 0),
        "plink_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("plink_balance" >= 0),
        "patom_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("patom_balance" >= 0),
        "pbat_balance"  INTEGER NOT NULL DEFAULT 0
                        CONSTRAINT "insufficient balance" CHECK ("pbat_balance" >= 0),
        "pxtz_balance"  INTEGER NOT NULL DEFAULT 0 
                        CONSTRAINT "insufficient balance" CHECK ("pxtz_balance" >= 0)
);
CREATE INDEX IF NOT EXISTS "idx_address_balances_address_id" ON "pn_addresses"("address");
`
