package creational

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFileClone(t *testing.T) {
	originalFile := &File{Name: "OriginalFile"}
	cloneFile := originalFile.Clone().(*File)

	if cloneFile.Name != "OriginalFile_Clone" {
		t.Errorf("Expected clone name 'OriginalFile_Clone', but got '%s'", cloneFile.Name)
	}

	if originalFile == cloneFile {
		t.Error("Expected original and clone to be distinct objects")
	}
}

func TestFolderClone(t *testing.T) {
	originalFile1 := &File{Name: "File1"}
	originalFile2 := &File{Name: "File2"}
	originalFolder := &Folder{
		Name:     "OriginalFolder",
		Children: []Inode{originalFile1, originalFile2},
	}

	cloneFolder := originalFolder.Clone().(*Folder)

	if cloneFolder.Name != "OriginalFolder_Clone" {
		t.Errorf("Expected clone folder name 'OriginalFolder_Clone', but got '%s'", cloneFolder.Name)
	}

	if len(cloneFolder.Children) != 2 {
		t.Errorf("Expected 2 children in the cloned folder, but got %d", len(cloneFolder.Children))
	}

	if cloneFolder.Children[0].(*File).Name != "File1_Clone" || cloneFolder.Children[1].(*File).Name != "File2_Clone" {
		t.Errorf("Expected cloned children names 'File1_Clone' and 'File2_Clone', but got '%s' and '%s'",
			cloneFolder.Children[0].(*File).Name, cloneFolder.Children[1].(*File).Name)
	}

	if originalFolder == cloneFolder {
		t.Error("Expected original and clone folders to be distinct objects")
	}

	if originalFolder.Children[0] == cloneFolder.Children[0] || originalFolder.Children[1] == cloneFolder.Children[1] {
		t.Error("Expected the children of the original and cloned folder to be distinct objects")
	}
}

func TestPrint(t *testing.T) {
	originalFile := &File{Name: "File1"}
	originalFolder := &Folder{
		Name:     "Folder1",
		Children: []Inode{originalFile},
	}

	var buf bytes.Buffer
	originalFolder.Print("  ")
	expectedOutput := "Folder1\n  File1\n"

	fmt.Fprint(&buf, originalFolder)
	if buf.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nBut got:\n%s", expectedOutput, buf.String())
	}
}
