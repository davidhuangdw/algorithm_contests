package google_kickstart.`2017`.B

import kotlin.math.absoluteValue

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun computeMin(points: List<Pair<Double, Double>>): Double {
        if (points.size <= 1) return .0
        val diffs = points.map { it.first - points[0].first to it.second }
        val sum = diffs.sumOf { it.second } / 2
        var pre = 0.0
        var (l, r) = 0 to 0
        for ((i, d) in diffs.withIndex()) {
            pre += d.second
            if (pre <= sum) l = i
            if (pre >= sum) {
                r = i
                break
            }
        }
        fun computeAt(j: Int) = diffs.sumOf {
            (diffs[j].first - it.first).absoluteValue * it.second
        }
        return minOf(computeAt(l), computeAt(r))
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val P = (0 until N).map {
            split_input().map { s -> s.toDouble() }
        }

        fun testset2(): Double {
            val minus = P.map { (x, y, w) ->
                (x - y) / 2 to w
            }.sortedBy { it.first }

            val add = P.map { (x, y, w) ->
                (x + y) / 2 to w
            }.sortedBy { it.first }
            return computeMin(minus) + computeMin(add)
        }
        println("Case #${it}: ${testset2()}")
    }
}
