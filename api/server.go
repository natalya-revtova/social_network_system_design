package api

type Server struct{}

//	@title		Social network for travelers API
//	@version	1.0
//	@BasePath	/api
func NewServer() *Server {
	return &Server{}
}
