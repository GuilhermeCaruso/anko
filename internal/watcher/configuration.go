package watcher

// Watcher as well as Configuration File is the central configuration
// of the project, used to control the entire anko flow.
type Watcher struct {
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
	//ACT_INIT is the flag to start the service
	ACT_INIT = "init"
	//ACT_STOP is the flag to stop the service
	ACT_STOP = "stop"
	//ACT_RESET is the flag to reset the service
	ACT_RESET = "reset"
)

// New is a simple Watcher initializer
func New(args Watcher) *Watcher {
	return &args
}
