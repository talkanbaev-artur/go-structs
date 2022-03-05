package gostructs_test

import (
	"testing"

	gostructs "github.com/talkanbaev-artur/go-structs"
	"gotest.tools/v3/assert"
)

func TestSet_Insert(t *testing.T) {
	type fields struct {
		initial []interface{}
	}
	type args struct {
		i interface{}
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		final_count int
	}{
		{
			name:        "Empty stuct",
			fields:      fields{[]interface{}{}},
			args:        args{1},
			final_count: 1,
		},
		{
			name:        "Non-empty struct",
			fields:      fields{[]interface{}{1, 2}},
			args:        args{3},
			final_count: 3,
		},
		{
			name:        "Duplicates",
			fields:      fields{[]interface{}{1, 2}},
			args:        args{2},
			final_count: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := gostructs.NewSet()
			s.Insert(tt.fields.initial...)
			s.Insert(tt.args.i)
			assert.Equal(t, s.Size(), tt.final_count)
		})
	}
}

func TestNewSet(t *testing.T) {
	s := gostructs.NewSet()
	assert.Equal(t, s.Size(), 0)
}
