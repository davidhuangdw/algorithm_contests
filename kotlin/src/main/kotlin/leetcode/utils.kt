package leetcode

import org.junit.jupiter.api.Assertions.assertArrayEquals
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*
import kotlin.collections.ArrayDeque

class SortedMultiSet<T>(val map: TreeMap<T, Int> = TreeMap()): MutableSet<T> by map.keys{
    override fun add(element: T): Boolean {
        map[element] = map.getOrDefault(element, 0) + 1
        return true
    }

    override fun remove(element: T): Boolean {
        if(element !in map) return false
        val cnt = map[element]!!
        if(cnt == 1) map.remove(element)
        else map[element] = cnt - 1
        return true
    }
}

class ListNode(var `val`: Int) {
    var next: ListNode? = null

    companion object {
        fun fromArr(arr: IntArray): ListNode? {
            var last: ListNode? = null
            for (v in arr.reversed()) {
                val x = ListNode(v)
                x.next = last
                last = x
            }
            return last
        }

        fun toArr(head: ListNode?): IntArray {
            val arr = mutableListOf<Int>()
            var cur = head
            while (cur != null) {
                arr.add(cur.`val`)
                cur = cur.next
            }
            return arr.toIntArray()
        }
    }
}

class TreeNode(var `val`: Int) {
    var left: TreeNode? = null
    var right: TreeNode? = null

    companion object {
        fun fromPreOrderList(arr: List<Int?>): TreeNode? {
            var (i, n) = 0 to arr.size
            fun build(): TreeNode? = if (i < n && arr[i] != null) {
                val root = TreeNode(arr[i]!!)
                i += 1
                root.left = build()
                root.right = build()
                root
            } else {
                i += 1
                null
            }
            return build()
        }

        fun toPreOrderList(root: TreeNode?): List<Int?> {
            val arr = mutableListOf<Int?>()
            fun dfs(x: TreeNode?) {
                if (x == null) {
                    arr.add(null)
                } else {
                    arr.add(x.`val`)
                    dfs(x.left)
                    dfs(x.right)
                }
            }
            dfs(root)
            while (arr.isNotEmpty() && arr.last() == null)
                arr.removeLast()
            return arr
        }

        fun fromLayerOrderList(arr: List<Int?>): TreeNode? {
            var n = arr.size
            fun build(i: Int): TreeNode? = if (i < n && arr[i] != null) {
                val root = TreeNode(arr[i]!!)
                root.left = build(i * 2 + 1)
                root.right = build(i * 2 + 2)
                root
            } else null
            return build(0)
        }

        fun toLayerdOrderList(root: TreeNode?): List<Int?> {
            val arr = mutableListOf<Int?>()
            var que = ArrayDeque<TreeNode?>()
            que.add(root)
            while (que.isNotEmpty()) {
                val x = que.removeFirst()
                if (x == null)
                    arr.add(null)
                else {
                    arr.add(x.`val`)
                    que.add(x.left)
                    que.add(x.right)
                }
            }
            while (arr.isNotEmpty() && arr.last() == null)
                arr.removeLast()
            return arr
        }
    }
}

fun parseArray(s: String): List<String> {
    val input = s.trim()
    val subs = mutableListOf<String>()
    var extra = 0
    var l = 0
    for ((i, ch) in input.withIndex()) {
        when (ch) {
            '[' -> {
                extra += 1
                if (extra == 1) l = i + 1
            }
            ']' -> {
                extra -= 1
                if (extra == 0) {
                    if(l != i)
                        subs.add(input.substring(l, i).trim())
                    l = i + 1
                }
            }
            ',' -> if (extra == 1) {
                subs.add(input.substring(l, i).trim())
                l = i + 1
            }
        }
    }
    return subs
}

inline fun <reified T> parseNullableTypedList(s: String, convert: (String) -> T) =
    parseArray(s).map { if (it.trim() == "null") null else convert(it) }

inline fun <reified T> parseTypedArray(s: String, convert: (String) -> T): Array<T> =
    parseNullableTypedList(s, convert).map { it!! }.toTypedArray()

fun parseNullableIntList(s: String) = parseNullableTypedList(s, String::toInt)


fun parseIntArray(s: String) = parseTypedArray(s, String::toInt).toIntArray()
fun parse2dIntArray(s: String) = parseTypedArray(s) { parseIntArray(it) }

fun parseLongArray(s: String) = parseTypedArray(s, String::toLong).toLongArray()
fun parse2dLongArray(s: String) = parseTypedArray(s) { parseLongArray(it) }

fun parseDoubleArray(s: String) = parseTypedArray(s, String::toDouble).toDoubleArray()
fun parse2dDoubleArray(s: String) = parseTypedArray(s) { parseDoubleArray(it) }

class TestLeetUtils {
    @Test
    fun test1() {
        assertArrayEquals(intArrayOf(1, 2, 3, 4), parseIntArray("  [1, 2, 3, 4] "))

        val input = " [ [0, 1, 2], [13, 14, 15 ], [16, 17, 18] ] "
        var expect = arrayOf(
            intArrayOf(0, 1, 2),
            intArrayOf(13, 14, 15),
            intArrayOf(16, 17, 18),
        )
        for ((a, b) in expect.zip(parse2dIntArray(input)))
            assertArrayEquals(a, b)
    }

    @Test
    fun testListNode() {
        val arr = parseIntArray(" [1, 3, 5, 7] ")
        println(arr.toList())
        val head = ListNode.fromArr(arr)
        val converted = ListNode.toArr(head)
        assertArrayEquals(arr, converted)
    }

    @Test
    fun testTreeNode() {
        val list = parseNullableIntList("[0,-3,9,-10,null,5]")
        println(list)
        var root: TreeNode? = null
        root = TreeNode.fromPreOrderList(list)
        assertEquals(list, TreeNode.toPreOrderList(root))

        root = TreeNode.fromLayerOrderList(list)
        assertEquals(list, TreeNode.toLayerdOrderList(root))

    }

    @Test
    fun testSortedMultiSet(){
        val x = SortedMultiSet<Int>()
        val vs = listOf(5, 2, 3, 1, 1, 2, 4, 3)
        for(v in vs) x.add(v)
        x.remove(1)
        assertEquals(1, x.minOrNull())
        x.remove(1)
        assertEquals(2, x.minOrNull())
        x.remove(3)
        x.remove(2)
        assertEquals(2, x.minOrNull())
        x.remove(2)
        assertEquals(3, x.minOrNull())
        x.remove(3)
        assertEquals(4, x.minOrNull())
        x.remove(4)
        assertEquals(5, x.minOrNull())
    }
}
