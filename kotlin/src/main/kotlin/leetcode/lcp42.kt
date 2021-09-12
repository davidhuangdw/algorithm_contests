package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class CircleGameLcp42 {
    fun circleGame(toys: Array<IntArray>, circles: Array<IntArray>, r: Int): Int {
        val cs = hashSetOf<Pair<Int, Int>>()
        for ((x, y) in circles) {
            cs.add(x to y)
        }
        fun sq(x: Int) = 1L * x * x
        var cnt = 0
        for ((tx, ty, tr) in toys) {
            val diff = r - tr
            fun exist(): Boolean {
                for (x in tx - diff..tx + diff)
                    for (y in ty - diff..ty + diff)
                        if (x to y in cs && sq(x - tx) + sq(y - ty) <= sq(diff))
                            return true
                return false
            }
            if (exist()) cnt += 1
        }
        return cnt
    }


    @Test
    fun test1() {
        assertEquals(1, circleGame(parse2dIntArray("[[3,3,1],[3,2,1]]"), parse2dIntArray("[[4,3]]"), 2))
        assertEquals(2, circleGame(parse2dIntArray("[[1,3,2],[4,3,1],[7,1,2]]"), parse2dIntArray("[[1,0],[3,3]]"), 4))
    }
}
