
MEMORY
{
  FLASH_TEXT (rx) : ORIGIN = 0x08000000, LENGTH = 128K
  RAM (xrw)       : ORIGIN = 0x20000000, LENGTH = 36K
}

_stack_size = 2K;

INCLUDE "targets/arm.ld"
