package google_kickstart.`2021`.D

import java.util.*

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun testset1(case: Int) {
        val (N, M) = input_ints()
        val points = hashSetOf<Int>()

        repeat(N) {
            val (l, r) = input_ints()
            for (p in l..r)
                points.add(p)
        }

        val res = mutableListOf<Int>()
        val S = input_ints()
        for (s in S) {
            var x = -1
            if (s in points) {
                x = s
            } else {
                val l = (s downTo 1).firstOrNull { it in points } ?: -10000
                val r = (s..1000).firstOrNull { it in points } ?: 10000
                if (s - l <= r - s) x = l
                else x = r
            }

            points.remove(x)
            res.add(x)
        }

        println("Case #${case}: ${res.joinToString(" ")}")
    }

    fun testset2(case: Int) {
        val (N, M) = input_ints()
        val points = mutableListOf<Long>()
        val lefts = hashSetOf<Long>()
        val rights = hashSetOf<Long>()

        repeat(N) {
            val (l, r) = input_longs()
            points.add(l)
            if (l != r)
                points.add(r)
            lefts.add(l)
            rights.add(r)
        }

        val ord = points.toSortedSet() as TreeSet

        fun removeRightPoint(x: Long) {
            ord.remove(x)
            rights.remove(x)
            if (x in lefts)
                lefts.remove(x)
            else {
                //insert ,x-1]
                ord.add(x - 1)
                rights.add(x - 1)
            }
        }

        fun removeLeftPoint(x: Long) {
            ord.remove(x)
            lefts.remove(x)
            if (x in rights)
                rights.remove(x)
            else {
                //insert [x+1,
                ord.add(x + 1)
                lefts.add(x + 1)
            }
        }

        fun insertLeftPoint(x: Long) {
            assert(x !in lefts)
//            if(x !in rights)
            ord.add(x)
            lefts.add(x)
        }

        fun insertRightPoint(x: Long) {
            assert(x !in rights)
//            if(x !in lefts)
            ord.add(x)
            rights.add(x)
        }

        val res = mutableListOf<Long>()
        val S = input_longs()
        for (s in S) {
            var x = -1L
            if (s in lefts) {
                x = s
                removeLeftPoint(x)
            } else if (s in rights) {
                x = s
                removeRightPoint(x)
            } else {
                val l = ord.lower(s)
                val r = ord.higher(s)
                if (l != null && (r == null || s - l <= r - s)) {
                    if (l in rights) {
                        removeRightPoint(l)
                        x = l
                    } else {
                        x = s
                        insertRightPoint(x - 1)
                        insertLeftPoint(x + 1)
                    }
                } else {
                    if (r in lefts) {
                        x = r
                        removeLeftPoint(r)
                    } else {
                        x = s
                        insertRightPoint(x - 1)
                        insertLeftPoint(x + 1)
                    }
                }
            }
            res.add(x)
        }

        println("Case #${case}: ${res.joinToString(" ")}")
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        testset2(it)
    }
}
