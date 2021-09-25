fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val pre = readLine()!!.map { it - '0' }.scan(0, Int::plus)
        val l = (N + 1) / 2
        var max = 0
        for (i in l..N)
            max = maxOf(max, pre[i] - pre[i - l])
        println("Case #${it}: $max")
    }
}
