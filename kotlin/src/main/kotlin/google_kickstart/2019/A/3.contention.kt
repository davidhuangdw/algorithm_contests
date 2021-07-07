package google_kickstart.`2019`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    data class Point(val x: Int, val i: Int, val d: Int) : Comparable<Point> {
        override fun compareTo(other: Point): Int {
            if (x != other.x) return x - other.x
            return 0
        }
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, Q) = input_ints()
        val points = mutableListOf<Point>()
        for ((i, list) in (0 until Q).map { input_ints() }.sortedWith(compareBy({ it[1] - it[0] })).withIndex()) {
            val (l, r) = list
            points.add(Point(l - 1, i, 1))
            points.add(Point(r, i, -1))
        }
        val done = hashMapOf<Int, Int>()

        points.sort()

        var min = Int.MAX_VALUE
        repeat(Q) {
            // greedy: pick the one max{seats_if_it_is_the_last_request}
            // (the segments whose count==1) == the seats if it's the last request
            var running = hashSetOf<Int>()
            var seats = MutableList(Q) { 0 }
            var pre = -1
            for (p in points) {
                if (p.i in done) continue
                if (running.size == 1) seats[running.first()] += p.x - pre
                if (p.d > 0) {
                    running.add(p.i)
                } else {
                    running.remove(p.i)
                }
                pre = p.x
            }
            val i = (0 until Q).maxByOrNull { seats[it] }!!
            done[i] = seats[i]
            min = minOf(min, seats[i])
        }
        println("Case #${it}: $min")
    }
    // todo: testset 2
}
