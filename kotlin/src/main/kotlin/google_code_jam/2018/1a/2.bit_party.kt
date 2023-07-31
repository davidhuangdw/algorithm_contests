package google_code_jam.`2018`.`1a`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (R, B, C) = input_ints()
        val Ca = (1..C).map { input_longs() } //.sortedBy{it[1]}
        fun set1(): Long {
            var min = Long.MAX_VALUE
            fun dfs(i: Int, r: Int, b: Int, max: Long) {
                if (b == 0) {
                    min = minOf(min, max)
                }
                if (i >= C || r == 0)
                    return
                val (m, s, p) = Ca[i]
                val u = minOf(m.toInt(), b)
                for (n in if (r == 1) u..u else 1..u) {
                    val t = n * s + p // bug: should use Long
                    dfs(i + 1, r - 1, b - n, maxOf(max, t))
                }
                dfs(i + 1, r, b, max)
            }
            dfs(0, R, B, 0L)
            return min
        }

        fun set2(): Long {
            var (l, r) = 1L to Long.MAX_VALUE
            fun check(t: Long): Boolean {
                val pick = Ca.map { (m, s, p) ->
                    if (p > t)
                        0L  // bug: should be 0L not -1L, why? because allow to pick less than R!!
                    else {
                        minOf(m, (t - p) / s)
                    }
                }.sortedDescending()
                return pick.subList(0, R).sum() >= B
            }
            while (l <= r) {
                val m = l + (r - l) / 2
                if (check(m)) r = m - 1
                else l = m + 1
            }
            return l
        }

        val res = set2()
        println("Case #${it}: $res")
    }
}
