package atcoder.a2

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }

    val coins = listOf(1, 5, 10, 50, 100, 500).reversed()
    val cnt = input_ints().reversed().toMutableList()
    val (N) = input_ints()
    val X = input_ints()
    println(if (X.all { x ->
            var r = x
            for((i, c) in coins.withIndex()) {
                val k = minOf(r / c, cnt[i])
                r -= k * c
                cnt[i] -= k
            }
            r == 0
        }) "Yes" else "No")
}