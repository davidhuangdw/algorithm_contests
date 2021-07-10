package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CombinationSum39 {
    fun combinationSum(candidates: IntArray, target: Int): List<List<Int>> {
        candidates.sort()
        val reach = List(target + 1) { hashSetOf<Int>() }
        for (i in 0..target) {
            if (i > 0 && reach[i].isEmpty()) continue
            for (c in candidates)
                if (c + i > target) break
                else reach[c + i].add(c)
        }

        val res = mutableListOf<List<Int>>()
        val cur = mutableListOf<Int>()
        fun dfs(v: Int, since: Int) {
            if (v == 0) {
                res.add(cur.toList())
            } else {
                for (c in reach[v]) {
                    if (c < since) continue
                    cur.add(c)
                    dfs(v - c, c)
                    cur.removeAt(cur.size - 1)
                }
            }
        }
        dfs(target, 0)
        return res
    }


    @Test
    fun test1() {
        assertEquals(
            parse2dIntArray("[[7],[2,2,3]]").map { it.toList() }.toSet(),
            combinationSum(parseIntArray("[2,3,6,7]"), 7).toSet()
        )
        assertEquals(
            parse2dIntArray("[[2,2,2,2],[2,3,3],[3,5]]").map { it.toList() }.toSet(),
            combinationSum(parseIntArray("[2,3,5]"), 8).toSet()
        )
    }
}
