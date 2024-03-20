package cmd

import (
	"archiver-go/lib/vlc"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
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

	packed := vlc.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}

}

func unpackedFileName(path string) string {
	// path = /path/to/file/myFile.txt
	fileName := filepath.Base(path)               // myFile.txt
	ext := filepath.Ext(fileName)                 // txt
	baseName := strings.TrimSuffix(fileName, ext) // 'myFile.txt' - '.txt' = 'myFile'

	return baseName + "." + unpackedExtension
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
