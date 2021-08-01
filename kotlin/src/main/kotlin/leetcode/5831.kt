package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NumberOfWeeks5831 {
    fun numberOfWeeks(milestones: IntArray): Long {
        val a = milestones.map { it.toLong() }
        var sum = a.sum()
        val x = a.maxOrNull()!!
        val y = sum - x

        return sum - maxOf(0, x - y - 1)
    }


    @Test
    fun test1() {
        assertEquals(97, numberOfWeeks(parseIntArray("[1,10,7,1,7,2,10,10,355359359]")))
    }
}
