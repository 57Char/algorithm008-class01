package cong_wei_dao_tou_da_yin_lian_biao_lcof

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(reversePrint(head.Next), head.Val)
}

func reversePrintV2(head *ListNode) []int {
	res := []int{}
	if head == nil {
		return res
	}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
