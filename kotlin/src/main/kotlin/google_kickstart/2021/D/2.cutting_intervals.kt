package google_kickstart.`2021`.D

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, C) = input_longs()

        data class Point(val x: Long, val left: Boolean) : Comparable<Point> {
            override fun compareTo(other: Point): Int {
                if (x != other.x) return compareValues(x, other.x);
                return if (!left || other.left) -1 else 1
            }

        }

        val points = mutableListOf<Point>()

        repeat(N.toInt()) {
            val l = input_longs()
            points.add(Point(l[0], true))
            points.add(Point(l[1], false))
        }

        points.sort()


        val cut_len = mutableListOf<Pair<Long, Long>>()  // count to length
        var i = 0
        val n = points.size
        var cur = 0L
        var pre_x = points[0].x
        while (i < n) {
            val x = points[i].x
            var j = i
            var (st, ed) = 0 to 0
            while (j < n && points[j].x == x) {
                if (points[j].left) st += 1
                else ed += 1
                j += 1
            }
            if (x != pre_x && cur > 0) {
                cut_len.add(cur to (x - pre_x - 1))
//                println("$x: ${cut_len.last()}")
                cut_len.add(cur - ed to 1)
//                println("$x: ${cut_len.last()}")
            }
            cur -= ed
            cur += st
            i = j
            pre_x = x
        }

        var rem = C
        var res = N
        for ((count, len) in cut_len.sortedBy { -it.first }) {
            val integers = minOf(rem, len)
            res += integers * count
            rem -= integers
            if (rem == 0L) break
        }
        println("Case #${it}: ${res}")
    }

}
