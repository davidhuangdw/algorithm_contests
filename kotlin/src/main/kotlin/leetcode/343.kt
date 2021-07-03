package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IntegerBreak {
    fun integerBreak(n: Int): Int {
        var max = IntArray(n) { it }
        for (i in 2 until n)
            for (x in 1..i / 2)
                max[i] = maxOf(max[i], max[x] * max[i - x])
        return (1..n / 2).fold(0) { pre, i -> maxOf(pre, max[i] * max[n - i]) }
    }

    @Test
    fun test1() {
        assertEquals(36, integerBreak(10))
        assertEquals(1, integerBreak(2))
    }
}
