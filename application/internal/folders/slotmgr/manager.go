package slotmgr

import (
	"tde/internal/folders/copymod"
	"tde/internal/folders/discovery"
	"tde/internal/utilities"
	models "tde/models/program"

	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

type Slot string

type path = string

type slots struct {
	free     []Slot
	assigned map[models.CandidateID]Slot
}

// slot manager is to reuse existing copies of the module for next generation
type SlotManager struct {
	sample   path
	combined *discovery.CombinedDetails
	tmp      path // reserved in instantiation. all
	slots    slots
}

func New(sample path, combined *discovery.CombinedDetails) *SlotManager {
	s := SlotManager{
		sample:   sample,
		combined: combined,
		slots: slots{
			free:     []Slot{},
			assigned: map[models.CandidateID]Slot{}},
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
	basename := strings.Join(utilities.StringFold(strings.ReplaceAll(uuid.String(), "-", ""), 2), "/")
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

// empty slots have the same contents of sample module, it just doesn't have the candidate content
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

func (s *SlotManager) assignCandidateToASlot(candidate *models.Candidate) (slot Slot) {
	s.slots.free, slot = utilities.SlicePop(s.slots.free)
	s.slots.assigned[candidate.UUID] = slot
	return slot
}

func (s *SlotManager) printToFile(candidate *models.Candidate) error {
	implementationFile := filepath.Join(s.tmp,
		string(s.slots.assigned[candidate.UUID]),
		s.combined.Package.PathInModule(),
		filepath.Base(s.combined.Target.Path),
	)
	err := replaceSectionInFile(implementationFile, s.combined.Target.LineStart, s.combined.Target.LineEnd+1, candidate.File)
	if err != nil {
		return fmt.Errorf("overwriting the new part: %w", err)
	}
	return nil
}

func (s *SlotManager) placeCandidate(candidate *models.Candidate) (*Slot, error) {
	if len(s.slots.free) == 0 {
		if err := s.createEmptySlot(); err != nil {
			return nil, fmt.Errorf("could not create new slot: %w", err)
		}
	}
	slot := s.assignCandidateToASlot(candidate)
	if err := s.printToFile(candidate); err != nil {
		return nil, fmt.Errorf("printing candidate body to file in slot: %w", err)
	}
	return &slot, nil
}

func (s *SlotManager) PlaceCandidatesIntoSlots(candidates []*models.Candidate) error {
	for _, candidate := range candidates {
		if _, err := s.placeCandidate(candidate); err != nil {
			return fmt.Errorf("putting candidate into a slot: %w", err)
		}
	}
	return nil
}

func (s *SlotManager) FreeAllSlots() error {
	var slots = maps.Values(s.slots.assigned)
	s.slots.free = append(s.slots.free, slots...)
	maps.Clear(s.slots.assigned)

	org := filepath.Join(s.sample, s.combined.Package.PathInModule(), filepath.Base(s.combined.Target.Path))
	for _, slot := range slots {
		dst := filepath.Join(s.tmp, string(slot), s.combined.Package.PathInModule(), filepath.Base(s.combined.Target.Path))
		if err := copymod.CopyFile(org, dst); err != nil {
			return fmt.Errorf("restoring the target file for the slot %q: %w", slot, err)
		}
	}
	return nil
}

func (s *SlotManager) GetPackagePathForCandidate(candidateID models.CandidateID) string {
	return filepath.Join(s.tmp,
		string(s.slots.assigned[candidateID]),
		s.combined.Package.PathInModule(),
	)
}

func (s *SlotManager) Print() {
	if len(s.slots.free) > 0 {
		fmt.Println("Free: (slots)")
		for i, slot := range s.slots.free {
			fmt.Printf("%03d %s\n", i, string(slot))
		}
	}
	if len(s.slots.free) > 0 && len(s.slots.assigned) > 0 {
		fmt.Println("----")
	}
	if len(s.slots.assigned) > 0 {
		fmt.Println("Assigned: (candidate-id:slot)")
		for id, slot := range s.slots.assigned {
			fmt.Printf("%s:%s\n", id, string(slot))
		}
	}
}
