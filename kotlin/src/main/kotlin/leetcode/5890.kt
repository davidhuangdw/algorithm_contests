package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinimumMoves5890 {
    fun minimumMoves(s: String): Int {
        val n = s.length
        var i = 0
        var cnt = 0
        while (i < n) {
            if (s[i] == 'X') {
                cnt += 1
                i += 2
            }
            i += 1
        }
        return cnt
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
