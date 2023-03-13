package view

type success struct {
	Message string `json:"message"`
}

func ToRespSuccess() *success {
	return &success{
		Message: "Transfer success",
	}
}
