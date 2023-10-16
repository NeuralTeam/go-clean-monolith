package middleware

// ---------------------------------------------------------------------------------------------------------------------

// Middleware description
type Middleware struct {
	Logger   LoggerMiddleware
	CORS     CORSMiddleware
	Recovery RecoveryMiddleware
}

// NewMiddleware description
func NewMiddleware(
	logger LoggerMiddleware,
	cors CORSMiddleware,
	recovery RecoveryMiddleware,
) Middleware {
	return Middleware{
		Logger:   logger,
		CORS:     cors,
		Recovery: recovery,
	}
}

// ---------------------------------------------------------------------------------------------------------------------
