package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CheckMove5827 {
    fun checkMove(board: Array<CharArray>, rMove: Int, cMove: Int, color: Char): Boolean {
        val n = 8
        fun valid(i: Int, j: Int) = i in 0 until n && j in 0 until n
        val other = if (color == 'B') 'W' else 'B'
        for ((di, dj) in listOf(0 to 1, 0 to -1, 1 to 0, -1 to 0, 1 to 1, -1 to -1, 1 to -1, -1 to 1)) {
            var cnt = 0
            var (i, j) = rMove + di to cMove + dj
            while (valid(i, j) && board[i][j] == other) {
                i += di
                j += dj
                cnt += 1
            }
            if (cnt > 0 && valid(i, j) && board[i][j] == color)
                return true
        }
        return false
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
