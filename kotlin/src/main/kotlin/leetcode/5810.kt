package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CanMerge {
    fun canMerge(trees: List<TreeNode?>): TreeNode? {
        val leaves = hashMapOf<Int, Pair<Int, TreeNode>>()
        val root = hashMapOf<Int, Int>()
        val remain = hashSetOf<Int>()

        val par = hashMapOf<Int, Int>()
        fun get(v: Int): Int {
            if (par[v] != v)
                par[v] = get(par[v]!!)
            return par[v]!!
        }

        for ((i, t) in trees.withIndex()) {
            if (t == null) continue
            remain.add(i)
            val v = t.`val`
            par[-i] = -i
            root[v] = i
            for (ch in listOf(t.left, t.right)) {
                if (ch == null) continue
                val cv = ch.`val`
                if (cv in par) return null
                par[cv] = -i
                if (cv in leaves) return null
                else leaves[cv] = i to ch
            }
        }

        for ((lv, i_node) in leaves) {
            if (lv !in root) continue
            val ri = root[lv]!!
            if (ri !in remain) return null

            remain.remove(ri)
            val rnode = trees[ri]!!
            if (get(lv) == get(-ri)) return null
            par[-ri] = lv
            val lnode = i_node.second
            lnode.left = rnode.left
            lnode.right = rnode.right
        }

        if (remain.size != 1) return null
        val cur_root = trees[remain.first()]!!
        var pre = 0
        fun valid(node: TreeNode?): Boolean {
            if (node == null) return true
            if (!valid(node!!.left)) return false
            if (pre >= node.`val`) return false
            pre = node.`val`
            return valid(node.right)
        }
        return if (valid(cur_root)) cur_root else null
    }


    @Test
    fun test1() {
        assertEquals(
            null, canMerge(
                listOf(
                    TreeNode.fromLayerOrderList(parseNullableIntList("[1,null,3]")),
                    TreeNode.fromLayerOrderList(parseNullableIntList("[3,1]")),
                    TreeNode.fromLayerOrderList(parseNullableIntList("[4,2]")),
                )
            )
        )

        assertEquals(
            parseNullableIntList("[3,2,5,1,null,4]"),
            canMerge(
                listOf(
                    TreeNode.fromLayerOrderList(parseNullableIntList("[2,1]")),
                    TreeNode.fromLayerOrderList(parseNullableIntList("[3,2,5]")),
                    TreeNode.fromLayerOrderList(parseNullableIntList("[5,4]")),
                )
            )!!.toLayerdOrderList()
        )
    }
}
