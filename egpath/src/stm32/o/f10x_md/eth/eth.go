// Peripheral: ETH_Periph  Ethernet MAC.
// Instances:
//  ETH  mmap.ETH_BASE
// Registers:
//  0x00   32  MACCR
//  0x04   32  MACFFR
//  0x08   32  MACHTHR
//  0x0C   32  MACHTLR
//  0x10   32  MACMIIAR
//  0x14   32  MACMIIDR
//  0x18   32  MACFCR
//  0x1C   32  MACVLANTR   8.
//  0x28   32  MACRWUFFR   11.
//  0x2C   32  MACPMTCSR
//  0x38   32  MACSR       15.
//  0x3C   32  MACIMR
//  0x40   32  MACA0HR
//  0x44   32  MACA0LR
//  0x48   32  MACA1HR
//  0x4C   32  MACA1LR
//  0x50   32  MACA2HR
//  0x54   32  MACA2LR
//  0x58   32  MACA3HR
//  0x5C   32  MACA3LR     24.
//  0x100  32  MMCCR       65.
//  0x104  32  MMCRIR
//  0x108  32  MMCTIR
//  0x10C  32  MMCRIMR
//  0x110  32  MMCTIMR     69.
//  0x14C  32  MMCTGFSCCR  84.
//  0x150  32  MMCTGFMSCCR
//  0x168  32  MMCTGFCR
//  0x194  32  MMCRFCECR
//  0x198  32  MMCRFAECR
//  0x1C4  32  MMCRGUFCR
//  0x700  32  PTPTSCR
//  0x704  32  PTPSSIR
//  0x708  32  PTPTSHR
//  0x70C  32  PTPTSLR
//  0x710  32  PTPTSHUR
//  0x714  32  PTPTSLUR
//  0x718  32  PTPTSAR
//  0x71C  32  PTPTTHR
//  0x720  32  PTPTTLR
//  0x1000 32  DMABMR
//  0x1004 32  DMATPDR
//  0x1008 32  DMARPDR
//  0x100C 32  DMARDLAR
//  0x1010 32  DMATDLAR
//  0x1014 32  DMASR
//  0x1018 32  DMAOMR
//  0x101C 32  DMAIER
//  0x1020 32  DMAMFBOCR
//  0x1048 32  DMACHTDR
//  0x104C 32  DMACHRDR
//  0x1050 32  DMACHTBAR
//  0x1054 32  DMACHRBAR
// Import:
//  stm32/o/f10x_md/mmap
package eth

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.