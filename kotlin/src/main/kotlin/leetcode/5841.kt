package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LolutlongestObstacleCourseAtEachPositionion5841 {
    fun longestObstacleCourseAtEachPosition(obstacles: IntArray): IntArray {
        val n = obstacles.size
        val min_value = MutableList(n) { Int.MAX_VALUE }
        val longest = IntArray(n)
        for ((i, v) in obstacles.withIndex()) {
            var (l, r) = 0 to n - 1
            while (l <= r) {
                val md = (l + r) / 2
                if (min_value[md] <= v)
                    l = md + 1
                else
                    r = md - 1
            }
            min_value[l] = v
            longest[i] = l + 1
        }
        return longest
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
