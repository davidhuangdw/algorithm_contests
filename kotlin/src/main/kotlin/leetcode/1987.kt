package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaximumRemovals {
    fun maximumRemovals(s: String, p: String, removable: IntArray): Int {
        val n = p.length
        val pre = MutableList(n) { 0 }

        for (i in 1 until n) {
            var k = pre[i - 1]
            while (k > 0 && p[k] != p[i])
                k = pre[k - 1]
            pre[i] = if (p[k] == p[i]) k + 1 else 0
        }

        fun found(k: Int): Boolean {
            if (s.length - k < p.length)
                return false

            val rm = removable.slice(0 until k).toHashSet()
            var j = 0
            for ((i, ch) in s.withIndex()) {
                if (i in rm) continue
                if (ch == p[j]) {
                    j += 1
                    if (j == p.length) return true
                }
            }
            return false
        }

        var (l, r) = 1 to removable.size
        while (l <= r) {
            val md = (l + r) ushr 1
            if (found(md)) l = md + 1
            else r = md - 1
        }
        return r
    }


    @Test
    fun test1() {
        assertEquals(2, maximumRemovals("abcacb", "ab", parseIntArray("[3,1,0]")))
        assertEquals(1, maximumRemovals("abcbddddd", "abcd", parseIntArray("[3,2,1,4,5,6]")))
        assertEquals(0, maximumRemovals("abcab", "abc", parseIntArray("[0,1,2,3,4]")))

    }
}
