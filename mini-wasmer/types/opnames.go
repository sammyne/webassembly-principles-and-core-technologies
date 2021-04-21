package types

var opnames = make([]string, 256)

func init() {
	opnames[OpcodeUnreachable] = "unreachable"
	opnames[OpcodeNop] = "nop"
	opnames[OpcodeBlock] = "block"
	opnames[OpcodeLoop] = "loop"
	opnames[OpcodeIf] = "if"
	opnames[OpcodeElse] = "else"
	opnames[OpcodeEnd] = "end"
	opnames[OpcodeBr] = "br"
	opnames[OpcodeBrIf] = "br_if"
	opnames[OpcodeBrTable] = "br_table"
	opnames[OpcodeReturn] = "return"
	opnames[OpcodeCall] = "call"
	opnames[OpcodeCallIndirect] = "call_indirect"
	opnames[OpcodeDrop] = "drop"
	opnames[OpcodeSelect] = "select"
	opnames[OpcodeLocalGet] = "local.get"
	opnames[OpcodeLocalSet] = "local.set"
	opnames[OpcodeLocalTee] = "local.tee"
	opnames[OpcodeGlobalGet] = "global.get"
	opnames[OpcodeGlobalSet] = "global.set"
	opnames[OpcodeI32Load] = "i32.load"
	opnames[OpcodeI64Load] = "i64.load"
	opnames[OpcodeF32Load] = "f32.load"
	opnames[OpcodeF64Load] = "f64.load"
	opnames[OpcodeI32Load8S] = "i32.load8_s"
	opnames[OpcodeI32Load8U] = "i32.load8_u"
	opnames[OpcodeI32Load16S] = "i32.load16_s"
	opnames[OpcodeI32Load16U] = "i32.load16_u"
	opnames[OpcodeI64Load8S] = "i64.load8_s"
	opnames[OpcodeI64Load8U] = "i64.load8_u"
	opnames[OpcodeI64Load16S] = "i64.load16_s"
	opnames[OpcodeI64Load16U] = "i64.load16_u"
	opnames[OpcodeI64Load32S] = "i64.load32_s"
	opnames[OpcodeI64Load32U] = "i64.load32_u"
	opnames[OpcodeI32Store] = "i32.store"
	opnames[OpcodeI64Store] = "i64.store"
	opnames[OpcodeF32Store] = "f32.store"
	opnames[OpcodeF64Store] = "f64.store"
	opnames[OpcodeI32Store8] = "i32.store8"
	opnames[OpcodeI32Store16] = "i32.store16"
	opnames[OpcodeI64Store8] = "i64.store8"
	opnames[OpcodeI64Store16] = "i64.store16"
	opnames[OpcodeI64Store32] = "i64.store32"
	opnames[OpcodeMemorySize] = "memory.size"
	opnames[OpcodeMemoryGrow] = "memory.grow"
	opnames[OpcodeI32Const] = "i32.const"
	opnames[OpcodeI64Const] = "i64.const"
	opnames[OpcodeF32Const] = "f32.const"
	opnames[OpcodeF64Const] = "f64.const"
	opnames[OpcodeI32Eqz] = "i32.eqz"
	opnames[OpcodeI32Eq] = "i32.eq"
	opnames[OpcodeI32Ne] = "i32.ne"
	opnames[OpcodeI32LtS] = "i32.lt_s"
	opnames[OpcodeI32LtU] = "i32.lt_u"
	opnames[OpcodeI32GtS] = "i32.gt_s"
	opnames[OpcodeI32GtU] = "i32.gt_u"
	opnames[OpcodeI32LeS] = "i32.le_s"
	opnames[OpcodeI32LeU] = "i32.le_u"
	opnames[OpcodeI32GeS] = "i32.ge_s"
	opnames[OpcodeI32GeU] = "i32.ge_u"
	opnames[OpcodeI64Eqz] = "i64.eqz"
	opnames[OpcodeI64Eq] = "i64.eq"
	opnames[OpcodeI64Ne] = "i64.ne"
	opnames[OpcodeI64LtS] = "i64.lt_s"
	opnames[OpcodeI64LtU] = "i64.lt_u"
	opnames[OpcodeI64GtS] = "i64.gt_s"
	opnames[OpcodeI64GtU] = "i64.gt_u"
	opnames[OpcodeI64LeS] = "i64.le_s"
	opnames[OpcodeI64LeU] = "i64.le_u"
	opnames[OpcodeI64GeS] = "i64.ge_s"
	opnames[OpcodeI64GeU] = "i64.ge_u"
	opnames[OpcodeF32Eq] = "f32.eq"
	opnames[OpcodeF32Ne] = "f32.ne"
	opnames[OpcodeF32Lt] = "f32.lt"
	opnames[OpcodeF32Gt] = "f32.gt"
	opnames[OpcodeF32Le] = "f32.le"
	opnames[OpcodeF32Ge] = "f32.ge"
	opnames[OpcodeF64Eq] = "f64.eq"
	opnames[OpcodeF64Ne] = "f64.ne"
	opnames[OpcodeF64Lt] = "f64.lt"
	opnames[OpcodeF64Gt] = "f64.gt"
	opnames[OpcodeF64Le] = "f64.le"
	opnames[OpcodeF64Ge] = "f64.ge"
	opnames[OpcodeI32Clz] = "i32.clz"
	opnames[OpcodeI32Ctz] = "i32.ctz"
	opnames[OpcodeI32PopCnt] = "i32.popcnt"
	opnames[OpcodeI32Add] = "i32.add"
	opnames[OpcodeI32Sub] = "i32.sub"
	opnames[OpcodeI32Mul] = "i32.mul"
	opnames[OpcodeI32DivS] = "i32.div_s"
	opnames[OpcodeI32DivU] = "i32.div_u"
	opnames[OpcodeI32RemS] = "i32.rem_s"
	opnames[OpcodeI32RemU] = "i32.rem_u"
	opnames[OpcodeI32And] = "i32.and"
	opnames[OpcodeI32Or] = "i32.or"
	opnames[OpcodeI32Xor] = "i32.xor"
	opnames[OpcodeI32Shl] = "i32.shl"
	opnames[OpcodeI32ShrS] = "i32.shr_s"
	opnames[OpcodeI32ShrU] = "i32.shr_u"
	opnames[OpcodeI32Rotl] = "i32.rotl"
	opnames[OpcodeI32Rotr] = "i32.rotr"
	opnames[OpcodeI64Clz] = "i64.clz"
	opnames[OpcodeI64Ctz] = "i64.ctz"
	opnames[OpcodeI64PopCnt] = "i64.popcnt"
	opnames[OpcodeI64Add] = "i64.add"
	opnames[OpcodeI64Sub] = "i64.sub"
	opnames[OpcodeI64Mul] = "i64.mul"
	opnames[OpcodeI64DivS] = "i64.div_s"
	opnames[OpcodeI64DivU] = "i64.div_u"
	opnames[OpcodeI64RemS] = "i64.rem_s"
	opnames[OpcodeI64RemU] = "i64.rem_u"
	opnames[OpcodeI64And] = "i64.and"
	opnames[OpcodeI64Or] = "i64.or"
	opnames[OpcodeI64Xor] = "i64.xor"
	opnames[OpcodeI64Shl] = "i64.shl"
	opnames[OpcodeI64ShrS] = "i64.shr_s"
	opnames[OpcodeI64ShrU] = "i64.shr_u"
	opnames[OpcodeI64Rotl] = "i64.rotl"
	opnames[OpcodeI64Rotr] = "i64.rotr"
	opnames[OpcodeF32Abs] = "f32.abs"
	opnames[OpcodeF32Neg] = "f32.neg"
	opnames[OpcodeF32Ceil] = "f32.ceil"
	opnames[OpcodeF32Floor] = "f32.floor"
	opnames[OpcodeF32Trunc] = "f32.trunc"
	opnames[OpcodeF32Nearest] = "f32.nearest"
	opnames[OpcodeF32Sqrt] = "f32.sqrt"
	opnames[OpcodeF32Add] = "f32.add"
	opnames[OpcodeF32Sub] = "f32.sub"
	opnames[OpcodeF32Mul] = "f32.mul"
	opnames[OpcodeF32Div] = "f32.div"
	opnames[OpcodeF32Min] = "f32.min"
	opnames[OpcodeF32Max] = "f32.max"
	opnames[OpcodeF32CopySign] = "f32.copysign"
	opnames[OpcodeF64Abs] = "f64.abs"
	opnames[OpcodeF64Neg] = "f64.neg"
	opnames[OpcodeF64Ceil] = "f64.ceil"
	opnames[OpcodeF64Floor] = "f64.floor"
	opnames[OpcodeF64Trunc] = "f64.trunc"
	opnames[OpcodeF64Nearest] = "f64.nearest"
	opnames[OpcodeF64Sqrt] = "f64.sqrt"
	opnames[OpcodeF64Add] = "f64.add"
	opnames[OpcodeF64Sub] = "f64.sub"
	opnames[OpcodeF64Mul] = "f64.mul"
	opnames[OpcodeF64Div] = "f64.div"
	opnames[OpcodeF64Min] = "f64.min"
	opnames[OpcodeF64Max] = "f64.max"
	opnames[OpcodeF64CopySign] = "f64.copysign"
	opnames[OpcodeI32WrapI64] = "i32.wrap_i64"
	opnames[OpcodeI32TruncF32S] = "i32.trunc_f32_s"
	opnames[OpcodeI32TruncF32U] = "i32.trunc_f32_u"
	opnames[OpcodeI32TruncF64S] = "i32.trunc_f64_s"
	opnames[OpcodeI32TruncF64U] = "i32.trunc_f64_u"
	opnames[OpcodeI64ExtendI32S] = "i64.extend_i32_s"
	opnames[OpcodeI64ExtendI32U] = "i64.extend_i32_u"
	opnames[OpcodeI64TruncF32S] = "i64.trunc_f32_s"
	opnames[OpcodeI64TruncF32U] = "i64.trunc_f32_u"
	opnames[OpcodeI64TruncF64S] = "i64.trunc_f64_s"
	opnames[OpcodeI64TruncF64U] = "i64.trunc_f64_u"
	opnames[OpcodeF32ConvertI32S] = "f32.convert_i32_s"
	opnames[OpcodeF32ConvertI32U] = "f32.convert_i32_u"
	opnames[OpcodeF32ConvertI64S] = "f32.convert_i64_s"
	opnames[OpcodeF32ConvertI64U] = "f32.convert_i64_u"
	opnames[OpcodeF32DemoteF64] = "f32.demote_f64"
	opnames[OpcodeF64ConvertI32S] = "f64.convert_i32_s"
	opnames[OpcodeF64ConvertI32U] = "f64.convert_i32_u"
	opnames[OpcodeF64ConvertI64S] = "f64.convert_i64_s"
	opnames[OpcodeF64ConvertI64U] = "f64.convert_i64_u"
	opnames[OpcodeF64PromoteF32] = "f64.promote_f32"
	opnames[OpcodeI32ReinterpretF32] = "i32.reinterpret_f32"
	opnames[OpcodeI64ReinterpretF64] = "i64.reinterpret_f64"
	opnames[OpcodeF32ReinterpretI32] = "f32.reinterpret_i32"
	opnames[OpcodeF64ReinterpretI64] = "f64.reinterpret_i64"
	opnames[OpcodeI32Extend8S] = "i32.extend8_s"
	opnames[OpcodeI32Extend16S] = "i32.extend16_s"
	opnames[OpcodeI64Extend8S] = "i64.extend8_s"
	opnames[OpcodeI64Extend16S] = "i64.extend16_s"
	opnames[OpcodeI64Extend32S] = "i64.extend32_s"
	opnames[OpcodeTruncSat] = "trunc_sat"
}

func GetOpname(opcode byte) (string, bool) {
	v := opnames[opcode]
	return v, v != ""
}
