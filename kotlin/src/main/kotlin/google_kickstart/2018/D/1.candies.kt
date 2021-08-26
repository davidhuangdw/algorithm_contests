package google_kickstart.`2018`.D

import java.util.*

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (NN, OO, D) = input_longs()
        val (N, O) = NN.toInt() to OO.toInt()
        val line = input_ints()
        var (x, y) = line.subList(0, 2).map { it.toLong() }
        val (A, B, C, M, L) = line.subList(2, line.size)

        val S = MutableList(N) { i ->
            when (i) {
                0 -> x
                1 -> y
                else -> {
                    x = y.also { y = (A * y + B * x + C) % M }
                    y
                }
            } + L
        }
//        println(S)

        var max = Long.MIN_VALUE
        var count = sortedMapOf<Long, Int>() as TreeMap

        val pre = S.scan(0L, Long::plus)
        var r = 1
        var odd_cnt = 0
        for (i in 1..N) {
            while (r <= N && odd_cnt <= O) {
                val is_odd = S[r - 1] % 2 != 0L
                if (is_odd && odd_cnt == O) break            // bug: allow add more after odd_cnt <= O, as long as odd_cnt <= O
                if (is_odd)
                    odd_cnt += 1
                count[pre[r]] = count.getOrDefault(pre[r], 0) + 1
                r += 1
            }


            val up = count.floorKey(D + pre[i - 1])
            if (up != null) {
//                println("$i ${r-1} ${up - pre[i-1]} $odd_cnt")
                max = maxOf(max, up - pre[i - 1])
            }

            // remove i
            if (r <= i) r += 1      // bug: maybe not included
            else {
                if (S[i - 1] % 2 != 0L)
                    odd_cnt -= 1
                if (count[pre[i]] == 1)
                    count.remove(pre[i])
                else
                    count[pre[i]] = count[pre[i]]!! - 1
            }
        }


        println("Case #${it}: ${if (max == Long.MIN_VALUE) "IMPOSSIBLE" else max}")
    }
}
