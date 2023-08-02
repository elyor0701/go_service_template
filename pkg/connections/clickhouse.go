package connections

import (
	"errors"
	"fmt"
	"ucode_go_query_service/pkg/logger"

	clickhouse "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

type Credentials struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func NewClickhouseDb(cfg Credentials) (*sqlx.DB, error) {
	chUrl := fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&database=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)
	fmt.Println("conn string::", fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&database=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	))

	chConn, err := sqlx.Open("clickhouse", chUrl)
	if err != nil {
		fmt.Println("error -> Connection Clickhouse ->", err)
		return nil, errors.New("could not connect to clickhouse")
	}

	if err := chConn.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Println("Could not connect to clickhouse: ", logger.Any("exception", exception))
			return nil, errors.New("could not connect to clickhouse")
		} else {
			fmt.Println("Could not connect to clickhouse: ", logger.Error(err))
			return nil, errors.New("could not connect to clickhouse")
		}
	}
	return chConn, nil
}
