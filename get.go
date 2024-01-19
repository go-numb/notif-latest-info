package notif_latest_info

type Geter interface {
	Get() error
	MakeMsg() (string, error)
}
