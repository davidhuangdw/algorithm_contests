package google_kickstart.`2017`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun abs(a: Long) = if (a >= 0) a else -a
    data class Point(val x: Long, val y: Long, val z: Long) {
        fun minWith(other: Point) = Point(minOf(x, other.x), minOf(y, other.y), minOf(z, other.z))
        fun maxWith(other: Point) = Point(maxOf(x, other.x), maxOf(y, other.y), maxOf(z, other.z))
        fun corners(other: Point): List<Point> {
            val ps = mutableListOf<Point>()
            for (a in listOf(x, other.x))
                for (b in listOf(y, other.y))
                    for (c in listOf(z, other.z)) {
                        ps.add(Point(a, b, c))
                    }
            return ps
        }

        fun comb(other: Point): List<Pair<Point, Point>> {
            val ps = corners(other)
            val res = mutableListOf<Pair<Point, Point>>()
            for ((i, x) in ps.withIndex())
                for (j in i + 1 until ps.size) {
                    res.add(x to ps[j])
                }
            return res
        }
    }

    data class Star(val min: Point, val max: Point) {
        fun maxDim(other: Point): Long {
            return min.corners(max).maxOf {
                maxOf(abs(it.x - other.x), abs(it.y - other.y), abs(it.z - other.z))
            }
        }
    }


    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()

        val stars = (0 until N).map {
            val (x, y, z, r) = input_longs()
            Star(Point(x - r, y - r, z - r), Point(x + r, y + r, z + r))
        }
        var minP = stars.map { it.min }.reduce { acc, l -> acc.minWith(l) }
        var maxP = stars.map { it.max }.reduce { acc, r -> acc.maxWith(r) }

        var ML = maxOf(maxP.x - minP.x, maxP.y - minP.y, maxP.z - minP.z)
        var min = ML
        for ((pa, pb) in minP.comb(maxP)) { // bug: need consider all comb, not just (minP, maxP) case
            var (l, r) = 1L to ML
            while (l <= r) {
                var m = (l + r) / 2
                if (stars.all { s -> (s.maxDim(pa) <= m || s.maxDim(pb) <= m) }) {
                    r = m - 1
                } else {
                    l = m + 1
                }
            }
            min = minOf(min, l)
        }

        println("Case #${it}: $min")
    }
}
