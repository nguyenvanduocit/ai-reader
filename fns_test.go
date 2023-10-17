package main

import (
	"testing"
)

func TestImportBook(t *testing.T) {
	type args struct {
		srcPath     string
		destBaseDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				srcPath:     "testdata/epub/book.epub",
				destBaseDir: "testdata/out",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ImportBook(tt.args.srcPath, tt.args.destBaseDir); (err != nil) != tt.wantErr {
				t.Errorf("ImportBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateDestFullPath(t *testing.T) {
	type args struct {
		srcPath     string
		destBaseDir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				srcPath:     "testdata/epub/book.epub",
				destBaseDir: "testdata/out",
			},
			want: "testdata/out/book",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateDestFullPath(tt.args.srcPath, tt.args.destBaseDir); got != tt.want {
				t.Errorf("GenerateDestFullPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnzipToDest(t *testing.T) {
	type args struct {
		zipPath  string
		destPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				zipPath:  "testdata/epub/book.epub",
				destPath: "testdata/out/book",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnzipToDest(tt.args.zipPath, tt.args.destPath); (err != nil) != tt.wantErr {
				t.Errorf("UnzipToDest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseEpubContainer(t *testing.T) {
	type args struct {
		epubDirPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				epubDirPath: "testdata/out/book",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseEpubContainer(tt.args.epubDirPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseEpubContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetPackFilePath(t *testing.T) {
	type args struct {
		epubDirPath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				epubDirPath: "testdata/out/book",
			},
			want:    "testdata/out/book/EPUB/package.opf",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPackFilePath(tt.args.epubDirPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPackFilePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPackFilePath() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseEpubPackageData(t *testing.T) {
	type args struct {
		packFilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				packFilePath: "testdata/out/book/EPUB/package.opf",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseEpubPackageData(tt.args.packFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseEpubPackageData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
