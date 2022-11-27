package pkg

type Response struct {
	Body interface{} `json:"body"`
}

// func JSONresponse(r *echo.Response, code int, body interface{}) error {
// 	//r.Header().Set("Access-Control-Allow-Credentials", "true") // TODO это в мидлвару
// 	//r.Header().Set("Access-Control-Allow-Origin", "http://localhost")
// 	//r.Header().Set("Access-Control-Allow-Headers", "*")
// 	//r.Header().Set("Content-Type", "application/json")

// 	response, err := json.Marshal(Response{
// 		Body: body,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	r.WriteHeader(code)
// 	r.Write(response)
// 	return nil
// }

// func ErrorResponse(r *echo.Response, code int, errorMessage string) error {
// 	return JSONresponse(r, code, Error{
// 		Description: errorMessage,
// 	})
// }
