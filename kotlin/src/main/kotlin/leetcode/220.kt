package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ContainsNearbyAlmostDuplicate {
    fun containsNearbyAlmostDuplicate(nums: IntArray, k: Int, t: Int): Boolean {
        fun abs(x: Long) = if (x >= 0) x else -x
        fun getKey(v: Int) = if (t == 0) v else v / t
        val buckets = hashMapOf<Int, Int>()
        for ((i, v) in nums.withIndex()) {
            val key = getKey(v)
            for (d in -1..1)
                if (key + d in buckets && abs(0L + buckets[key + d]!! - v) <= t)
                    return true
            buckets[key] = v
            if (i - k >= 0)
                buckets.remove(getKey(nums[i - k]))
        }
        return false
    }


    @Test
    fun test1() {
        assertEquals(false, containsNearbyAlmostDuplicate(parseIntArray("[1,5,9,1,5,9]"), 2, 3))
        assertEquals(true, containsNearbyAlmostDuplicate(parseIntArray("[1,2,3,1]"), 3, 0))
        assertEquals(true, containsNearbyAlmostDuplicate(parseIntArray("[1,0,1,1]"), 1, 2))
        assertEquals(false, containsNearbyAlmostDuplicate(parseIntArray("[1,2]"), 0, 1))
    }
}
