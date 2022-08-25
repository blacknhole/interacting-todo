package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	fmt.Println("Building tool...")
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()
	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "Test task number 1"
	taskNum := "1"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-task", task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		exp := fmt.Sprintf("  1: %s\n", task)
		if string(out) != exp {
			t.Errorf("Expected %q, got %q instead\n", exp, string(out))
		}
	})

	t.Run("CompleteItem", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", taskNum)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteItem", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-delete", taskNum)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
}
