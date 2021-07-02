package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class SortedListToBST {
    fun sortedListToBST(head: ListNode?): TreeNode? {
        var len = 0
        var cur = head
        while (cur != null) {
            cur = cur.next
            len += 1
        }

        if (len == 0) return null

        cur = head
        fun buildSub(k: Int): TreeNode? {   // build from cur's next k nodes
            if (k == 0) return null
            val md = k / 2
            val left = buildSub(md)
            val root = TreeNode(cur!!.`val`)
            root.left = left
            cur = cur!!.next
            root.right = buildSub(k - 1 - md)
            return root
        }
        return buildSub(len)
    }

    @Test
    fun test1() {
        val head = ListNode.fromArr(parseIntArray("[-10,-3,0,5,9]"))
        assertEquals(
            parseNullableIntList("[0,-3,9,-10,null,5]"),
            TreeNode.toLayerdOrderList(sortedListToBST(head))
        )
    }
}