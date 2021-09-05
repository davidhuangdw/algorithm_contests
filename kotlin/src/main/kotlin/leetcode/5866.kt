package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class GcdSort5866 {
    fun gcdSort(nums: IntArray): Boolean {
        val n = nums.size
        val U = nums.maxOrNull()!!
        val primes = mutableListOf<Int>()
        val comp = MutableList(U + 1) { false }
        for (x in 2..U) {
            if (!comp[x]) {
                primes.add(x)
                for (y in x * 2..U step x)
                    comp[y] = true
            }
        }

        val m = primes.size
        val par = MutableList(n + m) { it }
        fun get(i: Int): Int {
            if (par[i] != i)
                par[i] = get(par[i])
            return par[i]
        }

        fun union(i: Int, j: Int) {
            val (ri, rj) = get(i) to get(j)
            if (ri != rj)
                par[ri] = rj
        }


        for ((i, x) in nums.withIndex()) {
            for ((j, p) in primes.withIndex()) {
                if (p > x) break
                if (x % p == 0)
                    union(i, n + j)
            }
        }

        val order = (0 until n).sortedBy { nums[it] }
        for ((i, t) in order.withIndex())
            if (get(i) != get(t))
                return false
        return true
    }


    @Test
    fun test1() {
        assertEquals(true, gcdSort(parseIntArray("[7,21,3]")))
    }
}
