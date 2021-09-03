package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class rightSideView199 {
    fun rightSideView(root: TreeNode?): List<Int> {
        val view = mutableListOf<Int>()
        val que = ArrayDeque<Pair<Int, TreeNode>>()
        if(root != null) que.add(1 to root)
        while(que.isNotEmpty()){
            val (h, u) = que.removeFirst()
            if(view.size < h)
                view.add(u.`val`)
            for(ch in listOf(u.right, u.left)){
                if(ch != null)
                    que.add(h+1 to ch)
            }
        }
        return view
    }

    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
