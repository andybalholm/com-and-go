#include <runtime.h>
#include <cgocall.h>
void runtime·asmstdcall(void *c);

void ·cSyscall(WinCall *c) {
	runtime·cgocall(runtime·asmstdcall, c);
}
