// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/notnulldev/e-com-svelte/ent/product"
	"github.com/notnulldev/e-com-svelte/ent/user"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetPrice sets the "price" field.
func (pc *ProductCreate) SetPrice(f float64) *ProductCreate {
	pc.mutation.SetPrice(f)
	return pc
}

// SetPreviewURL sets the "preview_url" field.
func (pc *ProductCreate) SetPreviewURL(s string) *ProductCreate {
	pc.mutation.SetPreviewURL(s)
	return pc
}

// SetNillablePreviewURL sets the "preview_url" field if the given value is not nil.
func (pc *ProductCreate) SetNillablePreviewURL(s *string) *ProductCreate {
	if s != nil {
		pc.SetPreviewURL(*s)
	}
	return pc
}

// SetCategories sets the "categories" field.
func (pc *ProductCreate) SetCategories(s []string) *ProductCreate {
	pc.mutation.SetCategories(s)
	return pc
}

// SetImages sets the "images" field.
func (pc *ProductCreate) SetImages(s []string) *ProductCreate {
	pc.mutation.SetImages(s)
	return pc
}

// SetImagesStorage sets the "imagesStorage" field.
func (pc *ProductCreate) SetImagesStorage(s string) *ProductCreate {
	pc.mutation.SetImagesStorage(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProductCreate) SetDescription(s string) *ProductCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetStock sets the "stock" field.
func (pc *ProductCreate) SetStock(i int) *ProductCreate {
	pc.mutation.SetStock(i)
	return pc
}

// SetStockReserved sets the "stock_reserved" field.
func (pc *ProductCreate) SetStockReserved(i int) *ProductCreate {
	pc.mutation.SetStockReserved(i)
	return pc
}

// SetSellerID sets the "seller" edge to the User entity by ID.
func (pc *ProductCreate) SetSellerID(id int) *ProductCreate {
	pc.mutation.SetSellerID(id)
	return pc
}

// SetNillableSellerID sets the "seller" edge to the User entity by ID if the given value is not nil.
func (pc *ProductCreate) SetNillableSellerID(id *int) *ProductCreate {
	if id != nil {
		pc = pc.SetSellerID(*id)
	}
	return pc
}

// SetSeller sets the "seller" edge to the User entity.
func (pc *ProductCreate) SetSeller(u *User) *ProductCreate {
	return pc.SetSellerID(u.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.PreviewURL(); !ok {
		v := product.DefaultPreviewURL
		pc.mutation.SetPreviewURL(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Product.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Product.price"`)}
	}
	if v, ok := pc.mutation.Price(); ok {
		if err := product.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Product.price": %w`, err)}
		}
	}
	if _, ok := pc.mutation.PreviewURL(); !ok {
		return &ValidationError{Name: "preview_url", err: errors.New(`ent: missing required field "Product.preview_url"`)}
	}
	if v, ok := pc.mutation.PreviewURL(); ok {
		if err := product.PreviewURLValidator(v); err != nil {
			return &ValidationError{Name: "preview_url", err: fmt.Errorf(`ent: validator failed for field "Product.preview_url": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Categories(); !ok {
		return &ValidationError{Name: "categories", err: errors.New(`ent: missing required field "Product.categories"`)}
	}
	if _, ok := pc.mutation.Images(); !ok {
		return &ValidationError{Name: "images", err: errors.New(`ent: missing required field "Product.images"`)}
	}
	if _, ok := pc.mutation.ImagesStorage(); !ok {
		return &ValidationError{Name: "imagesStorage", err: errors.New(`ent: missing required field "Product.imagesStorage"`)}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Product.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := product.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Product.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Stock(); !ok {
		return &ValidationError{Name: "stock", err: errors.New(`ent: missing required field "Product.stock"`)}
	}
	if _, ok := pc.mutation.StockReserved(); !ok {
		return &ValidationError{Name: "stock_reserved", err: errors.New(`ent: missing required field "Product.stock_reserved"`)}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(product.Table, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := pc.mutation.PreviewURL(); ok {
		_spec.SetField(product.FieldPreviewURL, field.TypeString, value)
		_node.PreviewURL = value
	}
	if value, ok := pc.mutation.Categories(); ok {
		_spec.SetField(product.FieldCategories, field.TypeJSON, value)
		_node.Categories = value
	}
	if value, ok := pc.mutation.Images(); ok {
		_spec.SetField(product.FieldImages, field.TypeJSON, value)
		_node.Images = value
	}
	if value, ok := pc.mutation.ImagesStorage(); ok {
		_spec.SetField(product.FieldImagesStorage, field.TypeString, value)
		_node.ImagesStorage = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Stock(); ok {
		_spec.SetField(product.FieldStock, field.TypeInt, value)
		_node.Stock = value
	}
	if value, ok := pc.mutation.StockReserved(); ok {
		_spec.SetField(product.FieldStockReserved, field.TypeInt, value)
		_node.StockReserved = value
	}
	if nodes := pc.mutation.SellerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.SellerTable,
			Columns: []string{product.SellerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_products = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
