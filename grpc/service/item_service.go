// grpc/service/item_service.go
package service

import (
    "context"
    "grpc1/database"
    "grpc1/model"
    pb "grpc1/grpc/proto"
)

type ItemServiceServer struct {
    pb.UnimplementedItemServiceServer // Tambahkan ini
}

// GetItemById mengembalikan item berdasarkan ID.
func (s *ItemServiceServer) GetItemById(ctx context.Context, req *pb.ItemRequest) (*pb.Item, error) {
    db, err := database.GetDB()
    if err != nil {
        return nil, err
    }
    defer db.Close()

    var item model.Item
    err = db.QueryRow("SELECT id, deskripsi_item, harga_beli, stok FROM item WHERE id = ?", req.Id).Scan(&item.ID, &item.DeskripsiItem, &item.HargaBeli, &item.Stok)
    if err != nil {
        return nil, err
    }

    return &pb.Item{
        Id:            int32(item.ID),
        DeskripsiItem: item.DeskripsiItem,
        HargaBeli:     item.HargaBeli,
        Stok:          int32(item.Stok),
    }, nil
}

// CreateItem menambahkan item baru ke database.
func (s *ItemServiceServer) CreateItem(ctx context.Context, req *pb.Item) (*pb.Item, error) {
    db, err := database.GetDB()
    if err != nil {
        return nil, err
    }
    defer db.Close()

    // Insert item baru ke database
    stmt, err := db.Prepare("INSERT INTO item (deskripsi_item, harga_beli, stok) VALUES (?, ?, ?)")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    result, err := stmt.Exec(req.DeskripsiItem, req.HargaBeli, req.Stok)
    if err != nil {
        return nil, err
    }

    lastID, err := result.LastInsertId()
    if err != nil {
        return nil, err
    }

    return &pb.Item{
        Id:            int32(lastID),
        DeskripsiItem: req.DeskripsiItem,
        HargaBeli:     req.HargaBeli,
        Stok:          req.Stok,
    }, nil
}

// UpdateItem memperbarui informasi item berdasarkan ID.
func (s *ItemServiceServer) UpdateItem(ctx context.Context, req *pb.Item) (*pb.Item, error) {
    db, err := database.GetDB()
    if err != nil {
        return nil, err
    }
    defer db.Close()

    // Update item di database
    stmt, err := db.Prepare("UPDATE item SET deskripsi_item = ?, harga_beli = ?, stok = ? WHERE id = ?")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    _, err = stmt.Exec(req.DeskripsiItem, req.HargaBeli, req.Stok, req.Id)
    if err != nil {
        return nil, err
    }

    return &pb.Item{
        Id:            req.Id,
        DeskripsiItem: req.DeskripsiItem,
        HargaBeli:     req.HargaBeli,
        Stok:          req.Stok,
    }, nil
}

// DeleteItem menghapus item berdasarkan ID.
func (s *ItemServiceServer) DeleteItem(ctx context.Context, req *pb.ItemRequest) (*pb.Empty, error) {
    db, err := database.GetDB()
    if err != nil {
        return nil, err
    }
    defer db.Close()

    // Hapus item dari database
    stmt, err := db.Prepare("DELETE FROM item WHERE id = ?")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    _, err = stmt.Exec(req.Id)
    if err != nil {
        return nil, err
    }

    return &pb.Empty{}, nil
}
