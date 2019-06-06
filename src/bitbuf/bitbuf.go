package bitbuf;

var (
  bufByte = byte(0)
  bufSize = byte(0)
  buffer = make([]byte, 0, 1024)
)

func PushBit(val byte) {
  bufSize++
  bufByte = (bufByte<<1)|(val&1)
  if bufSize == 8 {
    buffer = append(buffer, bufByte)
    bufSize = 0
    bufByte = 0
  }
}

func PushByte(val byte) {
  for i := 7; i >= 0; i-- {
    PushBit((val&(1<<byte(i)))>>byte(i))
  }
}

func PushWord(val uint16) {
  PushByte(byte(val>>8))
  PushByte(byte(val))
}

func PushBuf(newBuf []byte) {
  for _, word := range(newBuf) {
    PushByte(word)
  }
}

func ReadBit() byte {
  if bufSize == 0 {
    if len(buffer) <= 2 {
      bufSize = buffer[1]
      bufByte = buffer[0]
      buffer = buffer[2:]
    } else {
      bufByte = buffer[0]
      bufSize = 8
      buffer = buffer[1:]
    }
  }
  bufSize--
  return (bufByte&(1<<bufSize))>>bufSize
}

func ReadByte() byte {
  res := byte(0)
  for i := 0; i < 8; i++ {
    res = (res<<1)|ReadBit()
  }
  return res
}

func ReadWord() uint16 {
  return (uint16(ReadByte())<<8)|uint16(ReadByte())
}

func ReadAll() []byte {
  buffer = append(buffer, bufByte)
  buffer = append(buffer, bufSize)
  return buffer
}

func BufSize() int {
  return len(buffer)+2
}

func HasBit() bool {
  return len(buffer) > 0 || bufSize > 0
}
