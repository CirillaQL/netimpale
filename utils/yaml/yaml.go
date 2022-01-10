package yaml

type Client struct {
	TCP struct {
		Host string `yaml:"host"`
		Port uint32 `yaml:"port"`
	}
	UDP struct {
		Host string `yaml:"host"`
		Port uint32 `yaml:"port"`
	}
}

type Server struct {
}
