package slotmgr

import (
	"fmt"
	"tde/internal/folders/copymod"
	"tde/internal/folders/types"
	"tde/internal/utilities"
	models "tde/models/program"

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
	modulePath  path
	testDetails *types.TestDetails
	tmp         path // reserved in instantiation. all
	slots       slots
}

func New(modulePath string, testDetails *types.TestDetails) *SlotManager {
	s := SlotManager{
		modulePath:  modulePath,
		testDetails: testDetails,
		slots: slots{
			free:     []Slot{},
			assigned: map[models.CandidateID]Slot{}},
	}
	s.createMainFolder()
	return &s
}

// returns eg. 65/36/f1/24/b8/56/5a/ad/8c/cc/22/ea/c3/7d/8e/63
func (s *SlotManager) genNewSlotPath() (path, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("can't create a uuid")
	}
	basename := strings.Join(utilities.StringFold(strings.ReplaceAll(uuid.String(), "-", ""), 2), "/")
	return basename, nil
}

func (s *SlotManager) createMainFolder() error {
	path, err := os.MkdirTemp(os.TempDir(), "tde.runner-folders.*")
	if err != nil {
		return errors.New("failed to create main folder for slot_manager in temp directory")
	}
	s.tmp = path
	return nil
}

func (s *SlotManager) createModuleDuplicate() error {
	newSlotPath, err := s.genNewSlotPath()
	if err != nil {
		return errors.Wrap(err, "genNewstring")
	}

	// fmt.Println("Original module duplicated:", path)
	err = copymod.Copy(s.modulePath, filepath.Join(s.tmp, newSlotPath), true, []string{}, []string{}, []string{}, false)
	if err != nil {
		return errors.Wrap(err, "copy_module.Module")
	}

	s.slots.free = append(s.slots.free, Slot(newSlotPath))
	return nil
}

func (s *SlotManager) assignCandidateToASlot(candidateID models.CandidateID) (slot Slot) {
	s.slots.free, slot = utilities.SlicePop(s.slots.free)
	s.slots.assigned[candidateID] = slot
	// fmt.Println("assigning", candidateID, "to folder", choosen)
	return
}

func (s *SlotManager) printToFile(candidate *models.Candidate) error {
	implementationFile := filepath.Join(s.tmp,
		string(s.slots.assigned[candidate.UUID]),
		s.testDetails.ImplFuncFile,
	)
	f, err := os.Create(implementationFile)
	if err != nil {
		return errors.Wrap(err, "open implementation file to overwrite")
	}
	defer f.Close()
	_, err = f.Write(candidate.File)
	if err != nil {
		return errors.Wrap(err, "Write")
	}
	return nil
}

func (s *SlotManager) placeCandidate(candidate *models.Candidate) {
	if len(s.slots.free) == 0 {
		s.createModuleDuplicate()
	}
	s.assignCandidateToASlot(candidate.UUID)
	s.printToFile(candidate)
}

func (s *SlotManager) PlaceCandidatesIntoSlots(candidates []*models.Candidate) {
	for _, candidate := range candidates {
		s.placeCandidate(candidate)
	}
}

func (s *SlotManager) FreeAllSlots() {
	s.slots.free = append(s.slots.free, maps.Values(s.slots.assigned)...)
}

func (s *SlotManager) GetPackagePathForCandidate(candidateID models.CandidateID) string {
	return filepath.Join(s.tmp,
		string(s.slots.assigned[candidateID]),
		s.testDetails.PackagePath,
	)
}

func (s *SlotManager) Print() {
	fmt.Println("Free: (slots)")
	for i, slot := range s.slots.free {
		fmt.Printf("%03d %s\n", i, string(slot))
	}
	fmt.Println("----")
	fmt.Println("Assigned: (candidate-id:slot)")
	for id, slot := range s.slots.assigned {
		fmt.Printf("%s %s\n", id, string(slot))
	}
	fmt.Println("----")
}
