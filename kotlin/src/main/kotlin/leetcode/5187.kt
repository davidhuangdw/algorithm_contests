package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinimumPerimeter5187 {
    fun minimumPerimeter(neededApples: Long): Long {
        fun calc(n: Long) = n * (n + 1) + 2 * (n * (n + 1) * n) + n * (n + 1) * (2 * n + 1)
        var (l, r) = 1L to 1e6.toLong()
        while (l <= r) {
            val md = (l + r) / 2
            if (calc(md) >= neededApples)
                r = md - 1
            else
                l = md + 1
        }
        return l * 8
    }


    @Test
    fun test1() {
        assertEquals(5040, minimumPerimeter(1000000000L))
    }
}
