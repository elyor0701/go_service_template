package clickhouse

import (
	"fmt"
	"strings"
	pb "ucode_go_query_service/genproto/query_service"
	"ucode_go_query_service/pkg/helper"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAll(db *sqlx.DB, in *pb.Query, bodyMap map[string]interface{}) (map[string]interface{}, error) {

	query, ok := bodyMap["body"]
	if !ok {
		fmt.Println("Query not found")
		return nil, status.Error(codes.Internal, "query not found")
	}
	vars := map[string]interface{}{}
	// replace variables {{some_field}} to some_field =  :some_field
	for _, variable := range in.GetVariables() {
		var val string
		if variable.GetValue() != "" {
			val = variable.GetValue()
		} else {
			val = variable.GetDefault()
		}
		query = strings.ReplaceAll(query.(string), "{{"+variable.GetKey()+"}}", ":"+variable.GetKey())
		vars[variable.GetKey()] = val
	}
	namedQuery, args := helper.ReplaceQueryParams(query.(string), vars)
	fmt.Println("aaa:::::", namedQuery)
	rows, err := db.DB.Query(namedQuery, args...)
	if err != nil {
		fmt.Println("err in query", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("Error while get columns of rows")
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()
	var allMaps []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}
		err = rows.Scan(pointers...)
		resultMap := make(map[string]interface{})
		for i, val := range values {
			fmt.Printf("Adding key=%s val=%v\n", columns[i], val)
			typeOfInterface := fmt.Sprintf("%T", val)
			if typeOfInterface == "[]string" && len(val.([]string)) > 0 {
				fmt.Println("check", len(val.([]string)))
				firstElement := val.([]string)[0]
				fmt.Println("val::", val)
				fmt.Println("first:::", firstElement)
				if firstElement != "" && firstElement[:1] == "{" {
					convertedMap, err := helper.ConvertStringToMap(val.([]string))
					if err != nil {
						return nil, errors.Wrap(err, "error while converting string to map")
					}
					resultMap[columns[i]] = convertedMap
					continue
				}
			} else if typeOfInterface == "uint8" && val != nil {
				if val.(uint8) == uint8(1) {
					val = true
				} else {
					val = false
				}
			}
			resultMap[columns[i]] = val
		}
		allMaps = append(allMaps, resultMap)
	}
	fmt.Println(allMaps)

	return map[string]interface{}{"objects": allMaps}, nil
}
