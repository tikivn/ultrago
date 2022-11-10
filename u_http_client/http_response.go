package u_http_client

type HttpResponse struct {
	status   int
	response []byte
	err      error
}
