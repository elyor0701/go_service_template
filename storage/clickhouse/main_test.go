package clickhouse_test

import (
	"ucode_go_query_service/storage/clickhouse"

	_ "github.com/ClickHouse/clickhouse-go"
)

var (
	strg clickhouse.StorageCHI
)

//func TestMain(m *testing.M) {
//
//	cfg := config.Load()
//
//	clickhouseUrl := fmt.Sprintf("tcp://%v:%v?username=%v&password=%v&database=%v&read_timeout=10&write_timeout=18",
//		cfg.ClickHouseHost,
//		cfg.ClickHousePort,
//		cfg.ClickHouseUser,
//		cfg.ClickHousePassword,
//		cfg.ClickHouseDatabase,
//	)
//
//	clickhouseConn, err := sqlx.Connect("clickhouse", clickhouseUrl)
//	if err != nil {
//		fmt.Print(err)
//	}
//
//	strg = storage.NewStorageCH(clickhouseConn)
//
//	os.Exit(m.Run())
//}
