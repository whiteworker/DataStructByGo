package main
func copyRandomList(head *Node) *Node {
    res := new(Node)
    p := res
    m := map[*Node]*Node{}
    curr := head
    for curr !=nil{
        p.Next=&Node{Val:curr.Val}
        m[curr]=p.Next
        p=p.Next
        curr=curr.Next
    }
    curr = head
    p= res.Next
    for curr !=nil{
        p.Random=m[curr.Random]
        p=p.Next
        curr=curr.Next
    }
    return res.Next
}