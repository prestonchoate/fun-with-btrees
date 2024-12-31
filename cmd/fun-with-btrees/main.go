package main

import (
	"fmt"

	"github.com/prestonchoate/fun-with-btrees/internal/data"
)

func main() {
  btree := data.Btree{
    Root: &data.Leaf{
      Value: 1,
      Left: &data.Leaf{
        Value: 2,
        Left: &data.Leaf{
          Value: 4,
          Left: nil,
          Right: nil,
        },
      },
      Right: &data.Leaf{
        Value: 3,
        Left: &data.Leaf{
          Value: 5,
          Left: nil,
          Right: nil,
        },
        Right: &data.Leaf{
          Value: 6,
          Left: nil,
          Right: nil,
        },
      },
    },
  }

  curQueue := data.NewQueue[data.Leaf]()
  nextQueue := data.NewQueue[data.Leaf]()
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
