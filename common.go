package common

var (
	CurrentServerName = ""
)

const (
	AdminServerName = "admin"
	ApiServerName   = "api"
	ZcronServerName = "zcron"
	ChatServerName  = "chat"
)

var NullStringSlice = []string{}

// ---------- 通用结构 ----------

type TKStrVStr struct {
	K string `json:"k"`
	V string `json:"v"`
}

type IIDName interface {
	GetID() string
	SetName(name string)
}

type TIDName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (this *TIDName) Set(id, name string) {
	this.ID = id
	this.Name = name
}

func (this *TIDName) SetName(name string) { this.Name = name }

func (this *TIDName) GetID() string { return this.ID }
