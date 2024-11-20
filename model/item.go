// model/item.go
package model

type Item struct {
    ID            int    `json:"id"`
    DeskripsiItem string `json:"deskripsi_item"`
    HargaBeli     string `json:"harga_beli"`
    Stok          int    `json:"stok"`
}
