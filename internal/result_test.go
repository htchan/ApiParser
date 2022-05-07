package internal

import (
	"testing"
)

func TestData(t *testing.T) {
	d := Data{ "exist": "existing data" }
	t.Run("Data", func (t *testing.T) {
		t.Run("existing key", func (t *testing.T) {
			t.Parallel()
			if result := d.Data("exist"); result != "existing data" {
				t.Errorf(`want "%v", got "%v"`, "existing data", result)
			}
		})

		t.Run("not existing key", func (t *testing.T) {
			t.Parallel()
			if result := d.Data("not_exist"); result != "" {
				t.Errorf(`want "%v", got "%v"`, "", result)
			}
		})
	})
}

func TestResult(t *testing.T) {
	r := Result{
		Data: Data{ "exist": "data" },
		Items: []Data{ Data{ "items": "data" }},
	}

	t.Run("GetData", func (t *testing.T) {
		t.Run("exist key", func (t *testing.T) {
			if result := r.GetData("exist"); result != "data" {
				t.Errorf(`want "%v, got "%v"`, "data", result)
			}
		})

		t.Run("not exist key", func (t *testing.T) {
			if result := r.GetData("not_exist"); result != "" {
				t.Errorf(`want "%v, got "%v"`, "", result)
			}
		})
	})

	t.Run("GetItems", func (t *testing.T) {
		t.Run("within length", func (t *testing.T) {
			if result := r.GetItems(0); result.Data("items") != "data" {
				t.Errorf(`want "%v", got "%v"`, r.Items[0], result)
			}
		})

		t.Run("out of length", func (t *testing.T) {
			if result := r.GetItems(1); result.Data("exist") != "" {
				t.Errorf(`want "%v", got "%v"`, Data{}, result)
			}
		})
	})
}