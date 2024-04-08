package typesense

import (
	"fmt"
	"testing"
)

type Kek int64

type A struct {
	ID   string `json:"id,index"`
	Name string `json:"name,index"`
	Kek  *Kek   `json:"kek"`
	B
}

type C struct {
	ID    string  `json:"id,index"`
	CName *string `json:"cname,index"`
	Be    *B      `json:"be"`
	B
}

type B struct {
	LocId string   `json:"loc_id"`
	OrgId string   `json:"org_id"`
	Users []string `json:"users"`
}

func TestCollection(t *testing.T) {
	c := Client{}
	c.CreateCollection(nil, "kek", A{})
	fmt.Println("______")
	c.CreateCollection(nil, "kek", C{})

}
