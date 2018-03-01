package etcd

type (
	EtcdBackend struct {
		HostGroupId       uint64 `db:"hostgroup_id,omitempty" json:"hostgroup_id"`
		HostName          string `db:"hostname" json:"hostname"`
		Port              uint64 `db:"port" json:"port"`
		Status            string `db:"status" json:"status"`
		Weight            uint64 `db:"weight" json:"weight"`
		Compression       uint64 `db:"compression" json:"compression"`
		MaxConnections    uint64 `db:"max_connections" json:"max_connections"`
		MaxReplicationLag uint64 `db:"max_replication_lag" json:"max_replication_lag"`
		UseSsl            uint64 `db:"use_ssl" json:"use_ssl"`
		MaxLatencyMs      uint64 `db:"max_latency_ms" json:"max_latency_ms"`
		Comment           string `db:"comment" json:"comment"`
	}
)

// init a new servers.
func NewServer(hostgroup_id uint64, hostname string, port uint64) (*EtcdBackend, error) {
	newsrv := new(EtcdBackend)

	newsrv.HostGroupId = hostgroup_id
	newsrv.HostName = hostname
	newsrv.Port = port

	newsrv.Status = "ONLINE"
	newsrv.Weight = 1000
	newsrv.Compression = 0
	newsrv.MaxConnections = 10000
	newsrv.MaxReplicationLag = 0
	newsrv.UseSsl = 0
	newsrv.MaxLatencyMs = 0
	newsrv.Comment = ""

	return newsrv, nil
}

// set servers status
func (srvs *EtcdBackend) SetServerStatus(status string) {
	switch status {
	case "ONLINE":
		srvs.Status = "ONLINE"
	case "SHUNNED":
		srvs.Status = "SHUNNED"
	case "OFFLINE_SOFT":
		srvs.Status = "OFFLINE_SOFT"
	case "OFFLINE_HARD":
		srvs.Status = "OFFLINE_HARD"
	default:
		srvs.Status = "ONLINE"
	}
}

// set servers weight
func (srvs *EtcdBackend) SetServerWeight(weight uint64) {
	srvs.Weight = weight
}

// set servers compression
func (srvs *EtcdBackend) SetServerCompression(compression uint64) {
	srvs.Compression = compression
}

// set servers max_connections
func (srvs *EtcdBackend) SetServerMaxConnection(max_connections uint64) {
	if max_connections >= 10000 {
		srvs.MaxConnections = 10000
	} else {
		srvs.MaxConnections = max_connections
	}
}

// set servers max_replication_lag
func (srvs *EtcdBackend) SetServerMaxReplicationLag(max_replication_lag uint64) {
	if max_replication_lag > 126144000 {
		srvs.MaxReplicationLag = 1261440000
	} else {
		srvs.MaxReplicationLag = max_replication_lag
	}
}

// set servers use_ssl
func (srvs *EtcdBackend) SetServerUseSSL(use_ssl uint64) {
	if use_ssl >= 1 {
		srvs.UseSsl = 1
	} else {
		srvs.UseSsl = 0
	}
}

// set servers max_latency_ms
func (srvs *EtcdBackend) SetServerMaxLatencyMs(max_latency_ms uint64) {
	srvs.MaxLatencyMs = max_latency_ms
}

// set servers comment
func (srvs *EtcdBackend) SetServersComment(comment string) {
	srvs.Comment = comment
}
