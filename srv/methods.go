package srv

func (s *APIServer) jrpcMethods() jrpc.MethodMap {
	return jrpc.MethodMap{
		"get-rich-list":          s.getRichList,
		"get-global-rich-list":   s.getGlobalRichList,
		"get-miner-distribution": s.getMiningDominance,
		"get-bank":               s.getBank,
		"get-transactions":       s.getTransactions(false),
		"get-transaction-status": s.getTransactionStatus,
		"get-transaction":        s.getTransactions(true),
		"get-pegnet-balances":    s.getPegnetBalances,
		"get-pegnet-issuance":    s.getPegnetIssuance,
		"send-transaction":       s.sendTransaction,

		"get-sync-status": s.getSyncStatus,
		"properties":      s.properties,

		"get-pegnet-rates": s.getPegnetRates,
	}

}
