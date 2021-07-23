package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxCompatibilitySum {
    fun maxCompatibilitySum(students: Array<IntArray>, mentors: Array<IntArray>): Int {
        val m = students.size
        val match = List(m) { MutableList(m) { 0 } }
        for (i in 0 until m)
            for (j in 0 until m) {
                for ((a, b) in students[i].zip(mentors[j]))
                    if (a == b) match[i][j] += 1
            }

        val assign = (0 until m).toMutableList()
        fun subMax(i: Int): Int {
            var sm = 0
            for (j in i until m) {
                assign[i] = assign[j].also { assign[j] = assign[i] }
                sm = maxOf(sm, match[i][assign[i]] + subMax(i + 1))
                assign[i] = assign[j].also { assign[j] = assign[i] }
            }
            return sm
        }
        return subMax(0)
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
