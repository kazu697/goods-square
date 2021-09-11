package repository

import (
	"goods-square/internal/domain/model"
	"goods-square/internal/domain/repository"
	"goods-square/pkg/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProductRepo_List(t *testing.T) {
	type args struct {
		repository.ProductListQuery
	}

	type state struct {
		products model.Products
	}

	type want struct {
		products model.Products
	}

	tests := []struct {
		name    string
		args    args
		state   state
		want    want
		wantErr bool
	}{
		{
			name: "success with no filter",
			args: args{},
			state: state{
				products: model.Products{
					{
						ID:   1,
						Type: "type",
						Name: "name",
					},
					// {
					// 	ID:   2,
					// 	Type: "type2",
					// 	Name: "name2",
					// },
				},
			},
			want: want{
				products: model.Products{
					{
						ID:   1,
						Type: "type",
						Name: "name",
					},
					{
						ID:   2,
						Type: "type2",
						Name: "name2",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := testutil.TestSQLClient()
			// testutil.SetUp(db)
			defer db.Close()

			// for _, p := range tt.state.products {
			// 	if err := db.Create(dto.NewProductFromEntity(p)); err != nil {
			// 		t.Error(err)
			// 		return
			// 	}
			// }

			repo := NewProduct()
			got, err := repo.List(testutil.NewTestContext(), db, &tt.args.ProductListQuery)
			if err != nil {
				assert.Equal(t, tt.wantErr, err != nil)
				return
			}

			for i, g := range got {
				assert.Equal(t, tt.want.products[i].ID, g.ID)
				assert.Equal(t, tt.want.products[i].Name, g.Name)
				assert.Equal(t, tt.want.products[i].Type, g.Type)
			}
		})
	}
}
