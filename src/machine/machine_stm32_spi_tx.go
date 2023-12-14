//go:build stm32 && stm32g0

package machine

import (
	"device/stm32"
	"runtime/volatile"
	"unsafe"
)

// Peripheral abstraction layer for SPI on the stm32 family

// Tx handles read/write operation for SPI interface. Since SPI is a syncronous write/read
// interface, there must always be the same number of bytes written as bytes read.
// The Tx method knows about this, and offers a few different ways of calling it.
//
// This form sends the bytes in tx buffer, putting the resulting bytes read into the rx buffer.
// Note that the tx and rx buffers must be the same size:
//
//	spi.Tx(tx, rx)
//
// This form sends the tx buffer, ignoring the result. Useful for sending "commands" that return zeros
// until all the bytes in the command packet have been received:
//
//	spi.Tx(tx, nil)
//
// This form sends zeros, putting the result into the rx buffer. Good for reading a "result packet":
//
//	spi.Tx(nil, rx)
func (spi SPI) Tx(tx, rx []byte) error {
	nw := len(tx)
	nr := len(rx)
	ntx := nw
	if ntx < nr {
		ntx = nr
	}

	sr := &spi.Bus.SR
	dr := (*volatile.Register8)(unsafe.Pointer(&spi.Bus.DR.Reg))
	iw := 0
	ir := 0
	var b byte
	for {
		// Try to write a byte, if buffer signals that it can be written to.
		for ntx > 0 && sr.HasBits(stm32.SPI_SR_TXE) {
			if iw < nw {
				dr.Set(tx[iw])
				iw++
			} else {
				dr.Set(0)
			}
			ntx--
			for sr.HasBits(stm32.SPI_SR_RXNE) {
				b = byte(dr.Get())
				if ir < nr {
					rx[ir] = b
					ir++
				}
			}
		}
		for sr.HasBits(stm32.SPI_SR_RXNE) {
			b = byte(dr.Get())
			if ir < nr {
				rx[ir] = b
				ir++
			}
		}
		if ntx == 0 && ir == nr {
			break
		}
	}

	// wait for SPI bus busy bit (BSY) to be clear to indicate synchronous
	// transfer complete. this will effectively prevent this Transfer() function
	// from being capable of maintaining high-bandwidth communication throughput,
	// but it will help guarantee stability on the bus.
	for spi.Bus.SR.HasBits(stm32.SPI_SR_BSY) {
	}

	// clear the overrun flag (only in full-duplex mode)
	if !spi.Bus.CR1.HasBits(stm32.SPI_CR1_RXONLY | stm32.SPI_CR1_BIDIMODE | stm32.SPI_CR1_BIDIOE) {
		spi.Bus.SR.Get()
	}
	return nil
}
