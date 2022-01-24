package helper

const posOp = 0
const sizeCode = 6
const sizeA = 8
const sizeB = 9
const sizeC = 9
const sizeBx = 18
const sizesBx = 18

const maxArgsA = (1 << sizeA) - 1
const maxArgsB = (1 << sizeB) - 1
const maxArgsC = (1 << sizeC) - 1
const maxArgBx = (1 << sizeBx) - 1
const maxArgSbx = maxArgBx >> 1

const (
	OP_MOVE     int = iota /*      A B     R(A) := R(B)                            */
	OP_LOADK               /*     A Bx    R(A) := Kst(Bx)                          */
	OP_LOADBOOL            /*  A B C   R(A) := (Bool)B; if (C) pc++                */
	OP_LOADNIL             /*   A B     R(A) := ... := R(B) := nil                 */
	OP_GETUPVAL            /*  A B     R(A) := UpValue[B]                          */

	OP_GETGLOBAL  /* A Bx    R(A) := Gbl[Kst(Bx)]                            */
	OP_GETTABLE   /*  A B C   R(A) := R(B)[RK(C)]                             */

	OP_SETGLOBAL  /* A Bx    Gbl[Kst(Bx)] := R(A)                            */
	OP_SETUPVAL   /*  A B     UpValue[B] := R(A)                              */
	OP_SETTABLE   /*  A B C   R(A)[RK(B)] := RK(C)                            */

	OP_NEWTABLE /*  A B C   R(A) := {} (size = BC)                         */

	OP_SELF /*      A B C   R(A+1) := R(B); R(A) := R(B)[RK(C)]             */

	OP_ADD /*       A B C   R(A) := RK(B) + RK(C)                           */
	OP_SUB /*       A B C   R(A) := RK(B) - RK(C)                           */
	OP_MUL /*       A B C   R(A) := RK(B) * RK(C)                           */
	OP_DIV /*       A B C   R(A) := RK(B) / RK(C)                           */
	OP_MOD /*       A B C   R(A) := RK(B) % RK(C)                           */
	OP_POW /*       A B C   R(A) := RK(B) ^ RK(C)                           */
	OP_UNM /*       A B     R(A) := -R(B)                                   */
	OP_NOT /*       A B     R(A) := not R(B)                                */
	OP_LEN /*       A B     R(A) := length of R(B)                          */

	OP_CONCAT /*    A B C   R(A) := R(B).. ... ..R(C)                       */

	OP_JMP /*       sBx     pc+=sBx                                 */

	OP_EQ /*        A B C   if ((RK(B) == RK(C)) ~= A) then pc++            */
	OP_LT /*        A B C   if ((RK(B) <  RK(C)) ~= A) then pc++            */
	OP_LE /*        A B C   if ((RK(B) <= RK(C)) ~= A) then pc++            */

	OP_TEST    /*      A C     if not (R(A) <=> C) then pc++                   */
	OP_TESTSET /*   A B C   if (R(B) <=> C) then R(A) := R(B) else pc++     */

	OP_CALL     /*      A B C   R(A) ... R(A+C-2) := R(A)(R(A+1) ... R(A+B-1)) */
	OP_TAILCALL /*  A B C   return R(A)(R(A+1) ... R(A+B-1))              */
	OP_RETURN   /*    A B     return R(A) ... R(A+B-2)      (see note)      */

	OP_FORLOOP /*   A sBx   R(A)+=R(A+2);
	     if R(A) <?= R(A+1) then { pc+=sBx; R(A+3)=R(A) }*/
	OP_FORPREP /*   A sBx   R(A)-=R(A+2); pc+=sBx                           */

	OP_TFORLOOP /*  A C     R(A+3) ... R(A+3+C) := R(A)(R(A+1) R(A+2));
	    if R(A+3) ~= nil then { pc++; R(A+2)=R(A+3); }  */
	OP_SETLIST /*   A B C   R(A)[(C-1)*FPF+i] := R(A+i) 1 <= i <= B        */

	OP_CLOSE   /*     A       close all variables in the stack up to (>=) R(A)*/
	OP_CLOSURE /*   A Bx    R(A) := closure(KPROTO[Bx] R(A) ... R(A+n))  */

	OP_VARARG /*     A B     R(A) R(A+1) ... R(A+B-1) = vararg            */
)
const opCodeMax = OP_VARARG

type opArgMode int

const (
	opArgModeN opArgMode = iota
	opArgModeU
	opArgModeR
	opArgModeK
)

