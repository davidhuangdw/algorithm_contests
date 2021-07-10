package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class RemoveNthFromEnd19 {
    fun removeNthFromEnd(head: ListNode?, n: Int): ListNode? {
        val hd = ListNode(0, head)
        var a = hd
        var b = hd
        repeat(n) {
            if (b.next == null) return hd.next
            b = b.next!!
        }
        while (b.next != null) {
            a = a.next!!
            b = b.next!!
        }
        a.next = a.next!!.next
        return hd.next
    }

    @Test
    fun test1() {
        assertEquals(
            parseIntArray("[1,2,3,5]").toList(),
            ListNode.toArr(removeNthFromEnd(ListNode.fromArr(parseIntArray("[1,2,3,4,5]")), 2)).toList()
        )
        assertEquals(
            parseIntArray("[]").toList(),
            ListNode.toArr(removeNthFromEnd(ListNode.fromArr(parseIntArray("[1]")), 1)).toList()
        )
        assertEquals(
            parseIntArray("[1]").toList(),
            ListNode.toArr(removeNthFromEnd(ListNode.fromArr(parseIntArray("[1,2]")), 1)).toList()
        )
    }
}
