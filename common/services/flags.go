package services

import "flag"

type ServiceFlags struct {
	fs *flag.FlagSet

	address string
	port    int
}

const (
	defaultIpAddress = "localhost"
	defaultPort      = 8080
)

func NewServiceFlags() *ServiceFlags {
	serviceFlags := &ServiceFlags{
		fs: flag.NewFlagSet("service_flags", flag.PanicOnError),
	}

	serviceFlags.fs.StringVar(&serviceFlags.address, "address", defaultIpAddress, "service ip address")
	serviceFlags.fs.IntVar(&serviceFlags.port, "port", defaultPort, "service port")

	return serviceFlags
}

func (s *ServiceFlags) Init(args []string) error {
	return s.fs.Parse(args)
}

func (s *ServiceFlags) GetAddress() string {
	return s.address
}

func (s *ServiceFlags) GetPort() int {
	return s.port
}
