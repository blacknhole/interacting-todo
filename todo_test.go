package todo_test

import (
	"os"
	"testing"

	"pragprog.com/rggo/interacting/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

func TestCompelete(t *testing.T) {
	l := todo.List{}
	taskName := "New task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("New task should not be completed.")
	}
	l.Complete(1)
	if !l[0].Done {
		t.Errorf("New task should be completed.")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}
	for _, v := range tasks {
		l.Add(v)
	}
	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Task)
	}

	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)
	}

}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}
	taskName := "New task"
	l1.Add(taskName)
	// t.Logf("l1 :%+v\n", l1)
	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
	}
	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temporary file: %s", err)
	}

	defer os.Remove(tf.Name())
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}
	// t.Logf("l2 :%+v\n", l2)
	if l2[0].Task != l1[0].Task {
		t.Errorf("Task %q should match task %q.", l2[0].Task, l1[0].Task)
		// t.Logf("l2: %+v\nl1 :%+v", l2, l1)
	}
}
