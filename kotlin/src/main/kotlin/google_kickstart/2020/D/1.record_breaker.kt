package google_kickstart.`2020`.D

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    fun input_longs() = split_input().map { it.toLong() }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (n) = input_ints()
        val a = input_ints()
        var max = -1
        var cnt = 0
        for (i in 0 until n) {
            if (a[i] > max && (i == n - 1 || a[i] > a[i + 1]))
                cnt += 1
            max = maxOf(max, a[i])
        }
        println("Case #${it}: $cnt")
    }
}
