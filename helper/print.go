package helper

import (
	"fmt"
	"strings"
)

func (fp *FunctionProto) String() string {
	return fp.str(1, 0)
}

func (fp *FunctionProto) str(level int, count int) string {
	indent := strings.Repeat("  ", level-1)
	buf := []string{}
	buf = append(buf, fmt.Sprintf("%v; function [%v] definition (level %v)\n",
		indent, count, level))
	buf = append(buf, fmt.Sprintf("%v; %v upvalues, %v params, %v stacks\n",
		indent, fp.NumUpvalues, fp.NumParameters, fp.NumUsedRegisters))
	for reg, linfo := range fp.DbgLocals {
		buf = append(buf, fmt.Sprintf("%v.local %v ; %v\n", indent, linfo.Name, reg))
	}
	for reg, upvalue := range fp.DbgUpvalues {
		buf = append(buf, fmt.Sprintf("%v.upvalue %v ; %v\n", indent, upvalue, reg))
	}
	for reg, conzt := range fp.Constants {
		buf = append(buf, fmt.Sprintf("%v.const %v ; %v\n", indent, conzt, reg))
	}
	buf = append(buf, "\n")

	protono := 0
	for no, code := range fp.Code {
		inst := OpCode(code)
		if inst == OP_CLOSURE {
			buf = append(buf, "\n")
			buf = append(buf, fp.FunctionPrototypes[protono].str(level+1, protono))
			buf = append(buf, "\n")
			protono++
		}
		buf = append(buf, fmt.Sprintf("%v[%03d] %v (line:%v)\n",
			indent, no+1, opToString(code), fp.DbgSourcePositions[no]))

	}
	buf = append(buf, fmt.Sprintf("%v; end of function\n", indent))
	return strings.Join(buf, "")
}

