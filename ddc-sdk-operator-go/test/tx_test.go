package test

import (
	"fmt"
	"testing"
)

func TestGetTransByHash(t *testing.T) {
	fmt.Println(client.GetTransByHash("0xd5cb1c52265379bf0d3a3e900f9bb54e45ce5a7188041e32136aa9bcb3488308"))
}
func TestGetTransReceipt(t *testing.T) {
	fmt.Println(client.GetTransReceipt("0x10e8d4e5f30b05d268b1b61ba89d85d8c8010153b21db00efe14f4656bf0b92b"))
}
func TestGetTransStatus(t *testing.T) {
	fmt.Println(client.GetTransStatus("0xc848f2cd49658cca7963c13d9218872bc667d3d0c4f8d1fc6ffb14df3cd029c7"))
}
