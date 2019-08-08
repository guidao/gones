package cpu

type CPU struct {
	ACC uint8
	X   uint8
	Y   uint8
	//状态寄存器
	P  uint8
	PC uint8
	SP uint8
}
