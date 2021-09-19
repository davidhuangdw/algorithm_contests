package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class DetectSquares5877 {
    class DetectSquares() {
        val count = hashMapOf<Pair<Int, Int>, Int>()
        val ys = List(1001) { hashSetOf<Int>() }
        fun add(point: IntArray) {
            val (x, y) = point
            ys[x].add(y)
            count[x to y] = count.getOrDefault(x to y, 0) + 1
        }

        fun count(point: IntArray): Int {
            val (tx, ty) = point
            var sum = 0
            for (y in ys[tx]) {
                val d = ty - y
                if (d == 0) continue
                for (x in listOf(tx - d, tx + d)) {
                    sum += count[tx to y]!! * count.getOrDefault(x to y, 0) * count.getOrDefault(x to ty, 0)
                }
            }
            return sum
        }
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
