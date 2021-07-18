package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxGeneticDifference {
    fun maxGeneticDifference(parents: IntArray, queries: Array<IntArray>): IntArray {
        val n = parents.size
        val node_to_qi = List(n) { mutableListOf<Int>() }
        for ((i, q) in queries.withIndex())
            node_to_qi[q[0]].add(i)

        var rooti = 0
        val children = List(n) { hashSetOf<Int>() }
        for ((i, p) in parents.withIndex()) {
            if (p < 0) rooti = i
            else children[p].add(i)
        }

        data class Node(var cnt: Int = 0) {
            val nxt = MutableList<Node?>(2) { null }
            fun add(x: Int, k: Int, diff: Int = 1) {
                cnt += diff
                if (k >= 0) {       // bug!!: not "k-1 >= 0"
                    val b = x ushr k and 1
                    if (nxt[b] == null)
                        nxt[b] = Node()
                    nxt[b]!!.add(x, k - 1, diff)
                }
            }

            fun exist(k: Int, bit: Int) = nxt[bit] != null && nxt[bit]!!.cnt > 0
        }

        val root = Node(0)


        val res = MutableList(queries.size) { 0 }
        val K = 22
        fun dfs(i: Int) {
            // insert i
            root.add(i, K - 1, 1)

            for (qi in node_to_qi[i]) {
                var cur = 0
                val v = queries[qi][1]
                var node = root
                for (k in K - 1 downTo 0) {
                    cur = cur shl 1
                    val b = v shr k and 1 xor 1
                    if (node.exist(k, b)) {
                        node = node.nxt[b]!!
                        cur += 1
                    } else {
                        node = node.nxt[b xor 1]!!
                    }
                }
                res[qi] = cur
            }
            for (j in children[i])
                dfs(j)

            // remove i
            root.add(i, K - 1, -1)
        }
        dfs(rooti)
        return res.toIntArray()
    }


    @Test
    fun test1() {
        assertEquals(
            parseIntArray("[2,3,7]").toList(),
            maxGeneticDifference(parseIntArray("[-1,0,1,1]"), parse2dIntArray("[[0,2],[3,2],[2,5]]")).toList()
        )
        assertEquals(
            parseIntArray("[6,14,7]").toList(),
            maxGeneticDifference(parseIntArray("[3,7,-1,2,0,7,0,2]"), parse2dIntArray("[[4,6],[1,15],[0,5]]")).toList()
        )
    }

}
