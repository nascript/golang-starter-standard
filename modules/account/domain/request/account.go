package request

type (
	LoginRequest struct {
		// e.g Email, Password
	}

	RegisterRequest struct {
		// e.g Email, Password, Name, Username, ...
	}
)

func (l *LoginRequest) Validate() error {
	return nil
}

func (r *RegisterRequest) Validate() error {
	return nil
}
