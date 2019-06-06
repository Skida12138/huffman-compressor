package priorityqueue;

type HasPriority interface {
  Swap(int, int)
  Higher(int, int) bool
  Len() int
}

func ShiftUp(que HasPriority, i int) {
  i++
  for i > 1 && que.Higher(i-1, i/2-1) {
    que.Swap(i-1, i/2-1)
    i /= 2
  }
}

func ShiftDown(que HasPriority, i int) {
  i++
  for {
    switch {
    case i*2-1 < que.Len() && que.Higher(i*2-1, i-1):
      que.Swap(i*2-1, i-1)
      i = i*2
    case i*2 < que.Len() && que.Higher(i*2, i-1):
      que.Swap(i*2, i-1)
      i = i*2+1
    default:
      return
    }
  }
}

func Build(que HasPriority) {
  for i := 1; i < que.Len(); i++ {
    ShiftUp(que, i)
  }
}
