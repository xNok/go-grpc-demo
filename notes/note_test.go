package notes

import (
	"reflect"
	"testing"
)

func TestPage_save(t *testing.T) {

	lorem := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")

	type fields struct {
		Title string
		Body  []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Happy Path: save page to disk",
			fields: fields{
				Title: "a_note",
				Body:  lorem,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Note{
				Title: tt.fields.Title,
				Body:  tt.fields.Body,
			}
			if err := SaveToDisk(p, "../testdata"); (err != nil) != tt.wantErr {
				t.Errorf("Note.save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_loadPage(t *testing.T) {

	lorem := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")

	type args struct {
		keyword string
		folder  string
	}
	tests := []struct {
		name    string
		args    args
		want    *Note
		wantErr bool
	}{
		{
			name: "Happy Path: read lorem.txt",
			args: args{
				keyword: "ipsum",
				folder:  "../testdata",
			},
			want: &Note{
				Title: "lorem",
				Body:  lorem,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadFromDisk(tt.args.keyword, tt.args.folder)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
