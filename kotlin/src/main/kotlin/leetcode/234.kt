package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IsPalindrome234 {
    fun isPalindrome(head: ListNode?): Boolean {
        val hd = ListNode(0, head)

        var (n, cur) = 0 to hd.next
        while (cur != null) {
            n += 1
            cur = cur.next
        }

        if (n <= 1) return true

        var hf = hd
        repeat((n + 1) / 2) {
            hf = hf.next!!
        }

        var r = hf.next
        hf.next = null
        while (r != null) {
            val nr = r.next
            r.next = hf.next
            hf.next = r
            r = nr
        }

        var l = hd.next!!
        r = hf.next!!
        repeat(n / 2) {
            if (l.`val` != r!!.`val`) return false
            l = l.next!!
            r = r!!.next
        }
        return true
    }


    @Test
    fun test1() {
        assertEquals(false, isPalindrome(ListNode.fromArr(parseIntArray("[1,2]"))))
        assertEquals(true, isPalindrome(ListNode.fromArr(parseIntArray("[1,2,2,1]"))))
        assertEquals(true, isPalindrome(ListNode.fromArr(parseIntArray("[1,2,3,2,1]"))))
        assertEquals(true, isPalindrome(ListNode.fromArr(parseIntArray("[1,1]"))))
    }
}
