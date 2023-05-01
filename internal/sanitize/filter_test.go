package sanitize

import (
	"testing"

	"github.com/CAMELNINGA/cdc-postgres.git/internal/models"
	"gotest.tools/v3/assert"
)

func TestPositiveFilterHandler(t *testing.T) {

	filterHandeler := &FilterHandler{
		Table: "test",
		filterColumns: map[string]string{
			"test": "test",
		},
	}

	tests := []struct {
		names   string
		in      *models.ActionData
		out     *models.ActionData
		wantOut bool
	}{
		{
			names: "basick test",
			in: &models.ActionData{
				Kind:  models.ActionKindInsert,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
			},
			out: &models.ActionData{
				Kind:  models.ActionKindInsert,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test2", Value: "test2"},
				},
			},
			wantOut: true,
		},
		{
			names: "basick test another table",
			in: &models.ActionData{
				Kind:  models.ActionKindInsert,
				Table: "test1",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
			},
			out: &models.ActionData{
				Kind:  models.ActionKindInsert,
				Table: "test1",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
			},
			wantOut: true,
		},
		{
			names: "full clear insert test",
			in: &models.ActionData{
				Kind:  models.ActionKindInsert,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
				},
			},
			wantOut: false,
		},
		{
			names: "full update test",
			in: &models.ActionData{
				Kind:  models.ActionKindUpdate,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
				OldColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
			},
			out: &models.ActionData{
				Kind:  models.ActionKindUpdate,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test2", Value: "test2"},
				},
				OldColumns: []models.Column{
					{Name: "test2", Value: "test2"},
				},
			},
			wantOut: true,
		},
		{
			names: "full clear update test",
			in: &models.ActionData{
				Kind:  models.ActionKindUpdate,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
				},
				OldColumns: []models.Column{
					{Name: "test", Value: "test"},
				},
			},
			wantOut: false,
		},
		{
			names: "full delete test",
			in: &models.ActionData{
				Kind:  models.ActionKindDelete,
				Table: "test",
				OldColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
			},
			out: &models.ActionData{
				Kind:  models.ActionKindDelete,
				Table: "test",
				OldColumns: []models.Column{
					{Name: "test2", Value: "test2"},
				},
			},
			wantOut: true,
		},
		{
			names: "full clear delete test",
			in: &models.ActionData{
				Kind:  models.ActionKindDelete,
				Table: "test",
				OldColumns: []models.Column{
					{Name: "test", Value: "test"},
				},
			},
			wantOut: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.names, func(t *testing.T) {
			got := filterHandeler.Handle(tt.in)
			if tt.wantOut {
				assert.Equal(t, got.Table, tt.out.Table)
				switch got.Kind {
				case models.ActionKindInsert:

					for i, v := range got.NewColumns {
						assert.Equal(t, v, tt.out.NewColumns[i])

					}
				case models.ActionKindUpdate:
					for i, v := range got.NewColumns {
						assert.Equal(t, v, tt.out.NewColumns[i])

					}
					for i, v := range got.OldColumns {
						assert.Equal(t, v, tt.out.OldColumns[i])

					}
				case models.ActionKindDelete:
					for i, v := range got.OldColumns {
						assert.Equal(t, v, tt.out.OldColumns[i])

					}

				}

			} else {
				assert.Assert(t, got == nil)
			}
		})
	}

}

func TestAllTAblesFoilter(t *testing.T) {
	filterHandeler := &FilterHandler{
		Table: "*",
		filterColumns: map[string]string{
			"test": "test",
		},
	}
	tests := []struct {
		names   string
		in      *models.ActionData
		out     *models.ActionData
		wantOut bool
	}{
		{
			names: "basick test",
			in: &models.ActionData{
				Kind:  models.ActionKindInsert,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
			},
			out: &models.ActionData{
				Kind:  models.ActionKindInsert,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test2", Value: "test2"},
				},
			},
			wantOut: true,
		},
		{
			names: "basick test another table",
			in: &models.ActionData{
				Kind:  models.ActionKindInsert,
				Table: "test1",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
			},
			out: &models.ActionData{
				Kind:  models.ActionKindInsert,
				Table: "test1",
				NewColumns: []models.Column{
					{Name: "test2", Value: "test2"},
				},
			},
			wantOut: true,
		},
		{
			names: "full clear insert test",
			in: &models.ActionData{
				Kind:  models.ActionKindInsert,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
				},
			},
			wantOut: false,
		},
		{
			names: "full update test",
			in: &models.ActionData{
				Kind:  models.ActionKindUpdate,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
				OldColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
			},
			out: &models.ActionData{
				Kind:  models.ActionKindUpdate,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test2", Value: "test2"},
				},
				OldColumns: []models.Column{
					{Name: "test2", Value: "test2"},
				},
			},
			wantOut: true,
		},
		{
			names: "full clear update test",
			in: &models.ActionData{
				Kind:  models.ActionKindUpdate,
				Table: "test",
				NewColumns: []models.Column{
					{Name: "test", Value: "test"},
				},
				OldColumns: []models.Column{
					{Name: "test", Value: "test"},
				},
			},
			wantOut: false,
		},
		{
			names: "full delete test",
			in: &models.ActionData{
				Kind:  models.ActionKindDelete,
				Table: "test",
				OldColumns: []models.Column{
					{Name: "test", Value: "test"},
					{Name: "test2", Value: "test2"},
				},
			},
			out: &models.ActionData{
				Kind:  models.ActionKindDelete,
				Table: "test",
				OldColumns: []models.Column{
					{Name: "test2", Value: "test2"},
				},
			},
			wantOut: true,
		},
		{
			names: "full clear delete test",
			in: &models.ActionData{
				Kind:  models.ActionKindDelete,
				Table: "test",
				OldColumns: []models.Column{
					{Name: "test", Value: "test"},
				},
			},
			wantOut: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.names, func(t *testing.T) {
			got := filterHandeler.Handle(tt.in)
			if tt.wantOut {
				assert.Equal(t, got.Table, tt.out.Table)
				switch got.Kind {
				case models.ActionKindInsert:

					for i, v := range got.NewColumns {
						assert.Equal(t, v, tt.out.NewColumns[i])

					}
				case models.ActionKindUpdate:
					for i, v := range got.NewColumns {
						assert.Equal(t, v, tt.out.NewColumns[i])

					}
					for i, v := range got.OldColumns {
						assert.Equal(t, v, tt.out.OldColumns[i])

					}
				case models.ActionKindDelete:
					for i, v := range got.OldColumns {
						assert.Equal(t, v, tt.out.OldColumns[i])

					}

				}

			} else {
				assert.Assert(t, got == nil)
			}
		})
	}
}
