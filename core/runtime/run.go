package runtime

import (
	"fmt"
	"reflect"
	"wasm/core/IR"
	"wasm/types"
	"wasm/utils"
)

func (vm *VM) Run(functionNameOrID interface{}, params ...interface{}) (err error) {
	//TODO check all type assertion
	defer utils.CatchError(&err)
	frame := vm.GetCurrentFrame()
	if frame == nil {
		panic("Frame stack overflow.")
	}
	funcIndex := 0
	switch functionNameOrID.(type) {
	case string:
		funcIndex, err = vm.Module.GetFuncIndexWithName(functionNameOrID.(string))
		if err != nil {
			vm.panic(err)
		}
	case int:
		funcIndex = functionNameOrID.(int)
	default:
		panic("functionNameOrID must be a string or an int")
	}

	Params := make([]IR.InterfaceValue, 0)
	//int8,int16,int32 --> int32;  uint8,uint16,uint32 --> uint32
	//int,int64 --> int64; uint,uint64 --> uint64
	for _, p := range params {
		kind := reflect.TypeOf(p).Kind()
		switch kind {
		case reflect.Int8:
			Params = append(Params, &Value{IR.TypeI32, int32(p.(int8))})
		case reflect.Uint8:
			Params = append(Params, &Value{IR.TypeI32, uint32(p.(uint8))})
		case reflect.Int16:
			Params = append(Params, &Value{IR.TypeI32, int32(p.(int16))})
		case reflect.Uint16:
			Params = append(Params, &Value{IR.TypeI32, uint32(p.(uint16))})
		case reflect.Int32:
			Params = append(Params, &Value{IR.TypeI32, p.(int32)})
		case reflect.Uint32:
			Params = append(Params, &Value{IR.TypeI32, p.(uint32)})

		case reflect.Int:
			Params = append(Params, &Value{IR.TypeI64, int64(p.(int))})
		case reflect.Uint:
			Params = append(Params, &Value{IR.TypeI64, uint64(p.(uint))})
		case reflect.Int64:
			Params = append(Params, &Value{IR.TypeI64, p.(int64)})
		case reflect.Uint64:
			Params = append(Params, &Value{IR.TypeI64, p.(uint64)})

		case reflect.Float32:
			Params = append(Params, &Value{IR.TypeF32, p.(float32)})
		case reflect.Float64:
			Params = append(Params, &Value{IR.TypeF64, p.(float64)})

		default:
			panic("Parameter type not valid,only support (u)int8,u(int16),u(int32),u(int),u(int64),float32,float64")
		}
	}

	err = frame.Init(funcIndex, vm, Params)
	endIndice := vm.FunctionCodes[funcIndex].EndIndice
	if err != nil {
		vm.panic(err)
	}
	ifStack := utils.Stack{}
	//Execute
	for {
		lastPC := frame.PC
		ins := frame.Instruction[frame.PC]

		switch ins.Op.Code {
		case IR.OPCunreachable:
			vm.panic("unreachable executed")
		case IR.OPCbr:
			td := ins.Imm.(*IR.BranchImm).TargetDepth
			tarInsIndex := endIndice[int(td)]
			frame.advanceTo(tarInsIndex)
		case IR.OPCbr_if:
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			con, _ := frame.Stack.Pop()
			v := con.Value().(int32)
			if v == 1 {
				td := ins.Imm.(*IR.BranchImm).TargetDepth
				tarInsIndex := endIndice[int(td)]
				frame.advanceTo(tarInsIndex)
			} else {
				frame.advance(1)
			}

		case IR.OPCbr_table:
		case IR.OPCreturn_:
			var retV IR.InterfaceValue
			vm.CurrentFrame -= 1
			hasResult := !frame.Stack.Empty()
			if hasResult {
				//has return value
				retV, _ = frame.Stack.Pop()
				//TODO check frame.FunctionSig with retV
			}
			if vm.CurrentFrame == -1 {
				if hasResult {
					vm.ReturnValue = retV
				}
				return nil
			} else {
				if hasResult {
					frame.Stack.Push(retV)
				}
				frame = vm.GetCurrentFrame()
			}
		case IR.OPCcall:
			frame.advance(1)
			funcIndex := ins.Imm.(*IR.FunctionImm).FunctionIndex
			vm.CurrentFrame += 1
			if vm.CurrentFrame >= MAXFRAME {
				vm.panic(types.ErrBeyondMaxFrame)
			}

			fType := vm.Module.Types[int(vm.FunctionCodes[funcIndex].Type.Index)]
			paraCount := fType.Params.NumElems
			if frame.Stack.Len() < int(paraCount) {
				vm.panic(types.ErrStackSizeErr)
			}
			params := make([]IR.InterfaceValue, paraCount)
			for elemIndex := range fType.Params.Elems {
				v, _ := frame.Stack.Pop()
				params[int(paraCount)-1-elemIndex] = v
			}

			frame = vm.GetCurrentFrame()
			err = frame.Init(int(funcIndex), vm, params)
			if err != nil {
				vm.panic(err)
			}

		case IR.OPCcall_indirect:
		case IR.OPCdrop:
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			frame.Stack.Pop()
			frame.advance(1)
		case IR.OPCselect:
			if frame.Stack.Len() < 3 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Pop()
			b, _ := frame.Stack.Pop()
			c, _ := frame.Stack.Pop()
			if IsZero(c) {
				frame.Stack.Push(b)
			} else {
				frame.Stack.Push(a)
			}
			frame.advance(1)

		case IR.OPCget_local:
			index := ins.Imm.(*IR.GetOrSetVariableImm).VariableIndex
			if index >= uint64(len(frame.Locals)) {
				vm.panic("get_local index out of range")
			}
			frame.Stack.Push(frame.Locals[index])
			frame.advance(1)
		case IR.OPCset_local:
			index := ins.Imm.(*IR.GetOrSetVariableImm).VariableIndex
			if index >= uint64(len(frame.Locals)) {
				vm.panic("set_local index out of range")
			}
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Pop()
			frame.Locals[index] = a
			frame.advance(1)
		case IR.OPCtee_local:
			index := ins.Imm.(*IR.GetOrSetVariableImm).VariableIndex
			if index >= uint64(len(frame.Locals)) {
				vm.panic("tee_local index out of range")
			}
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Top()
			frame.Locals[index] = a
			frame.advance(1)
		case IR.OPCget_global:
			index := ins.Imm.(*IR.GetOrSetVariableImm).VariableIndex
			if index >= uint64(len(vm.Global)) {
				vm.panic("get_local index out of range")
			}
			frame.Stack.Push(vm.Global[index])
			frame.advance(1)
		case IR.OPCset_global:
			index := ins.Imm.(*IR.GetOrSetVariableImm).VariableIndex
			if index >= uint64(len(vm.Global)) {
				vm.panic("set_local index out of range")
			}
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Pop()
			vm.Global[index] = a
			frame.advance(1)
		case IR.OPCtable_get:
		case IR.OPCtable_set:
		case IR.OPCthrow_:
		case IR.OPCrethrow:
		case IR.OPCnop:
			frame.advance(1)
		case IR.OPCi32_load:
			err = i32_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 4, true)
		case IR.OPCi64_load:
			err = i64_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 8, true)
		case IR.OPCf32_load:
			err = float_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 4)
		case IR.OPCf64_load:
			err = float_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 8)
		case IR.OPCi32_load8_s:
			err = i32_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 1, true)
		case IR.OPCi32_load8_u:
			err = i32_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 1, false)
		case IR.OPCi32_load16_s:
			err = i32_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 2, true)
		case IR.OPCi32_load16_u:
			err = i32_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 2, false)
		case IR.OPCi64_load8_s:
			err = i64_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 1, true)
		case IR.OPCi64_load8_u:
			err = i64_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 1, false)
		case IR.OPCi64_load16_s:
			err = i64_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 2, true)
		case IR.OPCi64_load16_u:
			err = i64_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 2, false)
		case IR.OPCi64_load32_s:
			err = i64_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 4, true)
		case IR.OPCi64_load32_u:
			err = i64_load(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 4, false)
		case IR.OPCi32_store:
			err = i32_store(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 4)
		case IR.OPCi64_store:
			err = i64_store(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 8)
		case IR.OPCf32_store:
			err = float_store(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 4)
		case IR.OPCf64_store:
			err = float_store(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 8)
		case IR.OPCi32_store8:
			err = i32_store(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 1)
		case IR.OPCi32_store16:
			err = i32_store(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 2)
		case IR.OPCi64_store8:
			err = i64_store(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 1)
		case IR.OPCi64_store16:
			err = i64_store(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 2)
		case IR.OPCi64_store32:
			err = i64_store(vm, frame, ins.Imm.(*IR.LoadOrStoreImm).Offset, 4)
		case IR.OPCmemory_size:
		case IR.OPCmemory_grow:
		case IR.OPCi32_const:
			defConst(vm, frame, I32_CONST)
		case IR.OPCi64_const:
			defConst(vm, frame, I64_CONST)
		case IR.OPCf32_const:
			defConst(vm, frame, F32_CONST)
		case IR.OPCf64_const:
			defConst(vm, frame, F64_CONST)
		case IR.OPCi32_eqz:
			err = eqz(vm, frame, I32_EQZ)
		case IR.OPCi32_eq:
			err = i32_compare(vm, frame, CMP_EQ, true)
		case IR.OPCi32_ne:
			err = i32_compare(vm, frame, CMP_NE, true)
		case IR.OPCi32_lt_s:
			err = i32_compare(vm, frame, CMP_LT, true)
		case IR.OPCi32_lt_u:
			err = i32_compare(vm, frame, CMP_LT, false)
		case IR.OPCi32_gt_s:
			err = i32_compare(vm, frame, CMP_GT, true)
		case IR.OPCi32_gt_u:
			err = i32_compare(vm, frame, CMP_GT, false)
		case IR.OPCi32_le_s:
			err = i32_compare(vm, frame, CMP_LE, true)
		case IR.OPCi32_le_u:
			err = i32_compare(vm, frame, CMP_LE, false)
		case IR.OPCi32_ge_s:
			err = i32_compare(vm, frame, CMP_GE, true)
		case IR.OPCi32_ge_u:
			err = i32_compare(vm, frame, CMP_GE, false)

		case IR.OPCi64_eqz:
			err = eqz(vm, frame, I64_EQZ)
		case IR.OPCi64_eq:
			err = i64_compare(vm, frame, CMP_EQ, true)
		case IR.OPCi64_ne:
			err = i64_compare(vm, frame, CMP_NE, true)
		case IR.OPCi64_lt_s:
			err = i64_compare(vm, frame, CMP_LT, true)
		case IR.OPCi64_lt_u:
			err = i64_compare(vm, frame, CMP_LT, false)
		case IR.OPCi64_gt_s:
			err = i64_compare(vm, frame, CMP_GT, true)
		case IR.OPCi64_gt_u:
			err = i64_compare(vm, frame, CMP_GT, false)
		case IR.OPCi64_le_s:
			err = i64_compare(vm, frame, CMP_LE, true)
		case IR.OPCi64_le_u:
			err = i64_compare(vm, frame, CMP_LE, false)
		case IR.OPCi64_ge_s:
			err = i64_compare(vm, frame, CMP_GE, true)
		case IR.OPCi64_ge_u:
			err = i64_compare(vm, frame, CMP_GE, false)
		case IR.OPCf32_eq:
			err = f32_compare(vm, frame, CMP_EQ)
		case IR.OPCf32_ne:
			err = f32_compare(vm, frame, CMP_NE)
		case IR.OPCf32_lt:
			err = f32_compare(vm, frame, CMP_LT)
		case IR.OPCf32_gt:
			err = f32_compare(vm, frame, CMP_GT)
		case IR.OPCf32_le:
			err = f32_compare(vm, frame, CMP_LE)
		case IR.OPCf32_ge:
			err = f32_compare(vm, frame, CMP_GE)
		case IR.OPCf64_eq:
			err = f64_compare(vm, frame, CMP_EQ)
		case IR.OPCf64_ne:
			err = f64_compare(vm, frame, CMP_NE)
		case IR.OPCf64_lt:
			err = f64_compare(vm, frame, CMP_LT)
		case IR.OPCf64_gt:
			err = f64_compare(vm, frame, CMP_GT)
		case IR.OPCf64_le:
			err = f64_compare(vm, frame, CMP_LE)
		case IR.OPCf64_ge:
			err = f64_compare(vm, frame, CMP_GE)

		case IR.OPCi32_clz:
			err = clz(vm, frame, I32_CLZ)
		case IR.OPCi32_ctz:
			err = ctz(vm, frame, I32_CTZ)
		case IR.OPCi32_popcnt:
			err = popcnt(vm, frame, I32_POPCNT)
		case IR.OPCi32_add:
			err = i32_arith(vm, frame, ARITH_ADD)
		case IR.OPCi32_sub:
			err = i32_arith(vm, frame, ARITH_SUB)
		case IR.OPCi32_mul:
			err = i32_arith(vm, frame, ARITH_MUL)
		case IR.OPCi32_div_s:
			err = i32_arith(vm, frame, ARITH_DIV)
		case IR.OPCi32_div_u:
			err = i32_arith(vm, frame, ARITH_DIV)
		case IR.OPCi32_rem_s:
			err = i32_arith(vm, frame, ARITH_REM)
		case IR.OPCi32_rem_u:
			err = i32_arith(vm, frame, ARITH_REM)
		case IR.OPCi32_and_:
			err = i32_arith(vm, frame, ARITH_AND)
		case IR.OPCi32_or_:
			err = i32_arith(vm, frame, ARITH_OR)
		case IR.OPCi32_xor_:
			err = i32_arith(vm, frame, ARITH_XOR)
		case IR.OPCi32_shl:
			err = i32_arith(vm, frame, ARITH_SHL)
		case IR.OPCi32_shr_s:
			err = i32_arith(vm, frame, ARITH_SHR)
		case IR.OPCi32_shr_u:
			err = i32_arith(vm, frame, ARITH_SHR)
		case IR.OPCi32_rotl:
			err = i32_arith(vm, frame, ARITH_ROTL)
		case IR.OPCi32_rotr:
			err = i32_arith(vm, frame, ARITH_ROTR)

		case IR.OPCi64_clz:
			err = clz(vm, frame, I64_CLZ)
		case IR.OPCi64_ctz:
			err = ctz(vm, frame, I64_CTZ)
		case IR.OPCi64_popcnt:
			err = popcnt(vm, frame, I64_POPCNT)
		case IR.OPCi64_add:
			err = i64_arith(vm, frame, ARITH_AND)
		case IR.OPCi64_sub:
			err = i64_arith(vm, frame, ARITH_SUB)
		case IR.OPCi64_mul:
			err = i64_arith(vm, frame, ARITH_MUL)
		case IR.OPCi64_div_s:
			err = i64_arith(vm, frame, ARITH_DIV)
		case IR.OPCi64_div_u:
			err = i64_arith(vm, frame, ARITH_DIV)
		case IR.OPCi64_rem_s:
			err = i64_arith(vm, frame, ARITH_REM)
		case IR.OPCi64_rem_u:
			err = i64_arith(vm, frame, ARITH_REM)
		case IR.OPCi64_and_:
			err = i64_arith(vm, frame, ARITH_AND)
		case IR.OPCi64_or_:
			err = i64_arith(vm, frame, ARITH_OR)
		case IR.OPCi64_xor_:
			err = i64_arith(vm, frame, ARITH_XOR)
		case IR.OPCi64_shl:
			err = i64_arith(vm, frame, ARITH_SHL)
		case IR.OPCi64_shr_s:
			err = i64_arith(vm, frame, ARITH_SHR)
		case IR.OPCi64_shr_u:
			err = i64_arith(vm, frame, ARITH_SHR)
		case IR.OPCi64_rotl:
			err = i64_arith(vm, frame, ARITH_ROTL)
		case IR.OPCi64_rotr:
			err = i64_arith(vm, frame, ARITH_ROTR)

		case IR.OPCf32_abs:
			err = f32Abs(vm, frame)
		case IR.OPCf32_neg:
			err = f32Neg(vm, frame)
		case IR.OPCf32_ceil:
			err = f32Ceil(vm, frame)
		case IR.OPCf32_floor:
			err = f32Floor(vm, frame)
		case IR.OPCf32_trunc:
			err = f32Trunc(vm, frame)
		case IR.OPCf32_nearest:
			err = f32Nearest(vm, frame)
		case IR.OPCf32_sqrt:
			err = f32Sqrt(vm, frame)

		case IR.OPCf32_add:
			err = f32_arith(vm, frame, ARITH_ADD)
		case IR.OPCf32_sub:
			err = f32_arith(vm, frame, ARITH_SUB)
		case IR.OPCf32_mul:
			err = f32_arith(vm, frame, ARITH_MUL)
		case IR.OPCf32_div:
			err = f32_arith(vm, frame, ARITH_DIV)
		case IR.OPCf32_min:
			err = f32_arith(vm, frame, ARITH_MIN)
		case IR.OPCf32_max:
			err = f32_arith(vm, frame, ARITH_MAX)
		case IR.OPCf32_copysign:
			err = f32copySign(vm, frame)

		case IR.OPCf64_abs:
			err = f64Abs(vm, frame)
		case IR.OPCf64_neg:
			err = f64Neg(vm, frame)
		case IR.OPCf64_ceil:
			err = f64Ceil(vm, frame)
		case IR.OPCf64_floor:
			err = f64Floor(vm, frame)
		case IR.OPCf64_trunc:
			err = f64Trunc(vm, frame)
		case IR.OPCf64_nearest:
			err = f64Nearest(vm, frame)
		case IR.OPCf64_sqrt:
			err = f64Sqrt(vm, frame)

		case IR.OPCf64_add:
			err = f64_arith(vm, frame, ARITH_ADD)
		case IR.OPCf64_sub:
			err = f64_arith(vm, frame, ARITH_SUB)
		case IR.OPCf64_mul:
			err = f64_arith(vm, frame, ARITH_MUL)
		case IR.OPCf64_div:
			err = f64_arith(vm, frame, ARITH_DIV)
		case IR.OPCf64_min:
			err = f64_arith(vm, frame, ARITH_MIN)
		case IR.OPCf64_max:
			err = f64_arith(vm, frame, ARITH_MAX)
		case IR.OPCf64_copysign:
			err = f64copySign(vm, frame)

		case IR.OPCi32_wrap_i64:
		case IR.OPCi32_trunc_s_f32:
		case IR.OPCi32_trunc_u_f32:
		case IR.OPCi32_trunc_s_f64:
		case IR.OPCi32_trunc_u_f64:
		case IR.OPCi64_extend_s_i32:
		case IR.OPCi64_extend_u_i32:
		case IR.OPCi64_trunc_s_f32:
		case IR.OPCi64_trunc_u_f32:
		case IR.OPCi64_trunc_s_f64:
		case IR.OPCi64_trunc_u_f64:
		case IR.OPCf32_convert_s_i32:
		case IR.OPCf32_convert_u_i32:
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Pop()
			m, ok := a.Value().(uint32)
			if !ok {
				vm.panic(fmt.Sprintf(types.ErrTypeAssertion, "uint32"))
			}
			v := float32(m)
			frame.Stack.Push(&Value{IR.TypeF32, v})
			frame.advance(1)
		case IR.OPCf32_convert_s_i64:
		case IR.OPCf32_convert_u_i64:
		case IR.OPCf32_demote_f64:
		case IR.OPCf64_convert_s_i32:
		case IR.OPCf64_convert_u_i32:
		case IR.OPCf64_convert_s_i64:
		case IR.OPCf64_convert_u_i64:
		case IR.OPCf64_promote_f32:
		case IR.OPCi32_reinterpret_f32:
		case IR.OPCi64_reinterpret_f64:
		case IR.OPCf32_reinterpret_i32:
		case IR.OPCf64_reinterpret_i64:
		case IR.OPCi32_extend8_s:
		case IR.OPCi32_extend16_s:
		case IR.OPCi64_extend8_s:
		case IR.OPCi64_extend16_s:
		case IR.OPCi64_extend32_s:
		case IR.OPCref_null:
		case IR.OPCref_isnull:
		case IR.OPCref_func:

		case IR.OPCblock:
			frame.advance(1)
		case IR.OPCloop:
			frame.advance(1)
		case IR.OPCif_:
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			con, _ := frame.Stack.Pop()
			if !IsZero(con) {
				ifStack.Push(ins.Index)
				frame.advance(1)
			} else {
				//advanctTo else or end
				nextIndex := ins.MatchedIndex
				if nextIndex == -1 {
					vm.panic("if without else or end")
				}
				frame.advanceTo(nextIndex)
			}
		case IR.OPCelse_:
			if ifStack.Empty() {
				frame.advance(1)
			} else {
				frame.advanceTo(ins.MatchedIndex)
			}

		case IR.OPCend:
			switch ins.MatchedIndex {
			case types.LastOpcode:
				//if ins.Index == len(frame.Instruction)-1 {
				var retV IR.InterfaceValue
				vm.CurrentFrame -= 1
				hasResult := !frame.Stack.Empty()
				if frame.FunctionDef.FunctionType.Results.NumElems > 0 && !hasResult {
					panic("This function should have return value,but stack is empty")
				}
				if hasResult {
					//has return value
					retV, _ = frame.Stack.Pop()
					//TODO check frame.FunctionSig with retV
				}
				if vm.CurrentFrame == -1 {
					if hasResult {
						vm.ReturnValue = retV
					}
					return nil
				} else {
					frame = vm.GetCurrentFrame()
					if hasResult {
						//TODO assert frame.Stack is empty
						frame.Stack.Push(retV)
					}
				}
			case -1:
				vm.panic("end ins matched index illegal")
			default:
				//TODO pop value stack since block or if
				mindex := ins.MatchedIndex
				switch frame.Instruction[mindex].Op.Code {
				case IR.OPCblock:
					frame.advance(1)
				case IR.OPCif_:
					ifStack.Pop()
					frame.advance(1)
				case IR.OPCloop:
					frame.advanceTo(mindex + 1)
				default:
					vm.panic("end ins not point to block/if/loop")
				}
			}

		case IR.OPCtry_:
		case IR.OPCcatch_:
		case IR.OPCcatch_all:
		default:
			frame.runBinaryOp(vm, &ins)
		}
		if err != nil {
			vm.panic(err)
		}
		if lastPC == frame.PC {
			panicStr := fmt.Sprintf("PC not changed. PC: %d instruction: %v", frame.PC, frame.Instruction[frame.PC].Op.Name)
			vm.panic(panicStr)
		}
	}
}
