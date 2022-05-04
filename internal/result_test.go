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
		data: Data{ "exist": "data" },
		items: []Data{ Data{ "items": "data" }},
	}

	t.Run("Data", func (t *testing.T) {
		t.Run("exist key", func (t *testing.T) {
			if result := r.Data("exist"); result != "data" {
				t.Errorf(`want "%v, got "%v"`, "data", result)
			}
		})

		t.Run("not exist key", func (t *testing.T) {
			if result := r.Data("not_exist"); result != "" {
				t.Errorf(`want "%v, got "%v"`, "", result)
			}
		})
	})

	t.Run("Items", func (t *testing.T) {
		t.Run("within length", func (t *testing.T) {
			if result := r.Items(0); result.Data("items") != "data" {
				t.Errorf(`want "%v", got "%v"`, r.items[0], result)
			}
		})

		t.Run("out of length", func (t *testing.T) {
			if result := r.Items(1); result.Data("exist") != "" {
				t.Errorf(`want "%v", got "%v"`, Data{}, result)
			}
		})
	})
}