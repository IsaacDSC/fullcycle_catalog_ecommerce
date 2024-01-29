// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: products.sql

package sqlc

import (
	"context"
)

const getProducts = `-- name: GetProducts :many
SELECT id, code, name, image_url, price, description, active, category_id, created_at, updated_at, deleted_at FROM products
`

func (q *Queries) GetProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Name,
			&i.ImageUrl,
			&i.Price,
			&i.Description,
			&i.Active,
			&i.CategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}