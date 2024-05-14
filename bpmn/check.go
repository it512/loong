package bpmn

import "fmt"

type CheckFunc func(string, *TDefinitions) error

var checkFnMap map[string]CheckFunc = make(map[string]CheckFunc)

func RegisterCheckFunc(key string, fn CheckFunc) {
	checkFnMap[key] = fn
}

func Check(tag string, def *TDefinitions) error {
	if err := checkParallelGateway(tag, def); err != nil {
		return err
	}
	if err := checkDefinitions(tag, def); err != nil {
		return err
	}

	for _, fn := range checkFnMap {
		if err := fn(tag, def); err != nil {
			return err
		}
	}
	return nil
}

func checkDefinitions(tag string, def *TDefinitions) error {
	if def.Id == "" {
		return fmt.Errorf("tag: %s, 流程定义ID为空", tag)
	}
	return nil
}

func checkParallelGateway(tag string, def *TDefinitions) error {
	for _, gw := range def.Process.ParallelGateway {
		in := len(gw.IncomingAssociation)
		out := len(gw.OutgoingAssociation)

		if in == 0 || out == 0 {
			return fmt.Errorf("tag: %s, ParallelGateway %s 存在0个出入口， %d 个入口 %d 出口", tag, gw.Id, in, out)
		}

		if in == 1 && out == 1 {
			return fmt.Errorf("tag: %s, ParallelGateway %s 为一进一出", tag, gw.Id)
		}
	}

	return nil
}
