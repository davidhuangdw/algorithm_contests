package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinimumDifference5897 {
    fun minimumDifference(nums: IntArray): Int {
        val n = nums.size / 2
        val orders = List(n + 1) { sortedSetOf<Int>() }
        for (bits in 0 until (1 shl n)) {
            var s = 0
            var k = 0
            for (i in 0 until n)
                if (bits and (1 shl i) > 0) {
                    s += nums[i]
                    k += 1
                }
            orders[k].add(s)
        }


        var max_low = Int.MIN_VALUE
        val sum = nums.sum()
        val hf = if (sum % 2 == 0) sum / 2 else (sum - 1) / 2       // bug: negative's floor
        for (bits in 0 until (1 shl n)) {
            var s = 0
            var k = 0
            for (i in 0 until n)
                if (bits and (1 shl i) > 0) {
                    s += nums[n + i]
                    k += 1
                }
            val fl = orders[n - k].floor(hf - s)
//            println(orders[n-k].toList())
//            println("${hf-s}, $k, $fl, $s")
            if (fl != null)
                max_low = maxOf(max_low, s + fl)
        }
        return nums.sum() - 2 * max_low
    }


    @Test
    fun test1() {
        assertEquals(2, minimumDifference(parseIntArray("[3,9,7,3]")))
        assertEquals(72, minimumDifference(parseIntArray("[-36,36]")))
        assertEquals(0, minimumDifference(parseIntArray("[2,-1,0,4,-2,-9]")))
        assertEquals(
            1,
            minimumDifference(parseIntArray("[7772197,4460211,-7641449,-8856364,546755,-3673029,527497,-9392076,3130315,-5309187,-4781283,5919119,3093450,1132720,6380128,-3954678,-1651499,-7944388,-3056827,1610628,7711173,6595873,302974,7656726,-2572679,0,2121026,-5743797,-8897395,-9699694]"))
        )
    }
}
