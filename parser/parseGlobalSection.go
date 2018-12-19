package parser

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/utils"
)

func (p *Parser) globalSection(sec *Section) error {
	err := checkSection(sec, types.OrderGlobal)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)

	//1. num global ele
	var numGlobal uint32
	err = utils.DecodeVarInt(rd, 32, &numGlobal)
	if err != nil {
		return err
	}

	//2. global eles
	for i := 0; i < int(numGlobal); i++ {
		globalType, err := DecodeGlobalType(rd)
		if err != nil {
			return err
		}
		initExpression, err := DecodeInitializer(rd)
		if err != nil {
			return err
		}
		globalDef := types.GlobalDef{Type: globalType, Initializer: &initExpression}
		p.Module.Globals.Defs = append(p.Module.Globals.Defs, globalDef)
	}

	err = p.validateGlobal()
	return err
}

func (p *Parser) validateGlobal() error {
	logrus.Info("TODO: validateGlobal()")
	return nil
}
