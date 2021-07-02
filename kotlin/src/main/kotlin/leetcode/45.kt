package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class Jump2{
    fun jump(nums: IntArray): Int {
        var (n, cnt) = nums.size to 0
        var (l, r) = 0 to 0
        while(l<=r && r < n-1){
            cnt += 1
            val nr = (l..r).fold(Int.MIN_VALUE){pre, i -> maxOf(pre, i + nums[i])}
            l = (r+1).also{r = nr}
        }
        return cnt
    }

    @Test
    fun test1(){
        assertEquals(2, jump(parseIntArray("[2,3,1,1,4]")))
    }
}