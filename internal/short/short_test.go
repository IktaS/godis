package short

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileRepo_Init(t *testing.T) {
	type fields struct {
		db   string
		file *os.File
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "db file exist",
			fields: fields{
				db:   "db_test.txt",
				file: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileRepo{
				db:   tt.fields.db,
				file: tt.fields.file,
			}
			if err := f.Init(); (err != nil) != tt.wantErr {
				t.Errorf("FileRepo.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileRepo_Save(t *testing.T) {
	type fields struct {
		db   string
		file *os.File
	}
	type args struct {
		l *Link
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid key pair",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "key",
					val: "val",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid key",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "",
					val: "val",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "",
					val: "[val",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "",
					val: "]val",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "",
					val: "val[",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "",
					val: "val]",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "",
					val: "[val]",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "",
					val: "|val",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "",
					val: "val|",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "key",
					val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "key|",
					val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "|key",
					val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "[key",
					val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "]key",
					val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "key[",
					val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "key]",
					val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "[key]",
					val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid both key val",
			fields: fields{
				"db_test.txt",
				nil,
			},
			args: args{
				&Link{
					key: "",
					val: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileRepo{
				db:   tt.fields.db,
				file: tt.fields.file,
			}
			f.Init()
			err := f.Save(tt.args.l)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestFileRepo_Get(t *testing.T) {
	type fields struct {
		db   string
		file *os.File
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "find key-val",
			fields: fields{
				db:   "db_test.txt",
				file: nil,
			},
			args: args{
				key: "key",
			},
			want:    "val",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileRepo{
				db:   tt.fields.db,
				file: tt.fields.file,
			}
			f.Init()
			got, err := f.Get(tt.args.key)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("FileRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if got != tt.want {
			// 	t.Errorf("FileRepo.Get() = %v, want %v", got, tt.want)
			// }
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}
