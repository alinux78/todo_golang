package client

import (
	"fmt"
	"testing"

	"github.com/alinux78/todo/pkg/model"
)

func TestClient(t *testing.T) {
	client := NewClient("http://localhost:8080")
	_, err := client.GetAll()

	//TODO asserts "github.com/alecthomas/assert"

	//"go.testFlags": ["-v"]

	if err != nil {
		t.Fatalf("Error %v", err)
	}

	//DISCUSS - don't do this
	newItem := model.TodoItemInput{"new item", "no description"}
	createdItem, err := client.Create(newItem)

	if err != nil {
		t.Fatalf("Error %v", err)
	}

	fmt.Printf("created: %v\n", createdItem)

	_, err = client.GetAll()

	if err != nil {
		t.Fatalf("Error %v", err)
	}

}
