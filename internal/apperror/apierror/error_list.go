package apierror

var (
	ErrFromWalletRequired = New("from wallet is required", 400, "Transfer_001")
	ErrToWalletRequired   = New("tos wallet is required", 400, "Transfer_002")
	ErrAmountMismatch     = New("tos wallet and amounts are not same length", 400, "Transfer_003")
)
