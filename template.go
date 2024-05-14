package loong

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/it512/loong/bpmn"
)

func FileTemplates(root string, patterns ...string) Option {
	load := NewFileDirLoad(root, patterns...)
	return SetTemplates(load)
}

func SetTemplates(load TemplatesLoader) Option {
	tpls := Must(load.Load())
	return func(c *Config) {
		c.templates = tpls
	}
}

type TemplatesLoader interface {
	Load() (Templates, error)
}

type FileDirTemplatesLaod struct {
	root     string
	patterns []string
}

func NewFileDirLoad(root string, patterns ...string) *FileDirTemplatesLaod {
	return &FileDirTemplatesLaod{
		root:     root,
		patterns: patterns,
	}
}

func (l *FileDirTemplatesLaod) Load() (Templates, error) {
	mdh := make(Templates)
	err := fs.WalkDir(os.DirFS(l.root), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		for _, pattern := range l.patterns {
			if ok := Must(filepath.Match(pattern, path)); ok {
				absp := Must(filepath.Abs(filepath.Join(l.root, path)))
				d, raw, err := bpmn.LoadFormFile(absp)
				if err != nil {
					return err
				}
				t := newTemplate(d, raw)
				if _, has := mdh[t.ProcID]; has {
					return fmt.Errorf("%s is already exists", t.ProcID)
				}
				mdh[t.ProcID] = t
				break
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return mdh, nil
}

type Templates map[string]*Template

/*
func (t Templates) GetDefinitions(defID string) *bpmn.TDefinitions {
	for _, v := range t {
		if v.DefID == defID {
			return v.Definitions
		}
	}
	return nil
}
*/

func (t Templates) GetTemplate(procID string) *Template {
	return t[procID]
}

type Template struct {
	DefID       string
	DefName     string
	ProcID      string
	ProcName    string
	Definitions *bpmn.TDefinitions
	BpmnData    []byte

	cache map[string]BpmnElement
}

func NewTemplate(def *bpmn.TDefinitions) *Template {
	return newTemplate(def, nil)
}

func newTemplate(def *bpmn.TDefinitions, raw []byte) *Template {
	return &Template{
		DefID:       def.Id,
		DefName:     def.Name,
		ProcID:      def.Process.Id,
		ProcName:    def.Process.Name,
		Definitions: def,
		BpmnData:    raw,

		cache: make(map[string]BpmnElement),
	}
}

func (h *Template) FindNormalStartEvent() (bpmn.TStartEvent, bool) {
	return bpmn.FindNormalStartEvent(h.Definitions)
}

func (h *Template) FindUserTask(id string) (bpmn.TUserTask, bool) {
	return bpmn.Find(h.Definitions.Process.UserTasks, id)
}

func (h *Template) FindElementByID(id string) (ele BpmnElement, ok bool) {
	if ele, ok = h.cache[id]; !ok {
		if ele, ok = bpmn.FindElementById(h.Definitions, id); ok {
			h.cache[id] = ele
		}
	}
	return
}

func (h *Template) FindSequenceFlows(ids []string) []bpmn.TSequenceFlow {
	return bpmn.FindSequenceFlows(h.Definitions, ids)
}

func (h *Template) FindSequenceFlow(id string) (bpmn.TSequenceFlow, bool) {
	return bpmn.Find(h.Definitions.Process.SequenceFlows, id)
}

func (h *Template) FindError(id string) (bpmn.TError, bool) {
	return bpmn.Find(h.Definitions.Errors, id)
}
