package bpmn

import (
	"encoding/xml"
	"os"
)

func LoadFromBytes(xmlData []byte) (*TDefinitions, []byte, error) {
	var (
		err error
	)

	def := &TDefinitions{}
	if err = xml.Unmarshal(xmlData, def); err != nil {
		return def, xmlData, nil
	}

	return def, xmlData, nil
}

func LoadFormFile(filename string) (*TDefinitions, []byte, error) {
	var (
		xmlData []byte
		err     error
	)

	if xmlData, err = os.ReadFile(filename); err != nil {
		return nil, xmlData, err
	}
	return LoadFromBytes(xmlData)
}
