package google_code_jam.`2018`.qua

import kotlin.math.*

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val A = split_input()[0].toDouble()
        println("Case #${it}:")
        fun set1() {
            var x = PI / 4 - acos(A / sqrt(2.0))
//            println("$x ${PI/4}")
            val a = 0.5 * sin(x)
            val b = 0.5 * cos(x)
            println("0 0 0.5")
            println("$a -$b 0") // bug: -b
            println("$b $a 0")
        }
        // todo: set2
        set1()
    }
}
