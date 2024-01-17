package vote

import (
	"log"
	"testing"
)

func TestClientV2(t *testing.T) {
	// Test successful connection with host
	client, err := ClientV2(dbhost)
	if err != nil {
		t.Fatalf("ClientV2 failed to connect: %v", err)
	}
	defer CloseV2(client)

	// Test connection reuse with key
	key := "testKey"
	client2, err := ClientV2(dbhost, key)
	if err != nil {
		t.Fatalf("ClientV2 failed to connect with key: %v", err)
	}
	if client != client2 {
		t.Errorf("ClientV2 did not reuse connection for same key")
	}

	client3, err := ClientV2(dbhost, key)
	if err != nil {
		t.Fatalf("ClientV2 failed to connect with key: %v", err)
	}
	if client3 != client2 {
		t.Fatalf("ClientV3 did not reuse connection for same key ClientV2")
	}

	// Test error handling for invalid host
	_, err = ClientV2("invalid://host")
	if err == nil {
		t.Error("ClientV2 did not return error for invalid host")
	}
}

func TestJwt(t *testing.T) {

	data, err := ClaimJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMxNjI1MjQsIm5hbWUiOiJOYW1lMSBBbGljZSIsInVzZXJJZCI6InVzZXIxIn0.sHjTfXXEN4OEM7V4hi6aem3R4dJEEOr9FBgk-iRhT2I")
	if err != nil {
		t.Errorf("ClaimJWT Token Error")
	}

	userId := data["userId"].(string)
	name := data["name"].(string)
	if userId == "" && name == "" {
		t.Error("userId  name Error")

	}
	log.Println(userId)
	log.Println(name)

}
func TestGenTokenuser(t *testing.T) {
	genTokenuser()
}
