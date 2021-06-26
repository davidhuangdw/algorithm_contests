package google_kickstart.`2020`.C

import kotlin.math.sqrt

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val sq = (0..sqrt(1e8).toInt()).map { it * it }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val P = input_ints().scan(0, Int::plus)
        val cnt = hashMapOf<Int, Int>()
        cnt[0] = 1

        var (res, min) = 0L to 0
        for (v in P.subList(1, P.size)) {
            for (d in sq) {
                if (v - d < min) break
                res += cnt.getOrDefault(v - d, 0)
            }

            cnt[v] = cnt.getOrDefault(v, 0) + 1
            min = minOf(min, v)
        }
        println("Case #${it}: $res")
    }
}
