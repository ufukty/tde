package slot_manager

import (
	"path/filepath"
	"tde/internal/folders/copy_module"
	"tde/internal/folders/types"
	"tde/internal/utilities"
	"tde/models"

	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

type slots struct {
	free     []types.SlotPath
	assigned map[models.CandidateID]types.SlotPath
}

type Session struct {
	modulePath  types.AbsolutePath
	testDetails *types.TestDetails
	tmp         types.TempPath // reserved in instantiation. all
	slots       slots
}

// returns eg. 65/36/f1/24/b8/56/5a/ad/8c/cc/22/ea/c3/7d/8e/63
func (s *Session) genNewSlotPath() (types.SlotPath, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("can't create a uuid")
	}
	basename := strings.Join(utilities.StringFold(strings.ReplaceAll(uuid.String(), "-", ""), 2), "/")
	return types.SlotPath(basename), nil
}

func (s *Session) createMainFolder() error {
	path, err := os.MkdirTemp(os.TempDir(), "tde.runner-folders.*")
	if err != nil {
		return errors.New("failed to create main folder for slot_manager in temp directory")
	}
	s.tmp = types.TempPath(path)
	return nil
}

func (s *Session) createModuleDuplicate() error {
	newSlotPath, err := s.genNewSlotPath()
	if err != nil {
		return errors.Wrap(err, "genNewtypes.SlotPath")
	}

	// fmt.Println("Original module duplicated:", path)
	err = copy_module.Module(string(s.modulePath), string(s.tmp.Join(newSlotPath)), true, copy_module.DefaultSkipDirs)
	if err != nil {
		return errors.Wrap(err, "copy_module.Module")
	}

	s.slots.free = append(s.slots.free, newSlotPath)
	return nil
}

func (s *Session) assignCandidateToASlot(candidateID models.CandidateID) (slot types.SlotPath) {
	s.slots.free, slot = utilities.SlicePop(s.slots.free)
	s.slots.assigned[candidateID] = slot
	// fmt.Println("assigning", candidateID, "to folder", choosen)
	return
}

func (s *Session) printToFile(candidate *models.Candidate) error {
	slot := s.slots.assigned[candidate.UUID]
	implementationFile := s.tmp.FindInModulePath(slot, s.testDetails.ImplFuncFile)
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

func (s *Session) placeCandidate(candidate *models.Candidate) {
	if len(s.slots.free) == 0 {
		s.createModuleDuplicate()
	}
	s.assignCandidateToASlot(candidate.UUID)
	s.printToFile(candidate)
}

func NewSession(modulePath types.AbsolutePath, testDetails *types.TestDetails) *Session {
	s := Session{
		modulePath:  modulePath,
		testDetails: testDetails,
		slots: slots{
			free:     []types.SlotPath{},
			assigned: map[models.CandidateID]types.SlotPath{}},
	}
	s.createMainFolder()
	return &s
}

func (s *Session) PlaceCandidatesIntoSlots(candidates []*models.Candidate) {
	for _, candidate := range candidates {
		s.placeCandidate(candidate)
	}
}

func (s *Session) FreeAllSlots() {
	s.slots.free = append(s.slots.free, maps.Values(s.slots.assigned)...)
}

func (s *Session) GetPackagePathForCandidate(candidateID models.CandidateID) string {
	return filepath.Join(string(s.tmp), string(s.slots.assigned[candidateID]), string(s.testDetails.PackagePath))
}
