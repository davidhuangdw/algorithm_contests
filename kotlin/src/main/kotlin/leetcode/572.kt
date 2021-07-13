package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IsSubtree572 {
    fun isSubtree(root: TreeNode?, subRoot: TreeNode?): Boolean {
        fun nodeToStr(x: TreeNode?): String {
            return buildString {
                fun dfs(sub: TreeNode?) {
                    if (sub == null) {
                        append("|")
                    } else {
                        append("(")
                        append(sub.`val`)
                        dfs(sub.left)
                        dfs(sub.right)
                        append(")")
                    }
                }
                dfs(x)
            }
        }

        val s = nodeToStr(root)
        val p = nodeToStr(subRoot)

        // kmp
        val n = p.length
        val max_match = MutableList(n) { 0 }
        for (i in 1 until n) {
            var l = max_match[i - 1]
            while (l > 0 && p[l] != p[i]) {
                l = max_match[l - 1]
            }
            max_match[i] = if (p[l] == p[i]) l + 1 else 0
        }

        var pre = 0
        for ((i, ch) in s.withIndex()) {
            var k = pre
            while (k > 0 && p[k] != ch) {
                k = max_match[k - 1]
            }
            k = if (p[k] == ch) k + 1 else 0
            if (k == n)
                return true
            pre = k
        }
        return false
    }


    @Test
    fun test1() {
        assertEquals(
            true, isSubtree(
                TreeNode.fromLayerOrderList(parseNullableIntList("[3,4,5,1,2]")),
                TreeNode.fromLayerOrderList(parseNullableIntList("[4,1,2]")),
            )
        )
    }
}
