package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FindKthNumber440 {
    fun findKthNumber(n: Int, k: Int): Int {
        fun get_sub_count(pre: Int): Int {
            var cur = pre.toLong()
            var level = 1L
            var cnt = 0L
            while (cur <= n) {      // bug: cur maybe 10*n-1, might overflow
                cnt += minOf(level, n - cur + 1)
                level *= 10
                cur *= 10
            }
            return cnt.toInt()
        }

        var cur = 1
        var left = k - 1
        while (left > 0) {
            val sub = get_sub_count(cur)    // diff_count to right tree 'cur+1'
            if (sub <= left) {
                left -= sub
                cur += 1
            } else {
                left -= 1       // go down used up 1
                cur *= 10
            }
        }
        return cur
    }


    @Test
    fun test1() {
        assertEquals(1, findKthNumber(13, 1))
        assertEquals(35, findKthNumber(100, 30))
        assertEquals(10, findKthNumber(13, 2))
        assertEquals(416126219, findKthNumber(681692778, 351251360))
    }
}
