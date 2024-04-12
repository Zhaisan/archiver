package cmd

import (
	"archiver-go/lib/compression"
	"archiver-go/lib/compression/vlc"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New()
	default:
		cmd.PrintErr("unknown method")
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

	packed := decoder.Decode(data)

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
	rootCmd.AddCommand(unpackCmd)

	packCmd.Flags().StringP("method", "m", "", "decompression method: vlc")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
