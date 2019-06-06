package huffmantree;

import (
  "fmt"
  "../priorityqueue"
  "../bitbuf"
);

type TreeNode struct {
  Weight uint64
  Value byte
  Id byte
  Child0 *TreeNode
  Child1 *TreeNode
  Parent *TreeNode
}

func (node *TreeNode) ChildNum(child *TreeNode) byte {
  if node.Child0 != nil && node.Child0.Id == child.Id {
    return 0
  }
  if node.Child1 != nil && node.Child1.Id == child.Id {
    return 1
  }
  return 2
}

type TreeNodeQue []*TreeNode;

func (que TreeNodeQue) Higher(i1, i2 int) bool {
  return que[i1].Id < que[i2].Id
}

func (que TreeNodeQue) Swap(i1, i2 int) {
  que[i1], que[i2] = que[i2], que[i1]
}

func (que TreeNodeQue) Len() int {
  return len(que)
}

var fields = make(map[uint64]TreeNodeQue)

var creaNode = func() func(*TreeNode, *TreeNode, *TreeNode) *TreeNode {
  idCounter := byte(0)
  return func(child0, child1, parent *TreeNode) *TreeNode {
    newNode := TreeNode { 1, 0, idCounter, child0, child1, parent }
    if field, ok := fields[1]; ok {
      fields[1] = append(field, &newNode)
      priorityqueue.ShiftUp(fields[1], fields[1].Len()-1)
    } else {
      fields[1] = make(TreeNodeQue, 0, 1)
      fields[1] = append(fields[1], &newNode)
    }
    idCounter++
    return &newNode
  }
}()

func updNode(node *TreeNode) {
  switch {
  case node.Id == 255:
    parent := creaNode(node, nil, node.Parent)
    child1 := creaNode(nil, nil, parent)
    node.Parent = parent
    parent.Child1 = child1
    if parent.Parent != nil {
      if parent.Parent.Child0.Id == 255 {
        parent.Parent.Child0 = parent
        updNode(parent.Parent)
      } else {
        parent.Parent.Child1 = parent
        updNode(parent.Parent)
      }
    }
  case fields[node.Weight].Len() > 0 && fields[node.Weight][0].Id != node.Id && fields[node.Weight][0].ChildNum(node) == 2:
    field := fields[node.Weight]
    if node.Parent != nil {
      if node.Parent.Child0.Id == node.Id {
        node.Parent.Child0 = field[0]
      } else {
        node.Parent.Child1 = field[0]
      }
    }
    if field[0].Parent != nil {
      if field[0].Parent.Child0.Id == field[0].Id {
        field[0].Parent.Child0 = node
      } else {
        field[0].Parent.Child1 = node
      }
    }
    node.Parent, field[0].Parent = field[0].Parent, node.Parent
    node.Id, field[0].Id = field[0].Id, node.Id
    fallthrough
  case fields[node.Weight].Len() > 0 && fields[node.Weight][0].ChildNum(node) != 2:
    for i, val := range(fields[node.Weight]) {
      if val.Id == node.Id {
        fields[node.Weight].Swap(i, fields[node.Weight].Len()-1)
        fields[node.Weight] = fields[node.Weight][:fields[node.Weight].Len()-1]
        priorityqueue.ShiftDown(fields[node.Weight], i)
        break
      }
    }
    priorityqueue.ShiftDown(fields[node.Weight], 0)
    fallthrough
  default:
    if node.Id == fields[node.Weight][0].Id {
      fields[node.Weight].Swap(0, fields[node.Weight].Len()-1)
      fields[node.Weight] = fields[node.Weight][:fields[node.Weight].Len()-1]
      priorityqueue.ShiftDown(fields[node.Weight], 0)
    }
    node.Weight++
    fields[node.Weight] = append(fields[node.Weight], node)
    priorityqueue.ShiftUp(fields[node.Weight], fields[node.Weight].Len()-1)
    if node.Parent != nil {
      updNode(node.Parent)
    }
  }
}

func outputNode(node *TreeNode) {
  if node.Parent != nil {
    outputNode(node.Parent)
    bitbuf.PushBit(node.Parent.ChildNum(node))
  }
}

func Compress(cont []byte) []byte {
  nyt := TreeNode { 0, 0, 255, nil, nil, nil }
  vised := make(map[byte]*TreeNode)
  tot := len(cont)
  for prog, val := range(cont) {
    fmt.Printf("\rCompressing file...%d%% finished", prog*100/tot)
    if node, ok := vised[val]; ok {
      outputNode(node)
      updNode(node)
    } else {
      outputNode(&nyt)
      updNode(&nyt)
      bitbuf.PushByte(val)
      vised[val] = nyt.Parent.Child1
    }
  }
  compRat := float64(bitbuf.BufSize()*100)/float64(tot)
  fmt.Printf("\rCompression finished, compression rate: %.2f%%\n", compRat)
  return bitbuf.ReadAll()
}

func Decompress(cont []byte) []byte {
  root := &TreeNode { 0, 0, 255, nil, nil, nil }
  res := make([]byte, 0, 1024)
  curNode := root
  bitbuf.PushBuf(cont)
  prog := 0
  tot := len(cont)
  for {
    fmt.Printf("\rExtracting file...%d%% finished", prog*25/tot/2)
    if curNode.Id == 255 {
      word := bitbuf.ReadByte()
      res = append(res, word)
      prog += 8
      updNode(curNode)
      if root.Id == 255 {
        root = root.Parent
      }
      curNode.Parent.Child1.Value = word
      curNode = root
    } else if curNode.Child0 == nil && curNode.Child1 == nil {
      res = append(res, curNode.Value)
      updNode(curNode)
      curNode = root
    } else if !bitbuf.HasBit() {
      break
    } else if bitbuf.ReadBit() > 0 {
      curNode = curNode.Child1
      prog++
    } else {
      curNode = curNode.Child0
      prog++
    }
  }
  fmt.Printf("\rExtraction finished           \n")
  return res
}
