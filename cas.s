#include "go_asm.h"
#include "textflag.h"

// bool ·Cas64(uint64 *ptr, uint64 old, uint64 new)
// Atomically:
//      if(*val == old){
//              *val = new;
//              return 1;
//      } else {
//              return 0;
//      }
TEXT ·Cas64(SB), NOSPLIT, $0-25
	MOVD	ptr+0(FP), R0
	MOVD	old+8(FP), R1
	MOVD	new+16(FP), R2
	MOVBU	ARM64+const_offsetARM64HasATOMICS(SB), R4
	CBZ 	R4, load_store_loop
	MOVD	R1, R3
	CASALD	R3, (R0), R2
	CMP 	R1, R3
	CSET	EQ, R0
	MOVB	R0, ret+24(FP)
	RET
load_store_loop:
	LDAXR	(R0), R3
	CMP	R1, R3
	BNE	ok
	STLXR	R2, (R0), R3
	CBNZ	R3, load_store_loop
ok:
	CSET	EQ, R0
	MOVB	R0, ret+24(FP)
	RET