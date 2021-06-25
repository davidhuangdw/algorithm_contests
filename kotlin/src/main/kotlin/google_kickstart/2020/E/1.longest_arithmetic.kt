package google_kickstart.`2020`.E

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    fun input_longs() = split_input().map { it.toLong() }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (n) = input_ints()
        val a = input_ints()
        var pre = 2
        var res = 2
        for (i in 2 until n) {
            if (a[i] - a[i - 1] == a[i - 1] - a[i - 2]) {
                pre += 1
                res = maxOf(res, pre)
            } else
                pre = 2
        }
        println("Case #${it}: $res")
    }
}
