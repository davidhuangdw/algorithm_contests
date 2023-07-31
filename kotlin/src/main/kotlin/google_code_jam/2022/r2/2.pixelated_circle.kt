package google_code_jam.`2022`.r2

import kotlin.math.sqrt

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (R) = input_ints()
        fun round(q: Double) = if (q <= q.toInt() + 0.5) q.toInt() else q.toInt() + 1
        fun round_sqrt(x: Int) = round(sqrt(x * 1.0))
        fun set1(): Int {
            val r = minOf(R, 200)
            val M = 2 * r + 1
            val a = List(M) { MutableList(M) { 0 } }
            val b = List(M) { MutableList(M) { 0 } }
            for (x in -r..r)
                for (y in -r..r)
                    if (round_sqrt(x * x + y * y) <= r)
                        a[x + r][y + r] = 1
            for (k in 0..r) {
                for (x in -k..k) {
                    val y = round_sqrt(k * k - x * x)
                    for ((i, j) in listOf(x to y, x to -y, y to x, -y to x))
                        b[i + r][j + r] = 1
                }
            }
//            fun show(p: List<MutableList<Int>>){
//                for((i,row) in p.withIndex())
//                    println(row.joinToString(" "))
//            }
//            show(a)
//            show(b)
            var cnt = 0
            for (i in 0 until M)
                for (j in 0 until M)
                    if (a[i][j] != b[i][j]) {
                        cnt++
                    }
            return cnt
        }

        val r = set1()
        println("Case #${it}: $r")
    }
}
