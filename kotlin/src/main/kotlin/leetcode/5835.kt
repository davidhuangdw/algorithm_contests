package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxMatrixSum5835 {
    fun maxMatrixSum(matrix: Array<IntArray>): Long {
        var sum = 0L
        var cnt = 0
        var min = Int.MAX_VALUE
        for (row in matrix)
            for (v in row) {
                if (v >= 0) {
                    sum += v
                    min = minOf(min, v)
                } else {
                    sum += -v
                    cnt += 1
                    min = minOf(min, -v)
                }
            }
        return sum - if (cnt % 2 == 0) 0 else 2L * min
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
