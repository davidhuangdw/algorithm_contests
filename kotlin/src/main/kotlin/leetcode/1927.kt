package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class SumGame1927 {
    fun sumGame(num: String): Boolean {
        val n = num.length
        var (diff, m) = 0 to 0
        for (i in 0 until n) {
            val neg = if (i < n / 2) 1 else -1
            if (num[i] == '?')
                m -= neg
            else
                diff += neg * (num[i] - '0')
        }
        return !(m % 2 == 0 && diff == m / 2 * 9)
    }


    @Test
    fun test1() {
        assertEquals(false, sumGame("5023"))
        assertEquals(false, sumGame("?3295???"))
        assertEquals(true, sumGame("25??"))
    }
}
