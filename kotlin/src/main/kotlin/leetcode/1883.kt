package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinSkips {
    fun minSkips(dist: IntArray, speed: Int, hoursBefore: Int): Int {
        val target_time = 1L * hoursBefore * speed
        val n = dist.size
        val min = MutableList(n + 1) { 0L }

        fun ceil(x: Long) = (x + speed - 1) / speed * speed

        for (d in dist)
            for (k in n downTo 0) {
                min[k] = ceil(min[k] + d)
                if (k - 1 >= 0)
                    min[k] = minOf(min[k], min[k - 1] + d)
            }
        for (i in 0..n)
            if (min[i] <= target_time)
                return i
        return -1
    }


    @Test
    fun test1() {
        assertEquals(-1, minSkips(parseIntArray("[7,3,5,5]"), 1, 10))
        assertEquals(1, minSkips(parseIntArray("[1,3,2]"), 4, 2))
        assertEquals(2, minSkips(parseIntArray("[7,3,5,5]"), 2, 10))
    }
}
