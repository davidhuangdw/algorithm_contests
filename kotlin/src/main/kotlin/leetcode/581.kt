package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FindUnsortedSubarray{
    fun findUnsortedSubarray(nums: IntArray): Int {
        val n = nums.size
        var lmax = Int.MIN_VALUE
        var r = -1
        for((i, v) in nums.withIndex()){
            if(lmax > v) r = i
            lmax = maxOf(lmax, v)
        }

        var rmin = Int.MAX_VALUE
        var l = n
        for(i in n-1 downTo 0){
            val v = nums[i]
            if(v > rmin) l = i
            rmin = minOf(rmin, v)
        }
//        println("$l $r")
        return maxOf(0, r-l+1)
    }

    @Test
    fun test1(){
        assertEquals(5, findUnsortedSubarray(parseIntArray("[2,6,4,8,10,9,15]")))
        assertEquals(5, findUnsortedSubarray(parseIntArray("[1, 2, 8, 5, 4, 6, 7, 9]")))
    }
}