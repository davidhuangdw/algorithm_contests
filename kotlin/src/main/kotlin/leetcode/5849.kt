package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NumberOfGoodSubsets5849 {
    fun numberOfGoodSubsets(nums: IntArray): Int { // todo: fix it
        fun gcd(a: Int, b: Int): Int {
            var (x, y) = a to b
            while (y > 0) {
                x = y.also { y = x % y }
            }
            return x
        }

        val MOD = 1000_000_007
        val count = MutableList(31) { 0L }
        for (x in nums)
            count[x] += 1L
        for (v in 2..5) {
            var x = v * v
            for (y in x..30 step x)
                count[y] = 0
        }
        val vs = (2..30).filter { count[it] > 0 }
        val m = vs.size
        val conflicts = List(m) { mutableListOf<Int>() }
        for ((i, x) in vs.withIndex()) {
            for (j in 0 until i)
                if (gcd(x, vs[j]) > 1)
                    conflicts[i].add(j)
        }

        var sum = 0L
        for (bits in 1 until (1 shl m)) {
            val sub = mutableListOf<Int>()
            var invalid = false
            for (i in 0 until m)
                if (bits and (1 shl i) > 0) {
                    for (j in conflicts[i])
                        if (bits and (1 shl j) > 0) {
                            invalid = true
                            break
                        }
                    if (invalid) break
                    sub.add(vs[i])
                }
            if (invalid) continue
            var s = 1L
            for (x in sub)
                s = s * count[x] % MOD
            sum = (sum + s) % MOD
        }

        repeat(count[1].toInt()) {
            sum = sum * 2 % MOD
        }
        return sum.toInt()
    }


    @Test
    fun test1() {
        assertEquals(6, numberOfGoodSubsets(parseIntArray("[1,2,3,4]")))
        assertEquals(5, numberOfGoodSubsets(parseIntArray("[4,2,3,15]")))
        assertEquals(
            5368,
            numberOfGoodSubsets(parseIntArray("[5,10,1,26,24,21,24,23,11,12,27,4,17,16,2,6,1,1,6,8,13,30,24,20,2,19]"))
        )
    }
}
