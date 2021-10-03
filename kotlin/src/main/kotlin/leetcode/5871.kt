package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class Construct2DArray5871 {
    fun construct2DArray(original: IntArray, m: Int, n: Int): Array<IntArray> {
        val l = original.size
        if (l != m * n)
            return arrayOf()
        return (0 until m).map { i -> IntArray(n) { j -> original[i * n + j] } }.toTypedArray()
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
