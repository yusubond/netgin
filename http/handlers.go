package http

func (s *Server) initHandlers() {
	s.POST("/auth/v1/register", s.RegisterUser)
	s.GET("/user", s.GetUserInfo)
}
