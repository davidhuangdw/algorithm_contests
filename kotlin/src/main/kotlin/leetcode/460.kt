package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LFUCache460 {
    class LFUCache(val capacity: Int) {
        class DoubleLinkNode(
            val key: Int = -1,
            var v: Int = -1,
            var prev: DoubleLinkNode? = null,
            var next: DoubleLinkNode? = null,
            with_head: Boolean = false
        ) {
            var count: Int = 1

            val head: DoubleLinkNode?
            val tail: DoubleLinkNode?
            var size: Int = 0

            init {
                if (with_head) {
                    head = DoubleLinkNode(-1)
                    tail = DoubleLinkNode(-1, prev = head)
                    head.next = tail
                    tail.prev = head
                } else {
                    head = null
                    tail = null
                }
            }

            fun unlink() {
                prev!!.next = next
                next!!.prev = prev
                next = null
                prev = null
            }
        }

        val rows_head = DoubleLinkNode()
        val count_row_node = hashMapOf<Int, DoubleLinkNode>()
        val key_node = hashMapOf<Int, DoubleLinkNode>()
        var size = 0

        init {
            rows_head.next = DoubleLinkNode(prev = rows_head)
        }

        fun get(key: Int): Int {
            if (key !in key_node) return -1
            val cur = key_node[key]!!
            val prev_row = removeNode(cur)
            cur.count += 1
            insertNode(cur, prev_row)
            return cur.v
        }

        fun put(key: Int, value: Int) {
            if(capacity == 0) return
            if (key in key_node) {
                val cur = key_node[key]!!
                cur.v = value
                val prev_row = removeNode(cur)

                cur.count += 1
                insertNode(cur, prev_row)
            } else {
                val cur = DoubleLinkNode(key, value)
                key_node[cur.key] = cur

                if (size == capacity) {
                    removeOne()
                }
                insertNode(cur, rows_head)
            }
        }

        fun removeOne() {
            val row = rows_head.next!!
            val node = row.tail!!.prev!!
            key_node.remove(node.key)
            removeNode(node)
        }

        fun removeNode(cur: DoubleLinkNode): DoubleLinkNode {
            val row = count_row_node[cur.count]!!
            var prev_row = row
            cur.unlink()
            row.size -= 1
            if (row.size == 0) {
                prev_row = row.prev!!
                row.unlink()
                rows_head.size -= 1
                count_row_node.remove(cur.count)
            }
            size -= 1
            return prev_row
        }

        fun insertNode(node: DoubleLinkNode, prev_row: DoubleLinkNode) {
            if (node.count !in count_row_node) {
                val row = DoubleLinkNode(with_head = true)
                count_row_node[node.count] = row
                insertNodeToListHead(row, prev_row)
                rows_head.size += 1
            }
            val row = count_row_node[node.count]!!
            insertNodeToListHead(node, row.head!!)
            row.size += 1
            size += 1
        }

        fun insertNodeToListHead(node: DoubleLinkNode, head: DoubleLinkNode) {
            node.next = head.next
            node.prev = head
            node.next!!.prev = node
            node.prev!!.next = node
        }

        fun debug() {
            val order = mutableListOf<Pair<Int, Int>>()
            println(size)
            var row = rows_head.next!!
            repeat(rows_head.size) {
                var cur = row.head!!.next!!
                repeat(row.size) {
                    print("${cur.key}(${cur.count}) ")
                    cur = cur.next!!
                }
                row = row.next!!
                print(" | ")
            }
            println()
            println("-".repeat(100))
        }
    }
    // todo: a key insight: how to get min_freq - if min_req removed, then min_freq+=1 - no need linklist for min_freq


    @Test
    fun test1() {
        val x = LFUCache(2)
        x.debug()
        println("put(1, 1)")
        x.put(1, 1)
        x.debug()
        println("put(2, 2)")
        x.put(2, 2)
        x.debug()
        println("put(2, 2)")
        x.put(2, 2)
        x.debug()
        println("get(1)")
        assertEquals(1, x.get(1))
        x.debug()
        println("put(3, 3)")
        x.put(3, 3)
        x.debug()
        println("get(2)")
        assertEquals(-1, x.get(2))
        x.debug()
        println("get(3)")
        assertEquals(3, x.get(3))
        x.debug()
        println("put(4, 4)")
        x.put(4, 4)
        x.debug()
        println("get(1)")
        assertEquals(-1, x.get(1))
        x.debug()
        println("get(3)")
        assertEquals(3, x.get(3))
        x.debug()
        println("get(4)")
        assertEquals(4, x.get(4))
        x.debug()
    }
}
