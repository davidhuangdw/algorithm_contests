package google_kickstart.`2019`.Practice

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val S = readLine()!!
        val pre = S.map { it - '0' }.scan(0, Int::plus)

        val k = (N + 1) / 2
        var max = 0
        for ((i, v) in pre.withIndex()) {
            if (i - k >= 0)
                max = maxOf(max, v - pre[i - k])
        }
        println("Case #${it}: $max")
    }
}
