package catalog

const (
	// ErrorRouteNotFound is the error message for an invalid API route
	ErrorRouteNotFound = "the api route could not be found"
	// ErrorLoadingConfig is the error message when the configuration cannot be loaded
	ErrorLoadingConfig = "cannot load configuration from environment"
	// ErrorInitLogger is the error message when the logger cannot be initialized
	ErrorInitLogger = "cannot initialize the logger"
	// ErrorMarshalJSONResponse is the error message when JSON cannot be marshaled for http response
	ErrorMarshalJSONResponse = "json could not be marshaled: review status code"
	// ErrorBadRequest is the error message for an internal server error
	ErrorBadRequest = "bad request, unable to process"
)
