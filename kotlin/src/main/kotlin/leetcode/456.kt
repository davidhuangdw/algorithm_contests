package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class Find132pattern {
    fun find132pattern(nums: IntArray): Boolean {
        var lmin = nums[0]
        val st = ArrayDeque<Pair<Int, Int>>()
        for (i in 1 until nums.size) {
            val v = nums[i]
            while (st.isNotEmpty() && st.last().second <= v)
                st.removeLast()
            if (st.isNotEmpty() && v in st.last().first + 1 until st.last().second)
                return true
            lmin = minOf(lmin, v)
            if (lmin < v)
                st.add(lmin to v)
        }
        return false
    }


    @Test
    fun test1() {
        assertEquals(false, find132pattern(parseIntArray("[1,2,3,4]")))
        assertEquals(true, find132pattern(parseIntArray("[3,1,4,2]")))
        assertEquals(true, find132pattern(parseIntArray("[-1,3,2,0]")))
    }
}
