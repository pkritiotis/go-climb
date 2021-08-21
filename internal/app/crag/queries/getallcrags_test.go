package queries

import (
	"errors"
	"github.com/pkritiotis/go-climb/internal/domain/crag"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCragsQueryHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		repo crag.Repository
	}
	tests := []struct {
		name   string
		fields fields
		want   []CragQueryResult
		err    error
	}{
		{
			name: "happy path - no crag with no errors - should return crag",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetAll").Return([]crag.Crag{}, nil)
					return mp
				}(),
			},
			want: []CragQueryResult(nil),
			err:  nil,
		},
		{
			name: "happy path - 1 crag with no errors - should return crag",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetAll").Return([]crag.Crag{{ID: mockUUID}}, nil)
					return mp
				}(),
			},
			want: []CragQueryResult{{ID: mockUUID}},
			err:  nil,
		},
		{
			name: "get crags errors - should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetAll").Return([]crag.Crag{{ID: mockUUID}}, errors.New("get crags error"))
					return mp
				}(),
			},
			want: []CragQueryResult(nil),
			err:  errors.New("get crags error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := getAllCragsQueryHandler{
				repo: tt.fields.repo,
			}
			got, err := h.Handle()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestNewGetAllCragsQueryHandler(t *testing.T) {
	type args struct {
		repo crag.Repository
	}
	tests := []struct {
		name string
		args args
		want GetAllCragsQueryHandler
	}{
		{
			name: "should create handler",
			args: args{
				repo: crag.MockRepository{},
			},
			want: getAllCragsQueryHandler{
				repo: crag.MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGetAllCragsQueryHandler(tt.args.repo)
			assert.Equal(t, tt.want, got)
		})
	}
}