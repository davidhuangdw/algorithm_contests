package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CountPalindromicSubsequence {
    fun countPalindromicSubsequence(s: String): Int {
        val n = s.length
        val left = MutableList(n) { hashSetOf<Char>() }

        var cur = hashSetOf<Char>()
        for ((i, ch) in s.withIndex()) {
            left[i] = cur.toHashSet()
            cur.add(ch)
        }

        var res = hashSetOf<Pair<Char, Char>>() // middle to side
        cur.clear()
        for (i in n - 1 downTo 0) {
            val ch = s[i]
            for (both in cur.intersect(left[i]))
                res.add(ch to both)
            cur.add(s[i])
        }

        return res.size
    }


    @Test
    fun test1() {
        assertEquals(3, countPalindromicSubsequence("aabca"))
        assertEquals(4, countPalindromicSubsequence("bbcbaba"))
    }
}
