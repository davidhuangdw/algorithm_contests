package google_kickstart.`2019`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        val P = input_ints()
        val A = input_ints()
        val color_pos = hashMapOf<Int, MutableList<Int>>()

        for ((c, p) in A.zip(P)) {
            if (c !in color_pos)
                color_pos[c] = mutableListOf()
            color_pos[c]!!.add(p)
        }

        // idea: extra info: 0 - all return, 1 - has one no return
        var pre = List(2) { MutableList(K + 1) { if (it == 0) 0 else 1e9.toInt() } }
        for ((c, row) in color_pos) {
            row.sort()
            for (k in K downTo 1) {
                for ((i, p) in row.withIndex()) {
                    if (i + 1 > k) break
                    pre[0][k] = minOf(pre[0][k], pre[0][k - (i + 1)] + p * 2)
                    pre[1][k] = minOf(pre[1][k], pre[0][k - (i + 1)] + p, pre[1][k - (i + 1)] + p * 2)
                }
            }
        }
        println("Case #${it}: ${pre[1][K]}")
    }
}
