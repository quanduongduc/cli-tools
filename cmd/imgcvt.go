package cmd

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(imgcvtCmd)
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

var imgcvtCmd = &cobra.Command{
  Use:   "imgcvt [image path]",
  Short: "Converts an image to JPEG format",
  Long:  `imgcvt is a CLI tool that converts an image to JPEG format and deletes the original file if the conversion is successful.`,
  Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    err := convertToJPG(args[0])
    if err != nil {
      fmt.Println("Error converting image:", err)
      os.Exit(1)
    }
    fmt.Println("Image converted successfully")
  },
}

func convertToJPG(inputPath string) error {
    originalFile, err := os.Open(inputPath)
    if err != nil {
        return err
    }
    defer originalFile.Close()

    img, _, err := image.Decode(originalFile)
    if err != nil {
        return err
    }

    resizedImg := resize.Resize(0, 0, img, resize.Lanczos3)

    newPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".jpg"
    newFile, err := os.Create(newPath)
    if err != nil {
        return err
    }
    defer newFile.Close()

    // Encode the image to JPEG and save it
    err = jpeg.Encode(newFile, resizedImg, &jpeg.Options{Quality: jpeg.DefaultQuality})
    if err != nil {
        return err
    }

    // Delete the original file
    err = os.Remove(inputPath)
    if err != nil {
        return err
    }

    return nil
}