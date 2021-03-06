package main

import "fmt"

type List struct {
    head *Node
    tail *Node

}

func (l *List) First() *Node {
     return l.head
}

func (l *List) Last() *Node {
    return l.tail
}

func (l *List) Push(value int) {
    
     node := &Node{value: value}
     
     if l.head == nil {
         l.head = node
     } else {
         l.tail.next = node
         node.prev = l.tail
     }
     l.tail = node
}

type Node struct {
    value int
    next *Node
    prev *Node
}

func (n *Node) Next() *Node {
    return n.next
}

func (n *Node) Prev() *Node {
    return n.prev
}

func main() {

   l := &List{}
   l.Push(1)
   l.Push(2)
   l.Push(3)
   l.Push(4)
   l.Push(5)
   l.Push(6)
   
   /* 
   n := l.First()
   fmt.Println(n.value)
   n =  n.Next()
   fmt.Println(n.value) */
   
   n := l.First()
   
   for {
       fmt.Println(n.value)
       if n = n.Next(); n == nil {
           break
       }
       
   }
   
   n = l.Last()
   for {
        fmt.Println(n.value)
        n = n.Prev(); 
        if n == nil {
           break
        }
       
    }
}