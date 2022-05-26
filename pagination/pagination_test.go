package pagination

import (
	"testing"
)

func TestPagination_GetOffset(t *testing.T) {
	type fields struct {
		Limit      int
		Page       int
		TotalRows  int64
		TotalPages int
		Rows       interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "first-page",
			fields: fields{
				Limit: 10,
				Page:  1,
			},
			want: 0,
		},
		{
			name: "second-page",
			fields: fields{
				Limit: 10,
				Page:  2,
			},
			want: 10,
		},
		{
			name: "third-page",
			fields: fields{
				Limit: 10,
				Page:  3,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pagination{
				Limit:      tt.fields.Limit,
				Page:       tt.fields.Page,
				TotalRows:  tt.fields.TotalRows,
				TotalPages: tt.fields.TotalPages,
				Rows:       tt.fields.Rows,
			}
			if got := p.GetOffset(); got != tt.want {
				t.Errorf("Pagination.GetOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_GetLimit(t *testing.T) {
	type fields struct {
		Limit      int
		Page       int
		TotalRows  int64
		TotalPages int
		Rows       interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "default-limit",
			fields: fields{},
			want:   10,
		},
		{
			name:   "custom-limit-0",
			fields: fields{},
			want:   10,
		},
		{
			name: "custom-limit-20",
			fields: fields{
				Limit: 20,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pagination{
				Limit:      tt.fields.Limit,
				Page:       tt.fields.Page,
				TotalRows:  tt.fields.TotalRows,
				TotalPages: tt.fields.TotalPages,
				Rows:       tt.fields.Rows,
			}
			if got := p.GetLimit(); got != tt.want {
				t.Errorf("Pagination.GetLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_GetPage(t *testing.T) {
	type fields struct {
		Limit      int
		Page       int
		TotalRows  int64
		TotalPages int
		Rows       interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "empty-page",
			fields: fields{},
			want:   1,
		},
		{
			name: "custom-page-0",
			fields: fields{
				Page: 0,
			},
			want: 1,
		},
		{
			name: "custom-page-5",
			fields: fields{
				Page: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pagination{
				Limit:      tt.fields.Limit,
				Page:       tt.fields.Page,
				TotalRows:  tt.fields.TotalRows,
				TotalPages: tt.fields.TotalPages,
				Rows:       tt.fields.Rows,
			}
			if got := p.GetPage(); got != tt.want {
				t.Errorf("Pagination.GetPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
