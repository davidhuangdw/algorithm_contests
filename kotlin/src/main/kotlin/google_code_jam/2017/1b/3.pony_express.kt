package google_code_jam.`2017`.`1b`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (N, _Q) = input_ints()
        val E = MutableList(N) { 0L }
        val S = MutableList(N) { 0L }
        for (i in 0 until N) {
            val (e, s) = input_longs()
            E[i] = e
            S[i] = s
        }
        val D = (1..N).map { input_longs() }
        val Q = (1.._Q).map { input_ints() }

        fun set2(): String {
            // todo: fix s2 WA
            fun min(a: Long, b: Long) = if (b < 0 || (a in 1..b)) a else b
            val dis = List(N) { D[it].toMutableList() }
            for (i in 0 until N)
                dis[i][i] = 0L
            for (k in 0 until N)
                for (i in 0 until N)
                    for (j in 0 until N) {
                        if (i != j && dis[i][k] > 0 && dis[k][j] > 0) {
                            val d = dis[i][k] + dis[k][j]
                            if (d <= E[i])
                                dis[i][j] = min(dis[i][j], d)
                        }
                    }

            val MAX = 1e18
            val T = List(N) { i ->
                dis[i].map { d ->
                    if (d > 0) {
                        d * 1.0 / S[i]
                    } else MAX
                }
            }
            val t = List(N) { T[it].toMutableList() }
            for (k in 0 until N) {
                for (i in 0 until N)
                    for (j in 0 until N)
                        if (i != j && t[i][k] != MAX && t[k][j] != MAX) {
                            t[i][j] = minOf(t[i][j], t[i][k] + t[k][j])
                        }
            }

            val res = Q.map { (a, b) -> t[a - 1][b - 1] }
            return res.joinToString(" ")
        }

        val r = set2()
        println("Case #${it}: $r")
    }
}
