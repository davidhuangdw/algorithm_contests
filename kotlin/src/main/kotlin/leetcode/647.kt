package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CountSubstrings {
    fun countSubstrings(s: String): Int {
        val n = s.length
        var cnt = 0
        for (i in 0 until n) {
            for (d in 0..1) {
                var (l, r) = i to i + d
                while (0 <= l && r < n) {
                    if (s[l] != s[r]) break
//                    println(s.substring(l, r+1))
                    cnt += 1
                    l -= 1
                    r += 1
                }
            }

        }
        return cnt
    }

    @Test
    fun test1() {
        assertEquals(3, countSubstrings("abc"))
        assertEquals(6, countSubstrings("666"))
        assertEquals(10, countSubstrings("abcbddb"))
    }
}
