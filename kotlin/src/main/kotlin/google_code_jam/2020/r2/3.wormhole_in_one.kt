package google_code_jam.`2020`.r2

import kotlin.math.absoluteValue

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun gcd(aa: Long, bb: Long): Long {
        var a = aa.absoluteValue
        var b = bb.absoluteValue
        while (a != 0L) {
            a = (b % a).also { b = a }
        }
        return b
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val P = (1..N).map { input_longs() }
        val slopes = hashSetOf<Pair<Long, Long>>()
        for (i in 0 until N) {
            for (j in i + 1 until N) {
                var (a, b) = P[i][0] - P[j][0] to P[i][1] - P[j][1]
                if (a < 0) {
                    a = -a
                    b = -b
                }
                val g = gcd(a, b)
                a /= g
                b /= g
                slopes.add(a to b)
            }
        }

        var max = minOf(N, 3)
        for ((a, b) in slopes) {
            val done = MutableList(N) { false }
            var sum = 0
            for ((i, p) in P.withIndex()) {
                if (done[i])
                    continue
                var cnt = 1
                done[i] = true
                for (j in i + 1 until N)
                    if (!done[j] && (P[j][0] - p[0]) * b == (P[j][1] - p[1]) * a) {
                        done[j] = true
                        cnt++
                    }
                if (cnt > 1)
                    sum += cnt
            }
            sum -= sum % 2
            max = maxOf(max, minOf(N, sum + 2))
        }
        println("Case #${it}: $max")
    }
}
