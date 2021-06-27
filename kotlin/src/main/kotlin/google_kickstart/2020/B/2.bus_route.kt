package google_kickstart.`2020`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, D) = input_longs()
        var cur = D
        for (v in input_longs().reversed())
            cur = cur / v * v
        println("Case #${it}: $cur")
    }
}
