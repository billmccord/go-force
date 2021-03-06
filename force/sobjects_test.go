package force

import (
	"math/rand"
	"testing"
	"time"

	"github.com/nimajalali/go-force/sobjects"
)

const (
	AccountId      = "001i000000RxW18"
	CustomObjectId = "a00i0000009SPer"
)

type CustomSObject struct {
	sobjects.BaseSObject
	Active    bool   `force:"Active__c"`
	AccountId string `force:"Account__c"`
}

func (t *CustomSObject) ApiName() string {
	return "CustomObject__c"
}

func init() {
	initTest()
}

func TestDescribeSObject(t *testing.T) {
	acc := &sobjects.Account{}

	desc, err := DescribeSObject(acc)
	if err != nil {
		t.Fatalf("Cannot retrieve SObject Description for Account SObject: %v", err)
	}

	t.Logf("SObject Description for Account Retrieved: %+v", desc)
}

func TestGetSObject(t *testing.T) {
	// Test Standard Object
	acc := &sobjects.Account{}

	err := GetSObject(AccountId, acc)
	if err != nil {
		t.Fatalf("Cannot retrieve SObject Account: %v", err)
	}

	t.Logf("SObject Account Retrieved: %+v", acc)

	// Test Custom Object
	customObject := &CustomSObject{}

	err = GetSObject(CustomObjectId, customObject)
	if err != nil {
		t.Fatalf("Cannot retrieve SObject CustomObject: %v", err)
	}

	t.Logf("SObject CustomObject Retrieved: %+v", customObject)
}

func TestUpdateSObject(t *testing.T) {
	// Need some random text for updating a field.
	rand.Seed(time.Now().UTC().UnixNano())
	someText := randomString(10)

	// Test Standard Object
	acc := &sobjects.Account{}
	acc.Name = someText

	err := UpdateSObject(AccountId, acc)
	if err != nil {
		t.Fatalf("Cannot update SObject Account: %v", err)
	}

	// Read back and verify
	err = GetSObject(AccountId, acc)
	if err != nil {
		t.Fatalf("Cannot retrieve SObject Account: %v", err)
	}

	if acc.Name != someText {
		t.Fatalf("Update SObject Account failed. Failed to persist.")
	}

	t.Logf("Updated SObject Account: %+v", acc)
}

func TestInsertDeleteSObject(t *testing.T) {
	objectId := insertSObject(t)
	deleteSObject(t, objectId)
}

func insertSObject(t *testing.T) string {
	// Need some random text for name field.
	rand.Seed(time.Now().UTC().UnixNano())
	someText := randomString(10)

	// Test Standard Object
	acc := &sobjects.Account{}
	acc.Name = someText

	resp, err := InsertSObject(acc)
	if err != nil {
		t.Fatalf("Insert SObject Account failed: %v", err)
	}

	if len(resp.Id) == 0 {
		t.Fatalf("Insert SObject Account failed to return Id: %+v", resp)
	}

	return resp.Id
}

func deleteSObject(t *testing.T, id string) {
	// Test Standard Object
	acc := &sobjects.Account{}

	err := DeleteSObject(id, acc)
	if err != nil {
		t.Fatalf("Delete SObject Account failed: %v", err)
	}

	// Read back and verify
	err = GetSObject(id, acc)
	if err == nil {
		t.Fatalf("Delete SObject Account failed, was able to retrieve deleted object: %+v", acc)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
