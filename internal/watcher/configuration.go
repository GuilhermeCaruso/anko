package watcher

type WatcherConfig struct {
	RootPath         string
	Extensions       []string
	Files            []string
	IsOpen           *bool
	Language         string
	AppPath          string
	DoneChan         chan bool
	SysOS            string
	DispatcherChan   chan string
	selectedLanguage *Language
}

const (
	ACT_INIT  = "init"
	ACT_STOP  = "stop"
	ACT_RESET = "reset"
)

func New(args WatcherConfig) *WatcherConfig {
	return &args
}
