package cmd

import (
	"archiver-go/lib/vlc"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcPackCmd = &cobra.Command{
	Use: "vlc",
	Short: "Pack file using variable-length code",
	Run: pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := vlc.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}

}

func packedFileName(path string) string {
	// path = /path/to/file/myFile.txt
	fileName := filepath.Base(path) // myFile.txt
	ext := filepath.Ext(fileName) // txt
	baseName := strings.TrimSuffix(fileName, ext) // 'myFile.txt' - '.txt' = 'myFile'

	return baseName + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcPackCmd)
}
