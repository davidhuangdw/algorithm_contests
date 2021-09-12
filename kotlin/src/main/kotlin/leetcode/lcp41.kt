package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FlipChessLcp41 {
    fun flipChess(chessboard: Array<String>): Int {
        val (n, m) = chessboard.size to chessboard[0].length
        fun valid(i: Int, j: Int) = i in 0 until n && j in 0 until m
        fun start(si: Int, sj: Int): Int {
            val b = (0 until n).map { chessboard[it].toMutableList() }
            b[si][sj] = 'X'
            val que = mutableListOf(si to sj)
            var cnt = 0
            while (que.isNotEmpty()) {
                val (ci, cj) = que.removeAt(que.size - 1)
                for (di in -1..1)
                    for (dj in -1..1) {
                        if (di == 0 && dj == 0) continue        // bug: di+dj == 0 !!!
                        var (i, j) = ci + di to cj + dj
                        var o = 0
                        while (valid(i, j) && b[i][j] == 'O') {
                            i += di
                            j += dj
                            o += 1
                        }
                        if (valid(i, j) && b[i][j] == 'X') {
                            cnt += o
                            repeat(o) {
                                i -= di
                                j -= dj
                                b[i][j] = 'X'
                                que.add(i to j)
                            }

                        }
                    }
            }
            return cnt
        }

        var ret = 0
        for (i in 0 until n)
            for (j in 0 until m)
                if (chessboard[i][j] == '.')
                    ret = maxOf(ret, start(i, j))
        return ret
    }


    @Test
    fun test1() {
        assertEquals(4, flipChess(arrayOf("X..X", "O.O.", "OOX.", "....", "....")))
        assertEquals(3, flipChess(arrayOf("....X.", "....X.", "XOOO..", "......", "......")))
    }
}
