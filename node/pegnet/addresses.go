package pegnet

import (
	"database/sql"
	"fmt"
	"github.com/pegnet/pegnetd/fat/fat2"
	"strings"
)

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

func (p *Pegnet) SelectIssuances() (map[fat2.PTicker]uint64, error) {
	issuanceMap := make(map[fat2.PTicker]uint64, int(fat2.PTickerMax))
	for i := fat2.PTickerInvalid + 1; i < fat2.PTickerMax; i++ {
		issuanceMap[i] = 0
	}
	// Can't make pointers of map elements, so a temporary array must be used
	issuances := make([]uint64, int(fat2.PTickerMax))
	queryFmt := `SELECT %v FROM pn_addresses`
	var sb strings.Builder
	for i := fat2.PTickerInvalid + 1; i < fat2.PTickerMax-1; i++ {
		tickerLower := strings.ToLower(i.String())
		sb.WriteString(fmt.Sprintf("IFNULL(SUM(%s_balance), 0), ", tickerLower))
	}
	tickerLower := strings.ToLower((fat2.PTickerMax - 1).String())
	sb.WriteString(fmt.Sprintf("IFNULL(SUM(%s_balance), 0) ", tickerLower))
	err := p.DB.QueryRow(fmt.Sprintf(queryFmt, sb.String())).Scan(
		&issuances[fat2.PTickerPEG],
		&issuances[fat2.PTickerUSD],
		&issuances[fat2.PTickerEUR],
		&issuances[fat2.PTickerJPY],
		&issuances[fat2.PTickerGBP],
		&issuances[fat2.PTickerCAD],
		&issuances[fat2.PTickerCHF],
		&issuances[fat2.PTickerINR],
		&issuances[fat2.PTickerSGD],
		&issuances[fat2.PTickerCNY],
		&issuances[fat2.PTickerHKD],
		&issuances[fat2.PTickerKRW],
		&issuances[fat2.PTickerBRL],
		&issuances[fat2.PTickerPHP],
		&issuances[fat2.PTickerMXN],
		&issuances[fat2.PTickerXAU],
		&issuances[fat2.PTickerXAG],
		&issuances[fat2.PTickerXBT],
		&issuances[fat2.PTickerETH],
		&issuances[fat2.PTickerLTC],
		&issuances[fat2.PTickerRVN],
		&issuances[fat2.PTickerXBC],
		&issuances[fat2.PTickerFCT],
		&issuances[fat2.PTickerBNB],
		&issuances[fat2.PTickerXLM],
		&issuances[fat2.PTickerADA],
		&issuances[fat2.PTickerXMR],
		&issuances[fat2.PTickerDASH],
		&issuances[fat2.PTickerZEC],
		&issuances[fat2.PTickerDCR],
		// V4 Additions
		&issuances[fat2.PTickerAUD],
		&issuances[fat2.PTickerNZD],
		&issuances[fat2.PTickerSEK],
		&issuances[fat2.PTickerNOK],
		&issuances[fat2.PTickerRUB],
		&issuances[fat2.PTickerZAR],
		&issuances[fat2.PTickerTRY],
		&issuances[fat2.PTickerEOS],
		&issuances[fat2.PTickerLINK],
		&issuances[fat2.PTickerATOM],
		&issuances[fat2.PTickerBAT],
		&issuances[fat2.PTickerXTZ],
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return issuanceMap, nil
		}
		return nil, err
	}
	for i := fat2.PTickerInvalid + 1; i < fat2.PTickerMax; i++ {
		issuanceMap[i] = issuances[i]
	}
	return issuanceMap, nil
}
