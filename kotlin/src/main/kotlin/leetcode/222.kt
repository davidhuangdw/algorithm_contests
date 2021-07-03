package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CountNodes {
    fun countNodes(root: TreeNode?): Int {
        var (lh, rh) = 0 to 0
        var x = root
        while (x != null) {
            lh += 1
            x = x.left
        }

        x = root
        while (x != null) {
            rh += 1
            x = x.right
        }

        if (lh == rh) return (1 shl rh) - 1

        fun reachable(x: Int): Boolean {
            var (cur, k) = root to lh - 2
            for (k in lh - 2 downTo 0) {
                if (cur == null) return false
                cur = if (x ushr k and 1 == 0) cur.left else cur.right
            }
            return cur != null
        }
        var (l, r) = 1 shl (lh - 1) to (1 shl lh) - 1
        while (l <= r) {
            val md = (l + r) ushr 1
            if (reachable(md)) l = md + 1
            else r = md - 1
        }
        return r
    }

    @Test
    fun test1() {
        assertEquals(6, countNodes(TreeNode.fromLayerOrderList(parseNullableIntList("[1,2,3,4,5,6]"))))
        assertEquals(0, countNodes(TreeNode.fromLayerOrderList(parseNullableIntList("[]"))))
        assertEquals(1, countNodes(TreeNode.fromLayerOrderList(parseNullableIntList("[1]"))))
        assertEquals(8, countNodes(TreeNode.fromLayerOrderList(parseNullableIntList("[1,2,3,4,5,6,7,8]"))))
        assertEquals(7, countNodes(TreeNode.fromLayerOrderList(parseNullableIntList("[1,2,3,4,5,6,7]"))))
    }
}
