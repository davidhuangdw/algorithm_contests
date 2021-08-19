package google_kickstart.`2018`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val line = input_ints()
        val (N, V, H, A, B) = line
        val (C, D, E, F, M) = line.subList(5, line.size)
        val points = mutableListOf(V to H)
        var (pv, ph) = V to H
        repeat(N - 1) {
            val v = (A * pv + B * ph + C) % M
            val h = (D * pv + E * ph + F) % M
            points.add(v to h)
            pv = v
            ph = h
        }
        points.sortWith(compareBy { it.first })

        fun testset1(): Long {
            val count = List(M) { MutableList(M) { 0L } }
            for ((v, h) in points)
                count[v][h] += 1L
            for (i in 0 until M)
                for (j in 0 until M)
                    count[i][j] += (if (i > 0) count[i - 1][j] else 0) + (if (j > 0) count[i][j - 1] else 0) - (if (i > 0 && j > 0) count[i - 1][j - 1] else 0)

            fun sub_count(i1: Int, i2: Int, j1: Int, j2: Int): Long { // (i1, i2] & (j1, j2]
                return if (i1 < i2 && j1 < j2) {
                    count[i2][j2] - count[i2][j1] - count[i1][j2] + count[i1][j1]
                } else 0L
            }

            var sum = 0L
            for ((i, x) in points.withIndex()) {
                val (vx, hx) = x
                for (j in i + 1 until N) {
                    val (vy, hy) = points[j]
                    sum += (j - i - 1) - sub_count(minOf(vx, vy), maxOf(vx, vy) - 1, minOf(hx, hy), maxOf(hx, hy) - 1)
                }
            }
            return sum
        }

        fun testset2(): Long {   // todo: fix WA for large case
            val hs = hashMapOf<Int, Int>()
            for ((i, h) in points.map { it.second }.toHashSet().sorted().withIndex()) {
                hs[h] = i
            }
            val H = hs.size

            val left = mutableListOf<Pair<Long, Long>>()
            var tree = MutableList(H + 1) { 0 }
            var sz = 0L
            fun query(xx: Int): Long {
                var sum = 0L
                var x = xx + 1
                while (x > 0) {
                    sum += tree[x]
                    x -= x and -x
                }
                return sum
            }

            fun add(xx: Int) {
                var x = xx + 1
                while (x <= H) {
                    tree[x] += 1
                    x += x and -x
                }
                sz += 1
            }

            var i = 0
            while (i < N) {
                var j = i
                while (j < N && points[j].first == points[i].first) {
                    val h = points[j].second
                    left.add(query(hs[h]!! - 1) to sz - query(hs[h]!!))
                    j += 1
                }
                for (k in i until j)
                    add(hs[points[k].second]!!)
                i = j
            }

            var sum = 1L * N * (N - 1) * (N - 2) / 6
            sz = 0
            tree = MutableList(H + 1) { 0 }
            i = N - 1
            while (i >= 0) {
                var j = i
                while (j >= 0 && points[j].first == points[i].first) {
                    val h = points[j].second
                    val (ra, rb) = query(hs[h]!! - 1) to sz - query(hs[h]!!)
                    val (la, lb) = left[j]
                    sum -= ra * lb + rb * la
                    j -= 1
                }
                for (k in j + 1..i)
                    add(hs[points[k].second]!!)
                i = j
            }
            return sum
        }

        println("Case #${it}: ${testset2()}")
    }
}
