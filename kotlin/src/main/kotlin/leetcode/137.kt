package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class SingleNumber {
    fun singleNumber(nums: IntArray): Int {
        var res = 0
        for (k in 0 until 32) {
            var cnt = 0
            for (v in nums)
                cnt += v ushr k and 1
            res += (cnt % 3) shl k
        }
        return res
    }


    @Test
    fun test1() {
        assertEquals(3, singleNumber(parseIntArray("[2,2,3,2]")))
        assertEquals(99, singleNumber(parseIntArray("[0,1,0,1,0,1,99]")))
    }
}
