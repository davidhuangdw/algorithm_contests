package google_code_jam.`2022`.qua

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val SUM = 1000_000
    val T = readLine()!!.toInt()
    (1..T).forEach {
        var mins = MutableList(4) { SUM }
        (1..3).forEach {
            input_ints().mapIndexed { i, v -> mins[i] = minOf(mins[i], v) }
        }
        var remain = SUM
        mins.mapIndexed { i, v ->
            val cost = minOf(v, remain)
            mins[i] = cost
            remain -= cost
        }
        val res = if (remain == 0) mins.joinToString(" ") else "IMPOSSIBLE"
        println("Case #${it}: $res")
    }
}
