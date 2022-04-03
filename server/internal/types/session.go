package types

type Member struct {
	ID    string `json:"id"`
	Name  string `json:"displayname"`
	Admin bool   `json:"admin"`
	Muted bool   `json:"muted"`
}

type Session interface {
	ID() string
	Name() string
	Admin() bool
	Muted() bool
	Connected() bool
	Member() *Member

	SetMuted(muted bool)
	SetName(name string) error
	SetConnected(connected bool) error
	SetSocket(socket WebSocket) error

	Address() string
	Send(v interface{}) error
}

type SessionManager interface {
	New(id string, admin bool, socket WebSocket) Session

	HasHost() bool
	IsHost(id string) bool
	SetHost(id string) error
	GetHost() (Session, bool)
	ClearHost()

	Has(id string) bool
	Get(id string) (Session, bool)

	Members() []*Member
	Destroy(id string)
	Clear() error

	OnDestroy(listener func(id string, session Session))
	OnCreated(listener func(id string, session Session))
	OnConnected(listener func(id string, session Session))
}
