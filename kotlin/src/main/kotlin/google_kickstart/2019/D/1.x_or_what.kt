package google_kickstart.`2019`.D

import java.util.*

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val odd = List(1024) {
        var x = it
        var k = 0
        while (x > 0) {
            x -= x and -x
            k = k xor 1
        }
        k
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, Q) = input_ints()
        val A = input_ints().toMutableList()
        val queries = (1..Q).map { input_ints() }

        fun testset1(): MutableList<Int> {
            val res = mutableListOf<Int>()
            for ((i, v) in queries) {
                A[i] = v
                val pre = A.runningFold(0, Int::xor)
                var longest = 0
                for (d in 0..1) {
                    val l = (0..N).firstOrNull { odd[pre[it]] == d }
                    val r = (N downTo 0).firstOrNull { odd[pre[it]] == d }
                    if (r != null && l != null)
                        longest = maxOf(longest, r - l)
                }
                res.add(longest)
            }
            return res
        }

        fun testset2(): MutableList<Int> {
            val res = mutableListOf<Int>()
            val odd_pos = (1..N).filter { odd[A[it - 1]] % 2 == 1 }.toSortedSet()
            for ((i, v) in queries) {
                if (odd[A[i] xor v] % 2 == 1) {
                    if (odd[A[i]] % 2 == 1) odd_pos.remove(i + 1)
                    else odd_pos.add(i + 1)
                }
                res.add(
                    if (odd_pos.size % 2 == 0) N
                    else maxOf(odd_pos.last() - 1, N - odd_pos.first())
                )
                A[i] = v
            }
            return res
        }

        val res = testset2().joinToString(" ")
        println("Case #${it}: $res")
    }
}
