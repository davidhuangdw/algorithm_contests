package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class InterchangeableRectangles {
    fun interchangeableRectangles(rectangles: Array<IntArray>): Long {
        fun gcd(x: Int, y: Int): Int {
            var (a, b) = x to y
            while(b != 0){
                a = b.also{b = a % b}
            }
            return a
        }
        val count = hashMapOf<Pair<Int, Int>, Int>()
        var sum = 0L
        for((x, y) in rectangles){
            val g = gcd(x, y)
            val (a, b) = x/g to y/g
            val c = count.getOrDefault(a to b, 0)
            sum += c
            count[a to b] = c + 1
        }
        return sum
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
