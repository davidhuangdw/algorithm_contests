package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class SmallestMissingValueSubtree {
    fun smallestMissingValueSubtree(parents: IntArray, nums: IntArray): IntArray {
        val n = parents.size
        val ms = MutableList(n) { 1 }   // only need to update 1's path
        val chs = List(n) { mutableListOf<Int>() }
        for ((i, p) in parents.withIndex())
            if (p >= 0) {
                chs[p].add(i)
            }

        var u = -1
        for ((i, v) in nums.withIndex())
            if (v == 1) {
                u = i
                break
            }

        val vis = hashSetOf<Int>()
        fun collect_sub(u: Int) {
            if (nums[u] in vis) return
            vis.add(nums[u])
            for (v in chs[u])
                collect_sub(v)
        }

        var mex = 1
        while (u >= 0) {
            collect_sub(u)
            while (mex in vis)
                mex += 1
            ms[u] = mex
            u = parents[u]
        }
        return ms.toIntArray()
    }


    @Test
    fun test1() {
        assertEquals(
            listOf(5, 1, 1, 1), smallestMissingValueSubtree(
                parseIntArray("[-1,0,0,2]"), parseIntArray("[1,2,3,4]")
            ).toList()
        )
    }
}
