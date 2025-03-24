package more

type Configuratoin func(s *Server) error
type Server struct {
	Host string
	Port int
}

func NewServer(cfgs ...Configuratoin) (*Server, error) {
	s := &Server{}
	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func WithPort(port int) Configuratoin {
	return func(s *Server) error {
		s.Port = port
		return nil
	}
}

func WithHost(host string) Configuratoin {
	return func(s *Server) error {
		s.Host = host
		return nil
	}
}
