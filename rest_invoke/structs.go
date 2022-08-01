package rest_invoke

type Lambda struct {
	Payload  Payload     `json:"payload"`
	Region   string      `json:"region"`
	Function string      `json:"function"`
	Object   interface{} `json:"object"`
}

type Payload struct {
	Path                  string      `json:"path"`
	HTTPMethod            string      `json:"httpMethod"`
	Body                  string      `json:"body"`
	IsBase64Encoded       string      `json:"isBase64Encoded"`
	QueryString           interface{} `json:"queryStringParameters"`
	MultiValueQueryString interface{} `json:"multiValueQueryStringParameters"`
	PathParameters        interface{} `json:"pathParameters"`
	StageVariables        interface{} `json:"stageVariables"`
	Headers               interface{} `json:"headers"`
}

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
	Headers    struct {
		ContentType string `json:"Content-Type"`
	} `json:"headers"`
	MultiValueHeaders struct {
		ContentType []string `json:"Content-Type"`
	} `json:"multiValueHeaders"`
}
