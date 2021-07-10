package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class Rotate48 {
    fun rotate(matrix: Array<IntArray>): Unit {
        val n = matrix.size
        for (i in 0 until n / 2)
            for (j in i until n - 1 - i) {
                var (ni, nj) = i to j
                repeat(3) {
                    ni = nj.also { nj = n - 1 - ni }
                    matrix[i][j] = matrix[ni][nj].also {
                        matrix[ni][nj] = matrix[i][j]
                    }
                }
            }
    }


    @Test
    fun test1() {
        var matrix = parse2dIntArray("[[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]")
        rotate(matrix)
        assertEquals(parse2dIntArray(
            "[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]"
        ).map { it.toList() },
            matrix.map { it.toList() }
        )

        matrix = parse2dIntArray("[[1]]")
        rotate(matrix)
        assertEquals(parse2dIntArray(
            "[[1]]"
        ).map { it.toList() },
            matrix.map { it.toList() }
        )

        matrix = parse2dIntArray("[[1,2],[3,4]]")
        rotate(matrix)
        assertEquals(parse2dIntArray(
            "[[3,1],[4,2]]"
        ).map { it.toList() },
            matrix.map { it.toList() }
        )
    }
}
