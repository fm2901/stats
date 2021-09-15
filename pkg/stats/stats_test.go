package stats

import (
	"reflect"
	"testing"

	"github.com/fm2901/bank/v2/pkg/types"
)

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "home"},
		{ID: 4, Category: "fun"},
	}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "home"},
		{ID: 4, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
	}
	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "auto"},
		{ID: 3, Category: "home"},
		{ID: 4, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "auto"},
	}
	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}


func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000},
		{ID: 2, Category: "auto", Amount: 2_000},
		{ID: 3, Category: "home", Amount: 3_000},
		{ID: 4, Category: "fun", Amount: 4_000},
	}
	expected := map[types.Category]types.Money{
		"auto" : 3000,
		"home" : 3000,
		"fun"  : 4000,
	}
	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000},
		{ID: 2, Category: "auto", Amount: 2_000},
		{ID: 3, Category: "home", Amount: 3_000},
		{ID: 4, Category: "fun", Amount: 4_000},
	}
	expected := map[types.Category]types.Money{
		"auto" : 1500,
		"home" : 3000,
		"fun"  : 4000,
	}
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}