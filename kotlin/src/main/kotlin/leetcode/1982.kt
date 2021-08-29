package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class RecoverArray1982 {
    fun recoverArray(n: Int, sums: IntArray): IntArray {
        fun divide(ss: List<Int>, x: Int): List<Int>? {
            val d = if (x > 0) x else -x
            val a = mutableListOf<Int>()
            val b = mutableListOf<Int>()
            val count = hashMapOf<Int, Int>()
            for (v in ss)
                count[v] = count.getOrDefault(v, 0) + 1
            for (v in ss) {
                if (count[v]!! > 0) {
                    if (count.getOrDefault(v + d, 0) > 0) {
                        count[v] = count[v]!! - 1
                        count[v + d] = count[v + d]!! - 1
                        if (count[v]!! < 0 || count[v + d]!! < 0)
                            return null
                        a.add(v)
                        b.add(v + d)
                    } else return null
                }
            }
            return if (x >= 0) a else b
        }

        val path = mutableListOf<Int>()
        fun dfs(ss: List<Int>): Boolean {
            if (0 !in ss) return false
            if (ss.size == 1) return true
            for (v in listOf(ss[0] - ss[1], ss[1] - ss[0])) {     // proof
                val part = divide(ss, v)
                if (part != null) {
                    path.add(v)
                    if (dfs(part))
                        return true
                    path.removeAt(path.size - 1)
                }
            }
            return false
        }
        dfs(sums.sorted())
        return path.toIntArray()
    }


    @Test
    fun test1() {
        assertEquals(listOf(-1, -2, 3), recoverArray(3, parseIntArray("[-3,-2,-1,0,0,1,2,3]")).toList())
        assertEquals(listOf(0, 0), recoverArray(2, parseIntArray("[0,0,0,0]")).toList())
        assertEquals(
            listOf(0, -1, 4, 5),
            recoverArray(4, parseIntArray("[0,0,5,5,4,-1,4,9,9,-1,4,3,4,8,3,8]")).toList()
        )
    }
}
