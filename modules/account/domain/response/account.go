package response

type (
	IDPResponse struct {
		// Identity Provider Response from linkedin, google, microsoft
		ID     string
		Name   string
		Avatar string
	}
)
