package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class Flatten {
    fun flatten(root: TreeNode?): Unit {
        var cur = root
        while(cur != null){
            if(cur.left != null){
                var left_end: TreeNode = cur.left!!
                while(left_end.right != null)
                    left_end = left_end.right!!
                left_end.right = cur.right
                cur.right = cur.left
                cur.left = null
            }
            cur = cur.right
        }
    }

    fun flatten2(root: TreeNode?): Unit {
        val st = mutableListOf<TreeNode?>()
        var pre: TreeNode? = null
        var cur = root
        while(cur != null || st.isNotEmpty()){
            cur = if(cur == null){
                st.removeAt(st.size-1)
            }else{
                if(pre != null){
                    pre.left = null
                    pre.right = cur
                }
                pre = cur
                st.add(cur.right)
                cur.left
            }
        }
    }

    fun flatten1(root: TreeNode?): Unit {
        var pre: TreeNode? = null
        fun dfs(cur: TreeNode?){
            if(cur == null) return
            if(pre != null) {
                pre!!.right = cur
                pre!!.left = null
            }
            val (l, r) = cur.left to cur.right
            pre = cur
            dfs(l)
            dfs(r)
        }
        dfs(root)
    }


    @Test
    fun test1() {
        var root: TreeNode? = null

        root = TreeNode.fromLayerOrderList(parseNullableIntList("[1,2,5,3,4,null,6]"))
        flatten(root)
        assertEquals(parseNullableIntList("[1,null,2,null,3,null,4,null,5,null,6]"),
            root!!.toLayerdOrderList())

        root = TreeNode.fromLayerOrderList(parseNullableIntList("[]"))
        flatten(root)
        assertEquals(null, root)

        root = TreeNode.fromLayerOrderList(parseNullableIntList("[0]"))
        flatten(root)
        assertEquals(parseNullableIntList("[0]"),
            root!!.toLayerdOrderList())
    }
}
