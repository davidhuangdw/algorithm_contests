package google_code_jam.`2022`.qua

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val MAX = 2000_000_000L
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val F = input_longs()
        val P = input_ints()
        val ch_min = MutableList(N) { MAX }
        var sum = 0L
        for (i in N - 1 downTo 0) {
            if (ch_min[i] == MAX)
                ch_min[i] = 0
            val cur = maxOf(ch_min[i], F[i])
            sum += maxOf(0, cur - ch_min[i])
            val j = P[i] - 1
            if (j >= 0)
                ch_min[j] = minOf(ch_min[j], cur)
        }
        println("Case #${it}: $sum")
    }
}
