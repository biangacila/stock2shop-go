package io

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

func connect() {
	cluster := setupCluster() //gocql.NewCluster(CusterHost1, CusterHost2, CusterHost3)
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println("Cassandra cluster.CreateSession err : ", err)
	}
	Session = session
}

func setupCluster() *gocql.ClusterConfig {
	cluster := gocql.NewCluster(Servers...)
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	cluster.Compressor = &gocql.SnappyCompressor{}
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{NumRetries: 3}
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = 10 * time.Second
	cluster.Timeout = 10 * time.Second
	cluster.ProtoVersion = 4
	cluster.DisableInitialHostLookup = true
	return cluster
}
