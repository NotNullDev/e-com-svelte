// Code generated by ent, DO NOT EDIT.

package product

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/notnulldev/e-com-svelte/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// PreviewURL applies equality check predicate on the "preview_url" field. It's identical to PreviewURLEQ.
func PreviewURL(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPreviewURL, v))
}

// ImagesStorage applies equality check predicate on the "imagesStorage" field. It's identical to ImagesStorageEQ.
func ImagesStorage(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldImagesStorage, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDescription, v))
}

// Stock applies equality check predicate on the "stock" field. It's identical to StockEQ.
func Stock(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStock, v))
}

// StockReserved applies equality check predicate on the "stock_reserved" field. It's identical to StockReservedEQ.
func StockReserved(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStockReserved, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldName, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldPrice, v))
}

// PreviewURLEQ applies the EQ predicate on the "preview_url" field.
func PreviewURLEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPreviewURL, v))
}

// PreviewURLNEQ applies the NEQ predicate on the "preview_url" field.
func PreviewURLNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldPreviewURL, v))
}

// PreviewURLIn applies the In predicate on the "preview_url" field.
func PreviewURLIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldPreviewURL, vs...))
}

// PreviewURLNotIn applies the NotIn predicate on the "preview_url" field.
func PreviewURLNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldPreviewURL, vs...))
}

// PreviewURLGT applies the GT predicate on the "preview_url" field.
func PreviewURLGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldPreviewURL, v))
}

// PreviewURLGTE applies the GTE predicate on the "preview_url" field.
func PreviewURLGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldPreviewURL, v))
}

// PreviewURLLT applies the LT predicate on the "preview_url" field.
func PreviewURLLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldPreviewURL, v))
}

// PreviewURLLTE applies the LTE predicate on the "preview_url" field.
func PreviewURLLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldPreviewURL, v))
}

// PreviewURLContains applies the Contains predicate on the "preview_url" field.
func PreviewURLContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldPreviewURL, v))
}

// PreviewURLHasPrefix applies the HasPrefix predicate on the "preview_url" field.
func PreviewURLHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldPreviewURL, v))
}

// PreviewURLHasSuffix applies the HasSuffix predicate on the "preview_url" field.
func PreviewURLHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldPreviewURL, v))
}

// PreviewURLEqualFold applies the EqualFold predicate on the "preview_url" field.
func PreviewURLEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldPreviewURL, v))
}

// PreviewURLContainsFold applies the ContainsFold predicate on the "preview_url" field.
func PreviewURLContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldPreviewURL, v))
}

// ImagesStorageEQ applies the EQ predicate on the "imagesStorage" field.
func ImagesStorageEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldImagesStorage, v))
}

// ImagesStorageNEQ applies the NEQ predicate on the "imagesStorage" field.
func ImagesStorageNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldImagesStorage, v))
}

// ImagesStorageIn applies the In predicate on the "imagesStorage" field.
func ImagesStorageIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldImagesStorage, vs...))
}

// ImagesStorageNotIn applies the NotIn predicate on the "imagesStorage" field.
func ImagesStorageNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldImagesStorage, vs...))
}

// ImagesStorageGT applies the GT predicate on the "imagesStorage" field.
func ImagesStorageGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldImagesStorage, v))
}

// ImagesStorageGTE applies the GTE predicate on the "imagesStorage" field.
func ImagesStorageGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldImagesStorage, v))
}

// ImagesStorageLT applies the LT predicate on the "imagesStorage" field.
func ImagesStorageLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldImagesStorage, v))
}

// ImagesStorageLTE applies the LTE predicate on the "imagesStorage" field.
func ImagesStorageLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldImagesStorage, v))
}

// ImagesStorageContains applies the Contains predicate on the "imagesStorage" field.
func ImagesStorageContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldImagesStorage, v))
}

// ImagesStorageHasPrefix applies the HasPrefix predicate on the "imagesStorage" field.
func ImagesStorageHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldImagesStorage, v))
}

// ImagesStorageHasSuffix applies the HasSuffix predicate on the "imagesStorage" field.
func ImagesStorageHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldImagesStorage, v))
}

// ImagesStorageEqualFold applies the EqualFold predicate on the "imagesStorage" field.
func ImagesStorageEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldImagesStorage, v))
}

// ImagesStorageContainsFold applies the ContainsFold predicate on the "imagesStorage" field.
func ImagesStorageContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldImagesStorage, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldDescription, v))
}

// StockEQ applies the EQ predicate on the "stock" field.
func StockEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStock, v))
}

// StockNEQ applies the NEQ predicate on the "stock" field.
func StockNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldStock, v))
}

// StockIn applies the In predicate on the "stock" field.
func StockIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldStock, vs...))
}

// StockNotIn applies the NotIn predicate on the "stock" field.
func StockNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldStock, vs...))
}

// StockGT applies the GT predicate on the "stock" field.
func StockGT(v int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldStock, v))
}

// StockGTE applies the GTE predicate on the "stock" field.
func StockGTE(v int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldStock, v))
}

// StockLT applies the LT predicate on the "stock" field.
func StockLT(v int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldStock, v))
}

// StockLTE applies the LTE predicate on the "stock" field.
func StockLTE(v int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldStock, v))
}

// StockReservedEQ applies the EQ predicate on the "stock_reserved" field.
func StockReservedEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStockReserved, v))
}

// StockReservedNEQ applies the NEQ predicate on the "stock_reserved" field.
func StockReservedNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldStockReserved, v))
}

// StockReservedIn applies the In predicate on the "stock_reserved" field.
func StockReservedIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldStockReserved, vs...))
}

// StockReservedNotIn applies the NotIn predicate on the "stock_reserved" field.
func StockReservedNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldStockReserved, vs...))
}

// StockReservedGT applies the GT predicate on the "stock_reserved" field.
func StockReservedGT(v int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldStockReserved, v))
}

// StockReservedGTE applies the GTE predicate on the "stock_reserved" field.
func StockReservedGTE(v int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldStockReserved, v))
}

// StockReservedLT applies the LT predicate on the "stock_reserved" field.
func StockReservedLT(v int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldStockReserved, v))
}

// StockReservedLTE applies the LTE predicate on the "stock_reserved" field.
func StockReservedLTE(v int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldStockReserved, v))
}

// HasSeller applies the HasEdge predicate on the "seller" edge.
func HasSeller() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SellerTable, SellerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSellerWith applies the HasEdge predicate on the "seller" edge with a given conditions (other predicates).
func HasSellerWith(preds ...predicate.User) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newSellerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		p(s.Not())
	})
}
