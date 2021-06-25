package google_kickstart.`2020`.D

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    fun input_longs() = split_input().map { it.toLong() }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (k) = input_ints()
        val a = input_ints().toMutableList()
        var n = 1
        for (i in 1 until k)
            if (a[i] != a[i - 1]) {
                a[n] = a[i]
                n += 1
            }

        var pre = 0
        var cnt = 0
        for (i in 1 until n) {
            if (i == n - 1 || (a[i - 1] < a[i] && a[i] > a[i + 1]) || (a[i - 1] > a[i] && a[i] < a[i + 1])) {
                cnt += (i - pre) / 4
                pre = i
            }
        }

        println("Case #${it}: $cnt")
    }
}
