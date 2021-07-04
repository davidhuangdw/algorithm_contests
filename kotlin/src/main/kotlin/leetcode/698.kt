package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CanPartitionKSubsets {
    fun canPartitionKSubsets(nums: IntArray, k: Int): Boolean {
        // (DP by push)idea: dp[bits] -- can be divided by 'part' and with only one group <= 'part'
        //  1. if bits deserve a try, then add a number without exceeding also deserve a try
        //  2. answer is dp[full] because last group must equals 'part'

        val n = nums.size
        nums.sort()
        val part = nums.sum() / k
        if (part * k != nums.sum() || nums[0] > part) return false

        val sum = IntArray(1 shl n)
        val dp = BooleanArray(1 shl n)
        dp[0] = true
        for(bits in 0 until (1 shl n)){
            if(!dp[bits]) continue
            for(i in 0 until n){
                if((1 shl i) and bits != 0) continue
                val nxt = bits or (1 shl i)
                if(dp[nxt]) continue
                if(sum[bits] % part + nums[i] > part)
                    break
                sum[nxt] = sum[bits] + nums[i]
                dp[nxt] = true
            }
        }
        return dp[(1 shl n) - 1]
    }

    fun canPartitionKSubsets1(nums: IntArray, k: Int): Boolean {
        val n = nums.size
        nums.sortDescending()
        val sub = nums.sum() / k
        if (sub * k != nums.sum() || nums[0] > sub) return false

        val poss_bits = List(n) { mutableListOf<Int>() }
        for (bits in 1 until (1 shl n) - 1) {
            var sum = 0
            var fst = -1
            for (i in 0 until n)
                if (bits and (1 shl i) != 0) {
                    if (fst == -1) fst = i
                    sum += nums[i]
                }
            if (sum == sub) poss_bits[fst].add(bits)
        }

        val failed = hashSetOf<Pair<Int, Int>>()
        fun test(used_bits: Int, cnt: Int): Boolean {
            if (cnt == 1) return true
            if (used_bits to cnt in failed) return false

            val low = (0 until n).first { (1 shl it) and used_bits == 0 }
            for (b in poss_bits[low]) {
                if (b and used_bits == 0 && test(used_bits or b, cnt - 1))
                    return true
            }
            failed.add(used_bits to cnt)
            return false
        }
        return test(0, k)
    }


    @Test
    fun test1() {
        assertEquals(true, canPartitionKSubsets(parseIntArray("[4, 3, 2, 3, 5, 2, 1]"), 4))
    }
}
