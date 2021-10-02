fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        var (R, C, K) = input_ints()
        K -= 1
        val G = (0 until R).map { readLine()!!.map { if (it == 'X') 1 else 0 } }
        val cols = (0 until C).map { c -> (0 until R).map { r -> G[r][c] }.sum() }
        val up = 1 + cols.map { if (it <= K) 0 else 1 }.sum()
        val down = 1 + cols.map { if (it <= R - K - 1) 0 else 1 }.sum()
        println(cols)
        println(up)
        println(down)

        val min_cnt = minOf(
            G[K].sum(),
            1 + cols.map { if (it <= K) 0 else 1 }.sum(),
            1 + cols.map { if (it <= R - K - 1) 0 else 1 }.sum(),
        )
        // todo: fix it


        println("Case #${it}: $min_cnt")
    }
}
