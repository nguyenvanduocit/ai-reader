package main

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GenerateBookIDFromPath(srcPath string) string {
	ebookFileName := path.Base(srcPath)
	nameWithoutExt := strings.TrimSuffix(ebookFileName, path.Ext(ebookFileName))
	return nameWithoutExt
}

func GenerateDestFullPath(srcPath string, destBaseDir string) string {
	bookID := GenerateBookIDFromPath(srcPath)
	destPath := path.Join(destBaseDir, bookID)
	return destPath
}

func UnzipToDest(zipPath string, destPath string) error {
	archive, err := zip.OpenReader(zipPath)
	if err != nil {
		return errors.WithMessage(err, "failed to open zip file")
	}
	defer archive.Close()
	for _, f := range archive.File {
		filePath := filepath.Join(destPath, f.Name)
		if !strings.HasPrefix(filePath, filepath.Clean(destPath)+string(os.PathSeparator)) {
			return errors.Errorf("%s: illegal file path", filePath)
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return errors.WithMessage(err, "failed to create directory")
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return errors.WithMessage(err, "failed to open file")
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return errors.WithMessage(err, "failed to open file in archive")
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return errors.WithMessage(err, "failed to copy file")
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	return nil
}

func ImportBook(srcPath string, bookBasePath string) (string, error) {
	destPath := GenerateDestFullPath(srcPath, bookBasePath)
	if err := UnzipToDest(srcPath, destPath); err != nil {
		return "", errors.WithMessage(err, "failed to unzip file")
	}

	return destPath, nil
}

func IndexBook(epubDirPath string) error {
	// generate cache data for this book, save under the book directory
	containerFilePath := GetContainerFilePath(epubDirPath)
	container, err := ParseEpubContainer(containerFilePath)
	if err != nil {
		return errors.WithMessage(err, "failed to parse epub container")
	}

	packFilePath := path.Join(epubDirPath, container.Rootfiles.Rootfile.FullPath)

	pack, err := ParseEpubPackageData(packFilePath)
	if err != nil {
		return errors.WithMessage(err, "failed to parse epub package")
	}

	if pack == nil {
		return errors.New("pack is nil")
	}

	epubRootDir := path.Dir(container.Rootfiles.Rootfile.FullPath)

	indexData, err := GenerateBookIndexData(pack, epubRootDir)
	if err != nil {
		return errors.WithMessage(err, "failed to generate index data")
	}

	// write cache data to json index.json
	if err := SaveIndexFile(indexData, epubDirPath); err != nil {
		return errors.WithMessage(err, "failed to save index file")
	}

	return nil
}

func GenerateBookIndexData(packData *EpubPackXml, baseDir string) (*BookIndexData, error) {
	indexData := &BookIndexData{
		Title:  packData.Metadata.Title,
		Author: packData.Metadata.Creator.Text,
	}

	for _, item := range packData.Spine.Itemref {
		for _, manifestItem := range packData.Manifest.Item {
			if item.Idref == manifestItem.ID {
				indexData.FileList = append(indexData.FileList, path.Join(baseDir, manifestItem.Href))
			}
		}
	}

	return indexData, nil
}

func SaveIndexFile(indexData *BookIndexData, epubDirPath string) error {
	payload, _ := json.MarshalIndent(indexData, "", "  ")
	indexFilePath := path.Join(epubDirPath, "META-INF", "index.json")
	if err := os.WriteFile(indexFilePath, payload, 0644); err != nil {
		return errors.WithMessage(err, "failed to write index file")
	}

	return nil
}

// GetContainerFilePath returns the path to the container.xml file in the epub directory
// epubDirPath is the path to the directory containing the META-INF directory
func GetContainerFilePath(epubDirPath string) string {
	return path.Join(epubDirPath, "META-INF", "container.xml")
}

// ParseEpubContainer parses the container.xml file in the epub directory
// epubDirPath is the path to the directory containing the META-INF directory
func ParseEpubContainer(containerFilePath string) (*EpubContainerXml, error) {
	containerFile, err := os.Open(containerFilePath)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to open container file")
	}
	defer containerFile.Close()

	xmlDecoder := xml.NewDecoder(containerFile)
	var container EpubContainerXml
	if err := xmlDecoder.Decode(&container); err != nil {
		return nil, errors.WithMessage(err, "failed to decode container file")
	}

	return &container, nil
}

// ParseEpubPackageData parses the content.opf file in the epub directory
// packFilePath is the path to the content.opf file
func ParseEpubPackageData(packFilePath string) (*EpubPackXml, error) {
	packFile, err := os.Open(packFilePath)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to open package file")
	}
	defer packFile.Close()

	xmlDecoder := xml.NewDecoder(packFile)
	var pack EpubPackXml
	if err := xmlDecoder.Decode(&pack); err != nil {
		return nil, errors.WithMessage(err, "failed to decode package file")
	}

	return &pack, nil
}
