/**
 * Definition for singly-linked list.
 * LeetCode Problem - 83
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 

func deleteDuplicates(head *ListNode) *ListNode {
    var temp *ListNode = head
    for temp != nil{
        if temp.Next != nil && temp.Next.Val == temp.Val{
            temp.Next = temp.Next.Next
        }else{
            temp = temp.Next
        }
    }
    return head
}