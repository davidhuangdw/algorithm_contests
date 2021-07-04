package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class EliminateMaximum {
    fun eliminateMaximum(dist: IntArray, speed: IntArray): Int {
        val time = dist.zip(speed) { a, b -> 1.0 * a / b }.sorted()
        val n = time.size
        for (i in 1 until n) {
//            val dis = time[i] - time[i-1]
            if (time[i] < i + 1e-6)
                return i
        }
        return n
    }


    @Test
    fun test1() {
        assertEquals(3, eliminateMaximum(parseIntArray("[4,2,3]"), parseIntArray("[2,1,1]")))
        assertEquals(3, eliminateMaximum(parseIntArray("[1,3,4]"), parseIntArray("[1,1,1]")))
        assertEquals(1, eliminateMaximum(parseIntArray("[1,1, 2,3]"), parseIntArray("[1,1,1,1]")))
        assertEquals(1, eliminateMaximum(parseIntArray("[3,2,4]"), parseIntArray("[5,3,2]")))
    }
}
