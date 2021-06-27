package google_kickstart.`2020`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val H = input_ints()
        val cnt = (1..N - 2).count { i -> H[i - 1] < H[i] && H[i] > H[i + 1] }
        println("Case #${it}: $cnt")
    }
}
