package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class Codec449 {
    class Codec() {
        // Encodes a URL to a shortened URL.
        fun serialize(root: TreeNode?): String {
            val list = mutableListOf<String>()
            fun ser(node: TreeNode?) {
                if (node == null) list.add("|")
                else {
                    list.add(node!!.`val`.toString())
                    ser(node.left)
                    ser(node.right)
                }
            }
            ser(root)
            return list.joinToString(",")
        }

        // Decodes your encoded data to tree.
        fun deserialize(data: String): TreeNode? {
            val list = data.split(",")
            var (i, n) = 0 to list.size
            fun des(): TreeNode? {
                i += 1
                if (list[i - 1] == "|")
                    return null
                val cur = TreeNode(list[i - 1].toInt())
                cur.left = des()
                cur.right = des()
                return cur
            }
            return des()
        }
    }


    @Test
    fun test1() {
        val x = Codec()
        assertEquals(
            parseNullableIntList("[2,1,3]"),
            x.deserialize(x.serialize(TreeNode.fromLayerOrderList(parseNullableIntList("[2,1,3]"))))!!
                .toLayerdOrderList()
        )
        assertEquals(
            null,
            x.deserialize(x.serialize(TreeNode.fromLayerOrderList(parseNullableIntList("[]"))))
        )
    }
}
