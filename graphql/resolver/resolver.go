// graphql/resolver/resolver.go
package resolver

import (
    "context"
    "grpc1/graphql/generated"
    "grpc1/graphql/models"
    "grpc1/grpc/proto" // gRPC client
)

type Resolver struct {
    ItemServiceClient proto.ItemServiceClient
}

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input models.NewItem) (*models.Item, error) {
    req := &proto.Item{
        DeskripsiItem: input.DeskripsiItem,
        HargaBeli:     input.HargaBeli,
        Stok:          int32(input.Stok),
    }

    resp, err := r.ItemServiceClient.CreateItem(ctx, req)
    if err != nil {
        return nil, err
    }

    return &models.Item{
        ID:            int(resp.Id),
        DeskripsiItem: resp.DeskripsiItem,
        HargaBeli:     resp.HargaBeli,
        Stok:          int(resp.Stok),
    }, nil
}

// UpdateItem is the resolver for the updateItem field.
func (r *mutationResolver) UpdateItem(ctx context.Context, input models.UpdateItem) (*models.Item, error) {
    req := &proto.Item{
        Id:            int32(input.ID),
        DeskripsiItem: input.DeskripsiItem,
        HargaBeli:     input.HargaBeli,
        Stok:          int32(input.Stok),
    }

    resp, err := r.ItemServiceClient.UpdateItem(ctx, req)
    if err != nil {
        return nil, err
    }

    return &models.Item{
        ID:            int(resp.Id),
        DeskripsiItem: resp.DeskripsiItem,
        HargaBeli:     resp.HargaBeli,
        Stok:          int(resp.Stok),
    }, nil
}

// DeleteItem is the resolver for the deleteItem field.
func (r *mutationResolver) DeleteItem(ctx context.Context, id int) (*models.Item, error) {
    req := &proto.ItemRequest{Id: int32(id)}

    // Hapus item melalui gRPC service
    _, err := r.ItemServiceClient.DeleteItem(ctx, req)
    if err != nil {
        return nil, err
    }

    return &models.Item{
        ID: int(id), // Mengembalikan ID yang telah dihapus
    }, nil
}

// GetItemByID is the resolver for the getItemById field.
func (r *queryResolver) GetItemByID(ctx context.Context, id int) (*models.Item, error) {
    req := &proto.ItemRequest{Id: int32(id)}
    resp, err := r.ItemServiceClient.GetItemById(ctx, req)
    if err != nil {
        return nil, err
    }

    return &models.Item{
        ID:            int(resp.Id),
        DeskripsiItem: resp.DeskripsiItem,
        HargaBeli:     resp.HargaBeli,
        Stok:          int(resp.Stok),
    }, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
