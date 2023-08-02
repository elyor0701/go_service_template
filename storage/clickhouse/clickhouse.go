package clickhouse

import (
	"fmt"
	"strings"
	"ucode_go_query_service/genproto/common_messages"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

func NewClickhouseConn(creds *common_messages.ClickhouseCredentials) (*sqlx.DB, error) {
	chUrl := fmt.Sprintf(`tcp://%s:%s?username=%s&password=%s&database=%s`,
		creds.GetHost(),
		strings.Trim(creds.GetPort(), ":"),
		creds.GetUsername(),
		creds.GetPassword(),
		creds.GetDatabase(),
	)

	chConn, err := sqlx.Open("clickhouse", chUrl)
	if err != nil {
		return nil, err
	}

	if err = chConn.Ping(); err != nil {
		return nil, err
	}

	return chConn, nil
}
