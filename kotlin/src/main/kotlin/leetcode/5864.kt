package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NumberOfWeakCharacters5864 {
    fun numberOfWeakCharacters(a: Array<IntArray>): Int {
        val n = a.size
        a.sortBy { it[0] }
        var i = n - 1
        var sum = 0
        var rmax = -1
        while (i >= 0) {
            var j = i
            var sub_max = -1
            while (j >= 0 && a[j][0] == a[i][0]) {
                if (a[j][1] < rmax)
                    sum += 1
                sub_max = maxOf(sub_max, a[j][1])
                j -= 1
            }
            rmax = maxOf(rmax, sub_max)
            i = j
        }
        return sum
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
