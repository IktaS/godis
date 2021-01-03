package short

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileRepo_Save(t *testing.T) {
	type args struct {
		l *Link
	}
	tests := []struct {
		name     string
		setup    func(*testing.T) *fileRepo
		teardown func(*testing.T, *fileRepo)
		args     args
		wantErr  bool
	}{
		{
			name: "valid key pair",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "key",
					Val: "val",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid key",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "",
					Val: "val",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "",
					Val: "[val",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "",
					Val: "]val",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "",
					Val: "val[",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "",
					Val: "val]",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "",
					Val: "[val]",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "",
					Val: "|val",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid key",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "",
					Val: "val|",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "key",
					Val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "key|",
					Val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "|key",
					Val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "[key",
					Val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "]key",
					Val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "key[",
					Val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "key]",
					Val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "[key]",
					Val: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid both key val",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			args: args{
				&Link{
					Key: "",
					Val: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.setup(t)
			err := f.Save(tt.args.l)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			tt.teardown(t, f)
		})
	}
}

func TestFileRepo_Get(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*testing.T) *fileRepo
		teardown func(*testing.T, *fileRepo)
		key      string
		want     string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "find key-val",
			setup: func(t *testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				_, err = f.file.Write([]byte("key | val\n"))
				if err != nil {
					t.Fatal(err)
				}
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			key:     "key",
			want:    "val",
			wantErr: false,
		},
		{
			name: "find non existent key-val",
			setup: func(*testing.T) *fileRepo {
				f := &fileRepo{
					db:   "db_test.txt",
					file: nil,
				}
				file, err := os.OpenFile(f.db, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					t.Fatal(err)
				}
				f.file = file
				return f
			},
			teardown: func(t *testing.T, f *fileRepo) {
				e := os.Remove(f.db)
				if e != nil {
					t.Fatal(e)
				}
			},
			key:     "rando",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.setup(t)
			got, err := f.Get(tt.key)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got, tt.name)
			tt.teardown(t, f)
		})
	}
}