type opType int

const (
	opTypeABC = iota
	opTypeABx
	opTypeASbx
)

type opProp struct {
	Name     string
	IsTest   bool
	SetRegA  bool
	ModeArgB opArgMode
	ModeArgC opArgMode
	Type     opType
}

var opProps = []opProp{
	opProp{"MOVE", false, true, opArgModeR, opArgModeN, opTypeABC},
	opProp{"LOADK", false, true, opArgModeK, opArgModeN, opTypeABx},
	opProp{"LOADBOOL", false, true, opArgModeU, opArgModeU, opTypeABC},
	opProp{"LOADNIL", false, true, opArgModeR, opArgModeN, opTypeABC},
	opProp{"GETUPVAL", false, true, opArgModeU, opArgModeN, opTypeABC},
	opProp{"GETGLOBAL", false, true, opArgModeK, opArgModeN, opTypeABx},
	opProp{"GETTABLE", false, true, opArgModeR, opArgModeK, opTypeABC},
	opProp{"SETGLOBAL", false, false, opArgModeK, opArgModeN, opTypeABx},
	opProp{"SETUPVAL", false, false, opArgModeU, opArgModeN, opTypeABC},
	opProp{"SETTABLE", false, false, opArgModeK, opArgModeK, opTypeABC},
	opProp{"NEWTABLE", false, true, opArgModeU, opArgModeU, opTypeABC},
	opProp{"SELF", false, true, opArgModeR, opArgModeK, opTypeABC},
	opProp{"ADD", false, true, opArgModeK, opArgModeK, opTypeABC},
	opProp{"SUB", false, true, opArgModeK, opArgModeK, opTypeABC},
	opProp{"MUL", false, true, opArgModeK, opArgModeK, opTypeABC},
	opProp{"DIV", false, true, opArgModeK, opArgModeK, opTypeABC},
	opProp{"MOD", false, true, opArgModeK, opArgModeK, opTypeABC},
	opProp{"POW", false, true, opArgModeK, opArgModeK, opTypeABC},
	opProp{"UNM", false, true, opArgModeR, opArgModeN, opTypeABC},
	opProp{"NOT", false, true, opArgModeR, opArgModeN, opTypeABC},
	opProp{"LEN", false, true, opArgModeR, opArgModeN, opTypeABC},
	opProp{"CONCAT", false, true, opArgModeR, opArgModeR, opTypeABC},
	opProp{"JMP", false, false, opArgModeR, opArgModeN, opTypeASbx},
	opProp{"EQ", true, false, opArgModeK, opArgModeK, opTypeABC},
	opProp{"LT", true, false, opArgModeK, opArgModeK, opTypeABC},
	opProp{"LE", true, false, opArgModeK, opArgModeK, opTypeABC},
	opProp{"TEST", true, true, opArgModeR, opArgModeU, opTypeABC},
	opProp{"TESTSET", true, true, opArgModeR, opArgModeU, opTypeABC},
	opProp{"CALL", false, true, opArgModeU, opArgModeU, opTypeABC},
	opProp{"TAILCALL", false, true, opArgModeU, opArgModeU, opTypeABC},
	opProp{"RETURN", false, false, opArgModeU, opArgModeN, opTypeABC},
	opProp{"FORLOOP", false, true, opArgModeR, opArgModeN, opTypeASbx},
	opProp{"FORPREP", false, true, opArgModeR, opArgModeN, opTypeASbx},
	opProp{"TFORLOOP", true, false, opArgModeN, opArgModeU, opTypeABC},
	opProp{"SETLIST", false, false, opArgModeU, opArgModeU, opTypeABC},
	opProp{"CLOSE", false, false, opArgModeN, opArgModeN, opTypeABC},
	opProp{"CLOSURE", false, true, opArgModeU, opArgModeN, opTypeABx},
	opProp{"VARARG", false, true, opArgModeU, opArgModeN, opTypeABC},
}

func OpCode(i uint32) int { return int(i >> posOp & (1 << sizeCode - 1)) }
func A(i uint32) int       { return int(i >> sizeA & maxArgsA) }
func B(i uint32) int       { return int(i >> sizeB & maxArgsB) }
func C(i uint32) int       { return int(i >> sizeC & maxArgsC) }
func Bx(i uint32) int      { return int(i >> sizeBx & maxArgBx) }
func Sbx(i uint32) int     { return int(i >> sizesBx & maxArgBx) - maxArgSbx }
