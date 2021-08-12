package google_kickstart.`2019`.F

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (V) = input_ints()
        val adj = List(V) { mutableListOf<Int>() }
        val B = input_longs()
        repeat(V - 1) {
            val (u, v) = input_ints()
            adj[u - 1].add(v - 1)
            adj[v - 1].add(u - 1)
        }
        val par = MutableList(V) { -1 }
        fun getPar(u: Int) {
            for (v in adj[u])
                if (v != par[u]) {
                    par[v] = u
                    getPar(v)
                }
        }
        getPar(0)
        for (i in 0 until V)
            adj[i].remove(par[i])


        val MIN = Long.MIN_VALUE
        val dp = List(V) { List(2) { MutableList(2) { MIN } } }
        fun calc(u: Int, par_on: Int, on: Int): Long {
            if (dp[u][par_on][on] == MIN) {
                dp[u][par_on][on] = if (on == 1 || par_on == 1) {
                    var sum = B[u]
                    for (v in adj[u])
                        sum += maxOf(calc(v, on, 0), calc(v, on, 1))
                    sum
                } else {
                    //dark
                    var mx = 0L
                    for (v in adj[u])
                        mx += calc(v, on, 0)

                    //light
                    val n = adj[u].size
                    if (n > 0) {
                        var light = B[u]
                        val pre = List(n) { MutableList(2) { 0L } }
                        for ((i, v) in adj[u].withIndex()) {
                            if (i == 0) {
                                pre[i][0] = calc(v, on, 0)
                                pre[i][1] = calc(v, on, 1)
                            } else {
                                pre[i][0] = pre[i - 1][0] + calc(v, on, 0)
                                pre[i][1] = maxOf(
                                    pre[i - 1][0] + calc(v, on, 1),
                                    pre[i - 1][1] + calc(v, on, 0),
                                    pre[i - 1][1] + calc(v, on, 1),
                                )
                            }
                        }
                        light += pre[n - 1][1]

                        mx = maxOf(mx, light)
                    }
                    mx
                }
            }
            return dp[u][par_on][on]
        }

        println("Case #${it}: ${maxOf(calc(0, 0, 0), calc(0, 0, 1))}")
    }
}
