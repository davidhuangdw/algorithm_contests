package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinNumber45 {
    fun minNumber(nums: IntArray): String {
        return nums.map{it.toString()}.sortedWith(Comparator(){ o1, o2 ->
            compareValues(o1+o2, o2+o1)
        }).joinToString("")
    }

    @Test
    fun test1() {
        assertEquals("3033459", minNumber(parseIntArray("[3,30,34,5,9]")))
    }
}
