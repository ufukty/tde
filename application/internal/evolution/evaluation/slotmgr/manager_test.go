package slotmgr

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	models "tde/models/program"

	"github.com/google/uuid"
)

func getTestPackage() (mod string, pkg *list.Package, err error) {
	if pkgs, err := list.ListPackagesInDir("testdata/words"); err != nil {
		return "", nil, fmt.Errorf("listing the test package: %w", err)
	} else {
		return pkgs.First().Module.Dir, pkgs.First(), nil
	}
}

func Test_CodePlacement(t *testing.T) {
	mod, pkg, err := getTestPackage()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	sample, err := inject.WithCreatingSample(mod, pkg, "TDE_WordReverse")
	if err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}

	subjects := models.Subjects{}
	subjects.Add(&models.Subject{
		Sid:  "00000000-0000-0000-0000-000000000000",
		Code: []byte("Hello world"),
	})

	var sm = New(sample, pkg.PathInModule(), "words.go")
	if err := sm.PlaceSubjectsIntoSlots(subjects); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}

	file := filepath.Join(sm.GetPackagePathForSubject("00000000-0000-0000-0000-000000000000"), "words.go")
	content, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(fmt.Errorf("assert-prep-2: %w", err))
	}
	if string(content) != "Hello world" {
		t.Fatal(fmt.Errorf("assert: expected %q, got %q", "Hello world", string(content)))
	}
}

func Test_AssignAndFree(t *testing.T) {
	mod, pkg, err := getTestPackage()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	sample, err := inject.WithCreatingSample(mod, pkg, "TDE_WordReverse")
	if err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}
	fmt.Println("sample module dir:", sample)

	var subjects = models.Subjects{}
	for i := 0; i < 10; i++ {
		subjects.Add(&models.Subject{
			Sid:  models.Sid(uuid.New().String()),
			Code: []byte(`hello world`),
		})
	}

	var sm = New(sample, pkg.PathInModule(), "words.go")
	if err := sm.PlaceSubjectsIntoSlots(subjects); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}

	fmt.Println("slot manager tmp:", sm.tmp)
	sm.Print()

	if len(sm.slots.assigned) < 10 {
		t.Fatal("assert. less then 10 slots assigned")
	}
	if len(sm.slots.free) != 0 {
		t.Fatal("assert. there should be no free slots")
	}

	sm.FreeAllSlots()
	sm.Print()

	if len(sm.slots.assigned) != 0 {
		t.Fatal("assert. there should be no assigned slots")
	}
	if len(sm.slots.free) < 10 {
		t.Fatal("assert. less then 10 slots freed")
	}
}
