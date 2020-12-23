package config

type piccolo struct {
	ServerAddress    string `goconf:"piccolo:piccolo_server_address"`     // ServerAddress : IP address of server which installed flute module
	ServerPort       int64  `goconf:"piccolo:piccolo_server_port"`        // ServerPort : Listening port number of flute module
	RequestTimeoutMs int64  `goconf:"piccolo:piccolo_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for GraphQL request to flute module
}

var Piccolo piccolo
