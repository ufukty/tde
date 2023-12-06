package slotmgr

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"tde/internal/evolution/evaluation/copymod"
	"tde/internal/evolution/models"
	"tde/internal/utilities/slicew"
	"tde/internal/utilities/strw"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

type Slot string

type path = string

type slots struct {
	free     []Slot
	assigned map[models.Sid]Slot
}

// slot manager is to reuse existing copies of the module for next generation
type SlotManager struct {
	sample         path
	pkgPathInMod   path
	targetFilename path
	tmp            path // reserved in instantiation. all
	slots          slots
}

func New(sample, pkgPathInMod, targetFilename path) *SlotManager {
	s := SlotManager{
		sample:         sample,
		pkgPathInMod:   pkgPathInMod,
		targetFilename: targetFilename,
		slots: slots{
			free:     []Slot{},
			assigned: map[models.Sid]Slot{}},
	}
	s.createMainFolder()
	return &s
}

// returns eg. 65/36/f1/24/b8/56/5a/ad/8c/cc/22/ea/c3/7d/8e/63
func (s *SlotManager) genNewSlotPath() (path, error) {
	uuid, err := uuid.NewUUID() // UUIDv1 has choosen because the id never leaves the same-device or compared with ids produced in another device
	if err != nil {
		return "", errors.New("can't create a uuid")
	}
	basename := strings.Join(strw.Fold(strings.ReplaceAll(uuid.String(), "-", ""), 2), "/")
	return basename, nil
}

func (s *SlotManager) createMainFolder() error {
	path, err := os.MkdirTemp(os.TempDir(), "deepthinker-slots-*")
	if err != nil {
		return errors.New("failed to create main folder for slot_manager in temp directory")
	}
	s.tmp = path
	return nil
}

// empty slots have the same contents of sample module, it just doesn't have the subject content
func (s *SlotManager) createEmptySlot() error {
	newSlotPath, err := s.genNewSlotPath()
	if err != nil {
		return errors.Wrap(err, "generating slot path")
	}
	var abs = filepath.Join(s.tmp, newSlotPath)
	err = os.MkdirAll(abs, 0755)
	if err != nil {
		return fmt.Errorf("creating parent dirs for new slot: %w", err)
	}
	err = copymod.CopyModule(abs, s.sample, true, []string{}, []string{}, []string{}, false)
	if err != nil {
		return errors.Wrap(err, "copying the contents of sample module into the slot")
	}
	s.slots.free = append(s.slots.free, Slot(newSlotPath))
	return nil
}

func (s *SlotManager) assignSubjectToASlot(subject *models.Subject) (slot Slot) {
	s.slots.free, slot = slicew.Pop(s.slots.free)
	s.slots.assigned[subject.Sid] = slot
	return slot
}

func (s *SlotManager) printToFile(subject *models.Subject) error {
	implementationFile := filepath.Join(
		s.tmp,
		string(s.slots.assigned[subject.Sid]),
		s.pkgPathInMod,
		s.targetFilename,
	)
	f, err := os.Create(implementationFile)
	if err != nil {
		return fmt.Errorf("opening the target file for overwrite: %w", err)
	}
	defer f.Close()
	_, err = f.Write(subject.Code)
	if err != nil {
		return fmt.Errorf("overwriting the target file with subject: %w", err)
	}
	return nil
}

func (s *SlotManager) placeSubject(subj *models.Subject) (*Slot, error) {
	if len(s.slots.free) == 0 {
		if err := s.createEmptySlot(); err != nil {
			return nil, fmt.Errorf("could not create new slot: %w", err)
		}
	}
	slot := s.assignSubjectToASlot(subj)
	if err := s.printToFile(subj); err != nil {
		return nil, fmt.Errorf("printing subject body to file in slot: %w", err)
	}
	return &slot, nil
}

func (s *SlotManager) PlaceSubjectsIntoSlots(subjects models.Subjects) error {
	for _, subj := range subjects {
		if _, err := s.placeSubject(subj); err != nil {
			return fmt.Errorf("putting subject into a slot: %w", err)
		}
	}
	return nil
}

func (s *SlotManager) FreeAllSlots() error {
	var slots = maps.Values(s.slots.assigned)
	s.slots.free = append(s.slots.free, slots...)
	maps.Clear(s.slots.assigned)

	org := filepath.Join(
		s.sample,
		s.pkgPathInMod,
		s.targetFilename,
	)
	for _, slot := range slots {
		dst := filepath.Join(s.tmp, string(slot), s.pkgPathInMod, s.targetFilename)
		if err := copymod.CopyFile(org, dst); err != nil {
			return fmt.Errorf("restoring the target file for the slot %q: %w", slot, err)
		}
	}
	return nil
}

func (s *SlotManager) GetPackagePathForSubject(sid models.Sid) string {
	return filepath.Join(
		s.tmp,
		string(s.slots.assigned[sid]),
		s.pkgPathInMod,
	)
}

func (s *SlotManager) Print() {
	fmt.Println("Printing SlotManager details:")
	fmt.Printf("  Sample dir: %s\n", s.sample)
	fmt.Printf("  Temp dir  : %s\n", s.tmp)
	if len(s.slots.free) > 0 {
		fmt.Println("  Free      : (index:slots)")
		for i, slot := range s.slots.free {
			fmt.Printf("    %03d %s\n", i, string(slot))
		}
	}
	if len(s.slots.free) > 0 && len(s.slots.assigned) > 0 {
		fmt.Println()
	}
	if len(s.slots.assigned) > 0 {
		fmt.Println("  Assigned  : (sid:slot)")
		for id, slot := range s.slots.assigned {
			fmt.Printf("    %s:%s\n", id, string(slot))
		}
	}
}
