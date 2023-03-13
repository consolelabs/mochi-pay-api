package apierror

var (
	ErrFromWalletRequired  = New("from wallet is required", 400, "Transfer_001")
	ErrToWalletRequired    = New("tos wallet is required", 400, "Transfer_002")
	ErrAmountMismatch      = New("tos wallet and amounts are not same length", 400, "Transfer_003")
	ErrTokenNotSupport     = New("token not supported", 400, "Transfer_004")
	ErrInsufficientBalance = New("insufficient balance", 400, "Transfer_005")
)

var (
	Code400 = "API_400"
	Code404 = "API_404"
	Code500 = "API_500"
)
