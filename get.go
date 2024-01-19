package notif_latest_info

// 以下を満たす型をサポートします
type Getter interface {
	Get() error
	MakeMsg() (string, error)
}
