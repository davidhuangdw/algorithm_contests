package google_code_jam.`2020`.`1c`

import kotlin.math.absoluteValue

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val S = split_input()
        fun sol(): String{
            var (x, y) = S[0].toInt() to S[1].toInt()
            if(x == 0 && y == 0) return "0"
            for((i, ch) in S[2].withIndex()){
                when(ch){
                    'S' -> y--
                    'N' -> y++
                    'W' -> x--
                    'E' -> x++
                }
                if(x.absoluteValue + y.absoluteValue <= i+1)
                    return (i+1).toString()
            }
            return "IMPOSSIBLE"
        }
        val res = sol()
        println("Case #${it}: $res")
    }
}
