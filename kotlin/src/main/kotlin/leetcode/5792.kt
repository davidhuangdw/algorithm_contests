package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CountTriples {
    fun countTriples(n: Int): Int {
        val sq = (1..n).map { it * it }.toHashSet()
        var cnt = 0
        for (a in 1..n) {
            for (b in 1..n)
                if (a * a + b * b in sq) cnt += 1
        }
        return cnt
    }


    @Test
    fun test1() {
        assertEquals(2, countTriples(5))
        assertEquals(4, countTriples(10))
    }
}
