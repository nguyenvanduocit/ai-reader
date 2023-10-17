package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shibukawa/configdir"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"path"
)

// App struct
type App struct {
	ctx         context.Context
	dataDirPath string
	bookDirPath string
}

func NewApp() *App {

	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	dataDir := configdir.New("aiocean", "epub-reader")
	firstDir := dataDir.QueryFolders(configdir.Global)[0]
	if err := firstDir.MkdirAll(); err != nil {
		runtime.LogErrorf(a.ctx, "Failed to create data directory: %s", err)
		return
	}

	a.dataDirPath = firstDir.Path
	a.bookDirPath = path.Join(a.dataDirPath, "books")
	runtime.LogDebugf(a.ctx, "Data directory: %s", a.dataDirPath)
}

// SetupDataDirectory check if the data directory exists and if not, create it
func (a *App) SetupDataDirectory() error {
	// check if the data directory exists

	return nil
}

func (a *App) ImportBook(epubPath string) string {
	savedBookPath, err := ImportBook(epubPath, a.bookDirPath)
	if err != nil {
		return fmt.Sprintf("Failed to import %s; err: %s", epubPath, err.Error())
	}

	if err := IndexBook(savedBookPath); err != nil {
		return fmt.Sprintf("Failed to index %s; err: %s", epubPath, err.Error())
	}

	return fmt.Sprintf("Imported %s", epubPath)
}

func (a *App) OpenEpubDialog() (string, error) {
	filepath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Open Epub",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Epub Files (*.epub)",
				Pattern:     "*.epub",
			},
		},
	})

	if err != nil {
		return "", errors.WithMessage(err, "failed to open epub dialog")
	}

	return filepath, nil
}
