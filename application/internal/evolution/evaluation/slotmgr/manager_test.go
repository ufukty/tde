package slotmgr

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"tde/internal/evolution/evaluation/discovery"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	models "tde/models/program"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

func Test_SlotManager_AssignAndFree(t *testing.T) {
	var (
		pkgs   list.Packages
		sample string
		err    error
	)
	if pkgs, err = list.ListPackagesInDir("testdata/words"); err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	var pkg = pkgs.First()
	if sample, err = inject.WithCreatingSample(pkg.Module.Dir, pkg, "TDE_WordReverse"); err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}
	fmt.Println("sample module dir:", sample)

	var candidates = []*models.Candidate{}
	for i := 0; i < 10; i++ {
		candidates = append(candidates, &models.Candidate{
			UUID: models.CandidateID(uuid.New().String()),
			File: []byte(`hello world`),
		})
	}

	combined, err := discovery.CombinedDetailsForTest("testdata/words", "TDE_WordReverse")
	if err != nil {
		t.Fatal(fmt.Errorf("prep, getting combined details: %w", err))
	}

	var sm = New(sample, combined)
	if err := sm.PlaceCandidatesIntoSlots(candidates); err != nil {
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

func checksum(path string) (string, error) {
	fh, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("opening file: %w", err)
	}
	defer fh.Close()
	hash := md5.New()
	_, err = io.Copy(hash, fh)
	if err != nil {
		return "", errors.Wrap(err, "Error calculating MD5 checksum")
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func Test_SlotManager_ComparingTargetFileAfterAssignAndFree(t *testing.T) {
	var (
		pkgs   list.Packages
		sample string
		err    error
	)
	if pkgs, err = list.ListPackagesInDir("testdata/words"); err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	var pkg = pkgs.First()
	if sample, err = inject.WithCreatingSample(pkg.Module.Dir, pkg, "TDE_WordReverse"); err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}

	fmt.Println("sample module dir:", sample)

	var candidates = []*models.Candidate{
		{
			UUID: models.CandidateID(uuid.New().String()),
			File: []byte(""),
		},
	}
	combined, err := discovery.CombinedDetailsForTest("testdata/words", "TDE_WordReverse")
	if err != nil {
		t.Fatal(fmt.Errorf("prep, getting combined details: %w", err))
	}

	originalHash, err := checksum(filepath.Join(sample, combined.Package.PathInModule(), filepath.Base(combined.Target.Path)))
	if err != nil {
		t.Fatal(fmt.Errorf("prep, hash original target file: %w", err))
	}

	var sm = New(sample, combined)
	if err := sm.PlaceCandidatesIntoSlots(candidates); err != nil {
		t.Fatal(fmt.Errorf("act 1: %w", err))
	}

	var assignedSlot = string(maps.Values(sm.slots.assigned)[0])
	var targetFileInSlot = filepath.Join(sm.tmp, assignedSlot, combined.Package.PathInModule(), filepath.Base(combined.Target.Path))
	fmt.Println("target file path in slot:", targetFileInSlot)
	modifiedHash, err := checksum(targetFileInSlot)
	if err != nil {
		t.Fatal(fmt.Errorf("assert 1 prep, hash assigned target file: %w", err))
	}

	if originalHash == modifiedHash {
		t.Fatal(fmt.Errorf("assert 1, got same hashes for original file and assigned slot"))
	}

	err = sm.FreeAllSlots()
	if err != nil {
		t.Fatal(fmt.Errorf("act 2, freeing slots: %w", err))
	}
	restoredHash, err := checksum(targetFileInSlot)
	if err != nil {
		t.Fatal(fmt.Errorf("assert 2 prep, hash freed target file: %w", err))
	}

	if originalHash != restoredHash {
		t.Fatal(fmt.Errorf("assert 2, got different hashes for original file and restored slot"))
	}

}
