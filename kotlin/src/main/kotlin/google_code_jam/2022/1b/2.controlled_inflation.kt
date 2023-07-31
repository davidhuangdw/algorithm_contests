package google_code_jam.`2022`.`1b`

import kotlin.math.absoluteValue

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, P) = input_ints()
        var dp = listOf(0L to 0L, 0L to 0L)
        for(_i in 0 until N){
            val X = input_longs()
            val (l, r) = X.minOrNull()!! to X.maxOrNull()!!
            val ed = listOf(l, r)
            dp = (0..1).map{i ->
                val (a, b) = ed[i] to ed [1-i]
                val sum = dp.minOf{p ->
                    val (sum, st) = p
                    sum + (st - a).absoluteValue + (a-b).absoluteValue
                }
                sum to b
            }
        }
        val res = dp.minOf{it.first}

        println("Case #${it}: $res")
    }
}
