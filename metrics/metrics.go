package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Events = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sp_events",
			Help: ": number of events received from network",
		},
		[]string{"type"},
	)

	ConnNumbers = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sp_conn_connection_numbers",
			Help: ": number of connections",
		},
		[]string{"type"})

	ConnReconnection = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sp_conn_reconnection_counters",
			Help: ": number of re-connections from each ip address",
		},
		[]string{"ip_address"},
	)

	InboundSpeed = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pp_inbound_speed",
			Help: ": inbound speed summarized from slice related traffics",
		},
		[]string{"in_speed"})

	OutboundSpeed = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pp_outbound_speed",
			Help: ": outbound speed summarized from slice related traffics",
		},
		[]string{"out_speed"})

	TaskCount = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pp_task_cnt",
			Help: ": count of ongoing tasks",
		},
		[]string{"task_cnt"})

	StoredSliceCount = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pp_stored_slices_cnt",
			Help: ": count of stored slices",
		},
		[]string{"stored_slices_cnt"})

	RpcReqCount = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pp_rpc_req_cnt",
			Help: ": count of rpc requests",
		},
		[]string{"rpc_req_cnt"})
)
