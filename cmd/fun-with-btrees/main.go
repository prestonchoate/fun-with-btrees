package main

import (
	"fmt"

	"github.com/prestonchoate/fun-with-btrees/internal/data"
)

func main() {
  btree := data.Btree{
    Root: &data.Node{
      Value: 1,
      Left: &data.Node{
        Value: 2,
        Left: &data.Node{
          Value: 4,
          Left: nil,
          Right: nil,
        },
      },
      Right: &data.Node{
        Value: 3,
        Left: &data.Node{
          Value: 5,
          Left: nil,
          Right: nil,
        },
        Right: &data.Node{
          Value: 6,
          Left: nil,
          Right: nil,
        },
      },
    },
  }

  curQueue := data.NewQueue[data.Node]()
  nextQueue := data.NewQueue[data.Node]()
  curQueue.Enqueue(*btree.Root)

  for true {
    if curQueue.IsEmpty() && nextQueue.IsEmpty() {
      break
    }
    for !curQueue.IsEmpty() {
      leaf := curQueue.Dequeue()
      if leaf.Left != nil {
        nextQueue.Enqueue(*leaf.Left)
      }
      if leaf.Right != nil {
        nextQueue.Enqueue(*leaf.Right)
      }
      fmt.Print(leaf.Value, " ")
    }
    curQueue, nextQueue = nextQueue, curQueue
    fmt.Println()
  }
}
