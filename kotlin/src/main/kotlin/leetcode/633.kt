package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import kotlin.math.sqrt

class JudgeSquareSum {
    fun judgeSquareSum(c: Int): Boolean {
        var (l, r) = 0L to sqrt(c.toDouble()).toLong()
        while (l <= r) {
            val sum = l * l + r * r
            if (c - sum == 0L)
                return true
            else if (sum < c)
                l += 1
            else r -= 1
        }
        return false
    }

    fun judgeSquareSum1(c: Int): Boolean {
        var x = 0L
        val sq = hashSetOf<Long>()
        while (x * x <= c) {
            sq.add(x * x)
            x += 1
        }
        x = 0L
        while (x * x <= c) {
            if (c - x * x in sq)
                return true
            x += 1
        }
        return false
    }

    @Test
    fun test1() {
        assertEquals(true, judgeSquareSum(5))
        assertEquals(false, judgeSquareSum(3))
        assertEquals(false, judgeSquareSum(2147482647))
    }
}