func opToString(inst uint32) string {
	op := int(inst >> 0 & (1 << 6 - 1))
	if op > opCodeMax {
		return ""
	}
	prop := &(opProps[op])

	arga := A(inst)
	argb := B(inst)
	argc := C(inst)
	argbx := Bx(inst)
	argsbx := Sbx(inst)

	buf := ""
	switch prop.Type {
	case opTypeABC:
		buf = fmt.Sprintf("%s      |  %d, %d, %d", prop.Name, arga, argb, argc)
	case opTypeABx:
		buf = fmt.Sprintf("%s      |  %d, %d", prop.Name, arga, argbx)
	case opTypeASbx:
		buf = fmt.Sprintf("%s      |  %d, %d", prop.Name, arga, argsbx)
	}

	switch op {
	case OP_MOVE:
		buf += fmt.Sprintf("; R(%v) := R(%v)", arga, argb)
	case OP_LOADK:
		buf += fmt.Sprintf("; R(%v) := Kst(%v)", arga, argbx)
	case OP_LOADBOOL:
		buf += fmt.Sprintf("; R(%v) := (Bool)%v; if (%v) pc++", arga, argb, argc)
	case OP_LOADNIL:
		buf += fmt.Sprintf("; R(%v) := ... := R(%v) := nil", arga, argb)
	case OP_GETUPVAL:
		buf += fmt.Sprintf("; R(%v) := UpValue[%v]", arga, argb)
	case OP_GETGLOBAL:
		buf += fmt.Sprintf("; R(%v) := Gbl[Kst(%v)]", arga, argbx)
	case OP_GETTABLE:
		buf += fmt.Sprintf("; R(%v) := R(%v)[RK(%v)]", arga, argb, argc)
	case OP_SETGLOBAL:
		buf += fmt.Sprintf("; Gbl[Kst(%v)] := R(%v)", argbx, arga)
	case OP_SETUPVAL:
		buf += fmt.Sprintf("; UpValue[%v] := R(%v)", argb, arga)
	case OP_SETTABLE:
		buf += fmt.Sprintf("; R(%v)[RK(%v)] := RK(%v)", arga, argb, argc)
	case OP_NEWTABLE:
		buf += fmt.Sprintf("; R(%v) := {} (size = BC)", arga)
	case OP_SELF:
		buf += fmt.Sprintf("; R(%v+1) := R(%v); R(%v) := R(%v)[RK(%v)]", arga, argb, arga, argb, argc)
	case OP_ADD:
		buf += fmt.Sprintf("; R(%v) := RK(%v) + RK(%v)", arga, argb, argc)
	case OP_SUB:
		buf += fmt.Sprintf("; R(%v) := RK(%v) - RK(%v)", arga, argb, argc)
	case OP_MUL:
		buf += fmt.Sprintf("; R(%v) := RK(%v) * RK(%v)", arga, argb, argc)
	case OP_DIV:
		buf += fmt.Sprintf("; R(%v) := RK(%v) / RK(%v)", arga, argb, argc)
	case OP_MOD:
		buf += fmt.Sprintf("; R(%v) := RK(%v) %% RK(%v)", arga, argb, argc)
	case OP_POW:
		buf += fmt.Sprintf("; R(%v) := RK(%v) ^ RK(%v)", arga, argb, argc)
	case OP_UNM:
		buf += fmt.Sprintf("; R(%v) := -R(%v)", arga, argb)
	case OP_NOT:
		buf += fmt.Sprintf("; R(%v) := not R(%v)", arga, argb)
	case OP_LEN:
		buf += fmt.Sprintf("; R(%v) := length of R(%v)", arga, argb)
	case OP_CONCAT:
		buf += fmt.Sprintf("; R(%v) := R(%v).. ... ..R(%v)", arga, argb, argc)
	case OP_JMP:
		buf += fmt.Sprintf("; pc+=%v", argsbx)
	case OP_EQ:
		buf += fmt.Sprintf("; if ((RK(%v) == RK(%v)) ~= %v) then pc++", argb, argc, arga)
	case OP_LT:
		buf += fmt.Sprintf("; if ((RK(%v) <  RK(%v)) ~= %v) then pc++", argb, argc, arga)
	case OP_LE:
		buf += fmt.Sprintf("; if ((RK(%v) <= RK(%v)) ~= %v) then pc++", argb, argc, arga)
	case OP_TEST:
		buf += fmt.Sprintf("; if not (R(%v) <=> %v) then pc++", arga, argc)
	case OP_TESTSET:
		buf += fmt.Sprintf("; if (R(%v) <=> %v) then R(%v) := R(%v) else pc++", argb, argc, arga, argb)
	case OP_CALL:
		buf += fmt.Sprintf("; R(%v) ... R(%v+%v-2) := R(%v)(R(%v+1) ... R(%v+%v-1))", arga, arga, argc, arga, arga, arga, argb)
	case OP_TAILCALL:
		buf += fmt.Sprintf("; return R(%v)(R(%v+1) ... R(%v+%v-1))", arga, arga, arga, argb)
	case OP_RETURN:
		buf += fmt.Sprintf("; return R(%v) ... R(%v+%v-2)", arga, arga, argb)
	case OP_FORLOOP:
		buf += fmt.Sprintf("; R(%v)+=R(%v+2); if R(%v) <?= R(%v+1) then { pc+=%v; R(%v+3)=R(%v) }", arga, arga, arga, arga, argsbx, arga, arga)
	case OP_FORPREP:
		buf += fmt.Sprintf("; R(%v)-=R(%v+2); pc+=%v", arga, arga, argsbx)
	case OP_TFORLOOP:
		buf += fmt.Sprintf("; R(%v+3) ... R(%v+3+%v) := R(%v)(R(%v+1) R(%v+2)); if R(%v+3) ~= nil then { pc++; R(%v+2)=R(%v+3); }", arga, arga, argc, arga, arga, arga, arga, arga, arga)
	case OP_SETLIST:
		buf += fmt.Sprintf("; R(%v)[(%v-1)*FPF+i] := R(%v+i) 1 <= i <= %v", arga, argc, arga, argb)
	case OP_CLOSE:
		buf += fmt.Sprintf("; close all variables in the stack up to (>=) R(%v)", arga)
	case OP_CLOSURE:
		buf += fmt.Sprintf("; R(%v) := closure(KPROTO[%v] R(%v) ... R(%v+n))", arga, argbx, arga, arga)
	case OP_VARARG:
		buf += fmt.Sprintf(";  R(%v) R(%v+1) ... R(%v+%v-1) = vararg", arga, arga, arga, argb)
	}
	return buf
}
