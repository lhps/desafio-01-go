package model_test

import (
	uuid "github.com/satori/go.uuid"
	"testing"

	"github.com/lhps/desafio-01/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewProduct(t *testing.T) {

	name := "Produto de teste"
	description := "Descrição do produto de teste"
	price := 1.99

	product, err := model.NewProduct(name, description, price)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(product.ID))
	require.Equal(t, product.Name, name)
	require.Equal(t, product.Description, description)
	require.Equal(t, product.Price, price)

	_, err = model.NewProduct(name, "", price)
	require.NotNil(t, err)
	_, err = model.NewProduct("", description, price)
	require.NotNil(t, err)
	_, err = model.NewProduct(name, description, 0)
	require.NotNil(t, err)
}
