package basic

import (
	"os"
	"testing"
)

func TestParser_SetFile(t *testing.T) {
	type fields struct {
		filename  string
		file      *os.File
		configMap map[string]string
	}
	type args struct {
		filename string
	}

	type test struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	// Test 1: File exists
	//
	dumfile, _ := os.Create("/tmp/dum.txt")
	dumfile.Close()

	test1 := test{"File exists",
		fields{},
		args{"/tmp/dum.txt"},
		false,
	}

	// Test 2 file does not exist
	//
	test2 := test{"File exists",
		fields{},
		args{"/tmp/dum2.txt"},
		true,
	}

	tests := []test{
		test1, test2,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				filename:  tt.fields.filename,
				file:      tt.fields.file,
				configMap: tt.fields.configMap,
			}
			if err := p.SetFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Parser.SetFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	os.Remove("/tmp/dum.txt")
}
