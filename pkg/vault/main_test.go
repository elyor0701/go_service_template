package vaultclient_test

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"testing"
// 	"time"
// 	vaultclient "ucode/ucode_go_api_ref_service/pkg/vault"

// 	"github.com/joho/godotenv"
// )

// func TestVaultClient(t *testing.T) {

// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		t.Log(err)
// 		t.FailNow()
// 	}
// 	var (
// 		roleID      string = os.Getenv("ROLE_ID")
// 		secretID    string = os.Getenv("SECRET_ID")
// 		mountPath   string = os.Getenv("MOUNT_PATH")
// 		secretPath  string = os.Getenv("SECRET_PATH")
// 		address     string = os.Getenv("VAULT_ADDRESS")
// 		countErrors int    = 0
// 	)

// 	defer func() {
// 		fmt.Println("errors count:", countErrors)
// 	}()

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	vc, err := vaultclient.NewClient(ctx, vaultclient.NewClientArgs{
// 		RoleID:    roleID,
// 		SecretID:  secretID,
// 		MountPath: mountPath,
// 		Address:   address,
// 	})
// 	if err != nil {
// 		t.Log(err)
// 		t.FailNow()
// 	}

// 	go func() {
// 		err := vc.RenewToken(ctx)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}()

// 	ticker := time.NewTicker(time.Millisecond * 100)
// 	time.AfterFunc(time.Minute*60, func() {
// 		ticker.Stop()
// 		fmt.Println("Stopping ticker")
// 	})

// 	for range ticker.C {
// 		err = vc.Put(ctx, secretPath+"/team19/project1/mongo", map[string]interface{}{
// 			"host": "0.0.0.0",
// 			"port": 80,
// 		})

// 		if err != nil {
// 			countErrors++
// 			t.Log(err)
// 			t.Fail()
// 		}

// 		secret, err := vc.Get(ctx, secretPath+"/team19/project1/mongo")
// 		if err != nil {
// 			countErrors++
// 			t.Log(err)
// 			t.Fail()
// 		}

// 		if bytes, err := json.Marshal(secret); err == nil {
// 			fmt.Println("seccessfully got:", string(bytes))
// 		}
// 	}

// }
