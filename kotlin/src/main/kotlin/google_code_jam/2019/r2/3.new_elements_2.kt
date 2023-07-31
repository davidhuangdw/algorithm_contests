package google_code_jam.`2019`.r2

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun gcd(_x: Long, _y: Long): Long {
        var (x, y) = _x to _y
        while (x != 0L) {
            x = y % x.also { y = x }
        }
        return y
    }

    fun reduce(p: Pair<Long, Long>): Pair<Long, Long> {
        val (x, y) = p
        if (x == 0L) return 0L to 1L
        if (y == 0L) return 1L to 0L
        val g = gcd(x, y)
        return x / g to y / g
    }

    fun cmp(a: Pair<Long, Long>, b: Pair<Long, Long>): Int {
        val (ax, ay) = a
        val (bx, by) = b
        if(ax * by == ay * bx) return 0
        var d = if(ax * by > ay * bx) 1 else -1
        if (ay * by < 0L) d = -d
        return d
    }
    val (T) = input_ints()
    (1..T).forEach {
        val (N) = input_ints()
        val P = (1..N).map {
            val (x, y) = input_longs()
            x to y
        }

        fun set(): String {
            var l = 1L to 1000_000_001L
            var r = 1000_000_001L to 1L

            for (i in 0 until N)
                for (j in i + 1 until N) {
                    val (ax, ay) = P[i]
                    val (bx, by) = P[j]
                    var (dx, dy) = bx - ax to by - ay
                    if (dx * dy >= 0) {
                        if (!(dx >= 0 && dy >= 0))
                            return "IMPOSSIBLE"
                    } else {
                        if (dx > 0) {
                            val v = reduce(-dy to dx)
                            if (cmp(l, v) < 0)
                                l = v
                        } else {
                            val v = reduce(dy to -dx)
                            if (cmp(v, r) < 0)
                                r = v
                        }
                    }
                }
            if (cmp(l, r) >= 0)
                return "IMPOSSIBLE"

            fun set1(): String {
                for (x in 1..202L)
                    for (y in 1..202L) {
                        val v = x to y
                        if (cmp(l, v) < 0 && cmp(v, r) < 0)
                            return "$x $y"
                    }
                return ""
            }
            fun set2(): String{
                // todo: fix set2: because bigInteger
                l = (r.second to r.first).also { r = l.second to l.first }

                fun aboveL(x: Long): Long {
                    var (ly, ry) = 0L to 2000_000_002L
                    while (ly <= ry) {
                        val my = (ly + ry) / 2
                        if (cmp(my to x, l) <= 0) ly = my + 1
                        else ry = my - 1
                    }
                    return ly
                }
                var (lx, rx) = 1L to 2000_000_002L
                while (lx <= rx) {
                    val mx = (lx + rx) / 2
                    val my = aboveL(mx)
                    if (cmp(my to mx, r) < 0) rx = mx - 1
                    else lx = mx + 1
                }
                val (x, y) = lx to aboveL(lx)
                return "$x $y"
            }
            return set2()
        }
        val res = set()
        println("Case #${it}: $res")
    }
}
