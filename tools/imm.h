#define enum_imms(visit_imm) \
visit_imm(AtomicLoadOrStoreImm)\
visit_imm(BranchImm)\
visit_imm(BranchTableImm)\
visit_imm(CallIndirectImm)\
visit_imm(ControlStructureImm)\
visit_imm(DataSegmentAndMemImm)\
visit_imm(DataSegmentImm)\
visit_imm(ElemSegmentAndTableImm)\
visit_imm(ElemSegmentImm)\
visit_imm(ExceptionTypeImm)\
visit_imm(FunctionImm)\
visit_imm(GetOrSetVariableImm)\
visit_imm(LaneIndexImm)\
visit_imm(LiteralImm_F32)\
visit_imm(LiteralImm_F64)\
visit_imm(LiteralImm_I32)\
visit_imm(LiteralImm_I64)\
visit_imm(LiteralImm_V128)\
visit_imm(LoadOrStoreImm)\
visit_imm(MemoryImm)\
visit_imm(NoImm)\
visit_imm(RethrowImm)\
visit_imm(ShuffleImm_16)\
visit_imm(TableImm)
