package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LockingTree5848 {
    class LockingTree(parent: IntArray) {
        val par = parent
        val n = parent.size
        val ch = List(n) { mutableListOf<Int>() }

        init {
            for ((i, p) in parent.withIndex()) {
                if (p != -1)
                    ch[p].add(i)
            }
        }

        val id = MutableList(n) { -1 }

        fun lock(num: Int, user: Int): Boolean {
            return if (id[num] == -1) {
                id[num] = user
                true
            } else false
        }

        fun unlock(num: Int, user: Int): Boolean {
            return if (id[num] == user) {
                id[num] = -1
                true
            } else false
        }

        fun upgrade(num: Int, user: Int): Boolean {
            if (id[num] != -1)
                return false
            val que = mutableListOf(num)
            var p = par[num]
            while (p != -1) {
                if (id[p] != -1)
                    return false
                p = par[p]
            }

            var found = false
            while (que.isNotEmpty()) {
                val u = que.removeAt(que.size - 1)
                if (id[u] != -1) {
                    found = true
                    id[u] = -1
                }
                for (v in ch[u])
                    que.add(v)
            }
            return if (found) {
                id[num] = user
                true
            } else false

        }
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
