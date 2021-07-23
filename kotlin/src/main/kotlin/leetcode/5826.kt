package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class DeleteDuplicateFolder {
    fun deleteDuplicateFolder(paths: List<List<String>>): List<List<String>> {
        data class Node(var repeat: Boolean = false) {
            val ch = hashMapOf<String, Node>()
        }

        val root = Node()
        for (p in paths) {
            var cur = root
            for (x in p) {
                if (x !in cur.ch)
                    cur.ch[x] = Node()
                cur = cur.ch[x]!!
            }
        }

        val exist = hashMapOf<String, Node>()
        fun calc(u: Node): String {
            val subs = mutableListOf<String>()
            for ((k, v) in u.ch) {
                subs.add("${k}${calc(v)}")
            }
            val cur = subs.sorted().joinToString("", prefix = "(", postfix = ")")
            if (u.ch.size > 0) {
                if (cur in exist) {
                    exist[cur]!!.repeat = true
                    u.repeat = true
                } else exist[cur] = u
            }
            return cur
        }
        calc(root)

        val res = mutableListOf<List<String>>()
        val pre = mutableListOf<String>()
        fun output(u: Node) {
            if (u.repeat) return
            if (pre.size > 0)
                res.add(pre.toList())
            for ((k, v) in u.ch) {
                pre.add(k)
                output(v)
                pre.removeAt(pre.size - 1)
            }
        }
        output(root)
        return res
    }

    @Test
    fun test1() {
        assertEquals(
            parse2dStringList("[[\"b\"],[\"b\",\"w\"],[\"b\",\"z\"],[\"a\"],[\"a\",\"z\"]]").toHashSet(),
            deleteDuplicateFolder(parse2dStringList("[[\"a\"],[\"a\",\"x\"],[\"a\",\"x\",\"y\"],[\"a\",\"z\"],[\"b\"],[\"b\",\"x\"],[\"b\",\"x\",\"y\"],[\"b\",\"z\"],[\"b\",\"w\"]]")).toHashSet()
        )
        assertEquals(
            parse2dStringList("[]").toHashSet(),
            deleteDuplicateFolder(parse2dStringList("[[\"a\"],[\"a\",\"x\"],[\"a\",\"x\",\"y\"],[\"a\",\"z\"],[\"b\"],[\"b\",\"x\"],[\"b\",\"x\",\"y\"],[\"b\",\"z\"]]")).toHashSet()
        )
        assertEquals(
            parse2dStringList("[[\"d\"],[\"d\",\"a\"]]").toHashSet(),
            deleteDuplicateFolder(parse2dStringList("[[\"a\"],[\"c\"],[\"d\"],[\"a\",\"b\"],[\"c\",\"b\"],[\"d\",\"a\"]]")).toHashSet()
        )
    }
}
